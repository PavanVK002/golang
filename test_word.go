package main

import (
	"fmt"
	"unicode"
	"strings"
	"math"
)
func remove_space(str string) string {

	var result []rune
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			if i==len(runes)-1{
				continue
			}
			// Skip consecutive spaces
			for i < len(runes)-1 && runes[i+1] == ' ' {
				i++
			}
			result = append(result, ' ')
		} else {
			result = append(result, runes[i])
		}
	}
	  // Remove trailing space if present
	  if len(result) > 0 && result[len(result)-1] == ' ' {
        result = result[:len(result)-1]
    }
	if result[0] == ' '{
		result =result[1:]
	}
	return string(result)
			
}
func isValid(s string) string {
	freq := make(map[rune]int)
	 for _, value := range s {
		 freq[value]++
	 }
 
	 // Count frequencies of frequencies
	 freqFreq := make(map[int]int)
	 for _, value := range freq {
		 freqFreq[value]++
	 }
 
	 fmt.Println(freq)
	//  fmt.Println(len(freqFreq))
	 // If all characters have the same frequency, return "YES"
	 if len(freqFreq) == 1 {
		 return "YES"
	 }
 
 
	 if len(freqFreq) > 2 {
		 return "NO"
	 }
 
 
	 maxFreq, minFreq := math.MinInt32, math.MaxInt32
	//  fmt.Println(maxFreq, minFreq)
	 for _,freq := range freqFreq {
		 if freq > maxFreq {
			 maxFreq = freq
			//  fmt.Println("x",i, freq)
		 }
		 if freq < minFreq {
			 minFreq = freq
			//  fmt.Println("j",i, freq)
		 }
	 }
 
	 if (freqFreq[maxFreq] == 1 && maxFreq == minFreq+1) || (freqFreq[minFreq] == 1 && minFreq == 1) {
		 return "YES"
	 }
 
	 // Otherwise, return "NO"
	 return "NO"
 
 }
func main(){
	phrase := "  multiple   whitespaces"
	var new_p []rune
    for _,val := range phrase{
	if unicode.IsSpace(val) || unicode.IsPunct(val) && val!='\'' || val=='$' || val=='^'{
		
			val =' '
		
	}
		
		new_p =append(new_p,val)
    }
	

	fmt.Println(string(new_p)) //prints:
	d := remove_space(string(new_p))
	fmt.Println(d) 
	// for _,value := range d{
	// 	fmt.Println(string(value))
	// }
	d1 := strings.Split(d, " ") 
	fmt.Println(d1) 
	for i,k:=range d1{
		fmt.Println(i,k)
	}
	// 	fmt.Println(i,value)
	// fmt.Println(isValid("aabbc"))
}
	
