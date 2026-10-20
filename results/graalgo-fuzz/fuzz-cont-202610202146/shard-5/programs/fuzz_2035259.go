package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{391}
    v2 := []uint32{4294967295, 14, 65535}
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    if false {
        var v5 bool = (!((127 > 42) || (36 < 1)))
        _ = v5
    }
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint64 = 2
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = float32(float64(0.1))
    _ = len(v1)
    var v10 float32 = (0.0 + (v6 / v6))
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float64 = ((3.14159 / (1.5 / ((0.0 + ((3.14159 + 1.5) * 0.1)) * float64((float32(1e10) / (0.1 + float32(0.1))))))) + 0.1)
        _ = len(v1)
        if (128 == fi11) {
            v12 = ((float64(((0.1 * (float32(float64(v6)) - (v6 + 1.5))) / (v6 + 3.14159))) * float64(float32((v12 + (v12 + -2.25))))) * v12)
            fmt.Println(v10, v12)
            var v13 float32 = float32(v12)
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    var v14 bool = (42 >= -9223372036854775808)
    _ = len(v4)
    v14 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
