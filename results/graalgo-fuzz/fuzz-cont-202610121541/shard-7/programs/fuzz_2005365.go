package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    v1 := []string{"hello", "foo bar", ""}
    v2 := []uint32{}
    v3 := []string{"", (("a" + ("hello" + "a")) + ("\n\t" + ("foo bar" + "\n\t"))), "a", "foo bar"}
    var v4 string = "foo bar"
    var v5 int64 = -9223372036854775808
    var v6 float32 = (((1e10 / 1e10) * 1e10) * 0.1)
    v7 := GenS1{F0: 2}
    v6 = 1.5
    v4 = v4
    v5 += 0
    var v8 uint16 = 65535
    for fi9 := 0; fi9 < 15; fi9++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi9
    }
    v4 = "hello"
    v6 = 0.0
    var v10 int8 = 79
    fmt.Println(v4, v10, v6)
    v11 := GenS2{F0: (v8 / ((255) | 1))}
    switch 1 {
    case 2147483647:
        var v12 int8 = (v10 * -46)
        _ = v12
    case 616:
        var v13 int8 = (v10 | (^(v10)))
        _ = v13
    case -737:
        fmt.Println(v6, v6)
    }
    var v14 string = v4
    v15 := GenS1{F0: 2147483647}
    if false {
        switch 1 {
        case -798:
            fmt.Println(v14, v14)
        case -2147483648:
            v16 := GenS3{F0: 2, F1: 888, F2: 42}
            _ = v16
        default:
            v17 := GenS4{F0: v8, F1: 42, F2: v4}
            _ = v17
        }
    }
    v14 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
