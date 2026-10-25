package main

import (
    "fmt"
)

func genFn0(p0 string, p1 uint8) int32 {
    p1 -= p1
    var v7 float64 = 0.1
    fmt.Println(v7, v7)
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    var v4 float32 = (v2 / v2)
    var v5 int32 = 1
    v6 := []int{1}
    var v8 uint32 = 754
    v8 = (^(v8))
    if true {
        fmt.Println(v4)
        var v9 uint8 = 1
        var v10 uint = 2
        _ = v10
        _ = v9
    }
    v4 = v4
    if len(v6) > 0 {
        v6[0] = -1
    }
    var v11 int = 127
    var v12 string = "foo bar"
    if false {
        delete(v3, (v8 - v8))
        var v13 int8 = int8(v11)
        _ = v13
    }
    v2 += float32(((0.1 + (3.14159 / 0.1)) + (0.0 - 3.14159)))
    fmt.Println(v11)
    v12 = ("" + "")
    var v14 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
