# Terraform Provider for Meta Networks

See the [Meta Networks Provider documentation](docs/index.md) to get started using the provider.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or higher
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)
- [GOPATH](http://golang.org/doc/code.html#GOPATH)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/mataneine/terraform-provider-metanetworks`

```shell
$ mkdir -p $GOPATH/src/github.com/mataneine; cd $GOPATH/src/github.com/mataneine
$ git clone git@github.com:mataneine/terraform-provider-metanetworks
```

Enter the provider directory and build the provider

```shell
$ cd $GOPATH/src/github.com/mataneine/terraform-provider-metanetworks
$ make build
```

## Using the provider

```terraform
terraform {
  required_providers {
    metanetworks = {
      source  = "FabioAntunes/metanetworks"
      version = "1.0.0-pre-2.4"
    }
  }
}

provider "metanetworks" {
  org = "example_organization"
}

# Example resource configuration
resource "metanetworks_resource" "example" {
  # ...
}
```

## Developing the Provider

Enter the provider directory.

```sh
$ cd $GOPATH/src/github.com/mataneine/terraform-provider-metanetworks
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/src/github.com/mataneine/terraform-provider-metanetworks/bin` directory.

```shell
$ make build
```

To use it locally you can install [this provider as plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin)

```shell
$ cd $GOPATH/src/github.com/FabioAntunes/terraform-provider-metanetworks
$ make install
```

And then use it on your terraform:

```terraform
provider "metanetworks" {
  org = "example_organization"
}

terraform {
  required_providers {
    metanetworks = {
      source  = "localhost/FabioAntunes/metanetworks"
      version = "1.0.0-pre-2.4"
    }
  }
}

# Example resource configuration
resource "metanetworks_resource" "example" {
  # ...
}
```

In order to test the provider, you can simply run `make test`.

```shell
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

**Note:** Acceptance tests create real resources, and often cost money to run.

```shell
$ make testacc
```
