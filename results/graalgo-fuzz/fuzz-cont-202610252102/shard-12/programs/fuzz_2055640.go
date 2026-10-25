package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8; F1 int32 }
type GenS12 struct { F0 int8; F1 uint8; F2 int }
type GenS13 struct { F0 uint8; F1 int8 }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int; F1 uint8; F2 uint16 }
type GenS17 struct { F0 uint8; F1 int64 }
type GenS18 struct { F0 int64; F1 uint }

func main() {
    var v1 int32 = 1
    v2 := []string{}
    var v3 bool = true
    var v4 int16 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v5 uint = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = float32(3.14159)
        _ = fi6
        _ = v7
    }
    var v8 int8 = 32
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 = v3
        _ = fi10
    }
    var v11 uint = v5
    if (false || (float32(((0.0 / 3.14159) - -2.25)) == (((-2.25 / 1.0) * float32(0.0)) / 1e10))) {
        var v12 uint32 = 2
        _ = v12
    }
    if len(v2) > 0 {
        v2[0] = ((("hello" + "foo bar") + "a") + ("hello" + "a"))
    }
    if false {
        switch 803 {
        case -2147483648:
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(v5, v3)
                switch 710 {
                case 4294967295:
                    for fi14 := 0; fi14 < 10; fi14++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v15 uint32 = 568
                        v8 = v8
                        _ = fi14
                        _ = v15
                    }
                case 808:
                    v16 := GenS1{F0: (fi13 - -1), F1: fi13, F2: 372}
                    _ = v16
                default:
                    switch (fi13 % ((fi13) | 1)) {
                    case -871:
                        v11 = (^(v11))
                    case 8:
                        v17 := GenS1{F0: (fi13 % ((32768) | 1)), F1: fi13, F2: 1}
                        _ = v17
                    case 1:
                        _ = len(v2)
                    default:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v5 += v5
                            v8 -= v8
                            switch (^(fi13)) {
                            case -1:
                                if ((v5 + (v11 << ((827) & 63))) != v11) {
                                    var v19 uint = (v11 + (v11 % ((v11) | 1)))
                                    if v3 {
                                        fmt.Println(v5, fi13)
                                    } else {
                                        v20 := GenS2{F0: 2147483647, F1: (v1 % ((v1) | 1)), F2: (v4 * 0)}
                                        if v3 {
                                            switch "" {
                                            case "\n\t":
                                                var v21 int16 = (v4 / ((v4) | 1))
                                                _ = v21
                                            case "foo bar":
                                                var v22 int16 = (v4 << (((v19 * (^(v11)))) & 63))
                                                _ = v22
                                            case "":
                                                v1 += v1
                                            default:
                                                v23 := GenS3{F0: (fi13 * 42), F1: 9223372036854775807, F2: 42}
                                                _ = v23
                                            }
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = ("" + "")
                                            }
                                            if true {
                                                fmt.Println(v1)
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    var v24 uint64 = 50
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                    var v25 uint16 = v9
                                                    _ = v24
                                                    _ = v25
                                                case 127:
                                                    v9 -= 1
                                                }
                                                fmt.Println(v5, v3)
                                            }
                                        }
                                        _ = v20
                                    }
                                    _ = v19
                                }
                            case -2147483648:
                                if (v3 && true) {
                                    if true {
                                        fmt.Println(v11)
                                    }
                                } else {
                                    if v3 {
                                        if true {
                                            switch ("a" + "foo bar") {
                                            case "":
                                                if (0.1 != 0.0) {
                                                    switch fi18 {
                                                    case 4294967295:
                                                        v9 -= 42
                                                    case 0:
                                                        var v26 int16 = v4
                                                        _ = v26
                                                    default:
                                                        var v27 int16 = -32768
                                                        _ = v27
                                                    }
                                                    switch "" {
                                                    case "a":
                                                        v28 := GenS4{F0: 4294967295, F1: v11}
                                                        _ = v28
                                                    default:
                                                        var v29 int64 = -9223372036854775808
                                                        _ = v29
                                                    }
                                                }
                                            default:
                                                var v30 uint = 502
                                                _ = v30
                                            }
                                            if (42 == 1) {
                                                var v31 int8 = v8
                                                fmt.Println(v31, v5)
                                                _ = v31
                                            } else {
                                                if (!(1.5 == 3.14159)) {
                                                    v32 := GenS1{F0: fi18, F1: 597, F2: uint64(v5)}
                                                    v33 := GenS5{F0: (v9 - (v9 / (((^(v9))) | 1))), F1: fi13}
                                                    var v34 int16 = 1000
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    if true {
                                                        var v35 uint16 = 1
                                                        _ = v35
                                                    }
                                                    if (!(v3 && (!true))) {
                                                        switch ("a" + "foo bar") {
                                                        case "hello":
                                                            v2 = append(v2, "\n\t")
                                                        }
                                                        v36 := GenS2{F0: 4294967295, F1: (v1 & 2147483647), F2: int16(v9)}
                                                        switch fi13 {
                                                        case 4294967295:
                                                            v11 -= (v5 | (v11 * (v5 & 42)))
                                                        case 127:
                                                            var v37 int = (fi13 >> ((v5) & 63))
                                                            _ = v37
                                                        default:
                                                            var v38 int64 = -532
                                                            _ = v38
                                                        }
                                                        _ = v36
                                                    } else {
                                                        switch (fi13 + (^(fi13))) {
                                                        case 42:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case 9223372036854775807:
                                                            v3 = v3
                                                        case 2147483647:
                                                            var v39 uint32 = 138
                                                            _ = v39
                                                        }
                                                        v8 += (v8 | (v8 + v8))
                                                        v3 = ((v1 % (((-(v1))) | 1)) >= v1)
                                                    }
                                                    v40 := GenS5{F0: (v9 << ((v11) & 63)), F1: 4294967295}
                                                    if v3 {
                                                        v5 = 33
                                                        switch "" {
                                                        case "\n\t":
                                                            if (2 > (^(v9))) {
                                                                var v41 uint32 = 889
                                                                v2 = append(v2, ("foo bar" + "foo bar"))
                                                                if (((true && (2 <= 0)) || true) || true) {
                                                                    if v3 {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (("\n\t" + "\n\t") + "hello")
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v11, fi18)
                                                                            _ = len(v2)
                                                                            var v42 int = (fi18 * -935)
                                                                            v43 := GenS3{F0: fi13, F1: 42, F2: 4294967295}
                                                                            v11 += 255
                                                                            _ = v42
                                                                            _ = v43
                                                                        } else {
                                                                            var v44 uint16 = v9
                                                                            fmt.Println(v9, v11)
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                    if (v41 == v41) {
                                                                        if (v5 == (v11 & 2)) {
                                                                            if (v3 || false) {
                                                                                fmt.Println(fi13, v8)
                                                                            }
                                                                        } else {
                                                                            var v45 int32 = v1
                                                                            v45 -= (-(v1))
                                                                            _ = v45
                                                                        }
                                                                        v46 := GenS5{F0: 255, F1: 4294967295}
                                                                        switch "" {
                                                                        case "":
                                                                            if v3 {
                                                                                fmt.Println(v4, fi13, v8)
                                                                                v2 = append(v2, "a")
                                                                                v47 := GenS5{F0: v9, F1: (fi13 >> ((v5) & 63))}
                                                                                v11 += 255
                                                                                var v48 int = (fi13 + (fi18 & fi18))
                                                                                if (v48 < (fi18 ^ (^(fi13)))) {
                                                                                    v49 := GenS6{F0: v11, F1: 42, F2: 84}
                                                                                    if (((v1 | v1) <= -27) && true) {
                                                                                        v11 += (v5 & 805)
                                                                                    } else {
                                                                                        var v50 float64 = 3.14159
                                                                                        if v3 {
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                if (!((^(v11)) < 0)) {
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = ("foo bar" + ("" + (((("" + "hello") + ("hello" + ("" + ("foo bar" + "\n\t")))) + ("a" + "hello")) + "\n\t")))
                                                                                                    }
                                                                                                    var v51 uint32 = v41
                                                                                                    fmt.Println(v8)
                                                                                                    _ = v51
                                                                                                } else {
                                                                                                    var v52 int = fi18
                                                                                                    switch v52 {
                                                                                                    case 0:
                                                                                                        switch fi18 {
                                                                                                        case 42:
                                                                                                            var v53 uint = 312
                                                                                                            _ = v53
                                                                                                        case -226:
                                                                                                            var v54 uint64 = 0
                                                                                                            _ = v54
                                                                                                        default:
                                                                                                            fmt.Println(fi18)
                                                                                                        }
                                                                                                    case 42:
                                                                                                        fmt.Println(v48, v41, v50)
                                                                                                    default:
                                                                                                        v3 = (!((v41 + 2147483648) == (v41 + (v41 | 42))))
                                                                                                    }
                                                                                                    var v55 int32 = (-(v1))
                                                                                                    v1 = 0
                                                                                                    _ = v52
                                                                                                    _ = v55
                                                                                                }
                                                                                            }
                                                                                            var v56 string = (("" + "hello") + ("a" + "foo bar"))
                                                                                            _ = v56
                                                                                        } else {
                                                                                            var v57 uint8 = 127
                                                                                            if (!false) {
                                                                                                var v58 int32 = 1
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v57 = (v57 & uint8(v9))
                                                                                                _ = v58
                                                                                            }
                                                                                            _ = v57
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(v11)
                                                                                            var v59 uint16 = v9
                                                                                            v60 := GenS1{F0: int(v5), F1: 4294967295, F2: 41}
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        } else {
                                                                                            v3 = (!v3)
                                                                                            v61 := GenS7{F0: uint32(v11), F1: (v41 << (((^(v5))) & 63))}
                                                                                            v62 := GenS8{F0: 42}
                                                                                            _ = len(v2)
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        }
                                                                                        _ = v50
                                                                                    }
                                                                                    _ = v49
                                                                                } else {
                                                                                    if ((3.14159 > (((1.5 + 3.14159) * 0.1) + 1.5)) && v3) {
                                                                                        if (419 == 898) {
                                                                                            var v63 bool = (1 <= 248)
                                                                                            v63 = v63
                                                                                            _ = v63
                                                                                        }
                                                                                    } else {
                                                                                        var v64 int64 = 1
                                                                                        _ = v64
                                                                                    }
                                                                                    v1 = int32(v4)
                                                                                }
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                        default:
                                                                            var v65 float32 = 3.14159
                                                                            _ = v65
                                                                        }
                                                                        _ = v46
                                                                    } else {
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "\n\t":
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case "a":
                                                                            v66 := GenS7{F0: (v41 / ((v41) | 1)), F1: (v41 % ((v41) | 1))}
                                                                            _ = v66
                                                                        default:
                                                                            switch fi18 {
                                                                            case -9223372036854775808:
                                                                                v67 := GenS9{F0: (v9 % (((v9 + v9)) | 1)), F1: -1}
                                                                                _ = v67
                                                                            case -561:
                                                                                var v68 int16 = v4
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        var v69 string = "hello"
                                                                        var v70 bool = v3
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                } else {
                                                                    var v71 int = -2147483648
                                                                    _ = v71
                                                                }
                                                                _ = v41
                                                            }
                                                            v8 = -1
                                                            v72 := GenS8{F0: 2}
                                                            _ = v72
                                                        case "a":
                                                            switch fi18 {
                                                            case -96:
                                                                v3 = v3
                                                            }
                                                        case "foo bar":
                                                            var v73 uint = (v11 - v11)
                                                            if false {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v74 uint16 = 255
                                                                v11 = v5
                                                                _ = v74
                                                            } else {
                                                                var v75 int8 = (v8 * v8)
                                                                if (v73 > v73) {
                                                                    if v3 {
                                                                        if v3 {
                                                                            var v76 int16 = (v4 / (((^(v4))) | 1))
                                                                            switch (fi13 + fi18) {
                                                                            case 1:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case 42:
                                                                                if v3 {
                                                                                    var v77 uint16 = (v9 + (v9 >> ((v5) & 63)))
                                                                                    var v78 int = (fi13 - 128)
                                                                                    v77 = (v77 % ((2) | 1))
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                } else {
                                                                                    switch fi18 {
                                                                                    case 1:
                                                                                        if false {
                                                                                            fmt.Println(v8, fi13, v1)
                                                                                            v2 = append(v2, ((("" + "hello") + ("hello" + "hello")) + "\n\t"))
                                                                                        } else {
                                                                                            fmt.Println(fi18, v11)
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = ("hello" + ("foo bar" + "foo bar"))
                                                                                            }
                                                                                        }
                                                                                    case 0:
                                                                                        fmt.Println(v75, v8)
                                                                                    }
                                                                                }
                                                                            case -756:
                                                                                var v79 int32 = -1
                                                                                _ = v79
                                                                            default:
                                                                                if (1 > int64(v4)) {
                                                                                    if (!v3) {
                                                                                        v2 = append(v2, "foo bar")
                                                                                    }
                                                                                    if (((0.0 + ((0.0 * 1e10) - (1e10 + (1.5 / 1.5)))) + 3.14159) != 0.1) {
                                                                                        if false {
                                                                                            var v80 int8 = (v75 >> ((v73) & 63))
                                                                                            _ = v80
                                                                                        } else {
                                                                                            var v81 uint32 = 0
                                                                                            v11 = (v73 / ((255) | 1))
                                                                                            _ = v81
                                                                                        }
                                                                                        fmt.Println(v73, v4)
                                                                                        if (3.14159 > float64(((((-2.25 * ((0.0 * 0.1) / 1.0)) * (0.1 - 1e10)) + 1.5) - 0.1))) {
                                                                                            v2 = append(v2, ("" + ""))
                                                                                        }
                                                                                    }
                                                                                    v75 = v8
                                                                                }
                                                                            }
                                                                            var v82 uint16 = (v9 ^ v9)
                                                                            _ = v76
                                                                            _ = v82
                                                                        }
                                                                    }
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                _ = v75
                                                            }
                                                            var v83 uint16 = (v9 & 65535)
                                                            _ = v73
                                                            _ = v83
                                                        }
                                                        var v84 int16 = (v4 % ((32767) | 1))
                                                        _ = v84
                                                    }
                                                    _ = v40
                                                }
                                                if (fi13 == (fi13 << (((v11 | (v5 + (v5 + 4294967295)))) & 63))) {
                                                    var v85 string = "\n\t"
                                                    var v86 int64 = -485
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    if false {
                                                        v87 := GenS10{F0: 1}
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            switch (^(fi13)) {
                                                            case -653:
                                                                var v88 float32 = 0.1
                                                                var v89 string = "foo bar"
                                                                _ = v88
                                                                _ = v89
                                                            case 32767:
                                                                v1 += (^(v1))
                                                            default:
                                                                if (v5 != (^(v5))) {
                                                                    if (3.14159 <= 0.1) {
                                                                        var v90 int32 = (^(v1))
                                                                        _ = v90
                                                                    }
                                                                    if (v4 == v4) {
                                                                        v91 := GenS4{F0: 4294967295, F1: (v11 * v5)}
                                                                        switch fi13 {
                                                                        case -261:
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                v4 = (v4 + v4)
                                                                            case "foo bar":
                                                                                v5 = v5
                                                                            }
                                                                        case 42:
                                                                            switch (-(fi18)) {
                                                                            case 9223372036854775807:
                                                                                var v92 int16 = v4
                                                                                var v93 uint32 = uint32(v9)
                                                                                _ = v92
                                                                                _ = v93
                                                                            case 127:
                                                                                if (int8(v1) != v8) {
                                                                                    fmt.Println(v11)
                                                                                } else {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = ((("foo bar" + ("\n\t" + ("foo bar" + ("" + "\n\t")))) + ("foo bar" + "foo bar")) + ("\n\t" + "hello"))
                                                                                    }
                                                                                    v94 := GenS11{F0: 0, F1: v1}
                                                                                    v2 = append(v2, "")
                                                                                    _ = v94
                                                                                }
                                                                            case -422:
                                                                                _ = len(v2)
                                                                            }
                                                                        default:
                                                                            if ((v1 - int32(fi18)) < (v1 | (v1 * 0))) {
                                                                                var v95 int = 42
                                                                                v1 = (v1 / ((v1) | 1))
                                                                                var v96 float32 = (((1e10 / -2.25) / (0.1 + (0.1 - float32((float64(0.0) * (0.0 - 1e10)))))) + 3.14159)
                                                                                v96 = v96
                                                                                _ = v95
                                                                                _ = v96
                                                                            } else {
                                                                                _ = len(v2)
                                                                                fmt.Println(fi18, v9)
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                        var v97 int = fi18
                                                                        _ = v91
                                                                        _ = v97
                                                                    } else {
                                                                        var v98 float32 = float32(-2.25)
                                                                        _ = v98
                                                                    }
                                                                    v99 := GenS5{F0: 1, F1: -9223372036854775808}
                                                                    _ = v99
                                                                } else {
                                                                    var v100 string = ("\n\t" + "")
                                                                    fmt.Println(fi13, v8)
                                                                    _ = v100
                                                                }
                                                            }
                                                        case "a":
                                                            v11 = (v11 << (((v11 << ((4294967295) & 63))) & 63))
                                                        case "":
                                                            var v101 float32 = 0.1
                                                            _ = v101
                                                        }
                                                        _ = v87
                                                    } else {
                                                        v102 := GenS2{F0: 80, F1: v1, F2: v4}
                                                        var v103 uint8 = 219
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    v104 := GenS12{F0: v8, F1: 255, F2: 1}
                                                    _ = v104
                                                }
                                            }
                                            fmt.Println(v8)
                                            fmt.Println(fi18, v1, v9)
                                            v105 := GenS9{F0: (v9 >> ((v5) & 63)), F1: 42}
                                            if v3 {
                                                switch "foo bar" {
                                                case "hello":
                                                    v106 := GenS13{F0: 2, F1: v8}
                                                    _ = v106
                                                case "\n\t":
                                                    v9 -= (v9 & (^(v9)))
                                                case "a":
                                                    v11 = (v5 % ((v11) | 1))
                                                }
                                            }
                                            if v3 {
                                                v1 = v1
                                                v2 = append(v2, "foo bar")
                                                v107 := GenS14{F0: -128, F1: 1}
                                                _ = v107
                                            }
                                            _ = v105
                                        } else {
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = "foo bar"
                                                }
                                                if false {
                                                    v2 = append(v2, "hello")
                                                }
                                            } else {
                                                v108 := GenS10{F0: 63}
                                                _ = len(v2)
                                                _ = v108
                                            }
                                            v9 += (^(v9))
                                            switch (("a" + "a") + (("" + (("foo bar" + "foo bar") + ("a" + "foo bar"))) + "hello")) {
                                            case "foo bar":
                                                var v109 float64 = float64((0.0 / 0.1))
                                                _ = v109
                                            default:
                                                if v3 {
                                                    v110 := GenS15{F0: 18446744073709551615}
                                                    _ = v110
                                                }
                                            }
                                            var v111 uint64 = 4294967295
                                            v2 = append(v2, ("\n\t" + "foo bar"))
                                            _ = v111
                                        }
                                    }
                                    fmt.Println(v11, v5, v1)
                                }
                            case 9223372036854775807:
                                switch "foo bar" {
                                case "hello":
                                    v112 := GenS16{F0: int(v4), F1: 100, F2: v9}
                                    v9 = (^(v9))
                                    _ = v112
                                }
                            }
                            _ = fi18
                        }
                    }
                }
                if (1e10 <= (0.1 + 1e10)) {
                    for fi113 := 0; fi113 < 17; fi113++ {
                        v2 = append(v2, ((("hello" + ("foo bar" + "a")) + "hello") + ""))
                        _ = fi113
                    }
                    v114 := GenS17{F0: 255, F1: -279}
                    _ = v114
                }
                _ = fi13
            }
        case 767:
            for fi115 := 0; fi115 < 7; fi115++ {
                v116 := GenS18{F0: 1, F1: 4294967295}
                if (18446744073709551615 < uint64(v5)) {
                    var v117 int32 = 42
                    v8 += 100
                    switch fi115 {
                    case 32767:
                        v4 += int16(v9)
                    case 127:
                        var v118 int16 = (v4 ^ 1000)
                        _ = v118
                    case -9223372036854775808:
                        if ((true || true) && v3) {
                            for fi119 := 0; fi119 < 15; fi119++ {
                                var v120 int = (fi115 >> ((v5) & 63))
                                _ = fi119
                                _ = v120
                            }
                        } else {
                            switch (fi115 & (fi115 >> (((v5 - v11)) & 63))) {
                            case 128:
                                if v3 {
                                    fmt.Println(v1, v11)
                                }
                            }
                            var v121 uint16 = (v9 << (((v5 | uint(v9))) & 63))
                            _ = v121
                        }
                    default:
                        v122 := GenS15{F0: 2}
                        _ = v122
                    }
                    _ = v117
                } else {
                    v1 += v1
                }
                _ = fi115
                _ = v116
            }
        case -9223372036854775808:
            var v123 int = 127
            _ = v123
        }
    } else {
        for fi124 := 0; fi124 < 1; fi124++ {
            var v125 string = "hello"
            fmt.Println(v5, v1, v11)
            fmt.Println(v9, v3)
            _ = fi124
            _ = v125
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
