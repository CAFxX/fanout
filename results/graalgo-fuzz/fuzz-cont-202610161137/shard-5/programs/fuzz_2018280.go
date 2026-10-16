package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint32; F1 string; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 bool; F2 string }
type GenS9 struct { F0 int16; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64; F1 int; F2 uint }
func genFn0() string {
    var v4 uint = 255
    var v5 string = "a"
    _ = v4
    _ = v5
    return ("\n\t" + v5)
}
func genFn1(p0 uint16) int8 {
    var v6 uint = 381
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{65535, 0}
    v2 := make(map[int]string)
    var v3 uint16 = 32767
    fmt.Println(v3)
    switch 9223372036854775807 {
    case 5:
        var v7 int8 = 100
        _ = v7
    case 0:
        v3 -= (v3 >> ((18446744073709551615) & 63))
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int32 = -610
        _ = len(v2)
        v9 = (v9 % ((int32(fi8)) | 1))
        _ = fi8
        _ = v9
    }
    v3 -= (v3 << ((65535) & 63))
    if true {
        fmt.Println(v3)
        _ = len(v2)
        v10 := GenS1{F0: 42}
        _ = v10
    }
    var v11 int64 = -520
    v12 := GenS2{F0: 1}
    v11 -= (^(v11))
    var v13 int64 = -1
    switch 127 {
    case 9223372036854775807:
        if (((v3 > v3) || (!false)) && false) {
            for fi14 := 0; fi14 < 2; fi14++ {
                if (!false) {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 uint = uint(v3)
                        v11 -= v13
                        if false {
                            var v17 uint8 = 1
                            _ = len(v2)
                            _ = v17
                        } else {
                            var v18 float32 = 0.1
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = len(v2)
                    fmt.Println(v13)
                }
                switch "a" {
                case "foo bar":
                    if (0.0 <= (((0.0 - (3.14159 - 0.1)) - 1.5) / 0.1)) {
                        v3 = 2
                    }
                case "\n\t":
                    v19 := v2[0]
                    _ = v19
                }
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 int64 = 1
                    _ = fi20
                    _ = v21
                }
                _ = fi14
            }
            var v22 uint16 = (v3 * 1)
            _ = v22
        } else {
            var v23 uint = 144
            v24 := GenS3{F0: "\n\t", F1: (234 != 2147483648), F2: (^(v3))}
            _ = v23
            _ = v24
        }
    default:
        if true {
            if (true && (true && false)) {
                if true {
                    v25 := GenS4{F0: 42, F1: "\n\t", F2: 972}
                    v26 := GenS1{F0: 4294967295}
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 int = 32768
                if false {
                    if false {
                        switch (v27 << ((255) & 63)) {
                        case 1:
                            if (32767 == -667) {
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    v29 := GenS5{F0: "hello"}
                                    var v30 int64 = (^(v13))
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                if true {
                                    _ = len(v2)
                                } else {
                                    if (349 == v27) {
                                        v11 = (v13 + (v11 ^ -482))
                                        for fi31 := 0; fi31 < 20; fi31++ {
                                            if (163 < 128) {
                                                var v32 uint8 = 255
                                                v27 += fi31
                                                _ = v32
                                            } else {
                                                var v33 int = fi31
                                                fmt.Println(v13, v27)
                                                var v34 uint64 = 2
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = fi31
                                        }
                                        if (v27 < v27) {
                                            for fi35 := 0; fi35 < 15; fi35++ {
                                                var v36 uint8 = 255
                                                var v37 uint8 = 42
                                                _ = fi35
                                                _ = v36
                                                _ = v37
                                            }
                                            if false {
                                                v38 := v2[v27]
                                                _ = v38
                                            } else {
                                                for fi39 := 0; fi39 < 15; fi39++ {
                                                    for fi40 := 0; fi40 < 1; fi40++ {
                                                        if true {
                                                            v41 := GenS5{F0: ("hello" + "")}
                                                            v42 := GenS6{F0: -115, F1: 18446744073709551615}
                                                            v13 -= (v13 - v11)
                                                            if (18446744073709551615 < 880) {
                                                                v13 = 364
                                                                if (!(-66 >= -1)) {
                                                                    v3 += (^(v3))
                                                                } else {
                                                                    v11 -= v11
                                                                }
                                                                if true {
                                                                    fmt.Println(v11, v13)
                                                                    v27 = fi40
                                                                    v3 = (^(v3))
                                                                    v3 += (^(v3))
                                                                    var v43 float64 = 1e10
                                                                    _ = v43
                                                                }
                                                                v3 -= 404
                                                                v3 = 0
                                                                var v44 uint = 620
                                                                _ = v44
                                                            } else {
                                                                v45 := GenS6{F0: 100, F1: 457}
                                                                v3 = v3
                                                                _ = v45
                                                            }
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            fmt.Println(v27)
                                                            fmt.Println(fi39, fi39, v3)
                                                            v27 = (v27 | (fi39 << ((0) & 63)))
                                                            v46 := GenS5{F0: "hello"}
                                                            var v47 int64 = (^(v13))
                                                            fmt.Println(v27, fi39)
                                                            if ((v3 - 2) >= (v3 + (v3 | (^(v3))))) {
                                                                v11 = v47
                                                                var v48 uint = 350
                                                                _ = v48
                                                            } else {
                                                                if ((fi40 & 32768) > 32768) {
                                                                    fmt.Println(v27, v3)
                                                                    var v49 uint64 = 18446744073709551615
                                                                    var v50 uint = 4294967295
                                                                    var v51 int64 = 0
                                                                    var v52 uint32 = 2147483648
                                                                    _ = v49
                                                                    _ = v50
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                var v53 int64 = v11
                                                                _ = v53
                                                            }
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        v54 := GenS7{F0: 1}
                                                        v55 := GenS8{F0: 218, F1: (!(1 < 746)), F2: ""}
                                                        var v56 uint8 = 128
                                                        _ = fi40
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    _ = fi39
                                                }
                                                if (!(((3.14159 * 0.1) + (0.0 + 0.1)) == 3.14159)) {
                                                    var v57 uint8 = 30
                                                    v58 := GenS3{F0: "foo bar", F1: (902 < 65535), F2: v3}
                                                    var v59 uint64 = 1
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                            }
                                            if (!false) {
                                                var v60 float64 = -2.25
                                                v11 += int64(v27)
                                                v3 += v3
                                                _ = v60
                                            } else {
                                                var v61 float32 = -2.25
                                                var v62 int = (v27 - v27)
                                                _ = v61
                                                _ = v62
                                            }
                                        } else {
                                            _ = len(v2)
                                        }
                                    }
                                    var v63 int = v27
                                    _ = v63
                                }
                                var v64 int64 = (^(v13))
                                _ = v64
                            }
                        case 127:
                            switch (v27 * int(v3)) {
                            case 1:
                                var v65 float32 = 1e10
                                _ = v65
                            case 804:
                                for fi66 := 0; fi66 < 4; fi66++ {
                                    for fi67 := 0; fi67 < 5; fi67++ {
                                        switch 691 {
                                        case -1:
                                            if (!true) {
                                                v68 := GenS2{F0: 42}
                                                var v69 float64 = 3.14159
                                                _ = v68
                                                _ = v69
                                            }
                                        case 32768:
                                            switch "\n\t" {
                                            case "":
                                                v70 := GenS1{F0: fi66}
                                                _ = v70
                                            default:
                                                delete(v2, (fi66 * 32768))
                                            }
                                        case -9223372036854775808:
                                            if ((!(42 == 255)) || true) {
                                                var v71 uint8 = 2
                                                switch v27 {
                                                case -420:
                                                    var v72 uint64 = 992
                                                    _ = v72
                                                case 299:
                                                    var v73 uint32 = 351
                                                    _ = v73
                                                }
                                                _ = v71
                                            } else {
                                                if false {
                                                    fmt.Println(fi67, fi67)
                                                    var v74 int = fi67
                                                    if true {
                                                        if true {
                                                            v75 := GenS9{F0: 32767, F1: -2147483648, F2: 1}
                                                            _ = v75
                                                        } else {
                                                            var v76 int16 = 1000
                                                            var v77 float64 = 1e10
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v74
                                                }
                                                var v78 uint = 4294967295
                                                fmt.Println(v27)
                                                _ = v78
                                            }
                                        default:
                                            var v79 bool = ((v3 / ((v3) | 1)) >= 1)
                                            _ = v79
                                        }
                                        _ = fi67
                                    }
                                    for fi80 := 0; fi80 < 15; fi80++ {
                                        v3 = 255
                                        v1 = append(v1, 901)
                                        _ = fi80
                                    }
                                    for fi81 := 0; fi81 < 19; fi81++ {
                                        if false {
                                            fmt.Println(v11, v27)
                                        }
                                        switch fi81 {
                                        case 32767:
                                            v13 -= 0
                                        case -606:
                                            _ = len(v1)
                                        default:
                                            v82 := GenS5{F0: "foo bar"}
                                            _ = v82
                                        }
                                        v83 := GenS10{F0: 2}
                                        _ = fi81
                                        _ = v83
                                    }
                                    _ = fi66
                                }
                            default:
                                var v84 int = 127
                                _ = v84
                            }
                        default:
                            v85 := GenS1{F0: v27}
                            _ = v85
                        }
                        var v86 float64 = (0.0 + 1e10)
                        _ = v86
                    }
                    fmt.Println(v11)
                    var v87 uint64 = 1
                    _ = v87
                } else {
                    v88 := GenS5{F0: ""}
                    v89 := v2[(v27 << ((42) & 63))]
                    _ = v88
                    _ = v89
                }
                _ = v27
            }
            if (v11 > v11) {
                fmt.Println(v11, v13)
            } else {
                v3 += v3
                if ((81 >= 1) && false) {
                    for fi90 := 0; fi90 < 16; fi90++ {
                        v91 := GenS11{F0: 700, F1: int(v13), F2: 42}
                        fmt.Println(fi90, v3)
                        _ = fi90
                        _ = v91
                    }
                }
            }
            v92 := GenS2{F0: 42}
            _ = v92
        }
    }
    var v93 float32 = 3.14159
    v93 = 1e10
    _ = len(v2)
    v93 = v93
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v93
}
