# Goptos

A fine grained reactive web framework in Go.

*Inspired by Leptos and SolidJS*

# Usage

To learn Goptos please refer to the [docs](https://github.com/goptos/docs). This README is primarily concerned with contributor guidance.

> N.B you won't need to follow any steps details here if only want to use Goptos, instead follow the getting started guide.

# Contributing

## Setup

The provided Makefile does most of the heavy lifting here. To setup your dev environment we use [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers) and have provided the dev container configuration within this repo.

> Configuring Dev Containers is outside of the scope of this document.

Begin by cloning this repo and optionally use VS Code to 'Reopen in Dev Container'.

Running `Make pre` will install prerequisites including Go and the Goptos CLI. This will be run for you if using a Dev Container.

Running `Make goptos` will clone the additional Goptos repositories for you. This is also already run when using a Dev Container.

You will need add the additional Goptos directory to your dev workspace. If using Dev Containers this will reload the session.

You should end up with a file structure as follows:

```
app
  > .devcontainer
  > .vscode
  > dist
  > src
  .gitignore
  index.html
  LICENSE
  Makefile
  README.md
goptos
  > ast
  > cli
  > lexer
  > parser
  > runtime
  > system
  > utils
```

> If cloning from the main branch each go.mod file will have a `replace` pointing locally to the cloned Goptos repositories. Ensure these paths are correct for your dev environment.

### Verbose

Goptos utils provides a `Verbose` struct to easily print debug statements. This will read the environment variable `GOPTOS_VERBOSE` to determine how much to print. Set this to your desired amount, eg:

`export GOPTOS_VERBOSE=4`

> Dom console debug statements should be printed using level 5 or above.

## Submitting code

This will likely be done through pull request, however, more information to come here...




