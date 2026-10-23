package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 uint16, p1 int) int64 {
    if false {
        v7 := []int{-2147483648, p1, -916, (-(p1))}
        if false {
            for fi8 := 0; fi8 < 10; fi8++ {
                p0 += p0
                _ = fi8
            }
        }
        _ = v7
    }
    if ((true || ((^(p0)) > 32767)) && (124 > 0)) {
        for fi9 := 0; fi9 < 16; fi9++ {
            switch (p1 >> ((4294967295) & 63)) {
            case 238:
                v10 := GenS1{F0: 0, F1: 42}
                _ = v10
            default:
                if false {
                    if (true && (!false)) {
                        v11 := []int{p1, (fi9 / ((301) | 1)), (p1 ^ (fi9 - (p1 | 9223372036854775807)))}
                        v11 = append(v11, (fi9 + (fi9 >> ((255) & 63))))
                        _ = v11
                    }
                } else {
                    var v12 uint16 = p0
                    if false {
                        switch (fi9 | (-(p1))) {
                        case 9223372036854775807:
                            if false {
                                switch fi9 {
                                case 879:
                                    for fi13 := 0; fi13 < 13; fi13++ {
                                        p0 = 444
                                        var v14 bool = true
                                        _ = fi13
                                        _ = v14
                                    }
                                }
                            }
                        default:
                            p0 -= (p0 | (v12 & v12))
                        }
                    }
                    if false {
                        switch "hello" {
                        case "":
                            p0 -= (v12 / ((1) | 1))
                        case "foo bar":
                            v15 := []int{}
                            _ = v15
                        case "hello":
                            var v16 string = "hello"
                            _ = v16
                        default:
                            var v17 int = fi9
                            _ = v17
                        }
                    } else {
                        var v18 uint32 = 242
                        _ = v18
                    }
                    _ = v12
                }
            }
            fmt.Println(p1)
            _ = fi9
        }
    } else {
        v19 := GenS2{F0: 2, F1: 252, F2: 79}
        if false {
            var v20 int16 = 1000
            _ = v20
        }
        v21 := GenS1{F0: 0, F1: (^(p0))}
        _ = v19
        _ = v21
    }
    p1 += (p1 >> ((65535) & 63))
    if (-77 == -12) {
        switch int(p0) {
        case 32768:
            v22 := []int{42}
            _ = v22
        default:
            if (!false) {
                switch p1 {
                case -805:
                    switch 4294967295 {
                    case 128:
                        switch "" {
                        case "\n\t":
                            var v23 uint = 0
                            _ = v23
                        case "hello":
                            fmt.Println(p0)
                        case "":
                            v24 := []int{}
                            _ = v24
                        }
                    case 2147483647:
                        var v25 string = "foo bar"
                        _ = v25
                    case -2147483648:
                        p0 = 2
                    default:
                        var v26 uint32 = 42
                        _ = v26
                    }
                }
                var v27 float32 = float32(1e10)
                _ = v27
            }
        }
        p1 += p1
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 0
    v3 := []string{"", "foo bar"}
    v4 := make(map[string]string)
    var v5 uint = (^(v2))
    var v6 int64 = 0
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS3{F0: 435, F1: (^(fi28)), F2: genFn0(2, fi28)}
        v5 += (v5 - (v5 & (v5 | v5)))
        _ = fi28
        _ = v29
    }
    if false {
        var v30 int16 = 495
        fmt.Println(v2, v6)
        _ = v30
    } else {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 uint16 = 42
            fmt.Println(v2, v32)
            var v33 float32 = (1e10 * 1.5)
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println(v6, v6, v2)
    _ = len(v1)
    v6 = (v6 % (((v6 & (v6 << ((42) & 63)))) | 1))
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS2{F0: v2, F1: 253, F2: (v2 & v5)}
            if (false && (451 >= 255)) {
                var v37 int = (fi34 << ((v5) & 63))
                _ = v37
            } else {
                if (848 == 42) {
                    if (float64((0.1 * 1e10)) != -2.25) {
                        v3 = append(v3, "a")
                    }
                    var v38 int64 = v6
                    v39 := GenS4{F0: (true || true)}
                    var v40 int = 1
                    switch fi34 {
                    case 9223372036854775807:
                        var v41 int8 = -1
                        _ = v41
                    default:
                        v2 -= 18446744073709551615
                    }
                    _ = v38
                    _ = v39
                    _ = v40
                }
                v6 -= (v6 % ((v6) | 1))
                var v42 uint = (v5 + v5)
                _ = v42
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi35
            _ = v36
        }
        if true {
            for fi43 := 0; fi43 < 2; fi43++ {
                switch fi34 {
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -465:
                    var v44 bool = (0.0 == 1.5)
                    _ = v44
                }
                var v45 int = fi34
                v6 += genFn0(925, v45)
                _ = fi43
                _ = v45
            }
        } else {
            _ = len(v1)
            _ = len(v3)
        }
        _ = fi34
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        var v47 int32 = 848
        var v48 uint64 = 851
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 uint = v2
            switch ("" + "") {
            case "hello":
                v51 := v4["foo bar"]
                _ = v51
            case "":
                _ = len(v3)
                var v52 uint8 = 42
                v47 = v47
                if (false || (46 >= 105)) {
                    v47 -= (^(v47))
                    v50 += v50
                    var v53 uint32 = 963
                    fmt.Println(v48, v5)
                    var v54 string = "foo bar"
                    if false {
                        if true {
                            v53 += (v53 % (((v53 - 65535)) | 1))
                        }
                    } else {
                        switch (fi49 % (((fi49 << ((v50) & 63))) | 1)) {
                        case 0:
                            if len(v3) > 0 {
                                v3[0] = v54
                            }
                        case 894:
                            fmt.Println(fi49, fi46, fi49)
                        default:
                            var v55 int32 = 0
                            _ = v55
                        }
                        v56 := GenS5{F0: 29}
                        v3 = append(v3, (v54 + "hello"))
                        _ = v56
                    }
                    var v57 float32 = ((1.5 + float32(float64(1.5))) - 3.14159)
                    v58 := GenS5{F0: v50}
                    _ = v53
                    _ = v54
                    _ = v57
                    _ = v58
                } else {
                    fmt.Println(v52)
                    delete(v1, "\n\t")
                    var v59 uint = v2
                    _ = v59
                }
                _ = v52
            case "a":
                if (42 <= v6) {
                    _ = len(v1)
                    var v60 bool = (false && true)
                    v60 = v60
                    _ = v60
                } else {
                    if false {
                        if (((0 < 1000) || (!false)) || false) {
                            switch fi49 {
                            case 32767:
                                var v61 float32 = ((float32(0.1) + 0.1) * (0.0 / (float32(((1.5 - (0.0 * -2.25)) + 0.0)) / 1.5)))
                                _ = v61
                            case -648:
                                if (4294967295 == 802) {
                                    v50 = v2
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                            v50 += v5
                            var v62 int32 = -2147483648
                            _ = v62
                        } else {
                            v63 := GenS3{F0: 2147483647, F1: fi46, F2: (v6 / (((v6 - 9)) | 1))}
                            v64 := GenS2{F0: v2, F1: 255, F2: v2}
                            _ = v63
                            _ = v64
                        }
                        fmt.Println(v50)
                        if (fi49 != (fi49 ^ -2147483648)) {
                            v6 = v6
                            fmt.Println(v5, v5)
                        } else {
                            v5 += (v5 % (((v5 * v50)) | 1))
                            var v65 uint16 = 65535
                            _ = v65
                        }
                    }
                }
            }
            v5 += (v50 >> (((v2 * v5)) & 63))
            _ = fi49
            _ = v50
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    if (!false) {
        switch 4294967295 {
        case -1:
            v6 -= 42
        case -193:
            v2 = (v5 ^ (v2 / ((65535) | 1)))
        case 248:
            switch 32768 {
            case 42:
                v66 := GenS2{F0: (v2 * v2), F1: 128, F2: v2}
                _ = v66
            case -1:
                v5 += (^(v2))
            default:
                switch 0 {
                case -966:
                    if false {
                        v6 = (^(v6))
                        if (true || (823 >= 2)) {
                            if ((495 != 1) || (false && true)) {
                                if false {
                                    v3 = append(v3, "foo bar")
                                    if (!(false && true)) {
                                        if (!(-128 == 7)) {
                                            for fi67 := 0; fi67 < 16; fi67++ {
                                                var v68 int8 = -88
                                                _ = fi67
                                                _ = v68
                                            }
                                            v6 = -1
                                        }
                                        var v69 uint = (v5 ^ v2)
                                        v70 := GenS6{F0: 2147483648, F1: 128}
                                        _ = v69
                                        _ = v70
                                    }
                                    for fi71 := 0; fi71 < 3; fi71++ {
                                        for fi72 := 0; fi72 < 8; fi72++ {
                                            v5 -= v2
                                            _ = fi72
                                        }
                                        switch fi71 {
                                        case -1:
                                            fmt.Println(v5, v6)
                                        case -680:
                                            fmt.Println(v6)
                                        }
                                        _ = fi71
                                    }
                                } else {
                                    for fi73 := 0; fi73 < 16; fi73++ {
                                        var v74 int64 = 599
                                        if false {
                                            fmt.Println(fi73, v2, v74)
                                            var v75 uint32 = 585
                                            _ = v75
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                v76 := GenS5{F0: v2}
                                                _ = v76
                                            case "\n\t":
                                                var v77 string = "foo bar"
                                                _ = v77
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = (("hello" + "hello") + ("" + "a"))
                                                }
                                            default:
                                                var v78 uint = 65535
                                                _ = v78
                                            }
                                            if (!(967 == 32767)) {
                                                v6 = (v74 / (((v74 << (((v5 | v2)) & 63))) | 1))
                                                if false {
                                                    delete(v1, "hello")
                                                }
                                            }
                                        }
                                        v79 := GenS6{F0: 2147483648, F1: 42}
                                        _ = fi73
                                        _ = v74
                                        _ = v79
                                    }
                                }
                                var v80 int64 = (v6 - (v6 >> (((v2 - 4294967295)) & 63)))
                                _ = v80
                            } else {
                                v6 = v6
                                var v81 float64 = (0.1 * 3.14159)
                                _ = v81
                            }
                        } else {
                            v5 = 0
                            var v82 uint8 = 1
                            _ = v82
                        }
                    } else {
                        if false {
                            if false {
                                var v83 int8 = -13
                                var v84 float64 = -2.25
                                switch "foo bar" {
                                case "hello":
                                    var v85 uint16 = 22
                                    _ = v85
                                default:
                                    if (568 >= -565) {
                                        fmt.Println(v5, v6)
                                        v1["foo bar"] = 2147483648
                                    }
                                }
                                _ = v83
                                _ = v84
                            } else {
                                v2 += (v5 & 837)
                            }
                            for fi86 := 0; fi86 < 6; fi86++ {
                                var v87 float64 = 1.5
                                v87 += ((1e10 + float64(-2.25)) * v87)
                                _ = fi86
                                _ = v87
                            }
                        }
                        v88 := GenS7{F0: 1, F1: v5, F2: (("foo bar" + ("hello" + ("a" + ("\n\t" + "a")))) == "a")}
                        v89 := GenS4{F0: true}
                        _ = v88
                        _ = v89
                    }
                case -984:
                    v90 := GenS4{F0: true}
                    _ = v90
                case 42:
                    _ = len(v3)
                default:
                    fmt.Println(v5)
                }
            }
        }
        v2 = v5
        if (true && ((42 > 2) && (!(false || false)))) {
            for fi91 := 0; fi91 < 7; fi91++ {
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v92 int16 = 804
                    _ = v92
                } else {
                    var v93 int64 = v6
                    _ = v93
                }
                _ = len(v4)
                var v94 uint64 = 2
                _ = fi91
                _ = v94
            }
        }
    }
    switch ("hello" + ("a" + "\n\t")) {
    case "\n\t":
        for fi95 := 0; fi95 < 4; fi95++ {
            fmt.Println(v2, v5)
            for fi96 := 0; fi96 < 9; fi96++ {
                v6 -= -9223372036854775808
                v6 -= -9223372036854775808
                v3 = append(v3, "")
                _ = fi96
            }
            _ = fi95
        }
    case "foo bar":
        v97 := GenS5{F0: 1}
        _ = v97
    }
    var v98 uint = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
