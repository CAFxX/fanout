package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 string; F1 int; F2 uint }
type GenS6 struct { F0 int16; F1 uint32; F2 int32 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 string) string {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 128, F1: true}
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint32 = 4294967295
            var v11 int64 = 9223372036854775807
            _ = fi9
            _ = v10
            _ = v11
        }
    case "foo bar":
        var v12 float64 = 3.14159
        _ = v12
    case "a":
        if (521 < 42) {
            p0 = p0
            if true {
                if (255 < 127) {
                    v13 := GenS1{F0: 127, F1: (true || true)}
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int8 = 42
                        v16 := []int{2147483647, 579}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                } else {
                    var v17 int = 128
                    p0 = ""
                    _ = v17
                }
                if false {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if true {
                            var v19 uint32 = 305
                            _ = v19
                        }
                        var v20 uint = 42
                        _ = fi18
                        _ = v20
                    }
                    p0 = p0
                } else {
                    v21 := []int{-344, 32767, 128}
                    _ = v21
                }
                var v22 bool = false
                _ = v22
            } else {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS2{F0: -776, F1: ((-1 != -1) && (0.0 != (3.14159 / 0.1)))}
                    v25 := GenS3{F0: 127}
                    p0 = p0
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        } else {
            p0 = p0
            var v26 float64 = 1.5
            fmt.Println(p0)
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := []int{fi27}
        v29 := GenS3{F0: 174}
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 uint = 2
            p0 = (p0 + "\n\t")
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = p0
    return p0
}
func genFn1() uint64 {
    var v32 uint32 = 2
    switch 127 {
    case -351:
        v33 := GenS4{F0: -1, F1: 42}
        _ = v33
    case 32767:
        var v34 uint16 = 65535
        _ = v34
    default:
        var v35 int64 = -1
        _ = v35
    }
    _ = v32
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []int{600}
    v4 := []string{("hello" + ("" + "\n\t")), ("hello" + "\n\t"), "hello"}
    var v5 int8 = 127
    var v6 uint16 = 230
    v36 := GenS3{F0: 42}
    if (685 != 1000) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 -= (v6 / (((v6 ^ (v6 ^ (v6 >> ((1) & 63))))) | 1))
        if (42 >= 18446744073709551615) {
            if (!(!(!((4294967295 >= 344) && (!(9223372036854775807 == 133)))))) {
                var v37 uint16 = (v6 >> ((540) & 63))
                v38 := GenS2{F0: -354, F1: false}
                _ = v37
                _ = v38
            }
            fmt.Println(v6)
            var v39 bool = false
            _ = v39
        } else {
            v6 -= (v6 / ((v6) | 1))
            var v40 int16 = -1
            var v41 float64 = 0.1
            _ = v40
            _ = v41
        }
    }
    var v42 int = 128
    if true {
        if (!(2147483648 != 42)) {
            v43 := GenS1{F0: uint8(v6), F1: true}
            var v44 int8 = 42
            if false {
                v44 = 71
                v45 := GenS5{F0: "hello", F1: int(v6), F2: 0}
                _ = v45
            } else {
                for fi46 := 0; fi46 < 20; fi46++ {
                    var v47 int8 = v44
                    _ = fi46
                    _ = v47
                }
                switch "hello" {
                case "\n\t":
                    fmt.Println(v44, v5, v42)
                case "":
                    for fi48 := 0; fi48 < 1; fi48++ {
                        var v49 int64 = -5
                        v44 = v5
                        _ = fi48
                        _ = v49
                    }
                }
            }
            _ = v43
            _ = v44
        } else {
            v42 -= (v42 << ((65535) & 63))
            switch 32768 {
            case 937:
                for fi50 := 0; fi50 < 12; fi50++ {
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 bool = (!(1 == -1))
                        _ = fi51
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 9; fi53++ {
                        fmt.Println(fi50)
                        if false {
                            var v54 int = fi50
                            _ = v54
                        }
                        _ = fi53
                    }
                    _ = fi50
                }
            case -2147483648:
                var v55 uint32 = 42
                _ = v55
            default:
                for fi56 := 0; fi56 < 1; fi56++ {
                    var v57 int32 = 0
                    _ = fi56
                    _ = v57
                }
            }
            v42 += v42
        }
        for fi58 := 0; fi58 < 15; fi58++ {
            _ = len(v4)
            v59 := GenS6{F0: 1000, F1: 4294967295, F2: 1}
            _ = fi58
            _ = v59
        }
    } else {
        v6 = v6
        for fi60 := 0; fi60 < 20; fi60++ {
            v42 = v42
            if true {
                fmt.Println(v6, v6, v42)
            }
            _ = fi60
        }
    }
    switch "hello" {
    case "hello":
        for fi61 := 0; fi61 < 11; fi61++ {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v62 := v1[527]
                v63 := GenS3{F0: 255}
                _ = v62
                _ = v63
            }
            _ = fi61
        }
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        for fi64 := 0; fi64 < 13; fi64++ {
            if true {
                for fi65 := 0; fi65 < 12; fi65++ {
                    var v66 int16 = 1
                    _ = fi65
                    _ = v66
                }
                var v67 uint64 = 4294967295
                v68 := GenS4{F0: 2147483647, F1: 2}
                _ = v67
                _ = v68
            }
            for fi69 := 0; fi69 < 14; fi69++ {
                switch "" {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    v42 = (fi69 & fi69)
                    fmt.Println(v5, v42, v42)
                case "\n\t":
                    v70 := GenS4{F0: -361, F1: (v6 + v6)}
                    _ = v70
                default:
                    _ = len(v4)
                }
                _ = fi69
            }
            _ = fi64
        }
    }
    var v71 float32 = 0.0
    fmt.Println(v71, v6, v42)
    switch v42 {
    case -2147483648:
        fmt.Println(v5)
    default:
        var v72 int = v42
        _ = v72
    }
    if (!((0.0 - 3.14159) <= 1e10)) {
        v6 = (^(v6))
        for fi73 := 0; fi73 < 13; fi73++ {
            var v74 uint32 = 2147483647
            _ = fi73
            _ = v74
        }
    } else {
        if true {
            v71 += -2.25
        } else {
            var v75 string = "\n\t"
            v42 = v42
            var v76 float32 = 1.5
            _ = v75
            _ = v76
        }
        v5 += (v5 ^ (^(v5)))
        if false {
            switch 2147483647 {
            case 32767:
                if ((!(((v6 != (v6 / ((v6) | 1))) || (v5 >= (v5 % (((-(v5))) | 1)))) || false)) || false) {
                    var v77 uint16 = v6
                    fmt.Println(v5, v42)
                    v78 := GenS2{F0: 1, F1: (1 != -1)}
                    _ = v77
                    _ = v78
                }
            case -915:
                if false {
                    var v79 string = "a"
                    var v80 int16 = 0
                    var v81 string = v79
                    _ = v79
                    _ = v80
                    _ = v81
                } else {
                    v6 = (v6 / (((v6 ^ v6)) | 1))
                    switch (v42 - (v42 << ((2) & 63))) {
                    case 9223372036854775807:
                        fmt.Println(v6)
                    default:
                        if false {
                            for fi82 := 0; fi82 < 15; fi82++ {
                                v71 = ((((v71 * ((v71 * float32(0.1)) * 0.0)) / v71) / v71) * v71)
                                v83 := GenS7{F0: 431, F1: 966}
                                v6 += (^(v6))
                                _ = fi82
                                _ = v83
                            }
                            v84 := GenS2{F0: -379, F1: false}
                            v71 -= (0.0 - -2.25)
                            _ = v84
                        }
                    }
                    v85 := GenS2{F0: -1, F1: (((v71 / ((float32(float64(v71)) - 0.0) + (float32(1.5) * v71))) / (3.14159 * -2.25)) <= v71)}
                    _ = v85
                }
            }
        } else {
            v5 += v5
        }
    }
    switch v42 {
    case 2147483647:
        v6 += 1
    case 42:
        var v86 uint = 65535
        _ = v86
    default:
        var v87 uint8 = 127
        _ = v87
    }
    for fi88 := 0; fi88 < 7; fi88++ {
        for fi89 := 0; fi89 < 11; fi89++ {
            if (((2147483647 > (v42 & (v42 & (fi89 * (fi88 >> ((989) & 63)))))) && (1 == 255)) && true) {
                if (765 <= 4294967295) {
                    v4 = append(v4, "hello")
                    v42 += fi88
                    if (!true) {
                        if false {
                            var v90 string = "\n\t"
                            v5 += -1
                            var v91 int8 = v5
                            v92 := GenS2{F0: -1, F1: (!false)}
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            _ = v90
                            _ = v91
                            _ = v92
                        } else {
                            v93 := GenS8{F0: 2}
                            _ = v93
                        }
                        if (!((3.14159 + v71) < (float32(0.0) / 1.0))) {
                            var v94 string = ("" + "foo bar")
                            if false {
                                switch fi88 {
                                case 289:
                                    v5 = (v5 / ((1) | 1))
                                default:
                                    fmt.Println(v42, fi89, v42)
                                }
                            }
                            var v95 string = v94
                            _ = v94
                            _ = v95
                        }
                        var v96 bool = true
                        _ = v96
                    } else {
                        fmt.Println(fi88)
                    }
                }
            }
            var v97 uint64 = 4294967295
            _ = len(v3)
            v98 := GenS9{F0: v6}
            fmt.Println(fi88, v6)
            _ = fi89
            _ = v97
            _ = v98
        }
        v6 -= 42
        _ = fi88
    }
    if false {
        if (false || true) {
            v99 := GenS2{F0: 795, F1: (-351 <= 32767)}
            v100 := GenS9{F0: v6}
            _ = v100
            _ = v99
        }
        var v101 float64 = 0.0
        _ = v101
    }
    var v102 int64 = 42
    var v103 int8 = (v5 - v5)
    delete(v1, 65535)
    switch "a" {
    case "hello":
        v103 -= v5
    case "foo bar":
        if (!false) {
            fmt.Println(v6)
            if (v6 != 65535) {
                var v104 uint8 = 0
                fmt.Println(v104)
                v105 := GenS5{F0: "a", F1: int(v5), F2: 65535}
                _ = v104
                _ = v105
            } else {
                v106 := v2[42]
                for fi107 := 0; fi107 < 10; fi107++ {
                    var v108 int8 = (v5 >> ((0) & 63))
                    var v109 float32 = v71
                    if (v6 < v6) {
                        var v110 uint8 = 0
                        var v111 int16 = 1000
                        var v112 uint = 65535
                        _ = v110
                        _ = v111
                        _ = v112
                    }
                    _ = fi107
                    _ = v108
                    _ = v109
                }
                v6 += (v6 * 1)
                _ = v106
            }
        } else {
            v113 := v2[864]
            _ = v113
        }
    default:
        v114 := GenS10{F0: v6}
        _ = v114
    }
    v71 -= v71
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v71
}
