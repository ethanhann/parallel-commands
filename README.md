# dev

Simple Golang program to run multiple shell commands in parallel.

## Configuration

The default config file name is `dev.json`.

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
./dev 
```

or 