package main

func cmmdc(a string, b string) int{
	x1 := len(a)
	x2 := len(b)

	for x1 != x2{
		if x1 > x2{
			x1 -= x2
		}else{
			x2 -= x1
		}
	}
	return x1
}

func gcdOfStrings(str1 string, str2 string) string {
	cmmdc := cmmdc(str1, str2)
	for i := 0; i < len(str1); i+= cmmdc {
		if str1[0:cmmdc] != str1[i:i+cmmdc] {
			return ""
		}
	}
	for i := 0; i < len(str2); i+= cmmdc {
		if str2[0:cmmdc] != str2[i:i+cmmdc] {
			return ""
		}
	}
	if(str1[0:cmmdc] != str2[0:cmmdc]){
		return ""
	}
	return str1[:cmmdc]
}
