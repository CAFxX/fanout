package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }

func main() {
    var v1 float64 = 0.0
    var v2 uint = 18446744073709551615
    var v3 int64 = -800
    var v4 uint64 = 288
    var v5 uint = (v2 * (v2 + v2))
    v2 = (v5 - v2)
    v1 -= v1
    var v6 uint = 1
    v7 := []int{}
    switch 4294967295 {
    case 42:
        if ("\n\t" != "") {
            v6 = (^(v5))
            fmt.Println(v4, v2)
        } else {
            switch "foo bar" {
            case "":
                var v8 uint64 = v4
                _ = v8
            case "\n\t":
                fmt.Println(v4)
            case "a":
                if (0 == -1) {
                    v9 := GenS1{F0: -365, F1: v6, F2: 188}
                    _ = v9
                } else {
                    var v10 int32 = 0
                    _ = v10
                }
            }
            var v11 int8 = 42
            _ = v11
        }
    case -290:
        v1 = (v1 / (v1 / (((v1 - (float64((0.1 + ((1.5 / 1.0) + float32((v1 + v1))))) * (((v1 + v1) + (0.0 - v1)) + ((v1 + (((v1 * float64(float32(1e10))) - -2.25) * 1.5)) + v1)))) * v1) * float64(1.5))))
    }
    fmt.Println(v4)
    var v12 float32 = 3.14159
    var v13 bool = true
    v14 := GenS2{F0: (v6 & 18446744073709551615), F1: 59, F2: -563}
    fmt.Println(v1)
    v1 -= ((v1 * 0.0) - (v1 - 3.14159))
    var v15 uint16 = 42
    if (!v13) {
        switch 342 {
        case 127:
            v1 -= -2.25
        case -360:
            if false {
                fmt.Println(v1)
                v16 := GenS1{F0: 0, F1: 0, F2: 42}
                _ = v16
            } else {
                v17 := GenS1{F0: 1, F1: (v2 % (((v6 >> ((v2) & 63))) | 1)), F2: 0}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v17
            }
        case -9223372036854775808:
            var v18 string = "hello"
            _ = v18
        default:
            v4 = v4
        }
        fmt.Println(v5, v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
