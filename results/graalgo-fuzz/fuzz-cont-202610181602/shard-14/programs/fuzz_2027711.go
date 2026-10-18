package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []string{("" + "\n\t"), "foo bar", "\n\t", "\n\t"}
    v2 := make(map[uint32]int)
    var v3 bool = (251 <= 1000)
    var v4 uint8 = 2
    var v5 uint8 = 1
    v6 := GenS1{F0: 23}
    var v7 uint8 = 255
    var v8 int32 = 2147483647
    v5 -= (v7 / ((v4) | 1))
    v4 -= (v5 << ((822) & 63))
    var v9 int64 = 0
    v10 := GenS1{F0: 203}
    var v11 int16 = -231
    fmt.Println(v4, v4, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
