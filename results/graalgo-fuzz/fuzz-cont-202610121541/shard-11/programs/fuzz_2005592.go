package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 0.0
    var v2 int32 = -595
    var v3 int16 = -547
    var v4 float64 = 0.1
    var v5 int32 = (^(v2))
    var v6 float32 = 1e10
    v7 := []int{-9223372036854775808, -2147483648, -796, 32768}
    var v8 int64 = 294
    if true {
        var v9 uint32 = 2147483647
        v7 = append(v7, 32767)
        _ = v9
    }
    v2 += 0
    var v10 float32 = v1
    if false {
        v4 = 0.0
        if false {
            v10 -= 3.14159
            v4 = (v4 * v4)
        }
    } else {
        switch "" {
        case "foo bar":
            var v11 float64 = v4
            _ = v11
        }
    }
    v7 = append(v7, -2147483648)
    var v12 int16 = (v3 >> ((42) & 63))
    var v13 uint8 = 128
    v13 = (v13 & (v13 + (v13 & 108)))
    var v14 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
