package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
func genFn0(p0 uint32) int64 {
    v6 := GenS1{F0: 1, F1: 2147483647}
    var v7 int8 = -55
    switch "a" {
    case "a":
        v8 := GenS2{F0: ((-1 < 811) || (-1 >= -434)), F1: 90, F2: 71}
        _ = v8
    }
    v9 := GenS2{F0: (true || (!false)), F1: 127, F2: 65535}
    fmt.Println(v7, v7)
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 210
    var v2 uint8 = 174
    var v3 float32 = ((-2.25 / -2.25) / 1e10)
    var v4 int32 = 0
    v5 := make(map[int]int)
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v3, fi10, v4)
        var v11 uint64 = 18446744073709551615
        fmt.Println(v2, v11, v11)
        _ = fi10
        _ = v11
    }
    v12 := v5[127]
    for fi13 := 0; fi13 < 3; fi13++ {
        if (("a" != ("a" + "\n\t")) || false) {
            for fi14 := 0; fi14 < 17; fi14++ {
                if false {
                    v3 = (0.1 * ((-2.25 / (0.0 / (v3 / v3))) * ((1.5 - (v3 + v3)) * v3)))
                    v15 := []int{-732, fi14}
                    fmt.Println(v3, fi14)
                    _ = v15
                } else {
                    if (true || (((v4 | (-(v4))) <= v4) || (9223372036854775807 <= v12))) {
                        v4 -= (-(v4))
                        var v16 uint8 = v2
                        var v17 int8 = -97
                        var v18 int8 = -1
                        fmt.Println(fi13, v2, v18)
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = len(v5)
                }
                var v19 uint = 2
                _ = fi14
                _ = v19
            }
            v20 := GenS1{F0: 1, F1: (v12 | (-(v12)))}
            _ = v20
        }
        _ = fi13
    }
    v12 = v12
    var v21 int = v12
    fmt.Println(v1, v4, v1)
    fmt.Println(v21, v1)
    v22 := GenS1{F0: 929, F1: 700}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
