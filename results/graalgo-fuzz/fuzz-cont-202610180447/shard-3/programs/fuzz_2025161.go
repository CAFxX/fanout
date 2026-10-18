package main

import (
    "fmt"
)


func main() {
    var v1 float64 = 0.1
    var v2 int64 = 692
    var v3 int64 = -1
    var v4 uint32 = 2147483647
    var v5 int64 = (v2 * v3)
    v5 -= v3
    fmt.Println(v5, v4)
    v6 := []int{763, 128, 127}
    v5 = (^(v2))
    if len(v6) > 0 {
        v6[0] = 644
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v1, v1)
    v4 += v4
    v6 = append(v6, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
