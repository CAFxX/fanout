package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }

func main() {
    var v1 uint64 = 0
    var v2 int = 32767
    v3 := make(map[int]uint32)
    var v4 uint64 = v1
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    _ = len(v5)
    v7 := []int{-897}
    var v8 uint8 = 213
    v9 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: 2}
    v2 -= v2
    var v10 uint = 0
    v11 := GenS1{F0: (v10 << (((v10 / ((v10) | 1))) & 63)), F1: 2147483647, F2: 44}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint64 = v1
        v10 = v10
        _ = fi12
        _ = v13
    }
    var v14 int = 0
    if ((v2 - (^(v14))) > 128) {
        var v15 bool = false
        v4 += v4
        _ = v15
    }
    delete(v6, (^(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
