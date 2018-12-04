package main

import (
	"errors"
	"fmt"
	"os"

	vaultapi "github.com/hashicorp/vault/api"
	flatten "github.com/jeremywohl/flatten"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vault-get"
	app.Usage = "Get a value from Vault"
	app.Version = fmt.Sprintf("0.6.0")
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "vault_host",
			Usage:  "Vault host url",
			EnvVar: "VAULT_HOST",
		},
		cli.StringFlag{
			Name:   "vault_auth",
			Usage:  "Vault auth: defaults to 'token' (can be set explicitly with vault_token) or 'userpass' with vault_username + vault_password",
			EnvVar: "VAULT_AUTH",
			Value:  "token",
		},
		cli.StringFlag{
			Name:   "vault_token",
			Usage:  "Vault token (used if vault_auth is token)",
			EnvVar: "VAULT_TOKEN",
		},
		cli.StringFlag{
			Name:   "vault_username",
			Usage:  "Vault username (used if vault_auth is userpass)",
			EnvVar: "VAULT_USERNAME",
		},
		cli.StringFlag{
			Name:   "vault_password",
			Usage:  "Vault password (used if vault_auth is userpass)",
			EnvVar: "VAULT_PASSWORD",
		},
		cli.StringFlag{
			Name:   "vault_path",
			Usage:  "Vault path of the secret. eg. secret/my-secret",
			EnvVar: "VAULT_PATH",
		},
	}

	app.Action = func(cli *cli.Context) error {
		if len(cli.String("vault_host")) == 0 {
			return errors.New("No Vault host provided")
		}

		if cli.String("vault_auth") == "userpass" {
		    if len(cli.String("vault_username")) == 0 {
			return errors.New("No Vault username provided")
		    }

		    if len(cli.String("vault_password")) == 0 {
			return errors.New("No Vault password provided")
		    }

		} else if cli.String("vault_auth") == "token" {
		    if len(cli.String("vault_token")) == 0 {
			return errors.New("No token provided")
		    }
		}

		if len(cli.String("vault_path")) == 0 {
			return errors.New("No Vault path provided")
		}

		client, err := vaultapi.NewClient(&vaultapi.Config{Address: cli.String("vault_host")})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from vault: %s", err)
			os.Exit(1)
		}

		options := map[string]interface{}{
			"password": cli.String("vault_password"),
			"max_ttl": "7200",
			"ttl": "7200",
		}

		logical := client.Logical()
		if cli.String("vault_auth") == "userpass" {
		    path := fmt.Sprintf("auth/userpass/login/%s", cli.String("vault_username"))
		    secret, err := logical.Write(path, options)
		    if err != nil {
			fmt.Fprintf(os.Stderr, "error getting token from vault: %s", err)
			os.Exit(1)
		    }
		    client, _ = vaultapi.NewClient(&vaultapi.Config{Address: cli.String("vault_host")})
		    client.SetToken(secret.Auth.ClientToken)

		} else if cli.String("vault_auth") == "token" {
		    client, _ = vaultapi.NewClient(&vaultapi.Config{Address: cli.String("vault_host")})
		    client.SetToken(cli.String("vault_token"))
		}

		logical = client.Logical()

		vaultSecret, err := logical.Read(cli.String("vault_path"))
		if vaultSecret == nil {
			fmt.Fprintf(os.Stderr, "Error retrieving data: path is wrong or not complete")
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from vault: %s", err)
			os.Exit(1)
		}

		data := flatten.Flatten(vaultSecret.Data, "", flatten.RailsStyle)
		output := "export "
		for vkey, vvalue := range data {
			output += vkey + "=" + fmt.Sprint(vvalue) + " "
		}
		fmt.Printf(output)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
