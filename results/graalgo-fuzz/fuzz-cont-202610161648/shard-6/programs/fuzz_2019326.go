package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }

func main() {
    var v1 float32 = (1e10 - 1.5)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 int = 749
    var v5 uint32 = 2
    var v6 uint16 = 44
    var v7 float32 = v1
    var v8 bool = false
    _ = len(v3)
    var v9 int = v4
    var v10 int64 = 42
    fmt.Println(v9, v10, v9)
    v11 := v3[v4]
    var v12 int32 = 42
    v13 := GenS1{F0: v4, F1: (v10 << ((255) & 63)), F2: -1}
    if true {
        var v14 int32 = (v12 / (((^(v12))) | 1))
        if v8 {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int = -9223372036854775808
                var v17 int64 = v10
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: -563, F1: int64(v14), F2: -128}
            _ = v18
        }
        v12 += 0
        _ = v14
    } else {
        v9 += (v9 << ((4294967295) & 63))
    }
    v19 := GenS1{F0: (^(v9)), F1: (v10 * (^(v10))), F2: 116}
    v12 -= (v12 + (v12 >> ((42) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
