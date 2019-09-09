package bar

import "github.com/andig/foo"

type Bar struct {
}

func (b *Bar) Bar() {
	f := new(foo.Foo)
	f.V1()
}
