package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 int64) uint16 {
    var v6 uint16 = 223
    var v7 float32 = float32(3.14159)
    v8 := GenS1{F0: 1, F1: 729, F2: 4294967295}
    v9 := []int{0}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v6
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    var v3 float64 = float64((1.5 + 1e10))
    var v4 uint8 = 0
    var v5 int = -788
    if false {
        var v10 uint64 = 2
        _ = v10
    }
    v3 = (v3 + ((float64((1e10 - float32((((v3 - v3) / 1e10) - v3)))) + (0.0 * float64(1.5))) * 0.1))
    fmt.Println(v5, v1)
    v3 = v3
    v11 := GenS1{F0: 517, F1: 0, F2: 4294967295}
    fmt.Println(v3)
    v5 = v5
    v12 := v2[42]
    v3 = ((float64(1e10) * float64((1e10 + 3.14159))) * v3)
    var v13 bool = false
    var v14 int32 = v1
    fmt.Println(v1, v14, v14)
    delete(v2, 380)
    v15 := GenS1{F0: 368, F1: 42, F2: 42}
    switch v5 {
    case 32768:
        v2[2147483647] = v5
    default:
        var v16 int32 = v14
        _ = v16
    }
    switch v5 {
    case 32767:
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := []int{(v5 / ((430) | 1)), v12, (fi18 & (v5 ^ (v12 >> ((437) & 63))))}
                _ = fi18
                _ = v19
            }
            v20 := GenS1{F0: 2, F1: 4294967295, F2: 2}
            _ = fi17
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
