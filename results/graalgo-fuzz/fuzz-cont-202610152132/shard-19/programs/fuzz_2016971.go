package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int32) int64 {
    v4 := GenS1{F0: 0}
    var v5 int32 = p0
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    v3 := []int{-621, 1}
    v6 := GenS1{F0: 100}
    _ = len(v3)
    var v7 int = -574
    v2 = ""
    var v8 int16 = -906
    var v9 bool = true
    var v10 uint32 = 65535
    if true {
        var v11 float32 = (-2.25 / 1e10)
        _ = v11
    } else {
        var v12 int = (v7 / ((9223372036854775807) | 1))
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
