// To create a new style
package main

import (
	"fmt"
	"strings"
)

type decorator struct {
	styles []Style
}

func NewDecorator(styles ...Style) *decorator {
	d := &decorator{}
	for _, style := range styles {
		d.styles = append(d.styles, style)
	}
	return d
}

func (d *decorator) String(str string) string {
	var styleString strings.Builder
	for _, style := range d.styles {
		styleString.WriteString(string(style))
	}
	return fmt.Sprintf("%s%s", styleString.String(), str)
}
