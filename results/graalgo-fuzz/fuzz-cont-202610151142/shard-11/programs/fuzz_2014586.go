package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint32; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int; F2 bool }
type GenS11 struct { F0 uint8; F1 string }
type GenS12 struct { F0 int64; F1 uint32 }
type GenS13 struct { F0 uint; F1 int8; F2 int8 }

func main() {
    var v1 int16 = 1
    v2 := make(map[int]int)
    var v3 uint = 4294967295
    var v4 uint = v3
    var v5 int64 = 27
    var v6 int = 2147483647
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint32 = 2147483648
            var v10 string = "hello"
            v10 = ""
            _ = len(v2)
            var v11 bool = ((1e10 < 3.14159) && true)
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        if ("" == (("" + ("foo bar" + ("\n\t" + ""))) + "a")) {
            if true {
                switch "foo bar" {
                case "foo bar":
                    v3 = 1
                case "hello":
                    var v12 bool = (!(float32(0.1) == 1e10))
                    _ = v12
                default:
                    switch (fi7 % ((9223372036854775807) | 1)) {
                    case 128:
                        if (42 < 100) {
                            v13 := []int{}
                            if (735 == 1) {
                                v14 := GenS1{F0: v6}
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v15 int16 = v1
                                _ = v14
                                _ = v15
                            } else {
                                if true {
                                    v5 -= -1
                                    v16 := GenS2{F0: 695}
                                    _ = v16
                                }
                                if (42 == -1) {
                                    switch (^(v6)) {
                                    case 1:
                                        var v17 uint32 = 65535
                                        _ = v17
                                    case 32767:
                                        if false {
                                            v4 += 4294967295
                                            for fi18 := 0; fi18 < 14; fi18++ {
                                                switch -961 {
                                                case 42:
                                                    if true {
                                                        _ = len(v13)
                                                    } else {
                                                        v19 := GenS3{F0: (614 != 736), F1: v5, F2: 126}
                                                        var v20 bool = ((!false) || (!(!(!false))))
                                                        _ = v19
                                                        _ = v20
                                                    }
                                                case 32767:
                                                    switch fi7 {
                                                    case 32768:
                                                        fmt.Println(fi18, v3, v5)
                                                    case 128:
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    _ = len(v13)
                                                                    var v21 string = ""
                                                                    _ = v21
                                                                }
                                                                v22 := GenS1{F0: 4294967295}
                                                                v23 := GenS2{F0: 2}
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            v24 := GenS4{F0: -32768}
                                                            _ = v24
                                                        } else {
                                                            var v25 int = fi18
                                                            _ = v25
                                                        }
                                                    case -554:
                                                        fmt.Println(fi18)
                                                    default:
                                                        fmt.Println(v3, v3, v5)
                                                    }
                                                case 32768:
                                                    var v26 uint32 = 854
                                                    _ = v26
                                                default:
                                                    if ((3.14159 + 1.5) >= 1e10) {
                                                        var v27 int8 = -70
                                                        v28 := GenS5{F0: false, F1: 2, F2: (fi7 * v6)}
                                                        fmt.Println(v27)
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                }
                                                var v29 uint32 = 590
                                                if false {
                                                    if true {
                                                        var v30 uint64 = 976
                                                        var v31 uint16 = 125
                                                        v32 := GenS2{F0: (v29 / ((2147483647) | 1))}
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    v4 += 0
                                                }
                                                _ = fi18
                                                _ = v29
                                            }
                                            v5 = (v5 + v5)
                                        } else {
                                            var v33 int32 = -2147483648
                                            v33 = v33
                                            fmt.Println(fi7)
                                            _ = v33
                                        }
                                    case 4294967295:
                                        delete(v2, -1)
                                    }
                                }
                            }
                            for fi34 := 0; fi34 < 8; fi34++ {
                                var v35 float32 = 0.0
                                v6 += (fi34 & fi7)
                                _ = fi34
                                _ = v35
                            }
                            _ = v13
                        } else {
                            v36 := GenS6{F0: 65535}
                            fmt.Println(v6, v4, v5)
                            _ = v36
                        }
                    default:
                        _ = len(v2)
                    }
                }
                v4 -= v3
                for fi37 := 0; fi37 < 1; fi37++ {
                    delete(v2, fi7)
                    _ = fi37
                }
            }
            v38 := GenS7{F0: v6, F1: 127}
            for fi39 := 0; fi39 < 12; fi39++ {
                v40 := GenS8{F0: int(v3), F1: 0}
                v1 -= 188
                v1 -= (v1 >> ((v4) & 63))
                v41 := []int{}
                delete(v2, fi7)
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v38
        }
        _ = fi7
    }
    var v42 int8 = 0
    v43 := GenS5{F0: (!false), F1: 474, F2: int(v4)}
    if true {
        _ = len(v2)
        if ((1e10 < ((1e10 + float64((1e10 + (float32(1.5) - 1.5)))) * (-2.25 - 1.5))) || false) {
            var v44 int64 = v5
            if (!(42 == 65535)) {
                v44 = -9223372036854775808
            } else {
                var v45 string = ""
                var v46 int64 = 42
                _ = v45
                _ = v46
            }
            _ = v44
        } else {
            v47 := GenS9{F0: 42}
            _ = v47
        }
        v6 -= (v6 & 127)
    }
    switch -1 {
    case 32767:
        if (376 != 2) {
            switch v6 {
            case 9223372036854775807:
                var v48 int8 = 100
                _ = v48
            case -176:
                if ("hello" <= "hello") {
                    if false {
                        fmt.Println(v6)
                    } else {
                        fmt.Println(v3, v1, v4)
                    }
                    v49 := []int{-245, (v6 ^ (v6 >> (((^(v4))) & 63)))}
                    _ = v49
                }
            default:
                var v50 uint64 = 2
                _ = v50
            }
            fmt.Println(v42, v6, v5)
        }
    case 273:
        for fi51 := 0; fi51 < 4; fi51++ {
            for fi52 := 0; fi52 < 20; fi52++ {
                var v53 uint64 = 42
                v53 = 4294967295
                v54 := GenS1{F0: (fi52 - -2147483648)}
                var v55 uint32 = 65535
                _ = fi52
                _ = v53
                _ = v54
                _ = v55
            }
            _ = fi51
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v56 int32 = -911
        _ = v56
    case "a":
        var v57 uint64 = 42
        _ = v57
    case "\n\t":
        var v58 float32 = (1e10 + (1.5 * 3.14159))
        _ = v58
    default:
        fmt.Println(v1)
    }
    delete(v2, (v6 / (((-(v6))) | 1)))
    for fi59 := 0; fi59 < 6; fi59++ {
        for fi60 := 0; fi60 < 20; fi60++ {
            switch "foo bar" {
            case "\n\t":
                var v61 string = "\n\t"
                _ = v61
            }
            v62 := []int{fi60, fi60, (-(fi59))}
            _ = fi60
            _ = v62
        }
        v63 := GenS10{F0: 42, F1: (fi59 | fi59), F2: (((v1 << (((^(v4))) & 63)) > v1) || false)}
        _ = fi59
        _ = v63
    }
    for fi64 := 0; fi64 < 13; fi64++ {
        v65 := GenS11{F0: 255, F1: "\n\t"}
        var v66 float32 = float32(3.14159)
        v42 += 1
        _ = fi64
        _ = v65
        _ = v66
    }
    v42 = v42
    v67 := GenS4{F0: int16(v5)}
    switch (v6 % (((v6 & (v6 % ((v6) | 1)))) | 1)) {
    case -9223372036854775808:
        for fi68 := 0; fi68 < 2; fi68++ {
            v4 += (v4 >> ((v4) & 63))
            _ = fi68
        }
    case 42:
        if (false && (594 > 42)) {
            v6 = int(v42)
            var v69 uint64 = 278
            for fi70 := 0; fi70 < 4; fi70++ {
                var v71 uint32 = 4294967295
                _ = fi70
                _ = v71
            }
            _ = v69
        } else {
            for fi72 := 0; fi72 < 17; fi72++ {
                for fi73 := 0; fi73 < 20; fi73++ {
                    var v74 bool = (127 == 127)
                    if (("hello" + ("foo bar" + "")) < ("hello" + "")) {
                        var v75 uint8 = 1
                        v76 := []int{-9223372036854775808}
                        var v77 int = (v6 & v6)
                        _ = v75
                        _ = v76
                        _ = v77
                    } else {
                        switch ("\n\t" + ("hello" + "hello")) {
                        case "foo bar":
                            var v78 int32 = 2147483647
                            _ = v78
                        }
                        v5 += v5
                        var v79 int = -309
                        v79 = 2147483647
                        fmt.Println(v4)
                        var v80 uint16 = 198
                        var v81 int64 = -268
                        _ = v79
                        _ = v80
                        _ = v81
                    }
                    v82 := []int{128, (fi72 / (((fi72 ^ fi72)) | 1))}
                    _ = fi73
                    _ = v74
                    _ = v82
                }
                fmt.Println(v1, v4, v6)
                _ = fi72
            }
            v83 := GenS3{F0: (!true), F1: (-(v5)), F2: 2}
            switch "foo bar" {
            case "":
                v84 := GenS6{F0: 255}
                _ = v84
            default:
                var v85 int64 = 0
                _ = v85
            }
            _ = v83
        }
    default:
        var v86 float64 = 1e10
        _ = v86
    }
    var v87 bool = ((v5 + v5) <= (^(v5)))
    for fi88 := 0; fi88 < 19; fi88++ {
        var v89 int16 = 32767
        _ = fi88
        _ = v89
    }
    if v87 {
        if (v87 && (("\n\t" > "hello") || ((260 == 953) && (32767 <= 0)))) {
            if v87 {
                var v90 int32 = 272
                _ = v90
            } else {
                v5 = v5
            }
            if v87 {
                v87 = v87
            } else {
                for fi91 := 0; fi91 < 14; fi91++ {
                    v92 := v2[42]
                    switch "\n\t" {
                    case "":
                        v93 := GenS4{F0: v1}
                        _ = v93
                    case "foo bar":
                        v94 := GenS3{F0: true, F1: (v5 * v5), F2: 222}
                        _ = v94
                    case "hello":
                        switch (fi91 & (fi91 / (((v6 + int(v5))) | 1))) {
                        case 128:
                            v1 += 1000
                        case 32767:
                            switch "hello" {
                            case "hello":
                                switch (fi91 - v92) {
                                case -742:
                                    v95 := GenS12{F0: (v5 * (v5 + v5)), F1: 2147483647}
                                    _ = v95
                                default:
                                    var v96 int64 = v5
                                    _ = v96
                                }
                            }
                        }
                    }
                    _ = fi91
                    _ = v92
                }
                var v97 float64 = (1e10 + 3.14159)
                _ = v97
            }
            v98 := GenS7{F0: v6, F1: 1}
            _ = v98
        }
        if (true || v87) {
            var v99 bool = v87
            if (-1 == v42) {
                var v100 uint8 = 2
                switch "\n\t" {
                case "":
                    v4 = v4
                case "a":
                    var v101 int64 = (v5 >> ((v4) & 63))
                    _ = v101
                default:
                    v42 = v42
                }
                v6 = -514
                _ = v100
            } else {
                if (true || v99) {
                    for fi102 := 0; fi102 < 12; fi102++ {
                        v103 := []int{-545}
                        _ = fi102
                        _ = v103
                    }
                    v104 := GenS13{F0: 18446744073709551615, F1: int8(v3), F2: (v42 << ((4294967295) & 63))}
                    _ = v104
                }
            }
            var v105 float64 = 3.14159
            _ = v105
            _ = v99
        }
    } else {
        v106 := GenS8{F0: v6, F1: v1}
        _ = v106
    }
    v107 := GenS13{F0: 4294967295, F1: v42, F2: (-(v42))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v67
    _ = v87
}
