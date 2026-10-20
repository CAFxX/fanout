package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 int64; F2 int64 }
type GenS11 struct { F0 uint16; F1 int64 }
func genFn0(p0 string, p1 bool, p2 bool) uint16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int64 = -9223372036854775808
            fmt.Println(v6, p0, fi4)
            var v7 bool = false
            _ = fi5
            _ = v6
            _ = v7
        }
        p1 = p1
        p2 = p2
        _ = fi4
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{(fi8 * 128), fi8, fi8}
        p1 = p2
        var v10 int64 = 397
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 2
    for fi12 := 0; fi12 < 5; fi12++ {
        v11 += (v11 & (^(v11)))
        for fi13 := 0; fi13 < 10; fi13++ {
            if (645 > (fi12 ^ fi12)) {
                switch (fi12 ^ (fi13 + int(v11))) {
                case -2147483648:
                    p1 = p1
                case 127:
                    if (0 >= 326) {
                        var v14 int64 = -1
                        var v15 string = p0
                        v16 := []int{fi13, fi13, fi13, 32768}
                        var v17 bool = p2
                        v18 := GenS1{F0: 255, F1: 0, F2: -1}
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint32 = 2147483648
                        v20 := GenS2{F0: 342}
                        v11 += 1
                        _ = v19
                        _ = v20
                    }
                default:
                    p2 = p1
                }
                var v21 uint32 = 42
                if false {
                    v22 := []int{}
                    _ = v22
                } else {
                    var v23 float32 = (-2.25 - -2.25)
                    _ = v23
                }
                _ = v21
            }
            _ = fi13
        }
        v24 := GenS3{F0: 633, F1: -1, F2: 1}
        _ = fi12
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 904
}

