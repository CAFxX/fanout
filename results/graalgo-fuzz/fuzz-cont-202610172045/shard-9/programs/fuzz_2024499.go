package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 int = -2147483648
    var v2 uint = 198
    v3 := make(map[string]uint32)
    var v4 int32 = -838
    var v5 float64 = (-2.25 * 0.1)
    var v6 float64 = -2.25
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 5; fi8++ {
        v5 = (v5 - 0.0)
        _ = fi8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch ("a" + (("hello" + ("" + "\n\t")) + ("hello" + "a"))) {
        case "hello":
            var v10 bool = false
            _ = v10
        case "a":
            var v11 int16 = -1
            _ = v11
        default:
            v6 -= ((0.1 / v6) / float64(1e10))
        }
        if false {
            v12 := GenS1{F0: ("" + "foo bar"), F1: 25, F2: -279}
            _ = v12
        } else {
            fmt.Println(v1, fi9)
            v2 += (v2 * v2)
        }
        v2 -= v2
        _ = fi9
    }
    v1 = v1
    var v13 uint16 = (v7 << ((v2) & 63))
    var v14 int64 = -9223372036854775808
    fmt.Println(v7, v4)
    v2 = (v2 / ((v2) | 1))
    v7 = (v13 / ((uint16(v2)) | 1))
    v15 := GenS2{F0: v1}
    if (!true) {
        v16 := v3[("\n\t" + "foo bar")]
        _ = v16
    } else {
        v6 -= v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
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
