package main

import (
    "fmt"
)

func genFn0(p0 int16) int {
    p0 -= (p0 - p0)
    p0 -= p0
    _ = p0
    return -725
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    var v4 int64 = (-(v3))
    fmt.Println(v3, v4)
    var v5 float32 = 0.0
    delete(v2, 2147483648)
    fmt.Println(v5)
    var v6 uint8 = 255
    var v7 int16 = 611
    var v8 int32 = 2147483647
    v3 = v3
    if false {
        for fi9 := 0; fi9 < 10; fi9++ {
            v3 = 9223372036854775807
            _ = fi9
        }
    } else {
        var v10 int8 = 100
        _ = v10
    }
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
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
