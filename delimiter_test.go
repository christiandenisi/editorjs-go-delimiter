package delimiter

import (
    "testing"
    "github.com/christiandenisi/editorjs-go"
)

func TestRenderDelimiter(t *testing.T) {
    block := editorjs.Block[DelimiterData]{}
    ctx := &editorjs.Context{}

    out, err := RenderDelimiter(block, ctx)
    if err != nil {
        t.Fatalf("RenderDelimiter returned an error: %v", err)
    }

    expected := "<hr>"
    if out != expected {
        t.Errorf("unexpected output: got %q, want %q", out, expected)
    }
}