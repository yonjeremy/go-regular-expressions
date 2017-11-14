//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
	"math/rand"
	"time"
	// "regexp"

)

func main(){
    // variable for user input
    var input string

    rand.Seed(time.Now().UTC().UnixNano())


	

    // prompt and get user input
    // fmt.Println("Enter String input:")
    // fmt.Scanf("%s\n", &input)
	input = "People say I look like both my mother and father."
    fmt.Printf("%s\n%s\n",input,elizaResponse(input))
	input = "Father was a teacher."
    fmt.Printf("%s\n%s\n",input,elizaResponse(input))
	input = "I was my father’s favourite."
    fmt.Printf("%s\n%s\n",input,elizaResponse(input))
	input = "I’m looking forward to the weekend."
    fmt.Printf("%s\n%s\n",input,elizaResponse(input))
	input = "My grandfather was French!"
    fmt.Printf("%s\n%s\n",input,elizaResponse(input))
 
}

func  elizaResponse(input string) string{

	// matchedFather, err := regexp.MatchString("\\bfather\\b", input)
	// fmt.Println(matchedFather, err)
	
	generalResponse := [] string{"I’m not sure what you’re trying to say. Could you explain it to me?",
									"How does that make you feel?",
									"Why do you say that?"	}

	return generalResponse[rand.Intn(3)]
	
}