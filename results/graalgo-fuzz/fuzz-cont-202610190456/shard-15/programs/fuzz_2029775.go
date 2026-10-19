package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool; F2 uint }
type GenS6 struct { F0 int64; F1 string; F2 uint8 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int64 }
func genFn0(p0 uint8, p1 uint64) uint {
    fmt.Println(p1)
    var v5 int16 = -1
    if ((p1 - (p1 | 18446744073709551615)) == 42) {
        fmt.Println(p1, v5, p0)
    } else {
        p0 += 2
    }
    v6 := GenS1{F0: (("hello" + "a") == "\n\t"), F1: -387}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int8) bool {
    var v7 float32 = -2.25
    v8 := []int{1}
    switch ("" + ("\n\t" + "")) {
    case "a":
        if false {
            switch ("" + "a") {
            case "hello":
                v9 := GenS1{F0: (true || true), F1: 987}
                _ = v9
            default:
                switch -689 {
                case 2147483647:
                    switch "hello" {
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 128:
                            var v10 uint = 2
                            _ = v10
                        case -278:
                            var v11 int = 688
                            _ = v11
                        case 0:
                            switch int(p0) {
                            case 2147483647:
                                var v12 int8 = p0
                                _ = v12
                            case -1:
                                v13 := GenS1{F0: false, F1: 1}
                                _ = v13
                            }
                        default:
                            v14 := GenS2{F0: 0, F1: false, F2: p0}
                            _ = v14
                        }
                    }
                default:
                    v15 := GenS2{F0: 42, F1: false, F2: 0}
                    _ = v15
                }
            }
            if (128 == 32768) {
                var v16 string = ("foo bar" + "\n\t")
                _ = v16
            }
            _ = len(v8)
        } else {
            fmt.Println(p0)
            if true {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int64 = 1
                    var v19 uint = 1
                    var v20 int64 = v18
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
            v8 = append(v8, 0)
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return (0 >= int64(p0))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    v4 := []uint32{353}
    if (0 <= 0) {
        if true {
            v2 = v2
            v2 = (v2 << ((92) & 63))
        }
        if true {
            var v21 float64 = 0.0
            var v22 string = ("" + "a")
            _ = v21
            _ = v22
        } else {
            var v23 int32 = -555
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: (v2 > (v2 * (v2 % (((v2 ^ (v2 % (((-(v2))) | 1)))) | 1)))), F1: -823}
            v2 = (v2 >> ((899) & 63))
            v4 = append(v4, 251)
            _ = fi24
            _ = v25
        }
    }
    v2 -= v2
    var v26 uint = 255
    v26 += (v26 + v26)
    var v27 string = "hello"
    v28 := GenS3{F0: 128}
    if (false && (42 <= 36)) {
        for fi29 := 0; fi29 < 4; fi29++ {
            switch v27 {
            case "hello":
                for fi30 := 0; fi30 < 12; fi30++ {
                    delete(v3, 0)
                    _ = fi30
                }
            case "\n\t":
                if (false || (true && true)) {
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                }
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 uint16 = 32767
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := GenS3{F0: 153}
                if (!false) {
                    switch fi33 {
                    case 1:
                        switch "\n\t" {
                        case "\n\t":
                            v26 += v26
                        case "a":
                            switch (fi33 ^ fi33) {
                            case -9223372036854775808:
                                v35 := v3[2147483648]
                                _ = v35
                            case 99:
                                if (42 >= (v2 << ((0) & 63))) {
                                    if true {
                                        delete(v1, (v27 + v27))
                                        v26 = v26
                                    }
                                    v2 = 0
                                } else {
                                    v36 := GenS1{F0: false, F1: 0}
                                    if (!(false || ((42 < 2) && false))) {
                                        if false {
                                            v4 = append(v4, 871)
                                            fmt.Println(v2)
                                            v37 := GenS4{F0: v27}
                                            var v38 bool = false
                                            if v38 {
                                                fmt.Println(v27)
                                                v27 = "a"
                                            } else {
                                                var v39 uint16 = 1
                                                fmt.Println(v39)
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                                _ = v39
                                            }
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v27 = v27
                                            fmt.Println(fi33, v2)
                                            var v40 uint32 = 2147483647
                                            _ = v40
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 2147483648
                                        }
                                    }
                                    var v41 int = (fi29 << ((0) & 63))
                                    v26 -= v26
                                    fmt.Println(v26, v27, v41)
                                    _ = v36
                                    _ = v41
                                }
                                v42 := GenS5{F0: false, F1: false, F2: (v26 / ((v26) | 1))}
                                v4 = append(v4, 2147483648)
                                if false {
                                    v43 := GenS2{F0: 2, F1: (0.1 != (0.1 / (1e10 - (0.1 * (0.1 * 1.5))))), F2: -1}
                                    _ = v43
                                } else {
                                    var v44 uint = 0
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    _ = v44
                                }
                                v2 = v2
                                _ = v42
                            case 221:
                                if (143 == 128) {
                                    if true {
                                        v45 := GenS3{F0: 42}
                                        var v46 bool = (v27 < v27)
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 string = (("a" + "") + "foo bar")
                                        _ = v47
                                    }
                                    v2 = v2
                                    var v48 uint = 65535
                                    _ = v48
                                } else {
                                    var v49 uint16 = 489
                                    _ = v49
                                }
                            }
                        }
                    default:
                        if (-2.25 > ((3.14159 + 1e10) / float32(-2.25))) {
                            if ((((-9223372036854775808 < 1) && (-1 >= v2)) || true) && false) {
                                var v50 string = v27
                                if (v2 != (v2 << ((genFn0(106, 4294967295)) & 63))) {
                                    v51 := v1["foo bar"]
                                    _ = len(v3)
                                    var v52 float64 = (1.5 * 3.14159)
                                    v53 := GenS3{F0: 0}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                _ = v50
                            } else {
                                v26 = v26
                                switch v27 {
                                case "a":
                                    v2 += (v2 & (v2 & v2))
                                case "":
                                    v27 = (v27 + ("\n\t" + v27))
                                case "foo bar":
                                    switch "foo bar" {
                                    case "a":
                                        var v54 string = (v27 + v27)
                                        _ = v54
                                    case "":
                                        if true {
                                            if ((1e10 * 0.1) >= 1.5) {
                                                v26 = (v26 ^ genFn0(127, 0))
                                            } else {
                                                fmt.Println(fi29, v2, v27)
                                                if (0 > -1) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    delete(v3, 91)
                                                }
                                                if (!(float64(0.0) == 0.0)) {
                                                    var v55 string = (((v27 + v27) + v27) + "foo bar")
                                                    var v56 uint64 = 1
                                                    _ = v55
                                                    _ = v56
                                                }
                                            }
                                            fmt.Println(v26, fi33, v27)
                                        }
                                    default:
                                        if ((^(v2)) == 1) {
                                            var v57 int8 = int8(fi33)
                                            _ = v57
                                        }
                                    }
                                default:
                                    var v58 uint16 = 158
                                    _ = v58
                                }
                                if false {
                                    if true {
                                        v59 := GenS6{F0: 0, F1: v27, F2: 255}
                                        _ = v59
                                    }
                                    if false {
                                        v1["\n\t"] = 725
                                        if true {
                                            v60 := GenS7{F0: -1, F1: 42}
                                            v61 := GenS2{F0: 65535, F1: true, F2: -128}
                                            v26 += 255
                                            _ = v60
                                            _ = v61
                                        }
                                    } else {
                                        var v62 uint64 = 42
                                        _ = v62
                                    }
                                    var v63 int16 = (v2 << ((v26) & 63))
                                    _ = v63
                                } else {
                                    fmt.Println(v26, fi33, fi33)
                                    v64 := GenS8{F0: 0}
                                    _ = v64
                                }
                            }
                            v65 := GenS1{F0: true, F1: 1}
                            _ = v65
                        } else {
                            var v66 int64 = -1
                            _ = v66
                        }
                    }
                    switch (fi29 & 32767) {
                    case 42:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    }
                    var v67 uint16 = 32767
                    _ = v67
                } else {
                    if (!true) {
                        if (true && false) {
                            v68 := GenS5{F0: true, F1: (-1 <= -25), F2: (v26 >> ((255) & 63))}
                            _ = v68
                        } else {
                            var v69 uint16 = 32767
                            var v70 int8 = 0
                            if (-2.25 != 0.1) {
                                v69 -= v69
                                v26 = v26
                                v71 := GenS4{F0: v27}
                                v72 := GenS6{F0: -1, F1: "\n\t", F2: 128}
                                var v73 float32 = (1e10 * 3.14159)
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            _ = v69
                            _ = v70
                        }
                        v2 -= v2
                        var v74 float32 = 3.14159
                        _ = v74
                    } else {
                        v2 += v2
                        v26 = (^(v26))
                        v75 := GenS8{F0: -128}
                        _ = v75
                    }
                    switch fi33 {
                    case 1:
                        v76 := GenS9{F0: v26, F1: 396}
                        _ = v76
                    case 590:
                        fmt.Println(v26, v26)
                    }
                    v2 += (v2 * (v2 + (v2 ^ (v2 ^ (v2 >> ((42) & 63))))))
                    v2 -= v2
                }
                v27 = "\n\t"
                _ = fi33
                _ = v34
            }
            _ = fi29
        }
        for fi77 := 0; fi77 < 1; fi77++ {
            var v78 int64 = -173
            _ = fi77
            _ = v78
        }
        if len(v4) > 0 {
            v4[0] = 510
        }
    } else {
        var v79 uint16 = 1
        _ = v79
    }
    if true {
        v26 += v26
        if false {
            for fi80 := 0; fi80 < 5; fi80++ {
                if false {
                    var v81 uint8 = 73
                    _ = v81
                }
                switch fi80 {
                case -9223372036854775808:
                    v3[0] = fi80
                default:
                    var v82 uint8 = 42
                    _ = v82
                }
                _ = fi80
            }
            var v83 int32 = -2147483648
            _ = v83
        } else {
            v84 := v1[("hello" + "\n\t")]
            _ = v84
        }
    }
    var v85 int16 = v2
    v85 += -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v85
}
