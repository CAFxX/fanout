package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    v1 := []string{"a", "", "foo bar", "foo bar"}
    var v2 uint64 = 364
    var v3 bool = (564 == 359)
    v4 := []int{-1}
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 = append(v4, fi5)
        var v6 uint64 = (v2 & (v2 | (v2 + 4294967295)))
        v6 -= 42
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 42
    var v8 int16 = 0
    v4 = append(v4, -359)
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: -1, F1: 2, F2: 2147483647}
        v7 -= v7
        if (false || (!true)) {
            fmt.Println(v9, v3, v9)
        } else {
            v3 = v3
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v8, v7)
    var v12 float32 = v9
    v1 = append(v1, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
