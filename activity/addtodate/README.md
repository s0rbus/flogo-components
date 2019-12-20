# Add to Date
This activity provides your Flogo app the ability to add a specified number of units to a date

Based on original code from 'retgits':
https://github.com/retgits/flogo-components/blob/master/activity/addtodate

This version has been updated for flogo 0.9.2


## Installation

```bash
flogo install github.com/s0rbus/flogo-components/activity/addtodate
```
Link for flogo web:
```
https://github.com/s0rbus/flogo-components/activity/addtodate
```

## Schema
Inputs and Outputs:

```json
{
"input":[
      {
        "name": "number",
        "type": "integer",
        "reuired": true
      },
      {
        "name": "units",
        "type": "string",
        "allowed" : ["years", "months", "days"]
      },
      {
        "name": "date",
        "type": "string"
      }
    ],
    "output": [
      {
        "name": "result",
        "type": "string"
      }
    ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| number | The number of units to add to the date |
| units  | The units to add (allowed values are years, months and days) |
| date   | The date to add the units to (must be in the format YYYY-MM-DD). If this is blank, the current date will be chosen |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| result    | The new date (will be in the format YYYY-MM-DD) |
