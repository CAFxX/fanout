package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = 417
    var v2 uint16 = 442
    var v3 int8 = -1
    v4 := []uint32{}
    var v5 float64 = 0.0
    var v6 uint16 = v2
    v5 = v5
    if len(v4) > 0 {
        v4[0] = 215
    }
    switch -634 {
    case 127:
        fmt.Println(v1)
    case -2147483648:
        v3 += (v3 >> ((1) & 63))
    }
    if (true && (1e10 < ((0.1 - 3.14159) + (1.5 / 0.1)))) {
        fmt.Println(v2, v2, v5)
        if false {
            var v7 uint64 = 336
            fmt.Println(v5)
            _ = v7
        }
        v1 = v1
    }
    var v8 bool = false
    var v9 bool = (255 > 0)
    v5 = ((v5 / -2.25) + ((v5 - v5) * 1e10))
    v8 = true
    v10 := GenS1{F0: int8(v1)}
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            v9 = v9
            var v12 float32 = (float32(0.1) + float32(v5))
            _ = fi11
            _ = v12
        }
        var v13 uint8 = 2
        _ = v13
    }
    var v14 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
