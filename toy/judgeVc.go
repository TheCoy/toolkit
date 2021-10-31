package toy

import "strconv"

func JudgeVC(did string) bool {
    if len(did) != 32 {
        return false
    }

    src := did[22:25]
    srcInt,  err := strconv.ParseUint(src, 16, 64)
    if err != nil {
        return false
    }
    //0x7FE
    targetInt := uint64(2046)

    return (srcInt & targetInt) == 2
}
