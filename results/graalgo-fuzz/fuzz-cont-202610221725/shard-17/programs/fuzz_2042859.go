package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    var v3 uint = 301
    var v4 uint16 = 2
    var v5 int8 = 1
    fmt.Println(v5, v5, v5)
    if (((!true) || true) && (float64((3.14159 / 0.1)) > 0.1)) {
        v3 = (v3 & v3)
        v6 := []int{}
        _ = v6
    }
    var v7 uint32 = v1
    v4 += (v4 | v4)
    v8 := GenS1{F0: 4294967295, F1: "a", F2: (v1 + v1)}
    if (!((0 > 1) || (42 != -596))) {
        v9 := []int{}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        v10 := GenS1{F0: 42, F1: "\n\t", F2: 855}
        _ = v10
        _ = v9
    }
    var v11 uint8 = uint8(v7)
    var v12 uint = uint(v11)
    v1 = (v7 % ((v1) | 1))
    v11 = (v11 & 2)
    var v13 float64 = 1e10
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
