package converter_test

import (
	"testing"

	"github.com/Geektree0101/clean-swift-scaffold/internal/converter"
)

// MARK: - Test Double

type CopyrightStub struct {
	GetSuccessStub string
	GetErrorStub   error
}

func (c *CopyrightStub) Get() (string, error) {

	return c.GetSuccessStub, c.GetErrorStub
}

// MARK: - Test Case

func TestCopyright(t *testing.T) {

	t.Run("get copyright", func(t *testing.T) {
		// given
		sut := converter.CopyrightImpl{}

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
