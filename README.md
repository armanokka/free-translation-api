## The Free Translation API  written in Golang 🚀
Documentation is located in `docs` folder
#### 🔥 Live Example
Base URL: http://188.225.74.17/api/v1/ ![](http://188.225.74.17/uptime)

|            |                                         |
|------------|-----------------------------------------|
| Swagger UI | http://188.225.74.17/swagger/index.html |
| Jaeger UI  | http://188.225.74.17:16686/             |

#### 👨‍💻  Full list what has been used:
* [clean architecture](https://github.com/evrone/go-clean-template) - clean Architecture template for Golang services
* [testcontainers-go](github.com/testcontainers/testcontainers-go) - test containers for integration testing
* [otlptracegrpc](go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc) - OTLP span exporter using gRPC
* [cleanenv](github.com/ilyakaznacheev/cleanenv) - environment configuration reader
* [go-redis](https://github.com/go-redis/redis) - type-safe Redis client for Golang
* [otel](https://go.opentelemetry.io/otel) -  Go implementation of OpenTelemetry
* [jaeger](https://www.jaegertracing.io/) - open-source tracing platform
* [docker compose](https://docs.docker.com/compose/) - Docker compose
* [go-colorable](github.com/mattn/go-colorable) - colorful logging
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gin](https://github.com/gin-gonic/gin) - web framework
* [docker](https://www.docker.com/) - Docker
* [swag](https://github.com/swaggo/swag) - Swagger
* [grpc](https://grpc.io/) - gRPC
* [gorm](https://gorm.io/) - ORM
* [zap](https://github.com/uber-go/zap) - logger

### Running
Edit `.env.dev` file to your credentials. Then rename it to `.env` and run:

    make init

### Requirements
- go (1.20+)
- make
- docker
- docker compose


## License

MIT
