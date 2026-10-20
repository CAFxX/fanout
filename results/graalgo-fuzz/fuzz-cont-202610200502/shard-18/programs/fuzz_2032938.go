package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS10 struct { F0 int64; F1 int16; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 string = "\n\t"
    var v3 int = -9223372036854775808
    v4 := make(map[string]bool)
    v5 := []int{(-(v3))}
    for fi6 := 0; fi6 < 2; fi6++ {
        if (false || false) {
            var v7 uint32 = 2
            fmt.Println(v7, fi6, v3)
            _ = v7
        } else {
            delete(v1, (fi6 | (fi6 % ((v3) | 1))))
            if true {
                v4[v2] = ((173 < 42) || true)
            }
            var v8 int8 = 1
            _ = v8
        }
        if false {
            if true {
                fmt.Println(v2, fi6)
            }
        } else {
            var v9 uint16 = 32767
            _ = len(v1)
            var v10 float32 = 3.14159
            _ = v10
            _ = v9
        }
        if (false || false) {
            if true {
                if false {
                    v2 = v2
                    v11 := GenS1{F0: 42, F1: 0}
                    _ = v11
                } else {
                    if true {
                        for fi12 := 0; fi12 < 17; fi12++ {
                            var v13 float64 = 1e10
                            if false {
                                if (((2147483647 >= 0) || false) || false) {
                                    v13 = 3.14159
                                    v13 = (v13 + v13)
                                    var v14 uint = 0
                                    _ = v14
                                }
                            }
                            _ = fi12
                            _ = v13
                        }
                        v15 := GenS1{F0: 32767, F1: 42}
                        v2 = v2
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            if (312 <= 65535) {
                                v2 = (("\n\t" + "a") + v2)
                            } else {
                                fmt.Println(v3, fi16)
                                v2 = ""
                            }
                            _ = fi16
                        }
                        if false {
                            fmt.Println(v3, v3)
                            v17 := GenS1{F0: 65535, F1: 2}
                            if (false && (!false)) {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    if len(v5) > 0 {
                                        v5[0] = (fi18 << ((255) & 63))
                                    }
                                    fmt.Println(v3)
                                    _ = fi18
                                }
                                if (!true) {
                                    var v19 int = (v3 >> ((2) & 63))
                                    for fi20 := 0; fi20 < 4; fi20++ {
                                        fmt.Println(v19, v3)
                                        var v21 bool = (!false)
                                        _ = fi20
                                        _ = v21
                                    }
                                    _ = v19
                                }
                                v3 += (^(v3))
                            } else {
                                v1[2147483647] = v2
                                if true {
                                    v22 := GenS1{F0: 0, F1: 128}
                                    if false {
                                        v3 = v3
                                        if true {
                                            v3 = (fi6 + (fi6 % (((v3 * 202)) | 1)))
                                            var v23 string = v2
                                            v3 += (v3 << ((18446744073709551615) & 63))
                                            _ = v23
                                        }
                                        _ = len(v4)
                                    }
                                    switch fi6 {
                                    case -99:
                                        for fi24 := 0; fi24 < 10; fi24++ {
                                            delete(v1, v3)
                                            v3 = 32767
                                            var v25 string = v2
                                            switch v25 {
                                            case "foo bar":
                                                if len(v5) > 0 {
                                                    v5[0] = v3
                                                }
                                            }
                                            var v26 uint64 = 0
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                    default:
                                        v3 = (fi6 >> ((1) & 63))
                                    }
                                    _ = v22
                                } else {
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        if false {
                                            var v28 int = (fi6 | fi6)
                                            fmt.Println(fi27, fi6)
                                            switch v2 {
                                            case "":
                                                v3 = (fi6 | (v3 % ((fi27) | 1)))
                                            case "foo bar":
                                                var v29 int64 = 42
                                                _ = v29
                                            case "hello":
                                                switch (v3 % (((fi27 ^ (-(fi27)))) | 1)) {
                                                case 310:
                                                    v30 := GenS2{F0: "\n\t", F1: -1, F2: ((v2 + v2) > ((((("foo bar" + (v2 + (v2 + v2))) + ("\n\t" + ("a" + v2))) + ("a" + "")) + v2) + ((((v2 + "a") + (("hello" + ("foo bar" + "")) + "a")) + ("hello" + v2)) + v2)))}
                                                    _ = v30
                                                default:
                                                    if len(v5) > 0 {
                                                        v5[0] = 1
                                                    }
                                                    if (0.1 >= (0.1 / ((3.14159 + (1e10 + -2.25)) / 1.5))) {
                                                        v28 += -533
                                                    }
                                                    v28 = fi27
                                                }
                                            }
                                            var v31 uint32 = 407
                                            var v32 uint64 = 4294967295
                                            v3 = v3
                                            v28 = int(v31)
                                            _ = v28
                                            _ = v31
                                            _ = v32
                                        } else {
                                            var v33 int64 = 42
                                            if (false && false) {
                                                var v34 string = (v2 + v2)
                                                _ = v34
                                            } else {
                                                v2 = (v2 + v2)
                                                if ("hello" != ("foo bar" + "")) {
                                                    var v35 int16 = -109
                                                    var v36 uint32 = 262
                                                    _ = v35
                                                    _ = v36
                                                }
                                                fmt.Println(v2)
                                            }
                                            if false {
                                                delete(v4, v2)
                                                var v37 int = (fi27 + fi6)
                                                v33 = (-(v33))
                                                _ = v37
                                            } else {
                                                var v38 int8 = 38
                                                switch 127 {
                                                case -1:
                                                    v39 := GenS1{F0: 255, F1: 127}
                                                    _ = v39
                                                case 42:
                                                    switch "" {
                                                    case "":
                                                        switch ("" + "a") {
                                                        case "foo bar":
                                                            v40 := GenS1{F0: 32767, F1: 184}
                                                            _ = v40
                                                        case "a":
                                                            var v41 float32 = -2.25
                                                            v42 := GenS3{F0: 18446744073709551615}
                                                            var v43 uint16 = 759
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        case "hello":
                                                            switch ((v2 + "hello") + "") {
                                                            case "hello":
                                                                _ = len(v5)
                                                            }
                                                        }
                                                        var v44 uint = 42
                                                        v44 = (^(v44))
                                                        v45 := GenS4{F0: 4294967295, F1: 1}
                                                        _ = len(v5)
                                                        v46 := GenS4{F0: 2147483647, F1: 243}
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    case "a":
                                                        var v47 uint8 = 42
                                                        _ = v47
                                                    case "\n\t":
                                                        var v48 uint = 65535
                                                        _ = v48
                                                    default:
                                                        var v49 uint8 = 42
                                                        _ = v49
                                                    }
                                                    if false {
                                                        v33 += v33
                                                        switch v3 {
                                                        case 1:
                                                            _ = len(v1)
                                                        case 127:
                                                            v3 = fi6
                                                        case -2147483648:
                                                            switch (v3 / (((v3 ^ 9223372036854775807)) | 1)) {
                                                            case 9223372036854775807:
                                                                fmt.Println(fi6, v2)
                                                            case -17:
                                                                v50 := GenS3{F0: 512}
                                                                _ = v50
                                                            case 552:
                                                                switch ((v2 + v2) + v2) {
                                                                case "hello":
                                                                    if true {
                                                                        var v51 int32 = 184
                                                                        var v52 int16 = 0
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                case "a":
                                                                    v38 += v38
                                                                }
                                                            default:
                                                                switch "a" {
                                                                case "":
                                                                    var v53 uint8 = 128
                                                                    _ = v53
                                                                case "a":
                                                                    switch (fi27 / (((fi27 << ((1) & 63))) | 1)) {
                                                                    case 0:
                                                                        switch ("foo bar" + (v2 + (("\n\t" + v2) + v2))) {
                                                                        case "hello":
                                                                            var v54 int32 = 1
                                                                            _ = v54
                                                                        case "":
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if ((float64(0.1) + ((1.5 * float64(float32(0.1))) + 0.1)) < 1e10) {
                                                                            fmt.Println(v3, v2)
                                                                            v55 := GenS4{F0: 1, F1: 1}
                                                                            _ = v55
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            v33 -= 0
                                                                            if true {
                                                                                if ((v3 << ((1) & 63)) != -2147483648) {
                                                                                    switch v2 {
                                                                                    case "hello":
                                                                                        fmt.Println(v3)
                                                                                    case "foo bar":
                                                                                        v33 -= v33
                                                                                    }
                                                                                    if (false || false) {
                                                                                        switch 273 {
                                                                                        case 491:
                                                                                            switch int(v33) {
                                                                                            case -2147483648:
                                                                                                var v56 int = fi27
                                                                                                _ = v56
                                                                                            case 127:
                                                                                                if ((442 <= 200) || false) {
                                                                                                    var v57 int8 = 0
                                                                                                    _ = v57
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v38, v3)
                                                                                        }
                                                                                        if false {
                                                                                            v58 := GenS5{F0: v2, F1: 127, F2: 1}
                                                                                            if true {
                                                                                                fmt.Println(v3)
                                                                                                if (255 == 402) {
                                                                                                    switch v2 {
                                                                                                    case "\n\t":
                                                                                                        v38 += 127
                                                                                                    default:
                                                                                                        _ = len(v5)
                                                                                                        switch (fi27 & -753) {
                                                                                                        case 1:
                                                                                                            v59 := GenS5{F0: (v2 + "foo bar"), F1: 127, F2: 876}
                                                                                                            _ = v59
                                                                                                        case 0:
                                                                                                            v38 += 42
                                                                                                        }
                                                                                                    }
                                                                                                    var v60 uint32 = 47
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    v3 -= v3
                                                                                                }
                                                                                                var v61 int32 = 1
                                                                                                _ = v61
                                                                                            } else {
                                                                                                v62 := GenS4{F0: 65535, F1: 42}
                                                                                                v2 = v2
                                                                                                _ = v62
                                                                                            }
                                                                                            var v63 string = (v2 + (v2 + v2))
                                                                                            var v64 int16 = 410
                                                                                            if (true && (0 == v33)) {
                                                                                                v38 -= (v38 >> ((18446744073709551615) & 63))
                                                                                                v65 := GenS6{F0: "foo bar"}
                                                                                                _ = v65
                                                                                            }
                                                                                            _ = v58
                                                                                            _ = v63
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi6, v33)
                                                                                var v66 float32 = float32((-2.25 * (3.14159 * (0.1 * 0.1))))
                                                                                _ = v66
                                                                            }
                                                                        }
                                                                    case 127:
                                                                        v67 := GenS5{F0: v2, F1: 42, F2: 42}
                                                                        _ = v67
                                                                    default:
                                                                        v5 = append(v5, fi6)
                                                                        fmt.Println(v33)
                                                                        if false {
                                                                            if true {
                                                                                if true {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (fi6 % (((^(fi27))) | 1))
                                                                                    }
                                                                                } else {
                                                                                    var v68 bool = true
                                                                                    _ = v68
                                                                                }
                                                                                switch (fi6 % (((v3 | fi6)) | 1)) {
                                                                                case 32767:
                                                                                    var v69 int16 = 864
                                                                                    _ = v69
                                                                                case 157:
                                                                                    fmt.Println(v2, fi27, v33)
                                                                                default:
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                }
                                                                                v70 := GenS7{F0: 42}
                                                                                _ = v70
                                                                            }
                                                                            fmt.Println(fi27, v2, v2)
                                                                        } else {
                                                                            var v71 uint8 = 252
                                                                            v3 -= (fi27 & (-(v3)))
                                                                            v72 := GenS8{F0: v3}
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    if (127 == 2) {
                                                                        if true {
                                                                            v2 = v2
                                                                            v2 = v2
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    v3 -= (fi6 % ((-545) | 1))
                                                }
                                                v2 = v2
                                                _ = v38
                                            }
                                            _ = v33
                                        }
                                        v73 := GenS1{F0: 255, F1: 78}
                                        var v74 uint16 = 65535
                                        _ = fi27
                                        _ = v73
                                        _ = v74
                                    }
                                    switch v2 {
                                    case "foo bar":
                                        switch v2 {
                                        case "hello":
                                            var v75 bool = false
                                            _ = v75
                                        case "":
                                            v2 = "a"
                                        }
                                    case "":
                                        var v76 uint16 = 2
                                        _ = v76
                                    case "a":
                                        v77 := GenS4{F0: 4294967295, F1: 42}
                                        _ = v77
                                    }
                                    fmt.Println(v3)
                                }
                            }
                            _ = v17
                        } else {
                            v3 = v3
                            v78 := GenS9{F0: 331, F1: 1, F2: 1000}
                            v79 := GenS3{F0: 0}
                            _ = v78
                            _ = v79
                        }
                    }
                    for fi80 := 0; fi80 < 17; fi80++ {
                        v3 = 42
                        v4["\n\t"] = (-291 <= -32768)
                        _ = fi80
                    }
                }
                v3 += 274
            }
            v3 += -550
            _ = len(v5)
        } else {
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (9223372036854775807 < -52) {
                    var v81 bool = true
                    _ = v81
                }
            }
            v82 := GenS8{F0: 32768}
            var v83 uint = 1
            _ = v82
            _ = v83
        }
        _ = fi6
    }
    v2 = (v2 + (v2 + (v2 + (v2 + (v2 + (("\n\t" + v2) + "\n\t"))))))
    v84 := GenS9{F0: 42, F1: 32767, F2: -400}
    _ = len(v4)
    v85 := v1[(v3 / ((v3) | 1))]
    v85 = v2
    if (!(false || (2 == 0))) {
        if true {
            if false {
                v86 := v4[v2]
                _ = v86
            } else {
                var v87 int = v3
                _ = len(v4)
                v3 += v87
                _ = v87
            }
        }
        var v88 uint32 = 431
        _ = v88
    }
    v89 := GenS10{F0: -9223372036854775808, F1: 1000, F2: v3}
    if len(v5) > 0 {
        v5[0] = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v89
}
