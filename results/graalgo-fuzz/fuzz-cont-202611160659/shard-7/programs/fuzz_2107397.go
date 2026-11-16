package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 uint16) uint64 {
    if false {
        var v4 uint64 = 1
        _ = v4
    } else {
        p0 += p0
        var v5 float32 = -2.25
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    return 571
}

func main() {
    var v1 int32 = 202
    var v2 uint64 = 693
    var v3 int16 = -32768
    switch ("foo bar" + "hello") {
    case "foo bar":
        fmt.Println(v1)
    case "a":
        if (9223372036854775807 >= 1) {
            v1 -= v1
            var v6 uint32 = 10
            _ = v6
        }
    case "hello":
        if (128 != 12) {
            fmt.Println(v3)
        } else {
            v1 = (v1 - -1)
            if true {
                v7 := []int{-9223372036854775808, 32767}
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int8 = 1
                    _ = fi8
                    _ = v9
                }
                if true {
                    var v10 uint = 255
                    _ = v10
                } else {
                    _ = len(v7)
                    var v11 bool = ((v2 & v2) <= genFn0(2))
                    _ = v11
                }
                _ = v7
            }
            v1 += v1
        }
    default:
        v3 = (v3 / (((v3 >> ((4294967295) & 63))) | 1))
    }
    var v12 float64 = (1.5 * ((1e10 * 1e10) - (0.0 / -2.25)))
    var v13 int64 = 9223372036854775807
    v13 += v13
    var v14 string = ("a" + ("foo bar" + "a"))
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            if (!(false || false)) {
                var v17 int = fi16
                if (2 == 1) {
                    var v18 uint16 = 501
                    _ = v18
                }
                _ = v17
            } else {
                v19 := []int{fi16, (fi16 | fi15), fi15}
                fmt.Println(v12, fi16, fi15)
                _ = v19
            }
            _ = fi16
        }
        switch v14 {
        case "\n\t":
            if ((v3 / ((1000) | 1)) <= v3) {
                var v20 string = v14
                _ = v20
            }
        default:
            fmt.Println(v14)
        }
        v21 := []int{9223372036854775807, 32767}
        _ = fi15
        _ = v21
    }
    v22 := []int{int(v2), -566, 9223372036854775807, 128}
    v23 := GenS1{F0: 127}
    v24 := GenS2{F0: (-(v13)), F1: -4, F2: 42}
    fmt.Println(v12, v2)
    if len(v22) > 0 {
        v22[0] = 1
    }
    v3 -= -1
    v1 = (v1 | 191)
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 int16 = v3
        _ = fi25
        _ = v26
    }
    v2 = v2
    if (int(v3) != 4294967295) {
        v12 = (v12 * v12)
    } else {
        var v27 float64 = ((-2.25 * v12) - v12)
        v3 -= (v3 * (v3 & (^(v3))))
        _ = v27
    }
    var v28 uint64 = genFn0(42)
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
}
