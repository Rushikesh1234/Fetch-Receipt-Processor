package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(receipt map[string]interface{}) int {
	points := 0

	retailer := receipt["retailer"].(string)
	for _, ch := range retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt["total"].(string), 64)
	if total == math.Trunc(total) {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	items := receipt["items"].([]interface{})
	points += (len(items) / 2) * 5

	for _, item := range items {
		itemMap := item.(map[string]interface{})
		desc := strings.TrimSpace(itemMap["shortDescription"].(string))
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(itemMap["price"].(string), 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt["purchaseDate"].(string))
	if date.Day()%2 != 0 {
		points += 6
	}

	timeStr := receipt["purchaseTime"].(string)
	purchaseTime, _ := time.Parse("15:04", timeStr)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
