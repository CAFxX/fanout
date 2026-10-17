package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2147483647
    var v3 int16 = -1
    v4 := make(map[int]uint32)
    var v5 int = 127
    var v6 uint16 = 42
    var v7 int32 = 0
    var v8 string = "a"
    _ = len(v4)
    var v9 int = (v5 << ((18446744073709551615) & 63))
    delete(v4, (v9 + (v9 / (((v9 % (((^(v9))) | 1))) | 1))))
    v8 = ("" + (("foo bar" + (v8 + v8)) + v8))
    var v10 int64 = v1
    v1 -= v1
    v11 := GenS1{F0: 133, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
