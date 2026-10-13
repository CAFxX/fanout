package main

import (
    "fmt"
)

func genFn0() uint64 {
    var v5 uint8 = 107
    var v6 float64 = (float64(((((1.5 / 1e10) * 1e10) - (1.5 - (1e10 / 1.5))) + 0.1)) + 3.14159)
    _ = v5
    _ = v6
    return 143
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]string)
    var v3 uint32 = 483
    v4 := make(map[uint32]bool)
    v1 = (v1 % ((v1) | 1))
    fmt.Println(v1, v1)
    v1 = (v1 + v1)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = fi8
            var v10 int = (fi8 >> ((0) & 63))
            fmt.Println(v1, fi8)
            v11 := []int{-88, (fi8 >> ((422) & 63)), 128, v10}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    v12 := []int{32767, -9223372036854775808, 4294967295, 127}
    switch 0 {
    case 0:
        fmt.Println(v1)
    case 1:
        if true {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint16 = uint16(v3)
                _ = fi13
                _ = v14
            }
        } else {
            if false {
                v1 += (v1 >> ((799) & 63))
            } else {
                var v15 uint8 = 127
                _ = v15
            }
            v3 += 583
        }
    default:
        var v16 uint16 = 2
        _ = v16
    }
    var v17 int32 = -487
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
