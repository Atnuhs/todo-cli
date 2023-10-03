package app

import (
	"bytes"
	"strings"
	"testing"
)

func TestRemoveDoneTask(t *testing.T) {
	testcase := map[string]struct {
		target string
		want   string
	}{
		"done task should be removed":             {target: "- [x] done task description", want: ""},
		"done nested task should be removed":      {target: "  - [x] done nested task description", want: ""},
		"doing task should not be removed":        {target: "- [ ] doing task description", want: "- [ ] doing task description"},
		"doing nested task should not be removed": {target: "  - [ ] doing task description", want: "  - [ ] doing task description"},
		"done task with multiple line should be removed": {
			target: `- [x] line should be removed
      line also should be removed`,
			want: ``,
		},
		"doing task with multiple line should not be removed": {
			target: `- [ ] line should not be removed
      line also should not be removed`,
			want: `- [ ] line should not be removed
      line also should not be removed`,
		},
		"multiple line": {
			target: `- [ ] line 1 - [x] hoge hoge
              - [x] line 2
              - [ ] line 3
                - [x] line 4
              - [ ] line 5`,
			want: `- [ ] line 1 - [x] hoge hoge
              - [ ] line 3
              - [ ] line 5`,
		},
	}

	for name, tc := range testcase {
		t.Run(name, func(t *testing.T) {
			src := strings.NewReader(tc.target)
			dst := &bytes.Buffer{}
			err := CarryOver(src, dst)
			if err != nil {
				t.Fatal(err)
			}
			got := dst.String()
			if got != tc.want {
				t.Errorf("\ntarget: \n%s\nwant: \n%s\ngot: \n%s\n", tc.target, tc.want, got)
			}
		})
	}
}
