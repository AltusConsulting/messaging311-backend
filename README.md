# messaging-backend

This is a backend for pushing notifications when reports are created with [Smart Cities](https://gitlab.priv.als/witcher/Smart_Cities/commits/open-source) application.

## Getting Started

To start using this project first you must install [Golang] (https://golang.org/dl/). 


### Prerequisites

The following dependencies are used in this project:

* [Gin-Gonic](https://github.com/gin-gonic/gin) - HTTP Web Framework 
* [Viper](https://github.com/spf13/viper) - Library used for giving a configuration solution
* [Resty](https://github.com/go-resty/resty) - Library of HTTP & REST Client 

In your [config.toml](config.toml) file you can configure the key server used for pushing notifications once you created an account on [Firebase web page](https://firebase.google.com/).
```toml
[fcm]
host = "https://fcm.googleapis.com/fcm/send"
keyServer = "PUT HERE YOUR KEY SERVER"

[server]
port = 9003
```


### Installing

Once Golang is installed clone this repository and run the following command at your project root
```
go get
```
This will download all the dependencies required. 


## Deployment

Once you have installed Golang, go to the project's root and execute the following command in your terminal:

```
go run main.go
```

You could use [Postman](https://www.getpostman.com/postman) for testing all endpoints of this backend.

## Contributing

...

## Versioning

This is the initial version We use [SemVer](http://semver.org/) for versioning.

## Authors

* [Altus Consulting Software Team] (https://github.com/AltusConsulting)

## License

...