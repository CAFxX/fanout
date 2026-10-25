package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint8; F2 int }
type GenS5 struct { F0 uint32; F1 bool; F2 string }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint16; F1 string; F2 int8 }
type GenS10 struct { F0 int32; F1 uint }
func genFn0(p0 uint16) string {
    v4 := GenS1{F0: -57}
    p0 = p0
    var v5 uint16 = (p0 | (p0 + 39))
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = 65535
        var v8 uint64 = 2
        fmt.Println(fi6, v5)
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = 590
    _ = p0
    _ = v4
    _ = v5
    return "a"
}

func main() {
    var v1 float64 = ((1e10 + 3.14159) + -2.25)
    v2 := make(map[int]int)
    v3 := []int{32768, -209, 128}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: 100}
    v10 := GenS2{F0: "\n\t", F1: -118, F2: "foo bar"}
    var v11 int32 = 2147483647
    v1 -= 1.5
    v12 := GenS3{F0: "foo bar"}
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 string = "foo bar"
        var v15 int = fi13
        v11 = 2147483647
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                v11 = 1
                var v18 int64 = -126
                v19 := GenS3{F0: "\n\t"}
                _ = fi17
                _ = v18
                _ = v19
            }
            v11 = (v11 / ((v11) | 1))
            _ = fi16
        }
        v20 := GenS4{F0: 65535, F1: 0, F2: -126}
        switch 194 {
        case 32767:
            switch "hello" {
            case "\n\t":
                delete(v2, -291)
            case "hello":
                v21 := GenS1{F0: 1}
                _ = v21
            default:
                var v22 int32 = 262
                _ = v22
            }
        case 128:
            if false {
                var v23 float64 = (0.1 / 1.0)
                v24 := GenS2{F0: "a", F1: 100, F2: "a"}
                _ = v23
                _ = v24
            } else {
                var v25 int = -753
                switch (("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + ("hello" + "a")))) + "foo bar") {
                case "hello":
                    v26 := v2[(v25 * (v25 >> ((0) & 63)))]
                    _ = v26
                case "":
                    if true {
                        var v27 int = 1
                        delete(v2, 42)
                        v28 := GenS5{F0: 2, F1: ((2 <= 42) || false), F2: ("foo bar" + "\n\t")}
                        _ = v27
                        _ = v28
                    } else {
                        switch 127 {
                        case 42:
                            fmt.Println(v11)
                        default:
                            v1 -= v1
                        }
                        switch (-(v25)) {
                        case 1:
                            v29 := GenS4{F0: 155, F1: 203, F2: -156}
                            _ = v29
                        case 9223372036854775807:
                            fmt.Println(v25)
                        case 42:
                            for fi30 := 0; fi30 < 20; fi30++ {
                                delete(v2, 320)
                                var v31 int8 = 0
                                _ = fi30
                                _ = v31
                            }
                        default:
                            switch (v25 << ((725) & 63)) {
                            case 49:
                                var v32 int32 = (v11 * v11)
                                _ = v32
                            default:
                                delete(v2, v25)
                            }
                        }
                    }
                default:
                    v33 := v2[v25]
                    _ = v33
                }
                switch "a" {
                case "foo bar":
                    for fi34 := 0; fi34 < 5; fi34++ {
                        _ = len(v3)
                        _ = fi34
                    }
                case "hello":
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 uint8 = 1
                        _ = fi35
                        _ = v36
                    }
                case "a":
                    v25 -= (v25 >> ((71) & 63))
                default:
                    var v37 uint32 = 647
                    _ = v37
                }
                _ = v25
            }
        default:
            delete(v2, -2147483648)
        }
        _ = v20
    } else {
        v38 := GenS4{F0: 51, F1: 150, F2: 212}
        var v39 float32 = 1.5
        _ = v38
        _ = v39
    }
    v40 := GenS6{F0: (v11 ^ 0), F1: -9223372036854775808}
    v11 = 42
    if true {
        if true {
            for fi41 := 0; fi41 < 3; fi41++ {
                var v42 int8 = 100
                _ = fi41
                _ = v42
            }
            switch "" {
            case "\n\t":
                v11 = (v11 / (((v11 ^ v11)) | 1))
            case "":
                delete(v2, 1)
            default:
                for fi43 := 0; fi43 < 17; fi43++ {
                    v44 := GenS2{F0: "", F1: -1, F2: "hello"}
                    _ = fi43
                    _ = v44
                }
            }
        } else {
            v2[-389] = -569
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v45 int = -2147483648
        _ = v45
    } else {
        if true {
            fmt.Println(v1, v11)
            v1 += 1e10
        } else {
            v46 := GenS7{F0: -703, F1: 605}
            if (!(467 > 65535)) {
                if false {
                    for fi47 := 0; fi47 < 10; fi47++ {
                        switch ("a" + "\n\t") {
                        case "":
                            if false {
                                v11 = v11
                                var v48 float32 = 0.1
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    var v50 int16 = -32768
                                    v48 = float32((v1 + float64(v48)))
                                    var v51 int64 = -9223372036854775808
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                                _ = v48
                            }
                        case "foo bar":
                            var v52 uint8 = 112
                            _ = v52
                        default:
                            v53 := GenS8{F0: (fi47 >> ((619) & 63)), F1: 18446744073709551615}
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 18; fi54++ {
                            fmt.Println(v1, v1, fi54)
                            v2[(fi54 + (fi54 * (fi54 ^ fi54)))] = (fi47 - fi47)
                            _ = len(v3)
                            _ = fi54
                        }
                        if (-416 == 0) {
                            v55 := GenS1{F0: -29}
                            v56 := GenS9{F0: 615, F1: "\n\t", F2: 1}
                            _ = v55
                            _ = v56
                        }
                        _ = fi47
                    }
                }
                if (!((943 >= 1) && (4294967295 <= 156))) {
                    if true {
                        v57 := GenS10{F0: (v11 & v11), F1: 0}
                        switch 32768 {
                        case 1:
                            for fi58 := 0; fi58 < 9; fi58++ {
                                if true {
                                    if false {
                                        fmt.Println(v11)
                                        var v59 int32 = (v11 % (((-(v11))) | 1))
                                        if false {
                                            v11 -= (-(v59))
                                            var v60 int32 = v11
                                            _ = v60
                                        } else {
                                            for fi61 := 0; fi61 < 2; fi61++ {
                                                if (!((195 < 127) || (true || (false || (-1 < -75))))) {
                                                    v1 += ((v1 / v1) - (v1 * (((float64(0.1) - ((1e10 / v1) * v1)) * (-2.25 + v1)) / v1)))
                                                    delete(v2, fi61)
                                                    v62 := GenS9{F0: 65535, F1: ("hello" + ("" + ("a" + ""))), F2: 100}
                                                    _ = v62
                                                } else {
                                                    v59 -= 1
                                                    if false {
                                                        switch ("a" + "a") {
                                                        case "foo bar":
                                                            v3 = append(v3, (fi61 + (^(fi58))))
                                                        case "a":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v63 := GenS3{F0: "a"}
                                                                _ = v63
                                                            case "a":
                                                                var v64 int16 = -32768
                                                                _ = v64
                                                            case "hello":
                                                                v65 := GenS9{F0: 0, F1: "hello", F2: -123}
                                                                _ = v65
                                                            default:
                                                                v59 += v59
                                                            }
                                                        }
                                                        var v66 uint8 = 95
                                                        switch 42 {
                                                        case 32767:
                                                            if false {
                                                                v2[(^(fi61))] = fi58
                                                                if ("a" < "foo bar") {
                                                                    var v67 uint32 = 65535
                                                                    _ = v67
                                                                }
                                                                if (v11 == v59) {
                                                                    switch 2147483647 {
                                                                    case -1:
                                                                        switch ("a" + "foo bar") {
                                                                        case "foo bar":
                                                                            v68 := GenS2{F0: (("hello" + "hello") + "\n\t"), F1: 1, F2: "foo bar"}
                                                                            _ = v68
                                                                        }
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (fi61 << ((0) & 63))
                                                                        }
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                } else {
                                                                    var v69 uint = 625
                                                                    var v70 int16 = 32767
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            } else {
                                                                fmt.Println(v66)
                                                            }
                                                            fmt.Println(v1, fi61)
                                                            var v71 int64 = 708
                                                            var v72 uint16 = 1
                                                            if false {
                                                                v3 = append(v3, fi61)
                                                                v73 := GenS8{F0: fi61, F1: 255}
                                                                _ = v73
                                                            }
                                                            switch int(v71) {
                                                            case 9223372036854775807:
                                                                var v74 int16 = -905
                                                                _ = v74
                                                            }
                                                            v59 = 2147483647
                                                            _ = v71
                                                            _ = v72
                                                        case 127:
                                                            fmt.Println(v66)
                                                        case -460:
                                                            if true {
                                                                v66 = (v66 | v66)
                                                                var v75 uint32 = 2
                                                                var v76 int32 = v11
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                        }
                                                        _ = v66
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(fi61)
                                                } else {
                                                    v1 = ((v1 * v1) / 3.14159)
                                                    switch "" {
                                                    case "a":
                                                        if len(v3) > 0 {
                                                            v3[0] = 32768
                                                        }
                                                    }
                                                    v3 = append(v3, fi58)
                                                }
                                                if true {
                                                    v77 := GenS7{F0: fi61, F1: 4294967295}
                                                    _ = v77
                                                }
                                                _ = fi61
                                            }
                                        }
                                        _ = v59
                                    }
                                    v11 += v11
                                }
                                switch "foo bar" {
                                case "":
                                    v78 := GenS9{F0: 32767, F1: "\n\t", F2: 100}
                                    _ = v78
                                case "\n\t":
                                    var v79 int = (fi58 & fi58)
                                    _ = v79
                                }
                                _ = fi58
                            }
                        case -499:
                            var v80 int = -392
                            _ = v80
                        case -9223372036854775808:
                            var v81 bool = (!(!((2147483647 <= -644) || true)))
                            _ = v81
                        }
                        v11 = -1
                        _ = v57
                    }
                } else {
                    v1 += (v1 + 0.1)
                    switch ("" + ("" + "hello")) {
                    case "hello":
                        for fi82 := 0; fi82 < 1; fi82++ {
                            v3 = append(v3, fi82)
                            _ = fi82
                        }
                    case "":
                        if true {
                            var v83 int32 = v11
                            _ = v83
                        } else {
                            fmt.Println(v11, v1)
                            fmt.Println(v11, v1)
                        }
                    case "foo bar":
                        v84 := GenS4{F0: 1, F1: 127, F2: -1}
                        _ = v84
                    default:
                        if (9223372036854775807 != -2147483648) {
                            fmt.Println(v11, v1)
                        } else {
                            if false {
                                var v85 int64 = 42
                                _ = v85
                            } else {
                                if ((true && false) && (!(18446744073709551615 > 65535))) {
                                    var v86 int8 = -10
                                    _ = v86
                                } else {
                                    fmt.Println(v11, v11)
                                }
                                v11 += (v11 - 1)
                                for fi87 := 0; fi87 < 18; fi87++ {
                                    var v88 int64 = 42
                                    _ = fi87
                                    _ = v88
                                }
                            }
                            _ = len(v3)
                            if (1.5 <= 0.0) {
                                v89 := GenS10{F0: v11, F1: 392}
                                _ = v89
                            }
                        }
                    }
                }
                var v90 int8 = 127
                _ = v90
            } else {
                var v91 string = "hello"
                var v92 int = -2147483648
                _ = v91
                _ = v92
            }
            _ = v46
        }
        for fi93 := 0; fi93 < 15; fi93++ {
            if (!("a" >= "\n\t")) {
                if (!true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v2)
                } else {
                    v11 = 1
                }
            }
            var v94 int8 = 100
            _ = fi93
            _ = v94
        }
        v1 = v1
    }
    var v95 int64 = 969
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v40
    _ = v9
    _ = v95
}
