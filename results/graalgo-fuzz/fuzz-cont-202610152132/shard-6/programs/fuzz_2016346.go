package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 int16; F1 int; F2 int }
func genFn0(p0 int8, p1 uint, p2 uint64) int8 {
    v4 := []int{32767, -1, 1, 42}
    fmt.Println(p1, p2, p1)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint64 = (^(p2))
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -128
}

func main() {
    var v1 bool = (9223372036854775807 < -625)
    var v2 float32 = ((0.1 * 0.1) - (1.5 * float32(0.1)))
    var v3 uint = 309
    v2 -= v2
    if v1 {
        v7 := []int{187}
        var v8 float32 = v2
        _ = v7
        _ = v8
    } else {
        if v1 {
            var v9 bool = (false && (v1 && v1))
            _ = v9
        } else {
            v10 := GenS1{F0: 198, F1: 1, F2: 989}
            switch 101 {
            case -500:
                v11 := []int{-1, 90, 2147483647}
                _ = v11
            case 2147483647:
                v1 = (1 <= 42)
            }
            _ = v10
        }
    }
    fmt.Println(v3, v1, v1)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint8 = 154
        if ((609 <= 4294967295) || true) {
            if (1 <= 18446744073709551615) {
                v13 = v13
                v14 := []int{}
                var v15 float32 = v2
                _ = v14
                _ = v15
            } else {
                fmt.Println(v1, v2)
                v3 += 2
                v1 = ((0.0 / ((3.14159 + 1e10) / (((0.1 / (1e10 + float64((-2.25 / float32(0.1))))) - (0.1 * (1.5 + 3.14159))) - 0.1))) >= ((3.14159 / -2.25) * 3.14159))
            }
            var v16 uint = v3
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS1{F0: -536, F1: -1, F2: -845}
    v18 := []int{4294967295, 4294967295, -2147483648}
    _ = len(v18)
    v18 = append(v18, -2147483648)
    v1 = v1
    v3 -= v3
    for fi19 := 0; fi19 < 4; fi19++ {
        v3 += (v3 * v3)
        _ = fi19
    }
    v20 := GenS2{F0: 1000, F1: -957, F2: 42}
    var v21 bool = ((v3 ^ v3) == 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
