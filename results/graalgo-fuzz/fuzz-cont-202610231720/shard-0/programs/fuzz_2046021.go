package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint, p1 uint) uint16 {
    fmt.Println(p1)
    switch -9223372036854775808 {
    case 245:
        v5 := GenS1{F0: 1}
        _ = v5
    case 0:
        switch "foo bar" {
        case "hello":
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 int8 = 1
                _ = fi6
                _ = v7
            }
        case "\n\t":
            fmt.Println(p1, p0)
        }
    default:
        if (701 == 362) {
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int16 = 42
                    v11 := []int{fi9, (fi9 * fi9), fi8}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                p0 += 875
                switch ("hello" + "") {
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi8)
                    case "\n\t":
                        p1 += (p1 | p1)
                    case "hello":
                        switch 0 {
                        case 289:
                            var v12 int16 = -768
                            _ = v12
                        case 127:
                            fmt.Println(p1, p1, p1)
                        case -1:
                            fmt.Println(fi8)
                        }
                    default:
                        p1 += 255
                    }
                }
                _ = fi8
            }
            if ((0 == 77) && false) {
                fmt.Println(p1)
                v13 := GenS2{F0: 869, F1: 4294967295, F2: 18446744073709551615}
                _ = v13
            }
            v14 := []int{-2147483648, -2147483648}
            _ = v14
        }
    }
    var v15 int8 = -76
    _ = p0
    _ = p1
    _ = v15
    return 458
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := []int{}
    v4 := make(map[string]bool)
    var v16 uint8 = 255
    v17 := GenS1{F0: 1}
    v16 -= v16
    var v18 int8 = 42
    if ((false || true) || true) {
        if (540 <= 255) {
            var v19 uint32 = 2147483648
            switch -2147483648 {
            case 128:
                v1[(^(v19))] = -2147483648
            case 0:
                var v20 int = 42
                _ = v20
            case -2147483648:
                var v21 int = 746
                _ = v21
            }
            var v22 uint16 = uint16(v19)
            _ = v19
            _ = v22
        } else {
            for fi23 := 0; fi23 < 6; fi23++ {
                v18 -= 127
                for fi24 := 0; fi24 < 6; fi24++ {
                    switch -294 {
                    case 961:
                        fmt.Println(v18, v16)
                    }
                    _ = fi24
                }
                v16 = v16
                _ = fi23
            }
            var v25 uint64 = 376
            _ = v25
        }
        var v26 int64 = 42
        _ = v26
    } else {
        v18 -= -128
        if false {
            _ = len(v3)
            v18 += -120
        }
    }
    var v27 uint32 = 566
    var v28 uint64 = 335
    switch "a" {
    case "":
        fmt.Println(v16, v16, v28)
    default:
        v27 = v27
    }
    switch 4294967295 {
    case 32768:
        var v29 float64 = 1.5
        _ = v29
    default:
        v28 = 4294967295
    }
    v28 = (v28 / (((^(v28))) | 1))
    var v30 bool = false
    v16 = v16
    v3 = append(v3, 9223372036854775807)
    var v31 int16 = -1
    v30 = ((v27 ^ (v27 << ((861) & 63))) <= 4294967295)
    v32 := GenS3{F0: (!v30), F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
