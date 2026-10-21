package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    var v3 float64 = float64((0.1 + 3.14159))
    v4 := make(map[int]int)
    v1 = (v1 << ((2) & 63))
    v3 += v3
    v5 := GenS1{F0: "hello", F1: ""}
    var v6 int64 = (v1 << ((665) & 63))
    v3 = (v3 / 0.1)
    var v7 int = int(v6)
    v8 := GenS1{F0: ("foo bar" + "\n\t"), F1: "a"}
    v9 := GenS1{F0: "foo bar", F1: (("foo bar" + "hello") + "a")}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
