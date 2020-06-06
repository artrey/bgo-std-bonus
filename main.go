package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := purchase * percent / 100 / 100
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
