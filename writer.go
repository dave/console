package console

import "github.com/gopherjs/gopherjs/js"

func init() {
	w := &Writer{}
	os.Stdout = w
	os.Stderr = w
}

type Writer struct {
	Top bool
	pre *js.Object
}

func (w *Writer) Write(b []byte) (int, error) {
	if w.pre == nil {
		doc := js.Global.Get("document")
		if doc.Call("getElementsByTagName", "pre").Length() > 0 {
			w.pre = doc.Call("getElementsByTagName", "pre").Index(0)
		} else {
			w.pre = doc.Call("createElement", "pre")
			body := doc.Call("getElementsByTagName", "body").Index(0)
			body.Call("appendChild", w.pre)
		}
	}
	if w.Top {
		w.pre.Set("innerHTML", string(b)+w.pre.Get("innerHTML").String())
	} else {
		w.pre.Set("innerHTML", w.pre.Get("innerHTML").String()+string(b))
	}
	return len(b), nil
}
