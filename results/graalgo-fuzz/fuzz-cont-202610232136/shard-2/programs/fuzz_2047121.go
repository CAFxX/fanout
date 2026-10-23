package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
func genFn0() string {
    v5 := []int{}
    var v6 uint32 = 42
    var v7 int64 = 1
    var v8 int32 = 42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + "a")
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]int)
    v3 := []uint32{2147483648, 2}
    v4 := make(map[uint32]bool)
    v9 := GenS1{F0: 176, F1: 65535, F2: 0}
    v10 := GenS1{F0: 2147483647, F1: 0, F2: 255}
    var v11 int32 = 683
    var v12 uint = 7
    v13 := GenS2{F0: 127, F1: 2147483647, F2: 1}
    for fi14 := 0; fi14 < 15; fi14++ {
        v12 += v12
        if true {
            var v15 int64 = -848
            var v16 uint64 = uint64(v12)
            _ = v15
            _ = v16
        }
        v1 += (((v1 / v1) - v1) - v1)
        _ = fi14
    }
    var v17 uint64 = 18446744073709551615
    fmt.Println(v12, v17, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
