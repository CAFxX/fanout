package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int16; F1 int16 }
func genFn0(p0 bool) int16 {
    p0 = (!(115 >= 0))
    var v5 string = "hello"
    _ = p0
    _ = v5
    return -1
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    var v3 int32 = -2147483648
    v4 := make(map[int]uint32)
    if (true && false) {
        switch "foo bar" {
        case "a":
            for fi6 := 0; fi6 < 4; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 uint64 = 42
                    v4[fi7] = 4294967295
                    _ = fi7
                    _ = v8
                }
                fmt.Println(v3, fi6)
                _ = fi6
            }
        case "hello":
            v9 := GenS1{F0: -1, F1: v3}
            _ = v9
        default:
            v10 := GenS2{F0: 42, F1: 42, F2: -743}
            _ = v10
        }
        v3 -= (v3 >> ((42) & 63))
        if true {
            _ = len(v1)
            v3 += v3
            if true {
                _ = len(v1)
                var v11 int64 = v2
                _ = v11
            }
        } else {
            fmt.Println(v3, v3)
            var v12 int8 = 0
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS2{F0: 32767, F1: 42, F2: genFn0(true)}
            _ = fi13
            _ = v14
        }
        v3 -= (v3 % ((v3) | 1))
        v15 := GenS3{F0: 454, F1: (^(v3)), F2: v3}
        _ = v15
    }
    switch -9223372036854775808 {
    case 32768:
        fmt.Println(v2)
    case 46:
        v16 := v4[-1]
        _ = v16
    }
    fmt.Println(v2)
    var v17 uint8 = 2
    var v18 uint16 = 42
    delete(v4, 128)
    v19 := GenS4{F0: 32767}
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(fi20, v17, v3)
        v17 = v17
        _ = fi20
    }
    v21 := GenS5{F0: 1, F1: v17}
    _ = len(v4)
    fmt.Println(v17, v3)
    switch ("a" + ("foo bar" + "a")) {
    case "\n\t":
        v17 = 0
    case "hello":
        v22 := GenS6{F0: 0, F1: 32767}
        _ = v22
    default:
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS2{F0: -392, F1: 42, F2: -32768}
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS7{F0: genFn0(false), F1: -32768}
                _ = fi25
                _ = v26
            }
            if false {
                var v27 int = fi23
                fmt.Println(v3, v2, v17)
                delete(v4, (^(fi23)))
                _ = v27
            } else {
                fmt.Println(fi23, v17, v2)
            }
            _ = fi23
            _ = v24
        }
    }
    var v28 float64 = ((1.5 - -2.25) * ((float64(-2.25) / (-2.25 / -2.25)) + float64((-2.25 - ((3.14159 * -2.25) / (((1.5 / (1e10 / 0.1)) + (0.0 / 3.14159)) * (float32(3.14159) * 1.5)))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}
