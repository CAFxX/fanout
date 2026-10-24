package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int; F1 bool; F2 uint32 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int16; F1 int16 }
type GenS13 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS14 struct { F0 int }
func genFn0(p0 string) uint16 {
    var v5 uint64 = 1
    var v6 string = "hello"
    v6 = p0
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -545
    var v3 float64 = ((-2.25 + ((1.5 + 0.0) * float64(((float32(3.14159) * 1e10) * 0.1)))) / (-2.25 / float64(v1)))
    var v4 float64 = ((1e10 / ((v3 - (0.0 * float64(1.5))) - -2.25)) + (v3 - v3))
    if (-787 > 4294967295) {
        var v7 int8 = 117
        v8 := []int{1, -1, -1, 127}
        _ = v7
        _ = v8
    }
    v4 = v3
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                v2 = (v2 & -2147483648)
            } else {
                v11 := []int{}
                _ = v11
            }
            _ = fi10
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := []int{}
            if ((2147483647 == v2) && (0 <= 2)) {
                switch "foo bar" {
                case "hello":
                    if ((!true) || (false && (v1 >= v1))) {
                        var v14 string = "foo bar"
                        _ = v14
                    }
                case "":
                    fmt.Println(v4)
                default:
                    var v15 uint32 = 374
                    _ = v15
                }
            } else {
                v4 += (v3 - v3)
                if true {
                    fmt.Println(v1, fi12)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v16 int32 = (v2 - 1)
                    var v17 uint = 65535
                    if ((!(0 <= 2)) && false) {
                        if true {
                            if false {
                                v3 = 3.14159
                                v18 := GenS1{F0: -17}
                                if len(v13) > 0 {
                                    v13[0] = fi9
                                }
                                fmt.Println(v3, v17, v16)
                                var v19 uint16 = 832
                                v20 := GenS2{F0: fi9, F1: ""}
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                var v21 uint64 = 856
                                _ = v21
                            }
                            switch (fi9 << ((839) & 63)) {
                            case 127:
                                var v22 bool = (42 > 18446744073709551615)
                                var v23 float64 = float64((v1 * (v1 * v1)))
                                switch (("foo bar" + (("hello" + ("" + ("\n\t" + "\n\t"))) + "a")) + "") {
                                case "hello":
                                    if false {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        v17 = (v17 - (v17 & v17))
                                    }
                                case "":
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                default:
                                    if (false && true) {
                                        v1 = (v1 - v1)
                                        switch "" {
                                        case "":
                                            v24 := GenS1{F0: 42}
                                            _ = v24
                                        case "a":
                                            v3 += (((-2.25 - 0.1) * v4) - float64(((v1 * 3.14159) * float32(((float64(v1) * v23) * 1e10)))))
                                        case "foo bar":
                                            var v25 float32 = v1
                                            _ = v25
                                        }
                                        switch fi9 {
                                        case 4294967295:
                                            v26 := GenS2{F0: (fi12 ^ (fi12 | (fi12 >> ((v17) & 63)))), F1: ((("foo bar" + ("hello" + ("a" + "hello"))) + "\n\t") + "\n\t")}
                                            _ = v26
                                        case -535:
                                            if ((v17 / ((42) | 1)) > 18446744073709551615) {
                                                var v27 int8 = -128
                                                v3 = float64((0.0 / 1e10))
                                                var v28 uint8 = 7
                                                switch fi12 {
                                                case 127:
                                                    v22 = false
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    v17 = 973
                                                }
                                                if (!v22) {
                                                    v29 := GenS1{F0: 760}
                                                    var v30 uint64 = 461
                                                    if (1.5 != (float32((v3 / v23)) - float32(float64(0.0)))) {
                                                        var v31 int16 = 1
                                                        switch (("" + "foo bar") + "") {
                                                        case "":
                                                            fmt.Println(fi12, v22)
                                                        case "a":
                                                            var v32 uint64 = (^(v30))
                                                            _ = v32
                                                        case "foo bar":
                                                            v13 = append(v13, (fi9 | -2147483648))
                                                        }
                                                        fmt.Println(v23)
                                                        _ = v31
                                                    } else {
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                        var v33 int16 = 0
                                                        _ = v33
                                                    }
                                                    _ = v29
                                                    _ = v30
                                                }
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v34 := GenS3{F0: 1, F1: -146, F2: (fi12 >= fi9)}
                                                if v22 {
                                                    fmt.Println(v1, v23, v4)
                                                } else {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    switch (-(fi9)) {
                                                    case 326:
                                                        v35 := GenS4{F0: 4294967295, F1: (true && (!false))}
                                                        _ = v35
                                                    default:
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                        v36 := GenS5{F0: 1, F1: (false || false), F2: 2147483647}
                                                        _ = v36
                                                    }
                                                }
                                                var v37 uint64 = 4294967295
                                                _ = v34
                                                _ = v37
                                            }
                                        default:
                                            v3 = v23
                                        }
                                        v4 = v23
                                        var v38 string = "a"
                                        if (!v22) {
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            var v39 int64 = -1
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            _ = v39
                                        } else {
                                            switch fi9 {
                                            case 927:
                                                v22 = (v23 < -2.25)
                                            case 4294967295:
                                                if false {
                                                    var v40 uint32 = 2
                                                    _ = v40
                                                } else {
                                                    _ = len(v13)
                                                    if v22 {
                                                        if (true || false) {
                                                            v2 -= v16
                                                            switch v38 {
                                                            case "a":
                                                                v38 = v38
                                                            case "hello":
                                                                v23 -= ((0.1 / (1e10 - (float64((float32(float64((v1 / (float32(float64(((float32(float64((((v1 - (v1 * v1)) + (v1 / ((((float32((((float64(0.0) * float64((1.5 - ((v1 * v1) + v1)))) * (v4 * (v4 - 0.0))) / float64(float32(float64(v1))))) + v1) * (v1 / -2.25)) / (3.14159 - v1)) + v1))) + (v1 / 0.1)))) + (v1 * float32(float64(v1)))) * float32(((v3 * v4) / v3))))) * (v1 - v1))))) + 3.14159)) + (v23 + (v4 * (v3 + 1.5)))))) - (1e10 - v23))
                                                            case "\n\t":
                                                                var v41 uint16 = 50
                                                                _ = v41
                                                            }
                                                        }
                                                    } else {
                                                        v42 := GenS3{F0: 107, F1: -1, F2: v22}
                                                        v38 = ((v38 + (v38 + "\n\t")) + v38)
                                                        _ = v42
                                                    }
                                                }
                                            }
                                            fmt.Println(fi9, v38)
                                        }
                                        if (!true) {
                                            v22 = (354 == 959)
                                            fmt.Println(v17, v23, v2)
                                        }
                                        _ = len(v13)
                                        _ = v38
                                    }
                                }
                                if v22 {
                                    v2 -= 1
                                    if len(v13) > 0 {
                                        v13[0] = -901
                                    }
                                    v17 -= v17
                                    if len(v13) > 0 {
                                        v13[0] = 32767
                                    }
                                    var v43 uint16 = 32767
                                    var v44 int64 = 42
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        var v45 float32 = v1
                                        _ = v45
                                    case "hello":
                                        if len(v13) > 0 {
                                            v13[0] = 127
                                        }
                                    case "a":
                                        fmt.Println(v2, v16, v16)
                                    }
                                    if (!(!true)) {
                                        var v46 int64 = (v44 - int64(v43))
                                        var v47 int8 = 127
                                        _ = v46
                                        _ = v47
                                    } else {
                                        v1 = 0.0
                                    }
                                    var v48 float64 = (v3 - (v4 * float64(-2.25)))
                                    v49 := GenS1{F0: (v44 << ((122) & 63))}
                                    var v50 float64 = 1.5
                                    var v51 int64 = 532
                                    _ = v43
                                    _ = v44
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                } else {
                                    v52 := GenS6{F0: "a", F1: v22}
                                    if (164 >= 9223372036854775807) {
                                        _ = len(v13)
                                        switch -947 {
                                        case 2147483647:
                                            if len(v13) > 0 {
                                                v13[0] = (fi12 ^ -1)
                                            }
                                            switch 0 {
                                            case 32767:
                                                if ((-(fi12)) >= 128) {
                                                    v53 := GenS7{F0: 9223372036854775807, F1: v22}
                                                    _ = v53
                                                } else {
                                                    if (v22 || ((!(1 >= 18446744073709551615)) && v22)) {
                                                        if ((v2 ^ (v2 - int32(v17))) >= 2147483647) {
                                                            var v54 uint8 = 1
                                                            v55 := GenS2{F0: 127, F1: "a"}
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    } else {
                                                        fmt.Println(v3, fi9, v1)
                                                        if (((127 <= 144) && v22) || v22) {
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        } else {
                                                            v22 = v22
                                                        }
                                                        var v56 uint64 = 42
                                                        _ = v56
                                                    }
                                                }
                                            default:
                                                v4 = v3
                                            }
                                            _ = len(v13)
                                        case 1:
                                            v1 = v1
                                        default:
                                            switch "a" {
                                            case "foo bar":
                                                var v57 float32 = v1
                                                _ = v57
                                            }
                                        }
                                        fmt.Println(v22)
                                        v58 := GenS8{F0: 42}
                                        if true {
                                            var v59 uint = (v17 + v17)
                                            _ = v59
                                        } else {
                                            v60 := GenS5{F0: 32767, F1: (v2 == (v2 & v16)), F2: 4294967295}
                                            if v22 {
                                                if v22 {
                                                    var v61 uint64 = 4294967295
                                                    v62 := GenS9{F0: 127, F1: -33}
                                                    fmt.Println(v4, v23, v23)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                v1 = float32(((3.14159 - -2.25) * (v4 + 1.5)))
                                                var v63 string = ("foo bar" + "foo bar")
                                                _ = v63
                                            } else {
                                                var v64 float32 = 1.5
                                                _ = v64
                                            }
                                            fmt.Println(fi9)
                                            _ = v60
                                        }
                                        _ = v58
                                    } else {
                                        var v65 uint8 = 2
                                        v13 = append(v13, (-(fi9)))
                                        switch (fi12 / (((fi9 % ((fi9) | 1))) | 1)) {
                                        case 127:
                                            _ = len(v13)
                                        default:
                                            switch ("\n\t" + "") {
                                            case "a":
                                                _ = len(v13)
                                            case "\n\t":
                                                v66 := GenS9{F0: 73, F1: 6}
                                                _ = v66
                                            case "foo bar":
                                                switch (fi9 >> ((65535) & 63)) {
                                                case 32768:
                                                    var v67 int32 = int32(v17)
                                                    _ = v67
                                                case -9223372036854775808:
                                                    var v68 uint = (^(v17))
                                                    _ = v68
                                                case 9223372036854775807:
                                                    var v69 bool = (13 <= 1)
                                                    _ = v69
                                                }
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(v65, v16, fi9)
                                                }
                                            }
                                        }
                                        _ = v65
                                    }
                                    _ = v52
                                }
                                v70 := GenS10{F0: (v17 % ((v17) | 1))}
                                _ = v22
                                _ = v23
                                _ = v70
                            case -1:
                                v71 := GenS4{F0: 2, F1: false}
                                _ = v71
                            default:
                                v72 := GenS6{F0: "a", F1: (-1 >= 73)}
                                _ = v72
                            }
                        } else {
                            if (!((!(!true)) && false)) {
                                v13 = append(v13, 127)
                                fmt.Println(v2)
                            } else {
                                var v73 int64 = -9223372036854775808
                                _ = v73
                            }
                            if true {
                                fmt.Println(fi9, v2)
                                var v74 string = (("hello" + (((("\n\t" + "a") + (("hello" + "\n\t") + "")) + ("hello" + "\n\t")) + "a")) + "foo bar")
                                _ = v74
                            }
                            var v75 uint16 = genFn0("a")
                            _ = v75
                        }
                        var v76 bool = (0.1 >= (-2.25 / v1))
                        _ = v76
                    }
                    var v77 int8 = 124
                    fmt.Println(v1)
                    fmt.Println(v3)
                    _ = v16
                    _ = v17
                    _ = v77
                } else {
                    v78 := GenS7{F0: (fi9 + (^(fi9))), F1: false}
                    _ = v78
                }
                if len(v13) > 0 {
                    v13[0] = -580
                }
                v1 -= 3.14159
                var v79 string = ("hello" + "")
                _ = v79
            }
            _ = fi12
            _ = v13
        }
        for fi80 := 0; fi80 < 1; fi80++ {
            v81 := []int{(fi80 + fi9), 42}
            if len(v81) > 0 {
                v81[0] = fi9
            }
            v4 = v4
            if len(v81) >= 2 {
                v81 = v81[1:len(v81)]
            }
            v82 := GenS5{F0: -2147483648, F1: ((5 > -1) && false), F2: 2147483647}
            _ = fi80
            _ = v81
            _ = v82
        }
        _ = fi9
    }
    fmt.Println(v1, v2, v4)
    v83 := []int{32768, -102}
    var v84 int16 = 42
    for fi85 := 0; fi85 < 3; fi85++ {
        for fi86 := 0; fi86 < 17; fi86++ {
            fmt.Println(fi86)
            if len(v83) > 0 {
                v83[0] = (fi85 - fi85)
            }
            v4 -= v3
            _ = fi86
        }
        var v87 uint64 = 487
        _ = fi85
        _ = v87
    }
    var v88 float32 = (v1 / (v1 / v1))
    v89 := GenS5{F0: 143, F1: ((!(0 != 42)) || true), F2: 4294967295}
    var v90 uint8 = 255
    switch int(v2) {
    case 0:
        v90 += 127
    case 32767:
        if false {
            v83 = append(v83, 504)
        } else {
            var v91 uint = 1
            _ = v91
        }
    case 467:
        for fi92 := 0; fi92 < 6; fi92++ {
            v83 = append(v83, (fi92 * (fi92 % (((fi92 % ((fi92) | 1))) | 1))))
            v93 := GenS11{F0: 0}
            if true {
                var v94 float32 = 0.1
                var v95 int64 = -9223372036854775808
                _ = v94
                _ = v95
            }
            _ = fi92
            _ = v93
        }
    default:
        _ = len(v83)
    }
    var v96 bool = false
    v88 -= (v1 * (((v1 + (v1 + v88)) / float32(v4)) / 1.5))
    v97 := GenS2{F0: 0, F1: "a"}
    if v96 {
        if false {
            v83 = append(v83, 4294967295)
            fmt.Println(v96, v84)
            if v96 {
                if len(v83) > 0 {
                    v83[0] = 2147483647
                }
                v88 -= (v88 + (0.0 + v1))
                for fi98 := 0; fi98 < 13; fi98++ {
                    switch (fi98 << ((727) & 63)) {
                    case 91:
                        v83 = append(v83, (fi98 * 1))
                    default:
                        v99 := GenS12{F0: v84, F1: v84}
                        _ = v99
                    }
                    v100 := GenS13{F0: 1, F1: 2147483648, F2: 2147483647}
                    _ = fi98
                    _ = v100
                }
            } else {
                v84 = (-(v84))
                v2 = (v2 % (((^(v2))) | 1))
                var v101 string = "foo bar"
                _ = v101
            }
        } else {
            var v102 float32 = (((v1 * 0.1) - 3.14159) + (v88 * ((1.5 - v88) + ((v1 / v1) - 0.1))))
            _ = v102
        }
        var v103 string = "\n\t"
        for fi104 := 0; fi104 < 2; fi104++ {
            v105 := GenS10{F0: 2}
            _ = fi104
            _ = v105
        }
        _ = v103
    } else {
        var v106 int16 = v84
        v107 := GenS14{F0: 32768}
        _ = v106
        _ = v107
    }
    var v108 uint32 = 42
    v83 = append(v83, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v88
    _ = v89
    _ = v90
    _ = v96
    _ = v97
}
