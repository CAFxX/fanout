package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 int16, p1 uint, p2 string) uint32 {
    v5 := GenS1{F0: 159, F1: 65535, F2: 42}
    p0 = p0
    if (598 > -9223372036854775808) {
        p0 = 0
        p2 = p2
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        if true {
            v7 := []int{(fi6 - (fi6 % ((fi6) | 1)))}
            _ = v7
        }
        _ = fi6
    }
    p1 += 961
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 377
}
func genFn1(p0 int8) int16 {
    v8 := []int{-953, 9223372036854775807}
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint16 = 2
            v11 := GenS1{F0: (^(v10)), F1: 253, F2: 2}
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(p0)
    } else {
        p0 += (p0 ^ -110)
        if false {
            p0 += 100
            if false {
                var v12 int8 = (p0 + 32)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: 811, F1: 879, F2: 2}
                    var v15 uint = 255
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                p0 = -10
                p0 = (p0 & -33)
                p0 -= p0
            }
        } else {
            fmt.Println(p0)
            _ = len(v8)
            if len(v8) > 0 {
                v8[0] = 2147483647
            }
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint16 = 702
            var v18 uint32 = 397
            fmt.Println(v17, v17, v18)
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 float64 = 1e10
    _ = p0
    _ = v19
    _ = v8
    return -354
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v1 = (v1 << ((18446744073709551615) & 63))
    v1 = v1
    var v20 uint8 = 89
    v21 := GenS2{F0: 0, F1: 42}
    v22 := GenS2{F0: -388, F1: 1}
    var v23 bool = (((v1 | 980) < v1) || false)
    switch "" {
    case "hello":
        var v24 uint8 = (^(v20))
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint8 = (v20 / ((2) | 1))
            v27 := []int{(fi25 >> ((0) & 63))}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 uint8 = 51
    fmt.Println(v28)
    var v29 uint32 = genFn0(500, 815, "")
    fmt.Println(v28, v23)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
