package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 uint = 65535
    v3 := []uint32{2147483648}
    var v4 bool = (float32(0.0) == 1.5)
    var v5 uint = 42
    fmt.Println(v5, v5, v2)
    if true {
        var v6 int8 = 1
        var v7 uint32 = 949
        fmt.Println(v4)
        _ = v6
        _ = v7
    } else {
        var v8 int16 = -1
        _ = v8
    }
    v9 := v1[-12]
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 = append(v3, 2147483648)
        _ = fi10
    }
    var v11 float64 = ((1.5 + float64(float32((0.1 * 3.14159)))) - 3.14159)
    for fi12 := 0; fi12 < 6; fi12++ {
        if v4 {
            fmt.Println(fi12, v11, v5)
        }
        _ = fi12
    }
    v3 = append(v3, 1)
    var v13 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
