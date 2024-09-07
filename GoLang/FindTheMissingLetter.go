package main

func FindMissingLetter(chars []rune) rune {

	for i := range chars[:len(chars)-1] {
		if chars[i]+1 != chars[i+1] {
			return chars[i]+1
		}
	}

	return 'a'
}