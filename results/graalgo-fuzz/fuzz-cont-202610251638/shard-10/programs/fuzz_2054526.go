package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 string; F2 int32 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 uint32; F2 string }

func main() {
    var v1 int8 = 1
    var v2 int8 = v1
    var v3 bool = true
    v4 := []string{"a", "hello"}
    v5 := []int{}
    v6 := GenS1{F0: v3, F1: 0, F2: 897}
    switch "foo bar" {
    case "\n\t":
        var v7 int8 = (v1 % ((-128) | 1))
        _ = v7
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            switch ((("foo bar" + "\n\t") + "\n\t") + "a") {
            case "foo bar":
                switch ("hello" + "hello") {
                case "a":
                    var v8 bool = ((float32(0.1) * 1e10) <= 3.14159)
                    _ = v8
                case "hello":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        if v3 {
                            if (42 != 0) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        } else {
                            var v10 float32 = (1.5 - (0.1 * 0.1))
                            var v11 int32 = 0
                            _ = v10
                            _ = v11
                        }
                        _ = fi9
                    }
                }
            case "":
                switch "a" {
                case "":
                    fmt.Println(v2, v1, v2)
                default:
                    v12 := GenS1{F0: false, F1: 255, F2: 0}
                    _ = v12
                }
            }
        case "foo bar":
            v13 := GenS1{F0: v3, F1: 2, F2: 42}
            _ = v13
        }
    default:
        v2 = (v2 * (v1 / (((^(v2))) | 1)))
    }
    _ = len(v5)
    for fi14 := 0; fi14 < 2; fi14++ {
        switch fi14 {
        case 949:
            if (v2 >= 0) {
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v15 int8 = v1
                    _ = v15
                case "":
                    if (!false) {
                        if (v3 || (fi14 > (^(fi14)))) {
                            v16 := GenS1{F0: (-2.25 <= (float32(float64((((0.1 / (0.1 / (1.5 + ((0.1 / 3.14159) + 3.14159)))) + float32(0.0)) - (-2.25 + 0.1)))) + 1e10)), F1: 0, F2: 2}
                            var v17 uint16 = 903
                            for fi18 := 0; fi18 < 7; fi18++ {
                                var v19 int32 = 42
                                switch "\n\t" {
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "a":
                                    v20 := GenS1{F0: v3, F1: 146, F2: 4294967295}
                                    _ = v20
                                case "foo bar":
                                    v21 := GenS1{F0: v3, F1: 832, F2: 884}
                                    if v3 {
                                        var v22 string = ("a" + "foo bar")
                                        if (0.1 <= (0.1 - 0.1)) {
                                            var v23 int64 = 9223372036854775807
                                            if v3 {
                                                v22 = v22
                                                var v24 int16 = 1000
                                                _ = v24
                                            }
                                            v25 := GenS2{F0: 9223372036854775807}
                                            fmt.Println(v17, v19, v2)
                                            _ = v23
                                            _ = v25
                                        } else {
                                            if v3 {
                                                v22 = v22
                                                v26 := GenS1{F0: false, F1: 2, F2: 98}
                                                _ = len(v4)
                                                var v27 uint = 1
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                        _ = v22
                                    }
                                    _ = v21
                                }
                                _ = fi18
                                _ = v19
                            }
                            _ = v16
                            _ = v17
                        } else {
                            v28 := GenS1{F0: v3, F1: 0, F2: 834}
                            fmt.Println(v1)
                            _ = v28
                        }
                    } else {
                        v29 := GenS2{F0: -1}
                        for fi30 := 0; fi30 < 20; fi30++ {
                            switch "foo bar" {
                            case "foo bar":
                                var v31 uint16 = 42
                                _ = v31
                            }
                            if v3 {
                                var v32 int = (fi30 | (fi30 % ((4294967295) | 1)))
                                var v33 int = fi30
                                _ = v32
                                _ = v33
                            } else {
                                var v34 uint16 = 401
                                _ = v34
                            }
                            _ = fi30
                        }
                        v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                        _ = v29
                    }
                case "\n\t":
                    v35 := GenS1{F0: false, F1: 18446744073709551615, F2: 715}
                    _ = v35
                }
            } else {
                for fi36 := 0; fi36 < 10; fi36++ {
                    v2 = v2
                    _ = fi36
                }
                v1 = (v1 + 61)
            }
        case -612:
            switch fi14 {
            case 32768:
                v37 := GenS2{F0: -1}
                _ = v37
            }
        }
        _ = fi14
    }
    v2 = v1
    var v38 float32 = float32((((0.0 - 0.0) + (0.1 + float64((1.5 * 1.5)))) / 3.14159))
    v1 = v1
    _ = len(v5)
    var v39 float32 = -2.25
    if len(v5) > 0 {
        v5[0] = 151
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 uint64 = 855
        _ = fi40
        _ = v41
    }
    var v42 uint16 = 1
    if (v3 || true) {
        v43 := GenS3{F0: 2147483647}
        var v44 string = ("hello" + ("" + ("foo bar" + (("hello" + "hello") + ("\n\t" + "hello")))))
        _ = v43
        _ = v44
    } else {
        if ("foo bar" > "") {
            v38 += (v39 / v38)
            var v45 bool = true
            _ = len(v5)
            _ = v45
        }
        var v46 bool = v3
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 int64 = 42
            if v3 {
                for fi49 := 0; fi49 < 14; fi49++ {
                    fmt.Println(v3)
                    var v50 int16 = 267
                    var v51 uint16 = (^(v42))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v52 := GenS4{F0: (("hello" + "\n\t") + "a"), F1: "hello", F2: 2147483647}
                    switch (fi47 | fi49) {
                    case -2147483648:
                        switch "a" {
                        case "foo bar":
                            v46 = false
                        default:
                            v5 = append(v5, 4294967295)
                        }
                    }
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
            }
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v38 = v38
    if (1e10 <= 0.1) {
        switch -2147483648 {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 4294967295:
            var v53 uint = 1
            _ = v53
        default:
            v54 := GenS2{F0: -701}
            _ = v54
        }
    } else {
        switch 32768 {
        case -2147483648:
            for fi55 := 0; fi55 < 13; fi55++ {
                switch "" {
                case "a":
                    v56 := GenS5{F0: (v42 % ((895) | 1)), F1: 2147483647, F2: 32767}
                    _ = v56
                case "hello":
                    var v57 string = "\n\t"
                    _ = v57
                default:
                    if (796 <= 109) {
                        switch int(v42) {
                        case -2147483648:
                            v1 -= 34
                        }
                    } else {
                        var v58 float64 = (3.14159 / 1.5)
                        v59 := GenS4{F0: "hello", F1: "\n\t", F2: 42}
                        _ = v58
                        _ = v59
                    }
                }
                _ = fi55
            }
        case 61:
            var v60 int16 = 1
            _ = v60
        case 118:
            var v61 uint8 = 127
            _ = v61
        default:
            switch "\n\t" {
            case "foo bar":
                if v3 {
                    v5 = append(v5, 42)
                }
            case "":
                var v62 int32 = 0
                _ = v62
            }
        }
        v4 = append(v4, ("\n\t" + "\n\t"))
    }
    if false {
        if v3 {
            if (((v3 || v3) && (!((-534 < 1000) || false))) || ("" <= "hello")) {
                for fi63 := 0; fi63 < 3; fi63++ {
                    if (v2 <= (v2 ^ (v2 * 14))) {
                        v2 -= (v2 << ((2) & 63))
                        var v64 float32 = (v38 - 0.0)
                        if (((int16(fi63) != -726) || (v3 || ("" < ""))) || false) {
                            _ = len(v5)
                        } else {
                            var v65 int = fi63
                            _ = v65
                        }
                        _ = v64
                    } else {
                        if true {
                            if v3 {
                                if (!v3) {
                                    switch fi63 {
                                    case 2147483647:
                                        v2 -= -128
                                    case -799:
                                        for fi66 := 0; fi66 < 9; fi66++ {
                                            v3 = v3
                                            _ = fi66
                                        }
                                    case 32768:
                                        v67 := GenS2{F0: 42}
                                        _ = v67
                                    default:
                                        var v68 string = "hello"
                                        _ = v68
                                    }
                                    var v69 int = 32768
                                    _ = v69
                                } else {
                                    v70 := GenS6{F0: v2, F1: (v42 / ((v42) | 1))}
                                    v42 = (v42 & v42)
                                    _ = v70
                                }
                            } else {
                                for fi71 := 0; fi71 < 6; fi71++ {
                                    v72 := GenS1{F0: false, F1: 4294967295, F2: 42}
                                    _ = fi71
                                    _ = v72
                                }
                                for fi73 := 0; fi73 < 11; fi73++ {
                                    fmt.Println(v2)
                                    _ = fi73
                                }
                            }
                        } else {
                            v74 := GenS1{F0: v3, F1: 480, F2: 143}
                            _ = v74
                        }
                    }
                    _ = fi63
                }
                fmt.Println(v1, v39)
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v75 := GenS2{F0: int64(v42)}
            _ = v75
        } else {
            if (0 != -128) {
                for fi76 := 0; fi76 < 5; fi76++ {
                    var v77 int8 = 1
                    switch ("hello" + "hello") {
                    case "hello":
                        var v78 bool = false
                        _ = v78
                    default:
                        switch fi76 {
                        case 0:
                            var v79 uint64 = 1
                            _ = v79
                        case -9223372036854775808:
                            var v80 uint16 = v42
                            _ = v80
                        }
                    }
                    _ = fi76
                    _ = v77
                }
            }
            v1 = 127
            var v81 float64 = ((0.0 + 0.1) / (0.1 / float64((float32((float64(v39) * (0.0 - 3.14159))) - v38))))
            _ = v81
        }
        var v82 string = "a"
        var v83 string = "\n\t"
        _ = v82
        _ = v83
    } else {
        var v84 float64 = (1.5 + 0.1)
        if true {
            var v85 uint16 = v42
            v39 = 1e10
            switch -1 {
            case 40:
                for fi86 := 0; fi86 < 9; fi86++ {
                    if v3 {
                        if v3 {
                            if (!v3) {
                                var v87 uint16 = 32767
                                var v88 int64 = 801
                                v89 := GenS7{F0: ("foo bar" + "")}
                                _ = v87
                                _ = v88
                                _ = v89
                            }
                        }
                    } else {
                        v90 := GenS5{F0: v85, F1: 582, F2: (v85 / ((42) | 1))}
                        if v3 {
                            v91 := GenS7{F0: "a"}
                            _ = v91
                        }
                        switch "hello" {
                        case "\n\t":
                            v38 -= v38
                        case "foo bar":
                            for fi92 := 0; fi92 < 1; fi92++ {
                                var v93 int8 = 12
                                _ = fi92
                                _ = v93
                            }
                        default:
                            if ((v39 + v39) < (1e10 - 0.1)) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi94 := 0; fi94 < 5; fi94++ {
                                    v84 = float64(float32(v84))
                                    var v95 int32 = 824
                                    switch ("a" + (("hello" + ((("" + (("\n\t" + "a") + "hello")) + "hello") + "\n\t")) + "\n\t")) {
                                    case "hello":
                                        var v96 uint8 = 127
                                        _ = v96
                                    case "a":
                                        var v97 uint32 = 487
                                        _ = v97
                                    case "":
                                        var v98 uint64 = 0
                                        if (0 >= 255) {
                                            switch -129 {
                                            case 9223372036854775807:
                                                var v99 int = fi86
                                                _ = v99
                                            default:
                                                var v100 string = "a"
                                                _ = v100
                                            }
                                            var v101 int64 = -1
                                            v42 -= (v85 & v85)
                                            v102 := GenS8{F0: "", F1: 891, F2: "hello"}
                                            v2 = v2
                                            if len(v5) > 0 {
                                                v5[0] = fi94
                                            }
                                            v3 = (v95 == v95)
                                            var v103 float32 = ((v38 * -2.25) * v38)
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                        }
                                        _ = v98
                                    }
                                    _ = fi94
                                    _ = v95
                                }
                            } else {
                                _ = len(v5)
                            }
                        }
                        _ = v90
                    }
                    if len(v5) > 0 {
                        v5[0] = fi86
                    }
                    _ = fi86
                }
            default:
                for fi104 := 0; fi104 < 20; fi104++ {
                    for fi105 := 0; fi105 < 16; fi105++ {
                        v106 := GenS6{F0: 0, F1: (^(v42))}
                        v2 = -94
                        _ = fi105
                        _ = v106
                    }
                    var v107 int16 = 662
                    _ = fi104
                    _ = v107
                }
            }
            _ = v85
        } else {
            fmt.Println(v39, v39)
            var v108 uint32 = 1
            _ = v108
        }
        _ = v84
    }
    for fi109 := 0; fi109 < 9; fi109++ {
        switch "" {
        case "":
            var v110 uint = 18446744073709551615
            _ = v110
        case "hello":
            if v3 {
                v111 := GenS8{F0: "foo bar", F1: 65535, F2: ""}
                _ = v111
            }
        }
        v42 -= 65535
        _ = fi109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
