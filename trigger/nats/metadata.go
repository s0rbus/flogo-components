package nats

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Host string `md:"host,required"`
	Port int `md:"port,required"`
}

type HandlerSettings struct {
	Topic string `md:"topic,required"`
}

type Output struct {
	Message string `md:"message"`
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.AnOutput, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": o.Message,
	}
}

