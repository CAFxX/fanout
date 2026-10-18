package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 int16, p1 uint8, p2 uint) int16 {
    fmt.Println(p2, p1, p0)
    if true {
        fmt.Println(p0)
        v6 := []int{-9223372036854775808}
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        if (false || (true || (uint64(p2) <= 174))) {
            if true {
                v8 := GenS1{F0: (p0 << ((p2) & 63)), F1: 0, F2: -78}
                _ = v8
            }
        } else {
            switch fi7 {
            case 4294967295:
                p1 += (^(p1))
            }
        }
        if (((fi7 ^ 4294967295) <= fi7) && (0.0 > 0.0)) {
            v9 := GenS1{F0: 1, F1: 0, F2: -2147483648}
            p0 += p0
            _ = v9
        }
        var v10 uint8 = 89
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 bool = ((0.1 > -2.25) || false)
    var v5 string = "a"
    var v11 int = -9223372036854775808
    v1 += -1
    v12 := []int{v11, v11}
    fmt.Println(v1)
    v5 = (v5 + "a")
    var v13 float32 = (3.14159 / float32((3.14159 * 3.14159)))
    if true {
        v5 = v5
    } else {
        v14 := GenS1{F0: -32768, F1: (v1 / (((v1 & (v1 % ((v1) | 1)))) | 1)), F2: 431}
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v13 -= ((v13 * 1e10) * 0.1)
        _ = fi15
    }
    switch v11 {
    case 127:
        v16 := GenS1{F0: 0, F1: v1, F2: 0}
        _ = v16
    default:
        var v17 float32 = v13
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            switch v5 {
            case "foo bar":
                v11 = -89
            case "hello":
                v20 := GenS2{F0: 562, F1: 2, F2: 255}
                _ = v20
            default:
                v4 = (-1 < 1)
            }
            _ = len(v3)
            fmt.Println(v1, v13)
            _ = fi19
        }
        _ = fi18
    }
    var v21 float32 = (v13 - float32(float64((v13 - 1.5))))
    for fi22 := 0; fi22 < 2; fi22++ {
        if false {
            v1 -= (v1 / ((v1) | 1))
        }
        fmt.Println(v4, v11)
        _ = fi22
    }
    v1 += (^(v1))
    _ = len(v12)
    for fi23 := 0; fi23 < 17; fi23++ {
        v1 = (v1 << ((65535) & 63))
        fmt.Println(v1, v5)
        _ = fi23
    }
    switch 4294967295 {
    case 128:
        var v24 int32 = 314
        _ = v24
    default:
        var v25 float32 = v13
        _ = v25
    }
    switch v11 {
    case -9223372036854775808:
        var v26 uint8 = 42
        _ = v26
    }
    var v27 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
