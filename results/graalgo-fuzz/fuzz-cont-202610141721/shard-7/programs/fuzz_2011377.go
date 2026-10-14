package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 bool) int8 {
    var v6 int64 = 488
    fmt.Println(p0, p0)
    v6 += 1
    _ = p0
    _ = v6
    return -128
}

func main() {
    v1 := []int{4294967295, -890}
    var v2 uint8 = 16
    var v3 uint32 = 634
    var v4 float64 = ((float64(((0.1 / float32(0.1)) + 0.0)) - 1.5) / (0.1 + 0.1))
    var v5 uint16 = 771
    v7 := GenS1{F0: 0}
    if true {
        _ = len(v1)
        switch "\n\t" {
        case "a":
            switch 878 {
            case 32767:
                var v8 int8 = 0
                _ = v8
            case 127:
                var v9 float64 = (v4 / 1.5)
                _ = v9
            }
        case "foo bar":
            if true {
                if (32767 < 958) {
                    v10 := GenS1{F0: 723}
                    var v11 uint64 = 129
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    if len(v1) > 0 {
                        v1[0] = (fi12 * fi12)
                    }
                    v5 = v5
                    var v13 int = -109
                    _ = fi12
                    _ = v13
                }
            } else {
                v3 = 1
            }
        case "hello":
            fmt.Println(v5, v3, v2)
        default:
            v1 = append(v1, -479)
        }
    }
    var v14 uint32 = 0
    v14 = 0
    v15 := GenS2{F0: int64(v2), F1: uint(v2)}
    v1 = append(v1, 4294967295)
    var v16 int16 = 75
    fmt.Println(v16)
    var v17 uint64 = 1
    _ = len(v1)
    v2 = v2
    v16 -= 32767
    var v18 bool = true
    v4 -= v4
    fmt.Println(v16, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
