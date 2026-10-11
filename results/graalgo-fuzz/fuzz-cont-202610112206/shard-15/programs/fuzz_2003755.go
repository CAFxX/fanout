package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
func genFn0() uint8 {
    v4 := []int{-9223372036854775808, 32768, 9223372036854775807, 127}
    if false {
        var v5 int8 = 100
        _ = v5
    } else {
        var v6 uint64 = 2
        var v7 uint8 = 158
        v8 := GenS1{F0: (^(v7))}
        _ = v6
        _ = v7
        _ = v8
    }
    _ = v4
    return 255
}

func main() {
    v1 := []int{-2147483648, 9223372036854775807, 128}
    var v2 float32 = 3.14159
    var v3 int64 = 42
    var v9 uint32 = 42
    fmt.Println(v9, v2, v3)
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int16 = 1
        var v12 uint32 = 255
        var v13 uint = 643
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: 1, F1: 474, F2: int(v14)}
        if ("\n\t" == ("" + "hello")) {
            switch -9223372036854775808 {
            case 9223372036854775807:
                v3 += -9223372036854775808
            case 128:
                if (!(genFn0() >= 127)) {
                    v9 = v9
                    _ = len(v1)
                } else {
                    v17 := GenS1{F0: 67}
                    _ = v17
                }
            }
            switch -1 {
            case 0:
                var v18 uint32 = v9
                _ = v18
            case -425:
                v14 += v14
            default:
                fmt.Println(v3, v2, v14)
            }
            if ((v9 ^ v9) < (v9 << (((^(v14))) & 63))) {
                v1 = append(v1, (^(fi15)))
                if len(v1) > 0 {
                    v1[0] = fi15
                }
                if false {
                    if len(v1) > 0 {
                        v1[0] = fi15
                    }
                }
            }
        }
        _ = fi15
        _ = v16
    }
    var v19 uint64 = 574
    if (!((2 == 42) && (-1 < 0))) {
        var v20 int8 = 42
        var v21 float64 = -2.25
        fmt.Println(v3)
        _ = v20
        _ = v21
    } else {
        var v22 uint16 = 32767
        var v23 float64 = (1.5 / (0.0 - float64(0.1)))
        _ = len(v1)
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
