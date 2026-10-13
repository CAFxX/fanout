package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{0}
    var v3 int64 = -1
    var v4 int8 = 127
    var v5 uint16 = 1
    v5 = (v5 >> ((2) & 63))
    var v6 float32 = 0.0
    var v7 uint16 = v5
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: (v5 & v5)}
        var v10 uint32 = 2147483647
        fmt.Println(v3, v5)
        _ = fi8
        _ = v10
        _ = v9
    }
    v5 -= 694
    v5 -= 2
    var v11 float32 = v6
    for fi12 := 0; fi12 < 3; fi12++ {
        v5 += (v7 & 754)
        _ = fi12
    }
    v4 = -60
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
