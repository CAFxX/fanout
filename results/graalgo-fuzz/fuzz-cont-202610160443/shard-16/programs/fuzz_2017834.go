package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int; F1 int; F2 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int, p1 string, p2 bool) bool {
    p1 = "a"
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 bool = p2
        var v9 uint32 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 736
    v11 := []int{(p0 - p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return ((p0 & (p0 + p0)) < (^(p0)))
}
func genFn1(p0 int) int {
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(fi12)
        v13 := GenS1{F0: 573, F1: ("hello" + "foo bar")}
        _ = fi12
        _ = v13
    }
    p0 = p0
    fmt.Println(p0)
    switch p0 {
    case 128:
        var v14 uint = 255
        _ = v14
    }
    _ = p0
    return p0
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]string)
    v3 := []uint32{}
    var v4 uint32 = 991
    var v5 uint64 = 2
    var v6 int64 = 0
    v15 := GenS1{F0: (v5 << ((18446744073709551615) & 63)), F1: "foo bar"}
    for fi16 := 0; fi16 < 16; fi16++ {
        v6 = (v6 * (^(v6)))
        for fi17 := 0; fi17 < 15; fi17++ {
            if (1 <= (v5 << ((2) & 63))) {
                var v18 int64 = 1
                v6 = 0
                var v19 int8 = -85
                _ = v18
                _ = v19
            } else {
                var v20 float64 = 1.5
                _ = v20
            }
            if true {
                v21 := GenS2{F0: 427}
                if true {
                    v22 := v2["\n\t"]
                    switch fi16 {
                    case 852:
                        var v23 int64 = 305
                        _ = v23
                    default:
                        var v24 string = v22
                        _ = v24
                    }
                    _ = v22
                } else {
                    v6 += -9223372036854775808
                    v6 = 0
                    var v25 uint64 = v5
                    _ = v25
                }
                v5 = (v5 % (((v5 | v5)) | 1))
                _ = v21
            } else {
                if true {
                    v26 := GenS1{F0: (v5 % ((0) | 1)), F1: "hello"}
                    v27 := GenS2{F0: 1}
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v4)
                }
                v1 += (v1 % ((1000) | 1))
                var v28 float32 = 0.1
                _ = v28
            }
            _ = fi17
        }
        var v29 uint64 = (^(v5))
        _ = fi16
        _ = v29
    }
    if (false || false) {
        var v30 uint32 = v4
        switch "" {
        case "":
            v31 := v2["a"]
            _ = v31
        case "hello":
            var v32 float32 = 0.0
            _ = v32
        case "a":
            fmt.Println(v1, v30, v1)
        default:
            v33 := GenS3{F0: 42, F1: 732, F2: 637}
            _ = v33
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 uint16 = 2
            v36 := GenS4{F0: 0, F1: v4, F2: 93}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v30
    }
    v37 := GenS4{F0: -95, F1: v4, F2: -128}
    var v38 float64 = 3.14159
    switch "" {
    case "":
        for fi39 := 0; fi39 < 15; fi39++ {
            if (1 != -110) {
                var v40 int8 = 29
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 int64 = (^(v6))
                    _ = fi41
                    _ = v42
                }
                _ = v40
            } else {
                v1 = 43
                delete(v2, "")
                for fi43 := 0; fi43 < 2; fi43++ {
                    var v44 float64 = ((v38 * (v38 * (v38 + v38))) - ((v38 / (float64(0.1) - (1.5 * v38))) + v38))
                    if (true || (((687 != (v6 - (v6 % (((v6 >> ((1) & 63))) | 1)))) && false) && false)) {
                        var v45 bool = true
                        delete(v2, ("a" + "\n\t"))
                        if false {
                            var v46 float64 = v44
                            switch fi43 {
                            case 32767:
                                var v47 float64 = 0.1
                                _ = v47
                            case -1:
                                v48 := GenS4{F0: 1, F1: 2, F2: 0}
                                _ = v48
                            }
                            _ = v46
                        }
                        _ = v45
                    }
                    _ = fi43
                    _ = v44
                }
            }
            _ = fi39
        }
    case "foo bar":
        delete(v2, "foo bar")
    }
    v49 := GenS5{F0: 65535}
    v4 = v4
    if (!false) {
        var v50 uint8 = 0
        _ = v50
    }
    var v51 int64 = v6
    v2["\n\t"] = "hello"
    v52 := GenS6{F0: 4294967295, F1: -2147483648, F2: 817}
    if (116 <= 1) {
        v38 = (v38 + v38)
    } else {
        if true {
            if true {
                v53 := GenS2{F0: 806}
                _ = v53
            } else {
                if (true && true) {
                    v2[("hello" + "")] = "a"
                } else {
                    for fi54 := 0; fi54 < 20; fi54++ {
                        var v55 bool = true
                        v56 := GenS3{F0: 42, F1: 2, F2: 0}
                        v57 := GenS3{F0: 42, F1: 957, F2: 766}
                        _ = fi54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                }
                fmt.Println(v4, v5)
            }
            if (!(-9223372036854775808 == 42)) {
                if true {
                    v58 := GenS4{F0: 35, F1: 1, F2: 0}
                    var v59 string = "a"
                    _ = v58
                    _ = v59
                }
                var v60 uint = 1
                v61 := GenS3{F0: -2147483648, F1: 42, F2: 32767}
                _ = v60
                _ = v61
            } else {
                var v62 int64 = v51
                v63 := GenS3{F0: -2147483648, F1: 1, F2: 1}
                _ = v62
                _ = v63
            }
            if (!true) {
                _ = len(v3)
                v64 := GenS7{F0: 2}
                var v65 int32 = -1
                _ = v64
                _ = v65
            } else {
                v4 -= 0
                var v66 int64 = v6
                for fi67 := 0; fi67 < 13; fi67++ {
                    var v68 uint = 317
                    for fi69 := 0; fi69 < 3; fi69++ {
                        var v70 int32 = int32(v68)
                        fmt.Println(v1, v70)
                        _ = fi69
                        _ = v70
                    }
                    var v71 bool = false
                    _ = fi67
                    _ = v68
                    _ = v71
                }
                _ = v66
            }
        } else {
            var v72 float32 = 0.1
            _ = v72
        }
    }
    if true {
        _ = len(v3)
    } else {
        var v73 int8 = -97
        _ = v73
    }
    fmt.Println(v5, v38)
    if (((false && true) || true) || (("hello" < "a") && false)) {
        if false {
            var v74 uint = 65535
            _ = len(v3)
            _ = v74
        } else {
            v4 = (v4 - (v4 ^ 2147483647))
            fmt.Println(v51, v5, v38)
            if ((v1 & v1) >= v1) {
                if true {
                    switch 4294967295 {
                    case 127:
                        v5 = v5
                    default:
                        v75 := GenS8{F0: 42, F1: v51}
                        _ = v75
                    }
                } else {
                    for fi76 := 0; fi76 < 17; fi76++ {
                        var v77 uint8 = 2
                        var v78 uint8 = v77
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                }
                for fi79 := 0; fi79 < 9; fi79++ {
                    if false {
                        for fi80 := 0; fi80 < 1; fi80++ {
                            v5 += 485
                            _ = fi80
                        }
                        var v81 uint8 = 2
                        if (false || (!((0 > 728) || true))) {
                            for fi82 := 0; fi82 < 12; fi82++ {
                                v38 = ((v38 / 1.0) - (1.5 / 0.1))
                                v6 += v6
                                var v83 uint8 = (^(v81))
                                delete(v2, "foo bar")
                                v84 := GenS9{F0: 534, F1: (fi79 + (^(fi79))), F2: 2}
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                            v5 = 333
                        } else {
                            v2["foo bar"] = "a"
                            for fi85 := 0; fi85 < 18; fi85++ {
                                if (((("foo bar" + "") + (((("hello" + "") + ("foo bar" + "a")) + "a") + "")) + ("foo bar" + "\n\t")) >= "a") {
                                    if (!true) {
                                        var v86 uint64 = (v5 >> ((4294967295) & 63))
                                        _ = v86
                                    }
                                }
                                _ = fi85
                            }
                        }
                        _ = v81
                    } else {
                        v87 := GenS10{F0: v5}
                        var v88 bool = true
                        _ = v87
                        _ = v88
                    }
                    v51 = v6
                    v1 = v1
                    _ = fi79
                }
                for fi89 := 0; fi89 < 4; fi89++ {
                    var v90 bool = (v1 > v1)
                    _ = fi89
                    _ = v90
                }
            } else {
                v51 += -9223372036854775808
                v91 := GenS11{F0: -400, F1: 0}
                _ = v91
            }
        }
        var v92 int16 = (v1 - (v1 & -32768))
        if true {
            for fi93 := 0; fi93 < 13; fi93++ {
                v38 = v38
                v3 = append(v3, (v4 << ((818) & 63)))
                v38 += v38
                _ = fi93
            }
            if (!(!(!(v5 >= v5)))) {
                var v94 float32 = (float32((1.5 * 3.14159)) / float32(v38))
                for fi95 := 0; fi95 < 1; fi95++ {
                    if ((^(v1)) < (v1 - (v1 << ((725) & 63)))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        for fi96 := 0; fi96 < 9; fi96++ {
                            var v97 uint32 = (v4 & (v4 + 0))
                            v98 := GenS8{F0: 127, F1: (v51 - (v51 / ((v6) | 1)))}
                            fmt.Println(v4, v92)
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                    }
                    _ = fi95
                }
                _ = v94
            } else {
                v99 := GenS1{F0: v5, F1: "foo bar"}
                _ = v99
            }
            var v100 int16 = 32767
            _ = v100
        }
        _ = v92
    } else {
        switch -2147483648 {
        case 649:
            v101 := v2["hello"]
            _ = v101
        }
        if false {
            v102 := GenS12{F0: 190}
            _ = v102
        } else {
            if (v5 != v5) {
                v103 := GenS6{F0: 9223372036854775807, F1: 32767, F2: 997}
                _ = v103
            } else {
                if false {
                    var v104 uint16 = 1
                    _ = v104
                }
                v51 = v51
            }
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
