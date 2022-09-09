package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMyFunction_0(t *testing.T) {
	g := NewGomegaWithT(t)
	actual := MyFunction("awesome")
	g.Expect(actual).To(Equal("awesome"))
}
