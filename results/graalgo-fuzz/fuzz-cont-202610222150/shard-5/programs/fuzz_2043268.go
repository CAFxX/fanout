package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -128
    var v3 int16 = 32767
    v4 := make(map[uint32]string)
    var v5 int16 = (v3 % ((42) | 1))
    var v6 uint32 = 74
    var v7 int32 = 0
    fmt.Println(v3, v6)
    var v8 float64 = 0.0
    v9 := []int{2147483647}
    switch 4294967295 {
    case 15:
        v6 = (v6 << ((2) & 63))
    default:
        v6 = v6
    }
    fmt.Println(v2, v6, v3)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = 0.0
        _ = fi10
        _ = v11
    }
    v7 = v7
    var v12 int16 = (v5 % ((-1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
