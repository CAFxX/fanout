package main

import (
    "fmt"
)


func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 uint64 = v3
    for fi6 := 0; fi6 < 16; fi6++ {
        v3 = 18446744073709551615
        v7 := v2["\n\t"]
        _ = fi6
        _ = v7
    }
    v1 = (^(v1))
    v8 := []int{0, 0}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = 127
        _ = fi9
        _ = v10
    }
    delete(v4, 42)
    var v11 float64 = 1.5
    v3 -= (v5 * (^(v5)))
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    v11 = float64(3.14159)
    var v12 bool = false
    v1 = 65535
    for fi13 := 0; fi13 < 20; fi13++ {
        v11 = float64(-2.25)
        _ = fi13
    }
    var v14 float64 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
