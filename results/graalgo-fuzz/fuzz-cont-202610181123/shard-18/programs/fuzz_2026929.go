package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 int }

func main() {
    var v1 int32 = -349
    var v2 int = 42
    var v3 int = (v2 << ((0) & 63))
    fmt.Println(v2)
    var v4 bool = true
    var v5 float64 = (1.5 + float64(1e10))
    v6 := GenS1{F0: 2, F1: (v1 / (((v1 | v1)) | 1))}
    var v7 int64 = 42
    switch (v2 / (((v3 * (^(v2)))) | 1)) {
    case 128:
        v8 := GenS2{F0: 4294967295, F1: (v7 << ((882) & 63)), F2: (v3 + 1)}
        _ = v8
    default:
        if (0 <= (v1 / ((v1) | 1))) {
            var v9 float32 = -2.25
            _ = v9
        }
    }
    v3 += (^(v2))
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{32767, fi10, -495}
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
