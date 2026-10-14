package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 string }
func genFn0(p0 uint, p1 int16) int64 {
    v5 := GenS1{F0: 42}
    var v6 float32 = (float32(3.14159) - 0.0)
    switch 128 {
    case -1:
        switch 128 {
        case 2147483647:
            fmt.Println(p1)
        case 9223372036854775807:
            v6 += 3.14159
        }
    default:
        v7 := []int{}
        _ = v7
    }
    switch 32767 {
    case 411:
        v8 := GenS2{F0: p0, F1: 255, F2: 255}
        _ = v8
    case 9223372036854775807:
        fmt.Println(v6)
    case -643:
        var v9 uint8 = 128
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -126
}

func main() {
    var v1 float64 = ((3.14159 / 1.0) + (float64(1e10) * 0.1))
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := []string{""}
    for fi10 := 0; fi10 < 14; fi10++ {
        if (false && (0.0 > 1e10)) {
            var v11 bool = (((fi10 | fi10) != (-(fi10))) || true)
            _ = v11
        }
        _ = fi10
    }
    v1 = (-2.25 - v1)
    if (1 < 398) {
        switch 128 {
        case 4294967295:
            for fi12 := 0; fi12 < 16; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int = (fi12 * -714)
                    var v15 float32 = (3.14159 / ((3.14159 * (float32(v1) * 0.1)) / 1.5))
                    var v16 int8 = 0
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                v1 += (1e10 / 1.5)
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                _ = fi12
            }
        default:
            switch 9223372036854775807 {
            case 599:
                fmt.Println(v1)
            case 127:
                v17 := GenS1{F0: 127}
                _ = v17
            case 1:
                v1 -= -2.25
            default:
                v18 := GenS3{F0: 255, F1: 2}
                _ = v18
            }
        }
        fmt.Println(v1)
        var v19 int64 = -1
        _ = v19
    }
    switch 490 {
    case 913:
        var v20 uint32 = 908
        _ = v20
    case -635:
        fmt.Println(v1)
    default:
        v21 := GenS4{F0: 2147483647, F1: (127 != 32768)}
        _ = v21
    }
    var v22 int32 = -2147483648
    switch -184 {
    case -1:
        for fi23 := 0; fi23 < 20; fi23++ {
            delete(v3, "hello")
            for fi24 := 0; fi24 < 1; fi24++ {
                v22 = 163
                if (((true && (!true)) || true) && (!(255 == 255))) {
                    fmt.Println(v22, v1, v22)
                } else {
                    if true {
                        v1 += 3.14159
                    }
                    v25 := GenS2{F0: 255, F1: 2, F2: 65535}
                    v22 = (v22 % (((v22 + (v22 ^ 1))) | 1))
                    _ = v25
                }
                _ = fi24
            }
            _ = fi23
        }
    case 32767:
        if true {
            var v26 int16 = 1
            _ = v26
        }
    }
    switch -2147483648 {
    case 114:
        v1 -= ((float64(((1e10 * float32(0.1)) / (1.5 + float32(v1)))) - float64(1.5)) * 1.5)
    case 0:
        v27 := GenS3{F0: 122, F1: 464}
        _ = v27
    default:
        var v28 uint = 2
        _ = v28
    }
    if false {
        var v29 string = "hello"
        switch 1 {
        case -9223372036854775808:
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v1)
            case 4294967295:
                for fi30 := 0; fi30 < 4; fi30++ {
                    v2[fi30] = 128
                    v22 += 42
                    var v31 int64 = 0
                    _ = fi30
                    _ = v31
                }
            default:
                fmt.Println(v22, v22)
            }
        case 833:
            v32 := GenS5{F0: "a"}
            _ = v32
        }
        _ = len(v3)
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
