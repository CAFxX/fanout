package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint16, p1 int, p2 uint8) bool {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: 1, F1: -1}
        _ = fi4
        _ = v5
    }
    var v6 float64 = 1.5
    var v7 float64 = ((float64(1.5) * 3.14159) / float64(-2.25))
    switch ("foo bar" + (("\n\t" + "\n\t") + (("foo bar" + "a") + "hello"))) {
    case "a":
        v8 := GenS1{F0: 4294967295, F1: 42}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 19; fi9++ {
            v6 -= (v6 * v6)
            v10 := GenS1{F0: 42, F1: 1}
            _ = fi9
            _ = v10
        }
    case "hello":
        p2 -= p2
    default:
        v11 := []int{(p1 << ((2) & 63))}
        _ = v11
    }
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int, p1 uint32) int8 {
    fmt.Println(p1, p1)
    if (1 == 0) {
        var v12 uint64 = 1
        v13 := []int{p0, 128}
        _ = v12
        _ = v13
    }
    if true {
        var v14 int16 = -32768
        switch "foo bar" {
        case "":
            switch p0 {
            case 0:
                v15 := []int{p0, (p0 ^ (p0 - p0))}
                _ = v15
            case 2147483647:
                v14 -= (v14 - -1)
            case -9223372036854775808:
                fmt.Println(p0, v14, p1)
            }
        }
        fmt.Println(p1, p1)
        _ = v14
    }
    p1 += (p1 | p1)
    _ = p0
    _ = p1
    return 76
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = 100
    var v3 bool = true
    v1 -= v1
    v16 := []int{}
    switch -2147483648 {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 7; fi17++ {
            if len(v16) > 0 {
                v16[0] = 4294967295
            }
            _ = fi17
        }
    default:
        var v18 int = -2147483648
        _ = v18
    }
    v19 := GenS2{F0: 1, F1: (v1 * (v1 * 1)), F2: 1}
    if v3 {
        v1 = v1
        switch "hello" {
        case "a":
            if true {
                if v3 {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            v20 := GenS2{F0: 2147483647, F1: 65535, F2: 42}
                            v16 = append(v16, -9223372036854775808)
                            _ = v20
                        } else {
                            var v21 uint8 = 128
                            _ = v21
                        }
                    }
                    var v22 int32 = 42
                    v2 = genFn1(276, 2)
                    _ = v22
                } else {
                    v2 -= v2
                    if (1 > 255) {
                        for fi23 := 0; fi23 < 2; fi23++ {
                            v2 = v2
                            switch ("a" + ("hello" + "a")) {
                            case "":
                                var v24 uint16 = v1
                                _ = v24
                            case "\n\t":
                                if len(v16) > 0 {
                                    v16[0] = (fi23 / (((fi23 | (fi23 - 42))) | 1))
                                }
                            default:
                                var v25 float64 = (0.1 / (float64(1e10) + (((0.0 * -2.25) * (3.14159 * float64(-2.25))) / 1.0)))
                                _ = v25
                            }
                            _ = fi23
                        }
                        var v26 uint16 = (v1 + (v1 ^ v1))
                        _ = v26
                    } else {
                        v3 = v3
                    }
                }
                fmt.Println(v2)
                fmt.Println(v3, v2)
            }
        case "foo bar":
            for fi27 := 0; fi27 < 11; fi27++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi27
            }
        case "\n\t":
            v28 := GenS3{F0: 1, F1: "", F2: ("hello" + ("" + ("" + "hello")))}
            _ = v28
        default:
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 float64 = (float64(0.1) * (((3.14159 - (0.1 - 0.0)) + 0.1) * 0.0))
                v31 := GenS1{F0: 1, F1: -151}
                for fi32 := 0; fi32 < 18; fi32++ {
                    v1 += (v1 / ((v1) | 1))
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v3, v30, v2)
                    default:
                        switch ("foo bar" + ("foo bar" + "\n\t")) {
                        case "":
                            switch ("\n\t" + ("hello" + "\n\t")) {
                            case "\n\t":
                                v2 += v2
                                if v3 {
                                    var v33 int = fi29
                                    var v34 int32 = -794
                                    var v35 int32 = v34
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    _ = len(v16)
                                    fmt.Println(v30, v3)
                                    var v36 int8 = v2
                                    _ = v36
                                }
                            default:
                                var v37 uint8 = 2
                                _ = v37
                            }
                        case "a":
                            if v3 {
                                v30 += v30
                            } else {
                                v3 = v3
                                if (127 >= 100) {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    v38 := GenS1{F0: 59, F1: 0}
                                    _ = v38
                                }
                                fmt.Println(v2, v1)
                            }
                        case "foo bar":
                            v16 = append(v16, fi32)
                        default:
                            v2 = 100
                        }
                    }
                    v30 = (((1.5 / 3.14159) - v30) / v30)
                    _ = fi32
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    var v39 int8 = v2
    var v40 uint8 = 128
    if v3 {
        v3 = v3
    }
    var v41 float64 = (1e10 / 0.1)
    var v42 float64 = v41
    v3 = v3
    for fi43 := 0; fi43 < 5; fi43++ {
        switch fi43 {
        case 0:
            v44 := GenS1{F0: 65535, F1: 0}
            _ = v44
        default:
            _ = len(v16)
        }
        for fi45 := 0; fi45 < 18; fi45++ {
            v16 = append(v16, (fi45 - fi43))
            if v3 {
                v39 = v2
                if v3 {
                    if v3 {
                        var v46 float32 = 1.5
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        switch (fi43 >> ((4294967295) & 63)) {
                        case 1:
                            if (("foo bar" + ("foo bar" + "foo bar")) >= "foo bar") {
                                var v47 uint32 = 2
                                _ = v47
                            } else {
                                v16 = append(v16, fi45)
                                if (1 == 108) {
                                    fmt.Println(v40, v39, v42)
                                } else {
                                    var v48 uint8 = (v40 * v40)
                                    fmt.Println(v40)
                                    switch fi45 {
                                    case 0:
                                        switch ("hello" + "a") {
                                        case "foo bar":
                                            _ = len(v16)
                                            var v49 uint32 = 751
                                            _ = v49
                                        case "hello":
                                            var v50 float32 = v46
                                            _ = v50
                                        default:
                                            var v51 uint8 = v48
                                            _ = v51
                                        }
                                    case -2147483648:
                                        if len(v16) > 0 {
                                            v16[0] = (fi43 >> ((4294967295) & 63))
                                        }
                                    }
                                    if v3 {
                                        fmt.Println(v3, v2, v48)
                                        v16 = append(v16, fi45)
                                        fmt.Println(v2, v1, fi43)
                                        var v52 float32 = (1e10 + 0.1)
                                        _ = v52
                                    } else {
                                        switch fi45 {
                                        case -370:
                                            v53 := GenS2{F0: 2, F1: (^(v1)), F2: -158}
                                            _ = v53
                                        default:
                                            if true {
                                                v2 = v39
                                                fmt.Println(v46)
                                                switch "foo bar" {
                                                case "":
                                                    var v54 float64 = -2.25
                                                    _ = v54
                                                case "foo bar":
                                                    if len(v16) >= 2 {
                                                        v16 = v16[1:len(v16)]
                                                    }
                                                case "a":
                                                    v55 := GenS4{F0: 0}
                                                    _ = v55
                                                default:
                                                    v46 -= (((1e10 / ((v46 - v46) * v46)) + 3.14159) - v46)
                                                    if true {
                                                        v56 := GenS4{F0: -2147483648}
                                                        var v57 int32 = 1
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    if v3 {
                                                        if false {
                                                            v40 -= v40
                                                            switch fi45 {
                                                            case -1:
                                                                v58 := GenS4{F0: -901}
                                                                var v59 int32 = 0
                                                                v60 := GenS2{F0: 4294967295, F1: 0, F2: -1}
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            var v61 uint16 = (^(v1))
                                                            _ = v61
                                                        } else {
                                                            if (325 >= 42) {
                                                                v62 := GenS1{F0: 2, F1: int16(v1)}
                                                                _ = v62
                                                            }
                                                            if true {
                                                                if v3 {
                                                                    v40 = (v40 - v48)
                                                                }
                                                                var v63 uint64 = uint64(v40)
                                                                var v64 int8 = v2
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                v39 = v39
                                                            }
                                                            var v65 int64 = -1
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v46 += (v46 + v46)
                                                            }
                                                            _ = v65
                                                        }
                                                        switch fi43 {
                                                        case 0:
                                                            fmt.Println(v41, fi43, fi43)
                                                        case -2147483648:
                                                            var v66 uint = 984
                                                            var v67 uint16 = (^(v1))
                                                            _ = v66
                                                            _ = v67
                                                        case 689:
                                                            if len(v16) >= 2 {
                                                                v16 = v16[1:len(v16)]
                                                            }
                                                        }
                                                        var v68 float32 = 1.5
                                                        _ = v68
                                                    } else {
                                                        var v69 uint16 = v1
                                                        var v70 int64 = 557
                                                        v71 := GenS4{F0: int32(fi43)}
                                                        var v72 uint32 = 968
                                                        _ = v69
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                }
                                            } else {
                                                if (!false) {
                                                    v39 = v2
                                                }
                                                switch (-(fi45)) {
                                                case 4294967295:
                                                    v39 = 100
                                                case 32767:
                                                    var v73 uint = 42
                                                    _ = v73
                                                case -75:
                                                    var v74 uint16 = v1
                                                    _ = v74
                                                default:
                                                    if v3 {
                                                        v75 := GenS5{F0: 4294967295, F1: 1}
                                                        v76 := GenS1{F0: 42, F1: 32767}
                                                        v1 = v1
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        _ = len(v16)
                                                        fmt.Println(v40, v42)
                                                    }
                                                }
                                                fmt.Println(fi43)
                                                v40 += 1
                                            }
                                        }
                                        fmt.Println(v40, v46, v3)
                                    }
                                    _ = v48
                                }
                            }
                            _ = len(v16)
                        case 31:
                            if len(v16) > 0 {
                                v16[0] = fi43
                            }
                        default:
                            switch "hello" {
                            case "a":
                                v3 = v3
                            }
                        }
                        var v77 string = ""
                        if v3 {
                            v78 := GenS3{F0: uint32(v1), F1: v77, F2: v77}
                            _ = v78
                        }
                        v39 = (v39 / (((v39 | (^(v39)))) | 1))
                        if ((v2 / ((v39) | 1)) < (^(v2))) {
                            v79 := GenS4{F0: 2147483647}
                            var v80 uint = 4294967295
                            _ = v79
                            _ = v80
                        } else {
                            switch ("foo bar" + v77) {
                            case "a":
                                fmt.Println(v40, v1, v2)
                                fmt.Println(v1)
                            case "hello":
                                var v81 float32 = -2.25
                                _ = v81
                            }
                            v77 = "a"
                            v82 := GenS5{F0: 2, F1: 42}
                            _ = v82
                        }
                        _ = v46
                        _ = v77
                    } else {
                        v1 = v1
                        switch "\n\t" {
                        case "a":
                            if len(v16) > 0 {
                                v16[0] = (fi45 >> ((1) & 63))
                            }
                        default:
                            _ = len(v16)
                        }
                        v2 += -128
                    }
                } else {
                    fmt.Println(v42, v2, v42)
                }
                v39 = (v39 | (v2 + v2))
            } else {
                var v83 uint16 = v1
                switch ("" + (((("foo bar" + ("\n\t" + "a")) + "\n\t") + "") + "\n\t")) {
                case "hello":
                    v1 = 2
                case "a":
                    fmt.Println(v3, fi43, v39)
                default:
                    v39 += v2
                }
                _ = v83
            }
            _ = fi45
        }
        _ = len(v16)
        _ = fi43
    }
    if (42 != 738) {
        if true {
            v3 = ("" <= "a")
            var v84 float64 = (1.5 * v42)
            _ = v84
        } else {
            if (v40 < 52) {
                if len(v16) > 0 {
                    v16[0] = 127
                }
                if false {
                    v85 := GenS6{F0: (^(v40))}
                    _ = v85
                }
                for fi86 := 0; fi86 < 4; fi86++ {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi86
                }
            } else {
                switch 32768 {
                case 9223372036854775807:
                    var v87 int16 = 1000
                    _ = v87
                case 1:
                    for fi88 := 0; fi88 < 13; fi88++ {
                        v89 := GenS3{F0: 65535, F1: "\n\t", F2: ("" + "hello")}
                        _ = fi88
                        _ = v89
                    }
                case 32767:
                    var v90 float64 = (0.0 * float64(0.0))
                    _ = v90
                default:
                    v91 := GenS1{F0: 2147483647, F1: -32768}
                    _ = v91
                }
                v92 := GenS6{F0: v40}
                for fi93 := 0; fi93 < 8; fi93++ {
                    var v94 uint = 0
                    switch (fi93 % (((-(fi93))) | 1)) {
                    case -335:
                        var v95 uint8 = uint8(v94)
                        _ = v95
                    default:
                        fmt.Println(v39)
                    }
                    _ = fi93
                    _ = v94
                }
                _ = v92
            }
            var v96 int32 = 42
            _ = v96
        }
        for fi97 := 0; fi97 < 11; fi97++ {
            if v3 {
                for fi98 := 0; fi98 < 4; fi98++ {
                    var v99 uint32 = 2147483647
                    fmt.Println(v2)
                    v100 := GenS7{F0: 390}
                    _ = fi98
                    _ = v100
                    _ = v99
                }
                for fi101 := 0; fi101 < 1; fi101++ {
                    var v102 string = ""
                    _ = fi101
                    _ = v102
                }
                var v103 float64 = 0.0
                _ = v103
            } else {
                for fi104 := 0; fi104 < 5; fi104++ {
                    v16 = append(v16, 127)
                    _ = fi104
                }
                _ = len(v16)
                switch ("a" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2)
                case "":
                    var v105 int64 = -1
                    _ = v105
                default:
                    v3 = (((v41 * v42) + 0.1) < v41)
                }
            }
            _ = fi97
        }
        var v106 uint16 = (^(v1))
        _ = v106
    }
    var v107 uint8 = 2
    fmt.Println(v2, v39, v2)
    var v108 string = ("\n\t" + ((("a" + ("" + "a")) + "hello") + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v107
    _ = v108
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
