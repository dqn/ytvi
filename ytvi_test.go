package ytvi

import (
	"os"
	"testing"
)

func TestGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(os.Getenv("VIDEO_ID"))
	if err != nil {
		t.Fatal(err)
	}
}
