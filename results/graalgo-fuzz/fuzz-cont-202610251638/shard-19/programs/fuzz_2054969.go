package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int16 {
    var v6 uint32 = 0
    var v7 float32 = float32((1.5 * -2.25))
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := []int{128, 9223372036854775807, 2147483647, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 float64 = 1.5
    var v4 uint64 = 984
    v5 := []int{-820}
    var v8 uint = 331
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v8, v8)
        var v10 uint32 = 4294967295
        switch 42 {
        case 585:
            v4 = v4
        case -1:
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: (v8 / (((^(v8))) | 1))}
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 32768
    }
    if true {
        var v13 bool = (-718 >= 42)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            _ = len(v2)
            _ = fi14
        }
        if true {
            fmt.Println(v3)
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v4, fi15)
                _ = fi15
            }
            var v16 int32 = -631
            _ = v16
        } else {
            var v17 uint16 = 0
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 -= (^(v4))
        _ = fi18
    }
    v8 += v8
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println(v4)
    fmt.Println(v8, v3, v3)
    var v19 uint = 255
    var v20 int16 = 1
    v2[2147483647] = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
