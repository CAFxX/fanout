package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int8 = 1
    var v2 bool = (2147483648 == 2147483648)
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    fmt.Println(v2, v4, v2)
    if ("\n\t" < "\n\t") {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint = 317
            _ = fi5
            _ = v6
        }
        v1 += (v1 ^ v1)
        v4 = v4
    } else {
        v7 := GenS1{F0: 42}
        _ = v7
    }
    switch 664 {
    case -1:
        for fi8 := 0; fi8 < 8; fi8++ {
            v4 = v4
            _ = fi8
        }
    case 128:
        if v2 {
            v9 := GenS2{F0: v4, F1: -614}
            fmt.Println(v2, v1, v1)
            var v10 float64 = ((0.1 * 0.1) / 1e10)
            _ = v10
            _ = v9
        }
    case 804:
        v2 = v2
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            if v2 {
                v12 := []int{(fi11 | (fi11 * 2147483647)), (fi11 % ((42) | 1)), fi11, int(v4)}
                _ = v12
            }
            var v13 float64 = float64(0.0)
            _ = fi11
            _ = v13
        }
    }
    v2 = (65535 > 526)
    var v14 int64 = 42
    if false {
        if (!(v2 && v2)) {
            switch 128 {
            case -175:
                v15 := []int{}
                _ = v15
            case -9223372036854775808:
                v1 += -128
            case -1:
                v16 := []int{}
                _ = v16
            default:
                v17 := GenS3{F0: (!true)}
                _ = v17
            }
        } else {
            var v18 uint8 = 128
            v18 += 24
            var v19 uint = 42
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: 2147483647}
        v21 := v3[42]
        _ = v20
        _ = v21
    }
    v2 = v2
    if false {
        v4 -= v4
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint8 = 1
            v14 = (^(v14))
            _ = fi22
            _ = v23
        }
        if ((false || (1 > 1000)) && (v2 && v2)) {
            if true {
                fmt.Println(v4, v14, v14)
            }
            if v2 {
                var v24 bool = true
                _ = v24
            }
            var v25 int8 = 0
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
