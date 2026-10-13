package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 574
    var v3 int = -2147483648
    v2 -= -9223372036854775808
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(v3, v2, v3)
        _ = fi4
    }
    var v5 int64 = -9223372036854775808
    var v6 string = ("foo bar" + "a")
    if false {
        v7 := GenS1{F0: 609, F1: v5}
        var v8 int8 = 100
        _ = v7
        _ = v8
    } else {
        v2 = (v2 + (v2 & -197))
        var v9 uint64 = 4294967295
        _ = v9
    }
    var v10 int64 = (v5 % ((v5) | 1))
    var v11 uint32 = 4294967295
    var v12 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
