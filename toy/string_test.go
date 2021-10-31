package toy

import "testing"

func Test_SwapString(t *testing.T)  {
    src := "abcdefgh"
    t.Log(src)
    t.Log(SwapString(src, 4))
}
