package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 2
    var v3 int32 = 0
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 10
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 -= 94
        _ = fi7
    }
    var v8 uint8 = v6
    v9 := v4[510]
    v10 := GenS1{F0: -157}
    v9 = "a"
    var v11 int8 = -60
    if false {
        v12 := GenS1{F0: 32768}
        var v13 int = 9223372036854775807
        _ = v12
        _ = v13
    }
    var v14 float64 = v1
    v1 = (3.14159 - (v1 / v14))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
