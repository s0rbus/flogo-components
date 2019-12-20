package addtodate

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	Number int `md:"number,required"`
   Units string `md:"units,required"`
   Date string `md:"date"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	var err error
   r.Number, err = coerce.ToInt(values["number"])
   if err != nil {
      return err
   }
   r.Units, err = coerce.ToString(values["units"])
   if err != nil {
      return err
   }
   r.Date, err = coerce.ToString(values["date"])
   if err != nil {
      return err
   }
   return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"number": r.Number,
		"units": r.Units,
		"date": r.Date,
	}
}

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
