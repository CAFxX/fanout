package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 4
    v2 := []uint32{815, 0, 489, 4294967295}
    var v3 float64 = ((1.5 + ((1.5 - 3.14159) * (3.14159 + -2.25))) * (float64(0.1) + float64(3.14159)))
    var v4 uint64 = 4294967295
    v3 -= float64(float32(3.14159))
    v3 += (v3 - v3)
    var v5 bool = true
    var v6 int64 = -861
    v3 = v3
    fmt.Println(v5, v4, v5)
    if ((true && (0 < 4294967295)) && false) {
        v1 += v1
        v3 = (0.0 + v3)
        var v7 uint64 = (v4 * 863)
        _ = v7
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        fmt.Println(v5)
    }
    var v8 string = "\n\t"
    fmt.Println(v1, v6, v5)
    v5 = (42 == 128)
    var v9 uint = 42
    var v10 int32 = 1
    if len(v2) > 0 {
        v2[0] = 282
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
