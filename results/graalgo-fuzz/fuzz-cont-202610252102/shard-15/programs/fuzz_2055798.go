package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 bool; F2 int16 }
type GenS8 struct { F0 int32; F1 bool; F2 uint32 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 int8, p1 uint) uint32 {
    if (32768 <= 127) {
        v6 := []int{-1}
        var v7 int8 = 127
        v8 := GenS1{F0: 7, F1: 45, F2: "a"}
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 1, F1: 1, F2: "hello"}
        p1 += (^(p1))
        var v10 string = "foo bar"
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint32 = 125
        v13 := GenS2{F0: p0}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(fi14, p0)
            var v16 int16 = -32768
            if (!(42 != 939)) {
                if (556 >= 0) {
                    if (9223372036854775807 <= -95) {
                        var v17 uint64 = 239
                        _ = v17
                    }
                    var v18 string = ("" + "foo bar")
                    if true {
                        var v19 float32 = 1e10
                        v20 := []int{(fi14 >> ((p1) & 63)), fi14}
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint32 = 2147483648
                        _ = v21
                    }
                    v22 := GenS3{F0: (p0 / ((p0) | 1)), F1: (!true)}
                    _ = v18
                    _ = v22
                }
                if true {
                    if true {
                        var v23 bool = false
                        if v23 {
                            v24 := []int{(fi15 * fi15), (fi14 & fi14), -9223372036854775808}
                            v25 := GenS1{F0: 255, F1: (fi15 >> ((p1) & 63)), F2: ("a" + "hello")}
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS3{F0: p0, F1: (!true)}
                            var v27 string = ("\n\t" + ((("hello" + "a") + "foo bar") + ""))
                            var v28 int8 = (-(p0))
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v23
                    } else {
                        var v29 int32 = -2147483648
                        if (255 == 210) {
                            v30 := []int{fi14, (^(fi15))}
                            if (!(p1 <= p1)) {
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                                if len(v30) > 0 {
                                    v30[0] = (fi14 << ((656) & 63))
                                }
                                v30 = append(v30, fi15)
                            } else {
                                if len(v30) > 0 {
                                    v30[0] = fi15
                                }
                            }
                            var v31 uint8 = 2
                            v32 := GenS4{F0: (!true), F1: (^(p1)), F2: fi15}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    }
                }
                if true {
                    var v33 bool = false
                    if (-445 == 0) {
                        v34 := GenS5{F0: v33}
                        _ = v34
                    }
                    _ = v33
                } else {
                    v35 := GenS3{F0: -1, F1: ((true || true) && true)}
                    _ = v35
                }
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    return 187
}
func genFn1() uint64 {
    v36 := GenS5{F0: (!(-2147483648 >= 42))}
    if (((763 == 18446744073709551615) && true) && true) {
        var v37 int = 573
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := []int{128, v37, (fi38 / ((fi38) | 1)), 127}
            for fi40 := 0; fi40 < 14; fi40++ {
                v37 += fi40
                fmt.Println(fi38)
                switch "foo bar" {
                case "a":
                    v41 := GenS6{F0: 4294967295}
                    _ = v41
                case "hello":
                    var v42 uint16 = 537
                    _ = v42
                default:
                    if true {
                        if (!true) {
                            if (fi40 <= 4294967295) {
                                if (false && (true && (false || (!false)))) {
                                    fmt.Println(v37, fi38, v37)
                                    var v43 int = 32767
                                    _ = v43
                                }
                                v44 := GenS5{F0: false}
                                _ = v44
                            } else {
                                v37 += fi38
                            }
                            v37 += 2147483647
                        } else {
                            v37 -= fi38
                        }
                        if len(v39) > 0 {
                            v39[0] = fi40
                        }
                        v39 = append(v39, (v37 * v37))
                    } else {
                        switch "foo bar" {
                        case "":
                            switch (^(v37)) {
                            case -352:
                                if true {
                                    if true {
                                        v45 := GenS6{F0: 468}
                                        _ = v45
                                    }
                                    if false {
                                        if (false && true) {
                                            var v46 uint64 = 4294967295
                                            fmt.Println(fi40)
                                            if (false || ((true || ((32768 > v37) && false)) && (!true))) {
                                                fmt.Println(fi40)
                                            }
                                            _ = v46
                                        } else {
                                            if (42 != -32768) {
                                                v47 := GenS3{F0: 127, F1: true}
                                                var v48 bool = ((false || true) && true)
                                                _ = v47
                                                _ = v48
                                            } else {
                                                v49 := GenS1{F0: 255, F1: (fi40 | 501), F2: "foo bar"}
                                                _ = v49
                                            }
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            v37 += (fi40 * v37)
                                        }
                                        v37 += -302
                                    } else {
                                        if ((979 <= 65535) || (-1 > 42)) {
                                            if true {
                                                if false {
                                                    v37 = v37
                                                }
                                                v50 := GenS6{F0: 4294967295}
                                                var v51 float32 = 1e10
                                                _ = v50
                                                _ = v51
                                            }
                                            v52 := GenS5{F0: true}
                                            switch ((("foo bar" + "a") + ("a" + (("" + ("hello" + "a")) + "\n\t"))) + "hello") {
                                            case "foo bar":
                                                switch -675 {
                                                case 0:
                                                    v53 := GenS6{F0: 119}
                                                    _ = v53
                                                }
                                            case "a":
                                                var v54 uint64 = 4294967295
                                                _ = v54
                                            case "":
                                                if len(v39) > 0 {
                                                    v39[0] = fi40
                                                }
                                            }
                                            v37 -= (fi40 & (fi40 / ((fi38) | 1)))
                                            v37 += (fi38 / (((fi38 ^ (fi38 - (fi40 ^ fi38)))) | 1))
                                            v37 = fi40
                                            switch (fi38 ^ (fi40 * v37)) {
                                            case -159:
                                                if (true || (true || false)) {
                                                    if len(v39) > 0 {
                                                        v39[0] = fi40
                                                    }
                                                    if (-605 > 999) {
                                                        v55 := GenS6{F0: 2}
                                                        v37 = v37
                                                        v37 = 128
                                                        _ = v55
                                                    }
                                                }
                                            case 4294967295:
                                                var v56 int16 = 42
                                                _ = v56
                                            }
                                            _ = v52
                                        }
                                        v37 = 4294967295
                                        v57 := GenS7{F0: 27, F1: (470 < 65535), F2: -1}
                                        _ = v57
                                    }
                                    if (-1 >= 9223372036854775807) {
                                        if (724 < 1) {
                                            if ((0.1 + -2.25) <= 0.0) {
                                                if len(v39) >= 2 {
                                                    v39 = v39[1:len(v39)]
                                                }
                                                if (0.1 > (0.1 / (float64(-2.25) - 1e10))) {
                                                    switch "hello" {
                                                    case "hello":
                                                        v37 = fi38
                                                    case "\n\t":
                                                        switch v37 {
                                                        case 984:
                                                            if (-72 <= 127) {
                                                                switch ("hello" + "a") {
                                                                case "\n\t":
                                                                    if len(v39) >= 2 {
                                                                        v39 = v39[1:len(v39)]
                                                                    }
                                                                }
                                                                var v58 uint16 = 910
                                                                _ = v58
                                                            } else {
                                                                v39 = append(v39, v37)
                                                                if false {
                                                                    fmt.Println(fi38, v37, fi40)
                                                                    var v59 float64 = 3.14159
                                                                    if false {
                                                                        v60 := GenS8{F0: 1, F1: (!(-241 <= -1)), F2: 2147483647}
                                                                        fmt.Println(fi38, v37, v59)
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 int32 = 1
                                                                        var v62 int16 = 32767
                                                                        var v63 uint32 = 2147483648
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                    _ = v59
                                                                } else {
                                                                    fmt.Println(fi40, v37, fi38)
                                                                }
                                                                v37 -= -1
                                                            }
                                                        }
                                                    case "a":
                                                        var v64 int8 = 90
                                                        _ = v64
                                                    }
                                                    v37 -= fi38
                                                    v37 = 4294967295
                                                } else {
                                                    fmt.Println(fi38, v37, v37)
                                                }
                                            }
                                            var v65 string = "a"
                                            v37 -= fi38
                                            v37 = -2147483648
                                            _ = len(v39)
                                            _ = v65
                                        }
                                        var v66 int32 = 2147483647
                                        switch (-(v37)) {
                                        case -2147483648:
                                            if true {
                                                v67 := GenS9{F0: 1000, F1: false}
                                                _ = v67
                                            }
                                        case 2147483647:
                                            v68 := GenS10{F0: 0, F1: -1, F2: 32767}
                                            _ = v68
                                        case 127:
                                            var v69 int8 = 1
                                            _ = v69
                                        }
                                        _ = v66
                                    }
                                } else {
                                    fmt.Println(fi40, fi40, v37)
                                    if (65535 >= 0) {
                                        if (!(255 != 667)) {
                                            var v70 int8 = 42
                                            if (-227 == 2147483647) {
                                                var v71 uint = 1
                                                v71 = v71
                                                v72 := GenS10{F0: 1, F1: 0, F2: 2}
                                                _ = v71
                                                _ = v72
                                            }
                                            v70 = (v70 & v70)
                                            _ = v70
                                        }
                                    }
                                }
                            default:
                                switch "hello" {
                                case "foo bar":
                                    var v73 uint64 = 0
                                    _ = v73
                                }
                            }
                        case "foo bar":
                            v74 := GenS11{F0: 127, F1: 717, F2: 1}
                            _ = v74
                        case "hello":
                            fmt.Println(fi38)
                        }
                        if false {
                            var v75 int64 = 0
                            _ = v75
                        }
                    }
                }
                if len(v39) > 0 {
                    v39[0] = v37
                }
                switch "\n\t" {
                case "a":
                    fmt.Println(fi40, fi38)
                default:
                    if (!(4294967295 != 2)) {
                        switch ((("hello" + "") + "") + "hello") {
                        case "hello":
                            v76 := GenS6{F0: 4294967295}
                            _ = v76
                        case "":
                            switch v37 {
                            case 388:
                                fmt.Println(fi38, fi38, fi40)
                            case -9223372036854775808:
                                var v77 int8 = 45
                                _ = v77
                            }
                        }
                    } else {
                        if (((-2.25 + (3.14159 / 3.14159)) / (1e10 - ((float32((-2.25 / 1e10)) / 3.14159) + (3.14159 - -2.25)))) <= (3.14159 / -2.25)) {
                            if false {
                                var v78 string = "hello"
                                if true {
                                    var v79 float32 = (-2.25 * ((((float32(0.0) / (-2.25 - 0.0)) / (0.0 + 0.1)) + -2.25) * ((0.0 - ((-2.25 / 1.0) + 1e10)) + 3.14159)))
                                    var v80 bool = (v37 >= (^(v37)))
                                    _ = v79
                                    _ = v80
                                }
                                _ = v78
                            } else {
                                var v81 int16 = -1
                                if (0 >= 65535) {
                                    v81 += (v81 >> ((1) & 63))
                                    if len(v39) >= 2 {
                                        v39 = v39[1:len(v39)]
                                    }
                                    v37 -= -1
                                } else {
                                    var v82 uint16 = 99
                                    _ = v82
                                }
                                _ = v81
                            }
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                            _ = len(v39)
                            var v83 uint64 = 345
                            _ = v83
                        }
                        if (0 > 42) {
                            v37 = (fi40 & (^(fi40)))
                            var v84 int = 2147483647
                            fmt.Println(fi40, fi38, fi38)
                            _ = v84
                        } else {
                            v85 := GenS7{F0: 467, F1: (fi40 < v37), F2: -1}
                            _ = v85
                        }
                        switch v37 {
                        case 32768:
                            fmt.Println(fi38)
                        case -1:
                            var v86 bool = ((!false) || false)
                            _ = v86
                        default:
                            fmt.Println(fi38, fi40, fi40)
                        }
                    }
                }
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        var v87 int8 = 1
        v88 := GenS8{F0: -2147483648, F1: ((339 != 9223372036854775807) && true), F2: 65535}
        v87 += (v87 / ((v87) | 1))
        _ = v87
        _ = v88
    }
    _ = v36
    return 42
}

func main() {
    var v1 float64 = 0.0
    var v2 uint = 42
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 0
    for fi89 := 0; fi89 < 6; fi89++ {
        var v90 uint64 = 171
        _ = fi89
        _ = v90
    }
    v5 -= 32767
    for fi91 := 0; fi91 < 17; fi91++ {
        fmt.Println(v2)
        _ = fi91
    }
    var v92 float32 = 1e10
    for fi93 := 0; fi93 < 8; fi93++ {
        fmt.Println(v1)
        _ = fi93
    }
    for fi94 := 0; fi94 < 15; fi94++ {
        if true {
            var v95 uint16 = (v5 - v5)
            fmt.Println(v95, v95, v92)
            var v96 uint64 = 162
            _ = v95
            _ = v96
        } else {
            var v97 int = fi94
            var v98 uint32 = 2147483647
            switch "\n\t" {
            case "a":
                if (float64(1e10) >= 0.1) {
                    v2 -= 2
                } else {
                    fmt.Println(v97)
                }
            }
            _ = v97
            _ = v98
        }
        v1 += float64(((v92 * -2.25) + v92))
        _ = fi94
    }
    v1 = 1.5
    fmt.Println(v92, v1)
    fmt.Println(v92)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
}
