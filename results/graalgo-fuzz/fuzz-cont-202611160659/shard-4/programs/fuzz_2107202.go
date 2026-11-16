package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0() int64 {
    v5 := []int{-2147483648, 32768, -1, -9223372036854775808}
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint16 = 508
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 42
        var v10 uint8 = 42
        _ = fi6
        _ = v10
        _ = v9
    }
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = (^(v1))
    v3 := make(map[uint32]int)
    var v4 bool = true
    for fi11 := 0; fi11 < 2; fi11++ {
        switch 127 {
        case -9223372036854775808:
            v2 = v1
        default:
            var v12 int64 = 9223372036854775807
            _ = v12
        }
        v13 := GenS1{F0: v2, F1: (fi11 >> ((65535) & 63)), F2: 1}
        _ = fi11
        _ = v13
    }
    if v4 {
        v14 := GenS1{F0: v1, F1: 32767, F2: 0}
        var v15 float64 = 0.1
        _ = v14
        _ = v15
    }
    delete(v3, v2)
    delete(v3, v1)
    v16 := GenS1{F0: (v2 ^ 2), F1: 42, F2: 65535}
    v2 += v1
    v3[904] = 314
    fmt.Println(v4, v4, v1)
    v2 = (^(v2))
    fmt.Println(v4)
    var v17 bool = (1 <= 2)
    v18 := []int{-435}
    var v19 float32 = (1e10 - 1.5)
    fmt.Println(v4, v17)
    var v20 int = 42
    for fi21 := 0; fi21 < 8; fi21++ {
        v2 = v1
        _ = fi21
    }
    var v22 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
