<!--suppress ALL -->
<img height="310" alt="image" src="https://user-images.githubusercontent.com/16114089/121805566-0cd81280-cc4c-11eb-9b7d-b5f8a6db4f8d.png" align="right">

# MockingJay

MockingJay is a heavily asynchronous server software for Minecraft Bedrock Edition written in Go. This software was forked from [Dragonfly](https://github.com/df-mc/dragonfly). It is being modified to accomodate a less source-heavy development style, allowing for plugins to be developed and shared. MockingJay is generally meant for use as a library.

## Getting started
Running MockingJay requires at least **Go 1.18**. After starting the server through one of the methods below,
**ctrl+c** may be used to shut down the server. Also check out the [wiki](https://github.com/xJustJqy/MockingJay/wiki) for
more detailed info.

#### Installation as library
```
go mod init github.com/<user>/<module name>
go get github.com/xJustJqy/MockingJay
```

#### Installation of the latest commit
```
git clone https://github.com/xJustJqy/MockingJay
cd mockingjay
go run main.go
```

## Developer info
[![Go Reference](https://pkg.go.dev/badge/github.com/xJustJqy/MockingJay/server.svg)](https://pkg.go.dev/github.com/xJustJqy/MockingJay/server)

MockingJay features a well-documented codebase with an easy-to-use API. Documentation may be found
[here](https://pkg.go.dev/github.com/xJustJqy/MockingJay/server) and in the subpackages found by clicking *Directories*.

Publishing your project on GitHub? Consider adding the **[#mockingjay-server](https://github.com/topics/mockingjay-server)** topic to your
repository to improve visibility of your project.

## Contributing
Contributions are very welcome! Issues, pull requests and feature requests are highly appreciated. Opening a pull
request? Also have a read through the [CONTRIBUTING.md](https://github.com/xJustJqy/MockingJay/blob/master/.github/CONTRIBUTING.md) for more info.
