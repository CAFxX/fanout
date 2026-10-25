package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0() int {
    if (true && true) {
        var v6 bool = false
        v7 := []int{9223372036854775807}
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 103, F1: -1}
        var v9 int16 = 0
        v9 += (v9 ^ (-(v9)))
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 1}
    var v11 int = 1
    v12 := []int{}
    var v13 float32 = 3.14159
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return v11
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int64 = 248
    var v4 int32 = -789
    v5 := make(map[int]string)
    switch 583 {
    case -425:
        switch genFn0() {
        case 4294967295:
            var v14 float32 = 0.0
            _ = v14
        case 32767:
            v4 += 1
        }
    case 128:
        for fi15 := 0; fi15 < 19; fi15++ {
            if (280 <= 42) {
                switch 9223372036854775807 {
                case 42:
                    var v16 int = fi15
                    _ = v16
                }
            } else {
                switch fi15 {
                case -589:
                    var v17 int16 = 656
                    _ = v17
                default:
                    v4 = (-(v4))
                }
                v3 = v3
            }
            _ = fi15
        }
    default:
        v4 += v4
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v3 = v3
        var v19 float32 = 0.1
        for fi20 := 0; fi20 < 11; fi20++ {
            v4 += (v4 ^ (v4 >> ((255) & 63)))
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    if true {
        v2[9223372036854775807] = (!(255 >= 32767))
        v21 := v1[927]
        fmt.Println(v21, v4, v21)
        _ = v21
    }
    fmt.Println(v3, v3)
    v22 := []int{-13, 4294967295}
    v3 -= (v3 | v3)
    var v23 int8 = 94
    if true {
        switch "" {
        case "hello":
            v23 = (^(v23))
        }
    } else {
        if true {
            if true {
                var v24 int16 = int16(v23)
                _ = v24
            }
        }
        v4 -= int32(v23)
        switch 817 {
        case 32767:
            var v25 bool = false
            _ = v25
        case 32768:
            var v26 uint = 1
            _ = v26
        }
    }
    var v27 bool = (true || (207 == 167))
    var v28 float64 = 1e10
    v4 = (v4 / ((-893) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
