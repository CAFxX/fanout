package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = -2147483648
    var v2 uint = 590
    var v3 int8 = 1
    var v4 int64 = 0
    v5 := make(map[int]uint32)
    v6 := []int{749, v1, (^(v1)), (v1 % (((v1 - (v1 * (v1 / ((v1) | 1))))) | 1))}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int16 = 883
        _ = fi7
        _ = v8
    }
    v4 -= (v4 ^ (v4 << (((^(v2))) & 63)))
    _ = len(v6)
    v9 := GenS1{F0: 42}
    delete(v5, (v1 | v1))
    var v10 float64 = 1e10
    if true {
        switch v1 {
        case -56:
            var v11 float64 = ((v10 + v10) * 1e10)
            _ = v11
        case -1:
            var v12 int64 = (v4 & v4)
            _ = v12
        case 9223372036854775807:
            var v13 int16 = -32768
            _ = v13
        }
    }
    var v14 uint = v2
    v4 += -825
    v1 = v1
    var v15 uint = v14
    fmt.Println(v14, v10, v15)
    var v16 float64 = 1e10
    var v17 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
