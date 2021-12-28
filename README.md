# secret-service
A tiny and simple secret service


## run/dev
```bash
go run cmb/secretservice/main.go
```

As of now the http service defaults to port :3000.


## example

```
$ curl http://localhost:3000
404 page not found
$ curl http://localhost:3000/get
id not provided
$ curl "http://localhost:3000/get?id=andrew"
'andrew' not found
$ curl "http://localhost:3000/set?id=andrew" --data '{"key": "value"}'
{"key":"value"}
$ curl "http://localhost:3000/get?id=andrew"
{"key":"value"}
$ curl "http://localhost:3000/delete?id=andrew"
$ curl "http://localhost:3000/get?id=andrew"
'andrew' not found
$ curl "http://localhost:3000/set?id=andrew" --data '{"key": "value"}'
{"key":"value"}
$ curl "http://localhost:3000/get?id=andrew"
{"key":"value"}
```





