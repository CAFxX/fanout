package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 int32) int8 {
    fmt.Println(p0)
    switch 1 {
    case 856:
        if true {
            switch 42 {
            case 32767:
                for fi5 := 0; fi5 < 6; fi5++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v6 int8 = 0
                        _ = v6
                    case "foo bar":
                        p0 = (p0 + (p0 ^ (^(p0))))
                    default:
                        v7 := []int{fi5, fi5}
                        _ = v7
                    }
                    var v8 uint32 = 2147483647
                    switch "a" {
                    case "a":
                        var v9 uint16 = 32767
                        _ = v9
                    case "foo bar":
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v8, p0)
                        default:
                            if (!true) {
                                switch "hello" {
                                case "foo bar":
                                    if ((-128 <= -118) && (323 != 991)) {
                                        p0 += 154
                                    }
                                case "hello":
                                    v10 := []int{fi5, fi5, fi5, fi5}
                                    _ = v10
                                case "a":
                                    v8 += v8
                                default:
                                    p0 -= p0
                                }
                            } else {
                                var v11 uint8 = 1
                                v12 := GenS1{F0: 1, F1: 18446744073709551615}
                                v13 := []int{fi5, (fi5 - (fi5 >> ((800) & 63))), 0, (fi5 >> ((967) & 63))}
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                        }
                    case "\n\t":
                        var v14 int64 = 1
                        _ = v14
                    }
                    _ = fi5
                    _ = v8
                }
            case 2147483647:
                v15 := []int{-9223372036854775808, 1, 4294967295, 32768}
                _ = v15
            }
            switch -2147483648 {
            case -1:
                p0 = p0
            default:
                var v16 uint = 2
                _ = v16
            }
            fmt.Println(p0)
        }
    }
    _ = p0
    return -1
}
func genFn1(p0 uint64, p1 uint) uint {
    v17 := GenS1{F0: (p0 & 0), F1: (p0 * p0)}
    var v18 float32 = (float32(0.0) * 1.5)
    var v19 int32 = 766
    var v20 int = 0
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return (p1 - p1)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    fmt.Println(v2)
    v2 = -128
    var v21 int = 892
    v2 = v2
    if false {
        v21 += 32768
    }
    v22 := []int{}
    switch -988 {
    case 0:
        var v23 int64 = -9223372036854775808
        _ = v23
    case 32767:
        v2 -= (v2 % ((1) | 1))
    case 4294967295:
        v24 := GenS2{F0: (3.14159 < 3.14159), F1: 255}
        _ = v24
    default:
        v21 = (v21 + (v21 ^ 127))
    }
    var v25 uint = 704
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
