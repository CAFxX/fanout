package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 int8) uint64 {
    v5 := []int{9223372036854775807, p0, p0}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int8 = (p1 & p1)
        if (p0 >= 493) {
            if (1 <= 1) {
                v8 := GenS1{F0: 1, F1: -2147483648}
                _ = v8
            } else {
                p1 = p1
            }
        } else {
            p0 = fi6
        }
        v9 := GenS1{F0: (^(v7)), F1: 0}
        _ = fi6
        _ = v7
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (-1 > -94)
    var v3 int16 = -32768
    v4 := make(map[int]bool)
    v10 := v4[9223372036854775807]
    v11 := []int{42, 1}
    if (!(!(v2 && v2))) {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 float64 = ((0.0 - float64(0.0)) / (float64(float32((3.14159 - 3.14159))) - 1e10))
            _ = fi12
            _ = v13
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    } else {
        var v14 int = -2147483648
        var v15 int16 = v3
        _ = v14
        _ = v15
    }
    var v16 uint8 = 127
    v3 -= 42
    var v17 int16 = (v3 % ((1000) | 1))
    v18 := GenS1{F0: 0, F1: 14}
    if false {
        if true {
            v19 := GenS2{F0: 1}
            v17 += v17
            _ = v19
        }
        if (!true) {
            var v20 int8 = 1
            _ = v20
        } else {
            fmt.Println(v16)
            v10 = (!true)
            fmt.Println(v3, v10, v17)
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            switch -983 {
            case -2147483648:
                fmt.Println(v10)
            case 32768:
                v1[621] = fi21
            case -548:
                _ = len(v11)
            default:
                var v22 uint16 = 1
                _ = v22
            }
            _ = fi21
        }
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint = 607
        v25 := v4[(fi23 / ((fi23) | 1))]
        for fi26 := 0; fi26 < 14; fi26++ {
            if len(v11) > 0 {
                v11[0] = 4294967295
            }
            var v27 uint32 = 332
            if len(v11) > 0 {
                v11[0] = fi23
            }
            _ = fi26
            _ = v27
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v2, v2, v3)
        _ = fi28
    }
    fmt.Println(v16, v16, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
