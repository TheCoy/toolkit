package toy

import (
	"encoding/json"
	"testing"
)

func TestJsonCompare(t *testing.T) {
	ljstr := `{
		 "aa":123,
		 "bb":45,
		 "cc":[
			   {"aaa1": {"aaaa1": 1}},
			   {"a1":"abdeghjl"}
		
			 ]
		
		}`
	rjstr := `{
		 "aa":12,
		 "bb":45,
		 "cc":[
			   {"aaa1": {"aaaa1": "a"}},
			   {"a1":"abdeghjl"}
		
			 ]
		
		}`
	ljson := make(map[string]interface{})
	rjson := make(map[string]interface{})
	if err := json.Unmarshal([]byte(ljstr), &ljson); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(rjstr), &rjson); err != nil {
		t.Fatal(err)
	}

	diff, hasDiff := JsonCompare(ljson, rjson, 1, "aa")
	t.Log(hasDiff)
	t.Log(diff)
}
