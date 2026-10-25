package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint16; F1 bool; F2 int32 }
type GenS6 struct { F0 uint; F1 int64; F2 int64 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int8 = 42
    var v2 string = ""
    var v3 int16 = 222
    v4 := make(map[uint32]string)
    var v5 uint8 = 218
    var v6 int32 = 2147483647
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    v9 := v4[2147483647]
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v2 = "a"
            var v12 string = v2
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := []int{937, 1}
    v2 = ("hello" + v9)
    if ((1.5 * (-2.25 * float32((((1e10 / 3.14159) * -2.25) * (float64(float32((float64((-2.25 - 1e10)) - 0.1))) - ((float64(3.14159) - (float64((0.0 / 1.0)) + float64(float32(1.5)))) - 3.14159)))))) > 3.14159) {
        if false {
            var v14 int8 = (v1 & v1)
            var v15 int32 = (v6 >> ((1) & 63))
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                if true {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 uint64 = 42
                        var v19 uint = 0
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS1{F0: 0}
                    _ = v20
                }
                switch v9 {
                case "a":
                    v5 = 1
                }
                _ = len(v13)
                _ = fi16
            }
        }
        if (!(71 > -1)) {
            switch v2 {
            case "hello":
                for fi21 := 0; fi21 < 5; fi21++ {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if len(v13) > 0 {
                            v13[0] = -872
                        }
                        var v23 int32 = v6
                        var v24 uint16 = 2
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 int64 = 9223372036854775807
                    _ = fi21
                    _ = v25
                }
            case "foo bar":
                var v26 int32 = (v6 >> ((uint(v5)) & 63))
                _ = v26
            default:
                for fi27 := 0; fi27 < 8; fi27++ {
                    if false {
                        switch (fi27 >> ((4294967295) & 63)) {
                        case 32767:
                            if ((v6 + (v6 << ((18446744073709551615) & 63))) <= -879) {
                                v2 = v9
                            } else {
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    if false {
                                        v29 := GenS2{F0: 2147483647, F1: 2, F2: v9}
                                        _ = v29
                                    } else {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    }
                                    switch v2 {
                                    case "\n\t":
                                        fmt.Println(v2, fi28)
                                        v1 += v1
                                        v30 := v4[641]
                                        _ = v30
                                    case "":
                                        v31 := v4[65535]
                                        _ = v31
                                    default:
                                        v2 = v9
                                    }
                                    var v32 string = "hello"
                                    _ = fi28
                                    _ = v32
                                }
                            }
                        case 128:
                            v9 = v2
                        }
                    } else {
                        var v33 int = int(v6)
                        _ = v33
                    }
                    switch ((((v2 + (v2 + v2)) + v9) + v9) + ((("\n\t" + v2) + (((v9 + v2) + ("\n\t" + (v9 + "hello"))) + v9)) + "hello")) {
                    case "":
                        v1 += (v1 + 43)
                    default:
                        var v34 int32 = int32(fi27)
                        _ = v34
                    }
                    if true {
                        var v35 int8 = v1
                        if true {
                            v2 = v2
                        }
                        v3 = v3
                        _ = v35
                    } else {
                        var v36 uint = 2
                        for fi37 := 0; fi37 < 3; fi37++ {
                            v2 = v2
                            _ = fi37
                        }
                        _ = v36
                    }
                    _ = fi27
                }
            }
            v9 = v2
        } else {
            var v38 uint32 = 2147483648
            v3 = v3
            _ = v38
        }
        var v39 bool = (1 >= 2)
        _ = v39
    } else {
        v40 := GenS1{F0: 42}
        _ = v40
    }
    if false {
        fmt.Println(v6)
        switch -1 {
        case 127:
            switch 42 {
            case 0:
                var v41 float32 = 0.1
                _ = v41
            case 1:
                for fi42 := 0; fi42 < 12; fi42++ {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        if ((0.0 - float32(-2.25)) > (1.5 + 0.1)) {
                            var v44 uint32 = 1
                            _ = v44
                        } else {
                            switch fi42 {
                            case 4294967295:
                                v13 = append(v13, fi42)
                            case 1:
                                v6 -= v6
                            case 127:
                                var v45 int64 = 9223372036854775807
                                _ = v45
                            default:
                                if (2147483648 <= 551) {
                                    fmt.Println(v1, v5)
                                    var v46 bool = (291 > 18446744073709551615)
                                    _ = v46
                                }
                            }
                            v1 = (v1 / (((v1 / (((v1 + v1)) | 1))) | 1))
                        }
                        v4[448] = (v9 + v2)
                        var v47 int8 = (v1 / (((v1 - (v1 * 1))) | 1))
                        _ = fi43
                        _ = v47
                    }
                    _ = fi42
                }
            }
        case -9223372036854775808:
            if true {
                for fi48 := 0; fi48 < 2; fi48++ {
                    if true {
                        var v49 uint8 = v5
                        switch "\n\t" {
                        case "a":
                            var v50 uint = 92
                            _ = v50
                        case "hello":
                            if (!((-995 == v6) || ((false || true) && ((v6 > v6) || (!(true && true)))))) {
                                switch fi48 {
                                case 2147483647:
                                    v51 := GenS2{F0: (v6 - (v6 << ((157) & 63))), F1: 955, F2: v9}
                                    _ = v51
                                case 32768:
                                    var v52 int8 = v1
                                    _ = v52
                                case 1:
                                    fmt.Println(v6, v49, v3)
                                }
                            }
                        case "\n\t":
                            v13 = append(v13, fi48)
                        default:
                            v6 = (v6 * (v6 ^ int32(v3)))
                        }
                        for fi53 := 0; fi53 < 6; fi53++ {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            v6 -= (v6 - (v6 * -2147483648))
                            _ = fi53
                        }
                        _ = v49
                    } else {
                        v5 = v5
                        if true {
                            if false {
                                v9 = v9
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    if (255 > 0) {
                                        switch v2 {
                                        case "hello":
                                            switch fi54 {
                                            case -52:
                                                v5 += (v5 | v5)
                                            }
                                        default:
                                            var v55 uint8 = v5
                                            v56 := GenS3{F0: -1}
                                            v5 = 42
                                            v57 := GenS2{F0: v6, F1: 550, F2: "a"}
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 int64 = 42
                                        _ = v58
                                    } else {
                                        var v59 int64 = 1
                                        var v60 uint = 255
                                        if false {
                                            fmt.Println(fi48, v2, v6)
                                        } else {
                                            var v61 int8 = 37
                                            _ = v61
                                        }
                                        var v62 int32 = (v6 / (((v6 * v6)) | 1))
                                        _ = v59
                                        _ = v60
                                        _ = v62
                                    }
                                    if true {
                                        var v63 uint16 = 0
                                        v4[2] = v9
                                        _ = v63
                                    }
                                    _ = fi54
                                }
                                if (263 < 65535) {
                                    v3 = 32767
                                } else {
                                    v1 = v1
                                    _ = len(v4)
                                    var v64 int8 = v1
                                    _ = v64
                                }
                            }
                            for fi65 := 0; fi65 < 19; fi65++ {
                                if (-128 == (v1 >> ((0) & 63))) {
                                    var v66 int16 = (v3 >> ((1) & 63))
                                    _ = v66
                                } else {
                                    delete(v4, 2)
                                    var v67 uint32 = 2147483648
                                    if (65535 != 32767) {
                                        _ = len(v13)
                                        if ((v1 / ((v1) | 1)) <= 100) {
                                            var v68 uint8 = v5
                                            if (!false) {
                                                if (("hello" <= "foo bar") || true) {
                                                    var v69 uint8 = 127
                                                    var v70 int8 = int8(v67)
                                                    _ = v69
                                                    _ = v70
                                                }
                                                _ = len(v4)
                                                var v71 int64 = -9223372036854775808
                                                _ = v71
                                            }
                                            if ((v1 / ((v1) | 1)) <= (-(v1))) {
                                                var v72 bool = true
                                                var v73 float64 = (0.0 / 0.1)
                                                _ = v72
                                                _ = v73
                                            } else {
                                                v6 = 357
                                            }
                                            _ = v68
                                        } else {
                                            fmt.Println(fi65, v9)
                                            if false {
                                                var v74 int = fi48
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                                _ = v74
                                            }
                                            switch -885 {
                                            case 128:
                                                if (true || true) {
                                                    fmt.Println(v67)
                                                }
                                            case 2147483647:
                                                fmt.Println(fi48, v67)
                                            case 32767:
                                                v75 := GenS1{F0: 104}
                                                var v76 int32 = v6
                                                v77 := GenS2{F0: -2147483648, F1: 42, F2: v2}
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            }
                                        }
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    } else {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        v2 = v9
                                    }
                                    if (!(true || true)) {
                                        var v78 uint = 1
                                        fmt.Println(v1)
                                        v5 -= 255
                                        _ = v78
                                    } else {
                                        fmt.Println(v9)
                                    }
                                    var v79 int16 = v3
                                    _ = v67
                                    _ = v79
                                }
                                v6 -= v6
                                switch (fi65 ^ -9223372036854775808) {
                                case 4294967295:
                                    v80 := GenS4{F0: 1, F1: -566}
                                    _ = v80
                                case 42:
                                    if (65535 < 18446744073709551615) {
                                        v81 := GenS5{F0: 0, F1: (v3 != (v3 * (v3 >> ((18446744073709551615) & 63)))), F2: v6}
                                        var v82 uint32 = 565
                                        v9 = (v2 + v9)
                                        v2 = "a"
                                        _ = v81
                                        _ = v82
                                    }
                                case 9223372036854775807:
                                    if (!(v1 < (v1 << ((602) & 63)))) {
                                        _ = len(v13)
                                        var v83 uint16 = 32767
                                        var v84 uint16 = v83
                                        if (!(1.5 >= -2.25)) {
                                            fmt.Println(v5, v83, v5)
                                            var v85 int = -393
                                            v3 = int16(v6)
                                            _ = v85
                                        }
                                        _ = v83
                                        _ = v84
                                    } else {
                                        var v86 uint8 = v5
                                        if false {
                                            var v87 uint16 = 65535
                                            var v88 uint8 = v86
                                            var v89 uint64 = 42
                                            _ = v87
                                            _ = v88
                                            _ = v89
                                        }
                                        var v90 uint64 = 759
                                        if true {
                                            var v91 float32 = 1.5
                                            _ = v91
                                        } else {
                                            var v92 string = "\n\t"
                                            _ = v92
                                        }
                                        fmt.Println(v1)
                                        _ = v86
                                        _ = v90
                                    }
                                default:
                                    var v93 int16 = (v3 & 1)
                                    _ = v93
                                }
                                _ = fi65
                            }
                        } else {
                            var v94 uint = 489
                            v95 := GenS6{F0: 777, F1: 326, F2: 42}
                            _ = v94
                            _ = v95
                        }
                    }
                    _ = fi48
                }
            }
        case 42:
            v5 += 42
        default:
            var v96 uint16 = 1
            _ = v96
        }
    } else {
        _ = len(v4)
    }
    v3 -= int16(v1)
    v97 := GenS6{F0: 602, F1: 42, F2: 9223372036854775807}
    var v98 int8 = 43
    v2 = "\n\t"
    v1 = (v98 + v98)
    v98 += (-(v98))
    switch 9223372036854775807 {
    case -179:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    case 1:
        v99 := GenS7{F0: (-(v6))}
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    _ = v97
    _ = v98
}
