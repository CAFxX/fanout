package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 uint16 = 727
    var v4 uint64 = 42
    v1[4294967295] = false
    fmt.Println(v4)
    var v5 uint = 800
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 += v3
        switch (fi6 / (((fi6 % (((fi6 * (fi6 * 222))) | 1))) | 1)) {
        case 128:
            var v7 int32 = 554
            _ = v7
        default:
            var v8 int16 = -185
            _ = v8
        }
        _ = fi6
    }
    var v9 uint = 301
    v2["foo bar"] = 973
    var v10 uint32 = 2
    v10 += (v10 & v10)
    v5 = (^(v9))
    var v11 uint16 = v3
    v12 := []int{127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
