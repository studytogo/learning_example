package fasttemplate

import (
	"bufio"
	"fmt"
	"github.com/valyala/fasttemplate"
	"io"
	"strings"
)

func StudyFastTemplate() {

	s := `ewr{{resource}}a{{/resource}}dgfsg {{resource}}ggh{{/resource}}`

	r, err := fasttemplate.NewTemplate(s, "{{resource}}", "{{/resource}}")

	if err != nil {
		fmt.Println("111", err)
		return
	}

	var text string
	r.ExecuteFuncString(func(w io.Writer, tag string) (int, error) { text = tag; return 0, nil })
	fmt.Println(text)
	var values []string
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		values = append(values, line)
	}
	fmt.Println("--", values)
}
