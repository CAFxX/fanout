package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
func genFn0() bool {
    v7 := []int{-2147483648, 32768, -1, 32767}
    v8 := GenS1{F0: 128, F1: 364, F2: 1}
    _ = v7
    _ = v8
    return (1e10 < 0.1)
}

func main() {
    v1 := []int{197, 865, 42, 1}
    v2 := make(map[int]uint32)
    v3 := []uint32{363, 4294967295, 909, 1}
    var v4 uint64 = 4294967295
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    v4 = 568
    var v9 uint32 = 2147483648
    v4 += (v4 ^ v4)
    v10 := GenS1{F0: 246, F1: (v4 >> ((1) & 63)), F2: 44}
    v1 = append(v1, 1)
    _ = len(v3)
    for fi11 := 0; fi11 < 14; fi11++ {
        v9 += (v9 - v9)
        if false {
            switch "a" {
            case "\n\t":
                v4 += uint64(v9)
            case "a":
                v9 = (v9 * (v9 << ((2) & 63)))
            case "foo bar":
                fmt.Println(v9, fi11, v9)
            }
        } else {
            if true {
                var v12 float64 = 0.1
                _ = v12
            } else {
                if (0.0 < 1.5) {
                    if true {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v14 := GenS1{F0: 255, F1: (v4 >> ((uint(v9)) & 63)), F2: -1}
                            var v15 uint16 = 229
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        v16 := GenS2{F0: 4294967295}
                        v17 := GenS3{F0: 0}
                        fmt.Println(v4, fi11, v4)
                        _ = v16
                        _ = v17
                    }
                } else {
                    if false {
                        var v18 float64 = (1e10 - 0.0)
                        _ = v18
                    } else {
                        v9 -= v9
                    }
                    fmt.Println(v4, v4, v9)
                }
                switch (-(fi11)) {
                case 128:
                    var v19 int16 = 1000
                    _ = v19
                case 127:
                    fmt.Println(v9)
                case -9223372036854775808:
                    if (!(42 > 0)) {
                        var v20 uint = 1
                        v4 = (v4 + 4294967295)
                        _ = v20
                    }
                default:
                    v4 = v4
                }
                var v21 int32 = 42
                _ = v21
            }
            v9 -= (v9 << ((2) & 63))
        }
        v9 += (v9 ^ (v9 % (((v9 % ((0) | 1))) | 1)))
        _ = fi11
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int = (fi22 % ((42) | 1))
        v9 = v9
        switch "a" {
        case "\n\t":
            var v24 uint64 = 0
            _ = v24
        case "foo bar":
            v25 := GenS3{F0: 0}
            _ = v25
        case "hello":
            for fi26 := 0; fi26 < 6; fi26++ {
                if (!((v4 * (v4 % ((uint64(v9)) | 1))) < 4294967295)) {
                    if true {
                        fmt.Println(v9, v9)
                    }
                    switch (fi26 | 0) {
                    case 2147483647:
                        fmt.Println(v9)
                    case -621:
                        v23 = (fi26 - (fi26 % (((fi22 - (fi26 / ((v23) | 1)))) | 1)))
                    default:
                        if true {
                            fmt.Println(fi26, v9)
                            v27 := v2[-9223372036854775808]
                            _ = v27
                        } else {
                            v28 := v6[fi22]
                            _ = v28
                        }
                    }
                }
                if false {
                    var v29 uint8 = 128
                    _ = v29
                } else {
                    v9 = v9
                }
                if (-128 < -1) {
                    var v30 uint8 = 42
                    _ = v30
                }
                v23 += (fi26 - (v23 % (((-(fi22))) | 1)))
                _ = fi26
            }
        default:
            v4 = v4
        }
        _ = fi22
        _ = v23
    }
    var v31 int8 = -128
    var v32 uint16 = 255
    v31 -= v31
    if true {
        v32 = (v32 | (^(v32)))
        var v33 int32 = 82
        if (((true && false) && true) && (3.14159 > (1e10 / (1.5 - (3.14159 + -2.25))))) {
            for fi34 := 0; fi34 < 15; fi34++ {
                v2[(fi34 & fi34)] = 65535
                _ = fi34
            }
            _ = len(v1)
        } else {
            v3 = append(v3, (v9 - 2147483647))
        }
        _ = v33
    } else {
        for fi35 := 0; fi35 < 7; fi35++ {
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := GenS2{F0: 18446744073709551615}
                v9 -= v9
                _ = fi36
                _ = v37
            }
            fmt.Println(v9)
            _ = fi35
        }
        v32 = (v32 << ((0) & 63))
        v38 := GenS1{F0: uint8(v4), F1: v4, F2: v31}
        _ = v38
    }
    v31 += v31
    v4 += v4
    var v39 uint8 = 0
    v40 := GenS1{F0: (v39 * (v39 & (v39 * (v39 ^ 4)))), F1: v4, F2: v31}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
