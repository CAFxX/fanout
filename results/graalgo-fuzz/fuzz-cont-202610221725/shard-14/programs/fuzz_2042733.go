package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int16; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS10 struct { F0 bool }
func genFn0(p0 bool) uint64 {
    v4 := GenS1{F0: 18446744073709551615, F1: p0, F2: 42}
    p0 = p0
    if p0 {
        var v5 int = 32767
        var v6 uint64 = 42
        v5 = v5
        _ = v5
        _ = v6
    }
    v7 := []int{127, 9223372036854775807}
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: 0, F1: p0, F2: -632}
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS2{F0: 0, F1: 127, F2: 1}
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS3{F0: 4294967295}
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    default:
        var v13 uint32 = 65535
        _ = v13
    }
    _ = p0
    _ = v4
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = ("foo bar" + "\n\t")
    var v3 int8 = 100
    v14 := GenS3{F0: 2147483648}
    for fi15 := 0; fi15 < 9; fi15++ {
        v2 = "foo bar"
        v16 := GenS3{F0: 754}
        var v17 bool = (65535 < 832)
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint64 = 0
        _ = fi18
        _ = v19
    }
    v2 = (v2 + (v2 + v2))
    if (true || (!true)) {
        if false {
            if (-883 == 42) {
                v20 := GenS4{F0: 0, F1: v2}
                if false {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v2 = ((v2 + v2) + v2)
                        if false {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                if true {
                                    switch (fi22 | -1) {
                                    case 0:
                                        v3 += (v3 & (v3 | v3))
                                    case 128:
                                        switch (fi21 ^ fi22) {
                                        case -735:
                                            if true {
                                                if true {
                                                    v23 := GenS3{F0: 42}
                                                    fmt.Println(v3)
                                                    v24 := GenS3{F0: 2147483648}
                                                    _ = v23
                                                    _ = v24
                                                }
                                                v2 = (v2 + v2)
                                                v2 = (v2 + "")
                                            }
                                        case 4294967295:
                                            var v25 uint64 = 1
                                            _ = v25
                                        case 42:
                                            v3 = (v3 >> ((255) & 63))
                                        default:
                                            v1[((v2 + "a") + v2)] = ((-(v3)) == v3)
                                        }
                                    }
                                    if ((fi22 ^ 128) != fi22) {
                                        switch (fi22 % ((fi22) | 1)) {
                                        case -1:
                                            fmt.Println(v3)
                                        case 705:
                                            v2 = (v2 + v2)
                                        default:
                                            v2 = ("hello" + v2)
                                        }
                                    } else {
                                        if (((9223372036854775807 > -129) && (71 > 255)) && (0 != 519)) {
                                            v3 = v3
                                            var v26 uint64 = 817
                                            v27 := GenS3{F0: 1}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 int64 = 0
                                            if true {
                                                switch fi21 {
                                                case 9223372036854775807:
                                                    var v29 int = fi21
                                                    _ = v29
                                                case 585:
                                                    v30 := GenS2{F0: 32767, F1: (^(v3)), F2: 1}
                                                    _ = v30
                                                case -341:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v31 int64 = 1
                                                        _ = v31
                                                    case "a":
                                                        v32 := GenS5{F0: -340}
                                                        _ = v32
                                                    case "\n\t":
                                                        v3 -= v3
                                                    }
                                                }
                                                var v33 uint16 = 882
                                                v2 = v2
                                                _ = v33
                                            }
                                            _ = v28
                                        }
                                        v1["foo bar"] = true
                                    }
                                    var v34 int = (fi22 << ((4294967295) & 63))
                                    _ = v34
                                } else {
                                    if (18446744073709551615 >= 42) {
                                        if true {
                                            fmt.Println(fi22, v2, fi22)
                                        }
                                        if (42 <= 32767) {
                                            var v35 bool = true
                                            if v35 {
                                                v36 := GenS6{F0: v3, F1: -592, F2: true}
                                                var v37 bool = v35
                                                v38 := GenS7{F0: -9223372036854775808}
                                                var v39 uint = 18446744073709551615
                                                fmt.Println(v2)
                                                if (32767 == 0) {
                                                    v39 = (v39 * (v39 % ((65535) | 1)))
                                                    var v40 int64 = 149
                                                    _ = v40
                                                } else {
                                                    var v41 string = ("hello" + v2)
                                                    v42 := GenS6{F0: v3, F1: 32767, F2: v37}
                                                    _ = v41
                                                    _ = v42
                                                }
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v43 := []int{}
                                                    _ = v43
                                                case "hello":
                                                    v35 = v35
                                                }
                                                fmt.Println(v2)
                                                v2 = "hello"
                                                var v44 uint32 = 2147483647
                                                v35 = v35
                                                v45 := GenS2{F0: 1, F1: (v3 ^ (v3 << ((545) & 63))), F2: 255}
                                                _ = v44
                                                _ = v45
                                            }
                                            v46 := []int{fi22, (fi22 * -485), fi22, (-(fi21))}
                                            _ = v35
                                            _ = v46
                                        } else {
                                            v47 := []int{}
                                            v48 := v1[(((v2 + v2) + v2) + v2)]
                                            _ = v47
                                            _ = v48
                                        }
                                        v3 = (^(v3))
                                    } else {
                                        v49 := GenS5{F0: 999}
                                        var v50 float32 = (1e10 + 1.5)
                                        v51 := []int{}
                                        if false {
                                            var v52 int64 = 0
                                            if true {
                                                fmt.Println(v2, v50)
                                                if (0 != 18446744073709551615) {
                                                    var v53 uint = 255
                                                    v53 += (v53 / (((^(v53))) | 1))
                                                    fmt.Println(fi22)
                                                    _ = v53
                                                }
                                            } else {
                                                if (true && false) {
                                                    v50 += v50
                                                    v54 := GenS5{F0: 1}
                                                    _ = v54
                                                }
                                            }
                                            v50 = float32(float64((v50 * v50)))
                                            v52 += 1
                                            switch (v2 + ("\n\t" + v2)) {
                                            case "":
                                                switch (fi21 | (fi22 ^ -172)) {
                                                case 32768:
                                                    if true {
                                                        var v55 int8 = 42
                                                        v3 = v3
                                                        _ = v55
                                                    }
                                                }
                                            }
                                            _ = v52
                                        } else {
                                            var v56 int = 127
                                            _ = v56
                                        }
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                    if true {
                                        var v57 float64 = 0.1
                                        _ = v57
                                    }
                                }
                                _ = fi22
                            }
                        }
                        _ = fi21
                    }
                }
                _ = v20
            }
            var v58 int16 = 868
            _ = v58
        }
        v2 = (v2 + v2)
    } else {
        delete(v1, ((v2 + "") + v2))
        v2 = v2
        fmt.Println(v3, v3)
    }
    if false {
        switch 127 {
        case -337:
            if (false && true) {
                for fi59 := 0; fi59 < 19; fi59++ {
                    v60 := []int{(fi59 >> ((0) & 63))}
                    _ = fi59
                    _ = v60
                }
                var v61 int8 = 13
                if false {
                    v61 -= (v61 * v3)
                    for fi62 := 0; fi62 < 15; fi62++ {
                        if false {
                            switch fi62 {
                            case -453:
                                v1["hello"] = true
                            case 4294967295:
                                v61 = 42
                            case 1:
                                var v63 string = "hello"
                                _ = v63
                            default:
                                switch (fi62 & (fi62 ^ 32768)) {
                                case 32768:
                                    v64 := []int{fi62, -269, -470}
                                    _ = v64
                                default:
                                    if true {
                                        if false {
                                            if false {
                                                for fi65 := 0; fi65 < 10; fi65++ {
                                                    var v66 float64 = -2.25
                                                    if true {
                                                        v3 = v61
                                                        var v67 string = "hello"
                                                        _ = v67
                                                    } else {
                                                        if true {
                                                            var v68 int = (^(fi62))
                                                            _ = v68
                                                        } else {
                                                            v61 -= (v3 / ((v61) | 1))
                                                        }
                                                    }
                                                    fmt.Println(v3, v3)
                                                    _ = fi65
                                                    _ = v66
                                                }
                                                var v69 int32 = -1
                                                var v70 int16 = 32767
                                                _ = v69
                                                _ = v70
                                            } else {
                                                switch "a" {
                                                case "":
                                                    v61 += (^(v61))
                                                case "a":
                                                    switch fi62 {
                                                    case -443:
                                                        _ = len(v1)
                                                    case 2147483647:
                                                        var v71 string = ((v2 + "\n\t") + v2)
                                                        _ = v71
                                                    default:
                                                        v72 := GenS8{F0: 1}
                                                        _ = v72
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v61, fi62, v61)
                                                }
                                            }
                                            if true {
                                                v73 := GenS9{F0: 42, F1: 2147483647, F2: 255}
                                                var v74 string = v2
                                                var v75 int16 = 1000
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                            }
                                        } else {
                                            for fi76 := 0; fi76 < 15; fi76++ {
                                                var v77 uint32 = 42
                                                _ = fi76
                                                _ = v77
                                            }
                                            fmt.Println(fi62)
                                            var v78 int32 = 1
                                            _ = v78
                                        }
                                        v2 = ""
                                    } else {
                                        v61 -= (^(v3))
                                        v79 := []int{fi62, -456, 42}
                                        _ = v79
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v2, v2)
                            v61 += (v3 | (v61 % (((v61 & v61)) | 1)))
                        }
                        v3 += v3
                        _ = fi62
                    }
                } else {
                    v80 := GenS2{F0: 335, F1: (v61 / (((v61 * (v3 | v3))) | 1)), F2: 127}
                    var v81 int = 4294967295
                    _ = v80
                    _ = v81
                }
                _ = v61
            }
        case 0:
            switch -9223372036854775808 {
            case 793:
                var v82 float32 = float32(-2.25)
                _ = v82
            case 0:
                v2 = "\n\t"
            }
        case -509:
            if true {
                if false {
                    fmt.Println(v2)
                    v83 := []int{42, 42, 1000}
                    var v84 uint32 = 65535
                    _ = v83
                    _ = v84
                }
                switch v2 {
                case "hello":
                    v2 = "\n\t"
                }
                var v85 bool = true
                _ = v85
            } else {
                fmt.Println(v2)
                var v86 uint8 = 127
                v87 := GenS10{F0: false}
                _ = v86
                _ = v87
            }
        }
    } else {
        delete(v1, (v2 + (v2 + (v2 + v2))))
    }
    var v88 uint16 = 42
    v3 = v3
    for fi89 := 0; fi89 < 14; fi89++ {
        if false {
            v90 := GenS3{F0: 1}
            v88 += (v88 % ((v88) | 1))
            _ = v90
        } else {
            v88 -= (v88 ^ (v88 & 894))
            if true {
                for fi91 := 0; fi91 < 6; fi91++ {
                    var v92 int8 = (^(v3))
                    v93 := []int{(-(fi91))}
                    var v94 int = (fi89 + (fi89 << ((333) & 63)))
                    v95 := GenS10{F0: (!false)}
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            fmt.Println(v94, v88, v92)
                            fmt.Println(v94, v3)
                        }
                    }
                    _ = fi91
                    _ = v92
                    _ = v93
                    _ = v94
                    _ = v95
                }
            } else {
                if (false && true) {
                    fmt.Println(fi89)
                }
            }
            fmt.Println(v3, fi89, v2)
        }
        _ = fi89
    }
    if false {
        delete(v1, "")
        if (65535 <= 18446744073709551615) {
            v96 := GenS5{F0: -2147483648}
            for fi97 := 0; fi97 < 7; fi97++ {
                v3 = v3
                if false {
                    v3 = (-(v3))
                }
                v2 = ((v2 + v2) + (("a" + "foo bar") + ""))
                _ = fi97
            }
            switch 127 {
            case 128:
                for fi98 := 0; fi98 < 1; fi98++ {
                    delete(v1, ("a" + ((((((v2 + v2) + ((v2 + "\n\t") + "foo bar")) + "\n\t") + (v2 + (("" + "") + ""))) + v2) + ("a" + "foo bar"))))
                    _ = fi98
                }
            case -2147483648:
                v99 := GenS3{F0: 529}
                _ = v99
            default:
                delete(v1, v2)
            }
            _ = v96
        } else {
            v100 := v1[v2]
            v101 := GenS2{F0: 526, F1: v3, F2: 128}
            _ = v100
            _ = v101
        }
    } else {
        for fi102 := 0; fi102 < 16; fi102++ {
            v103 := GenS3{F0: 705}
            v88 += v88
            _ = fi102
            _ = v103
        }
    }
    fmt.Println(v3, v88)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v88
}
