package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 int16; F1 int; F2 int8 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS8 struct { F0 int64; F1 int; F2 uint }
type GenS9 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS10 struct { F0 uint32; F1 uint; F2 uint }
type GenS11 struct { F0 uint; F1 int; F2 int }
type GenS12 struct { F0 int8 }
func genFn0() bool {
    var v7 uint = 133
    fmt.Println(v7)
    fmt.Println(v7)
    _ = v7
    return (!(2 <= uint32(v7)))
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{"hello"}
    var v3 string = ("\n\t" + "")
    v4 := []string{v3, v3}
    var v5 uint8 = 127
    var v6 uint64 = 4294967295
    for fi8 := 0; fi8 < 4; fi8++ {
        v1 = 3.14159
        if (!true) {
            if false {
                if (float64((((0.1 * (v1 * ((0.0 - float32(float64(v1))) - -2.25))) * (((3.14159 * (v1 / 1.5)) - (((float32(1.5) + (v1 + (1e10 * -2.25))) / 1.0) / float32(1.5))) - 0.0)) * (v1 * v1))) != float64(((v1 / v1) / v1))) {
                    v4 = append(v4, (("" + v3) + ("\n\t" + v3)))
                } else {
                    v3 = v3
                    var v9 bool = false
                    switch (v3 + ("foo bar" + (v3 + ""))) {
                    case "a":
                        v1 = (v1 + (v1 + (v1 - 1e10)))
                    }
                    _ = v9
                }
                fmt.Println(v3, fi8)
                if false {
                    var v10 string = (("\n\t" + v3) + "\n\t")
                    var v11 float32 = v1
                    _ = v10
                    _ = v11
                } else {
                    v6 += (^(v6))
                    var v12 string = (((v3 + v3) + "hello") + v3)
                    _ = v12
                }
            } else {
                if (false || (true && ((757 >= int16(v6)) || true))) {
                    v1 = v1
                }
            }
            if ((!true) && false) {
                var v13 string = v3
                fmt.Println(fi8, v1)
                _ = v13
            } else {
                v1 = 0.1
                if false {
                    _ = len(v4)
                }
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: v6}
                var v16 uint8 = (v5 * (^(v5)))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if true {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 uint32 = 42
                v20 := GenS1{F0: (v6 % ((4294967295) | 1))}
                v21 := GenS2{F0: 778}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            switch (fi17 * fi17) {
            case 494:
                for fi22 := 0; fi22 < 11; fi22++ {
                    v6 -= (v6 * 690)
                    _ = fi22
                }
            case 9223372036854775807:
                switch v3 {
                case "foo bar":
                    fmt.Println(v5)
                case "hello":
                    switch v3 {
                    case "hello":
                        if true {
                            var v23 uint32 = 2147483647
                            if (!false) {
                                fmt.Println(v6)
                                var v24 int32 = 42
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    var v26 int64 = 567
                                    v6 += (v6 | (v6 % ((v6) | 1)))
                                    v5 = (^(v5))
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            } else {
                                switch (fi17 / (((-(fi17))) | 1)) {
                                case 127:
                                    for fi27 := 0; fi27 < 4; fi27++ {
                                        v1 += v1
                                        v28 := GenS2{F0: 483}
                                        if ((v5 >> ((65535) & 63)) <= v5) {
                                            v29 := GenS3{F0: (v6 - (v6 | (v6 * 2)))}
                                            if len(v2) > 0 {
                                                v2[0] = (v3 + (("\n\t" + "a") + v3))
                                            }
                                            if (283 == 255) {
                                                switch (v3 + "foo bar") {
                                                case "foo bar":
                                                    fmt.Println(v1, v1)
                                                case "\n\t":
                                                    switch v3 {
                                                    case "hello":
                                                        if len(v2) > 0 {
                                                            v2[0] = v3
                                                        }
                                                    case "":
                                                        v3 = "\n\t"
                                                    }
                                                case "a":
                                                    switch fi17 {
                                                    case 192:
                                                        v30 := GenS4{F0: (fi17 * (fi17 + fi17)), F1: (("\n\t" + v3) + "hello"), F2: 2}
                                                        _ = v30
                                                    case -939:
                                                        var v31 uint8 = (v5 & 127)
                                                        _ = v31
                                                    }
                                                default:
                                                    switch 2147483647 {
                                                    case 32768:
                                                        var v32 float64 = (-2.25 * (1e10 / ((-2.25 - 1e10) + 0.0)))
                                                        _ = v32
                                                    case -1:
                                                        _ = len(v2)
                                                    default:
                                                        fmt.Println(fi17)
                                                    }
                                                }
                                            }
                                            v3 = "hello"
                                            switch (fi17 * 2147483647) {
                                            case 4294967295:
                                                if (!(!false)) {
                                                    var v33 float32 = 0.0
                                                    _ = v33
                                                } else {
                                                    v2 = append(v2, v3)
                                                    fmt.Println(v1, fi17)
                                                }
                                            case 117:
                                                if ((3.14159 + v1) < float32(0.1)) {
                                                    v2 = append(v2, v3)
                                                    v3 = v3
                                                }
                                            default:
                                                var v34 uint8 = 255
                                                _ = v34
                                            }
                                            v23 += 61
                                            if (139 == 0) {
                                                var v35 string = "a"
                                                v35 = ("" + (v3 + "foo bar"))
                                                _ = v35
                                            } else {
                                                v4 = append(v4, "\n\t")
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v36 int16 = 32767
                                                v37 := GenS5{F0: (v36 << ((1) & 63)), F1: fi17, F2: -128}
                                                var v38 uint32 = (^(v23))
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = v29
                                        }
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                                v2 = append(v2, "\n\t")
                                switch 115 {
                                case 2147483647:
                                    var v39 uint32 = v23
                                    _ = v39
                                case 552:
                                    var v40 int = int(v5)
                                    _ = v40
                                case 42:
                                    fmt.Println(v1, v3, v23)
                                }
                            }
                            v6 -= 39
                            _ = v23
                        }
                    case "foo bar":
                        if (!(4294967295 > 255)) {
                            if ((((v1 + ((0.1 - ((0.0 / 1.0) / (v1 / float32(3.14159)))) - (v1 * 1.5))) - (float32(0.0) / 1.0)) + v1) != 3.14159) {
                                v3 = (v3 + v3)
                                v5 -= v5
                            }
                            if (-2147483648 > 262) {
                                if true {
                                    for fi41 := 0; fi41 < 12; fi41++ {
                                        v42 := GenS6{F0: -1, F1: 571}
                                        v5 += v5
                                        var v43 float64 = (float64(0.1) * (float64(v1) / 1.5))
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                fmt.Println(v1, v3, v5)
                            }
                            var v44 uint8 = uint8(v6)
                            _ = v44
                        }
                    }
                }
            case 125:
                v45 := GenS3{F0: (^(v6))}
                _ = v45
            }
            fmt.Println(v1)
        }
        _ = fi17
    }
    var v46 int = 128
    for fi47 := 0; fi47 < 18; fi47++ {
        v2 = append(v2, v3)
        switch (fi47 | v46) {
        case 1:
            v4 = append(v4, "a")
        case -1:
            var v48 int32 = -573
            _ = v48
        default:
            if ((!true) && (false || false)) {
                var v49 float64 = float64(v1)
                var v50 bool = false
                for fi51 := 0; fi51 < 9; fi51++ {
                    if v50 {
                        v49 = (v49 / v49)
                        v2 = append(v2, (((v3 + "\n\t") + ("a" + v3)) + v3))
                        v6 += 18446744073709551615
                    } else {
                        var v52 uint16 = 865
                        v53 := GenS3{F0: (v6 * 0)}
                        switch (fi51 + (-(fi51))) {
                        case 195:
                            var v54 float64 = v49
                            _ = v54
                        case -15:
                            if (false || v50) {
                                v49 = float64(3.14159)
                                v4 = append(v4, (v3 + "foo bar"))
                                v52 = v52
                            }
                        case -213:
                            var v55 uint = 404
                            _ = v55
                        }
                        fmt.Println(v1, v52, v50)
                        _ = v52
                        _ = v53
                    }
                    v3 = "a"
                    fmt.Println(v50, v46, v46)
                    _ = len(v2)
                    _ = fi51
                }
                _ = v49
                _ = v50
            } else {
                if (!(-32768 > 1000)) {
                    if (false && (v46 >= (v46 / ((-139) | 1)))) {
                        v56 := GenS7{F0: 0, F1: 479, F2: -41}
                        var v57 int8 = 0
                        var v58 string = v3
                        _ = v56
                        _ = v57
                        _ = v58
                    } else {
                        _ = len(v2)
                        switch (fi47 >> ((18446744073709551615) & 63)) {
                        case -438:
                            if false {
                                for fi59 := 0; fi59 < 5; fi59++ {
                                    v60 := GenS8{F0: 1, F1: fi59, F2: 538}
                                    _ = fi59
                                    _ = v60
                                }
                            }
                        case -9223372036854775808:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            fmt.Println(fi47, v3)
                        }
                        v6 -= (v6 % (((v6 * v6)) | 1))
                    }
                } else {
                    if ((v46 | fi47) >= fi47) {
                        v61 := GenS9{F0: 17, F1: (v5 % ((v5) | 1)), F2: 1}
                        var v62 uint16 = 829
                        v63 := GenS10{F0: 685, F1: 18446744073709551615, F2: 1}
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    if ((true || ((fi47 | (fi47 & 10)) != fi47)) && false) {
                        v3 = (((v3 + (v3 + v3)) + v3) + v3)
                        if true {
                            switch (v46 % ((v46) | 1)) {
                            case 128:
                                fmt.Println(v5, fi47)
                            case 860:
                                fmt.Println(v5, fi47)
                            }
                        }
                        var v64 uint32 = 2147483648
                        _ = v64
                    } else {
                        v1 -= 0.1
                        v3 = v3
                        v46 = (fi47 - 42)
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (("a" + (v3 + (v3 + v3))) + (((v3 + v3) + "\n\t") + v3)) {
                case "a":
                    for fi65 := 0; fi65 < 5; fi65++ {
                        var v66 uint64 = v6
                        var v67 uint8 = 2
                        _ = len(v2)
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                case "\n\t":
                    v5 = v5
                case "foo bar":
                    for fi68 := 0; fi68 < 2; fi68++ {
                        v6 = v6
                        _ = fi68
                    }
                }
            }
        }
        _ = fi47
    }
    var v69 string = ""
    switch (("hello" + "") + v69) {
    case "hello":
        v70 := GenS1{F0: v6}
        _ = v70
    case "":
        fmt.Println(v1, v6)
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        if true {
            v5 -= v5
        } else {
            v72 := GenS8{F0: 42, F1: (v46 % ((v46) | 1)), F2: 65535}
            var v73 float64 = (3.14159 - -2.25)
            var v74 int64 = 982
            _ = v72
            _ = v73
            _ = v74
        }
        v75 := GenS6{F0: -1, F1: 65535}
        fmt.Println(v1)
        _ = fi71
        _ = v75
    }
    v76 := GenS7{F0: 4294967295, F1: 65535, F2: 0}
    v77 := GenS9{F0: v5, F1: 127, F2: 1}
    var v78 int64 = -9223372036854775808
    v2 = append(v2, v3)
    switch v3 {
    case "foo bar":
        v79 := GenS9{F0: (^(v5)), F1: (v5 << ((255) & 63)), F2: (v78 / ((v78) | 1))}
        _ = v79
    case "":
        v80 := GenS9{F0: 2, F1: v5, F2: v78}
        _ = v80
    }
    v81 := GenS11{F0: 415, F1: (v46 >> ((1) & 63)), F2: 2147483647}
    switch (-(v46)) {
    case 32767:
        switch v46 {
        case 127:
            v82 := GenS9{F0: v5, F1: 2, F2: v78}
            _ = v82
        }
    case -362:
        if true {
            if true {
                if true {
                    if false {
                        v69 = v69
                        for fi83 := 0; fi83 < 19; fi83++ {
                            fmt.Println(v46)
                            switch v3 {
                            case "\n\t":
                                switch (v3 + (((v3 + ((("\n\t" + (v69 + v3)) + "foo bar") + "")) + ("\n\t" + v3)) + "")) {
                                case "":
                                    for fi84 := 0; fi84 < 8; fi84++ {
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        v3 = v3
                                        _ = fi84
                                    }
                                default:
                                    var v85 string = "hello"
                                    _ = v85
                                }
                            }
                            _ = fi83
                        }
                    } else {
                        v69 = ""
                        v78 -= -1
                        v6 += uint64(v5)
                    }
                    fmt.Println(v78, v3)
                    v3 = ((((v69 + v3) + ((v69 + "foo bar") + ((v3 + "hello") + (v69 + "foo bar")))) + v3) + ((v69 + v69) + v69))
                } else {
                    var v86 string = "a"
                    fmt.Println(v86)
                    if (((!false) || (0.0 >= 3.14159)) || false) {
                        for fi87 := 0; fi87 < 14; fi87++ {
                            var v88 int64 = (v78 << ((uint(v6)) & 63))
                            _ = fi87
                            _ = v88
                        }
                        switch v86 {
                        case "":
                            var v89 string = v86
                            _ = v89
                        case "\n\t":
                            var v90 float32 = 0.1
                            _ = v90
                        default:
                            fmt.Println(v6, v69, v46)
                        }
                    } else {
                        for fi91 := 0; fi91 < 10; fi91++ {
                            switch (v46 % ((192) | 1)) {
                            case 1:
                                if (((!false) || (v46 >= (fi91 + (fi91 % ((fi91) | 1))))) || false) {
                                    v2 = append(v2, "foo bar")
                                } else {
                                    var v92 float32 = 0.0
                                    _ = v92
                                }
                            case 42:
                                for fi93 := 0; fi93 < 9; fi93++ {
                                    v94 := GenS1{F0: uint64(v5)}
                                    _ = fi93
                                    _ = v94
                                }
                            case 32768:
                                var v95 int64 = (v78 | 42)
                                _ = v95
                            default:
                                if (4294967295 >= 467) {
                                    var v96 uint = 2
                                    _ = v96
                                }
                            }
                            switch -971 {
                            case -9223372036854775808:
                                _ = len(v2)
                            case 82:
                                v1 = v1
                            case 1:
                                _ = len(v2)
                            }
                            _ = fi91
                        }
                        var v97 uint8 = v5
                        if (true || true) {
                            if true {
                                fmt.Println(v5, v6, v3)
                            }
                        } else {
                            switch ((("" + ("\n\t" + "\n\t")) + v69) + "hello") {
                            case "foo bar":
                                fmt.Println(v5, v86)
                            }
                        }
                        _ = v97
                    }
                    _ = v86
                }
                var v98 string = (v69 + v3)
                _ = v98
            } else {
                switch "hello" {
                case "\n\t":
                    var v99 float64 = -2.25
                    _ = v99
                case "foo bar":
                    var v100 float32 = v1
                    _ = v100
                default:
                    for fi101 := 0; fi101 < 10; fi101++ {
                        if false {
                            var v102 uint32 = 65535
                            v4 = append(v4, ("foo bar" + v69))
                            if (true && false) {
                                var v103 int32 = -1
                                var v104 uint64 = 42
                                _ = v103
                                _ = v104
                            }
                            _ = v102
                        }
                        v105 := GenS7{F0: 713, F1: 569, F2: 42}
                        _ = fi101
                        _ = v105
                    }
                }
            }
            var v106 uint16 = 32767
            _ = v106
        }
    }
    var v107 string = "a"
    if (193 != v5) {
        var v108 uint8 = (v5 % ((v5) | 1))
        v2 = append(v2, ((("" + v3) + v69) + ""))
        v6 = 4294967295
        _ = v108
    } else {
        v109 := GenS12{F0: 1}
        _ = v109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v69
    _ = v76
    _ = v77
    _ = v78
    _ = v81
}
