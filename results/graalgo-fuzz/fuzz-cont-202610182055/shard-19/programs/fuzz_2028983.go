package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 int16, p1 int16, p2 uint16) int8 {
    v5 := []int{}
    v5 = append(v5, 326)
    var v6 uint64 = 42
    p0 -= -1
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 += v6
            if false {
                v6 -= v6
                fmt.Println(fi8)
                v9 := GenS1{F0: (128 > 0)}
                _ = v9
            }
            var v10 uint = 18446744073709551615
            v11 := GenS1{F0: false}
            var v12 int64 = int64(v10)
            if true {
                var v13 int8 = -1
                _ = v13
            }
            p2 = (p2 & (p2 & (^(p2))))
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 string = (((("hello" + "hello") + "foo bar") + ("hello" + "hello")) + "")
            _ = fi14
            _ = v15
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int64) uint8 {
    v16 := []int{42, 32768, 127, 1}
    fmt.Println(p0)
    if (0 < 0) {
        p0 = p0
        fmt.Println(p0)
        if (!(1.5 > (((((0.0 + 1.5) / (1.5 - 3.14159)) / (1.5 - float64(0.0))) + 1e10) * 1.5))) {
            if len(v16) > 0 {
                v16[0] = -230
            }
            if true {
                p0 += -9223372036854775808
                var v17 float32 = 0.1
                _ = v17
            }
        }
    }
    var v18 uint32 = 65535
    v16 = append(v16, 2147483647)
    _ = p0
    _ = v16
    _ = v18
    return 255
}

func main() {
    var v1 int8 = 42
    var v2 int = -468
    var v3 int8 = (v1 - (v1 + v1))
    var v4 float64 = 0.1
    v19 := GenS2{F0: 1, F1: v2}
    switch "foo bar" {
    case "a":
        if true {
            fmt.Println(v1, v2)
            v20 := []int{v2, -761}
            _ = v20
        }
    case "foo bar":
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := []int{v2, 0}
            var v23 int16 = -1
            fmt.Println(fi21)
            _ = fi21
            _ = v22
            _ = v23
        }
    case "":
        fmt.Println(v4, v1, v3)
    default:
        v24 := GenS2{F0: 32767, F1: (v2 ^ 35)}
        _ = v24
    }
    v1 = (v1 - v3)
    switch "hello" {
    case "\n\t":
        v4 = ((((1e10 - v4) - v4) / (1e10 * v4)) * -2.25)
    default:
        var v25 uint8 = 0
        _ = v25
    }
    v26 := []int{32768, 50, v2, v2}
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v3 = genFn0(32767, 1, 32767)
    v1 += v1
    v27 := GenS3{F0: 354, F1: 1, F2: 2147483647}
    v2 = (v2 * (-(v2)))
    v28 := GenS2{F0: 523, F1: (^(v2))}
    fmt.Println(v2)
    v29 := GenS2{F0: 32767, F1: -2147483648}
    v3 += 42
    v4 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