func main() {
    var v1 int = 32767
    var v2 float32 = (((0.0 * ((1e10 * float32((1e10 - ((3.14159 * 0.1) - (float64((1e10 - (-2.25 - (0.0 * 0.0)))) - 1e10))))) - -2.25)) / float32(float64((1e10 - -2.25)))) + 1e10)
    var v3 bool = (true && false)
    for fi25 := 0; fi25 < 5; fi25++ {
        switch 127 {
        case 32768:
            v26 := []int{fi25, 128, 32768}
            _ = v26
        case 2147483647:
            v27 := []int{(^(fi25)), 128, 32768}
            _ = v27
        }
        _ = fi25
    }
    v3 = ("" >= ("hello" + "foo bar"))
    for fi28 := 0; fi28 < 11; fi28++ {
        if (801 == 130) {
            fmt.Println(fi28)
        }
        _ = fi28
    }
    if v3 {
        for fi29 := 0; fi29 < 4; fi29++ {
            for fi30 := 0; fi30 < 13; fi30++ {
                if true {
                    v31 := []int{fi30, fi30, -857, (v1 | 0)}
                    var v32 int32 = -334
                    _ = v31
                    _ = v32
                }
                _ = fi30
            }
            _ = fi29
        }
    } else {
        v2 -= v2
        var v33 int32 = -1
        v33 = (v33 >> ((387) & 63))
        _ = v33
    }
    fmt.Println(v3, v2)
    switch v1 {
    case 737:
        v34 := GenS2{F0: 1}
        _ = v34
    case -915:
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 uint32 = 0
            if (-24 <= 100) {
                fmt.Println(v3, v1)
                var v37 bool = (-1 >= -1)
                v1 = (-(v1))
                _ = v37
            } else {
                if v3 {
                    v1 = (v1 | 0)
                    for fi38 := 0; fi38 < 4; fi38++ {
                        switch (fi38 + 4294967295) {
                        case 2147483647:
                            if ((255 >= 31) || false) {
                                var v39 uint32 = v36
                                switch v1 {
                                case -2147483648:
                                    v40 := GenS1{F0: 42, F1: -173, F2: 42}
                                    _ = v40
                                case -369:
                                    v3 = (!false)
                                default:
                                    v41 := []int{}
                                    _ = v41
                                }
                                _ = v39
                            }
                        case 474:
                            var v42 int16 = 158
                            _ = v42
                        case -681:
                            if (1 == 136) {
                                v43 := []int{v1, fi35, (v1 << ((2) & 63))}
                                switch "hello" {
                                case "":
                                    if true {
                                        v44 := GenS4{F0: 1, F1: 0}
                                        _ = v44
                                    } else {
                                        var v45 uint32 = (^(v36))
                                        var v46 string = "a"
                                        if (uint64(v36) == 42) {
                                            v46 = ((("\n\t" + v46) + ("" + ("hello" + v46))) + v46)
                                            fmt.Println(fi35, fi38, v3)
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    if ((!false) && (v36 <= (v36 * 31))) {
                                        v1 = 0
                                        if v3 {
                                            v47 := GenS3{F0: fi38, F1: -128, F2: 0}
                                            _ = v47
                                        } else {
                                            var v48 int64 = -891
                                            switch (v1 >> ((65535) & 63)) {
                                            case 127:
                                                v48 = (v48 << ((2) & 63))
                                            case -434:
                                                v2 = (v2 - (1e10 + 0.0))
                                            case 4294967295:
                                                v49 := GenS3{F0: v1, F1: -5, F2: 127}
                                                _ = v49
                                            default:
                                                var v50 uint8 = 107
                                                _ = v50
                                            }
                                            if (v3 || false) {
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                                fmt.Println(v36, v2, v3)
                                                var v51 float32 = ((3.14159 + 0.0) + (1.5 / ((((0.1 + v2) * 0.1) / 3.14159) / (v2 - ((0.0 * (v2 * ((0.0 + -2.25) * float32(1e10)))) + (((v2 + float32(float64((0.1 / v2)))) * (1e10 / (((v2 - v2) + float32(1.5)) * (1.5 + v2)))) * v2))))))
                                                var v52 int16 = -32768
                                                var v53 string = ("hello" + "a")
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            } else {
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                            }
                                            _ = v48
                                        }
                                        v54 := GenS5{F0: (63 > 130), F1: ("foo bar" + ("" + "a"))}
                                        _ = v54
                                    }
                                case "hello":
                                    v2 = 0.1
                                }
                                _ = v43
                            }
                        }
                        v2 = (v2 * 3.14159)
                        v2 = (v2 - v2)
                        if v3 {
                            if v3 {
                                fmt.Println(v36, fi38, v2)
                                if ((!v3) || (939 == 42)) {
                                    switch v1 {
                                    case 32767:
                                        v3 = (!(!v3))
                                    default:
                                        switch (v1 >> ((122) & 63)) {
                                        case 42:
                                            switch "a" {
                                            case "foo bar":
                                                if v3 {
                                                    var v55 int8 = -128
                                                    v56 := []int{(fi35 >> ((65535) & 63)), fi38}
                                                    v56 = append(v56, -9223372036854775808)
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    if true {
                                                        v2 -= v2
                                                        v3 = true
                                                        var v57 int16 = 32767
                                                        _ = v57
                                                    }
                                                }
                                            }
                                        case 945:
                                            v58 := GenS6{F0: 42, F1: fi35}
                                            _ = v58
                                        case 32768:
                                            var v59 float32 = v2
                                            _ = v59
                                        default:
                                            var v60 uint8 = 42
                                            _ = v60
                                        }
                                    }
                                }
                            }
                            var v61 int32 = -330
                            _ = v61
                        } else {
                            var v62 int = (-(fi35))
                            var v63 int32 = 0
                            v64 := []int{53, v62}
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                        v1 = fi38
                        _ = fi38
                    }
                    fmt.Println(v2, v3)
                } else {
                    var v65 bool = (((v2 * v2) != v2) && v3)
                    fmt.Println(v1)
                    v66 := GenS5{F0: v65, F1: (("" + "") + "hello")}
                    _ = v65
                    _ = v66
                }
                v67 := GenS3{F0: -1, F1: 1, F2: 122}
                v68 := GenS7{F0: 0, F1: -1}
                _ = v67
                _ = v68
            }
            _ = fi35
            _ = v36
        }
    case -356:
        for fi69 := 0; fi69 < 2; fi69++ {
            switch v1 {
            case 127:
                if (905 > 945) {
                    fmt.Println(v1, v3, v2)
                    for fi70 := 0; fi70 < 10; fi70++ {
                        if true {
                            var v71 float32 = (v2 - float32(1e10))
                            v72 := GenS4{F0: 18446744073709551615, F1: 1}
                            _ = v71
                            _ = v72
                        } else {
                            v3 = (v3 && (!(4294967295 != v1)))
                            var v73 float64 = 0.0
                            var v74 int8 = 42
                            v73 = 0.1
                            var v75 uint16 = 2
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        v76 := []int{}
                        _ = fi70
                        _ = v76
                    }
                }
            case 32768:
                for fi77 := 0; fi77 < 8; fi77++ {
                    v2 -= ((v2 / v2) / (((v2 + (1e10 + (0.1 / v2))) / v2) + ((1e10 / ((v2 / (float32((3.14159 - ((0.0 * (0.0 / (float64(v2) * 0.1))) / 0.1))) / float32(1.5))) + ((v2 * v2) + ((v2 / (1e10 - v2)) + v2)))) + v2)))
                    _ = fi77
                }
            default:
                if ((v3 || false) && (!(1 > 0))) {
                    v78 := []int{(v1 - v1), (v1 ^ 4294967295), fi69, 4294967295}
                    v79 := GenS8{F0: 0}
                    _ = v78
                    _ = v79
                }
            }
            _ = fi69
        }
    }
    var v80 uint32 = 56
    var v81 int32 = -1
    for fi82 := 0; fi82 < 17; fi82++ {
        for fi83 := 0; fi83 < 1; fi83++ {
            var v84 float32 = 0.1
            _ = fi83
            _ = v84
        }
        var v85 uint16 = 42
        v86 := []int{fi82}
        _ = fi82
        _ = v85
        _ = v86
    }
    if v3 {
        v87 := []int{v1, -28, (v1 + 128)}
        fmt.Println(v1, v81)
        if v3 {
            switch ("foo bar" + ("hello" + "\n\t")) {
            case "":
                var v88 uint32 = v80
                _ = v88
            case "\n\t":
                v89 := GenS2{F0: 881}
                _ = v89
            default:
                var v90 int32 = (v81 - v81)
                _ = v90
            }
            var v91 bool = v3
            v92 := GenS1{F0: 1, F1: v81, F2: 0}
            _ = v91
            _ = v92
        } else {
            if (!v3) {
                if len(v87) > 0 {
                    v87[0] = (v1 ^ 42)
                }
                v87 = append(v87, (^(v1)))
                var v93 uint16 = 255
                _ = v93
            } else {
                if v3 {
                    var v94 uint16 = 42
                    if ((("a" + "hello") >= "foo bar") && ((-(v1)) == v1)) {
                        _ = len(v87)
                        v95 := GenS9{F0: "a"}
                        v96 := GenS6{F0: 4294967295, F1: 128}
                        _ = v95
                        _ = v96
                    } else {
                        var v97 uint8 = 255
                        v87 = append(v87, 32768)
                        v98 := GenS10{F0: 1, F1: 9223372036854775807, F2: -9223372036854775808}
                        _ = v97
                        _ = v98
                    }
                    if len(v87) > 0 {
                        v87[0] = (v1 | v1)
                    }
                    _ = v94
                }
            }
            v99 := GenS5{F0: v3, F1: ("" + "")}
            if ((("foo bar" + ((("" + ("a" + "hello")) + (("foo bar" + "a") + "a")) + "foo bar")) + ("foo bar" + "hello")) < "hello") {
                v2 = v2
                v2 -= 1.5
                fmt.Println(v81, v1, v2)
            } else {
                var v100 float64 = (float64(float32(((float64(v2) * float64((float32(1e10) * 1.5))) - 3.14159))) * (1e10 - (1e10 / 1.0)))
                for fi101 := 0; fi101 < 18; fi101++ {
                    var v102 int32 = (v81 ^ (^(v81)))
                    v80 = v80
                    _ = fi101
                    _ = v102
                }
                _ = v100
            }
            _ = v99
        }
        _ = v87
    } else {
        if v3 {
            v103 := []int{122, 32767, 1}
            v80 -= (v80 & v80)
            var v104 float64 = 0.1
            _ = v103
            _ = v104
        }
        var v105 uint16 = 2
        v106 := GenS7{F0: (v105 & (v105 | v105)), F1: (v81 * v81)}
        _ = v105
        _ = v106
    }
    for fi107 := 0; fi107 < 9; fi107++ {
        for fi108 := 0; fi108 < 4; fi108++ {
            fmt.Println(v1, fi107, v1)
            v109 := GenS6{F0: 1, F1: fi107}
            _ = fi108
            _ = v109
        }
        if ("\n\t" != "a") {
            v1 = int(v80)
            if v3 {
                v1 += v1
                v110 := GenS11{F0: 470, F1: -268}
                v2 = (0.1 + 3.14159)
                _ = v110
            } else {
                var v111 int64 = -9223372036854775808
                _ = v111
            }
        } else {
            var v112 int16 = 1000
            _ = v112
        }
        _ = fi107
    }
    var v113 int = 462
    var v114 int32 = 1
    var v115 float64 = (0.0 / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v2
    _ = v3
    _ = v80
    _ = v81
}
