package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int8) int {
    p0 = (p0 | p0)
    p0 -= (p0 * (p0 * 92))
    var v6 float32 = 1e10
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    var v4 int64 = -1
    v5 := make(map[string]string)
    var v7 int = 387
    _ = len(v5)
    v8 := GenS1{F0: 427, F1: -325}
    var v9 uint16 = 32767
    v10 := v5[("foo bar" + ("" + ("\n\t" + "a")))]
    switch ((v10 + v10) + "hello") {
    case "\n\t":
        var v11 int64 = v4
        _ = v11
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int32 = 42
            var v14 int32 = -875
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := v5[v10]
        v4 = v4
        v1 = append(v1, v16)
        _ = fi15
        _ = v16
    }
    v1 = append(v1, ("" + v10))
    var v17 uint32 = 4294967295
    for fi18 := 0; fi18 < 14; fi18++ {
        v1 = append(v1, v10)
        var v19 int32 = -2147483648
        _ = fi18
        _ = v19
    }
    v7 = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
