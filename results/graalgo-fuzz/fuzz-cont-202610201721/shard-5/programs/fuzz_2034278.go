package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }

func main() {
    v1 := []int{}
    var v2 int16 = 1000
    var v3 int64 = -1
    var v4 uint64 = 805
    v5 := []uint32{65535, 2147483648}
    var v6 int16 = (v2 * (v2 % (((v2 >> ((42) & 63))) | 1)))
    if (18446744073709551615 != v4) {
        if true {
            var v7 uint = 42
            v4 += (^(v4))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 uint32 = 2147483647
                var v10 int = 251
                var v11 bool = (false || ((true || true) && (v10 >= -464)))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            fmt.Println(v2, v6)
        }
        if true {
            if false {
                for fi12 := 0; fi12 < 12; fi12++ {
                    fmt.Println(v2, v4)
                    v4 = v4
                    _ = fi12
                }
                v5 = append(v5, 2)
            } else {
                if false {
                    v3 += int64(v6)
                    v4 = (v4 | v4)
                }
            }
            if true {
                v13 := GenS1{F0: -1, F1: 19}
                _ = v13
            }
            var v14 int16 = v2
            _ = v14
        } else {
            var v15 uint8 = 1
            _ = v15
        }
    }
    if len(v1) > 0 {
        v1[0] = -546
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 float32 = -2.25
        _ = fi16
        _ = v17
    }
    v5 = append(v5, 602)
    var v18 uint32 = 4294967295
    v4 = v4
    v19 := GenS1{F0: 0, F1: 135}
    v6 -= (^(v2))
    v5 = append(v5, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
