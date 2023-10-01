package main

import "testing"

func TestRemoveDoneTask(t *testing.T) {
	testcase := map[string]struct {
		target string
		want   string
	}{
		"done task should be removed":             {target: "- [x] done task description", want: ""},
		"done nested task should be removed":      {target: "  - [x] done nested task description", want: ""},
		"doing task should not be removed":        {target: "- [ ] doing task description", want: "- [ ] doing task description"},
		"doing nested task should not be removed": {target: "  - [ ] doing task description", want: "  - [ ] doing task description"},
		"multiple line": {
			target: `- [ ] line 1
              - [x] line 2
              - [ ] line 3
                - [x] line 4
              - [ ] line 5`,
			want: `- [ ] line 1
              - [ ] line 3
              - [ ] line 5`,
		},
	}

	for name, tc := range testcase {
		t.Run(name, func(t *testing.T) {
			got := RemoveDoneTask(tc.target)
			if got != tc.want {
				t.Errorf("\ntarget: %s\nwant: %s\ngot: %s\n", tc.target, tc.want, got)
			}
		})
	}
}
