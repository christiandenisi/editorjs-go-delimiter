package delimiter

import "github.com/christiandenisi/editorjs-go"

// DelimiterData represents the (empty) data structure for a delimiter block.
type DelimiterData struct{}

// RenderDelimiter is the renderer function for the "delimiter" block.
// It returns a horizontal rule (<hr>) element.
func RenderDelimiter(b editorjs.Block[DelimiterData], ctx *editorjs.Context) (string, error) {
    return "<hr>", nil
}