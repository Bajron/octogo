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
