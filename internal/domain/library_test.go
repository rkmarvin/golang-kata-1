package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-kata-1/v2/internal/domain"
)

func TestLibraty_Authros(t *testing.T) {
	// when
	expected := []*domain.Author{
		{
			Email:     "test@m.ru",
			Firstname: "John",
			Lastname:  "Doe",
		},
	}
	l := domain.NewLibrary(expected)

	// then
	authors := l.Authors()

	// given
	assert.Equal(t, authors, expected)
}
