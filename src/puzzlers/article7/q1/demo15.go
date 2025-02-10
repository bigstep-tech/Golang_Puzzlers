package main

import "fmt"

/*
总结 from 通义千问，https://tongyi.aliyun.com/
容量是slice的一个重要属性，它决定了在不进行内存重新分配的情况下，slice最多能容纳多少元素。
合理设置slice的初始容量可以优化性能，减少不必要的内存分配操作。
使用make([]T, length, capacity)函数可以显式地指定slice的长度和容量，其中第三个参数是可选的，如果不提供，默认容量等于长度。
*/
func main() {
	// 示例1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	// 示例2。
	/*
		计算新Slice s4 的长度和容量
		长度（Length）：
		新slice s4 包含了原slice s3 中从索引3到索引5的元素，因此它的长度是3。
		所以，len(s4) 是 3。
		容量（Capacity）：
		新slice s4 的容量是从其起始索引（这里是3）到原slice s3 的末尾的元素数量。
		因为原slice s3 的总长度是8，所以 s4 的容量是从索引3到索引7（不包括8），共5个元素。
		所以，cap(s4) 是 5。
	*/
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()

	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)
}
