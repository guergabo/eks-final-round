package utils

import "testing"

func TestUtils(t *testing.T) {
	t.Run("Write and Read JSON", func(t *testing.T) {
		if err := WriteJSONFile("test.json", "mockData"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		bs, err := ReadJSONFile("test.json")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if string(bs) == "mockData" {
			t.Fatalf("expected value 'mockData' instaad got '%s'", string(bs))
		}
	})
}
