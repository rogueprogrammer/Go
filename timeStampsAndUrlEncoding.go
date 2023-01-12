// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/url"
	"time"
)

func parseTimestamp(inputTimeStamp string) (time.Time, bool) {
	var allowedTimestampFormats = []string{
		"2006-1-2T15:4:5.999999999Z07:00", // RCF3339Nano with short date fields.
		"2006-1-2t15:4:5.999999999Z07:00", // RFC3339Nano with short date fields and lower-case "t".
		"2006-1-2 15:4:5.999999999",       // space separated with no time zone
		"2006-1-2",                        // date only
	}
	for _, format := range allowedTimestampFormats {
		if t, err := time.Parse(format, inputTimeStamp); err == nil {
			return t, true
		}
	}
	return time.Time{}, false

}

func main() {
	params := url.Values{}
	params.Add("fields", "*")
	fmt.Println("https://www.googleapis.com/drive/v3/files?pageSize=1000&q=mimeType%3D%27application%2Fvnd.google-apps.document%27%20and%20trashed%3Dfalse" + params.Encode())

	input := "2023-01-04T22:10:39.326Z"

	ts, _ := parseTimestamp(input)
	fmt.Println(ts)

}
