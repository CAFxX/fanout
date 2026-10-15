package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 int) uint16 {
    switch -2147483648 {
    case 128:
        if true {
            fmt.Println(p0, p0)
            if false {
                v5 := GenS1{F0: 1}
                p1 -= (-(p1))
                fmt.Println(p0, p1)
                _ = v5
            } else {
                p1 += -1
                var v6 uint32 = 2147483647
                _ = v6
            }
            if false {
                var v7 int64 = 9223372036854775807
                _ = v7
            }
        } else {
            p0 += (p0 * p0)
            var v8 uint = uint(p0)
            _ = v8
        }
    case 42:
        p1 = 42
    case 1:
        p0 = (p0 | p0)
    }
    fmt.Println(p0)
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = -61
    var v2 uint32 = 4294967295
    var v3 string = ("foo bar" + (("\n\t" + "") + ("" + "\n\t")))
    var v4 int8 = (v1 ^ v1)
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{int(v2), fi11, (^(fi11))}
        v12 = append(v12, (fi11 | fi11))
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 1
    var v14 uint16 = 255
    switch "a" {
    case "hello":
        var v15 float32 = -2.25
        _ = v15
    case "a":
        switch v3 {
        case "foo bar":
            var v16 int8 = 54
            _ = v16
        case "":
            fmt.Println(v3)
        case "a":
            for fi17 := 0; fi17 < 15; fi17++ {
                v18 := []int{(fi17 / (((^(fi17))) | 1)), 9223372036854775807, (fi17 >> ((259) & 63))}
                v19 := GenS2{F0: -1}
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            v1 = -125
        }
    case "":
        if (!true) {
            fmt.Println(v2, v2, v13)
            switch ("" + v3) {
            case "foo bar":
                v10 = -2.25
            case "\n\t":
                var v20 uint16 = v14
                _ = v20
            default:
                switch 128 {
                case 4294967295:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 int16 = 32767
                        v23 := []int{-904, (fi21 / ((fi21) | 1))}
                        var v24 string = v3
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                default:
                    v1 -= v4
                }
            }
            v10 = (v10 / 1.5)
        } else {
            v25 := []int{}
            _ = v25
        }
    default:
        v26 := GenS2{F0: 2147483647}
        _ = v26
    }
    v10 += 0.0
    var v27 float64 = (((v10 * v10) - v10) * float64(-2.25))
    v27 += (float64(3.14159) * 3.14159)
    v28 := []int{4294967295, -2147483648, -2147483648, 32767}
    var v29 int32 = -775
    v28 = append(v28, 1)
    v30 := GenS1{F0: (v4 << ((0) & 63))}
    v31 := GenS3{F0: false, F1: 42}
    switch 0 {
    case 0:
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := GenS1{F0: (-(v4))}
            var v34 float64 = v10
            _ = fi32
            _ = v33
            _ = v34
        }
    case -1:
        for fi35 := 0; fi35 < 15; fi35++ {
            if true {
                v2 += 1
            }
            fmt.Println(fi35, v3, v10)
            _ = fi35
        }
    default:
        v4 -= v4
    }
    switch 954 {
    case -2147483648:
        v2 += v2
    case 32767:
        if (0 != 609) {
            for fi36 := 0; fi36 < 13; fi36++ {
                v2 = 42
                if (0.0 <= 0.1) {
                    var v37 float64 = ((v27 + v10) + 0.0)
                    _ = v37
                }
                _ = fi36
            }
        } else {
            var v38 string = "\n\t"
            _ = v38
        }
    case 127:
        fmt.Println(v3)
    default:
        if (104 >= -1) {
            if false {
                switch "\n\t" {
                case "":
                    var v39 float32 = 1e10
                    _ = v39
                default:
                    var v40 int32 = v29
                    _ = v40
                }
                var v41 uint64 = 4294967295
                _ = v41
            }
        }
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        var v43 int64 = -1
        _ = fi42
        _ = v43
    }
    v13 = (v13 % (((v13 + uint8(v2))) | 1))
    v28 = append(v28, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
