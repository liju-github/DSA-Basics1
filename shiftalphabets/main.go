package main

import "fmt"


func main()  {
	fmt.Println(shiftAlphabet("aBcdefghi!!11jklmno",1))
}

func shiftAlphabet(str string,n int) string  {

	newString := ""

   for _,char := range str{
		if char >= 'a' && char <='z' {
		   if char > 'z'{
			char-=26
		   }
		   char+=rune(n)
		   newString += string(char)
		}else if char >= 'A' && char <='Z'{
			if char > 'Z'{
				char-=26
			   }
			   char+=rune(n)
			   newString += string(char)
		}else{
			newString += string(char)
		}
   }
	
 
	return newString
}