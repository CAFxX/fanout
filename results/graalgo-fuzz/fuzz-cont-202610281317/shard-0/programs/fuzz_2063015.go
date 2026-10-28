package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int64) int16 {
    v6 := GenS1{F0: 42}
    p0 = 42
    p0 -= (p0 & p0)
    switch -2147483648 {
    case 2147483647:
        v7 := GenS2{F0: 4294967295, F1: -128}
        _ = v7
    case 42:
        for fi8 := 0; fi8 < 7; fi8++ {
            if (673 < 372) {
                var v9 uint16 = 1
                _ = v9
            }
            _ = fi8
        }
    }
    _ = p0
    _ = v6
    return 754
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]int)
    var v3 string = "\n\t"
    var v4 uint = 2
    v5 := make(map[uint32]string)
    var v10 float64 = -2.25
    var v11 int = 1
    var v12 bool = true
    v5[0] = v3
    var v13 uint64 = 4294967295
    var v14 int32 = 42
    fmt.Println(v4)
    fmt.Println(v14, v10, v10)
    v10 = v10
    if (1e10 <= 0.1) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v1 -= (v1 >> ((v4) & 63))
            v16 := GenS3{F0: 100, F1: v14}
            _ = fi15
            _ = v16
        }
    }
    v17 := []int{v11}
    var v18 uint = 1
    fmt.Println(v1, v11, v1)
    if len(v17) > 0 {
        v17[0] = (v11 << (((v4 - (v4 % (((v4 * (v4 / ((v4) | 1)))) | 1)))) & 63))
    }
    var v19 float64 = (((v10 * 0.0) * 0.0) + -2.25)
    var v20 int32 = v14
    if v12 {
        v17 = append(v17, -707)
    } else {
        var v21 int8 = 82
        _ = v21
    }
    v10 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
