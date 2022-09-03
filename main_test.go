package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestIsAwesome_SuccesAllLower(t *testing.T) {
	g := NewGomegaWithT(t)
	actual := IsAwesome("awesome")
	g.Expect(actual).To(BeTrue())
}

func TestIsAwesome_SuccesAllUpper(t *testing.T) {
	g := NewGomegaWithT(t)
	actual := IsAwesome("AWESOME")
	g.Expect(actual).To(BeTrue())
}

func TestIsAwesome_NotAwesome(t *testing.T) {
	g := NewGomegaWithT(t)
	actual := IsAwesome("boring")
	g.Expect(actual).To(BeFalse())
}
