package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 uint64 = 42
    var v5 string = "hello"
    v1 -= v1
    v6 := GenS1{F0: 42}
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v7 = v7
            v5 = "\n\t"
            switch fi8 {
            case 42:
                _ = len(v2)
            case 674:
                if len(v3) > 0 {
                    v3[0] = (fi8 & fi9)
                }
                v10 := GenS2{F0: 1, F1: true, F2: 1}
                _ = v10
            }
            _ = fi9
        }
        v2[(fi8 % ((200) | 1))] = fi8
        var v11 float32 = 0.0
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 1}
        fmt.Println(v1, v7, fi12)
        v4 = (v4 * (v4 / ((v4) | 1)))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            v7 += ((v7 / v7) - 3.14159)
            _ = fi15
        }
        var v16 float64 = v7
        var v17 uint = 968
        _ = fi14
        _ = v16
        _ = v17
    }
    var v18 bool = (1.5 != -2.25)
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 uint32 = 946
        var v21 int8 = -1
        fmt.Println(fi19)
        _ = fi19
        _ = v20
        _ = v21
    }
    v1 = (v1 * 1)
    var v22 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
