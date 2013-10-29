package fasta_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFasta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fasta Suite")
}
