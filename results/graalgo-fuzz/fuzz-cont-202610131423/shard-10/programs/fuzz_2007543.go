package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64; F1 bool; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32 }
func genFn0(p0 bool, p1 int64) uint32 {
    v5 := GenS1{F0: 840, F1: (!(1 != 255))}
    p0 = (-1 >= -2147483648)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (2147483648 <= 131) {
            var v7 int32 = -2147483648
            _ = v7
        }
        if p0 {
            if true {
                p1 = (p1 * (p1 | (p1 << ((0) & 63))))
            } else {
                v8 := []int{}
                _ = v8
            }
            v9 := GenS2{F0: fi6, F1: 1, F2: ("hello" + "")}
            _ = v9
        }
        _ = fi6
    }
    v10 := []int{0, 127}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 2
}
func genFn1(p0 int32, p1 uint32) string {
    v11 := GenS1{F0: 1, F1: (0.0 == -2.25)}
    p0 += p0
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: 5, F1: true}
        v14 := []int{-663, (^(fi12))}
        p0 -= p0
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                var v17 int64 = 9223372036854775807
                _ = v17
            } else {
                fmt.Println(fi15)
                v18 := GenS3{F0: -45, F1: 0}
                _ = v18
            }
            p0 = (^(p0))
            switch (fi16 & (fi16 / (((fi16 >> ((201) & 63))) | 1))) {
            case 9223372036854775807:
                if true {
                    fmt.Println(p1, fi15)
                    fmt.Println(p1, p0, p0)
                    p1 = 65535
                    v19 := []int{fi16, (fi16 | (^(fi16)))}
                    _ = v19
                } else {
                    v20 := []int{(fi15 % (((fi15 | fi15)) | 1))}
                    _ = v20
                }
            case 943:
                p1 += p1
            case -1:
                v21 := []int{(fi15 ^ 32768), fi15}
                _ = v21
            default:
                p1 += (^(p1))
            }
            _ = fi16
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v11
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := []int{-9223372036854775808, 244}
    v4 := make(map[uint32]bool)
    v22 := GenS3{F0: -128, F1: 42}
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            v23 -= v23
            _ = fi25
        }
        if (false || false) {
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 float32 = 1.5
                switch "foo bar" {
                case "":
                    var v28 uint8 = 1
                    var v29 int64 = 42
                    _ = v28
                    _ = v29
                case "a":
                    if true {
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v30 := GenS4{F0: 42, F1: 83}
                            v31 := GenS2{F0: (fi26 ^ (-(fi26))), F1: (fi26 & (fi26 + 507)), F2: "hello"}
                            var v32 int32 = 2147483647
                            v33 := GenS4{F0: 0, F1: 100}
                            _ = v30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                default:
                    v27 = v27
                }
                var v34 int8 = -1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "" {
                case "\n\t":
                    var v35 int = fi26
                    _ = v35
                case "hello":
                    v23 -= v23
                }
                _ = fi26
                _ = v27
                _ = v34
            }
            fmt.Println(fi24, v23)
        } else {
            var v36 int = 868
            v37 := GenS5{F0: v23, F1: ((1 <= 42) && false), F2: false}
            switch fi24 {
            case 32767:
                for fi38 := 0; fi38 < 6; fi38++ {
                    v3 = append(v3, fi24)
                    v39 := GenS6{F0: 42}
                    var v40 bool = true
                    if (v40 || (v40 && (!(0.0 <= 1.5)))) {
                        var v41 int16 = -1
                        var v42 int = fi24
                        _ = v41
                        _ = v42
                    }
                    var v43 uint64 = (v23 << ((42) & 63))
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v43
                }
            }
            _ = v36
            _ = v37
        }
        _ = fi24
    }
    _ = len(v1)
    fmt.Println(v23)
    if true {
        if true {
            switch 42 {
            case 4294967295:
                v23 = (v23 - (v23 >> ((69) & 63)))
            case 32768:
                v23 = (v23 & (v23 - (v23 ^ v23)))
            }
            var v44 int64 = 0
            for fi45 := 0; fi45 < 13; fi45++ {
                v44 = (-(v44))
                _ = fi45
            }
            _ = v44
        } else {
            v23 = 963
        }
        if false {
            v23 -= v23
            if (-2147483648 != -517) {
                v46 := GenS3{F0: 100, F1: 1}
                var v47 string = ("hello" + ("\n\t" + "\n\t"))
                fmt.Println(v23, v23)
                _ = v46
                _ = v47
            }
            fmt.Println(v23)
        }
    } else {
        for fi48 := 0; fi48 < 9; fi48++ {
            fmt.Println(v23, fi48)
            _ = fi48
        }
        for fi49 := 0; fi49 < 18; fi49++ {
            v50 := v2[fi49]
            v51 := GenS7{F0: 166}
            v50 += v50
            _ = fi49
            _ = v50
            _ = v51
        }
        switch 128 {
        case 11:
            var v52 int = -1
            _ = v52
        default:
            var v53 int64 = -9223372036854775808
            _ = v53
        }
    }
    v54 := GenS8{F0: 65535}
    switch 9223372036854775807 {
    case 0:
        switch -2147483648 {
        case 2147483647:
            v23 += v23
        }
    default:
        switch -2147483648 {
        case 128:
            for fi55 := 0; fi55 < 14; fi55++ {
                v23 = v23
                var v56 int8 = -128
                _ = fi55
                _ = v56
            }
        case 806:
            for fi57 := 0; fi57 < 13; fi57++ {
                for fi58 := 0; fi58 < 2; fi58++ {
                    v23 = (v23 >> ((4294967295) & 63))
                    v23 = (v23 >> ((827) & 63))
                    _ = fi58
                }
                v23 += (v23 / (((^(v23))) | 1))
                _ = fi57
            }
        case -238:
            for fi59 := 0; fi59 < 19; fi59++ {
                v23 = (v23 | v23)
                _ = fi59
            }
        default:
            var v60 uint32 = 2147483647
            _ = v60
        }
    }
    switch 2147483647 {
    case 449:
        if false {
            for fi61 := 0; fi61 < 8; fi61++ {
                fmt.Println(fi61, fi61)
                _ = fi61
            }
        }
    case -1:
        for fi62 := 0; fi62 < 3; fi62++ {
            var v63 uint16 = 32767
            v64 := GenS9{F0: 2147483648}
            switch fi62 {
            case -401:
                if false {
                    fmt.Println(v63, v63, v23)
                } else {
                    v63 += 363
                    delete(v4, 4294967295)
                }
            case -949:
                switch (fi62 * fi62) {
                case 473:
                    v23 = 42
                default:
                    v63 -= (^(v63))
                }
            case 1:
                v4[0] = (2147483647 != 65535)
            default:
                if (41 != 127) {
                    v63 = (v63 << ((255) & 63))
                    var v65 float32 = -2.25
                    fmt.Println(v63)
                    _ = v65
                } else {
                    fmt.Println(v23)
                }
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    var v66 int8 = 127
    v67 := GenS1{F0: uint16(v23), F1: false}
    v66 = v66
    var v68 int32 = -2147483648
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v69 uint16 = 24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v54
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
