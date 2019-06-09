# GO Starter
- Prepare :
```shell
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

- Run :

```shell
go run gilded-rose.go
```

- Run tests :

```shell
go test
```

- Run tests and coverage :

```shell
go test -coverprofile=coverage.out

go tool cover -html=coverage.out
```
