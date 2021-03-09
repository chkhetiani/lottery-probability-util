package output

import "fmt"

func Print(array [][]float64) {
	fmt.Print("----")
	for range array[0] {
		fmt.Print("-------")
	}
	fmt.Println()
	for correctCount, correct := range array {
		fmt.Print(fmt.Sprintf("%d | ",correctCount))
		for _, generated := range correct {
			fmt.Print(fmt.Sprintf("%.2f | ", generated))
		}
		fmt.Println()
		fmt.Print("----")
		for range correct {
			fmt.Print("-------")
		}
		fmt.Println()
	}


	for correctCount, correct := range array {
		fmt.Println(fmt.Sprintf("probability of %d correct is: %.6f%% (one in %.2f)",
			correctCount, correct[len(correct) - 1] * 100, float64(1) / correct[len(correct) - 1] ))
	}
}
