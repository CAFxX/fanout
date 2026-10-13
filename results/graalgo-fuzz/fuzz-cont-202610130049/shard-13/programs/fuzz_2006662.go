package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 float32 = v1
    var v5 int8 = -57
    var v6 int64 = 1
    v1 -= -2.25
    v7 := []int{}
    _ = len(v3)
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (fi9 ^ (fi9 - (^(fi9)))) {
            case 2147483647:
                v5 -= v5
            case 4294967295:
                var v11 float64 = float64(v8)
                _ = v11
            case 42:
                v8 += v1
            default:
                switch "\n\t" {
                case "\n\t":
                    if false {
                        v1 += v8
                    } else {
                        fmt.Println(v1)
                        v2 += (v2 >> ((405) & 63))
                    }
                }
            }
            var v12 bool = false
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v6 = (v6 & (v6 / ((0) | 1)))
            fmt.Println(v5)
            v2 = 2
            _ = fi13
        }
        _ = fi9
    }
    var v14 uint32 = 2147483648
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint8 = 42
        _ = fi15
        _ = v16
    }
    var v17 int64 = (v6 - (v6 << ((255) & 63)))
    v14 = v14
    if true {
        var v18 bool = (true || (((1.5 / (v1 + v4)) - v4) != (v1 * v8)))
        _ = v18
    } else {
        var v19 float32 = v8
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
