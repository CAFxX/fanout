package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32; F1 int64; F2 uint }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int64; F1 uint64; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int64; F1 string; F2 uint }
type GenS13 struct { F0 int8; F1 int }
func genFn0() int {
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 103}
            v9 := []int{-1, 32768}
            var v10 int = fi6
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        fmt.Println(fi6)
        switch (fi6 - fi6) {
        case 4294967295:
            var v11 string = "a"
            _ = v11
        default:
            v12 := []int{fi6, (fi6 >> ((249) & 63)), (-(fi6))}
            _ = v12
        }
        _ = fi6
    }
    v13 := GenS2{F0: 2147483647, F1: -2147483648}
    var v14 uint = 42
    var v15 int32 = 38
    v14 = 65535
    _ = v13
    _ = v14
    _ = v15
    return 801
}

func main() {
    var v1 int = 32768
    v2 := []uint32{}
    var v3 uint8 = 0
    var v4 int8 = -1
    var v5 string = ("foo bar" + "\n\t")
    v4 -= v4
    _ = len(v2)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 float32 = (((1e10 + 1.5) / 3.14159) / 0.1)
        var v18 float64 = -2.25
        if (false || (0 >= 9223372036854775807)) {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS3{F0: (!((v4 - (v4 >> ((323) & 63))) == (v4 * (v4 / ((v4) | 1))))), F1: 1, F2: 788}
                v21 := GenS3{F0: true, F1: 4294967295, F2: 887}
                _ = fi19
                _ = v20
                _ = v21
            }
            v1 += (^(v1))
            if false {
                if true {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    var v22 uint64 = 1
                    _ = v22
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v23 := GenS3{F0: true, F1: 4294967295, F2: 1}
                v18 = 0.0
                _ = v23
            }
        } else {
            var v24 bool = false
            fmt.Println(fi16)
            for fi25 := 0; fi25 < 16; fi25++ {
                if v24 {
                    var v26 uint16 = 0
                    var v27 uint64 = 59
                    v28 := GenS4{F0: -2147483648, F1: 9223372036854775807, F2: 1}
                    switch ("" + v5) {
                    case "\n\t":
                        var v29 uint16 = v26
                        _ = v29
                    case "":
                        fmt.Println(v26, v1, fi16)
                    }
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    if v24 {
                        _ = len(v2)
                        v17 -= v17
                    }
                    if ((42 == v4) && false) {
                        v24 = ((v17 <= 3.14159) || (!(80 == 4294967295)))
                    }
                    switch (v1 + fi25) {
                    case 127:
                        fmt.Println(v17, v1, v18)
                    default:
                        v1 = (v1 | v1)
                    }
                }
                if (!(!(2147483647 > 0))) {
                    var v30 uint64 = 4294967295
                    if true {
                        var v31 int8 = 42
                        if ((!true) || v24) {
                            if (int16(v30) <= -1) {
                                if v24 {
                                    var v32 uint32 = 2
                                    var v33 uint16 = 566
                                    _ = v32
                                    _ = v33
                                }
                                if (!v24) {
                                    switch "\n\t" {
                                    case "hello":
                                        v34 := GenS4{F0: -2147483648, F1: 170, F2: 42}
                                        _ = v34
                                    case "a":
                                        var v35 float64 = v18
                                        v36 := GenS4{F0: 2147483647, F1: -452, F2: 352}
                                        if (v3 >= 11) {
                                            v37 := GenS1{F0: v3}
                                            _ = v37
                                        } else {
                                            fmt.Println(v1, fi25, v35)
                                            v38 := GenS2{F0: 2147483648, F1: 1}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v38
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 390
                                        }
                                        var v39 int16 = -32768
                                        _ = v35
                                        _ = v36
                                        _ = v39
                                    }
                                } else {
                                    if v24 {
                                        var v40 int = 32767
                                        _ = v40
                                    }
                                    if (0 == 0) {
                                        var v41 float32 = v17
                                        var v42 uint = 4294967295
                                        _ = v41
                                        _ = v42
                                    }
                                    if (1000 <= 1) {
                                        var v43 int16 = 42
                                        var v44 uint64 = v30
                                        _ = len(v2)
                                        v4 = v4
                                        v18 = (v18 + (((v18 - (v18 - float64(v17))) - (float64((((1e10 - v17) * 1e10) - v17)) / (float64(v17) + (v18 + 3.14159)))) * (((float64(3.14159) - (v18 * -2.25)) - 3.14159) + 0.0)))
                                        _ = v43
                                        _ = v44
                                    }
                                    v5 = ((("" + v5) + v5) + v5)
                                    v1 -= 747
                                    if (v24 && (143 >= 42)) {
                                        switch (fi16 + genFn0()) {
                                        case 2147483647:
                                            fmt.Println(v24, v5)
                                        case -9223372036854775808:
                                            _ = len(v2)
                                        }
                                        switch "a" {
                                        case "":
                                            if (1 != 2) {
                                                if true {
                                                    if (v17 != v17) {
                                                        var v45 int64 = -401
                                                        if (v24 && v24) {
                                                            fmt.Println(fi16, fi16)
                                                            switch ("hello" + v5) {
                                                            case "hello":
                                                                v46 := GenS5{F0: 255, F1: (v5 + v5)}
                                                                _ = v46
                                                            case "a":
                                                                var v47 int16 = 520
                                                                _ = v47
                                                            case "":
                                                                v48 := GenS6{F0: 0, F1: v5}
                                                                _ = v48
                                                            }
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 4294967295
                                                        }
                                                        _ = len(v2)
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v4, v5)
                                                        switch v5 {
                                                        case "a":
                                                            v49 := GenS6{F0: 653, F1: "foo bar"}
                                                            _ = v49
                                                        case "":
                                                            var v50 uint64 = v30
                                                            _ = v50
                                                        }
                                                    }
                                                    v51 := GenS3{F0: false, F1: v30, F2: 849}
                                                    _ = v51
                                                } else {
                                                    v52 := GenS2{F0: 4294967295, F1: 997}
                                                    fmt.Println(fi16)
                                                    v53 := GenS7{F0: 1, F1: uint64(v3), F2: (v30 > 42)}
                                                    fmt.Println(v5)
                                                    _ = v52
                                                    _ = v53
                                                }
                                                switch (fi25 << ((661) & 63)) {
                                                case 0:
                                                    fmt.Println(v24)
                                                }
                                            }
                                        case "foo bar":
                                            v5 = ""
                                        }
                                    } else {
                                        fmt.Println(v30, v4, v17)
                                    }
                                    fmt.Println(fi16, v4)
                                    if (false && (-2147483648 != 42)) {
                                        if v24 {
                                            switch 2147483647 {
                                            case 462:
                                                v2 = append(v2, 2147483648)
                                            default:
                                                var v54 int64 = 9223372036854775807
                                                _ = v54
                                            }
                                            switch v5 {
                                            case "a":
                                                var v55 int16 = -623
                                                _ = v55
                                            }
                                            var v56 int32 = -1
                                            _ = v56
                                        } else {
                                            var v57 uint = 0
                                            switch ("\n\t" + ((v5 + "hello") + "")) {
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "":
                                                if (((111 >= -1) || (v5 > v5)) && false) {
                                                    fmt.Println(v24, v17, v31)
                                                }
                                            case "a":
                                                v1 += fi25
                                            default:
                                                v1 = fi16
                                            }
                                            _ = v57
                                        }
                                    }
                                }
                                v24 = (-376 <= -1)
                            } else {
                                v58 := GenS8{F0: 2}
                                fmt.Println(v30)
                                v5 = (("a" + "\n\t") + (v5 + (v5 + ((v5 + ("foo bar" + "foo bar")) + ""))))
                                _ = v58
                            }
                            var v59 uint = 18446744073709551615
                            _ = v59
                        } else {
                            v4 = (v31 | v4)
                        }
                        _ = v31
                    } else {
                        v30 = (v30 << ((244) & 63))
                    }
                    _ = v30
                }
                _ = len(v2)
                _ = fi25
            }
            _ = v24
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v4, v5)
    switch "\n\t" {
    case "foo bar":
        v4 = v4
    }
    if (!(-886 == 366)) {
        for fi60 := 0; fi60 < 10; fi60++ {
            v61 := GenS7{F0: 1, F1: 42, F2: false}
            v3 = 128
            var v62 uint64 = 0
            _ = fi60
            _ = v61
            _ = v62
        }
    } else {
        switch v5 {
        case "foo bar":
            if true {
                v3 = (v3 << ((0) & 63))
                if (!(18446744073709551615 > 18446744073709551615)) {
                    v2 = append(v2, 370)
                }
            } else {
                if false {
                    switch 2147483647 {
                    case 127:
                        fmt.Println(v3, v5)
                    }
                }
                if false {
                    var v63 uint64 = 0
                    var v64 int32 = 42
                    _ = v63
                    _ = v64
                } else {
                    v65 := GenS9{F0: 9223372036854775807, F1: 121}
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v65
                }
            }
        case "":
            for fi66 := 0; fi66 < 7; fi66++ {
                var v67 int16 = -126
                _ = fi66
                _ = v67
            }
        case "hello":
            if true {
                v4 = int8(v1)
            }
        default:
            v3 += (v3 & v3)
        }
        v5 = v5
    }
    v68 := GenS7{F0: 42, F1: 2, F2: false}
    if len(v2) > 0 {
        v2[0] = 922
    }
    if true {
        if false {
            if false {
                var v69 int = v1
                v70 := GenS10{F0: 11}
                _ = v69
                _ = v70
            }
            v2 = append(v2, 2)
        } else {
            if false {
                for fi71 := 0; fi71 < 4; fi71++ {
                    if (!false) {
                        for fi72 := 0; fi72 < 11; fi72++ {
                            var v73 int8 = v4
                            _ = fi72
                            _ = v73
                        }
                    }
                    switch 2147483647 {
                    case -193:
                        var v74 int32 = -2147483648
                        _ = v74
                    case -9223372036854775808:
                        _ = len(v2)
                    case -121:
                        if false {
                            var v75 int32 = -394
                            v76 := GenS11{F0: "a"}
                            switch fi71 {
                            case -9223372036854775808:
                                if ((v75 >> ((4294967295) & 63)) != -1) {
                                    v1 = genFn0()
                                    for fi77 := 0; fi77 < 12; fi77++ {
                                        v75 += v75
                                        fmt.Println(v1, v4)
                                        _ = fi77
                                    }
                                    v78 := GenS8{F0: (v3 * (^(v3)))}
                                    _ = v78
                                } else {
                                    v79 := GenS12{F0: 0, F1: "hello", F2: 18446744073709551615}
                                    v4 = v4
                                    _ = v79
                                }
                            }
                            _ = v75
                            _ = v76
                        }
                    }
                    v80 := GenS11{F0: v5}
                    _ = fi71
                    _ = v80
                }
                fmt.Println(v4)
                if len(v2) > 0 {
                    v2[0] = 0
                }
            }
        }
        if (!false) {
            if (203 <= 2147483647) {
                v81 := GenS13{F0: v4, F1: (v1 + (v1 | (^(v1))))}
                var v82 int16 = 1
                fmt.Println(v82)
                _ = v81
                _ = v82
            } else {
                v5 = v5
            }
        }
        var v83 uint8 = (v3 % ((v3) | 1))
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
}
