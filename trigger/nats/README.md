<!--
title: Nats
weight: 4701
-->
# Nats Trigger

This trigger subscribes to a topic on Nats (nats.io) server and listens for the messages.

### Flogo CLI
```bash
flogo install github.com/s0rbus/flogo-components/trigger/nats
```

## Configuration

### Setting :

| Name       | Type   | Description
|:---        | :---   | :---     
| host       | string | host of nats server - ***REQUIRED***
| port       | int    | port of nats server - ***REQUIRED***

### HandlerSettings:

| Name       | Type   | Description
|:---        | :---   | :---   
| topic      | string | The Nats topic on which to listen for messages

### Output:

| Name         | Type     | Description
|:---          | :---     | :---   
| message      | string   | The message that was consumed


## Examples

```json
{
  "triggers": [
    {
      "id": "flogo-nats",
      "ref": "github.com/s0rbus/flogo-components/trigger/nats",
      "settings": {
        "host" : "localhost",
        "port" : 4442 
      },
      "handlers": [
        {
          "settings": {
            "topic": "syslog",
          },
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:my_flow"
            }
          }
        }
      ]
    }
  ]
}
```
 
## Development

### Testing

To be done.....
