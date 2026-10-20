package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0(p0 uint64) int64 {
    v7 := GenS1{F0: 0, F1: (p0 % (((p0 << ((65535) & 63))) | 1))}
    if true {
        var v8 uint16 = 465
        if ((p0 >> ((2) & 63)) != (p0 << ((255) & 63))) {
            var v9 int = 4294967295
            var v10 float64 = 0.0
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = v11
        _ = v8
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 1
    v4 := make(map[int]bool)
    var v5 int64 = 0
    v6 := make(map[uint32]string)
    _ = len(v4)
    var v12 int32 = 42
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 += v3
        _ = fi13
    }
    v2 = (v5 - v2)
    v14 := GenS1{F0: -390, F1: v3}
    v15 := GenS1{F0: -131, F1: (v3 % ((18446744073709551615) | 1))}
    if true {
        var v16 uint64 = (v3 * v3)
        _ = v16
    }
    var v17 float32 = v1
    v18 := GenS1{F0: 128, F1: 1}
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := []int{(fi19 << ((18446744073709551615) & 63)), fi19, fi19, (fi19 >> ((316) & 63))}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
