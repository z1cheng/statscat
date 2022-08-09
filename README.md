<div align="center">

# Stats Cat 🐈

Git statistics but with a cat

[![StatsCat](https://img.shields.io/badge/Repo-Stats%20Cat%20%F0%9F%90%88-yellow)](https://github.com/z1cheng/statscat) [![Language:Go](https://img.shields.io/static/v1?label=Language&message=Go&color=blue&style=flat-square)](https://go.dev) [![License:MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)](https://opensource.org/licenses/MIT) 
</div>

Stats Cat🐈 is a CLI tool to get statistics of your all git repositories.

![example](docs/example.gif)

## Installation

You need to confirm you have configured **Golang environment** beforce installing Stats Cat.

Then just run the following command:

```bash
go install github.com/z1cheng/statscat@latest
```

## Usage

```
Usage:
  statscat [-d dir] [-a author] [--since since] 
statscat -d ~/Public/projects -a z1cheng --since 1.year
Examples:

    statscat  # get the statistics of all repositories in current directory
    statscat -d /directory -a author --since 1.week  # get the statistics of all repositories under /directory, author is author name, since is from 1 week ago

Flags:
  -a, --author string   author name to be calculated, default is all authors
  -d, --dir string      directory to be calculated, statscat will search recursively, default is current directory (default ".")
  -h, --help            help for statscat
      --since string    show stats more recent than a specific date
```


