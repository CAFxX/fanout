package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
func genFn0() uint16 {
    var v6 uint64 = 4294967295
    v6 += 271
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, fi8)
            _ = fi8
        }
        v6 = 1
        v6 -= 4294967295
        _ = fi7
    }
    v9 := GenS1{F0: -117, F1: -1, F2: -519}
    _ = v6
    _ = v9
    return 32767
}
func genFn1(p0 int64) int64 {
    v10 := GenS2{F0: -252}
    var v11 uint = 42
    v12 := []int{-9223372036854775808, -361, 488, 42}
    v11 = (v11 / (((v11 / ((224) | 1))) | 1))
    v13 := GenS2{F0: 947}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return (p0 % ((42) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2
    v3 := []int{-1, 9223372036854775807}
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    fmt.Println(v2)
    _ = len(v3)
    switch 42 {
    case 1:
        var v14 float32 = 3.14159
        _ = v14
    }
    switch "\n\t" {
    case "a":
        var v15 uint = 0
        _ = v15
    case "hello":
        v2 = 685
    case "\n\t":
        var v16 float64 = 1.5
        _ = v16
    default:
        var v17 string = ("\n\t" + "\n\t")
        _ = v17
    }
    fmt.Println(v2)
    v2 = v2
    for fi18 := 0; fi18 < 5; fi18++ {
        if (false && ((-490 > 0) || true)) {
            if (false && (!false)) {
                v2 = (^(v2))
                if true {
                    var v19 uint = 4294967295
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v19 = v19
                        _ = len(v3)
                        _ = fi20
                    }
                    _ = v19
                }
                if false {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = (^(fi18))
                        }
                        v2 = (v2 ^ (v2 + v2))
                    }
                }
            } else {
                v21 := GenS3{F0: -128}
                var v22 float32 = -2.25
                _ = v21
                _ = v22
            }
        } else {
            fmt.Println(fi18, fi18)
            v2 = v2
        }
        _ = fi18
    }
    var v23 int32 = 0
    v23 += (v23 << ((755) & 63))
    var v24 bool = (((3.14159 - 1.5) + ((0.0 - 0.1) - -2.25)) > 1.5)
    fmt.Println(v2, v23, v2)
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := v4[fi25]
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: v23}
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
