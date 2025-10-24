// Empty file
package piscine

func Rot14(s string) string {
	res := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			if int(ch-'a'+14) >= 26 {
				res += string(rune(ch + 14 - 26))
			} else {
				res += string(rune(ch + 14))
			}
		} else if ch >= 'A' && ch <= 'Z' {
			if int(ch-'A'+14) >= 26 {
				res += string(rune(ch + 14 - 26))
			} else {
				res += string(rune(ch + 14))
			}
		} else {
			res += string(ch)
		}
	}
	return res
}
