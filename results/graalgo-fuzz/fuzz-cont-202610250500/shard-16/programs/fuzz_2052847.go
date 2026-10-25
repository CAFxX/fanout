package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }

func main() {
    v1 := []uint32{42, 2}
    var v2 int8 = 86
    v3 := make(map[int]string)
    v2 = v2
    v4 := GenS1{F0: 2147483647, F1: 2147483647, F2: (18446744073709551615 < 355)}
    var v5 uint8 = 0
    v2 = int8(v5)
    v5 = (^(v5))
    v6 := GenS1{F0: 716, F1: 0, F2: false}
    if false {
        var v7 uint16 = 1
        v8 := GenS1{F0: 359, F1: 2, F2: true}
        _ = v7
        _ = v8
    } else {
        var v9 float64 = ((3.14159 + 1e10) + float64((((3.14159 + 0.0) / 3.14159) / 0.1)))
        _ = v9
    }
    switch "\n\t" {
    case "":
        v2 -= (v2 << ((42) & 63))
    }
    v3[1] = ("" + "\n\t")
    _ = len(v1)
    var v10 uint16 = 32767
    v11 := GenS1{F0: 2147483647, F1: 2, F2: (18446744073709551615 >= 856)}
    v12 := GenS1{F0: 65535, F1: 65535, F2: (true && (true && (!(1 < 2))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
