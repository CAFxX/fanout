package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 int32 }
func genFn0(p0 uint8, p1 uint64, p2 bool) uint32 {
    var v6 uint64 = (p1 * p1)
    p0 += p0
    v7 := GenS1{F0: -32768, F1: -715}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 95
    var v3 int32 = 811
    v4 := make(map[string]uint32)
    var v5 int16 = -1
    var v8 uint = 750
    switch "a" {
    case "":
        var v9 uint32 = 452
        _ = v9
    case "\n\t":
        switch 763 {
        case -2147483648:
            fmt.Println(v3, v8)
        case -1:
            var v10 string = "foo bar"
            _ = v10
        case -875:
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS1{F0: (v5 & v5), F1: (^(v5))}
                _ = fi11
                _ = v12
            }
        default:
            v13 := []int{}
            _ = v13
        }
    case "a":
        switch -849 {
        case -103:
            v14 := GenS2{F0: v3}
            _ = v14
        default:
            if false {
                var v15 int = 32768
                _ = v15
            } else {
                v5 = (v5 + v5)
                v8 += v8
                v16 := GenS3{F0: "a", F1: 2147483648}
                _ = v16
            }
        }
    default:
        for fi17 := 0; fi17 < 16; fi17++ {
            if (false || false) {
                if false {
                    var v18 string = ("hello" + "")
                    _ = v18
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    v20 := []int{}
                    _ = fi19
                    _ = v20
                }
                v21 := v4["\n\t"]
                _ = v21
            }
            _ = fi17
        }
    }
    if (!true) {
        var v22 uint16 = 0
        if false {
            v23 := v4[(("a" + "\n\t") + "foo bar")]
            for fi24 := 0; fi24 < 4; fi24++ {
                v2 -= uint8(v22)
                _ = fi24
            }
            _ = v23
        }
        _ = v22
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 float64 = 3.14159
            fmt.Println(v2, v3, v3)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    var v28 int16 = v5
    v3 = (v3 / ((v3) | 1))
    switch ("" + "foo bar") {
    case "a":
        _ = len(v4)
    case "hello":
        if false {
            v2 += v2
            v3 += 0
            v3 = 936
        } else {
            v1 += v1
        }
    default:
        if (2 != v8) {
            v29 := []int{128, 32767}
            v30 := v4[""]
            var v31 int32 = v3
            _ = v29
            _ = v30
            _ = v31
        } else {
            v32 := GenS4{F0: (2 > 42), F1: (v2 % ((2) | 1)), F2: v3}
            _ = v32
        }
    }
    v8 += (^(v8))
    v2 -= (v2 + (^(v2)))
    delete(v4, ("a" + "\n\t"))
    var v33 int8 = -128
    var v34 uint = 42
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(v2)
        _ = fi35
    }
    switch 748 {
    case 829:
        if false {
            if (-2147483648 >= 466) {
                var v36 bool = (9223372036854775807 >= 128)
                v36 = true
                switch -9223372036854775808 {
                case 4294967295:
                    var v37 bool = (1 > 2)
                    _ = v37
                case -338:
                    var v38 uint32 = 564
                    _ = v38
                }
                _ = v36
            }
        } else {
            v3 -= (v3 * (v3 + v3))
            switch "" {
            case "foo bar":
                v2 = (^(v2))
            case "hello":
                switch "a" {
                case "\n\t":
                    v39 := GenS5{F0: (v5 << (((v8 << ((v34) & 63))) & 63)), F1: (v5 * v28), F2: ("hello" + "a")}
                    _ = v39
                case "hello":
                    v40 := []int{0, -1, -957}
                    _ = v40
                }
            case "\n\t":
                v3 = -1000
            default:
                for fi41 := 0; fi41 < 14; fi41++ {
                    v3 = (v3 % ((v3) | 1))
                    switch "a" {
                    case "\n\t":
                        switch ("" + "a") {
                        case "hello":
                            v42 := []int{(fi41 ^ (fi41 - fi41))}
                            _ = v42
                        case "a":
                            for fi43 := 0; fi43 < 11; fi43++ {
                                var v44 int16 = (v28 * 276)
                                if ((!(int16(fi43) != v28)) && ((v2 >> ((v34) & 63)) == v2)) {
                                    v45 := GenS4{F0: false, F1: v2, F2: 1}
                                    _ = v45
                                } else {
                                    if false {
                                        if true {
                                            var v46 int64 = 1
                                            _ = v46
                                        } else {
                                            v47 := GenS6{F0: 248}
                                            v2 -= (v2 / ((v2) | 1))
                                            v2 -= (^(v2))
                                            _ = v47
                                        }
                                        if (false || false) {
                                            if (v8 <= v8) {
                                                fmt.Println(v44, v1)
                                                fmt.Println(v34)
                                                switch "hello" {
                                                case "foo bar":
                                                    v34 = (^(v34))
                                                }
                                            } else {
                                                _ = len(v4)
                                            }
                                            v4["\n\t"] = 204
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(v8, fi43)
                                            case "a":
                                                switch (fi43 + fi43) {
                                                case 42:
                                                    v8 += (v34 / ((uint(v2)) | 1))
                                                case 0:
                                                    v48 := GenS6{F0: 18446744073709551615}
                                                    _ = v48
                                                }
                                            case "":
                                                v28 = (v28 >> (((v34 / (((v34 & v34)) | 1))) & 63))
                                                if (!(918 == 255)) {
                                                    var v49 float32 = v1
                                                    v50 := []int{}
                                                    if false {
                                                        var v51 int64 = -148
                                                        v52 := GenS2{F0: v3}
                                                        v4[""] = genFn0(v2, uint64(v8), false)
                                                        if (81 != 902) {
                                                            if (!true) {
                                                                v51 -= -579
                                                                switch 1 {
                                                                case -2147483648:
                                                                    v53 := GenS5{F0: (^(v44)), F1: v28, F2: ("\n\t" + ("hello" + ""))}
                                                                    v54 := GenS6{F0: 796}
                                                                    var v55 uint16 = 491
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v55
                                                                case 4294967295:
                                                                    delete(v4, "foo bar")
                                                                    if (int32(v2) >= (v3 / (((v3 % ((v3) | 1))) | 1))) {
                                                                        v44 = v28
                                                                        var v56 float64 = float64(v49)
                                                                        var v57 int = 32767
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                default:
                                                                    var v58 uint = 2
                                                                    _ = v58
                                                                }
                                                                if len(v50) >= 2 {
                                                                    v50 = v50[1:len(v50)]
                                                                }
                                                            } else {
                                                                fmt.Println(v34, fi41)
                                                                switch "\n\t" {
                                                                case "":
                                                                    switch "a" {
                                                                    case "":
                                                                        v1 += float32(((0.1 / 0.1) - 1e10))
                                                                    default:
                                                                        fmt.Println(v51, v44, v8)
                                                                    }
                                                                case "foo bar":
                                                                    v50 = append(v50, 9223372036854775807)
                                                                case "hello":
                                                                    v5 = -4
                                                                }
                                                                if false {
                                                                    v59 := GenS1{F0: (v44 << (((v8 ^ (^(v34)))) & 63)), F1: (v5 & 42)}
                                                                    var v60 int = int(v33)
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    fmt.Println(v1, v3)
                                                                    if (v44 > v5) {
                                                                        if (778 != (v28 - (v28 / ((v5) | 1)))) {
                                                                            v8 = (v34 & 42)
                                                                        }
                                                                        var v61 uint8 = 152
                                                                        v2 += v2
                                                                        _ = v61
                                                                    } else {
                                                                        v5 += -1
                                                                        switch "" {
                                                                        case "a":
                                                                            switch "" {
                                                                            case "hello":
                                                                                v62 := GenS6{F0: 4294967295}
                                                                                _ = v62
                                                                            case "a":
                                                                                if ((v33 << ((uint(v2)) & 63)) != v33) {
                                                                                    v49 += ((-2.25 * float32(0.0)) / 1.5)
                                                                                    v5 = (^(v28))
                                                                                    fmt.Println(v44, v2, fi41)
                                                                                } else {
                                                                                    fmt.Println(v28, v34)
                                                                                    v63 := GenS7{F0: 65535}
                                                                                    _ = v63
                                                                                }
                                                                            case "":
                                                                                var v64 float32 = (((v1 * (-2.25 / 1.5)) * 0.0) / 3.14159)
                                                                                _ = v64
                                                                            }
                                                                        case "foo bar":
                                                                            if (!false) {
                                                                                v1 = float32(float64(float32((0.1 / 1.5))))
                                                                                if true {
                                                                                    _ = len(v4)
                                                                                    if false {
                                                                                        if len(v50) > 0 {
                                                                                            v50[0] = (^(fi41))
                                                                                        }
                                                                                        if false {
                                                                                            v5 = v5
                                                                                        } else {
                                                                                            if ((false || false) && ("a" != "foo bar")) {
                                                                                                if false {
                                                                                                    var v65 uint64 = 42
                                                                                                    _ = v65
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v66 := GenS7{F0: 547}
                                                                                        _ = v66
                                                                                    }
                                                                                    v33 = v33
                                                                                    var v67 float32 = -2.25
                                                                                    var v68 uint32 = genFn0(v2, 887, true)
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                }
                                                                            } else {
                                                                                switch (fi43 >> (((v34 ^ (v8 + 18446744073709551615))) & 63)) {
                                                                                case 32768:
                                                                                    if false {
                                                                                        v69 := GenS8{F0: (v3 - -2147483648), F1: v34}
                                                                                        fmt.Println(v2, v8)
                                                                                        _ = v69
                                                                                    } else {
                                                                                        v44 = v44
                                                                                    }
                                                                                case 4294967295:
                                                                                    v70 := v4[(((("a" + "hello") + "") + "\n\t") + (("" + "a") + "\n\t"))]
                                                                                    _ = v70
                                                                                default:
                                                                                    fmt.Println(fi41, v28)
                                                                                }
                                                                                if (!false) {
                                                                                    var v71 float64 = float64((((v1 * v49) - (v1 - (0.1 + 0.1))) + (((0.1 - v1) * -2.25) + (v49 / ((v49 + float32(0.0)) - v1)))))
                                                                                    _ = v71
                                                                                }
                                                                                var v72 uint = v34
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    var v73 int = (-(fi41))
                                                    if ((!false) || true) {
                                                        var v74 uint8 = (v2 ^ (v2 >> ((18446744073709551615) & 63)))
                                                        _ = v74
                                                    }
                                                    v75 := GenS5{F0: (v28 & v28), F1: (v28 << ((255) & 63)), F2: "\n\t"}
                                                    v76 := GenS9{F0: v8, F1: (v34 + (v8 >> ((v8) & 63)))}
                                                    _ = v49
                                                    _ = v50
                                                    _ = v73
                                                    _ = v75
                                                    _ = v76
                                                }
                                                v3 = v3
                                            }
                                        }
                                    } else {
                                        fmt.Println(v28)
                                        var v77 int16 = (v28 - v44)
                                        if true {
                                            if (false && (!false)) {
                                                var v78 int16 = 842
                                                _ = v78
                                            } else {
                                                v79 := []int{fi43, 128}
                                                if ("hello" < "\n\t") {
                                                    v80 := GenS10{F0: (fi43 + fi43)}
                                                    _ = v80
                                                }
                                                v81 := GenS9{F0: 0, F1: v34}
                                                _ = v79
                                                _ = v81
                                            }
                                            var v82 uint16 = 25
                                            if (("foo bar" + ("foo bar" + "")) < (("hello" + "\n\t") + ("" + "a"))) {
                                                v8 = (v8 | 65535)
                                                v83 := GenS10{F0: fi41}
                                                var v84 int16 = (v28 & v44)
                                                _ = v83
                                                _ = v84
                                            } else {
                                                v77 = v28
                                            }
                                            v85 := []int{(fi41 >> ((v8) & 63)), (fi43 << ((875) & 63)), fi41, 271}
                                            v86 := GenS7{F0: 2147483647}
                                            _ = v82
                                            _ = v85
                                            _ = v86
                                        } else {
                                            var v87 int8 = (-(v33))
                                            switch "a" {
                                            case "foo bar":
                                                v28 = (v44 | (v44 & (v44 | -32768)))
                                            }
                                            _ = len(v4)
                                            _ = v87
                                        }
                                        _ = v77
                                    }
                                }
                                _ = fi43
                                _ = v44
                            }
                        default:
                            v1 = v1
                        }
                    case "a":
                        for fi88 := 0; fi88 < 19; fi88++ {
                            v89 := []int{}
                            if (32767 < 760) {
                                var v90 uint8 = 42
                                v91 := GenS11{F0: "a", F1: -1}
                                _ = v90
                                _ = v91
                            } else {
                                v92 := v4[""]
                                var v93 int32 = v3
                                _ = v92
                                _ = v93
                            }
                            _ = fi88
                            _ = v89
                        }
                    case "":
                        delete(v4, "")
                    default:
                        for fi94 := 0; fi94 < 20; fi94++ {
                            var v95 int16 = v28
                            _ = fi94
                            _ = v95
                        }
                    }
                    _ = fi41
                }
            }
        }
    case 2147483647:
        fmt.Println(v34, v33, v2)
    default:
        var v96 uint64 = 1
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
