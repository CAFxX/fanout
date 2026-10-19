package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{0, 2}
    var v2 uint16 = 32767
    var v3 int64 = 0
    var v4 uint32 = 4294967295
    var v5 uint64 = 42
    v6 := []int{-9223372036854775808, 127, 9223372036854775807}
    var v7 uint16 = (v2 >> ((1) & 63))
    if (false && false) {
        var v8 uint64 = v5
        _ = v8
    } else {
        if len(v6) > 0 {
            v6[0] = -934
        }
        var v9 int = 9223372036854775807
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(fi10, fi10)
        var v11 uint = 574
        _ = fi10
        _ = v11
    }
    v4 -= 880
    var v12 float32 = 1.5
    if len(v1) > 0 {
        v1[0] = v4
    }
    v12 += (1.5 * 3.14159)
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
