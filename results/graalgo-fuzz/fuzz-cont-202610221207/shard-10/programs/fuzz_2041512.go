package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
type GenS6 struct { F0 int64; F1 int32; F2 string }
type GenS7 struct { F0 uint8; F1 string; F2 int64 }
type GenS8 struct { F0 int8; F1 uint32; F2 bool }
type GenS9 struct { F0 int64; F1 string; F2 uint16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 uint8) uint64 {
    p0 -= 42
    if true {
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := GenS1{F0: 354, F1: 32767, F2: 4294967295}
            var v6 int32 = 42
            _ = fi4
            _ = v5
            _ = v6
        }
        fmt.Println(p0)
    } else {
        switch "hello" {
        case "foo bar":
            var v7 uint32 = 479
            _ = v7
        case "hello":
            p0 = p0
        case "a":
            var v8 uint64 = 4294967295
            _ = v8
        default:
            v9 := []int{42, -2147483648}
            _ = v9
        }
        var v10 int64 = 461
        if (false && (1 >= -17)) {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint32 = 2147483647
                _ = fi11
                _ = v12
            }
            fmt.Println(p0, v10)
        } else {
            var v13 uint = 255
            _ = v13
        }
        _ = v10
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := []int{-708, 127, 32768}
    v14 := GenS1{F0: 2, F1: 1000, F2: 940}
    switch (("foo bar" + "") + "\n\t") {
    case "a":
        v15 := GenS2{F0: 0, F1: -915, F2: 1}
        _ = v15
    default:
        if false {
            var v16 int8 = 41
            _ = v16
        }
    }
    v17 := GenS3{F0: 0}
    switch "" {
    case "":
        var v18 uint32 = 4294967295
        _ = v18
    case "hello":
        if false {
            if (false && (!false)) {
                v19 := GenS4{F0: 4294967295, F1: 255}
                _ = v19
            } else {
                var v20 int64 = 42
                v20 -= (v20 ^ v20)
                _ = v20
            }
            v21 := GenS5{F0: ((37 > 2147483647) || false), F1: 4294967295, F2: 1}
            _ = v21
        } else {
            if true {
                var v22 uint = 2
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch ("a" + "foo bar") {
                    case "a":
                        switch "" {
                        case "":
                            var v24 uint32 = 156
                            _ = v24
                        case "\n\t":
                            for fi25 := 0; fi25 < 9; fi25++ {
                                v22 = (v22 / ((v22) | 1))
                                _ = fi25
                            }
                        case "hello":
                            if (!true) {
                                for fi26 := 0; fi26 < 10; fi26++ {
                                    var v27 bool = true
                                    v28 := GenS1{F0: 559, F1: 169, F2: 42}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        default:
                            v29 := GenS6{F0: 1, F1: 42, F2: ("" + "a")}
                            _ = v29
                        }
                    case "\n\t":
                        v30 := GenS7{F0: 0, F1: ("hello" + ((("a" + "hello") + "") + ("a" + ""))), F2: 1}
                        _ = v30
                    default:
                        var v31 int64 = 1
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v33 := v1[("" + "hello")]
                        v22 = v22
                        _ = fi32
                        _ = v33
                    }
                    if (!false) {
                        if (v22 > (v22 ^ v22)) {
                            if ((((-2.25 - ((1.5 / 1e10) * 0.0)) - 3.14159) != 1.5) || (0 >= (v22 % ((v22) | 1)))) {
                                if (0 <= 42) {
                                    if false {
                                        fmt.Println(fi23, fi23)
                                        v34 := GenS8{F0: -30, F1: 1, F2: true}
                                        var v35 int64 = -499
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 uint32 = 2
                                        v37 := GenS2{F0: v36, F1: 42, F2: (^(v36))}
                                        delete(v2, ("hello" + "foo bar"))
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 float32 = (1.5 + 1.5)
                                    _ = v38
                                } else {
                                    _ = len(v1)
                                    var v39 uint = (v22 >> ((1) & 63))
                                    var v40 float32 = 1.5
                                    _ = v39
                                    _ = v40
                                }
                                if (-2147483648 >= 1) {
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        fmt.Println(fi41, v22)
                                        _ = len(v1)
                                        v22 -= v22
                                        v22 -= (^(v22))
                                        fmt.Println(fi23, v22)
                                        _ = fi41
                                    }
                                    if (true || false) {
                                        var v42 uint64 = 1
                                        _ = v42
                                    }
                                    v43 := GenS1{F0: 2, F1: 42, F2: 293}
                                    _ = v43
                                }
                            } else {
                                switch -343 {
                                case -1:
                                    if true {
                                        v44 := GenS2{F0: 223, F1: int64(v22), F2: 4294967295}
                                        fmt.Println(v22, fi23)
                                        v45 := v1[""]
                                        _ = v44
                                        _ = v45
                                    }
                                default:
                                    fmt.Println(fi23)
                                }
                                if (((42 == 42) || true) || false) {
                                    for fi46 := 0; fi46 < 2; fi46++ {
                                        if true {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v47 := GenS4{F0: 946, F1: 1}
                                            var v48 bool = (-128 >= -38)
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = fi46
                                    }
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        var v49 float64 = 1.5
                                        v49 -= float64(-2.25)
                                        for fi50 := 0; fi50 < 1; fi50++ {
                                            var v51 uint64 = 1
                                            var v52 int64 = 852
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v49
                                    }
                                    v53 := GenS9{F0: 0, F1: ("a" + ("\n\t" + "a")), F2: 657}
                                    _ = v53
                                } else {
                                    for fi54 := 0; fi54 < 8; fi54++ {
                                        v22 = 777
                                        v22 -= v22
                                        var v55 bool = false
                                        _ = fi54
                                        _ = v55
                                    }
                                    v3 = append(v3, (-(fi23)))
                                    if false {
                                        var v56 int64 = -88
                                        var v57 int = fi23
                                        _ = v56
                                        _ = v57
                                    } else {
                                        if false {
                                            v22 += (v22 % (((v22 + v22)) | 1))
                                            v1["foo bar"] = (fi23 * fi23)
                                        } else {
                                            fmt.Println(v22, fi23)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = fi23
                                        }
                                        for fi58 := 0; fi58 < 3; fi58++ {
                                            fmt.Println(fi23, v22)
                                            v22 = (v22 & v22)
                                            _ = fi58
                                        }
                                    }
                                }
                                v59 := GenS10{F0: 441}
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 10; fi60++ {
                                _ = len(v3)
                                var v61 int32 = 2147483647
                                _ = fi60
                                _ = v61
                            }
                        }
                        delete(v2, "\n\t")
                        var v62 uint = 18446744073709551615
                        _ = v62
                    }
                    _ = fi23
                }
                _ = v22
            }
            var v63 uint32 = 15
            if (true && true) {
                v63 -= v63
            } else {
                v64 := GenS8{F0: 75, F1: v63, F2: true}
                var v65 int32 = 1
                var v66 uint8 = 255
                _ = v64
                _ = v65
                _ = v66
            }
            _ = v63
        }
    case "foo bar":
        var v67 float64 = 1.5
        _ = v67
    default:
        _ = len(v3)
    }
    switch "" {
    case "foo bar":
        for fi68 := 0; fi68 < 12; fi68++ {
            for fi69 := 0; fi69 < 4; fi69++ {
                var v70 uint8 = 255
                v70 -= v70
                var v71 uint16 = 131
                var v72 uint16 = (v71 - v71)
                if false {
                    var v73 int16 = 931
                    v74 := GenS8{F0: -8, F1: 65535, F2: true}
                    fmt.Println(v71)
                    v73 = -981
                    _ = v73
                    _ = v74
                }
                _ = fi69
                _ = v70
                _ = v71
                _ = v72
            }
            _ = len(v1)
            _ = fi68
        }
    case "\n\t":
        var v75 int16 = 823
        _ = v75
    default:
        var v76 float64 = 0.0
        _ = v76
    }
    v3 = append(v3, 9223372036854775807)
    var v77 int64 = -257
    if true {
        var v78 uint8 = 42
        _ = v78
    } else {
        var v79 uint8 = 196
        if false {
            v79 = (v79 - (v79 * (v79 << ((65535) & 63))))
        } else {
            switch 32768 {
            case 977:
                fmt.Println(v77, v77)
            case 127:
                v80 := GenS2{F0: 4294967295, F1: -431, F2: 0}
                _ = v80
            default:
                var v81 uint8 = (v79 * (v79 & v79))
                _ = v81
            }
            if len(v3) > 0 {
                v3[0] = 32768
            }
            v77 = (v77 * (^(v77)))
        }
        fmt.Println(v77)
        _ = v79
    }
    v77 += -1
    v77 += 0
    for fi82 := 0; fi82 < 16; fi82++ {
        var v83 string = "a"
        var v84 uint8 = 73
        _ = fi82
        _ = v83
        _ = v84
    }
    v85 := GenS9{F0: 9223372036854775807, F1: "", F2: 411}
    var v86 bool = (127 == -1)
    var v87 float32 = (1e10 * (1e10 - 1e10))
    switch -254 {
    case 128:
        var v88 int32 = -2147483648
        _ = v88
    case 32768:
        if len(v3) > 0 {
            v3[0] = -531
        }
    case 702:
        v89 := GenS11{F0: v86, F1: 65535, F2: 218}
        _ = v89
    }
    v90 := GenS6{F0: 574, F1: -2147483648, F2: (("foo bar" + "a") + "hello")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v77:", v77)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v77
    _ = v85
    _ = v86
    _ = v87
    _ = v90
}
