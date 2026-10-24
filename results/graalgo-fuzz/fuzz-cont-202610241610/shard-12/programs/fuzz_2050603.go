package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 uint8; F2 string }
type GenS11 struct { F0 int8; F1 uint16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 string; F1 uint16; F2 uint16 }
type GenS14 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64, p1 uint) int64 {
    v7 := []int{}
    fmt.Println(p0, p1)
    p0 += (^(p0))
    v8 := GenS1{F0: 127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 605
}
func genFn1() int {
    var v9 bool = (65535 >= 2)
    switch 128 {
    case -617:
        v10 := GenS2{F0: -729, F1: -78}
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 string = ("foo bar" + "a")
            if (v9 && (4294967295 < 742)) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v9 = v9
                    var v14 uint8 = 2
                    v15 := GenS3{F0: false, F1: 2, F2: fi11}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            if (4294967295 != 4294967295) {
                var v16 int16 = -32768
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 int = 526
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = fi11
            _ = v12
        }
    case 4294967295:
        var v19 int64 = 775
        _ = v19
    }
    v9 = (v9 || v9)
    if (v9 && ((("hello" < "") || (true || true)) || (v9 && true))) {
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS1{F0: fi20}
                v23 := []int{32768, (fi20 - fi21), fi20}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v9 = (true || false)
        var v24 int32 = -2147483648
        _ = v24
    } else {
        if (!false) {
            switch "hello" {
            case "\n\t":
                var v25 uint64 = 178
                _ = v25
            }
        } else {
            if (!(65535 <= 0)) {
                fmt.Println(v9)
            } else {
                fmt.Println(v9)
                v9 = (!true)
                for fi26 := 0; fi26 < 12; fi26++ {
                    fmt.Println(v9, v9)
                    if true {
                        v27 := []int{(fi26 << ((18446744073709551615) & 63))}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            var v29 uint8 = 164
                            _ = fi28
                            _ = v29
                        }
                        v9 = ((1.5 - (((3.14159 - (0.0 - -2.25)) / 1.5) - 3.14159)) > -2.25)
                        fmt.Println(fi26, fi26)
                    }
                    _ = fi26
                }
            }
        }
        if (-2.25 >= ((1.5 / 0.1) * 1e10)) {
            v30 := GenS4{F0: v9, F1: 597}
            var v31 float64 = (0.0 - 1e10)
            _ = v30
            _ = v31
        } else {
            switch "" {
            case "":
                switch 42 {
                case 32767:
                    var v32 int32 = -2147483648
                    _ = v32
                case 547:
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v34 := GenS1{F0: fi33}
                        v35 := []int{}
                        var v36 uint8 = 1
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                case 0:
                    v37 := []int{203}
                    _ = v37
                default:
                    v38 := GenS5{F0: 65535, F1: 255, F2: 65535}
                    _ = v38
                }
            case "a":
                if false {
                    fmt.Println(v9)
                } else {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        fmt.Println(v9, fi39)
                        v40 := GenS6{F0: -2147483648}
                        var v41 uint = 2
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
            case "hello":
                for fi42 := 0; fi42 < 1; fi42++ {
                    v9 = v9
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v9 = false
                        fmt.Println(fi42, fi42)
                        v44 := GenS7{F0: 1, F1: 614, F2: 2}
                        v9 = (1 <= 127)
                        _ = fi43
                        _ = v44
                    }
                    v9 = v9
                    _ = fi42
                }
            default:
                fmt.Println(v9)
            }
            if (!false) {
                fmt.Println(v9)
                v9 = v9
            } else {
                var v45 int16 = 23
                _ = v45
            }
        }
        fmt.Println(v9)
    }
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 0
    var v3 uint8 = 255
    var v4 string = "\n\t"
    var v5 int16 = (v2 << ((710) & 63))
    var v6 bool = true
    var v46 uint = 985
    v47 := []int{}
    if true {
        var v48 int = 410
        var v49 uint = (v46 + (v46 * uint(v3)))
        _ = v48
        _ = v49
    } else {
        var v50 uint64 = 2
        for fi51 := 0; fi51 < 11; fi51++ {
            if v6 {
                var v52 bool = v6
                v5 = v2
                _ = v52
            }
            v6 = v6
            _ = fi51
        }
        _ = v50
    }
    fmt.Println(v5, v5, v6)
    if true {
        v53 := v1["\n\t"]
        if true {
            fmt.Println(v53, v53)
            for fi54 := 0; fi54 < 12; fi54++ {
                var v55 float32 = (3.14159 + (3.14159 / float32((1.5 * ((((0.0 * 1e10) * -2.25) * 0.1) - 0.1)))))
                _ = fi54
                _ = v55
            }
        } else {
            v56 := GenS8{F0: v4}
            var v57 int16 = v5
            fmt.Println(v46, v46, v4)
            _ = v56
            _ = v57
        }
        if v53 {
            v6 = true
        } else {
            if (true && false) {
                if (456 <= 0) {
                    var v58 uint8 = v3
                    var v59 uint = 4294967295
                    _ = v58
                    _ = v59
                }
            } else {
                if true {
                    var v60 int16 = (v5 & v5)
                    var v61 string = (v4 + ("foo bar" + (v4 + v4)))
                    _ = v60
                    _ = v61
                } else {
                    if v53 {
                        v53 = (!v53)
                        for fi62 := 0; fi62 < 7; fi62++ {
                            _ = len(v47)
                            v63 := GenS2{F0: 42, F1: 42}
                            _ = fi62
                            _ = v63
                        }
                        var v64 uint8 = v3
                        _ = v64
                    } else {
                        _ = len(v47)
                        var v65 float64 = 0.1
                        _ = v65
                    }
                    v3 += uint8(v46)
                    if v53 {
                        v66 := GenS9{F0: (0.0 < (-2.25 * 1e10))}
                        _ = v66
                    } else {
                        if v6 {
                            var v67 uint8 = (v3 | (v3 & v3))
                            if false {
                                var v68 uint64 = 0
                                switch 32768 {
                                case 1:
                                    switch -2147483648 {
                                    case 807:
                                        if false {
                                            switch "a" {
                                            case "a":
                                                v5 += (v5 | v2)
                                            default:
                                                for fi69 := 0; fi69 < 20; fi69++ {
                                                    switch v4 {
                                                    case "":
                                                        if ((v67 <= 128) && (-49 > -1)) {
                                                            if v6 {
                                                                v5 = (v5 ^ 0)
                                                            } else {
                                                                _ = len(v47)
                                                            }
                                                            var v70 uint = v46
                                                            _ = v70
                                                        } else {
                                                            v46 -= 668
                                                            var v71 string = ("hello" + (v4 + "hello"))
                                                            _ = v71
                                                        }
                                                    case "hello":
                                                        var v72 int16 = v5
                                                        _ = v72
                                                    case "foo bar":
                                                        var v73 int8 = 11
                                                        _ = v73
                                                    default:
                                                        switch fi69 {
                                                        case 32767:
                                                            var v74 uint32 = 4294967295
                                                            _ = v74
                                                        default:
                                                            fmt.Println(v67)
                                                        }
                                                    }
                                                    switch fi69 {
                                                    case 128:
                                                        var v75 int16 = v5
                                                        _ = v75
                                                    case -2147483648:
                                                        delete(v1, (v4 + v4))
                                                    case 42:
                                                        v67 = 255
                                                    default:
                                                        var v76 uint = v46
                                                        _ = v76
                                                    }
                                                    v68 -= (^(v68))
                                                    _ = fi69
                                                }
                                            }
                                            switch 32768 {
                                            case 42:
                                                var v77 uint32 = 921
                                                _ = v77
                                            case 128:
                                                var v78 int8 = 42
                                                _ = v78
                                            default:
                                                switch 32767 {
                                                case -2147483648:
                                                    var v79 int32 = 0
                                                    _ = v79
                                                default:
                                                    v80 := v1["foo bar"]
                                                    _ = v80
                                                }
                                            }
                                            var v81 int16 = (v2 & (v5 * -1))
                                            _ = v81
                                        } else {
                                            _ = len(v47)
                                            v53 = v53
                                        }
                                    case 672:
                                        v82 := GenS5{F0: v46, F1: 42, F2: 172}
                                        _ = v82
                                    case 2147483647:
                                        var v83 int = 1
                                        _ = v83
                                    default:
                                        fmt.Println(v2, v3)
                                    }
                                default:
                                    var v84 int64 = 9223372036854775807
                                    _ = v84
                                }
                                _ = v68
                            }
                            switch 32767 {
                            case 4294967295:
                                v6 = true
                            case 1:
                                for fi85 := 0; fi85 < 4; fi85++ {
                                    for fi86 := 0; fi86 < 13; fi86++ {
                                        if v53 {
                                            v3 = 1
                                            if true {
                                                var v87 int16 = v2
                                                v5 = (-(v87))
                                                if (false && v53) {
                                                    v88 := GenS10{F0: 958, F1: 1, F2: "hello"}
                                                    _ = v88
                                                }
                                                v89 := GenS11{F0: 100, F1: 65535}
                                                var v90 int64 = 9223372036854775807
                                                _ = v87
                                                _ = v89
                                                _ = v90
                                            } else {
                                                var v91 bool = true
                                                v4 = (v4 + (v4 + v4))
                                                _ = v91
                                            }
                                            delete(v1, v4)
                                        }
                                        switch v4 {
                                        case "\n\t":
                                            v53 = false
                                        case "foo bar":
                                            var v92 int64 = genFn0(42, (v46 & (v46 >> ((v46) & 63))))
                                            _ = v92
                                        default:
                                            if (!(((!true) || true) && (61 <= 4294967295))) {
                                                var v93 uint = (v46 & (v46 / ((4294967295) | 1)))
                                                _ = v93
                                            }
                                        }
                                        _ = fi86
                                    }
                                    if false {
                                        if len(v47) > 0 {
                                            v47[0] = fi85
                                        }
                                    } else {
                                        var v94 uint16 = 2
                                        for fi95 := 0; fi95 < 13; fi95++ {
                                            v96 := GenS12{F0: -1}
                                            var v97 float64 = 3.14159
                                            v98 := GenS2{F0: -1, F1: 100}
                                            _ = fi95
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                        _ = v94
                                    }
                                    _ = fi85
                                }
                            case 32767:
                                v2 -= v5
                            default:
                                var v99 int8 = 22
                                _ = v99
                            }
                            _ = v67
                        }
                    }
                }
                for fi100 := 0; fi100 < 16; fi100++ {
                    v101 := GenS8{F0: v4}
                    _ = fi100
                    _ = v101
                }
                if v6 {
                    fmt.Println(v4, v4, v2)
                } else {
                    switch 32767 {
                    case -994:
                        v3 = 42
                    }
                    v2 -= 0
                }
            }
        }
        _ = v53
    }
    var v102 uint = (v46 / (((v46 / ((v46) | 1))) | 1))
    var v103 uint = (^(v46))
    var v104 int8 = -1
    var v105 uint8 = 115
    v106 := GenS13{F0: (((((v4 + v4) + v4) + v4) + (v4 + "a")) + ("\n\t" + v4)), F1: 42, F2: uint16(v105)}
    v102 -= (v103 & 794)
    var v107 int64 = 9223372036854775807
    v108 := GenS14{F0: v104, F1: (^(v105))}
    var v109 uint8 = v105
    var v110 int = -2147483648
    fmt.Println(v105)
    var v111 float32 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
