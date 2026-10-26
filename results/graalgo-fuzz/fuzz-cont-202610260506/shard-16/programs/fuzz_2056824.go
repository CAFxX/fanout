package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0() uint64 {
    var v5 float32 = float32(0.1)
    if true {
        var v6 int16 = -32768
        _ = v6
    }
    _ = v5
    return 1
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    v4 := []string{"\n\t"}
    v1 += (v1 << ((64) & 63))
    var v7 int = 127
    var v8 uint = 759
    v7 += 2147483647
    var v9 bool = ((((("hello" + "hello") + "") + "\n\t") == "a") && true)
    var v10 int = v7
    v7 = int(v8)
    v11 := GenS1{F0: v1, F1: v1}
    fmt.Println(v7)
    switch v7 {
    case 127:
        _ = len(v3)
    }
    fmt.Println(v8, v7)
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
