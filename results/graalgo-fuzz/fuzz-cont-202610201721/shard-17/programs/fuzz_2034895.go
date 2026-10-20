package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
func genFn0(p0 int32, p1 bool) uint {
    var v6 int32 = -780
    v7 := GenS1{F0: (p1 || (-9223372036854775808 != 0)), F1: 436}
    v8 := []int{-9223372036854775808, 42, -445}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int32 = p0
        _ = fi9
        _ = v10
    }
    if len(v8) > 0 {
        v8[0] = 128
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int32, p1 string, p2 int64) uint16 {
    var v11 int16 = -1
    var v12 uint64 = 1
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 311
}

func main() {
    var v1 int16 = 0
    var v2 int8 = 127
    var v3 int32 = 42
    var v4 bool = true
    v5 := make(map[int]string)
    if v4 {
        if (2 > 4294967295) {
            switch -1 {
            case 32767:
                v14 := GenS1{F0: v4, F1: 65535}
                _ = v14
            case -9223372036854775808:
                v15 := GenS2{F0: (^(v3)), F1: v1, F2: -1}
                _ = v15
            default:
                v16 := GenS2{F0: -125, F1: 42, F2: 1}
                _ = v16
            }
            v1 -= (v1 + (v1 * (v1 | 617)))
        } else {
            if ("a" <= "hello") {
                fmt.Println(v3, v1)
                var v17 uint64 = 134
                _ = v17
            }
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS1{F0: (!false), F1: 255}
                _ = fi18
                _ = v19
            }
        }
        var v20 string = ("a" + "a")
        v21 := GenS2{F0: 42, F1: 42, F2: 1}
        _ = v20
        _ = v21
    }
    var v22 float64 = (3.14159 / ((((0.1 * 1e10) + 3.14159) / 1.5) - 0.0))
    if v4 {
        if (("foo bar" > "\n\t") && false) {
            delete(v5, -1)
            if (4294967295 <= 1) {
                if v4 {
                    var v23 int64 = 785
                    v24 := GenS2{F0: -1, F1: (v1 >> ((65535) & 63)), F2: v23}
                    delete(v5, -463)
                    _ = v23
                    _ = v24
                }
            } else {
                var v25 int64 = 482
                _ = v25
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float32 = (0.0 * (float32(float64((0.0 / float32((v22 - v22))))) - (-2.25 + 3.14159)))
                delete(v5, (^(fi26)))
                _ = fi26
                _ = v27
            }
        } else {
            var v28 int64 = -9223372036854775808
            _ = v28
        }
        v2 = v2
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v4 = (((!v4) && v4) && v4)
        v4 = true
        if v4 {
            v30 := []int{fi29, 0, fi29, 9223372036854775807}
            _ = v30
        }
        _ = fi29
    }
    v31 := []int{}
    var v32 uint8 = 0
    var v33 float32 = -2.25
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 int = 2147483647
        fmt.Println(v35, v32)
        _ = fi34
        _ = v35
    }
    var v36 float64 = v22
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v4)
        v32 -= 1
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
