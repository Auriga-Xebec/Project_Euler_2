package main

func main() {
	fibo_1 := 0
	fibo_2 := 1
	current_it := 0
	sum := 0

	for i := 1; i < 34; i++ {
		current_it = fibo_1 + fibo_2
		fibo_1 = fibo_2
		fibo_2 = current_it
		//fmt.Println(fibo_2)
		if fibo_2 < 4000000 {
			if fibo_2%2 == 0 {
				sum += fibo_2
			} else {
				continue
			}
		} else {
			println(sum)
			break
		}
	}
}
