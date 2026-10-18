package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int, p2 uint16) int8 {
    var v5 string = "a"
    v6 := []int{4294967295, p1, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -83
}

func main() {
    var v1 float32 = ((0.0 + -2.25) + (1.5 / 1.0))
    var v2 int8 = 127
    var v3 int16 = -1
    var v4 uint64 = 2
    var v7 string = ("foo bar" + "hello")
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v3, fi8)
        _ = fi8
    }
    fmt.Println(v3, v1)
    fmt.Println(v2, v4, v3)
    var v9 int = 4294967295
    v1 = v1
    v2 = (v2 / ((genFn0(-1, v9, 171)) | 1))
    v10 := []int{(^(v9)), v9}
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
