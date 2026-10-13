package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 int8 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 string; F1 int8; F2 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int32; F1 uint16; F2 uint16 }

func main() {
    v1 := []int{-2147483648, 200, 9223372036854775807}
    var v2 int32 = 42
    var v3 uint16 = 255
    var v4 string = ""
    v5 := []uint32{2147483647, 65535}
    var v6 int = -609
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: -9223372036854775808, F1: 39}
            if (v3 < v3) {
                switch v6 {
                case 128:
                    var v10 int = v6
                    _ = v10
                case -9223372036854775808:
                    if ((false || ((^(fi8)) != -497)) || true) {
                        fmt.Println(v4, v4)
                        if true {
                            v3 -= 2
                        } else {
                            v4 = ""
                        }
                        if false {
                            var v11 uint = 65535
                            _ = v11
                        } else {
                            v12 := GenS2{F0: -1, F1: (v2 * v2), F2: -112}
                            _ = v12
                        }
                        switch fi8 {
                        case -1:
                            v6 -= -929
                        case 32767:
                            switch ("" + v4) {
                            case "\n\t":
                                switch fi8 {
                                case 287:
                                    fmt.Println(v6, v4, v6)
                                case -551:
                                    var v13 int64 = 1
                                    _ = v13
                                case 4294967295:
                                    if (18446744073709551615 == 1) {
                                        v14 := GenS3{F0: 4294967295, F1: "\n\t", F2: 255}
                                        fmt.Println(v6, v2, v4)
                                        v2 += v2
                                        _ = v14
                                    }
                                default:
                                    if (true && (v3 >= (v3 / ((1) | 1)))) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                            case "foo bar":
                                v2 = v2
                            case "hello":
                                var v15 bool = true
                                _ = v15
                            default:
                                var v16 uint32 = 0
                                _ = v16
                            }
                        }
                    }
                case -313:
                    var v17 string = (v4 + v4)
                    _ = v17
                }
            }
            _ = fi8
            _ = v9
        }
        if (1.5 > 3.14159) {
            for fi18 := 0; fi18 < 14; fi18++ {
                v2 += v2
                var v19 uint64 = 42
                if (float32(1.5) != 1.5) {
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                } else {
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                }
                _ = fi18
                _ = v19
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v20 int32 = v2
            _ = v20
        } else {
            v21 := GenS3{F0: 1, F1: v4, F2: 0}
            for fi22 := 0; fi22 < 18; fi22++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = "hello"
                _ = fi22
            }
            _ = v21
        }
        var v23 uint16 = v3
        _ = fi7
        _ = v23
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v24 := GenS2{F0: 127, F1: (v2 ^ v2), F2: 116}
    fmt.Println(v6, v2, v2)
    fmt.Println(v3, v4, v6)
    if (!true) {
        for fi25 := 0; fi25 < 20; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                fmt.Println(v3, v2)
                if false {
                    v27 := GenS3{F0: 2, F1: v4, F2: 65535}
                    var v28 uint = 356
                    if ((!(false && false)) || true) {
                        var v29 string = (v4 + "\n\t")
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                } else {
                    switch -612 {
                    case -340:
                        if (127 == 0) {
                            v1 = append(v1, (fi25 / ((v6) | 1)))
                        }
                    case 127:
                        if false {
                            var v30 int64 = -1
                            v31 := GenS1{F0: 9223372036854775807, F1: -90}
                            v32 := GenS2{F0: 1, F1: -1, F2: -128}
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            v6 = (fi26 + fi26)
                            var v33 int = 42
                            v4 = v4
                            var v34 uint16 = v3
                            if true {
                                v33 += -1
                                v34 -= v3
                            } else {
                                if false {
                                    v4 = "\n\t"
                                    var v35 bool = (((255 <= 0) || ((((v4 + v4) + "hello") + (("hello" + v4) + v4)) >= v4)) || (!((-(v2)) > v2)))
                                    _ = v35
                                } else {
                                    if (v3 != v34) {
                                        var v36 int64 = -1
                                        v2 = -1
                                        _ = v36
                                    } else {
                                        fmt.Println(v33, v6, v4)
                                        v34 = 1
                                        _ = len(v1)
                                    }
                                    if (4294967295 != 2) {
                                        if (((!(86 < 916)) || true) || false) {
                                            if (true || true) {
                                                if false {
                                                    v34 -= 715
                                                } else {
                                                    if true {
                                                        v37 := GenS3{F0: 65535, F1: "hello", F2: 877}
                                                        v38 := GenS4{F0: -1, F1: v4}
                                                        if (!(42 <= 42)) {
                                                            v39 := GenS2{F0: 42, F1: v2, F2: -1}
                                                            _ = v39
                                                        } else {
                                                            switch v4 {
                                                            case "foo bar":
                                                                v40 := GenS5{F0: 42}
                                                                _ = v40
                                                            case "a":
                                                                var v41 int64 = -667
                                                                _ = v41
                                                            case "hello":
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                            if false {
                                                                v4 = "hello"
                                                                var v42 uint16 = v3
                                                                var v43 int8 = 127
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                if false {
                                                                    v44 := GenS6{F0: v4, F1: 42, F2: (v34 - 611)}
                                                                    var v45 float32 = 1.5
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                                var v46 float32 = 3.14159
                                                                v47 := GenS7{F0: -955}
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        }
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        v34 = v3
                                                        v4 = "hello"
                                                    }
                                                    v48 := GenS7{F0: 1}
                                                    v34 = 42
                                                    v2 -= 719
                                                    _ = v48
                                                }
                                                v2 = v2
                                                fmt.Println(v3)
                                            } else {
                                                var v49 bool = true
                                                _ = v49
                                            }
                                            var v50 uint64 = 2
                                            v4 = v4
                                            _ = v50
                                        } else {
                                            fmt.Println(fi26)
                                        }
                                        var v51 uint8 = 229
                                        _ = v51
                                    } else {
                                        fmt.Println(fi26, v33)
                                    }
                                    v52 := GenS8{F0: true, F1: 128}
                                    _ = v52
                                }
                                if true {
                                    var v53 int = -127
                                    if ((^(v3)) == v34) {
                                        v6 = fi26
                                    } else {
                                        v54 := GenS9{F0: "hello", F1: 42, F2: ((v4 + ("a" + "foo bar")) + ((v4 + v4) + v4))}
                                        v33 = (fi25 ^ (fi25 - (v6 * v6)))
                                        switch (v53 % ((fi25) | 1)) {
                                        case 127:
                                            v3 -= (^(v34))
                                        case 628:
                                            var v55 int = 32768
                                            _ = v55
                                        }
                                        _ = v54
                                    }
                                    _ = v53
                                }
                                v33 = v6
                            }
                            _ = v33
                            _ = v34
                        }
                    case -2147483648:
                        v56 := GenS10{F0: 0}
                        _ = v56
                    default:
                        v6 = v6
                    }
                    v3 += v3
                }
                switch (fi26 / (((fi26 - fi25)) | 1)) {
                case -991:
                    var v57 uint16 = v3
                    _ = v57
                }
                _ = fi26
            }
            v3 = v3
            switch (v4 + v4) {
            case "\n\t":
                for fi58 := 0; fi58 < 11; fi58++ {
                    switch 4294967295 {
                    case 499:
                        if (48 >= 0) {
                            v3 -= v3
                        } else {
                            var v59 int32 = v2
                            if (true || (true || true)) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v59 -= (v2 % (((v2 ^ (-(v2)))) | 1))
                            _ = v59
                        }
                    }
                    v60 := GenS8{F0: false, F1: 176}
                    _ = fi58
                    _ = v60
                }
            case "a":
                switch v4 {
                case "":
                    var v61 uint = 986
                    _ = v61
                case "a":
                    switch "hello" {
                    case "foo bar":
                        var v62 float32 = 0.1
                        _ = v62
                    case "\n\t":
                        fmt.Println(v2)
                    case "":
                        if true {
                            fmt.Println(v2)
                            for fi63 := 0; fi63 < 7; fi63++ {
                                var v64 float32 = 1.5
                                switch 42 {
                                case 32767:
                                    v6 = (fi25 | v6)
                                case 128:
                                    var v65 float32 = ((1e10 + v64) + float32(float64((v64 - v64))))
                                    _ = v65
                                case 0:
                                    if (true && true) {
                                        var v66 uint16 = (v3 - v3)
                                        _ = v66
                                    }
                                }
                                _ = fi63
                                _ = v64
                            }
                        }
                    }
                default:
                    for fi67 := 0; fi67 < 12; fi67++ {
                        if ((2 > 4294967295) || (773 != 18446744073709551615)) {
                            v68 := GenS3{F0: 65535, F1: v4, F2: 0}
                            _ = len(v5)
                            v4 = ((v4 + (v4 + ((v4 + (("hello" + v4) + (v4 + ("foo bar" + v4)))) + v4))) + (v4 + ((v4 + (v4 + ((v4 + ((v4 + (((v4 + v4) + (v4 + (((v4 + ("foo bar" + v4)) + "") + ("" + v4)))) + "foo bar")) + ("a" + v4))) + v4))) + v4)))
                            _ = v68
                        } else {
                            fmt.Println(v6, fi25, v2)
                            var v69 int8 = 100
                            switch 567 {
                            case 1:
                                if true {
                                    v2 = (v2 + (v2 & (v2 & (v2 & v2))))
                                    _ = len(v5)
                                    if ((1.5 * (1e10 - float64(0.0))) == ((0.1 - -2.25) + float64(3.14159))) {
                                        var v70 float64 = 1e10
                                        if (3.14159 < -2.25) {
                                            switch (fi25 % ((-542) | 1)) {
                                            case 128:
                                                var v71 int64 = 0
                                                _ = v71
                                            case -971:
                                                v72 := GenS11{F0: (v3 >> ((463) & 63)), F1: 0, F2: 1000}
                                                _ = v72
                                            case 1:
                                                if true {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            default:
                                                fmt.Println(v69, v69)
                                            }
                                            var v73 uint32 = 2147483648
                                            switch ((("\n\t" + v4) + "hello") + (v4 + ("\n\t" + "a"))) {
                                            case "\n\t":
                                                var v74 float32 = -2.25
                                                _ = v74
                                            case "foo bar":
                                                fmt.Println(fi25, v3, fi67)
                                            case "a":
                                                if true {
                                                    var v75 float64 = v70
                                                    v69 -= v69
                                                    _ = v75
                                                } else {
                                                    if ((v6 - (-(fi25))) <= (fi25 / ((fi25) | 1))) {
                                                        var v76 float64 = v70
                                                        _ = v76
                                                    }
                                                    var v77 int32 = (-(v2))
                                                    _ = v77
                                                }
                                            }
                                            _ = v73
                                        }
                                        switch (fi67 << ((65535) & 63)) {
                                        case -338:
                                            if len(v1) > 0 {
                                                v1[0] = (v6 % ((v6) | 1))
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "\n\t":
                                                v78 := GenS4{F0: 42, F1: "hello"}
                                                _ = v78
                                            case "":
                                                v79 := GenS12{F0: 715}
                                                _ = v79
                                            default:
                                                var v80 int = fi67
                                                _ = v80
                                            }
                                            var v81 float64 = v70
                                            _ = v81
                                        }
                                        _ = v70
                                    }
                                }
                            }
                            _ = v69
                        }
                        var v82 int64 = 916
                        v2 = (v2 >> ((479) & 63))
                        v83 := GenS2{F0: 100, F1: v2, F2: 127}
                        _ = fi67
                        _ = v82
                        _ = v83
                    }
                }
            default:
                v3 = (v3 << ((0) & 63))
            }
            _ = fi25
        }
        var v84 bool = true
        var v85 uint64 = 860
        _ = v84
        _ = v85
    }
    v86 := GenS13{F0: 42, F1: (^(v3)), F2: (v3 >> ((4294967295) & 63))}
    var v87 int32 = -1
    var v88 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
    _ = v87
    _ = v88
}
