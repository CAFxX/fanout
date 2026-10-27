package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (((0.0 - (1.5 - 3.14159)) + -2.25) / (-2.25 - float32((0.1 - 1e10))))
    var v3 float64 = 0.0
    var v4 int = -9223372036854775808
    v5 := []uint32{2147483647, 773, 1}
    var v6 uint16 = 1
    v2 += 0.0
    var v7 uint = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = 255
    v7 += 2
    var v8 int32 = 42
    var v9 uint = v7
    v6 -= v6
    v2 += v2
    v10 := v1[v4]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
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
