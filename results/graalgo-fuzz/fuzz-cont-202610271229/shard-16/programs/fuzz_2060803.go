package main

import (
    "fmt"
)


func main() {
    var v1 float64 = 1e10
    v2 := []int{-479, -9223372036854775808, 128, -218}
    v3 := make(map[string]int)
    var v4 float64 = ((v1 * 0.0) - v1)
    v5 := make(map[int]uint32)
    var v6 uint64 = 196
    var v7 uint8 = 160
    fmt.Println(v1)
    fmt.Println(v7, v6)
    var v8 uint32 = 2147483648
    v1 = -2.25
    var v9 int64 = -1
    if true {
        v4 = float64(-2.25)
    }
    var v10 uint64 = (v6 % ((2) | 1))
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            var v12 uint8 = 107
            v13 := v3["\n\t"]
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    v3[("a" + "a")] = 128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
