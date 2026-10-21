package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint = 489
    var v2 uint32 = 65535
    v3 := make(map[uint32]uint32)
    v2 = (v2 - 194)
    var v4 float64 = (1.5 + (1.5 * 1.5))
    if false {
        var v5 uint16 = 619
        var v6 int8 = -57
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 += 1
            _ = fi7
        }
    }
    v4 = 3.14159
    v8 := GenS1{F0: v2, F1: false, F2: 255}
    switch -1 {
    case 127:
        fmt.Println(v2)
    case 4294967295:
        fmt.Println(v4, v4, v1)
    case -344:
        if (true || true) {
            v9 := GenS2{F0: 18446744073709551615, F1: -1}
            switch 2147483647 {
            case 296:
                fmt.Println(v2, v2, v4)
            case -2147483648:
                if (42 <= 52) {
                    v1 += (v1 << (((v1 % ((255) | 1))) & 63))
                    v10 := GenS3{F0: -32768}
                    v2 -= v2
                    _ = v10
                }
            default:
                var v11 uint32 = 4294967295
                _ = v11
            }
            _ = v9
        }
    }
    if false {
        v4 += v4
        for fi12 := 0; fi12 < 14; fi12++ {
            if (v1 > 42) {
                if (-128 > 100) {
                    v3[v2] = (^(v2))
                } else {
                    v13 := []int{(-(fi12)), (-(fi12))}
                    v4 -= 1e10
                    v4 = (3.14159 + 0.1)
                    _ = v13
                }
                if false {
                    v14 := []int{}
                    if false {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            var v16 uint64 = 42
                            var v17 float64 = v4
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 1; fi18++ {
                            switch fi12 {
                            case -9223372036854775808:
                                var v19 uint64 = 0
                                _ = v19
                            case 32767:
                                v20 := GenS1{F0: v2, F1: (-555 >= -2147483648), F2: 0}
                                _ = v20
                            }
                            var v21 int64 = int64(v2)
                            switch "hello" {
                            case "foo bar":
                                v3[99] = (v2 / (((^(v2))) | 1))
                            }
                            v4 = v4
                            _ = len(v3)
                            _ = fi18
                            _ = v21
                        }
                    }
                    _ = v14
                }
            } else {
                var v22 int32 = -2147483648
                v1 = 1
                _ = v22
            }
            v4 -= v4
            _ = fi12
        }
        var v23 int = -9223372036854775808
        _ = v23
    } else {
        if true {
            var v24 string = "\n\t"
            _ = v24
        }
        switch 127 {
        case 744:
            var v25 float32 = (1e10 + 1.5)
            _ = v25
        case 1:
            var v26 int32 = -2147483648
            _ = v26
        case 89:
            v1 = uint(v2)
        default:
            v27 := GenS4{F0: (v2 % ((v2) | 1)), F1: v1}
            _ = v27
        }
    }
    v28 := GenS5{F0: 255, F1: -922}
    v29 := GenS1{F0: (v2 << ((v1) & 63)), F1: false, F2: 255}
    v2 += 4294967295
    var v30 uint16 = 65535
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 bool = true
        switch "hello" {
        case "\n\t":
            v30 = 1
        case "foo bar":
            v2 -= 2147483647
        }
        var v33 uint32 = (v2 + 4294967295)
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := []int{128}
    v1 = 550
    switch 2147483647 {
    case 127:
        var v35 bool = true
        _ = v35
    case 32767:
        v34 = append(v34, 387)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v8
}
