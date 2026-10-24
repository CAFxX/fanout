package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }
func genFn0(p0 int16, p1 uint64, p2 bool) uint {
    v7 := []int{1, -885, 9223372036854775807}
    var v8 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 uint = 255
    v4 := make(map[uint32]string)
    var v5 string = "\n\t"
    var v6 int32 = 42
    var v9 int32 = v6
    fmt.Println(v9)
    v3 -= (v3 | (v3 & (^(v3))))
    switch 127 {
    case 127:
        if v2 {
            for fi10 := 0; fi10 < 12; fi10++ {
                v6 -= (v9 >> ((v3) & 63))
                v11 := []int{fi10, fi10, (fi10 << ((451) & 63)), (fi10 * -568)}
                if (!false) {
                    var v12 bool = (v2 || false)
                    var v13 uint32 = 1
                    _ = v12
                    _ = v13
                } else {
                    v5 = v5
                }
                _ = fi10
                _ = v11
            }
            var v14 bool = ((v2 && v2) && (genFn0(-71, 518, (v2 || false)) < 0))
            _ = v14
        }
    case -9223372036854775808:
        fmt.Println(v9)
    }
    var v15 uint8 = 241
    v9 -= (v6 << (((v3 & (v3 % ((v3) | 1)))) & 63))
    for fi16 := 0; fi16 < 2; fi16++ {
        v9 = 0
        v17 := []int{(fi16 | -564), fi16, int(v15), (fi16 - fi16)}
        var v18 uint16 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 406
    v20 := GenS1{F0: 4294967295, F1: (^(v3)), F2: 42}
    if ((float64(((-2.25 / float32(1.5)) * 0.1)) * 3.14159) != -2.25) {
        fmt.Println(v9)
    } else {
        var v21 bool = (v9 >= v9)
        _ = v21
    }
    v22 := []int{440, 2147483647, -2147483648, 4294967295}
    if ((v19 | 42) < v19) {
        if v2 {
            v23 := GenS1{F0: 506, F1: uint(v19), F2: 18446744073709551615}
            _ = v23
        }
    } else {
        if ((v3 ^ v3) == 338) {
            var v24 int8 = 0
            _ = v24
        }
        v6 = (v6 / ((v6) | 1))
    }
    v25 := GenS2{F0: 42, F1: 70, F2: (v6 / (((v9 / ((v6) | 1))) | 1))}
    _ = len(v4)
    for fi26 := 0; fi26 < 14; fi26++ {
        _ = len(v22)
        _ = fi26
    }
    v27 := v4[v19]
    var v28 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
