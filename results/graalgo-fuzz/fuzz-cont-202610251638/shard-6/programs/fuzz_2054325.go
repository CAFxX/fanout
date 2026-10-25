package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint8, p1 int, p2 int8) int32 {
    switch -2147483648 {
    case -9223372036854775808:
        v6 := GenS1{F0: 2}
        _ = v6
    case 665:
        v7 := []int{(p1 / (((p1 >> ((18446744073709551615) & 63))) | 1)), (-(p1))}
        _ = v7
    default:
        v8 := GenS2{F0: 100, F1: 1}
        _ = v8
    }
    fmt.Println(p1, p0, p2)
    var v9 string = "\n\t"
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 942
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (0.0 / 1.5)
    var v3 int8 = 127
    var v4 int = 2147483647
    v5 := []int{(-(v4)), v4, 32767, -526}
    v2 = ((v2 / ((1.5 - 1.5) / (v2 * v2))) * (float32(1.5) + v2))
    v4 += (v4 * v4)
    var v10 string = "\n\t"
    v3 = 127
    v10 = v10
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 42}
        var v13 int8 = (v3 - v3)
        delete(v1, 1)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 string = v10
    if ((true || (v3 == v3)) || true) {
        v5 = append(v5, (^(v4)))
        var v15 uint8 = 0
        v2 = -2.25
        _ = v15
    } else {
        v14 = "foo bar"
        var v16 float64 = float64(float32((0.0 / 1e10)))
        fmt.Println(v16, v16)
        _ = v16
    }
    fmt.Println(v14, v3, v4)
    var v17 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
