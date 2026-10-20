package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 int16; F2 int64 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 int16, p1 uint32) uint16 {
    switch 4294967295 {
    case 9223372036854775807:
        if true {
            var v7 float32 = 0.1
            fmt.Println(p1, p0, v7)
            if false {
                var v8 uint16 = 42
                _ = v8
            } else {
                v9 := GenS1{F0: 32767}
                switch 127 {
                case 2147483647:
                    v10 := GenS2{F0: 4294967295, F1: 1, F2: ""}
                    _ = v10
                case -900:
                    var v11 int64 = 0
                    _ = v11
                case -128:
                    p0 += -952
                default:
                    var v12 int8 = 117
                    _ = v12
                }
                _ = v9
            }
            _ = v7
        }
    case -9223372036854775808:
        var v13 float32 = 0.1
        _ = v13
    case -1:
        v14 := []int{-443, 0}
        _ = v14
    }
    if false {
        p0 += p0
        if (42 < 575) {
            p0 = 811
            if true {
                v15 := []int{32767, 2147483647, 46}
                _ = v15
            }
            if true {
                fmt.Println(p1, p1)
                p1 = 1
            }
        } else {
            v16 := GenS3{F0: 18446744073709551615, F1: p0, F2: 0}
            switch 128 {
            case 4294967295:
                p1 += (p1 * 179)
            default:
                fmt.Println(p0, p0)
            }
            _ = v16
        }
    }
    fmt.Println(p1)
    var v17 float32 = 0.0
    _ = p0
    _ = p1
    _ = v17
    return 65535
}
func genFn1() int16 {
    var v18 float32 = -2.25
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int16 = 32767
        _ = fi19
        _ = v20
    }
    v18 = v18
    v18 -= (v18 / 1.0)
    _ = v18
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 float32 = -2.25
    var v3 int16 = 1000
    v4 := make(map[uint32]uint32)
    var v5 float32 = float32(((((float64(((((float32(0.1) / 1e10) * v2) * ((((v2 - (v2 * (v2 / v2))) + 0.0) * v2) - ((v2 / (v2 - v2)) / 0.1))) - (3.14159 + v2))) + 3.14159) * ((0.1 + (-2.25 * (0.0 - 3.14159))) / 1.0)) + 0.0) - (-2.25 - (1.5 + 3.14159))))
    var v6 uint8 = 249
    var v21 int32 = 740
    v22 := []int{1, -2147483648, 418}
    if len(v22) > 0 {
        v22[0] = -663
    }
    var v23 float64 = float64(float32(1e10))
    v21 += v21
    var v24 int64 = (v1 % ((v1) | 1))
    v21 = -2147483648
    var v25 string = (("foo bar" + "") + "foo bar")
    _ = len(v22)
    if false {
        v26 := GenS4{F0: 26, F1: 0}
        _ = v26
    }
    var v27 string = (("hello" + "a") + ("hello" + v25))
    var v28 string = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
