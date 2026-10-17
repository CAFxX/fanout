package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 int; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 string; F1 int; F2 uint8 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 uint }
func genFn0(p0 int) int64 {
    p0 += -1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 string = ("hello" + ("\n\t" + "\n\t"))
        v7 := []int{fi5, p0}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (-75 >= 1) {
        switch -771 {
        case -9223372036854775808:
            v8 := []int{}
            _ = v8
        }
    } else {
        var v9 float32 = 1.5
        p0 = p0
        _ = v9
    }
    switch "\n\t" {
    case "a":
        v10 := GenS1{F0: 127, F1: -1, F2: 18446744073709551615}
        _ = v10
    case "\n\t":
        var v11 string = ("" + ("hello" + "hello"))
        _ = v11
    case "":
        fmt.Println(p0)
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                switch "a" {
                case "foo bar":
                    v14 := GenS1{F0: 1, F1: 1000, F2: 2}
                    _ = v14
                case "\n\t":
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            var v15 int32 = -1
                            _ = v15
                        case "hello":
                            p0 -= 0
                        case "a":
                            var v16 int32 = 0
                            _ = v16
                        default:
                            switch 42 {
                            case 4294967295:
                                p0 += fi12
                            }
                        }
                        p0 = fi13
                    }
                case "":
                    v17 := []int{(-(fi12)), (-(fi13)), (fi12 | fi12)}
                    _ = v17
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    _ = p0
    return -1
}
func genFn1(p0 string, p1 bool, p2 uint16) uint16 {
    p1 = false
    var v18 int = 2147483647
    var v19 uint16 = p2
    if false {
        var v20 uint32 = uint32(v19)
        _ = v20
    } else {
        switch "\n\t" {
        case "":
            v21 := []int{(^(v18)), 396}
            _ = v21
        default:
            var v22 int = (v18 + v18)
            _ = v22
        }
    }
    var v23 bool = (104 < 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v23
    return p2
}

func main() {
    v1 := []uint32{697}
    var v2 uint64 = 326
    v3 := make(map[uint32]bool)
    var v4 uint8 = 0
    fmt.Println(v2, v4)
    v24 := GenS1{F0: v4, F1: 1, F2: (v2 | (^(v2)))}
    switch ("hello" + ("hello" + "foo bar")) {
    case "hello":
        switch "\n\t" {
        case "a":
            v25 := GenS1{F0: v4, F1: -1, F2: (v2 % (((v2 | v2)) | 1))}
            _ = v25
        case "":
            if false {
                v4 = (v4 + v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v26 uint8 = 42
                v2 = 115
                if true {
                    var v27 float32 = (1.5 * (0.0 + (1.5 / 1.0)))
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if false {
                            for fi29 := 0; fi29 < 3; fi29++ {
                                if true {
                                    var v30 uint32 = 0
                                    switch ((("a" + "a") + "hello") + "foo bar") {
                                    case "foo bar":
                                        v26 -= 128
                                    default:
                                        var v31 int8 = 0
                                        _ = v31
                                    }
                                    _ = v30
                                } else {
                                    v26 = v4
                                    v26 = 128
                                }
                                var v32 float32 = (v27 - v27)
                                v32 = v32
                                fmt.Println(v4)
                                if false {
                                    var v33 uint = 0
                                    _ = v33
                                }
                                _ = fi29
                                _ = v32
                            }
                            var v34 uint = 400
                            _ = v34
                        }
                        v26 = v4
                        var v35 float64 = -2.25
                        _ = fi28
                        _ = v35
                    }
                    if ("a" < "") {
                        fmt.Println(v4, v4, v27)
                        if false {
                            for fi36 := 0; fi36 < 2; fi36++ {
                                var v37 uint = 1
                                for fi38 := 0; fi38 < 10; fi38++ {
                                    fmt.Println(v37, v27)
                                    if (genFn1(("\n\t" + "\n\t"), false, 65535) < genFn1("\n\t", true, 0)) {
                                        fmt.Println(fi36, v2)
                                        var v39 float32 = v27
                                        _ = v39
                                    } else {
                                        switch (fi38 * fi36) {
                                        case 797:
                                            if (false && true) {
                                                var v40 uint16 = 65535
                                                _ = v40
                                            } else {
                                                v41 := GenS1{F0: v26, F1: 42, F2: (v2 >> ((149) & 63))}
                                                fmt.Println(v27, v27, v37)
                                                if true {
                                                    v42 := GenS1{F0: v26, F1: -884, F2: (^(v2))}
                                                    _ = v42
                                                }
                                                var v43 uint64 = v2
                                                _ = v41
                                                _ = v43
                                            }
                                        case 2147483647:
                                            v1 = append(v1, 4294967295)
                                        case -562:
                                            switch (fi38 + fi36) {
                                            case 9223372036854775807:
                                                if false {
                                                    switch -1 {
                                                    case 0:
                                                        var v44 string = ("" + "a")
                                                        _ = v44
                                                    default:
                                                        if (false && true) {
                                                            v45 := GenS2{F0: ("foo bar" < "\n\t"), F1: 427, F2: int32(v37)}
                                                            v46 := v3[185]
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = 2
                                                            }
                                                        }
                                                    }
                                                    var v47 int16 = -900
                                                    _ = v47
                                                } else {
                                                    if true {
                                                        var v48 uint16 = 0
                                                        var v49 float64 = 3.14159
                                                        var v50 int32 = -738
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if (true || true) {
                                                            switch ("a" + "hello") {
                                                            case "":
                                                                fmt.Println(v26)
                                                            case "a":
                                                                var v51 uint16 = 568
                                                                _ = v51
                                                            }
                                                            fmt.Println(v27)
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v2 += (v2 % (((v2 | 4294967295)) | 1))
                                                    case "foo bar":
                                                        v52 := GenS3{F0: ("hello" + "a"), F1: 32767}
                                                        _ = v52
                                                    case "a":
                                                        var v53 int16 = -32768
                                                        _ = v53
                                                    }
                                                    v1 = append(v1, 1)
                                                    v54 := GenS4{F0: 2147483647, F1: fi36}
                                                    _ = v54
                                                }
                                                if true {
                                                    v2 = v2
                                                } else {
                                                    v55 := GenS3{F0: "", F1: -32768}
                                                    _ = v55
                                                }
                                            case 32767:
                                                var v56 int64 = -1
                                                v57 := GenS4{F0: 2, F1: fi36}
                                                v58 := GenS5{F0: v37, F1: (!false)}
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            case -404:
                                                if (42 > 555) {
                                                    if (((^(v37)) != 40) || false) {
                                                        v27 = v27
                                                        var v59 string = ("foo bar" + "\n\t")
                                                        var v60 int8 = 0
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        fmt.Println(v26, v37, v27)
                                                        var v61 uint64 = 754
                                                        if false {
                                                            var v62 int16 = -691
                                                            v2 = (v61 / (((v2 + (v2 % (((v61 * 4294967295)) | 1)))) | 1))
                                                            var v63 int = fi38
                                                            v2 = v2
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        _ = v61
                                                    }
                                                    v2 = v2
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        default:
                                            v3[0] = (433 != genFn1("a", ((v37 ^ (v37 - 335)) < 18446744073709551615), 0))
                                        }
                                        _ = len(v1)
                                        v64 := GenS6{F0: 444}
                                        _ = v64
                                    }
                                    _ = fi38
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                    _ = v27
                }
                _ = v26
            }
        default:
            var v65 int32 = 750
            _ = v65
        }
    case "a":
        v3[2147483648] = true
    default:
        var v66 float64 = ((0.1 + 3.14159) / 0.1)
        _ = v66
    }
    if len(v1) > 0 {
        v1[0] = 445
    }
    v4 -= 1
    fmt.Println(v4, v4)
    _ = len(v1)
    for fi67 := 0; fi67 < 20; fi67++ {
        if false {
            v2 += v2
            fmt.Println(v2)
        } else {
            var v68 int8 = -23
            v69 := GenS7{F0: v2, F1: -1, F2: "\n\t"}
            switch (("\n\t" + ("hello" + "\n\t")) + "") {
            case "hello":
                if false {
                    v4 = 0
                    if true {
                        v4 = 255
                        var v70 uint = 4294967295
                        _ = v70
                    } else {
                        if false {
                            v68 = (v68 / (((v68 | 26)) | 1))
                            for fi71 := 0; fi71 < 13; fi71++ {
                                v72 := v3[1]
                                fmt.Println(v4, v4, v4)
                                v73 := GenS8{F0: (128 > uint8(v2))}
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                        } else {
                            for fi74 := 0; fi74 < 15; fi74++ {
                                var v75 int64 = 337
                                _ = fi74
                                _ = v75
                            }
                        }
                        if true {
                            v76 := GenS1{F0: (^(v4)), F1: -1, F2: 42}
                            fmt.Println(fi67)
                            for fi77 := 0; fi77 < 20; fi77++ {
                                if true {
                                    if true {
                                        var v78 float64 = 0.0
                                        var v79 int16 = -32768
                                        _ = v78
                                        _ = v79
                                    } else {
                                        fmt.Println(fi67, v4)
                                    }
                                    v80 := GenS5{F0: 18446744073709551615, F1: true}
                                    var v81 int = fi77
                                    _ = v80
                                    _ = v81
                                }
                                _ = len(v1)
                                v68 = (v68 & v68)
                                _ = fi77
                            }
                            _ = v76
                        }
                        if ((!(false || true)) || ((v68 % ((42) | 1)) < (v68 / ((v68) | 1)))) {
                            v68 = 127
                        }
                    }
                } else {
                    for fi82 := 0; fi82 < 9; fi82++ {
                        fmt.Println(v2, v68)
                        v68 += 100
                        var v83 string = ""
                        var v84 uint64 = 42
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    var v85 int64 = int64(v2)
                    for fi86 := 0; fi86 < 13; fi86++ {
                        if false {
                            var v87 uint8 = (v4 ^ (v4 % ((128) | 1)))
                            v4 = (v87 << ((1) & 63))
                            fmt.Println(fi67)
                            _ = v87
                        }
                        _ = fi86
                    }
                    _ = v85
                }
            default:
                for fi88 := 0; fi88 < 12; fi88++ {
                    switch "a" {
                    case "a":
                        if false {
                            v68 = int8(v4)
                            var v89 uint = 1
                            _ = v89
                        } else {
                            var v90 uint8 = v4
                            v68 += 127
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = 2
                                }
                                v91 := GenS9{F0: (-(v68)), F1: ""}
                                if (!true) {
                                    var v92 uint32 = 499
                                    if (fi88 == (fi67 + (fi67 & (fi67 % (((fi88 ^ (fi88 | int(v92)))) | 1))))) {
                                        delete(v3, 1)
                                        fmt.Println(v4, fi88)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v2 += v2
                                    var v93 float32 = (-2.25 + 1e10)
                                    if true {
                                        v90 = (^(v4))
                                        v94 := GenS10{F0: "\n\t", F1: (-(fi88)), F2: 34}
                                        v68 += (v68 * int8(v4))
                                        _ = v94
                                    }
                                    _ = v92
                                    _ = v93
                                } else {
                                    var v95 uint = 65535
                                    fmt.Println(fi67, v95)
                                    v90 = 127
                                    v96 := v3[0]
                                    fmt.Println(fi88)
                                    v1 = append(v1, 780)
                                    _ = v95
                                    _ = v96
                                }
                                _ = v91
                            }
                            var v97 float32 = 1.5
                            var v98 uint32 = 784
                            if false {
                                if (357 != 255) {
                                    v99 := GenS11{F0: 682, F1: -493}
                                    var v100 uint = 4294967295
                                    _ = v100
                                    _ = v99
                                } else {
                                    v68 = v68
                                }
                            } else {
                                switch ("hello" + ("" + "hello")) {
                                case "\n\t":
                                    v101 := GenS3{F0: "\n\t", F1: 32767}
                                    _ = v101
                                case "a":
                                    switch "\n\t" {
                                    case "":
                                        if true {
                                            if true {
                                                var v102 uint32 = 0
                                                if ((v97 == v97) && (394 < 0)) {
                                                    fmt.Println(v4, fi88)
                                                } else {
                                                    fmt.Println(fi67)
                                                    v1 = append(v1, (v102 >> ((0) & 63)))
                                                    fmt.Println(v90, v102)
                                                }
                                                _ = v102
                                            }
                                        }
                                    case "a":
                                        var v103 string = "a"
                                        _ = v103
                                    }
                                case "":
                                    v98 += v98
                                default:
                                    v104 := GenS9{F0: v68, F1: "foo bar"}
                                    _ = v104
                                }
                            }
                            _ = v90
                            _ = v97
                            _ = v98
                        }
                    }
                    _ = fi88
                }
            }
            _ = v68
            _ = v69
        }
        _ = fi67
    }
    if false {
        v4 = v4
    }
    v2 += uint64(v4)
    v4 += (v4 >> ((42) & 63))
    v105 := GenS12{F0: uint(v4)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v105
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
