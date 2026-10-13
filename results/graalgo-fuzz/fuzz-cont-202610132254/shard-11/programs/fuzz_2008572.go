package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 bool }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }
type GenS6 struct { F0 bool }
func genFn0(p0 int8, p1 bool, p2 int16) int64 {
    switch -9223372036854775808 {
    case -1:
        var v4 float32 = 1.5
        _ = v4
    case -962:
        p2 -= (p2 + p2)
    case -882:
        v5 := GenS1{F0: 255}
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        if (true && (65535 != 0)) {
            p2 = (p2 % ((0) | 1))
            if p1 {
                switch "hello" {
                case "hello":
                    var v7 uint16 = 0
                    _ = v7
                }
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 549
}

func main() {
    var v1 int16 = -1
    var v2 float32 = 0.0
    v3 := []string{}
    fmt.Println(v1)
    v8 := GenS1{F0: 4294967295}
    if (!(!(("" == ("a" + "a")) || (1.5 < v2)))) {
        v9 := GenS2{F0: (285 != 1)}
        _ = v9
    }
    v2 = (v2 - v2)
    switch "\n\t" {
    case "\n\t":
        v3 = append(v3, "\n\t")
    default:
        switch ("hello" + "foo bar") {
        case "a":
            switch 1 {
            case 127:
                for fi10 := 0; fi10 < 10; fi10++ {
                    _ = len(v3)
                    v11 := GenS2{F0: false}
                    _ = fi10
                    _ = v11
                }
            case 436:
                v12 := GenS3{F0: 2147483647}
                _ = v12
            }
        case "\n\t":
            if (false && (32 >= -1)) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int16 = (v1 / ((v1) | 1))
                    if false {
                        var v15 float64 = 3.14159
                        v15 = float64((v2 + (v2 * (v2 * v2))))
                        var v16 uint8 = 42
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS4{F0: 2, F1: 42, F2: (!true)}
                        fmt.Println(v1, fi13)
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v1, v2)
            } else {
                v18 := GenS5{F0: (false && false), F1: v1, F2: 1}
                _ = v18
            }
        case "":
            var v19 uint8 = 42
            _ = v19
        default:
            var v20 float64 = 0.1
            _ = v20
        }
    }
    switch 4294967295 {
    case 42:
        var v21 bool = true
        _ = v21
    }
    v1 -= v1
    switch "foo bar" {
    case "":
        fmt.Println(v2, v2)
    case "foo bar":
        fmt.Println(v2)
    case "a":
        if (!(!(-1 >= -9223372036854775808))) {
            v22 := GenS2{F0: (!((!((!(1 != 18446744073709551615)) || (false && (0 <= 903)))) || (660 > 1)))}
            switch -765 {
            case 32768:
                var v23 int32 = -2147483648
                _ = v23
            case 42:
                v1 = v1
            case 134:
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 int16 = (v1 << ((2) & 63))
                    _ = fi24
                    _ = v25
                }
            default:
                var v26 int16 = v1
                _ = v26
            }
            if (!false) {
                var v27 float32 = (v2 - ((1.5 / v2) * (v2 - v2)))
                _ = v27
            } else {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                fmt.Println(v2, v2)
            }
            _ = v22
        }
    default:
        var v28 float32 = 1e10
        _ = v28
    }
    var v29 bool = true
    var v30 uint64 = 142
    v31 := GenS6{F0: (-1 <= -70)}
    v3 = append(v3, (("foo bar" + "hello") + "foo bar"))
    fmt.Println(v30)
    var v32 int32 = 362
    var v33 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v8
}
