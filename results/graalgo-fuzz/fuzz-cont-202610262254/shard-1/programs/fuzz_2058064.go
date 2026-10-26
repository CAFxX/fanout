package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 int64) uint64 {
    p0 -= (p0 + p0)
    var v5 int16 = -32768
    if false {
        if (true || (false || (1 >= 4294967295))) {
            switch "\n\t" {
            case "foo bar":
                var v6 float32 = ((-2.25 * -2.25) / float32(3.14159))
                _ = v6
            default:
                for fi7 := 0; fi7 < 10; fi7++ {
                    for fi8 := 0; fi8 < 2; fi8++ {
                        switch ("foo bar" + ("\n\t" + "foo bar")) {
                        case "a":
                            var v9 float64 = (-2.25 + 0.1)
                            _ = v9
                        case "hello":
                            v10 := []int{}
                            _ = v10
                        default:
                            var v11 float64 = 1e10
                            _ = v11
                        }
                        p0 -= -9223372036854775808
                        _ = fi8
                    }
                    if true {
                        if ((!((!false) && false)) || (true && (42 != 1))) {
                            v12 := []int{fi7, -1, (fi7 | fi7)}
                            _ = v12
                        } else {
                            v13 := []int{fi7, fi7, fi7, fi7}
                            var v14 int16 = -1
                            _ = v13
                            _ = v14
                        }
                        if (p0 >= p0) {
                            var v15 int16 = v5
                            v16 := []int{(fi7 + fi7), (fi7 & (-(fi7)))}
                            _ = v15
                            _ = v16
                        }
                    } else {
                        v5 -= v5
                    }
                    var v17 bool = false
                    _ = fi7
                    _ = v17
                }
            }
            switch "" {
            case "":
                p0 = p0
            case "\n\t":
                v18 := []int{}
                _ = v18
            case "a":
                switch "\n\t" {
                case "hello":
                    v19 := []int{778, 0, -2147483648, 1}
                    _ = v19
                default:
                    fmt.Println(v5)
                }
            default:
                p0 = -241
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                p0 -= p0
                _ = fi20
            }
        }
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int8, p1 int) int16 {
    v21 := []int{}
    p0 = (^(p0))
    v22 := GenS1{F0: p1, F1: 42}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 42
    var v3 uint8 = 255
    var v4 int8 = 1
    var v23 float32 = 0.0
    v2 = 0
    v24 := GenS1{F0: 32768, F1: v2}
    v4 += (v4 / ((100) | 1))
    var v25 int8 = v4
    for fi26 := 0; fi26 < 11; fi26++ {
        v1[0] = true
        fmt.Println(v3)
        v2 -= 70
        _ = fi26
    }
    v27 := []int{}
    v28 := v1[0]
    v4 = (v4 % (((v4 << ((4294967295) & 63))) | 1))
    for fi29 := 0; fi29 < 5; fi29++ {
        v28 = ((1e10 <= 1e10) || v28)
        v30 := GenS2{F0: 42, F1: "foo bar"}
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
