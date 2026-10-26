package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint16; F2 uint8 }
func genFn0(p0 uint32) uint16 {
    v5 := GenS1{F0: 1, F1: 32767, F2: 32767}
    var v6 uint = 18446744073709551615
    v7 := []int{128, 32767}
    v7 = append(v7, 42)
    if false {
        if len(v7) > 0 {
            v7[0] = 441
        }
        p0 = p0
    } else {
        var v8 uint16 = 42
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64) int {
    switch ("hello" + "a") {
    case "\n\t":
        var v9 uint32 = 4294967295
        _ = v9
    case "a":
        fmt.Println(p0)
    case "hello":
        var v10 float64 = ((0.0 - -2.25) / 3.14159)
        _ = v10
    default:
        v11 := GenS2{F0: -688, F1: -1, F2: 42}
        _ = v11
    }
    switch 127 {
    case 9223372036854775807:
        p0 = (p0 | (p0 & 747))
    case -9223372036854775808:
        p0 = (^(p0))
    }
    var v12 float64 = 3.14159
    v12 = (-2.25 * v12)
    var v13 uint32 = 1
    _ = p0
    _ = v12
    _ = v13
    return -482
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    var v3 int32 = 1
    v4 := []string{}
    fmt.Println(v2)
    switch 1 {
    case 127:
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int16 = 380
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS1{F0: genFn0(0), F1: v15, F2: 169}
                fmt.Println(v3)
                v18 := GenS3{F0: v3, F1: 2147483648}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int16 = v15
            _ = fi14
            _ = v15
            _ = v19
        }
    case 32768:
        if true {
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v2, v2, v2)
                fmt.Println(fi20)
                _ = fi20
            }
        } else {
            v2 -= v2
            switch 42 {
            case 1:
                var v21 string = ""
                _ = v21
            case 42:
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 int16 = 1
                    var v24 bool = true
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                v25 := GenS4{F0: 65535}
                _ = v25
            }
            fmt.Println(v3, v2)
        }
    }
    v26 := GenS5{F0: v3, F1: 973, F2: (v2 - 127)}
    var v27 uint64 = 4294967295
    v28 := GenS1{F0: 255, F1: int16(v27), F2: 90}
    var v29 uint = 83
    v2 += v2
    var v30 uint32 = 4294967295
    var v31 bool = ((v27 != v27) || (127 <= -44))
    v32 := GenS5{F0: (v3 % ((v3) | 1)), F1: 2, F2: (v2 ^ 188)}
    v33 := GenS4{F0: (v30 * v30)}
    switch "\n\t" {
    case "a":
        var v34 uint = 18446744073709551615
        _ = v34
    default:
        var v35 uint64 = 459
        _ = v35
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
