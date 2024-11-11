package sorting

func insertion_sort(input []int, len int) {
	for i:=1; i<len; i++ {
		key := input[i]

		j := i - 1
		for j>=0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
}
