// Package addtodate adds a specified number of units to a date.
package addtodate

import (
	"io/ioutil"
	"testing"

	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = &activity.Metadata{}
		err = activityMetadata.UnmarshalJSON(jsonMetadataBytes)
		if err != nil {
			panic("Error unmarshalling Json Metadata")
		}

		//activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	//tc := test.NewTestActivityContext(getActivityMetadata())
	//ref := activity.GetRef(&activity.Activity{})
	tc := test.NewActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("number", 8)
	tc.SetInput("units", "days")
	tc.SetInput("date", "2019-12-17")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, result, "2019-12-25")
}
