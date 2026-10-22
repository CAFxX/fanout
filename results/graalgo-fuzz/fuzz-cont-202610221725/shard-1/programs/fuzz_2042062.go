package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 int8) uint32 {
    p0 = p0
    var v4 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int32 = -906
    var v2 uint16 = 32767
    var v3 int16 = 42
    v5 := []int{-9223372036854775808, 32768, 4294967295}
    v6 := GenS1{F0: -1}
    v5 = append(v5, 2147483647)
    if (((((1e10 / (1.5 + float32(3.14159))) - 3.14159) - 1e10) - -2.25) < (0.0 - float32(-2.25))) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint16 = (^(v2))
            v9 := GenS1{F0: fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint64 = 4294967295
        var v11 uint32 = 122
        var v12 float64 = 0.0
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int64 = 656
    v14 := GenS2{F0: 32768}
    v15 := GenS2{F0: -1}
    var v16 float32 = float32(3.14159)
    v2 += 42
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int8 = -1
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
