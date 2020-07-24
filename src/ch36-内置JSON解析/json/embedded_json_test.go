package jsontest

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
}
