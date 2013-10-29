package fasta_test

import (
	. "github.com/lsimoneau/fasta"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fasta", func() {
  Describe("LoadString", func() {

    var (
      fasta Fasta
      err   error
    )

    Context("with single sequence strings", func() {

      BeforeEach(func() {
        singleString := ">Sequence_1\nATCCAGCT\n"
        fasta, err = LoadString(singleString)
        if err != nil {
          Fail(err.Error())
        }
      })

      It("has 1 sequence", func() {
        Expect(len(fasta.Sequences)).To(Equal(1))
      })

      It("has the correct dna sequence", func() {
        Expect(fasta.Sequences[0].Dna).To(Equal("ATCCAGCT"))
      })

    })

    Context("with multiple sequence strings", func() {

      BeforeEach(func() {
        multiString := ">Sequence_1\nATCCAGCT\n>Sequence_2\nGGGCAACT\n>Sequence_3\nATG\nGATCT\n"
        fasta, err = LoadString(multiString)
        if err != nil {
          Fail(err.Error())
        }
      })

      It("has the right number of sequences", func() {
        Expect(len(fasta.Sequences)).To(Equal(3))
      })

      It("has the correct dna sequence", func() {
        Expect(fasta.Sequences[0].Dna).To(Equal("ATCCAGCT"))
        Expect(fasta.Sequences[1].Dna).To(Equal("GGGCAACT"))
        Expect(fasta.Sequences[2].Dna).To(Equal("ATGGATCT"))
      })

    })
  })
})
