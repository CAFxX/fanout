package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 0
    var v2 int32 = 314
    v3 := make(map[int]int)
    var v4 uint32 = 2147483647
    var v5 uint = 255
    var v6 int8 = 74
    fmt.Println(v2)
    v6 = 42
    var v7 int16 = -32768
    fmt.Println(v6, v7, v1)
    v2 = int32(v4)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := v3[(fi8 << (((v5 + v5)) & 63))]
        _ = fi8
        _ = v9
    }
    delete(v3, 671)
    var v10 int64 = -882
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
