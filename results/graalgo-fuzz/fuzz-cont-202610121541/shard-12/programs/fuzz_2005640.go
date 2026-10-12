package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 366
    var v3 uint32 = 4294967295
    v4 := []string{"", "foo bar", "foo bar"}
    v5 := GenS1{F0: -1, F1: -310, F2: 127}
    v6 := GenS1{F0: 9223372036854775807, F1: 1, F2: 42}
    var v7 uint64 = v2
    v1 += 0.0
    v1 -= float32(0.0)
    fmt.Println(v7)
    v2 = uint64(v3)
    if (((-1 == 1) && (0 == 42)) || (0.0 > -2.25)) {
        fmt.Println(v7, v7, v1)
    } else {
        if true {
            var v8 uint16 = 32767
            var v9 int64 = 21
            _ = v8
            _ = v9
        } else {
            var v10 int64 = 1
            v11 := GenS1{F0: (v10 & v10), F1: 1, F2: -108}
            v4 = append(v4, ("\n\t" + "hello"))
            _ = v10
            _ = v11
        }
        fmt.Println(v7)
        v1 += v1
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
