package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    v3 := []uint32{433, 2147483648, 4294967295, 65535}
    var v4 uint64 = 1
    v5 := make(map[int]string)
    if false {
        if (2147483647 != 4294967295) {
            delete(v5, 9223372036854775807)
            var v6 uint32 = 2147483647
            v4 -= v4
            _ = v6
        }
        var v7 uint = 928
        _ = v7
    }
    if true {
        var v8 uint16 = 1
        _ = v8
    }
    var v9 uint64 = v4
    var v10 uint8 = v1
    fmt.Println(v9, v1)
    var v11 int8 = 100
    fmt.Println(v1, v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS1{F0: 58}
    var v13 float32 = (float32(float64((1.5 - 1e10))) * ((0.1 / 3.14159) * (0.0 + (-2.25 + 0.1))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
