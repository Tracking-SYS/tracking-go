# TRACKING SYSTEM
[![](https://travis-ci.com/Tracking-SYS/tracking-go.svg?branch=main)](https://travis-ci.com/github/Tracking-SYS/tracking-go/builds) 

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Tracking-SYS_tracking-go&metric=bugs)](https://sonarcloud.io/dashboard?id=Tracking-SYS_tracking-go) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Tracking-SYS_tracking-go&metric=code_smells)](https://sonarcloud.io/dashboard?id=Tracking-SYS_tracking-go) [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=Tracking-SYS_tracking-go&metric=ncloc)](https://sonarcloud.io/dashboard?id=Tracking-SYS_tracking-go) 

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Tracking-SYS_tracking-go&metric=security_rating)](https://sonarcloud.io/dashboard?id=Tracking-SYS_tracking-go) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=Tracking-SYS_tracking-go&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=Tracking-SYS_tracking-go)
### Installation

1. Set GOPRIVATE environment to download private package
```
export GOPRIVATE=github.com/tikivn
```
2. Compile & build execute binary file
```
make default
```
3. Up server
```
make run
```
4. Install migrate command
```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.darwin-amd64.tar.gz | tar xvz
mv migrate.darwin-amd64.tar.gz migrate
```
5. Run migration with version 1
```
./migrate.sh up 1
```
### Tips

Downgrade Go modules
```
go get -u github.com/apache/thrift@v0.13.0
go mod tidy
go clean -modcache
go mod vendor
```

### Kafka Setup
1. Edit configuration of kafka advertise listener to public interface (public IP of VM)
```
KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://broker:29092,PLAINTEXT_HOST://[public IP]:9092
```

### Locust 
* How to write locustfile
[Documentation](https://docs.locust.io/en/latest/writing-a-locustfile.html)

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg)](CODE_OF_CONDUCT.md)
