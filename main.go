package main

import (
	"fmt"
)

/**
 * Assign every lowercase letter a value, 1 for 1 to 26 for z.
 * Given a string of lowercase letters, find the sum of values of the letters in the string.
 */

func lettersum(value string) (sum int) {

	fmt.Println("lettersum(): received 'value' '" + value + "'.")
	for i := 0; i < len(value); i++ {
		character := int(value[i]) - 96
		if character > 0 && character < 27 {
			sum += character
		}
		// sum += mapper[character]
	}

	return
}

func main() {

	fmt.Println(lettersum(""))                         // 0
	fmt.Println(lettersum("a"))                        // 1
	fmt.Println(lettersum("z"))                        // 26
	fmt.Println(lettersum("cabOT"))                    //6
	fmt.Println(lettersum("excellent"))                // 100
	fmt.Println(lettersum("microspectrophotometries")) // 317

}
