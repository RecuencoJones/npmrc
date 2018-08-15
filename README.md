# npmrc

A CLI to manage multiple npmrc files

## Usage

```
npmrc <command> [flags]

Available commands:

use, u      Select a profile
view, v     View a profile
list, ls    List available profiles
edit, ed    Create or update profiles
copy, cp    Copy profiles
remove, rm  Remove a profile
help, h     Display this message
```

## Configuration

| Env variable | Default value | Description |
| --- | --- | --- |
| NPMRC_DIR | `~` | Directory where profiles will be stored |
| EDITOR | `vim` | Editor to use for creating/editing profiles |
| VIEWER | `cat` | Viewer to use for viewing profiles |
