package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int64 {
    v5 := []int{484, -263, -2147483648, 2147483647}
    switch 528 {
    case -9223372036854775808:
        v6 := GenS1{F0: 0}
        _ = v6
    }
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    var v4 uint = 0
    var v7 int64 = v3
    var v8 float32 = ((0.0 - (((-2.25 - 0.0) - 1.5) * float32((-2.25 * 1.5)))) / float32(float64(0.1)))
    v9 := []int{}
    switch "foo bar" {
    case "foo bar":
        if len(v9) > 0 {
            v9[0] = 2147483647
        }
    }
    v8 = v8
    var v10 float32 = (v8 + v8)
    fmt.Println(v3)
    var v11 uint64 = 4294967295
    delete(v2, "")
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
