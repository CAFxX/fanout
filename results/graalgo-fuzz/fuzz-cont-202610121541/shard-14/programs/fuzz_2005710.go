package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint; F2 int }
type GenS5 struct { F0 uint32 }
func genFn0() uint16 {
    switch "\n\t" {
    case "":
        var v4 int64 = -1
        _ = v4
    case "a":
        var v5 int16 = -32768
        _ = v5
    }
    var v6 float32 = 3.14159
    _ = v6
    return 2
}

func main() {
    var v1 uint8 = 2
    var v2 int = 554
    var v3 int8 = 100
    v7 := []int{v2}
    v7 = append(v7, v2)
    var v8 int8 = (^(v3))
    fmt.Println(v2)
    if ((42 <= 1) || false) {
        for fi9 := 0; fi9 < 6; fi9++ {
            v3 = (v3 * -128)
            fmt.Println(v8, v2)
            v3 -= (v3 << ((113) & 63))
            _ = fi9
        }
        if false {
            var v10 uint8 = (^(v1))
            if false {
                v7 = append(v7, (v2 / ((711) | 1)))
            } else {
                fmt.Println(v1, v10)
                var v11 int16 = -750
                fmt.Println(v3, v10, v10)
                _ = v11
            }
            _ = v10
        } else {
            if len(v7) > 0 {
                v7[0] = (^(v2))
            }
            fmt.Println(v3, v8, v1)
        }
        var v12 int32 = -1
        _ = v12
    } else {
        var v13 uint16 = 32767
        fmt.Println(v13)
        _ = v13
    }
    v2 -= -9223372036854775808
    if (-9223372036854775808 > 42) {
        if len(v7) > 0 {
            v7[0] = 0
        }
    } else {
        v14 := GenS1{F0: 42}
        _ = v14
    }
    v15 := GenS2{F0: -705}
    v16 := GenS3{F0: -1, F1: 1, F2: 387}
    var v17 int8 = int8(v1)
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    v1 = v1
    for fi18 := 0; fi18 < 13; fi18++ {
        if (false || false) {
            for fi19 := 0; fi19 < 14; fi19++ {
                v17 -= v8
                _ = fi19
            }
            v7 = append(v7, v2)
            var v20 float64 = 1.5
            _ = v20
        }
        switch "" {
        case "hello":
            v7 = append(v7, fi18)
        default:
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 string = ("\n\t" + "\n\t")
                v23 := GenS4{F0: 0, F1: 0, F2: (fi18 << ((0) & 63))}
                v24 := GenS5{F0: 1}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        }
        if ((-2.25 / -2.25) != (1.5 - (0.1 * 0.0))) {
            var v25 uint = 255
            _ = v25
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
