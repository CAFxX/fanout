package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }
type GenS2 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 int16 = 32767
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    var v6 uint = 0
    v6 = v6
    v7 := []int{32767, 0}
    v4 = true
    v8 := GenS1{F0: 0, F1: v4, F2: (v6 / ((255) | 1))}
    switch -9223372036854775808 {
    case -2147483648:
        v9 := GenS1{F0: 433, F1: (v3 == -32768), F2: v6}
        _ = v9
    }
    switch 2147483647 {
    case -1:
        _ = len(v7)
    case 42:
        v6 = (v6 >> ((v6) & 63))
    default:
        v6 = (v6 | 65535)
    }
    switch 2147483647 {
    case 9223372036854775807:
        v7 = append(v7, 128)
    case 32768:
        fmt.Println(v3, v3, v3)
    case -668:
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: -1, F1: (244 >= 906), F2: v6}
            _ = v10
        case "\n\t":
            v3 += (v3 * (^(v3)))
        case "hello":
            v6 -= v6
        default:
            v11 := GenS2{F0: (!(false || true)), F1: 128}
            _ = v11
        }
    default:
        switch ("hello" + "foo bar") {
        case "":
            if len(v7) > 0 {
                v7[0] = -1
            }
        }
    }
    switch ("" + "hello") {
    case "":
        v12 := GenS2{F0: (((v6 >> (((v6 >> ((0) & 63))) & 63)) > v6) || v4), F1: int(v3)}
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
