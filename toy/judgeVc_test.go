package toy

import "testing"

func BenchmarkJudgeVC(b *testing.B) {
    did := "0025e23740c000000f401700300000b7"
    for i :=0;i< b.N;i++{
        b.Log(JudgeVC(did))
    }
}
