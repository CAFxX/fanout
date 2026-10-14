package main

import (
    "fmt"
)


func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 uint = 18446744073709551615
    var v5 int16 = 1000
    v4 = v4
    var v6 uint16 = 1
    if len(v1) > 0 {
        v1[0] = 859
    }
    v6 = (v6 >> (((v4 ^ (^(v4)))) & 63))
    v7 := v3[190]
    for fi8 := 0; fi8 < 4; fi8++ {
        if (!true) {
            v5 = v5
        }
        _ = fi8
    }
    v5 += v5
    if (v4 < (v4 + (v4 % (((v4 >> ((v4) & 63))) | 1)))) {
        var v9 uint32 = 2
        _ = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
