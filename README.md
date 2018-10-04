# fasthttpcors test

## Pre-flight request

```shell
 curl -H "Origin: http://localhost1" -H "Access-Control-Request-Method: GET" -X OPTIONS --verbose http://localhost:8080/test
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> OPTIONS /test HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
> Origin: http://localhost1
> Access-Control-Request-Method: GET
> 
< HTTP/1.1 200 OK
< Server: fasthttp
< Date: Thu, 04 Oct 2018 23:02:26 GMT
< Content-Length: 0
< Access-Control-Allow-Origin: http://localhost1                     <--- OK
< Access-Control-Allow-Methods: GET
< 
* Connection #0 to host localhost left intact
```

## GET request

```shell
curl -v http://localhost:8080/test -H 'Origin: http://localhost1'
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /test HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
> Origin: http://localhost1
> 
< HTTP/1.1 200 OK
< Server: fasthttp
< Date: Thu, 04 Oct 2018 23:01:14 GMT
< Content-Type: text/plain; charset=utf-8
< Content-Length: 25
< Access-Control-Allow-Origin: http://localhost1                     <--- OK
< 
* Connection #0 to host localhost left intact
Requested path is "/test"
```

