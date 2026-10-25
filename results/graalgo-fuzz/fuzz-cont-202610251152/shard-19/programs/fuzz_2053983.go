package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 uint8 = 128
    v5 := []int{-374, 1, 128}
    v6 := []uint32{2147483647, 2147483647, 0}
    if (((1e10 * 0.0) != 3.14159) && (168 < 42)) {
        if (-735 >= 1) {
            var v7 int = 32767
            for fi8 := 0; fi8 < 10; fi8++ {
                if (!false) {
                    switch "a" {
                    case "\n\t":
                        for fi9 := 0; fi9 < 8; fi9++ {
                            var v10 float64 = 1e10
                            var v11 int8 = -1
                            v12 := GenS1{F0: "foo bar", F1: (0 != 364), F2: -1000}
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        v13 := GenS2{F0: 42}
                        _ = v13
                    case "":
                        for fi14 := 0; fi14 < 2; fi14++ {
                            fmt.Println(v4)
                            v15 := GenS3{F0: (fi8 + (-(fi14)))}
                            _ = fi14
                            _ = v15
                        }
                    default:
                        fmt.Println(v7)
                    }
                } else {
                    var v16 uint64 = 42
                    _ = v16
                }
                v4 = v4
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi8
            }
            _ = v7
        } else {
            var v17 uint = 255
            _ = v17
        }
    } else {
        v3 -= (v3 << ((557) & 63))
    }
    v4 = v4
    v4 = v4
    switch "foo bar" {
    case "foo bar":
        v3 = (v3 | (v3 % ((v3) | 1)))
    case "a":
        for fi18 := 0; fi18 < 1; fi18++ {
            delete(v1, (^(fi18)))
            _ = fi18
        }
    default:
        v4 += v4
    }
    v19 := GenS4{F0: 719}
    fmt.Println(v4, v4)
    if false {
        v3 = 0
        v3 = (v3 & (-(v3)))
        if false {
            v6 = append(v6, 920)
        }
    }
    if (true || (!false)) {
        fmt.Println(v3, v3)
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    var v21 bool = true
    var v22 float32 = ((-2.25 - 1e10) * 1e10)
    switch 428 {
    case 4294967295:
        v23 := GenS5{F0: (v4 << ((1) & 63))}
        _ = v23
    case 88:
        fmt.Println(v21, v21, v22)
    }
    var v24 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
