package tab

import (
	"testing"

	"github.com/brettbuddin/sm808/test/assert"
)

func TestLoad(t *testing.T) {
	tab, err := Load("testdata/good.txt")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(tab), 3)
	assert.Equal(t, tab["bassdrum"], []rune{'x', 'x', '_'})
	assert.Equal(t, tab["snaredrum"], []rune{'_', '_', '_', 'x', 'X', '_'})
	assert.Equal(t, tab["closedhihat"], []rune{'x', '_', '_', '_'})
}

func TestLoadWeirdFormatting(t *testing.T) {
	tab, err := Load("testdata/weird.txt")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(tab), 3)
	assert.Equal(t, tab["bassdrum"], []rune{'x', 'x', '_'})
	assert.Equal(t, tab["snaredrum"], []rune{'_', '_', '_', 'x', 'X', '_'})
	assert.Equal(t, tab["closedhihat"], []rune{'x', '_', '_', '_'})
}

func TestLoadBadVoiceName(t *testing.T) {
	_, err := Load("testdata/badvoice.txt")
	assert.NotEqual(t, err, nil)
}

func TestLoadBadSequence(t *testing.T) {
	_, err := Load("testdata/badsequence.txt")
	assert.NotEqual(t, err, nil)
}

func TestLoadDuplicateVoice(t *testing.T) {
	_, err := Load("testdata/duplicatevoice.txt")
	assert.NotEqual(t, err, nil)
}

func TestLoadMissingVoiceName(t *testing.T) {
	_, err := Load("testdata/missingvoicename.txt")
	assert.NotEqual(t, err, nil)
}
