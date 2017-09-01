# Vault-Get

A small application connects to [Vault](https://www.vaultproject.io/) and out puts the value of a key in a path using user/password authentication.

## Getting Started

Application is packed as a single binary, just download and run.

### Prerequisites

Nothing.

### Usage

The help section explains everything:

```
NAME:
   vault-get - Get a value from Vault

USAGE:
   vault-get [global options] command [command options] [arguments...]

VERSION:
   0.5.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --vault_host value      Vault host url [$VAULT_HOST]
   --vault_username value  Vault username [$VAULT_USERNAME]
   --vault_password value  Vault password [$VAULT_PASSWORD]
   --vault_path value      Vault path to get key from [$VAULT_PATH]
   --vault_key value       Vault key to get value from [$VAULT_KEY]
   --help, -h              show help
   --version, -v           print the version
```

## Deployment

Download a [release compatible to your OS](https://github.com/devops-israel/vault-get/releases) and run the application.

## Built With

* [Golang](https://golang.org/)
* [Cli](https://github.com/urfave/cli) - A simple, fast, and fun package for building command line apps in Go.
* [Vault API](github.com/hashicorp/vault/api)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md)

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/devops-israel/delete-aws-es-incidents/tags).

## Authors

* [**Josh Dvir**](https://github.com/joshdvir)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
