package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{("foo bar" + "a"), "a", "foo bar"}
    var v2 uint64 = 259
    v3 := []uint32{2147483647, 606, 2, 204}
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 1}
    var v7 uint16 = 573
    v7 += v7
    v8 := GenS1{F0: 4294967295}
    var v9 uint8 = 127
    var v10 uint32 = 990
    var v11 int = 177
    var v12 uint16 = (v7 | v7)
    v11 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
