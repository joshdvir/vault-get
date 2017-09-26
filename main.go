package main

import (
	"errors"
	"fmt"
	"os"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vault-get"
	app.Usage = "Get a value fron Vault"
	app.Version = fmt.Sprintf("0.5.0")
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "vault_host",
			Usage:  "Vault host url",
			EnvVar: "VAULT_HOST",
		},
		cli.StringFlag{
			Name:   "vault_username",
			Usage:  "Vault username",
			EnvVar: "VAULT_USERNAME",
		},
		cli.StringFlag{
			Name:   "vault_password",
			Usage:  "Vault password",
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

		if len(cli.String("vault_username")) == 0 {
			return errors.New("No Vault username provided")
		}

		if len(cli.String("vault_password")) == 0 {
			return errors.New("No Vault password provided")
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
		}

		path := fmt.Sprintf("auth/userpass/login/%s", cli.String("vault_username"))
		logical := client.Logical()
		secret, err := logical.Write(path, options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting token from vault: %s", err)
			os.Exit(1)
		}

		client, _ = vaultapi.NewClient(&vaultapi.Config{Address: cli.String("vault_host")})
		client.SetToken(secret.Auth.ClientToken)
		logical = client.Logical()

		vaultSecret, err := logical.Read(cli.String("vault_path"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from vault: %s", err)
			os.Exit(1)
		}

		output := "export "
		for vkey, vvalue := range vaultSecret.Data {
			output += vkey + "=" + vvalue.(string) + " "
		}
		fmt.Printf(output)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
