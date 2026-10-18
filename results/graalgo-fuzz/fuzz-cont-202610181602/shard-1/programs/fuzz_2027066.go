package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 1
    v2 := make(map[int]string)
    var v3 uint64 = (v1 >> ((65535) & 63))
    var v4 float64 = 0.1
    v1 = 60
    v3 -= (v1 | (v1 << ((428) & 63)))
    v4 += (v4 * (-2.25 + v4))
    var v5 float64 = ((v4 + v4) / v4)
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 -= 42
        v1 = 2
        _ = fi6
    }
    v7 := []int{32768, 127, 42, -1}
    for fi8 := 0; fi8 < 3; fi8++ {
        v1 = (v1 + v1)
        _ = fi8
    }
    v1 = (v1 + (v3 & v1))
    v5 = v4
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 float64 = (float64(0.1) / (v4 * float64((float32(v4) + 0.0))))
            _ = fi10
            _ = v11
        }
        v4 = v4
        v3 = v1
        _ = fi9
    }
    var v12 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
