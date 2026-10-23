package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int; F2 uint64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint32; F2 int16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32 }

func main() {
    var v1 float64 = (1e10 - (0.0 + (1.5 + float64(-2.25))))
    var v2 int64 = 9223372036854775807
    var v3 int32 = 42
    var v4 int = -9223372036854775808
    v5 := []string{}
    switch "hello" {
    case "hello":
        var v6 int64 = (v2 % ((42) | 1))
        _ = v6
    case "\n\t":
        if ((v1 < 0.0) && true) {
            var v7 string = ""
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(v4, fi8, v3)
                var v9 uint16 = 32767
                if false {
                    v9 = 255
                    if (65535 != 18446744073709551615) {
                        var v10 uint64 = 4294967295
                        _ = v10
                    } else {
                        v11 := GenS1{F0: (v4 ^ (fi8 * -2147483648)), F1: 42, F2: 56}
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: (fi12 / ((122) | 1)), F1: 0, F2: 2}
                        if true {
                            v14 := GenS2{F0: 42, F1: 100, F2: v7}
                            v5 = append(v5, (((v7 + "foo bar") + v7) + v7))
                            fmt.Println(v2, v4, v7)
                            _ = v14
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch (-(v4)) {
                            case -9223372036854775808:
                                v9 -= 32767
                            case -2147483648:
                                v7 = "hello"
                            default:
                                if len(v5) > 0 {
                                    v5[0] = ("a" + v7)
                                }
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                }
                _ = fi8
                _ = v9
            }
            var v15 uint32 = 388
            _ = v15
            _ = v7
        }
    case "foo bar":
        v4 += (v4 + v4)
    default:
        v3 = v3
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if (!false) {
        if ((("" == ("hello" + ("foo bar" + "foo bar"))) || ((v4 << ((18446744073709551615) & 63)) > (v4 ^ (-(v4))))) && true) {
            v16 := GenS2{F0: int8(v2), F1: 42, F2: ""}
            v17 := GenS3{F0: 1000}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v16
            _ = v17
        }
    }
    v18 := GenS4{F0: ("hello" + "\n\t")}
    var v19 int = (v4 >> ((18446744073709551615) & 63))
    v4 = v4
    fmt.Println(v1, v19)
    v20 := GenS3{F0: 790}
    v21 := GenS4{F0: "hello"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 -= 0
    switch ("foo bar" + ("a" + "\n\t")) {
    case "foo bar":
        var v22 int16 = -433
        _ = v22
    }
    if false {
        switch "" {
        case "foo bar":
            if true {
                if true {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            v1 += 1.5
                        }
                        if (true && (!(!false))) {
                            var v24 uint16 = 32767
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                v19 += (fi25 | fi23)
                                _ = fi25
                            }
                            var v26 uint64 = 1
                            _ = v26
                        }
                        _ = fi23
                    }
                    if true {
                        fmt.Println(v2)
                        if (v2 >= v2) {
                            var v27 uint = 1
                            switch "" {
                            case "":
                                v4 = v4
                            case "\n\t":
                                v27 -= v27
                            }
                            _ = v27
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v28 string = "hello"
                                _ = v28
                            case "":
                                if (42 != 2147483647) {
                                    for fi29 := 0; fi29 < 5; fi29++ {
                                        _ = len(v5)
                                        if len(v5) > 0 {
                                            v5[0] = "\n\t"
                                        }
                                        switch (^(fi29)) {
                                        case -9223372036854775808:
                                            for fi30 := 0; fi30 < 2; fi30++ {
                                                var v31 uint32 = 2
                                                _ = len(v5)
                                                _ = fi30
                                                _ = v31
                                            }
                                        case 0:
                                            fmt.Println(v1, v4)
                                        case 1:
                                            if (!false) {
                                                v32 := GenS3{F0: 628}
                                                _ = v32
                                            }
                                        }
                                        _ = fi29
                                    }
                                    fmt.Println(v1)
                                    var v33 bool = ((-(v2)) < 499)
                                    _ = v33
                                } else {
                                    v34 := GenS3{F0: -32768}
                                    v5 = append(v5, "\n\t")
                                    if (!true) {
                                        if false {
                                            for fi35 := 0; fi35 < 3; fi35++ {
                                                v36 := GenS2{F0: -1, F1: 127, F2: "foo bar"}
                                                v37 := GenS3{F0: 32767}
                                                v38 := GenS3{F0: 1000}
                                                _ = fi35
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            v2 = 216
                                            switch (v19 ^ v4) {
                                            case 254:
                                                v3 = v3
                                            case -19:
                                                v3 += -2147483648
                                            default:
                                                if false {
                                                    var v39 int8 = -128
                                                    v39 -= v39
                                                    _ = v39
                                                }
                                            }
                                        } else {
                                            for fi40 := 0; fi40 < 8; fi40++ {
                                                var v41 uint16 = 103
                                                _ = fi40
                                                _ = v41
                                            }
                                            _ = len(v5)
                                            v3 += 42
                                        }
                                        v2 = (v2 ^ 1)
                                    }
                                    _ = v34
                                }
                            }
                        }
                    } else {
                        var v42 uint8 = 42
                        v42 = 1
                        v3 = v3
                        _ = v42
                    }
                    v43 := GenS4{F0: "hello"}
                    _ = v43
                }
                if (26 != 4294967295) {
                    v4 -= -487
                } else {
                    switch "hello" {
                    case "a":
                        fmt.Println(v19, v4, v3)
                    case "\n\t":
                        fmt.Println(v1, v2)
                    case "":
                        fmt.Println(v1, v19)
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            v44 := GenS5{F0: 127, F1: (v4 & (v4 << ((626) & 63))), F2: 4294967295}
                            _ = v44
                        case "\n\t":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "":
                            v45 := GenS6{F0: 0, F1: true}
                            _ = v45
                        default:
                            v46 := GenS7{F0: (^(v3))}
                            _ = v46
                        }
                    }
                }
                v1 = (v1 * 0.0)
            } else {
                v19 -= 1
                var v47 int8 = 100
                _ = v47
            }
        case "":
            for fi48 := 0; fi48 < 14; fi48++ {
                for fi49 := 0; fi49 < 1; fi49++ {
                    v4 = (fi49 + v4)
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                    var v50 uint = 538
                    _ = fi49
                    _ = v50
                }
                var v51 uint64 = 18446744073709551615
                for fi52 := 0; fi52 < 15; fi52++ {
                    v53 := GenS5{F0: 171, F1: (v19 % ((v19) | 1)), F2: v51}
                    var v54 int16 = 32767
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi48
                _ = v51
            }
        default:
            if (true && (!false)) {
                if len(v5) > 0 {
                    v5[0] = (("foo bar" + "foo bar") + ("hello" + "a"))
                }
            } else {
                if true {
                    v5 = append(v5, "foo bar")
                } else {
                    v5 = append(v5, ("\n\t" + ""))
                }
                if true {
                    if (42 != 2) {
                        for fi55 := 0; fi55 < 10; fi55++ {
                            v56 := GenS2{F0: 42, F1: 100, F2: ("\n\t" + (("foo bar" + "\n\t") + ""))}
                            for fi57 := 0; fi57 < 2; fi57++ {
                                v19 = 128
                                var v58 float32 = float32(v1)
                                fmt.Println(v3, v3, v2)
                                switch "foo bar" {
                                case "a":
                                    var v59 int8 = 27
                                    _ = v59
                                case "\n\t":
                                    switch ("hello" + ("hello" + "foo bar")) {
                                    case "foo bar":
                                        v2 = v2
                                    case "a":
                                        if (!(!(v3 >= v3))) {
                                            v2 = -1
                                            v3 += v3
                                            _ = len(v5)
                                        } else {
                                            v58 = (((0.1 / float32(v1)) + (v58 * (-2.25 + 0.0))) + ((0.1 - v58) * ((v58 + (v58 + 0.1)) + float32((float64(3.14159) + v1)))))
                                        }
                                    case "\n\t":
                                        var v60 bool = false
                                        _ = v60
                                    }
                                }
                                v61 := GenS8{F0: "hello"}
                                _ = fi57
                                _ = v58
                                _ = v61
                            }
                            _ = fi55
                            _ = v56
                        }
                    } else {
                        for fi62 := 0; fi62 < 3; fi62++ {
                            fmt.Println(v2, v3)
                            switch -1 {
                            case -949:
                                for fi63 := 0; fi63 < 20; fi63++ {
                                    v64 := GenS4{F0: ("a" + "a")}
                                    _ = fi63
                                    _ = v64
                                }
                            case 757:
                                if false {
                                    for fi65 := 0; fi65 < 15; fi65++ {
                                        if false {
                                            if (v3 > v3) {
                                                var v66 uint = 65535
                                                _ = v66
                                            } else {
                                                v1 = v1
                                            }
                                            if true {
                                                fmt.Println(v1, v2, v19)
                                                v67 := GenS9{F0: "hello", F1: 2147483647, F2: -1}
                                                _ = v67
                                            } else {
                                                v1 = v1
                                            }
                                        } else {
                                            if (true || ((9223372036854775807 != (^(v2))) && false)) {
                                                var v68 int = fi62
                                                _ = v68
                                            } else {
                                                var v69 int = -1
                                                v69 = (-(v19))
                                                var v70 int32 = v3
                                                _ = v69
                                                _ = v70
                                            }
                                            if (!false) {
                                                var v71 int = v19
                                                var v72 uint8 = 255
                                                if (!false) {
                                                    if (-128 == 120) {
                                                        v2 -= -1
                                                    }
                                                    switch 4294967295 {
                                                    case 0:
                                                        v72 += v72
                                                    case 42:
                                                        if true {
                                                            _ = len(v5)
                                                        }
                                                    case 127:
                                                        var v73 int8 = 100
                                                        _ = v73
                                                    default:
                                                        fmt.Println(fi65)
                                                    }
                                                    var v74 bool = ((0 >= v19) && true)
                                                    _ = v74
                                                }
                                                v75 := GenS10{F0: 434}
                                                _ = v71
                                                _ = v72
                                                _ = v75
                                            } else {
                                                var v76 float64 = float64((((0.1 + 0.0) * (-2.25 + -2.25)) - -2.25))
                                                var v77 uint16 = 255
                                                _ = v76
                                                _ = v77
                                            }
                                            var v78 int32 = -1
                                            _ = v78
                                        }
                                        v4 = v4
                                        _ = fi65
                                    }
                                    switch (("foo bar" + ("foo bar" + "foo bar")) + "\n\t") {
                                    case "hello":
                                        for fi79 := 0; fi79 < 18; fi79++ {
                                            switch 128 {
                                            case 32767:
                                                v1 = -2.25
                                            default:
                                                var v80 uint64 = 0
                                                _ = v80
                                            }
                                            var v81 int32 = (v3 ^ -2147483648)
                                            if len(v5) > 0 {
                                                v5[0] = ""
                                            }
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(fi79)
                                            default:
                                                switch (-(fi79)) {
                                                case 42:
                                                    if true {
                                                        var v82 uint8 = 4
                                                        v3 = -1
                                                        _ = v82
                                                    }
                                                case 32768:
                                                    v83 := GenS11{F0: 1}
                                                    _ = v83
                                                }
                                                v1 -= 0.0
                                            }
                                            _ = fi79
                                            _ = v81
                                        }
                                    case "":
                                        var v84 uint64 = 4294967295
                                        _ = v84
                                    }
                                }
                            default:
                                var v85 uint8 = 2
                                _ = v85
                            }
                            _ = fi62
                        }
                        for fi86 := 0; fi86 < 18; fi86++ {
                            for fi87 := 0; fi87 < 5; fi87++ {
                                var v88 int64 = v2
                                if false {
                                    v4 = int(v2)
                                    v88 -= v2
                                    if (!true) {
                                        if (-128 > 100) {
                                            fmt.Println(v3)
                                            if false {
                                                var v89 int64 = v2
                                                if len(v5) > 0 {
                                                    v5[0] = ("hello" + ("\n\t" + "foo bar"))
                                                }
                                                var v90 uint32 = 0
                                                _ = v89
                                                _ = v90
                                            }
                                            v3 = 1
                                        } else {
                                            var v91 uint = 530
                                            _ = v91
                                        }
                                    } else {
                                        v92 := GenS2{F0: -79, F1: -128, F2: "a"}
                                        _ = v92
                                    }
                                } else {
                                    var v93 int32 = (v3 & (v3 / ((1) | 1)))
                                    _ = v93
                                }
                                if false {
                                    switch (v4 ^ (-(fi86))) {
                                    case 32767:
                                        v19 -= fi87
                                    }
                                    var v94 uint8 = 248
                                    var v95 int8 = 0
                                    _ = v94
                                    _ = v95
                                }
                                if false {
                                    var v96 bool = (!(fi86 != -1))
                                    _ = v96
                                }
                                var v97 int32 = v3
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v98 int32 = v97
                                _ = fi87
                                _ = v88
                                _ = v97
                                _ = v98
                            }
                            var v99 uint64 = 846
                            _ = fi86
                            _ = v99
                        }
                    }
                    v2 = (v2 << ((18446744073709551615) & 63))
                    v4 = 0
                } else {
                    v4 = v19
                    var v100 float64 = (v1 + ((((v1 / (0.1 * v1)) * (v1 + v1)) - v1) / v1))
                    for fi101 := 0; fi101 < 2; fi101++ {
                        v19 = 2147483647
                        var v102 float32 = 1.5
                        v103 := GenS1{F0: (v19 + (^(v4))), F1: 65535, F2: 187}
                        _ = fi101
                        _ = v102
                        _ = v103
                    }
                    _ = v100
                }
            }
        }
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v104 int16 = -1
            v1 = 1e10
            _ = v104
        }
    } else {
        v1 = v1
        var v105 bool = (v19 < (v19 % (((v19 ^ (v4 << ((0) & 63)))) | 1)))
        var v106 float64 = v1
        _ = v105
        _ = v106
    }
    _ = len(v5)
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
