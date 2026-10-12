package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int16; F2 uint16 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 string; F1 string; F2 string }

func main() {
    var v1 bool = false
    v2 := []int{128, 4294967295}
    v3 := make(map[string]uint32)
    var v4 int = 958
    v5 := GenS1{F0: 255, F1: v4}
    fmt.Println(v4, v1)
    v6 := v3["hello"]
    var v7 string = "foo bar"
    if (false && v1) {
        fmt.Println(v1, v4, v7)
    }
    switch (v4 * (v4 / ((v4) | 1))) {
    case 2147483647:
        if (float64((1e10 + 1.5)) > 1.5) {
            var v8 uint8 = 221
            if (!v1) {
                var v9 uint64 = 334
                v9 -= (v9 & 2)
                if ((v9 << ((4294967295) & 63)) <= 42) {
                    v10 := GenS1{F0: 0, F1: (v4 * int(v8))}
                    if false {
                        fmt.Println(v9, v1, v8)
                        var v11 int8 = 1
                        _ = v11
                    }
                    v7 = (v7 + v7)
                    _ = v10
                } else {
                    var v12 int = -604
                    for fi13 := 0; fi13 < 18; fi13++ {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            v6 += (v6 - 2147483647)
                            _ = fi14
                        }
                        _ = fi13
                    }
                    v15 := GenS2{F0: (v6 & (v6 / ((65535) | 1))), F1: (255 < 648)}
                    _ = v12
                    _ = v15
                }
                _ = v9
            }
            fmt.Println(v7, v8, v6)
            _ = v8
        } else {
            if false {
                for fi16 := 0; fi16 < 14; fi16++ {
                    fmt.Println(v4)
                    var v17 uint8 = 45
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi18
                }
                var v19 int8 = 0
                _ = v19
            }
        }
    default:
        delete(v3, v7)
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v1 {
            var v21 uint64 = 254
            var v22 uint16 = 42
            var v23 float64 = 0.1
            _ = v21
            _ = v22
            _ = v23
        } else {
            v24 := v3["a"]
            if (9223372036854775807 < -146) {
                v25 := GenS1{F0: 1, F1: (v4 - (fi20 + v4))}
                if ("" >= "foo bar") {
                    v26 := GenS2{F0: v24, F1: (521 != 2)}
                    switch v7 {
                    case "":
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v7 = v7
                            _ = fi27
                        }
                    default:
                        var v28 string = v7
                        _ = v28
                    }
                    v4 = v4
                    _ = v26
                } else {
                    v29 := GenS3{F0: true, F1: 0, F2: 42}
                    v1 = (v1 || (42 <= 220))
                    var v30 bool = (-1 != -9223372036854775808)
                    _ = v29
                    _ = v30
                }
                _ = v25
            } else {
                if (v4 > v4) {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        var v32 uint8 = 32
                        if v1 {
                            if v1 {
                                switch ("a" + ((("" + "a") + v7) + "a")) {
                                case "a":
                                    v33 := GenS4{F0: (v32 ^ (v32 * 42))}
                                    _ = v33
                                }
                                var v34 uint32 = 2147483647
                                _ = len(v3)
                                fmt.Println(fi31, v4, v6)
                                if ((32767 != 2) && v1) {
                                    _ = len(v3)
                                    if (v1 && (v1 || (2147483648 == (v34 / ((v34) | 1))))) {
                                        v32 = v32
                                        if ((v32 % (((v32 / ((255) | 1))) | 1)) < 144) {
                                            var v35 uint16 = 1
                                            var v36 int = (^(fi31))
                                            var v37 int = fi31
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                    } else {
                                        v7 = v7
                                        switch (v7 + v7) {
                                        case "foo bar":
                                            v3["hello"] = (v34 ^ 142)
                                        case "a":
                                            v1 = ((^(v34)) != v34)
                                        case "\n\t":
                                            v34 += (v6 / ((v34) | 1))
                                        default:
                                            var v38 int32 = -2147483648
                                            if false {
                                                fmt.Println(fi20, v7, fi20)
                                            }
                                            _ = v38
                                        }
                                    }
                                    if (!v1) {
                                        v1 = v1
                                        switch v7 {
                                        case "hello":
                                            switch v7 {
                                            case "":
                                                v39 := GenS3{F0: v1, F1: 100, F2: -740}
                                                _ = v39
                                            case "a":
                                                if false {
                                                    fmt.Println(v4, v1)
                                                } else {
                                                    var v40 int8 = 0
                                                    switch fi20 {
                                                    case 880:
                                                        v41 := GenS5{F0: v40, F1: 1, F2: 0}
                                                        _ = v41
                                                    case 198:
                                                        switch "" {
                                                        case "foo bar":
                                                            v7 = v7
                                                        case "\n\t":
                                                            var v42 float32 = 1e10
                                                            _ = v42
                                                        default:
                                                            delete(v3, (v7 + (v7 + (v7 + "\n\t"))))
                                                        }
                                                    case 128:
                                                        fmt.Println(v24)
                                                    default:
                                                        fmt.Println(v4, v6, v40)
                                                    }
                                                    var v43 uint32 = v24
                                                    fmt.Println(v43, v6, fi20)
                                                    v34 += (v24 ^ v24)
                                                    switch fi31 {
                                                    case 933:
                                                        if v1 {
                                                            var v44 uint32 = (v34 ^ (v6 | v43))
                                                            var v45 int8 = 42
                                                            var v46 string = "a"
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    case 128:
                                                        var v47 int16 = -828
                                                        _ = v47
                                                    case 9223372036854775807:
                                                        if false {
                                                            var v48 uint32 = (v43 ^ 74)
                                                            v2 = append(v2, 32767)
                                                            _ = v48
                                                        } else {
                                                            v49 := GenS6{F0: -2147483648, F1: uint8(v24)}
                                                            var v50 int = fi20
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    }
                                                    _ = v40
                                                    _ = v43
                                                }
                                            case "foo bar":
                                                var v51 int64 = 42
                                                _ = v51
                                            default:
                                                var v52 int = (fi31 / ((int(v34)) | 1))
                                                _ = v52
                                            }
                                        default:
                                            var v53 int8 = 42
                                            _ = v53
                                        }
                                        fmt.Println(fi20)
                                    }
                                } else {
                                    switch v7 {
                                    case "hello":
                                        switch -866 {
                                        case 1:
                                            v6 = (v6 % ((904) | 1))
                                        case 128:
                                            var v54 int32 = -611
                                            _ = v54
                                        case 42:
                                            if false {
                                                fmt.Println(v24, fi20, v32)
                                                v55 := GenS5{F0: 42, F1: 42, F2: 1}
                                                _ = v55
                                            }
                                        }
                                    case "foo bar":
                                        switch 128 {
                                        case -68:
                                            var v56 float64 = -2.25
                                            _ = v56
                                        case 32767:
                                            if (!(1.5 < 1.5)) {
                                                var v57 string = v7
                                                v58 := GenS7{F0: 255}
                                                v1 = v1
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                    case "":
                                        switch (v7 + ((v7 + "foo bar") + "")) {
                                        case "\n\t":
                                            var v59 uint64 = 2
                                            _ = v59
                                        case "":
                                            var v60 uint32 = v6
                                            _ = v60
                                        case "foo bar":
                                            var v61 int8 = -47
                                            _ = v61
                                        default:
                                            v7 = v7
                                        }
                                    }
                                    if v1 {
                                        var v62 uint8 = v32
                                        fmt.Println(v32)
                                        v63 := GenS8{F0: 1000, F1: 214, F2: 1}
                                        _ = v62
                                        _ = v63
                                    }
                                    var v64 int32 = -1
                                    _ = v64
                                }
                                _ = v34
                            } else {
                                switch "" {
                                case "a":
                                    v32 = 2
                                case "hello":
                                    v4 = (v4 % ((2147483647) | 1))
                                case "\n\t":
                                    switch v7 {
                                    case "":
                                        fmt.Println(v24, v32, v1)
                                    default:
                                        v1 = v1
                                    }
                                default:
                                    var v65 int = 1
                                    var v66 string = v7
                                    _ = v65
                                    _ = v66
                                }
                                v67 := v3["a"]
                                _ = v67
                            }
                            var v68 string = v7
                            _ = v68
                        }
                        if (3.14159 < 3.14159) {
                            v24 = (v6 * (v6 - 2147483648))
                            v6 -= v24
                        }
                        _ = fi31
                        _ = v32
                    }
                    for fi69 := 0; fi69 < 16; fi69++ {
                        if true {
                            v24 -= v24
                            fmt.Println(fi69, v7)
                        } else {
                            v4 = fi20
                        }
                        _ = fi69
                    }
                    var v70 string = (("\n\t" + ((v7 + v7) + v7)) + v7)
                    _ = v70
                }
                var v71 uint16 = 1
                var v72 int64 = -9223372036854775808
                _ = v71
                _ = v72
            }
            if v1 {
                var v73 uint32 = v6
                v74 := GenS1{F0: 668, F1: (^(fi20))}
                fmt.Println(fi20, v1, v7)
                _ = v73
                _ = v74
            } else {
                for fi75 := 0; fi75 < 16; fi75++ {
                    delete(v3, "a")
                    var v76 uint16 = 402
                    fmt.Println(v76, fi75, v7)
                    _ = fi75
                    _ = v76
                }
                switch (v4 / ((4294967295) | 1)) {
                case -1:
                    var v77 int = 0
                    _ = v77
                case 128:
                    if ((639 <= 255) || v1) {
                        switch (("\n\t" + "\n\t") + "\n\t") {
                        case "\n\t":
                            var v78 int16 = 180
                            _ = v78
                        case "":
                            var v79 float64 = 1e10
                            _ = v79
                        default:
                            _ = len(v3)
                        }
                    } else {
                        if ((!(18446744073709551615 <= 4294967295)) && (966 <= 1000)) {
                            v80 := GenS9{F0: -881, F1: fi20}
                            var v81 uint32 = v6
                            _ = v80
                            _ = v81
                        } else {
                            v82 := v3["hello"]
                            v24 = 457
                            _ = v82
                        }
                        v83 := GenS10{F0: v7, F1: ("\n\t" + ""), F2: v7}
                        _ = v83
                    }
                default:
                    v84 := GenS3{F0: (!v1), F1: 28, F2: -390}
                    _ = v84
                }
            }
            _ = v24
        }
        for fi85 := 0; fi85 < 12; fi85++ {
            fmt.Println(fi20)
            v86 := GenS3{F0: v1, F1: 1, F2: 9223372036854775807}
            _ = fi85
            _ = v86
        }
        _ = fi20
    }
    if len(v2) > 0 {
        v2[0] = (v4 * v4)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v4 {
    case 9223372036854775807:
        var v87 int64 = 0
        _ = v87
    case -2147483648:
        var v88 uint16 = 0
        _ = v88
    case 4294967295:
        v89 := GenS10{F0: "", F1: v7, F2: "\n\t"}
        _ = v89
    }
    switch ((("" + "\n\t") + v7) + "") {
    case "\n\t":
        if false {
            switch "a" {
            case "a":
                v6 = (^(v6))
            }
        }
    default:
        if true {
            fmt.Println(v7, v7)
        }
    }
    var v90 int64 = 0
    var v91 float32 = 1.5
    var v92 uint16 = 2
    var v93 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
