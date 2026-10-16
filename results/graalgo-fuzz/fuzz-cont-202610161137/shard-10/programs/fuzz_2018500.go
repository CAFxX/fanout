package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 uint, p2 int16) int64 {
    fmt.Println(p1)
    fmt.Println(p1, p1)
    if (9223372036854775807 == 0) {
        v5 := GenS1{F0: -1, F1: 0}
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{}
            var v8 uint64 = 18446744073709551615
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        v9 := GenS2{F0: p0}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 296
}

func main() {
    var v1 uint64 = 579
    v2 := make(map[int]int)
    var v3 uint32 = 813
    var v4 float32 = 1.5
    for fi10 := 0; fi10 < 12; fi10++ {
        v2[(fi10 ^ 982)] = fi10
        var v11 int16 = 32767
        fmt.Println(fi10, v11, v1)
        _ = fi10
        _ = v11
    }
    v4 += (v4 + (v4 / ((v4 * v4) / (3.14159 - v4))))
    v12 := GenS1{F0: 1, F1: int64(v1)}
    if true {
        var v13 uint16 = 361
        fmt.Println(v13, v3)
        _ = v13
    }
    var v14 int64 = -9223372036854775808
    var v15 int16 = 1000
    if false {
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := []int{(fi16 + 4294967295), (-(fi16)), int(v3)}
            var v18 uint8 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            if false {
                v4 -= v4
                if false {
                    v15 = v15
                    v1 = 2
                } else {
                    switch (-(fi19)) {
                    case 661:
                        var v20 uint32 = v3
                        _ = v20
                    default:
                        var v21 int32 = -2147483648
                        _ = v21
                    }
                    v4 = v4
                }
                var v22 int32 = 376
                _ = v22
            }
            _ = fi19
        }
    }
    var v23 float32 = float32((((float64(v4) / 0.1) + float64(v4)) * 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
