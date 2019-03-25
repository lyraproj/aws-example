<p align="center"><img src="docs/media/lyrabanner.png" alt="Lyra"></p>

## Example AWS plugin

[Lyra](https://github.com/lyraproj/lyra) is an open source workflow engine for provisioning and managing cloud native infrastructure. Using infrastructure as code, Lyra enables you to declaratively provision and manage public cloud, private cloud, and other API-backed resources as well as orchestrate imperative actions.

This repo contains an example plugin that provides AWS content. It implements a few resources only and is intended to be a learning resource for plugin creators.

This plugin is not intended for solving real-world problems! The content included with Lyra uses the Lyra Terraform Bridge makes [Terraform providers](https://github.com/terraform-providers) and is much more comprehensive.

### Build
The project requires [Go](https://golang.org/doc/install) 1.11 or higher, and [go modules](https://github.com/golang/go/wiki/Modules) to be on.

Build the project using make:

	make

When no targets are specified, the build will lint, test, compile and sanity-check..

## Contributing
We'd love to get contributions from you! For a quick guide, take a look at our guide to [contributing](CONTRIBUTING.md).
