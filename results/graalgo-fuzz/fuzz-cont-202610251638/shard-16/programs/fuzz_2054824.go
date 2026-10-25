package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0() int {
    v5 := GenS1{F0: -1}
    var v6 int16 = 32767
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = 42
        fmt.Println(v6, v6, v8)
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 int8 = 127
    var v2 int32 = -2147483648
    var v3 uint32 = 2
    var v4 int = 2147483647
    if (0.0 < (0.0 * (-2.25 * ((3.14159 - 0.0) / 1.0)))) {
        v4 -= genFn0()
        if (32767 >= 42) {
            v4 = (v4 ^ 595)
            if false {
                switch (v4 ^ (v4 / (((v4 & (v4 * v4))) | 1))) {
                case -422:
                    for fi9 := 0; fi9 < 16; fi9++ {
                        fmt.Println(v4, v1)
                        v3 += v3
                        _ = fi9
                    }
                case 915:
                    v10 := GenS1{F0: 0}
                    _ = v10
                case -293:
                    if (!true) {
                        fmt.Println(v3, v4)
                        fmt.Println(v2, v3, v1)
                    }
                default:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := []int{-9223372036854775808}
                        _ = fi11
                        _ = v12
                    }
                }
            } else {
                fmt.Println(v2)
                v13 := GenS2{F0: 611, F1: v2}
                _ = v13
            }
        } else {
            v3 = (v3 >> ((18446744073709551615) & 63))
            v2 += (v2 * v2)
            var v14 float64 = float64(3.14159)
            _ = v14
        }
        var v15 uint64 = 734
        _ = v15
    } else {
        fmt.Println(v1, v3, v3)
        for fi16 := 0; fi16 < 5; fi16++ {
            v4 = fi16
            _ = fi16
        }
        fmt.Println(v1)
    }
    v17 := []int{}
    fmt.Println(v4, v4, v4)
    v18 := GenS1{F0: 979}
    v2 = v2
    v2 = (v2 << ((1) & 63))
    var v19 float32 = 0.1
    var v20 float64 = -2.25
    var v21 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
