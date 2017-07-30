package dutch

// a partitioningFunction should return -1 0 or 1 if the parameter belongs
// to the first, second or third partition respectively.
type partitioningFunction func(int) int

// swap the content of the i-th cell and the j-th in the a array.
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// dutchFlag implements the simple dutch flag problem:
// an in place three-way partition of an array.
// param a: []int, the array to partition
// param getPartition: partitioningFunction, a function int -> int that takes
// 											in input an element and return the partition id
func DutchFlag(a []int, getPartition partitioningFunction) {
	// i = f + 1 where f is the index of the last item of partition #1, default f = -1
	// j = s + 1 where s is the index of the last item of partition #2, default s = -1
	// last = t - 1 where t is the first element of the partition #3, default t = len(a)
	i, j := 0, 0
	last := len(a) - 1

	for j <= last {
		// current evaluating index is j
		switch getPartition(a[j]) {
		case -1:
			// first partition
			swap(a, i, j)
			i++
			j++
		case 0:
			// second partition
			j++
		case 1:
			// third partition
			swap(a, j, last)
			last--
		}
	}
}

// getPartitionInt returns a value in {-1, 0, 1} representing
// the partition the item (int) belongs to in a three class partition.
// Current implementation: 1 returns -1, 2 returns 0; others returns 1;
func BasicOneTwoThreeGetPartition(item int) int {
	if item == 1 {
		return -1
	} else if item == 2 {
		return 0
	} else {
		return 1
	}
}
