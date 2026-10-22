package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = (v4 / ((v4) | 1))
        _ = fi5
        _ = v6
    }
    v1[529] = 1
    v4 = v4
    fmt.Println(v4)
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    var v7 int16 = 1
    var v8 uint8 = 127
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
