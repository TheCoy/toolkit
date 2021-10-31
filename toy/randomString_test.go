package toy

import "testing"

func TestGenRandomString(t *testing.T) {
    t.Log(GenRandomString())
}

func BenchmarkGenRandomString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        b.Log(GenRandomString())
    }
}

func TestHashBySha1(t *testing.T) {
    t.Log(HashBySha1(GenRandomString()))
}

func BenchmarkHashBySha1(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        b.Log(HashBySha1(GenRandomString()))
    }
}
