package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 2
    v2 := []uint32{0}
    v3 := []string{}
    var v4 uint64 = 310
    var v5 float64 = ((0.1 + 1e10) * 0.1)
    v6 := []int{302, 32768, 128, -2147483648}
    v7 := GenS1{F0: -1}
    v5 += 0.1
    v4 = v4
    v8 := GenS1{F0: 32767}
    if false {
        var v9 int = -9223372036854775808
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 += 255
            _ = fi10
        }
        _ = v9
    } else {
        v11 := GenS1{F0: 384}
        _ = v11
    }
    v4 = v4
    var v12 uint8 = 42
    var v13 float32 = (1.5 * 3.14159)
    v13 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
