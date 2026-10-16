package main

import (
    "fmt"
)

func genFn0(p0 uint8) int64 {
    var v6 float32 = 0.1
    v6 += v6
    _ = p0
    _ = v6
    return 472
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = -342
    var v3 int32 = -1
    var v4 uint32 = 2
    v5 := make(map[int]uint32)
    var v7 uint64 = 1
    v7 = v7
    v3 += v3
    fmt.Println(v7, v3, v3)
    switch 127 {
    case -190:
        delete(v5, 42)
    case 42:
        v4 = (v4 * (v4 & (v4 % ((v4) | 1))))
    }
    switch 4294967295 {
    case 4294967295:
        v1 = v1
    }
    v4 += v4
    if (false || (!false)) {
        var v8 int8 = 1
        fmt.Println(v7)
        _ = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
