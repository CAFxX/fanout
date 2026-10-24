package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 bool }
func genFn0() int16 {
    v7 := GenS1{F0: 2147483647, F1: 18446744073709551615}
    v8 := []int{}
    v9 := GenS1{F0: 42, F1: 1}
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    v4 := []int{9223372036854775807, 4294967295, 920, 164}
    v5 := []uint32{(^(v1)), v1, 1, (v1 * (v1 ^ 2))}
    var v6 uint8 = v2
    v6 = (^(v6))
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: fi10, F1: ((0.1 - float64(0.0)) == 1e10)}
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v6)
            delete(v3, fi12)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            v5 = append(v5, (v1 | (v1 & (v1 / ((v1) | 1)))))
            if (-1 == -68) {
                v1 = 0
                v1 += v1
            } else {
                switch (-(fi14)) {
                case 468:
                    fmt.Println(fi13)
                case -722:
                    var v15 int = fi13
                    _ = v15
                }
                switch -2147483648 {
                case -292:
                    fmt.Println(v1, v6)
                case 4294967295:
                    var v16 int64 = -1
                    _ = v16
                case -2147483648:
                    fmt.Println(fi14, v2, fi13)
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    var v17 string = "a"
    v18 := GenS1{F0: 42, F1: 410}
    if false {
        v19 := v3[2147483647]
        v20 := GenS1{F0: v1, F1: 960}
        _ = v19
        _ = v20
    } else {
        v1 = v1
        if false {
            fmt.Println(v2)
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
        }
    }
    var v21 int32 = -1
    var v22 int32 = 2147483647
    v23 := GenS1{F0: (v1 % ((v1) | 1)), F1: 0}
    v24 := GenS2{F0: 32767, F1: false}
    v1 -= (v1 - 478)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
