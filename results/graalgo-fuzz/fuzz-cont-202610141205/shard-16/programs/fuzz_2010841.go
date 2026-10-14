package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0() int16 {
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = 127
        v7 = v7
        _ = fi8
    }
    switch 127 {
    case -524:
        v7 += 111
    case 32768:
        v9 := []int{}
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: 0, F1: 681}
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int32 = 36
            v15 := []int{(fi13 | 9223372036854775807), int(v7)}
            fmt.Println(v7, v14, fi13)
            var v16 uint8 = v7
            if (-2.25 < (0.1 * 1.5)) {
                v17 := GenS1{F0: int64(v14), F1: 39}
                var v18 float32 = 1e10
                v15 = append(v15, fi11)
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    var v19 uint16 = 42
    _ = v19
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 uint) uint16 {
    p0 = 757
    p0 -= p0
    var v20 uint64 = 0
    _ = p0
    _ = p1
    _ = v20
    return 590
}

func main() {
    var v1 float64 = ((0.1 * 3.14159) / 1.0)
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    var v6 uint64 = 2
    v21 := []int{91, -265, -1}
    v5[606] = (!true)
    var v22 int8 = -1
    fmt.Println(v4)
    v23 := GenS1{F0: 9223372036854775807, F1: 65535}
    v1 = v1
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v4 = (v4 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
