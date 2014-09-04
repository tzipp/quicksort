package main

import "fmt"

func main() {
    list_one := []float64{6, 2, 32, 4, 1, 5, 8, 31}
    fmt.Println("Slice to sort:", list_one)
    fmt.Println(quicksort(list_one))


}

func quicksort(numbers []float64) []float64 {
    if len(numbers) == 0 || len(numbers) == 1 {
        return numbers
    }

    if len(numbers) == 2 {
        if numbers[0] > numbers[1] {
            numbers[0], numbers[1] = numbers[1], numbers[0]
        }
        return numbers
    }

    pivot := numbers[0]
    s_nums := make([]float64, 0, len(numbers))
    g_nums := make([]float64, 0, len(numbers))

    for _, num := range(numbers) {
        if num < pivot {
            s_nums = append(s_nums, num)
        }

        if num > pivot {
            g_nums = append(g_nums, num)
        }

    }
    
    s_nums = append(s_nums, pivot)
    
    
    return append(quicksort(s_nums), quicksort(g_nums)...)
}


