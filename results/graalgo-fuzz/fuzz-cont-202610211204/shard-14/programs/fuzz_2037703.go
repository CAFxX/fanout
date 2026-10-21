package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint8) int {
    v6 := GenS1{F0: 2}
    var v7 float32 = -2.25
    _ = p0
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 uint = 2
    v5 := []uint32{4294967295, 96}
    v5 = append(v5, 0)
    fmt.Println(v4)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint = v4
        var v10 uint64 = 2
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 -= 42
    if (2 != 127) {
        var v11 int32 = 42
        if false {
            v4 -= v4
        } else {
            var v12 float64 = float64(0.0)
            _ = v12
        }
        _ = v11
    }
    var v13 uint16 = 732
    v14 := GenS1{F0: (v13 / ((v13) | 1))}
    _ = len(v5)
    v15 := GenS1{F0: v13}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
