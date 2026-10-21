package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 int8) uint32 {
    p0 += (^(p0))
    p0 = p0
    var v7 int32 = 42
    _ = p0
    _ = v7
    return 985
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 float32 = 0.1
    var v4 uint32 = 42
    var v5 float32 = v3
    var v6 int = 1
    var v8 int8 = 100
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float64 = 3.14159
        v11 := GenS1{F0: 4294967295}
        switch (v6 << ((18446744073709551615) & 63)) {
        case 128:
            v3 = 3.14159
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v2[1] = v6
    v12 := []int{v6, v6, (v6 ^ v6)}
    v12 = append(v12, -2147483648)
    switch (v6 ^ v6) {
    case 1:
        var v13 int = (v6 ^ 32767)
        _ = v13
    case 9223372036854775807:
        v8 -= v8
    default:
        v8 = (v8 >> ((18446744073709551615) & 63))
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v6 -= fi14
        if false {
            fmt.Println(v4, v4, v3)
            var v15 uint32 = (v4 % (((^(v4))) | 1))
            _ = v15
        }
        var v16 uint16 = 255
        _ = fi14
        _ = v16
    }
    v17 := GenS2{F0: true, F1: (true && false)}
    if (0 == 885) {
        var v18 int64 = 615
        _ = len(v2)
        fmt.Println(v5, v18)
        _ = v18
    }
    v19 := v2[(v4 ^ v4)]
    v6 = v19
    v20 := GenS2{F0: (!false), F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
