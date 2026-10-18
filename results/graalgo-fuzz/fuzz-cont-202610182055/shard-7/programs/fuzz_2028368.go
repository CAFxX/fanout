package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 bool, p1 uint, p2 uint) int {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: fi6, F1: 42}
        if (false && (2 != 255)) {
            p1 = (p2 << ((p1) & 63))
            var v8 int8 = 1
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 int64 = -9223372036854775808
    var v10 int = -9223372036854775808
    var v11 uint32 = 806
    switch "a" {
    case "foo bar":
        v12 := GenS1{F0: (v10 / ((v10) | 1)), F1: 267}
        _ = v12
    case "\n\t":
        p0 = (2 > p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (-(v10))
}
func genFn1(p0 string, p1 string, p2 int) uint {
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            fmt.Println(fi13, fi13, fi13)
        }
        switch fi13 {
        case 9223372036854775807:
            if false {
                v14 := []int{-9223372036854775808}
                if ((0.0 < 1e10) || false) {
                    v15 := GenS2{F0: 796, F1: -1, F2: 194}
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v15
                } else {
                    var v16 float64 = float64((3.14159 * 0.1))
                    switch (fi13 >> ((18446744073709551615) & 63)) {
                    case 32768:
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    case 32767:
                        p2 -= (fi13 ^ genFn0(false, 881, 498))
                    case -9223372036854775808:
                        var v17 int16 = -590
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 18; fi18++ {
                            if (!(!((-9223372036854775808 <= -572) || ((44 <= 0) && true)))) {
                                var v19 uint32 = 42
                                if false {
                                    var v20 float64 = 0.0
                                    _ = v20
                                }
                                _ = v19
                            }
                            switch (p2 << ((42) & 63)) {
                            case 1:
                                p1 = "foo bar"
                            }
                            _ = fi18
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        var v21 int64 = 1
                        _ = v21
                    case "":
                        for fi22 := 0; fi22 < 15; fi22++ {
                            var v23 uint16 = 905
                            v23 = 42
                            _ = len(v14)
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = v16
                }
                _ = v14
            }
        case 0:
            var v24 uint = 690
            _ = v24
        case 32767:
            if (true || false) {
                var v25 uint = 566
                _ = v25
            } else {
                switch p1 {
                case "foo bar":
                    var v26 float32 = -2.25
                    _ = v26
                case "\n\t":
                    var v27 uint8 = 0
                    _ = v27
                case "a":
                    v28 := []int{(^(p2)), (fi13 + fi13), (fi13 / ((-715) | 1)), -9223372036854775808}
                    _ = v28
                default:
                    v29 := GenS3{F0: 127, F1: 2147483648}
                    _ = v29
                }
                v30 := GenS3{F0: 42, F1: 298}
                v31 := GenS1{F0: -589, F1: 137}
                _ = v30
                _ = v31
            }
        default:
            if true {
                fmt.Println(p0, p2)
                for fi32 := 0; fi32 < 17; fi32++ {
                    if false {
                        p2 -= 127
                    }
                    p2 = p2
                    var v33 string = p0
                    _ = fi32
                    _ = v33
                }
                var v34 int = (fi13 | (fi13 << ((2) & 63)))
                _ = v34
            } else {
                var v35 int8 = -51
                v36 := []int{1, (fi13 ^ genFn0(((0 < 519) || false), 0, 65535)), 42}
                _ = v35
                _ = v36
            }
        }
        _ = fi13
    }
    p2 -= (p2 | p2)
    fmt.Println(p0)
    v37 := GenS4{F0: 42, F1: -380, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    return 255
}

func main() {
    var v1 float64 = (3.14159 * (1e10 + (float64((0.0 + (0.1 + (0.1 * (float32(1e10) * (-2.25 * (float32(-2.25) * (-2.25 * ((((0.1 + 0.0) - 1e10) + 0.1) * 3.14159))))))))) - float64(0.1))))
    v2 := []uint32{656, 396, 42, 2147483648}
    var v3 uint8 = 104
    var v4 uint64 = 42
    v5 := make(map[int]uint32)
    v3 += v3
    v3 -= 0
    v38 := GenS4{F0: 2147483648, F1: 1, F2: -9223372036854775808}
    _ = len(v5)
    fmt.Println(v4, v3)
    v1 -= ((((float64(((0.1 + ((1.5 * 0.1) / 1e10)) + 1.5)) * v1) + v1) * (v1 / v1)) * v1)
    v4 += 645
    _ = len(v2)
    var v39 uint8 = (v3 << ((18446744073709551615) & 63))
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
