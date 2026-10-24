package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    var v3 string = "\n\t"
    var v4 uint64 = 18446744073709551615
    var v5 float64 = 1.5
    switch "foo bar" {
    case "hello":
        var v6 uint = 680
        _ = v6
    default:
        fmt.Println(v4, v4, v5)
    }
    if (!true) {
        v1 = append(v1, v3)
        if true {
            var v7 int32 = 0
            _ = v7
        }
    } else {
        var v8 uint32 = 42
        fmt.Println(v3)
        v4 -= (v4 + v4)
        _ = v8
    }
    var v9 int64 = 549
    var v10 uint = 0
    var v11 int16 = 1
    v12 := GenS1{F0: uint(v4), F1: 42}
    fmt.Println(v10)
    v9 += (v9 + -1)
    switch 0 {
    case 4294967295:
        var v13 int16 = v11
        _ = v13
    default:
        switch 32768 {
        case 127:
            for fi14 := 0; fi14 < 8; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS2{F0: v4, F1: -1}
                    _ = fi15
                    _ = v16
                }
                v9 = v9
                v3 = "a"
                _ = fi14
            }
        case -2147483648:
            var v17 bool = false
            _ = v17
        default:
            fmt.Println(v4)
        }
    }
    v18 := GenS3{F0: (v4 << ((v10) & 63))}
    fmt.Println(v10, v11)
    switch "foo bar" {
    case "\n\t":
        v5 -= 1e10
    case "":
        v4 = v4
    default:
        v19 := GenS3{F0: v4}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
