package converter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Geektree0101/clean-swift-scaffold/internal/model"
)

type HeaderConverter struct {
	org       string
	copyright string
	date      time.Time
}

func NewHeaderConverter(
	config *model.Config,
	date time.Time) *HeaderConverter {

	return &HeaderConverter{
		org:       config.Org,
		copyright: config.Copyright,
		date:      date,
	}
}

func (header *HeaderConverter) Render(source string, sceneName string) string {

	day := header.date.Day()
	month := int(header.date.Month())
	year := header.date.Year()

	dateStr := fmt.Sprintf("%d/%d/%d", day, month, year)

	var replacedSource string = source
	replacedSource = strings.ReplaceAll(replacedSource, "__SCENE_NAME__", sceneName)
	replacedSource = strings.ReplaceAll(replacedSource, "__ORGANIZATION__", header.org)
	replacedSource = strings.ReplaceAll(replacedSource, "__DATE__", dateStr)
	replacedSource = strings.ReplaceAll(replacedSource, "__YEAR__", strconv.Itoa(year))
	replacedSource = strings.ReplaceAll(replacedSource, "__COPYRIGHT__", header.copyright)

	return replacedSource
}
