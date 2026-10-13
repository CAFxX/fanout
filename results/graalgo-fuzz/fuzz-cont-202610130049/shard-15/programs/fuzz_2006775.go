package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 bool = (2 == 299)
    var v2 uint16 = 65535
    var v3 int16 = int16(v2)
    var v4 int8 = -128
    switch 127 {
    case -225:
        var v5 int = 4294967295
        _ = v5
    case -6:
        if (((!(355 != 2)) || (763 == v2)) || (v1 || v1)) {
            for fi6 := 0; fi6 < 8; fi6++ {
                if v1 {
                    var v7 uint32 = 4294967295
                    for fi8 := 0; fi8 < 15; fi8++ {
                        v9 := []int{}
                        v4 += v4
                        var v10 uint = uint(v7)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    if false {
                        if ((v1 || v1) && v1) {
                            if (127 != 42) {
                                switch 0 {
                                case 0:
                                    fmt.Println(v3, v4, v2)
                                }
                                for fi11 := 0; fi11 < 17; fi11++ {
                                    var v12 int = (fi6 ^ 0)
                                    switch ("\n\t" + ("foo bar" + ("foo bar" + "a"))) {
                                    case "":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v13 string = (("\n\t" + "foo bar") + "hello")
                                            _ = v13
                                        case "a":
                                            v14 := GenS1{F0: v7, F1: 255, F2: "a"}
                                            _ = v14
                                        default:
                                            var v15 int64 = -9223372036854775808
                                            _ = v15
                                        }
                                    case "hello":
                                        var v16 uint32 = (^(v7))
                                        v17 := []int{(fi11 / ((int(v2)) | 1)), v12, v12, 2147483647}
                                        fmt.Println(v16, v4)
                                        fmt.Println(fi6)
                                        _ = v16
                                        _ = v17
                                    }
                                    if (-2.25 >= (-2.25 + 1.5)) {
                                        v18 := []int{}
                                        v19 := GenS2{F0: (v3 | (v3 * -140)), F1: -9223372036854775808, F2: -642}
                                        if false {
                                            v4 -= v4
                                            v4 = -128
                                            var v20 uint32 = (v7 & v7)
                                            _ = v20
                                        }
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        switch v12 {
                                        case 988:
                                            if len(v18) > 0 {
                                                v18[0] = (^(fi11))
                                            }
                                        }
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v1 = ((^(v3)) < (v3 >> ((255) & 63)))
                                    }
                                    _ = fi11
                                    _ = v12
                                }
                            } else {
                                var v21 uint32 = 2
                                v1 = v1
                                if (v1 && v1) {
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        var v23 int = -1
                                        v1 = v1
                                        var v24 float64 = 1e10
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                } else {
                                    var v25 int64 = 0
                                    _ = v25
                                }
                                _ = v21
                            }
                            v26 := []int{127, fi6}
                            var v27 bool = v1
                            _ = v26
                            _ = v27
                        } else {
                            v28 := []int{fi6, fi6, (fi6 % ((fi6) | 1))}
                            v29 := GenS3{F0: (float32(0.1) == 3.14159), F1: (v3 | v3)}
                            if (true && v1) {
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    var v31 bool = v1
                                    if v1 {
                                        var v32 int32 = 0
                                        _ = v32
                                    }
                                    v31 = (42 > 0)
                                    _ = fi30
                                    _ = v31
                                }
                                if v1 {
                                    var v33 int8 = int8(v2)
                                    _ = v33
                                } else {
                                    v4 = v4
                                }
                            } else {
                                var v34 uint16 = (v2 - (v2 * v2))
                                _ = v34
                            }
                            _ = v28
                            _ = v29
                        }
                        fmt.Println(v7, v4)
                    }
                    _ = v7
                }
                var v35 float64 = 0.0
                v36 := GenS4{F0: v1}
                _ = fi6
                _ = v35
                _ = v36
            }
            var v37 uint16 = v2
            _ = v37
        }
    default:
        var v38 uint16 = (v2 & v2)
        _ = v38
    }
    var v39 float64 = (0.1 * 1e10)
    v40 := GenS1{F0: 2147483647, F1: 42, F2: ("\n\t" + ((("a" + "a") + "foo bar") + ("\n\t" + "")))}
    v4 = 1
    switch (("" + "foo bar") + "") {
    case "":
        switch "foo bar" {
        case "foo bar":
            v4 = v4
        case "hello":
            v3 -= v3
        default:
            switch -897 {
            case 32767:
                v41 := GenS1{F0: 2, F1: 178, F2: "foo bar"}
                _ = v41
            case 0:
                for fi42 := 0; fi42 < 17; fi42++ {
                    switch "a" {
                    case "a":
                        v43 := []int{}
                        _ = v43
                    default:
                        v1 = v1
                    }
                    if v1 {
                        v44 := []int{}
                        if (1 != 2147483647) {
                            var v45 int32 = 2147483647
                            v44 = append(v44, int(v2))
                            v45 += v45
                            _ = v45
                        } else {
                            v46 := GenS2{F0: v3, F1: 0, F2: -9223372036854775808}
                            for fi47 := 0; fi47 < 17; fi47++ {
                                var v48 uint16 = v2
                                v3 -= v3
                                _ = fi47
                                _ = v48
                            }
                            for fi49 := 0; fi49 < 9; fi49++ {
                                switch 1 {
                                case 128:
                                    var v50 bool = (!((v4 & (v4 << ((694) & 63))) > 42))
                                    _ = v50
                                case 4294967295:
                                    var v51 int64 = -855
                                    _ = v51
                                default:
                                    if v1 {
                                        var v52 uint = 1
                                        switch -9223372036854775808 {
                                        case 32768:
                                            v52 = v52
                                            switch 444 {
                                            case -2147483648:
                                                v1 = false
                                            }
                                            var v53 uint64 = 4294967295
                                            _ = v53
                                        case -708:
                                            v3 -= 32767
                                        case 9223372036854775807:
                                            if len(v44) > 0 {
                                                v44[0] = (fi42 * (fi49 ^ (fi42 ^ fi49)))
                                            }
                                        default:
                                            switch "" {
                                            case "\n\t":
                                                v39 = ((1e10 - v39) / ((0.1 + v39) - v39))
                                            case "a":
                                                v52 = (v52 | (v52 >> (((v52 * v52)) & 63)))
                                            default:
                                                if v1 {
                                                    var v54 uint64 = 1
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v52)
                                                    case "\n\t":
                                                        v4 += v4
                                                    case "":
                                                        var v55 int32 = 1
                                                        _ = v55
                                                    default:
                                                        v54 += v54
                                                    }
                                                    v56 := GenS2{F0: 93, F1: -9223372036854775808, F2: -1}
                                                    _ = v54
                                                    _ = v56
                                                }
                                            }
                                        }
                                        _ = v52
                                    } else {
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                    }
                                }
                                if (2147483648 == 1) {
                                    if len(v44) > 0 {
                                        v44[0] = (fi42 - (fi42 << ((490) & 63)))
                                    }
                                    v57 := GenS5{F0: (v2 | v2), F1: 2, F2: 32768}
                                    _ = v57
                                } else {
                                    switch fi42 {
                                    case 32767:
                                        fmt.Println(v4, fi42, v39)
                                    default:
                                        v58 := GenS6{F0: -2147483648, F1: fi42}
                                        _ = v58
                                    }
                                }
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                                _ = fi49
                            }
                            _ = v46
                        }
                        _ = v44
                    }
                    v2 = v2
                    _ = fi42
                }
            case 128:
                switch -984 {
                case -2147483648:
                    var v59 int64 = 83
                    _ = v59
                }
            default:
                for fi60 := 0; fi60 < 15; fi60++ {
                    switch (fi60 % (((fi60 & (fi60 * (-(fi60))))) | 1)) {
                    case 127:
                        for fi61 := 0; fi61 < 7; fi61++ {
                            fmt.Println(fi60)
                            v2 = 749
                            v2 -= v2
                            _ = fi61
                        }
                    case -661:
                        switch "foo bar" {
                        case "a":
                            for fi62 := 0; fi62 < 10; fi62++ {
                                var v63 int32 = 0
                                v64 := []int{}
                                v1 = false
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        case "\n\t":
                            v39 += (v39 - -2.25)
                        }
                    case 128:
                        v3 = v3
                    }
                    for fi65 := 0; fi65 < 17; fi65++ {
                        if (!v1) {
                            v39 = ((v39 * float64(0.0)) / v39)
                            v4 += (v4 + (v4 & 0))
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v66 := GenS5{F0: 703, F1: 65535, F2: fi65}
                                _ = v66
                            case "foo bar":
                                if v1 {
                                    var v67 int32 = -1
                                    _ = v67
                                } else {
                                    var v68 int8 = v4
                                    _ = v68
                                }
                            case "hello":
                                if (!v1) {
                                    fmt.Println(v2, fi65)
                                } else {
                                    if true {
                                        fmt.Println(v39, v1, v4)
                                        v69 := []int{fi60, (fi60 * (fi65 << ((65535) & 63))), fi60}
                                        _ = v69
                                    }
                                }
                            default:
                                switch "a" {
                                case "hello":
                                    v39 += v39
                                case "":
                                    fmt.Println(v3)
                                case "\n\t":
                                    if (v3 < (v3 - v3)) {
                                        v70 := []int{}
                                        _ = v70
                                    } else {
                                        v4 -= v4
                                        v1 = v1
                                        if v1 {
                                            v3 += v3
                                            v39 = (v39 * v39)
                                        }
                                    }
                                default:
                                    v71 := []int{}
                                    _ = v71
                                }
                            }
                            var v72 uint = 1
                            v73 := GenS3{F0: v1, F1: 1}
                            v74 := []int{(fi65 & fi65), (-(fi60)), fi60}
                            v4 -= (v4 << ((v72) & 63))
                            var v75 int8 = v4
                            if len(v74) >= 2 {
                                v74 = v74[1:len(v74)]
                            }
                            _ = v72
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        _ = fi65
                    }
                    _ = fi60
                }
            }
        }
    }
    v76 := []int{32767, 128}
    for fi77 := 0; fi77 < 12; fi77++ {
        for fi78 := 0; fi78 < 3; fi78++ {
            _ = len(v76)
            var v79 uint32 = 2147483647
            v80 := GenS6{F0: 909, F1: (fi78 / (((fi77 + (fi77 | (-(fi78))))) | 1))}
            _ = fi78
            _ = v79
            _ = v80
        }
        if v1 {
            fmt.Println(v3)
            if v1 {
                var v81 int16 = (v3 | v3)
                _ = v81
            }
        }
        _ = fi77
    }
    if ((v3 - -344) < v3) {
        var v82 float32 = 1e10
        if len(v76) > 0 {
            v76[0] = 9223372036854775807
        }
        _ = v82
    } else {
        for fi83 := 0; fi83 < 13; fi83++ {
            v3 -= 1
            v84 := GenS7{F0: 42, F1: 163}
            fmt.Println(v2, v1, v2)
            _ = fi83
            _ = v84
        }
        switch 9223372036854775807 {
        case -1:
            v4 += 42
        case 32768:
            var v85 int8 = v4
            _ = v85
        case 1:
            switch 128 {
            case -9223372036854775808:
                v2 = v2
            default:
                if (!false) {
                    v2 += v2
                    v76 = append(v76, 9223372036854775807)
                } else {
                    for fi86 := 0; fi86 < 11; fi86++ {
                        v3 += (^(v3))
                        for fi87 := 0; fi87 < 4; fi87++ {
                            var v88 uint16 = 1
                            switch fi86 {
                            case 42:
                                fmt.Println(v88, v4, v1)
                            case 14:
                                fmt.Println(fi87, v39)
                            }
                            v39 -= 1e10
                            v76 = append(v76, 32767)
                            v89 := GenS8{F0: ((fi86 >> ((326) & 63)) > (-(fi86)))}
                            var v90 uint32 = 4294967295
                            _ = fi87
                            _ = v88
                            _ = v89
                            _ = v90
                        }
                        _ = fi86
                    }
                }
            }
        default:
            if len(v76) > 0 {
                v76[0] = 2147483647
            }
        }
        var v91 uint32 = uint32(v2)
        _ = v91
    }
    fmt.Println(v3)
    for fi92 := 0; fi92 < 2; fi92++ {
        if v1 {
            for fi93 := 0; fi93 < 15; fi93++ {
                v76 = append(v76, (fi93 & (fi92 + (fi92 + fi92))))
                v94 := GenS9{F0: 380}
                _ = fi93
                _ = v94
            }
            v1 = (!(65535 != 194))
            fmt.Println(v1, v2)
        }
        var v95 uint64 = 1
        _ = fi92
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v76
}
