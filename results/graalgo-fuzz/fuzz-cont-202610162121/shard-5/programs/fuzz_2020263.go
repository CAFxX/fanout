package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 1
    var v3 uint64 = v2
    var v4 bool = false
    fmt.Println(v4, v2)
    fmt.Println(v4, v3)
    var v5 uint16 = 0
    if (0.0 < (1e10 + 1e10)) {
        delete(v1, 114)
        var v6 int16 = 0
        _ = v6
    }
    var v7 float64 = (((1e10 + (1.5 / float64(float32(3.14159)))) + (-2.25 * 0.1)) + (float64(0.0) * 0.1))
    if (2147483647 < 9223372036854775807) {
        fmt.Println(v5, v2)
        var v8 uint = 2
        switch -9223372036854775808 {
        case -282:
            fmt.Println(v7)
        case 247:
            v9 := GenS1{F0: 42, F1: 42, F2: 127}
            _ = v9
        }
        _ = v8
    }
    v4 = (42 >= 1)
    fmt.Println(v4)
    for fi10 := 0; fi10 < 4; fi10++ {
        if (!(!(-1 <= 0))) {
            v3 = 464
            var v11 float32 = 1.5
            _ = v11
        }
        if ((v2 + 18446744073709551615) <= v3) {
            v12 := GenS1{F0: 1, F1: 0, F2: 32767}
            if v4 {
                v2 += (v2 ^ (v3 | (v2 * 4294967295)))
                switch 32767 {
                case 32767:
                    _ = len(v1)
                case 2147483647:
                    v5 = v5
                }
            } else {
                var v13 float64 = 0.0
                v7 += ((1e10 / float64(0.1)) * 0.1)
                _ = len(v1)
                _ = v13
            }
            _ = v12
        } else {
            v2 = (v3 | v2)
            v14 := GenS2{F0: 2147483647, F1: (v5 % ((v5) | 1)), F2: 34}
            _ = v14
        }
        _ = fi10
    }
    v2 += 1
    v7 += v7
    v3 = (v3 % ((18446744073709551615) | 1))
    v15 := []int{}
    v4 = v4
    v15 = append(v15, 42)
    if v4 {
        _ = len(v1)
        v2 = 1
        var v16 int32 = 1
        _ = v16
    }
    var v17 bool = ((v5 % (((v5 / ((862) | 1))) | 1)) == v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
