package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
func genFn0() uint {
    v7 := GenS1{F0: 251}
    v8 := []int{-1, -678}
    _ = v7
    _ = v8
    return 192
}

func main() {
    var v1 uint8 = 2
    var v2 int = 463
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 float32 = float32(1e10)
    var v6 uint64 = 42
    var v9 uint8 = (v1 & 0)
    v10 := []int{(v2 / (((-(v2))) | 1))}
    var v11 float64 = (1.5 + float64(v5))
    var v12 float32 = v5
    var v13 int16 = 32767
    switch int(v9) {
    case 32767:
        var v14 uint32 = 4294967295
        _ = v14
    case 32768:
        var v15 uint32 = 2147483647
        _ = v15
    case 4294967295:
        v16 := GenS1{F0: (v6 ^ v6)}
        _ = v16
    }
    var v17 int32 = -687
    if (((v6 & v6) != 2) && (false || (!((-(v2)) < v2)))) {
        v18 := GenS2{F0: (v17 - (v17 ^ (-(v17)))), F1: ((!false) || (v17 >= (v17 / ((v17) | 1)))), F2: -128}
        var v19 uint16 = 474
        var v20 int8 = 100
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        switch "hello" {
        case "a":
            v22 := GenS3{F0: -121, F1: false, F2: 638}
            _ = v22
        case "\n\t":
            if ((false || false) && (true || (!(!false)))) {
                fmt.Println(v12)
                v1 = (^(v1))
                if (0 == 4294967295) {
                    if true {
                        var v23 int16 = 32767
                        _ = v23
                    }
                }
            }
        default:
            v10 = append(v10, (-(v2)))
        }
        if true {
            for fi24 := 0; fi24 < 5; fi24++ {
                switch (fi24 * fi24) {
                case 0:
                    var v25 string = "foo bar"
                    _ = v25
                }
                _ = fi24
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                v10 = append(v10, fi21)
                var v27 bool = true
                switch 2147483647 {
                case 4294967295:
                    switch (v2 ^ fi21) {
                    case -613:
                        v9 -= (^(v9))
                    case 4294967295:
                        switch 2147483647 {
                        case 42:
                            v28 := GenS3{F0: 7, F1: v27, F2: 32767}
                            _ = v28
                        }
                    }
                case 42:
                    fmt.Println(v9, v27)
                case 1:
                    v1 = (v1 % ((v9) | 1))
                default:
                    switch 2147483647 {
                    case 32768:
                        v13 = (v13 - (v13 << ((255) & 63)))
                    }
                }
                _ = fi26
                _ = v27
            }
            if len(v10) > 0 {
                v10[0] = (v2 & (fi21 | 32767))
            }
        }
        _ = fi21
    }
    switch v2 {
    case -9223372036854775808:
        var v29 uint64 = v6
        _ = v29
    default:
        var v30 int64 = 0
        _ = v30
    }
    var v31 uint64 = 249
    var v32 uint8 = v9
    fmt.Println(v13)
    v33 := GenS2{F0: -681, F1: (v5 < 1.5), F2: 104}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
