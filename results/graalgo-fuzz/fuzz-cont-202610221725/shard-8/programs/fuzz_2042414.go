package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    v3 := []uint32{752, 4294967295, 576}
    v1 = (!(!(0 >= 42)))
    switch "hello" {
    case "a":
        v1 = (!false)
    case "foo bar":
        var v4 int16 = 32767
        _ = v4
    }
    fmt.Println(v1)
    v5 := GenS1{F0: 32768}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v1, v1)
        _ = fi6
    }
    var v7 uint16 = 1
    v7 = 2
    v7 += 102
    v8 := v2[32768]
    v8 = (v8 | v8)
    var v9 uint8 = 2
    var v10 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
