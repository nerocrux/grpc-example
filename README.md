# grpc-example

* example for using go-grpc
* https://github.com/mattn/grpc-example

## generate pb.go

```
$ cd proto
$ protoc --go_out=plugins=grpc:. nogizaka.proto
```

## run server

```
$ go run server/main.go
```

## run client

### add member

```
$ go run client/main.go add 1 erika.ikuta 1997-1-22 1
```

### list member

```
$ go run client/main.go list
```
