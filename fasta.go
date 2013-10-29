package fasta

import (
  "bytes"
  "bufio"
  "os"
  "strings"
)

type Fasta struct {
  Sequences []Sequence
}

type Sequence struct {
  Name string
  Dna  string
}

func LoadString(input string) (Fasta, error) {
  var sequences []Sequence
  lines, err := readString(input)

  var currentSequence Sequence
  openSequence := false

  var buffer bytes.Buffer

  flushSequence := func() {
    if openSequence {
      currentSequence.Dna = buffer.String()
      buffer.Reset()
      a = append(a, currentSequence)
      openSequence = false
    }
  }

  for _, l := range lines {
    if l[:1] == ">" {
      flushSequence()
      currentSequence = Sequence{
        Name: l[1:len(l)],
      }
      openSequence = true
    } else {
      buffer.WriteString(l)
    }
  }
  flushSequence()

  return Fasta{a}, err
}

func readString(input string) ([]string, error) {
  var lines []string

  reader  := strings.NewReader(input)
  scanner := bufio.NewScanner(reader)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
