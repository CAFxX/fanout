package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() bool {
    v5 := GenS1{F0: 238}
    v6 := []int{-565, 127, 4294967295, 128}
    _ = v5
    _ = v6
    return (!true)
}
func genFn1(p0 uint8, p1 int8) uint32 {
    v7 := []int{32767, 32767, -773, 0}
    p1 = (p1 - (^(p1)))
    var v8 uint = 0
    fmt.Println(p0, v8)
    v7 = append(v7, 2147483647)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 int = 558
    var v2 float64 = -2.25
    v3 := make(map[string]bool)
    var v4 float64 = float64(((1e10 - 3.14159) / 1.5))
    v9 := []int{(v1 << ((1) & 63))}
    var v10 uint = 18446744073709551615
    fmt.Println(v1, v1, v4)
    _ = len(v3)
    var v11 uint64 = 1
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint = v10
        _ = fi12
        _ = v13
    }
    v14 := v3["hello"]
    for fi15 := 0; fi15 < 4; fi15++ {
        switch (v1 | v1) {
        case -716:
            var v16 int = int(v11)
            _ = v16
        case 127:
            var v17 uint32 = 4294967295
            _ = v17
        default:
            v11 += 1
        }
        var v18 uint64 = v11
        _ = fi15
        _ = v18
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
