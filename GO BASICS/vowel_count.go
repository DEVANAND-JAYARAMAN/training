package main
import (
    "fmt"
    "strings"
)
func main(){
    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)

    vowel_count := 0
    consonant_count := 0
    vowels:= "aeiouAEIOU"
    for _,val := range input{
        if(!strings.ContainsRune(vowels,val)){
            consonant_count++
        }else{
            vowel_count++;
        }
    }
    fmt.Printf("The number of vowels in given string %s is: %d\n",input,vowel_count)
    fmt.Printf("The number of consonants in given string %s is %d\n",input,consonant_count)
}