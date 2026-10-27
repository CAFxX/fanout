package main

import (
    "fmt"
)


func main() {
    v1 := []int{}
    var v2 uint = 853
    var v3 uint8 = 2
    v4 := []uint32{2147483647, 65535, 0, 2147483647}
    v4 = append(v4, 42)
    v1 = append(v1, 42)
    if (-868 == 1) {
        v2 += (v2 << ((v2) & 63))
    }
    var v5 uint16 = 913
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int8 = 109
        _ = fi6
        _ = v7
    }
    v4 = append(v4, 1)
    v3 = (v3 ^ (v3 | v3))
    var v8 bool = (true && true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
