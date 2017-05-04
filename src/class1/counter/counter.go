package counter

func LenOfString(str string) int {

	if (str=="") {
		return 0
	} else {
		return len(str) 
	}
}

func CountEven (str string) bool{
	if LenOfString(str) % 2 == 0 {
		return true 
	} else {
		return false 
	}
}


