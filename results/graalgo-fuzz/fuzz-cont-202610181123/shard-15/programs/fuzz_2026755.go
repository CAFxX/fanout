package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 1
    v3 := []int{32767, 0}
    v4 := make(map[int]string)
    var v5 bool = (!(-46 >= 42))
    var v6 int8 = 122
    v7 := v4[0]
    if false {
        var v8 float64 = v1
        _ = v8
    }
    if (v5 || v5) {
        var v9 int16 = 1
        _ = v9
    } else {
        if (2 == v2) {
            var v10 uint16 = (v2 << ((4294967295) & 63))
            v5 = (32768 < -605)
            _ = v10
        } else {
            v6 = (v6 * v6)
            var v11 bool = (v1 == v1)
            _ = v11
        }
        var v12 bool = true
        var v13 float32 = 3.14159
        _ = v12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v14 := GenS1{F0: 1}
    if len(v3) > 0 {
        v3[0] = 261
    }
    v6 = (v6 & (v6 + 100))
    var v15 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
