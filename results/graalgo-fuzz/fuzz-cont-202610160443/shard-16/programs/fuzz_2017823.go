package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 int64; F1 string; F2 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int8; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8, p1 uint8, p2 uint8) uint16 {
    p2 = (p2 & (p2 << ((482) & 63)))
    p2 = p1
    for fi6 := 0; fi6 < 19; fi6++ {
        p1 = (p1 + p2)
        _ = fi6
    }
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1() uint8 {
    if (-2147483648 >= 916) {
        v7 := GenS1{F0: 255, F1: 1, F2: 2147483647}
        var v8 int32 = -1
        _ = v7
        _ = v8
    } else {
        if false {
            var v9 string = "hello"
            if false {
                if (!false) {
                    if false {
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v11 := []int{fi10, fi10, fi10}
                            switch 271 {
                            case 9223372036854775807:
                                v9 = v9
                            case 498:
                                if (42 < 255) {
                                    if len(v11) > 0 {
                                        v11[0] = 32768
                                    }
                                }
                            default:
                                fmt.Println(v9, fi10)
                            }
                            v9 = ""
                            _ = fi10
                            _ = v11
                        }
                        var v12 int16 = 32767
                        _ = v12
                    }
                    var v13 uint8 = 58
                    _ = v13
                } else {
                    if true {
                        var v14 bool = (!true)
                        v15 := GenS1{F0: 558, F1: 2, F2: 392}
                        if v14 {
                            var v16 uint8 = 2
                            v17 := []int{-351, 0, 1}
                            if len(v17) > 0 {
                                v17[0] = -1
                            }
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                        _ = v15
                    }
                }
                fmt.Println(v9)
            } else {
                var v18 uint64 = 323
                v18 += (v18 % ((v18) | 1))
                var v19 int = -296
                _ = v18
                _ = v19
            }
            v9 = ("a" + v9)
            _ = v9
        } else {
            var v20 uint16 = 42
            _ = v20
        }
    }
    if (!false) {
        v21 := GenS1{F0: 641, F1: 42, F2: 0}
        var v22 float64 = -2.25
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS2{F0: 2, F1: (fi23 + (^(fi23))), F2: (!true)}
            var v25 int32 = 42
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v21
        _ = v22
    } else {
        var v26 int = 0
        v27 := []int{(v26 * 141), v26, 4294967295, (v26 << ((42) & 63))}
        if ("" == "hello") {
            var v28 int16 = 590
            if (genFn0(100, 1, 127) == 42) {
                fmt.Println(v28)
                var v29 uint8 = 128
                v27 = append(v27, -9223372036854775808)
                _ = v29
            } else {
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            }
            if (18446744073709551615 > 0) {
                v30 := GenS3{F0: 1, F1: ("" + ""), F2: 328}
                _ = len(v27)
                _ = v30
            }
            _ = v28
        } else {
            var v31 int8 = 93
            _ = len(v27)
            _ = v31
        }
        _ = v26
        _ = v27
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        if (1 > 42) {
            fmt.Println(fi32)
            if (6 <= 206) {
                var v33 float64 = 0.1
                _ = v33
            } else {
                var v34 int64 = 72
                switch "" {
                case "":
                    for fi35 := 0; fi35 < 18; fi35++ {
                        if (!true) {
                            var v36 int = fi35
                            _ = v36
                        }
                        v37 := GenS4{F0: 1, F1: 32767}
                        v38 := []int{-807, fi35}
                        _ = fi35
                        _ = v37
                        _ = v38
                    }
                default:
                    for fi39 := 0; fi39 < 15; fi39++ {
                        v40 := []int{9223372036854775807, (fi39 % ((32768) | 1)), fi39, (fi39 | (fi32 + 42))}
                        _ = fi39
                        _ = v40
                    }
                }
                v41 := []int{32768, -1, (fi32 << ((1) & 63))}
                _ = v34
                _ = v41
            }
        }
        var v42 string = "foo bar"
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := []int{}
            var v45 uint16 = 255
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi32
        _ = v42
    }
    return 128
}

func main() {
    var v1 uint = 85
    v2 := make(map[int]bool)
    var v3 uint64 = 829
    var v4 int16 = 32767
    v5 := []string{"foo bar", ("\n\t" + "hello"), "\n\t"}
    var v46 uint8 = 127
    v47 := GenS4{F0: (v4 % (((v4 | -1)) | 1)), F1: v4}
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v48 uint32 = 410
    switch 42 {
    case -560:
        switch 4294967295 {
        case 128:
            fmt.Println(v3, v46, v1)
        }
    case 128:
        v3 = (v3 ^ (v3 % ((2) | 1)))
    case 127:
        v48 -= (v48 & v48)
    default:
        var v49 int8 = 120
        _ = v49
    }
    fmt.Println(v48)
    var v50 int = -920
    var v51 int = 32767
    for fi52 := 0; fi52 < 4; fi52++ {
        for fi53 := 0; fi53 < 7; fi53++ {
            fmt.Println(v48)
            _ = fi53
        }
        fmt.Println(v50)
        switch int(v3) {
        case 128:
            fmt.Println(v3)
        case -1:
            if false {
                v46 = (^(v46))
                v51 += fi52
            } else {
                var v54 int32 = -1
                for fi55 := 0; fi55 < 2; fi55++ {
                    delete(v2, (v51 - (v51 & v50)))
                    _ = fi55
                }
                _ = v54
            }
        default:
            if true {
                for fi56 := 0; fi56 < 7; fi56++ {
                    var v57 uint16 = genFn0(127, uint8(v48), (v46 ^ uint8(v48)))
                    v51 = -2147483648
                    switch "hello" {
                    case "hello":
                        var v58 bool = true
                        _ = v58
                    }
                    _ = fi56
                    _ = v57
                }
            } else {
                if true {
                    var v59 uint = (^(v1))
                    v5 = append(v5, "foo bar")
                    _ = v59
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v60 int64 = -9223372036854775808
                    _ = v60
                }
            }
        }
        _ = fi52
    }
    fmt.Println(v50, v51)
    var v61 int8 = -128
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 uint64 = (v3 / ((v3) | 1))
        _ = fi62
        _ = v63
    }
    if true {
        if true {
            var v64 int64 = -749
            switch v50 {
            case -1:
                var v65 int8 = int8(v1)
                _ = v65
            case 4294967295:
                var v66 bool = ((^(v50)) != (v50 & (v51 << ((v1) & 63))))
                _ = v66
            case 9223372036854775807:
                var v67 int16 = 32767
                _ = v67
            }
            v68 := GenS4{F0: (-(v4)), F1: v4}
            _ = v64
            _ = v68
        } else {
            var v69 uint16 = 2
            switch "foo bar" {
            case "hello":
                for fi70 := 0; fi70 < 9; fi70++ {
                    var v71 uint32 = 867
                    _ = fi70
                    _ = v71
                }
            case "":
                v72 := GenS5{F0: v4, F1: (^(v46))}
                _ = v72
            case "\n\t":
                v1 = v1
            default:
                v48 = (v48 - v48)
            }
            _ = v69
        }
        v3 = (v3 / ((42) | 1))
    } else {
        for fi73 := 0; fi73 < 19; fi73++ {
            if true {
                if (false && ((!(0 > uint64(v48))) || false)) {
                    fmt.Println(v1, v1)
                    fmt.Println(v4, v50, v48)
                    if (((!false) && (int32(v3) >= 1)) && (((!(v61 > (v61 & (v61 * v61)))) && (true && true)) || (!false))) {
                        v74 := GenS6{F0: (v48 | v48), F1: 12, F2: 0}
                        _ = v74
                    }
                }
                switch v51 {
                case -9223372036854775808:
                    switch (fi73 + (v50 * (v50 >> ((25) & 63)))) {
                    case 32768:
                        v75 := GenS7{F0: 1, F1: 238, F2: (v3 | v3)}
                        _ = v75
                    case -917:
                        for fi76 := 0; fi76 < 12; fi76++ {
                            v77 := GenS8{F0: uint64(v1), F1: 0}
                            if (65535 == 42) {
                                v78 := GenS5{F0: v4, F1: (v46 & (v46 & (v46 >> (((v1 * 2)) & 63))))}
                                _ = v78
                            }
                            _ = fi76
                            _ = v77
                        }
                    default:
                        for fi79 := 0; fi79 < 2; fi79++ {
                            v4 = v4
                            var v80 string = "foo bar"
                            v81 := v2[(^(v51))]
                            _ = fi79
                            _ = v80
                            _ = v81
                        }
                    }
                }
                v82 := GenS3{F0: 466, F1: "foo bar", F2: v1}
                _ = v82
            } else {
                fmt.Println(v46, v61)
                var v83 float64 = 1.5
                _ = v83
            }
            v84 := GenS3{F0: -9223372036854775808, F1: "hello", F2: v1}
            v48 -= (v48 + v48)
            _ = fi73
            _ = v84
        }
        if (!(1e10 <= (1.5 + 1.5))) {
            v85 := GenS4{F0: -32768, F1: v4}
            _ = v85
        } else {
            fmt.Println(v4, v61, v46)
            for fi86 := 0; fi86 < 3; fi86++ {
                var v87 string = "foo bar"
                var v88 float32 = 1e10
                var v89 bool = true
                _ = fi86
                _ = v87
                _ = v88
                _ = v89
            }
            v90 := v2[v51]
            _ = v90
        }
        v91 := GenS8{F0: uint64(v46), F1: uint16(v1)}
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v61
}
