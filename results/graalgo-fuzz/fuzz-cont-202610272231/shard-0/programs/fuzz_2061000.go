package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint; F1 int64; F2 uint64 }
type GenS9 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 int32, p2 int16) bool {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{fi4, fi4, -775, -9223372036854775808}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int64 = 266
        _ = fi6
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}
func genFn1(p0 uint64, p1 bool, p2 uint) uint8 {
    p1 = (3.14159 != ((float64(-2.25) / ((0.1 - (0.1 - (float64(float32(3.14159)) + 0.0))) * ((0.0 - (-2.25 - 0.0)) * -2.25))) - (1.5 / 0.1)))
    fmt.Println(p2, p0)
    for fi9 := 0; fi9 < 12; fi9++ {
        p2 -= (p2 & p2)
        v10 := []int{fi9, fi9}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: p2, F1: (p0 | 2)}
        p1 = (2 == 458)
        p1 = (2 != 2147483647)
        _ = fi11
        _ = v12
    }
    switch "foo bar" {
    case "\n\t":
        var v13 bool = p1
        _ = v13
    default:
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := []int{0, fi14}
            if len(v15) > 0 {
                v15[0] = (fi14 << ((p2) & 63))
            }
            v16 := GenS2{F0: "foo bar"}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = -32768
    var v3 string = ("foo bar" + (v1 + ""))
    if (-282 < 712) {
        fmt.Println(v1)
        var v17 uint = 406
        v18 := []int{}
        _ = v17
        _ = v18
    } else {
        if true {
            v1 = v3
            v19 := GenS1{F0: 42, F1: 610}
            switch v3 {
            case "foo bar":
                var v20 uint8 = 0
                _ = v20
            default:
                if (872 >= 563) {
                    v2 = (v2 + (v2 * (v2 / (((v2 << ((0) & 63))) | 1))))
                    switch v1 {
                    case "foo bar":
                        v2 += v2
                    default:
                        if (1000 == v2) {
                            var v21 uint8 = 60
                            fmt.Println(v21, v2, v1)
                            v22 := GenS1{F0: 708, F1: 42}
                            _ = v21
                            _ = v22
                        } else {
                            v2 -= 42
                            for fi23 := 0; fi23 < 16; fi23++ {
                                if ((!(!true)) || true) {
                                    v24 := []int{(fi23 / (((fi23 ^ 4294967295)) | 1)), 0, fi23, 127}
                                    for fi25 := 0; fi25 < 18; fi25++ {
                                        v26 := GenS3{F0: 146, F1: 2, F2: (v1 + ("\n\t" + ("" + v3)))}
                                        _ = fi25
                                        _ = v26
                                    }
                                    _ = v24
                                } else {
                                    var v27 int64 = 9223372036854775807
                                    _ = v27
                                }
                                _ = fi23
                            }
                            v28 := GenS3{F0: 1, F1: 2147483648, F2: v3}
                            _ = v28
                        }
                    }
                }
            }
            _ = v19
        } else {
            var v29 uint32 = 0
            _ = v29
        }
        v30 := GenS1{F0: 226, F1: 562}
        var v31 bool = true
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        switch "" {
        case "":
            v3 = (v3 + v1)
        case "\n\t":
            var v33 int64 = -1
            _ = v33
        case "a":
            var v34 string = v3
            _ = v34
        default:
            v1 = ""
        }
        v35 := []int{fi32, fi32, fi32}
        v35 = append(v35, -844)
        _ = fi32
        _ = v35
    }
    var v36 string = ((v1 + v1) + ((v1 + ((v3 + (v1 + v1)) + (v3 + v3))) + v1))
    v37 := []int{4294967295}
    _ = len(v37)
    _ = len(v37)
    v38 := GenS1{F0: 86, F1: 4294967295}
    if false {
        switch v1 {
        case "foo bar":
            var v39 float64 = 0.0
            _ = v39
        default:
            if true {
                var v40 int8 = 56
                _ = v40
            }
        }
        var v41 int = -259
        _ = v41
    } else {
        var v42 float32 = float32(0.1)
        _ = v42
    }
    fmt.Println(v2, v36, v36)
    for fi43 := 0; fi43 < 11; fi43++ {
        switch "foo bar" {
        case "hello":
            var v44 int = int(v2)
            _ = v44
        case "a":
            if (2 >= 127) {
                if (false || (2147483647 <= (fi43 ^ 127))) {
                    if true {
                        var v45 uint32 = 714
                        if (-843 != -1) {
                            var v46 int = (fi43 ^ (fi43 - (fi43 ^ -638)))
                            _ = v46
                        }
                        _ = v45
                    } else {
                        if false {
                            switch -9223372036854775808 {
                            case 32767:
                                var v47 float32 = -2.25
                                _ = v47
                            case 128:
                                if true {
                                    fmt.Println(v1, fi43, v1)
                                    if true {
                                        v48 := GenS1{F0: 255, F1: 0}
                                        v49 := GenS4{F0: v3, F1: 32767, F2: 42}
                                        _ = v48
                                        _ = v49
                                    } else {
                                        fmt.Println(v2, v3, fi43)
                                    }
                                    for fi50 := 0; fi50 < 7; fi50++ {
                                        if (!(64 >= -128)) {
                                            v2 = v2
                                            if false {
                                                var v51 int64 = -372
                                                fmt.Println(fi43, fi43, fi43)
                                                var v52 uint16 = 2
                                                _ = v51
                                                _ = v52
                                            } else {
                                                v53 := GenS4{F0: "foo bar", F1: v2, F2: 4294967295}
                                                var v54 uint16 = 32767
                                                v37 = append(v37, (fi50 >> ((18446744073709551615) & 63)))
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        switch "a" {
                                        case "":
                                            if (3.14159 < 1.5) {
                                                if false {
                                                    var v55 uint = 432
                                                    switch (((v1 + (v1 + (v3 + v36))) + "a") + "a") {
                                                    case "hello":
                                                        v3 = v3
                                                    default:
                                                        if ((((("a" + (v36 + ("\n\t" + v36))) + (("\n\t" + (v1 + ("" + v3))) + "hello")) + v3) + (v3 + v1)) <= "a") {
                                                            var v56 uint64 = 2
                                                            var v57 float64 = 3.14159
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                    v3 = v3
                                                    _ = v55
                                                } else {
                                                    var v58 int64 = 533
                                                    _ = v58
                                                }
                                                var v59 int64 = 9223372036854775807
                                                var v60 int8 = int8(v2)
                                                _ = v59
                                                _ = v60
                                            } else {
                                                v61 := GenS4{F0: "\n\t", F1: v2, F2: 899}
                                                _ = v61
                                            }
                                        case "\n\t":
                                            v62 := GenS2{F0: "hello"}
                                            _ = v62
                                        case "hello":
                                            _ = len(v37)
                                        }
                                        v63 := GenS2{F0: v36}
                                        _ = fi50
                                        _ = v63
                                    }
                                }
                            case 32768:
                                var v64 uint8 = 1
                                _ = v64
                            default:
                                switch fi43 {
                                case 709:
                                    var v65 string = ("a" + "")
                                    _ = v65
                                default:
                                    if len(v37) >= 2 {
                                        v37 = v37[1:len(v37)]
                                    }
                                }
                            }
                            if (true || ((!((828 == 2147483648) || true)) || false)) {
                                switch v36 {
                                case "":
                                    fmt.Println(v1)
                                case "a":
                                    v66 := GenS1{F0: 252, F1: 90}
                                    _ = v66
                                case "\n\t":
                                    if false {
                                        fmt.Println(v36, v1)
                                        v1 = v3
                                        v67 := GenS1{F0: 255, F1: 42}
                                        _ = v67
                                    }
                                default:
                                    v36 = (v36 + "")
                                }
                            }
                        } else {
                            for fi68 := 0; fi68 < 18; fi68++ {
                                v3 = "foo bar"
                                if true {
                                    if len(v37) > 0 {
                                        v37[0] = (^(fi43))
                                    }
                                    v36 = v36
                                    v2 -= (v2 ^ (v2 % (((v2 - v2)) | 1)))
                                } else {
                                    fmt.Println(fi68, v2, v1)
                                }
                                fmt.Println(v2)
                                switch fi68 {
                                case 9223372036854775807:
                                    v2 = (v2 * 32767)
                                case -1:
                                    switch "a" {
                                    case "":
                                        v36 = ((v36 + "foo bar") + "a")
                                    case "\n\t":
                                        var v69 uint8 = 128
                                        _ = v69
                                    default:
                                        if len(v37) > 0 {
                                            v37[0] = (fi68 / ((32767) | 1))
                                        }
                                    }
                                }
                                _ = fi68
                            }
                            if false {
                                var v70 int64 = -9223372036854775808
                                _ = len(v37)
                                if (!false) {
                                    if true {
                                        for fi71 := 0; fi71 < 1; fi71++ {
                                            v72 := GenS5{F0: 330, F1: 2}
                                            _ = fi71
                                            _ = v72
                                        }
                                        var v73 int8 = 1
                                        v74 := GenS4{F0: v1, F1: int16(v70), F2: 330}
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                _ = v70
                            }
                        }
                        if len(v37) > 0 {
                            v37[0] = fi43
                        }
                        for fi75 := 0; fi75 < 9; fi75++ {
                            var v76 uint8 = 223
                            if false {
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                                v3 = v1
                                switch (v3 + v1) {
                                case "foo bar":
                                    v37 = append(v37, fi43)
                                default:
                                    v77 := GenS2{F0: v36}
                                    _ = v77
                                }
                                if (1e10 == (1.5 * float32(float64(3.14159)))) {
                                    if (0 <= 2) {
                                        v3 = v3
                                        v78 := GenS2{F0: ""}
                                        var v79 uint8 = v76
                                        _ = v78
                                        _ = v79
                                    }
                                }
                            } else {
                                v36 = ((v1 + "a") + v1)
                            }
                            switch 2147483647 {
                            case 32767:
                                if true {
                                    v1 = v36
                                }
                            }
                            fmt.Println(v36, fi43)
                            _ = fi75
                            _ = v76
                        }
                    }
                    fmt.Println(v36, v1)
                } else {
                    for fi80 := 0; fi80 < 19; fi80++ {
                        if (9223372036854775807 == 9223372036854775807) {
                            if false {
                                v81 := GenS6{F0: 127}
                                _ = v81
                            } else {
                                var v82 uint = 65535
                                v83 := GenS7{F0: 0, F1: 672, F2: 1}
                                _ = len(v37)
                                _ = v82
                                _ = v83
                            }
                            var v84 uint16 = 0
                            _ = v84
                        } else {
                            fmt.Println(v1, v3)
                        }
                        _ = len(v37)
                        _ = fi80
                    }
                    v85 := GenS2{F0: v1}
                    _ = len(v37)
                    _ = v85
                }
            } else {
                fmt.Println(v1, v36)
                var v86 bool = true
                _ = v86
            }
        case "":
            var v87 string = v3
            _ = v87
        default:
            for fi88 := 0; fi88 < 7; fi88++ {
                v3 = v36
                fmt.Println(v2)
                var v89 int8 = 61
                _ = fi88
                _ = v89
            }
        }
        v90 := GenS7{F0: 65535, F1: 18446744073709551615, F2: 127}
        for fi91 := 0; fi91 < 2; fi91++ {
            v3 = ""
            if true {
                v92 := GenS8{F0: 950, F1: 42, F2: 18446744073709551615}
                _ = len(v37)
                var v93 int64 = 9223372036854775807
                _ = v92
                _ = v93
            } else {
                v36 = ("hello" + ("hello" + "foo bar"))
            }
            _ = fi91
        }
        _ = fi43
        _ = v90
    }
    fmt.Println(v2)
    for fi94 := 0; fi94 < 18; fi94++ {
        fmt.Println(v2)
        _ = fi94
    }
    for fi95 := 0; fi95 < 19; fi95++ {
        v96 := GenS5{F0: 2, F1: 127}
        _ = fi95
        _ = v96
    }
    fmt.Println(v3)
    if false {
        v3 = "foo bar"
        if (((255 != 2) && false) || (((float64(0.0) - 1e10) + ((0.0 + (0.1 * (float64(1.5) / (-2.25 + 3.14159)))) * 1.5)) < ((0.0 * (1e10 * 1.5)) - 1.5))) {
            v37 = append(v37, 127)
            if true {
                if (-2.25 <= 0.1) {
                    if len(v37) > 0 {
                        v37[0] = 127
                    }
                } else {
                    v97 := GenS6{F0: 69}
                    var v98 uint64 = 653
                    if false {
                        fmt.Println(v36)
                    }
                    _ = v97
                    _ = v98
                }
                v1 = v1
                for fi99 := 0; fi99 < 12; fi99++ {
                    v100 := GenS9{F0: 1, F1: 180}
                    var v101 uint = 1
                    v102 := GenS4{F0: (v1 + ("" + "a")), F1: v2, F2: 1}
                    _ = fi99
                    _ = v100
                    _ = v101
                    _ = v102
                }
            }
            if true {
                var v103 uint = 65535
                v36 = "a"
                _ = v103
            }
        } else {
            if len(v37) > 0 {
                v37[0] = 75
            }
        }
        switch (("foo bar" + "hello") + ("" + ((("foo bar" + v3) + "") + v1))) {
        case "\n\t":
            var v104 string = ("hello" + v1)
            _ = v104
        case "hello":
            if (!true) {
                if len(v37) >= 2 {
                    v37 = v37[1:len(v37)]
                }
            }
        default:
            v37 = append(v37, 9223372036854775807)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
