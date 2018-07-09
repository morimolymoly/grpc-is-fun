# grpc-is-fun
golang grpc tutorial

## HelloWorld

### protoc
```
protoc --go_out=plugins=grpc:./ helloworld/pb/hello.proto
```

### Usage
```
docker-compose -f docker-compose.helloworld.yml build
docker-compose -f docker-compose.helloworld.yml up server
docker-compose -f docker-compose.helloworld.yml up client
```

### How to envoke API
I used grpcc for grpc client.
```
grpcc --proto helloworld/pb/hello.proto --address localhost:8100 -i

Connecting to pb.HelloWorldService on localhost:8100. Available globals:

  client - the client connection to HelloWorldService
    sayHello (SayHelloRequest, callback) returns HelloReply

  printReply - function to easily print a unary call reply (alias: pr)
  streamReply - function to easily print stream call replies (alias: sr)
  createMetadata - convert JS objects into grpc metadata instances (alias: cm)
  printMetadata - function to easily print a unary call's metadata (alias: pm)
HelloWorldService@localhost:8100> client.sayHello({name: "もりもりもりもり"}, pr)
EventEmitter {}
HelloWorldService@localhost:8100>
{
  "message": "Hello, もりもりもりもり"
}
```
