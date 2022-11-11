package main

import "log"

func main() {
	type testCase struct {
		name string
		args []string
		want string
	}

	tests := []testCase{
		{
			name: "has prefix",
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
	}

	for _, t := range tests {
		if got := longestCommonPrefix(t.args); got != t.want {
			log.Fatalf("failed %s, got: %s, want: %s", t.name, got, t.want)
		}

		log.Println("TEST SUCCESS " + t.name)
	}

}
