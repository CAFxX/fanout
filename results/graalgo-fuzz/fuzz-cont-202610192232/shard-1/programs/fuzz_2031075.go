package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    var v4 uint = (v3 >> (((v3 / ((4294967295) | 1))) & 63))
    v5 := make(map[int]uint32)
    v6 := GenS1{F0: int64(v4)}
    v4 = v3
    var v7 float32 = (0.1 / (float32(0.1) + 0.0))
    delete(v2, 128)
    v1 += v7
    fmt.Println(v3, v1, v3)
    _ = len(v5)
    var v8 uint8 = 5
    fmt.Println(v7, v1)
    v8 = 127
    v4 = v3
    v9 := []int{42, 127, -178, 9223372036854775807}
    if false {
        if (!(!true)) {
            fmt.Println(v1, v4)
            fmt.Println(v4, v8)
            v1 += float32(0.1)
        } else {
            fmt.Println(v3, v8)
        }
        v7 = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
