#What is gRPC?
A high performance, open source universal RPC framework,
Originally developed by Google

#What is RPC (Remote Procedure Calls?)
- It is a protocol that allows program to:  
    a. execute Procedure of another program located in other computer
    b. without developer explicitly coding the details for the remote interactions
- In the client code it's as if calling a server code function directly
- Client and server codes can be written in different languages 

#How gRPC Works?

1. Protocol Buffers (Protobuf): gRPC uses Protocol Buffers as its "Interface Definition Language(IDL)"
    to define structure of the message and services.
    - Usage: You define your services and message types in a .proto file, which is then compiled into code
    in your preferred programming languages.

2. Service Definition:
    Services and methods are defined in a .proto file using the Protobuf language.
    ```
    service Greeter {
        rpc SayHello (HelloRequest) returns (HelloReply);
    }

    message HelloRequest {
        string name = 1;
    }

    message HelloReply {
        string message = 1;
    }
    ```

3. Code Generation:
    The .proto file is compiled using the Protobuf compiler (protoc) to generate client and server code

4. Communication:
    gRPC uses HTTP/2 as its transport protocol, which provides features like multiplexing, flow control, header compression

5. Communication Patterns:
    1. Unary RPC: A single request and a single response. 
    2. Server Streaming RPC: The client sends a single request and receives a stream of responses.
    3. Client Streaming RPC: The client sends a stream of requests and receives a single response.
    4. Bidirectional Streaming RPC: Both the client and server send and receive streams of messages.

6. Authentication and Security:
    gRPC supports authentication and encryption using SSL/TLS

7. Load Balancing and Fault Tolerance:
    gRPC supports client-side load balancing, retries, and deadlines to handle distributed system complexities.