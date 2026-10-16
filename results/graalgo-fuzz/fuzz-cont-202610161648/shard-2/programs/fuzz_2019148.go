package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 int }

func main() {
    var v1 float64 = -2.25
    var v2 string = ("a" + "")
    v3 := []string{"foo bar"}
    var v4 uint32 = 101
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    v1 = v1
    v2 = (v2 + "")
    var v7 string = (v2 + "\n\t")
    switch "a" {
    case "hello":
        v8 := GenS1{F0: (v4 >> ((1) & 63))}
        _ = v8
    case "\n\t":
        fmt.Println(v4)
    }
    fmt.Println(v1, v4, v1)
    var v9 int32 = -463
    var v10 float32 = (-2.25 + (float32(3.14159) * float32(v1)))
    switch 127 {
    case -546:
        var v11 int32 = (^(v9))
        _ = v11
    case -507:
        if (992 < 0) {
            if len(v3) > 0 {
                v3[0] = (v2 + ((((v7 + "foo bar") + "hello") + v2) + ""))
            }
            switch v2 {
            case "":
                v5[((v7 + ("hello" + v2)) + v7)] = (false || (!true))
            case "hello":
                v12 := GenS2{F0: 9223372036854775807, F1: v9, F2: 32768}
                _ = v12
            case "foo bar":
                v3 = append(v3, "hello")
            default:
                v9 -= v9
            }
        }
    default:
        var v13 uint8 = 127
        _ = v13
    }
    fmt.Println(v2)
    var v14 int64 = 716
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
