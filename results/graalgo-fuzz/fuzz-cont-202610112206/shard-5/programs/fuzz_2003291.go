package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    var v4 int8 = 43
    var v5 uint = 412
    v6 := []string{}
    v7 := GenS1{F0: 2147483647, F1: v5}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 = (v4 + v4)
    if false {
        v4 = (^(v4))
        v2 = v2
        var v8 uint64 = 2
        _ = v8
    }
    v9 := v1["foo bar"]
    var v10 float32 = 0.1
    v11 := GenS1{F0: 2147483647, F1: (v5 | v5)}
    v2 = v2
    var v12 uint16 = 2
    if true {
        fmt.Println(v5, v2, v10)
        var v13 uint16 = 2
        v5 = (v5 | 42)
        _ = v13
    } else {
        v9 = (v9 + v9)
        if (1.5 > 3.14159) {
            var v14 int8 = v4
            if false {
                v4 -= v14
                _ = len(v6)
                switch v9 {
                case "":
                    v15 := GenS2{F0: 0}
                    _ = v15
                case "\n\t":
                    v12 += 0
                case "hello":
                    v10 -= v10
                default:
                    var v16 int32 = 698
                    _ = v16
                }
            }
            _ = v14
        }
        var v17 int64 = (v2 + (v2 * 9223372036854775807))
        _ = v17
    }
    switch v9 {
    case "a":
        var v18 uint = (v5 % ((129) | 1))
        _ = v18
    case "foo bar":
        var v19 string = (("hello" + v9) + ("a" + v9))
        _ = v19
    case "":
        fmt.Println(v2, v9)
    default:
        if true {
            var v20 uint32 = 2147483647
            fmt.Println(v2, v20, v9)
            _ = v20
        } else {
            v21 := GenS1{F0: 39, F1: v5}
            switch "foo bar" {
            case "":
                var v22 float64 = (-2.25 + float64(v10))
                _ = v22
            case "\n\t":
                v23 := GenS2{F0: 0}
                _ = v23
            default:
                var v24 uint8 = 42
                _ = v24
            }
            switch -960 {
            case -357:
                if len(v6) > 0 {
                    v6[0] = "a"
                }
            default:
                var v25 uint32 = 2
                _ = v25
            }
            _ = v21
        }
    }
    v10 = v10
    var v26 uint16 = (v12 * v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
