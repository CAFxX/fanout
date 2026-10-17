package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = -217
    v4 := make(map[string]string)
    var v5 int64 = 484
    v6 := make(map[uint32]int)
    fmt.Println(v5, v5)
    v5 = v5
    v5 -= (^(v5))
    var v7 int = -34
    v1[v7] = 4294967295
    v3 += (v3 % (((v3 & (v3 | (v3 << ((15) & 63))))) | 1))
    v3 = -1
    var v8 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
