package file

import "testing"

func TestPathInfo(t *testing.T) {
	base, name, extension := PathInfo("a/b.c.d")
	if base != "b.c.d" {
		t.Errorf("base %q; want 'b.c.d'", base)
	}
	if name != "b.c" {
		t.Errorf("name %q; want 'b.c'", name)
	}
	if extension != "d" {
		t.Errorf("extension %q; want 'd'", extension)
	}
}
