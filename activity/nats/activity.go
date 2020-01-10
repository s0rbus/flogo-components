package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{})

// Activity is a nats activity
type Activity struct {
	conn  *nats.Conn
	topic string
}

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	// Connect Options.
   opts := []nats.Option{nats.Name("NATS Flogo Publisher")}

	conn, err := nats.Connect(s.Server, opts...)
   if err != nil {
      return nil,err
   }

	act := &Activity{conn: conn, topic: s.Topic}

	return act, nil
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if input.Message == "" {
		return false, fmt.Errorf("no message to publish")
	}

	ctx.Logger().Debugf("sending NATS message")

	a.conn.Publish(a.topic, []byte(input.Message))
   a.conn.Flush()

   if err := a.conn.LastError(); err != nil {
      return false, err
   } else {
      ctx.Logger().Debugf("Published [%s] : '%s'\n", a.topic, input.Message)
   }


	//output := &Output{}
	//err = ctx.SetOutputObject(output)
	//if err != nil {
	//	return false, err
	//}

	return true, nil
}
