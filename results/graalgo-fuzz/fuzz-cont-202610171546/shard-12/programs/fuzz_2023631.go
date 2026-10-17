package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 0
    var v2 int = 1
    v3 := make(map[string]int)
    var v4 int16 = 42
    var v5 uint32 = 2147483647
    var v6 float32 = (3.14159 + 0.0)
    var v7 bool = (472 == v4)
    v4 -= 1
    v2 -= (^(v2))
    if ((v1 >> ((65535) & 63)) >= v1) {
        var v8 int32 = int32(v4)
        _ = v8
    }
    fmt.Println(v1)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float64 = 1e10
        v11 := []int{}
        switch v2 {
        case 184:
            var v12 uint8 = 1
            _ = v12
        case 32767:
            var v13 int = fi9
            _ = v13
        default:
            v1 = (^(v1))
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v4, v1)
    var v14 int32 = 299
    fmt.Println(v2, v7, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
