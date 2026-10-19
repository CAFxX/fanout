package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }

func main() {
    var v1 int16 = 42
    var v2 uint16 = 42
    var v3 int64 = -9223372036854775808
    var v4 uint64 = 0
    v5 := GenS1{F0: 60}
    if false {
        if true {
            switch ("a" + "\n\t") {
            case "":
                var v6 float64 = ((3.14159 - ((float64((1.5 - 1.5)) * (-2.25 / 1.0)) + -2.25)) + 0.1)
                _ = v6
            case "a":
                v4 += v4
            case "hello":
                v2 = 865
            }
            v7 := GenS1{F0: 127}
            v1 -= 42
            _ = v7
        } else {
            if true {
                var v8 uint16 = 65535
                fmt.Println(v1, v8, v3)
                _ = v8
            } else {
                fmt.Println(v3, v3)
                var v9 int = 128
                _ = v9
            }
            v3 += (v3 + (v3 >> ((65535) & 63)))
        }
        fmt.Println(v2, v4, v3)
        fmt.Println(v1)
    }
    fmt.Println(v2, v1, v4)
    v10 := GenS2{F0: 4294967295, F1: 2147483647, F2: v2}
    if (-41 != 100) {
        var v11 int = -9223372036854775808
        _ = v11
    } else {
        var v12 int8 = -81
        var v13 uint8 = 23
        _ = v12
        _ = v13
    }
    v14 := []int{}
    var v15 int8 = -117
    v1 -= (v1 - (-(v1)))
    var v16 int16 = (v1 * v1)
    v17 := GenS2{F0: 351, F1: -1, F2: 52}
    var v18 int8 = (^(v15))
    v19 := GenS2{F0: 2147483648, F1: -1, F2: (v2 / (((v2 % ((v2) | 1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
