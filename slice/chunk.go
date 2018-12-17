package slice

func ChunkInt(slice []int, size int) [][]int {
	length := len(slice)
	if size <= 0 {
		size = length
	}

	var datas [][]int
	for i := 0; i < length; i += size {
		endPoint := i + size
		if endPoint > length {
			endPoint = length
		}

		datas = append(datas, slice[i:endPoint])
	}

	return datas
}

func ChunkInt32(slice []int32, size int) [][]int32 {
	length := len(slice)
	if size <= 0 {
		size = length
	}

	var datas [][]int32
	for i := 0; i < length; i += size {
		endPoint := i + size
		if endPoint > length {
			endPoint = length
		}

		datas = append(datas, slice[i:endPoint])
	}

	return datas
}

func ChunkInt64(slice []int64, size int) [][]int64 {
	length := len(slice)
	if size <= 0 {
		size = length
	}

	var datas [][]int64
	for i := 0; i < length; i += size {
		endPoint := i + size
		if endPoint > length {
			endPoint = length
		}

		datas = append(datas, slice[i:endPoint])
	}

	return datas
}

func ChunkStr(slice []string, size int) [][]string {
	length := len(slice)
	if size <= 0 {
		size = length
	}

	var datas [][]string
	for i := 0; i < length; i += size {
		endPoint := i + size
		if endPoint > length {
			endPoint = length
		}

		datas = append(datas, slice[i:endPoint])
	}

	return datas
}