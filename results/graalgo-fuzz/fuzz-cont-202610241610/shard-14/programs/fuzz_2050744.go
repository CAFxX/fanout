package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 int; F2 int }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int; F1 int64; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int8 = -1
    var v2 float32 = (float32(3.14159) + 3.14159)
    v3 := make(map[uint32]string)
    v4 := []uint32{511, 249, 4294967295, 0}
    v5 := make(map[uint32]string)
    var v6 uint = 0
    for fi7 := 0; fi7 < 1; fi7++ {
        v1 = (v1 / (((v1 & v1)) | 1))
        _ = fi7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v6)
        var v9 int64 = 9223372036854775807
        _ = fi8
        _ = v9
    }
    var v10 int32 = -2147483648
    switch "foo bar" {
    case "hello":
        if true {
            _ = len(v4)
            var v11 int32 = v10
            _ = v11
        }
    case "a":
        v2 -= (0.0 - (-2.25 + float32(-2.25)))
    case "foo bar":
        var v12 uint8 = 148
        _ = v12
    }
    var v13 uint = 18446744073709551615
    v10 += (v10 | (^(v10)))
    if ((3.14159 + 0.0) == (0.1 - 1.5)) {
        if (9223372036854775807 <= 1) {
            v10 = (^(v10))
            v1 -= (v1 / (((v1 & (v1 + v1))) | 1))
            if (-2.25 <= (((3.14159 / 3.14159) + 0.0) + (0.1 + 1.5))) {
                if (false || true) {
                    v14 := GenS1{F0: 0, F1: v1}
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v2 -= v2
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v6 -= (v13 - v13)
                            _ = fi16
                        }
                        _ = fi15
                    }
                    if (32767 >= 32767) {
                        v1 = 0
                        v17 := GenS2{F0: 208, F1: -9223372036854775808, F2: 1000}
                        v18 := GenS3{F0: 32768, F1: int(v13), F2: 4294967295}
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                }
            } else {
                var v19 string = ("a" + "foo bar")
                v20 := GenS3{F0: 4294967295, F1: 0, F2: 946}
                _ = v19
                _ = v20
            }
        } else {
            var v21 uint32 = 170
            if true {
                var v22 int8 = v1
                if (-512 != 32767) {
                    v10 = -2147483648
                    for fi23 := 0; fi23 < 3; fi23++ {
                        if ((false || (!true)) || false) {
                            if true {
                                v2 += float32(1e10)
                            } else {
                                v24 := GenS4{F0: 32767, F1: 480}
                                if true {
                                    v22 -= 100
                                    if true {
                                        for fi25 := 0; fi25 < 2; fi25++ {
                                            v13 = v6
                                            switch ((("hello" + ("a" + "")) + ("hello" + "")) + "a") {
                                            case "":
                                                fmt.Println(v6)
                                            case "\n\t":
                                                v22 = 42
                                            }
                                            v22 += v22
                                            _ = fi25
                                        }
                                        v21 = (v21 << ((v13) & 63))
                                        v21 = 42
                                    }
                                    var v26 int16 = -1
                                    _ = v26
                                } else {
                                    delete(v5, v21)
                                    v27 := GenS5{F0: (fi23 - fi23), F1: -679, F2: 65535}
                                    v22 -= (v1 / ((-128) | 1))
                                    _ = v27
                                }
                                if (1e10 > v2) {
                                    var v28 uint64 = 0
                                    if (799 < (v21 % (((^(v21))) | 1))) {
                                        if (128 != 199) {
                                            v29 := GenS1{F0: 1, F1: v1}
                                            fmt.Println(v22)
                                            _ = v29
                                        } else {
                                            v10 = v10
                                            v30 := GenS5{F0: (fi23 | fi23), F1: 1, F2: 153}
                                            _ = v30
                                        }
                                    }
                                    _ = v28
                                } else {
                                    if false {
                                        var v31 uint64 = 522
                                        for fi32 := 0; fi32 < 11; fi32++ {
                                            fmt.Println(v1)
                                            _ = fi32
                                        }
                                        fmt.Println(v13)
                                        _ = v31
                                    }
                                    var v33 float32 = -2.25
                                    _ = v33
                                }
                                _ = v24
                            }
                        }
                        _ = fi23
                    }
                    if (!true) {
                        v13 += v13
                        v13 = v6
                    } else {
                        var v34 float32 = v2
                        v35 := GenS6{F0: (v10 & (v10 + (v10 % (((v10 * (v10 % (((-(v10))) | 1)))) | 1))))}
                        for fi36 := 0; fi36 < 8; fi36++ {
                            for fi37 := 0; fi37 < 19; fi37++ {
                                fmt.Println(v6)
                                _ = fi37
                            }
                            v5[v21] = ("" + ("hello" + "foo bar"))
                            _ = fi36
                        }
                        _ = v34
                        _ = v35
                    }
                }
                if false {
                    if false {
                        delete(v3, (v21 / (((^(v21))) | 1)))
                        if false {
                            var v38 uint = v13
                            _ = v38
                        }
                        v39 := GenS7{F0: 128, F1: int8(v6), F2: true}
                        _ = v39
                    } else {
                        fmt.Println(v21, v2, v1)
                        delete(v3, (^(v21)))
                    }
                    if (!true) {
                        for fi40 := 0; fi40 < 13; fi40++ {
                            var v41 bool = true
                            v42 := GenS4{F0: 555, F1: 4294967295}
                            var v43 uint = (v13 << (((v6 - uint(v21))) & 63))
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        if ((((false || (-224 == 1)) || false) && (v21 == 620)) || ((!(!false)) && true)) {
                            v22 = (v22 & (v1 % ((127) | 1)))
                        } else {
                            if (((1e10 / 0.1) >= (v2 - (-2.25 / v2))) && true) {
                                v10 = v10
                            }
                            if len(v4) > 0 {
                                v4[0] = 636
                            }
                        }
                        v22 = 127
                    } else {
                        if true {
                            v22 -= (v1 / (((v1 + int8(v13))) | 1))
                            v10 += v10
                            switch -381 {
                            case -2147483648:
                                var v44 int = 1
                                _ = v44
                            case 1:
                                delete(v3, (v21 - (^(v21))))
                            }
                        } else {
                            var v45 int16 = -907
                            v46 := GenS6{F0: v10}
                            if (2147483648 == 1) {
                                for fi47 := 0; fi47 < 12; fi47++ {
                                    v4 = append(v4, (v21 & 42))
                                    _ = fi47
                                }
                                v2 += float32(0.1)
                                v5[(v21 - v21)] = "hello"
                            } else {
                                for fi48 := 0; fi48 < 5; fi48++ {
                                    var v49 bool = (v10 != v10)
                                    fmt.Println(v6)
                                    v13 += v6
                                    _ = fi48
                                    _ = v49
                                }
                                var v50 uint16 = 401
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    for fi52 := 0; fi52 < 11; fi52++ {
                                        var v53 int32 = int32(v13)
                                        var v54 float64 = 1e10
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                    for fi55 := 0; fi55 < 4; fi55++ {
                                        var v56 int8 = (v22 / ((v22) | 1))
                                        v57 := GenS8{F0: 655, F1: ((false && true) && false), F2: uint32(v50)}
                                        v45 = (v45 & v45)
                                        if true {
                                            v58 := GenS4{F0: v45, F1: 657}
                                            var v59 uint16 = v50
                                            v10 = 42
                                            _ = v58
                                            _ = v59
                                        }
                                        if (1 < -129) {
                                            if (int8(v45) == (v22 / ((v1) | 1))) {
                                                v60 := GenS6{F0: v10}
                                                switch (("\n\t" + ("a" + "")) + ("" + ("hello" + "a"))) {
                                                case "foo bar":
                                                    if (v10 != (v10 - (v10 & -912))) {
                                                        if (!false) {
                                                            var v61 uint32 = (v21 * v21)
                                                            v50 = v50
                                                            _ = v61
                                                        }
                                                    }
                                                }
                                                _ = v60
                                            }
                                            var v62 uint8 = 207
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v63 uint64 = 853
                                                fmt.Println(v63, v2)
                                                _ = v63
                                            }
                                            _ = v62
                                        } else {
                                            v56 = (v1 >> (((^(v13))) & 63))
                                        }
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = fi51
                                }
                                _ = v50
                            }
                            _ = v45
                            _ = v46
                        }
                        for fi64 := 0; fi64 < 20; fi64++ {
                            v65 := GenS9{F0: 42}
                            if false {
                                _ = len(v4)
                                var v66 int64 = 1
                                _ = v66
                            }
                            if (!(!true)) {
                                if (-624 == 32767) {
                                    v2 += (v2 * 1e10)
                                    if (!true) {
                                        v21 -= v21
                                    } else {
                                        _ = len(v3)
                                        var v67 uint64 = 18446744073709551615
                                        fmt.Println(fi64, v6, v67)
                                        _ = v67
                                    }
                                    if false {
                                        if (!(true && ((true || (!((v21 <= v21) && false))) || (fi64 != 180)))) {
                                            if true {
                                                v68 := GenS10{F0: 2147483647, F1: 127}
                                                _ = len(v4)
                                                _ = v68
                                            } else {
                                                fmt.Println(v1, v10)
                                            }
                                        } else {
                                            var v69 int = fi64
                                            _ = v69
                                        }
                                        for fi70 := 0; fi70 < 2; fi70++ {
                                            v2 = 3.14159
                                            var v71 int = fi70
                                            v21 = v21
                                            _ = fi70
                                            _ = v71
                                        }
                                    } else {
                                        if (-9223372036854775808 < -9223372036854775808) {
                                            var v72 int16 = -32768
                                            if len(v4) > 0 {
                                                v4[0] = (v21 * (v21 ^ 65535))
                                            }
                                            if ((v21 << ((v13) & 63)) >= v21) {
                                                for fi73 := 0; fi73 < 8; fi73++ {
                                                    switch "hello" {
                                                    case "":
                                                        if true {
                                                            v74 := v5[(^(v21))]
                                                            v1 = v22
                                                            switch fi73 {
                                                            case -2147483648:
                                                                delete(v5, (v21 | (v21 & (^(v21)))))
                                                            case -492:
                                                                v10 = v10
                                                            case 293:
                                                                var v75 int8 = -110
                                                                _ = v75
                                                            default:
                                                                if true {
                                                                    v2 = 1.5
                                                                }
                                                            }
                                                            _ = v74
                                                        }
                                                    case "hello":
                                                        if (!true) {
                                                            var v76 uint = 65535
                                                            var v77 string = "hello"
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    }
                                                    fmt.Println(v72, v6)
                                                    if true {
                                                        v78 := GenS3{F0: 846, F1: fi64, F2: -719}
                                                        delete(v5, 868)
                                                        _ = v78
                                                    }
                                                    _ = fi73
                                                }
                                                var v79 uint32 = (v21 * uint32(v6))
                                                _ = v79
                                            }
                                            _ = v72
                                        }
                                        v4 = append(v4, v21)
                                    }
                                } else {
                                    _ = len(v4)
                                    if false {
                                        var v80 float64 = ((0.0 + 1.5) + 1.5)
                                        _ = v80
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    v5[(v21 & 65535)] = ""
                                }
                                _ = len(v5)
                            } else {
                                switch (fi64 & (fi64 + fi64)) {
                                case 9223372036854775807:
                                    for fi81 := 0; fi81 < 17; fi81++ {
                                        v21 = v21
                                        delete(v5, v21)
                                        var v82 uint8 = 1
                                        _ = fi81
                                        _ = v82
                                    }
                                case 1:
                                    v83 := GenS11{F0: 2}
                                    _ = v83
                                }
                            }
                            _ = fi64
                            _ = v65
                        }
                    }
                }
                _ = v22
            } else {
                delete(v5, v21)
                v21 = (^(v21))
                var v84 float32 = (v2 + v2)
                _ = v84
            }
            switch 127 {
            case -9223372036854775808:
                v85 := GenS10{F0: (v21 ^ (^(v21))), F1: 0}
                _ = v85
            case 32768:
                var v86 int32 = (^(v10))
                _ = v86
            }
            _ = v21
        }
    }
    var v87 uint64 = 42
    for fi88 := 0; fi88 < 7; fi88++ {
        for fi89 := 0; fi89 < 4; fi89++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                var v90 float64 = (1e10 / -2.25)
                _ = v90
            }
            if ((65535 >= 32767) && (!true)) {
                v10 += (^(v10))
                var v91 string = "\n\t"
                _ = v91
            } else {
                var v92 float32 = (v2 * -2.25)
                var v93 bool = (!(false && true))
                var v94 int16 = -401
                switch fi89 {
                case 404:
                    v95 := GenS12{F0: 0, F1: 42}
                    _ = v95
                default:
                    if v93 {
                        v93 = v93
                        var v96 int8 = 42
                        var v97 float32 = v2
                        _ = v96
                        _ = v97
                    } else {
                        v87 = (v87 >> (((v6 & 65535)) & 63))
                        if (((243 != 193) && (2 != 0)) && v93) {
                            switch fi88 {
                            case 609:
                                fmt.Println(v92, v94)
                            case 2147483647:
                                v98 := v5[42]
                                _ = v98
                            default:
                                _ = len(v5)
                            }
                        }
                    }
                    var v99 int64 = -723
                    _ = v99
                }
                _ = v92
                _ = v93
                _ = v94
            }
            _ = fi89
        }
        var v100 uint64 = (v87 ^ 853)
        _ = fi88
        _ = v100
    }
    v13 += (^(v13))
    v87 = 42
    var v101 int32 = 2147483647
    if false {
        var v102 int = -9223372036854775808
        fmt.Println(v1, v10, v101)
        var v103 bool = (208 >= 32767)
        _ = v102
        _ = v103
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v10 += (-(v101))
    var v104 uint32 = 65535
    var v105 uint16 = 836
    if (v2 > 0.0) {
        fmt.Println(v10, v6)
        for fi106 := 0; fi106 < 1; fi106++ {
            v4 = append(v4, (v104 / ((v104) | 1)))
            for fi107 := 0; fi107 < 1; fi107++ {
                v4 = append(v4, v104)
                if false {
                    fmt.Println(v2)
                    fmt.Println(v101, v2)
                    v2 += ((float32(-2.25) / 1.0) - 1e10)
                }
                v10 += (v101 & v101)
                _ = fi107
            }
            var v108 bool = (-9223372036854775808 == -9223372036854775808)
            _ = fi106
            _ = v108
        }
        var v109 float64 = (0.1 - 1e10)
        _ = v109
    } else {
        if false {
            var v110 int64 = -1
            _ = v110
        }
        v111 := GenS4{F0: 32767, F1: 4294967295}
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v101:", v101)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v10
    _ = v101
    _ = v104
    _ = v105
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v87
}
