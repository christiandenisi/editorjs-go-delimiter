# editorjs-go-delimiter

`editorjs-go-delimiter` is a plugin for [`editorjs-go`](https://github.com/christiandenisi/editorjs-go) that renders Editor.js `delimiter` blocks as horizontal rules.

## ✨ Features

- ✅ Outputs a semantic `<hr>` element
- ✅ Simple, zero-config renderer
- 🧩 Plug-and-play with `editorjs-go`

---

## 📦 Installation

```bash
go get github.com/christiandenisi/editorjs-go
go get github.com/christiandenisi/editorjs-go-delimiter
```

---

## 🚀 Usage

```go
package main

import (
    "fmt"
    "github.com/christiandenisi/editorjs-go"
    "github.com/christiandenisi/editorjs-go-delimiter"
)

func main() {
    jsonData := []byte(`{
        "blocks": [{
            "type": "delimiter",
            "data": {}
        }]
    }`)

    conv := editorjs.New()
    editorjs.Register(conv, "delimiter", delimiter.RenderDelimiter)

    html, err := conv.Convert(jsonData)
    if err != nil {
        panic(err)
    }

    fmt.Println(html)
}
```

---

## 📌 Notes

- This block has no fields and always renders `<hr>`

---

## 🧱 Compatibility

- Compatible with `editorjs-go` version 1.x

---

## 👤 License

MIT