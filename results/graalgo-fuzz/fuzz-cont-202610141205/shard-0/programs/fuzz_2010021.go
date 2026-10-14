package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint16; F1 int16; F2 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint32 }

func main() {
    v1 := []int{670, 42, 82, 0}
    var v2 uint = 495
    var v3 uint16 = 0
    v3 = v3
    if (true || false) {
        v4 := GenS1{F0: 1, F1: 255}
        if (false && ((9223372036854775807 >= 42) || true)) {
            var v5 uint = 18446744073709551615
            _ = v5
        } else {
            v2 -= v2
        }
        v1 = append(v1, -2147483648)
        _ = v4
    } else {
        v2 = v2
    }
    v6 := GenS1{F0: 4294967295, F1: v3}
    switch "a" {
    case "a":
        var v7 uint32 = uint32(v2)
        _ = v7
    case "\n\t":
        v3 -= v3
    case "hello":
        for fi8 := 0; fi8 < 13; fi8++ {
            if len(v1) > 0 {
                v1[0] = (fi8 << ((v2) & 63))
            }
            _ = fi8
        }
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            if (0.1 < ((0.1 / 0.1) / 1.0)) {
                var v10 uint8 = 128
                _ = v10
            } else {
                v3 = v3
                fmt.Println(fi9)
                var v11 uint16 = 42
                _ = v11
            }
            fmt.Println(v2, v3)
            v12 := GenS1{F0: 1, F1: (v3 << (((v2 << ((255) & 63))) & 63))}
            _ = fi9
            _ = v12
        }
    }
    fmt.Println(v3, v2)
    if (!(42 == 985)) {
        var v13 int32 = 0
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 uint16 = v3
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v2, fi16)
        v17 := GenS2{F0: 4294967295, F1: 32}
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 134
    var v19 uint = (v2 * v2)
    if true {
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 10; fi21++ {
                v1 = append(v1, 128)
                _ = fi21
            }
            _ = fi20
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 string = "hello"
            switch ((v23 + "a") + (((v23 + "foo bar") + v23) + ("hello" + ("" + v23)))) {
            case "foo bar":
                v18 -= v18
            case "":
                v3 = (v3 + 633)
            }
            _ = fi22
            _ = v23
        }
        if ((!(623 != -9223372036854775808)) || ((-2.25 * 0.1) >= (0.1 * (1.5 / 3.14159)))) {
            var v24 string = "foo bar"
            switch "a" {
            case "hello":
                for fi25 := 0; fi25 < 15; fi25++ {
                    v26 := GenS1{F0: 42, F1: 32767}
                    _ = fi25
                    _ = v26
                }
            case "":
                fmt.Println(v19, v3, v18)
            case "foo bar":
                var v27 float64 = float64(float32((1e10 + ((((1e10 + 0.1) / 1e10) + float64((3.14159 / 1e10))) / 0.1))))
                _ = v27
            default:
                if (0 != 1) {
                    var v28 int8 = 1
                    _ = v28
                } else {
                    v1 = append(v1, -710)
                    var v29 bool = (true || (2 > uint32(v2)))
                    var v30 int16 = -890
                    _ = v29
                    _ = v30
                }
            }
            _ = v24
        } else {
            var v31 uint = 0
            _ = v31
        }
    } else {
        fmt.Println(v18)
        v3 += 65535
        var v32 uint64 = 2
        _ = v32
    }
    var v33 uint32 = 42
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 string = ("\n\t" + ("\n\t" + "foo bar"))
        v36 := GenS3{F0: 4294967295, F1: 9223372036854775807}
        _ = fi34
        _ = v35
        _ = v36
    }
    if false {
        v37 := GenS3{F0: 42, F1: 800}
        v38 := GenS2{F0: 4294967295, F1: -24}
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 uint32 = 1
            v41 := GenS4{F0: -118}
            if (false && false) {
                var v42 int64 = 0
                v43 := GenS5{F0: 42}
                v44 := GenS6{F0: 42, F1: -1}
                _ = v42
                _ = v43
                _ = v44
            } else {
                var v45 uint16 = (v3 ^ v3)
                for fi46 := 0; fi46 < 1; fi46++ {
                    switch 42 {
                    case 0:
                        v40 -= (v33 & v40)
                    case -510:
                        v33 = v40
                    case 42:
                        v47 := GenS5{F0: fi46}
                        if len(v1) > 0 {
                            v1[0] = fi46
                        }
                        v48 := GenS6{F0: v2, F1: -32768}
                        _ = v47
                        _ = v48
                    default:
                        if true {
                            var v49 int16 = 207
                            if false {
                                if false {
                                    var v50 uint = (v2 % ((v19) | 1))
                                    switch "foo bar" {
                                    case "a":
                                        var v51 int64 = -1
                                        v50 = v2
                                        _ = v51
                                    case "\n\t":
                                        v19 -= (v50 - v19)
                                    case "hello":
                                        switch fi39 {
                                        case -496:
                                            switch (fi46 << (((v2 | (v19 % ((v2) | 1)))) & 63)) {
                                            case 1:
                                                v33 = v40
                                            case 128:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (false || ((false || false) && false)) {
                                                    v52 := GenS5{F0: -1}
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v53 int32 = 42
                                                        _ = v53
                                                    case "\n\t":
                                                        v54 := GenS2{F0: 4294967295, F1: 0}
                                                        _ = v54
                                                    }
                                                    if (100 >= 100) {
                                                        if true {
                                                            var v55 uint32 = (v40 - 2147483648)
                                                            v50 -= (v2 << ((v50) & 63))
                                                            _ = v55
                                                        } else {
                                                            var v56 uint8 = v18
                                                            _ = v56
                                                        }
                                                        var v57 float64 = (0.0 / (1e10 * float64(-2.25)))
                                                        _ = v57
                                                    }
                                                    _ = v52
                                                }
                                                fmt.Println(v19, fi46, v40)
                                            default:
                                                v33 = (^(v33))
                                            }
                                        case 4294967295:
                                            if true {
                                                fmt.Println(v19, v50)
                                            } else {
                                                v58 := GenS7{F0: v3, F1: 1}
                                                var v59 uint32 = v40
                                                _ = v58
                                                _ = v59
                                            }
                                        default:
                                            if (false || ((^(v45)) < v45)) {
                                                if true {
                                                    if (false && false) {
                                                        switch fi46 {
                                                        case 0:
                                                            var v60 float64 = ((1e10 + -2.25) / 0.1)
                                                            v61 := GenS8{F0: 255, F1: v49, F2: 1}
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    }
                                                } else {
                                                    if (0 < 1) {
                                                        if true {
                                                            if (!false) {
                                                                v1 = append(v1, fi46)
                                                                var v62 bool = true
                                                                v40 = (v33 % (((v33 >> ((65535) & 63))) | 1))
                                                                v3 = (v45 % (((v3 >> ((v19) & 63))) | 1))
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if (0 <= 1) {
                                                                    var v63 uint = v19
                                                                    _ = v63
                                                                }
                                                                v64 := GenS9{F0: 0}
                                                                _ = v62
                                                                _ = v64
                                                            } else {
                                                                v33 = (v33 % ((v40) | 1))
                                                                switch ("" + ((("\n\t" + "a") + "hello") + "foo bar")) {
                                                                case "foo bar":
                                                                    var v65 uint16 = (v3 * v45)
                                                                    _ = v65
                                                                case "hello":
                                                                    v19 += v2
                                                                case "a":
                                                                    v50 = 161
                                                                }
                                                                var v66 int32 = -169
                                                                fmt.Println(v33, v49, v66)
                                                                if false {
                                                                    fmt.Println(v33, v40)
                                                                } else {
                                                                    var v67 uint8 = (v18 / ((uint8(v2)) | 1))
                                                                    _ = v67
                                                                }
                                                                _ = v66
                                                            }
                                                            var v68 uint64 = 4294967295
                                                            v69 := GenS5{F0: (fi39 + (fi39 << ((uint(v45)) & 63)))}
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        v19 -= (v19 ^ v50)
                                                        v49 = (v49 >> ((637) & 63))
                                                    }
                                                    if (false || true) {
                                                        var v70 uint64 = 2
                                                        v49 = (v49 & v49)
                                                        v71 := GenS10{F0: 1, F1: "\n\t"}
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        var v72 uint16 = (v3 ^ v45)
                                                        var v73 uint8 = (v18 / ((uint8(v40)) | 1))
                                                        var v74 uint = v19
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                }
                                                switch fi39 {
                                                case 267:
                                                    var v75 int64 = 210
                                                    _ = v75
                                                case 9223372036854775807:
                                                    v50 = v19
                                                }
                                                var v76 int8 = 82
                                                _ = v76
                                            }
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            var v77 bool = true
                                            _ = v77
                                        }
                                        var v78 bool = true
                                        switch "a" {
                                        case "foo bar":
                                            v19 += (v50 << ((v50) & 63))
                                        case "\n\t":
                                            fmt.Println(fi46, v50, v50)
                                        }
                                        var v79 string = ""
                                        _ = v78
                                        _ = v79
                                    }
                                    if (100 >= 0) {
                                        v2 += v2
                                    }
                                    switch (^(fi39)) {
                                    case 128:
                                        v80 := GenS11{F0: (v33 <= 65535), F1: uint64(v18), F2: (v3 ^ (v3 | v45))}
                                        _ = v80
                                    case -824:
                                        v49 += -802
                                        v49 = (v49 & (v49 % ((v49) | 1)))
                                        v81 := GenS12{F0: (v3 + (v3 << (((v19 ^ (v50 * 1))) & 63)))}
                                        _ = v81
                                    case 32768:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v82 uint16 = v45
                                        _ = v82
                                    }
                                    _ = v50
                                }
                                v19 += uint(v33)
                            } else {
                                v49 = v49
                            }
                            _ = v49
                        }
                    }
                    var v83 uint8 = v18
                    _ = fi46
                    _ = v83
                }
                _ = v45
            }
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v37
        _ = v38
    } else {
        v2 = v2
        v84 := GenS13{F0: v33}
        _ = v84
    }
    _ = len(v1)
    if ((3.14159 == (0.0 + 1.5)) && true) {
        if (!true) {
            fmt.Println(v19, v33, v2)
            var v85 bool = true
            _ = v85
        } else {
            if (((0.0 * float32((((((0.0 - (-2.25 - float64(float32(3.14159)))) * 0.0) - 1e10) - (0.0 * 1.5)) * float64(3.14159)))) * ((-2.25 / (1.5 * float32(3.14159))) * (float32(0.1) / -2.25))) < 0.1) {
                for fi86 := 0; fi86 < 17; fi86++ {
                    if (true || true) {
                        var v87 uint = v2
                        v19 = v19
                        var v88 float32 = 1.5
                        _ = v87
                        _ = v88
                    } else {
                        if (-25 > 0) {
                            v33 = (^(v33))
                        } else {
                            var v89 bool = (((v2 & (v2 % ((65535) | 1))) > v2) && (false && true))
                            _ = v89
                        }
                    }
                    v1 = append(v1, fi86)
                    fmt.Println(v2, v3)
                    _ = fi86
                }
                fmt.Println(v19)
                switch -9223372036854775808 {
                case -332:
                    if false {
                        if false {
                            switch ("a" + "\n\t") {
                            case "a":
                                v19 = v2
                            case "foo bar":
                                if true {
                                    for fi90 := 0; fi90 < 4; fi90++ {
                                        v2 += uint(v3)
                                        v19 = (v2 % ((v2) | 1))
                                        fmt.Println(fi90)
                                        _ = fi90
                                    }
                                    for fi91 := 0; fi91 < 16; fi91++ {
                                        var v92 int = (fi91 + fi91)
                                        if (9223372036854775807 >= v92) {
                                            for fi93 := 0; fi93 < 15; fi93++ {
                                                if false {
                                                    var v94 float32 = 0.0
                                                    _ = v94
                                                }
                                                var v95 int = fi93
                                                v96 := GenS10{F0: -128, F1: ""}
                                                if (v33 != (v33 * (v33 & v33))) {
                                                    switch "" {
                                                    case "hello":
                                                        v92 -= fi93
                                                    case "\n\t":
                                                        v19 -= (v19 + v2)
                                                    case "foo bar":
                                                        v97 := GenS5{F0: (v92 * (v95 - v92))}
                                                        _ = v97
                                                    }
                                                    v3 += 0
                                                    var v98 uint = v2
                                                    _ = v98
                                                } else {
                                                    v33 = uint32(v3)
                                                    var v99 uint32 = 2
                                                    _ = v99
                                                }
                                                _ = fi93
                                                _ = v95
                                                _ = v96
                                            }
                                            var v100 bool = (!false)
                                            _ = v100
                                        }
                                        _ = fi91
                                        _ = v92
                                    }
                                } else {
                                    var v101 uint64 = 916
                                    _ = v101
                                }
                            }
                            var v102 int = -2147483648
                            _ = len(v1)
                            _ = v102
                        } else {
                            var v103 int8 = -128
                            switch (("\n\t" + "") + "") {
                            case "hello":
                                _ = len(v1)
                            default:
                                var v104 uint8 = 128
                                _ = v104
                            }
                            _ = v103
                        }
                    }
                case 32768:
                    var v105 int64 = 27
                    _ = v105
                case 4294967295:
                    fmt.Println(v19)
                }
            }
            fmt.Println(v3, v3)
            v2 = uint(v3)
        }
    } else {
        var v106 int32 = 0
        _ = v106
    }
    v33 = 963
    if false {
        var v107 uint = 956
        fmt.Println(v18, v19, v19)
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v6
}
