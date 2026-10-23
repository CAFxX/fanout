package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 676
    var v2 int32 = 1
    v3 := make(map[uint32]int)
    v4 := []int{-62, int(v1), 2147483647}
    v5 := make(map[uint32]int)
    var v6 float64 = float64(1.5)
    v2 = v2
    var v7 int32 = (v2 << ((97) & 63))
    v4 = append(v4, 128)
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            _ = len(v4)
            _ = fi8
        }
        v1 = 1
        var v9 int32 = (v2 << ((255) & 63))
        _ = v9
    } else {
        v1 += (v1 / ((734) | 1))
    }
    var v10 int16 = 32767
    v7 = (v7 | (v7 | -1))
    var v11 int = 128
    var v12 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
