package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string; F1 uint16; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint }
func genFn0(p0 int32, p1 uint32) bool {
    p1 = p1
    if true {
        var v7 int32 = (^(p0))
        _ = v7
    }
    var v8 int8 = 100
    if false {
        if (255 <= 696) {
            v9 := GenS1{F0: -1, F1: ("" + "foo bar")}
            _ = v9
        } else {
            p1 = (p1 | (p1 << ((65535) & 63)))
            v10 := GenS1{F0: 0, F1: "hello"}
            _ = v10
        }
        if true {
            if (808 >= 42) {
                fmt.Println(p1, v8, p1)
            } else {
                var v11 uint = 4294967295
                if true {
                    var v12 int64 = -9223372036854775808
                    _ = v12
                }
                var v13 int = 32767
                _ = v11
                _ = v13
            }
            if true {
                v8 -= (v8 % ((-80) | 1))
                var v14 int = 2147483647
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 float64 = 1.5
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        fmt.Println(p1)
                        v19 := GenS1{F0: 1000, F1: (("\n\t" + "foo bar") + "a")}
                        if ((!(1 > 9223372036854775807)) && false) {
                            v20 := GenS1{F0: 1000, F1: ("" + ("" + "hello"))}
                            v21 := GenS1{F0: 350, F1: "hello"}
                            _ = v20
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                    switch 127 {
                    case 128:
                        v22 := GenS1{F0: 422, F1: "hello"}
                        _ = v22
                    case 375:
                        v23 := GenS2{F0: (("\n\t" + "a") + "a"), F1: 42, F2: (fi17 << ((120) & 63))}
                        _ = v23
                    }
                    fmt.Println(v8)
                }
                _ = fi17
            }
        }
        v8 += 1
    } else {
        var v24 int64 = -782
        v25 := []int{-614, -9223372036854775808, -9223372036854775808}
        var v26 bool = false
        _ = v24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v8
    return false
}
func genFn1(p0 uint16, p1 int16, p2 int) int16 {
    v27 := GenS2{F0: "hello", F1: (p0 / (((p0 - p0)) | 1)), F2: p2}
    var v28 uint = 363
    var v29 float32 = (0.0 - 0.0)
    var v30 uint = (v28 >> ((v28) & 63))
    var v31 int = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    _ = v31
    return 32767
}

func main() {
    var v1 bool = false
    var v2 uint32 = 370
    v3 := []string{"foo bar", ("foo bar" + "")}
    var v4 uint64 = 18446744073709551615
    v5 := make(map[string]int)
    v6 := make(map[string]bool)
    var v32 bool = (((!v1) && v1) && (1.5 == 1e10))
    var v33 float32 = (3.14159 * 3.14159)
    fmt.Println(v4, v32)
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 int64 = -652
        _ = fi34
        _ = v35
    }
    _ = len(v5)
    v4 -= (v4 + v4)
    if v1 {
        fmt.Println(v33, v32, v32)
        v32 = v1
        if (4294967295 != 65535) {
            if (2 <= 2) {
                for fi36 := 0; fi36 < 9; fi36++ {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        v38 := GenS1{F0: 32767, F1: ""}
                        _ = fi37
                        _ = v38
                    }
                    v39 := GenS1{F0: 1, F1: "hello"}
                    var v40 uint16 = 42
                    _ = fi36
                    _ = v39
                    _ = v40
                }
                v1 = ((-105 <= -128) && true)
                v41 := GenS3{F0: (v2 & (v2 ^ (v2 | (^(v2)))))}
                _ = v41
            } else {
                switch ("hello" + "") {
                case "hello":
                    var v42 float32 = v33
                    _ = v42
                case "foo bar":
                    fmt.Println(v4)
                }
                var v43 int64 = 474
                var v44 int16 = 698
                _ = v43
                _ = v44
            }
            for fi45 := 0; fi45 < 13; fi45++ {
                switch fi45 {
                case 4294967295:
                    v32 = (88 <= 0)
                case 128:
                    fmt.Println(fi45, v2, v33)
                case -1:
                    if len(v3) > 0 {
                        v3[0] = ("" + "")
                    }
                }
                v2 -= (v2 >> ((376) & 63))
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                _ = fi45
            }
            var v46 string = "hello"
            _ = v46
        } else {
            if false {
                v32 = (v32 || (0 < 2))
                var v47 uint64 = (v4 ^ v4)
                _ = v47
            }
            v48 := v5["\n\t"]
            var v49 uint64 = 4294967295
            _ = v48
            _ = v49
        }
    }
    var v50 int64 = 363
    v32 = ((v2 <= v2) && v32)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            v51 := GenS3{F0: (^(v2))}
            _ = v51
        case "":
            v33 -= v33
        case "a":
            v1 = v32
        }
        v4 = v4
    } else {
        fmt.Println(v4, v33, v4)
    }
    if true {
        if (("\n\t" + "\n\t") < ((("hello" + ("\n\t" + ("" + "hello"))) + "\n\t") + ("hello" + ("" + "\n\t")))) {
            v3 = append(v3, "")
        } else {
            if false {
                v2 = v2
            } else {
                if (1e10 < 0.0) {
                    fmt.Println(v2, v4, v32)
                    fmt.Println(v1)
                }
                fmt.Println(v1)
                switch -379 {
                case 1:
                    var v52 uint16 = 255
                    _ = v52
                case -2147483648:
                    delete(v6, ("hello" + ""))
                default:
                    for fi53 := 0; fi53 < 9; fi53++ {
                        var v54 uint = 57
                        v55 := GenS1{F0: int16(v4), F1: ("foo bar" + "")}
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                }
            }
            switch -2147483648 {
            case -504:
                for fi56 := 0; fi56 < 2; fi56++ {
                    v57 := GenS3{F0: (^(v2))}
                    for fi58 := 0; fi58 < 19; fi58++ {
                        if (v32 || (float64((v33 * 1.5)) == -2.25)) {
                            v1 = ((((("a" + "") + (("a" + "hello") + "\n\t")) + "foo bar") + "foo bar") < "a")
                        } else {
                            if ((uint(v2) != 255) && v1) {
                                var v59 int64 = v50
                                _ = len(v5)
                                _ = v59
                            }
                            switch "" {
                            case "\n\t":
                                if v1 {
                                    var v60 uint8 = 0
                                    if v1 {
                                        v61 := v5[""]
                                        fmt.Println(v61, fi56, fi56)
                                        _ = v61
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    _ = len(v5)
                                    v5["foo bar"] = fi56
                                    _ = v60
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if v32 {
                                    v4 -= (^(v4))
                                    var v62 int32 = -2147483648
                                    _ = v62
                                } else {
                                    if v1 {
                                        var v63 int64 = 852
                                        _ = v63
                                    }
                                }
                            default:
                                v64 := GenS2{F0: (("a" + (("foo bar" + "foo bar") + "hello")) + "hello"), F1: 1, F2: (fi56 / ((fi58) | 1))}
                                _ = v64
                            }
                            switch fi58 {
                            case 395:
                                fmt.Println(v32)
                            case 1:
                                var v65 int32 = -516
                                _ = v65
                            case -332:
                                v32 = (0.0 <= 0.0)
                            default:
                                v1 = v1
                            }
                        }
                        var v66 int16 = -1
                        v4 -= 822
                        _ = fi58
                        _ = v66
                    }
                    _ = fi56
                    _ = v57
                }
            default:
                v4 += (v4 | (v4 % ((42) | 1)))
            }
            var v67 int32 = 0
            _ = v67
        }
    } else {
        v1 = (v2 >= v2)
        for fi68 := 0; fi68 < 1; fi68++ {
            v32 = false
            for fi69 := 0; fi69 < 2; fi69++ {
                if (v32 && (float32((3.14159 - 3.14159)) != v33)) {
                    v6["foo bar"] = (((1 < 991) && v32) || v1)
                    switch (fi68 & (fi68 >> ((353) & 63))) {
                    case 484:
                        if v1 {
                            fmt.Println(v33)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (14 <= 813) {
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                                v70 := GenS1{F0: -225, F1: "foo bar"}
                                var v71 int = (fi68 % ((fi69) | 1))
                                var v72 uint8 = 1
                                _ = len(v3)
                                var v73 string = ("a" + ("a" + ("a" + ((("" + "") + "foo bar") + "a"))))
                                var v74 bool = (-536 == (v50 ^ (v50 * v50)))
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                                _ = v74
                            } else {
                                if v32 {
                                    if ((!v32) && false) {
                                        if ((-(fi68)) <= (fi68 << ((65535) & 63))) {
                                            if true {
                                                var v75 int16 = 1
                                                v2 = 366
                                                v5[""] = (fi69 * fi68)
                                                _ = v75
                                            }
                                            v50 -= int64(v4)
                                            if (240 != 0) {
                                                if ((("a" + "a") >= "hello") && v32) {
                                                    var v76 uint8 = 127
                                                    v77 := GenS2{F0: "", F1: 42, F2: fi69}
                                                    var v78 float32 = 0.0
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    v50 -= 486
                                                    v79 := GenS1{F0: 42, F1: "a"}
                                                    v4 = (^(v4))
                                                    _ = v79
                                                }
                                            } else {
                                                v2 -= v2
                                                var v80 uint64 = v4
                                                var v81 uint64 = v80
                                                _ = v80
                                                _ = v81
                                            }
                                        }
                                        if (!false) {
                                            switch "\n\t" {
                                            case "hello":
                                                v82 := GenS3{F0: v2}
                                                _ = v82
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                            case "foo bar":
                                                if v32 {
                                                    _ = len(v3)
                                                    var v83 uint16 = 579
                                                    _ = v83
                                                } else {
                                                    var v84 int16 = genFn1(65535, 1, 0)
                                                    v1 = (!v1)
                                                    _ = v84
                                                }
                                            }
                                        }
                                        v85 := GenS4{F0: 1}
                                        _ = v85
                                    }
                                    if false {
                                        var v86 int16 = 1000
                                        v87 := GenS5{F0: v86, F1: fi68}
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                        _ = v86
                                        _ = v87
                                    }
                                    v2 = v2
                                }
                                if v1 {
                                    var v88 float32 = (-2.25 + float32(3.14159))
                                    _ = len(v6)
                                    _ = v88
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        var v89 int8 = 1
                                        _ = v89
                                    case "a":
                                        v33 -= (3.14159 * ((v33 - v33) / v33))
                                    case "hello":
                                        v90 := GenS5{F0: 1, F1: fi69}
                                        _ = v90
                                    }
                                    switch fi69 {
                                    case 1:
                                        _ = len(v3)
                                    case 9223372036854775807:
                                        v91 := GenS4{F0: genFn1(557, 1, (fi69 % (((-(fi69))) | 1)))}
                                        _ = v91
                                    case -739:
                                        v4 = v4
                                    }
                                }
                            }
                            if (64 == -1) {
                                var v92 float32 = 1e10
                                _ = v92
                            } else {
                                v6[""] = true
                                if true {
                                    delete(v6, "")
                                    switch 128 {
                                    case 0:
                                        var v93 uint16 = 32767
                                        _ = v93
                                    }
                                }
                                v94 := GenS6{F0: 1}
                                _ = v94
                            }
                        } else {
                            v32 = v32
                            fmt.Println(v2, v2)
                        }
                    case 4294967295:
                        v95 := GenS1{F0: int16(v50), F1: "a"}
                        _ = v95
                    case 127:
                        v96 := GenS6{F0: 65535}
                        _ = v96
                    default:
                        v50 -= v50
                    }
                } else {
                    if v1 {
                        var v97 int8 = 0
                        switch (^(fi68)) {
                        case 32767:
                            var v98 float64 = -2.25
                            _ = v98
                        case 32768:
                            var v99 uint16 = 42
                            _ = v99
                        }
                        _ = v97
                    }
                    var v100 uint32 = 785
                    if true {
                        if (0 == 79) {
                            fmt.Println(v50)
                        } else {
                            var v101 bool = v1
                            if ((("" + "a") + "hello") < "foo bar") {
                                fmt.Println(v32)
                                var v102 uint64 = v4
                                v103 := v6["foo bar"]
                                _ = v102
                                _ = v103
                            } else {
                                var v104 int = (fi68 << ((18446744073709551615) & 63))
                                _ = v104
                            }
                            _ = v101
                        }
                    }
                    _ = v100
                }
                _ = fi69
            }
            fmt.Println(v32)
            _ = fi68
        }
    }
    var v105 int64 = v50
    var v106 uint32 = v2
    for fi107 := 0; fi107 < 1; fi107++ {
        fmt.Println(v105)
        for fi108 := 0; fi108 < 14; fi108++ {
            fmt.Println(fi107)
            _ = fi108
        }
        _ = fi107
    }
    v105 = (v50 ^ v105)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
