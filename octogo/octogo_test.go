package octogo

import (
	//"image"
	//"image/color"
	"testing"
)

func TestInitialization(t *testing.T) {
	if encoders == nil {
		t.Error("encoders should be initialized")
	}
	if processors == nil {
		t.Error("processors should be initialized")
	}
}

func TestMeanHorizontal(t *testing.T) {

}

func TestGetModes(t *testing.T) {
	modes := GetModes()
	if len(modes) < 1 {
		t.Error("There should be at least one mode.")
	}

	if len(modes) != len(processors) {
		t.Error("There should be as much processors as modes.")
	}

	hasCopy := false
	for _, v := range modes {
		if v == "copy" {
			hasCopy = true
			break
		}
	}
	if !hasCopy {
		t.Error("There should be 'copy' in modes.")
	}
}
