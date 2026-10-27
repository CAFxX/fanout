package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{"foo bar", "\n\t", "a"}
    var v3 int8 = 120
    var v4 int32 = -2147483648
    v5 := make(map[int]uint32)
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v4 += -82
    }
    var v6 int16 = -855
    var v7 uint = 0
    var v8 uint16 = 32767
    fmt.Println(v4)
    var v9 int = 107
    if true {
        v7 -= (v7 >> ((0) & 63))
    }
    fmt.Println(v7)
    v10 := GenS1{F0: -109, F1: 42, F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
