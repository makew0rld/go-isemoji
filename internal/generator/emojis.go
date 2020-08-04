package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/makeworld-the-better-one/go-isemoji"
)

// Some elements adapted from files in
// https://github.com/enescakir/emoji/tree/master/internal/generator
// which are all under an MIT License, like this repo.

const emojiListURL = "https://unicode.org/Public/emoji/" + isemoji.EmojiVersion + "/emoji-test.txt"

// Modified from original to also capture status
var emojiRegex = regexp.MustCompile(`^(?m)(?P<code>[A-Z\d ]+[A-Z\d])\s+;\s+(?P<status>fully-qualified|minimally-qualified|unqualified|component)\s+#\s+.+\s+E\d+\.\d+ (?P<name>.+)$`)

var emojis = make(map[string][2]string) // Codepoints to [name, status]

func addEmoji(line string) {
	matches := emojiRegex.FindStringSubmatch(line)
	if matches == nil || len(matches) < 4 {
		return
	}
	code := matches[1]
	status := matches[2]
	name := matches[3]

	// Transform code into actual codepoints
	unicodes := []string{}
	for _, v := range strings.Split(code, " ") {
		u, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(fmt.Errorf("unknown unicode: %v", v))
		}
		unicodes = append(unicodes, string(u))
	}
	code = strings.Join(unicodes, "")

	emojis[code] = [2]string{name, status}
}

func processLines(url string, fn func(string)) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fn(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %v", err)
	}
	return nil
}

func saveEmojiMap(filename string) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("internal/generator/%s.tmpl", filename))
	if err != nil {
		return err
	}

	d := struct {
		Link     string
		Date     string
		MapLines string
	}{
		Link:     emojiListURL,
		Date:     time.Now().UTC().Format(time.RFC3339),
		MapLines: fmt.Sprintf("%#v", emojis),
	}

	var w bytes.Buffer
	if err = tmpl.Execute(&w, d); err != nil {
		return err
	}

	content, err := format.Source(w.Bytes())
	if err != nil {
		return fmt.Errorf("could not format file: %v", err)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}

func main() {
	err := processLines(emojiListURL, addEmoji)
	if err != nil {
		panic(err)
	}
	err = saveEmojiMap("data.go")
	if err != nil {
		panic(err)
	}
}
