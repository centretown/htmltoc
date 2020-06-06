package main

import (
	"testing"
)

func TestLocate(t *testing.T) {
	s := "12345//* hello there */12345/* good bye/**/54321"
	ss := comments[0].locate(s, 1)
	t.Log(ss.start, ss.end)
	t.Log(s[ss.start:ss.end])
	ss = comments[0].locate(s, ss.end)
	t.Log(ss.start, ss.end)
	t.Log(s[ss.start:ss.end])

	s = comments[0].strip(s)
	t.Log(s)

	s = "12345<!-- hello there -->12345\n"
	ss = comments[1].locate(s, 3)
	t.Log(ss.start, ss.end)
	t.Log(s[ss.start:ss.end])
	s = comments[1].strip(s)
	t.Log(s)

}

func TestStrip(t *testing.T) {
	s := "12345<!-- hello there -->12345" + "12345// hello there */12345\n" + "12345<!-- hello there -->12345\n"
	r := s
	for _, del := range comments {
		t.Log(r)
		r = del.strip(r)
		t.Log(r)
	}

	s = comments[0].strip(page)
	t.Log(s)
	s = comments[1].strip(s)
	t.Log(s)
}

func TestStripComments(t *testing.T) {
	s := comments.stripComments(page)
	t.Log(s)
}

func TestReduce(t *testing.T) {
	s := reduce(page)
	t.Log(s)
}
