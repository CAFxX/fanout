package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 uint, p1 uint16) int16 {
    p0 = p0
    if true {
        var v5 int32 = -1
        switch "hello" {
        case "":
            p1 = 2
        case "foo bar":
            p1 -= (p1 % (((p1 ^ 365)) | 1))
        default:
            var v6 uint32 = 42
            _ = v6
        }
        var v7 int = -915
        _ = v5
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{int(p0), (^(fi8)), fi8}
        v10 := GenS1{F0: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        var v11 uint64 = 243
        p0 = (p0 << (((p0 * 863)) & 63))
        _ = v11
    } else {
        switch "\n\t" {
        case "a":
            var v12 int16 = -852
            _ = v12
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint) uint64 {
    var v13 uint16 = 1
    fmt.Println(v13, v13)
    v14 := []int{1}
    v14 = append(v14, -9223372036854775808)
    _ = p0
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    var v1 uint64 = 981
    var v2 uint32 = 42
    var v3 int64 = -564
    v4 := []string{"\n\t", "", "\n\t", "hello"}
    v3 -= v3
    v2 = (v2 | (v2 - 4294967295))
    v1 += v1
    fmt.Println(v1, v1)
    if false {
        if true {
            for fi15 := 0; fi15 < 8; fi15++ {
                _ = len(v4)
                v1 += (v1 * v1)
                _ = fi15
            }
            v1 -= genFn1(513)
        } else {
            switch -863 {
            case -424:
                v16 := GenS2{F0: 128, F1: 849, F2: 812}
                _ = v16
            }
            switch "hello" {
            case "":
                if false {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1, v1, v2)
                    case "":
                        for fi17 := 0; fi17 < 4; fi17++ {
                            for fi18 := 0; fi18 < 15; fi18++ {
                                fmt.Println(v1)
                                _ = fi18
                            }
                            v19 := GenS2{F0: fi17, F1: 2, F2: fi17}
                            _ = fi17
                            _ = v19
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = ("" + ("\n\t" + (("a" + "foo bar") + "foo bar")))
                    }
                } else {
                    var v20 int = -903
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v21 uint32 = (^(v2))
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                if true {
                    fmt.Println(v1, v1, v3)
                    var v22 uint32 = 0
                    _ = v22
                } else {
                    v4 = append(v4, ("" + (("a" + "foo bar") + "hello")))
                }
            case "a":
                fmt.Println(v3, v1, v3)
            default:
                v3 -= (v3 % ((-1) | 1))
            }
        }
        var v23 int64 = (v3 / ((v3) | 1))
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        _ = len(v4)
        var v25 float64 = 1e10
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        for fi26 := 0; fi26 < 4; fi26++ {
            v4 = append(v4, ("\n\t" + "a"))
            var v27 float64 = (0.1 / (0.1 / ((float64(1.5) / 0.1) / 0.1)))
            _ = fi26
            _ = v27
        }
    case "hello":
        v2 += v2
    default:
        switch 128 {
        case -924:
            fmt.Println(v1, v2)
        case 394:
            if ((!false) || false) {
                v3 = 9223372036854775807
            }
        case 971:
            switch ("" + (("" + "\n\t") + "foo bar")) {
            case "":
                var v28 float64 = -2.25
                _ = v28
            case "foo bar":
                var v29 int64 = v3
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
