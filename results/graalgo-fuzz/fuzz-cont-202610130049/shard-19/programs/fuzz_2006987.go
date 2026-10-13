package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint16; F2 int8 }
type GenS7 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    var v3 int8 = 1
    var v4 string = (("hello" + "a") + "foo bar")
    var v5 int = -788
    v6 := []int{v5, v5, -643, (v5 >> ((1) & 63))}
    v5 += (v5 / ((-1) | 1))
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = 128
            v5 = -9223372036854775808
            v8 = v8
            _ = fi7
            _ = v8
        }
        var v9 float32 = 3.14159
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        switch (v4 + (v4 + v4)) {
        case "a":
            var v11 int16 = 42
            _ = v11
        }
        var v12 int8 = -128
        _ = fi10
        _ = v12
    }
    v13 := GenS1{F0: -2147483648, F1: (v4 + v4)}
    v4 = "hello"
    if len(v6) > 0 {
        v6[0] = v5
    }
    var v14 string = "a"
    _ = len(v2)
    var v15 uint32 = 607
    v16 := GenS1{F0: 42, F1: v4}
    v4 = (v14 + "\n\t")
    if len(v6) > 0 {
        v6[0] = 127
    }
    if false {
        v17 := GenS1{F0: 2147483647, F1: "a"}
        if false {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: -774, F1: v14}
                _ = fi18
                _ = v19
            }
            switch (v14 + v4) {
            case "foo bar":
                for fi20 := 0; fi20 < 3; fi20++ {
                    if (0 >= 42) {
                        if false {
                            var v21 float64 = 1.5
                            if ((true && (!(!(v15 > (^(v15)))))) || (false && true)) {
                                v22 := GenS1{F0: -683, F1: v4}
                                var v23 int32 = 1
                                _ = v22
                                _ = v23
                            } else {
                                v15 = 2147483647
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    var v25 uint16 = 0
                                    v14 = v14
                                    _ = fi24
                                    _ = v25
                                }
                                var v26 uint8 = 127
                                _ = v26
                            }
                            _ = v21
                        }
                        v27 := GenS1{F0: int32(v15), F1: v4}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(v14)
                            if (!false) {
                                v29 := GenS1{F0: 2147483647, F1: "\n\t"}
                                if (617 == 2) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                v30 := GenS1{F0: 0, F1: v14}
                                if (1e10 == float32((0.1 - ((float64(float32((-2.25 / 3.14159))) * 0.0) * (float64((1e10 + (((0.0 * ((((0.1 * 0.0) - 0.1) + -2.25) * 1.5)) / float32((0.1 - 1e10))) + 1.5))) / 3.14159))))) {
                                    var v31 string = "a"
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v31
                                } else {
                                    var v32 string = v14
                                    var v33 float64 = 0.0
                                    if true {
                                        v32 = (("foo bar" + "\n\t") + (v4 + (v4 + ("foo bar" + v4))))
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                var v34 uint = 18446744073709551615
                                switch 1 {
                                case 1:
                                    switch v5 {
                                    case 32768:
                                        var v35 uint = v34
                                        _ = v35
                                    case -2147483648:
                                        v36 := GenS2{F0: v14, F1: 42, F2: 42}
                                        _ = v36
                                    case 2147483647:
                                        var v37 uint16 = 788
                                        _ = v37
                                    }
                                case 32768:
                                    fmt.Println(v3, v3)
                                default:
                                    var v38 int16 = 326
                                    _ = v38
                                }
                                if (!(-2.25 > 3.14159)) {
                                    var v39 float64 = (3.14159 * 0.0)
                                    switch (v4 + "") {
                                    case "a":
                                        v4 = v14
                                        var v40 string = "a"
                                        v41 := GenS2{F0: "a", F1: 74, F2: 2}
                                        _ = v40
                                        _ = v41
                                    case "":
                                        switch v5 {
                                        case 128:
                                            fmt.Println(v5, v14)
                                            v6 = append(v6, -2147483648)
                                            if true {
                                                switch 32767 {
                                                case 1:
                                                    v1 = v3
                                                case 127:
                                                    switch v5 {
                                                    case 9223372036854775807:
                                                        v3 = 0
                                                    default:
                                                        switch v4 {
                                                        case "hello":
                                                            v5 -= v5
                                                            v42 := GenS1{F0: 0, F1: v4}
                                                            _ = v42
                                                        case "":
                                                            if true {
                                                                var v43 int = (fi20 | (v5 & (fi28 / ((1) | 1))))
                                                                switch v4 {
                                                                case "\n\t":
                                                                    v6 = append(v6, v5)
                                                                default:
                                                                    var v44 float64 = ((v39 / v39) - 0.1)
                                                                    _ = v44
                                                                }
                                                                v45 := GenS1{F0: 42, F1: ((v4 + "") + (v4 + v14))}
                                                                _ = v43
                                                                _ = v45
                                                            } else {
                                                                switch 127 {
                                                                case 9223372036854775807:
                                                                    var v46 uint32 = (v15 * 42)
                                                                    _ = v46
                                                                case -9223372036854775808:
                                                                    v1 += (^(v1))
                                                                case 0:
                                                                    var v47 string = v4
                                                                    _ = v47
                                                                default:
                                                                    fmt.Println(v1)
                                                                }
                                                                v14 = v4
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v3, v4)
                                                        default:
                                                            v48 := GenS3{F0: (-(v1)), F1: 267}
                                                            _ = v48
                                                        }
                                                    }
                                                case 32767:
                                                    v4 = ("foo bar" + (v4 + (v14 + ((v4 + (("\n\t" + (("\n\t" + v4) + "\n\t")) + "")) + ""))))
                                                }
                                                if ((v1 << (((v34 - v34)) & 63)) == (v1 * (v1 * (v3 ^ (v3 - v3))))) {
                                                    v39 = float64((3.14159 / 1.0))
                                                    v39 = v39
                                                }
                                                v49 := GenS3{F0: v3, F1: 18446744073709551615}
                                                _ = v49
                                            }
                                        case -471:
                                            _ = len(v6)
                                            var v50 uint = (^(v34))
                                            _ = v50
                                        }
                                    case "\n\t":
                                        v5 -= (fi28 >> ((v34) & 63))
                                    }
                                    _ = v39
                                }
                                if false {
                                    v5 = 2147483647
                                } else {
                                    _ = len(v6)
                                }
                                v1 -= v3
                                if true {
                                    switch v5 {
                                    case -2147483648:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v51 uint = 18446744073709551615
                                        _ = v51
                                    case -814:
                                        v3 = v3
                                    }
                                    switch (v5 % ((fi20) | 1)) {
                                    case 32768:
                                        v52 := GenS4{F0: int16(v34), F1: v34}
                                        _ = v52
                                    case 128:
                                        if (0 > 42) {
                                            v53 := GenS5{F0: (v3 & v3), F1: 127, F2: 18446744073709551615}
                                            var v54 uint32 = v15
                                            if (false && true) {
                                                v55 := GenS3{F0: (v3 & (v3 & 0)), F1: 42}
                                                var v56 float64 = 1.5
                                                _ = v55
                                                _ = v56
                                            } else {
                                                v6 = append(v6, (fi20 | (^(fi28))))
                                            }
                                            v4 = ("" + ("a" + (v14 + ((("" + (v4 + v4)) + v4) + "\n\t"))))
                                            var v57 uint16 = 2
                                            _ = v53
                                            _ = v54
                                            _ = v57
                                        } else {
                                            _ = len(v6)
                                            if ((-(v1)) < (v1 + 126)) {
                                                v58 := GenS3{F0: v1, F1: 2}
                                                v59 := GenS2{F0: v14, F1: 0, F2: 2}
                                                _ = v58
                                                _ = v59
                                            } else {
                                                switch v5 {
                                                case -865:
                                                    var v60 int64 = 9223372036854775807
                                                    _ = v60
                                                default:
                                                    switch (v14 + "a") {
                                                    case "hello":
                                                        v6 = append(v6, (fi28 & v5))
                                                    case "foo bar":
                                                        delete(v2, v5)
                                                    default:
                                                        v61 := v2[(-(fi20))]
                                                        _ = v61
                                                    }
                                                }
                                                var v62 int16 = -32768
                                                var v63 uint8 = 1
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    }
                                    switch (("a" + ((v14 + v14) + "\n\t")) + "") {
                                    case "":
                                        v64 := GenS4{F0: int16(fi28), F1: v34}
                                        _ = v64
                                    }
                                }
                                _ = v29
                                _ = v30
                                _ = v34
                            } else {
                                var v65 int32 = 0
                                fmt.Println(v65, v1, v65)
                                v66 := GenS6{F0: false, F1: 32767, F2: 42}
                                var v67 uint8 = 127
                                switch (v14 + v4) {
                                case "hello":
                                    var v68 int8 = v3
                                    _ = v68
                                case "a":
                                    v15 += v15
                                }
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            v6 = append(v6, (v5 * (fi20 & fi20)))
                            _ = fi28
                        }
                        if false {
                            if true {
                                var v69 int = fi20
                                _ = v69
                            }
                        }
                    }
                    var v70 float32 = -2.25
                    var v71 uint8 = 128
                    _ = fi20
                    _ = v70
                    _ = v71
                }
            case "a":
                v1 = (v1 << ((255) & 63))
            default:
                if (!true) {
                    if false {
                        v72 := GenS1{F0: 1, F1: v4}
                        if true {
                            var v73 int8 = v1
                            v15 -= (^(v15))
                            _ = v73
                        } else {
                            if ((float32(0.1) / 1.0) != 1.5) {
                                _ = len(v6)
                                var v74 uint8 = 255
                                _ = v74
                            }
                            v15 = v15
                        }
                        _ = v72
                    }
                }
            }
        } else {
            v2[32768] = (1e10 <= 0.1)
        }
        _ = v17
    } else {
        for fi75 := 0; fi75 < 8; fi75++ {
            v5 -= (-(v5))
            var v76 uint16 = 274
            if len(v6) > 0 {
                v6[0] = fi75
            }
            _ = fi75
            _ = v76
        }
    }
    if false {
        if len(v6) > 0 {
            v6[0] = v5
        }
    } else {
        if false {
            v77 := GenS6{F0: false, F1: 148, F2: v1}
            if (0 == 998) {
                v78 := GenS7{F0: 42, F1: 127}
                _ = v78
            } else {
                fmt.Println(v5)
                for fi79 := 0; fi79 < 8; fi79++ {
                    fmt.Println(v4, v15)
                    fmt.Println(v3, fi79, v4)
                    _ = fi79
                }
                fmt.Println(v14, v1, v5)
            }
            _ = v77
        }
    }
    delete(v2, (^(v5)))
    switch -638 {
    case -2147483648:
        for fi80 := 0; fi80 < 4; fi80++ {
            if false {
                v5 = fi80
            }
            v15 += (v15 + 0)
            v2[(fi80 ^ v5)] = (-98 > 626)
            _ = fi80
        }
    case 32767:
        for fi81 := 0; fi81 < 3; fi81++ {
            for fi82 := 0; fi82 < 2; fi82++ {
                v15 = 575
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
                var v83 uint = 776
                _ = fi82
                _ = v83
            }
            for fi84 := 0; fi84 < 6; fi84++ {
                v1 += (v3 | 127)
                fmt.Println(v4, v5, v4)
                v85 := GenS4{F0: -406, F1: 18446744073709551615}
                _ = fi84
                _ = v85
            }
            var v86 float32 = 0.1
            _ = fi81
            _ = v86
        }
    default:
        var v87 bool = (v15 == v15)
        _ = v87
    }
    if (298 == 0) {
        v6 = append(v6, 32768)
        var v88 uint64 = 42
        if false {
            var v89 uint = 4294967295
            _ = v89
        }
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
