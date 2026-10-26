package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint64; F2 uint }
type GenS7 struct { F0 uint64; F1 int64; F2 bool }
type GenS8 struct { F0 int8 }

func main() {
    var v1 uint16 = 543
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 uint32 = 4294967295
    v6 := []string{"foo bar", "foo bar", "foo bar"}
    var v7 float64 = (v2 * (1.5 / v2))
    v8 := GenS1{F0: 542}
    var v9 int64 = 42
    v1 -= (v1 ^ (v1 * 1))
    if (true || (1e10 == 0.0)) {
        var v10 int64 = 42
        if (127 < 1) {
            v11 := v3["hello"]
            for fi12 := 0; fi12 < 15; fi12++ {
                v10 = v10
                var v13 float64 = (v2 - v2)
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 127}
            _ = v11
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            switch "a" {
            case "":
                if true {
                    fmt.Println(fi15, fi15)
                    if true {
                        _ = len(v3)
                    }
                }
            case "a":
                _ = len(v6)
            case "hello":
                fmt.Println(fi15, v7, v1)
            }
            switch "" {
            case "foo bar":
                v16 := GenS2{F0: 228, F1: 1, F2: fi15}
                _ = v16
            default:
                if false {
                    v5 = (v5 ^ v5)
                    switch (("a" + "foo bar") + "foo bar") {
                    case "\n\t":
                        v7 -= float64(0.0)
                    default:
                        v2 += v2
                    }
                }
            }
            switch "\n\t" {
            case "hello":
                v17 := GenS1{F0: (fi15 % ((32767) | 1))}
                _ = v17
            default:
                if true {
                    switch fi15 {
                    case -918:
                        fmt.Println(v10)
                    case -9223372036854775808:
                        v18 := v4["foo bar"]
                        _ = v18
                    case -1:
                        var v19 int = (fi15 >> ((255) & 63))
                        _ = v19
                    default:
                        var v20 uint8 = 219
                        _ = v20
                    }
                    v21 := GenS2{F0: 1, F1: 208, F2: (fi15 | (fi15 - 4294967295))}
                    _ = v21
                } else {
                    var v22 uint64 = 4294967295
                    var v23 int8 = 87
                    _ = v22
                    _ = v23
                }
            }
            _ = fi15
        }
        _ = v10
    }
    var v24 float64 = v7
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS2{F0: 2, F1: 1, F2: (^(fi25))}
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: 0, F1: 756}
    var v28 bool = false
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 string = ""
        switch (fi29 / ((970) | 1)) {
        case 1:
            v30 = v30
        case -1:
            if v28 {
                if (!v28) {
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v32 := GenS4{F0: -2147483648, F1: 951}
                        v9 = (v9 - 1)
                        _ = fi31
                        _ = v32
                    }
                    _ = len(v4)
                } else {
                    var v33 int8 = 0
                    fmt.Println(v30, v28)
                    if v28 {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v35 := GenS1{F0: 128}
                            var v36 string = (v30 + (v30 + ""))
                            var v37 uint64 = uint64(v1)
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS4{F0: 1, F1: 2}
                        v39 := GenS1{F0: fi29}
                        _ = v38
                        _ = v39
                    } else {
                        v24 -= (0.0 + v7)
                        v28 = (0 > 42)
                    }
                    _ = v33
                }
                var v40 bool = (float64(1e10) > v7)
                for fi41 := 0; fi41 < 20; fi41++ {
                    v5 = v5
                    var v42 int8 = 1
                    if v28 {
                        var v43 uint32 = (^(v5))
                        var v44 uint = 4294967295
                        _ = v43
                        _ = v44
                    } else {
                        v45 := GenS2{F0: 42, F1: 127, F2: fi41}
                        var v46 uint16 = (v1 - v1)
                        switch (("" + "a") + (("\n\t" + "\n\t") + (v30 + ""))) {
                        case "\n\t":
                            if (-1 > -2147483648) {
                                if v40 {
                                    switch v30 {
                                    case "hello":
                                        var v47 int = 128
                                        _ = v47
                                    case "a":
                                        v1 = (v1 & (v46 * 65535))
                                    default:
                                        _ = len(v6)
                                    }
                                    if (v40 && v40) {
                                        switch fi29 {
                                        case -2147483648:
                                            v9 += (v9 - 300)
                                        case 32767:
                                            switch v30 {
                                            case "a":
                                                v2 -= v2
                                            }
                                        }
                                        v48 := GenS5{F0: 32768}
                                        _ = v48
                                    } else {
                                        if (127 > fi29) {
                                            v2 = v24
                                            delete(v4, v30)
                                        }
                                        if v40 {
                                            if v40 {
                                                if ((^(v42)) > (v42 >> ((734) & 63))) {
                                                    var v49 string = v30
                                                    _ = v49
                                                } else {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v5 = v5
                                                    case "a":
                                                        v30 = ("\n\t" + v30)
                                                    case "foo bar":
                                                        switch -546 {
                                                        case 0:
                                                            fmt.Println(v30)
                                                        case 128:
                                                            v50 := GenS3{F0: 255, F1: 255}
                                                            _ = v50
                                                        default:
                                                            if true {
                                                                v51 := GenS2{F0: 42, F1: 66, F2: fi29}
                                                                v5 = (^(v5))
                                                                _ = v51
                                                            } else {
                                                                if true {
                                                                    var v52 uint = 42
                                                                    var v53 int16 = -443
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                if v40 {
                                                                    if (!(float32(float64((-2.25 * ((3.14159 - float32(v2)) - (0.0 - 1e10))))) == 0.0)) {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        if v28 {
                                                                            switch fi29 {
                                                                            case 42:
                                                                                if false {
                                                                                    var v54 int32 = -178
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        var v55 uint8 = 45
                                                                                        _ = v55
                                                                                    case "a":
                                                                                        v40 = true
                                                                                    case "":
                                                                                        v7 -= (1e10 - 0.0)
                                                                                    }
                                                                                    fmt.Println(v46, v7, v24)
                                                                                    _ = v54
                                                                                } else {
                                                                                    if (!true) {
                                                                                        v56 := GenS5{F0: -1}
                                                                                        _ = v56
                                                                                    } else {
                                                                                        v57 := GenS6{F0: v42, F1: 746, F2: 2}
                                                                                        v6 = append(v6, v30)
                                                                                        _ = v57
                                                                                    }
                                                                                    if v28 {
                                                                                        switch (v30 + v30) {
                                                                                        case "a":
                                                                                            if ("a" == "a") {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = ("hello" + v30)
                                                                                                }
                                                                                                if (((float32(v24) - -2.25) * float32(-2.25)) <= ((3.14159 - 1e10) + -2.25)) {
                                                                                                    v2 -= float64(-2.25)
                                                                                                    v4[(v30 + "a")] = (v9 <= (v9 & (v9 / ((v9) | 1))))
                                                                                                }
                                                                                            }
                                                                                            var v58 int64 = (v9 << ((65535) & 63))
                                                                                            if (!(486 <= 1)) {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = v30
                                                                                                }
                                                                                            } else {
                                                                                                var v59 int64 = v58
                                                                                                _ = v59
                                                                                            }
                                                                                            _ = v58
                                                                                        case "":
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        switch ((v30 + (v30 + (v30 + ("hello" + ("a" + v30))))) + v30) {
                                                                                        case "":
                                                                                            v60 := GenS2{F0: 108, F1: 0, F2: fi41}
                                                                                            _ = v60
                                                                                        }
                                                                                        switch (v30 + "") {
                                                                                        case "\n\t":
                                                                                            if (9223372036854775807 != (v9 / ((int64(v5)) | 1))) {
                                                                                                v40 = true
                                                                                                switch -100 {
                                                                                                case -2147483648:
                                                                                                    v40 = v40
                                                                                                case 4294967295:
                                                                                                    var v61 uint8 = 0
                                                                                                    _ = v61
                                                                                                case -569:
                                                                                                    if len(v6) > 0 {
                                                                                                        v6[0] = "a"
                                                                                                    }
                                                                                                }
                                                                                                v46 += 849
                                                                                                v1 = (v1 << ((4294967295) & 63))
                                                                                                fmt.Println(v2)
                                                                                            } else {
                                                                                                v62 := GenS1{F0: (fi41 % ((1) | 1))}
                                                                                                v5 = (^(v5))
                                                                                                _ = v62
                                                                                            }
                                                                                        case "a":
                                                                                            var v63 int16 = 0
                                                                                            _ = v63
                                                                                        }
                                                                                        fmt.Println(v28)
                                                                                        delete(v3, v30)
                                                                                    } else {
                                                                                        v64 := GenS7{F0: 561, F1: (v9 ^ v9), F2: (fi41 == (fi29 >> ((280) & 63)))}
                                                                                        var v65 float32 = 3.14159
                                                                                        if v40 {
                                                                                            v30 = "hello"
                                                                                        } else {
                                                                                            if v40 {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = "hello"
                                                                                                }
                                                                                            }
                                                                                            v24 -= (0.1 + (((v2 * v2) / 1.5) * 1.5))
                                                                                        }
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                    var v66 float32 = 0.0
                                                                                    _ = v66
                                                                                }
                                                                                _ = len(v3)
                                                                                if (v40 || true) {
                                                                                    switch (fi29 * (fi29 & fi29)) {
                                                                                    case 127:
                                                                                        v7 += (v7 + v7)
                                                                                    case 9223372036854775807:
                                                                                        var v67 int16 = 400
                                                                                        _ = v67
                                                                                    case -482:
                                                                                        v68 := v3[v30]
                                                                                        _ = v68
                                                                                    default:
                                                                                        if (!(int16(v46) < 845)) {
                                                                                            v3[(v30 + "")] = (^(v5))
                                                                                            delete(v3, v30)
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v69 int32 = 42
                                                                                    fmt.Println(fi29)
                                                                                    var v70 int = (fi41 >> ((0) & 63))
                                                                                    var v71 int32 = v69
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                }
                                                                                v9 += v9
                                                                                var v72 int16 = 0
                                                                                _ = v72
                                                                            case -2147483648:
                                                                                v1 = (^(v46))
                                                                            default:
                                                                                var v73 uint = 4294967295
                                                                                _ = v73
                                                                            }
                                                                            v74 := GenS4{F0: -2147483648, F1: 189}
                                                                            switch (v30 + v30) {
                                                                            case "":
                                                                                v1 -= v1
                                                                            case "\n\t":
                                                                                fmt.Println(v28)
                                                                            case "foo bar":
                                                                                var v75 float64 = v2
                                                                                _ = v75
                                                                            }
                                                                            _ = v74
                                                                        } else {
                                                                            if v28 {
                                                                                v6 = append(v6, "hello")
                                                                                var v76 bool = (v30 < "foo bar")
                                                                                v77 := GenS3{F0: 1, F1: 0}
                                                                                _ = v76
                                                                                _ = v77
                                                                            } else {
                                                                                if v40 {
                                                                                    fmt.Println(v28, fi41)
                                                                                    switch (fi29 % ((fi29) | 1)) {
                                                                                    case 127:
                                                                                        var v78 int16 = 648
                                                                                        _ = v78
                                                                                    case 32767:
                                                                                        v79 := v4[v30]
                                                                                        _ = v79
                                                                                    default:
                                                                                        v80 := GenS2{F0: 1, F1: 204, F2: fi29}
                                                                                        _ = v80
                                                                                    }
                                                                                    v9 -= v9
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                } else {
                                                                                    delete(v3, "")
                                                                                    v1 -= v1
                                                                                }
                                                                            }
                                                                            var v81 int32 = -1
                                                                            fmt.Println(v5, v40)
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                    v82 := GenS8{F0: 1}
                                                                    if (1.5 > 1e10) {
                                                                        var v83 float64 = (0.0 + (v7 + v7))
                                                                        _ = v83
                                                                    } else {
                                                                        v7 = v2
                                                                    }
                                                                    _ = v82
                                                                } else {
                                                                    v84 := GenS2{F0: 42, F1: 42, F2: 986}
                                                                    v24 += 0.0
                                                                    _ = v84
                                                                }
                                                                switch (((v30 + v30) + v30) + "foo bar") {
                                                                case "\n\t":
                                                                    v85 := GenS3{F0: 1, F1: 4294967295}
                                                                    _ = v85
                                                                case "":
                                                                    fmt.Println(v40, v7)
                                                                case "a":
                                                                    v6 = append(v6, (((v30 + (v30 + "")) + ((v30 + v30) + (v30 + v30))) + (v30 + v30)))
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                switch v30 {
                                                case "\n\t":
                                                    fmt.Println(v24, v42)
                                                case "foo bar":
                                                    if (v28 || v40) {
                                                        v40 = (false && v28)
                                                    }
                                                case "a":
                                                    var v86 int16 = 0
                                                    _ = v86
                                                }
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        var v87 int64 = 9223372036854775807
                                        _ = v87
                                    }
                                    v9 -= v9
                                } else {
                                    if v28 {
                                        if (!v28) {
                                            var v88 int8 = int8(v9)
                                            _ = v88
                                        }
                                        var v89 float32 = ((0.1 - 1.5) + -2.25)
                                        _ = v89
                                    }
                                    var v90 bool = v40
                                    _ = v90
                                }
                                _ = len(v4)
                            }
                        default:
                            var v91 int16 = -1
                            _ = v91
                        }
                        _ = v45
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                }
                _ = v40
            } else {
                if false {
                    for fi92 := 0; fi92 < 5; fi92++ {
                        if true {
                            delete(v3, v30)
                        } else {
                            v9 += v9
                            fmt.Println(v24)
                        }
                        var v93 uint32 = 42
                        v94 := GenS6{F0: -54, F1: 0, F2: 1}
                        _ = fi92
                        _ = v93
                        _ = v94
                    }
                }
                var v95 int = fi29
                for fi96 := 0; fi96 < 1; fi96++ {
                    v3[v30] = (v5 ^ v5)
                    if (!v28) {
                        switch v30 {
                        case "\n\t":
                            v7 = float64((0.1 / 1.5))
                        default:
                            if ("\n\t" >= "foo bar") {
                                switch (v30 + (v30 + "hello")) {
                                case "\n\t":
                                    var v97 int = -9223372036854775808
                                    _ = v97
                                case "hello":
                                    fmt.Println(fi29, v7, v9)
                                case "":
                                    var v98 uint32 = (v5 % ((2147483648) | 1))
                                    _ = v98
                                }
                                fmt.Println(v7, v30)
                                v7 = v7
                            }
                        }
                        var v99 int64 = 9223372036854775807
                        _ = v99
                    }
                    v24 = ((((v7 - 1.5) + (v2 * v7)) - (3.14159 / v24)) / v24)
                    _ = fi96
                }
                _ = v95
            }
        }
        for fi100 := 0; fi100 < 1; fi100++ {
            v101 := GenS7{F0: 1, F1: v9, F2: (v5 >= (v5 / ((v5) | 1)))}
            v5 = v5
            _ = fi100
            _ = v101
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
