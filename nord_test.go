package nord

import "testing"

func TestParseHex(t *testing.T) {
	table := []struct {
		In, Out string
	}{
		{"#afc79e", "#afc79e"},
		{"#01AFCD", "#01afcd"},
		{"#a0c", "#aa00cc"},
	}

	for _, pair := range table {
		c, err := ParseHex(pair.In)
		out := c.Hex()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if out != pair.Out {
			t.Errorf("for %v expected %v got %v", pair.In, pair.Out, out)
		}
	}

}
