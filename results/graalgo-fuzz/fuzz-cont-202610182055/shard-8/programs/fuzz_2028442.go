package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int32 }
func genFn0() int64 {
    if false {
        switch "foo bar" {
        case "":
            var v6 bool = (!(true || ((!(!(true || ((((4294967295 < 1) && false) || (!((!(!false)) && ((-2.25 * 3.14159) > 3.14159)))) && (-60 >= 127))))) || (false || true))))
            _ = v6
        case "foo bar":
            var v7 uint16 = 935
            _ = v7
        }
        var v8 int8 = 1
        var v9 uint = 301
        _ = v8
        _ = v9
    } else {
        v10 := []int{127}
        for fi11 := 0; fi11 < 9; fi11++ {
            if true {
                var v12 int64 = 662
                fmt.Println(fi11, v12)
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                _ = v12
            }
            switch -1 {
            case 317:
                if len(v10) > 0 {
                    v10[0] = (fi11 << ((1) & 63))
                }
            case 32767:
                v10 = append(v10, fi11)
            default:
                switch fi11 {
                case 2147483647:
                    v13 := GenS1{F0: 32767, F1: "\n\t", F2: 94}
                    _ = v13
                case 32768:
                    if (796 != 2147483647) {
                        if len(v10) > 0 {
                            v10[0] = (fi11 >> ((1) & 63))
                        }
                        var v14 uint64 = 528
                        v15 := GenS2{F0: 42, F1: 2147483647, F2: (!false)}
                        _ = v14
                        _ = v15
                    }
                }
            }
            _ = fi11
        }
        if false {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int = (fi16 & fi16)
                _ = fi16
                _ = v17
            }
            switch (("" + ("hello" + "foo bar")) + "a") {
            case "foo bar":
                if len(v10) > 0 {
                    v10[0] = -1
                }
            case "":
                v10 = append(v10, -893)
            case "a":
                var v18 float32 = float32(float64((float32(1.5) + -2.25)))
                _ = v18
            default:
                switch 32768 {
                case 98:
                    v19 := GenS2{F0: 230, F1: 2, F2: false}
                    _ = v19
                }
            }
            var v20 uint = 65535
            _ = v20
        }
        _ = v10
    }
    v21 := []int{1, 32768}
    _ = len(v21)
    switch 4294967295 {
    case 2147483647:
        if false {
            if false {
                if false {
                    switch "hello" {
                    case "foo bar":
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v21 = append(v21, fi22)
                            switch "\n\t" {
                            case "a":
                                v23 := GenS2{F0: 54, F1: 42, F2: false}
                                _ = v23
                            case "foo bar":
                                var v24 int8 = 118
                                _ = v24
                            default:
                                v25 := GenS2{F0: 183, F1: 0, F2: (!true)}
                                _ = v25
                            }
                            for fi26 := 0; fi26 < 7; fi26++ {
                                v27 := GenS2{F0: 1, F1: 0, F2: false}
                                fmt.Println(fi22)
                                var v28 int64 = -1
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            _ = fi22
                        }
                    case "a":
                        if false {
                            v21 = append(v21, 42)
                        } else {
                            if ((false && (2147483647 < -1)) && true) {
                                if len(v21) > 0 {
                                    v21[0] = 9223372036854775807
                                }
                            }
                        }
                    }
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(fi29)
                    _ = fi29
                }
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 uint = 43
                _ = fi30
                _ = v31
            }
        } else {
            if len(v21) > 0 {
                v21[0] = -9223372036854775808
            }
            var v32 float32 = 3.14159
            for fi33 := 0; fi33 < 12; fi33++ {
                v32 += float32((0.1 - 3.14159))
                var v34 int64 = 42
                _ = fi33
                _ = v34
            }
            _ = v32
        }
    case -428:
        if (!true) {
            var v35 int64 = 0
            _ = v35
        }
    default:
        for fi36 := 0; fi36 < 19; fi36++ {
            fmt.Println(fi36)
            var v37 bool = (false && true)
            v38 := GenS2{F0: 127, F1: 42, F2: (v37 && ((150 <= -1) && ("foo bar" < ("\n\t" + ""))))}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    _ = v21
    return 0
}
func genFn1(p0 uint8, p1 uint16) uint8 {
    var v39 int8 = -84
    v39 = (v39 | (v39 | (-(v39))))
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := []int{-418, -820}
        for fi42 := 0; fi42 < 17; fi42++ {
            v43 := GenS1{F0: 1, F1: "a", F2: (p1 ^ (p1 ^ (^(p1))))}
            _ = fi42
            _ = v43
        }
        if true {
            fmt.Println(p1)
            v44 := GenS2{F0: (p0 - (p0 ^ p0)), F1: 2147483647, F2: false}
            _ = v44
        }
        _ = fi40
        _ = v41
    }
    var v45 int8 = v39
    _ = p0
    _ = p1
    _ = v39
    _ = v45
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    var v3 int = 803
    v4 := make(map[int]string)
    var v5 uint16 = 65535
    fmt.Println(v2)
    var v46 uint16 = (v5 << ((426) & 63))
    v47 := v4[474]
    v47 = v47
    fmt.Println(v47)
    switch (^(v3)) {
    case -295:
        fmt.Println(v3, v3)
    }
    if true {
        switch v47 {
        case "a":
            if true {
                var v48 int8 = -128
                v49 := GenS2{F0: genFn1(113, v5), F1: v2, F2: (!true)}
                _ = v48
                _ = v49
            } else {
                v50 := []int{}
                var v51 int32 = 501
                _ = v50
                _ = v51
            }
        case "foo bar":
            var v52 uint = 2
            _ = v52
        case "hello":
            var v53 uint64 = 1
            _ = v53
        }
        v46 = v5
    }
    v54 := []int{}
    for fi55 := 0; fi55 < 5; fi55++ {
        switch fi55 {
        case -9223372036854775808:
            var v56 float32 = 0.0
            _ = v56
        case 396:
            v57 := GenS1{F0: 42, F1: "\n\t", F2: (v46 & (v5 ^ v5))}
            _ = v57
        case 127:
            fmt.Println(v3)
        default:
            var v58 uint = 65535
            _ = v58
        }
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 int32 = -2147483648
            var v61 uint64 = uint64(v5)
            var v62 uint32 = (v2 | (^(v2)))
            fmt.Println(v46)
            if false {
                var v63 uint8 = 205
                if true {
                    v46 += (v46 >> ((18446744073709551615) & 63))
                    v64 := GenS3{F0: (v47 + v47)}
                    switch v47 {
                    case "":
                        v65 := v4[v3]
                        _ = v65
                    }
                    _ = v64
                }
                _ = v63
            } else {
                if false {
                    switch (v47 + v47) {
                    case "\n\t":
                        v61 = v61
                        v61 -= uint64(v46)
                        var v66 uint16 = v46
                        _ = v66
                    }
                }
                if true {
                    switch v47 {
                    case "hello":
                        if len(v54) >= 2 {
                            v54 = v54[1:len(v54)]
                        }
                    default:
                        switch 128 {
                        case 620:
                            switch fi59 {
                            case -9223372036854775808:
                                var v67 string = ("\n\t" + v47)
                                _ = len(v4)
                                v5 -= (v46 >> ((0) & 63))
                                _ = v67
                            case 9223372036854775807:
                                var v68 uint = 65535
                                _ = v68
                            default:
                                _ = len(v54)
                            }
                        case 1:
                            v47 = ((v47 + ((("a" + v47) + ((v47 + v47) + v47)) + ("a" + "a"))) + v47)
                        }
                    }
                    if (-2.25 > (3.14159 / 3.14159)) {
                        fmt.Println(fi55)
                        if true {
                            var v69 bool = ((v2 - 2) != (v62 << ((18446744073709551615) & 63)))
                            _ = v69
                        }
                    }
                } else {
                    fmt.Println(v2, v60)
                    v70 := GenS1{F0: -1, F1: (v47 + "a"), F2: (^(v5))}
                    switch v3 {
                    case 32767:
                        v61 -= (v61 / (((^(v61))) | 1))
                    default:
                        var v71 uint64 = 0
                        _ = v71
                    }
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                    var v72 int16 = 0
                    if false {
                        var v73 int32 = v60
                        var v74 float32 = 1e10
                        var v75 uint32 = v62
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                    v72 = 42
                    _ = v70
                    _ = v72
                }
                switch (v3 * 32768) {
                case -677:
                    switch 128 {
                    case -961:
                        v47 = "a"
                    }
                case 32768:
                    if ((^(v46)) != (v46 | (v5 | v5))) {
                        var v76 uint32 = (v62 & (^(v2)))
                        var v77 uint64 = (v61 ^ (v61 & (v61 ^ (v61 << ((915) & 63)))))
                        _ = v76
                        _ = v77
                    } else {
                        var v78 int64 = 0
                        _ = v78
                    }
                }
            }
            switch (fi59 % ((int(v60)) | 1)) {
            case -9223372036854775808:
                if false {
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v79 int = (fi55 << ((2) & 63))
                            _ = v79
                        case "":
                            var v80 uint16 = 255
                            _ = v80
                        case "a":
                            v3 -= fi55
                        }
                    }
                    v81 := GenS4{F0: false, F1: 1}
                    var v82 float64 = ((1e10 / 3.14159) / -2.25)
                    _ = v81
                    _ = v82
                }
            case 32768:
                fmt.Println(v2, fi59, v3)
            default:
                fmt.Println(v2, v47)
            }
            _ = fi59
            _ = v60
            _ = v61
            _ = v62
        }
        _ = fi55
    }
    v83 := GenS5{F0: 1}
    if len(v54) > 0 {
        v54[0] = (v3 * v3)
    }
    for fi84 := 0; fi84 < 20; fi84++ {
        v46 = v5
        v85 := v4[0]
        if (false && (((v46 & v46) <= (^(v5))) && false)) {
            if ((v2 - v2) < v2) {
                for fi86 := 0; fi86 < 9; fi86++ {
                    var v87 int8 = 100
                    _ = fi86
                    _ = v87
                }
            } else {
                var v88 uint64 = 42
                var v89 int32 = 85
                _ = v88
                _ = v89
            }
        }
        _ = fi84
        _ = v85
    }
    delete(v4, 575)
    switch ((v47 + "") + v47) {
    case "hello":
        v47 = v47
    default:
        v90 := GenS2{F0: 20, F1: v2, F2: (-11 < 42)}
        _ = v90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v54
    _ = v83
}
