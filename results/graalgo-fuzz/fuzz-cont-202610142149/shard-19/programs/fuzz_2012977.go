package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 uint64) uint64 {
    fmt.Println(p0, p1)
    fmt.Println(p2)
    v6 := GenS1{F0: 123, F1: -658, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 int = 469
    var v2 float64 = ((1.5 - 1.5) * (3.14159 + float64(float32((float64((0.1 / 3.14159)) * (1e10 * 1e10))))))
    v3 := make(map[uint32]string)
    var v4 uint32 = 0
    var v5 float64 = ((float64(1e10) + 1e10) / (v2 / (v2 - 0.1)))
    fmt.Println(v1, v1, v1)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    var v9 uint32 = 2147483647
    v1 -= -1
    v10 := GenS2{F0: 1}
    for fi11 := 0; fi11 < 16; fi11++ {
        v2 = -2.25
        v1 += 42
        _ = fi11
    }
    var v12 float64 = v2
    v13 := []int{655, -2147483648, v1, v1}
    fmt.Println(v4, v2)
    var v14 uint32 = (v9 << ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
