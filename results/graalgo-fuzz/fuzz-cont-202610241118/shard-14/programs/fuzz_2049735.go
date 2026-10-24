package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{0, 547, 473, 964}
    var v3 float64 = 1e10
    var v4 uint = 0
    var v5 int32 = 1
    v5 -= v5
    var v6 uint = 65535
    var v7 uint16 = 0
    fmt.Println(v6)
    v7 = v7
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int16 = -579
            var v11 float64 = v3
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    _ = len(v2)
    fmt.Println(v7)
    v4 = v4
    fmt.Println(v6, v7, v1)
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
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
