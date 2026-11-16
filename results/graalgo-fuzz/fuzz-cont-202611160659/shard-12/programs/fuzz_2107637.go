package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 int16, p2 bool) uint8 {
    if p2 {
        var v7 int = -1
        _ = v7
    }
    v8 := []int{-1, -50, 128, -2147483648}
    v9 := GenS1{F0: 65535}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    _ = v9
    return 107
}
func genFn1() int {
    v13 := []int{32768, 2147483647}
    var v14 int64 = -1
    _ = len(v13)
    switch 753 {
    case -2147483648:
        var v15 int16 = 32767
        _ = v15
    }
    v16 := GenS1{F0: 995}
    _ = v13
    _ = v14
    _ = v16
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{1, 1, 804, 42}
    v3 := []uint32{}
    var v4 uint64 = 783
    var v5 string = ("\n\t" + "a")
    v6 := make(map[uint32]int)
    if len(v3) > 0 {
        v3[0] = 42
    }
    v4 += v4
    var v17 bool = (true || false)
    var v18 uint8 = 1
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v19 float32 = 0.0
    v17 = true
    v20 := GenS2{F0: -128, F1: -48, F2: 42}
    v17 = (!v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
