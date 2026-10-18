package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 608
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    v1 = v1
    v6 := []int{630}
    v1 += (v1 - v1)
    if (v1 == v1) {
        v1 = (v1 % ((2) | 1))
        v1 = v1
    } else {
        v1 += (v1 & v1)
    }
    v1 -= 1
    fmt.Println(v1)
    var v7 uint = 514
    var v8 int = -581
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v1 = v1
        var v10 uint32 = 1
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
