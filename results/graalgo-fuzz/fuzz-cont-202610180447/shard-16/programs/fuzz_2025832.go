package main

import (
    "fmt"
)


func main() {
    var v1 int32 = -458
    v2 := make(map[uint32]int)
    var v3 uint32 = 726
    var v4 float64 = (1e10 - (0.0 / float64((float32(0.1) + 1e10))))
    if (true && true) {
        v1 += (v1 >> ((0) & 63))
    }
    v5 := []int{}
    var v6 uint32 = (v3 - v3)
    var v7 uint64 = 1
    fmt.Println(v6)
    for fi8 := 0; fi8 < 3; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v7 = v7
        switch "\n\t" {
        case "foo bar":
            v1 = (v1 >> ((4294967295) & 63))
        default:
            var v9 int32 = (^(v1))
            _ = v9
        }
        _ = fi8
    }
    var v10 uint8 = 49
    fmt.Println(v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
