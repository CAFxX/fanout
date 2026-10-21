package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string; F1 int8; F2 uint }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 int; F2 int32 }
type GenS6 struct { F0 bool }
func genFn0(p0 string) bool {
    if (!true) {
        if ((42 == 128) || false) {
            var v4 uint32 = 2
            _ = v4
        }
        p0 = (p0 + p0)
    } else {
        v5 := GenS1{F0: false, F1: 4294967295}
        if (1e10 > 0.1) {
            p0 = (((p0 + p0) + "foo bar") + p0)
        } else {
            var v6 string = p0
            v7 := GenS1{F0: false, F1: 2}
            _ = v6
            _ = v7
        }
        _ = v5
    }
    var v8 int = -721
    var v9 uint64 = 1
    var v10 uint = 115
    var v11 int64 = -9223372036854775808
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (!(383 == (^(v10))))
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483647
    var v12 uint8 = 1
    var v13 bool = false
    v14 := GenS1{F0: (v3 < 457), F1: v3}
    if ("hello" < "foo bar") {
        switch "" {
        case "a":
            switch -9223372036854775808 {
            case 1:
                v15 := GenS1{F0: v13, F1: (v3 * 421)}
                _ = v15
            default:
                if (!(!("\n\t" > "\n\t"))) {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 string = "a"
                        _ = fi16
                        _ = v17
                    }
                    var v18 uint8 = 1
                    var v19 uint16 = 1
                    _ = v18
                    _ = v19
                } else {
                    v1 += ((3.14159 * v1) - v1)
                    var v20 int = 0
                    _ = v20
                }
            }
        case "foo bar":
            if v13 {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 uint = 929
                    fmt.Println(fi21, v12)
                    _ = fi21
                    _ = v22
                }
                if (65535 <= uint(v3)) {
                    var v23 bool = v13
                    if (2147483647 != 128) {
                        v24 := GenS2{F0: 65535, F1: false}
                        v25 := GenS1{F0: v13, F1: v3}
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                } else {
                    fmt.Println(v3, v13, v12)
                    v3 += (v3 - (v3 / ((v3) | 1)))
                    v13 = (v3 != 2147483647)
                }
            }
        case "hello":
            if false {
                var v26 uint32 = (v3 + v3)
                _ = v26
            } else {
                switch "" {
                case "foo bar":
                    var v27 int32 = 1
                    _ = v27
                case "hello":
                    v28 := v2[v3]
                    _ = v28
                case "\n\t":
                    var v29 float32 = ((float32(3.14159) / 0.1) / 1.0)
                    _ = v29
                default:
                    v30 := GenS3{F0: ("" + "foo bar"), F1: -128, F2: 0}
                    _ = v30
                }
                v31 := GenS1{F0: (51 >= -128), F1: v3}
                for fi32 := 0; fi32 < 20; fi32++ {
                    v13 = ((v13 || (!(-96 == 9223372036854775807))) || (!v13))
                    _ = fi32
                }
                _ = v31
            }
        default:
            switch ("" + "hello") {
            case "foo bar":
                var v33 uint64 = 4294967295
                _ = v33
            default:
                v34 := GenS1{F0: (6 <= 610), F1: (v3 << ((966) & 63))}
                _ = v34
            }
        }
        if (true && (59 >= 753)) {
            fmt.Println(v3, v1, v3)
            if v13 {
                v35 := GenS2{F0: 255, F1: (1 == 42)}
                switch ("foo bar" + "") {
                case "":
                    var v36 int64 = 1
                    _ = v36
                case "hello":
                    fmt.Println(v1)
                default:
                    if true {
                        _ = len(v2)
                    } else {
                        switch ("\n\t" + (("\n\t" + "\n\t") + "a")) {
                        case "foo bar":
                            switch 32767 {
                            case -614:
                                var v37 uint16 = 222
                                _ = v37
                            case -1:
                                switch ("" + "\n\t") {
                                case "a":
                                    v38 := []int{1, 127, 9223372036854775807, 32767}
                                    _ = v38
                                case "\n\t":
                                    v13 = (true && false)
                                }
                            case 9223372036854775807:
                                if (((false || v13) && v13) && v13) {
                                    for fi39 := 0; fi39 < 15; fi39++ {
                                        for fi40 := 0; fi40 < 19; fi40++ {
                                            v41 := GenS1{F0: v13, F1: (v3 >> ((255) & 63))}
                                            v2[v3] = ""
                                            if v13 {
                                                var v42 bool = false
                                                var v43 uint32 = (v3 ^ (^(v3)))
                                                v13 = ((0 < -1) && (!v42))
                                                _ = v42
                                                _ = v43
                                            }
                                            if true {
                                                v44 := []int{(fi39 + (^(fi40))), fi40}
                                                fmt.Println(v12, v1, v13)
                                                _ = len(v2)
                                                _ = v44
                                            }
                                            switch (-(fi39)) {
                                            case 0:
                                                v12 += 42
                                            case 32768:
                                                v3 = v3
                                            }
                                            fmt.Println(v3)
                                            fmt.Println(v1, fi39, v1)
                                            fmt.Println(fi40)
                                            _ = fi40
                                            _ = v41
                                        }
                                        for fi45 := 0; fi45 < 11; fi45++ {
                                            v2[2] = "foo bar"
                                            delete(v2, (v3 | 4294967295))
                                            var v46 int = (fi45 << ((18446744073709551615) & 63))
                                            _ = fi45
                                            _ = v46
                                        }
                                        _ = fi39
                                    }
                                    for fi47 := 0; fi47 < 12; fi47++ {
                                        v2[v3] = "foo bar"
                                        v1 = v1
                                        for fi48 := 0; fi48 < 18; fi48++ {
                                            fmt.Println(v1)
                                            v12 = 131
                                            var v49 uint = 65535
                                            v12 = 1
                                            _ = fi48
                                            _ = v49
                                        }
                                        _ = fi47
                                    }
                                }
                            default:
                                switch 2147483647 {
                                case 4294967295:
                                    v50 := []int{0, -258, 42, 42}
                                    _ = v50
                                case -1:
                                    v51 := []int{}
                                    _ = v51
                                }
                            }
                        default:
                            v52 := GenS2{F0: 553, F1: ("hello" != "a")}
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 18; fi53++ {
                            v54 := []int{fi53, fi53, (fi53 ^ fi53)}
                            delete(v2, (v3 | (v3 | 2)))
                            v1 -= (v1 + v1)
                            _ = fi53
                            _ = v54
                        }
                    }
                }
                _ = v35
            }
            v55 := []int{4294967295}
            _ = v55
        }
    }
    var v56 int = 42
    fmt.Println(v1, v1, v12)
    v57 := GenS2{F0: 32767, F1: (1 < -1)}
    var v58 float64 = v1
    if false {
        v3 += v3
        v3 -= 2
        for fi59 := 0; fi59 < 3; fi59++ {
            if (true || (0.0 != 1e10)) {
                fmt.Println(fi59, v56, v3)
                v12 += v12
                v60 := GenS4{F0: 65535, F1: ((255 >= v12) || ((v1 - float64((((float32(float64((-2.25 * 0.0))) + 1e10) / 1.0) / -2.25))) != 1.5))}
                _ = v60
            } else {
                v1 = v1
            }
            v12 += v12
            var v61 bool = ((v3 * v3) <= v3)
            _ = fi59
            _ = v61
        }
    } else {
        fmt.Println(v13, v1, v12)
        if (1 > int8(v12)) {
            if (396 != 4294967295) {
                for fi62 := 0; fi62 < 10; fi62++ {
                    switch "\n\t" {
                    case "":
                        v58 += float64(1.5)
                    case "\n\t":
                        if (!true) {
                            if v13 {
                                for fi63 := 0; fi63 < 10; fi63++ {
                                    v64 := GenS2{F0: 902, F1: (1 <= 0)}
                                    v2[42] = ("\n\t" + "\n\t")
                                    _ = fi63
                                    _ = v64
                                }
                                var v65 int16 = 237
                                fmt.Println(v65, v3)
                                _ = v65
                            } else {
                                var v66 int8 = 100
                                _ = v66
                            }
                            v67 := []int{322, fi62, (fi62 % ((16) | 1)), v56}
                            _ = v67
                        } else {
                            v68 := []int{2147483647, (-(fi62))}
                            fmt.Println(v56, v12, v3)
                            fmt.Println(v13)
                            _ = v68
                        }
                    case "hello":
                        v3 = v3
                    }
                    if (v13 && v13) {
                        v58 = (float64(0.0) + 1e10)
                        fmt.Println(v3, v3, v3)
                        v12 = (v12 - 167)
                    } else {
                        delete(v2, v3)
                    }
                    _ = fi62
                }
                delete(v2, (^(v3)))
            } else {
                var v69 float32 = (((-2.25 * (1.5 - float32(-2.25))) + 1.5) * 0.0)
                var v70 float64 = ((v1 - ((v1 - float64(float32((float64(v69) - 1e10)))) - ((float64(v69) / v58) + (float64((0.1 / 3.14159)) * ((-2.25 / float64(((float32((3.14159 - (float64(1.5) + float64(1e10)))) - v69) / (v69 / (((1.5 + 0.1) / (1e10 - (1e10 * v69))) + (0.1 * float32(v1))))))) * 1e10))))) / v1)
                _ = v69
                _ = v70
            }
            v71 := GenS2{F0: 42, F1: v13}
            _ = v71
        } else {
            v58 += -2.25
        }
        for fi72 := 0; fi72 < 13; fi72++ {
            v73 := GenS1{F0: (1 != 0), F1: 737}
            if (-24 != 127) {
                v58 -= (((v1 + 3.14159) - v58) + 1.5)
            }
            for fi74 := 0; fi74 < 18; fi74++ {
                v56 = fi74
                fmt.Println(fi74, fi74)
                v75 := GenS1{F0: (!(v13 && (1 >= 0))), F1: 2}
                _ = fi74
                _ = v75
            }
            _ = fi72
            _ = v73
        }
    }
    fmt.Println(v3, v13)
    for fi76 := 0; fi76 < 4; fi76++ {
        v13 = true
        switch "hello" {
        case "\n\t":
            var v77 int32 = 0
            _ = v77
        case "a":
            for fi78 := 0; fi78 < 8; fi78++ {
                var v79 uint8 = 213
                _ = fi78
                _ = v79
            }
        case "foo bar":
            for fi80 := 0; fi80 < 15; fi80++ {
                switch 398 {
                case 127:
                    var v81 uint = 18446744073709551615
                    _ = v81
                default:
                    var v82 int64 = -1
                    _ = len(v2)
                    v1 = v1
                    _ = v82
                }
                switch ("" + ("foo bar" + "\n\t")) {
                case "a":
                    var v83 uint32 = 702
                    _ = v83
                default:
                    delete(v2, (v3 & (v3 + (v3 + v3))))
                }
                fmt.Println(v3, fi80)
                if (!v13) {
                    switch ("" + "") {
                    case "":
                        switch "hello" {
                        case "hello":
                            v84 := []int{v56, (fi76 % (((v56 % (((fi80 & 1)) | 1))) | 1)), v56, -181}
                            _ = v84
                        }
                    case "foo bar":
                        if v13 {
                            var v85 uint16 = 2
                            v86 := []int{fi80, -394}
                            _ = v85
                            _ = v86
                        } else {
                            v87 := GenS5{F0: (^(v3)), F1: (fi76 & 9223372036854775807), F2: -2147483648}
                            if (-1 >= -2147483648) {
                                fmt.Println(v13, v13, v56)
                                var v88 int = fi76
                                switch (fi80 & v56) {
                                case -494:
                                    fmt.Println(v13)
                                default:
                                    v89 := GenS6{F0: v13}
                                    _ = v89
                                }
                                _ = v88
                            } else {
                                v13 = (v13 && false)
                            }
                            _ = v87
                        }
                    }
                    var v90 uint16 = 65535
                    _ = v90
                }
                v13 = (0 <= 907)
                v3 += v3
                _ = fi80
            }
        default:
            for fi91 := 0; fi91 < 2; fi91++ {
                v92 := []int{(fi76 << ((42) & 63)), 0}
                _ = fi91
                _ = v92
            }
        }
        var v93 int64 = 638
        _ = fi76
        _ = v93
    }
    v58 = 3.14159
    v56 = -9223372036854775808
    var v94 uint16 = 152
    v95 := GenS2{F0: (v94 * (v94 | (^(v94)))), F1: v13}
    var v96 int8 = 87
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v58
    _ = v94
    _ = v95
    _ = v96
}
