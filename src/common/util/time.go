package util

import "time"

// TStr2Unix convert time string to timestamp.
// Time string must comply with time.DateTime standard,
// otherwise the returned timestamp is 0.
func TStr2Unix(source string) int64 {
	parse, err := time.ParseInLocation(time.DateTime, source, time.Local)
	if err != nil {
		return 0
	}

	return parse.Unix()
}

func Unix2Time(source int64) string {

	return time.Unix(source, 0).Format(time.DateTime)
}

func ConvertTimeStr(source time.Time) string {

	return source.Format(time.DateTime)
}
