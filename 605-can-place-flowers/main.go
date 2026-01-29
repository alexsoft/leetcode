package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		left := i == 0 || flowerbed[i-1] == 0
		right := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if flowerbed[i] == 0 && left && right {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
