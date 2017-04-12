# goREST
Golang simple library for message passing using REST


[![Build Status](https://travis-ci.org/alruiz12/goREST.svg?branch=master)](https://travis-ci.org/alruiz12/goREST)
[![codecov](https://codecov.io/gh/alruiz12/goREST/branch/master/graph/badge.svg)](https://codecov.io/gh/alruiz12/goREST)
[![Code Health](https://landscape.io/github/alruiz12/goREST/master/landscape.svg?style=flat)](https://landscape.io/github/alruiz12/goREST/master)

___
## Install
* Download Go and follow [the official guidelines for Code organization](https://golang.org/doc/code.html#Organization) 

* Clone, download the project or use go get:
`go get github.com/alruiz12/goREST`

* Find the Main.go file under the project root

* Use your preferred network interface (loopback default) or IP address for the server:
```go
serverIP:=config.GetMyIP("lo")
```

* Change default ports if necessary:
```go
serverPort:="8888"
clientPort:="8080"
```

* Modify times at your convenience:
```go
var interval time.Duration=2
var finishTime time.Duration=9
```

* Modify the message you want to send to the server:
```go
message:="hello!"
```

