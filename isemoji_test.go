package isemoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO

func TestIsEmoji(t *testing.T) {
	assert.Equal(t, true, IsEmoji("🤗"), "U+1F917 is an emoji")
	assert.Equal(t, true, IsEmoji("👨🏼‍🦰"), "1F468 F3FC 200D F9B0 is an emoji")
	assert.Equal(t, true, IsEmoji("🏻"), "U+1F3FB is a component emoji")
	assert.Equal(t, false, IsEmoji("👱‍♀"), "1F471 200D 2640 is a minimally-qualified emoji")
	assert.Equal(t, false, IsEmoji("test"), "the string \"test\" is not an emoji")
}

func TestIsEmojiNonStrict(t *testing.T) {
	assert.Equal(t, true, IsEmojiNonStrict("🤗"), "U+1F917 is an emoji")
	assert.Equal(t, true, IsEmojiNonStrict("👨🏼‍🦰"), "1F468 F3FC 200D F9B0 is an emoji")
	assert.Equal(t, true, IsEmoji("🏻"), "U+1F3FB is a component emoji")
	assert.Equal(t, true, IsEmojiNonStrict("👱‍♀"), "1F471 200D 2640 is a minimally-qualified emoji")
	assert.Equal(t, true, IsEmojiNonStrict("☠"), "U+2620 is an unqualified emoji")
	assert.Equal(t, false, IsEmojiNonStrict("test"), "the string \"test\" is not an emoji")
}

func TestName(t *testing.T) {
	assert.Equal(t, "man: medium-light skin tone, red hair", Name("👨🏼‍🦰"), "that is the name")
	assert.Equal(t, "", Name("test"), "\"test\" is not an emoji with a name")
}

// TestLatest tests a newly-added emoji to ensure the proper version is being used.
func TestLatest(t *testing.T) {
	assert.Equal(t, true, IsEmoji("😶‍🌫️"), "1F636 200D 1F32B FE0F was added in Emoji 13.1")
	assert.Equal(t, "face in clouds", Name("😶‍🌫️"), "that is the name of an emoji added in 13.1")
}
