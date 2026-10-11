package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint = 997
    v2 := []uint32{751, 4294967295, 623, 2}
    var v3 bool = false
    var v4 int64 = -495
    var v5 float64 = (0.1 - 1.5)
    fmt.Println(v4)
    switch -32 {
    case 42:
        switch "hello" {
        case "\n\t":
            if false {
                switch 4294967295 {
                case 258:
                    var v6 bool = false
                    _ = v6
                case 0:
                    if (187 == 2) {
                        if ((("" + "hello") + "a") > "foo bar") {
                            v7 := GenS1{F0: 18446744073709551615, F1: (-165 <= -591)}
                            var v8 int = -325
                            v8 -= v8
                            _ = v7
                            _ = v8
                        }
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = len(v2)
                _ = len(v2)
            }
        case "foo bar":
            var v9 int8 = 127
            _ = v9
        }
    case 127:
        v10 := GenS2{F0: v3, F1: -1}
        _ = v10
    case 0:
        v11 := GenS1{F0: 4294967295, F1: (uint8(v1) >= 255)}
        _ = v11
    }
    _ = len(v2)
    fmt.Println(v4, v5)
    if false {
        var v12 uint8 = 122
        _ = v12
    } else {
        var v13 uint64 = 1
        if (!v3) {
            v5 += v5
        }
        _ = v13
    }
    var v14 int16 = 32767
    fmt.Println(v3, v5, v1)
    var v15 int = 32767
    if (!v3) {
        fmt.Println(v1, v1, v14)
        if (2 >= 1) {
            v14 = 747
        }
    } else {
        if ("foo bar" > "\n\t") {
            v16 := GenS2{F0: ((-2.25 - (-2.25 / 1.0)) == 1.5), F1: -1}
            v4 = (v4 - v4)
            var v17 int32 = -537
            _ = v16
            _ = v17
        }
        var v18 float64 = v5
        _ = v18
    }
    var v19 float64 = v5
    var v20 int = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
