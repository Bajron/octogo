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

func TestGetEncoders(t *testing.T) {
	encs := GetEncoders()
	if len(encs) < 1 {
		t.Error("There should be at least one encoder")
	}

	hasPng := false
	for _, v := range encs {
		if v == "png" {
			hasPng = true
			break
		}
	}
	if !hasPng {
		t.Error("There should be PNG encoder.")
	}

	if len(encs) != len(encoders) {
		t.Error("There should be as many encoders as registered.")
	}
}
