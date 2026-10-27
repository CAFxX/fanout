package main

import (
    "fmt"
)


func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    v3 := []uint32{}
    var v4 int8 = 84
    var v5 float64 = -2.25
    v4 += (v4 - v4)
    delete(v2, "\n\t")
    v5 = (v5 + (3.14159 / (v5 / v5)))
    var v6 float32 = ((-2.25 * 0.0) + 3.14159)
    var v7 string = ((("a" + "") + "foo bar") + "a")
    var v8 bool = false
    v6 += (v6 + v6)
    var v9 int32 = 2147483647
    var v10 int64 = 229
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!((v10 & v10) < v10)) {
        var v11 int = -822
        _ = v11
    }
    v10 += v10
    v2[v7] = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
