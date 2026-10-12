package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint32; F2 int8 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8 }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: 65535, F1: 1, F2: 9223372036854775807}
    p0 += 0
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (^(fi7)) {
        case 424:
            var v8 int16 = 42
            _ = v8
        default:
            p0 = (p0 % (((fi7 + 32767)) | 1))
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 int, p1 uint16) uint8 {
    var v9 uint16 = p1
    switch 42 {
    case 0:
        v10 := GenS1{F0: v9, F1: 279, F2: -9223372036854775808}
        _ = v10
    case 1:
        var v11 float64 = 0.1
        _ = v11
    default:
        var v12 string = "\n\t"
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v9, fi13)
        var v14 uint16 = (p1 & (p1 * 65535))
        var v15 uint8 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    switch ((("foo bar" + "foo bar") + "\n\t") + "\n\t") {
    case "a":
        v9 += p1
    case "":
        if true {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := []int{fi16, (fi16 | 179), p0}
                p1 = (v9 >> ((18446744073709551615) & 63))
                fmt.Println(p1, fi16)
                _ = fi16
                _ = v17
            }
            var v18 int8 = 34
            _ = v18
        }
    default:
        var v19 uint16 = (p1 & 255)
        _ = v19
    }
    if (42 > 0) {
        var v20 int64 = 0
        switch p0 {
        case 32768:
            p0 += p0
        }
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            if true {
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int32 = 42
                    if false {
                        if true {
                            v23 += -2147483648
                        } else {
                            fmt.Println(p1, fi21, v23)
                        }
                        v24 := []int{fi22, 127, p0}
                        v25 := GenS1{F0: (v9 % (((^(v9))) | 1)), F1: 4294967295, F2: int64(fi22)}
                        _ = v24
                        _ = v25
                    }
                    v23 = v23
                    _ = fi22
                    _ = v23
                }
                var v26 int = 32767
                _ = v26
            }
            p0 += 4294967295
            _ = fi21
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 127
}

