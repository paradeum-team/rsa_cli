package utils

import (
	"bfs_cli_rsa/common/dict"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

func isHex(str string) bool {
	check := regexp.MustCompile(`^[0-9a-f]+$`).MatchString
	return check(str)
}

func IsValidAFID(afid string) bool {
	if len(afid) != dict.AFIdLength {
		return false
	}

	if !isHex(afid) {
		return false
	}

	return true
}

func IsValidGFID(agfid string) bool {
	if len(agfid) != dict.GfIdLength {
		return false
	}

	if !isHex(agfid) {
		return false
	}

	return true
}

func KeyForValue(m map[int]string, v string) int {
	for k, x := range m {
		if x == v {
			return k
		}
	}
	return -1
}


// randStringRunes returns a random string of length n from an array of characters.
func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func GetCurrentDate()(dateLen8 string){
	currentDate :=time.Now().Format(dict.SysTimeFmt4compact)[:8]
	return currentDate
}