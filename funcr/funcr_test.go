package funcr

import (
	"encoding/json"
	"fmt"
	"testing"
)

type substr string

func ptrint(i int) *int {
	return &i
}
func ptrstr(s string) *string {
	return &s
}

func TestPretty(t *testing.T) {
	cases := []interface{}{
		"strval",
		substr("substrval"),
		true,
		false,
		int(93),
		int8(93),
		int16(93),
		int32(93),
		int64(93),
		int(-93),
		int8(-93),
		int16(-93),
		int32(-93),
		int64(-93),
		uint(93),
		uint8(93),
		uint16(93),
		uint32(93),
		uint64(93),
		uintptr(93),
		float32(93.76),
		float64(93.76),
		ptrint(93),
		ptrstr("pstrval"),
		[]int{9, 3, 7, 6},
		[4]int{9, 3, 7, 6},
		struct {
			Int    int
			String string
		}{
			93, "seventy-six",
		},
		map[string]int{
			"nine": 3,
		},
		map[substr]int{
			"nine": 3,
		},
		fmt.Errorf("error"),
		struct {
			X int `json:"x"`
			Y int `json:"y"`
		}{
			93, 76,
		},
		struct {
			X []int
			Y map[int]int
			Z struct{ P, Q int }
		}{
			[]int{9, 3, 7, 6},
			map[int]int{9: 3},
			struct{ P, Q int }{9, 3},
		},
		[]struct{ X, Y string }{
			{"nine", "three"},
			{"seven", "six"},
		},
	}

	for i, tc := range cases {
		ours := pretty(tc)
		std, err := json.Marshal(tc)
		if err != nil {
			t.Errorf("[%d]: unexpected error: %v", i, err)
		}
		if ours != string(std) {
			t.Errorf("[%d]: expected %q, got %q", i, std, ours)
		}
	}
}