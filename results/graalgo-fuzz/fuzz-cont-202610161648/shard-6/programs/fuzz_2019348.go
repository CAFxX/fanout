package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 uint32 = 2147483648
    var v4 uint64 = 1
    v5 := make(map[int]bool)
    v6 := []uint32{(v3 | (v3 & v3)), (v3 - 65535)}
    var v7 int8 = -128
    switch (("foo bar" + "\n\t") + ("hello" + "")) {
    case "foo bar":
        v7 = v7
    default:
        var v8 bool = ((3.14159 / ((1.5 - 3.14159) + float64(1.5))) < 1e10)
        _ = v8
    }
    var v9 bool = false
    var v10 int32 = 2147483647
    delete(v1, v3)
    if (4294967295 == 255) {
        v9 = false
        var v11 int = -400
        if (((1e10 + (1.5 / 3.14159)) - (0.1 / 1.5)) >= (((0.0 / (3.14159 / 1.5)) - 1.5) + ((0.0 * 0.0) - 3.14159))) {
            v12 := GenS1{F0: 255, F1: v11, F2: 1}
            fmt.Println(v10, v10)
            _ = v12
        } else {
            v13 := GenS2{F0: 42}
            fmt.Println(v10, v7)
            _ = v13
        }
        _ = v11
    }
    fmt.Println(v3, v10, v4)
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
