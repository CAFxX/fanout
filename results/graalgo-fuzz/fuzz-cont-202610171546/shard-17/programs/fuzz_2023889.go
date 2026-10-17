package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 2147483648
    var v6 int32 = 2147483647
    switch int(v1) {
    case 244:
        v5 = (^(v5))
    case 4294967295:
        _ = len(v3)
    }
    v1 = (v1 >> ((2) & 63))
    v6 = (^(v6))
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint32 = v5
        fmt.Println(v5, v8)
        _ = fi7
        _ = v8
    }
    var v9 uint = 4294967295
    v6 -= v6
    v10 := GenS1{F0: 7, F1: 681}
    var v11 float64 = 0.0
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int = (fi12 + (fi12 + -294))
        v5 += v5
        _ = fi12
        _ = v13
    }
    v14 := []int{42, -659, 127}
    var v15 uint32 = v5
    var v16 int = 9223372036854775807
    var v17 uint = (v9 & 653)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
