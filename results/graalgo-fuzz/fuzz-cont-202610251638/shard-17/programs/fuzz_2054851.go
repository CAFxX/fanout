package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int32; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }
type GenS7 struct { F0 string; F1 bool; F2 int8 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 bool; F1 bool; F2 uint }
type GenS10 struct { F0 int; F1 int16; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 uint32 }
type GenS13 struct { F0 bool; F1 int32; F2 uint32 }
func genFn0(p0 int) uint8 {
    v4 := GenS1{F0: 38}
    v5 := GenS1{F0: 127}
    if false {
        v6 := GenS2{F0: "foo bar", F1: p0, F2: 2}
        _ = v6
    } else {
        var v7 int32 = 1
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = -96
            p0 += (^(fi8))
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    p0 = (p0 << ((255) & 63))
    _ = p0
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 int16 = 76
    v2 := make(map[uint32]string)
    var v3 string = ("a" + "foo bar")
    if (!false) {
        if (!false) {
            if false {
                v10 := GenS2{F0: v3, F1: 128, F2: 443}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v2[992] = "hello"
                    _ = fi11
                }
                switch (v3 + "") {
                case "hello":
                    v12 := []int{-264}
                    _ = v12
                case "a":
                    v13 := GenS2{F0: v3, F1: 32767, F2: 726}
                    _ = v13
                }
                if (v3 > v3) {
                    v14 := GenS1{F0: 100}
                    _ = v14
                }
            }
            v3 = "foo bar"
            var v15 string = "foo bar"
            _ = v15
        } else {
            v16 := []int{2147483647, 408, -2147483648, 138}
            v1 += 722
            var v17 float32 = ((-2.25 / (((float32(1.5) - -2.25) / 1.0) / 1.5)) + 1.5)
            _ = v16
            _ = v17
        }
        delete(v2, 717)
    }
    v1 = v1
    v1 -= v1
    if (false || true) {
        v18 := GenS2{F0: (v3 + "a"), F1: -465, F2: 42}
        v1 = (v1 & 266)
        var v19 int64 = 368
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                if true {
                    if (false || (!(!(1e10 > 0.1)))) {
                        v22 := GenS3{F0: 37}
                        v3 = (v3 + "a")
                        switch (fi20 | (^(fi20))) {
                        case 2147483647:
                            if (!true) {
                                if false {
                                    v1 += v1
                                    v23 := GenS4{F0: v1, F1: -1, F2: 42}
                                    var v24 uint64 = 100
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 uint = 1
                                    if true {
                                        if (true || false) {
                                            v26 := []int{(^(fi21)), (fi21 * (^(fi21))), (^(fi20))}
                                            v27 := GenS5{F0: 539}
                                            _ = len(v2)
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v28 := GenS6{F0: -250, F1: 2, F2: v1}
                                            var v29 uint16 = 65535
                                            v30 := GenS6{F0: int64(v1), F1: 65535, F2: 1000}
                                            v2[65] = ""
                                            if true {
                                                v3 = v3
                                                var v31 uint = 18446744073709551615
                                                if ((v31 != 65535) && true) {
                                                    var v32 int8 = 96
                                                    if true {
                                                        if false {
                                                            v33 := []int{}
                                                            v29 = 2
                                                            _ = v33
                                                        }
                                                        _ = len(v2)
                                                    } else {
                                                        if ((v1 >> ((v31) & 63)) != v1) {
                                                            var v34 int32 = -16
                                                            v32 += -128
                                                            if true {
                                                                if true {
                                                                    v35 := GenS6{F0: 42, F1: (v31 / (((v31 << (((v31 + (v25 % (((v31 - (v25 | 465))) | 1)))) & 63))) | 1)), F2: 42}
                                                                    _ = v35
                                                                } else {
                                                                    _ = len(v2)
                                                                }
                                                                v25 = (^(v25))
                                                            } else {
                                                                if false {
                                                                    v31 = (v31 % ((65535) | 1))
                                                                } else {
                                                                    if ((((v29 < (v29 & (v29 ^ (^(v29))))) || false) && (!(true || (uint32(v31) == 2)))) || (int(v25) == 1)) {
                                                                        if true {
                                                                            v25 = v31
                                                                        }
                                                                        delete(v2, 2147483648)
                                                                        _ = len(v2)
                                                                    }
                                                                    switch fi20 {
                                                                    case 127:
                                                                        var v36 int16 = (v1 ^ v1)
                                                                        _ = v36
                                                                    case 32767:
                                                                        v31 = uint(v29)
                                                                    default:
                                                                        v37 := []int{(fi21 | (fi21 ^ int(v31))), fi20, (fi21 & (fi21 << ((198) & 63))), fi21}
                                                                        _ = v37
                                                                    }
                                                                }
                                                                var v38 int64 = int64(v32)
                                                                _ = v38
                                                            }
                                                            _ = v34
                                                        } else {
                                                            v39 := []int{fi20, (fi20 - fi21)}
                                                            switch v3 {
                                                            case "a":
                                                                var v40 uint32 = 592
                                                                _ = v40
                                                            }
                                                            var v41 uint64 = 18446744073709551615
                                                            _ = v39
                                                            _ = v41
                                                        }
                                                        v42 := GenS7{F0: (("" + v3) + (((v3 + v3) + "") + v3)), F1: false, F2: (v32 / ((v32) | 1))}
                                                        _ = v42
                                                    }
                                                    switch v3 {
                                                    case "":
                                                        var v43 uint32 = 0
                                                        _ = v43
                                                    default:
                                                        switch "a" {
                                                        case "a":
                                                            if true {
                                                                if (-2147483648 < 206) {
                                                                    var v44 uint16 = (v29 << (((v25 * v25)) & 63))
                                                                    v45 := GenS1{F0: -114}
                                                                    _ = v44
                                                                    _ = v45
                                                                } else {
                                                                    fmt.Println(v3, fi21, v1)
                                                                }
                                                                if true {
                                                                    v46 := []int{fi20, -1, 9223372036854775807, (fi20 & fi20)}
                                                                    _ = v46
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        delete(v2, 620)
                                                                    }
                                                                }
                                                                v47 := GenS8{F0: v3, F1: (v32 - 127)}
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                                _ = v31
                                            } else {
                                                var v48 uint8 = 255
                                                _ = v48
                                            }
                                            if false {
                                                var v49 int = fi21
                                                _ = v49
                                            } else {
                                                if (!(1 == 1)) {
                                                    fmt.Println(v3, v3)
                                                }
                                                v50 := GenS6{F0: 42, F1: (v25 & 255), F2: (v1 + 32767)}
                                                _ = v50
                                            }
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        v51 := GenS9{F0: false, F1: false, F2: v25}
                                        _ = v51
                                    } else {
                                        v52 := GenS10{F0: fi20, F1: (v1 * (v1 % ((1000) | 1))), F2: 127}
                                        var v53 float32 = 3.14159
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v25
                                }
                                var v54 int16 = -8
                                fmt.Println(v3, fi20, fi20)
                                v54 += v54
                                _ = v54
                            }
                        case -2147483648:
                            v1 -= 0
                        case 0:
                            v55 := v2[426]
                            _ = v55
                        }
                        _ = v22
                    }
                    if true {
                        switch fi21 {
                        case -1:
                            v3 = ((("\n\t" + v3) + v3) + "a")
                        }
                    } else {
                        fmt.Println(fi21)
                        switch (v3 + "foo bar") {
                        case "a":
                            fmt.Println(v3, v3, fi20)
                        case "foo bar":
                            switch fi20 {
                            case 430:
                                switch (fi21 + 32767) {
                                case -1:
                                    v56 := []int{(fi20 << ((4294967295) & 63)), 128, fi21}
                                    _ = v56
                                case 4294967295:
                                    if (true && (!false)) {
                                        v57 := []int{}
                                        v57 = append(v57, -3)
                                        v2[42] = ("" + "foo bar")
                                        _ = v57
                                    }
                                case 0:
                                    fmt.Println(v3)
                                }
                            case 2147483647:
                                v58 := GenS3{F0: 46}
                                _ = v58
                            default:
                                if (true && ((47 != 100) || false)) {
                                    var v59 float64 = 0.0
                                    _ = v59
                                }
                            }
                        }
                    }
                } else {
                    v1 = (v1 - (v1 | (v1 * (v1 - (v1 >> ((235) & 63))))))
                }
                var v60 int64 = 9223372036854775807
                switch v3 {
                case "hello":
                    switch v3 {
                    case "\n\t":
                        v61 := []int{127, fi20}
                        _ = v61
                    case "foo bar":
                        v62 := []int{0}
                        _ = v62
                    case "hello":
                        v3 = ((v3 + v3) + v3)
                    }
                case "a":
                    delete(v2, 2147483648)
                case "foo bar":
                    v63 := v2[1]
                    _ = v63
                }
                _ = fi21
                _ = v60
            }
            v3 = ""
            _ = fi20
        }
        v64 := []int{-9223372036854775808}
        v1 -= (-(v1))
        _ = v64
    }
    v65 := []int{188, -9223372036854775808, -2147483648, 32768}
    for fi66 := 0; fi66 < 19; fi66++ {
        if false {
            var v67 uint64 = 705
            var v68 int = 4294967295
            _ = v67
            _ = v68
        } else {
            v69 := GenS9{F0: (!true), F1: (false && false), F2: 127}
            _ = v69
        }
        var v70 int8 = -128
        switch 636 {
        case 128:
            for fi71 := 0; fi71 < 7; fi71++ {
                v70 -= (v70 ^ (v70 << ((0) & 63)))
                v1 -= v1
                switch (fi66 % (((fi71 | 32768)) | 1)) {
                case -9223372036854775808:
                    v72 := GenS2{F0: (v3 + (v3 + v3)), F1: -2147483648, F2: 579}
                    _ = v72
                case -2147483648:
                    var v73 string = ""
                    _ = v73
                default:
                    v70 = v70
                }
                _ = fi71
            }
        case 669:
            v74 := GenS11{F0: 42}
            _ = v74
        }
        _ = fi66
        _ = v70
    }
    var v75 int8 = -96
    v75 -= -100
    if ((false || ((!false) || false)) && (1 != 65535)) {
        switch v3 {
        case "":
            for fi76 := 0; fi76 < 12; fi76++ {
                v3 = v3
                _ = fi76
            }
        case "a":
            switch 14 {
            case 32767:
                v1 = v1
            }
        default:
            var v77 uint64 = 4294967295
            _ = v77
        }
        v65 = append(v65, 32768)
        var v78 int = 9223372036854775807
        _ = v78
    }
    var v79 uint8 = 127
    var v80 int8 = (v75 / ((42) | 1))
    for fi81 := 0; fi81 < 13; fi81++ {
        for fi82 := 0; fi82 < 3; fi82++ {
            v65 = append(v65, fi81)
            if false {
                v83 := GenS11{F0: 527}
                v65 = append(v65, -71)
                v79 = (v79 >> ((42) & 63))
                switch 95 {
                case -150:
                    v84 := GenS6{F0: 0, F1: 4294967295, F2: 0}
                    if len(v65) >= 2 {
                        v65 = v65[1:len(v65)]
                    }
                    _ = v84
                default:
                    if false {
                        fmt.Println(v75, v75, v79)
                    } else {
                        v2[65535] = v3
                        var v85 int16 = v1
                        v1 = (v1 + 32767)
                        _ = v85
                    }
                }
                v3 = "foo bar"
                _ = v83
            }
            if len(v65) > 0 {
                v65[0] = 131
            }
            _ = fi82
        }
        _ = fi81
    }
    var v86 int64 = 0
    v75 = 1
    if len(v65) >= 2 {
        v65 = v65[1:len(v65)]
    }
    if false {
        v87 := GenS5{F0: 2147483648}
        var v88 uint16 = 1
        _ = v87
        _ = v88
    } else {
        if (!(!(292 == 767))) {
            switch "" {
            case "hello":
                var v89 int = 397
                _ = v89
            default:
                var v90 uint = 596
                _ = v90
            }
            var v91 uint8 = (v79 / ((v79) | 1))
            _ = v91
        } else {
            v92 := GenS12{F0: v75, F1: 0}
            v3 = "foo bar"
            v75 -= (v80 >> ((18446744073709551615) & 63))
            _ = v92
        }
    }
    v3 = ((("\n\t" + (v3 + (("hello" + "a") + "foo bar"))) + (v3 + v3)) + v3)
    v93 := GenS13{F0: true, F1: 1, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v75:", v75)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v75
    _ = v79
    _ = v80
    _ = v86
    _ = v93
}
