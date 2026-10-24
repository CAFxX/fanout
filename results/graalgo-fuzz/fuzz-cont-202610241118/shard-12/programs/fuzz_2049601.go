package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint = 4294967295
    v2 := []string{"\n\t", (("foo bar" + "") + "\n\t")}
    v3 := make(map[uint32]int)
    var v4 int32 = -2147483648
    var v5 float32 = 0.0
    var v6 float32 = float32((0.0 * 1e10))
    var v7 int32 = v4
    switch "" {
    case "hello":
        _ = len(v2)
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        v8 := GenS1{F0: 656, F1: v1}
        _ = v8
    default:
        _ = len(v3)
    }
    v5 += v6
    v4 = (v4 << ((1) & 63))
    fmt.Println(v6, v4)
    var v9 uint64 = 399
    v7 -= -1
    v5 = v5
    v10 := v3[2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
