package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 int16 = 286
    var v2 int64 = 0
    var v3 uint64 = 4294967295
    var v4 int64 = -588
    if false {
        v5 := []int{}
        _ = v5
    } else {
        var v6 int8 = int8(v2)
        v7 := GenS1{F0: 127}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 13}
    if true {
        var v9 float64 = 0.1
        _ = v9
    }
    v10 := []int{-100, -2147483648, 4294967295}
    if (!((v1 / (((v1 * v1)) | 1)) != v1)) {
        if len(v10) > 0 {
            v10[0] = -536
        }
    } else {
        if (!(true || true)) {
            v11 := GenS2{F0: 9223372036854775807, F1: 472}
            v12 := GenS3{F0: 127, F1: ("a" + ("\n\t" + "a"))}
            v3 = v3
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: v4, F1: 524}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS4{F0: 1}
        switch (fi14 << ((42) & 63)) {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 float32 = 0.0
                v10 = append(v10, int(v2))
                v3 = (v3 % ((90) | 1))
                switch (("" + "a") + "hello") {
                case "hello":
                    _ = len(v10)
                case "\n\t":
                    var v18 string = ""
                    _ = v18
                default:
                    v3 = 4294967295
                    v19 := GenS4{F0: 0}
                    var v20 int64 = (v4 + v4)
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = v3
                _ = fi16
                _ = v17
                _ = v21
            }
        case 302:
            v1 = (v1 % ((v1) | 1))
        case 0:
            v4 = v2
        default:
            fmt.Println(v4, v3, v2)
        }
        var v22 float64 = 3.14159
        _ = fi14
        _ = v15
        _ = v22
    }
    v23 := GenS4{F0: 18446744073709551615}
    var v24 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
