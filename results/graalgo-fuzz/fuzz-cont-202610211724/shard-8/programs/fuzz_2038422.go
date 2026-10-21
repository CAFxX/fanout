package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{977, 4294967295, 42, 42}
    var v3 int8 = 127
    v4 := make(map[int]uint32)
    var v5 float64 = (((0.0 + ((0.1 / 1e10) / -2.25)) + (3.14159 + float64(3.14159))) * (((1.5 / 3.14159) - float64(0.0)) / 0.1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 uint = 65535
    fmt.Println(v6)
    var v7 string = "foo bar"
    var v8 bool = (true || true)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 string = v7
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = v9
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    fmt.Println(v7, v9, v7)
    v9 = ((v7 + v9) + v9)
    v3 += v3
    v3 -= v3
    v4[0] = 1
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
