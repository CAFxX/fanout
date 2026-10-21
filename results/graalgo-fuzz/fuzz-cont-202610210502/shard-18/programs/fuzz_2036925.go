package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint) uint16 {
    v7 := GenS1{F0: 34, F1: -902, F2: p1}
    if (!false) {
        var v8 uint8 = 234
        p1 = p1
        if ((p1 | 338) == (p1 % (((p1 % ((p1) | 1))) | 1))) {
            var v9 string = ""
            switch v9 {
            case "a":
                p2 = 0
            case "foo bar":
                v10 := []int{128, -9223372036854775808, 471, 32767}
                _ = v10
            case "":
                var v11 int32 = 42
                _ = v11
            default:
                p0 = 18446744073709551615
            }
            _ = v9
        }
        _ = v8
    }
    v12 := []int{32768, -2147483648, -9223372036854775808, 2147483647}
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, p1, p1)
        for fi14 := 0; fi14 < 11; fi14++ {
            switch (-(fi14)) {
            case 32767:
                var v15 string = "foo bar"
                var v16 uint8 = 145
                v15 = v15
                var v17 int = (-(fi14))
                fmt.Println(v17, fi13, fi14)
                var v18 int64 = 42
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            case -236:
                fmt.Println(p1, p1, p1)
            case -771:
                switch "\n\t" {
                case "hello":
                    v19 := GenS1{F0: 2, F1: 0, F2: p1}
                    _ = v19
                case "a":
                    p2 = 4294967295
                case "":
                    if true {
                        var v20 uint32 = 575
                        var v21 int32 = 735
                        p2 = p2
                        _ = v20
                        _ = v21
                    }
                default:
                    fmt.Println(fi13, fi13)
                }
            }
            if (!((p0 + p0) == (^(p0)))) {
                fmt.Println(p0, p1, fi13)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v12 = append(v12, fi13)
            } else {
                fmt.Println(p1, p0, fi14)
                var v22 uint32 = 4294967295
                _ = v22
            }
            v23 := GenS1{F0: 128, F1: -32768, F2: (p1 & uint16(p2))}
            _ = fi14
            _ = v23
        }
        _ = fi13
    }
    if true {
        var v24 int8 = -128
        fmt.Println(p1, p2)
        var v25 uint64 = 0
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 161
    var v3 uint64 = 42
    var v4 bool = (!true)
    var v5 bool = ((6 < 68) && false)
    var v6 int32 = 2147483647
    v26 := []int{}
    if (42 != 534) {
        var v27 float32 = float32(float64((float32(3.14159) - (3.14159 * (0.0 - 0.1)))))
        if v5 {
            v4 = v5
            switch 32767 {
            case -737:
                v28 := GenS2{F0: 1, F1: -868}
                _ = v28
            case -9223372036854775808:
                fmt.Println(v3)
            case 4294967295:
                var v29 int64 = v2
                _ = v29
            default:
                v6 = v6
            }
            fmt.Println(v3)
        }
        _ = v27
    }
    var v30 float64 = 1.5
    var v31 float64 = v30
    v6 = v6
    v2 = (v2 * (v2 & 1))
    var v32 uint32 = 1
    var v33 int64 = 1
    var v34 string = ("foo bar" + ("" + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
