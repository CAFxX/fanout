package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 uint = 868
    v4 := []string{("foo bar" + ""), ""}
    var v5 string = ("hello" + "a")
    var v6 int64 = 1
    v5 = "\n\t"
    var v7 uint = v3
    v1[4294967295] = (329 <= 0)
    for fi8 := 0; fi8 < 5; fi8++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        _ = fi8
    }
    v9 := GenS1{F0: false}
    var v10 string = (v5 + v5)
    v11 := GenS2{F0: (((v10 + "\n\t") + "foo bar") + v5), F1: 42, F2: (v6 * (v6 * v6))}
    fmt.Println(v7, v5, v6)
    v10 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
