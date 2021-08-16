package converter

import (
	"time"
)

type SourceConverter struct {
	usecases   []string
	sourcePath string
	testPath   string
	date       time.Time
}

func NewSourceConverter(
	usecases []string,
	sourcePath string,
	testPath string,
	date time.Time) *SourceConverter {

	return &SourceConverter{
		usecases:   usecases,
		sourcePath: sourcePath,
		testPath:   testPath,
		date:       date,
	}
}
