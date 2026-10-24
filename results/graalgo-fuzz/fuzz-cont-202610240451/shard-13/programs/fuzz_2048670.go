package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -1
    var v3 uint8 = 2
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    v6 := []int{}
    v7 := GenS1{F0: 1, F1: 255, F2: 42}
    var v8 int16 = -1
    delete(v1, "hello")
    var v9 int32 = v2
    if (v3 != v3) {
        var v10 int = -9223372036854775808
        _ = len(v6)
        if (!false) {
            if (!(4294967295 == 4294967295)) {
                fmt.Println(v3)
                fmt.Println(v3, v8, v10)
            }
        }
        _ = v10
    } else {
        v11 := GenS1{F0: -9223372036854775808, F1: 741, F2: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(fi12, v8)
        v13 := GenS2{F0: (v9 ^ v9)}
        _ = fi12
        _ = v13
    }
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
