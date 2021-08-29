package converter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type HeaderConverter struct {
	copyright             Copyright
	targetProjectName     string
	copyrightDefaultValue string
	date                  time.Time
}

func NewHeaderConverter(
	copyright Copyright,
	config *model.Config,
	date time.Time) *HeaderConverter {

	return &HeaderConverter{
		copyright:             copyright,
		targetProjectName:     config.TargetProjectName,
		copyrightDefaultValue: config.Copyright,
		date:                  date,
	}
}

func (header *HeaderConverter) Render(source string, sceneName string) string {

	day := header.date.Day()
	month := int(header.date.Month())
	year := header.date.Year()

	dateStr := fmt.Sprintf("%d/%d/%d", day, month, year)

	copyright, err := header.copyright.Get()

	if err != nil {
		copyright = header.copyrightDefaultValue
	}

	var replacedSource string = source
	replacedSource = strings.ReplaceAll(replacedSource, "__SCENE_NAME__", sceneName)
	replacedSource = strings.ReplaceAll(replacedSource, "__TARGET_PROJECT_NAME__", header.targetProjectName)
	replacedSource = strings.ReplaceAll(replacedSource, "__DATE__", dateStr)
	replacedSource = strings.ReplaceAll(replacedSource, "__YEAR__", strconv.Itoa(year))
	replacedSource = strings.ReplaceAll(replacedSource, "__COPYRIGHT__", copyright)

	return replacedSource
}
