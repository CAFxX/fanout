package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 uint16; F2 int8 }
type GenS6 struct { F0 int; F1 int64; F2 uint8 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint16; F1 int64; F2 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0(p0 string) int16 {
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: -1}
            switch fi7 {
            case -1:
                var v9 int16 = -1
                _ = v9
            case 42:
                p0 = p0
            case 4294967295:
                var v10 string = p0
                _ = v10
            }
            v11 := GenS2{F0: 32767}
            _ = fi7
            _ = v11
            _ = v8
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            v5 += (v5 % (((v5 - 4294967295)) | 1))
            var v13 uint16 = 0
            var v14 int64 = -462
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    var v15 int8 = -101
    _ = p0
    _ = v15
    _ = v5
    return -529
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 18446744073709551615
    var v3 int32 = 2147483647
    var v4 bool = (!(!false))
    v16 := []int{4294967295, 1, 128}
    v2 += v2
    var v17 int8 = -1
    v18 := GenS3{F0: "\n\t", F1: (v17 >> ((v2) & 63))}
    var v19 uint32 = 554
    v2 -= 255
    if v4 {
        switch 128 {
        case 32767:
            v20 := GenS3{F0: "", F1: (v17 << ((v2) & 63))}
            _ = v20
        case -9223372036854775808:
            var v21 float32 = 0.1
            _ = v21
        case 462:
            if (!v4) {
                if ((((float32((1.5 * float64(0.1))) * -2.25) + (((((1.5 + 3.14159) / 3.14159) - 3.14159) / float32(0.1)) * ((1e10 - -2.25) - -2.25))) * 3.14159) >= (1e10 / 0.1)) {
                    fmt.Println(v4, v17)
                    if (!v4) {
                        var v22 int32 = v3
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v24 := GenS4{F0: -32768}
                            var v25 float32 = 3.14159
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        v4 = (((-2.25 + 0.0) + (1.5 / 1e10)) == 0.1)
                        _ = v22
                    }
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v2 -= 480
                        _ = fi26
                    }
                }
                var v27 int32 = v3
                _ = v27
            } else {
                var v28 uint8 = 1
                _ = v28
            }
        }
        var v29 int64 = 0
        var v30 uint32 = (v19 * 65535)
        _ = v29
        _ = v30
    } else {
        var v31 int16 = -731
        if ((v17 * -66) >= -1) {
            fmt.Println(v17, v19)
            fmt.Println(v3, v3, v19)
        } else {
            fmt.Println(v4)
        }
        _ = v31
    }
    var v32 uint16 = 255
    v16 = append(v16, 9223372036854775807)
    if false {
        for fi33 := 0; fi33 < 13; fi33++ {
            for fi34 := 0; fi34 < 5; fi34++ {
                switch "hello" {
                case "\n\t":
                    v35 := v1[fi33]
                    _ = v35
                case "hello":
                    var v36 uint8 = 0
                    _ = v36
                case "foo bar":
                    if ((v32 / (((v32 >> ((987) & 63))) | 1)) > (v32 + 287)) {
                        if v4 {
                            switch fi34 {
                            case -2147483648:
                                var v37 uint32 = 4294967295
                                _ = v37
                            }
                        }
                    } else {
                        if (1e10 > 0.0) {
                            if true {
                                fmt.Println(fi34, fi34)
                            }
                            fmt.Println(v2, fi34)
                        }
                    }
                default:
                    if v4 {
                        v38 := GenS2{F0: (v32 + (^(v32)))}
                        if false {
                            v17 = v17
                        }
                        _ = v38
                    } else {
                        v39 := v1[fi34]
                        _ = v39
                    }
                }
                switch "\n\t" {
                case "":
                    fmt.Println(v17, v4)
                default:
                    fmt.Println(v2, fi34, v3)
                }
                fmt.Println(v4, v17)
                _ = fi34
            }
            switch fi33 {
            case 127:
                fmt.Println(fi33, v32)
            default:
                for fi40 := 0; fi40 < 18; fi40++ {
                    v32 += (v32 + v32)
                    var v41 uint32 = v19
                    v4 = false
                    var v42 int32 = (v3 - (-(v3)))
                    if (!false) {
                        v17 -= (v17 | v17)
                        var v43 int32 = (v3 ^ v42)
                        var v44 uint16 = (v32 % (((v32 | (v32 % ((0) | 1)))) | 1))
                        _ = v43
                        _ = v44
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            var v45 uint32 = v19
            _ = fi33
            _ = v45
        }
        v3 += v3
        if (!(2 >= 9)) {
            v46 := GenS1{F0: 32768}
            var v47 uint32 = (v19 ^ 429)
            _ = v46
            _ = v47
        } else {
            v48 := GenS1{F0: 606}
            if v4 {
                switch "a" {
                case "":
                    var v49 uint64 = 997
                    _ = v49
                }
            } else {
                switch "" {
                case "\n\t":
                    fmt.Println(v3, v4, v4)
                case "hello":
                    for fi50 := 0; fi50 < 9; fi50++ {
                        fmt.Println(v17, v2)
                        var v51 float64 = 1e10
                        v2 -= uint(v19)
                        _ = fi50
                        _ = v51
                    }
                case "foo bar":
                    var v52 string = ""
                    _ = v52
                default:
                    _ = len(v1)
                }
                v19 += (v19 - (v19 % ((v19) | 1)))
            }
            _ = v48
        }
    } else {
        switch -2147483648 {
        case 42:
            var v53 uint32 = (^(v19))
            _ = v53
        default:
            var v54 int8 = -128
            _ = v54
        }
        switch 42 {
        case 42:
            if len(v16) > 0 {
                v16[0] = 2147483647
            }
        case 128:
            v55 := GenS5{F0: 96, F1: (^(v32)), F2: v17}
            _ = v55
        case 2147483647:
            v3 = 42
        default:
            v4 = (793 < 946)
        }
        if false {
            v19 = v19
        } else {
            var v56 int64 = 0
            if (363 >= 1) {
                if v4 {
                    var v57 int8 = 100
                    var v58 int16 = 1
                    switch -135 {
                    case 127:
                        var v59 int = -1
                        _ = v59
                    case 32767:
                        _ = len(v1)
                    case 43:
                        if (9223372036854775807 != v56) {
                            _ = len(v1)
                        } else {
                            if true {
                                v60 := GenS3{F0: ("a" + "hello"), F1: v57}
                                if v4 {
                                    v61 := GenS6{F0: 9223372036854775807, F1: v56, F2: 104}
                                    v62 := GenS7{F0: (^(v3)), F1: (v17 > (v17 * (v57 / ((int8(v32)) | 1))))}
                                    _ = v61
                                    _ = v62
                                }
                                var v63 uint8 = 128
                                _ = v60
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 3; fi64++ {
                                fmt.Println(v19, v19, v3)
                                for fi65 := 0; fi65 < 16; fi65++ {
                                    switch ((("\n\t" + "foo bar") + "a") + "foo bar") {
                                    case "":
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    }
                                    fmt.Println(fi65, v17, fi65)
                                    v66 := GenS2{F0: v32}
                                    _ = fi65
                                    _ = v66
                                }
                                _ = fi64
                            }
                            v67 := GenS1{F0: -9223372036854775808}
                            _ = v67
                        }
                    default:
                        fmt.Println(v3, v2)
                    }
                    _ = v57
                    _ = v58
                } else {
                    v3 = (v3 * 1)
                    for fi68 := 0; fi68 < 2; fi68++ {
                        for fi69 := 0; fi69 < 12; fi69++ {
                            v2 = (v2 - v2)
                            v56 = (v56 >> (((v2 & 217)) & 63))
                            if false {
                                v32 += (v32 ^ 2)
                                v16 = append(v16, fi69)
                            }
                            _ = fi69
                        }
                        _ = fi68
                    }
                }
                v3 -= (-(v3))
            } else {
                if false {
                    var v70 int16 = 42
                    _ = v70
                } else {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    if v4 {
                        switch "" {
                        case "a":
                            if v4 {
                                var v71 uint32 = v19
                                v16 = append(v16, 4294967295)
                                switch 4294967295 {
                                case -2147483648:
                                    v72 := GenS7{F0: -801, F1: (18446744073709551615 < 731)}
                                    _ = v72
                                case -452:
                                    v73 := GenS8{F0: v56, F1: ((^(v32)) > 901)}
                                    _ = v73
                                }
                                _ = v71
                            } else {
                                v32 += v32
                                var v74 int16 = 0
                                _ = v74
                            }
                        case "foo bar":
                            fmt.Println(v17)
                        case "":
                            v75 := v1[-516]
                            _ = v75
                        }
                        if true {
                            for fi76 := 0; fi76 < 10; fi76++ {
                                fmt.Println(v32)
                                for fi77 := 0; fi77 < 17; fi77++ {
                                    v78 := GenS8{F0: -1, F1: (0 >= 135)}
                                    if ((true && false) && v4) {
                                        v79 := GenS9{F0: 2, F1: 775, F2: (fi77 % (((fi77 & (-(fi76)))) | 1))}
                                        var v80 int16 = 32767
                                        if v4 {
                                            var v81 uint64 = 362
                                            v4 = (!("a" == "foo bar"))
                                            var v82 uint16 = (^(v32))
                                            if len(v16) > 0 {
                                                v16[0] = 32768
                                            }
                                            _ = v81
                                            _ = v82
                                        } else {
                                            v3 -= (v3 / ((v3) | 1))
                                        }
                                        _ = v79
                                        _ = v80
                                    }
                                    _ = fi77
                                    _ = v78
                                }
                                var v83 int = fi76
                                _ = fi76
                                _ = v83
                            }
                            v56 -= (v56 % ((int64(v32)) | 1))
                            v84 := GenS10{F0: 122}
                            _ = v84
                        }
                        for fi85 := 0; fi85 < 2; fi85++ {
                            for fi86 := 0; fi86 < 15; fi86++ {
                                switch (fi85 >> ((v2) & 63)) {
                                case 172:
                                    v3 = 62
                                case -822:
                                    var v87 uint64 = 84
                                    _ = v87
                                }
                                switch fi86 {
                                case -2147483648:
                                    if v4 {
                                        var v88 bool = (0.0 == 3.14159)
                                        _ = v88
                                    }
                                case -44:
                                    var v89 uint64 = 0
                                    _ = v89
                                case -1:
                                    fmt.Println(v32, v19)
                                default:
                                    v56 += (v56 & (v56 | (v56 % ((-1) | 1))))
                                }
                                _ = fi86
                            }
                            _ = fi85
                        }
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
                v90 := GenS2{F0: (^(v32))}
                _ = v90
            }
            _ = v56
        }
    }
    v32 += v32
    switch -579 {
    case 1:
        v1[1] = ((-2.25 - 3.14159) < float32(0.1))
    case 657:
        for fi91 := 0; fi91 < 16; fi91++ {
            var v92 uint32 = v19
            _ = fi91
            _ = v92
        }
    case 90:
        switch -966 {
        case 1:
            for fi93 := 0; fi93 < 13; fi93++ {
                switch (fi93 - int(v3)) {
                case 0:
                    v16 = append(v16, fi93)
                case 1:
                    for fi94 := 0; fi94 < 17; fi94++ {
                        delete(v1, (fi93 + (fi93 % (((fi94 - 960)) | 1))))
                        v3 = (v3 | (^(v3)))
                        if v4 {
                            var v95 uint16 = 1
                            if (32767 > v95) {
                                v17 -= (v17 * (v17 % ((42) | 1)))
                                var v96 float32 = (3.14159 / ((float32(-2.25) - 3.14159) - (0.0 + 0.1)))
                                _ = len(v1)
                                _ = v96
                            } else {
                                v32 = uint16(v2)
                                v97 := GenS6{F0: fi93, F1: 42, F2: 0}
                                v32 += (v32 | 99)
                                v98 := GenS6{F0: -9223372036854775808, F1: -452, F2: 140}
                                _ = v97
                                _ = v98
                            }
                            if v4 {
                                if true {
                                    fmt.Println(v2)
                                    if (-52 == 447) {
                                        switch ("\n\t" + ("foo bar" + "")) {
                                        case "hello":
                                            fmt.Println(fi93)
                                        }
                                        v19 -= 2147483647
                                        var v99 int = fi93
                                        _ = v99
                                    }
                                    v19 = 2147483647
                                    v32 += v95
                                }
                                if (!v4) {
                                    v100 := GenS9{F0: (v95 - (^(v32))), F1: -9223372036854775808, F2: fi94}
                                    _ = v100
                                } else {
                                    v95 -= 255
                                }
                                v101 := GenS1{F0: fi93}
                                _ = v101
                            }
                            _ = v95
                        }
                        var v102 uint32 = v19
                        _ = fi94
                        _ = v102
                    }
                default:
                    var v103 float32 = 0.0
                    _ = v103
                }
                v104 := v1[(fi93 & (fi93 - (fi93 >> ((uint(v19)) & 63))))]
                _ = fi93
                _ = v104
            }
        case 131:
            v105 := GenS11{F0: 0}
            _ = v105
        }
    default:
        for fi106 := 0; fi106 < 14; fi106++ {
            _ = len(v1)
            _ = fi106
        }
    }
    for fi107 := 0; fi107 < 5; fi107++ {
        for fi108 := 0; fi108 < 19; fi108++ {
            v109 := GenS12{F0: v17, F1: v32, F2: 127}
            if v4 {
                v16 = append(v16, fi108)
            } else {
                var v110 float32 = (-2.25 - (-2.25 + 0.0))
                var v111 uint = v2
                _ = v110
                _ = v111
            }
            _ = fi108
            _ = v109
        }
        _ = fi107
    }
    v19 = (v19 | v19)
    for fi112 := 0; fi112 < 13; fi112++ {
        for fi113 := 0; fi113 < 19; fi113++ {
            _ = len(v1)
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = fi113
        }
        v32 = 674
        var v114 float64 = ((1.5 + 1e10) + 1e10)
        _ = fi112
        _ = v114
    }
    fmt.Println(v17, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