func main() {
    var v1 uint = 2
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    var v4 float32 = 3.14159
    var v5 string = "a"
    for fi27 := 0; fi27 < 8; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            switch 32767 {
            case 128:
                fmt.Println(v4, v4)
                v5 = v5
                var v29 bool = false
                _ = v29
            case 1:
                fmt.Println(fi27)
            }
            _ = fi28
        }
        v1 = v1
        _ = fi27
    }
    v30 := []int{}
    if false {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint8 = genFn1(-9223372036854775808, 255)
            var v33 uint = (^(v1))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    if ((!(127 < 1)) || false) {
        var v34 int32 = 0
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v5)
        var v36 float32 = (1.5 / v4)
        _ = fi35
        _ = v36
    }
    switch -292 {
    case 32768:
        var v37 float64 = -2.25
        _ = v37
    case -168:
        var v38 int8 = -11
        _ = v38
    default:
        switch -440 {
        case 0:
            v39 := GenS2{F0: 311, F1: ("" + v5)}
            _ = v39
        case -9223372036854775808:
            var v40 int32 = -2147483648
            _ = v40
        default:
            var v41 uint = (v1 % (((v1 - 1)) | 1))
            _ = v41
        }
    }
    v4 = v4
    if false {
        switch 42 {
        case -9223372036854775808:
            v1 = 2
        case -198:
            var v42 int8 = -20
            _ = v42
        default:
            switch "\n\t" {
            case "a":
                for fi43 := 0; fi43 < 13; fi43++ {
                    if (-128 <= 0) {
                        v44 := GenS2{F0: 127, F1: ("hello" + v5)}
                        for fi45 := 0; fi45 < 16; fi45++ {
                            if len(v30) > 0 {
                                v30[0] = (fi43 / (((fi43 % (((^(fi43))) | 1))) | 1))
                            }
                            if (1 < 42) {
                                if ((^(fi45)) <= 798) {
                                    var v46 float32 = float32(((0.0 * 1.5) - 3.14159))
                                    var v47 int64 = 0
                                    _ = v46
                                    _ = v47
                                } else {
                                    fmt.Println(v5, fi45)
                                    var v48 uint = (^(v1))
                                    v49 := GenS3{F0: (!(false && true)), F1: 0}
                                    _ = v48
                                    _ = v49
                                }
                                switch "hello" {
                                case "":
                                    switch fi43 {
                                    case -1:
                                        var v50 uint8 = 0
                                        _ = v50
                                    case 32768:
                                        switch -171 {
                                        case 32767:
                                            v4 = float32((1e10 / (3.14159 * 1.5)))
                                            v4 -= 1e10
                                        case -2147483648:
                                            var v51 uint64 = 42
                                            _ = v51
                                        case 11:
                                            if ((false && (!false)) && (1 != 1)) {
                                                v52 := GenS3{F0: false, F1: 42}
                                                var v53 uint32 = 2147483647
                                                v54 := GenS4{F0: 18446744073709551615}
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            } else {
                                                var v55 int16 = -1
                                                switch v5 {
                                                case "hello":
                                                    switch v5 {
                                                    case "foo bar":
                                                        if (false && (1 != -2147483648)) {
                                                            v2 = (v2 + genFn0(fi43))
                                                        } else {
                                                            var v56 int64 = v2
                                                            _ = v56
                                                        }
                                                        if len(v30) > 0 {
                                                            v30[0] = 0
                                                        }
                                                    case "a":
                                                        if false {
                                                            fmt.Println(v4, fi45, v4)
                                                            v57 := GenS1{F0: 111, F1: 30, F2: genFn0((fi45 >> ((v1) & 63)))}
                                                            _ = v57
                                                        } else {
                                                            v58 := GenS5{F0: 138, F1: 1}
                                                            _ = v58
                                                        }
                                                    case "":
                                                        switch fi45 {
                                                        case 42:
                                                            if false {
                                                                v4 += v4
                                                                var v59 int = 4294967295
                                                                _ = len(v30)
                                                                v5 = v5
                                                                v2 += (v2 + genFn0(int(v55)))
                                                                _ = v59
                                                            } else {
                                                                v1 += v1
                                                            }
                                                        case 0:
                                                            v60 := GenS6{F0: (fi43 <= (-(fi45))), F1: 4294967295, F2: int8(v55)}
                                                            _ = v60
                                                        default:
                                                            if len(v30) >= 2 {
                                                                v30 = v30[1:len(v30)]
                                                            }
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v61 int = (fi45 << (((v1 | 1)) & 63))
                                                    _ = v61
                                                }
                                                _ = v55
                                            }
                                        default:
                                            v62 := GenS5{F0: genFn1(fi43, 0), F1: 42}
                                            _ = v62
                                        }
                                    default:
                                        var v63 int16 = -32768
                                        _ = v63
                                    }
                                case "\n\t":
                                    if (!false) {
                                        if (2 > 42) {
                                            var v64 uint16 = 255
                                            v4 = v4
                                            _ = v64
                                        }
                                        v30 = append(v30, (fi43 & (^(fi43))))
                                        var v65 string = "hello"
                                        _ = v65
                                    }
                                }
                            } else {
                                v66 := GenS4{F0: 18446744073709551615}
                                var v67 uint = (v1 | (v1 << (((^(v1))) & 63)))
                                v68 := GenS7{F0: 2147483647, F1: (v1 % ((v1) | 1))}
                                v69 := GenS8{F0: -1}
                                var v70 bool = false
                                _ = v66
                                _ = v67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            if (false || (((float32(0.0) - v4) >= (float32(3.14159) * v4)) || (197 >= 80))) {
                                var v71 uint8 = 202
                                v71 += genFn1(909, uint16(v1))
                                v5 = v5
                                _ = v71
                            }
                            _ = fi45
                        }
                        _ = v44
                    } else {
                        var v72 uint64 = 42
                        v73 := GenS6{F0: true, F1: 27, F2: 100}
                        var v74 uint64 = 399
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    v2 += v2
                    _ = fi43
                }
            }
        }
        var v75 uint16 = 2
        _ = v75
    }
    for fi76 := 0; fi76 < 7; fi76++ {
        if false {
            for fi77 := 0; fi77 < 4; fi77++ {
                v78 := GenS9{F0: 4294967295}
                _ = fi77
                _ = v78
            }
            v79 := GenS4{F0: 944}
            _ = v79
        }
        if (4294967295 > 42) {
            var v80 int64 = v2
            if (127 == 2) {
                var v81 float32 = ((1.5 - 3.14159) / (1e10 - float32(0.1)))
                if true {
                    switch -2147483648 {
                    case -9223372036854775808:
                        if ((v1 / (((v1 / ((v1) | 1))) | 1)) != v1) {
                            var v82 int32 = 42
                            v83 := GenS10{F0: int8(v1)}
                            _ = v82
                            _ = v83
                        } else {
                            for fi84 := 0; fi84 < 20; fi84++ {
                                var v85 int = 32767
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v3)
                                case "a":
                                    v86 := v3[4294967295]
                                    _ = v86
                                case "":
                                    v2 += (v2 | v2)
                                }
                                _ = fi84
                                _ = v85
                            }
                            delete(v3, 948)
                        }
                    case 1:
                        var v87 int16 = -1
                        _ = v87
                    case 42:
                        var v88 uint = (v1 % ((42) | 1))
                        _ = v88
                    default:
                        var v89 uint16 = 2
                        _ = v89
                    }
                    var v90 int16 = 781
                    v1 += 2
                    _ = v90
                }
                _ = v81
            }
            _ = v80
        }
        _ = fi76
    }
    v91 := GenS5{F0: uint8(v1), F1: 1}
    v92 := v3[2]
    v93 := GenS7{F0: v92, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v91
    _ = v92
    _ = v93
}
