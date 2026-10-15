package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint64 }
func genFn0() uint64 {
    var v5 uint64 = 1
    v5 = v5
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float32 = float32(-2.25)
        var v8 int16 = -417
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -2147483648 {
    case -9223372036854775808:
        v5 += v5
    default:
        var v9 int8 = -1
        _ = v9
    }
    _ = v5
    return v5
}

func main() {
    var v1 int64 = -1
    v2 := []int{}
    var v3 int8 = 127
    var v4 string = "foo bar"
    var v10 int32 = 42
    if false {
        switch (v4 + (v4 + v4)) {
        case "a":
            fmt.Println(v1)
        case "hello":
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: 0, F1: v1, F2: 737}
                _ = len(v2)
                fmt.Println(v3, v4, v10)
                _ = fi11
                _ = v12
            }
        default:
            var v13 int64 = (^(v1))
            _ = v13
        }
    }
    var v14 uint = 2
    v15 := GenS1{F0: 0, F1: v1, F2: 306}
    var v16 int32 = (^(v10))
    if false {
        v17 := GenS1{F0: genFn0(), F1: (v1 & -9223372036854775808), F2: 1}
        _ = v17
    }
    fmt.Println(v3)
    var v18 uint8 = 117
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
