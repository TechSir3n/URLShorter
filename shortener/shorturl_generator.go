package shortener

import (
	"encoding/base64"
)


func GenerateShortLink(urlLong, userID string) string {
	url := []byte(urlLong + userID)
	str := base64.StdEncoding.EncodeToString(url)
	return str[:8] // up to 8 characters maximum
}
