package main

// Написать программу, которая конкурентно рассчитает значение 
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
)

func Square(num int, ch chan int) {
	ch <- num * num
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	ch := make(chan int, len(nums))
	defer close(ch)

	for i := range nums {
		go Square(nums[i], ch)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}

/************* 2 вариант *****************************************
import (
	"fmt"
	"time"
)

func Square(num int) {
	fmt.Printf("Квадрат числа %d равен %d\n", num, num * num)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	for i := range nums {
		go Square(nums[i])
	}
	time.Sleep(time.Millisecond)
}
*/

/************* 3 вариант - пакет "sync" ***************************
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Квадрат числа %d равен %d\n", num, num * num)
		}(num)
	}

	wg.Wait()
}
*/
