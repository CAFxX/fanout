package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    var v4 int64 = -872
    v5 := make(map[uint32]int)
    _ = len(v2)
    if (v4 <= (v4 % ((v4) | 1))) {
        if false {
            var v6 bool = (!(42 != 330))
            _ = v6
        } else {
            v3 = (v3 & 318)
        }
    }
    v7 := GenS1{F0: 1, F1: 1}
    fmt.Println(v4)
    var v8 int64 = v4
    fmt.Println(v4, v3)
    v3 -= 1
    var v9 string = "hello"
    var v10 uint = 255
    v4 -= (v8 >> ((v10) & 63))
    v4 = v8
    var v11 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
