package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 bool }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint; F1 string; F2 uint8 }
func genFn0(p0 uint16) int {
    p0 -= p0
    var v7 int = 2147483647
    p0 = (p0 * p0)
    _ = p0
    _ = v7
    return (v7 >> ((191) & 63))
}
func genFn1(p0 int32, p1 uint64, p2 uint8) string {
    v8 := GenS1{F0: 65535}
    var v9 int16 = 32767
    v10 := GenS2{F0: 434}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 int16 = -504
    var v2 int16 = v1
    var v3 uint64 = 4294967295
    var v4 float64 = 3.14159
    v5 := make(map[uint32]uint32)
    v6 := []string{"foo bar"}
    var v11 bool = ((42 <= 2) && ((v4 * (v4 / v4)) >= ((1.5 * float64(1.5)) - (((1e10 / 1.5) - 1e10) * ((1e10 * v4) * v4)))))
    var v12 uint8 = 88
    delete(v5, 407)
    v3 -= v3
    v13 := GenS2{F0: 336}
    if (float32(3.14159) >= (1e10 - float32((-2.25 * v4)))) {
        if v11 {
            fmt.Println(v12, v11)
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            fmt.Println(v12)
        }
        var v14 bool = true
        v6 = append(v6, "")
        _ = v14
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            delete(v5, 65535)
            switch (fi15 % (((^(fi15))) | 1)) {
            case -774:
                for fi16 := 0; fi16 < 1; fi16++ {
                    v11 = ((v3 == v3) && true)
                    if v11 {
                        var v17 int64 = 1
                        _ = v17
                    }
                    v18 := GenS2{F0: 42}
                    _ = fi16
                    _ = v18
                }
            }
            _ = fi15
        }
        switch 1 {
        case 98:
            if (0 <= 284) {
                v5[722] = 1
                var v19 int16 = v1
                _ = v19
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 = (v2 * -32768)
                v11 = v11
            }
        case 42:
            v1 = (v2 / (((^(v2))) | 1))
        }
    }
    var v20 float32 = 0.1
    if (false || (v12 != (v12 & 1))) {
        var v21 int16 = 42
        _ = v21
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint32 = 902
            var v24 uint32 = (v23 % ((v23) | 1))
            var v25 float64 = v4
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = len(v6)
    }
    _ = len(v5)
    _ = len(v6)
    if v11 {
        v12 += 2
        v26 := GenS2{F0: 501}
        for fi27 := 0; fi27 < 11; fi27++ {
            switch "" {
            case "":
                v11 = (-1 >= 42)
            case "\n\t":
                for fi28 := 0; fi28 < 1; fi28++ {
                    v12 += (v12 | (v12 % (((^(v12))) | 1)))
                    v29 := GenS3{F0: true, F1: -713, F2: (true && v11)}
                    v30 := GenS3{F0: ((((((((v20 - v20) - 0.0) + v20) - ((1.5 * v20) / (-2.25 * (v20 + -2.25)))) - v20) + v20) > ((float32(v4) - (v20 - (0.1 * v20))) * 3.14159)) && v11), F1: -1, F2: ((v2 | 32767) != 1000)}
                    fmt.Println(v3, fi27, v20)
                    fmt.Println(v2, fi28)
                    v31 := GenS3{F0: v11, F1: 42, F2: v11}
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            case "hello":
                delete(v5, 65535)
            }
            _ = fi27
        }
        _ = v26
    } else {
        var v32 uint8 = (^(v12))
        _ = v32
    }
    v20 += ((1.5 + float32(3.14159)) - (((float32(float64(float32(1.5))) - ((v20 / 1.0) * 0.1)) / (v20 - 3.14159)) / v20))
    if (4 < 42) {
        switch 9223372036854775807 {
        case 32767:
            for fi33 := 0; fi33 < 12; fi33++ {
                for fi34 := 0; fi34 < 5; fi34++ {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        if v11 {
                            v35 := GenS4{F0: 2, F1: 0, F2: 0}
                            fmt.Println(v11, v12, v11)
                            v12 += (v12 * (v12 | 127))
                            _ = v35
                        } else {
                            if len(v6) > 0 {
                                v6[0] = ""
                            }
                            var v36 uint = 720
                            fmt.Println(v2)
                            _ = v36
                        }
                    case "":
                        v37 := GenS5{F0: 42}
                        _ = v37
                    }
                    v1 += v2
                    var v38 int = 128
                    v39 := GenS5{F0: 10}
                    _ = fi34
                    _ = v38
                    _ = v39
                }
                _ = fi33
            }
        default:
            for fi40 := 0; fi40 < 14; fi40++ {
                v2 = (v1 | 32767)
                if len(v6) > 0 {
                    v6[0] = (("a" + "foo bar") + "foo bar")
                }
                _ = fi40
            }
        }
        var v41 uint8 = v12
        for fi42 := 0; fi42 < 7; fi42++ {
            switch (fi42 - genFn0(42)) {
            case 2147483647:
                v43 := GenS1{F0: 2}
                _ = v43
            case 9223372036854775807:
                v11 = v11
            case -1:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            if v11 {
                if v11 {
                    if (!true) {
                        switch fi42 {
                        case -2147483648:
                            var v44 int8 = -128
                            _ = v44
                        case 32768:
                            for fi45 := 0; fi45 < 11; fi45++ {
                                switch genFn0(42) {
                                case 9223372036854775807:
                                    fmt.Println(v4, fi42)
                                case -1:
                                    v4 -= (0.1 + v4)
                                case -42:
                                    v6 = append(v6, ("foo bar" + "foo bar"))
                                default:
                                    var v46 int64 = -9223372036854775808
                                    switch "foo bar" {
                                    case "a":
                                        var v47 uint16 = uint16(v41)
                                        _ = v47
                                    case "\n\t":
                                        var v48 bool = (!true)
                                        _ = v48
                                    case "":
                                        if (v4 > (((v4 - 1e10) - v4) / (v4 - 3.14159))) {
                                            var v49 uint = 42
                                            _ = v49
                                        } else {
                                            var v50 uint = 0
                                            v50 -= v50
                                            _ = v50
                                        }
                                    }
                                    var v51 float32 = (0.1 * ((float32((3.14159 / 1.0)) + v20) - float32(v4)))
                                    v6 = append(v6, "foo bar")
                                    if v11 {
                                        if (v11 || (v51 != v51)) {
                                            switch ("" + "hello") {
                                            case "a":
                                                switch (fi42 + (fi42 / ((fi45) | 1))) {
                                                case -2147483648:
                                                    var v52 string = "foo bar"
                                                    _ = v52
                                                case -842:
                                                    var v53 uint16 = 32
                                                    if v11 {
                                                        fmt.Println(v51, fi45)
                                                    } else {
                                                        v54 := GenS4{F0: 281, F1: 2147483647, F2: -96}
                                                        v4 -= -2.25
                                                        _ = v54
                                                    }
                                                    v55 := GenS6{F0: (^(v46)), F1: 0}
                                                    _ = v53
                                                    _ = v55
                                                case 128:
                                                    if (v2 >= v1) {
                                                        var v56 int32 = 2147483647
                                                        v57 := v5[42]
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        if v11 {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        } else {
                                                            var v58 uint = 1
                                                            v46 = (v46 + 1)
                                                            _ = v58
                                                        }
                                                        v59 := GenS6{F0: (v46 * (v46 & (v46 + 1))), F1: 1}
                                                        v1 += v1
                                                        _ = v59
                                                    }
                                                    v20 += ((0.1 - 3.14159) / v20)
                                                default:
                                                    v60 := GenS3{F0: (!(65535 < 2)), F1: v46, F2: ("hello" == "")}
                                                    _ = v60
                                                }
                                            case "foo bar":
                                                if (!(v46 == v46)) {
                                                    v12 -= v41
                                                } else {
                                                    switch fi42 {
                                                    case 32768:
                                                        if (v46 < v46) {
                                                            var v61 int64 = -1
                                                            if (!false) {
                                                                _ = len(v6)
                                                                var v62 uint = 255
                                                                _ = v62
                                                            }
                                                            if true {
                                                                v41 = v41
                                                                v41 += 98
                                                                switch fi45 {
                                                                case 9223372036854775807:
                                                                    v2 = v2
                                                                    var v63 int16 = (v2 - 0)
                                                                    if (!(557 != 1)) {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            var v64 int32 = 225
                                                                            _ = v64
                                                                        case "foo bar":
                                                                            var v65 uint32 = 1
                                                                            _ = v65
                                                                        default:
                                                                            var v66 int = fi45
                                                                            _ = v66
                                                                        }
                                                                    } else {
                                                                        v67 := GenS7{F0: -9223372036854775808}
                                                                        if (!true) {
                                                                            v68 := v5[0]
                                                                            v5[(v68 >> ((88) & 63))] = (v68 + (v68 - 65535))
                                                                            v2 = v63
                                                                            var v69 int32 = 38
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                    _ = v63
                                                                }
                                                            }
                                                            _ = v61
                                                        } else {
                                                            v51 = v51
                                                            var v70 uint32 = 2147483648
                                                            _ = v70
                                                        }
                                                    case 32767:
                                                        var v71 string = ("a" + "a")
                                                        v41 = (v12 >> ((169) & 63))
                                                        v1 = v1
                                                        _ = v71
                                                    case 0:
                                                        var v72 int = 127
                                                        v3 += (v3 / ((18446744073709551615) | 1))
                                                        _ = v72
                                                    default:
                                                        v1 = v1
                                                    }
                                                    var v73 uint = 42
                                                    _ = v73
                                                }
                                            default:
                                                if true {
                                                    switch int(v2) {
                                                    case 495:
                                                        if false {
                                                            v2 += (v2 >> ((642) & 63))
                                                            var v74 uint64 = (v3 ^ 436)
                                                            _ = v74
                                                        } else {
                                                            v2 = (v1 << ((65535) & 63))
                                                            v75 := GenS6{F0: (v46 << ((2) & 63)), F1: 2}
                                                            _ = v75
                                                        }
                                                    case -364:
                                                        v76 := GenS7{F0: v46}
                                                        fmt.Println(v3)
                                                        _ = v76
                                                    case -804:
                                                        v11 = true
                                                    }
                                                    v77 := GenS7{F0: v46}
                                                    _ = v77
                                                } else {
                                                    switch (fi42 | (^(fi45))) {
                                                    case 0:
                                                        if v11 {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v2)
                                                            case "hello":
                                                                var v78 uint32 = 1
                                                                _ = v78
                                                            default:
                                                                if v11 {
                                                                    _ = len(v6)
                                                                    v79 := GenS8{F0: -173, F1: v46}
                                                                    _ = v79
                                                                }
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        delete(v5, 2147483647)
                                                    case 42:
                                                        var v80 string = "\n\t"
                                                        _ = v80
                                                    default:
                                                        switch "foo bar" {
                                                        case "":
                                                            v81 := GenS4{F0: 331, F1: 867, F2: -1}
                                                            _ = v81
                                                        case "a":
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    }
                                                    var v82 int64 = (v46 | (v46 - v46))
                                                    _ = len(v5)
                                                    _ = v82
                                                }
                                            }
                                        }
                                        var v83 uint8 = 42
                                        fmt.Println(v4)
                                        var v84 int32 = -1
                                        var v85 float64 = (v4 / 1.0)
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v46
                                    _ = v51
                                }
                                if v11 {
                                    delete(v5, 4294967295)
                                    switch ((("foo bar" + "hello") + (("hello" + "") + (("hello" + "foo bar") + (("a" + "a") + ("\n\t" + "hello"))))) + "\n\t") {
                                    case "a":
                                        v20 -= v20
                                    }
                                }
                                _ = fi45
                            }
                        }
                    } else {
                        var v86 int32 = 2147483647
                        var v87 uint = 42
                        _ = v86
                        _ = v87
                    }
                    fmt.Println(v41)
                    var v88 uint64 = v3
                    _ = v88
                }
            } else {
                for fi89 := 0; fi89 < 15; fi89++ {
                    if v11 {
                        var v90 uint16 = uint16(v12)
                        v12 = 127
                        v91 := GenS9{F0: 2, F1: ((("\n\t" + ("a" + ("hello" + "\n\t"))) + "foo bar") + "hello"), F2: (^(v41))}
                        _ = v90
                        _ = v91
                    }
                    var v92 uint64 = v3
                    _ = len(v6)
                    _ = fi89
                    _ = v92
                }
            }
            v93 := GenS6{F0: 436, F1: 4294967295}
            _ = fi42
            _ = v93
        }
        _ = v41
    }
    var v94 uint = 644
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
}
