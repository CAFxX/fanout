package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 100
    v2 := make(map[int]int)
    var v3 int16 = 450
    v4 := []uint32{}
    v5 := make(map[int]string)
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = 2147483647
            v9 += v9
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v6 = (v6 * (v6 / 1.0))
    fmt.Println(v1, v6)
    var v10 int8 = (v1 >> ((658) & 63))
    var v11 int8 = 1
    v3 -= (v3 - 32767)
    var v12 int8 = 0
    if (65535 < 2) {
        var v13 uint8 = 2
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
