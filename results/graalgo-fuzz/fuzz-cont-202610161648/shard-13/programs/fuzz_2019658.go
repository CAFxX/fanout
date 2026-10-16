package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int8; F2 uint }

func main() {
    v1 := []int{}
    var v2 uint8 = 2
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            _ = len(v1)
            _ = len(v1)
            if false {
                v2 = (v2 + 205)
            }
            fmt.Println(v2)
            v6 := GenS1{F0: -1, F1: fi4}
            var v7 int16 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            if false {
                v2 = 38
                var v9 uint8 = v2
                var v10 uint16 = 204
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: -128, F1: -980}
            _ = fi8
            _ = v11
        }
        _ = fi4
    }
    v3[-991] = "hello"
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint64 = 18446744073709551615
        if (false && false) {
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := v3[73]
                _ = fi14
                _ = v15
            }
        } else {
            if (1 >= -463) {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint = 255
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 6; fi18++ {
                    fmt.Println(v2, v13)
                    v19 := GenS2{F0: 2147483648, F1: (fi18 << ((4294967295) & 63)), F2: v2}
                    delete(v3, (fi18 | 42))
                    _ = fi18
                    _ = v19
                }
            } else {
                fmt.Println(fi12)
                v2 -= (v2 & uint8(v13))
                fmt.Println(fi12, v13)
            }
            v13 += 589
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint16 = 0
                _ = fi20
                _ = v21
            }
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v13)
        case "a":
            switch "hello" {
            case "foo bar":
                v13 = 1
            default:
                for fi22 := 0; fi22 < 8; fi22++ {
                    if false {
                        if true {
                            var v23 uint32 = 2147483647
                            _ = v23
                        } else {
                            v24 := GenS1{F0: 42, F1: fi12}
                            switch "hello" {
                            case "a":
                                v25 := GenS1{F0: 42, F1: (-(fi22))}
                                _ = v25
                            case "foo bar":
                                v13 = (v13 + v13)
                            case "\n\t":
                                var v26 int32 = 20
                                _ = v26
                            }
                            _ = v24
                        }
                        if (9223372036854775807 <= -1) {
                            switch (fi12 << ((42) & 63)) {
                            case 128:
                                var v27 int32 = -1
                                _ = v27
                            case 1:
                                if false {
                                    v28 := GenS1{F0: 37, F1: -863}
                                    _ = v28
                                }
                            case 0:
                                v13 += (v13 * (v13 & (^(v13))))
                            }
                            if (611 < 2) {
                                v29 := GenS2{F0: 2, F1: -2147483648, F2: 255}
                                _ = v29
                            } else {
                                v13 = v13
                                var v30 int64 = 1
                                _ = v30
                            }
                            v31 := GenS1{F0: int8(v13), F1: fi22}
                            v13 = uint64(v2)
                            if true {
                                if true {
                                    v2 += v2
                                    switch "" {
                                    case "a":
                                        v1 = append(v1, 4294967295)
                                    case "":
                                        v13 = v13
                                    default:
                                        v2 += v2
                                    }
                                }
                                if false {
                                    if (42 > 0) {
                                        v13 = v13
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v32 := GenS3{F0: 2, F1: 2147483647}
                                        _ = v32
                                    } else {
                                        v33 := GenS4{F0: true, F1: 42, F2: 2}
                                        var v34 float64 = ((((float64(0.0) + (1.5 * ((float64((-2.25 - 1.5)) / 3.14159) / float64(1e10)))) + 1e10) + ((0.0 + 1.5) - ((3.14159 + 0.0) * 0.1))) * float64(1.5))
                                        v35 := v3[(fi12 / (((-(fi22))) | 1))]
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    v13 -= 18446744073709551615
                                } else {
                                    var v36 uint = 711
                                    if (1 <= 32767) {
                                        var v37 float32 = -2.25
                                        v13 -= (v13 - 2)
                                        v38 := GenS4{F0: false, F1: 0, F2: uint16(v36)}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v1 = append(v1, fi22)
                                    }
                                    v13 = v13
                                    switch "a" {
                                    case "":
                                        v36 -= 255
                                        var v39 int64 = -66
                                        _ = v39
                                    case "\n\t":
                                        v40 := GenS1{F0: -13, F1: fi22}
                                        var v41 uint32 = 65535
                                        v13 -= (v13 ^ 0)
                                        _ = v40
                                        _ = v41
                                    case "a":
                                        delete(v3, fi12)
                                    }
                                    var v42 uint16 = 29
                                    _ = v36
                                    _ = v42
                                }
                            } else {
                                v13 = (v13 & (v13 * 760))
                            }
                            var v43 uint32 = 18
                            if false {
                                v43 = (v43 >> ((1) & 63))
                            }
                            _ = v31
                            _ = v43
                        }
                        v13 = 208
                    } else {
                        if (false && (2 == 18446744073709551615)) {
                            v44 := GenS5{F0: 18446744073709551615}
                            _ = v44
                        }
                    }
                    _ = fi22
                }
            }
        case "foo bar":
            for fi45 := 0; fi45 < 9; fi45++ {
                v13 += (^(v13))
                _ = fi45
            }
        }
        _ = fi12
        _ = v13
    }
    v2 = (v2 + (v2 >> ((400) & 63)))
    if true {
        if ((!false) && (32767 > 65535)) {
            v2 = 127
            switch 532 {
            case -1:
                var v46 int16 = 755
                _ = v46
            case 0:
                var v47 int32 = 2147483647
                _ = v47
            case -952:
                fmt.Println(v2)
            default:
                var v48 float32 = 0.1
                _ = v48
            }
            for fi49 := 0; fi49 < 19; fi49++ {
                switch "" {
                case "foo bar":
                    for fi50 := 0; fi50 < 11; fi50++ {
                        if true {
                            v2 += (v2 / (((v2 ^ 159)) | 1))
                            v2 = (v2 * v2)
                            v2 -= 1
                        } else {
                            if (!false) {
                                v2 -= v2
                            } else {
                                var v51 float32 = -2.25
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "" {
                                case "\n\t":
                                    v52 := GenS6{F0: 1, F1: ("a" + ("foo bar" + "\n\t"))}
                                    _ = v52
                                case "foo bar":
                                    switch ("hello" + "") {
                                    case "\n\t":
                                        if (65535 != 2147483648) {
                                            var v53 int64 = 966
                                            delete(v3, fi49)
                                            v51 += 1.5
                                            _ = v53
                                        } else {
                                            var v54 int32 = 2147483647
                                            v55 := GenS1{F0: 42, F1: (fi49 - (-(fi50)))}
                                            _ = len(v3)
                                            _ = v54
                                            _ = v55
                                        }
                                        fmt.Println(v2, v2)
                                        delete(v3, fi50)
                                    case "hello":
                                        v56 := v3[-9223372036854775808]
                                        _ = v56
                                    default:
                                        if (true || (false || (!(0 != 241)))) {
                                            v2 += (v2 - (v2 / (((v2 << ((18446744073709551615) & 63))) | 1)))
                                        } else {
                                            fmt.Println(v2)
                                        }
                                    }
                                }
                                _ = v51
                            }
                        }
                        v2 += (v2 << ((4294967295) & 63))
                        var v57 int64 = 992
                        _ = fi50
                        _ = v57
                    }
                }
                _ = fi49
            }
        }
        for fi58 := 0; fi58 < 10; fi58++ {
            v2 += (^(v2))
            _ = fi58
        }
        var v59 int = int(v2)
        _ = v59
    } else {
        if (0 > 37) {
            for fi60 := 0; fi60 < 4; fi60++ {
                v3[fi60] = ("foo bar" + ("a" + "foo bar"))
                if false {
                    if (-2.25 <= (3.14159 + 0.1)) {
                        fmt.Println(v2, fi60)
                        switch "foo bar" {
                        case "hello":
                            v61 := GenS4{F0: false, F1: -1, F2: 42}
                            _ = v61
                        case "foo bar":
                            var v62 bool = false
                            _ = v62
                        case "a":
                            var v63 uint = 0
                            _ = v63
                        }
                        if true {
                            var v64 float64 = float64(1.5)
                            var v65 uint32 = 65535
                            _ = v64
                            _ = v65
                        } else {
                            if (42 <= -96) {
                                var v66 uint8 = (v2 >> ((4294967295) & 63))
                                if ((!(true && true)) || (!false)) {
                                    v66 += 1
                                    v66 += 2
                                    v67 := GenS7{F0: "foo bar"}
                                    _ = v67
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi60 ^ -9223372036854775808)
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if true {
                                        v66 += v66
                                        switch fi60 {
                                        case -9223372036854775808:
                                            for fi68 := 0; fi68 < 2; fi68++ {
                                                var v69 uint32 = 4294967295
                                                fmt.Println(fi60)
                                                _ = fi68
                                                _ = v69
                                            }
                                        case 1000:
                                            fmt.Println(v66)
                                        }
                                        v1 = append(v1, (fi60 * (fi60 - (fi60 ^ 727))))
                                    }
                                }
                                for fi70 := 0; fi70 < 7; fi70++ {
                                    var v71 uint32 = 2
                                    if true {
                                        v2 -= 248
                                        var v72 int = (-(fi60))
                                        _ = v72
                                    } else {
                                        if ((604 <= 842) || (v71 >= 2)) {
                                            var v73 uint64 = 1
                                            v73 += (v73 | v73)
                                            _ = v73
                                        }
                                        v74 := GenS4{F0: false, F1: 38, F2: 0}
                                        v2 = v66
                                        _ = v74
                                    }
                                    _ = fi70
                                    _ = v71
                                }
                                _ = v66
                            }
                            var v75 int64 = 42
                            v2 = v2
                            _ = v75
                        }
                    } else {
                        if (42 > 0) {
                            v76 := GenS5{F0: 1}
                            var v77 uint16 = 32767
                            _ = v76
                            _ = v77
                        } else {
                            fmt.Println(v2)
                        }
                        if (!(true && ((704 != 2) && (1000 < 512)))) {
                            v2 = (v2 & (v2 / ((127) | 1)))
                            var v78 uint64 = 461
                            _ = v78
                        }
                        switch ("hello" + ("" + ("hello" + "\n\t"))) {
                        case "a":
                            if (fi60 != (fi60 / (((fi60 - fi60)) | 1))) {
                                if (!true) {
                                    switch (-(fi60)) {
                                    case -2147483648:
                                        for fi79 := 0; fi79 < 14; fi79++ {
                                            v80 := GenS8{F0: 32767, F1: 0, F2: 65535}
                                            if true {
                                                v1 = append(v1, fi79)
                                                v2 = v2
                                            }
                                            _ = fi79
                                            _ = v80
                                        }
                                    case -573:
                                        v2 += 122
                                    }
                                    var v81 int = fi60
                                    _ = len(v3)
                                    _ = v81
                                }
                                _ = len(v1)
                            }
                        default:
                            var v82 float32 = 3.14159
                            _ = v82
                        }
                    }
                    for fi83 := 0; fi83 < 15; fi83++ {
                        v84 := GenS2{F0: 451, F1: fi60, F2: v2}
                        fmt.Println(fi83, v2)
                        _ = fi83
                        _ = v84
                    }
                }
                _ = fi60
            }
        }
        var v85 string = "a"
        _ = v85
    }
    fmt.Println(v2)
    v2 = 128
    v86 := GenS4{F0: ((0 > 656) || false), F1: -1, F2: 2}
    var v87 uint32 = 424
    var v88 float32 = (1.5 * (0.0 + ((-2.25 + (float32(float64(1.5)) / float32((-2.25 * 0.1)))) + -2.25)))
    var v89 uint = 18446744073709551615
    v90 := GenS1{F0: 100, F1: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
}
