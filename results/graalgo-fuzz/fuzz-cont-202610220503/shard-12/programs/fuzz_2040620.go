package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (1.5 + 3.14159)
    v2 := []uint32{2147483647, 2147483647, 1, 2}
    v3 := []uint32{}
    v4 := []int{183}
    var v5 float64 = (0.0 / (0.0 + ((1e10 + (1.5 + -2.25)) / -2.25)))
    v1 = v1
    var v6 float64 = v5
    var v7 float64 = v6
    var v8 float32 = (v1 + v1)
    for fi9 := 0; fi9 < 17; fi9++ {
        switch -2147483648 {
        case -2147483648:
            v1 += (v1 * (0.1 * (v1 / (1e10 + v1))))
        default:
            var v10 float32 = v8
            _ = v10
        }
        _ = fi9
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    v7 -= v6
    if (-315 > -32768) {
        var v11 uint = 0
        var v12 uint16 = 1
        _ = v11
        _ = v12
    }
    v8 = v8
    v6 = float64((float32(3.14159) - (float32(v7) * (3.14159 - 1e10))))
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 bool = ((!(true && (1 != -1))) && false)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int64 = 42
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
