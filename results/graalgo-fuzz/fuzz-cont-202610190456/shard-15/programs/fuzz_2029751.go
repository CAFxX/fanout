package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int8; F1 int8; F2 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 string; F2 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64; F1 int; F2 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 uint16) string {
    v6 := GenS1{F0: (p0 % ((p0) | 1))}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return "a"
}
func genFn1() uint64 {
    var v7 string = "\n\t"
    v8 := GenS1{F0: 363}
    if ((1.5 >= (1e10 / 0.1)) && true) {
        v7 = v7
        for fi9 := 0; fi9 < 7; fi9++ {
            v7 = v7
            v7 = ("\n\t" + "\n\t")
            _ = fi9
        }
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            if (!(!true)) {
                if (0.1 > 3.14159) {
                    var v11 string = (v7 + v7)
                    var v12 uint8 = 42
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS2{F0: 2, F1: 110}
                    v14 := GenS3{F0: 65535}
                    var v15 uint32 = 65535
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if (844 > 42) {
                    fmt.Println(fi10)
                } else {
                    if true {
                        if true {
                            if true {
                                if ("a" == ("\n\t" + (("a" + v7) + v7))) {
                                    if true {
                                        for fi16 := 0; fi16 < 17; fi16++ {
                                            var v17 int32 = -1
                                            v17 = v17
                                            fmt.Println(fi16, v7)
                                            _ = fi16
                                            _ = v17
                                        }
                                        switch fi10 {
                                        case 42:
                                            v7 = "hello"
                                        case 208:
                                            switch (v7 + "foo bar") {
                                            case "foo bar":
                                                v7 = v7
                                            }
                                        case 32767:
                                            for fi18 := 0; fi18 < 11; fi18++ {
                                                v7 = "\n\t"
                                                v7 = "\n\t"
                                                var v19 string = "\n\t"
                                                _ = fi18
                                                _ = v19
                                            }
                                        default:
                                            switch fi10 {
                                            case -2147483648:
                                                v7 = v7
                                            }
                                        }
                                    } else {
                                        v20 := GenS4{F0: 0, F1: 36}
                                        _ = v20
                                    }
                                    if (false && (127 < 0)) {
                                        v7 = ("\n\t" + v7)
                                    }
                                    if false {
                                        for fi21 := 0; fi21 < 18; fi21++ {
                                            var v22 uint64 = 1
                                            var v23 int16 = -632
                                            fmt.Println(fi21, v7, fi21)
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                }
                                v7 = v7
                                var v24 string = ""
                                _ = v24
                            } else {
                                if false {
                                    if true {
                                        for fi25 := 0; fi25 < 17; fi25++ {
                                            v7 = (("hello" + "hello") + "\n\t")
                                            v7 = ("foo bar" + (("foo bar" + "hello") + "hello"))
                                            v7 = (("foo bar" + v7) + v7)
                                            var v26 int = 127
                                            v27 := GenS5{F0: 110, F1: 78, F2: int8(fi10)}
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                        v7 = ""
                                        v28 := GenS6{F0: false}
                                        _ = v28
                                    } else {
                                        var v29 string = ""
                                        _ = v29
                                    }
                                    v30 := GenS3{F0: 137}
                                    if ((!false) || (!(!false))) {
                                        v7 = ((v7 + "\n\t") + (v7 + ""))
                                    } else {
                                        if (3.14159 < 1.5) {
                                            for fi31 := 0; fi31 < 3; fi31++ {
                                                var v32 int8 = -1
                                                v33 := GenS7{F0: -455, F1: ("foo bar" + v7), F2: 42}
                                                switch "a" {
                                                case "a":
                                                    v34 := GenS3{F0: 65535}
                                                    _ = v34
                                                case "\n\t":
                                                    v35 := GenS8{F0: false}
                                                    _ = v35
                                                default:
                                                    switch 128 {
                                                    case 32767:
                                                        fmt.Println(fi10, v7, fi31)
                                                    case 545:
                                                        if false {
                                                            var v36 float32 = 0.0
                                                            var v37 int = 42
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            switch v7 {
                                                            case "hello":
                                                                v32 += v32
                                                            }
                                                            v38 := []int{fi10, (fi31 & fi31), (fi10 << ((42) & 63)), (fi10 * -706)}
                                                            v32 = (v32 | 90)
                                                            _ = v38
                                                        }
                                                    case 9223372036854775807:
                                                        var v39 int = (fi31 << ((1) & 63))
                                                        _ = v39
                                                    }
                                                }
                                                v32 = (v32 << ((477) & 63))
                                                var v40 uint = 65535
                                                v41 := GenS9{F0: 1}
                                                if (0.1 <= float64(0.1)) {
                                                    v40 += (v40 >> (((v40 - v40)) & 63))
                                                    v7 = (v7 + v7)
                                                }
                                                _ = fi31
                                                _ = v32
                                                _ = v33
                                                _ = v40
                                                _ = v41
                                            }
                                        } else {
                                            v42 := []int{32768}
                                            _ = v42
                                        }
                                        fmt.Println(v7, v7)
                                    }
                                    _ = v30
                                } else {
                                    var v43 uint8 = 163
                                    v44 := []int{fi10, 32767, 961, (fi10 + fi10)}
                                    _ = v43
                                    _ = v44
                                }
                            }
                            var v45 int16 = -1
                            _ = v45
                        } else {
                            v46 := []int{(fi10 | (fi10 & -591)), (fi10 + fi10), (^(fi10))}
                            if (!true) {
                                switch v7 {
                                case "a":
                                    var v47 uint8 = 0
                                    _ = v47
                                case "\n\t":
                                    v7 = v7
                                case "":
                                    var v48 int32 = 42
                                    _ = v48
                                default:
                                    switch ((v7 + v7) + (v7 + "")) {
                                    case "\n\t":
                                        _ = len(v46)
                                    case "":
                                        switch 9223372036854775807 {
                                        case -690:
                                            switch 247 {
                                            case -1:
                                                for fi49 := 0; fi49 < 3; fi49++ {
                                                    var v50 bool = false
                                                    fmt.Println(v50, fi49, fi49)
                                                    if true {
                                                        switch v7 {
                                                        case "hello":
                                                            switch -1 {
                                                            case 42:
                                                                if false {
                                                                    v51 := GenS10{F0: 469, F1: (fi49 >> ((0) & 63)), F2: 1}
                                                                    if len(v46) > 0 {
                                                                        v46[0] = (fi10 - fi49)
                                                                    }
                                                                    v52 := GenS9{F0: 2147483647}
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    if len(v46) > 0 {
                                                                        v46[0] = -683
                                                                    }
                                                                    var v53 bool = (!(v50 && v50))
                                                                    _ = v53
                                                                }
                                                                switch (fi49 - fi49) {
                                                                case 161:
                                                                    if len(v46) >= 2 {
                                                                        v46 = v46[1:len(v46)]
                                                                    }
                                                                default:
                                                                    v50 = (!(!(2147483647 > 2147483647)))
                                                                }
                                                            case 32767:
                                                                var v54 int16 = 1
                                                                _ = v54
                                                            default:
                                                                v55 := GenS9{F0: -1}
                                                                _ = v55
                                                            }
                                                        case "":
                                                            fmt.Println(v7, fi10)
                                                        case "\n\t":
                                                            var v56 uint = 42
                                                            _ = v56
                                                        }
                                                        var v57 float32 = 0.0
                                                        if (-149 >= -32768) {
                                                            var v58 int8 = 86
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    }
                                                    _ = fi49
                                                    _ = v50
                                                }
                                            case 0:
                                                v59 := GenS11{F0: 42}
                                                _ = v59
                                            default:
                                                v7 = "foo bar"
                                            }
                                        case 42:
                                            v7 = v7
                                        case -343:
                                            v7 = v7
                                        }
                                    }
                                }
                                var v60 int = fi10
                                _ = v60
                            } else {
                                if len(v46) > 0 {
                                    v46[0] = -768
                                }
                                fmt.Println(v7, v7)
                                fmt.Println(v7, fi10)
                            }
                            v61 := GenS1{F0: 255}
                            _ = v46
                            _ = v61
                        }
                        v7 = "a"
                        v62 := GenS2{F0: 1, F1: 980}
                        _ = v62
                    }
                }
                switch v7 {
                case "":
                    if false {
                        v7 = ((((v7 + (v7 + v7)) + v7) + "hello") + "")
                        v63 := []int{}
                        if false {
                            if len(v63) > 0 {
                                v63[0] = (fi10 << ((1) & 63))
                            }
                        } else {
                            for fi64 := 0; fi64 < 6; fi64++ {
                                fmt.Println(fi10, fi64, v7)
                                _ = fi64
                            }
                            if (0 == 1) {
                                v7 = v7
                                var v65 bool = (255 <= 42)
                                _ = len(v63)
                                _ = v65
                            }
                        }
                        _ = v63
                    } else {
                        v7 = "\n\t"
                        switch "foo bar" {
                        case "\n\t":
                            v66 := []int{fi10}
                            _ = v66
                        default:
                            v7 = ""
                        }
                    }
                default:
                    v7 = v7
                }
            } else {
                v7 = ((v7 + "") + "")
                v67 := GenS12{F0: 0, F1: 2}
                _ = v67
            }
            _ = fi10
        }
        if false {
            switch v7 {
            case "":
                var v68 int16 = 32767
                _ = v68
            }
            var v69 string = v7
            v69 = (v69 + v69)
            _ = v69
        } else {
            var v70 uint32 = 231
            if true {
                if ((1.5 / -2.25) == 0.0) {
                    v71 := GenS7{F0: 0, F1: v7, F2: 42}
                    _ = v71
                } else {
                    var v72 int64 = -9223372036854775808
                    _ = v72
                }
                switch 32767 {
                case -849:
                    fmt.Println(v7)
                default:
                    v7 = v7
                }
                if false {
                    if (3.14159 > ((float32(1.5) + -2.25) - 1.5)) {
                        if (!(!(true && false))) {
                            var v73 int64 = -170
                            for fi74 := 0; fi74 < 8; fi74++ {
                                for fi75 := 0; fi75 < 11; fi75++ {
                                    if (!(true || true)) {
                                        v76 := GenS13{F0: 42, F1: 18446744073709551615, F2: 65535}
                                        _ = v76
                                    } else {
                                        v77 := []int{fi74, 118}
                                        v78 := GenS10{F0: 42, F1: (fi75 - (fi74 << ((434) & 63))), F2: 42}
                                        _ = v77
                                        _ = v78
                                    }
                                    switch fi75 {
                                    case -285:
                                        v7 = v7
                                        v79 := GenS13{F0: 242, F1: 2, F2: 384}
                                        _ = v79
                                    case -1:
                                        v80 := []int{(fi74 + fi75)}
                                        _ = v80
                                    case -708:
                                        v81 := []int{fi75, (^(fi75)), (fi75 + (-(fi74)))}
                                        _ = v81
                                    }
                                    v7 = ("foo bar" + "a")
                                    _ = fi75
                                }
                                v82 := GenS8{F0: true}
                                _ = fi74
                                _ = v82
                            }
                            fmt.Println(v7)
                            _ = v73
                        } else {
                            var v83 int64 = 173
                            v83 += (v83 ^ 42)
                            _ = v83
                        }
                    } else {
                        v84 := []int{127}
                        v84 = append(v84, 127)
                        _ = v84
                    }
                } else {
                    var v85 int64 = 9223372036854775807
                    v85 += v85
                    _ = v85
                }
            }
            _ = v70
        }
        var v86 float64 = -2.25
        _ = v86
    }
    var v87 int64 = -51
    _ = v7
    _ = v8
    _ = v87
    return 111
}

func main() {
    v1 := []int{-9223372036854775808, 32768, 128, 402}
    var v2 int = 0
    var v3 float32 = ((-2.25 / 3.14159) * float32(0.0))
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    if len(v1) > 0 {
        v1[0] = v2
    }
    for fi88 := 0; fi88 < 16; fi88++ {
        v3 = v3
        v2 -= 32767
        var v89 uint64 = 4294967295
        _ = fi88
        _ = v89
    }
    var v90 float64 = float64(0.0)
    if false {
        var v91 string = "a"
        switch 999 {
        case 9223372036854775807:
            if (!false) {
                switch 32767 {
                case 128:
                    var v92 int64 = 42
                    _ = v92
                case 32767:
                    var v93 string = "hello"
                    _ = v93
                }
            }
        case 4294967295:
            v90 = v90
        default:
            var v94 bool = (v91 != (v91 + "a"))
            _ = v94
        }
        _ = v91
    }
    fmt.Println(v2, v90)
    if (888 <= 2) {
        for fi95 := 0; fi95 < 5; fi95++ {
            v3 = v3
            for fi96 := 0; fi96 < 12; fi96++ {
                v97 := GenS13{F0: 42, F1: 76, F2: 42}
                fmt.Println(fi95)
                var v98 float32 = (3.14159 - v3)
                switch v2 {
                case -1:
                    var v99 uint8 = 1
                    _ = v99
                case 128:
                    var v100 int8 = 42
                    _ = v100
                case 173:
                    var v101 int16 = 624
                    _ = v101
                }
                _ = fi96
                _ = v97
                _ = v98
            }
            _ = fi95
        }
        v90 -= float64(-2.25)
    }
    if len(v1) > 0 {
        v1[0] = (v2 + 9223372036854775807)
    }
    var v102 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v90:", v90)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
