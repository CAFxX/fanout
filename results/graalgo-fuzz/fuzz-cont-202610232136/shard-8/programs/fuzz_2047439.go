package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = -863
    }
    var v7 uint32 = 148
    v8 := GenS1{F0: 137}
    v6 = append(v6, 32767)
    v9 := v4[2147483647]
    var v10 uint16 = 1
    var v11 uint64 = 301
    if (0.0 >= float32(1e10)) {
        if (!((v10 << ((42) & 63)) > (v10 / ((v10) | 1)))) {
            v7 -= 0
        } else {
            var v12 int32 = 42
            _ = v12
        }
        switch -161 {
        case 2147483647:
            var v13 int = -7
            _ = v13
        case 1:
            var v14 uint = uint(v10)
            _ = v14
        case 127:
            var v15 uint = 714
            _ = v15
        default:
            if ((v11 | (v11 ^ 18446744073709551615)) > (v11 | (v11 - v11))) {
                v7 -= 0
                if true {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 int64 = -845
                        _ = fi16
                        _ = v17
                    }
                    var v18 uint64 = v11
                    _ = v18
                }
                _ = len(v1)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
