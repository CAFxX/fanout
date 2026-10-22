package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 int64, p1 uint32) uint64 {
    switch ("" + "foo bar") {
    case "a":
        p0 = (p0 >> ((606) & 63))
    case "":
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 = (^(p0))
            _ = fi6
        }
    default:
        p0 = p0
    }
    var v7 float64 = float64(0.1)
    switch -832 {
    case -21:
        fmt.Println(v7, v7)
    case 9223372036854775807:
        var v8 uint32 = (p1 - (p1 >> ((1) & 63)))
        _ = v8
    case -1:
        switch "hello" {
        case "\n\t":
            for fi9 := 0; fi9 < 16; fi9++ {
                if (0 != (fi9 / (((fi9 ^ fi9)) | 1))) {
                    p0 = (-(p0))
                    v10 := []int{}
                    _ = v10
                }
                _ = fi9
            }
        case "hello":
            var v11 uint = 2
            _ = v11
        case "":
            v12 := GenS1{F0: 32768, F1: 42}
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            if (true || false) {
                p1 += p1
            }
            v14 := []int{fi13, fi13, (-(fi13))}
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 460
}
func genFn1() uint32 {
    var v15 int64 = -974
    v15 += -1
    if ("hello" >= ("a" + ("a" + ("hello" + "foo bar")))) {
        v16 := GenS1{F0: -9223372036854775808, F1: 20}
        v17 := []int{2147483647, 128}
        _ = v16
        _ = v17
    } else {
        v15 = (v15 - (v15 + (v15 / ((v15) | 1))))
    }
    if false {
        fmt.Println(v15)
        if true {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{(fi18 >> ((354) & 63)), fi18, (fi18 - (fi18 / ((fi18) | 1))), (fi18 | 4294967295)}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                fmt.Println(fi18, fi18)
                _ = fi18
                _ = v19
            }
        } else {
            var v20 int32 = 839
            _ = v20
        }
    } else {
        v15 -= (v15 ^ (v15 | 838))
        for fi21 := 0; fi21 < 13; fi21++ {
            if true {
                v15 += v15
                v15 -= -946
                v15 = v15
            } else {
                var v22 int = (fi21 & (-(fi21)))
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 int64 = (v15 % ((v15) | 1))
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
            _ = fi21
        }
    }
    _ = v15
    return 2147483647
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = (float32(-2.25) * (((1e10 / 1.5) + (-2.25 / 1.0)) / -2.25))
    var v3 float32 = 1.5
    v4 := []uint32{959, 394, 1, 1}
    v5 := []int{523, -267, -2147483648, 127}
    switch "foo bar" {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = 1
        }
    case "\n\t":
        v2 = 1.5
    case "hello":
        fmt.Println(v3, v3)
    default:
        if len(v5) > 0 {
            v5[0] = -371
        }
    }
    if (true || true) {
        var v25 float32 = 0.1
        v1 -= (-(v1))
        _ = len(v4)
        _ = v25
    } else {
        var v26 string = ""
        _ = v26
    }
    switch ("\n\t" + ("a" + "a")) {
    case "":
        for fi27 := 0; fi27 < 2; fi27++ {
            v2 -= v2
            v1 = (v1 | (v1 / ((v1) | 1)))
            var v28 uint32 = 0
            _ = fi27
            _ = v28
        }
    default:
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
    }
    var v29 float64 = float64(1.5)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v29 += float64((1.5 + (1e10 / 1.0)))
        v30 := GenS1{F0: 32768, F1: 65535}
        _ = v30
    }
    if false {
        var v31 string = "foo bar"
        _ = v31
    } else {
        fmt.Println(v3, v1, v2)
        v1 = (v1 | v1)
    }
    v4 = append(v4, 42)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    v29 += (((1.5 * 1.5) / v29) / (0.0 * v29))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
