package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 bool = (true && (350 != 604))
    var v4 float32 = (1e10 - (0.1 + 1e10))
    var v5 float32 = v4
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            v3 = v3
        }
        if (0 <= 42) {
            v5 = 0.0
            var v7 float64 = float64(float32(-2.25))
            _ = v7
        }
        _ = fi6
    }
    if v3 {
        for fi8 := 0; fi8 < 10; fi8++ {
            v4 = -2.25
            v3 = v3
            var v9 string = "a"
            _ = fi8
            _ = v9
        }
    }
    var v10 int32 = -1
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int64 = 9223372036854775807
        _ = fi11
        _ = v12
    }
    if false {
        var v13 float64 = 1.5
        _ = v13
    }
    v3 = v3
    v5 -= (((-2.25 / (float32(0.1) / ((v4 * 0.1) + v4))) / -2.25) / 3.14159)
    var v14 bool = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
