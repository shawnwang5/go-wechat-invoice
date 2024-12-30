package utils

import "strconv"

// 移除 []string 重复的内容
func UniqueStringSlice(arr []string) []string {
	seen := make(map[string]struct{})
	result := []string{}
	for _, value := range arr {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}       // 记录这个元素
			result = append(result, value) // 添加到结果切片
		}
	}
	return result
}

// 移除 []uint 重复的内容
func UniqueUintSlice(arr []uint) []uint {
	seen := make(map[string]uint)
	result := []uint{}
	for _, value := range arr {
		valueStr := strconv.Itoa(int(value))
		if _, exists := seen[valueStr]; !exists {
			seen[valueStr] = value         // 记录这个元素
			result = append(result, value) // 添加到结果切片
		}
	}
	return result
}

// 移除 []int 重复的内容
func UniqueIntSlice(arr []int) []int {
	seen := make(map[string]int)
	result := []int{}
	for _, value := range arr {
		valueStr := strconv.Itoa(value)
		if _, exists := seen[valueStr]; !exists {
			seen[valueStr] = value         // 记录这个元素
			result = append(result, value) // 添加到结果切片
		}
	}
	return result
}
