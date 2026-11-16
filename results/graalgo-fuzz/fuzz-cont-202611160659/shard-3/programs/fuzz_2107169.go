package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
func genFn0(p0 int16, p1 int64, p2 bool) int16 {
    p0 += p0
    if p2 {
        if (!p2) {
            if (243 > 127) {
                v6 := []int{9223372036854775807, 1, -1, -2147483648}
                _ = v6
            } else {
                fmt.Println(p1, p0, p2)
                var v7 int8 = 85
                _ = v7
            }
            v8 := []int{-583, 259}
            if false {
                var v9 uint = 610
                switch int(v9) {
                case -65:
                    v8 = append(v8, 127)
                case 563:
                    fmt.Println(v9)
                }
                _ = v9
            } else {
                var v10 uint64 = 720
                var v11 float32 = (1e10 - (3.14159 + 1.5))
                fmt.Println(p0)
                _ = v10
                _ = v11
            }
            _ = v8
        } else {
            fmt.Println(p2, p1)
            v12 := []int{}
            var v13 string = "foo bar"
            _ = v12
            _ = v13
        }
        p0 -= p0
    } else {
        switch "\n\t" {
        case "a":
            for fi14 := 0; fi14 < 14; fi14++ {
                if (true || (1 == 2)) {
                    var v15 uint = 255
                    v16 := []int{fi14, (^(fi14)), 0, int(p0)}
                    _ = v15
                    _ = v16
                } else {
                    var v17 float64 = float64((0.0 * (0.0 - -2.25)))
                    _ = v17
                }
                p2 = p2
                _ = fi14
            }
        case "foo bar":
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS1{F0: (-(fi18)), F1: 65535}
                v20 := []int{}
                p0 = (p0 * (p0 >> ((4294967295) & 63)))
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS1{F0: -1, F1: 1}
            _ = v21
        }
        var v22 float64 = 0.1
        _ = v22
    }
    fmt.Println(p1, p2, p0)
    for fi23 := 0; fi23 < 12; fi23++ {
        p1 -= (p1 % ((p1) | 1))
        fmt.Println(p1, p2, fi23)
        _ = fi23
    }
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 bool = false
    var v2 uint8 = 255
    var v3 int16 = 1
    var v4 int32 = 1
    var v5 int32 = -2147483648
    var v24 uint64 = 185
    var v25 string = ""
    var v26 uint16 = 32767
    v27 := []int{int(v2), -9223372036854775808, 128, -2147483648}
    v28 := GenS2{F0: (v2 & 0), F1: ("\n\t" + v25)}
    fmt.Println(v2)
    switch "foo bar" {
    case "":
        switch (v25 + v25) {
        case "hello":
            v24 = v24
        case "\n\t":
            fmt.Println(v1)
        case "a":
            v26 = (v26 >> ((18446744073709551615) & 63))
        }
    default:
        v29 := GenS3{F0: v1, F1: genFn0(v3, -1, true), F2: -1}
        _ = v29
    }
    v24 = v24
    fmt.Println(v25)
    var v30 uint8 = uint8(v26)
    var v31 int8 = 42
    v32 := GenS3{F0: v1, F1: int16(v5), F2: int64(v31)}
    var v33 uint16 = (v26 << ((2) & 63))
    var v34 uint64 = (v24 / (((v24 + (v24 ^ (^(v24))))) | 1))
    v26 = (v26 % ((v26) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
