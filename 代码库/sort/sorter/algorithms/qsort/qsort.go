package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		// j指向从右往左第一个小于temp的元素
		// 或者j = p - 1
		for j >= p && values[j] >= temp {
			j--
		}
		// 将小于元素赋值给p指向的temp元素，将p指向小于元素本来位置
		if j >= p {
			values[p] = values[j]
			p = j
		}
		// 从左往右，找到第一个大于temp的元素
		// 或者i > p
		if values[i] <= temp && i <= p {
			i++
		}
		// 将大于元素赋值给p指向的temp元素，将p指向大于元素本来位置
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	// 将p指向元素赋值为temp
	values[p] = temp
	// 如果左半部分还没有排好序
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	// 如果右半部分还没有排好序
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
