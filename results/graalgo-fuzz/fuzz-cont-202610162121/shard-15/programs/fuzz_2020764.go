package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint32; F2 int16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 string, p1 int8, p2 uint32) int16 {
    p1 = (p1 & p1)
    var v4 uint = 4294967295
    if true {
        var v5 string = (p0 + ("a" + p0))
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    var v3 uint32 = 42
    for fi6 := 0; fi6 < 7; fi6++ {
        v3 += 4294967295
        if false {
            if false {
                v2 = v2
                for fi7 := 0; fi7 < 5; fi7++ {
                    v3 = v3
                    _ = len(v1)
                    v2 += v2
                    _ = fi7
                }
            }
            v3 = 2
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint32 = v3
                v2 = (^(v2))
                _ = fi8
                _ = v9
            }
            v10 := []int{fi6}
            _ = v10
        }
        v2 += 1
        _ = fi6
    }
    fmt.Println(v3, v3)
    var v11 uint32 = 0
    fmt.Println(v3)
    if (1e10 == 1e10) {
        v12 := GenS1{F0: 440, F1: -1, F2: 88}
        switch "foo bar" {
        case "foo bar":
            v3 = (v3 | 115)
        case "hello":
            var v13 int64 = v2
            _ = v13
        default:
            v14 := GenS2{F0: 255}
            _ = v14
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float32 = 0.1
        v17 := v1[fi15]
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        var v18 int = -158
        if (!((0.0 / 0.1) <= 3.14159)) {
            v19 := GenS1{F0: 685, F1: 284, F2: 127}
            v2 = (v2 ^ (-(v2)))
            delete(v1, (v18 - v18))
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            v11 = 313
            for fi21 := 0; fi21 < 2; fi21++ {
                _ = len(v1)
                v11 -= 2147483648
                var v22 float32 = float32(1e10)
                delete(v1, -322)
                fmt.Println(fi20, v22, v18)
                var v23 int = (^(fi21))
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int32 = 1
            _ = fi20
            _ = v24
        }
        _ = v18
    } else {
        v11 += v3
    }
    v25 := GenS2{F0: 32767}
    switch "\n\t" {
    case "":
        switch ("hello" + "foo bar") {
        case "":
            v26 := []int{-1, int(v2)}
            _ = v26
        }
    case "a":
        fmt.Println(v3)
    case "\n\t":
        if true {
            switch 4294967295 {
            case 127:
                if false {
                    v11 += (v11 % ((v11) | 1))
                    if (65535 != 2) {
                        var v27 uint64 = 4294967295
                        v28 := GenS2{F0: 0}
                        _ = v27
                        _ = v28
                    }
                    var v29 bool = false
                    _ = v29
                } else {
                    v3 += (^(v11))
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v2 -= (v2 % ((v2) | 1))
                        v1[127] = 128
                        v3 += v3
                        _ = fi30
                    }
                }
            default:
                if (0 >= 0) {
                    switch ("" + "hello") {
                    case "\n\t":
                        v31 := v1[0]
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 20; fi32++ {
                        fmt.Println(v11)
                        for fi33 := 0; fi33 < 5; fi33++ {
                            var v34 bool = false
                            var v35 int8 = 1
                            var v36 uint8 = 213
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = fi32
                    }
                } else {
                    switch ("" + "a") {
                    case "a":
                        var v37 int32 = -319
                        _ = v37
                    case "\n\t":
                        v38 := GenS2{F0: 797}
                        _ = v38
                    case "hello":
                        v39 := []int{}
                        _ = v39
                    }
                    var v40 uint32 = (^(v11))
                    switch "hello" {
                    case "\n\t":
                        if (false || (v3 < (v11 << ((4294967295) & 63)))) {
                            v41 := GenS3{F0: 211, F1: ("\n\t" + "foo bar")}
                            _ = v41
                        } else {
                            if false {
                                if true {
                                    v42 := GenS4{F0: 961}
                                    if (711 <= 0) {
                                        if true {
                                            v11 -= 2147483648
                                            v43 := []int{127, 965}
                                            v43 = append(v43, -2147483648)
                                            _ = v43
                                        }
                                        v44 := GenS1{F0: 4294967295, F1: (v2 ^ (v2 << ((876) & 63))), F2: 100}
                                        v40 -= (v11 % (((v40 - 428)) | 1))
                                        _ = v44
                                    }
                                    _ = v42
                                }
                                var v45 int16 = genFn0("foo bar", 42, 2)
                                _ = v45
                            }
                            v46 := GenS3{F0: 9223372036854775807, F1: ""}
                            if (!(true || false)) {
                                v40 -= (v40 ^ (v11 << ((2) & 63)))
                                var v47 uint32 = (v40 << ((0) & 63))
                                _ = v47
                            } else {
                                var v48 int = 2147483647
                                for fi49 := 0; fi49 < 19; fi49++ {
                                    fmt.Println(v2, v11, v40)
                                    v50 := []int{fi49, fi49}
                                    v11 += 42
                                    _ = fi49
                                    _ = v50
                                }
                                var v51 int = 32767
                                _ = v48
                                _ = v51
                            }
                            _ = v46
                        }
                    case "a":
                        for fi52 := 0; fi52 < 16; fi52++ {
                            fmt.Println(v40)
                            if true {
                                v40 -= v40
                                v11 += (v3 & 4294967295)
                            } else {
                                if true {
                                    var v53 float32 = -2.25
                                    _ = v53
                                }
                                v54 := []int{}
                                v2 += int64(fi52)
                                _ = v54
                            }
                            _ = fi52
                        }
                    default:
                        switch "hello" {
                        case "a":
                            for fi55 := 0; fi55 < 8; fi55++ {
                                fmt.Println(v40)
                                v56 := GenS1{F0: 172, F1: (v2 >> ((0) & 63)), F2: 127}
                                _ = fi55
                                _ = v56
                            }
                        case "\n\t":
                            if true {
                                var v57 int64 = (v2 % ((42) | 1))
                                var v58 uint = 4294967295
                                _ = v57
                                _ = v58
                            }
                        default:
                            var v59 int8 = 127
                            _ = v59
                        }
                    }
                    _ = v40
                }
            }
            switch 128 {
            case 4294967295:
                for fi60 := 0; fi60 < 13; fi60++ {
                    v61 := GenS3{F0: (fi60 ^ 42), F1: ("hello" + "\n\t")}
                    _ = fi60
                    _ = v61
                }
            case -2147483648:
                if false {
                    var v62 int = -2147483648
                    v2 = (v2 >> ((18446744073709551615) & 63))
                    v63 := GenS3{F0: (v62 & v62), F1: ("hello" + ("\n\t" + "a"))}
                    _ = v62
                    _ = v63
                } else {
                    v64 := GenS5{F0: "", F1: (v3 * v11), F2: -32768}
                    v65 := GenS3{F0: 2147483647, F1: ("a" + "foo bar")}
                    switch "" {
                    case "a":
                        var v66 int16 = 0
                        _ = v66
                    default:
                        v67 := GenS6{F0: v2}
                        _ = v67
                    }
                    _ = v64
                    _ = v65
                }
            case -665:
                for fi68 := 0; fi68 < 19; fi68++ {
                    v11 += v11
                    switch fi68 {
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "hello":
                            v3 = 2147483648
                        case "foo bar":
                            if ((false || false) || true) {
                                fmt.Println(v3)
                                switch 32768 {
                                case -752:
                                    fmt.Println(v3, v11, v3)
                                case 32768:
                                    var v69 uint32 = 2
                                    _ = v69
                                default:
                                    v2 = v2
                                }
                                v2 = v2
                            } else {
                                for fi70 := 0; fi70 < 7; fi70++ {
                                    v71 := GenS3{F0: (fi70 / (((fi68 / ((fi70) | 1))) | 1)), F1: (("hello" + "hello") + "\n\t")}
                                    _ = len(v1)
                                    v72 := []int{(fi68 / (((^(fi70))) | 1)), fi68, (fi70 | fi68)}
                                    if (((!(1e10 < 0.1)) && false) || (true || (1 != 65535))) {
                                        if (!(!(489 >= 18446744073709551615))) {
                                            var v73 uint16 = 678
                                            v3 = (v11 / ((1) | 1))
                                            delete(v1, fi68)
                                            _ = v73
                                        } else {
                                            var v74 string = "\n\t"
                                            switch "" {
                                            case "a":
                                                if (!(false || ((false || (65535 >= 65535)) || (18446744073709551615 >= 242)))) {
                                                    fmt.Println(fi70)
                                                }
                                                v75 := GenS3{F0: fi68, F1: v74}
                                                v3 = (v3 << ((18446744073709551615) & 63))
                                                _ = v75
                                            case "":
                                                v2 -= (v2 / (((v2 | 1)) | 1))
                                            }
                                            _ = v74
                                        }
                                        v1[(-(fi68))] = fi68
                                    } else {
                                        var v76 uint8 = 2
                                        v77 := v1[fi70]
                                        _ = v76
                                        _ = v77
                                    }
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                }
                            }
                        case "\n\t":
                            v11 = v3
                        default:
                            if false {
                                var v78 uint8 = 182
                                for fi79 := 0; fi79 < 14; fi79++ {
                                    var v80 uint8 = v78
                                    _ = fi79
                                    _ = v80
                                }
                                v2 = (v2 | v2)
                                _ = v78
                            } else {
                                delete(v1, fi68)
                                for fi81 := 0; fi81 < 7; fi81++ {
                                    v3 = 4294967295
                                    _ = fi81
                                }
                            }
                        }
                    }
                    v2 += 0
                    _ = fi68
                }
            default:
                if true {
                    fmt.Println(v3)
                    v82 := GenS2{F0: 183}
                    for fi83 := 0; fi83 < 10; fi83++ {
                        v11 = v11
                        v84 := v1[fi83]
                        for fi85 := 0; fi85 < 7; fi85++ {
                            var v86 int64 = v2
                            var v87 int = v84
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        _ = fi83
                        _ = v84
                    }
                    _ = v82
                } else {
                    v88 := GenS1{F0: 725, F1: v2, F2: 127}
                    _ = v88
                }
            }
            v89 := GenS4{F0: 1}
            _ = v89
        } else {
            var v90 uint = 232
            _ = v90
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
}
