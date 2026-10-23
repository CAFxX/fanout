package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 411
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    var v4 int16 = 1000
    var v5 uint64 = (v2 ^ (v2 >> ((436) & 63)))
    v1 = (v1 + (v1 % ((v1) | 1)))
    var v6 int8 = -51
    v5 += v2
    fmt.Println(v5)
    fmt.Println(v4, v6, v4)
    var v7 uint64 = 42
    var v8 float32 = 0.0
    if (true && false) {
        var v9 uint32 = 1
        if (!(1 == 0)) {
            fmt.Println(v4, v5, v1)
        }
        _ = v9
    } else {
        v2 = v2
        v5 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
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
