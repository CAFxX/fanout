package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]bool)
    v3 := []string{"foo bar", "a", "foo bar", ("\n\t" + "a")}
    if false {
        v4 := GenS1{F0: 833, F1: -1}
        var v5 float32 = -2.25
        _ = v4
        _ = v5
    }
    var v6 float64 = v1
    var v7 uint = 18446744073709551615
    var v8 uint = (^(v7))
    fmt.Println(v7, v6, v6)
    switch 128 {
    case 9223372036854775807:
        var v9 int8 = 127
        _ = v9
    case -9223372036854775808:
        var v10 int32 = -725
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v6, fi11)
            _ = fi11
        }
    }
    if (false || false) {
        var v12 uint64 = 122
        v1 = float64(0.0)
        _ = v12
    }
    v13 := GenS2{F0: 9223372036854775807, F1: "a"}
    if true {
        v14 := GenS1{F0: v7, F1: 0}
        for fi15 := 0; fi15 < 1; fi15++ {
            v8 = v7
            v16 := GenS2{F0: 1, F1: ("foo bar" + "")}
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        var v17 uint16 = 2
        _ = v17
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
