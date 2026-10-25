package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }

func main() {
    var v1 int32 = 1
    var v2 bool = (3.14159 == 1e10)
    var v3 int = 42
    v4 := make(map[string]int)
    v3 = v3
    _ = len(v4)
    fmt.Println(v3)
    var v5 int8 = 0
    var v6 int8 = (v5 / (((v5 % ((v5) | 1))) | 1))
    _ = len(v4)
    var v7 int8 = -1
    if ((32767 < 790) && (32768 < v3)) {
        v6 = 127
        v8 := GenS1{F0: 570, F1: v7}
        _ = v8
    }
    v6 = v7
    var v9 uint32 = 0
    if v2 {
        fmt.Println(v3, v3)
        v6 += (^(v5))
        v6 -= v6
    } else {
        v7 = 1
    }
    _ = len(v4)
    _ = len(v4)
    v10 := GenS2{F0: (v3 - v3), F1: v1, F2: 13}
    v11 := GenS2{F0: v3, F1: 2147483647, F2: 18446744073709551615}
    v5 = (v7 / (((^(v6))) | 1))
    fmt.Println(v5, v2, v9)
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := []int{}
        _ = len(v4)
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
