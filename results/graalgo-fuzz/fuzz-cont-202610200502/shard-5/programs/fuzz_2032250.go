package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 bool = (!true)
    var v4 uint32 = 0
    delete(v1, 9223372036854775807)
    fmt.Println(v3, v4)
    v4 += v4
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6, v5)
        fmt.Println(fi6, v3, fi6)
        var v7 int = (-(fi6))
        _ = fi6
        _ = v7
    }
    v5 -= v5
    v8 := GenS1{F0: -1}
    var v9 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
