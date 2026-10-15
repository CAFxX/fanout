package main

import (
    "fmt"
)


func main() {
    var v1 float64 = ((float64(1e10) + 1.5) - (-2.25 + float64(0.1)))
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    v1 = 1e10
    var v4 uint32 = 65535
    _ = len(v3)
    if (127 != 42) {
        var v5 float64 = float64((0.0 / 0.1))
        _ = v5
    } else {
        var v6 uint64 = 716
        v2 += (v2 >> ((42) & 63))
        _ = v6
    }
    v1 = v1
    v7 := []int{9223372036854775807, 0}
    fmt.Println(v2)
    v4 = v4
    switch "a" {
    case "foo bar":
        v2 = 222
    case "":
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    fmt.Println(v1, v2)
    v2 += v2
    v1 += float64((0.1 + float32(((v1 - (v1 * (v1 - v1))) + (3.14159 - v1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
