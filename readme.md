# Twitchy Chat

A clone of Twitch chat, with added skbidification of text messages!:

Original: `hi there`

After post processing: `Hey bestie! ✨💅  Fr fr, hi bestie!  No cap, slayyy.  Busssin.  🤪`

# Solution Design

## At a glance
We are building a high throughput, low latency, live message chat service.

## Functional requirements
1. Chat:
   * Message sending and receiving in real-time.
   * Skibidification of messages.
   * User should be able to scroll back in time to read previous messages before they joined the chat.
2. User accounts:
   * Only registered users can send messages.

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
   * This should not be a primary concern for this system, as it is primarily used to send and store nonsense.