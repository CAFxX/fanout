package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 uint16; F2 int64 }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]string)
    v3 := []string{"", "", "hello", (("\n\t" + ("\n\t" + "a")) + ("a" + ((("foo bar" + "\n\t") + "") + ("a" + "a"))))}
    var v4 float64 = ((-2.25 - (0.1 + -2.25)) - float64(float32((0.0 * float64((-2.25 * (1e10 / 1.5)))))))
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            var v7 int = fi6
            var v8 int = 32768
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 42, F1: (v1 * v1), F2: 230}
            var v10 string = "foo bar"
            var v11 int = fi6
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            switch (fi12 & fi12) {
            case 2147483647:
                switch (fi12 >> ((87) & 63)) {
                case -734:
                    v13 := GenS1{F0: 1, F1: (-(v1)), F2: 0}
                    _ = v13
                }
            case 32768:
                fmt.Println(fi12, fi6)
            case 1:
                fmt.Println(v4, v1)
            default:
                switch "a" {
                case "hello":
                    v14 := GenS2{F0: (v1 << ((2) & 63))}
                    _ = v14
                case "a":
                    var v15 uint64 = 18446744073709551615
                    if ((1.5 >= 0.1) || (!true)) {
                        v15 = 2
                        if ((v4 <= v4) && (true || (645 <= 15))) {
                            v15 -= v15
                            var v16 int64 = 0
                            switch "a" {
                            case "a":
                                switch "a" {
                                case "a":
                                    var v17 float32 = float32(-2.25)
                                    _ = v17
                                case "":
                                    switch fi6 {
                                    case 32767:
                                        var v18 bool = false
                                        _ = v18
                                    case -830:
                                        if ("foo bar" == (("hello" + "hello") + "foo bar")) {
                                            v19 := GenS1{F0: 65535, F1: (v1 >> ((569) & 63)), F2: fi6}
                                            _ = v19
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                case "foo bar":
                                    if (2147483647 < 1) {
                                        v20 := GenS3{F0: 2}
                                        fmt.Println(v1)
                                        _ = v20
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v21 := v2["\n\t"]
                                            _ = v21
                                        case "foo bar":
                                            v15 += v15
                                        case "a":
                                            fmt.Println(fi6, fi12)
                                        }
                                        var v22 bool = (763 < 612)
                                        _ = v22
                                    }
                                default:
                                    v23 := GenS4{F0: ((v1 & (v1 * (v1 ^ -32768))) < 0)}
                                    _ = v23
                                }
                            case "\n\t":
                                v16 = v16
                            case "foo bar":
                                var v24 int = (fi6 & fi6)
                                _ = v24
                            default:
                                if true {
                                    v1 -= 42
                                    v25 := GenS2{F0: (v1 * (v1 * (v1 * (v1 % ((1000) | 1)))))}
                                    fmt.Println(v16, v15)
                                    _ = v25
                                } else {
                                    var v26 bool = true
                                    _ = v26
                                }
                            }
                            _ = v16
                        }
                        if len(v3) > 0 {
                            v3[0] = (("" + "a") + "hello")
                        }
                    }
                    v27 := GenS5{F0: -32768}
                    v3 = append(v3, "foo bar")
                    if (1e10 < 1e10) {
                        var v28 int8 = 55
                        _ = len(v2)
                        var v29 int64 = 1
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint64 = v15
                        v1 = -362
                        switch "\n\t" {
                        case "\n\t":
                            v4 = v4
                        case "":
                            var v31 uint = 937
                            _ = v31
                        case "foo bar":
                            v32 := v5[42]
                            _ = v32
                        }
                        _ = v30
                    }
                    _ = v15
                    _ = v27
                }
                if ((v1 & v1) >= (v1 / ((v1) | 1))) {
                    v2[""] = "a"
                    var v33 float32 = (0.0 / 3.14159)
                    var v34 uint32 = 295
                    var v35 float64 = v4
                    if false {
                        v36 := v5[(^(fi6))]
                        var v37 int = (fi12 & -502)
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS5{F0: v1}
                    fmt.Println(v1, v35)
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v38
                } else {
                    v4 = v4
                    var v39 int32 = -532
                    _ = v39
                }
                v1 += 639
                var v40 float32 = 1e10
                _ = v40
            }
            _ = fi12
        }
        for fi41 := 0; fi41 < 17; fi41++ {
            if ((!true) && (!false)) {
                v3 = append(v3, "a")
                if true {
                    var v42 uint32 = 42
                    if (!(-1 > -453)) {
                        v43 := v5[(fi6 & fi41)]
                        v1 = (v1 & (v1 + v1))
                        if (!true) {
                            v44 := GenS3{F0: 469}
                            v4 = v4
                            var v45 string = "hello"
                            _ = v44
                            _ = v45
                        }
                        _ = v43
                    } else {
                        var v46 bool = false
                        if false {
                            var v47 uint32 = v42
                            if false {
                                v42 -= v47
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    var v48 int16 = v1
                                    _ = v48
                                case "":
                                    var v49 bool = (4294967295 <= 42)
                                    _ = v49
                                case "foo bar":
                                    v50 := GenS3{F0: 65535}
                                    _ = v50
                                }
                            }
                            var v51 bool = (34 == 42)
                            _ = v47
                            _ = v51
                        }
                        var v52 bool = (1 < 1)
                        _ = v46
                        _ = v52
                    }
                    _ = v42
                } else {
                    if true {
                        switch fi6 {
                        case -496:
                            if (!(true || false)) {
                                if false {
                                    v1 = -32768
                                } else {
                                    if (!false) {
                                        switch (((("hello" + "") + "a") + ("" + "")) + "") {
                                        case "hello":
                                            var v53 int = (fi6 + (fi6 << ((642) & 63)))
                                            _ = v53
                                        case "a":
                                            fmt.Println(fi6)
                                            delete(v2, "a")
                                            var v54 uint = 4294967295
                                            _ = v54
                                        case "\n\t":
                                            var v55 uint = 325
                                            _ = v55
                                        default:
                                            var v56 uint32 = 2147483648
                                            _ = v56
                                        }
                                    } else {
                                        var v57 bool = (((v1 ^ v1) != 1000) || true)
                                        if (1 < 2) {
                                            var v58 uint8 = 127
                                            var v59 string = "foo bar"
                                            if false {
                                                v3 = append(v3, v59)
                                                var v60 uint32 = 42
                                                fmt.Println(v60, v1, fi41)
                                                _ = v60
                                            } else {
                                                fmt.Println(v57)
                                                v1 = -1
                                            }
                                            _ = v58
                                            _ = v59
                                        }
                                        var v61 uint64 = 42
                                        v62 := GenS6{F0: (v1 & (-(v1))), F1: -321}
                                        v4 += v4
                                        var v63 int32 = 1
                                        _ = v57
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                }
                                var v64 int32 = 870
                                _ = v64
                            }
                        case 465:
                            var v65 int8 = 42
                            _ = v65
                        default:
                            v4 += float64((1.5 + (3.14159 - 3.14159)))
                        }
                        v4 += v4
                    }
                }
            } else {
                v4 = ((((0.1 + v4) * (float64(-2.25) - v4)) - v4) + (v4 - v4))
                var v66 int = (fi6 ^ -1)
                v1 -= 1000
                var v67 int64 = 189
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v66
                _ = v67
            }
            var v68 uint64 = 30
            if len(v3) > 0 {
                v3[0] = (("\n\t" + "hello") + ("foo bar" + "a"))
            }
            _ = fi41
            _ = v68
        }
        _ = fi6
    }
    var v69 float32 = float32(v4)
    for fi70 := 0; fi70 < 12; fi70++ {
        v3 = append(v3, ("" + ""))
        v71 := GenS4{F0: (968 != 10)}
        _ = fi70
        _ = v71
    }
    v4 = v4
    switch ("" + (("\n\t" + "hello") + "\n\t")) {
    case "foo bar":
        fmt.Println(v69, v69)
    case "\n\t":
        var v72 uint = 2
        _ = v72
    default:
        if (1 == 2147483647) {
            v1 = v1
        }
    }
    if (!(-682 > (v1 >> ((890) & 63)))) {
        var v73 bool = true
        _ = v73
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v74 uint64 = 1
    fmt.Println(v69, v1, v4)
    var v75 int64 = 9223372036854775807
    var v76 int64 = (^(v75))
    if true {
        v74 = 18446744073709551615
    }
    var v77 string = "\n\t"
    var v78 uint16 = 32767
    if (true && (42 == 4294967295)) {
        v79 := GenS1{F0: 42, F1: v1, F2: 32768}
        _ = v79
    }
    var v80 uint8 = 127
    for fi81 := 0; fi81 < 7; fi81++ {
        for fi82 := 0; fi82 < 13; fi82++ {
            v80 = v80
            _ = fi82
        }
        v1 = (v1 - (v1 + v1))
        v83 := v2[((("" + (v77 + v77)) + "") + "a")]
        _ = fi81
        _ = v83
    }
    if true {
        v4 += (float64(float32(((v4 / (float64(float32(float64((1.5 * v69)))) * v4)) + (0.0 + 0.0)))) - v4)
        var v84 uint = 926
        _ = v84
    } else {
        for fi85 := 0; fi85 < 10; fi85++ {
            var v86 uint32 = 2
            _ = fi85
            _ = v86
        }
        for fi87 := 0; fi87 < 2; fi87++ {
            if false {
                if (((v77 + v77) + (v77 + "\n\t")) >= "\n\t") {
                    switch (-(fi87)) {
                    case -2147483648:
                        delete(v2, (((v77 + (v77 + "")) + "hello") + (v77 + v77)))
                    case -186:
                        v88 := GenS7{F0: fi87}
                        _ = v88
                    case 4294967295:
                        for fi89 := 0; fi89 < 10; fi89++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi89
                        }
                    default:
                        v90 := GenS8{F0: -1, F1: v78, F2: v76}
                        _ = v90
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v1, v69)
                    default:
                        v91 := GenS6{F0: v1, F1: (v75 + (v75 * v76))}
                        _ = v91
                    }
                    var v92 int32 = 42
                    _ = v92
                }
            }
            v93 := GenS7{F0: fi87}
            _ = fi87
            _ = v93
        }
    }
    v94 := GenS2{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v69:", v69)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v80
    _ = v94
}
