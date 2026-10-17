package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 1
    var v3 int8 = 1
    var v4 int8 = (v3 / ((v3) | 1))
    var v5 int32 = 42
    v6 := []int{127, int(v5), -551, -1}
    if ((-32768 <= 1000) && false) {
        v5 -= (^(v5))
        var v7 uint8 = 1
        _ = v7
    }
    var v8 uint16 = 32767
    fmt.Println(v5, v2, v8)
    v8 += 907
    var v9 uint32 = v2
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = (false || false)
        _ = fi10
        _ = v11
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    v12 := GenS1{F0: 18446744073709551615, F1: -338}
    var v13 int = 127
    v9 = v2
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
