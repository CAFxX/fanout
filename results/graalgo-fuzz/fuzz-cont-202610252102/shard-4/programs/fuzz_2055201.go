package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool) int {
    p0 = p0
    v6 := GenS1{F0: 42}
    v7 := []int{331, 48, 2147483647, -426}
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{819}
    var v3 float32 = 0.1
    var v4 int = 4294967295
    v5 := []uint32{}
    var v8 float64 = 0.0
    fmt.Println(v1, v8, v1)
    v9 := GenS1{F0: 65535}
    var v10 uint16 = 277
    var v11 uint32 = 2147483648
    var v12 float64 = float64(v3)
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = 457
        switch genFn0(((v14 / ((1) | 1)) > (v14 + (v14 % (((^(v14))) | 1))))) {
        case 2:
            var v15 float32 = -2.25
            _ = v15
        case -2147483648:
            fmt.Println(v14, v8)
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v10)
    if len(v5) > 0 {
        v5[0] = 913
    }
    if true {
        fmt.Println(v4)
        v3 -= v3
    } else {
        var v16 float64 = (v12 - v8)
        _ = v16
    }
    v17 := GenS1{F0: 62}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
