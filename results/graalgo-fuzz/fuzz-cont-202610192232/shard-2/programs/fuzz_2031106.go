package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8 }
func genFn0() uint {
    v4 := GenS1{F0: -1, F1: 0}
    var v5 int64 = 0
    v5 = (v5 ^ -869)
    _ = v4
    _ = v5
    return 970
}
func genFn1(p0 bool, p1 uint) int32 {
    v6 := GenS1{F0: 42, F1: 0}
    var v7 int8 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -435
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -348
    var v3 float64 = -2.25
    fmt.Println(v3)
    v3 -= v3
    v2 -= v2
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint16 = 2
            v11 := []int{2147483647, (fi9 % ((4294967295) | 1)), fi9, (^(fi9))}
            if (1 == 239) {
                var v12 int16 = -43
                v13 := GenS1{F0: (v2 % (((-(v2))) | 1)), F1: 42}
                if true {
                    fmt.Println(v2, v2)
                    v3 = 1.5
                    var v14 float32 = 0.1
                    _ = v14
                }
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v2 += v2
    v15 := GenS2{F0: -74}
    if false {
        v2 -= (v2 + (v2 | v2))
    } else {
        var v16 uint = 943
        var v17 int64 = 9223372036854775807
        _ = v16
        _ = v17
    }
    switch 609 {
    case -2147483648:
        var v18 int8 = -128
        _ = v18
    case 1:
        var v19 uint8 = 2
        _ = v19
    }
    var v20 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
