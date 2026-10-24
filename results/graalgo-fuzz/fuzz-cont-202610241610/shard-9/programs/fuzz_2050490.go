package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 int8) uint32 {
    v4 := GenS1{F0: 42, F1: 149}
    var v5 uint8 = 128
    v5 = (^(v5))
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 += v5
        _ = fi6
    }
    fmt.Println(v5)
    _ = p0
    _ = v4
    _ = v5
    return 19
}

func main() {
    var v1 uint8 = 161
    var v2 string = ""
    v3 := make(map[int]bool)
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        v2 = ""
        v8 := v3[(fi7 >> ((562) & 63))]
        switch fi7 {
        case -2147483648:
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: 65535, F1: -1}
                _ = fi9
                _ = v10
            }
        case -476:
            var v11 int32 = 2147483647
            _ = v11
        default:
            var v12 int64 = 9223372036854775807
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v13 int = 4294967295
    var v14 uint16 = 889
    v15 := GenS2{F0: v13, F1: 220, F2: -507}
    v16 := []int{}
    var v17 int8 = 45
    switch (v13 / ((-9223372036854775808) | 1)) {
    case -1:
        var v18 int16 = 1000
        _ = v18
    }
    fmt.Println(v14, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
