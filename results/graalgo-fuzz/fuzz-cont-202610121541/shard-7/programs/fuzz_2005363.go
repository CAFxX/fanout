package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 83
    var v3 int8 = -1
    v4 := []int{-9223372036854775808, 32767}
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(v3)
        v3 = (v2 % (((v2 >> ((53) & 63))) | 1))
        fmt.Println(v2)
        _ = fi5
    }
    var v6 float64 = -2.25
    var v7 int32 = 0
    var v8 int = 35
    if len(v4) > 0 {
        v4[0] = v8
    }
    if (!(!(1000 == 1))) {
        switch ("" + "foo bar") {
        case "\n\t":
            delete(v1, 2147483648)
        case "foo bar":
            v6 = (v6 / v6)
        }
    }
    v9 := GenS1{F0: 42, F1: 0}
    v2 += (v3 >> ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
