package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int; F1 int64; F2 int64 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 string; F1 int; F2 uint64 }
type GenS12 struct { F0 uint }
func genFn0(p0 int32, p1 int8, p2 string) uint32 {
    var v4 bool = true
    var v5 uint16 = 0
    p2 = p2
    v6 := GenS1{F0: ((p2 + p2) + "hello"), F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 bool, p1 uint16, p2 bool) uint64 {
    p0 = (!(-2.25 >= 0.0))
    var v7 float64 = ((0.0 * 0.1) * (float64((-2.25 - 3.14159)) * 1e10))
    v8 := []int{32767, 4294967295, 0}
    v9 := GenS2{F0: 42, F1: 2, F2: 26}
    var v10 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]int)
    v3 := []uint32{11, 5, 986}
    var v11 int16 = 0
    switch 127 {
    case 36:
        v1 = v1
    case 4294967295:
        if (0 == 9223372036854775807) {
            if false {
                v1 = 295
            } else {
                var v12 bool = ("\n\t" < "\n\t")
                v12 = false
                v1 = v1
                _ = v12
            }
            switch "\n\t" {
            case "\n\t":
                var v13 float64 = float64((1e10 - float32(((-2.25 - (-2.25 / -2.25)) - (0.1 - 1e10)))))
                _ = v13
            case "hello":
                v11 += v11
            }
        } else {
            if (!true) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        var v16 int = fi14
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                v11 -= -506
                for fi17 := 0; fi17 < 9; fi17++ {
                    if (-2.25 != -2.25) {
                        v18 := GenS1{F0: "hello", F1: 1}
                        _ = len(v3)
                        _ = v18
                    } else {
                        fmt.Println(v1, fi17, v1)
                        switch (fi17 + 942) {
                        case 32768:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    var v19 uint32 = 42
                    if (1.5 > 3.14159) {
                        var v20 uint16 = 42
                        if (v19 >= v19) {
                            var v21 float32 = 3.14159
                            for fi22 := 0; fi22 < 17; fi22++ {
                                if (((fi22 * (-(fi22))) >= (fi17 ^ 4294967295)) && true) {
                                    v19 += (v19 << ((42) & 63))
                                    var v23 uint32 = (^(v19))
                                    _ = v23
                                } else {
                                    if false {
                                        var v24 int32 = (v1 | v1)
                                        _ = v24
                                    } else {
                                        v19 -= 1
                                        if true {
                                            switch ((("a" + "\n\t") + ("foo bar" + ("a" + "foo bar"))) + (((("" + ((("foo bar" + "foo bar") + "hello") + ("hello" + ""))) + "a") + ("hello" + "a")) + ("" + "hello"))) {
                                            case "hello":
                                                var v25 int = (fi17 ^ 899)
                                                _ = v25
                                            case "":
                                                v26 := GenS2{F0: -9223372036854775808, F1: 18446744073709551615, F2: 1}
                                                _ = v26
                                            case "a":
                                                switch "hello" {
                                                case "a":
                                                    var v27 uint = 0
                                                    _ = v27
                                                case "foo bar":
                                                    var v28 uint = 996
                                                    _ = v28
                                                default:
                                                    fmt.Println(v11, v19)
                                                }
                                            }
                                            v29 := GenS3{F0: 0, F1: (^(v19))}
                                            _ = v29
                                        }
                                    }
                                    v19 += v19
                                }
                                fmt.Println(fi22)
                                delete(v2, (fi22 % (((fi17 - 0)) | 1)))
                                _ = fi22
                            }
                            _ = v21
                        } else {
                            var v30 int = fi17
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v20)
                            _ = v30
                        }
                        var v31 int = fi17
                        _ = v20
                        _ = v31
                    } else {
                        if ((3.14159 != 3.14159) || true) {
                            for fi32 := 0; fi32 < 5; fi32++ {
                                switch (fi17 ^ fi17) {
                                case 2147483647:
                                    var v33 uint64 = 4294967295
                                    _ = v33
                                case -942:
                                    fmt.Println(v19, v11)
                                case -764:
                                    v1 = -269
                                }
                                _ = fi32
                            }
                        }
                        v34 := GenS3{F0: 384, F1: v19}
                        _ = v34
                    }
                    _ = fi17
                    _ = v19
                }
            } else {
                var v35 uint16 = 196
                v11 -= 339
                _ = v35
            }
            fmt.Println(v11)
        }
    case -248:
        for fi36 := 0; fi36 < 2; fi36++ {
            for fi37 := 0; fi37 < 16; fi37++ {
                delete(v2, 0)
                switch "hello" {
                case "":
                    var v38 float32 = 3.14159
                    _ = v38
                case "foo bar":
                    if true {
                        v39 := v2[(fi37 - (fi36 | (fi37 >> ((255) & 63))))]
                        _ = v39
                    } else {
                        v1 = v1
                        v40 := GenS2{F0: 1, F1: 4294967295, F2: 42}
                        var v41 string = "foo bar"
                        _ = v40
                        _ = v41
                    }
                default:
                    fmt.Println(fi36)
                }
                v11 = (v11 / ((v11) | 1))
                _ = fi37
            }
            v1 = (v1 % (((v1 + (v1 & -2147483648))) | 1))
            v42 := v2[fi36]
            _ = fi36
            _ = v42
        }
    default:
        v11 = v11
    }
    v2[1] = 9223372036854775807
    v43 := GenS1{F0: "hello", F1: 2}
    if (true || false) {
        if false {
            if (false || false) {
                for fi44 := 0; fi44 < 4; fi44++ {
                    if true {
                        v1 = 42
                        for fi45 := 0; fi45 < 17; fi45++ {
                            if true {
                                v11 = (-(v11))
                            }
                            v46 := GenS4{F0: v1, F1: 1}
                            _ = fi45
                            _ = v46
                        }
                        v11 += v11
                    } else {
                        var v47 int8 = 42
                        for fi48 := 0; fi48 < 3; fi48++ {
                            if (1e10 > 0.0) {
                                var v49 float64 = 1.5
                                var v50 string = "\n\t"
                                _ = v49
                                _ = v50
                            } else {
                                v51 := GenS3{F0: genFn1(true, 1, (false && (("a" + (("hello" + (("a" + "\n\t") + (("foo bar" + "") + "hello"))) + ("hello" + "foo bar"))) == ""))), F1: 2}
                                var v52 int16 = v11
                                v53 := GenS5{F0: 18446744073709551615, F1: 2147483647}
                                if false {
                                    if false {
                                        if false {
                                            var v54 uint16 = 398
                                            fmt.Println(v52, v11, v52)
                                            v54 = 1
                                            _ = v54
                                        }
                                    } else {
                                        v1 = int32(fi44)
                                        var v55 uint32 = 538
                                        _ = v55
                                    }
                                }
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            switch "foo bar" {
                            case "a":
                                var v56 int64 = 695
                                _ = v56
                            case "foo bar":
                                v11 = (v11 | v11)
                            case "hello":
                                var v57 uint8 = 2
                                _ = v57
                            }
                            v47 -= 100
                            _ = fi48
                        }
                        var v58 int = 32768
                        _ = v47
                        _ = v58
                    }
                    if (true && (42 < 65535)) {
                        v3 = append(v3, 1)
                    } else {
                        fmt.Println(v1, v11, v11)
                        v1 -= (v1 & v1)
                    }
                    _ = fi44
                }
            }
            switch "foo bar" {
            case "":
                fmt.Println(v11, v11)
            case "foo bar":
                var v59 int16 = v11
                _ = v59
            case "\n\t":
                var v60 int16 = v11
                _ = v60
            default:
                if false {
                    v1 = v1
                    var v61 bool = false
                    _ = v61
                } else {
                    v62 := GenS6{F0: 9223372036854775807, F1: 779, F2: -1}
                    v63 := GenS3{F0: genFn1(true, 1, true), F1: 0}
                    var v64 int16 = v11
                    _ = v62
                    _ = v63
                    _ = v64
                }
            }
        }
    }
    v65 := GenS7{F0: -118, F1: 42}
    for fi66 := 0; fi66 < 13; fi66++ {
        v67 := GenS7{F0: 100, F1: 1}
        switch 9223372036854775807 {
        case 128:
            v1 = 638
        case 2147483647:
            var v68 uint64 = genFn1(false, 32767, true)
            _ = v68
        }
        if false {
            for fi69 := 0; fi69 < 16; fi69++ {
                v1 = (v1 * (v1 + v1))
                _ = fi69
            }
            if true {
                _ = len(v2)
                v11 = (-(v11))
            } else {
                delete(v2, fi66)
                v11 -= (v11 % ((v11) | 1))
            }
        } else {
            if true {
                delete(v2, (fi66 & fi66))
                switch fi66 {
                case -2147483648:
                    fmt.Println(v11, fi66, v11)
                case -342:
                    for fi70 := 0; fi70 < 15; fi70++ {
                        _ = len(v3)
                        _ = fi70
                    }
                default:
                    v71 := GenS7{F0: -3, F1: 1}
                    _ = v71
                }
                v72 := v2[(fi66 - (fi66 | fi66))]
                _ = v72
            }
        }
        _ = fi66
        _ = v67
    }
    v2[32768] = -9223372036854775808
    v11 = 1
    var v73 int32 = 42
    v73 = (v73 ^ 42)
    var v74 uint32 = 2
    v75 := GenS8{F0: 0}
    v11 += (v11 & (-(v11)))
    var v76 int32 = (-(v1))
    if (("a" + "\n\t") > "a") {
        var v77 bool = true
        for fi78 := 0; fi78 < 12; fi78++ {
            fmt.Println(fi78)
            _ = fi78
        }
        _ = v77
    } else {
        fmt.Println(v1, v76, v76)
    }
    if ((false && false) || true) {
        switch 0 {
        case 9223372036854775807:
            if false {
                v79 := GenS8{F0: 25}
                _ = v79
            }
        case 0:
            if (false || (0 < 616)) {
                var v80 uint8 = 0
                v81 := GenS9{F0: true, F1: ("" + "\n\t")}
                for fi82 := 0; fi82 < 6; fi82++ {
                    if false {
                        var v83 uint64 = 2
                        _ = v83
                    }
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v74, fi82)
                    case "a":
                        v11 = v11
                    case "foo bar":
                        if false {
                            var v84 int64 = 9223372036854775807
                            _ = v84
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi82
                }
                _ = v80
                _ = v81
            } else {
                delete(v2, 371)
                v85 := GenS10{F0: 1, F1: true}
                if (false && (-2147483648 == v1)) {
                    if false {
                        var v86 int64 = -1
                        v87 := GenS11{F0: "\n\t", F1: 128, F2: 2}
                        _ = v86
                        _ = v87
                    } else {
                        for fi88 := 0; fi88 < 13; fi88++ {
                            v74 += 640
                            _ = fi88
                        }
                        for fi89 := 0; fi89 < 1; fi89++ {
                            var v90 uint16 = uint16(v74)
                            for fi91 := 0; fi91 < 5; fi91++ {
                                fmt.Println(v11, v11, v76)
                                _ = fi91
                            }
                            _ = fi89
                            _ = v90
                        }
                    }
                }
                _ = v85
            }
        case -33:
            switch "hello" {
            case "\n\t":
                v92 := GenS12{F0: 255}
                _ = v92
            case "a":
                for fi93 := 0; fi93 < 17; fi93++ {
                    if false {
                        switch "" {
                        case "foo bar":
                            for fi94 := 0; fi94 < 11; fi94++ {
                                switch fi94 {
                                case 815:
                                    var v95 int = fi94
                                    _ = v95
                                case -1:
                                    var v96 float32 = 3.14159
                                    _ = v96
                                case 4294967295:
                                    var v97 uint = 0
                                    _ = v97
                                default:
                                    var v98 uint = 2
                                    var v99 uint8 = 3
                                    _ = v98
                                    _ = v99
                                }
                                _ = fi94
                            }
                        }
                        fmt.Println(v1)
                        fmt.Println(v76, fi93)
                    }
                    var v100 uint16 = 210
                    _ = fi93
                    _ = v100
                }
            case "foo bar":
                fmt.Println(v76, v76, v76)
            }
        }
        if len(v3) > 0 {
            v3[0] = genFn0((v73 % ((0) | 1)), -128, "")
        }
        var v101 uint32 = genFn0(v73, 127, "a")
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v43
    _ = v65
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
