package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int64; F1 uint32; F2 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int; F1 uint64; F2 int }
type GenS11 struct { F0 int16; F1 int16 }

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    var v3 int8 = -26
    var v4 float64 = 0.0
    var v5 float64 = (v4 * v4)
    if true {
        fmt.Println(v3, v1)
        var v6 uint32 = 4294967295
        _ = v6
    }
    var v7 string = v1
    v3 += -59
    if false {
        if true {
            if (-32768 == -32768) {
                var v8 int64 = -1
                v9 := GenS1{F0: 42, F1: 2147483647}
                v10 := []int{962, 32768, int(v3), 127}
                _ = v10
                _ = v8
                _ = v9
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch (fi11 ^ (-(fi11))) {
                    case -161:
                        fmt.Println(fi11)
                    case 127:
                        var v12 uint8 = 255
                        _ = v12
                    }
                    if (4294967295 >= 940) {
                        delete(v2, (("" + "foo bar") + (v7 + "")))
                        v13 := []int{fi11, fi11, fi11}
                        _ = v13
                    }
                    v4 += 1e10
                    _ = fi11
                }
                var v14 int32 = -2147483648
                var v15 string = (v1 + ("hello" + v7))
                _ = v14
                _ = v15
            }
            fmt.Println(v4, v1, v5)
        } else {
            delete(v2, (((v1 + "foo bar") + (("foo bar" + (v1 + (("hello" + "\n\t") + (((v7 + "foo bar") + v1) + v7)))) + "")) + v7))
        }
        if (false && false) {
            var v16 uint64 = 0
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v1, v3, v3)
            if true {
                v18 := GenS2{F0: 1, F1: 255, F2: 2}
                _ = v18
            } else {
                v19 := []int{fi17, fi17, fi17}
                if false {
                    v7 = v7
                }
                _ = v19
            }
            var v20 float32 = (1.5 + float32(((1.5 * (v4 + 0.1)) + (v5 + float64((0.0 - (0.1 / 3.14159)))))))
            _ = fi17
            _ = v20
        }
    } else {
        v21 := v2[(v7 + v7)]
        switch 454 {
        case 371:
            var v22 uint64 = 18446744073709551615
            _ = v22
        case 0:
            fmt.Println(v21, v3, v7)
        case 128:
            if (38 != v3) {
                v23 := []int{704, 1, -23}
                switch "\n\t" {
                case "":
                    var v24 string = ("\n\t" + (v1 + v7))
                    _ = v24
                case "\n\t":
                    v4 = v5
                default:
                    v25 := GenS1{F0: -1, F1: 2147483647}
                    _ = v25
                }
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v23
            } else {
                if v21 {
                    fmt.Println(v7, v4)
                }
            }
        }
        _ = v21
    }
    var v26 uint32 = 0
    v27 := GenS1{F0: 42, F1: -1}
    switch -1 {
    case 603:
        v1 = "hello"
    case -2147483648:
        var v28 string = ((v1 + "\n\t") + v7)
        _ = v28
    default:
        var v29 bool = (!true)
        _ = v29
    }
    v2[("a" + v1)] = true
    var v30 bool = ((false || false) && (0.0 > 0.1))
    v4 = 0.0
    fmt.Println(v7)
    if v30 {
        switch 32767 {
        case -129:
            var v31 float32 = 0.1
            _ = v31
        case 32768:
            var v32 float64 = (v5 - 1e10)
            _ = v32
        case -2147483648:
            v33 := GenS3{F0: -1}
            _ = v33
        }
    }
    switch -2147483648 {
    case 2147483647:
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 float32 = ((1e10 - 1.5) * (1.5 / (float32((v5 + ((float64(3.14159) + v4) + v5))) / (((float32((v4 + v5)) + float32(v4)) - (0.0 * 3.14159)) - 1.5))))
            if true {
                v36 := GenS3{F0: 0}
                v2[(("" + v7) + ((v1 + (v1 + v7)) + (v1 + ("a" + v1))))] = v30
                fmt.Println(v4, v30)
                _ = v36
            }
            _ = fi34
            _ = v35
        }
    case 9223372036854775807:
        for fi37 := 0; fi37 < 15; fi37++ {
            switch (v1 + "foo bar") {
            case "foo bar":
                if true {
                    if false {
                        for fi38 := 0; fi38 < 19; fi38++ {
                            if true {
                                if v30 {
                                    if (-1 == -1) {
                                        switch "hello" {
                                        case "a":
                                            var v39 uint8 = 202
                                            _ = v39
                                        case "":
                                            v40 := GenS4{F0: 4294967295, F1: v7, F2: 9223372036854775807}
                                            _ = v40
                                        }
                                    }
                                    fmt.Println(v4, v7, fi37)
                                } else {
                                    fmt.Println(v7, v3, v3)
                                    v41 := GenS5{F0: (((((480 > 4294967295) || v30) || false) || ((v5 * v5) == v5)) || ((v3 < v3) || (196 != 110))), F1: 255, F2: 2}
                                    v42 := GenS2{F0: (v26 ^ (v26 % (((v26 - v26)) | 1))), F1: 4, F2: v26}
                                    _ = v41
                                    _ = v42
                                }
                                switch "a" {
                                case "":
                                    var v43 float32 = (-2.25 * 0.0)
                                    _ = v43
                                case "\n\t":
                                    var v44 bool = (!((3.14159 - float32(1.5)) > ((0.0 * 0.1) - (1.5 / -2.25))))
                                    _ = v44
                                case "foo bar":
                                    v45 := v2[v7]
                                    _ = v45
                                }
                            } else {
                                var v46 uint64 = 294
                                fmt.Println(v1)
                                _ = v46
                            }
                            _ = fi38
                        }
                        switch (v1 + v7) {
                        case "hello":
                            fmt.Println(v30)
                        }
                    }
                } else {
                    var v47 int = 32767
                    _ = v47
                }
            case "":
                if (82 < 0) {
                    var v48 int8 = v3
                    _ = v48
                } else {
                    _ = len(v2)
                    if (9223372036854775807 < 42) {
                        for fi49 := 0; fi49 < 6; fi49++ {
                            if (true || (!true)) {
                                v50 := []int{(^(fi37))}
                                v4 -= v4
                                _ = v50
                            }
                            if (!v30) {
                                switch -467 {
                                case -1:
                                    v2[(v7 + v1)] = true
                                case 32767:
                                    var v51 float32 = ((float32(0.0) * 0.1) / -2.25)
                                    _ = v51
                                case 127:
                                    if v30 {
                                        var v52 float32 = 0.1
                                        var v53 uint32 = (v26 & (v26 - (v26 % ((2) | 1))))
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 float32 = 3.14159
                                    v55 := GenS6{F0: 1, F1: -1}
                                    _ = v54
                                    _ = v55
                                default:
                                    var v56 bool = false
                                    _ = v56
                                }
                                fmt.Println(v5, v3, v3)
                            } else {
                                if v30 {
                                    v2[("foo bar" + v1)] = (!((42 <= -1) || v30))
                                }
                                var v57 int32 = 42
                                _ = v57
                            }
                            _ = fi49
                        }
                        var v58 int32 = -134
                        for fi59 := 0; fi59 < 3; fi59++ {
                            if (!v30) {
                                var v60 float64 = float64((0.0 + -2.25))
                                if (804 == 4294967295) {
                                    if v30 {
                                        v26 = (^(v26))
                                        if true {
                                            v60 += (((1.5 * 0.0) - v4) + v60)
                                            v61 := GenS6{F0: 2, F1: 42}
                                            _ = v61
                                        }
                                        v62 := GenS4{F0: (^(v26)), F1: (v7 + ("\n\t" + ("foo bar" + v1))), F2: -157}
                                        _ = v62
                                    } else {
                                        fmt.Println(v7)
                                        v60 = v60
                                    }
                                    v63 := GenS7{F0: 983, F1: v26, F2: 128}
                                    _ = v63
                                } else {
                                    v2["foo bar"] = v30
                                }
                                var v64 int32 = v58
                                _ = v60
                                _ = v64
                            }
                            var v65 string = v1
                            var v66 int16 = -1
                            _ = fi59
                            _ = v65
                            _ = v66
                        }
                        _ = v58
                    }
                    if true {
                        for fi67 := 0; fi67 < 2; fi67++ {
                            if (1e10 == 1e10) {
                                fmt.Println(v3)
                                var v68 uint16 = 2
                                if v30 {
                                    if (true && (580 >= 65535)) {
                                        switch -2147483648 {
                                        case 1:
                                            v3 += (v3 - (v3 & v3))
                                        case -223:
                                            switch (v7 + v7) {
                                            case "a":
                                                if v30 {
                                                    if false {
                                                        v26 = v26
                                                        var v69 uint64 = 1
                                                        switch -2147483648 {
                                                        case -1:
                                                            var v70 uint = 4294967295
                                                            _ = v70
                                                        default:
                                                            v71 := []int{fi67, fi67, (fi37 << ((uint(v26)) & 63))}
                                                            _ = v71
                                                        }
                                                        _ = v69
                                                    }
                                                    var v72 string = ("\n\t" + v1)
                                                    var v73 uint64 = 313
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    var v74 uint64 = 157
                                                    if v30 {
                                                        v1 = v1
                                                        fmt.Println(v74)
                                                        var v75 uint16 = 473
                                                        _ = v75
                                                    } else {
                                                        if (2147483647 < v26) {
                                                            var v76 uint16 = (v68 >> ((341) & 63))
                                                            _ = v76
                                                        }
                                                        v77 := GenS2{F0: 0, F1: 1, F2: v26}
                                                        var v78 float64 = 0.1
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    _ = v74
                                                }
                                            case "hello":
                                                v79 := GenS8{F0: v7}
                                                _ = v79
                                            case "foo bar":
                                                switch ((((v7 + v7) + "\n\t") + ("a" + (v7 + "a"))) + v7) {
                                                case "hello":
                                                    v80 := []int{fi67}
                                                    _ = v80
                                                default:
                                                    v81 := GenS3{F0: (-(v3))}
                                                    _ = v81
                                                }
                                            }
                                        case 537:
                                            v82 := GenS9{F0: 438}
                                            _ = v82
                                        default:
                                            var v83 int32 = 1
                                            _ = v83
                                        }
                                    }
                                    fmt.Println(v26, v68)
                                    v3 = (v3 ^ (v3 & 42))
                                    v84 := []int{-372, (fi37 * 412), (fi67 >> ((4294967295) & 63)), (fi37 >> ((65535) & 63))}
                                    if v30 {
                                        switch fi67 {
                                        case 1:
                                            v85 := GenS10{F0: (fi67 - fi67), F1: 4294967295, F2: fi37}
                                            _ = v85
                                        case -9223372036854775808:
                                            var v86 float64 = v4
                                            _ = v86
                                        case -926:
                                            v30 = v30
                                            v1 = v1
                                            var v87 int64 = -173
                                            _ = v87
                                        }
                                        var v88 uint32 = v26
                                        _ = v88
                                    } else {
                                        fmt.Println(v5, v26, v68)
                                        fmt.Println(v3, fi67)
                                        switch (fi67 ^ (fi67 ^ (fi37 / (((fi67 & (fi67 & (fi67 ^ (fi67 / ((127) | 1)))))) | 1)))) {
                                        case -72:
                                            var v89 uint16 = 2
                                            _ = v89
                                        case 128:
                                            v84 = append(v84, fi37)
                                            var v90 string = "a"
                                            _ = v90
                                        case -9223372036854775808:
                                            if len(v84) >= 2 {
                                                v84 = v84[1:len(v84)]
                                            }
                                        }
                                    }
                                    v30 = (!v30)
                                    _ = v84
                                } else {
                                    var v91 int8 = (v3 / ((v3) | 1))
                                    _ = v91
                                }
                                _ = v68
                            }
                            v4 -= float64(float32(1.5))
                            v26 = (v26 | v26)
                            _ = fi67
                        }
                        v92 := GenS6{F0: 2, F1: 42}
                        if true {
                            fmt.Println(v26)
                        }
                        _ = v92
                    } else {
                        if (-1 == -9223372036854775808) {
                            var v93 uint32 = (v26 >> ((18446744073709551615) & 63))
                            v4 = v4
                            if v30 {
                                var v94 uint32 = v26
                                _ = v94
                            }
                            _ = v93
                        }
                        delete(v2, v7)
                    }
                }
            case "a":
                for fi95 := 0; fi95 < 1; fi95++ {
                    var v96 int32 = 1
                    _ = fi95
                    _ = v96
                }
            default:
                var v97 float64 = v5
                _ = v97
            }
            switch 32767 {
            case 2147483647:
                var v98 uint = 699
                _ = v98
            case 4294967295:
                v99 := []int{}
                _ = v99
            case 622:
                switch v1 {
                case "\n\t":
                    if (42 != 826) {
                        for fi100 := 0; fi100 < 8; fi100++ {
                            switch fi100 {
                            case 32768:
                                if v30 {
                                    var v101 uint = 0
                                    var v102 int8 = (v3 - (v3 % ((0) | 1)))
                                    var v103 float32 = 1.5
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                } else {
                                    var v104 uint8 = 128
                                    _ = v104
                                }
                            case 127:
                                var v105 float32 = (0.0 + 1.5)
                                _ = v105
                            case 2147483647:
                                v106 := GenS11{F0: 1000, F1: 32767}
                                _ = v106
                            }
                            v107 := GenS8{F0: ("" + (v1 + v1))}
                            v1 = v1
                            _ = fi100
                            _ = v107
                        }
                        fmt.Println(fi37, fi37, fi37)
                        v108 := v2[("hello" + v7)]
                        _ = v108
                    } else {
                        var v109 int16 = -774
                        var v110 uint16 = 255
                        _ = v109
                        _ = v110
                    }
                }
            }
            _ = fi37
        }
    default:
        var v111 int64 = 590
        _ = v111
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
}
