package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -78
    var v3 uint16 = 2
    var v4 int64 = 42
    var v5 uint32 = 4294967295
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 42, F1: 2, F2: 2147483647}
    v5 -= v5
    var v8 uint16 = v3
    v2 = (-(v2))
    var v9 int64 = (v4 | v4)
    v10 := GenS2{F0: 26, F1: "\n\t", F2: (v4 >> ((670) & 63))}
    v11 := GenS1{F0: 297, F1: 1, F2: v2}
    v12 := []int{}
    var v13 uint32 = (v5 << ((614) & 63))
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 65535, F1: (v3 - v3), F2: v2}
        for fi16 := 0; fi16 < 13; fi16++ {
            switch int(v3) {
            case -744:
                if (0 > 0) {
                    var v17 int = 610
                    _ = v17
                }
            case -2147483648:
                var v18 uint8 = 0
                _ = v18
            default:
                v12 = append(v12, (fi16 ^ fi16))
            }
            _ = fi16
        }
        if (v8 >= (v3 % ((v8) | 1))) {
            if true {
                v9 = v4
            } else {
                v3 += 42
            }
        } else {
            switch fi14 {
            case 32768:
                v19 := GenS3{F0: v2, F1: 42}
                _ = v19
            case -545:
                switch int(v8) {
                case 127:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        switch fi20 {
                        case 830:
                            var v21 bool = ((v9 % (((v4 & (v4 / ((v9) | 1)))) | 1)) < (v9 << ((18446744073709551615) & 63)))
                            _ = v21
                        case 899:
                            if len(v12) > 0 {
                                v12[0] = (-(fi14))
                            }
                        }
                        _ = fi20
                    }
                }
            case 160:
                var v22 float64 = 0.1
                _ = v22
            }
            if true {
                fmt.Println(v5, fi14)
                var v23 bool = false
                v6["hello"] = v13
                _ = v23
            }
            v2 = -1
        }
        _ = fi14
        _ = v15
    }
    v4 = v4
    switch int(v13) {
    case 284:
        var v24 string = "foo bar"
        _ = v24
    default:
        var v25 int8 = int8(v9)
        _ = v25
    }
    if false {
        _ = len(v1)
        if true {
            v4 -= 600
        } else {
            v2 = 114
            v5 = v13
            var v26 uint32 = (v5 / ((v5) | 1))
            _ = v26
        }
    } else {
        v5 += (v5 | 65535)
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
