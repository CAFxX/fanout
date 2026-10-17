package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{106, 0, 4294967295}
    var v2 bool = (0.1 < (1.5 + 0.1))
    v3 := make(map[uint32]uint32)
    var v4 int32 = 583
    v5 := make(map[uint32]int)
    var v6 uint16 = 724
    v2 = true
    fmt.Println(v4)
    if v2 {
        fmt.Println(v6, v4)
        var v7 uint32 = 2
        _ = v7
    }
    var v8 int = int(v6)
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 bool = false
            var v13 int = (v8 / ((v8) | 1))
            var v14 int = (-(v8))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 int = 45
        var v16 string = (v9 + v9)
        _ = fi10
        _ = v15
        _ = v16
    }
    v17 := v3[2147483647]
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
