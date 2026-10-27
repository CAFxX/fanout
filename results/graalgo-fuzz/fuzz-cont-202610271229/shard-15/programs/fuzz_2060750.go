package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (0.0 - (1.5 + ((3.14159 - 1e10) + ((float32(1.5) + -2.25) * -2.25))))
    var v2 int8 = 100
    var v3 uint = 276
    v4 := make(map[int]bool)
    var v5 float64 = 1.5
    var v6 int64 = 42
    v6 -= (v6 - (v6 * -9223372036854775808))
    var v7 int = 128
    fmt.Println(v3, v6)
    var v8 int16 = 308
    v1 = 1.5
    var v9 uint64 = 1
    var v10 uint = (^(v3))
    var v11 int16 = (-(v8))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
