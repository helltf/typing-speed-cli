# typing-speed-cli

This simple cli Tool is meant to test your typing speed inside your terminal.
The cli is build with golang and cobra as the main library.

## Installation

To install this tool run following command

```sh
git clone git@github.com:helltf/typing-speed-cli.git
```

## Running 

To run the typing speed test cd into the folder and run the following command

```sh
go run . play
```

## Config 

The config currently contains 3 keys

```typescript
space: string
unit: "cpm" | "cps" | "wpm"
cursor: boolean
```

Set your configuration via: 

```sh
go run . config set --key=value
```
