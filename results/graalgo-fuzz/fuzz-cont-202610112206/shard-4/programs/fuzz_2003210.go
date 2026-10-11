package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int) uint16 {
    v4 := GenS1{F0: 227, F1: 0}
    var v5 bool = false
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 1; fi6++ {
        if (65535 < 2) {
            v7 := GenS1{F0: 2, F1: 18446744073709551615}
            _ = v7
        } else {
            var v8 float64 = 0.0
            var v9 int32 = -943
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int64 = 1
                p0 += (fi10 ^ 308)
                v12 := []int{fi10, fi6, (fi10 << ((2) & 63)), p0}
                var v13 uint64 = 0
                v13 -= v13
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int = -284
    var v2 float64 = 1e10
    var v3 uint = 42
    v14 := []int{(^(v1))}
    fmt.Println(v1, v3, v2)
    v15 := GenS2{F0: 264, F1: 738}
    switch (v1 | 127) {
    case 32767:
        _ = len(v14)
    }
    v3 -= (v3 * 79)
    switch "hello" {
    case "hello":
        var v16 int = v1
        _ = v16
    case "\n\t":
        v17 := GenS2{F0: 42, F1: 1}
        _ = v17
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v18 := GenS2{F0: uint64(v3), F1: 524}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
