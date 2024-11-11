package sorting

func selection_sort(input []int, len int){
	for i := 0; i < len-1; i++ {
		smallest := i
		for j := i + 1; j < len; j++ {
			if input[j] < input[smallest] {
				smallest = j
			}
		}

		if smallest != i {
			tmp := input[i]
			input[i] = input[smallest]
			input[smallest] = tmp
		}
	}
}
