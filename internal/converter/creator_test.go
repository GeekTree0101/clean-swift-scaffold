package converter_test

import (
	"testing"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
)

// MARK: - Test Double

type CreatorStub struct {
	GetSuccessStub string
	GetErrorStub   error
}

func (c CreatorStub) Get() (string, error) {

	return c.GetSuccessStub, c.GetErrorStub
}

// MARK: - Test Case

func TestCreator(t *testing.T) {

	t.Run("get Creator", func(t *testing.T) {
		// given
		sut := converter.CreatorImpl{}

		// when
		out, err := sut.Get()

		// then
		if err != nil {
			t.Error(err)
		}

		if len(out) == 0 {
			t.Errorf("output should be longer than 0, output: %s", out)
		}
	})
}
