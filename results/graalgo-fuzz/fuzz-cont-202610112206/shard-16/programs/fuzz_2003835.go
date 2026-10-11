package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() int16 {
    v4 := GenS1{F0: 1, F1: 2, F2: 626}
    v5 := GenS1{F0: 890, F1: 42, F2: 890}
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{64}
    v2 := make(map[int]int)
    var v3 uint64 = 1
    v3 = 4294967295
    var v6 uint64 = v3
    fmt.Println(v6, v3)
    var v7 uint = 18446744073709551615
    v3 = 1
    if len(v1) > 0 {
        v1[0] = 2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 int32 = -2147483648
    v6 += (v6 << (((v7 & v7)) & 63))
    var v9 string = "\n\t"
    var v10 int64 = -1
    for fi11 := 0; fi11 < 2; fi11++ {
        _ = len(v1)
        v8 = (v8 | v8)
        _ = fi11
    }
    v12 := GenS2{F0: (v10 & v10), F1: false}
    v13 := GenS1{F0: 545, F1: (v7 & v7), F2: (v3 | v3)}
    v2[127] = 32767
    v2[4294967295] = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!true) {
        var v14 float32 = (((0.0 * 1e10) * -2.25) + 1.5)
        v10 += v10
        v14 = (v14 + -2.25)
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
