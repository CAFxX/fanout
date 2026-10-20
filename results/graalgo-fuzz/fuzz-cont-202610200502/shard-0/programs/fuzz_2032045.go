package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 int16, p1 int32) uint32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{-1, -2147483648}
        var v8 uint32 = 4294967295
        v8 += (v8 << ((0) & 63))
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{128}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 float64 = ((1.5 * ((0.0 * -2.25) + 1e10)) + 1e10)
    var v2 string = "a"
    v3 := []uint32{2147483647, 1}
    v4 := make(map[int]string)
    var v5 int16 = 309
    delete(v4, 42)
    v5 = (v5 >> ((351) & 63))
    var v10 uint8 = 2
    var v11 bool = true
    fmt.Println(v5, v10, v5)
    v10 = (v10 ^ 2)
    fmt.Println(v1, v10, v10)
    v3 = append(v3, 2147483647)
    if (2 > (v10 << ((366) & 63))) {
        v1 = float64((0.1 / 0.1))
        _ = len(v4)
        for fi12 := 0; fi12 < 13; fi12++ {
            if v11 {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v2 = v2
                    v14 := GenS1{F0: 2147483647}
                    _ = fi13
                    _ = v14
                }
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            v1 = 1e10
            _ = fi12
        }
    } else {
        var v15 float64 = (v1 / v1)
        var v16 int32 = -316
        v10 += 184
        _ = v15
        _ = v16
    }
    var v17 float32 = 1.5
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println(v11)
    v18 := GenS2{F0: false, F1: 42, F2: (v5 & -32768)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
