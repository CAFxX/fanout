package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 476
    var v2 int16 = -955
    v3 := []uint32{}
    var v4 int32 = -1
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v5 int32 = v4
    v3 = append(v3, 940)
    v5 = v5
    if (false || true) {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(v5)
            fmt.Println(v5)
            _ = fi6
        }
    }
    var v7 uint32 = 0
    var v8 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
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
