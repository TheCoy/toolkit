package toy

import (
    "strconv"
    "testing"
)

func TestToSnakeCase(t *testing.T) {
    str := "/HelloWorldHaloFlowerStd"
    t.Log(ToSnakeCase(str))
}

func TestMatchTime(t *testing.T) {
    str := "/data/user/0/com.baidu.homework/cache/a7d7d217170ce1f8fd4ada2536a03b06_1622620135_temp.jar"
    tstr := MatchTime(str)
    if tint, err := strconv.Atoi(tstr); err == nil {
        t.Log(tint)
    }
}