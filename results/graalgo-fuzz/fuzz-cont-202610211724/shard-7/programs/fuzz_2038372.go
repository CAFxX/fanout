package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
func genFn0() int32 {
    var v5 int64 = -9223372036854775808
    if (3.14159 > 1.5) {
        v6 := GenS1{F0: 127, F1: 1000, F2: (-1 < 54)}
        _ = v6
    } else {
        fmt.Println(v5)
    }
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := make(map[uint32]int)
    var v4 uint32 = 89
    var v7 int32 = genFn0()
    v8 := GenS1{F0: 164, F1: -24, F2: (42 >= 124)}
    v2 = (v2 | 105)
    var v9 bool = false
    var v10 uint8 = uint8(v4)
    v7 += 308
    v11 := []int{32767, 9223372036854775807}
    _ = len(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
