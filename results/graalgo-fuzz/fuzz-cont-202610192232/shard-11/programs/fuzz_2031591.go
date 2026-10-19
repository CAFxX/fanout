package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int; F1 uint8; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int8; F1 bool; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint32; F1 uint8; F2 string }
type GenS18 struct { F0 int32; F1 uint16 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint16 }
func genFn0() int8 {
    if false {
        v5 := GenS1{F0: 511}
        _ = v5
    } else {
        var v6 bool = false
        var v7 int32 = 42
        var v8 int = 9223372036854775807
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS2{F0: ("a" + "foo bar"), F1: -1}
        switch "\n\t" {
        case "\n\t":
            v11 := GenS3{F0: 128, F1: true}
            _ = v11
        case "":
            for fi12 := 0; fi12 < 5; fi12++ {
                fmt.Println(fi9)
                var v13 string = "a"
                v14 := GenS4{F0: v13, F1: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    if true {
        var v15 uint8 = 142
        _ = v15
    } else {
        v16 := GenS4{F0: "", F1: 35}
        var v17 uint = 553
        v18 := GenS5{F0: -1}
        _ = v16
        _ = v17
        _ = v18
    }
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 245
    var v3 int8 = 100
    var v4 uint32 = 319
    if true {
        switch (("" + (("" + ("hello" + ("foo bar" + ("\n\t" + "a")))) + "")) + "\n\t") {
        case "":
            var v19 int16 = 236
            _ = v19
        default:
            v20 := []int{128, 127, 42}
            _ = v20
        }
        var v21 int = 128
        _ = v21
    }
    var v22 int32 = 0
    v23 := GenS6{F0: v2, F1: v2, F2: 2}
    if true {
        v24 := GenS7{F0: (2147483647 > 251), F1: 2147483647}
        _ = v24
    }
    var v25 float64 = 0.1
    switch int(v22) {
    case -787:
        v25 -= ((((1e10 / 1.5) * v25) + float64(1e10)) * (((float64(0.0) * v25) + (float64((float32(v25) + 1.5)) * -2.25)) * v25))
    default:
        var v26 int32 = (-(v22))
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := v1[fi27]
        v4 = v28
        v2 = (v2 | v2)
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        _ = len(v1)
        _ = fi29
    }
    delete(v1, 43)
    if true {
        var v30 int8 = 100
        v31 := []int{0, 1, -828}
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                v25 -= (v25 * (-2.25 * (((v25 + v25) - v25) - v25)))
                var v34 uint64 = 740
                var v35 bool = true
                if false {
                    switch ("" + "hello") {
                    case "\n\t":
                        if v35 {
                            var v36 uint64 = (v34 / (((v34 % (((v34 - (v34 - 776))) | 1))) | 1))
                            _ = v36
                        } else {
                            switch ("" + "\n\t") {
                            case "\n\t":
                                if false {
                                    v37 := GenS8{F0: 32767, F1: (^(v34))}
                                    switch fi33 {
                                    case 9223372036854775807:
                                        v38 := GenS3{F0: uint8(v2), F1: true}
                                        _ = v38
                                    case 4294967295:
                                        var v39 uint = 1
                                        _ = v39
                                    case 980:
                                        v25 += v25
                                    }
                                    _ = v37
                                }
                            case "":
                                var v40 uint64 = v34
                                _ = v40
                            }
                            var v41 uint = 4294967295
                            _ = v41
                        }
                    case "foo bar":
                        if v35 {
                            v42 := GenS9{F0: fi33, F1: 42, F2: 1000}
                            var v43 uint32 = v4
                            v44 := GenS10{F0: genFn0()}
                            switch "foo bar" {
                            case "hello":
                                v34 += (v34 - v34)
                            default:
                                if ((1.5 * v25) >= 0.1) {
                                    v4 = v43
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v45 float64 = v25
                                        _ = v45
                                    }
                                } else {
                                    fmt.Println(v25, v34, fi33)
                                    v46 := GenS11{F0: v30, F1: v35, F2: 2}
                                    v47 := GenS8{F0: -32768, F1: (v34 / (((v34 | (v34 + v34))) | 1))}
                                    _ = v46
                                    _ = v47
                                }
                            }
                            var v48 int16 = -32768
                            v43 = v43
                            v49 := GenS12{F0: 9223372036854775807}
                            _ = v42
                            _ = v43
                            _ = v44
                            _ = v48
                            _ = v49
                        } else {
                            fmt.Println(v2)
                            if (-1 == -1) {
                                var v50 float32 = (1e10 - 1e10)
                                fmt.Println(v25)
                                var v51 int64 = -173
                                switch ((("hello" + "hello") + "a") + "foo bar") {
                                case "foo bar":
                                    v3 = v30
                                case "":
                                    if v35 {
                                        v1[fi32] = 135
                                        var v52 bool = false
                                        _ = len(v31)
                                        _ = v52
                                    }
                                case "\n\t":
                                    _ = len(v1)
                                }
                                v22 = (-(v22))
                                fmt.Println(v34, v51, v22)
                                v1[-9223372036854775808] = v4
                                fmt.Println(v22, v51)
                                fmt.Println(v35, v30)
                                if v35 {
                                    v31 = append(v31, 176)
                                } else {
                                    var v53 uint16 = (^(v2))
                                    v31 = append(v31, fi33)
                                    _ = len(v31)
                                    _ = v53
                                }
                                var v54 uint32 = (^(v4))
                                if true {
                                    var v55 bool = v35
                                    var v56 string = "\n\t"
                                    fmt.Println(v4, v54)
                                    var v57 int8 = (v30 & (v30 & v30))
                                    v58 := GenS12{F0: (v51 / (((v51 * v51)) | 1))}
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                var v59 int8 = -1
                                var v60 uint8 = 162
                                delete(v1, -2147483648)
                                _ = v50
                                _ = v51
                                _ = v54
                                _ = v59
                                _ = v60
                            }
                        }
                    }
                    if (v25 < 0.0) {
                        v30 = -3
                        if v35 {
                            v35 = v35
                            v34 += 1
                        } else {
                            v22 -= v22
                            if v35 {
                                switch (fi32 * fi33) {
                                case 42:
                                    switch (fi33 ^ fi32) {
                                    case 2147483647:
                                        var v61 int8 = (v30 * 57)
                                        _ = v61
                                    }
                                case 127:
                                    switch (("a" + "foo bar") + ("foo bar" + "a")) {
                                    case "":
                                        var v62 int8 = 0
                                        _ = v62
                                    case "hello":
                                        switch 127 {
                                        case -9223372036854775808:
                                            if false {
                                                v31 = append(v31, (fi32 + int(v30)))
                                            } else {
                                                v4 = 0
                                                if true {
                                                    var v63 int = (fi32 + 42)
                                                    _ = v63
                                                }
                                                v64 := GenS13{F0: (^(v22)), F1: v35}
                                                _ = v64
                                            }
                                        default:
                                            if (!v35) {
                                                fmt.Println(v2, v25, v4)
                                                var v65 uint64 = (v34 | 18446744073709551615)
                                                _ = v65
                                            } else {
                                                v35 = (uint32(v34) <= 659)
                                            }
                                        }
                                    }
                                case -567:
                                    _ = len(v31)
                                }
                                switch (("foo bar" + ("hello" + "foo bar")) + "") {
                                case "":
                                    var v66 int64 = 0
                                    _ = v66
                                case "a":
                                    _ = len(v31)
                                case "foo bar":
                                    switch "hello" {
                                    case "\n\t":
                                        var v67 int8 = (v30 & 31)
                                        _ = v67
                                    case "foo bar":
                                        switch (fi32 >> ((255) & 63)) {
                                        case 9223372036854775807:
                                            switch (fi33 & (^(fi32))) {
                                            case -194:
                                                v68 := GenS14{F0: 161}
                                                _ = v68
                                            case 32767:
                                                v69 := GenS12{F0: 875}
                                                _ = v69
                                            case -1:
                                                if (!true) {
                                                    var v70 int32 = v22
                                                    _ = v70
                                                } else {
                                                    var v71 int64 = -651
                                                    var v72 int16 = -656
                                                    switch "" {
                                                    case "\n\t":
                                                        var v73 int = (fi33 * (-(fi33)))
                                                        _ = v73
                                                    case "":
                                                        if (584 != 18446744073709551615) {
                                                            var v74 uint8 = 127
                                                            delete(v1, fi33)
                                                            fmt.Println(v25, fi32, v22)
                                                            _ = v74
                                                        }
                                                    default:
                                                        var v75 uint32 = v4
                                                        _ = v75
                                                    }
                                                    _ = v71
                                                    _ = v72
                                                }
                                            }
                                        case 32767:
                                            v76 := GenS6{F0: v2, F1: (v2 ^ v2), F2: 2}
                                            var v77 uint8 = 0
                                            var v78 uint = 4294967295
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                        default:
                                            fmt.Println(v25, v25, v3)
                                        }
                                    case "a":
                                        v79 := GenS15{F0: -1}
                                        _ = v79
                                    }
                                default:
                                    if v35 {
                                        var v80 float32 = 1e10
                                        if v35 {
                                            fmt.Println(v4, v4, v25)
                                            if v35 {
                                                v81 := GenS16{F0: v34}
                                                fmt.Println(v30)
                                                v2 += (v2 << ((18446744073709551615) & 63))
                                                _ = v81
                                            } else {
                                                fmt.Println(v34, v25)
                                            }
                                        }
                                        var v82 int = fi33
                                        v83 := GenS17{F0: (v4 + v4), F1: 42, F2: (("\n\t" + "\n\t") + ("a" + (("hello" + "a") + "foo bar")))}
                                        v84 := GenS18{F0: (v22 & (v22 % (((v22 * v22)) | 1))), F1: 434}
                                        _ = v80
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    } else {
                                        var v85 string = "foo bar"
                                        v4 = 1
                                        _ = v85
                                    }
                                }
                            }
                            v86 := GenS5{F0: (v22 * (v22 << ((18446744073709551615) & 63)))}
                            _ = v86
                        }
                        v87 := GenS19{F0: 0}
                        _ = v87
                    }
                }
                switch int(v3) {
                case 4294967295:
                    v88 := GenS7{F0: (2147483647 != (fi32 / (((fi33 % ((-528) | 1))) | 1))), F1: fi32}
                    _ = v88
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            fmt.Println(v22, v25, v30)
            fmt.Println(v4, fi32, v2)
            _ = fi32
        }
        _ = v30
        _ = v31
    } else {
        for fi89 := 0; fi89 < 8; fi89++ {
            v90 := GenS20{F0: (^(v2))}
            for fi91 := 0; fi91 < 12; fi91++ {
                var v92 int32 = 0
                _ = fi91
                _ = v92
            }
            _ = fi89
            _ = v90
        }
        switch "a" {
        case "a":
            if (true || true) {
                var v93 uint8 = 128
                v3 += v3
                for fi94 := 0; fi94 < 11; fi94++ {
                    v3 = (^(v3))
                    switch "" {
                    case "\n\t":
                        v95 := v1[int(v22)]
                        _ = v95
                    case "hello":
                        var v96 float32 = float32(v25)
                        _ = v96
                    }
                    _ = fi94
                }
                _ = v93
            }
        case "\n\t":
            v22 -= -1
        case "foo bar":
            for fi97 := 0; fi97 < 4; fi97++ {
                var v98 uint64 = 984
                if false {
                    v3 = v3
                    if true {
                        v3 = 127
                        v3 += 100
                    }
                }
                _ = fi97
                _ = v98
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
