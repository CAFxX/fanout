package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint32, p1 uint) bool {
    var v6 int16 = 349
    v7 := []int{955, 32768}
    var v8 bool = (1e10 > 0.0)
    var v9 int16 = v6
    if true {
        switch 4294967295 {
        case -405:
            fmt.Println(v6)
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (42 != int8(p0))
}

func main() {
    var v1 float64 = (1.5 + float64(((1.5 + (((1.5 + ((3.14159 - -2.25) - 3.14159)) * 3.14159) / ((float32(1.5) + -2.25) + 1e10))) - (0.0 / 0.1))))
    var v2 uint8 = 1
    v3 := []int{0, 529}
    v4 := []uint32{335, 775, 65535}
    v5 := []int{-533, -1, 42, 9223372036854775807}
    for fi10 := 0; fi10 < 11; fi10++ {
        if len(v5) > 0 {
            v5[0] = fi10
        }
        _ = fi10
    }
    switch 32768 {
    case 42:
        switch -339 {
        case 9223372036854775807:
            v11 := GenS1{F0: 1, F1: -2147483648, F2: -128}
            _ = v11
        }
    case 878:
        for fi12 := 0; fi12 < 6; fi12++ {
            if true {
                fmt.Println(fi12)
            }
            _ = fi12
        }
    default:
        var v13 int16 = 692
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 bool = (v1 > 0.0)
        fmt.Println(v2, fi14)
        var v16 bool = false
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: -32, F1: 2147483647, F2: 127}
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint64 = 100
        fmt.Println(v2, v19, v2)
        _ = fi18
        _ = v19
    }
    var v20 int32 = -1
    v1 += v1
    var v21 int64 = int64(v2)
    v20 -= v20
    if false {
        if len(v5) > 0 {
            v5[0] = 210
        }
        var v22 uint16 = 2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v22
    } else {
        _ = len(v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v23 int = 127
    v24 := GenS2{F0: 2147483647, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
