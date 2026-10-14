package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := []string{("" + "foo bar"), "\n\t"}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 int32 = 42
    if (!true) {
        fmt.Println(v4)
    }
    var v5 uint8 = 128
    var v6 uint8 = v5
    if (0 < 930) {
        v7 := GenS1{F0: 42}
        v5 += (v5 * (v5 >> ((2) & 63)))
        v1 = append(v1, "")
        _ = v7
    } else {
        fmt.Println(v4, v6, v4)
    }
    v8 := GenS2{F0: -128, F1: 694}
    var v9 string = ("foo bar" + "foo bar")
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    fmt.Println(v5, v4)
    var v11 int32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
