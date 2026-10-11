package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 int64; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 int32; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint8; F1 int32; F2 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 string; F1 int16; F2 uint16 }
type GenS13 struct { F0 uint64; F1 uint; F2 uint16 }
func genFn0() bool {
    if (2 < 42) {
        v5 := []int{32767, 32768, 42}
        var v6 uint64 = 569
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
        _ = v6
    }
    var v7 uint32 = 0
    v8 := GenS1{F0: 702, F1: 255}
    switch 2147483647 {
    case 989:
        switch "\n\t" {
        case "\n\t":
            v9 := []int{}
            _ = v9
        case "a":
            if (true || false) {
                if (211 == 42) {
                    if true {
                        v7 -= (v7 >> ((2) & 63))
                        v7 = 2147483647
                    } else {
                        var v10 string = "a"
                        var v11 uint32 = 2147483648
                        _ = v10
                        _ = v11
                    }
                    v7 += (^(v7))
                    var v12 uint32 = 494
                    _ = v12
                }
                fmt.Println(v7)
            } else {
                var v13 float32 = (0.1 - -2.25)
                if (!(42 >= 2)) {
                    v14 := GenS1{F0: 127, F1: 865}
                    _ = v14
                }
                _ = v13
            }
        }
    case -204:
        v15 := []int{}
        _ = v15
    default:
        v7 = (v7 & 65535)
    }
    switch 962 {
    case -443:
        var v16 int = 2147483647
        _ = v16
    }
    _ = v7
    _ = v8
    return (-651 >= 958)
}
func genFn1(p0 int64, p1 uint8, p2 int) int32 {
    for fi17 := 0; fi17 < 14; fi17++ {
        p2 = (fi17 / ((-123) | 1))
        if false {
            v18 := GenS2{F0: 100, F1: p0, F2: ((("hello" + (("a" + "\n\t") + ("\n\t" + "a"))) + "foo bar") + "a")}
            var v19 int64 = p0
            _ = v18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(p0, p0)
    p0 -= (p0 ^ (p0 / (((p0 - p0)) | 1)))
    v20 := GenS1{F0: 0, F1: 400}
    if ("hello" > "\n\t") {
        v21 := GenS3{F0: -102}
        fmt.Println(p2, p2)
        p0 += 159
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return -1
}

func main() {
    var v1 bool = (0 < 2)
    var v2 string = "hello"
    var v3 uint32 = 356
    var v4 uint32 = (v3 + (v3 | v3))
    var v22 float32 = 1.5
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 float32 = v22
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if false {
            v26 := GenS4{F0: (v3 - (v3 + (v3 >> ((812) & 63)))), F1: 100, F2: uint64(v3)}
            _ = v26
        } else {
            if false {
                v27 := GenS4{F0: v3, F1: 90, F2: 506}
                v2 = "a"
                for fi28 := 0; fi28 < 1; fi28++ {
                    v29 := []int{-497, 32767, (fi28 - fi28)}
                    if v1 {
                        var v30 string = ""
                        if (v2 != "\n\t") {
                            fmt.Println(v30, v4, v3)
                        }
                        _ = v30
                    } else {
                        switch v2 {
                        case "foo bar":
                            var v31 string = v2
                            _ = v31
                        case "a":
                            v1 = (!(-69 >= -697))
                        case "\n\t":
                            v3 -= v4
                        }
                    }
                    var v32 string = v2
                    _ = fi28
                    _ = v29
                    _ = v32
                }
                _ = v27
            } else {
                var v33 uint = 689
                if v1 {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        switch "\n\t" {
                        case "a":
                            v4 = v3
                        case "foo bar":
                            var v35 float64 = 0.0
                            _ = v35
                        }
                        var v36 bool = (375 > 2)
                        fmt.Println(v3, fi25, v4)
                        _ = fi34
                        _ = v36
                    }
                    v4 += (^(v3))
                    v37 := GenS4{F0: 4294967295, F1: -27, F2: 0}
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := []int{0, 271, (^(fi38))}
                        v40 := GenS5{F0: 2}
                        v1 = v1
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 int64 = 1
                    _ = v41
                }
                if ((false && ((1e10 * (float32((0.0 - float64((v22 + v22)))) / -2.25)) < 0.1)) && v1) {
                    v42 := GenS6{F0: -1, F1: -1, F2: 0}
                    _ = v42
                }
                _ = v33
            }
        }
        fmt.Println(v3)
        _ = fi25
    }
    fmt.Println(v22, v3, v1)
    switch 4294967295 {
    case 128:
        v43 := GenS6{F0: -1, F1: -196, F2: int64(v4)}
        _ = v43
    case -2147483648:
        var v44 int16 = 1000
        _ = v44
    case -695:
        var v45 int16 = -32768
        _ = v45
    default:
        var v46 string = "hello"
        _ = v46
    }
    switch 127 {
    case -648:
        var v47 int8 = 127
        _ = v47
    case 4294967295:
        if v1 {
            v48 := GenS2{F0: 42, F1: -91, F2: ((v2 + "a") + v2)}
            for fi49 := 0; fi49 < 15; fi49++ {
                for fi50 := 0; fi50 < 2; fi50++ {
                    v3 -= v3
                    v51 := GenS5{F0: 234}
                    _ = fi50
                    _ = v51
                }
                v52 := GenS4{F0: v3, F1: 100, F2: 18446744073709551615}
                for fi53 := 0; fi53 < 6; fi53++ {
                    fmt.Println(v2, v1, fi49)
                    switch 853 {
                    case 42:
                        v3 = 626
                    default:
                        v54 := []int{}
                        _ = v54
                    }
                    if (-2.25 > -2.25) {
                        fmt.Println(fi53)
                        var v55 int64 = -846
                        var v56 int = fi49
                        v3 = v3
                        v57 := GenS6{F0: -128, F1: 2147483647, F2: -258}
                        _ = v55
                        _ = v56
                        _ = v57
                    } else {
                        if v1 {
                            if true {
                                v58 := GenS6{F0: -128, F1: genFn1(317, 128, fi53), F2: 9223372036854775807}
                                if (!(v1 && true)) {
                                    var v59 int64 = 42
                                    v60 := []int{}
                                    _ = v59
                                    _ = v60
                                } else {
                                    v61 := GenS7{F0: -9223372036854775808}
                                    v62 := GenS8{F0: 183, F1: fi49}
                                    var v63 int8 = 42
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                _ = v58
                            } else {
                                var v64 uint32 = v4
                                var v65 float32 = v22
                                switch v2 {
                                case "a":
                                    var v66 float32 = ((((1e10 + (v65 / ((float32((float64(3.14159) / -2.25)) * (((v65 * (float32((((-2.25 * 1e10) / (float64(1.5) * float64(float32((-2.25 * 1.5))))) + 1e10)) - v22)) + 1e10) + ((1e10 * v22) + 1.5))) - (v22 + v65)))) + v22) - v22) / v65)
                                    _ = v66
                                case "\n\t":
                                    fmt.Println(v1, v64, v64)
                                default:
                                    switch -1 {
                                    case 4294967295:
                                        v64 -= (v3 % (((^(v4))) | 1))
                                    case 1:
                                        v65 -= (float32(0.1) / v65)
                                    default:
                                        v67 := []int{(fi49 - fi53)}
                                        _ = v67
                                    }
                                }
                                var v68 uint8 = 2
                                v69 := GenS9{F0: (v68 % (((v68 * (v68 | v68))) | 1)), F1: -1, F2: (v22 != (v65 + -2.25))}
                                _ = v64
                                _ = v65
                                _ = v68
                                _ = v69
                            }
                        } else {
                            if v1 {
                                v70 := []int{}
                                var v71 uint = 0
                                var v72 uint8 = 208
                                if false {
                                    v72 = v72
                                    var v73 uint = (v71 + v71)
                                    v74 := GenS7{F0: int64(fi53)}
                                    _ = v73
                                    _ = v74
                                } else {
                                    var v75 int = fi49
                                    switch 0 {
                                    case -108:
                                        v70 = append(v70, fi53)
                                        var v76 uint32 = 753
                                        switch v2 {
                                        case "hello":
                                            if (v1 && true) {
                                                v77 := GenS10{F0: 18446744073709551615}
                                                _ = v77
                                            }
                                            var v78 uint32 = (v76 ^ 2)
                                            v22 = (v22 / v22)
                                            _ = v78
                                        case "":
                                            switch fi49 {
                                            case 2147483647:
                                                var v79 uint16 = 222
                                                _ = v79
                                            case -69:
                                                switch v75 {
                                                case 811:
                                                    v80 := GenS8{F0: -530, F1: (v75 >> ((2) & 63))}
                                                    _ = v80
                                                }
                                            }
                                        default:
                                            v70 = append(v70, v75)
                                        }
                                        _ = v76
                                    default:
                                        fmt.Println(fi53, v22, v72)
                                    }
                                    var v81 float32 = float32(1e10)
                                    _ = v75
                                    _ = v81
                                }
                                if v1 {
                                    v70 = append(v70, fi49)
                                }
                                _ = v70
                                _ = v71
                                _ = v72
                            } else {
                                var v82 int = 14
                                _ = v82
                            }
                            v83 := []int{(fi49 / ((42) | 1)), fi53, (fi53 + 128), -827}
                            v2 = v2
                            _ = v83
                        }
                        var v84 int8 = 14
                        _ = v84
                    }
                    if v1 {
                        v22 = (v22 * 1.5)
                    } else {
                        fmt.Println(v2)
                        v85 := []int{42, 128, 626, -564}
                        if (!v1) {
                            switch v2 {
                            case "":
                                v85 = append(v85, (fi53 * fi49))
                            default:
                                var v86 int16 = 32767
                                _ = v86
                            }
                            fmt.Println(v2, v3)
                            v4 = (v3 / (((^(v4))) | 1))
                        }
                        _ = v85
                    }
                    switch (v2 + "a") {
                    case "hello":
                        if false {
                            v87 := GenS11{F0: 100, F1: 1}
                            _ = v87
                        } else {
                            if (4294967295 >= 707) {
                                v88 := GenS12{F0: "", F1: 1000, F2: 1}
                                _ = v88
                            }
                            v4 += (^(v3))
                            var v89 uint64 = 0
                            _ = v89
                        }
                    }
                    _ = fi53
                }
                _ = fi49
                _ = v52
            }
            v2 = v2
            _ = v48
        } else {
            if v1 {
                fmt.Println(v4)
                v22 = (3.14159 + v22)
                v90 := GenS13{F0: 4294967295, F1: 66, F2: 1}
                _ = v90
            } else {
                for fi91 := 0; fi91 < 18; fi91++ {
                    v22 = 0.1
                    switch v2 {
                    case "":
                        var v92 int32 = -320
                        _ = v92
                    }
                    _ = fi91
                }
            }
            v93 := []int{603, -836, -212}
            var v94 int = 42
            _ = v93
            _ = v94
        }
    default:
        switch 272 {
        case -48:
            if true {
                switch ((v2 + (v2 + "a")) + (v2 + (v2 + v2))) {
                case "hello":
                    v95 := []int{-603}
                    _ = v95
                case "foo bar":
                    switch v2 {
                    case "foo bar":
                        v4 += 889
                    default:
                        var v96 uint8 = 0
                        _ = v96
                    }
                case "\n\t":
                    switch ("\n\t" + v2) {
                    case "foo bar":
                        v3 -= (v3 >> ((98) & 63))
                    }
                }
                for fi97 := 0; fi97 < 4; fi97++ {
                    var v98 string = "a"
                    _ = fi97
                    _ = v98
                }
            } else {
                for fi99 := 0; fi99 < 3; fi99++ {
                    v100 := []int{fi99}
                    _ = fi99
                    _ = v100
                }
            }
        default:
            for fi101 := 0; fi101 < 15; fi101++ {
                v2 = (("hello" + "") + v2)
                _ = fi101
            }
        }
    }
    v102 := []int{-31, -656, int(v3)}
    var v103 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
