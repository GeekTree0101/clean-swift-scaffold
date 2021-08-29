package converter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type HeaderConverter struct {
	creator               Creator
	targetProjectName     string
	copyrightDefaultValue string
	date                  time.Time
}

func NewHeaderConverter(
	creator Creator,
	config *model.Config,
	date time.Time) *HeaderConverter {

	return &HeaderConverter{
		creator:               creator,
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

	creator, err := header.creator.Get()

	if err != nil {
		creator = "clean-swift-scaffold"
	}

	var replacedSource string = source
	replacedSource = strings.ReplaceAll(replacedSource, "__SCENE_NAME__", sceneName)
	replacedSource = strings.ReplaceAll(replacedSource, "__TARGET_PROJECT_NAME__", header.targetProjectName)
	replacedSource = strings.ReplaceAll(replacedSource, "__DATE__", dateStr)
	replacedSource = strings.ReplaceAll(replacedSource, "__YEAR__", strconv.Itoa(year))
	replacedSource = strings.ReplaceAll(replacedSource, "__COPYRIGHT__", header.copyrightDefaultValue)
	replacedSource = strings.ReplaceAll(replacedSource, "__CREATOR__", creator)

	return replacedSource
}
