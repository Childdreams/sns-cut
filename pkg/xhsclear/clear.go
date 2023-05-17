package xhsclear

import "regexp"

var emojiRe = regexp.MustCompile(`\[.{1,10}(R|r)\]`)

// EmojiFilter 表情过滤器
func EmojiFilter(content []byte) []byte {
	return emojiRe.ReplaceAll(content, []byte{})
}

var tagRe = regexp.MustCompile(`\#{1,2}[^\#|\s|\pP]{1,10}`)

// TagFilter 标签过滤器
func TagFilter(content []byte) []byte {
	return tagRe.ReplaceAll(content, []byte{})
}

//CleanAll 将所有小红书特有标签清除
func CleanAll(content string) string {
	return CleanBySpecifyFilter(content, EmojiFilter, TagFilter, AtFilter)
}

// CleanBySpecifyFilter 通过制定的过滤器来过滤内容
func CleanBySpecifyFilter(content string, filters ...func(content []byte) []byte) string {
	contentBytes := []byte(content)
	for _, f := range filters {
		contentBytes = f(contentBytes)
	}
	return string(contentBytes)
}

var atRe = regexp.MustCompile(`\@.[^\@|\s|\pP]{2,30}`)

// AtFilter @
func AtFilter(content []byte) []byte {
	return atRe.ReplaceAll(content, []byte{})
}
