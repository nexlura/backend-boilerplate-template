package utilities

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"net"
	"strconv"
	"strings"
	"time"
)

func GenerateUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func GeneratePropertyReferenceNumber() string {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentMonth := currentTime.Month()

	// get the month & year
	month := getFormattedMonth(int(currentMonth))
	year := fmt.Sprintf("%d", currentYear)[2:]

	// get the random 5 digits
	randomValues := fmt.Sprint(time.Now().Nanosecond())[:5]

	// concat the strings
	generatedNumber := fmt.Sprintf("%s%s%s", year, month, randomValues)

	return generatedNumber
}

func getFormattedMonth(month int) string {
	switch month {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	default:
		return strconv.Itoa(month)
	}
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateTemporaryPassword(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func GetCurrentTime(value string) string {
	const layout = "Mon Jan 2 15:04:05 UTC 2006"
	t, _ := time.Parse(layout, value)
	return t.Format(time.RFC850)
}

func GetMediaIdFromLink(link string, separator string) string {
	// split the avatar to get the key for minio
	splitLink := strings.Split(link, separator)[1]

	// remove the '/' character from the start of the string
	return strings.Replace(splitLink, "/", "", -1)
}

func GetArrayFromCommaStrings(value string) []string {
	arr := strings.FieldsFunc(value, func(r rune) bool {
		return r == ','
	})

	return arr
}

func DeleteStringElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// MOVING SLICE ELEMENTS TO A NEW POSITION START
func insertInt[T any](array []T, value T, index int) []T {
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func removeInt[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func MoveElement[T any](array []T, srcIndex int, dstIndex int) []T {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

// END MOVING SLICE ELEMENTS TO A NEW POSITION
