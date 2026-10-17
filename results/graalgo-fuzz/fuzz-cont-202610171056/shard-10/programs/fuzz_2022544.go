package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 uint64, p2 int16) uint16 {
    v4 := []int{}
    var v5 int16 = (p2 / (((p2 * int16(p0))) | 1))
    if (42 >= (p0 & 0)) {
        if len(v4) > 0 {
            v4[0] = 570
        }
        var v6 float32 = 1.5
        _ = v6
    } else {
        var v7 bool = false
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 128
            v9 -= (^(v9))
            var v10 uint16 = 56
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        p0 += (p0 - (p0 ^ (p0 * p0)))
        p2 += p2
        _ = fi11
    }
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    v12 := []int{9223372036854775807, 2147483647, 32767}
    switch "" {
    case "a":
        fmt.Println(v2)
    default:
        fmt.Println(v2)
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 float32 = 3.14159
        v2 -= 845
        v14 -= v14
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 string = "a"
        _ = fi15
        _ = v16
    }
    v2 = 0
    var v17 uint8 = 0
    v17 += (v17 >> ((297) & 63))
    v2 -= (v2 >> ((0) & 63))
    for fi18 := 0; fi18 < 8; fi18++ {
        switch (fi18 * (fi18 | 2147483647)) {
        case 2147483647:
            switch fi18 {
            case 2147483647:
                var v19 uint = 1
                _ = v19
            }
        }
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
