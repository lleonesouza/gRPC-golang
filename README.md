# protogRPC

the proto files to generate the protobuffer code in go using protoc:

```
protoc service.proto
```

# protoRPC - server

protoRPC server - receive the proto-buffer by tpc, compute and send to the client.

Actions:
Multiply - compute a * b
Add - compute a + b

# protoRPC - client

protoRPC client - receive the proto-buffer by tpc and send via http.

Routes:
GET: '/multiply/:a/:b'
GET: '/add/:a/:b'