package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 549
    var v2 int8 = -1
    v3 := make(map[uint32]bool)
    var v4 int = 42
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 2147483648
    var v7 int64 = 42
    v7 = (^(v7))
    for fi8 := 0; fi8 < 16; fi8++ {
        v3[(v6 & v6)] = true
        _ = fi8
    }
    var v9 int8 = int8(v1)
    v7 = (^(v7))
    fmt.Println(v7)
    var v10 float64 = ((0.0 + ((1e10 / (1.5 / (3.14159 * ((((1.5 / (3.14159 * float64((1.5 - (float32(float64((0.0 / 1.5))) / 3.14159))))) - (0.1 - float64((0.1 + 3.14159)))) + (1.5 - 0.0)) + float64(1e10))))) + 1e10)) * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
