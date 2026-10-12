package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    var v1 int32 = -2147483648
    var v2 string = ("" + "\n\t")
    var v3 uint32 = 1
    v4 := []uint32{(^(v3)), (^(v3))}
    var v5 uint = 255
    var v6 uint32 = 782
    if ("hello" < "foo bar") {
        if true {
            fmt.Println(v5)
            _ = len(v4)
            v1 = v1
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v6, v3, v6)
            _ = fi7
        }
    }
    var v8 float64 = -2.25
    switch -9223372036854775808 {
    case 128:
        var v9 int32 = (v1 ^ (v1 * (v1 | (v1 ^ v1))))
        _ = v9
    case 0:
        if (((v5 * v5) > v5) && ("hello" <= v2)) {
            if (!true) {
                v3 -= (v6 ^ (v3 & (^(v6))))
                for fi10 := 0; fi10 < 18; fi10++ {
                    fmt.Println(v1, v5, v3)
                    if (false || true) {
                        var v11 uint = v5
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v4 = append(v4, 42)
                            v13 := GenS1{F0: 0, F1: "\n\t", F2: fi12}
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    } else {
                        switch ("a" + ("a" + "a")) {
                        case "":
                            if (!false) {
                                v14 := GenS1{F0: 300, F1: v2, F2: fi10}
                                switch fi10 {
                                case 42:
                                    var v15 uint32 = v3
                                    _ = v15
                                }
                                _ = v14
                            }
                        }
                        switch 0 {
                        case 856:
                            var v16 int32 = v1
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 14; fi17++ {
                                if true {
                                    v18 := GenS1{F0: 1, F1: v2, F2: 128}
                                    _ = v18
                                }
                                v4 = append(v4, v3)
                                _ = fi17
                            }
                        }
                    }
                    _ = fi10
                }
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v19 float32 = -2.25
    if len(v4) > 0 {
        v4[0] = v6
    }
    v20 := GenS2{F0: -1, F1: 42}
    v21 := GenS3{F0: 1, F1: 32767}
    v6 = 4294967295
    v22 := GenS3{F0: 18446744073709551615, F1: 1}
    if true {
        switch v2 {
        case "a":
            v4 = append(v4, 42)
        case "":
            switch v2 {
            case "hello":
                v5 = 1
            default:
                v23 := GenS4{F0: 32767, F1: (v3 / (((^(v6))) | 1))}
                _ = v23
            }
        case "foo bar":
            var v24 int64 = -9223372036854775808
            _ = v24
        }
        if (false && false) {
            fmt.Println(v6, v5, v1)
        } else {
            if len(v4) > 0 {
                v4[0] = (^(v3))
            }
            var v25 int8 = 77
            _ = v25
        }
        fmt.Println(v1)
    } else {
        var v26 int32 = (v1 << ((v5) & 63))
        _ = v26
    }
    v27 := GenS4{F0: 42, F1: v6}
    var v28 int64 = 0
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS4{F0: 1, F1: v6}
        _ = fi29
        _ = v30
    }
    switch 2147483647 {
    case 458:
        v31 := GenS1{F0: 255, F1: "foo bar", F2: 355}
        _ = v31
    case 42:
        for fi32 := 0; fi32 < 20; fi32++ {
            if true {
                for fi33 := 0; fi33 < 12; fi33++ {
                    if (-128 == 1) {
                        var v34 int32 = v1
                        if false {
                            var v35 int = fi33
                            _ = v35
                        } else {
                            v2 = (((v2 + v2) + v2) + v2)
                            var v36 string = (("a" + "\n\t") + v2)
                            _ = v36
                        }
                        var v37 bool = ((fi33 - (fi33 & fi32)) != fi33)
                        _ = v34
                        _ = v37
                    } else {
                        if false {
                            fmt.Println(fi33, v8)
                            _ = len(v4)
                            fmt.Println(fi32, fi33, v5)
                        } else {
                            v38 := GenS1{F0: 0, F1: v2, F2: (fi33 | -9223372036854775808)}
                            if (1 <= 350) {
                                if (v28 < 42) {
                                    switch fi33 {
                                    case 1:
                                        if false {
                                            if false {
                                                v8 = float64((v19 * 1e10))
                                            }
                                            fmt.Println(v8, v1, v2)
                                            v5 += 65535
                                        }
                                    case 2147483647:
                                        v39 := GenS4{F0: -32768, F1: v3}
                                        _ = v39
                                    }
                                    var v40 int32 = (v1 ^ v1)
                                    _ = v40
                                }
                                var v41 float32 = v19
                                switch -1 {
                                case -667:
                                    var v42 string = ""
                                    _ = v42
                                }
                                _ = v41
                            } else {
                                switch fi32 {
                                case 4294967295:
                                    switch fi33 {
                                    case -2147483648:
                                        var v43 string = "hello"
                                        var v44 uint64 = 1
                                        if true {
                                            var v45 uint64 = uint64(v6)
                                            var v46 uint8 = 1
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v43 = "hello"
                                            switch "foo bar" {
                                            case "foo bar":
                                                v47 := GenS5{F0: 1, F1: 255, F2: (v28 & (v28 / ((int64(v1)) | 1)))}
                                                _ = v47
                                            case "hello":
                                                switch (fi33 / ((fi33) | 1)) {
                                                case 128:
                                                    var v48 float32 = v19
                                                    _ = v48
                                                case 2147483647:
                                                    var v49 float64 = float64(v19)
                                                    _ = v49
                                                case 32768:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            case "":
                                                v4 = append(v4, (^(v6)))
                                                switch fi33 {
                                                case 32767:
                                                    if false {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (v5 >= (v5 * v5)) {
                                                            v50 := GenS6{F0: v6, F1: v3}
                                                            if false {
                                                                var v51 int = (^(fi32))
                                                                _ = v51
                                                            } else {
                                                                v52 := GenS3{F0: (^(v44)), F1: 65535}
                                                                switch (fi32 - (fi32 & fi33)) {
                                                                case 9223372036854775807:
                                                                    v3 -= uint32(v5)
                                                                default:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        switch v2 {
                                                                        case "\n\t":
                                                                            switch 500 {
                                                                            case 1:
                                                                                v8 = float64((v19 * v19))
                                                                            case 4294967295:
                                                                                var v53 int = (fi32 << (((v5 + 18446744073709551615)) & 63))
                                                                                _ = v53
                                                                            case -284:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            default:
                                                                                if (true || (false && false)) {
                                                                                    _ = len(v4)
                                                                                    if true {
                                                                                        switch (fi32 & (fi33 + 0)) {
                                                                                        case -818:
                                                                                            v44 -= v44
                                                                                        case 310:
                                                                                            v6 = (v3 & v3)
                                                                                        case 42:
                                                                                            if (2 <= 1) {
                                                                                                var v54 uint8 = 1
                                                                                                if (true || false) {
                                                                                                    if false {
                                                                                                        v19 = ((-2.25 - v19) / v19)
                                                                                                        var v55 int32 = v1
                                                                                                        _ = v55
                                                                                                    } else {
                                                                                                        v43 = v43
                                                                                                        if true {
                                                                                                            if (!false) {
                                                                                                                if true {
                                                                                                                    var v56 float64 = float64(v19)
                                                                                                                    v44 = v44
                                                                                                                    var v57 float32 = ((((v19 / 1.5) * ((v19 - v19) / float32((3.14159 - 1.5)))) + (1.5 / 0.1)) / (v19 - ((v19 * ((v19 / -2.25) / v19)) - v19)))
                                                                                                                    v3 = v6
                                                                                                                    _ = v56
                                                                                                                    _ = v57
                                                                                                                }
                                                                                                                fmt.Println(v8)
                                                                                                                fmt.Println(v8, v3, v19)
                                                                                                                if (!(32767 <= (fi33 & (fi32 ^ int(v1))))) {
                                                                                                                    v43 = v43
                                                                                                                    v58 := GenS7{F0: (v28 ^ v28), F1: v43}
                                                                                                                    switch "foo bar" {
                                                                                                                    case "hello":
                                                                                                                        fmt.Println(v1)
                                                                                                                    case "a":
                                                                                                                        v59 := GenS5{F0: -32768, F1: 204, F2: 665}
                                                                                                                        _ = v59
                                                                                                                    case "foo bar":
                                                                                                                        if false {
                                                                                                                            v4 = append(v4, 2147483648)
                                                                                                                            fmt.Println(v6)
                                                                                                                            v6 = 281
                                                                                                                        } else {
                                                                                                                            v4 = append(v4, (v3 - uint32(v54)))
                                                                                                                            if len(v4) > 0 {
                                                                                                                                v4[0] = (^(v3))
                                                                                                                            }
                                                                                                                            switch v43 {
                                                                                                                            case "a":
                                                                                                                                v28 -= v28
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                                fmt.Println(v5, v19)
                                                                                                            } else {
                                                                                                                var v60 float64 = float64((float32(((v8 * ((-2.25 / v8) * (((v8 / (v8 - v8)) * (((v8 / (v8 + v8)) / v8) / float64(v19))) * v8))) * (1.5 * 0.0))) * (3.14159 - 0.0)))
                                                                                                                switch ("" + v43) {
                                                                                                                case "hello":
                                                                                                                    if true {
                                                                                                                        _ = len(v4)
                                                                                                                        var v61 uint = v5
                                                                                                                        if false {
                                                                                                                            if false {
                                                                                                                                if false {
                                                                                                                                    v28 = v28
                                                                                                                                    if (!true) {
                                                                                                                                        switch "hello" {
                                                                                                                                        case "a":
                                                                                                                                            v62 := GenS8{F0: 1}
                                                                                                                                            _ = v62
                                                                                                                                        case "":
                                                                                                                                            if len(v4) > 0 {
                                                                                                                                                v4[0] = v6
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        if true {
                                                                                                                                            if len(v4) > 0 {
                                                                                                                                                v4[0] = (v3 * (v6 | v3))
                                                                                                                                            }
                                                                                                                                            var v63 int8 = 127
                                                                                                                                            v19 -= ((1.5 * v19) + v19)
                                                                                                                                            _ = v63
                                                                                                                                        }
                                                                                                                                        if len(v4) > 0 {
                                                                                                                                            v4[0] = uint32(v61)
                                                                                                                                        }
                                                                                                                                        var v64 bool = true
                                                                                                                                        _ = v64
                                                                                                                                    }
                                                                                                                                    v19 = 0.0
                                                                                                                                }
                                                                                                                                v65 := GenS9{F0: (^(v44)), F1: ((^(v1)) != 777)}
                                                                                                                                v19 = v19
                                                                                                                                v4 = append(v4, (v6 | 807))
                                                                                                                                v66 := GenS1{F0: 32767, F1: v43, F2: fi33}
                                                                                                                                if false {
                                                                                                                                    if len(v4) > 0 {
                                                                                                                                        v4[0] = (v6 / ((v6) | 1))
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if ((-710 <= 688) || false) {
                                                                                                                                    if true {
                                                                                                                                        var v67 int64 = (-(v28))
                                                                                                                                        v68 := GenS3{F0: v44, F1: 1}
                                                                                                                                        if false {
                                                                                                                                            v69 := GenS10{F0: 32767, F1: v44, F2: -128}
                                                                                                                                            switch (fi32 + (fi33 / ((fi32) | 1))) {
                                                                                                                                            case 113:
                                                                                                                                                if true {
                                                                                                                                                    if (v8 != v60) {
                                                                                                                                                        var v70 int8 = -1
                                                                                                                                                        v1 += 1
                                                                                                                                                        v71 := GenS7{F0: v67, F1: "\n\t"}
                                                                                                                                                        if (int64(fi33) < (v28 << (((v5 * v61)) & 63))) {
                                                                                                                                                            switch v43 {
                                                                                                                                                            case "foo bar":
                                                                                                                                                                v72 := GenS5{F0: -32768, F1: 0, F2: (v67 | v67)}
                                                                                                                                                                _ = v72
                                                                                                                                                            case "a":
                                                                                                                                                                var v73 int32 = (v1 + v1)
                                                                                                                                                                _ = v73
                                                                                                                                                            }
                                                                                                                                                            v74 := GenS11{F0: false, F1: 966}
                                                                                                                                                            switch (fi32 | 32767) {
                                                                                                                                                            case -1:
                                                                                                                                                                v44 = 330
                                                                                                                                                            default:
                                                                                                                                                                v4 = append(v4, v3)
                                                                                                                                                            }
                                                                                                                                                            _ = v74
                                                                                                                                                        }
                                                                                                                                                        _ = v70
                                                                                                                                                        _ = v71
                                                                                                                                                    } else {
                                                                                                                                                        v6 = v6
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    v4 = append(v4, v3)
                                                                                                                                                }
                                                                                                                                            case -9223372036854775808:
                                                                                                                                                if false {
                                                                                                                                                    var v75 int = 42
                                                                                                                                                    _ = v75
                                                                                                                                                } else {
                                                                                                                                                    var v76 uint16 = 255
                                                                                                                                                    _ = v76
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                v77 := GenS6{F0: v3, F1: uint32(v54)}
                                                                                                                                                _ = v77
                                                                                                                                            }
                                                                                                                                            if false {
                                                                                                                                                if true {
                                                                                                                                                    var v78 int8 = 127
                                                                                                                                                    _ = v78
                                                                                                                                                }
                                                                                                                                                fmt.Println(v2)
                                                                                                                                            } else {
                                                                                                                                                var v79 bool = false
                                                                                                                                                _ = v79
                                                                                                                                            }
                                                                                                                                            _ = v69
                                                                                                                                        }
                                                                                                                                        _ = v67
                                                                                                                                        _ = v68
                                                                                                                                    }
                                                                                                                                    v80 := GenS7{F0: v28, F1: v2}
                                                                                                                                    _ = v80
                                                                                                                                } else {
                                                                                                                                    v61 = v61
                                                                                                                                }
                                                                                                                                v1 += (v1 / (((v1 - (v1 + (v1 * int32(v5))))) | 1))
                                                                                                                                _ = v65
                                                                                                                                _ = v66
                                                                                                                            } else {
                                                                                                                                v60 -= v8
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v61
                                                                                                                    } else {
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = 2147483647
                                                                                                                        }
                                                                                                                        if (65535 <= 42) {
                                                                                                                            v43 = (("hello" + v2) + "hello")
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v81 uint32 = (^(v3))
                                                                                                                    _ = v81
                                                                                                                case "":
                                                                                                                    var v82 int32 = (v1 * (v1 + (v1 << ((42) & 63))))
                                                                                                                    _ = v82
                                                                                                                default:
                                                                                                                    var v83 int = (fi33 % (((fi32 % ((fi33) | 1))) | 1))
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                _ = v60
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v84 uint32 = v6
                                                                                                                _ = v84
                                                                                                            }
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = v6
                                                                                                            }
                                                                                                        } else {
                                                                                                            v44 = 0
                                                                                                            if false {
                                                                                                                v44 = (v44 >> ((677) & 63))
                                                                                                                v2 = (v43 + "")
                                                                                                            }
                                                                                                            if false {
                                                                                                                v19 = v19
                                                                                                            } else {
                                                                                                                var v85 float32 = 0.0
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (v3 & (v3 & (^(v6))))
                                                                                                                }
                                                                                                                v86 := GenS12{F0: v2}
                                                                                                                _ = v85
                                                                                                                _ = v86
                                                                                                            }
                                                                                                        }
                                                                                                        if (!true) {
                                                                                                            var v87 uint64 = 0
                                                                                                            v1 -= 0
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            _ = v87
                                                                                                        } else {
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            var v88 string = v2
                                                                                                            _ = v88
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = v3
                                                                                                    }
                                                                                                    var v89 int32 = int32(v28)
                                                                                                    v54 = v54
                                                                                                    if false {
                                                                                                        v90 := GenS13{F0: (v6 << ((18446744073709551615) & 63)), F1: (v2 + "")}
                                                                                                        fmt.Println(v43, v89, v44)
                                                                                                        _ = v90
                                                                                                    } else {
                                                                                                        fmt.Println(v2, v6, v89)
                                                                                                    }
                                                                                                    if (v5 == (v5 << ((v5) & 63))) {
                                                                                                        v91 := GenS7{F0: (v28 >> (((v5 / ((71) | 1))) & 63)), F1: v2}
                                                                                                        if false {
                                                                                                            if (-74 > -86) {
                                                                                                                if false {
                                                                                                                    var v92 float64 = 1e10
                                                                                                                    v6 += 0
                                                                                                                    _ = v92
                                                                                                                } else {
                                                                                                                    var v93 uint8 = (^(v54))
                                                                                                                    v94 := GenS14{F0: v54}
                                                                                                                    fmt.Println(fi32, fi32, v2)
                                                                                                                    _ = v93
                                                                                                                    _ = v94
                                                                                                                }
                                                                                                                var v95 bool = ((-(v28)) < (v28 | v28))
                                                                                                                _ = v95
                                                                                                            } else {
                                                                                                                var v96 string = "\n\t"
                                                                                                                v97 := GenS11{F0: true, F1: 1}
                                                                                                                switch (fi32 ^ (^(fi33))) {
                                                                                                                case 1:
                                                                                                                    v96 = "foo bar"
                                                                                                                case 32767:
                                                                                                                    if len(v4) > 0 {
                                                                                                                        v4[0] = v6
                                                                                                                    }
                                                                                                                case 4294967295:
                                                                                                                    v1 -= v1
                                                                                                                    var v98 int64 = v28
                                                                                                                    _ = v98
                                                                                                                default:
                                                                                                                    var v99 int16 = int16(v5)
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                _ = v96
                                                                                                                _ = v97
                                                                                                            }
                                                                                                            var v100 string = (v43 + "a")
                                                                                                            if true {
                                                                                                                if len(v4) >= 2 {
                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                }
                                                                                                                var v101 uint64 = v44
                                                                                                                v102 := GenS15{F0: 247, F1: 2, F2: 2}
                                                                                                                _ = v101
                                                                                                                _ = v102
                                                                                                            } else {
                                                                                                                fmt.Println(v89, v5)
                                                                                                                v44 -= 42
                                                                                                                if (v19 > ((float32(v8) / (-2.25 * v19)) / 1e10)) {
                                                                                                                    v103 := GenS2{F0: 32767, F1: v44}
                                                                                                                    v43 = (v43 + v2)
                                                                                                                    _ = v103
                                                                                                                } else {
                                                                                                                    var v104 int32 = v1
                                                                                                                    v105 := GenS10{F0: 2, F1: 794, F2: 127}
                                                                                                                    _ = v104
                                                                                                                    _ = v105
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v100
                                                                                                        }
                                                                                                        if (65535 <= v5) {
                                                                                                            switch (fi32 | int(v6)) {
                                                                                                            case -1:
                                                                                                                var v106 int8 = -128
                                                                                                                _ = v106
                                                                                                            case 128:
                                                                                                                var v107 uint8 = v54
                                                                                                                _ = v107
                                                                                                            default:
                                                                                                                v43 = v43
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch fi33 {
                                                                                                            case -923:
                                                                                                                v54 -= v54
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    if (!((v3 * 386) <= v3)) {
                                                                                                                        var v108 bool = (fi32 == (fi32 | (fi32 | fi33)))
                                                                                                                        v109 := GenS5{F0: -667, F1: 0, F2: (v28 & int64(v54))}
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                        _ = v108
                                                                                                                        _ = v109
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            var v110 int32 = int32(fi33)
                                                                                                                            _ = v110
                                                                                                                        }
                                                                                                                        var v111 uint64 = v44
                                                                                                                        if true {
                                                                                                                            fmt.Println(v8, fi32, v44)
                                                                                                                            var v112 uint32 = (^(v3))
                                                                                                                            var v113 int16 = -32768
                                                                                                                            _ = v112
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                        _ = v111
                                                                                                                    }
                                                                                                                    var v114 int8 = 42
                                                                                                                    v8 = float64(float32(1.5))
                                                                                                                    _ = v114
                                                                                                                } else {
                                                                                                                    fmt.Println(v5)
                                                                                                                }
                                                                                                            }
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = 2
                                                                                                            }
                                                                                                            var v115 int16 = -32768
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        _ = v91
                                                                                                    }
                                                                                                    _ = v89
                                                                                                }
                                                                                                _ = v54
                                                                                            } else {
                                                                                                var v116 float32 = v19
                                                                                                _ = v116
                                                                                            }
                                                                                        }
                                                                                        var v117 int8 = int8(v3)
                                                                                        _ = v117
                                                                                    }
                                                                                    if (true && (!((!(true && false)) || (!false)))) {
                                                                                        switch v2 {
                                                                                        case "a":
                                                                                            v28 += v28
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                var v118 string = "a"
                                                                                                if false {
                                                                                                    var v119 int64 = (v28 / ((v28) | 1))
                                                                                                    _ = v119
                                                                                                } else {
                                                                                                    var v120 float32 = v19
                                                                                                    v5 += (^(v5))
                                                                                                    fmt.Println(v28, v118, v19)
                                                                                                    v44 = v44
                                                                                                    v5 = 18446744073709551615
                                                                                                    if (v1 > v1) {
                                                                                                        fmt.Println(v6, v6, v2)
                                                                                                        fmt.Println(v8)
                                                                                                    } else {
                                                                                                        switch ("foo bar" + "") {
                                                                                                        case "":
                                                                                                            v121 := GenS2{F0: 46, F1: (^(v44))}
                                                                                                            _ = v121
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        v43 = v43
                                                                                                                        var v122 float32 = v19
                                                                                                                        var v123 uint32 = v3
                                                                                                                        _ = v122
                                                                                                                        _ = v123
                                                                                                                    } else {
                                                                                                                        var v124 int32 = (v1 - (v1 % (((v1 - (v1 / ((int32(v28)) | 1)))) | 1)))
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    v44 -= (^(v44))
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v120
                                                                                                }
                                                                                                v118 = v43
                                                                                                _ = v118
                                                                                            }
                                                                                        case "hello":
                                                                                            v3 = (v3 | 2147483647)
                                                                                        default:
                                                                                            v28 -= (^(v28))
                                                                                        }
                                                                                        if false {
                                                                                            v125 := GenS11{F0: true, F1: (v3 + v6)}
                                                                                            fmt.Println(v2, v44)
                                                                                            switch v43 {
                                                                                            case "":
                                                                                                if (false && (0 > 32767)) {
                                                                                                    var v126 uint16 = 255
                                                                                                    _ = v126
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v127 := GenS14{F0: 255}
                                                                                                _ = v127
                                                                                            }
                                                                                            _ = v125
                                                                                        }
                                                                                    }
                                                                                    var v128 float32 = v19
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(fi33, fi33, v8)
                                                                        default:
                                                                            switch fi33 {
                                                                            case -629:
                                                                                v6 = v6
                                                                            default:
                                                                                switch (fi33 * fi33) {
                                                                                case 128:
                                                                                    switch int(v44) {
                                                                                    case 4294967295:
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v3
                                                                                        }
                                                                                    case -1:
                                                                                        if (((!(395 == (v1 ^ 2147483647))) && (801 >= -1)) && true) {
                                                                                            var v129 int16 = -363
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = v6
                                                                                            }
                                                                                            v5 = 854
                                                                                            _ = v129
                                                                                        } else {
                                                                                            v130 := GenS4{F0: 1000, F1: (v3 ^ v3)}
                                                                                            _ = v130
                                                                                        }
                                                                                    default:
                                                                                        v131 := GenS16{F0: (v28 / ((v28) | 1))}
                                                                                        _ = v131
                                                                                    }
                                                                                case 1:
                                                                                    v5 = v5
                                                                                }
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(fi33)
                                                                    }
                                                                }
                                                                fmt.Println(v8, v28)
                                                                _ = v52
                                                            }
                                                            _ = v50
                                                        }
                                                        var v132 int16 = -32768
                                                        _ = v132
                                                    }
                                                case 2147483647:
                                                    if (true || true) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (!((^(v6)) >= 4294967295)) {
                                                            v28 = (v28 ^ (v28 & 1))
                                                        } else {
                                                            switch fi33 {
                                                            case -343:
                                                                fmt.Println(v6)
                                                            case 2147483647:
                                                                var v133 uint8 = 57
                                                                _ = v133
                                                            case 1:
                                                                if len(v4) > 0 {
                                                                    v4[0] = uint32(v5)
                                                                }
                                                            default:
                                                                var v134 int8 = 0
                                                                _ = v134
                                                            }
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                v28 += v28
                                            default:
                                                v43 = ("foo bar" + v43)
                                            }
                                            fmt.Println(fi32, v6)
                                        }
                                        _ = v43
                                        _ = v44
                                    case 2147483647:
                                        var v135 int32 = -1
                                        _ = v135
                                    case 1:
                                        v6 -= (v6 ^ v3)
                                    default:
                                        var v136 int16 = 1000
                                        _ = v136
                                    }
                                case 150:
                                    v3 = v6
                                case -1:
                                    switch (^(fi32)) {
                                    case 127:
                                        fmt.Println(v19, v28)
                                    case 722:
                                        if (0.1 >= v8) {
                                            fmt.Println(fi33, v1)
                                            var v137 int8 = -1
                                            v6 = v3
                                            v137 = v137
                                            _ = v137
                                        }
                                    }
                                default:
                                    if (-3 >= -97) {
                                        v3 = (v6 | (v6 >> (((v5 >> ((4294967295) & 63))) & 63)))
                                    } else {
                                        if false {
                                            var v138 int64 = (v28 ^ (-(v28)))
                                            _ = v138
                                        }
                                        if (!true) {
                                            v19 = (0.1 + float32(1e10))
                                            var v139 int32 = v1
                                            _ = v139
                                        } else {
                                            if true {
                                                if true {
                                                    v6 = v6
                                                }
                                                var v140 bool = ((!(!true)) && true)
                                                var v141 string = (((v2 + "\n\t") + "\n\t") + v2)
                                                _ = v140
                                                _ = v141
                                            }
                                            switch fi33 {
                                            case -769:
                                                if true {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v142 float64 = (float64((((v19 * float32(1e10)) - v19) + v19)) + (((float64((v19 / (v19 * v19))) / (3.14159 * v8)) - 3.14159) + float64(v19)))
                                                    v143 := GenS5{F0: 495, F1: 42, F2: v28}
                                                    v144 := GenS1{F0: 801, F1: v2, F2: 128}
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                } else {
                                                    v8 -= v8
                                                }
                                            case -2147483648:
                                                if false {
                                                    fmt.Println(v1)
                                                    if true {
                                                        v6 -= v3
                                                        var v145 string = v2
                                                        _ = v145
                                                    }
                                                }
                                            case -892:
                                                v28 = v28
                                                switch fi33 {
                                                case 404:
                                                    fmt.Println(v28)
                                                default:
                                                    var v146 uint16 = 1
                                                    _ = v146
                                                }
                                                _ = len(v4)
                                            }
                                            v5 += (v5 >> ((v5) & 63))
                                        }
                                        v147 := GenS16{F0: int64(v3)}
                                        _ = v147
                                    }
                                }
                                v4 = append(v4, 2)
                                if (true || true) {
                                    v4 = append(v4, v3)
                                } else {
                                    fmt.Println(v3, v8, v28)
                                    var v148 string = v2
                                    _ = len(v4)
                                    _ = v148
                                }
                            }
                            var v149 int64 = (v28 ^ -9223372036854775808)
                            v150 := GenS9{F0: 461, F1: false}
                            v3 += v6
                            _ = v149
                            _ = v150
                            _ = v38
                        }
                        switch "" {
                        case "hello":
                            v151 := GenS15{F0: 2, F1: (v5 & v5), F2: 0}
                            _ = v151
                        case "foo bar":
                            v1 -= v1
                        default:
                            var v152 string = (v2 + (v2 + ""))
                            _ = v152
                        }
                    }
                    v4 = append(v4, (v3 / (((v6 << ((v5) & 63))) | 1)))
                    if len(v4) > 0 {
                        v4[0] = 838
                    }
                    if ((fi33 << ((v5) & 63)) == (fi32 << ((v5) & 63))) {
                        v1 -= (v1 >> ((42) & 63))
                    } else {
                        if (uint8(v5) == 128) {
                            fmt.Println(v3, v2)
                            v153 := GenS17{F0: v1, F1: true, F2: (v28 >> ((v5) & 63))}
                            _ = v153
                        } else {
                            fmt.Println(v1)
                            v4 = append(v4, v3)
                            var v154 int16 = int16(v28)
                            _ = v154
                        }
                        if false {
                            v2 = v2
                            if len(v4) > 0 {
                                v4[0] = (v3 >> ((65535) & 63))
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    var v155 float32 = (0.1 / 3.14159)
                    _ = fi33
                    _ = v155
                }
                v156 := GenS16{F0: (v28 - 0)}
                _ = v156
            } else {
                var v157 uint32 = (^(v3))
                _ = v157
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v6 = 2147483647
            _ = fi32
        }
    default:
        v1 -= -2147483648
    }
    for fi158 := 0; fi158 < 10; fi158++ {
        if false {
            var v159 int16 = 1000
            _ = v159
        } else {
            v28 += (^(v28))
            v160 := GenS5{F0: 0, F1: 2, F2: (-(v28))}
            if true {
                v19 -= (((v19 - v19) * (v19 / v19)) / (v19 + v19))
                for fi161 := 0; fi161 < 5; fi161++ {
                    v1 = (v1 % ((v1) | 1))
                    var v162 uint8 = 42
                    if false {
                        if ((v162 & v162) <= (v162 >> ((18446744073709551615) & 63))) {
                            var v163 int8 = 94
                            v3 -= (v6 & v6)
                            v3 = (^(v6))
                            _ = v163
                        } else {
                            v6 = (^(v6))
                        }
                        if true {
                            switch v2 {
                            case "a":
                                v28 += v28
                            }
                        } else {
                            v28 = v28
                            var v164 int8 = 1
                            v165 := GenS16{F0: (v28 * -1)}
                            _ = v164
                            _ = v165
                        }
                        v166 := GenS7{F0: (v28 * v28), F1: v2}
                        _ = v166
                    } else {
                        v4 = append(v4, uint32(v5))
                        if len(v4) > 0 {
                            v4[0] = 483
                        }
                        if true {
                            var v167 int64 = v28
                            _ = v167
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v168 := GenS8{F0: 2}
                            _ = v168
                        }
                    }
                    _ = fi161
                    _ = v162
                }
                for fi169 := 0; fi169 < 14; fi169++ {
                    var v170 int16 = 1000
                    var v171 uint = (v5 % ((18446744073709551615) | 1))
                    _ = fi169
                    _ = v170
                    _ = v171
                }
            } else {
                if true {
                    var v172 uint8 = 127
                    _ = v172
                }
                v1 = v1
                if true {
                    v3 -= (v6 ^ 26)
                } else {
                    if (true && false) {
                        fmt.Println(v28, v19, v6)
                        v3 = (v3 | v6)
                        v8 = (v8 / 1.0)
                    } else {
                        v1 = 1
                    }
                }
            }
            _ = v160
        }
        _ = fi158
    }
    var v173 float32 = float32(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v173:", v173)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v173
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
