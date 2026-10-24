package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := []int{}
    var v3 float32 = 3.14159
    var v4 int64 = 0
    var v5 float64 = 0.0
    v5 += v5
    v4 += 0
    v4 += 42
    var v6 float64 = 1e10
    var v7 uint64 = 406
    var v8 float64 = float64(v3)
    var v9 int64 = -691
    if true {
        switch 32768 {
        case 1:
            v10 := GenS1{F0: 0, F1: 2, F2: 4294967295}
            _ = v10
        case 127:
            var v11 uint8 = 1
            _ = v11
        case 0:
            v2 = append(v2, 9223372036854775807)
        }
        v4 = v4
        v12 := GenS1{F0: 94, F1: 95, F2: 2}
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS1{F0: 4, F1: 166, F2: v7}
        var v15 int = (fi13 / (((fi13 + fi13)) | 1))
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint8 = 2
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
