package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    fmt.Println(v1, v1)
    var v4 uint8 = 128
    var v5 float32 = float32((((0.1 - 0.0) / 3.14159) / 1e10))
    delete(v2, "foo bar")
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{(fi6 + (fi6 / (((fi6 / (((-(fi6))) | 1))) | 1))), -793}
        _ = fi6
        _ = v7
    }
    switch ("" + ("" + "")) {
    case "a":
        if v1 {
            v8 := GenS1{F0: 2, F1: "", F2: v1}
            _ = v8
        } else {
            var v9 uint = 4294967295
            var v10 uint64 = 1
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int = 9223372036854775807
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
    default:
        v5 = v5
    }
    if (-552 == (v3 >> ((65535) & 63))) {
        var v13 float32 = (v5 * v5)
        var v14 int = -2147483648
        _ = v13
        _ = v14
    } else {
        v15 := GenS2{F0: 2, F1: ((((("\n\t" + "foo bar") + "hello") + "a") + "\n\t") + "a")}
        _ = v15
    }
    v16 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
