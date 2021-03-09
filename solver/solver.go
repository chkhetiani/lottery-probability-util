package solver

import (
	"github.com/chkhetiani/markov/input"
)

func Solve(data *input.Data) [][] float64 {
	array := initArray(data.Chosen + 2, data.Generated + 1) // height data.Chosen // width data.Generated

	array[0][0] = 1

	for correct := 0; correct <= data.Chosen; correct++ {
		for generated := correct; generated < data.Generated; generated++ {
			data.Total -= generated
			leftCorrect := float64(data.Chosen - correct)
			leftNonCorrect := float64(data.Total) - leftCorrect

			thisProb := array[correct][generated]
			array[correct][generated + 1] += thisProb * (leftNonCorrect / float64(data.Total))
			array[correct + 1][generated + 1] += thisProb * (leftCorrect / float64(data.Total))
			data.Total += generated
		}
	}

	return array
}


func initArray(height int, width int) [][]float64 {
	var result [][]float64
	for i := 0; i < height; i++ {
		var slice []float64
		for j := 0; j < width; j++ {
			slice = append(slice, 0)
		}
		result = append(result, slice)
	}

	return result
}