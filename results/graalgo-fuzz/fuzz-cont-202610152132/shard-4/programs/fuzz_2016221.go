package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 uint16, p1 int8) int {
    switch -2147483648 {
    case -1:
        if true {
            var v6 bool = ((!true) && false)
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: (fi7 + 9223372036854775807)}
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := []int{fi7, -972, fi9, fi9}
                    _ = fi9
                    _ = v10
                }
                if (v6 || ((fi7 * fi7) > int(p1))) {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v6, v6)
                    case "hello":
                        fmt.Println(p0, p0, p0)
                    }
                    v11 := []int{fi7, fi7, fi7}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v11
                } else {
                    p1 += (p1 ^ p1)
                    v6 = v6
                }
                _ = fi7
                _ = v8
            }
            var v12 int8 = p1
            _ = v12
            _ = v6
        } else {
            var v13 int16 = -1
            _ = v13
        }
    case -219:
        fmt.Println(p1)
    }
    p0 = (p0 - (p0 << ((0) & 63)))
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-9223372036854775808, 490, 2147483647}
    var v3 uint64 = 42
    var v4 int16 = 1000
    v5 := []int{-2147483648, -9223372036854775808, 42}
    var v14 int8 = 127
    if (573 != 42) {
        switch "" {
        case "":
            for fi15 := 0; fi15 < 14; fi15++ {
                switch genFn0(65535, v14) {
                case 2147483647:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v4 -= (v4 / (((v4 / ((744) | 1))) | 1))
                        _ = fi16
                    }
                case -1:
                    v17 := GenS2{F0: -882, F1: 1}
                    _ = v17
                case 128:
                    var v18 int32 = -822
                    _ = v18
                default:
                    var v19 bool = false
                    _ = v19
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    v3 += 187
                    _ = fi20
                }
                _ = fi15
            }
        default:
            var v21 int32 = int32(v4)
            _ = v21
        }
    } else {
        if true {
            if ((!true) || ((32768 == genFn0(32767, v14)) && (42 > 65535))) {
                v22 := GenS3{F0: 42, F1: uint16(v3), F2: -128}
                fmt.Println(v4, v3, v3)
                v3 = 42
                _ = v22
            } else {
                v4 -= v4
                if false {
                    switch "a" {
                    case "a":
                        if false {
                            var v23 bool = false
                            v2 = append(v2, 32768)
                            if (1 != v4) {
                                v23 = (v23 || v23)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = v23
                        } else {
                            var v24 string = "foo bar"
                            var v25 int32 = -79
                            for fi26 := 0; fi26 < 20; fi26++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi26
                            }
                            _ = v24
                            _ = v25
                        }
                    case "":
                        v4 += 42
                    }
                    fmt.Println(v4, v4)
                    v27 := GenS3{F0: 128, F1: 910, F2: v14}
                    _ = v27
                }
            }
            v28 := GenS1{F0: -206}
            fmt.Println(v14, v3, v3)
            _ = v28
        }
    }
    v29 := GenS4{F0: 888, F1: true}
    fmt.Println(v3, v4)
    fmt.Println(v4, v3)
    var v30 string = ("\n\t" + "")
    v31 := GenS1{F0: -620}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 uint64 = (v3 & 72)
    v4 = v4
    if (!(1 == 575)) {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 uint64 = v32
            _ = fi33
            _ = v34
        }
        v3 -= v3
        v4 += (v4 + (v4 << ((979) & 63)))
    }
    if (3.14159 != (1.5 / (float64(float32((-2.25 - 1.5))) - -2.25))) {
        v35 := GenS2{F0: 627, F1: 4294967295}
        _ = v35
    }
    v36 := GenS5{F0: 255, F1: (!false)}
    if ((128 > 1) && false) {
        v37 := GenS6{F0: 0, F1: -821, F2: 4294967295}
        _ = v37
    } else {
        var v38 uint8 = 127
        v14 += (v14 * (v14 / ((1) | 1)))
        _ = v38
    }
    v3 -= (v3 - (v3 / ((1) | 1)))
    fmt.Println(v30, v14, v30)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
