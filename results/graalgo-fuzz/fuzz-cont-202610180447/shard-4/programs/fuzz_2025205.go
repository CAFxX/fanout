package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8; F2 string }

func main() {
    v1 := []uint32{103, 0, 2, 65535}
    var v2 string = ("a" + "\n\t")
    var v3 int16 = 1
    var v4 uint32 = 2147483647
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 9; fi6++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = "hello"
        _ = fi6
    }
    var v7 int64 = 1
    v3 -= 42
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: (v2 + v2)}
        _ = fi8
        _ = v9
    }
    if (true || true) {
        v10 := GenS1{F0: v2}
        v4 -= (^(v4))
        _ = v10
    }
    fmt.Println(v4, v3, v3)
    switch v2 {
    case "hello":
        var v11 int64 = (v7 ^ (^(v7)))
        _ = v11
    default:
        fmt.Println(v4, v4, v7)
    }
    v2 = "foo bar"
    fmt.Println(v3)
    v12 := GenS2{F0: (0 >= (v4 | v4)), F1: 127, F2: ("a" + "\n\t")}
    switch v2 {
    case "a":
        var v13 bool = ((v3 + v3) > 0)
        _ = v13
    }
    v14 := GenS1{F0: "hello"}
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
