package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    v3 := []uint32{4294967295}
    var v4 uint64 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v2)
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v2, v2, fi7)
        _ = fi7
    }
    var v8 uint = 796
    var v9 int8 = v2
    v10 := GenS1{F0: 1, F1: -1, F2: 42}
    var v11 string = "\n\t"
    var v12 int8 = (v2 + 100)
    v9 -= v2
    var v13 string = "foo bar"
    var v14 float64 = 1.5
    if true {
        fmt.Println(v6, v9, v11)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
