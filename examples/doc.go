// Package playground registers an HTTP handler at "/compile" that
// proxies requests to the golang.org playground service.
package playground

// Author represents the person who wrote and/or is presenting the document.
type Author struct {
	    Elem []Elem
}

// TextElem returns the first text elements of the author details.
// This is used to display the author' name, job title, and company
// without the contact details.
func (p *Author) TextElem() (elems []Elem) {
