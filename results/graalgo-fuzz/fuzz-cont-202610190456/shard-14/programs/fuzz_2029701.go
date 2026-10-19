package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 0
    v3 := make(map[int]bool)
    var v4 uint64 = 2
    v5 := make(map[int]uint32)
    v2 = v2
    var v6 int32 = v2
    fmt.Println(v4, v6, v6)
    v4 = 0
    fmt.Println(v2, v6, v2)
    v4 = v4
    v7 := GenS1{F0: 2, F1: 65535, F2: (v6 ^ (-(v6)))}
    fmt.Println(v2, v6)
    v8 := []int{9223372036854775807, 32768, -9223372036854775808}
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: 1, F1: v2, F2: 1}
        var v11 int = (fi9 | (-(fi9)))
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if (false || true) {
            var v13 int = 32768
            _ = v13
        }
        fmt.Println(fi12, v2, v2)
        v6 = (v6 >> ((901) & 63))
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
