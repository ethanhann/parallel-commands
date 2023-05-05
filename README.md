# parallel-commands

Simple Golang program to run multiple shell commands in parallel.

## Installation

Install:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/ethanhann/parallel-commands/HEAD/install.sh)"
```

Uninstall:

```bash
rm /usr/local/bin/pc
```

## Configuration

The default config file name is `commands.json`.

Config file format:

```json
{
  "commands": [
    "COMMAND_NAME"
  ]
}
```

Example:

```json
{
  "commands": [
    "ls -l",
    "ls -lhp"
  ]
}
```

## Usage

Simply
```shell
./pc
```

Or, with a custom config file:

```shell
./pc
```
