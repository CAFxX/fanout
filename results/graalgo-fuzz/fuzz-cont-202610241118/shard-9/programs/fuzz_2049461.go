package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32; F1 uint16; F2 int16 }
func genFn0() uint32 {
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = 42
        switch (fi5 ^ (fi5 ^ (fi5 & 4294967295))) {
        case 2147483647:
            var v7 uint8 = 128
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v8 string = "a"
    v9 := []int{4294967295, 1}
    _ = v4
    _ = v8
    _ = v9
    return 42
}
func genFn1() string {
    for fi10 := 0; fi10 < 14; fi10++ {
        if (759 >= -1) {
            var v11 uint8 = 127
            for fi12 := 0; fi12 < 20; fi12++ {
                if false {
                    v11 += 18
                    if true {
                        if true {
                            v11 += 127
                        }
                        v13 := GenS1{F0: v11}
                        if false {
                            v11 = v11
                        } else {
                            v14 := GenS2{F0: 1, F1: (0 != 255), F2: 1}
                            _ = v14
                        }
                        _ = v13
                    } else {
                        fmt.Println(v11, fi12)
                        switch ("foo bar" + (("a" + "hello") + ("" + "a"))) {
                        case "\n\t":
                            if (1.5 == 0.0) {
                                switch int(v11) {
                                case -46:
                                    var v15 int8 = -128
                                    switch fi10 {
                                    case -9223372036854775808:
                                        var v16 float64 = (3.14159 + 1.5)
                                        _ = v16
                                    case 9223372036854775807:
                                        v15 = v15
                                    default:
                                        var v17 int = (fi12 % ((1) | 1))
                                        _ = v17
                                    }
                                    var v18 int64 = -9223372036854775808
                                    _ = v15
                                    _ = v18
                                default:
                                    var v19 float32 = 0.0
                                    _ = v19
                                }
                                var v20 int32 = -2147483648
                                if true {
                                    switch "a" {
                                    case "foo bar":
                                        v21 := GenS1{F0: (v11 % ((135) | 1))}
                                        _ = v21
                                    default:
                                        v20 -= v20
                                    }
                                } else {
                                    v22 := GenS3{F0: "foo bar"}
                                    fmt.Println(fi10)
                                    v23 := []int{}
                                    _ = v22
                                    _ = v23
                                }
                                _ = v20
                            }
                        default:
                            v24 := GenS4{F0: 4294967295}
                            _ = v24
                        }
                    }
                    var v25 float64 = float64(1e10)
                    _ = v25
                }
                v26 := GenS1{F0: (^(v11))}
                switch "foo bar" {
                case "":
                    v27 := []int{(fi12 >> ((42) & 63)), fi12, fi10}
                    _ = v27
                case "a":
                    var v28 int8 = -1
                    _ = v28
                case "foo bar":
                    v11 -= 161
                default:
                    v29 := GenS3{F0: "\n\t"}
                    _ = v29
                }
                v11 += (v11 % (((^(v11))) | 1))
                v30 := GenS5{F0: 0}
                var v31 bool = ((42 <= 36) || false)
                _ = fi12
                _ = v26
                _ = v30
                _ = v31
            }
            _ = v11
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS2{F0: 244, F1: ((!true) || false), F2: -1}
            fmt.Println(fi10, fi10)
            v34 := []int{(^(fi10)), (fi10 | fi10), (fi10 - fi10), fi32}
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint = 2
            _ = fi35
            _ = v36
        }
        _ = fi10
    }
    var v37 uint16 = 255
    _ = v37
    return "hello"
}

func main() {
    v1 := []int{32767, 0, 383}
    var v2 uint = 538
    var v3 uint32 = 42
    v38 := GenS6{F0: 1, F1: genFn0()}
    for fi39 := 0; fi39 < 10; fi39++ {
        if (42 < -128) {
            var v40 uint8 = 127
            _ = v40
        } else {
            if (9223372036854775807 <= (fi39 + fi39)) {
                fmt.Println(v2, fi39, v2)
                var v41 int = fi39
                v42 := GenS2{F0: 0, F1: true, F2: 63}
                _ = v41
                _ = v42
            }
            for fi43 := 0; fi43 < 14; fi43++ {
                v44 := GenS7{F0: 1, F1: 42, F2: 127}
                _ = fi43
                _ = v44
            }
        }
        if (65535 < 1) {
            v2 = v2
        }
        if true {
            var v45 uint16 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v45
        } else {
            v3 = (v3 + v3)
            v1 = append(v1, fi39)
        }
        _ = fi39
    }
    if true {
        var v46 uint32 = genFn0()
        fmt.Println(v2, v3)
        _ = v46
    } else {
        v3 = v3
    }
    if (!(1 == -128)) {
        var v47 string = "foo bar"
        v48 := GenS5{F0: 445}
        _ = v47
        _ = v48
    }
    _ = len(v1)
    if (100 != -1) {
        if true {
            v2 -= v2
            if (-1 < 1) {
                for fi49 := 0; fi49 < 18; fi49++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v50 := GenS8{F0: 255, F1: -114}
                    _ = fi49
                    _ = v50
                }
            } else {
                var v51 float64 = 0.1
                if false {
                    for fi52 := 0; fi52 < 5; fi52++ {
                        v2 = v2
                        if (9223372036854775807 < fi52) {
                            fmt.Println(fi52, v2)
                            for fi53 := 0; fi53 < 12; fi53++ {
                                var v54 int = fi52
                                var v55 uint16 = 499
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        } else {
                            switch "" {
                            case "a":
                                var v56 int8 = 42
                                _ = v56
                            case "":
                                fmt.Println(v3, v51)
                            case "hello":
                                v2 = v2
                            default:
                                v2 -= v2
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = len(v1)
                        }
                        _ = fi52
                    }
                }
                fmt.Println(v2, v51)
                _ = v51
            }
        } else {
            var v57 float64 = 0.1
            for fi58 := 0; fi58 < 6; fi58++ {
                var v59 float32 = ((((3.14159 + ((float32(-2.25) + 3.14159) + (0.0 - (1.5 * -2.25)))) + 1e10) * 0.0) * ((float32(3.14159) + (1e10 / 3.14159)) + 0.0))
                var v60 uint = (v2 % (((v2 + v2)) | 1))
                if (!false) {
                    var v61 bool = true
                    v62 := GenS6{F0: fi58, F1: (v3 & v3)}
                    var v63 uint32 = genFn0()
                    _ = v61
                    _ = v62
                    _ = v63
                } else {
                    v57 -= (-2.25 - v57)
                    v2 = uint(v3)
                }
                _ = fi58
                _ = v59
                _ = v60
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v57
        }
        v64 := GenS8{F0: 29, F1: 39}
        _ = v64
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        fmt.Println(v2, v3, fi65)
        v3 -= genFn0()
        _ = fi65
    }
    if (!true) {
        var v66 string = "a"
        fmt.Println(v2, v66)
        _ = v66
    }
    if true {
        v67 := GenS9{F0: v2}
        if ("\n\t" != "a") {
            var v68 float32 = -2.25
            if (((1 >= 1) && (!true)) || false) {
                for fi69 := 0; fi69 < 5; fi69++ {
                    var v70 bool = false
                    v71 := GenS10{F0: (v3 << ((896) & 63)), F1: 42}
                    v68 = ((((((1.5 * 0.1) - 0.0) + (v68 * float32(1.5))) / v68) + (-2.25 + (float32(3.14159) + -2.25))) + v68)
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                var v72 uint16 = 423
                _ = v72
            }
            v3 = v3
            _ = v68
        }
        _ = v67
    } else {
        for fi73 := 0; fi73 < 11; fi73++ {
            var v74 bool = (!true)
            fmt.Println(v74, v3, v2)
            _ = fi73
            _ = v74
        }
        v75 := GenS11{F0: 1, F1: 42, F2: 1}
        _ = v75
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v76 uint = (^(v2))
        _ = v76
    } else {
        if (!(9223372036854775807 != 42)) {
            for fi77 := 0; fi77 < 3; fi77++ {
                if true {
                    for fi78 := 0; fi78 < 15; fi78++ {
                        switch 32768 {
                        case 9223372036854775807:
                            v2 -= (v2 >> (((v2 | v2)) & 63))
                        case -779:
                            fmt.Println(v2, fi78, v2)
                        case -1:
                            if len(v1) > 0 {
                                v1[0] = (fi78 ^ (fi77 + (fi78 / (((fi77 * 32768)) | 1))))
                            }
                        default:
                            v79 := GenS4{F0: v2}
                            _ = v79
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi78
                    }
                }
                for fi80 := 0; fi80 < 8; fi80++ {
                    v2 -= 42
                    v2 = v2
                    var v81 int8 = 127
                    _ = fi80
                    _ = v81
                }
                if false {
                    if (968 > 42) {
                        fmt.Println(v2, v2, fi77)
                    } else {
                        var v82 int8 = 59
                        for fi83 := 0; fi83 < 17; fi83++ {
                            switch 32768 {
                            case 32767:
                                var v84 uint = 4294967295
                                _ = v84
                            case -111:
                                var v85 uint64 = 0
                                _ = v85
                            }
                            var v86 uint8 = 67
                            _ = fi83
                            _ = v86
                        }
                        _ = v82
                    }
                    if (!false) {
                        for fi87 := 0; fi87 < 6; fi87++ {
                            fmt.Println(v2, fi87, v2)
                            _ = len(v1)
                            var v88 uint8 = 0
                            _ = fi87
                            _ = v88
                        }
                    } else {
                        v3 -= 2147483648
                        switch 0 {
                        case 346:
                            var v89 string = (("" + "foo bar") + "\n\t")
                            _ = v89
                        case -1:
                            v3 += v3
                        case 0:
                            for fi90 := 0; fi90 < 10; fi90++ {
                                fmt.Println(fi90, fi77, v2)
                                if false {
                                    fmt.Println(v2, v3, fi77)
                                    if (true && (-32768 == 62)) {
                                        v91 := GenS12{F0: 42}
                                        switch fi77 {
                                        case -733:
                                            switch -599 {
                                            case 32767:
                                                var v92 float64 = float64(0.1)
                                                _ = v92
                                            case 32768:
                                                var v93 int = (fi90 - fi90)
                                                _ = v93
                                            }
                                        case 480:
                                            if false {
                                                fmt.Println(fi90)
                                            } else {
                                                v3 = 2147483648
                                                var v94 int8 = 42
                                                v2 -= v2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v94 = v94
                                                _ = v94
                                            }
                                        case 32768:
                                            v95 := GenS13{F0: (v3 / ((v3) | 1))}
                                            _ = v95
                                        default:
                                            if true {
                                                if false {
                                                    v96 := GenS1{F0: 42}
                                                    _ = len(v1)
                                                    _ = v96
                                                } else {
                                                    if (fi90 >= 32767) {
                                                        var v97 int32 = 0
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v97
                                                    }
                                                    if (((v3 * (v3 >> (((v2 << ((2) & 63))) & 63))) > v3) && true) {
                                                        v1 = append(v1, fi90)
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi90 >> ((1) & 63))
                                                        }
                                                    }
                                                    var v98 uint32 = 203
                                                    _ = v98
                                                }
                                            } else {
                                                if (true || (1e10 < 3.14159)) {
                                                    var v99 uint = v2
                                                    _ = v99
                                                }
                                            }
                                        }
                                        v3 += 2147483647
                                        if (!(!false)) {
                                            var v100 int8 = 1
                                            v3 = (v3 + v3)
                                            _ = v100
                                        } else {
                                            v3 -= (v3 + genFn0())
                                            v101 := GenS14{F0: ""}
                                            fmt.Println(fi77)
                                            _ = v101
                                        }
                                        switch (fi77 % ((-93) | 1)) {
                                        case -9223372036854775808:
                                            switch ("hello" + ("hello" + ("foo bar" + ""))) {
                                            case "\n\t":
                                                v2 = uint(v3)
                                            }
                                            var v102 uint32 = v3
                                            _ = v102
                                        case 1:
                                            v2 = 459
                                        }
                                        _ = v91
                                    } else {
                                        switch (fi77 - 9223372036854775807) {
                                        case -1:
                                            v103 := GenS4{F0: (v2 ^ 551)}
                                            _ = v103
                                        case 773:
                                            fmt.Println(fi77, v3)
                                        default:
                                            v3 = v3
                                        }
                                        v104 := GenS12{F0: -681}
                                        _ = v104
                                    }
                                    var v105 string = "foo bar"
                                    _ = v105
                                }
                                _ = fi90
                            }
                        default:
                            v106 := GenS15{F0: -2147483648, F1: 1, F2: 0}
                            _ = v106
                        }
                    }
                }
                _ = fi77
            }
        } else {
            for fi107 := 0; fi107 < 2; fi107++ {
                var v108 int = fi107
                _ = fi107
                _ = v108
            }
        }
    }
    var v109 int32 = 0
    v1 = append(v1, 32768)
    for fi110 := 0; fi110 < 4; fi110++ {
        var v111 bool = false
        _ = fi110
        _ = v111
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v38
}
