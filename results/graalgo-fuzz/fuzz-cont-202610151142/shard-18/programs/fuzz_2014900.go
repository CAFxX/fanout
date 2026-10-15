package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 bool; F2 int8 }
func genFn0(p0 int8) int {
    for fi6 := 0; fi6 < 5; fi6++ {
        if false {
            fmt.Println(p0, p0)
            var v7 int = fi6
            _ = v7
        } else {
            v8 := GenS1{F0: (p0 * 127), F1: (-(p0))}
            _ = v8
        }
        v9 := GenS1{F0: (p0 << ((42) & 63)), F1: (p0 + (^(p0)))}
        if (-9223372036854775808 >= -9223372036854775808) {
            p0 = p0
        }
        _ = fi6
        _ = v9
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    return 440
}

func main() {
    var v1 bool = false
    var v2 uint32 = 763
    var v3 uint32 = v2
    var v4 int64 = -610
    v5 := []string{"a", "hello", "a", "a"}
    var v10 float64 = 0.0
    var v11 float32 = (1.5 / 0.1)
    var v12 uint = 726
    v13 := GenS1{F0: 100, F1: 18}
    for fi14 := 0; fi14 < 13; fi14++ {
        v3 = (v3 - (v3 / ((0) | 1)))
        v15 := GenS2{F0: genFn0(-91)}
        _ = fi14
        _ = v15
    }
    if v1 {
        fmt.Println(v4)
        var v16 uint64 = 66
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 uint64 = v16
                _ = fi18
                _ = v19
            }
            fmt.Println(v2, fi17)
            _ = fi17
        }
        _ = v16
    } else {
        v12 += (v12 % ((v12) | 1))
        switch -1 {
        case 128:
            for fi20 := 0; fi20 < 14; fi20++ {
                switch "foo bar" {
                case "hello":
                    if true {
                        var v21 uint = 681
                        _ = v21
                    }
                case "\n\t":
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                case "a":
                    var v22 uint8 = 255
                    _ = v22
                default:
                    v11 -= v11
                }
                _ = fi20
            }
        }
    }
    if (-1 != 1) {
        var v23 int = genFn0(-65)
        _ = v23
    } else {
        switch "foo bar" {
        case "hello":
            v2 -= uint32(v12)
        default:
            var v24 uint16 = 42
            _ = v24
        }
        v2 = (v3 + (v3 | v3))
    }
    v12 = (v12 & v12)
    var v25 uint32 = 773
    v10 = (1.5 * ((v10 + float64(-2.25)) - v10))
    var v26 int8 = 100
    if (v10 != 1.5) {
        v27 := GenS3{F0: 397, F1: (("foo bar" + (("\n\t" + "a") + "a")) + ("\n\t" + ""))}
        switch -1 {
        case -1:
            var v28 bool = (v2 != v2)
            _ = v28
        case -2147483648:
            if false {
                v5 = append(v5, "foo bar")
                v26 = (v26 * v26)
                _ = len(v5)
            }
        }
        _ = v27
    } else {
        if (2147483647 != 1) {
            var v29 int16 = int16(v12)
            if v1 {
                var v30 int8 = 59
                if (0 < 42) {
                    v12 = v12
                    v25 = 1
                }
                _ = v30
            } else {
                _ = len(v5)
            }
            var v31 uint8 = 1
            _ = v29
            _ = v31
        }
        var v32 uint16 = 42
        _ = v32
    }
    if v1 {
        switch -9223372036854775808 {
        case -2147483648:
            v5 = append(v5, "foo bar")
        }
    }
    v33 := GenS4{F0: (v25 + 2147483647), F1: (158 > 2), F2: v26}
    v4 = int64(v3)
    v11 -= float32(float64(0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
