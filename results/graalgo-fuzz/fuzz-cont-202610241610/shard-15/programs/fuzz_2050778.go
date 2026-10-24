package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []uint32{2, 42}
    var v2 bool = (-1 == 388)
    v3 := []string{((("" + "a") + "foo bar") + "foo bar"), "\n\t", ("" + "foo bar"), ""}
    var v4 int = 4294967295
    var v5 int64 = -9223372036854775808
    var v6 int64 = v5
    v5 = (-(v5))
    fmt.Println(v4)
    if v2 {
        v1 = append(v1, 0)
        var v7 int8 = 100
        _ = v7
    } else {
        var v8 uint64 = 0
        _ = v8
    }
    var v9 uint64 = 139
    fmt.Println(v9)
    v2 = (949 >= 0)
    switch "" {
    case "\n\t":
        if true {
            v10 := GenS1{F0: 255, F1: 23}
            v5 -= -9223372036854775808
            _ = v10
        } else {
            v5 += v5
            fmt.Println(v9, v9, v9)
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
