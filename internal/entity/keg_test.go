package entity

import (
	"encoding/json"
	"testing"
)

func TestKegDto(t *testing.T) {
	kdto := KegDto{}
	blob, err := json.MarshalIndent(&kdto, "", "\t")
	if err != nil {
		t.Logf("%+v\n", kdto)
		t.Fatal(err)
	}
	t.Log(string(blob))
}
