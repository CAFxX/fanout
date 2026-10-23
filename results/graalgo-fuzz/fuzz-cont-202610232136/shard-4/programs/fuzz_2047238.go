package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int = 0
    v4 := GenS1{F0: 2}
    var v5 int16 = -227
    v6 := GenS1{F0: 1}
    var v7 uint = 780
    var v8 string = "a"
    v7 -= 65535
    v7 -= 18446744073709551615
    v3 += v3
    v9 := GenS1{F0: 0}
    fmt.Println(v7)
    v10 := GenS1{F0: 42}
    var v11 int8 = 1
    var v12 int16 = 32767
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
