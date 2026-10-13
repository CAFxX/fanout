package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]int)
    var v2 int64 = -450
    var v3 bool = (0 == 568)
    v4 := make(map[uint32]uint32)
    fmt.Println(v2, v3)
    var v5 int32 = 953
    if v3 {
        delete(v4, 2147483647)
        v6 := v1[-9223372036854775808]
        _ = v6
    }
    if (((float64(1.5) + ((0.1 + -2.25) * 0.1)) != (3.14159 - 0.1)) && v3) {
        fmt.Println(v3)
        v7 := []int{4294967295, 653, 2147483647, 18}
        _ = v7
    } else {
        var v8 float64 = 0.1
        _ = v8
    }
    var v9 uint = 4294967295
    v5 = (v5 - -1)
    var v10 int64 = (v2 ^ (v2 / ((-1) | 1)))
    var v11 int64 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
