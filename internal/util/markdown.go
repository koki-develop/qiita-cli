package util

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

func ReadMarkdown(r io.Reader, frontMatter interface{}) (string, error) {
	scanner := bufio.NewScanner(r)
	var fmLines []string
	var mdLines []string
	var fm bool

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimRightFunc(line, unicode.IsSpace) == "---" {
			fm = !fm
			continue
		}

		if fm {
			fmLines = append(fmLines, line)
		} else {
			mdLines = append(mdLines, line)
		}
	}

	md := strings.TrimSpace(strings.Join(mdLines, "\n"))
	if len(fmLines) == 0 {
		return md, nil
	}

	if err := yaml.Unmarshal([]byte(strings.Join(fmLines, "\n")), frontMatter); err != nil {
		return "", err
	}

	return md, nil
}

func WriteMarkdown(w io.Writer, md string, frontMatter interface{}) error {
	y, err := yaml.Marshal(frontMatter)
	if err != nil {
		return err
	}
	if _, err := w.Write([]byte("---\n")); err != nil {
		return err
	}
	if _, err := w.Write(y); err != nil {
		return err
	}
	if _, err := w.Write([]byte("---\n")); err != nil {
		return err
	}
	if _, err := w.Write([]byte{'\n'}); err != nil {
		return err
	}
	if _, err := w.Write([]byte(md)); err != nil {
		return err
	}

	return nil
}
