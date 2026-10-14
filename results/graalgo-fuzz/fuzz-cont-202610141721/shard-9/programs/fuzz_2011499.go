package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint16; F1 bool }
type GenS13 struct { F0 int16; F1 bool }
type GenS14 struct { F0 int8; F1 int; F2 uint8 }

func main() {
    var v1 string = ""
    v2 := make(map[int]int)
    var v3 int8 = -65
    var v4 uint16 = 281
    v5 := []string{("a" + "\n\t")}
    var v6 string = v1
    v7 := GenS1{F0: 18446744073709551615}
    for fi8 := 0; fi8 < 5; fi8++ {
        v3 -= v3
        v6 = ((v6 + "") + "hello")
        v9 := GenS1{F0: 604}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 67}
    v11 := v2[4294967295]
    if len(v5) > 0 {
        v5[0] = v1
    }
    v12 := GenS2{F0: 42}
    switch 829 {
    case 32768:
        if false {
            if true {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS2{F0: 0}
                    v3 -= 127
                    v4 -= v4
                    _ = fi13
                    _ = v14
                }
                switch 42 {
                case 42:
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v16 := GenS1{F0: 1}
                        _ = fi15
                        _ = v16
                    }
                case 824:
                    var v17 uint32 = 2147483648
                    _ = v17
                default:
                    if (56 <= 1) {
                        var v18 uint = 255
                        v11 = 898
                        _ = v18
                    }
                }
            }
            var v19 int32 = -2147483648
            _ = v19
        }
    case 738:
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint64 = 4
            if false {
                v22 := v2[fi20]
                fmt.Println(v21, v1, v3)
                _ = v22
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v5 = append(v5, "hello")
                var v23 uint8 = 255
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    default:
        fmt.Println(v11, v1)
    }
    if (1 > v11) {
        var v24 uint64 = 42
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 uint16 = v4
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        var v27 int64 = 733
        if true {
            var v28 int64 = v27
            var v29 int32 = -312
            _ = v28
            _ = v29
        }
        var v30 uint32 = 1
        _ = v27
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        if true {
            var v32 bool = (!true)
            v6 = v6
            switch (v1 + v6) {
            case "foo bar":
                var v33 uint64 = 4294967295
                _ = v33
            }
            _ = v32
        } else {
            switch (v6 + "foo bar") {
            case "":
                switch int(v4) {
                case -706:
                    v4 = 32767
                case -530:
                    var v34 float32 = (0.0 - 1.5)
                    _ = v34
                }
            case "hello":
                v35 := GenS3{F0: -1, F1: true}
                _ = v35
            }
            var v36 int64 = 337
            var v37 string = ("foo bar" + "foo bar")
            _ = v36
            _ = v37
        }
        var v38 float64 = (1.5 / 1e10)
        v39 := GenS4{F0: 2, F1: ((0.1 < -2.25) && (1 > (v11 >> ((65535) & 63))))}
        _ = fi31
        _ = v38
        _ = v39
    }
    var v40 int16 = 42
    if (v11 <= int(v40)) {
        var v41 int64 = 42
        _ = v41
    } else {
        for fi42 := 0; fi42 < 10; fi42++ {
            v3 = -1
            if true {
                v43 := GenS3{F0: int32(v4), F1: (v40 == (v40 * (v40 % ((v40) | 1))))}
                v1 = (v1 + ((("" + ((((v6 + v6) + "a") + v1) + v6)) + "\n\t") + v6))
                v44 := GenS3{F0: 2147483647, F1: (v1 < v6)}
                _ = v43
                _ = v44
            } else {
                switch (v11 + 9223372036854775807) {
                case 35:
                    var v45 int64 = -9223372036854775808
                    _ = v45
                case -993:
                    v46 := GenS1{F0: 2}
                    _ = v46
                case 32768:
                    for fi47 := 0; fi47 < 14; fi47++ {
                        switch v6 {
                        case "":
                            v48 := GenS5{F0: (fi47 ^ fi47)}
                            _ = v48
                        }
                        switch "" {
                        case "hello":
                            if true {
                                var v49 int16 = (v40 - -915)
                                if true {
                                    fmt.Println(v49)
                                    fmt.Println(fi47, v1, v1)
                                } else {
                                    if true {
                                        v1 = (v6 + v6)
                                    } else {
                                        switch v6 {
                                        case "":
                                            v3 -= -118
                                        case "a":
                                            switch fi42 {
                                            case 147:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v6 = v6
                                                v50 := GenS6{F0: 2, F1: (true || false)}
                                                _ = v50
                                            case -1:
                                                var v51 uint16 = 65535
                                                fmt.Println(v3)
                                                switch (fi47 << ((1) & 63)) {
                                                case 638:
                                                    v52 := GenS7{F0: v40, F1: (v1 + v6)}
                                                    _ = v52
                                                }
                                                _ = v51
                                            case 1:
                                                var v53 uint8 = 127
                                                _ = v53
                                            default:
                                                v54 := GenS3{F0: -575, F1: (65535 < 0)}
                                                _ = v54
                                            }
                                        default:
                                            fmt.Println(v3)
                                        }
                                        var v55 int8 = (v3 * (-(v3)))
                                        if false {
                                            if false {
                                                switch v6 {
                                                case "a":
                                                    var v56 float32 = (0.0 + (((1e10 - -2.25) + 1.5) + -2.25))
                                                    _ = v56
                                                }
                                                var v57 int16 = v40
                                                _ = v57
                                            } else {
                                                var v58 string = ((v6 + (v1 + "hello")) + v6)
                                                if len(v5) > 0 {
                                                    v5[0] = ((("hello" + (v6 + v6)) + (v58 + "\n\t")) + v1)
                                                }
                                                v6 = "foo bar"
                                                v59 := GenS5{F0: v11}
                                                _ = v58
                                                _ = v59
                                            }
                                        } else {
                                            switch (v11 % (((v11 / (((fi42 + (^(fi47)))) | 1))) | 1)) {
                                            case 128:
                                                v11 = (^(fi47))
                                            }
                                            if ((fi42 << ((250) & 63)) < fi47) {
                                                var v60 uint16 = (v4 | v4)
                                                _ = v60
                                            }
                                        }
                                        var v61 int16 = 1000
                                        _ = v55
                                        _ = v61
                                    }
                                    var v62 uint8 = 127
                                    if true {
                                        v63 := GenS1{F0: 255}
                                        switch fi47 {
                                        case 128:
                                            switch v1 {
                                            case "\n\t":
                                                if false {
                                                    switch (v1 + (v1 + v6)) {
                                                    case "":
                                                        v64 := GenS7{F0: v40, F1: (v1 + "\n\t")}
                                                        _ = v64
                                                    default:
                                                        var v65 uint64 = 562
                                                        _ = v65
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    if (-1 != -841) {
                                                        v66 := GenS8{F0: fi47, F1: v11}
                                                        v6 = v6
                                                        var v67 int = v11
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        v68 := GenS3{F0: -1, F1: (1 <= -311)}
                                                        _ = v68
                                                    }
                                                }
                                            case "hello":
                                                v69 := GenS9{F0: (v4 + 1)}
                                                _ = v69
                                            }
                                        case 127:
                                            v1 = "a"
                                            if (!((0.1 - 1.5) >= float32(3.14159))) {
                                                fmt.Println(v49, v4, fi47)
                                                v70 := GenS9{F0: 255}
                                                _ = v70
                                            }
                                            v71 := GenS10{F0: (^(v49)), F1: (v49 / ((v49) | 1))}
                                            _ = v71
                                        case 32768:
                                            v72 := GenS11{F0: 940, F1: (1 <= 0)}
                                            _ = v72
                                        default:
                                            v49 = int16(v3)
                                        }
                                        _ = v63
                                    }
                                    _ = v62
                                }
                                _ = v49
                            }
                        case "a":
                            v73 := GenS12{F0: v4, F1: (!((1 <= 207) || false))}
                            _ = v73
                        default:
                            if true {
                                var v74 uint = 42
                                _ = v74
                            } else {
                                v4 = (v4 * (v4 ^ 42))
                                fmt.Println(v3, fi42, v6)
                                v75 := GenS13{F0: v40, F1: true}
                                _ = v75
                            }
                        }
                        var v76 int64 = -931
                        v1 = v6
                        if true {
                            v3 -= (v3 * (v3 << ((42) & 63)))
                            switch v6 {
                            case "foo bar":
                                v4 -= 993
                            case "\n\t":
                                if true {
                                    var v77 uint64 = 1
                                    _ = v77
                                }
                                switch (^(fi47)) {
                                case 128:
                                    v78 := GenS4{F0: (v4 + (v4 / ((2) | 1))), F1: (false && (101 < 2))}
                                    _ = v78
                                case 0:
                                    v4 -= v4
                                case 32768:
                                    var v79 float32 = ((3.14159 + (0.0 / float32((0.1 - 1.5)))) - -2.25)
                                    _ = v79
                                }
                            case "hello":
                                var v80 string = "a"
                                _ = v80
                            default:
                                var v81 int16 = -779
                                _ = v81
                            }
                        } else {
                            v82 := GenS5{F0: 0}
                            v40 -= -32768
                            _ = v82
                        }
                        if (!(0 != (v76 + v76))) {
                            if false {
                                fmt.Println(v3, v40, fi47)
                                v6 = v6
                                v1 = v1
                            } else {
                                if false {
                                    var v83 uint = 1
                                    var v84 uint64 = 42
                                    _ = v83
                                    _ = v84
                                }
                                fmt.Println(v3, v3, v4)
                                if true {
                                    v11 += fi47
                                    v76 -= -1
                                    delete(v2, (-(v11)))
                                }
                            }
                            v1 = (v1 + "hello")
                            var v85 int = (v11 % (((fi47 << ((255) & 63))) | 1))
                            _ = v85
                        }
                        _ = fi47
                        _ = v76
                    }
                }
            }
            _ = fi42
        }
        if true {
            v40 += 920
            switch v1 {
            case "foo bar":
                for fi86 := 0; fi86 < 3; fi86++ {
                    v87 := v2[v11]
                    v3 = v3
                    _ = fi86
                    _ = v87
                }
            }
            switch (v11 ^ 32768) {
            case -1:
                var v88 float64 = ((0.0 - 1e10) - (0.0 / -2.25))
                _ = v88
            case 127:
                v89 := v2[42]
                _ = v89
            case 2147483647:
                switch -797 {
                case 128:
                    var v90 int64 = 579
                    _ = v90
                case 32767:
                    v91 := GenS4{F0: v4, F1: ((v40 % ((32767) | 1)) >= 168)}
                    _ = v91
                case 9223372036854775807:
                    v11 = (v11 << ((65535) & 63))
                default:
                    if false {
                        fmt.Println(v11, v11)
                        var v92 uint = 1
                        _ = v92
                    }
                }
            default:
                for fi93 := 0; fi93 < 13; fi93++ {
                    v40 = (v40 / ((v40) | 1))
                    v94 := GenS8{F0: fi93, F1: fi93}
                    fmt.Println(v40, v40, v40)
                    _ = fi93
                    _ = v94
                }
            }
        }
    }
    var v95 int16 = (v40 >> ((378) & 63))
    for fi96 := 0; fi96 < 9; fi96++ {
        var v97 int8 = 42
        fmt.Println(fi96, v95)
        _ = fi96
        _ = v97
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v98 bool = true
    v95 = (v95 + 42)
    for fi99 := 0; fi99 < 17; fi99++ {
        if ((-2.25 * float32((3.14159 - ((1e10 / -2.25) / (0.0 + (((0.1 + 0.0) - ((((((float64(0.0) + float64(1e10)) + (-2.25 - ((float64((((1.5 * (0.0 + 0.0)) + 3.14159) * (0.0 * (-2.25 / 1e10)))) * 0.0) * ((0.0 - (-2.25 / 1.0)) - 3.14159)))) - 0.1) - -2.25) + 1e10) / 0.1)) * 0.1)))))) <= (float32((((3.14159 + 1e10) + 0.0) + (1.5 - 3.14159))) / (1e10 - 0.1))) {
            if v98 {
                v100 := GenS6{F0: uint64(v4), F1: (!(!(65535 <= 901)))}
                var v101 uint64 = 4294967295
                switch (v11 >> ((1) & 63)) {
                case -80:
                    v95 = v95
                default:
                    var v102 uint8 = 2
                    _ = v102
                }
                _ = v100
                _ = v101
            } else {
                v4 = (^(v4))
                _ = len(v5)
            }
        } else {
            v3 = 70
            for fi103 := 0; fi103 < 20; fi103++ {
                var v104 string = (v1 + "\n\t")
                var v105 bool = v98
                v40 -= (v95 ^ (v40 / ((-32768) | 1)))
                if (0 >= 100) {
                    var v106 uint = 299
                    v107 := GenS14{F0: v3, F1: 42, F2: 1}
                    _ = v106
                    _ = v107
                } else {
                    var v108 int64 = -9223372036854775808
                    v40 = (^(v95))
                    _ = v108
                }
                var v109 int64 = 1
                _ = fi103
                _ = v104
                _ = v105
                _ = v109
            }
        }
        _ = fi99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v95
    _ = v98
}
