# effective-octo-parakeet
Trying, improving and understading of redis, helm and stuff alike

There are two services comunicating using Redis Pub:

- receiver
- sender

## Receiver

Using go-chi as router, we receive messages from the 'sender', we store them it in Redis as well.
We expose http endpoints to see all/specific messages.

## Sender

Using go-chi as router, we expose http endpoints where we send messages that in turn will be passed to the 'receiver'

Techonoligies:

- Helm, for installing redis
- Go
- Redis PubSub and as DB
- K8s
- Docker

