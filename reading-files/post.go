package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return scanner.Text()[len(tagName):]
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
	}, nil
}
