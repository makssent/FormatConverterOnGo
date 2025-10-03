package convertor

import "testing"

func TestConvertStub(t *testing.T) {
	in := []byte("test")
	out, err := Convert(in, "json", "yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(out) != "test" {
		t.Fatalf("expected 'test', got %s", out)
	}
}
