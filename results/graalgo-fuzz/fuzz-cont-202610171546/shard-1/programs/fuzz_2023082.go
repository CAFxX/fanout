package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 int32, p1 string) uint8 {
    var v5 uint32 = 720
    var v6 int8 = 127
    if (128 >= 127) {
        v7 := GenS1{F0: 2147483648, F1: ((0.0 / 1.0) <= ((((((-2.25 + (1e10 - (float64(1.5) + (1e10 + 1e10)))) / 1.0) + float64(float32(0.0))) - 1.5) * 3.14159) * float64((0.0 / 3.14159)))), F2: -32768}
        var v8 uint16 = 32767
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int64 = 9223372036854775807
            var v11 uint64 = 313
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    v12 := GenS2{F0: 288, F1: "foo bar"}
    var v13 uint16 = 65535
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v5
    _ = v6
    return 176
}

func main() {
    var v1 uint = 65535
    v2 := []uint32{1, 1, 34}
    var v3 float64 = ((1.5 / (0.1 / 1e10)) - (3.14159 * 1.5))
    var v4 uint64 = 302
    if false {
        fmt.Println(v1, v1, v3)
        var v14 string = ("foo bar" + "hello")
        if (v14 > "a") {
            if (0 > 1) {
                v14 = v14
                var v15 int16 = -32768
                _ = v15
            } else {
                var v16 uint64 = 958
                var v17 int64 = 50
                v14 = ("foo bar" + v14)
                _ = v16
                _ = v17
            }
            fmt.Println(v3)
            v18 := GenS3{F0: 2}
            _ = v18
        }
        _ = v14
    }
    if (-2.25 < -2.25) {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS2{F0: 1, F1: "a"}
            if len(v2) > 0 {
                v2[0] = 317
            }
            fmt.Println(v4)
            _ = fi19
            _ = v20
        }
        var v21 bool = false
        _ = v21
    } else {
        var v22 float32 = (((0.0 - 1e10) * 0.0) / 0.1)
        if (false && ((!(396 >= -894)) && true)) {
            var v23 int16 = 42
            _ = v23
        }
        _ = v22
    }
    v4 += 4294967295
    v24 := GenS2{F0: 2, F1: ((("a" + "a") + "foo bar") + ("a" + "\n\t"))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v25 := GenS4{F0: 32767, F1: v1}
    if len(v2) > 0 {
        v2[0] = 13
    }
    v1 = (v1 * (^(v1)))
    v26 := GenS5{F0: uint16(v4), F1: (v1 | (v1 ^ (v1 - (v1 % (((v1 + (v1 - (v1 / ((4294967295) | 1))))) | 1)))))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
