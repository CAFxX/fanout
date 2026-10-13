package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := []uint32{497, 4294967295}
    var v2 uint8 = 127
    var v3 uint64 = 782
    var v4 uint8 = v2
    v5 := GenS1{F0: 1, F1: 471}
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = 936
        _ = fi6
        _ = v7
    }
    var v8 int8 = -91
    switch 953 {
    case 32768:
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int = fi9
            var v11 bool = false
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        var v12 uint = 822
        _ = v12
    }
    v4 = (v2 - (^(v4)))
    var v13 int32 = -2147483648
    v8 -= (v8 / ((v8) | 1))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int32 = (v13 % ((v13) | 1))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: v8, F1: 257}
    if false {
        v8 = (v8 % ((int8(v3)) | 1))
        v8 = (^(v8))
    }
    switch -1 {
    case -916:
        _ = len(v1)
    case -676:
        fmt.Println(v13, v13)
    case 42:
        for fi17 := 0; fi17 < 19; fi17++ {
            if (((v4 / ((42) | 1)) == v2) || true) {
                v4 = (v4 / ((v4) | 1))
                v4 -= (v4 ^ (v4 * (v4 * 125)))
            } else {
                fmt.Println(v4, v13)
                if true {
                    v8 = (v8 / ((100) | 1))
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v3 -= (v3 % ((0) | 1))
            }
            if (42 >= v8) {
                _ = len(v1)
            } else {
                var v18 uint16 = 2
                _ = v18
            }
            var v19 int64 = -306
            _ = fi17
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: "foo bar"}
        _ = fi20
        _ = v21
    }
    if false {
        switch ("a" + "foo bar") {
        case "\n\t":
            v4 -= (v4 - (v4 * v4))
        case "a":
            v4 += v2
        default:
            fmt.Println(v13, v3)
        }
    }
    var v22 int32 = (-(v13))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
