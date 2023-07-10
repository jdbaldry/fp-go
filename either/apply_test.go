package either

import (
	"testing"

	M "github.com/ibm/fp-go/monoid/testing"
	N "github.com/ibm/fp-go/number"
	"github.com/stretchr/testify/assert"
)

func TestApplySemigroup(t *testing.T) {

	sg := ApplySemigroup[string](N.SemigroupSum[int]())

	la := Left[string, int]("a")
	lb := Left[string, int]("b")
	r1 := Right[string](1)
	r2 := Right[string](2)
	r3 := Right[string](3)

	assert.Equal(t, la, sg.Concat(la, lb))
	assert.Equal(t, lb, sg.Concat(r1, lb))
	assert.Equal(t, la, sg.Concat(la, r2))
	assert.Equal(t, lb, sg.Concat(r1, lb))
	assert.Equal(t, r3, sg.Concat(r1, r2))
}

func TestApplicativeMonoid(t *testing.T) {

	m := ApplicativeMonoid[string](N.MonoidSum[int]())

	la := Left[string, int]("a")
	lb := Left[string, int]("b")
	r1 := Right[string](1)
	r2 := Right[string](2)
	r3 := Right[string](3)

	assert.Equal(t, la, m.Concat(la, m.Empty()))
	assert.Equal(t, lb, m.Concat(m.Empty(), lb))
	assert.Equal(t, r1, m.Concat(r1, m.Empty()))
	assert.Equal(t, r2, m.Concat(m.Empty(), r2))
	assert.Equal(t, r3, m.Concat(r1, r2))
}

func TestApplicativeMonoidLaws(t *testing.T) {
	m := ApplicativeMonoid[string](N.MonoidSum[int]())
	M.AssertLaws(t, m)([]Either[string, int]{Left[string, int]("a"), Right[string](1)})
}