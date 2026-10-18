package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }

func main() {
    var v1 int = 128
    var v2 int16 = 42
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 uint64 = 0
    v6 := []int{v1, v1, (-(v1)), 1}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint32 = 664
        v8 = 499
        _ = len(v3)
        _ = fi7
        _ = v8
    }
    v5 = 753
    _ = len(v3)
    var v9 uint8 = 1
    fmt.Println(v5)
    v10 := GenS1{F0: 1, F1: 9223372036854775807}
    var v11 uint64 = 18446744073709551615
    switch (v1 & int(v9)) {
    case 0:
        fmt.Println(v5, v5, v9)
    default:
        v12 := GenS1{F0: 118, F1: -373}
        _ = v12
    }
    if (9223372036854775807 == 42) {
        v1 = (v1 | (v1 ^ 32768))
    }
    _ = len(v4)
    fmt.Println(v11, v5)
    v13 := GenS2{F0: "", F1: -208, F2: 65535}
    v2 = -1
    v11 = 42
    var v14 uint8 = 0
    v5 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
