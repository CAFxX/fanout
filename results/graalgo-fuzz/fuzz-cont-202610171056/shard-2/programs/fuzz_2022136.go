package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int8 }

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]string)
    var v3 bool = false
    v4 := []int{}
    var v5 bool = v3
    var v6 uint16 = 466
    v7 := GenS1{F0: ("a" + ("hello" + ((("hello" + "") + "") + ""))), F1: 672, F2: (!(0 != -508))}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int32 = 1
            var v11 uint16 = (v6 ^ v6)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v6 = 786
            _ = fi13
        }
        v2[156] = "a"
        for fi14 := 0; fi14 < 14; fi14++ {
            if (42 != 100) {
                v2[2147483648] = ("\n\t" + ((("" + "a") + ((("hello" + ("foo bar" + "foo bar")) + "\n\t") + "hello")) + "foo bar"))
                v5 = v5
            }
            _ = fi14
        }
        _ = fi12
    }
    if true {
        if ((float64(1.5) - (3.14159 + ((1.5 * (3.14159 / 0.1)) - 0.1))) == (1.5 + 0.0)) {
            for fi15 := 0; fi15 < 4; fi15++ {
                if false {
                    var v16 int32 = -971
                    var v17 int = 630
                    v1 = v17
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 bool = (1.5 == -2.25)
                        var v20 int64 = 42
                        v4 = append(v4, 0)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                if (1.5 >= float32(-2.25)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v21 uint8 = 34
                    var v22 int8 = 6
                    _ = v21
                    _ = v22
                }
                v23 := GenS2{F0: 255}
                _ = fi15
                _ = v23
            }
            var v24 int64 = -1
            v25 := GenS3{F0: 65535, F1: -87, F2: "\n\t"}
            _ = v24
            _ = v25
        } else {
            if len(v4) > 0 {
                v4[0] = (v1 / (((v1 >> ((861) & 63))) | 1))
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                if (!(v3 && v5)) {
                    v5 = true
                    for fi27 := 0; fi27 < 13; fi27++ {
                        var v28 bool = v5
                        fmt.Println(v6)
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint16 = 0
                    _ = v29
                } else {
                    v6 = (v6 / ((65535) | 1))
                }
                v3 = v3
                var v30 string = ("hello" + ("" + "hello"))
                _ = fi26
                _ = v30
            }
            fmt.Println(v1, v3, v6)
        }
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 int = (fi31 * (fi31 | (v1 ^ (fi31 | (v1 | 1)))))
            for fi33 := 0; fi33 < 20; fi33++ {
                if ((42 <= 42) && (float64(3.14159) != 0.0)) {
                    if (1.5 <= float32((3.14159 - (float64(1.5) - 1.5)))) {
                        v1 -= v1
                        if false {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v1)
                            case "":
                                v4 = append(v4, -1)
                            }
                        } else {
                            var v34 uint32 = 0
                            var v35 uint = 18446744073709551615
                            v36 := GenS1{F0: "", F1: 4294967295, F2: (1e10 > float32((float64(((((0.0 - ((0.1 * (float32((1e10 / 3.14159)) + 0.0)) + 0.1)) - 3.14159) * (-2.25 * 1e10)) - (1.5 + ((float32(float64(((3.14159 + 0.1) * 1.5))) - float32((-2.25 + 0.0))) / 3.14159)))) + ((3.14159 * 3.14159) - 0.0))))}
                            v1 = v1
                            v3 = v5
                            var v37 uint = v35
                            var v38 uint64 = 435
                            switch (fi33 << ((v35) & 63)) {
                            case 9223372036854775807:
                                v35 -= (^(v37))
                            default:
                                v39 := GenS2{F0: 0}
                                _ = v39
                            }
                            v34 += (v34 & v34)
                            v5 = (1 != 1)
                            var v40 float64 = -2.25
                            if false {
                                if (89 == 1) {
                                    v41 := GenS1{F0: "", F1: v38, F2: true}
                                    v42 := GenS1{F0: "\n\t", F1: v38, F2: (v5 || (32767 <= fi31))}
                                    _ = v41
                                    _ = v42
                                }
                            } else {
                                var v43 int32 = 1
                                _ = v43
                            }
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v40
                        }
                        var v44 uint32 = 65535
                        _ = v44
                    }
                }
                if false {
                    fmt.Println(v6, v3)
                } else {
                    if len(v4) > 0 {
                        v4[0] = 127
                    }
                }
                v5 = (60 < -2147483648)
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
        v6 = v6
        switch v1 {
        case 9223372036854775807:
            var v45 uint16 = (v6 << ((18446744073709551615) & 63))
            _ = v45
        case -416:
            var v46 uint32 = 627
            _ = v46
        case 0:
            if (v1 == 32768) {
                if v3 {
                    v3 = false
                    _ = len(v2)
                    var v47 uint8 = 127
                    _ = v47
                }
                if (!v3) {
                    if v5 {
                        var v48 float32 = 1e10
                        v6 = v6
                        delete(v2, 2147483648)
                        _ = v48
                    }
                    v1 = (v1 * v1)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                var v49 float64 = (-2.25 / (float64((1e10 / 1.0)) + float64(3.14159)))
                v49 += 1.5
                _ = v49
            }
        default:
            fmt.Println(v3)
        }
    }
    var v50 uint16 = (^(v6))
    fmt.Println(v50)
    var v51 int8 = -63
    if v3 {
        for fi52 := 0; fi52 < 16; fi52++ {
            fmt.Println(v6, fi52, v50)
            for fi53 := 0; fi53 < 2; fi53++ {
                v6 = (^(v50))
                if v3 {
                    switch ("hello" + ((((("foo bar" + "") + "") + "") + ("hello" + "")) + "hello")) {
                    case "":
                        var v54 bool = v3
                        _ = len(v2)
                        fmt.Println(v50, v6, v50)
                        _ = len(v2)
                        _ = v54
                    case "\n\t":
                        v51 = (v51 ^ v51)
                    default:
                        var v55 int8 = (v51 ^ (v51 - int8(fi53)))
                        _ = v55
                    }
                    if (v6 > (v50 >> ((4294967295) & 63))) {
                        v5 = v3
                        var v56 int8 = v51
                        _ = v56
                    }
                    v57 := GenS4{F0: -73}
                    switch (v1 ^ v1) {
                    case -751:
                        switch (fi53 / ((fi53) | 1)) {
                        case 474:
                            v1 += int(v50)
                        case 32767:
                            var v58 int16 = 42
                            _ = v58
                        case 262:
                            var v59 float64 = -2.25
                            _ = v59
                        default:
                            var v60 float64 = 1.5
                            _ = v60
                        }
                    case 127:
                        v5 = (-277 <= -989)
                    case -9223372036854775808:
                        switch fi52 {
                        case 43:
                            v50 = (v50 ^ v6)
                        case 484:
                            if true {
                                switch -722 {
                                case 2147483647:
                                    v5 = ((1 <= 32767) && (1.5 < ((1.5 - 3.14159) / (-2.25 - (((1.5 + 0.0) + 3.14159) - 1.5)))))
                                case 4294967295:
                                    if ((-(fi52)) <= 2147483647) {
                                        if (794 != 255) {
                                            fmt.Println(fi53)
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                var v61 int16 = 0
                                                _ = v61
                                            case "":
                                                v4 = append(v4, (v1 / ((127) | 1)))
                                            case "\n\t":
                                                fmt.Println(fi52, v51, fi52)
                                                v62 := v2[2147483648]
                                                v63 := v2[65535]
                                                v64 := GenS3{F0: 2147483647, F1: v1, F2: v63}
                                                v65 := GenS3{F0: 1, F1: (^(v1)), F2: v63}
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                            v66 := GenS2{F0: 0}
                                            _ = v66
                                        }
                                        if v5 {
                                            if v5 {
                                                var v67 int32 = -1
                                                var v68 int8 = (-(v51))
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        fmt.Println(v1, v51, v5)
                                    }
                                case 128:
                                    if (!(42 > -786)) {
                                        switch "hello" {
                                        case "":
                                            v4 = append(v4, (fi52 & (fi52 - 0)))
                                        }
                                        var v69 string = ("\n\t" + ("\n\t" + "\n\t"))
                                        _ = v69
                                    } else {
                                        v1 += fi53
                                        var v70 uint32 = 42
                                        switch fi53 {
                                        case 858:
                                            var v71 string = "a"
                                            _ = v71
                                        case 168:
                                            v72 := GenS5{F0: 117}
                                            _ = v72
                                        }
                                        _ = v70
                                    }
                                }
                                if v3 {
                                    var v73 uint16 = 65535
                                    v74 := GenS2{F0: 0}
                                    v50 = 714
                                    fmt.Println(v5, v73, v51)
                                    var v75 int32 = -2147483648
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    fmt.Println(v50, v5, v51)
                                    var v76 string = ("hello" + "")
                                    v76 = (v76 + (v76 + "foo bar"))
                                    _ = v76
                                }
                                fmt.Println(v51, v3)
                                var v77 int16 = 1000
                                v78 := GenS3{F0: 98, F1: v1, F2: ((("\n\t" + "\n\t") + (("a" + ("a" + "a")) + (("foo bar" + ("" + "a")) + "hello"))) + "foo bar")}
                                var v79 int8 = (v51 % ((v51) | 1))
                                if (v51 < (^(v51))) {
                                    if v3 {
                                        v51 = -92
                                        v77 = v77
                                        fmt.Println(v50)
                                    }
                                    switch ("foo bar" + (("\n\t" + "a") + ("hello" + "a"))) {
                                    case "":
                                        delete(v2, 42)
                                    }
                                } else {
                                    v80 := GenS4{F0: v77}
                                    _ = v80
                                }
                                if true {
                                    switch ("a" + "a") {
                                    case "hello":
                                        if (42 > 0) {
                                            var v81 int32 = -2147483648
                                            v82 := GenS5{F0: 2}
                                            var v83 string = "a"
                                            _ = v81
                                            _ = v82
                                            _ = v83
                                        }
                                    case "":
                                        var v84 uint8 = 2
                                        _ = v84
                                    case "a":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        fmt.Println(v5, v77)
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v6)
                                    } else {
                                        v77 = (v77 % ((v77) | 1))
                                    }
                                }
                                if ("a" <= "\n\t") {
                                    var v85 uint64 = 2
                                    v5 = true
                                    _ = v85
                                } else {
                                    var v86 uint16 = (v6 + (^(v6)))
                                    _ = v86
                                }
                                v6 = v6
                                _ = v77
                                _ = v78
                                _ = v79
                            } else {
                                fmt.Println(v51)
                            }
                        default:
                            v87 := GenS2{F0: 1}
                            _ = v87
                        }
                    }
                    v5 = (0 != 0)
                    v88 := GenS6{F0: 2}
                    _ = v57
                    _ = v88
                } else {
                    v1 += fi53
                }
                var v89 int = (fi53 | (-(fi52)))
                _ = fi53
                _ = v89
            }
            var v90 bool = v3
            _ = fi52
            _ = v90
        }
    } else {
        if (-320 > -2147483648) {
            var v91 int32 = 42
            switch (("foo bar" + "") + "") {
            case "hello":
                fmt.Println(v3, v1)
            case "\n\t":
                if (v91 >= v91) {
                    v50 += (v50 & (v50 - 0))
                }
            }
            _ = v91
        }
        var v92 int = v1
        _ = v92
    }
    for fi93 := 0; fi93 < 16; fi93++ {
        switch (v1 + -9223372036854775808) {
        case 73:
            fmt.Println(v6, v51, v50)
        case 1:
            v5 = (((-2.25 + (((0.0 - (0.1 / 1e10)) * (0.1 * -2.25)) + 3.14159)) / float64(-2.25)) < (0.1 - 0.0))
        case 261:
            if v3 {
                v94 := GenS1{F0: (("\n\t" + "a") + "\n\t"), F1: 2, F2: (v51 == v51)}
                _ = v94
            }
        }
        for fi95 := 0; fi95 < 11; fi95++ {
            v3 = v5
            v96 := GenS3{F0: 820, F1: fi95, F2: (("a" + ("a" + "")) + "\n\t")}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi95
            _ = v96
        }
        v97 := GenS7{F0: ("a" + "hello"), F1: v51}
        _ = fi93
        _ = v97
    }
    v98 := GenS2{F0: 255}
    for fi99 := 0; fi99 < 20; fi99++ {
        v1 = 1
        fmt.Println(v51, v6, v50)
        var v100 int8 = 11
        _ = fi99
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v98
}
