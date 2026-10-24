package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint }
func genFn0() int32 {
    var v6 int16 = 1000
    v7 := GenS1{F0: 1}
    v6 = -32768
    _ = v6
    _ = v7
    return 915
}
func genFn1(p0 int, p1 uint16, p2 int32) uint32 {
    v8 := []int{(p0 >> ((1) & 63))}
    fmt.Println(p0, p1)
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 917
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 0.1
    var v3 float64 = (v2 / -2.25)
    var v4 uint32 = 856
    var v5 int64 = 1
    v9 := []int{-1, -468}
    var v10 int32 = (v1 / ((genFn0()) | 1))
    v5 -= v5
    switch "foo bar" {
    case "\n\t":
        v2 = v3
    default:
        var v11 int8 = 100
        _ = v11
    }
    fmt.Println(v1)
    switch 32767 {
    case 127:
        if false {
            switch ("foo bar" + "") {
            case "hello":
                if (!(false && ((true || true) || (false || true)))) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            case "\n\t":
                v9 = append(v9, 1)
            default:
                switch ("a" + ("hello" + "hello")) {
                case "":
                    _ = len(v9)
                default:
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 uint32 = v4
                        _ = fi12
                        _ = v13
                    }
                }
            }
            v14 := GenS1{F0: 699}
            if (928 != 2) {
                if ((((-2.25 * (v3 * v3)) + v2) - (0.1 / v3)) > 3.14159) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                v15 := GenS1{F0: 855}
                var v16 uint32 = 2147483648
                _ = v15
                _ = v16
            }
            _ = v14
        }
    case -40:
        v2 += v2
    case 9223372036854775807:
        for fi17 := 0; fi17 < 19; fi17++ {
            v5 = (v5 | v5)
            if (true || true) {
                v4 += (v4 & v4)
            } else {
                var v18 uint8 = 128
                if true {
                    v19 := GenS2{F0: -1, F1: 1}
                    v1 -= (v10 - (v1 & genFn0()))
                    if false {
                        v20 := GenS2{F0: 1, F1: 42}
                        v10 += v1
                        _ = v20
                    }
                    _ = v19
                }
                _ = v18
            }
            _ = fi17
        }
    }
    if true {
        v5 = 9223372036854775807
    }
    v21 := GenS2{F0: 42, F1: 210}
    if true {
        if true {
            for fi22 := 0; fi22 < 13; fi22++ {
                fmt.Println(v1, v10, v4)
                var v23 uint32 = (v4 + v4)
                _ = fi22
                _ = v23
            }
            if len(v9) > 0 {
                v9[0] = -693
            }
        }
        var v24 string = "\n\t"
        _ = v24
    }
    v4 = v4
    v25 := GenS3{F0: 42}
    v4 = (v4 % ((v4) | 1))
    var v26 int32 = (v10 & (v10 / (((v10 << ((4294967295) & 63))) | 1)))
    if (!(false && true)) {
        for fi27 := 0; fi27 < 10; fi27++ {
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(fi27, v3, fi27)
                var v29 int32 = genFn0()
                if false {
                    v4 += (v4 + (^(v4)))
                } else {
                    var v30 uint32 = (^(v4))
                    if len(v9) > 0 {
                        v9[0] = 32767
                    }
                    _ = v30
                }
                if true {
                    var v31 uint64 = 984
                    v26 = 897
                    _ = v31
                }
                v5 -= v5
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                fmt.Println(v10, v10)
                var v32 int16 = -905
                _ = fi28
                _ = v29
                _ = v32
            }
            v33 := GenS3{F0: 644}
            fmt.Println(v10)
            _ = fi27
            _ = v33
        }
        if false {
            v34 := GenS2{F0: 129, F1: 42}
            _ = v34
        }
    }
    v5 = v5
    fmt.Println(v5, v26, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
