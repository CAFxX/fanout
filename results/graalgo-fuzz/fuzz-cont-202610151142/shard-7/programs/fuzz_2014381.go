package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 24
    var v3 uint64 = 0
    var v4 uint = 4294967295
    v5 := make(map[string]int)
    var v6 float64 = ((3.14159 / 1.5) + ((float64((1e10 * (1e10 + 1.5))) + (3.14159 - -2.25)) * (-2.25 * 1e10)))
    v7 := GenS1{F0: 0, F1: (v2 | (v2 >> (((v4 >> ((v4) & 63))) & 63))), F2: 1}
    fmt.Println(v3, v2)
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v5["hello"] = (-(fi9))
            v4 += (v4 << ((42) & 63))
            _ = fi9
        }
        var v10 int16 = 42
        v6 -= v6
        _ = fi8
        _ = v10
    }
    var v11 int = -45
    fmt.Println(v6)
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: 4294967295, F1: (v2 - 0), F2: 17}
        var v14 uint16 = 535
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float64 = 0.0
    var v16 float32 = 3.14159
    var v17 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
