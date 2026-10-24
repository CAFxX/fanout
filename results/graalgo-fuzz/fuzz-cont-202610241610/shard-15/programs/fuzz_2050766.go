package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 uint32; F2 string }
type GenS4 struct { F0 int64; F1 int16; F2 uint }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 string; F2 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 int32 }
func genFn0(p0 int16, p1 uint8, p2 string) uint16 {
    v7 := GenS1{F0: 0, F1: 42}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 bool = (true && true)
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int64 = -9223372036854775808
    var v5 int32 = 42
    var v6 uint32 = 2
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v1 = (false && false)
            _ = fi10
        }
        var v11 uint16 = 65535
        _ = fi9
        _ = v11
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        if (32767 == -955) {
            switch ("\n\t" + (("" + ("a" + (("" + "\n\t") + ("hello" + "a")))) + "foo bar")) {
            case "\n\t":
                fmt.Println(v1, v5, v6)
            default:
                v4 += (v4 >> ((584) & 63))
            }
            v6 = (v6 - v6)
            v12 := GenS1{F0: genFn0(-468, 255, "hello"), F1: (v5 / (((-(v5))) | 1))}
            _ = v12
        }
    default:
        v4 = (-(v4))
    }
    var v13 float64 = 1.5
    v3[("" + "hello")] = 141
    if false {
        if true {
            v14 := v3["hello"]
            var v15 int64 = 42
            _ = v14
            _ = v15
        } else {
            var v16 int64 = (v4 / (((v4 - (-(v4)))) | 1))
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            switch (("foo bar" + ("" + "foo bar")) + "a") {
            case "foo bar":
                switch ("a" + "foo bar") {
                case "":
                    switch "\n\t" {
                    case "":
                        var v18 uint32 = (v6 - (v6 >> ((717) & 63)))
                        _ = v18
                    case "foo bar":
                        if (!v1) {
                            var v19 int32 = (v5 % (((v5 * v5)) | 1))
                            _ = v19
                        }
                    case "a":
                        v4 -= (v4 | (v4 * (v4 % (((v4 | (v4 & (-(v4))))) | 1))))
                    }
                default:
                    var v20 int64 = v4
                    _ = v20
                }
            case "hello":
                v5 = (v5 | v5)
            case "":
                if v1 {
                    v6 = (v6 - 116)
                    v21 := GenS1{F0: 764, F1: v5}
                    _ = v21
                }
            default:
                var v22 float64 = (0.0 / (v13 * (0.0 - 1.5)))
                _ = v22
            }
            _ = fi17
        }
    }
    var v23 uint8 = uint8(v6)
    switch 4294967295 {
    case -9223372036854775808:
        v24 := GenS2{F0: -175, F1: 0, F2: 127}
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS3{F0: (v23 | (^(v23))), F1: (v6 << ((65535) & 63)), F2: ""}
        if v1 {
            v27 := GenS1{F0: 483, F1: v5}
            _ = v27
        } else {
            v4 = 113
            var v28 uint64 = 2
            _ = v28
        }
        v29 := []int{fi25, (fi25 | (fi25 / ((546) | 1))), 4294967295}
        _ = fi25
        _ = v26
        _ = v29
    }
    switch "foo bar" {
    case "hello":
        v13 = (0.1 * v13)
    case "\n\t":
        if v1 {
            v1 = ((1 > 1) && ("\n\t" <= ("" + ("" + "hello"))))
            var v30 uint32 = v6
            _ = v30
        } else {
            delete(v2, -86)
            _ = len(v2)
            v31 := []int{-1}
            _ = v31
        }
    default:
        var v32 uint = 0
        _ = v32
    }
    v33 := v2[-56]
    var v34 uint32 = v6
    v1 = (65535 == 2)
    var v35 string = (("a" + v33) + (v33 + v33))
    for fi36 := 0; fi36 < 20; fi36++ {
        v23 = v23
        var v37 uint = 1
        if true {
            switch v35 {
            case "\n\t":
                var v38 int32 = (v5 % ((-636) | 1))
                _ = v38
            case "hello":
                v2[(fi36 + 259)] = v33
            case "a":
                v39 := GenS2{F0: v5, F1: 2, F2: fi36}
                _ = v39
            default:
                var v40 int8 = -91
                _ = v40
            }
            for fi41 := 0; fi41 < 4; fi41++ {
                switch v33 {
                case "a":
                    var v42 int8 = 24
                    _ = v42
                default:
                    var v43 int = -104
                    _ = v43
                }
                switch v35 {
                case "foo bar":
                    v44 := []int{fi41}
                    _ = v44
                }
                if (375 >= (v5 / (((v5 + 2147483647)) | 1))) {
                    var v45 string = ("a" + v33)
                    if false {
                        v13 = 1e10
                        var v46 bool = (241 <= 2)
                        _ = v46
                    }
                    fmt.Println(fi41, v45)
                    var v47 uint32 = 65
                    fmt.Println(v4)
                    if v1 {
                        fmt.Println(fi36, v47)
                        v48 := GenS4{F0: v4, F1: 948, F2: (v37 * (^(v37)))}
                        _ = v48
                    }
                    _ = v45
                    _ = v47
                } else {
                    if v1 {
                        v49 := GenS5{F0: (v4 << (((v37 >> (((v37 * (v37 - 255))) & 63))) & 63)), F1: (v6 ^ v34)}
                        if ((!v1) && false) {
                            if ((v4 ^ v4) >= (v4 % (((v4 & (-(v4)))) | 1))) {
                                v50 := GenS6{F0: v4, F1: (^(v23))}
                                var v51 string = "hello"
                                switch "hello" {
                                case "\n\t":
                                    var v52 float32 = 0.0
                                    _ = v52
                                case "foo bar":
                                    if (1 < 0) {
                                        var v53 uint64 = uint64(v37)
                                        fmt.Println(v6, fi36)
                                        v51 = ((v35 + (v33 + v51)) + v33)
                                        _ = v53
                                    } else {
                                        v54 := []int{127, fi41, (fi41 | fi41), (^(fi36))}
                                        var v55 uint16 = genFn0(32767, v23, "a")
                                        var v56 int32 = v5
                                        v37 -= v37
                                        if v1 {
                                            var v57 string = "hello"
                                            v55 -= v55
                                            _ = v57
                                        } else {
                                            if false {
                                                v58 := GenS3{F0: v23, F1: uint32(v37), F2: "hello"}
                                                v33 = v51
                                                _ = v58
                                            }
                                            fmt.Println(v56)
                                            if (!true) {
                                                v3[(v33 + "hello")] = (fi41 ^ (fi36 >> ((uint(v23)) & 63)))
                                            }
                                        }
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                case "a":
                                    fmt.Println(v33, v5, fi36)
                                default:
                                    v13 = (v13 + (v13 - v13))
                                }
                                _ = v50
                                _ = v51
                            } else {
                                fmt.Println(v23, v1)
                            }
                            switch "a" {
                            case "foo bar":
                                if false {
                                    var v59 string = ""
                                    var v60 uint = (v37 >> ((65535) & 63))
                                    v35 = (v35 + v35)
                                    _ = v59
                                    _ = v60
                                }
                            case "a":
                                if (!(!false)) {
                                    fmt.Println(fi41, v33)
                                    v61 := GenS7{F0: -875, F1: v5}
                                    _ = v61
                                }
                            case "hello":
                                switch v33 {
                                case "":
                                    v4 -= (v4 << ((v37) & 63))
                                case "a":
                                    fmt.Println(v13)
                                case "\n\t":
                                    delete(v3, "a")
                                default:
                                    switch fi36 {
                                    case -747:
                                        if false {
                                            v5 = (v5 / (((v5 << ((v37) & 63))) | 1))
                                        } else {
                                            v62 := GenS1{F0: 0, F1: (v5 | int32(v4))}
                                            _ = v62
                                        }
                                    case 0:
                                        var v63 uint = (v37 >> ((uint(v23)) & 63))
                                        _ = v63
                                    default:
                                        fmt.Println(v33, fi36, v5)
                                    }
                                    v64 := GenS8{F0: 127}
                                    v65 := v2[fi41]
                                    _ = v64
                                    _ = v65
                                }
                            }
                            if v1 {
                                var v66 int32 = (v5 ^ 0)
                                _ = v66
                            } else {
                                v23 = (v23 ^ 182)
                                if (v4 < (v4 >> (((v37 & v37)) & 63))) {
                                    var v67 uint = (v37 | (v37 & (v37 - v37)))
                                    var v68 uint32 = v6
                                    switch "a" {
                                    case "\n\t":
                                        if v1 {
                                            var v69 uint = 65535
                                            v33 = (v35 + v35)
                                            _ = v69
                                        } else {
                                            v6 -= (v6 & 2147483647)
                                            var v70 uint32 = 2
                                            v71 := GenS9{F0: (!((fi41 % ((fi36) | 1)) < fi41)), F1: ("foo bar" + v35), F2: v4}
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                    var v72 float32 = float32(1e10)
                                    _ = v67
                                    _ = v68
                                    _ = v72
                                } else {
                                    fmt.Println(v37, v33, v4)
                                    var v73 uint64 = 2
                                    var v74 int = fi41
                                    _ = v73
                                    _ = v74
                                }
                                var v75 int64 = (v4 << ((255) & 63))
                                _ = v75
                            }
                            var v76 uint64 = 570
                            _ = v76
                        }
                        _ = v49
                    }
                    var v77 string = (((v33 + v35) + "hello") + "foo bar")
                    _ = v77
                }
                _ = fi41
            }
            v78 := []int{fi36, int(v37)}
            _ = v78
        }
        _ = fi36
        _ = v37
    }
    switch "hello" {
    case "foo bar":
        switch "a" {
        case "hello":
            v13 = (((float64(-2.25) - v13) * v13) + (v13 * (3.14159 + v13)))
        case "a":
            switch 42 {
            case 42:
                if v1 {
                    var v79 string = (v35 + ("hello" + ("\n\t" + v33)))
                    _ = v79
                }
            case 2147483647:
                for fi80 := 0; fi80 < 18; fi80++ {
                    if true {
                        if (((0.0 * 1.5) - 0.1) >= 1e10) {
                            v23 = (^(v23))
                            v81 := GenS8{F0: -128}
                            _ = v81
                        } else {
                            v82 := GenS10{F0: 401}
                            _ = v82
                        }
                        var v83 uint8 = v23
                        _ = v83
                    } else {
                        var v84 int64 = int64(v5)
                        switch (fi80 ^ fi80) {
                        case -355:
                            v13 = v13
                        case 9223372036854775807:
                            v85 := GenS1{F0: 255, F1: 2147483647}
                            _ = v85
                        case -9223372036854775808:
                            if true {
                                switch fi80 {
                                case 0:
                                    for fi86 := 0; fi86 < 2; fi86++ {
                                        v87 := GenS11{F0: v23, F1: (v5 << ((0) & 63))}
                                        _ = fi86
                                        _ = v87
                                    }
                                default:
                                    v35 = v33
                                }
                                switch (fi80 | 4294967295) {
                                case 4294967295:
                                    if false {
                                        var v88 int32 = (v5 << ((77) & 63))
                                        _ = v88
                                    } else {
                                        for fi89 := 0; fi89 < 19; fi89++ {
                                            var v90 uint64 = 2
                                            _ = fi89
                                            _ = v90
                                        }
                                        v91 := GenS4{F0: (v84 & v84), F1: 42, F2: 556}
                                        _ = v91
                                    }
                                case -9223372036854775808:
                                    var v92 int = 127
                                    _ = v92
                                case 0:
                                    v93 := []int{(^(fi80)), 42, fi80, fi80}
                                    _ = v93
                                default:
                                    v94 := []int{(fi80 / ((fi80) | 1)), (-(fi80)), (fi80 << ((252) & 63))}
                                    _ = v94
                                }
                                fmt.Println(v33, v34)
                            }
                        default:
                            v84 = -1
                        }
                        _ = v84
                    }
                    switch fi80 {
                    case 32767:
                        var v95 int16 = 436
                        _ = v95
                    }
                    _ = fi80
                }
            case 375:
                for fi96 := 0; fi96 < 5; fi96++ {
                    if (42 == 255) {
                        fmt.Println(v13)
                    } else {
                        switch ("\n\t" + "a") {
                        case "a":
                            if true {
                                var v97 int8 = 4
                                if (1.5 < 0.0) {
                                    var v98 float32 = ((float32(float64(1e10)) / 1.5) * 0.0)
                                    var v99 uint64 = 0
                                    fmt.Println(v6)
                                    _ = v98
                                    _ = v99
                                }
                                _ = v97
                            } else {
                                v100 := []int{}
                                v4 += v4
                                v101 := GenS3{F0: 127, F1: v34, F2: "\n\t"}
                                _ = v100
                                _ = v101
                            }
                        case "\n\t":
                            v13 -= 1e10
                        }
                        fmt.Println(v5, v23)
                    }
                    _ = fi96
                }
            }
        case "\n\t":
            switch 772 {
            case 32768:
                v34 -= (v6 << ((65535) & 63))
            case 4294967295:
                var v102 float64 = (0.1 / v13)
                _ = v102
            }
        default:
            v5 -= v5
        }
    default:
        v13 -= v13
    }
    v103 := []int{-1, 127, 0, 1}
    fmt.Println(v23, v5)
    v2[2147483647] = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v103:", len(v103), v103)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
