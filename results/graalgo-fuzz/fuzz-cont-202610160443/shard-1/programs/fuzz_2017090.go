package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []string{"", "", "hello"}
    v2 := make(map[string]bool)
    v3 := []int{32767, 279, -9223372036854775808, -1}
    v4 := []int{2147483647, 32767, 1}
    var v5 int8 = 3
    v6 := make(map[string]uint32)
    v5 += v5
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v5, v5)
        v8 := GenS1{F0: false}
        v5 -= v5
        _ = fi7
        _ = v8
    }
    v5 = v5
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: true}
        v11 := GenS1{F0: (42 <= 42)}
        if len(v4) > 0 {
            v4[0] = (fi9 / (((fi9 << ((4294967295) & 63))) | 1))
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 += 0
    v12 := GenS1{F0: (1 != -32768)}
    v13 := GenS2{F0: 1, F1: -455}
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS2{F0: 2, F1: 1}
            _ = len(v4)
            var v17 float32 = -2.25
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            if false {
                var v19 uint64 = 4294967295
                _ = len(v2)
                _ = v19
            }
            _ = fi18
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
