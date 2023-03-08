/** * Created with VS Code. * User: Dimas Anom Priyayi * Date: 07/03/23 * Time: 19:23  * ID Reg : 1955617840-940 * Challanges 2 **/
package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	var j int = 0
	for j < 5 {
		fmt.Println("Nilai j = ", j)
		j++
	}

	for pos, char := range "САШАРВО" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	for {
		if j++; j <= (10) {
			fmt.Println("Nilai j = ", j)
		} else {
			break
		}
	}
}
