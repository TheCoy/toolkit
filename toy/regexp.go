package toy

import (
    "fmt"
    "regexp"
    "strings"
)

var (
    matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
    matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
    matchTime     = regexp.MustCompile("(.+)_([0-9]+)_temp.jar")
)

// ToSnakeCase method change string to snakecase
func ToSnakeCase(str string) string {
    fmt.Println(str)
    snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
    fmt.Println(snake)
    snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    fmt.Println(snake)
    return strings.ToLower(snake)
}


func MatchTime(str string) string {
    res := matchTime.ReplaceAllString(str, "${2}")

    return res
}