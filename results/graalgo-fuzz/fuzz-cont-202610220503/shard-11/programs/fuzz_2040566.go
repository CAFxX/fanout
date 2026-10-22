package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0(p0 int) int32 {
    v7 := []int{-2147483648, -673}
    _ = len(v7)
    fmt.Println(p0)
    p0 += (p0 + (p0 >> ((0) & 63)))
    p0 = (p0 % ((-9223372036854775808) | 1))
    _ = p0
    _ = v7
    return -673
}
func genFn1(p0 uint8, p1 int8, p2 uint8) int {
    v8 := GenS1{F0: 32768}
    v9 := GenS1{F0: 32767}
    v10 := []int{}
    var v11 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 int64 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    var v6 bool = (!false)
    v6 = (v3 < (^(v3)))
    var v12 float32 = (((-2.25 / 1.0) * 3.14159) / ((0.1 * 1.5) / 1.0))
    v3 -= v3
    if ((v3 | (^(v3))) >= v3) {
        delete(v4, "a")
    } else {
        switch 32768 {
        case -2147483648:
            if (false || (9223372036854775807 <= -9223372036854775808)) {
                _ = len(v4)
                fmt.Println(v3, v3)
                if (!v6) {
                    v6 = ((v3 - v3) <= (v3 * (^(v3))))
                    v3 = v3
                    var v13 float64 = 3.14159
                    _ = v13
                } else {
                    v2 = append(v2, 819)
                }
            } else {
                v14 := GenS1{F0: 42}
                if (42 <= 111) {
                    v3 -= (v3 % (((-(v3))) | 1))
                    v15 := GenS2{F0: 1}
                    if false {
                        var v16 int16 = -1
                        _ = v16
                    } else {
                        v12 -= ((float32(1e10) + 3.14159) / (v12 + float32(((float64(v12) / (0.0 - float64((1.5 - (1e10 + v12))))) - (3.14159 - 1e10)))))
                        var v17 float32 = v12
                        v18 := GenS2{F0: v3}
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
                _ = v14
            }
        default:
            v19 := GenS1{F0: 245}
            _ = v19
        }
        switch 32768 {
        case 4294967295:
            switch "hello" {
            case "\n\t":
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := GenS2{F0: (v3 * v3)}
                    _ = fi20
                    _ = v21
                }
            case "":
                v12 -= v12
            }
        case 9223372036854775807:
            v3 = -399
        }
        switch 32767 {
        case -816:
            switch 77 {
            case -2147483648:
                v6 = (2 > 18446744073709551615)
            default:
                switch -1 {
                case 596:
                    var v22 int64 = v3
                    _ = v22
                case 748:
                    v6 = v6
                }
            }
        case 1:
            v23 := GenS1{F0: 9223372036854775807}
            _ = v23
        case -2147483648:
            v12 = v12
        }
    }
    var v24 uint32 = 2
    if (2 != 32767) {
        var v25 uint8 = 2
        _ = v25
    }
    switch 4294967295 {
    case -625:
        v26 := GenS2{F0: (v3 ^ v3)}
        _ = v26
    case -651:
        var v27 float64 = 1.5
        _ = v27
    }
    v3 += 42
    v24 += (v24 * (v24 % (((^(v24))) | 1)))
    switch "hello" {
    case "\n\t":
        v28 := GenS3{F0: 0, F1: 4294967295, F2: v24}
        _ = v28
    case "":
        v6 = ((("foo bar" + ("" + "hello")) == "hello") && true)
    case "foo bar":
        var v29 int32 = 515
        _ = v29
    }
    var v30 uint8 = 200
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 float32 = v12
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
