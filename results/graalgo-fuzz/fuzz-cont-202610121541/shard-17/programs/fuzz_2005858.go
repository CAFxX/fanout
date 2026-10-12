package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int32 = -443
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int = 869
    v5 := GenS1{F0: (!false)}
    var v6 float64 = 3.14159
    var v7 uint32 = 2147483648
    v8 := []int{}
    if false {
        _ = len(v8)
        if (!((0.0 != 3.14159) && true)) {
            v7 -= (v7 >> ((300) & 63))
            var v9 uint64 = 18446744073709551615
            _ = v9
        } else {
            v10 := GenS2{F0: 835}
            v11 := GenS3{F0: 483, F1: -5}
            if true {
                v6 += (v6 + (1.5 + (1e10 / 3.14159)))
                if false {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 uint32 = 2147483647
                        _ = fi12
                        _ = v13
                    }
                }
            } else {
                switch (("hello" + ("hello" + ("hello" + "hello"))) + "") {
                case "\n\t":
                    switch v4 {
                    case 842:
                        var v14 uint32 = v7
                        _ = v14
                    case 996:
                        for fi15 := 0; fi15 < 14; fi15++ {
                            for fi16 := 0; fi16 < 9; fi16++ {
                                var v17 int32 = 1
                                v1 = v1
                                _ = fi16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 15; fi18++ {
                                v19 := GenS1{F0: (-360 == -1)}
                                v20 := GenS4{F0: int16(v7), F1: 2}
                                v6 = v6
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = fi15
                        }
                    default:
                        switch v4 {
                        case -713:
                            fmt.Println(v1, v7, v6)
                        case 2147483647:
                            v21 := GenS2{F0: uint(v7)}
                            _ = v21
                        default:
                            fmt.Println(v1, v1)
                        }
                    }
                default:
                    var v22 uint64 = 2
                    _ = v22
                }
                for fi23 := 0; fi23 < 16; fi23++ {
                    if false {
                        var v24 int8 = 100
                        var v25 int8 = v24
                        var v26 uint = 42
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS3{F0: 565, F1: -128}
                    _ = fi23
                    _ = v27
                }
            }
            _ = v10
            _ = v11
        }
        var v28 int64 = 42
        _ = v28
    }
    if false {
        if (!((!(true || false)) && true)) {
            v4 = (v4 >> ((1) & 63))
        } else {
            v7 = (^(v7))
        }
    }
    switch v4 {
    case -594:
        var v29 int16 = 32767
        _ = v29
    case 127:
        v30 := GenS3{F0: 336, F1: 7}
        _ = v30
    case 4294967295:
        switch "" {
        case "a":
            var v31 uint8 = 128
            _ = v31
        default:
            v32 := GenS4{F0: -401, F1: 1}
            _ = v32
        }
    }
    v1 -= (v1 / ((-2147483648) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
