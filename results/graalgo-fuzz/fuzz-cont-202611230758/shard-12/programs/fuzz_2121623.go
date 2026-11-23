package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int16; F2 uint32 }
type GenS11 struct { F0 int64; F1 string; F2 uint32 }
type GenS12 struct { F0 uint32; F1 uint16 }
type GenS13 struct { F0 string; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 0
    var v2 int = -220
    v3 := make(map[uint32]string)
    var v4 int64 = 505
    switch v2 {
    case -168:
        v5 := GenS1{F0: 2147483648}
        _ = v5
    case 2147483647:
        fmt.Println(v4)
    default:
        var v6 int64 = (v4 >> ((0) & 63))
        _ = v6
    }
    fmt.Println(v4, v1, v4)
    var v7 float32 = (-2.25 + -2.25)
    if (true && true) {
        var v8 uint16 = 1
        _ = v8
    } else {
        v1 = (^(v1))
        if (!false) {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int32 = 1
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(v1, v4, v2)
                    v12 := GenS2{F0: fi9, F1: 127}
                    fmt.Println(v2, v10, v4)
                    if false {
                        _ = len(v3)
                        var v13 uint8 = 255
                        var v14 uint8 = (v13 & (^(v13)))
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint8 = 128
                        var v16 uint = 824
                        _ = v15
                        _ = v16
                    }
                    if (!((true && false) && true)) {
                        var v17 float64 = 1.5
                        switch v2 {
                        case -811:
                            v18 := GenS3{F0: (v4 - (v4 << ((2) & 63)))}
                            _ = v18
                        case 9223372036854775807:
                            var v19 float32 = v7
                            _ = v19
                        case -2147483648:
                            switch (fi11 * fi11) {
                            case 32768:
                                var v20 int = fi11
                                _ = v20
                            case 128:
                                v21 := []int{}
                                _ = v21
                            default:
                                switch 582 {
                                case -576:
                                    fmt.Println(v4)
                                default:
                                    if (false || ((v10 + v10) == v10)) {
                                        fmt.Println(v17)
                                    } else {
                                        var v22 int = -1
                                        var v23 int16 = 1000
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                        default:
                            if (!true) {
                                var v24 uint16 = 65535
                                _ = v24
                            } else {
                                var v25 uint64 = (v1 % (((v1 & (v1 % (((v1 << ((4294967295) & 63))) | 1)))) | 1))
                                if (v17 > v17) {
                                    var v26 int = (-(v2))
                                    fmt.Println(v7)
                                    var v27 int = v2
                                    var v28 uint32 = 127
                                    var v29 int64 = v4
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 float64 = 0.1
                                    v2 = (fi11 - 538)
                                    v17 += v17
                                    _ = v30
                                }
                                var v31 uint = 625
                                _ = v25
                                _ = v31
                            }
                        }
                        _ = v17
                    }
                    var v32 uint = 255
                    _ = fi11
                    _ = v12
                    _ = v32
                }
                v33 := GenS1{F0: 2}
                _ = fi9
                _ = v10
                _ = v33
            }
            var v34 uint64 = (v1 << ((462) & 63))
            if false {
                if false {
                    if (true || false) {
                        fmt.Println(v1, v4, v7)
                        if false {
                            v35 := GenS4{F0: v34}
                            _ = v35
                        } else {
                            v1 = v34
                            var v36 uint8 = 0
                            _ = v36
                        }
                    }
                    var v37 int16 = 32767
                    _ = v37
                }
                fmt.Println(v4, v34, v7)
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 uint = 2
                    switch v2 {
                    case 635:
                        delete(v3, 899)
                    }
                    _ = fi38
                    _ = v39
                }
            } else {
                v40 := []int{v2, (-(v2)), -113, v2}
                v2 += 1
                if (607 != 65535) {
                    var v41 int8 = -1
                    var v42 string = ("foo bar" + "hello")
                    v43 := GenS5{F0: 18, F1: v2}
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    v1 = (v34 & 488)
                }
                _ = v40
            }
            _ = v34
        } else {
            for fi44 := 0; fi44 < 8; fi44++ {
                if ((2 >= (^(v1))) || false) {
                    v7 = (3.14159 + v7)
                    v1 = v1
                    if (255 != 568) {
                        for fi45 := 0; fi45 < 9; fi45++ {
                            switch "hello" {
                            case "\n\t":
                                if true {
                                    _ = len(v3)
                                    var v46 uint64 = (v1 >> ((264) & 63))
                                    _ = v46
                                } else {
                                    var v47 uint16 = 709
                                    var v48 uint8 = 30
                                    switch fi45 {
                                    case 11:
                                        fmt.Println(v1)
                                    case 42:
                                        v49 := GenS3{F0: (v4 & v4)}
                                        _ = v49
                                    case 9223372036854775807:
                                        var v50 uint16 = (v47 / ((v47) | 1))
                                        _ = v50
                                    default:
                                        v51 := GenS1{F0: 2}
                                        _ = v51
                                    }
                                    var v52 bool = (int(v48) <= v2)
                                    if v52 {
                                        v53 := GenS2{F0: (fi44 & fi45), F1: 127}
                                        var v54 uint32 = 619
                                        _ = v53
                                        _ = v54
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v7 = ((-2.25 * (v7 - (float32(float64((((float32((3.14159 + 1e10)) * (v7 + ((1e10 / 1e10) / ((v7 + (float32(float64(v7)) * v7)) + 1e10)))) - v7) - v7))) / float32(3.14159)))) - 1.5)
                                        case "":
                                            if v52 {
                                                v55 := GenS6{F0: 1000}
                                                v3[4294967295] = ""
                                                switch ("hello" + "") {
                                                case "foo bar":
                                                    fmt.Println(fi45)
                                                    v56 := GenS5{F0: 0, F1: (fi44 | (v2 - v2))}
                                                    _ = v56
                                                }
                                                v57 := GenS5{F0: -128, F1: -1}
                                                _ = v55
                                                _ = v57
                                            } else {
                                                v2 = (fi45 - (fi45 - v2))
                                            }
                                        default:
                                            if false {
                                                v58 := []int{fi45, fi44, fi45}
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                                if v52 {
                                                    v2 += (v2 % (((fi44 - fi44)) | 1))
                                                    var v59 int8 = 127
                                                    v60 := GenS5{F0: int8(v47), F1: fi44}
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v61 := GenS7{F0: 1, F1: (-(v4))}
                                                    _ = v61
                                                }
                                                _ = v58
                                            }
                                        }
                                        var v62 uint32 = uint32(v48)
                                        var v63 uint8 = 1
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v47
                                    _ = v48
                                    _ = v52
                                }
                            case "hello":
                                v1 = (v1 | 100)
                            default:
                                v7 += v7
                            }
                            switch ("" + "foo bar") {
                            case "hello":
                                fmt.Println(v4, fi45, fi44)
                            case "":
                                if false {
                                    var v64 uint = 626
                                    _ = v64
                                } else {
                                    v65 := GenS8{F0: 89}
                                    v2 += (fi45 % (((fi44 % (((v2 << ((4294967295) & 63))) | 1))) | 1))
                                    var v66 int = (v2 / (((fi45 * fi45)) | 1))
                                    _ = v65
                                    _ = v66
                                }
                                if false {
                                    switch (v2 / ((v2) | 1)) {
                                    case -9223372036854775808:
                                        v4 += v4
                                    case 32768:
                                        v67 := []int{v2, fi44, (fi45 & -148)}
                                        _ = v67
                                    case 128:
                                        var v68 uint64 = (v1 >> ((18446744073709551615) & 63))
                                        _ = v68
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                            case "a":
                                v3[539] = ""
                            default:
                                var v69 uint32 = 2147483647
                                _ = v69
                            }
                            var v70 uint = 537
                            _ = fi45
                            _ = v70
                        }
                    } else {
                        var v71 float32 = v7
                        _ = v71
                    }
                }
                _ = fi44
            }
            v72 := []int{(v2 * (v2 + 9223372036854775807)), -1, v2, (v2 << ((18446744073709551615) & 63))}
            _ = v72
        }
        if false {
            var v73 int8 = 1
            v4 = (v4 / (((-(v4))) | 1))
            v74 := []int{(v2 + -233)}
            _ = v73
            _ = v74
        }
    }
    v2 -= (v2 | v2)
    for fi75 := 0; fi75 < 2; fi75++ {
        v76 := GenS9{F0: "a"}
        var v77 uint64 = (v1 - (^(v1)))
        _ = fi75
        _ = v76
        _ = v77
    }
    v2 -= v2
    v2 = (v2 / ((v2) | 1))
    for fi78 := 0; fi78 < 12; fi78++ {
        var v79 uint64 = v1
        for fi80 := 0; fi80 < 13; fi80++ {
            var v81 bool = (-26 >= -1)
            var v82 float64 = 1e10
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = fi78
        _ = v79
    }
    delete(v3, 959)
    for fi83 := 0; fi83 < 2; fi83++ {
        if (-164 >= 1000) {
            if true {
                for fi84 := 0; fi84 < 17; fi84++ {
                    v2 += 951
                    v85 := []int{fi84}
                    _ = fi84
                    _ = v85
                }
                var v86 string = "\n\t"
                v87 := GenS10{F0: true, F1: 32767, F2: 42}
                _ = v86
                _ = v87
            } else {
                v4 -= 1
                for fi88 := 0; fi88 < 20; fi88++ {
                    fmt.Println(v4, fi88)
                    v89 := GenS11{F0: (v4 & v4), F1: "\n\t", F2: 4294967295}
                    _ = fi88
                    _ = v89
                }
                if false {
                    for fi90 := 0; fi90 < 15; fi90++ {
                        var v91 int16 = 432
                        var v92 string = "foo bar"
                        if false {
                            v93 := GenS12{F0: 195, F1: 0}
                            _ = v93
                        } else {
                            var v94 int16 = 1000
                            v94 -= (v94 * -566)
                            v7 = (3.14159 * 1.5)
                            _ = v94
                        }
                        _ = fi90
                        _ = v91
                        _ = v92
                    }
                    v95 := v3[2]
                    var v96 uint8 = 2
                    _ = v95
                    _ = v96
                }
            }
            v4 = v4
        } else {
            v97 := v3[596]
            fmt.Println(fi83, v2)
            _ = v97
        }
        v98 := GenS13{F0: "a", F1: 42, F2: 65535}
        _ = fi83
        _ = v98
    }
    if ((((v7 - ((float32(((0.1 / 1e10) / 3.14159)) / v7) + 1.5)) + 0.1) / v7) != 0.0) {
        var v99 float64 = ((((-2.25 / float64(v7)) / (1e10 / 1.5)) - (1.5 + (1e10 * float64(v7)))) + (3.14159 - ((-2.25 - 1.5) / 1.5)))
        v100 := GenS12{F0: 628, F1: 42}
        if (466 <= 32767) {
            v1 = (v1 + v1)
            fmt.Println(v2)
            v101 := []int{v2, (v2 / ((9223372036854775807) | 1)), (v2 + -2147483648), v2}
            _ = v101
        }
        _ = v100
        _ = v99
    } else {
        v102 := []int{(v2 - (v2 - v2)), v2}
        _ = v102
    }
    v7 = 1e10
    fmt.Println(v4, v4, v7)
    var v103 bool = (1e10 <= v7)
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
