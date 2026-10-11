package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 string = "a"
    var v2 float32 = float32(0.0)
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    var v5 uint64 = 1
    v6 := []int{32767, -18, 127, 127}
    v5 += v5
    v5 = (^(v5))
    v3 -= 100
    v2 -= v2
    v3 = (v3 >> ((65535) & 63))
    if false {
        var v7 uint64 = 4294967295
        var v8 uint16 = 20
        v5 -= v5
        _ = v7
        _ = v8
    } else {
        fmt.Println(v5, v3)
        var v9 int64 = 115
        _ = v9
    }
    v1 = "a"
    v3 = (v3 >> ((65535) & 63))
    v1 = "a"
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: 2147483647}
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
