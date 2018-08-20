[![](https://travis-ci.com/RecuencoJones/npmrc.svg?branch=develop)](https://travis-ci.com/RecuencoJones/npmrc)

# npmrc

A CLI to manage multiple npmrc files

## Installation from source

```
git clone https://github.com/RecuencoJones/npmrc
cd npmrc
go get -t ./...
go install ./...
```

## Installation from releases

- Download the file for your distribution
- Extract it to get `npmrc`
- Make it available in your path

Happy `.npmrc` switching!

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
version     Display version
```

## Configuration

| Env variable | Default value | Description |
| --- | --- | --- |
| NPMRC_DIR | `~` | Directory where profiles will be stored |
| EDITOR | `vi` | Editor to use for creating/editing profiles |
| VIEWER | `cat` | Viewer to use for viewing profiles |

## Profile name rules

- Profile name must not start with `.`
- Profile name must not contain spaces
- Profile name must be of at least 1 character length

[Check regex at regexr.com](https://regexr.com/3u5rg)
