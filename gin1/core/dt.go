package core

import (
	"errors"
	"time"

	"github.com/qzakwani/go/gin1/settings"
)

const frmt = time.RFC3339

func DtNow() string {
	if settings.UseTZ {
		loc, _ := time.LoadLocation(settings.TimeZone)
		return time.Now().In(loc).Format(frmt)
	} else {
		return time.Now().UTC().Format(frmt)

	}
}
func DtComp(dt1 string, comp rune, dt2 string) (bool, error) {
	_dt1, _ := time.Parse(frmt, dt1)
	_dt2, _ := time.Parse(frmt, dt2)
	switch comp {
	case '>':
		return _dt1.After(_dt2), nil
	case '<':
		return _dt1.Before(_dt2), nil
	case '=':
		return _dt1.Equal(_dt2), nil
	default:
		return false, errors.New("unsupported operation")
	}
}
