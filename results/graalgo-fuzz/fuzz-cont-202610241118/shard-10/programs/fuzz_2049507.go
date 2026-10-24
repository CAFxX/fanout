package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS7 struct { F0 int16 }
func genFn0() bool {
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 4; fi6++ {
            if (!false) {
                var v7 uint16 = 139
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint64 = 758
                    v10 := []int{-2147483648, 2147483647, fi8, fi6}
                    var v11 int32 = -2147483648
                    if len(v10) > 0 {
                        v10[0] = (fi6 << ((1) & 63))
                    }
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 int8 = 100
                _ = v12
                _ = v7
            }
            if (0.0 != 3.14159) {
                var v13 uint = 301
                var v14 int32 = 805
                _ = v13
                _ = v14
            }
            switch "\n\t" {
            case "\n\t":
                v15 := []int{fi6, fi6, fi6}
                _ = v15
            case "foo bar":
                var v16 int16 = 32767
                _ = v16
            case "a":
                var v17 uint = 0
                _ = v17
            default:
                for fi18 := 0; fi18 < 4; fi18++ {
                    fmt.Println(fi18)
                    _ = fi18
                }
            }
            _ = fi6
        }
    case "\n\t":
        for fi19 := 0; fi19 < 16; fi19++ {
            if (!(!false)) {
                if false {
                    if (!(61 == 4294967295)) {
                        if (1 == 18446744073709551615) {
                            for fi20 := 0; fi20 < 12; fi20++ {
                                var v21 uint32 = 1
                                _ = fi20
                                _ = v21
                            }
                            switch fi19 {
                            case 127:
                                v22 := GenS1{F0: 1}
                                _ = v22
                            case 0:
                                var v23 float64 = ((0.0 * 3.14159) - (float64(0.1) - (1e10 * 3.14159)))
                                _ = v23
                            case 100:
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    var v25 bool = false
                                    v25 = (122 == -1)
                                    _ = fi24
                                    _ = v25
                                }
                            default:
                                switch fi19 {
                                case 32767:
                                    fmt.Println(fi19)
                                case -2147483648:
                                    var v26 uint32 = 65535
                                    _ = v26
                                }
                            }
                        } else {
                            v27 := []int{fi19}
                            var v28 int = fi19
                            fmt.Println(v28, v28)
                            _ = v27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 16; fi29++ {
                            var v30 string = "\n\t"
                            var v31 uint64 = 920
                            var v32 bool = false
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    var v33 uint32 = 0
                    switch ("a" + "foo bar") {
                    case "":
                        v33 = (v33 / (((^(v33))) | 1))
                    case "\n\t":
                        v34 := GenS2{F0: (v33 + v33), F1: int64(fi19)}
                        _ = v34
                    default:
                        for fi35 := 0; fi35 < 16; fi35++ {
                            if true {
                                v36 := GenS3{F0: 0, F1: 938}
                                _ = v36
                            }
                            switch "a" {
                            case "":
                                var v37 uint32 = v33
                                _ = v37
                            default:
                                var v38 uint8 = 42
                                _ = v38
                            }
                            v33 += (v33 | (v33 | 391))
                            _ = fi35
                        }
                    }
                    _ = v33
                }
                fmt.Println(fi19)
            } else {
                v39 := []int{fi19, (fi19 >> ((523) & 63)), fi19}
                v40 := GenS2{F0: 145, F1: 161}
                _ = v39
                _ = v40
            }
            _ = fi19
        }
    case "foo bar":
        if true {
            var v41 int32 = 0
            v41 = v41
            _ = v41
        }
    default:
        var v42 int8 = -1
        _ = v42
    }
    v43 := GenS3{F0: 2, F1: -433}
    v44 := []int{9223372036854775807, 1, 32768}
    _ = v43
    _ = v44
    return (!(((0 >= 2147483647) && (false && false)) || (9223372036854775807 < 42)))
}
func genFn1() uint8 {
    var v45 uint8 = 128
    for fi46 := 0; fi46 < 6; fi46++ {
        for fi47 := 0; fi47 < 11; fi47++ {
            switch ((("\n\t" + "a") + ("foo bar" + ("foo bar" + "a"))) + "a") {
            case "foo bar":
                v45 += (v45 >> ((992) & 63))
            case "a":
                v48 := []int{42, fi47}
                _ = v48
            case "":
                var v49 int8 = 0
                _ = v49
            default:
                v45 += v45
            }
            var v50 int = fi46
            _ = fi47
            _ = v50
        }
        _ = fi46
    }
    if true {
        v45 += (^(v45))
        var v51 uint8 = (v45 + (v45 * v45))
        _ = v51
    } else {
        v52 := []int{127, 4294967295, 127}
        _ = len(v52)
        _ = v52
    }
    var v53 uint8 = (v45 / (((^(v45))) | 1))
    if (true && ((v53 >> ((513) & 63)) >= (v45 * (^(v53))))) {
        var v54 int8 = 127
        fmt.Println(v54, v54, v45)
        v53 = (v53 - v45)
        _ = v54
    }
    _ = v45
    _ = v53
    return v53
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    v3 := make(map[string]bool)
    v4 := []int{-316, -192, 32767}
    v5 := make(map[string]bool)
    switch 127 {
    case 501:
        if false {
            fmt.Println(v2)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case -9223372036854775808:
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 uint = 521
            _ = fi55
            _ = v56
        }
    }
    v2 = 0
    var v57 uint8 = 131
    var v58 uint = 2
    switch 319 {
    case 32767:
        var v59 uint64 = 42
        _ = v59
    case 1:
        switch -918 {
        case 4294967295:
            for fi60 := 0; fi60 < 17; fi60++ {
                v58 = v58
                var v61 string = "foo bar"
                switch fi60 {
                case 4294967295:
                    for fi62 := 0; fi62 < 13; fi62++ {
                        var v63 uint16 = v2
                        v64 := GenS1{F0: -9223372036854775808}
                        v58 += 600
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                case 32768:
                    for fi65 := 0; fi65 < 5; fi65++ {
                        var v66 int = (^(fi65))
                        v67 := GenS1{F0: 288}
                        fmt.Println(fi60, v58)
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                case 0:
                    for fi68 := 0; fi68 < 8; fi68++ {
                        switch 2147483647 {
                        case 0:
                            v57 -= v57
                        case -893:
                            v58 -= (^(v58))
                        case 614:
                            if (32767 > -32768) {
                                switch ((v61 + (v61 + v61)) + v61) {
                                case "foo bar":
                                    v2 += v2
                                case "":
                                    v69 := GenS4{F0: fi68, F1: 2147483647}
                                    _ = v69
                                }
                                v70 := v3[(v61 + v61)]
                                _ = v70
                            } else {
                                v61 = (v61 + "a")
                                var v71 bool = true
                                _ = v71
                            }
                        }
                        _ = fi68
                    }
                }
                _ = fi60
                _ = v61
            }
        case 9223372036854775807:
            v72 := GenS2{F0: 1, F1: 9223372036854775807}
            _ = v72
        case 0:
            for fi73 := 0; fi73 < 8; fi73++ {
                switch "hello" {
                case "\n\t":
                    if false {
                        v74 := GenS5{F0: v58, F1: 208}
                        var v75 uint64 = 641
                        _ = v74
                        _ = v75
                    } else {
                        if (954 == -626) {
                            if true {
                                v76 := GenS2{F0: 502, F1: 1}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v76
                            } else {
                                for fi77 := 0; fi77 < 10; fi77++ {
                                    _ = len(v4)
                                    _ = fi77
                                }
                                v2 -= v2
                                switch fi73 {
                                case 32768:
                                    v4 = append(v4, fi73)
                                case 32767:
                                    fmt.Println(v57, v58)
                                case 0:
                                    var v78 string = "foo bar"
                                    _ = v78
                                }
                            }
                            for fi79 := 0; fi79 < 7; fi79++ {
                                var v80 uint8 = genFn1()
                                var v81 int64 = 9223372036854775807
                                var v82 uint16 = 150
                                _ = fi79
                                _ = v80
                                _ = v81
                                _ = v82
                            }
                        } else {
                            for fi83 := 0; fi83 < 15; fi83++ {
                                if true {
                                    if (true && (42 >= v2)) {
                                        v57 = genFn1()
                                        if (!false) {
                                            v58 += (^(v58))
                                            if (1.5 > 3.14159) {
                                                v4 = append(v4, 1)
                                                switch (-(fi83)) {
                                                case 4294967295:
                                                    switch (-(fi73)) {
                                                    case 32768:
                                                        v2 = v2
                                                        v84 := GenS6{F0: 0, F1: 1, F2: 1}
                                                        var v85 uint = (v58 * (v58 & v58))
                                                        _ = v84
                                                        _ = v85
                                                    case 0:
                                                        v2 = v2
                                                    }
                                                case -820:
                                                    v58 = v58
                                                default:
                                                    v86 := GenS7{F0: 0}
                                                    _ = v86
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = fi73
                                                }
                                            } else {
                                                switch ("a" + ("\n\t" + "foo bar")) {
                                                case "foo bar":
                                                    fmt.Println(v58)
                                                default:
                                                    var v87 uint = (v58 * 255)
                                                    _ = v87
                                                }
                                                var v88 string = "\n\t"
                                                v57 += v57
                                                switch v88 {
                                                case "":
                                                    v57 -= 255
                                                }
                                                v89 := GenS2{F0: 4294967295, F1: -9223372036854775808}
                                                _ = v88
                                                _ = v89
                                            }
                                            switch ("hello" + "\n\t") {
                                            case "\n\t":
                                                var v90 int64 = -9223372036854775808
                                                _ = v90
                                            case "foo bar":
                                                v1[((("foo bar" + "\n\t") + "a") + "\n\t")] = (-(fi73))
                                            case "a":
                                                if len(v4) > 0 {
                                                    v4[0] = fi73
                                                }
                                            }
                                        } else {
                                            v2 = (^(v2))
                                            fmt.Println(v57, fi73, fi73)
                                        }
                                        fmt.Println(v58, v2, v2)
                                    }
                                    var v91 uint = v58
                                    _ = v91
                                }
                                v4 = append(v4, (fi73 << ((42) & 63)))
                                _ = fi83
                            }
                            for fi92 := 0; fi92 < 5; fi92++ {
                                v2 = (^(v2))
                                fmt.Println(v57, fi92, v57)
                                _ = fi92
                            }
                        }
                        fmt.Println(v58, v2, fi73)
                        var v93 string = ""
                        _ = v93
                    }
                case "":
                    v94 := GenS5{F0: 42, F1: v57}
                    _ = v94
                case "a":
                    var v95 uint64 = 1
                    _ = v95
                }
                v2 = v2
                v58 = 4294967295
                _ = fi73
            }
        default:
            v2 -= (v2 & v2)
        }
    }
    var v96 uint64 = uint64(v57)
    fmt.Println(v57, v57)
    var v97 float32 = (1.5 / (1e10 / 1e10))
    for fi98 := 0; fi98 < 20; fi98++ {
        var v99 int32 = 42
        var v100 int8 = 0
        _ = fi98
        _ = v100
        _ = v99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v96
    _ = v97
}
