package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 int32; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    var v3 int32 = 1
    var v4 string = "\n\t"
    v5 := []string{}
    var v6 uint = 255
    v7 := GenS1{F0: (false || ((v2 - (v2 & uint32(v6))) <= (v2 / ((v2) | 1)))), F1: v4}
    v6 -= v6
    fmt.Println(v3, v3)
    v2 -= v2
    v3 -= (v3 | -2147483648)
    _ = len(v5)
    fmt.Println(v4)
    v8 := GenS1{F0: ((3.14159 * 1e10) <= 3.14159), F1: "foo bar"}
    v9 := GenS2{F0: -959, F1: v3, F2: 42}
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint64 = 944
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 449
    fmt.Println(v2, v2)
    if (false && true) {
        for fi13 := 0; fi13 < 17; fi13++ {
            v12 -= v12
            var v14 uint64 = 18446744073709551615
            _ = fi13
            _ = v14
        }
        v6 -= v6
        var v15 uint = 255
        _ = v15
    } else {
        v5 = append(v5, "hello")
        var v16 uint32 = v2
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: (v3 & (v3 - 0)), F2: -1}
    v5 = append(v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
