package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 int8 = 5
    var v4 int32 = 2147483647
    var v5 int64 = 1
    v3 = (v3 + (v3 + v3))
    v4 = v4
    v4 = 2147483647
    if false {
        v6 := v2[("foo bar" + ("hello" + ("foo bar" + ("a" + ""))))]
        v4 = -2147483648
        var v7 int64 = v5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 737
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 += v5
        _ = fi9
    }
    v10 := GenS1{F0: 4294967295, F1: v5}
    v11 := GenS2{F0: -53, F1: 375}
    switch "" {
    case "":
        v12 := v2[("" + "hello")]
        _ = v12
    case "\n\t":
        v13 := GenS1{F0: 2147483648, F1: (v5 >> ((902) & 63))}
        _ = v13
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 float64 = 0.1
            v16 := GenS3{F0: (v4 / ((-1) | 1))}
            fmt.Println(v1)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 int32 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
