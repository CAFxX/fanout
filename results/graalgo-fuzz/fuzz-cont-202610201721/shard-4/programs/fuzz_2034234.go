package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 32767
    var v2 int64 = 1
    v3 := make(map[string]int)
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    var v6 float32 = ((((0.0 * 0.0) * 1e10) * (3.14159 + ((1.5 - 1.5) + 3.14159))) / 1.0)
    for fi7 := 0; fi7 < 7; fi7++ {
        v1 -= v1
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    fmt.Println(v1, v4)
    var v10 int8 = 100
    fmt.Println(v4, v9)
    v10 += v10
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
