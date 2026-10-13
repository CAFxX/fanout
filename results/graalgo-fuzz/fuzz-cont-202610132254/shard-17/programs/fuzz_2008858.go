package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 int) uint64 {
    var v6 string = "foo bar"
    var v7 uint = 778
    p0 -= (-(p0))
    v6 = v6
    var v8 float32 = 0.0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int8) uint {
    p0 = p0
    var v9 int64 = 915
    if false {
        var v10 uint16 = 368
        var v11 int = 42
        v12 := GenS1{F0: 2, F1: 0}
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v9)
    }
    var v13 uint64 = 288
    if (true && (!(116 <= 42))) {
        var v14 bool = false
        _ = v14
    } else {
        if (!(true && ((v13 + v13) >= v13))) {
            v15 := []int{}
            v16 := GenS1{F0: genFn0(9223372036854775807), F1: 1000}
            for fi17 := 0; fi17 < 10; fi17++ {
                if (false || (128 <= 128)) {
                    if (2 <= 0) {
                        v13 = (v13 ^ v13)
                        for fi18 := 0; fi18 < 8; fi18++ {
                            if (2147483647 > 2) {
                                v9 += (v9 + v9)
                            }
                            var v19 bool = true
                            _ = fi18
                            _ = v19
                        }
                    }
                    fmt.Println(v9, fi17)
                }
                for fi20 := 0; fi20 < 5; fi20++ {
                    fmt.Println(v9, fi17, p0)
                    v13 += (v13 / ((v13) | 1))
                    switch (fi20 & -862) {
                    case 2147483647:
                        var v21 int = fi20
                        _ = v21
                    case 128:
                        p0 = (p0 / ((p0) | 1))
                    }
                    _ = fi20
                }
                fmt.Println(v9, p0)
                _ = fi17
            }
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v13
    _ = v9
    return 766
}

func main() {
    var v1 float64 = (0.0 * 1e10)
    v2 := []int{32767}
    var v3 int = 4294967295
    var v4 string = ("\n\t" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t")))
    var v5 int64 = 42
    for fi22 := 0; fi22 < 16; fi22++ {
        if true {
            switch "" {
            case "a":
                if len(v2) > 0 {
                    v2[0] = v3
                }
            case "foo bar":
                var v23 uint32 = 926
                _ = v23
            }
            if true {
                var v24 int = (v3 | v3)
                _ = v24
            }
            var v25 string = "\n\t"
            _ = v25
        }
        v26 := GenS1{F0: 42, F1: -32768}
        for fi27 := 0; fi27 < 2; fi27++ {
            v3 -= (v3 >> ((42) & 63))
            _ = fi27
        }
        _ = fi22
        _ = v26
    }
    if (!(974 <= 65535)) {
        v1 = 0.0
        for fi28 := 0; fi28 < 9; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                switch -46 {
                case 0:
                    var v30 int8 = 1
                    _ = v30
                }
                _ = fi29
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                v3 += 127
                switch v4 {
                case "hello":
                    v2 = append(v2, (fi31 << ((18446744073709551615) & 63)))
                    if ((0 >= 127) || false) {
                        v32 := GenS2{F0: 127, F1: 388, F2: 4294967295}
                        _ = v32
                    }
                    switch (v3 * 127) {
                    case -2147483648:
                        v5 = 1
                        v33 := GenS3{F0: genFn0(-322)}
                        _ = v33
                    case 517:
                        var v34 uint32 = 353
                        _ = v34
                    }
                    var v35 bool = false
                    var v36 int32 = 42
                    _ = v35
                    _ = v36
                case "":
                    fmt.Println(fi31, v3, v1)
                }
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v4)
                    switch v4 {
                    case "a":
                        fmt.Println(fi28, v1)
                    default:
                        v37 := GenS3{F0: 343}
                        _ = v37
                    }
                } else {
                    var v38 uint = 65535
                    if false {
                        if true {
                            v39 := GenS4{F0: v4, F1: -50}
                            fmt.Println(v38, v4, v1)
                            var v40 uint16 = 765
                            _ = v39
                            _ = v40
                        } else {
                            fmt.Println(v4, v3)
                            switch v4 {
                            case "\n\t":
                                v41 := GenS2{F0: 125, F1: 530, F2: v38}
                                _ = v41
                            case "foo bar":
                                v38 = (^(v38))
                            default:
                                v3 = (fi28 | (v3 & (v3 - 128)))
                            }
                            switch "" {
                            case "foo bar":
                                if ((-2147483648 == -1) && (449 > 2)) {
                                    if (v5 == v5) {
                                        fmt.Println(fi28)
                                        if true {
                                            if (0 == 0) {
                                                fmt.Println(v3)
                                            } else {
                                                if (false && false) {
                                                    fmt.Println(fi28, v1, fi28)
                                                    var v42 uint32 = 2147483647
                                                    var v43 int64 = (-(v5))
                                                    var v44 uint64 = 479
                                                    switch fi31 {
                                                    case 127:
                                                        var v45 uint16 = 2
                                                        _ = v45
                                                    case -522:
                                                        v46 := GenS4{F0: v4, F1: 42}
                                                        _ = v46
                                                    case 104:
                                                        var v47 uint16 = 83
                                                        v38 += (v38 ^ 2)
                                                        _ = v47
                                                    }
                                                    _ = v42
                                                    _ = v43
                                                    _ = v44
                                                }
                                            }
                                            v38 -= v38
                                            v48 := GenS5{F0: "foo bar", F1: 908}
                                            _ = v48
                                        } else {
                                            var v49 float64 = v1
                                            switch v4 {
                                            case "a":
                                                var v50 int64 = -1
                                                _ = v50
                                            case "hello":
                                                var v51 float64 = float64(float32(0.0))
                                                _ = v51
                                            case "\n\t":
                                                if false {
                                                    var v52 uint = (v38 ^ v38)
                                                    _ = v52
                                                }
                                                var v53 float32 = 0.0
                                                switch fi28 {
                                                case 4294967295:
                                                    switch 343 {
                                                    case 0:
                                                        fmt.Println(v4, v1)
                                                        v1 = v1
                                                    case -42:
                                                        fmt.Println(v53, v5)
                                                    default:
                                                        fmt.Println(v3, v53)
                                                    }
                                                default:
                                                    if (v53 < v53) {
                                                        switch fi31 {
                                                        case 134:
                                                            var v54 float64 = -2.25
                                                            _ = v54
                                                        case 127:
                                                            var v55 float32 = v53
                                                            _ = v55
                                                        case 32768:
                                                            v49 = v49
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v56 := GenS6{F0: (v5 - (v5 - (v5 / ((-1) | 1)))), F1: (fi28 + fi28)}
                                                            v57 := GenS7{F0: v4}
                                                            v38 = v38
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        var v58 int16 = 42
                                                        _ = v58
                                                    } else {
                                                        if false {
                                                            switch "foo bar" {
                                                            case "":
                                                                var v59 int32 = -2147483648
                                                                _ = v59
                                                            }
                                                            v60 := GenS6{F0: 926, F1: (-(fi31))}
                                                            v53 = float32(v49)
                                                            _ = v60
                                                        } else {
                                                            fmt.Println(v4, v1)
                                                        }
                                                        _ = len(v2)
                                                        switch (fi28 - 32767) {
                                                        case 0:
                                                            var v61 int = v3
                                                            _ = v61
                                                        case 1:
                                                            v62 := GenS5{F0: "a", F1: -2147483648}
                                                            _ = v62
                                                        default:
                                                            v4 = "\n\t"
                                                        }
                                                    }
                                                }
                                                _ = v53
                                            }
                                            _ = v49
                                        }
                                    }
                                    var v63 string = ""
                                    var v64 uint = v38
                                    _ = v63
                                    _ = v64
                                } else {
                                    v4 = v4
                                    var v65 float32 = 1.5
                                    _ = v65
                                }
                            case "\n\t":
                                switch v4 {
                                case "a":
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (-1 != 0) {
                                            var v66 uint64 = 42
                                            _ = v66
                                        } else {
                                            var v67 uint8 = 1
                                            _ = v67
                                        }
                                    } else {
                                        var v68 uint32 = 4294967295
                                        _ = v68
                                    }
                                case "hello":
                                    if (fi28 > (fi31 >> ((787) & 63))) {
                                        v3 = 42
                                        switch (fi28 ^ v3) {
                                        case 0:
                                            v69 := GenS8{F0: v5, F1: (^(v38))}
                                            _ = v69
                                        default:
                                            switch (fi28 * v3) {
                                            case -2147483648:
                                                var v70 uint16 = 2
                                                _ = v70
                                            case 42:
                                                var v71 int32 = 0
                                                var v72 string = v4
                                                if true {
                                                    fmt.Println(fi31, v3)
                                                }
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                    }
                                default:
                                    v73 := GenS9{F0: (v3 >> ((genFn1(-128)) & 63)), F1: v4, F2: v38}
                                    _ = v73
                                }
                            default:
                                switch 32767 {
                                case 171:
                                    var v74 uint = (v38 - v38)
                                    _ = v74
                                default:
                                    v75 := GenS5{F0: "\n\t", F1: 0}
                                    _ = v75
                                }
                            }
                        }
                        v38 = 65535
                    }
                    _ = v38
                }
                _ = fi31
            }
            _ = fi28
        }
        if false {
            var v76 string = ("foo bar" + (v4 + ("\n\t" + ("foo bar" + (("" + (v4 + ("\n\t" + "\n\t"))) + (v4 + "\n\t"))))))
            _ = v76
        }
    } else {
        fmt.Println(v1, v1)
    }
    for fi77 := 0; fi77 < 5; fi77++ {
        if true {
            v1 += 1e10
            if len(v2) > 0 {
                v2[0] = v3
            }
            if (4294967295 == genFn1(0)) {
                switch (("" + ((("" + (v4 + "")) + ("\n\t" + "\n\t")) + "hello")) + v4) {
                case "":
                    for fi78 := 0; fi78 < 13; fi78++ {
                        v4 = v4
                        _ = fi78
                    }
                case "foo bar":
                    for fi79 := 0; fi79 < 16; fi79++ {
                        if false {
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = v3
                                }
                                var v80 int64 = v5
                                _ = v80
                            }
                        }
                        var v81 int16 = 32767
                        _ = fi79
                        _ = v81
                    }
                default:
                    _ = len(v2)
                }
                for fi82 := 0; fi82 < 20; fi82++ {
                    var v83 int16 = -1
                    var v84 bool = true
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                fmt.Println(v4, v5, v1)
            }
        } else {
            fmt.Println(v1, fi77, v4)
            for fi85 := 0; fi85 < 3; fi85++ {
                v86 := GenS2{F0: 127, F1: 32767, F2: 18446744073709551615}
                _ = fi85
                _ = v86
            }
            var v87 uint8 = 42
            _ = v87
        }
        fmt.Println(v3, v1)
        _ = fi77
    }
    var v88 float64 = -2.25
    var v89 uint64 = 563
    if true {
        if (!false) {
            var v90 int32 = -1
            _ = v90
        }
        var v91 uint = genFn1(-1)
        _ = v91
    } else {
        v4 = v4
    }
    var v92 float64 = v88
    var v93 int64 = (^(v5))
    v2 = append(v2, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((v88 <= -2.25) && (!false)) {
        var v94 bool = (65535 == 743)
        v95 := GenS1{F0: v89, F1: -1}
        _ = v94
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v89
    _ = v92
    _ = v93
}
