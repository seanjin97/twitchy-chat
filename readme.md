# Twitchy Chat

A clone of Twitch chat, with added skbidification of text messages and chat rooms!

Original: `hi there`

After post processing: `Hey bestie! ✨💅  Fr fr, hi bestie!  No cap, slayyy.  Busssin.  🤪`

# Solution Design

## At a glance
We are building a high throughput, low latency, live message chat service.

## Functional requirements
1. Chat:
   * There will be a preset number of chatrooms, each with a fixed preset of skibidification style.
   * Multiple users can join a chatroom.
   * Message sending and receiving in real-time.
   * Skibidification of messages.
   * On joining a chat room, the last 10 messages are fetched.
   * ~~Users should be able to search for old messages using a keyword, which will then return you a list of messages containing that keyword.~~ (Assume the need to store messages for long term.)
2. User accounts:
   * Only registered users can send messages.
   * Non-registered users can join chat rooms, and only read messages.

## Non-functional requirements
1. Scalability:
   * The system should be able to handle a large number of concurrent users and messages.
2. Performance:
    * High throughput: The system should be able to handle a high volume of messages per second.
    * Messages should be delivered in real-time with low latency.
    * The skibidification process should be fast enough to not introduce significant latency.
    * The system should be able to handle spikes in traffic, such as during popular chat sessions.
    * We are more concerned with write performance than read performance. Since, in a live chat, there are typically more messages being sent than read at any given time.
3. Availability (is the system often online?):
   * The system should be highly available, with minimal downtime.
4. Reliability (does the system often perform it's core functions):
   * Messages should not be lost in case of failures.
   * Messaging
5. Consistency:
   * During high traffic, it is acceptable for messages to be delivered out of order.
   * Consistency is not a primary concern, as long as messages are eventually delivered.
6. Security:
   * This should not be a primary concern for this system, as all messages are nonsense.

## Capacity planning
### We expect:

#### General

> No. of chatrooms: 3
> 
> Average number of concurrent users per chatroom: 5,000
> 
> Maximum number of concurrent users per chatroom: 50,000
> 
> Average number of messages per user per chatroom: 50
>
> Average session duration: 30 mins
> 
> Average size of a message: 150 bytes

#### Message sending
> Average number of messages sent per chatroom: 5000 * 50 = 250,000
> 
> Maximum number of messages sent per chatroom (assuming no. of messages sent stay the same): 50,000 * 50 = 2,500,000
> 
> Average no. of messages sent in total: 3 * 250,000 = 750,000
> 
> Max no. of messages sent in total: 3 * 2,500,000 = 7,500,000
> 
> Average number of messages sent per chatroom per second: 250,000 / (30 * 60s) = 139 messages/s
> 
> Maximum number of messages sent per chatroom per second:  2,5000,000 / (30m * 60s) = 13,889 messages/s
>
> Average qps for message sending: 750,000 / (30 * 60s) = 417 messages/s
> 
> Maximum qps for message sending: 7,500,000 / (30 * 60s) = 41,667 messages/s
> 
> Maximum throughput of messages per second: 41,667 * 150 bytes = ~6 MB/s
> 
> Maximum throughput of messages per day: 6 MB/s * 86400s = ~518 GB/day

#### Message reading
> Max no. of read queries per chatroom in a session: 50,000
> 
> Max no. of read queries in a session: 3 * 50,000 = 150,000
> 
> Max no. of read queries per session per second: 150,000 / (30 * 60) = 84 queries/s


### Capacity analysis
We observe that we generally have a high write to read ratio. We will need to design our system to handle high write throughput.

## System architecture
### Components
1. Client:
   * Simple client side rendered framework with good 3rd party library support such as React would be sufficient.
2. API & Websocket Server:
    * REST API server for user authentication, chatroom management, and message sending.
      * Why not gRPC or tRPC?
        * gRPC: For our use this is overkill, since we don't need extremely high performance for the general sign in, register, create chat APIs. Generally, our system will not have many API complex API routes. We don't need to manage the type complexity with tRPC.
        * tRPC: It does give nice DX, but generally it won't help with performance since it is just a layer on HTTP. Generally, our system will not have many API complex API routes. We don't need to manage the type complexity with tRPC.
    * WebSocket server for real-time message delivery.
      * Why not IRC?: We will need to enshitify an original user's messages on the server side, which would require 2 way communication and more fine-grained controls and logic over the processing behaviour!
    * This company is also has low funding (im poor), we will need to keep our server costs low.
      * We will need a programming language that is performant, has good concurrency support, and has good library support for WebSocket and REST API.
    * A good candidate would be golang for REST APIs and a WebSocket server.
3. Skibidification service:
   * This is a core component of the system, as it is responsible for transforming user messages into skbidified (nonsense) messages.
   * It should be able to handle high throughput and low latency.
   * The skibidification process is an API call to a LLM model.
   * The LLM is a small self hosted, open sourced model to keep costs low and to keep throughput and latency high.
   * The Skibidification service should be stateless and horizontally scalable.
   * A good candidate for the LLM would be LLaMA 3.2 cluster.
4. Data store:
   * We won't need strong consistency.
   * We will require horizontal scaling to support high write throughput.
   * Most of our data will be timeseries data (messages with timestamps).
   * A good candidate would be a NoSQL database like Cassandra.
5. Message broker:
   * The core of the application is the live chat functionality. In order to support max 42k messages/s, we are looking at a message broker that can support thousands of requests per second
   * As we are looking to increase the scale of this platform, we do expect the number of messages to increase, and we want a message broker that can scale horizontally.
   * A good candidate would be Apache Kafka.


## Questions
1. How many connections can a single WebSocket server handle?
   * A single WebSocket server can handle around ?? connections. We will need to scale horizontally to support more connections.
2. How many messages can a single Skibidification service handle?
    * A single Skibidification service can handle around ?? messages per second. We will need to scale horizontally to support more queries.
3. How many messages can a single database instance handle?
    * A single database instance can handle around ?? messages per second.
4. How many messages can a single message broker instance handle?
    * A single message broker instance can handle around ?? messages per second.