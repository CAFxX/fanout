package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 float64 = (1.5 / 1.5)
    v3 := []int{-1, -328, -2147483648, -9223372036854775808}
    v4 := []int{}
    var v5 int64 = 1
    var v6 int16 = 0
    v1 = (v1 + 2)
    var v7 uint32 = 622
    v6 = 42
    if (true || false) {
        var v8 bool = (v6 == v6)
        v1 -= v1
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int8 = 0
            if len(v4) > 0 {
                v4[0] = 0
            }
            v7 = 2
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v11 uint32 = v7
    switch ("foo bar" + "a") {
    case "foo bar":
        fmt.Println(v6, v1, v11)
    case "\n\t":
        if true {
            var v12 int8 = -121
            v13 := GenS1{F0: 255, F1: false, F2: 2}
            _ = v12
            _ = v13
        } else {
            if len(v3) > 0 {
                v3[0] = 311
            }
        }
    case "":
        v1 -= 42
    }
    switch 1 {
    case 1:
        var v14 string = "a"
        _ = v14
    case 246:
        fmt.Println(v7, v5)
    default:
        if false {
            var v15 uint32 = (v11 - 2)
            fmt.Println(v5)
            var v16 string = "foo bar"
            _ = v15
            _ = v16
        }
    }
    v3 = append(v3, -2147483648)
    var v17 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
