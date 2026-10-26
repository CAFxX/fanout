package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint) int {
    p0 -= 0
    if (741 > 35) {
        if false {
            v4 := []int{0, 4294967295, 5, -428}
            p0 += (p0 | (p0 % ((p0) | 1)))
            _ = v4
        }
    } else {
        switch 9223372036854775807 {
        case 4294967295:
            switch "\n\t" {
            case "\n\t":
                v5 := []int{-271, -2147483648, -1, 32767}
                _ = v5
            case "foo bar":
                switch ("foo bar" + "a") {
                case "\n\t":
                    if true {
                        switch "foo bar" {
                        case "a":
                            for fi6 := 0; fi6 < 1; fi6++ {
                                p0 += p0
                                v7 := []int{-9223372036854775808}
                                _ = fi6
                                _ = v7
                            }
                        case "":
                            if true {
                                p0 -= p0
                                v8 := GenS1{F0: (^(p0))}
                                fmt.Println(p0)
                                _ = v8
                            }
                        case "foo bar":
                            var v9 uint64 = 1
                            _ = v9
                        default:
                            v10 := GenS1{F0: (^(p0))}
                            _ = v10
                        }
                        p0 = (p0 << (((p0 & (p0 | p0))) & 63))
                    }
                case "foo bar":
                    v11 := GenS1{F0: p0}
                    _ = v11
                case "hello":
                    fmt.Println(p0)
                default:
                    var v12 uint16 = 0
                    _ = v12
                }
            default:
                v13 := []int{-947}
                _ = v13
            }
        }
        v14 := GenS1{F0: (p0 | p0)}
        fmt.Println(p0)
        _ = v14
    }
    _ = p0
    return -1
}
func genFn1(p0 int16) int16 {
    p0 += p0
    switch "" {
    case "":
        p0 = p0
    case "\n\t":
        v15 := GenS2{F0: 42, F1: 42}
        _ = v15
    case "hello":
        if true {
            v16 := GenS3{F0: 255, F1: 151, F2: 978}
            _ = v16
        }
    }
    _ = p0
    return p0
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]uint32)
    var v3 uint64 = 182
    v3 -= (^(v3))
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint64 = v3
        var v19 uint64 = v18
        _ = fi17
        _ = v18
        _ = v19
    }
    v1 = 0
    var v20 int16 = 32767
    v21 := []int{-890, -52, 9223372036854775807, -1}
    v20 = genFn1(v20)
    var v22 int32 = 535
    v1 = v1
    if true {
        if (128 == (v1 >> ((170) & 63))) {
            v1 -= v1
            v3 = v3
            var v23 uint64 = (v3 & 2)
            _ = v23
        } else {
            var v24 int = -2147483648
            _ = v24
        }
    } else {
        v25 := GenS4{F0: 296, F1: (v3 << ((18446744073709551615) & 63))}
        _ = v25
    }
    var v26 int32 = v22
    if (!(("\n\t" > "") && false)) {
        fmt.Println(v3, v20)
        var v27 uint64 = (v3 - uint64(v1))
        v26 -= -2147483648
        _ = v27
    } else {
        if (0 >= 32768) {
            v20 -= v20
            for fi28 := 0; fi28 < 20; fi28++ {
                v20 -= v20
                _ = fi28
            }
            if false {
                var v29 int32 = 2147483647
                v20 = (^(v20))
                _ = v29
            }
        } else {
            switch "hello" {
            case "hello":
                fmt.Println(v26)
            default:
                delete(v2, -392)
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 string = "hello"
                _ = fi30
                _ = v31
            }
        }
        fmt.Println(v1, v22, v26)
        var v32 uint8 = 128
        _ = v32
    }
    fmt.Println(v3, v22)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
}
