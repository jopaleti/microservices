# How to define a protocol message?🚀
```
    Name of the message: UpperCammelCase
    Name of the field: lower_snake_case
    Data types:
        - string, bool, bytes
        - float, double
        - int32, int64, uint32, uint64, sint32, sint64 etc...
    Data type can be user defined enums
```

### Project Folder structure
```go
    protobuf/
        pb/
        proto/
            example.proto
```

### How To Generate Go Code And GRPC Code in pb/ Folder?
Run ... make gen
Manually run below command
```go
    protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
```

# In-Memory Storage in Computer Science 🧠💻
In-memory storage refers to the practice of storing data 
directly in a computer's RAM 
### Why Use In-Memory Storage?
1. **Speed 🚀:** RAM is significantly faster than hard drives
or SSDs. By storing data in memory, programs can access and manipulate it almost instantaneously.
2. **Efficiency ⚡**: Reduces the time spent on disk I/O operations, leading to quicker response times 
and better overall performance.

# ⏳ Context Timeout in Go
In Go, context is used to manage deadlines, cancellations across
API boundaries and goroutines.It is useful for avoiding long-running 
operations that might hang or take long time.

### 🛠️ How Context Timeout Works?
1. Creating a Context with Timeout:
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```
- **context.WithTimeout** creates a context that automatically cancels after the specified duration (e.g., 5 seconds).
- **defer cancel()** ensures the context is canceled when you're done.

2. Using the Context:
```go
select {
case <-ctx.Done(): // Context was canceled or timed out
    fmt.Println("Operation canceled or timed out:", ctx.Err())
}
```
- **<-ctx.Done()** listens for cancellation or timeout signals from the context.
- **ctx.Err()** provides the reason for cancellation or timeout.

# 🚀 Now, Let's Use In-Memory Storage with Context in Go
## 📊 Data Overview
We'll be handling laptop data generated from a gRPC-based microservice. 
This data includes various laptops categorized by brands and operating systems.

## 🎯 Objective 
Our goal is to efficiently save and process this laptop data in the 
computer's RAM (In-Memory storage). We will leverage Go’s context 
package to manage cancellation and timeouts, ensuring robust handling 
of client-server interactions.

## 🎯 Aim
We should be able to manage the context in such a way that:
1. **Data Integrity**: No laptop data is saved to In-Memory storage if the 
context exceeds its timeout duration.

2. Cancellation Handling: If the context is canceled (either due to a timeout
 or user cancellation), no data should be saved.
```
Example of context timeout canceled::
Client:> context.WithTimeout(context.Background(), 4*time.Second)
Server:> time.Sleep(5*time.Second)
Explanation: Server sleeps for 5 seconds and client request should
be terminated after 4 minutes, that means before server wake up the 
client request would have terminated and when such happen, you are to 
determine how the server should respond to such event either render it 
useless or still make the data persistent.
```
