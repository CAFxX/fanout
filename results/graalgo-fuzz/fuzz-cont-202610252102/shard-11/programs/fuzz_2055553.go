package main

import (
    "fmt"
)


func main() {
    var v1 bool = ((698 <= -2147483648) && (1 <= -578))
    v2 := []uint32{285, 290, 1}
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v1 = (618 < 994)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 bool = (v1 && v1)
        _ = fi6
        _ = v7
    }
    var v8 int16 = 1000
    v2 = append(v2, 829)
    switch "hello" {
    case "hello":
        var v9 uint64 = 2
        _ = v9
    case "":
        v1 = true
    case "a":
        v8 -= v8
    }
    fmt.Println(v8)
    v8 = (v8 << ((0) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
