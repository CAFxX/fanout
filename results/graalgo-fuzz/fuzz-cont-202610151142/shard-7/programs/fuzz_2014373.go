package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int8; F1 int64; F2 int16 }
func genFn0() uint16 {
    var v6 uint = 4294967295
    if (!(v6 > (v6 >> ((v6) & 63)))) {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 541, F1: false}
            if (-12 <= 100) {
                v9 := GenS1{F0: 378, F1: (!true)}
                for fi10 := 0; fi10 < 12; fi10++ {
                    v6 = v6
                    var v11 uint16 = 2
                    if false {
                        if false {
                            var v12 float32 = (-2.25 / 1e10)
                            _ = v12
                        } else {
                            if ((("foo bar" + ("hello" + (("a" + "a") + ((("foo bar" + "\n\t") + (("" + "foo bar") + (("\n\t" + "\n\t") + "hello"))) + ("foo bar" + ("a" + ("" + (("hello" + "a") + "\n\t")))))))) + ("\n\t" + "\n\t")) >= ("foo bar" + (("" + "foo bar") + ""))) {
                                var v13 string = "a"
                                v6 += 255
                                _ = v13
                            }
                            v14 := []int{}
                            v6 += (v6 / (((v6 * v6)) | 1))
                            _ = v14
                        }
                        var v15 string = "a"
                        if (v11 < v11) {
                            v16 := GenS1{F0: 42, F1: true}
                            _ = v16
                        } else {
                            fmt.Println(v11, fi7, fi10)
                            v17 := []int{775}
                            v15 = (v15 + (v15 + "foo bar"))
                            _ = v17
                        }
                        var v18 float32 = 1.5
                        v19 := []int{}
                        _ = v15
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS1{F0: 42, F1: false}
                        _ = v20
                    }
                    _ = fi10
                    _ = v11
                }
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int32 = 2147483647
                    v6 = (v6 ^ v6)
                    var v23 uint32 = 2147483648
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v9
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                v6 = v6
                _ = fi24
            }
            _ = fi7
            _ = v8
        }
        if (!true) {
            var v25 int32 = 42
            v26 := GenS2{F0: 829, F1: ("\n\t" < ("foo bar" + "hello")), F2: 278}
            fmt.Println(v25)
            _ = v25
            _ = v26
        } else {
            v27 := []int{359, 1}
            _ = v27
        }
    }
    v28 := []int{}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    _ = v28
    _ = v6
    return 293
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 42
    var v3 string = ""
    var v4 uint = 1
    var v5 uint32 = 1
    var v29 uint64 = 18446744073709551615
    var v30 int16 = -758
    v31 := GenS1{F0: v30, F1: true}
    if false {
        var v32 int = 505
        var v33 int64 = -1
        _ = v32
        _ = v33
    }
    fmt.Println(v4, v5, v4)
    for fi34 := 0; fi34 < 20; fi34++ {
        v35 := GenS2{F0: 0, F1: true, F2: (fi34 << ((v4) & 63))}
        if (!((!false) && (((^(v4)) != (v4 + (v4 << ((v4) & 63)))) && (false && true)))) {
            var v36 float32 = 0.1
            v37 := GenS1{F0: (v30 & (^(v30))), F1: ((v2 << ((v4) & 63)) > 56)}
            v38 := GenS2{F0: 43, F1: false, F2: (-(fi34))}
            _ = v36
            _ = v37
            _ = v38
        } else {
            var v39 int8 = v2
            _ = v39
        }
        switch fi34 {
        case 1:
            for fi40 := 0; fi40 < 14; fi40++ {
                v41 := GenS3{F0: ((0 < 0) && (true || true)), F1: 2147483647}
                switch 1 {
                case 42:
                    switch ((v3 + (v3 + "a")) + v3) {
                    case "hello":
                        v42 := GenS4{F0: -1, F1: 687}
                        _ = v42
                    case "foo bar":
                        if true {
                            switch "hello" {
                            case "":
                                v1[fi34] = (v5 + 1)
                            case "foo bar":
                                v43 := GenS2{F0: 0, F1: true, F2: fi34}
                                _ = v43
                            case "hello":
                                v44 := GenS4{F0: v30, F1: 0}
                                _ = v44
                            default:
                                v5 = 2
                            }
                        }
                    case "a":
                        v2 = (v2 | v2)
                    default:
                        v2 = (v2 - v2)
                    }
                case 128:
                    v3 = v3
                case -9223372036854775808:
                    fmt.Println(v5, fi40)
                default:
                    if (32767 > 0) {
                        var v45 uint64 = (^(v29))
                        switch v3 {
                        case "hello":
                            if false {
                                v30 -= (v30 % ((32767) | 1))
                                v29 += 829
                                v46 := GenS2{F0: 664, F1: (!(true && false)), F2: 0}
                                switch (fi40 | fi34) {
                                case 32768:
                                    if false {
                                        v30 = (^(v30))
                                        v47 := []int{}
                                        if false {
                                            v4 = v4
                                        }
                                        _ = v47
                                    }
                                default:
                                    fmt.Println(v29, v5)
                                }
                                v48 := GenS4{F0: v30, F1: 42}
                                switch 790 {
                                case 1:
                                    switch (v3 + (v3 + v3)) {
                                    case "\n\t":
                                        v49 := []int{fi34}
                                        var v50 float32 = 0.0
                                        _ = v49
                                        _ = v50
                                    case "a":
                                        v45 -= v29
                                    }
                                default:
                                    var v51 float64 = 1e10
                                    _ = v51
                                }
                                _ = v46
                                _ = v48
                            } else {
                                fmt.Println(v30, v3, v29)
                            }
                        case "a":
                            var v52 int = fi34
                            _ = v52
                        case "\n\t":
                            if (0 == 2147483647) {
                                if ((v45 >> (((^(v4))) & 63)) <= 2) {
                                    var v53 uint16 = 265
                                    v3 = "hello"
                                    _ = v53
                                } else {
                                    v54 := []int{}
                                    if false {
                                        fmt.Println(v2, v5, v29)
                                        v55 := GenS2{F0: 32767, F1: false, F2: fi34}
                                        _ = v55
                                    } else {
                                        v3 = v3
                                        v29 = v29
                                        fmt.Println(v45, v2)
                                    }
                                    if (false || false) {
                                        fmt.Println(v4, fi34)
                                    }
                                    _ = v54
                                }
                            } else {
                                fmt.Println(v30, v45, v45)
                            }
                        }
                        v56 := v1[-265]
                        v57 := GenS1{F0: v30, F1: (42 > 209)}
                        v58 := []int{(fi40 % ((-490) | 1)), 42}
                        var v59 uint = (v4 | v4)
                        fmt.Println(fi40, v3, v4)
                        _ = v45
                        _ = v56
                        _ = v57
                        _ = v58
                        _ = v59
                    } else {
                        v60 := []int{(-(fi34)), int(v2)}
                        if (v2 == (^(v2))) {
                            if len(v60) >= 2 {
                                v60 = v60[1:len(v60)]
                            }
                        }
                        if false {
                            fmt.Println(v5)
                            var v61 int8 = v2
                            var v62 uint = v4
                            _ = v61
                            _ = v62
                        }
                        _ = v60
                    }
                }
                _ = fi40
                _ = v41
            }
        }
        _ = fi34
        _ = v35
    }
    for fi63 := 0; fi63 < 18; fi63++ {
        fmt.Println(v29, v30, v29)
        v4 = v4
        _ = fi63
    }
    if (255 > 255) {
        if true {
            switch v3 {
            case "foo bar":
                v64 := GenS3{F0: false, F1: 0}
                _ = v64
            default:
                v30 += 86
            }
        }
        if true {
            for fi65 := 0; fi65 < 19; fi65++ {
                switch -855 {
                case 42:
                    v66 := v1[fi65]
                    _ = v66
                case -11:
                    v67 := []int{fi65, (fi65 - fi65)}
                    _ = v67
                }
                _ = fi65
            }
            v68 := []int{619, -641}
            var v69 bool = ((v3 != v3) && (!false))
            _ = v68
            _ = v69
        } else {
            var v70 bool = (!(!false))
            _ = v70
        }
        v71 := v1[42]
        _ = v71
    } else {
        v72 := GenS2{F0: genFn0(), F1: (((((float32((1e10 - 1e10)) * 0.1) / (float32(float64(1e10)) + 3.14159)) + 3.14159) - -2.25) > 1e10), F2: -394}
        _ = v72
    }
    switch "foo bar" {
    case "hello":
        var v73 int32 = 42
        _ = v73
    case "a":
        v74 := GenS5{F0: -599, F1: v3}
        _ = v74
    default:
        v4 = (v4 * v4)
    }
    var v75 int32 = 42
    for fi76 := 0; fi76 < 9; fi76++ {
        var v77 int8 = v2
        fmt.Println(v30)
        _ = fi76
        _ = v77
    }
    var v78 uint16 = 255
    var v79 int16 = (-(v30))
    var v80 uint8 = 95
    if (!true) {
        switch "" {
        case "foo bar":
            for fi81 := 0; fi81 < 12; fi81++ {
                v82 := []int{}
                var v83 float64 = -2.25
                var v84 int32 = v75
                _ = fi81
                _ = v82
                _ = v83
                _ = v84
            }
        case "a":
            for fi85 := 0; fi85 < 1; fi85++ {
                fmt.Println(v2)
                var v86 int64 = 9223372036854775807
                _ = fi85
                _ = v86
            }
        case "\n\t":
            v29 += 1
        }
        v87 := v1[-1]
        _ = v87
    } else {
        _ = len(v1)
        var v88 uint16 = (v78 << ((42) & 63))
        _ = v88
    }
    switch (v3 + v3) {
    case "foo bar":
        v75 += v75
    case "\n\t":
        fmt.Println(v3, v5)
    default:
        if ((v75 >> ((256) & 63)) <= v75) {
            fmt.Println(v29, v3, v79)
        } else {
            for fi89 := 0; fi89 < 19; fi89++ {
                switch v3 {
                case "foo bar":
                    for fi90 := 0; fi90 < 19; fi90++ {
                        if ((float64(0.0) - -2.25) != (float64((1.5 - ((-2.25 * (((1.5 / -2.25) - -2.25) * (0.0 - 3.14159))) / 0.1))) + (0.0 - (3.14159 / ((-2.25 + 0.1) - 1.5))))) {
                            switch v3 {
                            case "a":
                                delete(v1, fi90)
                            case "\n\t":
                                var v91 uint16 = (v78 / ((42) | 1))
                                _ = v91
                            case "":
                                if (true && (1 != (v2 % (((v2 % (((-(v2))) | 1))) | 1)))) {
                                    if (!(uint32(v4) != v5)) {
                                        if false {
                                            var v92 uint32 = (v5 % (((v5 * uint32(v78))) | 1))
                                            _ = v92
                                        } else {
                                            if true {
                                                var v93 int16 = (v30 - int16(v5))
                                                var v94 bool = ((v80 & 255) > uint8(v4))
                                                _ = v93
                                                _ = v94
                                            }
                                        }
                                        switch 0 {
                                        case -2147483648:
                                            if false {
                                                fmt.Println(v2, v29)
                                                if ((!((v80 + v80) == uint8(v5))) && true) {
                                                    switch fi90 {
                                                    case -9223372036854775808:
                                                        v95 := GenS4{F0: v79, F1: v75}
                                                        _ = v95
                                                    default:
                                                        var v96 int16 = (v79 | 42)
                                                        _ = v96
                                                    }
                                                }
                                                var v97 string = ""
                                                _ = v97
                                            }
                                        }
                                        v98 := []int{-1, -447, 42}
                                        _ = v98
                                    } else {
                                        v99 := GenS2{F0: uint16(v4), F1: ((v75 >> ((v4) & 63)) == 42), F2: fi90}
                                        v79 = v79
                                        _ = v99
                                    }
                                }
                            }
                        }
                        var v100 int64 = 9223372036854775807
                        var v101 uint32 = (v5 + uint32(v4))
                        _ = fi90
                        _ = v100
                        _ = v101
                    }
                case "hello":
                    var v102 bool = (!false)
                    _ = v102
                case "a":
                    var v103 bool = false
                    _ = v103
                default:
                    v1[0] = 302
                }
                var v104 uint16 = (v78 % ((v78) | 1))
                _ = fi89
                _ = v104
            }
            v105 := GenS6{F0: -5, F1: 0, F2: int16(v78)}
            _ = v105
        }
    }
    var v106 int16 = (-(v30))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v106
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v75
    _ = v78
    _ = v79
    _ = v80
}
