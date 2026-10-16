package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint64, p1 bool) int16 {
    var v6 int64 = 663
    if (2147483648 <= 642) {
        fmt.Println(p1)
        if ((-1 < 1) && false) {
            v7 := []int{4294967295, 795, -1}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint64 = p0
                _ = fi8
                _ = v9
            }
        }
        v10 := GenS1{F0: 4294967295}
        _ = v10
    }
    p1 = p1
    v11 := GenS1{F0: 363}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -1
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 uint64 = 4294967295
    v12 := GenS2{F0: -488, F1: 0}
    v13 := GenS2{F0: 32767, F1: 468}
    v4[((("foo bar" + "\n\t") + "\n\t") + ("\n\t" + "\n\t"))] = 795
    var v14 int = 1
    var v15 uint = 65535
    if ((-1 == 42) && (78 >= 1)) {
        v16 := GenS1{F0: v15}
        fmt.Println(v2, v15, v14)
        var v17 bool = (!(!false))
        _ = v16
        _ = v17
    }
    v18 := []int{809, v14}
    var v19 uint8 = 255
    v20 := GenS1{F0: v15}
    v21 := GenS3{F0: uint8(v5), F1: v5}
    v19 -= (^(v19))
    switch (v14 >> (((v15 * v15)) & 63)) {
    case 32768:
        for fi22 := 0; fi22 < 14; fi22++ {
            v19 -= (v19 % (((v19 | 128)) | 1))
            v19 += (v19 - 205)
            _ = fi22
        }
    case 924:
        v15 = v15
    case -1:
        v2 = (v2 - (-(v2)))
    }
    var v23 int8 = 0
    v19 -= (v19 % ((v19) | 1))
    if false {
        var v24 uint16 = 42
        delete(v4, "\n\t")
        switch "hello" {
        case "a":
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        default:
            switch (v14 ^ -1) {
            case 0:
                v2 += v2
            default:
                if (1e10 <= 3.14159) {
                    v5 = (v5 << ((v15) & 63))
                    v25 := GenS1{F0: v15}
                    _ = v25
                } else {
                    v14 = (-(v14))
                    v26 := v1[(v14 & 32767)]
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v2 += (v2 % ((v2) | 1))
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v15, v2)
                        case "\n\t":
                            if (false || (true || true)) {
                                switch fi27 {
                                case 460:
                                    v28 := GenS2{F0: 0, F1: 32767}
                                    _ = v28
                                case 0:
                                    var v29 uint16 = uint16(v19)
                                    _ = v29
                                case 42:
                                    _ = len(v18)
                                default:
                                    v30 := GenS2{F0: 1, F1: -32768}
                                    _ = v30
                                }
                            }
                        default:
                            v15 += v15
                        }
                        _ = fi27
                    }
                    _ = v26
                }
            }
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
