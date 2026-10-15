package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", "a", "", ""}
    var v3 int16 = -1
    var v4 int32 = -444
    var v5 uint32 = 4294967295
    v4 += 0
    var v6 bool = ("hello" < "\n\t")
    v5 -= v5
    v7 := GenS1{F0: 1, F1: "foo bar", F2: 42}
    v8 := v1[42]
    v2 = append(v2, "a")
    v5 = (v5 >> ((255) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
