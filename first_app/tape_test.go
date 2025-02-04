package poker

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	want := "abc"
	tape := &tape{file}
	tape.Write([]byte(want))

	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
