package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 uint32) int32 {
    v4 := GenS1{F0: -861, F1: 92, F2: -2147483648}
    var v5 int = 4294967295
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint64 = 0
        v8 := GenS2{F0: 42, F1: 2, F2: false}
        switch ("\n\t" + (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))) {
        case "\n\t":
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint8 = 42
                v10 += v10
                _ = fi9
                _ = v10
            }
        case "hello":
            v11 := []int{fi6, (fi6 * 27), 125, fi6}
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: genFn0(65535), F1: -1, F2: 42}
                v7 += v7
                var v14 uint32 = 4294967295
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v15 := []int{127, 9223372036854775807, -914, 852}
    switch 42 {
    case -1:
        if true {
            switch 127 {
            case -809:
                var v16 string = "a"
                _ = v16
            case -2147483648:
                v17 := GenS3{F0: 744, F1: 52}
                _ = v17
            case -547:
                for fi18 := 0; fi18 < 11; fi18++ {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    v19 := GenS3{F0: 2, F1: 1}
                    switch "\n\t" {
                    case "a":
                        var v20 uint8 = 218
                        _ = v20
                    default:
                        var v21 int = 777
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
            }
        }
    case 42:
        v22 := GenS1{F0: 0, F1: -74, F2: -9223372036854775808}
        _ = v22
    default:
        var v23 int64 = -1
        _ = v23
    }
    var v24 float32 = 1e10
    _ = v15
    _ = v24
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{}
    v3 := []string{("\n\t" + "a"), "foo bar", "a", ("foo bar" + ("hello" + (("\n\t" + (("" + "a") + "foo bar")) + "\n\t")))}
    var v25 uint16 = 0
    var v26 uint16 = v25
    v1 = -2.25
    switch 42 {
    case 42:
        var v27 uint32 = 307
        _ = v27
    case 4294967295:
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := GenS1{F0: 297, F1: 1, F2: fi28}
            v1 += (v1 * (v1 * ((((v1 / v1) / v1) - ((float32(0.1) / v1) - ((v1 * 0.1) * ((v1 - ((v1 * (v1 * 1.5)) / (v1 / (v1 + 3.14159)))) + float32(0.0))))) * (v1 * v1))))
            var v30 int64 = -9223372036854775808
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    v25 -= (^(v26))
    if true {
        fmt.Println(v26, v1)
        var v31 uint = 641
        _ = v31
    }
    if (false || false) {
        v25 = v26
        if false {
            if len(v2) > 0 {
                v2[0] = -1
            }
            v26 += v26
            fmt.Println(v26, v1, v26)
        } else {
            v32 := GenS3{F0: 2, F1: -83}
            _ = v32
        }
    }
    if false {
        if (false || (468 <= 4294967295)) {
            var v33 int64 = 9223372036854775807
            v26 += (v26 >> ((912) & 63))
            _ = v33
        }
        v34 := GenS4{F0: (v25 | (v26 ^ (v26 << ((4294967295) & 63)))), F1: -702}
        v35 := GenS4{F0: v25, F1: -502}
        _ = v34
        _ = v35
    } else {
        fmt.Println(v25, v26)
        v36 := GenS5{F0: false}
        if true {
            switch -9223372036854775808 {
            case 128:
                switch 4294967295 {
                case 4294967295:
                    var v37 int8 = -1
                    _ = v37
                case 2147483647:
                    var v38 int16 = 32767
                    _ = v38
                case -1:
                    for fi39 := 0; fi39 < 17; fi39++ {
                        v1 -= v1
                        _ = fi39
                    }
                }
            case 42:
                if (!false) {
                    var v40 uint64 = 0
                    _ = v40
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch ("foo bar" + ("a" + ("a" + "foo bar"))) {
                    case "hello":
                        v41 := GenS5{F0: (!true)}
                        _ = v41
                    case "foo bar":
                        v42 := GenS4{F0: (^(v25)), F1: 9223372036854775807}
                        _ = v42
                    case "a":
                        if false {
                            var v43 uint64 = 0
                            v44 := GenS3{F0: 65535, F1: -128}
                            fmt.Println(v43)
                            _ = v43
                            _ = v44
                        } else {
                            _ = len(v3)
                            for fi45 := 0; fi45 < 12; fi45++ {
                                var v46 uint16 = v26
                                v1 += v1
                                _ = fi45
                                _ = v46
                            }
                        }
                    }
                    var v47 int16 = 0
                    _ = v47
                }
            }
            switch 2147483647 {
            case -2147483648:
                switch 32768 {
                case 714:
                    v25 += (^(v26))
                case 986:
                    var v48 bool = false
                    _ = v48
                case -9223372036854775808:
                    if true {
                        v49 := GenS3{F0: 2147483648, F1: 100}
                        _ = v49
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            }
            var v50 int32 = 1
            _ = v50
        } else {
            _ = len(v3)
            for fi51 := 0; fi51 < 2; fi51++ {
                v1 -= v1
                _ = fi51
            }
            for fi52 := 0; fi52 < 7; fi52++ {
                for fi53 := 0; fi53 < 9; fi53++ {
                    var v54 int32 = genFn0(0)
                    if ((!true) || (42 == 1)) {
                        fmt.Println(fi53)
                    } else {
                        if false {
                            if (0 > -1) {
                                if (-2147483648 == fi52) {
                                    if true {
                                        switch -116 {
                                        case -331:
                                            switch "foo bar" {
                                            case "":
                                                var v55 bool = (!false)
                                                _ = v55
                                            case "hello":
                                                var v56 float64 = float64((v1 / (0.0 - (float32(-2.25) + float32((0.1 - 3.14159))))))
                                                _ = v56
                                            default:
                                                if (!true) {
                                                    var v57 int32 = (^(v54))
                                                    var v58 float64 = 0.1
                                                    _ = v57
                                                    _ = v58
                                                }
                                                var v59 string = (("hello" + ("\n\t" + "a")) + "\n\t")
                                                _ = v59
                                            }
                                        case 4294967295:
                                            var v60 uint32 = 65535
                                            _ = v60
                                        default:
                                            _ = len(v3)
                                        }
                                    }
                                    if (127 <= 1) {
                                        fmt.Println(v1)
                                    } else {
                                        if ((3.14159 / 1.5) > -2.25) {
                                            v61 := GenS2{F0: 4294967295, F1: 2, F2: true}
                                            v1 -= (((v1 - v1) + 1.5) / v1)
                                            v62 := GenS3{F0: 65535, F1: -81}
                                            _ = v61
                                            _ = v62
                                        } else {
                                            fmt.Println(v54)
                                            if (-2.25 != 1e10) {
                                                var v63 uint = 0
                                                _ = v63
                                            } else {
                                                v64 := GenS6{F0: (v26 - v25)}
                                                fmt.Println(v26)
                                                _ = v64
                                            }
                                            if (v1 < v1) {
                                                _ = len(v3)
                                                v26 = v25
                                                var v65 uint32 = 537
                                                v1 -= v1
                                                v65 = v65
                                                v1 = v1
                                                v1 = v1
                                                v2 = append(v2, 32768)
                                                _ = v65
                                            }
                                            var v66 int32 = v54
                                            _ = v66
                                        }
                                        var v67 int16 = -742
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v67
                                    }
                                }
                                fmt.Println(fi53)
                                var v68 float32 = v1
                                var v69 uint32 = 42
                                _ = v68
                                _ = v69
                            }
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            _ = len(v2)
                            if (520 < 234) {
                                switch fi53 {
                                case 32767:
                                    v3 = append(v3, ((("hello" + "a") + "hello") + "a"))
                                }
                                var v70 bool = true
                                switch 32768 {
                                case 721:
                                    var v71 string = ("a" + "hello")
                                    _ = v71
                                case 127:
                                    var v72 int32 = v54
                                    _ = v72
                                case -784:
                                    if ("foo bar" != ("a" + ("" + ("foo bar" + (("" + "\n\t") + ("\n\t" + "")))))) {
                                        v25 = v25
                                        var v73 uint64 = 2
                                        var v74 float64 = -2.25
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 float32 = v1
                                        var v76 float32 = (float32(0.1) - float32(1e10))
                                        fmt.Println(v25, v25, v75)
                                        _ = len(v3)
                                        _ = v75
                                        _ = v76
                                    }
                                default:
                                    switch fi52 {
                                    case -872:
                                        if ((0.0 != 1e10) || v70) {
                                            if (fi52 < fi52) {
                                                switch ("hello" + "foo bar") {
                                                case "":
                                                    if ((((((-2.25 - v1) / 3.14159) + v1) / 0.1) * -2.25) <= ((v1 * ((((float32((((float64((3.14159 * 1.5)) + ((3.14159 + 1.5) - 0.0)) / (1.5 + (1e10 * (0.1 / float64((3.14159 - 0.1)))))) + float64((v1 / ((1e10 + -2.25) * -2.25))))) * (v1 - ((((1e10 / v1) + 0.0) - ((v1 - v1) + 1.5)) * v1))) / 1e10) + (((v1 / v1) * v1) - ((v1 + v1) / v1))) * (v1 * 3.14159))) - -2.25)) {
                                                        if false {
                                                            var v77 int64 = 0
                                                            var v78 uint16 = v26
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    }
                                                case "hello":
                                                    var v79 uint = 0
                                                    _ = v79
                                                case "\n\t":
                                                    v3 = append(v3, "")
                                                default:
                                                    var v80 uint16 = 287
                                                    _ = v80
                                                }
                                            } else {
                                                var v81 int16 = 763
                                                _ = v81
                                            }
                                            var v82 uint8 = 128
                                            _ = v82
                                        }
                                    case 2147483647:
                                        var v83 uint64 = 2
                                        var v84 uint8 = 42
                                        _ = v83
                                        _ = v84
                                    case -1:
                                        if v70 {
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                        }
                                    default:
                                        if true {
                                            v25 -= v25
                                        }
                                    }
                                }
                                _ = v70
                            } else {
                                if false {
                                    fmt.Println(fi53, v1, v25)
                                }
                                v54 += v54
                            }
                        }
                        v54 = v54
                        switch (-(fi52)) {
                        case 2147483647:
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                        case -9223372036854775808:
                            switch "hello" {
                            case "foo bar":
                                v85 := GenS6{F0: 2}
                                _ = v85
                            case "a":
                                if false {
                                    v54 += (v54 >> ((42) & 63))
                                }
                            case "":
                                fmt.Println(fi53, v1, v54)
                            default:
                                var v86 uint16 = 0
                                _ = v86
                            }
                        case 187:
                            var v87 bool = (false || (("hello" + ((("\n\t" + "foo bar") + ("a" + "foo bar")) + "\n\t")) > "a"))
                            _ = v87
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi53 * fi53)
                    }
                    _ = fi53
                    _ = v54
                }
                _ = fi52
            }
        }
        _ = v36
    }
    for fi88 := 0; fi88 < 13; fi88++ {
        fmt.Println(fi88)
        var v89 uint16 = 1
        _ = fi88
        _ = v89
    }
    v90 := GenS4{F0: (v25 + 32767), F1: -1}
    _ = len(v3)
    v91 := GenS3{F0: 65535, F1: 127}
    fmt.Println(v1)
    for fi92 := 0; fi92 < 5; fi92++ {
        v93 := GenS2{F0: 107, F1: 2147483647, F2: (true || (!true))}
        if len(v2) > 0 {
            v2[0] = fi92
        }
        var v94 uint16 = v25
        _ = fi92
        _ = v93
        _ = v94
    }
    if ((v25 + (v26 << ((4294967295) & 63))) >= 720) {
        switch ("a" + "\n\t") {
        case "foo bar":
            fmt.Println(v1, v25)
        case "a":
            switch 32768 {
            case 1:
                switch 949 {
                case 652:
                    if (genFn0(230) >= 1) {
                        v95 := GenS2{F0: 540, F1: 47, F2: true}
                        v26 += (v25 & v26)
                        _ = v95
                    }
                default:
                    fmt.Println(v26, v25)
                }
            case -476:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
        case "hello":
            for fi96 := 0; fi96 < 3; fi96++ {
                v97 := GenS6{F0: v26}
                fmt.Println(v1, v26)
                var v98 uint = 255
                _ = fi96
                _ = v97
                _ = v98
            }
        }
        fmt.Println(v1)
        switch 4294967295 {
        case 42:
            fmt.Println(v26)
        default:
            for fi99 := 0; fi99 < 16; fi99++ {
                v100 := GenS7{F0: 0, F1: 257}
                if (-2147483648 <= 437) {
                    v101 := GenS4{F0: (v26 ^ v25), F1: -821}
                    _ = v101
                }
                v2 = append(v2, 32768)
                _ = fi99
                _ = v100
            }
        }
    }
    var v102 float64 = -2.25
    var v103 uint64 = 42
    switch "" {
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if false {
            v104 := GenS8{F0: 544, F1: "\n\t", F2: "hello"}
            fmt.Println(v103, v102)
            _ = v104
        } else {
            for fi105 := 0; fi105 < 14; fi105++ {
                fmt.Println(v102, v26)
                _ = fi105
            }
        }
    default:
        v25 = (v25 + (^(v25)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v90
    _ = v91
}
