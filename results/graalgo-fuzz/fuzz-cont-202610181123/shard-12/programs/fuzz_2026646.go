package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 uint; F1 string; F2 bool }
type GenS12 struct { F0 uint64; F1 int64; F2 int }
type GenS13 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS14 struct { F0 uint; F1 int16; F2 uint }
type GenS15 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    v4 := []string{("" + "foo bar")}
    var v5 bool = false
    v6 := make(map[int]bool)
    v5 = true
    var v7 uint16 = 563
    fmt.Println(v3)
    v8 := GenS1{F0: ("hello" + "hello")}
    if (29 == 20) {
        fmt.Println(v3)
        var v9 uint64 = 172
        _ = v9
    } else {
        var v10 int8 = 0
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = v10
    }
    switch 277 {
    case 1:
        if v5 {
            switch -386 {
            case -9223372036854775808:
                if (1 >= 42) {
                    v5 = false
                    if false {
                        v11 := GenS2{F0: (v3 * uint32(v7)), F1: (v7 << ((18446744073709551615) & 63))}
                        var v12 bool = true
                        switch ("a" + "foo bar") {
                        case "a":
                            for fi13 := 0; fi13 < 10; fi13++ {
                                for fi14 := 0; fi14 < 3; fi14++ {
                                    v15 := GenS1{F0: "a"}
                                    _ = fi14
                                    _ = v15
                                }
                                var v16 int8 = -41
                                _ = fi13
                                _ = v16
                            }
                        }
                        _ = v11
                        _ = v12
                    } else {
                        switch "a" {
                        case "":
                            _ = len(v1)
                        default:
                            var v17 int32 = -2147483648
                            _ = v17
                        }
                        if (-1 > int(v3)) {
                            if v5 {
                                v18 := GenS1{F0: "hello"}
                                _ = v18
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        var v19 uint64 = 39
                        _ = v19
                    }
                }
            default:
                v20 := GenS3{F0: 2, F1: -9223372036854775808}
                _ = v20
            }
        }
    case 127:
        if len(v4) > 0 {
            v4[0] = ("foo bar" + ("\n\t" + "hello"))
        }
    default:
        for fi21 := 0; fi21 < 15; fi21++ {
            if v5 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v3 = (v3 ^ (v3 / ((v3) | 1)))
                v7 -= (v7 / (((v7 >> ((4294967295) & 63))) | 1))
            }
            if (!(42 < 1)) {
                fmt.Println(v7, v3)
                if (1e10 == (0.1 / 1.0)) {
                    var v22 int = fi21
                    _ = v22
                }
            } else {
                v23 := GenS2{F0: 368, F1: v7}
                delete(v6, (fi21 >> ((42) & 63)))
                v24 := GenS4{F0: fi21, F1: 1, F2: (v5 || false)}
                _ = v23
                _ = v24
            }
            fmt.Println(v5, v5, v5)
            _ = fi21
        }
    }
    switch ("hello" + "hello") {
    case "hello":
        v5 = (1e10 >= -2.25)
    case "a":
        var v25 uint32 = (v3 >> ((uint(v3)) & 63))
        _ = v25
    default:
        if (2 < 210) {
            if ((1 >= 2) || (((!(true || v5)) && v5) || (!(0 == 18446744073709551615)))) {
                v5 = false
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 uint64 = 42
                    _ = fi26
                    _ = v27
                }
                var v28 uint16 = 972
                _ = v28
            } else {
                v29 := GenS5{F0: 32767, F1: 2147483647}
                _ = v29
            }
            v30 := GenS6{F0: "", F1: (v7 | v7)}
            _ = v30
        }
    }
    var v31 uint8 = 2
    v32 := GenS6{F0: "a", F1: v7}
    if (v3 == 2) {
        if (1.5 < 1e10) {
            if ((v7 & 255) < v7) {
                switch "" {
                case "":
                    for fi33 := 0; fi33 < 5; fi33++ {
                        var v34 bool = false
                        fmt.Println(v5, v34)
                        switch "foo bar" {
                        case "hello":
                            v34 = v5
                        }
                        _ = fi33
                        _ = v34
                    }
                case "foo bar":
                    v35 := GenS7{F0: 1, F1: 589}
                    _ = v35
                }
                v31 = uint8(v7)
            }
            for fi36 := 0; fi36 < 9; fi36++ {
                if ("" != "hello") {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        var v38 string = "hello"
                        _ = fi37
                        _ = v38
                    }
                }
                switch 1 {
                case -1:
                    v3 -= v3
                case 350:
                    var v39 int16 = 32767
                    _ = v39
                }
                _ = fi36
            }
            v40 := GenS1{F0: "foo bar"}
            _ = v40
        } else {
            var v41 uint32 = v3
            _ = v41
        }
    }
    if (-2.25 < 1e10) {
        var v42 float32 = 0.1
        v31 -= v31
        _ = v42
    }
    v5 = (812 <= 1)
    v7 = v7
    var v43 bool = v5
    if (32767 >= 0) {
        for fi44 := 0; fi44 < 6; fi44++ {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            v31 = v31
            _ = fi44
        }
    } else {
        if false {
            var v45 uint16 = 0
            for fi46 := 0; fi46 < 2; fi46++ {
                for fi47 := 0; fi47 < 11; fi47++ {
                    if true {
                        v48 := GenS8{F0: 65535}
                        _ = v48
                    } else {
                        if (v5 || (896 < 4294967295)) {
                            switch "\n\t" {
                            case "a":
                                v43 = (42 >= 255)
                            case "foo bar":
                                fmt.Println(v7, v45, v31)
                            }
                            v3 -= (v3 + (v3 << ((42) & 63)))
                            var v49 uint = 212
                            delete(v1, v3)
                            var v50 uint16 = 42
                            _ = v49
                            _ = v50
                        }
                        v4 = append(v4, ("foo bar" + ("\n\t" + "")))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v5 {
                            if (1 < -1) {
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                if v43 {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch (fi47 >> ((346) & 63)) {
                                    case 1:
                                        delete(v1, (v3 + 2147483647))
                                        fmt.Println(fi46, v31, v45)
                                        v7 += uint16(v31)
                                    case 444:
                                        fmt.Println(v45, fi47)
                                    }
                                }
                                v7 = v45
                                if v43 {
                                    if (743 <= 2) {
                                        v51 := GenS9{F0: int8(v3), F1: ("\n\t" + "foo bar")}
                                        if true {
                                            v7 += v45
                                            fmt.Println(v3)
                                        } else {
                                            v7 = v7
                                        }
                                        if true {
                                            var v52 string = "\n\t"
                                            var v53 uint16 = (v7 << ((27) & 63))
                                            v6[(fi47 >> ((65535) & 63))] = (42 == 0)
                                            _ = v52
                                            _ = v53
                                        } else {
                                            switch ("" + "hello") {
                                            case "\n\t":
                                                v54 := GenS10{F0: 42, F1: 42}
                                                _ = v54
                                            case "a":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "":
                                                v3 = (v3 << ((18446744073709551615) & 63))
                                            default:
                                                v5 = (!true)
                                            }
                                            v6[(-(fi47))] = false
                                            if ((0.0 + 0.1) < (float32(3.14159) - 3.14159)) {
                                                v55 := GenS10{F0: 0, F1: 435}
                                                _ = v55
                                            } else {
                                                var v56 uint64 = 42
                                                _ = v56
                                            }
                                        }
                                        v45 += v45
                                        _ = v51
                                    } else {
                                        if v5 {
                                            var v57 int32 = -1
                                            _ = v57
                                        } else {
                                            v45 += (v45 - (^(v7)))
                                            switch ("" + ("a" + "hello")) {
                                            case "a":
                                                v58 := GenS6{F0: "hello", F1: (v7 | v7)}
                                                _ = v58
                                            case "hello":
                                                v59 := GenS5{F0: 42, F1: -805}
                                                _ = v59
                                            default:
                                                v60 := GenS11{F0: 1, F1: "", F2: true}
                                                _ = v60
                                            }
                                            if v43 {
                                                fmt.Println(v31)
                                                var v61 uint32 = v3
                                                if (1 != -310) {
                                                    if false {
                                                        v61 += 3
                                                        if (-32768 != -89) {
                                                            var v62 int32 = -2147483648
                                                            var v63 int = fi47
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            var v64 int64 = 42
                                                            var v65 float64 = -2.25
                                                            var v66 string = "\n\t"
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        var v67 uint = 164
                                                        _ = v67
                                                    }
                                                    var v68 bool = (32767 == fi47)
                                                    v3 -= (v3 * (v3 & v61))
                                                    _ = v68
                                                } else {
                                                    fmt.Println(fi47, v31)
                                                }
                                                _ = v61
                                            }
                                            if (-55 < -103) {
                                                var v69 uint16 = v7
                                                fmt.Println(fi47, v31)
                                                _ = v69
                                            }
                                        }
                                        var v70 int64 = 0
                                        _ = v70
                                    }
                                }
                                switch (fi46 & fi47) {
                                case 1:
                                    v43 = (((^(v31)) <= 119) && false)
                                case 313:
                                    switch (fi47 >> ((uint(v31)) & 63)) {
                                    case 1:
                                        v71 := v2[(v3 & (v3 + uint32(v31)))]
                                        _ = v71
                                    case -9223372036854775808:
                                        fmt.Println(v7, fi47, fi47)
                                    default:
                                        v4 = append(v4, ("foo bar" + (("a" + "a") + ("hello" + (("hello" + ("" + ("a" + ("" + ("hello" + "foo bar"))))) + ("\n\t" + ""))))))
                                        var v72 uint32 = v3
                                        var v73 float32 = (-2.25 * -2.25)
                                        var v74 uint16 = v7
                                        v31 += (v31 / (((v31 * v31)) | 1))
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                case -2147483648:
                                    var v75 int = int(v45)
                                    _ = v75
                                }
                            } else {
                                if false {
                                    fmt.Println(v31, v31, v7)
                                    var v76 string = (("" + (((("a" + "\n\t") + (("\n\t" + "\n\t") + "hello")) + "\n\t") + (("foo bar" + "hello") + "hello"))) + (("a" + "\n\t") + "foo bar"))
                                    v77 := GenS12{F0: 0, F1: 42, F2: (fi47 << ((18446744073709551615) & 63))}
                                    _ = v76
                                    _ = v77
                                } else {
                                    switch fi46 {
                                    case -593:
                                        v78 := GenS3{F0: uint16(v3), F1: -1}
                                        _ = v78
                                    }
                                    v45 += (v7 & v7)
                                    v79 := GenS13{F0: 722, F1: uint64(v31), F2: 1}
                                    _ = v79
                                }
                            }
                            switch "a" {
                            case "":
                                fmt.Println(v31)
                            case "a":
                                v80 := GenS3{F0: v45, F1: -9223372036854775808}
                                _ = v80
                            case "hello":
                                v45 -= v45
                                var v81 int64 = 132
                                var v82 int16 = -90
                                if true {
                                    v83 := GenS13{F0: 18446744073709551615, F1: 311, F2: (v82 % (((v82 & -716)) | 1))}
                                    var v84 int = int(v3)
                                    if (v5 && ((v3 ^ 1) != (v3 | 1))) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        var v85 uint = 0
                                        switch (("hello" + (("\n\t" + "") + "a")) + "a") {
                                        case "\n\t":
                                            v86 := GenS14{F0: 0, F1: 1, F2: v85}
                                            v87 := v2[v3]
                                            switch ("" + ("a" + ("foo bar" + "\n\t"))) {
                                            case "":
                                                var v88 uint8 = v31
                                                _ = v88
                                            }
                                            if (!v5) {
                                                var v89 float64 = 0.0
                                                _ = v89
                                            }
                                            v6[fi46] = v87
                                            var v90 uint16 = 136
                                            if (1.5 > ((0.1 + float32(1.5)) + 0.1)) {
                                                v45 = 32767
                                            } else {
                                                var v91 uint64 = 18446744073709551615
                                                if false {
                                                    var v92 uint16 = 65535
                                                    var v93 uint32 = 971
                                                    _ = v92
                                                    _ = v93
                                                }
                                                _ = v91
                                            }
                                            var v94 uint8 = 0
                                            var v95 uint32 = 0
                                            _ = v86
                                            _ = v87
                                            _ = v90
                                            _ = v94
                                            _ = v95
                                        default:
                                            fmt.Println(v81, fi46, v84)
                                        }
                                        fmt.Println(v82, v81, v84)
                                        v7 = (v7 / ((v7) | 1))
                                        _ = v85
                                    }
                                    if v5 {
                                        fmt.Println(v84)
                                        v7 = 32767
                                        v45 -= (v45 | v45)
                                    } else {
                                        v3 += (v3 - v3)
                                        fmt.Println(v5, v82)
                                        v6[fi47] = (((v82 ^ v82) == (v82 | int16(fi47))) || (!v43))
                                        switch (v84 << ((2) & 63)) {
                                        case 32768:
                                            v84 -= (v84 - (^(v84)))
                                        case 4294967295:
                                            v84 += int(v3)
                                        }
                                    }
                                    var v96 uint64 = 484
                                    v7 -= (v7 + (v45 << ((65535) & 63)))
                                    var v97 uint64 = v96
                                    _ = v83
                                    _ = v84
                                    _ = v96
                                    _ = v97
                                }
                                v43 = v5
                                _ = v81
                                _ = v82
                            }
                        } else {
                            var v98 uint64 = 18446744073709551615
                            _ = v98
                        }
                        v45 = (v7 << ((1) & 63))
                        var v99 bool = v43
                        _ = v99
                    }
                    v31 = v31
                    var v100 int16 = -824
                    _ = fi47
                    _ = v100
                }
                v101 := GenS15{F0: 539, F1: v5}
                _ = fi46
                _ = v101
            }
            v1[v3] = 2
            _ = v45
        }
        fmt.Println(v43, v7)
        var v102 float64 = -2.25
        _ = v102
    }
    fmt.Println(v3)
    var v103 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
