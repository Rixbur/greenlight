# green light

### Running the app
```
go run ./cmd/api
```

### curl
#### curl defaults to GET when we don't pass -X
```sh
curl localhost:4000/v1/movies/123
curl localhost:4000/v1/health-check
```

#### curl defaults to POST when we pass -d
`-i is used to include the response headers in the output`
```sh
BODY='{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}'
curl -i -d $BODY localhost:4000/v1/movies
curl -i -X POST -d $BODY localhost:4000/v1/movies # same as above, just explicitly specifying POST
```