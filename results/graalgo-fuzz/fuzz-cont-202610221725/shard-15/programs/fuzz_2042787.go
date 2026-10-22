package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := []int{-536, 383}
    if (false || false) {
        v5 := v2[("a" + "")]
        _ = len(v4)
        _ = v5
    } else {
        v1 -= (v1 + 65535)
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 int8 = 16
    v8 := GenS1{F0: 257, F1: 128}
    if (79 == 2147483647) {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: (fi9 % (((^(fi9))) | 1))}
            _ = fi9
            _ = v10
        }
    } else {
        fmt.Println(v1, v1)
        var v11 int32 = 2147483647
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: -874, F1: 42}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if (!(0 <= 466)) {
            v1 -= v1
        } else {
            fmt.Println(v1)
            for fi15 := 0; fi15 < 20; fi15++ {
                if (false || ((true && true) && false)) {
                    if true {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = fi15
                            }
                            var v16 uint64 = 731
                            if len(v4) > 0 {
                                v4[0] = 236
                            }
                            _ = v16
                        } else {
                            var v17 uint16 = 65535
                            fmt.Println(fi14, fi14)
                            _ = v17
                        }
                        if (!true) {
                            v1 -= v1
                            v18 := GenS1{F0: -740, F1: fi15}
                            v1 = (v1 % ((4294967295) | 1))
                            _ = v18
                        }
                    }
                }
                v4 = append(v4, fi15)
                switch "foo bar" {
                case "hello":
                    v4 = append(v4, int(v1))
                }
                var v19 int64 = 42
                _ = fi15
                _ = v19
            }
            switch "\n\t" {
            case "foo bar":
                v20 := GenS1{F0: 1, F1: fi14}
                _ = v20
            default:
                delete(v3, fi14)
            }
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v7 -= (-(v7))
        _ = fi21
    }
    v1 -= v1
    var v22 int32 = 1
    var v23 int32 = (v22 >> ((255) & 63))
    v1 -= (v1 / (((v1 / ((18446744073709551615) | 1))) | 1))
    _ = len(v3)
    fmt.Println(v23)
    var v24 int16 = -934
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
