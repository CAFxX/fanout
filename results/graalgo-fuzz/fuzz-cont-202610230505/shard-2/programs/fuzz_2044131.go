package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 int16, p2 bool) int64 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{(-(fi6)), (fi6 * (fi6 << ((18446744073709551615) & 63)))}
        var v8 uint = 2
        if len(v7) > 0 {
            v7[0] = fi6
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: p0}
    p2 = (65535 != 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 ^ (-(p0)))
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 0
    var v3 uint8 = 255
    var v4 float32 = (-2.25 * ((v1 + v1) - (float32((((0.0 * (-2.25 - 1.5)) / -2.25) / (((float64(float32((1e10 / -2.25))) + (1.5 * -2.25)) + float64((1e10 / (-2.25 * v1)))) + 0.1))) * (v1 * v1))))
    v5 := make(map[uint32]uint32)
    var v10 int16 = 1000
    v11 := []int{127, -9223372036854775808, 128, 9223372036854775807}
    v3 = (v3 & v3)
    v12 := GenS1{F0: 42}
    v4 = v4
    fmt.Println(v1, v10)
    v10 -= -32768
    v13 := GenS2{F0: 832}
    v14 := v5[2]
    fmt.Println(v14, v10)
    v2 -= uint64(v14)
    if ("\n\t" == "hello") {
        var v15 int8 = 77
        v16 := GenS1{F0: 909}
        _ = v15
        _ = v16
    }
    var v17 float64 = 1.5
    fmt.Println(v14)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
