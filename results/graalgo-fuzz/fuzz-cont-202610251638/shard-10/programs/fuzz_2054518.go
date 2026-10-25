package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }

func main() {
    var v1 uint16 = 0
    v2 := []int{-895, 9223372036854775807}
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint32 = 1
        _ = fi5
        _ = v6
    }
    v1 = (^(v1))
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = -667
        _ = fi8
        _ = v9
    }
    if false {
        if true {
            _ = len(v4)
            for fi10 := 0; fi10 < 8; fi10++ {
                v7 -= 4294967295
                _ = fi10
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 := GenS1{F0: -1, F1: -755, F2: ("hello" + "foo bar")}
    v7 = (v7 + (v7 * v7))
    _ = len(v4)
    fmt.Println(v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
