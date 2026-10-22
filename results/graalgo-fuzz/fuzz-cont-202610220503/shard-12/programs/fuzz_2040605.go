package main

import (
    "fmt"
)


func main() {
    var v1 uint = 4294967295
    var v2 float32 = 0.0
    var v3 uint16 = 42
    v4 := []uint32{758, 2147483648, 783}
    var v5 int = 32767
    v1 += v1
    v4 = append(v4, 2147483647)
    v3 += (v3 * (^(v3)))
    if (0 > v1) {
        var v6 uint32 = 495
        v6 -= uint32(v3)
        fmt.Println(v6, v5, v2)
        _ = v6
    } else {
        v1 -= (v1 * 791)
    }
    var v7 uint32 = 4294967295
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = 0.0
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 19; fi10++ {
        v7 += (v7 & (v7 * (^(v7))))
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    v7 = v7
    v1 += uint(v3)
    var v12 uint64 = 1
    var v13 float64 = (float64(((1e10 + (v2 + (float32(1.5) * (v2 / ((3.14159 / float32(-2.25)) / v2))))) / (v2 / (v2 / v2)))) - (1e10 / float64((v2 * (float32(((-2.25 / 1.5) * 0.0)) * float32(1e10))))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
