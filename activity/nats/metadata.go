package nats

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Server string `md:"server,required"`
	Topic string `md:"topic,required"`
}

type Input struct {
   Message string `md:"message,required"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
   i.Message, err = coerce.ToString(values["message"])
   return err
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
	}
}

/*
type Output struct {
	Result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
   o.Result, err = coerce.ToString(values["result"])
   if err != nil {
      return err
   }
   return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
*/

