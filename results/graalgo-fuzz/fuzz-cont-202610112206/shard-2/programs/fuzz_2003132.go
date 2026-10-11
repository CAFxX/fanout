package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 bool }
func genFn0(p0 bool, p1 uint8, p2 bool) uint32 {
    var v5 int8 = 100
    switch -1 {
    case 369:
        switch ("\n\t" + "a") {
        case "hello":
            v6 := []int{}
            _ = v6
        case "\n\t":
            var v7 int = 128
            _ = v7
        default:
            if p0 {
                if p2 {
                    v8 := []int{4294967295}
                    switch "" {
                    case "\n\t":
                        for fi9 := 0; fi9 < 14; fi9++ {
                            for fi10 := 0; fi10 < 6; fi10++ {
                                v5 = -128
                                _ = fi10
                            }
                            _ = fi9
                        }
                    default:
                        for fi11 := 0; fi11 < 19; fi11++ {
                            for fi12 := 0; fi12 < 5; fi12++ {
                                var v13 float32 = 3.14159
                                fmt.Println(fi12, p1, fi12)
                                if false {
                                    var v14 float32 = v13
                                    v15 := GenS1{F0: 2, F1: p1}
                                    fmt.Println(fi11)
                                    _ = v14
                                    _ = v15
                                }
                                var v16 int64 = 9223372036854775807
                                var v17 uint16 = 65535
                                v18 := GenS1{F0: uint(p1), F1: p1}
                                _ = fi12
                                _ = v13
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v5)
                            _ = fi11
                        }
                    }
                    p0 = true
                    _ = v8
                }
                v19 := GenS1{F0: 4294967295, F1: 255}
                _ = v19
            }
        }
    case 32767:
        v20 := []int{0, -648, -9223372036854775808, -2147483648}
        _ = v20
    case 849:
        var v21 uint = 338
        _ = v21
    default:
        v5 += (^(v5))
    }
    var v22 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 int64, p2 int) int32 {
    if true {
        var v23 uint8 = 1
        _ = v23
    } else {
        fmt.Println(p0)
    }
    if false {
        v24 := []int{p2, 809, 32768, p2}
        _ = v24
    }
    var v25 float64 = 1.5
    if (false || true) {
        v26 := GenS2{F0: (p1 - p1), F1: p0}
        _ = v26
    } else {
        for fi27 := 0; fi27 < 18; fi27++ {
            p0 = 1
            v28 := []int{4294967295, p2, p2, p2}
            _ = fi27
            _ = v28
        }
        var v29 uint32 = 4294967295
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 2147483647
}

func main() {
    var v1 uint = 2
    v2 := []uint32{42, 2147483647, 473, 491}
    v3 := make(map[uint32]int)
    var v4 int = 32767
    v30 := GenS3{F0: 42, F1: 678, F2: false}
    v1 = v1
    var v31 int = v4
    if (!(0 > 218)) {
        v31 -= v4
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        switch "" {
        case "hello":
            for fi33 := 0; fi33 < 2; fi33++ {
                v1 = v1
                fmt.Println(fi32)
                var v34 string = "a"
                fmt.Println(fi33, v31, fi33)
                _ = fi33
                _ = v34
            }
        }
        _ = fi32
    }
    _ = len(v2)
    var v35 int64 = 42
    v1 -= (v1 + 2)
    v31 += (v31 >> ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
}
