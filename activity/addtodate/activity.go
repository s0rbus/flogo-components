package addtodate

import (
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	//ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	date := time.Now()

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if input.Date != "" {
		date, _ = time.Parse("2006-01-02", input.Date)
	}

	switch input.Units {
   case "days":
      date = date.AddDate(0, 0, 1*input.Number)
   case "months":
      date = date.AddDate(0, 1*input.Number, 0)
   case "years":
      date = date.AddDate(1*input.Number, 0, 0)
   }

	output := &Output{Result: date.Format("2006-01-02")}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
