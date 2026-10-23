package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 bool = (429 < 2147483647)
    v4 := []uint32{42, 42, 0, 0}
    v2 -= (v2 | v2)
    var v5 uint = 2
    fmt.Println(v3, v2)
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v5)
        var v7 int16 = -32768
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    var v8 float64 = -2.25
    var v9 string = "foo bar"
    fmt.Println(v8, v2)
    if false {
        v1[588] = -1
    } else {
        v10 := GenS1{F0: uint64(v2)}
        var v11 uint16 = v2
        _ = v10
        _ = v11
    }
    v5 = (^(v5))
    if false {
        v9 = "a"
        delete(v1, 395)
    }
    v12 := GenS1{F0: 42}
    var v13 int = 4294967295
    var v14 float64 = (v8 - ((1.5 * 1e10) * (v8 + (3.14159 - v8))))
    v15 := GenS2{F0: 649}
    v3 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
