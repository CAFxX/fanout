package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    v1 := []string{}
    var v2 uint = 2
    var v3 uint64 = 18446744073709551615
    var v4 int16 = -32768
    var v5 bool = (2147483647 >= 258)
    if (42 <= -1) {
        v5 = (3.14159 != -2.25)
        v6 := GenS1{F0: 65535, F1: uint16(v3), F2: 42}
        switch "a" {
        case "hello":
            v1 = append(v1, "a")
        default:
            var v7 float64 = ((float64(1e10) / 0.1) / (-2.25 + 1.5))
            _ = v7
        }
        _ = v6
    }
    v5 = true
    var v8 int8 = -50
    if (0 > v2) {
        v5 = true
    } else {
        var v9 bool = (810 <= -1)
        _ = v9
    }
    v2 -= v2
    var v10 int8 = v8
    v11 := GenS1{F0: 2, F1: 65535, F2: 243}
    fmt.Println(v5, v8, v5)
    if v5 {
        if v5 {
            v1 = append(v1, ("" + "a"))
            v4 = (v4 | (v4 - v4))
        }
        v12 := GenS1{F0: 504, F1: 424, F2: 102}
        _ = v12
    }
    if ((v3 >= 0) && (((!v5) && v5) || ((v5 && (2147483647 < 65535)) && false))) {
        v1 = append(v1, "\n\t")
    }
    v13 := GenS2{F0: 128, F1: (v2 << ((852) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
