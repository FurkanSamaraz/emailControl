package email

import "regexp"

func CheckEmail(mail string) bool {
	match, _ := regexp.MatchString("[^@]+@[^@]+\\.[^@]+", mail)
	return match
}
