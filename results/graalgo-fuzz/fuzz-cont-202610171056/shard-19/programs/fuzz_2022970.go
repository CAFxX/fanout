package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int32; F1 uint; F2 int16 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 int16; F1 uint; F2 uint }
type GenS11 struct { F0 int16; F1 int8; F2 int64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int }
func genFn0(p0 uint32) uint32 {
    var v7 int32 = 1
    v8 := GenS1{F0: (v7 >> ((2) & 63)), F1: "hello", F2: uint(p0)}
    v9 := GenS2{F0: -659}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 855
    v2 := []string{}
    var v3 int = -691
    var v4 int = 2147483647
    var v5 uint64 = 646
    v6 := []string{}
    v1 = v1
    v10 := GenS1{F0: 42, F1: ("\n\t" + ("a" + ("hello" + ((("\n\t" + "foo bar") + "\n\t") + "")))), F2: 985}
    switch "\n\t" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "hello")
        }
    case "foo bar":
        v4 = -106
    default:
        var v11 uint8 = 24
        _ = v11
    }
    _ = len(v6)
    var v12 int8 = 60
    v13 := GenS1{F0: -100, F1: "\n\t", F2: 1}
    if ((!false) || (true || (true || (695 != genFn0(42))))) {
        var v14 float64 = float64((3.14159 * -2.25))
        fmt.Println(v5, v14)
        _ = v14
    }
    v4 += v3
    var v15 int64 = 0
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            if (4294967295 == 4294967295) {
                var v18 uint = 42
                fmt.Println(v4)
                _ = v18
            } else {
                switch "hello" {
                case "\n\t":
                    v4 = v3
                case "a":
                    _ = len(v6)
                case "":
                    var v19 int16 = -406
                    _ = v19
                default:
                    if (2147483647 <= -770) {
                        if (("foo bar" + "foo bar") >= "a") {
                            v5 = v5
                        }
                        if (false || false) {
                            fmt.Println(v5, v4, v1)
                            fmt.Println(v12)
                            v20 := GenS3{F0: 0, F1: "hello"}
                            _ = v20
                        } else {
                            switch (v4 | (-(fi16))) {
                            case -9223372036854775808:
                                v2 = append(v2, "")
                            case 321:
                                v21 := GenS1{F0: -2147483648, F1: "foo bar", F2: 4294967295}
                                _ = v21
                            default:
                                var v22 int16 = 176
                                _ = v22
                            }
                            if true {
                                v3 += (fi17 >> ((973) & 63))
                            }
                            v1 = (v1 / (((v1 & 2)) | 1))
                            var v23 uint = 2
                            _ = v23
                        }
                        v6 = append(v6, "\n\t")
                    } else {
                        if ((v15 > (^(v15))) || (-1 == 174)) {
                            var v24 uint64 = (v5 ^ (v5 >> ((707) & 63)))
                            _ = v24
                        } else {
                            if true {
                                switch ("foo bar" + ("\n\t" + "foo bar")) {
                                case "":
                                    v2 = append(v2, "")
                                    var v25 string = ("" + (("foo bar" + "\n\t") + (("" + "\n\t") + "\n\t")))
                                    if false {
                                        if true {
                                            v1 = (^(v1))
                                            v1 = v1
                                        }
                                    } else {
                                        v15 = 42
                                    }
                                    _ = v25
                                case "\n\t":
                                    if false {
                                        switch 9223372036854775807 {
                                        case 9223372036854775807:
                                            var v26 uint8 = 136
                                            _ = v26
                                        default:
                                            var v27 uint8 = 128
                                            _ = v27
                                        }
                                        if (!(!(42 <= 128))) {
                                            if len(v6) > 0 {
                                                v6[0] = ("hello" + "a")
                                            }
                                            if true {
                                                v5 -= (v5 | (v5 * 532))
                                            }
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v4, fi16, v5)
                                var v28 uint64 = (v5 ^ v5)
                                if false {
                                    fmt.Println(v4)
                                }
                                if (1 >= -1) {
                                    v5 += v28
                                    var v29 float64 = 1.5
                                    var v30 float64 = (v29 - (v29 + 3.14159))
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS4{F0: 1, F1: 0, F2: (v1 | (v1 % (((v1 << ((2) & 63))) | 1)))}
                                    fmt.Println(v28, v4, v3)
                                    if (!((!((65535 != uint32(v1)) || (1 > -1))) && (((!false) && (2 >= 294)) || (!false)))) {
                                        if len(v6) > 0 {
                                            v6[0] = "foo bar"
                                        }
                                        v12 = 1
                                        if true {
                                            if true {
                                                v32 := GenS4{F0: v1, F1: 18446744073709551615, F2: v1}
                                                var v33 string = "a"
                                                _ = v32
                                                _ = v33
                                            }
                                            v1 += v1
                                            v34 := GenS1{F0: 0, F1: "hello", F2: 255}
                                            v35 := GenS3{F0: (v12 ^ (^(v12))), F1: ""}
                                            var v36 string = "a"
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        } else {
                                            v28 -= (v5 + (v28 >> ((18446744073709551615) & 63)))
                                        }
                                        if (!((-468 < (v4 / (((v4 % ((561) | 1))) | 1))) || false)) {
                                            v12 += (v12 ^ (v12 * v12))
                                            v3 += v4
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v1, v3)
                                            case "foo bar":
                                                v37 := GenS1{F0: 42, F1: "a", F2: 665}
                                                _ = v37
                                            case "":
                                                if true {
                                                    _ = len(v6)
                                                    fmt.Println(fi16, v12, fi16)
                                                    if (0 > 690) {
                                                        v38 := GenS4{F0: 65535, F1: 2, F2: v1}
                                                        v6 = append(v6, "\n\t")
                                                        v2 = append(v2, "foo bar")
                                                        _ = v38
                                                    } else {
                                                        if false {
                                                            v39 := GenS5{F0: (v4 ^ v4), F1: 1, F2: 2147483647}
                                                            switch (("" + "\n\t") + (("" + ("hello" + "hello")) + ("hello" + ("" + ("foo bar" + (("a" + "a") + ((("" + (("a" + ("a" + "")) + ("a" + "foo bar"))) + "\n\t") + "hello"))))))) {
                                                            case "a":
                                                                if (552 <= 0) {
                                                                    v12 = 0
                                                                    if false {
                                                                        fmt.Println(v15)
                                                                        if (false || true) {
                                                                            if true {
                                                                                var v40 int64 = (-(v15))
                                                                                _ = v40
                                                                            } else {
                                                                                v41 := GenS6{F0: 255, F1: (^(v3))}
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v5 = (v5 * 395)
                                                                                }
                                                                                var v42 uint8 = 0
                                                                                _ = v41
                                                                                _ = v42
                                                                            }
                                                                        }
                                                                        v43 := GenS7{F0: v12, F1: (fi16 * fi17)}
                                                                        v2 = append(v2, "")
                                                                        var v44 string = ""
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    fmt.Println(v12, v4, v3)
                                                                    var v45 uint64 = 42
                                                                    _ = v45
                                                                } else {
                                                                    fmt.Println(v3, v4)
                                                                    if false {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    var v46 int32 = -1
                                                                    v47 := GenS8{F0: int32(v15), F1: 851, F2: 0}
                                                                    if (true && false) {
                                                                        var v48 uint32 = 61
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            var v49 int16 = 1
                                                                            _ = v49
                                                                        default:
                                                                            var v50 float32 = 1.5
                                                                            _ = v50
                                                                        }
                                                                        v28 += (v28 + v28)
                                                                        _ = v48
                                                                    } else {
                                                                        v51 := GenS9{F0: (v28 | (v5 * (v28 ^ v28))), F1: v15}
                                                                        if (!("foo bar" >= "\n\t")) {
                                                                            var v52 string = (("foo bar" + "a") + "foo bar")
                                                                            if false {
                                                                                v4 = v3
                                                                                v15 += (v15 + v15)
                                                                                var v53 uint8 = 248
                                                                                _ = v53
                                                                            }
                                                                            var v54 bool = ((v12 + (-(v12))) >= -128)
                                                                            var v55 bool = v54
                                                                            _ = v52
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        var v56 bool = false
                                                                        _ = v51
                                                                        _ = v56
                                                                    }
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            case "foo bar":
                                                                var v57 uint32 = 65535
                                                                _ = v57
                                                            case "":
                                                                switch (^(v4)) {
                                                                case 42:
                                                                    v58 := GenS10{F0: 0, F1: uint(v5), F2: 236}
                                                                    _ = v58
                                                                case 1:
                                                                    v15 = v15
                                                                case -739:
                                                                    v6 = append(v6, ("hello" + ""))
                                                                }
                                                            default:
                                                                fmt.Println(fi17)
                                                            }
                                                            _ = v39
                                                        }
                                                        var v59 float32 = 3.14159
                                                        v59 = (v59 - v59)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if (false || true) {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(v1)
                                                                case "a":
                                                                    v60 := GenS5{F0: (fi16 - (v4 >> ((0) & 63))), F1: 279, F2: -1}
                                                                    _ = v60
                                                                case "hello":
                                                                    var v61 uint16 = (v1 >> ((0) & 63))
                                                                    _ = v61
                                                                }
                                                                var v62 uint32 = 2
                                                                _ = v62
                                                            }
                                                        }
                                                        _ = v59
                                                    }
                                                    if (468 >= fi17) {
                                                        if (1e10 > 1e10) {
                                                            v6 = append(v6, "\n\t")
                                                        } else {
                                                            if len(v6) > 0 {
                                                                v6[0] = ("a" + "foo bar")
                                                            }
                                                            v63 := GenS11{F0: 1, F1: (v12 / ((v12) | 1)), F2: (v15 + v15)}
                                                            fmt.Println(fi17, v5)
                                                            _ = v63
                                                        }
                                                        v64 := GenS12{F0: int16(v5)}
                                                        _ = v64
                                                    }
                                                } else {
                                                    var v65 string = "a"
                                                    var v66 uint32 = 65535
                                                    _ = v65
                                                    _ = v66
                                                }
                                            default:
                                                var v67 string = "hello"
                                                _ = v67
                                            }
                                        }
                                        v5 = 735
                                        if false {
                                            var v68 int = (-(fi16))
                                            v28 = (v28 & (v28 * v28))
                                            fmt.Println(fi17)
                                            _ = v68
                                        } else {
                                            v69 := GenS11{F0: 0, F1: -31, F2: v15}
                                            var v70 uint16 = (v1 % ((255) | 1))
                                            v6 = append(v6, ("" + "a"))
                                            _ = v69
                                            _ = v70
                                        }
                                        var v71 uint16 = v1
                                        fmt.Println(v12, v3, v12)
                                        v72 := GenS11{F0: 1000, F1: -128, F2: v15}
                                        if (true && false) {
                                            var v73 int64 = (-(v15))
                                            if ("foo bar" == "foo bar") {
                                                if len(v6) > 0 {
                                                    v6[0] = "\n\t"
                                                }
                                            }
                                            v6 = append(v6, "hello")
                                            _ = v73
                                        }
                                        if (!(1 == 981)) {
                                            fmt.Println(fi16)
                                            v5 = (v28 & 2)
                                        } else {
                                            v74 := GenS12{F0: 189}
                                            switch v3 {
                                            case -9223372036854775808:
                                                var v75 string = "hello"
                                                _ = v75
                                            default:
                                                v1 -= 255
                                            }
                                            fmt.Println(v5, fi16)
                                            _ = v74
                                        }
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v2 = append(v2, ("hello" + "\n\t"))
                                    }
                                    var v76 int64 = v15
                                    _ = v31
                                    _ = v76
                                }
                                _ = v28
                            }
                            var v77 uint32 = 4294967295
                            var v78 uint16 = v1
                            _ = v77
                            _ = v78
                        }
                        if (false && (1 <= -32768)) {
                            var v79 uint16 = (v1 ^ v1)
                            _ = v79
                        }
                        v80 := GenS13{F0: (^(v4))}
                        _ = v80
                    }
                    switch "a" {
                    case "foo bar":
                        if ((true || (128 >= 1)) && true) {
                            var v81 int16 = 42
                            var v82 int16 = (v81 << ((4294967295) & 63))
                            _ = v81
                            _ = v82
                        }
                    }
                }
            }
            var v83 bool = true
            _ = fi17
            _ = v83
        }
        v84 := GenS13{F0: (v3 << ((414) & 63))}
        _ = fi16
        _ = v84
    }
    v85 := GenS12{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
