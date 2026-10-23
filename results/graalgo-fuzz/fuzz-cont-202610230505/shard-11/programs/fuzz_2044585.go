package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    var v3 int8 = 79
    var v4 uint8 = 2
    var v5 int8 = (-(v3))
    var v6 int64 = int64(v2)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int32 = (v2 % (((v2 + v2)) | 1))
        var v9 uint = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 2147483647, F1: 9223372036854775807}
    var v11 uint8 = (v4 / ((248) | 1))
    var v12 uint8 = (v11 / ((v11) | 1))
    var v13 float32 = 3.14159
    v12 = v4
    fmt.Println(v4, v6)
    fmt.Println(v6, v13)
    v6 = 42
    var v14 float64 = (0.1 + 0.1)
    fmt.Println(v3, v5)
    v12 = (v4 | v12)
    v4 = v4
    var v15 string = ""
    v16 := []int{141}
    var v17 string = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
