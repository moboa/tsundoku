# Tsundoku
Tsundoku is a command-line manga downloader written in Go.

## Supported sites
* [Mangareader](https://www.mangareader.net/)
* [Mangapark](https://www.mangapark.me/)
* [Mangapanda](https://www.mangapanda.com)
* [Mangastream](https://readms.net/)

## Installation
* Install go from [https://golang.org/](https://golang.org/)
* Run the following commands in a bash terminal
```
$ go get github.com/moboa/tsundoku
$ cd ~/go/src/github.com/moboa/tsundoku/cmd/tsundoku
$ go build
# Add the following line to .bashrc to use it beyond one terminal session
$ alias tsundoku='~/go/src/github.com/moboa/tsundoku/cmd/tsundoku'
```

## Usage
`tsundoku [--help/-h] [--version/-v] [--output/-o] [<url>]`

For example,
```
$ tsundoku -o kuroko https://www.mangareader.net/kuroko-no-basket/1/2
Downloaded 56 pages from https://www.mangareader.net/kuroko-no-basket/1/2
```
will download chapter 1 of `Kuroko No Basket`, create a `kuroko` folder if one does not exists,
and store the image files in that folder.
