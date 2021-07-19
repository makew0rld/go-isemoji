// Package isemoji helps you determine whether a string is an emoji.
package isemoji

const EmojiVersion = "14.0"

// IsEmoji returns true if the provided string contains one fully-qualified emoji.
// Component emojis are also valid.
//
// See the follow document for definitions of fully-qualified and component:
//	https://www.unicode.org/reports/tr51/#def_qualified_emoji_character
func IsEmoji(s string) bool {
	for k, v := range emojis {
		if s == k && (v[1] == "fully-qualified" || v[1] == "component") {
			return true
		}
	}
	return false
}

// IsEmojiNonStrict returns true if the provided string contains one emoji.
// fully-qualified, minimally-qualified, unqualified, and component emojis all return true.
//
// See the follow document for definitions of those terms:
//	https://www.unicode.org/reports/tr51/#def_qualified_emoji_character
func IsEmojiNonStrict(s string) bool {
	for k := range emojis {
		if s == k {
			return true
		}
	}
	return false
}

// Name returns the Unicode name for the emoji.
// It returns an empty string if the provided string is not an emoji.
func Name(s string) string {
	return emojis[s][0]
}
