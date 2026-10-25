package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS7 struct { F0 string; F1 int8; F2 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 uint64; F1 string; F2 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint; F1 int64; F2 uint32 }
type GenS14 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 624
    var v3 int32 = 1
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(v2, v3, v1)
        _ = fi4
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    v2 -= 2147483647
    var v7 uint16 = 65535
    v8 := []int{0, 2147483647, -822}
    if false {
        v3 = v2
        if len(v8) > 0 {
            v8[0] = -172
        }
    }
    fmt.Println(v1, v3)
    var v9 uint64 = 517
    var v10 uint = 65535
    if true {
        var v11 uint16 = 2
        if ((3.14159 + 3.14159) > (-2.25 - (float32((0.1 + 3.14159)) * 3.14159))) {
            var v12 int = -82
            switch (v12 & v12) {
            case -354:
                var v13 uint8 = 1
                _ = v13
            default:
                var v14 uint = (^(v10))
                _ = v14
            }
            _ = v12
        } else {
            var v15 int16 = 0
            for fi16 := 0; fi16 < 15; fi16++ {
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 uint = (v10 / (((v10 + (v10 - 946))) | 1))
                    _ = fi17
                    _ = v18
                }
                if false {
                    switch int(v9) {
                    case 681:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v9 -= v9
                            var v20 uint32 = 2
                            if true {
                                v10 -= (v10 ^ v10)
                                switch 42 {
                                case -939:
                                    if true {
                                        var v21 float64 = 1.5
                                        v3 = (^(v2))
                                        _ = v21
                                    }
                                case 4294967295:
                                    _ = len(v8)
                                default:
                                    var v22 uint32 = v20
                                    _ = v22
                                }
                            } else {
                                v23 := GenS1{F0: (v15 % (((v15 & (v15 % ((v15) | 1)))) | 1)), F1: -81}
                                v20 = (v20 ^ v20)
                                switch ((("a" + "hello") + "\n\t") + "hello") {
                                case "a":
                                    v7 += 42
                                case "\n\t":
                                    var v24 int8 = -128
                                    _ = v24
                                case "foo bar":
                                    fmt.Println(v11)
                                }
                                _ = v23
                            }
                            v25 := GenS2{F0: v1}
                            fmt.Println(v20, v2, v20)
                            if true {
                                if false {
                                    var v26 float64 = (3.14159 / 1.5)
                                    _ = len(v8)
                                    _ = v26
                                } else {
                                    v8 = append(v8, (fi16 / ((fi19) | 1)))
                                    v27 := GenS3{F0: v2}
                                    _ = v27
                                }
                                v28 := GenS1{F0: (v15 + (^(v15))), F1: (^(v15))}
                                _ = v28
                            }
                            _ = fi19
                            _ = v20
                            _ = v25
                        }
                    case 2147483647:
                        if ((^(v7)) < (v11 - uint16(v10))) {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                v10 -= (v10 & (v10 << ((851) & 63)))
                                v30 := GenS1{F0: 516, F1: v15}
                                _ = fi29
                                _ = v30
                            }
                            v31 := GenS4{F0: (v2 ^ v2), F1: -1, F2: (v9 & 2)}
                            if len(v8) > 0 {
                                v8[0] = (-(fi16))
                            }
                            _ = v31
                        } else {
                            if ((v9 - (v9 << (((v10 ^ (v10 * v10))) & 63))) < (v9 / (((v9 - v9)) | 1))) {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    v10 = 297
                                    if (1e10 != 3.14159) {
                                        v33 := GenS5{F0: v10, F1: v15}
                                        switch fi16 {
                                        case -9223372036854775808:
                                            v2 += 484
                                        default:
                                            var v34 float32 = 1.5
                                            _ = v34
                                        }
                                        if ((v2 | v3) >= (v3 & (v2 / (((v3 * v2)) | 1)))) {
                                            if true {
                                                if (!((v9 - v9) < (v9 + (v9 & v9)))) {
                                                    var v35 int64 = 1
                                                    _ = v35
                                                } else {
                                                    var v36 uint32 = 164
                                                    var v37 float32 = ((0.1 + (1e10 / (((float32(3.14159) + -2.25) - 0.0) / 1.5))) + 1.5)
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                            v3 = v2
                                            var v38 int64 = 1
                                            _ = v38
                                        } else {
                                            v39 := GenS6{F0: 194, F1: 65535, F2: (v2 ^ v3)}
                                            var v40 uint8 = 0
                                            var v41 int8 = -128
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        if len(v8) > 0 {
                                            v8[0] = (fi16 | (fi16 * 127))
                                        }
                                        var v42 float64 = 0.1
                                        if true {
                                            v43 := GenS6{F0: 905, F1: 722, F2: 271}
                                            _ = v43
                                        } else {
                                            fmt.Println(v10, v2)
                                            var v44 bool = true
                                            _ = v44
                                        }
                                        _ = v33
                                        _ = v42
                                    }
                                    _ = fi32
                                }
                                v45 := GenS7{F0: "a", F1: 127, F2: (v7 % (((v11 & (v7 ^ v7))) | 1))}
                                fmt.Println(v7)
                                _ = v45
                            }
                            v10 = v10
                            if false {
                                if false {
                                    switch "foo bar" {
                                    case "":
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        case "foo bar":
                                            switch "hello" {
                                            case "hello":
                                                if (v9 != v9) {
                                                    if false {
                                                        switch (^(fi16)) {
                                                        case 2147483647:
                                                            v7 += v11
                                                        case -9223372036854775808:
                                                            v46 := GenS8{F0: (fi16 >> ((14) & 63))}
                                                            _ = v46
                                                        case -1:
                                                            if (488 == 42) {
                                                                var v47 int16 = -32768
                                                                var v48 uint32 = 185
                                                                if true {
                                                                    for fi49 := 0; fi49 < 2; fi49++ {
                                                                        v50 := GenS9{F0: v1, F1: (^(v1))}
                                                                        _ = fi49
                                                                        _ = v50
                                                                    }
                                                                    v8 = append(v8, -740)
                                                                    if true {
                                                                        var v51 uint16 = v7
                                                                        _ = v51
                                                                    } else {
                                                                        var v52 bool = true
                                                                        _ = v52
                                                                    }
                                                                }
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                        default:
                                                            v3 = v2
                                                        }
                                                        if true {
                                                            if false {
                                                                v1 = 255
                                                                v9 = uint64(v10)
                                                            } else {
                                                                v53 := GenS10{F0: (v9 | v9), F1: ("" + "foo bar"), F2: -1}
                                                                if (0 > 626) {
                                                                    fmt.Println(v10, v2)
                                                                    switch (fi16 & fi16) {
                                                                    case 9223372036854775807:
                                                                        switch 32767 {
                                                                        case 9223372036854775807:
                                                                            if len(v8) > 0 {
                                                                                v8[0] = fi16
                                                                            }
                                                                        case 2147483647:
                                                                            v10 = 65535
                                                                        }
                                                                    }
                                                                    v11 -= (v11 | 255)
                                                                }
                                                                var v54 uint32 = 935
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            if true {
                                                                var v55 uint = v10
                                                                _ = v55
                                                            } else {
                                                                for fi56 := 0; fi56 < 2; fi56++ {
                                                                    var v57 int64 = 9223372036854775807
                                                                    _ = fi56
                                                                    _ = v57
                                                                }
                                                            }
                                                            v7 += v11
                                                        } else {
                                                            var v58 uint32 = 0
                                                            fmt.Println(v1, v10, v15)
                                                            _ = v58
                                                        }
                                                    } else {
                                                        v59 := GenS2{F0: 2}
                                                        _ = v59
                                                    }
                                                    v11 += v7
                                                } else {
                                                    v1 = 1
                                                }
                                            case "foo bar":
                                                v1 -= 127
                                            case "a":
                                                switch (fi16 + fi16) {
                                                case 4294967295:
                                                    v7 += v7
                                                case 2147483647:
                                                    v10 = (v10 + v10)
                                                case -2147483648:
                                                    v60 := GenS2{F0: (^(v1))}
                                                    _ = v60
                                                default:
                                                    if false {
                                                        var v61 float32 = -2.25
                                                        if (!(!(!(true || (v61 != v61))))) {
                                                            if (!false) {
                                                                var v62 int8 = 127
                                                                v10 = (v10 - 620)
                                                                _ = v62
                                                            }
                                                            if len(v8) > 0 {
                                                                v8[0] = (fi16 + 127)
                                                            }
                                                            if (1e10 > (1.5 + v61)) {
                                                                var v63 bool = true
                                                                for fi64 := 0; fi64 < 16; fi64++ {
                                                                    v8 = append(v8, 973)
                                                                    if (!false) {
                                                                        _ = len(v8)
                                                                        var v65 string = "hello"
                                                                        v11 += (v7 - 690)
                                                                        v8 = append(v8, fi16)
                                                                        v66 := GenS11{F0: 4294967295}
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                    var v67 string = "foo bar"
                                                                    var v68 int16 = v15
                                                                    v9 -= v9
                                                                    _ = fi64
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                _ = v63
                                                            } else {
                                                                switch "a" {
                                                                case "hello":
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v69 uint = v10
                                                                        _ = v69
                                                                    case "a":
                                                                        v70 := GenS12{F0: fi16}
                                                                        _ = v70
                                                                    }
                                                                case "":
                                                                    var v71 uint = (v10 % (((v10 * (v10 & v10))) | 1))
                                                                    _ = v71
                                                                }
                                                                fmt.Println(v1, v1)
                                                                v7 = (^(v11))
                                                            }
                                                        }
                                                        _ = v61
                                                    }
                                                }
                                            }
                                        case "a":
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                v72 := GenS13{F0: (v10 & (v10 / (((v10 & (v10 * (^(v10))))) | 1))), F1: -9223372036854775808, F2: 220}
                                                _ = v72
                                            case "foo bar":
                                                fmt.Println(fi16, v3, v2)
                                            default:
                                                v3 = (v3 | v2)
                                            }
                                        }
                                    }
                                }
                                v10 = (v10 - uint(v9))
                                switch (((("a" + ("a" + ("\n\t" + "a"))) + "foo bar") + ("hello" + ("a" + "\n\t"))) + "a") {
                                case "hello":
                                    var v73 int32 = (v2 | -171)
                                    _ = v73
                                }
                            } else {
                                if (!(true && (18446744073709551615 != uint(v9)))) {
                                    var v74 float64 = 3.14159
                                    v2 = (v2 + (-(v3)))
                                    _ = v74
                                } else {
                                    if true {
                                        v75 := GenS2{F0: v1}
                                        _ = v75
                                    } else {
                                        v3 -= (v3 + v3)
                                    }
                                    var v76 int64 = -9223372036854775808
                                    v7 -= (v11 * v11)
                                    _ = v76
                                }
                            }
                        }
                    default:
                        for fi77 := 0; fi77 < 6; fi77++ {
                            v78 := GenS13{F0: (v10 | (v10 & v10)), F1: -1, F2: 2147483648}
                            fmt.Println(v1)
                            v1 = 0
                            _ = fi77
                            _ = v78
                        }
                    }
                    if (!(!false)) {
                        switch "hello" {
                        case "a":
                            v79 := GenS14{F0: 2, F1: 2147483647}
                            _ = v79
                        default:
                            v80 := GenS3{F0: (v2 << ((1) & 63))}
                            _ = v80
                        }
                        _ = len(v8)
                        var v81 int = fi16
                        _ = v81
                    }
                }
                var v82 uint8 = uint8(v9)
                _ = fi16
                _ = v82
            }
            _ = v15
        }
        v83 := GenS12{F0: 9223372036854775807}
        _ = v11
        _ = v83
    } else {
        var v84 uint = 4294967295
        if false {
            var v85 uint8 = (v1 % ((v1) | 1))
            var v86 int32 = 112
            var v87 int32 = v3
            _ = v85
            _ = v86
            _ = v87
        }
        _ = v84
    }
    var v88 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v88:", v88)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v88
    _ = v9
}
