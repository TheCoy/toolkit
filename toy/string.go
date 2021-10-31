package toy

func SwapString(src string, num int) string {
    dst := []byte(src)
    length := len(dst) - 1

    for index := range dst {
        if index >= num  {
            break
        }
        dst[index], dst[length-index] = dst[length-index], dst[index]
    }
    return string(dst)
}
