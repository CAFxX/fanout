package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int8 = -112
    v2 := []int{-837, -9223372036854775808, 4294967295, -1}
    var v3 uint8 = 1
    if (!false) {
        v2 = append(v2, 32767)
        v4 := GenS1{F0: (!false), F1: 4294967295}
        _ = v4
    } else {
        switch -9223372036854775808 {
        case 0:
            fmt.Println(v1, v1)
        }
        _ = len(v2)
    }
    v5 := GenS1{F0: false, F1: 42}
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 += (v1 ^ (v1 * v1))
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, v3, fi6)
        case "foo bar":
            var v7 string = "\n\t"
            _ = v7
        }
        var v8 uint8 = (v3 - (v3 >> ((1) & 63)))
        _ = fi6
        _ = v8
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            var v9 uint8 = v3
            _ = v9
        }
        var v10 uint8 = v3
        _ = v10
    } else {
        switch "" {
        case "a":
            if true {
                var v11 int16 = -1
                _ = v11
            }
        case "":
            var v12 float32 = 3.14159
            _ = v12
        default:
            fmt.Println(v1)
        }
    }
    var v13 uint16 = 65535
    v14 := GenS1{F0: true, F1: uint64(v13)}
    var v15 float64 = 3.14159
    var v16 int32 = -2147483648
    v3 -= v3
    v2 = append(v2, -9223372036854775808)
    if len(v2) > 0 {
        v2[0] = 127
    }
    fmt.Println(v13)
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 float64 = (1e10 - 0.0)
        v15 = float64((0.1 * float32(((v15 - v18) / v18))))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v1, fi19, fi19)
        _ = fi19
    }
    var v20 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
