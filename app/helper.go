package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"time"
)

const TimeFormat = "20060102"

var rDoneTask = regexp.MustCompile(`^\s*- \[x\] `)

func TimeToBasename(t time.Time) string {
	return t.Format(TimeFormat)
}

func BaseNameToTime(basename string) (time.Time, error) {
	return time.Parse(TimeFormat, basename)
}

func ShowContent(r io.Reader) string {
	var bf bytes.Buffer
	bf.ReadFrom(r)
	return bf.String()
}

func CarryOver(src io.Reader, dst io.Writer) error {
	r := bufio.NewReader(src)
	sp := -1
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if sp > 0 {
			rsp := regexp.MustCompile(fmt.Sprintf(`\s{%d}\S`, sp))
			for rsp.MatchString(line) {
				line, err = r.ReadString('\n')
				if err != nil && err != io.EOF {
					return err
				}
				if err == io.EOF {
					return nil
				}
			}
			sp = -1
		}

		if !rDoneTask.MatchString(line) {
			_, err := io.WriteString(dst, line)
			if err != nil {
				return err
			}
		} else {
			ns := rDoneTask.FindStringIndex(line)
			sp = ns[1]
		}
		if err == io.EOF {
			break
		}
	}
	return nil
}
