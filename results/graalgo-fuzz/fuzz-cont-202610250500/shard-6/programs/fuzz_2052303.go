package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0() uint32 {
    v5 := GenS1{F0: -1, F1: 864}
    var v6 uint = 4294967295
    var v7 uint = (v6 * (^(v6)))
    var v8 uint = (v7 * v6)
    fmt.Println(v7)
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 bool, p1 uint64) uint64 {
    v9 := []int{}
    var v10 int32 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (p1 + 4294967295)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 int = 42
    var v11 int8 = 100
    fmt.Println(v11, v11)
    fmt.Println(v4, v11)
    var v12 float32 = 0.0
    v13 := v1["hello"]
    _ = len(v3)
    switch (v4 * -2147483648) {
    case 339:
        for fi14 := 0; fi14 < 8; fi14++ {
            switch (-(v4)) {
            case 9223372036854775807:
                if false {
                    var v15 int32 = 42
                    _ = v15
                }
            case 4294967295:
                v12 -= float32((1e10 - (1e10 + 0.0)))
            default:
                if v13 {
                    fmt.Println(v4, v12)
                }
            }
            var v16 int32 = 0
            _ = fi14
            _ = v16
        }
    case 206:
        var v17 int = v4
        _ = v17
    case -866:
        v18 := GenS1{F0: 722, F1: 261}
        _ = v18
    default:
        _ = len(v2)
    }
    var v19 int32 = 1
    var v20 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
