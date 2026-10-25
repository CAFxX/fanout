package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    v1 := []int{64, 32768, 32768}
    var v2 float64 = float64(-2.25)
    var v3 int = 2147483647
    var v4 float64 = 3.14159
    v5 := []int{v3, v3}
    v6 := GenS1{F0: 50, F1: "", F2: ""}
    fmt.Println(v2)
    v2 += 0.1
    var v7 float32 = 0.1
    var v8 float32 = v7
    switch "" {
    case "foo bar":
        var v9 float64 = ((v2 + -2.25) / -2.25)
        _ = v9
    case "a":
        var v10 float64 = v2
        _ = v10
    case "":
        var v11 string = "\n\t"
        _ = v11
    default:
        v8 -= v8
    }
    var v12 uint8 = 2
    for fi13 := 0; fi13 < 3; fi13++ {
        switch "foo bar" {
        case "hello":
            if (!(9223372036854775807 == 667)) {
                v7 -= v8
                v3 += (v3 - (fi13 + (v3 & (-(fi13)))))
            } else {
                v3 += (v3 & (fi13 + (fi13 % ((v3) | 1))))
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 uint64 = 4294967295
                    v16 := GenS2{F0: 94, F1: 2, F2: 2}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 float64 = v4
                _ = v17
            }
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 42
            }
        default:
            if len(v1) > 0 {
                v1[0] = v3
            }
        }
        _ = fi13
    }
    if true {
        fmt.Println(v12)
        _ = len(v1)
    }
    v2 = v2
    fmt.Println(v8, v7, v8)
    v1 = append(v1, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
