package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }

func main() {
    var v1 float32 = 0.1
    v2 := []int{9223372036854775807, 186}
    var v3 int = -9223372036854775808
    var v4 bool = false
    v5 := make(map[uint32]uint32)
    v3 += (v3 & -9223372036854775808)
    v2 = append(v2, (^(v3)))
    v6 := GenS1{F0: 0, F1: -1, F2: ((v3 + (v3 % ((v3) | 1))) != (v3 / ((v3) | 1)))}
    _ = len(v5)
    if (321 != 1) {
        if (42 != 0) {
            _ = len(v5)
        } else {
            v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: v4}
            fmt.Println(v1, v4)
            for fi8 := 0; fi8 < 14; fi8++ {
                v5[866] = 2147483647
                fmt.Println(fi8)
                v3 = 32767
                _ = fi8
            }
            _ = v7
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = ""
            v3 = (v3 % (((v3 + (-(v3)))) | 1))
            _ = fi9
            _ = v10
        }
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    v3 = (v3 + v3)
    var v11 int16 = 42
    delete(v5, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
