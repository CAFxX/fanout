package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8, p1 int64, p2 uint32) int {
    p2 += (p2 - 2147483647)
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            if (false && (0.0 <= -2.25)) {
                var v8 string = "a"
                _ = v8
            } else {
                switch 131 {
                case 32767:
                    var v9 uint8 = uint8(p2)
                    _ = v9
                default:
                    var v10 int32 = 0
                    _ = v10
                }
                p0 = p0
                if (false || (2 != 0)) {
                    p0 = (-(p0))
                } else {
                    fmt.Println(p2, p2, p1)
                    p1 -= (p1 | (p1 + int64(p0)))
                    p1 -= (p1 & 42)
                }
            }
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{42, (fi11 * fi11), 743, (-(fi11))}
        v13 := GenS1{F0: "a", F1: p1, F2: -9223372036854775808}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 106
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 int16 = 776
    v5 := make(map[int]uint32)
    var v6 bool = false
    v15 := v2[131]
    v6 = (1 <= v15)
    v6 = v6
    var v16 float64 = 1.5
    v17 := v2[203]
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(fi19, fi18, fi18)
            v20 := GenS2{F0: 933, F1: (v4 / ((v4) | 1))}
            _ = fi19
            _ = v20
        }
        var v21 uint8 = 42
        fmt.Println(fi18, v15, v4)
        _ = fi18
        _ = v21
    }
    v16 -= (1e10 * (-2.25 / 1.0))
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{}
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
