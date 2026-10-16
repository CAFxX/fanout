package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float32 = -2.25
    var v2 int = -1
    v3 := make(map[string]uint32)
    var v4 int16 = -32768
    v5 := []int{-2147483648}
    fmt.Println(v1, v1)
    v6 := GenS1{F0: 222}
    _ = len(v5)
    v7 := GenS1{F0: 42}
    v2 += 1
    var v8 uint32 = 81
    delete(v3, "a")
    if len(v5) > 0 {
        v5[0] = 32768
    }
    if (932 <= 1) {
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
