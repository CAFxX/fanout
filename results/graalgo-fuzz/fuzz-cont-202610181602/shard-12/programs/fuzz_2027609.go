package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint16 = 62
    var v2 int64 = 0
    var v3 int32 = 1
    v4 := make(map[string]int)
    var v5 uint8 = 128
    var v6 uint8 = v5
    fmt.Println(v3, v2)
    if false {
        switch -9223372036854775808 {
        case -591:
            fmt.Println(v1, v5, v3)
        case -465:
            v5 -= (v6 / ((v5) | 1))
        }
    } else {
        v6 += 1
    }
    if false {
        fmt.Println(v3)
        var v7 uint64 = 714
        fmt.Println(v6, v2, v2)
        _ = v7
    }
    v8 := []int{0}
    _ = len(v8)
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint32 = 0
        var v11 string = ("a" + "a")
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 = v5
    switch ("hello" + ("hello" + "hello")) {
    case "":
        _ = len(v8)
    case "a":
        var v12 bool = (280 > 4294967295)
        _ = v12
    case "hello":
        v13 := GenS1{F0: (v5 + v6)}
        _ = v13
    }
    v3 = (v3 ^ v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
