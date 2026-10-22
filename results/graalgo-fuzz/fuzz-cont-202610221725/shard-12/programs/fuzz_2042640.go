package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 uint64, p2 bool) uint64 {
    fmt.Println(p1)
    var v5 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1(p0 uint16) uint32 {
    fmt.Println(p0)
    var v6 float64 = 0.1
    _ = p0
    _ = v6
    return 310
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 42
    var v3 uint64 = 18446744073709551615
    var v4 int8 = -1
    v2 += -702
    if ((-2.25 * 0.0) == -2.25) {
        var v7 int8 = v4
        _ = v7
    }
    switch "a" {
    case "":
        var v8 uint8 = 127
        _ = v8
    case "hello":
        v2 = v2
    default:
        var v9 int64 = (v2 | -1)
        _ = v9
    }
    var v10 float64 = -2.25
    fmt.Println(v2, v10, v3)
    if ((v4 / (((v4 >> ((65535) & 63))) | 1)) != (v4 >> ((971) & 63))) {
        var v11 int16 = 32767
        _ = v11
    }
    fmt.Println(v4, v4)
    var v12 uint32 = 4294967295
    var v13 int64 = v2
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
