package morestrings

func Encript(s string)string  {
	key :="Roblox"
	r := []rune(s)
	maxLen := len(key)
	
	for i,j := 0,0; i < len(s); i++ {
		
		j = i%maxLen

		r[i] = int32(key[j] ^ s[i])
		
	}
	return string(r)
}