package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
func genFn0(p0 uint32) uint32 {
    p0 = 65535
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint = 18446744073709551615
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 754, F1: 223, F2: 0}
    _ = p0
    _ = v8
    return 65535
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 uint = 611
    var v5 uint16 = 32767
    fmt.Println(v5, v4, v5)
    fmt.Println(v4, v1)
    var v9 uint64 = 778
    fmt.Println(v5)
    v5 -= 710
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint64 = (v9 | 4294967295)
        _ = fi10
        _ = v11
    }
    v9 = (v9 ^ uint64(v4))
    var v12 int16 = int16(v5)
    v13 := GenS2{F0: 235, F1: int16(v9), F2: ("a" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
