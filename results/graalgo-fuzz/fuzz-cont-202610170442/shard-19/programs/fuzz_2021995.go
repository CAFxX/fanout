package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32 }
func genFn0(p0 string) string {
    fmt.Println(p0)
    v7 := GenS1{F0: 1, F1: 0, F2: "hello"}
    var v8 uint8 = 0
    v8 += 0
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v8 = v8
            if true {
                switch fi10 {
                case 297:
                    var v11 int = 32768
                    _ = v11
                case -576:
                    v12 := []int{fi9, (fi9 ^ (-(fi10))), -39, 32767}
                    _ = v12
                case 32767:
                    v13 := GenS1{F0: 4294967295, F1: 255, F2: "foo bar"}
                    _ = v13
                }
                fmt.Println(p0, fi9)
                if (!false) {
                    v14 := GenS1{F0: 1, F1: 32767, F2: "foo bar"}
                    v15 := []int{(^(fi10))}
                    _ = v14
                    _ = v15
                }
            }
            if (true && false) {
                if false {
                    var v16 string = ("a" + p0)
                    _ = v16
                }
            } else {
                switch "a" {
                case "\n\t":
                    fmt.Println(fi10)
                default:
                    v17 := []int{fi10}
                    _ = v17
                }
                if false {
                    v18 := GenS1{F0: 256, F1: 65535, F2: (p0 + p0)}
                    fmt.Println(fi10, p0, p0)
                    switch 128 {
                    case -923:
                        if true {
                            switch fi10 {
                            case -600:
                                fmt.Println(fi9, p0)
                            case -983:
                                var v19 int16 = 1
                                _ = v19
                            case 128:
                                var v20 int = fi10
                                _ = v20
                            }
                            switch (p0 + "hello") {
                            case "a":
                                v8 = v8
                            }
                        }
                        v21 := GenS2{F0: 2147483648, F1: ""}
                        _ = v21
                    case 490:
                        var v22 bool = true
                        _ = v22
                    case 796:
                        if ((-1 < 42) && (465 > 42)) {
                            fmt.Println(v8, fi10, v8)
                        } else {
                            fmt.Println(fi10, fi9, v8)
                        }
                    }
                    _ = v18
                } else {
                    fmt.Println(fi10, p0, p0)
                    var v23 float64 = (((0.0 + (3.14159 / 1e10)) * -2.25) - 0.1)
                    _ = v23
                }
                if (uint16(v8) >= 1) {
                    v8 += 1
                    if (996 != -782) {
                        switch p0 {
                        case "\n\t":
                            v24 := GenS3{F0: 2}
                            _ = v24
                        case "":
                            switch ((("a" + (p0 + "\n\t")) + p0) + p0) {
                            case "a":
                                p0 = "foo bar"
                            case "hello":
                                p0 = ""
                            case "foo bar":
                                if (1 > -1) {
                                    p0 = p0
                                    var v25 float32 = (3.14159 * 1.5)
                                    _ = v25
                                } else {
                                    v26 := GenS3{F0: 904}
                                    _ = v26
                                }
                            }
                        default:
                            var v27 int = fi9
                            _ = v27
                        }
                    }
                    if false {
                        var v28 string = ((("foo bar" + (((p0 + p0) + "a") + (p0 + p0))) + "") + (("a" + "\n\t") + ("" + p0)))
                        var v29 int16 = 1
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 65535, F1: 32767, F2: ""}
                        v31 := []int{(fi9 & 907), 32768, fi10}
                        _ = len(v31)
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 uint16 = 0
                    _ = v32
                }
            }
            var v33 int64 = 737
            var v34 int16 = -32768
            v8 = (^(v8))
            _ = fi10
            _ = v33
            _ = v34
        }
        v35 := []int{(fi9 / ((fi9) | 1)), (fi9 - 382), 4294967295}
        _ = fi9
        _ = v35
    }
    _ = p0
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int16, p1 string) string {
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := []int{}
        _ = fi36
        _ = v37
    }
    if false {
        if ((false || false) || (535 == 851)) {
            var v38 uint16 = 255
            _ = v38
        }
    }
    fmt.Println(p1, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return ((p1 + "") + p1)
}

