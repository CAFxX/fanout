package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 int8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint32; F1 uint8; F2 string }
type GenS8 struct { F0 int32; F1 int; F2 int8 }
type GenS9 struct { F0 int16; F1 uint32 }

func main() {
    var v1 float64 = -2.25
    v2 := []string{"\n\t", "a", (((("foo bar" + "a") + "a") + "foo bar") + "\n\t"), "foo bar"}
    var v3 int32 = -2147483648
    var v4 uint8 = 42
    v5 := []uint32{847, 2147483648, 2147483647, 2}
    v6 := GenS1{F0: -32768, F1: 962}
    var v7 int8 = 100
    var v8 uint64 = 996
    v1 = (v1 - (0.1 * ((((0.1 + v1) - v1) - ((-2.25 - v1) * 1.5)) + v1)))
    var v9 int64 = -737
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v8)
        var v11 int64 = (v9 % (((^(v9))) | 1))
        _ = fi10
        _ = v11
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v9 = (^(v9))
    }
    if true {
        v8 = v8
        v7 -= (-(v7))
        if false {
            for fi12 := 0; fi12 < 15; fi12++ {
                switch (^(fi12)) {
                case 840:
                    v3 = v3
                case -9223372036854775808:
                    var v13 bool = true
                    _ = v13
                case -881:
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 uint = 4294967295
                        _ = fi14
                        _ = v15
                    }
                }
                var v16 int = (fi12 - fi12)
                v17 := GenS1{F0: 32767, F1: -783}
                _ = fi12
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                if (((!(!true)) && true) || ((fi18 >> ((414) & 63)) >= (-(fi18)))) {
                    if true {
                        if false {
                            switch (fi18 | fi18) {
                            case 894:
                                v8 = 918
                            case 32767:
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    v20 := GenS1{F0: 0, F1: 1}
                                    fmt.Println(v7)
                                    var v21 string = "foo bar"
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            var v22 float64 = 1.5
                            _ = v22
                        }
                    } else {
                        var v23 int32 = v3
                        _ = v23
                    }
                    var v24 uint16 = 42
                    v9 -= 0
                    _ = v24
                } else {
                    if (v7 >= v7) {
                        var v25 int64 = (v9 | v9)
                        if (2147483647 != 671) {
                            v26 := GenS1{F0: -1, F1: 1000}
                            switch "\n\t" {
                            case "":
                                v25 = (v25 >> ((1) & 63))
                            default:
                                switch "" {
                                case "":
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        var v28 bool = true
                                        _ = fi27
                                        _ = v28
                                    }
                                case "a":
                                    switch "foo bar" {
                                    case "a":
                                        v2 = append(v2, (("\n\t" + "\n\t") + ("" + "\n\t")))
                                    }
                                case "foo bar":
                                    fmt.Println(v3)
                                default:
                                    v29 := GenS1{F0: -1, F1: -1}
                                    _ = v29
                                }
                            }
                            _ = v26
                        } else {
                            v30 := GenS1{F0: 247, F1: 1000}
                            for fi31 := 0; fi31 < 4; fi31++ {
                                var v32 bool = false
                                if v32 {
                                    var v33 int8 = v7
                                    if true {
                                        v9 += v25
                                        v4 += v4
                                        if (-2.25 == 1e10) {
                                            var v34 float64 = (((0.1 * float64(1.5)) - v1) * v1)
                                            _ = v34
                                        } else {
                                            v35 := GenS2{F0: (0 != 2147483648), F1: (v4 - v4), F2: (v4 % ((v4) | 1))}
                                            _ = v35
                                        }
                                    } else {
                                        var v36 uint8 = v4
                                        if ((1e10 - 1e10) == (v1 + float64(float32((1.5 * v1))))) {
                                            switch fi18 {
                                            case -1:
                                                var v37 int = 0
                                                _ = v37
                                            case 317:
                                                if v32 {
                                                    fmt.Println(v9)
                                                    var v38 float64 = 1.5
                                                    if v32 {
                                                        if ((^(v8)) >= 42) {
                                                            if v32 {
                                                                var v39 uint32 = 42
                                                                if false {
                                                                    v1 += 1.5
                                                                    var v40 int = fi18
                                                                    if v32 {
                                                                        v41 := GenS2{F0: true, F1: v36, F2: (v4 & (v4 << ((65535) & 63)))}
                                                                        v2 = append(v2, "foo bar")
                                                                        _ = v41
                                                                    }
                                                                    v2 = append(v2, "a")
                                                                    var v42 uint8 = (v4 * 128)
                                                                    if (((1.5 + 0.0) + -2.25) != 1.5) {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 42
                                                                        }
                                                                        v43 := GenS2{F0: v32, F1: v36, F2: v42}
                                                                        v40 = fi18
                                                                        _ = v43
                                                                    } else {
                                                                        switch (v40 % (((fi18 ^ fi31)) | 1)) {
                                                                        case 958:
                                                                            switch ("foo bar" + "hello") {
                                                                            case "\n\t":
                                                                                var v44 uint32 = 736
                                                                                _ = v44
                                                                            }
                                                                        case 1:
                                                                            var v45 int32 = (v3 % ((2147483647) | 1))
                                                                            _ = v45
                                                                        case 32768:
                                                                            v46 := GenS2{F0: true, F1: 42, F2: v42}
                                                                            _ = v46
                                                                        }
                                                                        _ = len(v5)
                                                                    }
                                                                    var v47 uint8 = v36
                                                                    _ = v40
                                                                    _ = v42
                                                                    _ = v47
                                                                }
                                                                _ = v39
                                                            } else {
                                                                v48 := GenS2{F0: (!v32), F1: (v4 - v4), F2: v36}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v48
                                                            }
                                                        } else {
                                                            var v49 int8 = (v7 << ((671) & 63))
                                                            _ = v49
                                                        }
                                                        v7 -= 100
                                                    } else {
                                                        v5 = append(v5, 65535)
                                                        v50 := GenS2{F0: v32, F1: (v4 << ((4294967295) & 63)), F2: v4}
                                                        var v51 int32 = v3
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = v38
                                                } else {
                                                    v52 := GenS2{F0: true, F1: v36, F2: (^(v4))}
                                                    v36 = 255
                                                    _ = v52
                                                }
                                            }
                                        }
                                        v53 := GenS3{F0: ((v1 / 3.14159) > v1), F1: -1, F2: (v7 + 6)}
                                        _ = v36
                                        _ = v53
                                    }
                                    var v54 float32 = float32(0.1)
                                    _ = v33
                                    _ = v54
                                }
                                var v55 uint32 = 2
                                _ = fi31
                                _ = v32
                                _ = v55
                            }
                            for fi56 := 0; fi56 < 12; fi56++ {
                                v57 := GenS4{F0: v9, F1: 0, F2: 42}
                                var v58 bool = (!true)
                                v25 -= v25
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            _ = v30
                        }
                        _ = v25
                    }
                    if true {
                        var v59 int8 = (v7 >> ((580) & 63))
                        var v60 int8 = 0
                        if (false && ((v9 / ((v9) | 1)) >= (v9 & v9))) {
                            v7 = v7
                            var v61 bool = ("\n\t" >= "hello")
                            var v62 int8 = v60
                            _ = v61
                            _ = v62
                        }
                        _ = v59
                        _ = v60
                    }
                    var v63 string = "\n\t"
                    _ = v63
                }
                _ = fi18
            }
        } else {
            switch ((("hello" + "\n\t") + "") + "foo bar") {
            case "foo bar":
                v9 -= (v9 & (v9 | -51))
            case "\n\t":
                for fi64 := 0; fi64 < 2; fi64++ {
                    _ = len(v5)
                    v5 = append(v5, 0)
                    _ = fi64
                }
            case "":
                var v65 int64 = (^(v9))
                _ = v65
            }
            for fi66 := 0; fi66 < 14; fi66++ {
                for fi67 := 0; fi67 < 2; fi67++ {
                    v68 := GenS3{F0: false, F1: 42, F2: v7}
                    switch int(v3) {
                    case -915:
                        if false {
                            if false {
                                v69 := GenS2{F0: false, F1: v4, F2: 233}
                                var v70 int16 = 947
                                v3 += (v3 & v3)
                                _ = v69
                                _ = v70
                            }
                        }
                        v1 = v1
                        v1 = 0.1
                    default:
                        var v71 string = ("a" + "\n\t")
                        _ = v71
                    }
                    _ = fi67
                    _ = v68
                }
                v5 = append(v5, 2147483647)
                _ = fi66
            }
            var v72 string = "foo bar"
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 2; fi73++ {
        v74 := GenS2{F0: (true || (!true)), F1: 127, F2: v4}
        fmt.Println(v4, v7)
        switch "" {
        case "\n\t":
            switch (fi73 - fi73) {
            case -641:
                var v75 int8 = 42
                _ = v75
            case 1:
                var v76 uint8 = (^(v4))
                _ = v76
            case 9223372036854775807:
                if ((((((1e10 - -2.25) - -2.25) - 0.0) * float32(-2.25)) + (1e10 + (0.1 - -2.25))) <= 0.0) {
                    v77 := GenS5{F0: -666}
                    _ = v77
                }
            }
        case "hello":
            var v78 int8 = v7
            _ = v78
        case "a":
            for fi79 := 0; fi79 < 5; fi79++ {
                v3 = v3
                if (false && false) {
                    if false {
                        v80 := GenS6{F0: "", F1: fi79}
                        var v81 int8 = v7
                        v82 := GenS4{F0: v9, F1: 2147483648, F2: 32767}
                        fmt.Println(v8, fi73)
                        switch -316 {
                        case -769:
                            var v83 uint16 = 599
                            _ = v83
                        case -450:
                            v7 += 49
                        case 32768:
                            fmt.Println(v4, v4)
                        default:
                            var v84 uint16 = uint16(v8)
                            v85 := GenS7{F0: 2147483647, F1: (v4 - v4), F2: ("foo bar" + "foo bar")}
                            var v86 uint16 = v84
                            v9 = 0
                            if true {
                                if (true || true) {
                                    fmt.Println(v9, v81, v84)
                                } else {
                                    v1 = float64(-2.25)
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v7 = (v7 >> ((42) & 63))
                            } else {
                                var v87 bool = ((v9 / ((v9) | 1)) == v9)
                                _ = v87
                            }
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        _ = v80
                        _ = v81
                        _ = v82
                    }
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    if (false && (523 < 4294967295)) {
                        v88 := GenS4{F0: (v9 ^ -86), F1: 860, F2: 107}
                        fmt.Println(v3, v8, v7)
                        _ = v88
                    } else {
                        switch 9223372036854775807 {
                        case 32767:
                            fmt.Println(v7)
                        default:
                            v2 = append(v2, ("" + "a"))
                        }
                        v5 = append(v5, 0)
                    }
                } else {
                    if false {
                        var v89 float64 = float64(((((1e10 + (3.14159 + (((3.14159 - 3.14159) * float32(float64(float32(v1)))) * 0.1))) + -2.25) / 3.14159) + 0.0))
                        _ = v89
                    } else {
                        v5 = append(v5, 0)
                    }
                }
                _ = fi79
            }
        }
        _ = fi73
        _ = v74
    }
    if true {
        var v90 float32 = ((1.5 * float32(float64(0.0))) / (1.5 + 3.14159))
        var v91 string = ""
        v92 := GenS8{F0: (v3 ^ (v3 + v3)), F1: int(v9), F2: (v7 & (v7 << ((65535) & 63)))}
        _ = v90
        _ = v91
        _ = v92
    } else {
        v93 := GenS8{F0: v3, F1: 127, F2: 100}
        var v94 int32 = v3
        fmt.Println(v7, v3, v4)
        _ = v93
        _ = v94
    }
    v8 -= (^(v8))
    switch 9223372036854775807 {
    case 741:
        var v95 uint32 = 4294967295
        _ = v95
    case 503:
        if false {
            for fi96 := 0; fi96 < 8; fi96++ {
                for fi97 := 0; fi97 < 20; fi97++ {
                    switch ("hello" + ("\n\t" + "")) {
                    case "\n\t":
                        var v98 int64 = (v9 / ((v9) | 1))
                        _ = v98
                    default:
                        v3 += (v3 + -2147483648)
                    }
                    if false {
                        var v99 uint64 = v8
                        v9 = 42
                        _ = v99
                    }
                    _ = fi97
                }
                _ = fi96
            }
            v7 = v7
            var v100 int16 = 0
            _ = v100
        } else {
            var v101 float64 = v1
            var v102 uint = 1
            if len(v2) > 0 {
                v2[0] = "a"
            }
            _ = v101
            _ = v102
        }
    case 4294967295:
        v9 = v9
    default:
        switch ("\n\t" + "hello") {
        case "foo bar":
            if false {
                var v103 bool = false
                v8 = (v8 - uint64(v4))
                var v104 int16 = -403
                _ = v103
                _ = v104
            } else {
                fmt.Println(v8)
                var v105 bool = ((-(v7)) == (v7 | (-(v7))))
                _ = v105
            }
        case "":
            v106 := GenS9{F0: int16(v4), F1: 872}
            _ = v106
        default:
            fmt.Println(v9)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
