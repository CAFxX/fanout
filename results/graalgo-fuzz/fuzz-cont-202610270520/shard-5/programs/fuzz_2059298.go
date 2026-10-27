package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 7; fi4++ {
        v1 = (v1 | v1)
        fmt.Println(v1)
        _ = fi4
    }
    var v5 float64 = (0.0 * (3.14159 + (0.1 - 3.14159)))
    var v6 uint = 277
    delete(v2, 2147483647)
    if (58 != 42) {
        switch -9223372036854775808 {
        case 128:
            if ((^(v1)) != (v1 / (((v1 & (v1 - 264))) | 1))) {
                var v7 int32 = -217
                var v8 int16 = 1
                _ = v7
                _ = v8
            } else {
                var v9 uint16 = 593
                _ = v9
            }
        default:
            if false {
                v6 -= v6
                var v10 int16 = 32767
                _ = v10
            }
        }
        v11 := GenS1{F0: 9223372036854775807, F1: (v1 * v1)}
        _ = v11
    }
    fmt.Println(v5, v5)
    var v12 string = ""
    fmt.Println(v5, v12)
    switch ("foo bar" + v12) {
    case "\n\t":
        var v13 float64 = float64((0.0 / 1.0))
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: 0, F1: (v6 ^ (v6 >> ((v6) & 63))), F2: (v6 << (((v6 * (v6 * 255))) & 63))}
        _ = v14
    }
    switch ((v12 + v12) + "a") {
    case "":
        v12 = (v12 + v12)
    }
    var v15 uint = v6
    var v16 bool = false
    switch 0 {
    case 1:
        v16 = ((42 == 1) && (42 >= 417))
    case 4294967295:
        fmt.Println(v5)
    default:
        if true {
            if v16 {
                v17 := []int{-118}
                v6 = 18446744073709551615
                _ = v17
            }
        }
    }
    switch 2147483647 {
    case 127:
        v18 := []int{0}
        _ = v18
    default:
        v6 = 782
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
