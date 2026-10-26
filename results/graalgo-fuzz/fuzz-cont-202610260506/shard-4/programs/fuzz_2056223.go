package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int64; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 uint8 }

func main() {
    var v1 uint16 = 178
    var v2 float64 = 0.1
    var v3 int32 = -1
    var v4 uint16 = 65535
    var v5 uint64 = 1
    v2 = 1e10
    v6 := GenS1{F0: int(v1)}
    switch -2147483648 {
    case 32767:
        v7 := GenS2{F0: 32768, F1: -9223372036854775808, F2: (1 == 62)}
        _ = v7
    case 2147483647:
        v8 := []int{0, 583, 32768, -1}
        _ = v8
    case 32768:
        if false {
            var v9 uint = 65535
            var v10 int64 = 467
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: 0, F1: -539, F2: false}
            _ = v11
        }
    default:
        v12 := GenS3{F0: 1000}
        _ = v12
    }
    if false {
        switch "foo bar" {
        case "a":
            var v13 int = 22
            _ = v13
        case "foo bar":
            var v14 bool = true
            _ = v14
        case "hello":
            fmt.Println(v3, v5)
        }
        var v15 int8 = 127
        var v16 int8 = (^(v15))
        _ = v15
        _ = v16
    } else {
        if (false || false) {
            switch ("hello" + "a") {
            case "foo bar":
                for fi17 := 0; fi17 < 5; fi17++ {
                    v5 += v5
                    _ = fi17
                }
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                if false {
                    v1 -= (v1 - (^(v4)))
                } else {
                    v19 := []int{int(v5), (fi18 >> ((2) & 63)), fi18, fi18}
                    _ = v19
                }
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 uint8 = 255
                    switch "a" {
                    case "":
                        v22 := GenS4{F0: 108, F1: (((v21 & (v21 - v21)) != (v21 & 1)) || false)}
                        _ = v22
                    case "foo bar":
                        var v23 float32 = 3.14159
                        _ = v23
                    default:
                        var v24 int32 = (v3 - (v3 * (v3 << ((18446744073709551615) & 63))))
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v5)
                _ = fi18
            }
        } else {
            v25 := []int{-2147483648}
            if true {
                if (255 == 42) {
                    var v26 uint32 = 1
                    var v27 int16 = -1
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    if false {
                        fmt.Println(v3, v5)
                        if true {
                            v1 = v1
                        } else {
                            var v29 int = fi28
                            _ = v29
                        }
                        var v30 int16 = -32768
                        _ = v30
                    } else {
                        var v31 uint32 = 263
                        _ = v31
                    }
                    v2 += float64(1.5)
                    if len(v25) > 0 {
                        v25[0] = 127
                    }
                    _ = fi28
                }
            }
            _ = v25
        }
        fmt.Println(v1)
    }
    if ((v5 ^ uint64(v4)) <= v5) {
        v5 -= (v5 * (^(v5)))
        var v32 uint32 = 42
        v4 = uint16(v5)
        _ = v32
    }
    v4 = v4
    var v33 int64 = -410
    fmt.Println(v1, v4)
    for fi34 := 0; fi34 < 15; fi34++ {
        switch "hello" {
        case "hello":
            if (!true) {
                v35 := GenS5{F0: v33, F1: 607, F2: 127}
                v36 := GenS6{F0: ("foo bar" + "foo bar"), F1: "a"}
                _ = v35
                _ = v36
            }
        case "foo bar":
            if ((v5 << ((4294967295) & 63)) > (v5 / (((^(v5))) | 1))) {
                var v37 uint8 = 128
                if false {
                    v3 = 149
                }
                var v38 int16 = -1
                _ = v37
                _ = v38
            } else {
                v5 -= (^(v5))
                v5 = (v5 - v5)
            }
        default:
            var v39 uint32 = 2147483647
            _ = v39
        }
        _ = fi34
    }
    fmt.Println(v2)
    v40 := GenS5{F0: (v33 >> ((0) & 63)), F1: 4294967295, F2: 0}
    switch "a" {
    case "":
        if (32767 > 1) {
            v4 += v1
            v2 = v2
        }
    }
    if ((v5 ^ (v5 & v5)) >= 18446744073709551615) {
        v3 = 42
        for fi41 := 0; fi41 < 15; fi41++ {
            if false {
                if true {
                    if false {
                        for fi42 := 0; fi42 < 4; fi42++ {
                            var v43 uint16 = (v4 << ((18446744073709551615) & 63))
                            v2 = v2
                            switch "\n\t" {
                            case "\n\t":
                                if (v2 < float64(0.0)) {
                                    v5 = v5
                                    var v44 int16 = -1
                                    if (!("hello" < "\n\t")) {
                                        if false {
                                            v44 = v44
                                            v45 := GenS5{F0: v33, F1: 42, F2: 100}
                                            v43 = (v43 >> ((63) & 63))
                                            _ = v45
                                        }
                                        v46 := []int{}
                                        v1 = (v1 * (v4 | (^(v1))))
                                        _ = v46
                                    }
                                    _ = v44
                                }
                            case "hello":
                                v47 := []int{fi42, -1}
                                _ = v47
                            default:
                                var v48 float32 = -2.25
                                _ = v48
                            }
                            _ = fi42
                            _ = v43
                        }
                    }
                } else {
                    for fi49 := 0; fi49 < 9; fi49++ {
                        var v50 int = (fi49 - (^(fi41)))
                        var v51 int16 = 1
                        v3 = (v3 + (-(v3)))
                        if (true || (0 != 65535)) {
                            if (true || false) {
                                switch "foo bar" {
                                case "hello":
                                    switch "\n\t" {
                                    case "\n\t":
                                        v4 = v4
                                    }
                                case "\n\t":
                                    var v52 uint64 = v5
                                    if (!("foo bar" == "foo bar")) {
                                        v2 = (float64(-2.25) / v2)
                                        if (true && true) {
                                            if true {
                                                var v53 uint64 = uint64(v1)
                                                if true {
                                                    if true {
                                                        fmt.Println(v53)
                                                        var v54 int = fi49
                                                        var v55 int64 = 1
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if false {
                                                                v51 = (-(v51))
                                                            }
                                                        case "foo bar":
                                                            if (true && (true || true)) {
                                                                v56 := GenS3{F0: (v51 << ((190) & 63))}
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v57 := GenS2{F0: 128, F1: (v33 ^ -445), F2: true}
                                                                    v50 = fi49
                                                                    _ = v57
                                                                case "":
                                                                    fmt.Println(v50)
                                                                }
                                                                _ = v56
                                                            } else {
                                                                switch ("" + "a") {
                                                                case "foo bar":
                                                                    var v58 uint8 = 2
                                                                    _ = v58
                                                                case "a":
                                                                    fmt.Println(v51, v1, v2)
                                                                    if true {
                                                                        v59 := []int{fi49, (fi49 % ((-218) | 1)), fi49, fi49}
                                                                        _ = v59
                                                                    } else {
                                                                        var v60 int16 = (v51 ^ int16(v33))
                                                                        v52 -= v5
                                                                        v51 = -32768
                                                                        _ = v60
                                                                    }
                                                                }
                                                                var v61 uint32 = 1
                                                                fmt.Println(v4, fi49, v53)
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                    var v62 int = fi41
                                                    _ = v62
                                                } else {
                                                    v1 -= v1
                                                }
                                                v63 := GenS7{F0: v4, F1: 65535, F2: 125}
                                                _ = v53
                                                _ = v63
                                            }
                                            v64 := []int{(fi49 ^ (v50 + fi49))}
                                            _ = v64
                                        }
                                        switch fi49 {
                                        case -1:
                                            v65 := GenS8{F0: 4294967295, F1: (v50 & fi41)}
                                            _ = v65
                                        case 4294967295:
                                            switch ("foo bar" + "") {
                                            case "hello":
                                                v1 = (v1 << ((0) & 63))
                                            default:
                                                if true {
                                                    if true {
                                                        var v66 float32 = 1e10
                                                        v67 := GenS9{F0: -128}
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        v68 := GenS10{F0: v33, F1: (fi41 | int(v51))}
                                                        v52 -= uint64(v1)
                                                        v69 := []int{fi41}
                                                        var v70 uint16 = (v4 << ((464) & 63))
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    fmt.Println(v1, v52, v1)
                                                }
                                            }
                                            v3 = (v3 >> ((59) & 63))
                                        case 9223372036854775807:
                                            v50 += int(v51)
                                        }
                                    } else {
                                        fmt.Println(v51, v5)
                                        var v71 uint32 = 2
                                        _ = v71
                                    }
                                    if ((!false) && (true && ((195 <= 127) || true))) {
                                        fmt.Println(v3)
                                        v72 := GenS11{F0: -128, F1: 3}
                                        if (132 <= 63) {
                                            if false {
                                                v73 := GenS6{F0: ("foo bar" + "hello"), F1: ("\n\t" + "hello")}
                                                _ = v73
                                            } else {
                                                var v74 uint16 = (v4 - (^(v4)))
                                                fmt.Println(v1, v2)
                                                _ = v74
                                            }
                                        } else {
                                            switch ("foo bar" + "foo bar") {
                                            case "foo bar":
                                                switch "" {
                                                case "a":
                                                    var v75 int16 = v51
                                                    _ = v75
                                                case "":
                                                    var v76 uint32 = 2147483648
                                                    _ = v76
                                                }
                                            case "a":
                                                var v77 uint32 = 993
                                                _ = v77
                                            case "hello":
                                                if (2147483647 <= uint32(v4)) {
                                                    v50 -= fi49
                                                    var v78 uint32 = 2147483648
                                                    v79 := []int{(-(fi49)), (fi41 - (fi49 << ((0) & 63)))}
                                                    _ = v78
                                                    _ = v79
                                                }
                                            default:
                                                fmt.Println(v4, v50)
                                            }
                                            v80 := []int{fi49, v50, -995}
                                            v1 = uint16(v52)
                                            fmt.Println(fi49, v3, fi49)
                                            _ = v80
                                        }
                                        _ = v72
                                    } else {
                                        v5 += v52
                                    }
                                    v81 := GenS9{F0: int8(v3)}
                                    _ = v52
                                    _ = v81
                                }
                            } else {
                                v51 -= v51
                                if true {
                                    if (2 >= 128) {
                                        v1 = v4
                                        v3 = (v3 << ((131) & 63))
                                        if true {
                                            var v82 int8 = 42
                                            v51 = (^(v51))
                                            v2 += 0.1
                                            _ = v82
                                        } else {
                                            if true {
                                                fmt.Println(fi41)
                                                var v83 uint8 = 0
                                                _ = v83
                                            }
                                        }
                                    } else {
                                        v3 -= (v3 * v3)
                                    }
                                } else {
                                    v84 := GenS3{F0: v51}
                                    if false {
                                        var v85 uint8 = uint8(v1)
                                        _ = v85
                                    } else {
                                        fmt.Println(v1, v1, fi49)
                                    }
                                    _ = v84
                                }
                            }
                            var v86 uint = uint(v5)
                            switch -215 {
                            case -2147483648:
                                fmt.Println(v5, v51, v51)
                                if false {
                                    v87 := GenS2{F0: (v50 + fi49), F1: (v33 + (v33 ^ (v33 ^ (v33 >> ((v86) & 63))))), F2: true}
                                    if (65535 <= 539) {
                                        v86 -= (^(v86))
                                        fmt.Println(v4)
                                    } else {
                                        v88 := GenS4{F0: v33, F1: true}
                                        fmt.Println(v86, v50)
                                        _ = v88
                                    }
                                    var v89 uint = 1
                                    _ = v87
                                    _ = v89
                                } else {
                                    v33 = int64(v1)
                                    var v90 bool = (false && (!(v86 <= v86)))
                                    _ = v90
                                }
                            default:
                                v91 := []int{fi49}
                                _ = v91
                            }
                            _ = v86
                        } else {
                            if true {
                                var v92 int64 = (v33 + v33)
                                _ = v92
                            }
                            fmt.Println(v33, v33, fi49)
                        }
                        var v93 uint16 = 1
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v93
                    }
                }
                v94 := []int{int(v5), fi41}
                v2 += ((v2 - float64(0.0)) * v2)
                _ = v94
            } else {
                switch "" {
                case "hello":
                    var v95 uint8 = 2
                    _ = v95
                }
                var v96 string = ""
                v97 := []int{0, (fi41 + fi41), fi41}
                _ = v96
                _ = v97
            }
            _ = fi41
        }
        switch -148 {
        case 128:
            for fi98 := 0; fi98 < 19; fi98++ {
                v5 += v5
                _ = fi98
            }
        case 32768:
            var v99 float64 = 1.5
            _ = v99
        case 777:
            if false {
                if true {
                    v4 = v4
                    v100 := []int{519, 42, -219}
                    _ = v100
                }
            }
        }
    }
    v101 := []int{2147483647, -875}
    v33 = (v33 + (v33 - (v33 & (-(v33)))))
    if true {
        v2 = (v2 + v2)
        switch ("hello" + "a") {
        case "a":
            fmt.Println(v3)
        case "":
            v102 := GenS10{F0: (v33 >> ((615) & 63)), F1: 128}
            _ = v102
        }
    } else {
        var v103 int32 = int32(v33)
        v104 := GenS12{F0: 4294967295, F1: (-(v103))}
        var v105 float64 = 1e10
        _ = v103
        _ = v104
        _ = v105
    }
    fmt.Println(v1, v5, v2)
    v106 := GenS13{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v106
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
