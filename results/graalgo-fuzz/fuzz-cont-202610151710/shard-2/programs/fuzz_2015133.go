package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{508, 2147483647, 43, 42}
    v2 := make(map[uint32]int)
    var v3 uint64 = 42
    var v4 uint8 = 42
    var v5 uint = 372
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7)
        var v8 float64 = v6
        _ = fi7
        _ = v8
    }
    var v9 uint8 = v4
    switch 917 {
    case 2147483647:
        _ = len(v2)
    case 4294967295:
        v3 = v3
    case 1:
        var v10 int = -9223372036854775808
        _ = v10
    }
    v5 += v5
    v1 = append(v1, 42)
    var v11 float32 = (float32((v6 - 1.5)) * 3.14159)
    var v12 uint64 = uint64(v5)
    if true {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint32 = 2
            var v15 int16 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        if (true || false) {
            fmt.Println(v9, v3)
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint = 65535
                _ = fi16
                _ = v17
            }
            var v18 int16 = 35
            _ = v18
        }
    } else {
        var v19 uint64 = 258
        v5 -= v5
        _ = v19
    }
    var v20 float32 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
