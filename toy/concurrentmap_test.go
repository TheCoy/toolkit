package toy

import (
    "strconv"
    "testing"
)

func TestJudgeItem(t *testing.T) {
    for i :=0 ;i < 100;i++ {
        go JudgeItem(strconv.Itoa(i))
    }

    val, ok := LoadItem("101")
    t.Logf("val:%+v\t, ok : %v\n", val, ok)
}