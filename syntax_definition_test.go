package ldif

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestExample(t *testing.T) {
	raw, _ := ioutil.ReadFile("testdata/example1.ldif")
	f := File([]rune(string(raw))).Best()

	t.Run("version", func(t *testing.T) {
		if f.GetSubNode("version:") == nil {
			t.Error("version not found")
			return
		}

		if nr := f.GetSubNode("version-number"); nr == nil {
			t.Error("version number not found")
		} else if v := nr.Value; string(v) != "1" {
			t.Errorf("version number was not 1, got %s", string(v))
		}
	})

	t.Run("records", func(t *testing.T) {
		records := f.GetSubNodes("ldif-attrval-record")
		if l := len(records); l != 2 {
			t.Errorf("did not find 2 records, got %d", l)
		}
	})
}

func TestExamples(t *testing.T) {
	for i := 1; i < 8; i++ {
		t.Run(fmt.Sprintf("example%d", i), func(t *testing.T) {
			raw, _ := ioutil.ReadFile("testdata/example1.ldif")
			if File([]rune(string(raw))) == nil {
				t.Errorf("could not parse ldif file: example%d.ldif", i)
			}
		})
	}
}
