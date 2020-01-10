# flogo-components
My collection of custom built flogo components

## Activities
* [addtodate](activity/addtodate): Add a specified number of days to either the current date or a chosen date
* [addtodate_legacy](activity/addtodate_legacy): As above legacy code version
* [echo](activity/echo): Simple echo based on flogo-project sample activity
* [nats](activity/nats): Send message to NATS server

## Triggers
* [nats](trigger/nats): NATS subscriber

## Functions
* [NotContains](function/notContainsFn): string does not contain substring

## Installation

#### Install Activity
Example: install **nats** activity
```bash
flogo install github.com/s0rbus/flogo-components/activity/nats
```

#### Install Trigger
Example: install **nats** trigger
```bash
flogo install github.com/s0rbus/flogo-components/trigger/nats
```

#### Install Function
Example: install **notContainsFn** function
```bash
flogo install github.com/s0rbus/flogo-components/function/notContainsFn
```

## License
Licensed under MIT license. See [LICENSE](LICENSE) for license text.


