# What is Microservices?
Microservices is an architectural style where an application is developed
as a collection of small, independent services that communicate with each 
other through well-defined APIs.

# Benefits of Microservices:
1.0 Scalability
2.0 Flexibility
3.0 Resilience: Failure in one service doesn’t necessarily bring down the entire system.
4.0 Faster Development: Teams can work on different services simultaneously, 
                        leading to faster development cycles.

# Challenges Faced by Microservices:
1.0 Managing service interactions
2.0 Ensuring data consistency
    Data consistency: It ensures that data remains accurate, reliable, and uniform across a system
3.0 Dealing with increased complexity in deployment and monitoring

# Best Ways to Establish Communication between Microservices?
1.0 gRPC
2.0 Message Brokers
    Description: Services communicate asynchronously through a message broker (e.g., RabbitMQ, Apache Kafka, or Amazon SQS).
    Pros: Decouples services, supports asynchronous processing, and can handle high volumes of messages.
    Cons: Adds complexity in terms of message management and requires handling of potential message loss or duplication.
3.0 Event Streaming
    Description: Services publish and subscribe to events using a streaming platform (e.g., Apache Kafka, AWS Kinesis).
    Pros: Real-time data processing, allows event-driven architectures, and handles high throughput.
    Cons: Requires management of event schemas and consumer state.
