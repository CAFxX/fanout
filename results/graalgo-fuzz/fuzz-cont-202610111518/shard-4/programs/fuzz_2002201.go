package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 int }

func main() {
    var v1 uint = 42
    var v2 float32 = 1e10
    v3 := []int{}
    fmt.Println(v2)
    var v4 uint8 = 128
    v1 -= (v1 << (((v1 & (v1 % (((v1 << (((v1 + v1)) & 63))) | 1)))) & 63))
    v3 = append(v3, -1)
    var v5 int = 1
    _ = len(v3)
    if (-1 > -9223372036854775808) {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 920, F1: 4294967295}
            if true {
                var v8 int8 = -1
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 uint8 = (v4 % ((v4) | 1))
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    v4 -= (v4 << (((^(v1))) & 63))
                    fmt.Println(v8, fi11, v1)
                    _ = fi11
                }
                _ = v8
            }
            var v12 uint64 = 379
            _ = fi6
            _ = v12
            _ = v7
        }
        switch 900 {
        case -332:
            v13 := GenS1{F0: 0, F1: v5}
            _ = v13
        case 0:
            v2 -= 3.14159
        }
    }
    if true {
        v14 := GenS2{F0: v5, F1: v5}
        switch v5 {
        case 32767:
            if true {
                switch v5 {
                case -2147483648:
                    v1 -= 65535
                }
                var v15 uint8 = 199
                _ = v15
            }
        case 32768:
            fmt.Println(v4, v2)
        default:
            fmt.Println(v5, v2, v2)
        }
        _ = v14
    }
    var v16 uint = (^(v1))
    fmt.Println(v16, v5)
    var v17 uint32 = 65535
    if true {
        if ((uint64(v17) != 0) || true) {
            for fi18 := 0; fi18 < 2; fi18++ {
                v2 = v2
                _ = fi18
            }
            if false {
                v5 = (v5 & (v5 % (((v5 % ((2147483647) | 1))) | 1)))
                var v19 bool = false
                fmt.Println(v17)
                _ = v19
            }
        } else {
            v20 := GenS1{F0: -1, F1: v5}
            v21 := GenS1{F0: 904, F1: (^(v5))}
            _ = v20
            _ = v21
        }
    }
    var v22 float32 = (v2 * v2)
    for fi23 := 0; fi23 < 1; fi23++ {
        v16 -= 0
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
