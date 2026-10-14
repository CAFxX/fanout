package main

import (
    "fmt"
)


func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 220
    var v3 uint16 = 2
    var v4 string = "a"
    v5 := []int{128, 731}
    fmt.Println(v4, v4)
    fmt.Println(v3, v4)
    _ = len(v5)
    var v6 float32 = 0.0
    if len(v5) > 0 {
        v5[0] = 42
    }
    v5 = append(v5, int(v2))
    v1 = v1
    var v7 uint64 = 0
    if (949 == uint(v3)) {
        if false {
            var v8 uint32 = 9
            v8 += uint32(v2)
            var v9 uint = 42
            _ = v8
            _ = v9
        }
        fmt.Println(v6, v1)
    } else {
        var v10 int32 = 0
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
