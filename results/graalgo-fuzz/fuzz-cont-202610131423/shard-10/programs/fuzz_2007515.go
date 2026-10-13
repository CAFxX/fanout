package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0(p0 uint8, p1 int16, p2 int8) int8 {
    v5 := GenS1{F0: "\n\t", F1: 0, F2: 2}
    switch "" {
    case "\n\t":
        switch 127 {
        case 127:
            p2 = (p2 + (p2 % ((p2) | 1)))
        default:
            fmt.Println(p1)
        }
    default:
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int16 = (p1 % ((int16(fi6)) | 1))
            var v8 uint64 = 713
            for fi9 := 0; fi9 < 19; fi9++ {
                switch 32768 {
                case 42:
                    v10 := []int{fi6, fi6}
                    _ = v10
                case 32768:
                    var v11 int64 = 506
                    _ = v11
                default:
                    var v12 uint32 = 4294967295
                    _ = v12
                }
                _ = fi9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        switch "foo bar" {
        case "":
            fmt.Println(fi13)
        }
        p0 = p0
        _ = fi13
    }
    switch ("a" + ((((("foo bar" + "foo bar") + "hello") + "\n\t") + "") + ("\n\t" + ""))) {
    case "":
        p0 -= 29
    case "a":
        var v14 int = 127
        _ = v14
    case "foo bar":
        p1 += (p1 % ((p1) | 1))
    default:
        if ((439 < 300) || (uint32(p0) != 65535)) {
            v15 := []int{9223372036854775807, -275, 0}
            if false {
                p0 = p0
                p0 -= p0
            } else {
                if (-39 != (p2 >> ((22) & 63))) {
                    if true {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                    var v16 int16 = p1
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if false {
                            var v18 uint32 = 2
                            _ = v18
                        }
                        p1 += 42
                        fmt.Println(p2)
                        _ = fi17
                    }
                }
                var v19 float64 = (3.14159 + float64(1e10))
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 int16 = (p1 ^ p1)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint = 854
                    switch "\n\t" {
                    case "":
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    case "hello":
                        switch fi22 {
                        case -480:
                            switch "a" {
                            case "":
                                if ((p0 * (p0 - uint8(v24))) < (p0 << ((v24) & 63))) {
                                    if true {
                                        var v25 int16 = p1
                                        p0 = (^(p0))
                                        _ = v25
                                    }
                                }
                            case "foo bar":
                                switch 42 {
                                case 4294967295:
                                    p0 += (^(p0))
                                case 127:
                                    if len(v15) > 0 {
                                        v15[0] = (-(fi22))
                                    }
                                case -334:
                                    _ = len(v15)
                                default:
                                    var v26 string = ("a" + ((("foo bar" + "") + "hello") + "\n\t"))
                                    _ = v26
                                }
                            case "a":
                                p1 = 1000
                            default:
                                var v27 int = -1
                                _ = v27
                            }
                        }
                    case "\n\t":
                        v15 = append(v15, fi22)
                        v28 := GenS1{F0: "", F1: 124, F2: 4294967295}
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
                var v29 int64 = -93
                var v30 uint = 433
                _ = fi22
                _ = v29
                _ = v30
            }
            _ = v15
        } else {
            for fi31 := 0; fi31 < 14; fi31++ {
                p2 -= p2
                _ = fi31
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                var v33 string = "hello"
                if false {
                    if (670 == uint32(p0)) {
                        v34 := GenS2{F0: 203, F1: (p0 | (^(p0))), F2: 255}
                        _ = v34
                    }
                }
                _ = fi32
                _ = v33
            }
            v35 := []int{-1, 42}
            _ = v35
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}
func genFn1(p0 uint64, p1 int8, p2 uint64) uint {
    fmt.Println(p1, p2, p0)
    switch ("a" + ("foo bar" + "foo bar")) {
    case "hello":
        v36 := GenS2{F0: 255, F1: 0, F2: 0}
        _ = v36
    case "":
        var v37 int32 = 656
        _ = v37
    case "\n\t":
        p1 += (p1 | genFn0(235, -1, (p1 & 100)))
    default:
        p1 = genFn0(116, -108, (p1 << ((711) & 63)))
    }
    v38 := []int{42, 0, 184}
    v39 := GenS1{F0: ("foo bar" + ("hello" + "")), F1: 271, F2: (p0 + 885)}
    p1 += (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v39
    return 265
}

func main() {
    var v1 int32 = 1
    var v2 int = 127
    var v3 int64 = 42
    var v4 int64 = (v3 >> ((18446744073709551615) & 63))
    switch "hello" {
    case "a":
        var v40 string = "hello"
        _ = v40
    case "":
        var v41 uint16 = 0
        _ = v41
    case "hello":
        v2 -= v2
    default:
        v42 := GenS2{F0: 128, F1: 128, F2: 65535}
        _ = v42
    }
    switch "foo bar" {
    case "a":
        v43 := []int{(v2 ^ v2)}
        _ = v43
    case "foo bar":
        fmt.Println(v3, v3)
    case "":
        var v44 uint32 = 2147483648
        _ = v44
    }
    var v45 uint64 = 0
    var v46 int8 = genFn0(1, -1, 42)
    v47 := []int{}
    v1 -= int32(v2)
    for fi48 := 0; fi48 < 14; fi48++ {
        switch "foo bar" {
        case "":
            if true {
                for fi49 := 0; fi49 < 14; fi49++ {
                    v47 = append(v47, (fi48 / ((483) | 1)))
                    var v50 bool = (("a" + "\n\t") > ("" + "\n\t"))
                    _ = fi49
                    _ = v50
                }
                var v51 int8 = v46
                _ = v51
            } else {
                var v52 int = v2
                v3 -= v4
                var v53 uint = 1
                _ = v52
                _ = v53
            }
        case "a":
            var v54 uint64 = 42
            _ = v54
        case "\n\t":
            var v55 int = (fi48 | 1)
            _ = v55
        default:
            v4 = (v3 >> ((uint(v45)) & 63))
        }
        v56 := GenS3{F0: (true || true), F1: 65535, F2: 65535}
        var v57 int16 = 42
        _ = fi48
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 12; fi58++ {
        var v59 uint8 = 172
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
