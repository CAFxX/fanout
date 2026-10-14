package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool) uint16 {
    v6 := []int{-1, 1}
    if true {
        var v7 float32 = -2.25
        fmt.Println(p0, v7)
        _ = v7
    }
    v8 := GenS1{F0: (true && p0), F1: 32767}
    p0 = p0
    _ = p0
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := []int{}
    var v2 int32 = 0
    v3 := []string{}
    var v4 int8 = -1
    var v5 float32 = (-2.25 + (0.0 + -2.25))
    if (2147483647 <= 348) {
        if (true && false) {
            v5 = v5
        } else {
            var v9 int64 = -511
            _ = len(v3)
            var v10 int = 1
            _ = v10
            _ = v9
        }
        var v11 uint8 = 128
        var v12 uint = 18446744073709551615
        _ = v11
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (0 >= 0) {
        var v13 uint64 = 42
        fmt.Println(v4, v5)
        if (!false) {
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    v2 += 0
                    fmt.Println(v13)
                    fmt.Println(v4, fi14, v13)
                    _ = fi15
                }
                if false {
                    v16 := GenS1{F0: true, F1: 376}
                    v4 = (v4 >> ((1) & 63))
                    _ = len(v3)
                    _ = v16
                }
                _ = fi14
            }
            v4 -= v4
        }
        _ = v13
    }
    v5 += v5
    v1 = append(v1, 128)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 -= v4
    fmt.Println(v5, v2, v5)
    var v17 int32 = v2
    if (1 <= 0) {
        var v18 uint16 = 398
        if false {
            var v19 uint32 = 65535
            _ = v19
        }
        _ = v18
    } else {
        v20 := GenS1{F0: (int64(v17) != -9223372036854775808), F1: 133}
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
