package models

import "testing"

func TestChecksValidation(t *testing.T) {
	i := &Instrument{}
	err := i.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
