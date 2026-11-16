package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 int8) string {
    var v7 int32 = -126
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("a" + "")
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 uint16 = 858
    var v5 bool = false
    v6 := make(map[string]int)
    for fi8 := 0; fi8 < 20; fi8++ {
        switch fi8 {
        case -366:
            v1 = (v1 * v1)
        }
        if (false && v5) {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: 42}
                _ = fi9
                _ = v10
            }
        } else {
            v11 := GenS1{F0: 0}
            _ = v11
        }
        _ = fi8
    }
    fmt.Println(v1)
    v4 = (v4 << ((1) & 63))
    v4 = v4
    var v12 float32 = 0.1
    v13 := []int{1, -2147483648, 2147483647, -2147483648}
    v1 = v1
    var v14 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
