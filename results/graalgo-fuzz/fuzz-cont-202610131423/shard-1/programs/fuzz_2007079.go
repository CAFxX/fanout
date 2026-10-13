package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint32) string {
    v7 := []int{602, 769, 529}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(fi8, p0)
            if (!true) {
                v7 = append(v7, fi9)
                if ((p0 << ((0) & 63)) <= 2) {
                    fmt.Println(fi9, p0)
                    var v10 int = fi9
                    if true {
                        var v11 int = -692
                        var v12 int8 = 10
                        _ = v11
                        _ = v12
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v13 int16 = 1000
                        if len(v7) > 0 {
                            v7[0] = (v10 % (((fi8 % (((v10 >> ((4294967295) & 63))) | 1))) | 1))
                        }
                        _ = v13
                    }
                    var v14 uint8 = 2
                    var v15 int = -9223372036854775808
                    _ = v10
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: fi9, F1: 54}
                    p0 -= (^(p0))
                    var v17 int = fi9
                    var v18 int8 = 100
                    _ = v16
                    _ = v17
                    _ = v18
                }
                v19 := GenS2{F0: 4294967295, F1: 100, F2: -559}
                var v20 uint16 = 42
                _ = v19
                _ = v20
            } else {
                if len(v7) > 0 {
                    v7[0] = fi9
                }
                fmt.Println(fi8, fi9, p0)
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi9
        }
        v21 := GenS3{F0: 1, F1: p0, F2: true}
        var v22 string = "a"
        _ = fi8
        _ = v21
        _ = v22
    }
    v23 := GenS4{F0: uint64(p0), F1: 42}
    _ = p0
    _ = v23
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int16 = 562
    v5 := make(map[string]string)
    var v6 uint32 = 1
    var v24 uint64 = 694
    switch "a" {
    case "hello":
        var v25 uint64 = v24
        _ = v25
    case "a":
        v26 := GenS5{F0: 0}
        _ = v26
    default:
        v24 = v24
    }
    v27 := GenS6{F0: 9223372036854775807, F1: 65535}
    var v28 float64 = -2.25
    fmt.Println(v6)
    switch -512 {
    case -706:
        v29 := v3["foo bar"]
        _ = v29
    case 1:
        v6 += 2147483648
    }
    var v30 float32 = ((((0.0 / -2.25) - (-2.25 + ((0.1 - 0.1) - 0.0))) + float32(v28)) - 0.0)
    if true {
        fmt.Println(v4)
    } else {
        for fi31 := 0; fi31 < 16; fi31++ {
            fmt.Println(v30, v28, fi31)
            v2[("" + "")] = fi31
            for fi32 := 0; fi32 < 14; fi32++ {
                v4 += (v4 ^ (^(v4)))
                var v33 bool = false
                var v34 int8 = 42
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    if false {
        if (v4 < -32768) {
            var v35 int = -9223372036854775808
            switch v35 {
            case 721:
                v28 = 3.14159
            case 128:
                if false {
                    var v36 int16 = (v4 - v4)
                    _ = v36
                }
            case 32767:
                var v37 uint64 = (v24 ^ v24)
                _ = v37
            default:
                var v38 int16 = 0
                _ = v38
            }
            var v39 int16 = v4
            _ = v35
            _ = v39
        }
        switch "a" {
        case "a":
            v40 := GenS5{F0: -1}
            _ = v40
        case "foo bar":
            if (42 <= 1) {
                v4 -= (v4 ^ 338)
            }
        default:
            _ = len(v1)
        }
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        delete(v3, "\n\t")
        v1[4294967295] = ("a" + ("a" + ""))
        switch fi41 {
        case -904:
            var v42 uint8 = 130
            _ = v42
        case 969:
            var v43 int8 = 42
            _ = v43
        default:
            v4 -= 0
        }
        _ = fi41
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        v45 := GenS6{F0: int64(v4), F1: uint16(v24)}
        v46 := []int{336, 4294967295}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 int = 1
    v48 := []int{v47, (-(v47)), -2147483648, v47}
    v30 -= (v30 - ((v30 + v30) + float32((v28 * 0.1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
