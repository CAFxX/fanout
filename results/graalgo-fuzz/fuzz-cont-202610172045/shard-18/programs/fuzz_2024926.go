package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (3.14159 + float32((((float64(3.14159) - -2.25) + (0.1 - float64(3.14159))) / 0.1)))
    v2 := []int{-1, 0, 4294967295}
    var v3 uint8 = 42
    v4 := make(map[uint32]int)
    var v5 uint32 = 509
    var v6 uint64 = 467
    v3 -= (v3 - uint8(v6))
    if (0.0 < (float32(0.1) + v1)) {
        delete(v4, uint32(v3))
    }
    var v7 bool = false
    fmt.Println(v6, v5)
    var v8 float32 = v1
    v6 += 4294967295
    for fi9 := 0; fi9 < 7; fi9++ {
        if (42 <= v6) {
            v2 = append(v2, 1)
            var v10 float64 = 3.14159
            _ = v10
        }
        _ = fi9
    }
    v3 = 1
    var v11 int32 = 1
    v3 = uint8(v5)
    var v12 int64 = 928
    fmt.Println(v7, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
