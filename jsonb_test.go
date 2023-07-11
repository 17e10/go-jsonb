package jsonb

import (
	"os"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	f, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err)
	}
	var filename = f.Name()
	f.Close()
	defer os.Remove(filename)

	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	john := person{"john", 20}

	if err := Save(filename, &john); err != nil {
		t.Fatal(err)
	}

	var readed person
	if err := Load(filename, &readed); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(readed, john) {
		t.Errorf("%s = %#v, want %#v", "person", readed, john)
	}
}
