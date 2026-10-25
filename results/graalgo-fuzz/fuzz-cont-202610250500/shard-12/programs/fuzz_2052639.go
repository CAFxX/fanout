package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 bool) int32 {
    var v7 uint8 = 127
    p0 -= p0
    var v8 uint8 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 500
}

func main() {
    v1 := []uint32{2}
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v6 := []string{"hello", "", (("" + "hello") + "\n\t")}
    fmt.Println(v2)
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v2, v2)
        var v10 float64 = float64(float32((float64(0.0) / 1.5)))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v2 = (^(v2))
        v4["a"] = false
        _ = fi11
    }
    var v12 uint16 = 0
    v6 = append(v6, "a")
    v2 = (v2 | 1)
    v2 = 825
    var v13 int16 = -797
    v14 := GenS1{F0: v12, F1: v12}
    v12 = v12
    v2 = v2
    v15 := GenS2{F0: 32768}
    v12 += (v12 ^ (v12 / ((v12) | 1)))
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
