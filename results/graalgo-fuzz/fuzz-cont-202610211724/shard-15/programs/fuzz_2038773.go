package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }
func genFn0(p0 uint32) int16 {
    var v7 int16 = 0
    fmt.Println(v7)
    _ = p0
    _ = v7
    return v7
}

func main() {
    var v1 int = -1
    v2 := make(map[int]bool)
    v3 := []string{"", "\n\t"}
    var v4 uint = 1
    v5 := make(map[string]string)
    var v6 int32 = -1
    var v8 float64 = -2.25
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int32 = (v6 + v6)
            v11 := GenS1{F0: 910, F1: 0, F2: 1}
            fmt.Println(v4)
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        var v12 uint64 = 2
        if true {
            fmt.Println(v6, v1, v4)
            var v13 float64 = float64(1e10)
            _ = v13
        } else {
            if false {
                var v14 uint16 = 42
                _ = v14
            }
        }
        _ = v12
    }
    v4 -= (v4 % ((v4) | 1))
    var v15 bool = true
    v16 := GenS1{F0: uint64(v4), F1: 0, F2: 0}
    v4 = 42
    switch "foo bar" {
    case "\n\t":
        v17 := GenS1{F0: 18446744073709551615, F1: -271, F2: 0}
        _ = v17
    default:
        v18 := GenS1{F0: 0, F1: 882, F2: -9223372036854775808}
        _ = v18
    }
    v19 := GenS1{F0: 252, F1: 9223372036854775807, F2: 9223372036854775807}
    var v20 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
