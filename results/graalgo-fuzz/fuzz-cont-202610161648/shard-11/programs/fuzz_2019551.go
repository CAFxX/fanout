package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int }
func genFn0() uint16 {
    var v4 uint64 = 591
    v4 += v4
    v4 = v4
    _ = v4
    return 546
}
func genFn1() int32 {
    v5 := GenS1{F0: 42, F1: -1}
    var v6 int64 = 42
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = (float32(0.0) + (1e10 + 0.1))
    v3 := make(map[uint32]int)
    if true {
        var v7 uint = 0
        v8 := GenS1{F0: 2147483647, F1: 187}
        _ = v7
        _ = v8
    } else {
        v1 -= -200
        var v9 uint = 2
        var v10 int16 = -1
        _ = v10
        _ = v9
    }
    if false {
        v11 := GenS2{F0: -1, F1: 255}
        v2 -= (v2 / v2)
        var v12 int64 = 1
        _ = v11
        _ = v12
    } else {
        v13 := GenS3{F0: 32767, F1: 131, F2: 4294967295}
        _ = v13
    }
    var v14 uint16 = 439
    if (100 >= 31) {
        fmt.Println(v1, v14)
    } else {
        switch "\n\t" {
        case "foo bar":
            v1 = (v1 - v1)
        case "":
            if true {
                v15 := GenS2{F0: -46, F1: (v14 ^ v14)}
                var v16 bool = false
                delete(v3, 2)
                _ = v15
                _ = v16
            } else {
                var v17 uint = 18446744073709551615
                _ = v17
            }
        case "a":
            fmt.Println(v1)
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v14, v2)
                fmt.Println(fi18, v14)
                _ = fi18
            }
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(v1, fi19, fi19)
            v14 -= (v14 & v14)
            var v20 int64 = (v1 << ((561) & 63))
            _ = fi19
            _ = v20
        }
    }
    if (v1 >= v1) {
        switch "a" {
        case "a":
            v1 += 1
        case "hello":
            v21 := []int{}
            _ = v21
        case "":
            fmt.Println(v2, v14)
        default:
            var v22 bool = false
            _ = v22
        }
        var v23 int64 = v1
        v2 += 1.5
        _ = v23
    } else {
        switch (("\n\t" + "a") + ("" + ("\n\t" + ""))) {
        case "a":
            var v24 int32 = -2147483648
            _ = v24
        default:
            if false {
                v25 := v3[708]
                _ = v25
            } else {
                switch (("\n\t" + "\n\t") + "") {
                case "":
                    v26 := []int{32767, 0, 127, 310}
                    _ = v26
                case "\n\t":
                    var v27 int64 = v1
                    _ = v27
                default:
                    if ((0 > -2147483648) || ((v14 * v14) >= v14)) {
                        v14 -= 65535
                    }
                }
            }
        }
    }
    if ((v14 < v14) || (!(true && false))) {
        v28 := []int{}
        v14 -= (v14 % (((^(v14))) | 1))
        _ = v28
    }
    var v29 int8 = -128
    v30 := GenS2{F0: 127, F1: v14}
    var v31 int = 517
    v32 := []int{}
    fmt.Println(v14, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
