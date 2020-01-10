# NATS

This activity provides your Flogo app the ability to send a message to a NATS server

## Installation

```bash
flogo install github.com/s0rbus/flogo-components/activity/nats
```
Link for flogo web:
```
https://github.com/s0rbus/flogo-components/activity/nats
```

## Configuration

### Settings:
| Name        | Type   | Description
|:---         | :---   | :---
| server      | string | NATS server (format is servername:port) - ***REQUIRED*** 
| topic       | string | topic on which to publish message - ***REQUIRED*** 

### Input
| Name        | Type   | Description
|:---         | :---   | :---
| message     | string | The message to send

## Example

The json below sends 'Hello from Flogo' to topic 'general' on a NATS server running on localhost

```json
{
  "id": "publish_nats_message",
  "name": "Publish message to NATS",
  "activity": {
     "ref": "github.com/s0rbus/flogo-components/activity/nats",
     "settings": {
        "server" : "localhost:4222",
        "topic" : "general"
     },
     "input": {
        "message": "Hello from Flogo"
     }
  }
}
```

## Development

### Testing

To run tests first set up a NATS server using the docker-compose file given below:

```yaml
version: '3.7'
services:
    server:
        image: nats
        container_name: natsio-server
        ports:
            - '4222:4222'
            - '8222:8222'

```

Then run the following command:

```bash
go test
```

