//Guessing game
//Author Jeremy Yon G00330435

package main

import (
    "fmt"
	"math/rand"
	"time"

)

func main(){
    // variable for user input
    var input string

    rand.Seed(time.Now().UTC().UnixNano())

	

    // prompt and get user input
    fmt.Println("Enter String input:")
    fmt.Scanf("%s\n", &input)

    fmt.Println(elizaResponse(input))
 
}

func  elizaResponse(input string) string{
	
	generalResponse := [] string{"I’m not sure what you’re trying to say. Could you explain it to me?",
									"How does that make you feel?",
									"Why do you say that?"	}

	return generalResponse[rand.Intn(3)]
	
}