package subtitles

import (
	"fmt"
	"io"
	"path/filepath"
	"time"
)

// Block struct
type Block struct {
	lines    []string
	startAt  time.Duration
	finishAt time.Duration
}

// NewBlock returns block instance
func NewBlock() *Block {
	return &Block{}
}

// Subtitle struct
type Subtitle struct {
	blocks []*Block
}

// NewSubtitle returns subtitle instance
func NewSubtitle() *Subtitle {
	return &Subtitle{}
}

// ErrInvalidSubtitle error
type ErrInvalidSubtitle struct {
	format string
	reason string
}

func (e *ErrInvalidSubtitle) Error() string {
	return fmt.Sprintf("Invalid %s subtitle: %s", e.format, e.reason)
}

// ErrUnsupportedExtension error
type ErrUnsupportedExtension struct {
	extension string
}

func (e *ErrUnsupportedExtension) Error() string {
	return fmt.Sprintf("Unsupported extension: %s", e.extension)
}

// Parser interface
type Parser interface {
	Parse(r io.Reader) (*Subtitle, error)
}

// ParserForFile returns parser for subtitle format
func ParserForFile(f string) (p Parser, err error) {
	fileExtension := filepath.Ext(f)

	switch fileExtension {
	case SRTExtension:
		p = NewSRTParser()
	default:
		err = &ErrUnsupportedExtension{
			extension: fileExtension,
		}
	}

	return p, err
}

// Load method
func Load(path string) (s *Subtitle, err error) {
	s = NewSubtitle()

	return s, err
}
