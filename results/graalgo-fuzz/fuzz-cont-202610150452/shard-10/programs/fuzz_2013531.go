package main

import (
    "fmt"
)

func genFn0(p0 uint8, p1 uint, p2 int32) uint8 {
    var v4 int = 4294967295
    var v5 uint32 = 0
    fmt.Println(p0, v4, v4)
    var v6 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 39
}

func main() {
    v1 := []int{4294967295}
    var v2 float32 = 1.5
    var v3 float64 = (1.5 + float64(-2.25))
    var v7 float32 = (-2.25 / 1.0)
    var v8 uint32 = 2147483648
    if len(v1) > 0 {
        v1[0] = 127
    }
    fmt.Println(v8)
    var v9 int8 = -128
    var v10 uint32 = v8
    _ = len(v1)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
