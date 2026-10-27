package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 uint64, p1 bool) int {
    v4 := GenS1{F0: (p0 | p0), F1: 4294967295, F2: -608}
    v5 := []int{-671, 127, 36}
    var v6 float32 = (-2.25 * 1e10)
    var v7 int8 = -91
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []string{"hello"}
    var v2 float32 = (3.14159 / (float32(0.0) - (3.14159 - -2.25)))
    var v3 uint32 = 2147483647
    fmt.Println(v2)
    v2 -= (-2.25 + (float32((0.1 - 3.14159)) / v2))
    fmt.Println(v3, v3)
    fmt.Println(v2, v3)
    v3 = v3
    for fi9 := 0; fi9 < 1; fi9++ {
        if (2 == 195) {
            if (!(127 > 42)) {
                v3 += (^(v3))
                if false {
                    var v10 int16 = 314
                    var v11 float64 = -2.25
                    v2 += v2
                    _ = v10
                    _ = v11
                }
            } else {
                fmt.Println(v2, v2, v3)
                v2 += (3.14159 - (v2 / 1.0))
            }
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v2 = (float32(((-2.25 * 3.14159) + 0.1)) / v2)
            _ = fi12
        }
        _ = fi9
    }
    v2 = (-2.25 * v2)
    v13 := GenS2{F0: genFn0(42, (233 == 783)), F1: 1, F2: v3}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
}
