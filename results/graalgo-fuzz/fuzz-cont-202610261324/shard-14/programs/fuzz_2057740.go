package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 string; F2 uint }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint16; F2 string }
func genFn0() int16 {
    var v7 int16 = 57
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int64 = 0
        v9 = 29
        switch "a" {
        case "\n\t":
            v10 := []int{(fi8 << ((2) & 63))}
            _ = v10
        case "foo bar":
            var v11 float64 = 0.1
            _ = v11
        case "":
            switch (fi8 ^ (fi8 ^ 32767)) {
            case 4294967295:
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch 2147483647 {
                    case 42:
                        v9 -= (v9 * (^(v9)))
                    }
                    v13 := []int{fi8, fi12, fi12}
                    _ = fi12
                    _ = v13
                }
            default:
                var v14 float32 = 0.0
                _ = v14
            }
        }
        _ = fi8
        _ = v9
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint64 = 121
            _ = fi16
            _ = v17
        }
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 3; fi18++ {
                fmt.Println(fi15, fi18)
                v19 := GenS1{F0: -1, F1: 9223372036854775807, F2: false}
                _ = fi18
                _ = v19
            }
        }
        if (!(false || (!false))) {
            v7 -= 42
        } else {
            var v20 int16 = (v7 ^ 42)
            v21 := []int{}
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    v22 := []int{}
    if true {
        for fi23 := 0; fi23 < 18; fi23++ {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            v24 := GenS2{F0: fi23, F1: 0, F2: 9223372036854775807}
            _ = fi23
            _ = v24
        }
        v7 = v7
    }
    _ = v22
    _ = v7
    return 1
}
func genFn1(p0 uint32) int64 {
    switch 32767 {
    case -1:
        p0 += 1
    default:
        p0 = 126
    }
    v25 := []int{397, 127}
    if false {
        for fi26 := 0; fi26 < 1; fi26++ {
            if false {
                switch ("a" + "hello") {
                case "hello":
                    switch (fi26 * (-(fi26))) {
                    case -1:
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    }
                default:
                    switch (fi26 - 0) {
                    case -1:
                        for fi27 := 0; fi27 < 10; fi27++ {
                            if len(v25) > 0 {
                                v25[0] = 771
                            }
                            _ = fi27
                        }
                    default:
                        switch (fi26 & fi26) {
                        case 102:
                            v25 = append(v25, -1)
                        case -9223372036854775808:
                            v25 = append(v25, 4294967295)
                        case 32767:
                            _ = len(v25)
                        default:
                            var v28 int = fi26
                            _ = v28
                        }
                    }
                }
                var v29 string = (("hello" + ("a" + "a")) + ("foo bar" + "a"))
                _ = v29
            } else {
                p0 = (^(p0))
                v30 := GenS1{F0: 42, F1: 42, F2: false}
                _ = v30
            }
            fmt.Println(fi26)
            _ = fi26
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(fi31, p0)
            _ = fi31
        }
    } else {
        var v32 int64 = 790
        _ = v32
    }
    var v33 uint16 = 350
    _ = p0
    _ = v25
    _ = v33
    return -1
}

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, (v1 & (v1 % ((v1) | 1))), (v1 & 42), v1}
    var v3 int64 = 0
    v4 := []uint32{(v1 % (((v1 - v1)) | 1)), (v1 / (((v1 >> ((2) & 63))) | 1)), v1}
    v5 := []int{42}
    v6 := []int{53, 128}
    v34 := GenS3{F0: 0}
    v1 = 116
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 bool = true
        v3 -= (v3 ^ (-(v3)))
        if v36 {
            if ((1e10 / 3.14159) != float64(1e10)) {
                v1 += v1
                v4 = append(v4, 2)
                for fi37 := 0; fi37 < 20; fi37++ {
                    v1 -= v1
                    fmt.Println(v36, v1)
                    _ = fi37
                }
            } else {
                var v38 int8 = -1
                fmt.Println(v36, v1)
                v1 = v1
                _ = v38
            }
            fmt.Println(v36, v1)
            var v39 float64 = 1.5
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 float64 = (3.14159 + 0.1)
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        var v43 uint16 = 255
        for fi44 := 0; fi44 < 10; fi44++ {
            fmt.Println(fi44, fi42, fi44)
            var v45 int8 = 42
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    fmt.Println(v3)
    var v46 int16 = -427
    fmt.Println(v3)
    for fi47 := 0; fi47 < 9; fi47++ {
        if false {
            var v48 int16 = -32768
            _ = v48
        } else {
            v49 := GenS2{F0: 9223372036854775807, F1: v3, F2: 9223372036854775807}
            if (false || true) {
                var v50 string = "\n\t"
                if true {
                    v51 := GenS4{F0: -128, F1: 0, F2: 42}
                    v46 += (-(v46))
                    _ = v51
                } else {
                    fmt.Println(v1)
                    for fi52 := 0; fi52 < 11; fi52++ {
                        v1 += v1
                        _ = len(v5)
                        if true {
                            var v53 uint16 = 1
                            v54 := GenS1{F0: v3, F1: genFn1(v1), F2: true}
                            if ((v53 % (((^(v53))) | 1)) <= 865) {
                                v46 = v46
                                if false {
                                    _ = len(v4)
                                    v50 = v50
                                    v46 -= 0
                                }
                                var v55 uint8 = 238
                                _ = v55
                            } else {
                                switch fi47 {
                                case 2147483647:
                                    v46 -= int16(v53)
                                case 316:
                                    v50 = (v50 + v50)
                                case -2147483648:
                                    v6 = append(v6, -1)
                                default:
                                    var v56 string = ((v50 + ("hello" + v50)) + ("foo bar" + v50))
                                    _ = v56
                                }
                                if (!(v46 == v46)) {
                                    if (334 == 126) {
                                        var v57 uint8 = 255
                                        _ = v57
                                    }
                                }
                            }
                            switch -434 {
                            case 483:
                                switch "" {
                                case "":
                                    v58 := GenS5{F0: 128, F1: "\n\t", F2: 2}
                                    _ = len(v2)
                                    v59 := GenS4{F0: -1, F1: 18446744073709551615, F2: v53}
                                    _ = v58
                                    _ = v59
                                case "a":
                                    v46 += genFn0()
                                case "\n\t":
                                    var v60 int32 = -2147483648
                                    _ = v60
                                default:
                                    fmt.Println(fi47, v1)
                                }
                            case -1:
                                var v61 int8 = -128
                                _ = v61
                            case -2147483648:
                                v1 += v1
                            default:
                                var v62 int8 = -1
                                _ = v62
                            }
                            v50 = v50
                            _ = v53
                            _ = v54
                        }
                        _ = fi52
                    }
                }
                _ = v50
            } else {
                v63 := GenS6{F0: "", F1: -224}
                var v64 int16 = genFn0()
                v46 = v46
                _ = v63
                _ = v64
            }
            var v65 int = (fi47 / ((fi47) | 1))
            _ = v49
            _ = v65
        }
        if len(v5) > 0 {
            v5[0] = fi47
        }
        _ = fi47
    }
    var v66 string = "foo bar"
    if ((v46 + v46) > -1) {
        var v67 int8 = 16
        _ = v67
    }
    for fi68 := 0; fi68 < 11; fi68++ {
        switch ("\n\t" + v66) {
        case "a":
            v69 := GenS7{F0: 61, F1: 688, F2: v66}
            _ = v69
        default:
            for fi70 := 0; fi70 < 15; fi70++ {
                v3 = v3
                v71 := GenS3{F0: 2}
                _ = fi70
                _ = v71
            }
        }
        _ = fi68
    }
    if ((-56 == 1) && ((int64(v46) > (^(v3))) && false)) {
        v72 := GenS2{F0: 127, F1: v3, F2: v3}
        _ = v72
    }
    v6 = append(v6, 4294967295)
    for fi73 := 0; fi73 < 18; fi73++ {
        if false {
            for fi74 := 0; fi74 < 18; fi74++ {
                v66 = v66
                v66 = "foo bar"
                _ = fi74
            }
            v3 += -1
        } else {
            fmt.Println(v46)
            for fi75 := 0; fi75 < 15; fi75++ {
                var v76 string = v66
                v46 += (v46 << ((2) & 63))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi75
                _ = v76
            }
        }
        _ = fi73
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v66
}
