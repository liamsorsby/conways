package utils

func Evaluate(dataSet [][]bool) [][]bool {
	var dataset [][]bool

	for dataIndex, b := range dataSet {
		var set []bool

		for index, state := range b {
			var aliveCount int
			aliveCount = checkState(dataSet, dataIndex-1, index-1, aliveCount)
			aliveCount = checkState(dataSet, dataIndex-1, index, aliveCount)
			aliveCount = checkState(dataSet, dataIndex-1, index+1, aliveCount)
			aliveCount = checkState(dataSet, dataIndex, index-1, aliveCount)
			aliveCount = checkState(dataSet, dataIndex, index+1, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index-1, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index, aliveCount)
			aliveCount = checkState(dataSet, dataIndex+1, index+1, aliveCount)

			if state {
				if aliveCount < 2 || aliveCount > 3 {
					set = append(set, false)
				} else {
					set = append(set, true)
				}
			} else {
				if aliveCount == 3 {
					set = append(set, true)
				} else {
					set = append(set, false)
				}
			}
		}

		dataset = append(dataset, set)
	}

	return dataset
}

func checkState(dataset [][]bool, firstIndex int, secondIndex int, aliveCount int) int {
	if firstIndex >= len(dataset) || firstIndex < 0 || secondIndex >= len(dataset[0]) || secondIndex < 0 {
		return aliveCount
	}

	if dataset[firstIndex][secondIndex] {
		return aliveCount + 1
	}

	return aliveCount
}
