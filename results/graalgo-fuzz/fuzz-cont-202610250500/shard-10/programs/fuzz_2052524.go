package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int; F2 string }
func genFn0(p0 int32) uint32 {
    var v5 int64 = -67
    switch 4294967295 {
    case -2147483648:
        var v6 uint8 = 2
        _ = v6
    }
    if (true && (1 != 255)) {
        var v7 int8 = 100
        switch 175 {
        case 0:
            v5 -= v5
        default:
            v8 := GenS1{F0: 125, F1: 636, F2: ("a" + ("a" + "foo bar"))}
            _ = v8
        }
        var v9 int16 = -523
        _ = v7
        _ = v9
    } else {
        if (false && (0 == v5)) {
            var v10 int = -325
            var v11 float64 = ((0.0 * ((1.5 - 3.14159) + 1.5)) / (((-2.25 + float64(0.0)) - ((0.0 + ((-2.25 - ((1e10 / 3.14159) * -2.25)) + (-2.25 - (3.14159 + 3.14159)))) / 1.0)) + 1e10))
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi12)
            var v13 float32 = 0.0
            v14 := GenS2{F0: (fi12 % ((-9223372036854775808) | 1))}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 float32 = 0.1
    var v2 int = 36
    var v3 int64 = 0
    var v4 bool = (1 < 4294967295)
    v15 := []int{(^(v2)), v2, v2, v2}
    v15 = append(v15, v2)
    v2 = 2147483647
    v1 -= (v1 * (-2.25 * v1))
    if (!v4) {
        if (!(!v4)) {
            v3 = (^(v3))
            v3 = v3
            for fi16 := 0; fi16 < 18; fi16++ {
                if (fi16 >= v2) {
                    if (100 == -26) {
                        v17 := GenS2{F0: v2}
                        var v18 uint16 = 65535
                        _ = v17
                        _ = v18
                    }
                    var v19 int32 = 2147483647
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v2)
                        var v21 float32 = float32(float64((3.14159 / 0.1)))
                        v22 := GenS2{F0: fi16}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = v19
                }
                _ = fi16
            }
        }
        if len(v15) > 0 {
            v15[0] = 127
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            v15 = append(v15, 9223372036854775807)
            _ = fi23
        }
    } else {
        fmt.Println(v3)
        v24 := GenS3{F0: 32767, F1: (-(v2)), F2: "hello"}
        var v25 int8 = 100
        _ = v24
        _ = v25
    }
    if (false && (2 < 2)) {
        if len(v15) > 0 {
            v15[0] = (-(v2))
        }
        fmt.Println(v1)
        v4 = ((568 <= 65535) || true)
    }
    var v26 bool = (v4 || false)
    for fi27 := 0; fi27 < 9; fi27++ {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        v26 = v4
        _ = fi27
    }
    var v28 int8 = 100
    if len(v15) > 0 {
        v15[0] = (v2 | (v2 >> ((65535) & 63)))
    }
    if (0 == 18446744073709551615) {
        fmt.Println(v2)
        fmt.Println(v4, v3)
        var v29 int64 = (v3 - (^(v3)))
        _ = v29
    } else {
        var v30 int16 = 32767
        _ = v30
    }
    var v31 int = 42
    _ = len(v15)
    switch -2147483648 {
    case 639:
        var v32 int64 = 42
        _ = v32
    }
    v28 = v28
    var v33 int8 = 83
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
