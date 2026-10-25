package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int16; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint; F1 uint; F2 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool) bool {
    if (32767 > 42) {
        switch "a" {
        case "a":
            var v5 int32 = 1
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: 2}
            _ = v6
        case "hello":
            switch "" {
            case "":
                v7 := GenS1{F0: 2147483647}
                _ = v7
            case "\n\t":
                p0 = ((846 < 255) || true)
            }
        default:
            fmt.Println(p0)
        }
        var v8 int32 = 42
        p0 = (-1 <= 32767)
        _ = v8
    }
    v9 := []int{-785}
    if true {
        var v10 int64 = 749
        v11 := GenS2{F0: -2147483648, F1: 0}
        switch 105 {
        case -898:
            var v12 uint = 1
            _ = v12
        case 32767:
            switch 1 {
            case -361:
                var v13 uint8 = 42
                _ = v13
            case -455:
                for fi14 := 0; fi14 < 17; fi14++ {
                    if ("a" != ("hello" + "a")) {
                        p0 = (42 >= 9223372036854775807)
                    } else {
                        p0 = (2 <= 65535)
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    fmt.Println(v10, p0)
                    _ = fi14
                }
            }
        }
        _ = v10
        _ = v11
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = v9
    return false
}
func genFn1(p0 int64) uint64 {
    switch "\n\t" {
    case "foo bar":
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS3{F0: fi15}
            var v17 int64 = p0
            v17 -= (p0 % ((p0) | 1))
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    p0 += (p0 & 42)
    _ = p0
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.1
    var v3 int16 = 32767
    var v4 float64 = 1.5
    v1[0] = -17
    if (0 < -1) {
        if true {
            fmt.Println(v4)
            switch "hello" {
            case "hello":
                v2 += v2
            default:
                switch ("hello" + ("foo bar" + "a")) {
                case "foo bar":
                    var v18 int32 = 42
                    _ = v18
                case "hello":
                    if ((-1 >= 42) && true) {
                        if ((false && false) && (32768 == 42)) {
                            v19 := GenS4{F0: "\n\t"}
                            if false {
                                v2 -= -2.25
                                if false {
                                    var v20 string = ""
                                    var v21 uint = 255
                                    v22 := GenS3{F0: 592}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 int16 = v3
                                    _ = v23
                                }
                            }
                            _ = v19
                        } else {
                            v4 += (0.0 / (((v4 + 3.14159) * ((float64(v2) / v4) * v4)) - 3.14159))
                            v24 := GenS2{F0: 0, F1: 42}
                            _ = v24
                        }
                        v25 := GenS5{F0: 1, F1: 1, F2: 471}
                        v26 := []int{525, -771, 847}
                        _ = v25
                        _ = v26
                    }
                case "":
                    v27 := GenS4{F0: ("hello" + ("" + ""))}
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 14; fi28++ {
                fmt.Println(fi28)
                _ = fi28
            }
        }
    }
    var v29 bool = (2 < genFn1(92))
    v29 = false
    _ = len(v1)
    var v30 uint64 = 0
    if ((true && v29) && (0 >= 2147483647)) {
        for fi31 := 0; fi31 < 5; fi31++ {
            _ = len(v1)
            v32 := v1[706]
            if false {
                fmt.Println(v29)
                v4 = 0.0
            }
            _ = fi31
            _ = v32
        }
        v33 := []int{127}
        if len(v33) > 0 {
            v33[0] = 681
        }
        _ = v33
    }
    if (((3.14159 + (-2.25 - -2.25)) - v2) == float32(v4)) {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int8 = 100
            if v29 {
                var v36 uint8 = 127
                v2 = -2.25
                v37 := GenS3{F0: (fi34 ^ fi34)}
                _ = v36
                _ = v37
            } else {
                switch 128 {
                case 69:
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := v1[918]
                        v2 = v2
                        _ = fi38
                        _ = v39
                    }
                default:
                    var v40 float32 = v2
                    _ = v40
                }
                v29 = (54 <= 32767)
                var v41 int8 = v35
                _ = v41
            }
            fmt.Println(v4, fi34, v35)
            _ = fi34
            _ = v35
        }
        var v42 int64 = 170
        v1[42] = -1
        _ = v42
    }
    if v29 {
        v43 := GenS6{F0: 9223372036854775807, F1: genFn1(int64(v30))}
        _ = v43
    }
    var v44 bool = (!v29)
    v1[498] = -403
    v45 := []int{}
    var v46 int = 9223372036854775807
    if (1 == 236) {
        var v47 int64 = -1
        _ = v47
    } else {
        fmt.Println(v46, v46, v29)
    }
    var v48 int = v46
    v30 = genFn1(9223372036854775807)
    for fi49 := 0; fi49 < 4; fi49++ {
        if false {
            v50 := GenS7{F0: 875, F1: 2, F2: "foo bar"}
            if v44 {
                for fi51 := 0; fi51 < 11; fi51++ {
                    if false {
                        if (!((!v44) || ((1e10 * float32(v4)) < v2))) {
                            var v52 bool = v29
                            _ = v52
                        } else {
                            var v53 float64 = (v4 + v4)
                            v44 = v29
                            _ = v53
                        }
                    } else {
                        v44 = v29
                        var v54 uint32 = 65535
                        _ = v54
                    }
                    _ = fi51
                }
                for fi55 := 0; fi55 < 4; fi55++ {
                    v56 := GenS6{F0: -9223372036854775808, F1: (v30 & v30)}
                    var v57 uint16 = 65535
                    if len(v45) > 0 {
                        v45[0] = fi55
                    }
                    v2 = (3.14159 - (0.0 - float32(v4)))
                    if (!(v29 && v29)) {
                        v58 := v1[2147483648]
                        _ = v58
                    } else {
                        var v59 int64 = -1
                        var v60 int = int(v3)
                        _ = len(v45)
                        _ = v59
                        _ = v60
                    }
                    v61 := GenS5{F0: (-(fi55)), F1: v3, F2: (v57 / ((v57) | 1))}
                    if (!("hello" == "hello")) {
                        var v62 float32 = float32(v4)
                        v3 += v3
                        fmt.Println(v30, v44)
                        _ = v62
                    } else {
                        switch int(v30) {
                        case 1:
                            var v63 int16 = -32768
                            _ = v63
                        case 2147483647:
                            v4 = (1e10 * ((0.0 - ((((float64(1e10) * (v4 + 3.14159)) + v4) / float64((-2.25 + v2))) * (float64(v2) * v4))) - v4))
                        }
                    }
                    v46 -= (v46 + fi49)
                    if false {
                        v48 = (fi55 - int(v3))
                        var v64 int32 = -2147483648
                        fmt.Println(v44, fi49)
                        _ = v64
                    } else {
                        var v65 uint16 = 42
                        if v44 {
                            _ = len(v45)
                            if (v57 == v57) {
                                v66 := GenS6{F0: 1, F1: 0}
                                _ = v66
                            } else {
                                var v67 uint = 65535
                                if (v67 > v67) {
                                    if ("\n\t" >= "foo bar") {
                                        v67 -= (v67 >> ((v67) & 63))
                                        var v68 float64 = (3.14159 / v4)
                                        v46 = v46
                                        _ = v68
                                    }
                                    var v69 int32 = 0
                                    if (!true) {
                                        _ = len(v45)
                                        v30 += v30
                                        var v70 float64 = (0.1 / (v4 + (-2.25 - float64(v2))))
                                        _ = v70
                                    } else {
                                        if false {
                                            fmt.Println(v65)
                                            var v71 float32 = 0.0
                                            _ = v71
                                        } else {
                                            var v72 uint16 = v65
                                            v67 = (v67 + (v67 - uint(v30)))
                                            if true {
                                                v45 = append(v45, -903)
                                                var v73 int16 = (v3 % (((v3 & v3)) | 1))
                                                _ = v73
                                            } else {
                                                v74 := GenS8{F0: -1}
                                                var v75 uint32 = 4294967295
                                                var v76 float32 = (float32(0.0) * v2)
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            _ = v72
                                        }
                                        var v77 float32 = (0.1 + (v2 * float32(1.5)))
                                        _ = v77
                                    }
                                    var v78 string = "a"
                                    if len(v45) > 0 {
                                        v45[0] = -2147483648
                                    }
                                    _ = v69
                                    _ = v78
                                }
                                v79 := GenS7{F0: (v67 - v67), F1: (^(v67)), F2: ("" + ("\n\t" + ((("a" + "a") + (((("" + "a") + "\n\t") + "") + ("" + "hello"))) + "a")))}
                                switch fi49 {
                                case 42:
                                    switch "hello" {
                                    case "":
                                        if true {
                                            v57 = 1
                                            v48 -= 20
                                        }
                                    default:
                                        var v80 uint8 = 231
                                        _ = v80
                                    }
                                case -793:
                                    fmt.Println(fi55, fi55, fi55)
                                case 287:
                                    switch 23 {
                                    case 0:
                                        var v81 uint16 = v57
                                        _ = v81
                                    default:
                                        var v82 int8 = 1
                                        var v83 int16 = v3
                                        v84 := GenS7{F0: uint(v30), F1: v67, F2: ""}
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    fmt.Println(v67)
                                    v48 += 1
                                }
                                if v44 {
                                    v85 := GenS9{F0: 0}
                                    if v44 {
                                        v57 = (v65 | (^(v65)))
                                        var v86 uint16 = 65535
                                        v86 += (v57 << (((v67 | (v67 + v67))) & 63))
                                        _ = v86
                                    } else {
                                        fmt.Println(v44, v2)
                                        v87 := GenS10{F0: 768, F1: 42}
                                        fmt.Println(v4, v67, v48)
                                        v67 -= (v67 / ((v67) | 1))
                                        var v88 int8 = 100
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = v85
                                }
                                _ = v67
                                _ = v79
                            }
                            delete(v1, 2)
                        }
                        _ = v65
                    }
                    switch "hello" {
                    case "hello":
                        v2 = (float32((0.1 - ((((3.14159 + (float64(1.5) + 0.0)) + 1e10) / ((float64(3.14159) / v4) * 3.14159)) / float64(1.5)))) + v2)
                    case "\n\t":
                        v29 = (-43 < 127)
                    }
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v61
                }
                var v89 bool = v44
                _ = v89
            }
            v90 := GenS8{F0: 967}
            _ = v50
            _ = v90
        } else {
            var v91 float64 = v4
            fmt.Println(v3, v46)
            _ = v91
        }
        var v92 float32 = -2.25
        var v93 int8 = 42
        _ = fi49
        _ = v92
        _ = v93
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v48
}
