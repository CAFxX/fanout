package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint16; F2 int64 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 int8, p2 uint32) int16 {
    v4 := []int{4294967295, 4294967295}
    var v5 uint64 = 737
    var v6 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 int8, p1 int, p2 bool) bool {
    if p2 {
        var v7 int = (p1 * p1)
        _ = v7
    } else {
        v8 := []int{2147483647, (p1 >> ((948) & 63)), p1}
        switch "a" {
        case "hello":
            v9 := GenS1{F0: 37}
            _ = v9
        }
        p0 -= p0
        _ = v8
    }
    switch ("foo bar" + (("\n\t" + "") + "a")) {
    case "\n\t":
        fmt.Println(p2, p1)
    case "foo bar":
        v10 := []int{}
        _ = v10
    case "":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 float32 = 0.0
            if p2 {
                p2 = p2
            }
            var v13 int16 = 1000
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        var v14 uint16 = 996
        _ = v14
    }
    p1 += (p1 >> ((42) & 63))
    p1 += (p1 ^ p1)
    if p2 {
        switch ("" + "a") {
        case "foo bar":
            for fi15 := 0; fi15 < 1; fi15++ {
                fmt.Println(fi15)
                _ = fi15
            }
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(fi16, p2, fi17)
                fmt.Println(p1, fi17)
                _ = fi17
            }
            v18 := GenS2{F0: p2}
            _ = fi16
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := []int{160, 9223372036854775807}
    var v2 string = "a"
    var v3 float32 = 0.1
    fmt.Println(v3, v3)
    v19 := GenS3{F0: 507, F1: 65535}
    switch 32767 {
    case 1:
        for fi20 := 0; fi20 < 4; fi20++ {
            if (true && true) {
                v2 = ""
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS4{F0: 685, F1: 680}
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (-278 < (^(fi20))) {
                        v3 -= ((v3 - float32((((float64(0.0) * 1e10) - 0.0) * 3.14159))) * 3.14159)
                        v24 := GenS2{F0: (!(!true))}
                        if true {
                            switch v2 {
                            case "":
                                fmt.Println(fi23)
                            default:
                                v25 := GenS4{F0: 303, F1: 4294967295}
                                v2 = "\n\t"
                                if (-1 <= 9223372036854775807) {
                                    var v26 int16 = 728
                                    var v27 uint64 = 689
                                    _ = v26
                                    _ = v27
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi23 / (((fi20 & (fi20 + 4294967295))) | 1))
                                    }
                                    v2 = v2
                                }
                                _ = v25
                            }
                            v28 := GenS4{F0: 4294967295, F1: 2}
                            if (236 >= genFn0(42, -22, 4294967295)) {
                                if false {
                                    v29 := GenS5{F0: 42}
                                    _ = v29
                                } else {
                                    fmt.Println(v3, fi23, v3)
                                    var v30 uint16 = 994
                                    var v31 uint8 = 127
                                    _ = v30
                                    _ = v31
                                }
                                if (852 > 2) {
                                    if (true && false) {
                                        switch -483 {
                                        case 32768:
                                            if true {
                                                switch (("" + v2) + v2) {
                                                case "":
                                                    v3 = ((1.5 + float32(float64(v3))) * (((v3 - ((v3 - (float32(3.14159) * 1e10)) / (v3 - 3.14159))) + v3) + v3))
                                                case "\n\t":
                                                    v32 := GenS4{F0: 0, F1: 4294967295}
                                                    _ = v32
                                                default:
                                                    v33 := GenS6{F0: 106, F1: 255, F2: 2}
                                                    _ = v33
                                                }
                                                v3 = v3
                                                var v34 float64 = 0.1
                                                _ = v34
                                            } else {
                                                v3 -= (0.0 - v3)
                                            }
                                        case 2147483647:
                                            var v35 int32 = 1
                                            _ = v35
                                        }
                                        var v36 int16 = 1000
                                        fmt.Println(v3)
                                        _ = v36
                                    }
                                    switch (^(fi20)) {
                                    case 2147483647:
                                        v2 = (v2 + ((v2 + v2) + v2))
                                    default:
                                        fmt.Println(v3, v2)
                                    }
                                    var v37 float64 = -2.25
                                    _ = v37
                                } else {
                                    if (!((false && (!false)) && (127 == 42))) {
                                        fmt.Println(fi20, fi23)
                                        v2 = "\n\t"
                                        var v38 uint16 = 255
                                        _ = v38
                                    } else {
                                        v39 := GenS4{F0: 581, F1: 4294967295}
                                        switch fi23 {
                                        case -2147483648:
                                            var v40 int8 = 1
                                            _ = v40
                                        case 413:
                                            fmt.Println(v2, v3)
                                        }
                                        v41 := GenS1{F0: 100}
                                        _ = v39
                                        _ = v41
                                    }
                                }
                                var v42 uint32 = 42
                                var v43 uint32 = 2
                                var v44 uint16 = 42
                                if false {
                                    fmt.Println(v42)
                                    v3 = v3
                                }
                                _ = v42
                                _ = v43
                                _ = v44
                            } else {
                                if true {
                                    v3 -= float32(((-2.25 / 1.0) / 1e10))
                                    if true {
                                        fmt.Println(v2, fi20, v3)
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(fi20, v2, v2)
                                        case "a":
                                            var v45 float64 = float64(v3)
                                            var v46 uint32 = 349
                                            _ = v45
                                            _ = v46
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = fi23
                                            }
                                        }
                                        if true {
                                            if (!true) {
                                                var v47 int16 = 32767
                                                if (2 <= 255) {
                                                    v3 = v3
                                                }
                                                var v48 int32 = 42
                                                v49 := GenS7{F0: (true && false), F1: 667, F2: int64(v48)}
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    v50 := GenS8{F0: 42, F1: 119}
                                                    _ = v50
                                                }
                                                v51 := GenS4{F0: 688, F1: 4294967295}
                                                var v52 uint8 = 255
                                                _ = v51
                                                _ = v52
                                            }
                                            var v53 int32 = 1
                                            _ = v53
                                        }
                                        var v54 uint8 = 2
                                        v55 := GenS7{F0: true, F1: 32767, F2: 9223372036854775807}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        v1 = append(v1, -9223372036854775808)
                                    }
                                }
                                fmt.Println(fi23, fi20)
                            }
                            _ = v28
                        }
                        _ = v24
                    }
                    v56 := GenS5{F0: 574}
                    _ = fi23
                    _ = v56
                }
            } else {
                switch ("foo bar" + v2) {
                case "hello":
                    var v57 int = fi20
                    _ = v57
                case "a":
                    v58 := GenS8{F0: -1, F1: 1}
                    _ = v58
                case "":
                    v3 = v3
                default:
                    var v59 uint16 = 861
                    _ = v59
                }
                v1 = append(v1, 32767)
                var v60 uint8 = 200
                _ = v60
            }
            v61 := GenS9{F0: 2}
            _ = fi20
            _ = v61
        }
    case 605:
        var v62 bool = false
        _ = v62
    case 42:
        v2 = v2
    }
    fmt.Println(v3)
    var v63 uint = 42
    var v64 int8 = -28
    var v65 uint8 = 0
    var v66 int32 = -2147483648
    v67 := GenS3{F0: 2, F1: 42}
    v1 = append(v1, 42)
    fmt.Println(v3)
    for fi68 := 0; fi68 < 5; fi68++ {
        if false {
            _ = len(v1)
            var v69 string = v2
            _ = v69
        } else {
            for fi70 := 0; fi70 < 13; fi70++ {
                switch (("\n\t" + v2) + "") {
                case "hello":
                    v2 = v2
                case "a":
                    if ((v64 & (-(v64))) < (v64 * 127)) {
                        v65 += (^(v65))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v71 float32 = (0.1 + v3)
                        var v72 int32 = (v66 - v66)
                        if ((fi70 <= fi70) || (2147483648 <= 2)) {
                            _ = len(v1)
                            if true {
                                v1 = append(v1, fi70)
                            }
                        } else {
                            switch fi70 {
                            case -239:
                                switch v2 {
                                case "a":
                                    v3 += v71
                                default:
                                    if (int32(v64) > v72) {
                                        v73 := GenS10{F0: v72, F1: ("" + "\n\t")}
                                        _ = v73
                                    } else {
                                        var v74 uint = v63
                                        _ = v74
                                    }
                                }
                            }
                            _ = len(v1)
                        }
                        _ = v71
                        _ = v72
                    } else {
                        var v75 float64 = 1e10
                        switch ("foo bar" + ("" + "a")) {
                        case "":
                            v63 += v63
                        default:
                            var v76 int64 = 42
                            _ = v76
                        }
                        _ = v75
                    }
                case "foo bar":
                    v66 = (v66 | int32(v65))
                }
                v77 := GenS3{F0: 1, F1: 0}
                fmt.Println(v66, v2)
                _ = fi70
                _ = v77
            }
            var v78 int16 = 32767
            if len(v1) > 0 {
                v1[0] = fi68
            }
            _ = v78
        }
        if false {
            v64 = (v64 * (v64 - (v64 + -1)))
        } else {
            _ = len(v1)
            v1 = append(v1, (fi68 << ((4294967295) & 63)))
        }
        v66 = 42
        _ = fi68
    }
    for fi79 := 0; fi79 < 13; fi79++ {
        for fi80 := 0; fi80 < 13; fi80++ {
            v81 := GenS6{F0: (-(v64)), F1: 0, F2: 42}
            var v82 int16 = -554
            fmt.Println(v2)
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = fi79
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
