package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
func genFn0() int16 {
    v4 := GenS1{F0: 47, F1: 0, F2: 1000}
    var v5 uint64 = 2
    _ = v4
    _ = v5
    return -799
}
func genFn1(p0 uint16, p1 uint16) uint64 {
    var v6 uint8 = 42
    v6 += v6
    var v7 float64 = 0.1
    if false {
        v8 := GenS2{F0: 1, F1: 573}
        v7 += v7
        var v9 int64 = 42
        _ = v8
        _ = v9
    }
    if false {
        v10 := GenS3{F0: 0, F1: 0, F2: 127}
        v11 := []int{0}
        v12 := GenS1{F0: 658, F1: 169, F2: -752}
        _ = v10
        _ = v11
        _ = v12
    } else {
        p0 -= (p0 / (((p1 / (((p0 / ((p0) | 1))) | 1))) | 1))
        v13 := []int{42, 32767}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 263
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []uint32{798, 65535, 667, 2147483648}
    var v14 int32 = 1
    fmt.Println(v14)
    fmt.Println(v14)
    var v15 int64 = -9223372036854775808
    v16 := GenS2{F0: 1, F1: 1}
    if (true || false) {
        v17 := v1["foo bar"]
        v18 := GenS3{F0: 9223372036854775807, F1: 42, F2: -99}
        _ = v17
        _ = v18
    } else {
        fmt.Println(v14)
        switch "foo bar" {
        case "":
            switch "a" {
            case "hello":
                v15 += v15
            case "a":
                var v19 bool = false
                _ = v19
            case "\n\t":
                v15 = v15
            default:
                if true {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 int16 = 42
                        var v22 string = "foo bar"
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            }
        case "hello":
            switch 128 {
            case 32767:
                for fi23 := 0; fi23 < 1; fi23++ {
                    _ = len(v3)
                    _ = fi23
                }
            default:
                fmt.Println(v15, v14)
            }
        default:
            var v24 uint = 974
            _ = v24
        }
        if ((v14 > v14) && true) {
            var v25 int8 = 1
            _ = v25
        }
    }
    var v26 uint8 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
}
