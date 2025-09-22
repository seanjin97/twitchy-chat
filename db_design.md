# Database design

There's only 2 main things we want to store.

1. User account
2. Messages

## User account

The query pattern is simply a directd look up based on the user's ID to fetch their details.

We don't really have any other query patterns.

```cql
CREATE TABLE users {
    user_id bigint,
    username TEXT,
    email TEXT
}
```

## Messages

The query pattern is to retrieve messages from a chatroom, with time ranges.

Generally, copying Discord since the query pattern is the same.

All IDs are snowflakes as we want the added convenience of being able to query using time ranges.

```cql
CREATE TABLE messages {
    chatroom_id bigint,
    bucket int,
    message_id bigint,
    author_id bigint,
    content text,
    PRIMARY KEY ((chatroom_id, bucket), message_id)
} WITH CLUSTERING ORDER BY (message_id DESC)

```
