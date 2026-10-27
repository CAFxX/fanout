package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }

func main() {
    v1 := []uint32{}
    v2 := []string{"\n\t", "a", "\n\t", "\n\t"}
    var v3 bool = true
    v4 := make(map[int]int)
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 128
    fmt.Println(v3)
    fmt.Println(v6, v6, v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = -1
        var v9 int32 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 += v5
    v6 -= 2
    v3 = false
    switch 127 {
    case 127:
        v3 = v3
    case 42:
        switch "a" {
        case "hello":
            v10 := GenS1{F0: 42, F1: 32767, F2: (v6 - (v6 * v6))}
            _ = v10
        case "a":
            var v11 uint64 = 512
            _ = v11
        default:
            v4[int(v6)] = -233
        }
    case -645:
        var v12 uint8 = v6
        _ = v12
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v3, v3)
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        var v13 bool = false
        _ = v13
    default:
        fmt.Println(v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
