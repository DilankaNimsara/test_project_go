package main

func printStarPatterns(len int) {
	for i := 0; i <= len; i++ {
		for j := 0; j < i; j++ {
			print(" *")
		}
		println()
	}

	println("\n=============================\n")

	for i := len; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			print(" *")
		}
		println()
	}

	println("\n=============================\n")

}
