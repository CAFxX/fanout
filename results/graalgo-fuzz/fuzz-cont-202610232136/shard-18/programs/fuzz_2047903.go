package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int; F2 uint }
type GenS5 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int64; F1 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 uint64; F1 uint }
func genFn0(p0 int) uint8 {
    v4 := GenS1{F0: 267}
    switch p0 {
    case 32767:
        for fi5 := 0; fi5 < 15; fi5++ {
            p0 += fi5
            v6 := GenS2{F0: 557}
            _ = fi5
            _ = v6
        }
    case 9223372036854775807:
        if false {
            if true {
                p0 = (p0 % ((4294967295) | 1))
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 uint16 = 0
                    for fi9 := 0; fi9 < 12; fi9++ {
                        fmt.Println(fi7, v8, fi7)
                        _ = fi9
                    }
                    var v10 uint = 992
                    _ = fi7
                    _ = v10
                    _ = v8
                }
            }
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                if (!(!(false || true))) {
                    v12 := []int{(fi11 + p0), (-(p0)), (p0 - p0), p0}
                    _ = v12
                }
                v13 := GenS3{F0: p0}
                var v14 float32 = 0.0
                _ = fi11
                _ = v13
                _ = v14
            }
        }
    case 128:
        v15 := GenS4{F0: -1, F1: (p0 + (p0 + p0)), F2: 255}
        _ = v15
    default:
        p0 = (p0 % ((0) | 1))
    }
    v16 := GenS3{F0: (-(p0))}
    p0 = (p0 % ((p0) | 1))
    if false {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 uint16 = 194
            _ = fi17
            _ = v18
        }
        var v19 int32 = 2147483647
        if (!true) {
            p0 = 128
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                if (!true) {
                    if (true && false) {
                        p0 = p0
                        v19 += v19
                    } else {
                        fmt.Println(v19)
                        v19 -= v19
                        var v21 uint8 = 51
                        _ = v21
                    }
                    switch ("\n\t" + ("" + ("\n\t" + "a"))) {
                    case "foo bar":
                        v19 += v19
                    default:
                        if true {
                            switch ((("hello" + ("" + "a")) + ("" + "a")) + "") {
                            case "foo bar":
                                v22 := GenS5{F0: 781, F1: 214, F2: -9223372036854775808}
                                _ = v22
                            case "a":
                                for fi23 := 0; fi23 < 16; fi23++ {
                                    p0 += fi23
                                    var v24 uint32 = 2147483647
                                    _ = fi23
                                    _ = v24
                                }
                            case "":
                                p0 += (fi20 / ((p0) | 1))
                            default:
                                var v25 uint32 = 2147483648
                                _ = v25
                            }
                            p0 = (p0 + p0)
                            var v26 uint8 = 72
                            _ = v26
                        } else {
                            var v27 float64 = (float64(((1e10 / ((3.14159 * float32(((0.0 - (float64((float32((1.5 * -2.25)) * 0.0)) / (0.1 - 1.5))) - (float64(3.14159) / (1e10 - float64((1.5 - 3.14159))))))) + (1e10 - 0.1))) * (3.14159 * 1.5))) * (float64(3.14159) + 3.14159))
                            _ = v27
                        }
                    }
                    var v28 uint16 = 13
                    _ = v28
                }
                _ = fi20
            }
            var v29 int = (p0 & 9223372036854775807)
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS6{F0: 650}
                var v32 bool = (!("a" == (("foo bar" + "\n\t") + (("a" + "a") + "hello"))))
                for fi33 := 0; fi33 < 9; fi33++ {
                    if true {
                        switch (-(fi30)) {
                        case 0:
                            var v34 uint8 = 252
                            var v35 int = 42
                            _ = v34
                            _ = v35
                        case -262:
                            var v36 string = "\n\t"
                            _ = v36
                        default:
                            p0 += 2147483647
                            fmt.Println(fi30)
                            var v37 uint16 = 0
                            if true {
                                v38 := GenS7{F0: 2}
                                v39 := GenS1{F0: (^(v37))}
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint8 = 42
                                _ = v40
                            }
                            _ = v37
                        }
                        if ("\n\t" > "") {
                            switch "hello" {
                            case "a":
                                v41 := []int{}
                                v42 := GenS8{F0: 73}
                                _ = v41
                                _ = v42
                            case "":
                                p0 = p0
                            default:
                                v29 += 32767
                            }
                            v19 -= (v19 & -2147483648)
                            v19 = v19
                        } else {
                            v43 := GenS3{F0: (p0 ^ (fi30 >> ((42) & 63)))}
                            var v44 bool = ((42 > 503) || (!(!(!(int64(v29) > -1)))))
                            v19 -= 1
                            _ = v43
                            _ = v44
                        }
                    }
                    v19 += v19
                    _ = fi33
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        _ = v19
    } else {
        p0 = p0
    }
    _ = p0
    _ = v16
    _ = v4
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 2
    var v3 int = 9223372036854775807
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 uint16 = 1
        _ = fi45
        _ = v46
    }
    if true {
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 float64 = (0.0 - 3.14159)
            v48 += (v48 * v48)
            fmt.Println(v3, fi47)
            _ = fi47
            _ = v48
        }
        v49 := GenS9{F0: 9223372036854775807, F1: 18446744073709551615}
        _ = v49
    } else {
        var v50 float64 = (1e10 * (float64(0.1) * float64(0.0)))
        switch 332 {
        case 42:
            var v51 bool = false
            _ = v51
        case 2147483647:
            var v52 uint64 = 755
            _ = v52
        case 1:
            fmt.Println(v50, v2)
        }
        for fi53 := 0; fi53 < 19; fi53++ {
            fmt.Println(fi53, v50, v2)
            for fi54 := 0; fi54 < 6; fi54++ {
                v2 = (v2 + (v2 + genFn0(-652)))
                v55 := GenS10{F0: 1, F1: 4294967295}
                delete(v1, "hello")
                v50 = v50
                var v56 uint16 = 140
                _ = fi54
                _ = v55
                _ = v56
            }
            _ = fi53
        }
        _ = v50
    }
    v2 += (^(v2))
    switch v3 {
    case -747:
        var v57 int = (v3 | 9223372036854775807)
        _ = v57
    }
    _ = len(v1)
    v58 := []int{v3, v3}
    var v59 int64 = 60
    var v60 uint = 990
    if (0 >= int64(v2)) {
        fmt.Println(v59, v60)
        v61 := GenS5{F0: 65535, F1: 2, F2: v59}
        fmt.Println(v60, v2, v2)
        _ = v61
    } else {
        if (true && true) {
            if false {
                v62 := GenS11{F0: -775}
                var v63 int32 = -2147483648
                _ = v62
                _ = v63
            } else {
                v3 -= v3
                if false {
                    v64 := GenS1{F0: 772}
                    v2 -= (v2 / ((uint8(v60)) | 1))
                    var v65 int32 = 306
                    _ = v64
                    _ = v65
                } else {
                    if ((v60 & (v60 - 1)) == v60) {
                        v3 = (v3 & (v3 | -479))
                    }
                    v1[("foo bar" + "a")] = false
                    v60 += 998
                }
            }
            if (false || (4294967295 >= 2)) {
                if true {
                    v60 = v60
                }
            } else {
                switch (-(v3)) {
                case 42:
                    switch ("hello" + (((("" + "hello") + "a") + "\n\t") + "foo bar")) {
                    case "":
                        for fi66 := 0; fi66 < 12; fi66++ {
                            fmt.Println(v59)
                            for fi67 := 0; fi67 < 16; fi67++ {
                                v68 := GenS2{F0: 42}
                                fmt.Println(v3, fi66)
                                var v69 int64 = v59
                                v58 = append(v58, 42)
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            for fi70 := 0; fi70 < 9; fi70++ {
                                v71 := GenS12{F0: (("a" + "\n\t") + "hello"), F1: (^(v60))}
                                _ = fi70
                                _ = v71
                            }
                            _ = fi66
                        }
                    case "foo bar":
                        if (0 > 1) {
                            v72 := GenS13{F0: 844, F1: v60}
                            _ = v72
                        }
                    }
                }
            }
            v59 = v59
        }
    }
    v73 := GenS5{F0: 2, F1: 681, F2: v59}
    v74 := GenS13{F0: 652, F1: (v60 << ((v60) & 63))}
    if false {
        v59 = v59
    }
    fmt.Println(v59)
    v60 = v60
    var v75 uint32 = 2147483647
    fmt.Println(v75)
    v76 := GenS7{F0: (v75 % (((v75 % (((v75 & v75)) | 1))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v60
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
