package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для генерации случайного массива чисел
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) // Генерация чисел от 0 до 99
	}
	return arr
}

// Функция для сортировки массива методом пузырька
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Функция для поиска числа в массиве
func searchNumber(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1 // Если число не найдено
}

func Rnd() {
	size := 10
	arr := generateRandomArray(size)

	fmt.Println("Исходный массив:", arr)

	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)

	target := arr[rand.Intn(size)] // Выбираем случайное число из массива для поиска
	index := searchNumber(arr, target)

	if index != -1 {
		fmt.Printf("Число %d найдено по индексу %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено\n", target)
	}
}
