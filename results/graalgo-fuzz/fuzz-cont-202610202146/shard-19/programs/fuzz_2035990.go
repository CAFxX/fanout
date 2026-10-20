package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 int; F2 uint16 }
type GenS11 struct { F0 uint16; F1 bool }
type GenS12 struct { F0 int8; F1 int; F2 int16 }
type GenS13 struct { F0 bool; F1 uint16; F2 bool }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int16; F1 uint64; F2 bool }
type GenS18 struct { F0 int; F1 int8 }

func main() {
    v1 := []int{128}
    var v2 uint8 = 0
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 32767
    var v5 string = "a"
    if true {
        var v6 uint = 692
        switch -1 {
        case -1:
            var v7 uint64 = 860
            _ = v7
        case 42:
            if ((!(282 < -2147483648)) || false) {
                var v8 int64 = v3
                v9 := GenS1{F0: true, F1: "foo bar"}
                _ = v8
                _ = v9
            } else {
                if (-2.25 >= -2.25) {
                    v4 = v4
                    v2 -= 0
                } else {
                    var v10 float32 = (1.5 + (-2.25 * 1e10))
                    v10 -= ((1.5 + v10) * v10)
                    v2 += v2
                    _ = v10
                }
                fmt.Println(v4, v2, v5)
                if false {
                    _ = len(v1)
                    v11 := GenS2{F0: 32767, F1: 300, F2: 42}
                    for fi12 := 0; fi12 < 9; fi12++ {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 int32 = -1
                            v15 := GenS3{F0: 551, F1: 630, F2: (v2 - (v2 % (((v2 >> ((4294967295) & 63))) | 1)))}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        v5 = v5
                        _ = fi12
                    }
                    _ = v11
                } else {
                    switch 2147483647 {
                    case -516:
                        if ((!((v3 / (((v3 / (((v3 >> ((v6) & 63))) | 1))) | 1)) >= v3)) && (100 > 100)) {
                            v16 := GenS4{F0: 249}
                            for fi17 := 0; fi17 < 15; fi17++ {
                                v4 = (v4 + v4)
                                v18 := GenS3{F0: 42, F1: 1, F2: v2}
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    switch (v5 + (v5 + (v5 + v5))) {
                                    case "a":
                                        switch fi17 {
                                        case 32768:
                                            var v20 int16 = 526
                                            _ = v20
                                        case 4294967295:
                                            if true {
                                                v4 = (v4 & v4)
                                            }
                                        default:
                                            v21 := GenS5{F0: 1}
                                            var v22 int64 = v3
                                            if false {
                                                var v23 uint64 = 0
                                                v2 -= 203
                                                _ = v23
                                            } else {
                                                v4 = 252
                                                fmt.Println(v3, v3)
                                                var v24 uint16 = uint16(v2)
                                                _ = v24
                                            }
                                            v25 := GenS6{F0: (v6 & (^(v6)))}
                                            _ = v21
                                            _ = v22
                                            _ = v25
                                        }
                                    case "hello":
                                        v6 += (v6 >> (((v6 + 2)) & 63))
                                    case "foo bar":
                                        v26 := GenS4{F0: 2147483647}
                                        _ = v26
                                    }
                                    switch (fi17 - 774) {
                                    case 32768:
                                        if false {
                                            v27 := GenS7{F0: false}
                                            fmt.Println(fi17)
                                            _ = v27
                                        }
                                    case -1:
                                        switch (("" + "") + "\n\t") {
                                        case "":
                                            if ((v6 >> (((v6 - v6)) & 63)) >= v6) {
                                                var v28 uint8 = v2
                                                if true {
                                                    fmt.Println(v28)
                                                    switch fi17 {
                                                    case -430:
                                                        switch v5 {
                                                        case "a":
                                                            switch v5 {
                                                            case "":
                                                                v2 = 245
                                                            case "\n\t":
                                                                if false {
                                                                    if (!true) {
                                                                        var v29 uint64 = 691
                                                                        if ((true && true) && ((!false) && false)) {
                                                                            switch v5 {
                                                                            case "hello":
                                                                                if true {
                                                                                    v30 := GenS8{F0: true}
                                                                                    v5 = (v5 + (((v5 + v5) + ("" + v5)) + "a"))
                                                                                    if true {
                                                                                        var v31 int = fi19
                                                                                        var v32 uint8 = (v2 | (v2 % (((v2 % (((v2 % (((v28 ^ 1)) | 1))) | 1))) | 1)))
                                                                                        v33 := GenS9{F0: "\n\t"}
                                                                                        _ = v31
                                                                                        _ = v32
                                                                                        _ = v33
                                                                                    } else {
                                                                                        var v34 int32 = 1
                                                                                        v1 = append(v1, fi17)
                                                                                        _ = v34
                                                                                    }
                                                                                    v28 -= uint8(v29)
                                                                                    v4 = 42
                                                                                    _ = v30
                                                                                } else {
                                                                                    var v35 uint64 = v29
                                                                                    _ = v35
                                                                                }
                                                                            default:
                                                                                v3 += v3
                                                                            }
                                                                            var v36 int = fi19
                                                                            if (true || (1 > 42)) {
                                                                                v6 += (^(v6))
                                                                                _ = len(v1)
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                        _ = v29
                                                                    } else {
                                                                        v37 := GenS8{F0: false}
                                                                        v38 := GenS10{F0: 1, F1: fi17, F2: (^(v4))}
                                                                        _ = v37
                                                                        _ = v38
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                } else {
                                                                    var v39 float32 = float32((-2.25 * -2.25))
                                                                    v5 = (((v5 + (v5 + v5)) + v5) + "")
                                                                    _ = v39
                                                                }
                                                            default:
                                                                v2 = (v2 >> ((347) & 63))
                                                            }
                                                        case "":
                                                            v5 = "foo bar"
                                                        case "foo bar":
                                                            fmt.Println(v2)
                                                        }
                                                        var v40 bool = false
                                                        v4 = v4
                                                        _ = v40
                                                    case 4294967295:
                                                        v41 := GenS11{F0: (v4 / (((v4 >> ((v6) & 63))) | 1)), F1: true}
                                                        var v42 uint16 = (v4 & v4)
                                                        _ = v41
                                                        _ = v42
                                                    case -618:
                                                        switch (("" + "\n\t") + v5) {
                                                        case "":
                                                            v5 = v5
                                                        case "hello":
                                                            v1 = append(v1, fi19)
                                                        }
                                                    default:
                                                        var v43 int16 = -696
                                                        _ = v43
                                                    }
                                                    fmt.Println(fi17, v4)
                                                } else {
                                                    v2 = (^(v28))
                                                    var v44 int8 = -128
                                                    _ = v44
                                                }
                                                _ = v28
                                            } else {
                                                switch v5 {
                                                case "foo bar":
                                                    var v45 uint32 = 42
                                                    _ = v45
                                                case "\n\t":
                                                    v46 := GenS8{F0: false}
                                                    _ = v46
                                                case "a":
                                                    if ((^(v6)) < (^(v6))) {
                                                        var v47 bool = false
                                                        _ = v47
                                                    }
                                                default:
                                                    v5 = v5
                                                }
                                                var v48 uint32 = 42
                                                v3 = 42
                                                v49 := GenS8{F0: (fi19 > 0)}
                                                if false {
                                                    if true {
                                                        if len(v1) > 0 {
                                                            v1[0] = fi19
                                                        }
                                                        v48 = v48
                                                    }
                                                    var v50 float64 = 3.14159
                                                    var v51 int8 = 42
                                                    _ = v50
                                                    _ = v51
                                                } else {
                                                    if (true && true) {
                                                        if len(v1) > 0 {
                                                            v1[0] = 127
                                                        }
                                                        var v52 uint16 = v4
                                                        _ = v52
                                                    }
                                                    if false {
                                                        var v53 string = ("foo bar" + (v5 + v5))
                                                        if len(v1) > 0 {
                                                            v1[0] = int(v2)
                                                        }
                                                        if ((v3 << ((v6) & 63)) != (v3 % ((-9223372036854775808) | 1))) {
                                                            v54 := GenS12{F0: 0, F1: fi19, F2: 286}
                                                            _ = v54
                                                        } else {
                                                            var v55 float32 = (((float32(1e10) / 3.14159) * 1.5) / float32((0.1 + 1.5)))
                                                            _ = len(v1)
                                                            if ((v2 + 127) > (v2 * uint8(v4))) {
                                                                var v56 float64 = float64(float32(-2.25))
                                                                _ = v56
                                                            }
                                                            _ = v55
                                                        }
                                                        var v57 string = v53
                                                        _ = v53
                                                        _ = v57
                                                    }
                                                    if (-534 != v3) {
                                                        var v58 bool = (v4 > v4)
                                                        switch (-(fi19)) {
                                                        case 127:
                                                            v3 = (-(v3))
                                                            v59 := GenS13{F0: (v58 || ((v4 - (v4 & (v4 / ((v4) | 1)))) != v4)), F1: (v4 << ((v6) & 63)), F2: ((^(v2)) != (v2 << (((v6 << (((^(v6))) & 63))) & 63)))}
                                                            _ = v59
                                                        case 2147483647:
                                                            var v60 int = (fi19 ^ (fi17 + fi19))
                                                            _ = v60
                                                        case -194:
                                                            v3 = 0
                                                            fmt.Println(fi17)
                                                            if v58 {
                                                                var v61 float32 = -2.25
                                                                _ = v61
                                                            }
                                                        default:
                                                            var v62 int8 = -128
                                                            _ = v62
                                                        }
                                                        _ = v58
                                                    }
                                                    v6 = v6
                                                    if (false || (!(true || true))) {
                                                        var v63 uint32 = (v48 >> ((v6) & 63))
                                                        var v64 int8 = 1
                                                        v2 -= v2
                                                        switch v5 {
                                                        case "\n\t":
                                                            v64 = (v64 * (v64 + int8(v48)))
                                                        case "foo bar":
                                                            if (v6 > (v6 / ((v6) | 1))) {
                                                                switch "hello" {
                                                                case "hello":
                                                                    if len(v1) > 0 {
                                                                        v1[0] = fi19
                                                                    }
                                                                case "a":
                                                                    _ = len(v1)
                                                                }
                                                                v4 = (v4 & (^(v4)))
                                                                v65 := GenS5{F0: uint64(v6)}
                                                                _ = v65
                                                            } else {
                                                                if true {
                                                                    var v66 int8 = (v64 | v64)
                                                                    fmt.Println(v48, v6)
                                                                    var v67 uint8 = uint8(v4)
                                                                    v68 := GenS3{F0: 4294967295, F1: 225, F2: v67}
                                                                    if (false && ((fi17 << ((0) & 63)) > (fi17 >> (((v6 | (v6 * v6))) & 63)))) {
                                                                        _ = len(v1)
                                                                        v69 := GenS14{F0: ((216 <= 18446744073709551615) && (0.0 < (0.1 + ((0.1 + 1.5) * -2.25))))}
                                                                        v63 += v48
                                                                        _ = v69
                                                                    } else {
                                                                        if (true || (true && true)) {
                                                                            v2 -= v2
                                                                        } else {
                                                                            v70 := GenS13{F0: true, F1: v4, F2: ((v4 >> ((v6) & 63)) <= v4)}
                                                                            if ((4294967295 < 2) || ((v2 != v2) && true)) {
                                                                                _ = len(v1)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v67, v2)
                                                                                v71 := GenS15{F0: -32768, F1: v5}
                                                                                v63 -= v48
                                                                                _ = v71
                                                                            }
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                    _ = v66
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    if true {
                                                                        if ((v6 * (^(v6))) != (v6 << ((v6) & 63))) {
                                                                            var v72 string = v5
                                                                            _ = v72
                                                                        } else {
                                                                            switch v5 {
                                                                            case "":
                                                                                fmt.Println(v4, v3, v48)
                                                                            case "foo bar":
                                                                                v48 += (v48 * 42)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                var v73 int32 = int32(v3)
                                                                                _ = v73
                                                                            case "hello":
                                                                                fmt.Println(v6, v5)
                                                                                _ = len(v1)
                                                                            default:
                                                                                var v74 uint32 = (v48 >> ((v6) & 63))
                                                                                _ = v74
                                                                            }
                                                                            fmt.Println(v48, v48)
                                                                            var v75 int64 = v3
                                                                            switch (fi17 - int(v6)) {
                                                                            case 9223372036854775807:
                                                                                _ = len(v1)
                                                                            }
                                                                            var v76 float32 = 1.5
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                        var v77 float32 = 1e10
                                                                        v78 := GenS11{F0: v4, F1: false}
                                                                        _ = v77
                                                                        _ = v78
                                                                    } else {
                                                                        v63 += v63
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        default:
                                                            v64 -= 127
                                                        }
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    fmt.Println(fi19, fi17)
                                                    var v79 bool = false
                                                    _ = v79
                                                }
                                                if true {
                                                    _ = len(v1)
                                                    var v80 int8 = -128
                                                    v48 = v48
                                                    _ = v80
                                                } else {
                                                    if (127 != -80) {
                                                        var v81 uint = v6
                                                        _ = v81
                                                    }
                                                }
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    case 9223372036854775807:
                                        if false {
                                            if true {
                                                if false {
                                                    switch (^(fi19)) {
                                                    case 331:
                                                        fmt.Println(fi17)
                                                    case -315:
                                                        fmt.Println(v6)
                                                    case -363:
                                                        _ = len(v1)
                                                    }
                                                    v2 += 14
                                                    switch v5 {
                                                    case "":
                                                        v82 := GenS14{F0: (!(v5 != v5))}
                                                        _ = v82
                                                    case "foo bar":
                                                        fmt.Println(v6, v4)
                                                    case "a":
                                                        if (false && false) {
                                                            if len(v1) > 0 {
                                                                v1[0] = fi17
                                                            }
                                                            v83 := GenS16{F0: 42}
                                                            _ = v83
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v84 float32 = 1e10
                                                v85 := GenS16{F0: 1}
                                                _ = v84
                                                _ = v85
                                            }
                                        } else {
                                            v86 := GenS17{F0: 40, F1: 18446744073709551615, F2: true}
                                            var v87 bool = (-1 < 99)
                                            _ = v86
                                            _ = v87
                                        }
                                    default:
                                        v5 = (v5 + "\n\t")
                                    }
                                    var v88 int32 = int32(v6)
                                    v5 = (v5 + v5)
                                    _ = fi19
                                    _ = v88
                                }
                                _ = fi17
                                _ = v18
                            }
                            for fi89 := 0; fi89 < 17; fi89++ {
                                v90 := GenS12{F0: -1, F1: (fi89 << ((v6) & 63)), F2: 0}
                                fmt.Println(v2)
                                v1 = append(v1, (-(fi89)))
                                _ = fi89
                                _ = v90
                            }
                            _ = v16
                        } else {
                            fmt.Println(v5)
                        }
                    case 32767:
                        var v91 int32 = -609
                        _ = v91
                    }
                    var v92 int8 = 42
                    switch v5 {
                    case "foo bar":
                        v93 := GenS11{F0: v4, F1: (!false)}
                        _ = v93
                    case "\n\t":
                        var v94 uint16 = (^(v4))
                        _ = v94
                    case "":
                        v95 := GenS1{F0: (4294967295 < 773), F1: v5}
                        _ = v95
                    }
                    _ = v92
                }
            }
        default:
            var v96 float32 = 1e10
            _ = v96
        }
        _ = v6
    }
    v2 += (^(v2))
    v1 = append(v1, 435)
    v97 := GenS18{F0: 127, F1: -128}
    switch (v5 + v5) {
    case "foo bar":
        v4 = 785
    }
    var v98 uint64 = 337
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = 65535
    if (!true) {
        _ = len(v1)
        v5 = "foo bar"
    }
    var v99 int = 4294967295
    var v100 int32 = -1
    var v101 int = int(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
    _ = v99
}
