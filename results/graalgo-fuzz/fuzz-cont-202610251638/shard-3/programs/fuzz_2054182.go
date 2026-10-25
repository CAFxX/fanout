package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]int)
    v3 := []string{("foo bar" + "\n\t"), "a"}
    v4 := make(map[uint32]bool)
    var v5 float32 = 1.5
    var v6 float32 = (v5 - 1.5)
    v2[526] = 818
    v6 = (((((float32(0.1) - 1e10) * -2.25) * v6) - 0.0) + v6)
    v1 -= (v1 * (v1 << ((v1) & 63)))
    var v7 bool = false
    var v8 int16 = 0
    if true {
        fmt.Println(v6)
        var v9 uint = (v1 << ((291) & 63))
        delete(v4, 1)
        _ = v9
    } else {
        var v10 uint8 = 1
        v5 = (((v6 - 1.5) / v6) * v5)
        v11 := GenS1{F0: 2147483648, F1: v8, F2: 646}
        _ = v10
        _ = v11
    }
    var v12 bool = (255 == 2)
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
