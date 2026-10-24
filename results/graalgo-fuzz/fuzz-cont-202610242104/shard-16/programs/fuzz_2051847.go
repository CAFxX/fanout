package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    v3 = v3
    if false {
        var v4 uint = 0
        _ = v4
    }
    switch -440 {
    case 128:
        v5 := GenS1{F0: 18446744073709551615, F1: v1, F2: 0}
        _ = v5
    }
    var v6 bool = true
    v3 += v3
    fmt.Println(v6)
    fmt.Println(v1)
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            delete(v2, 0)
            switch fi7 {
            case 4294967295:
                var v9 int16 = -603
                _ = v9
            }
            var v10 int = (fi8 << ((42) & 63))
            if false {
                switch "a" {
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        v11 := []int{v10, (v10 & v10), (v10 << ((820) & 63)), (fi7 >> ((42) & 63))}
                        _ = v11
                    case "foo bar":
                        if (!false) {
                            v12 := GenS2{F0: v3, F1: v6}
                            var v13 uint32 = 42
                            var v14 uint32 = v13
                            _ = v12
                            _ = v13
                            _ = v14
                        } else {
                            var v15 float32 = (3.14159 + float32((1e10 - float64(1e10))))
                            if (65535 > 65535) {
                                switch fi8 {
                                case 9223372036854775807:
                                    v2[0] = 1
                                case 32767:
                                    v15 -= v15
                                }
                            }
                            var v16 int16 = 195
                            _ = v15
                            _ = v16
                        }
                    }
                case "":
                    var v17 string = ("a" + ("hello" + "\n\t"))
                    _ = v17
                default:
                    var v18 uint8 = 1
                    _ = v18
                }
            }
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    var v19 uint16 = 1
    switch -9223372036854775808 {
    case -784:
        fmt.Println(v6)
    default:
        v20 := GenS3{F0: 2, F1: 42, F2: 523}
        _ = v20
    }
    fmt.Println(v19, v3, v3)
    v21 := v2[65535]
    for fi22 := 0; fi22 < 12; fi22++ {
        v21 = (v21 * (v21 % (((v21 >> ((42) & 63))) | 1)))
        var v23 uint16 = v19
        v24 := []int{}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = 892
    v21 -= (v21 ^ v21)
    delete(v2, 887)
    v3 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v6
}
