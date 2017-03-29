// tab package loads tablature files
package tab

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func Load(filepath string) (Tab, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := readLines(f)
	if err != nil {
		return nil, err
	}

	tab := Tab{}

	for i, l := range lines {
		if l[0] == '#' {
			continue
		}

		start, end := strings.Index(l, "["), strings.Index(l, "]")
		if start < 0 || end < 0 {
			return nil, fmt.Errorf("malformed voice on line #%d", i)
		}

		name := strings.TrimSpace(l[:start])
		if len(name) == 0 {
			return nil, fmt.Errorf("missing voice name on line #%d", i)
		}
		if !isValidVoice(name) {
			return nil, fmt.Errorf("unknown voice %q on line #%d", name, i)
		}
		if _, ok := tab[name]; ok {
			return nil, fmt.Errorf("multiple declarations of %q on line #%d", name, i)
		}
		pattern := removeSpaces(l[start+1 : end])
		if len(pattern) == 0 {
			return nil, fmt.Errorf("empty pattern on line #", i)
		}

		tab[name] = []rune{}
		for _, r := range pattern {
			tab[name] = append(tab[name], r)
		}
	}
	return tab, nil
}

func readLines(r io.Reader) ([]string, error) {
	var (
		reader = bufio.NewReader(r)
		lines  = []string{}
		line   []byte
	)
	for {
		segment, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return lines, err
		}
		line = append(line, segment...)
		if !isPrefix {
			lines = append(lines, string(line))
			line = []byte{}
		}
	}
	return lines, nil
}

func removeSpaces(v string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, v)
}
