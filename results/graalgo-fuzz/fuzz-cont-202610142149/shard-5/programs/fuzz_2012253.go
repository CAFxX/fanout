package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 int64 = 338
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v4 int8 = -128
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{-952, fi5, (fi5 << ((255) & 63))}
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    v8 := v2[65535]
    switch 0 {
    case 613:
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint32 = 42
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v1, fi11, v8)
                v3 = true
                var v12 bool = true
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v3 = ((fi13 >= (fi13 >> ((2) & 63))) && v3)
        switch "hello" {
        case "hello":
            if v3 {
                v1 += (v1 % ((v1) | 1))
                for fi14 := 0; fi14 < 15; fi14++ {
                    switch "hello" {
                    case "foo bar":
                        var v15 uint32 = 2147483647
                        _ = v15
                    default:
                        _ = len(v7)
                    }
                    delete(v2, 2)
                    var v16 int = fi13
                    v17 := GenS1{F0: (^(v1)), F1: 42, F2: 0}
                    switch "" {
                    case "hello":
                        v1 -= v1
                    default:
                        v1 += v1
                    }
                    _ = fi14
                    _ = v16
                    _ = v17
                }
                v4 = v4
            }
        }
        v4 -= v4
        _ = fi13
    }
    v8 = (v1 <= 9223372036854775807)
    switch -9223372036854775808 {
    case 1:
        switch 280 {
        case -9223372036854775808:
            if true {
                v3 = ((!v8) && ((float32(0.0) * float32(3.14159)) < 1e10))
                v3 = false
            } else {
                fmt.Println(v3, v4, v1)
            }
        case 4294967295:
            var v18 int8 = (v4 | (v4 / ((v4) | 1)))
            _ = v18
        case -740:
            v19 := GenS1{F0: (v1 - v1), F1: 42, F2: 42}
            _ = v19
        }
    case 2147483647:
        if false {
            v3 = ((v1 - v1) < (v1 | (v1 | (v1 / (((v1 % ((v1) | 1))) | 1)))))
            v7 = append(v7, 42)
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS1{F0: 1, F1: 2147483648, F2: 255}
                switch "a" {
                case "foo bar":
                    v22 := GenS2{F0: 364, F1: 201}
                    _ = v22
                case "a":
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if true {
                            fmt.Println(fi23, fi23)
                        }
                        var v24 int8 = (^(v4))
                        _ = fi23
                        _ = v24
                    }
                }
                if len(v7) > 0 {
                    v7[0] = fi20
                }
                _ = fi20
                _ = v21
            }
        }
    case 802:
        v25 := v2[42]
        _ = v25
    default:
        var v26 int32 = 42
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
