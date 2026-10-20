package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint32 = 313
    var v2 uint16 = 1
    var v3 float32 = (1.5 - 0.1)
    fmt.Println(v2)
    if false {
        if (!(!((v3 != v3) && (!(1 > 724))))) {
            v3 -= v3
        }
    } else {
        v4 := []int{2147483647, 9223372036854775807, 128, 9223372036854775807}
        v2 += (v2 + (v2 ^ 81))
        var v5 int64 = 976
        _ = v4
        _ = v5
    }
    var v6 float32 = 0.1
    fmt.Println(v2, v2, v6)
    v1 = (v1 - v1)
    v6 = (0.0 / 1.0)
    v7 := []int{128, -2147483648}
    v1 += (v1 * 2147483647)
    var v8 uint64 = 1
    _ = len(v7)
    if true {
        v9 := GenS1{F0: 65535, F1: 75}
        switch 1 {
        case -9223372036854775808:
            var v10 uint = 522
            _ = v10
        case 2147483647:
            var v11 string = "hello"
            _ = v11
        }
        var v12 int = 127
        _ = v12
        _ = v9
    }
    v13 := GenS1{F0: (v2 ^ (v2 + v2)), F1: 127}
    switch (("a" + "hello") + "foo bar") {
    case "":
        switch "" {
        case "\n\t":
            if len(v7) > 0 {
                v7[0] = -1
            }
        }
    case "hello":
        if len(v7) > 0 {
            v7[0] = 1
        }
    case "foo bar":
        var v14 float64 = 0.1
        _ = v14
    default:
        var v15 string = "a"
        _ = v15
    }
    v3 = float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
