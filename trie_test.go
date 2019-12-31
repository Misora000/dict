package dict

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TrieTestSuite struct {
	suite.Suite
	d *Dict
}

func (s *TrieTestSuite) SetupSuite() {
	s.d = NewDictionary()
	str := []string{
		"popo",
		"popo scam",
		"popo fuck",
		"portto",
		"pig",
		"cob",
	}
	for _, v := range str {
		s.d.Add(v)
	}
	s.d.Dump()
}

func (s *TrieTestSuite) TestTrie() {
	// Found, return is the exact node.
	n := s.d.Find("popo")
	s.Require().NotNil(n)
	s.Require().Equal(n.Word, "popo")

	n = s.d.Find("popo scam")
	s.Require().NotNil(n)
	s.Require().Equal(n.Word, "popo scam")

	// Not found, return is nil.
	n = s.d.Find("popo scam fuck")
	s.Require().Nil(n)

	n = s.d.Find("popo shit")
	s.Require().Nil(n)

	n = s.d.Find("po")
	s.Require().Nil(n)

	// Predict
	o := s.d.Predict("po")
	s.Require().Len(o, 4)
	s.Require().Equal(
		len(strings.Join([]string{"popo", "popo scam", "popo fuck", "portto"}, "")),
		len(strings.Join(o, "")))

	o = s.d.Predict("p")
	s.Require().Len(o, 5)
	s.Require().Equal(
		len(strings.Join([]string{"pig", "popo", "popo scam", "popo fuck", "portto"}, "")),
		len(strings.Join(o, "")))

	o = s.d.Predict("k")
	s.Require().Len(o, 0)
	s.Require().Equal(
		len(strings.Join([]string{}, "")),
		len(strings.Join(o, "")))

	o = s.d.Predict("popo ")
	s.Require().Len(o, 2)
	s.Require().Equal(
		len(strings.Join([]string{"popo scam", "popo fuck"}, "")),
		len(strings.Join(o, "")))
}

func TestTrie(t *testing.T) {
	suite.Run(t, new(TrieTestSuite))
}
