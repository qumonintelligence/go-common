package lang

import (
	"fmt"
	"testing"
)

type SampleStruct struct {
	IntValue int `json:"int_value"`
}

func TestGetAsNil(t *testing.T) {
	m := StringMap{}
	x := SampleStruct{}
	err := m.GetAs("x", &x)
	if err == nil {
		t.Errorf("Should return err: %v", err)
		return
	}
}

func TestGetAs(t *testing.T) {
	m := StringMap{}
	m.Set("x", map[string]interface{}{"int_value": 101})
	x := SampleStruct{}
	err := m.GetAs("x", &x)
	if err != nil {
		t.Errorf("Should not return err: %v", err)
		return
	}

	if x.IntValue != 101 {
		t.Errorf("Should have int_value 101")
		return
	}
}

func TestGetAsWitPathNotFound(t *testing.T) {
	m := StringMap{}
	x := SampleStruct{}
	err := m.GetAsByPath("x.y.z", &x)

	if err == nil {
		t.Errorf("Should return err: %v", err)
		return
	}
}

func TestGetAsWithPath(t *testing.T) {
	m := StringMap{}
	x := SampleStruct{}

	m.Set("geo", map[string]interface{}{"int_value": 100})
	err := m.GetAsByPath("geo", &x)
	if err != nil {
		t.Errorf("Should not return err: %v", err)
		return
	}

	if x.IntValue != 100 {
		t.Errorf("Should have intValue 100, got %v", x.IntValue)
		return
	}
}

func TestGetAsWithMultiPath(t *testing.T) {
	m := StringMap{}
	x := SampleStruct{}

	m.Set("geo", map[string]interface{}{"plugins": map[string]interface{}{"int_value": 100}})

	fmt.Printf("%v\n\n", m)

	err := m.GetAsByPath("geo.plugins", &x)
	if err != nil {
		t.Errorf("Should not return err: %v", err)
		return
	}

	if x.IntValue != 100 {
		t.Errorf("Should have intValue 100, got %v", x.IntValue)
		return
	}
}
