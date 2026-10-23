package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    var v3 float64 = 0.0
    v4 := []string{(("a" + "\n\t") + ""), "", "foo bar"}
    var v5 int64 = -1
    delete(v1, "hello")
    v6 := v1[("hello" + "hello")]
    var v7 uint32 = 363
    v5 += int64(v7)
    v4 = append(v4, "foo bar")
    v3 -= 0.0
    v8 := GenS1{F0: v5, F1: ("foo bar" + "a"), F2: "a"}
    v9 := GenS2{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
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