func main() {
    var v1 bool = true
    var v2 int16 = 0
    v3 := []string{("foo bar" + "hello")}
    var v4 int64 = -9223372036854775808
    v5 := []string{}
    v6 := make(map[uint32]uint32)
    v39 := GenS3{F0: 548}
    v5 = append(v5, "hello")
    v1 = v1
    var v40 uint16 = 2
    if true {
        for fi41 := 0; fi41 < 7; fi41++ {
            if false {
                v3 = append(v3, "")
                for fi42 := 0; fi42 < 20; fi42++ {
                    v40 = 2
                    var v43 uint16 = (v40 << ((45) & 63))
                    _ = fi42
                    _ = v43
                }
            }
            v4 = v4
            fmt.Println(v2, v1)
            _ = fi41
        }
    } else {
        var v44 float32 = float32(float64((-2.25 * 3.14159)))
        _ = v44
    }
    var v45 uint32 = 636
    switch -9223372036854775808 {
    case 342:
        switch "\n\t" {
        case "hello":
            if false {
                var v46 int = 108
                if false {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        delete(v6, v45)
                        v48 := GenS2{F0: (v45 ^ v45), F1: ""}
                        v49 := GenS2{F0: 42, F1: ("" + "hello")}
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    var v50 int8 = -1
                    _ = v50
                }
                switch 32767 {
                case -750:
                    fmt.Println(v40, v46)
                case 230:
                    if v1 {
                        if ((v45 & v45) != v45) {
                            var v51 int16 = v2
                            if v1 {
                                v52 := GenS2{F0: v45, F1: ("foo bar" + "")}
                                v40 -= 0
                                v53 := GenS4{F0: 1, F1: 1}
                                _ = v52
                                _ = v53
                            }
                            _ = v51
                        } else {
                            fmt.Println(v1)
                            switch "hello" {
                            case "hello":
                                var v54 int16 = v2
                                _ = v54
                            case "\n\t":
                                if (!v1) {
                                    var v55 int64 = (-(v4))
                                    _ = v55
                                }
                            default:
                                switch (("\n\t" + ("a" + ("" + "hello"))) + "hello") {
                                case "":
                                    _ = len(v5)
                                default:
                                    var v56 float64 = 0.0
                                    _ = v56
                                }
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v2 = (v2 & v2)
                    } else {
                        fmt.Println(v45, v40, v1)
                        switch v46 {
                        case 869:
                            if v1 {
                                v57 := GenS1{F0: (v45 & v45), F1: v40, F2: "\n\t"}
                                var v58 float64 = 1.5
                                var v59 int8 = 1
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        default:
                            for fi60 := 0; fi60 < 19; fi60++ {
                                var v61 float32 = float32(1.5)
                                for fi62 := 0; fi62 < 4; fi62++ {
                                    if (!(0 > 1)) {
                                        var v63 float64 = float64((v61 / v61))
                                        if (float64(v61) == (float64(v61) - 1.5)) {
                                            v4 += (v4 / (((v4 >> ((65535) & 63))) | 1))
                                            if v1 {
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            switch (v46 >> ((uint(v40)) & 63)) {
                                            case 238:
                                                var v64 bool = v1
                                                _ = v64
                                            default:
                                                var v65 bool = (v1 && v1)
                                                _ = v65
                                            }
                                        }
                                        var v66 int16 = (-(v2))
                                        _ = v63
                                        _ = v66
                                    }
                                    var v67 uint16 = v40
                                    v40 = 42
                                    _ = fi62
                                    _ = v67
                                }
                                _ = fi60
                                _ = v61
                            }
                        }
                        if v1 {
                            v68 := GenS5{F0: v45, F1: "\n\t"}
                            var v69 int16 = 216
                            var v70 uint64 = 1
                            _ = v68
                            _ = v69
                            _ = v70
                        }
                    }
                }
                _ = v46
            } else {
                fmt.Println(v1, v4)
                for fi71 := 0; fi71 < 19; fi71++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v40 += 153
                    for fi72 := 0; fi72 < 18; fi72++ {
                        var v73 float64 = 0.0
                        _ = fi72
                        _ = v73
                    }
                    _ = fi71
                }
                delete(v6, (v45 >> ((758) & 63)))
            }
        default:
            for fi74 := 0; fi74 < 18; fi74++ {
                for fi75 := 0; fi75 < 8; fi75++ {
                    var v76 uint64 = 0
                    v77 := GenS2{F0: v45, F1: "\n\t"}
                    switch "a" {
                    case "foo bar":
                        v78 := GenS5{F0: (^(v45)), F1: ("foo bar" + ("" + ("" + (("\n\t" + "") + "a"))))}
                        _ = v78
                    default:
                        v79 := GenS4{F0: 2, F1: -1}
                        _ = v79
                    }
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                _ = fi74
            }
        }
    default:
        v3 = append(v3, ((("a" + "") + "\n\t") + ""))
    }
    var v80 uint = 2
    v4 = 26
    v80 = (^(v80))
    for fi81 := 0; fi81 < 10; fi81++ {
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v40 = (v40 - 255)
        }
        var v82 string = "a"
        _ = fi81
        _ = v82
    }
    v5 = append(v5, "a")
    if (v1 || v1) {
        v80 = (v80 ^ (^(v80)))
    } else {
        v80 -= v80
        v83 := GenS5{F0: (^(v45)), F1: "\n\t"}
        _ = v83
    }
    for fi84 := 0; fi84 < 15; fi84++ {
        _ = len(v6)
        v2 += (v2 & (v2 ^ v2))
        _ = fi84
    }
    fmt.Println(v45, v1)
    if false {
        switch "\n\t" {
        case "\n\t":
            var v85 bool = true
            _ = v85
        case "a":
            switch 2147483647 {
            case -9223372036854775808:
                var v86 uint32 = 2
                _ = v86
            case -1:
                var v87 int = -2147483648
                _ = v87
            case 798:
                v88 := GenS6{F0: (v2 + v2), F1: (v2 * v2)}
                _ = v88
            }
        case "hello":
            v89 := v6[v45]
            _ = v89
        }
        if v1 {
            var v90 int32 = 121
            if ((19 != 62) || (uint16(v45) <= v40)) {
                if v1 {
                    for fi91 := 0; fi91 < 7; fi91++ {
                        if (((v4 / ((int64(v90)) | 1)) == (v4 ^ v4)) && (!(1.5 == 1.5))) {
                            for fi92 := 0; fi92 < 5; fi92++ {
                                if v1 {
                                    var v93 uint32 = (v45 % (((v45 * 541)) | 1))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v93
                                } else {
                                    var v94 uint = 291
                                    var v95 bool = ((v40 >> ((v80) & 63)) <= 32767)
                                    v96 := GenS7{F0: 127, F1: 2}
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                                v2 += v2
                                _ = fi92
                            }
                            for fi97 := 0; fi97 < 18; fi97++ {
                                v80 += v80
                                v98 := v6[(^(v45))]
                                v5 = append(v5, "foo bar")
                                fmt.Println(v98, v90)
                                _ = fi97
                                _ = v98
                            }
                        } else {
                            v4 = (v4 ^ v4)
                        }
                        var v99 uint32 = (v45 >> (((v80 / ((v80) | 1))) & 63))
                        _ = fi91
                        _ = v99
                    }
                    for fi100 := 0; fi100 < 18; fi100++ {
                        switch "foo bar" {
                        case "a":
                            v101 := GenS8{F0: 16}
                            _ = v101
                        }
                        _ = fi100
                    }
                    fmt.Println(v80)
                } else {
                    for fi102 := 0; fi102 < 14; fi102++ {
                        v103 := GenS7{F0: 255, F1: 860}
                        if len(v3) > 0 {
                            v3[0] = (("foo bar" + ((("foo bar" + "\n\t") + "foo bar") + ("hello" + "foo bar"))) + "")
                        }
                        v104 := v6[v45]
                        _ = fi102
                        _ = v103
                        _ = v104
                    }
                    fmt.Println(v45)
                }
                v105 := GenS9{F0: 0}
                _ = v105
            }
            v90 = v90
            _ = v90
        }
    }
    for fi106 := 0; fi106 < 14; fi106++ {
        v107 := GenS4{F0: 128, F1: 535}
        fmt.Println(fi106)
        _ = len(v6)
        _ = fi106
        _ = v107
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
    _ = v80
}
