package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int; F2 uint }
func genFn0() uint {
    var v6 string = "foo bar"
    v7 := []int{42}
    v6 = v6
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 string, p1 uint16) uint8 {
    var v8 uint8 = 127
    var v9 float32 = float32((1e10 + (1.5 - (((((0.1 / (-2.25 * 3.14159)) + 1e10) + 1.5) / (3.14159 - ((1.5 + 0.0) - (0.1 * float64(0.0))))) / 3.14159))))
    p0 = p0
    v8 = (v8 ^ 42)
    p0 = ("hello" + (p0 + "hello"))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (v8 ^ v8)
}

func main() {
    v1 := []int{-918, 127, 741}
    v2 := make(map[int]uint32)
    var v3 uint64 = 1
    var v4 bool = false
    v5 := []int{0, -883, 1, 1}
    v10 := GenS1{F0: 42}
    switch -13 {
    case 32767:
        _ = len(v5)
    case 2147483647:
        var v11 int = 127
        _ = v11
    case -851:
        var v12 bool = (!(-128 < 0))
        _ = v12
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            _ = len(v1)
            var v14 int8 = -128
            v15 := GenS2{F0: uint(v3), F1: fi13, F2: 237}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 int = 127
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        var v17 int64 = -921
        _ = v17
    }
    v3 += v3
    v16 += v16
    for fi18 := 0; fi18 < 18; fi18++ {
        v3 = v3
        var v19 int16 = -1
        if v4 {
            for fi20 := 0; fi20 < 15; fi20++ {
                v16 = fi20
                v21 := GenS2{F0: 18446744073709551615, F1: (fi18 * 42), F2: 255}
                _ = fi20
                _ = v21
            }
            fmt.Println(fi18)
        }
        _ = fi18
        _ = v19
    }
    v4 = (-64 != -82)
    switch v16 {
    case 158:
        if (v4 && v4) {
            var v22 uint64 = (v3 >> ((18446744073709551615) & 63))
            v3 += 2
            v16 = (v16 / ((v16) | 1))
            _ = v22
        }
    }
    v16 += v16
    var v23 uint16 = 266
    if ((v16 + v16) < v16) {
        var v24 uint16 = (^(v23))
        v25 := v2[(v16 % (((v16 - v16)) | 1))]
        fmt.Println(v4)
        _ = v24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
