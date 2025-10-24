package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// confirmation
	if !ConfirmBase(baseFrom, baseTo) {
		return ""
	}

	decimalNumber := converTodicemal(nbr, baseFrom)
	return convertToNewBase(decimalNumber, baseTo)
}

// this gunc convert the base to
func converTodicemal(number string, base string) int {
	res := 0

	for _, ch := range number {

		digitValue := -1
		for i, bch := range base {
			if ch == bch {
				digitValue = i
				break
			}
		}

		if digitValue == -1 { // character not found in base
			return 0
		}
		res = res*len(base) + digitValue
	}

	return res
}

// cinfirme if bases are true
func ConfirmBase(base1 string, base2 string) bool {
	if len(base1) < 2 || len(base2) < 2 {
		return false
	}
	// confirme base1
	for i := 0; i <= len(base1)-2; i++ {
		for j := i + 1; j <= len(base1)-1; j++ {
			if base1[i] == base1[j] {
				return false
			}
		}
	}
	for i := 0; i <= len(base2)-1; i++ {
		for j := i + 1; j <= len(base2)-2; j++ {
			if base2[i] == base2[j] {
				return false
			}
		}
	}
	return true
	// confirme base2
}

func convertToNewBase(nbr int, base string) string {
	res := ""
	negative := false
	if nbr < 0 {
		nbr = -nbr
		negative = true

	}
	for nbr > 0 {
		res = string(base[nbr%len(base)]) + res
		nbr = nbr / len(base)

	}
	if negative {
		res = "-" + res
	}
	return res
}
