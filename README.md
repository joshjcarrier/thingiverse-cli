# Thingiverse command line tool

Manages your Things from Thingiverse.com! Find things and back them up to your computer.

For more info run `thing help`.

## Install

Download the pre-compiled binary from the [Releases page](https://github.com/joshjcarrier/thingiverse-cli/releases/latest). Move the binary to somewhere on your \$PATH like `/usr/local/bin` to run it as `thing`, otherwise run it as `./thing`.

To configure providing authentication via the `--token` flag, create a `.thingrc` in your `$HOME` directory or in the same directory as the binary.

**.thingrc**

```yaml
token: <your-bearer-token>
```

Get a token by creating a Desktop app at https://www.thingiverse.com/apps/create and using the readonly app token.

## Recipes

### Get Things by search result, in Markdown format

```sh
thing search plant | xargs thing get -f md
```

### Get Things by username, in Markdown format

```sh
thing user joshc -f md
```

### Get a Thing by its id

```sh
thing get 3786462
```

## Dev workflow

### Prerequisites

1. Install Go if you don't have it.
2. Make sure this directory is placed under \$GOPATH/src/github.com/joshjcarrier/thingiverse-cli.

To build the binary, run `make`.
