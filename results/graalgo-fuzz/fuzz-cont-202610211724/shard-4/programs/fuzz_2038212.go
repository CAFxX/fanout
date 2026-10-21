package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{2147483647, 2}
    var v2 uint = 0
    var v3 float32 = 1e10
    var v4 uint16 = 32767
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v3 = float32(3.14159)
    var v5 int64 = -902
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 593 {
    case -584:
        v4 = v4
    case 128:
        var v6 uint16 = (v4 ^ (^(v4)))
        _ = v6
    case 2147483647:
        v5 += (v5 >> ((504) & 63))
    }
    var v7 float32 = 0.0
    var v8 uint32 = 2
    if len(v1) > 0 {
        v1[0] = v8
    }
    v1 = append(v1, (v8 % ((uint32(v2)) | 1)))
    v2 += 470
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int64 = v5
        var v11 float64 = float64(v7)
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 -= (^(v8))
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
