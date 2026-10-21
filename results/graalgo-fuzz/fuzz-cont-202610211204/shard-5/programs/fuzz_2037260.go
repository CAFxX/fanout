package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS11 struct { F0 int64; F1 int16; F2 int8 }
type GenS12 struct { F0 uint16 }

func main() {
    var v1 bool = true
    var v2 bool = v1
    var v3 float32 = -2.25
    var v4 uint32 = 603
    var v5 uint16 = 745
    switch "foo bar" {
    case "":
        switch "\n\t" {
        case "\n\t":
            v1 = (525 >= 65535)
        default:
            switch "hello" {
            case "":
                v6 := []int{-1, 128}
                _ = v6
            default:
                v7 := GenS1{F0: (32767 != (v5 + v5))}
                _ = v7
            }
        }
    case "foo bar":
        var v8 uint32 = (v4 ^ (^(v4)))
        _ = v8
    case "hello":
        v9 := GenS2{F0: -1, F1: 0}
        _ = v9
    default:
        var v10 uint = 533
        _ = v10
    }
    v3 = v3
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint = 1
        v1 = ((42 >= 32767) || (100 >= 0))
        if v1 {
            if false {
                if (!v2) {
                    v12 = (^(v12))
                    if v2 {
                        v13 := []int{9223372036854775807, (fi11 & (fi11 ^ (fi11 * (-(fi11))))), fi11}
                        _ = len(v13)
                        _ = v13
                    }
                } else {
                    v2 = false
                }
            }
        } else {
            var v14 uint8 = 42
            v4 = uint32(v12)
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v5, v5)
    var v15 string = ""
    if v2 {
        var v16 int = 2147483647
        v17 := GenS3{F0: 595}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        switch 602 {
        case -988:
            switch "a" {
            case "a":
                v4 = (^(v4))
            case "\n\t":
                v19 := GenS4{F0: v1}
                _ = v19
            default:
                switch 32768 {
                case -2147483648:
                    var v20 int8 = -8
                    _ = v20
                default:
                    switch "hello" {
                    case "":
                        switch "" {
                        case "":
                            if ((v3 / (v3 - float32(1.5))) != v3) {
                                v21 := []int{(fi18 & 42)}
                                if true {
                                    var v22 uint16 = v5
                                    v23 := GenS5{F0: 18446744073709551615, F1: 1, F2: (v4 + v4)}
                                    _ = v22
                                    _ = v23
                                } else {
                                    switch (v15 + v15) {
                                    case "\n\t":
                                        var v24 string = (v15 + ((v15 + "foo bar") + (("\n\t" + v15) + v15)))
                                        _ = v24
                                    case "foo bar":
                                        if v1 {
                                            var v25 int16 = 0
                                            _ = v25
                                        }
                                    case "hello":
                                        var v26 uint32 = 42
                                        _ = v26
                                    }
                                    fmt.Println(v4, v15)
                                }
                                _ = v21
                            }
                        case "\n\t":
                            if v1 {
                                v27 := []int{}
                                var v28 float32 = (0.0 - float32(((0.0 / (float64(float32(((1e10 / -2.25) / 1.5))) - -2.25)) / 1e10)))
                                _ = v27
                                _ = v28
                            } else {
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        v4 -= v4
                                    case 128:
                                        var v30 int64 = 32
                                        _ = v30
                                    }
                                    v5 += 894
                                    var v31 int64 = -1
                                    v5 = v5
                                    _ = fi29
                                    _ = v31
                                }
                            }
                        }
                    case "\n\t":
                        var v32 float32 = 1.5
                        _ = v32
                    default:
                        v33 := GenS6{F0: 2}
                        _ = v33
                    }
                }
            }
        case -184:
            var v34 uint = 1
            _ = v34
        case -1:
            for fi35 := 0; fi35 < 1; fi35++ {
                fmt.Println(fi35)
                _ = fi35
            }
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            v4 = (v4 ^ (v4 & (v4 * (v4 | 594))))
            if v1 {
                switch "hello" {
                case "":
                    if (1 > -2147483648) {
                        var v37 int32 = -481
                        var v38 int32 = v37
                        v15 = "foo bar"
                        _ = v37
                        _ = v38
                    }
                case "\n\t":
                    switch fi36 {
                    case 127:
                        v39 := GenS2{F0: 941, F1: 100}
                        _ = v39
                    case 799:
                        var v40 uint64 = 720
                        _ = v40
                    case 0:
                        var v41 int16 = -1
                        _ = v41
                    default:
                        v42 := GenS5{F0: 1, F1: int8(fi36), F2: 522}
                        _ = v42
                    }
                case "a":
                    var v43 int8 = 0
                    _ = v43
                }
            } else {
                var v44 uint32 = (^(v4))
                v45 := []int{fi18, -9223372036854775808, (fi18 & fi18), 0}
                switch "" {
                case "hello":
                    v46 := GenS7{F0: -32768, F1: (3.14159 == v3)}
                    _ = v46
                case "":
                    if len(v45) > 0 {
                        v45[0] = fi18
                    }
                case "foo bar":
                    switch "a" {
                    case "hello":
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                    case "\n\t":
                        var v47 int = (fi36 * (fi36 << ((368) & 63)))
                        _ = v47
                    }
                    var v48 string = (("hello" + ((v15 + v15) + "\n\t")) + v15)
                    _ = v48
                }
                var v49 uint64 = 0
                _ = v44
                _ = v45
                _ = v49
            }
            if (230 > 42) {
                if (v1 && v1) {
                    if true {
                        v2 = v2
                        fmt.Println(v1, v15, v3)
                        v3 = float32((1e10 + 0.0))
                    } else {
                        fmt.Println(fi18)
                        switch v15 {
                        case "a":
                            fmt.Println(v3)
                        }
                    }
                    v50 := GenS8{F0: 42, F1: (v4 - 532)}
                    _ = v50
                } else {
                    fmt.Println(v3, v5)
                    var v51 bool = v1
                    v15 = "hello"
                    _ = v51
                }
                fmt.Println(fi36, v3, v1)
            } else {
                var v52 int32 = -1
                switch ("" + v15) {
                case "hello":
                    var v53 int8 = -1
                    _ = v53
                default:
                    v54 := GenS7{F0: -32768, F1: false}
                    v2 = (v2 || (!v2))
                    _ = v54
                }
                _ = v52
            }
            _ = fi36
        }
        _ = fi18
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        for fi56 := 0; fi56 < 12; fi56++ {
            if (0.1 <= 0.1) {
                fmt.Println(v15, v3, v3)
                switch v15 {
                case "":
                    v4 += v4
                default:
                    v57 := []int{(fi56 ^ (fi55 / ((fi56) | 1))), (-(fi56))}
                    _ = v57
                }
                if true {
                    v2 = true
                    var v58 bool = (!(!v2))
                    _ = v58
                } else {
                    v59 := GenS1{F0: ((fi56 << ((4294967295) & 63)) != 429)}
                    _ = v59
                }
            } else {
                if true {
                    var v60 int8 = 42
                    v61 := []int{}
                    v62 := GenS8{F0: 0, F1: (v4 + 408)}
                    _ = v60
                    _ = v61
                    _ = v62
                } else {
                    v63 := GenS1{F0: v1}
                    if (float64((v3 / 1.5)) < 1.5) {
                        if (1 > 116) {
                            v5 = v5
                            if false {
                                if (v1 || true) {
                                    var v64 uint32 = (v4 - (v4 / ((v4) | 1)))
                                    var v65 string = "hello"
                                    _ = v64
                                    _ = v65
                                }
                                if ((-2147483648 >= fi56) || v1) {
                                    v66 := []int{fi55}
                                    v67 := GenS2{F0: 0, F1: -127}
                                    v1 = false
                                    switch (fi55 >> ((956) & 63)) {
                                    case -814:
                                        fmt.Println(v5)
                                    case 42:
                                        fmt.Println(v5, v3, v1)
                                    case 127:
                                        v68 := GenS6{F0: 667}
                                        _ = v68
                                    default:
                                        if len(v66) > 0 {
                                            v66[0] = -185
                                        }
                                    }
                                    _ = v66
                                    _ = v67
                                }
                                if true {
                                    v69 := []int{128}
                                    _ = v69
                                } else {
                                    var v70 float32 = (v3 - v3)
                                    var v71 uint = 490
                                    if (true && (fi55 == fi55)) {
                                        var v72 uint64 = 18446744073709551615
                                        v2 = ((v72 | v72) == (v72 & (v72 - v72)))
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                            } else {
                                v73 := GenS8{F0: 108, F1: (^(v4))}
                                _ = v73
                            }
                            if ((643 < 0) || (v1 && (65535 > 386))) {
                                v5 = 65535
                            }
                        } else {
                            switch "" {
                            case "a":
                                var v74 int64 = -849
                                _ = v74
                            case "foo bar":
                                if (127 >= 250) {
                                    switch "a" {
                                    case "a":
                                        var v75 bool = (v1 && v2)
                                        _ = v75
                                    case "":
                                        v76 := GenS6{F0: 42}
                                        _ = v76
                                    case "hello":
                                        fmt.Println(v1, fi56)
                                    }
                                } else {
                                    if (!(!true)) {
                                        fmt.Println(fi56, v2)
                                        var v77 uint = 4294967295
                                        switch "foo bar" {
                                        case "hello":
                                            if (v1 && (v77 != v77)) {
                                                if (-2.25 == 1.5) {
                                                    v4 -= 2
                                                }
                                                v78 := []int{}
                                                var v79 int8 = 0
                                                _ = v78
                                                _ = v79
                                            } else {
                                                var v80 int = (fi56 % ((fi56) | 1))
                                                var v81 uint16 = v5
                                                _ = v80
                                                _ = v81
                                            }
                                        case "a":
                                            v1 = v1
                                        case "":
                                            if v1 {
                                                if (!v1) {
                                                    var v82 uint32 = (v4 & (v4 | v4))
                                                    v83 := []int{}
                                                    fmt.Println(v1, v5)
                                                    _ = v82
                                                    _ = v83
                                                }
                                                fmt.Println(fi56, v2)
                                                if false {
                                                    var v84 float32 = float32(-2.25)
                                                    var v85 int16 = 0
                                                    v86 := []int{675}
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    var v87 uint8 = 0
                                                    _ = v87
                                                }
                                            }
                                        }
                                        _ = v77
                                    }
                                    if (-1 <= 1) {
                                        v3 = v3
                                    } else {
                                        v1 = (4294967295 < 0)
                                        v88 := GenS9{F0: -412}
                                        var v89 uint32 = (v4 % (((v4 / (((v4 << ((65535) & 63))) | 1))) | 1))
                                        v90 := []int{fi56, fi55}
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                    }
                                    fmt.Println(v4, v5)
                                }
                            default:
                                if v1 {
                                    v91 := []int{(fi56 << ((uint(v4)) & 63))}
                                    var v92 int32 = 2147483647
                                    _ = v91
                                    _ = v92
                                }
                            }
                            var v93 uint32 = 524
                            v15 = (("" + ("\n\t" + (((((v15 + v15) + v15) + v15) + ("" + ("\n\t" + (v15 + "")))) + v15))) + "\n\t")
                            switch "\n\t" {
                            case "hello":
                                var v94 uint64 = 0
                                _ = v94
                            case "\n\t":
                                v95 := GenS6{F0: 65535}
                                _ = v95
                            case "foo bar":
                                v3 -= v3
                            default:
                                var v96 int64 = 0
                                _ = v96
                            }
                            switch v15 {
                            case "a":
                                v97 := GenS10{F0: v93, F1: 4294967295, F2: v5}
                                _ = v97
                            case "":
                                v98 := []int{(^(fi55))}
                                _ = v98
                            case "\n\t":
                                fmt.Println(v2)
                            }
                            _ = v93
                        }
                        v99 := GenS11{F0: -9223372036854775808, F1: 42, F2: 42}
                        _ = v99
                    } else {
                        v100 := []int{4294967295, (fi56 + fi55)}
                        _ = v100
                    }
                    switch fi56 {
                    case 2147483647:
                        v2 = true
                    default:
                        v101 := GenS12{F0: (v5 | v5)}
                        _ = v101
                    }
                    v102 := []int{fi56, fi56, 9223372036854775807}
                    _ = v102
                    _ = v63
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(v1, fi55)
                case "hello":
                    switch (fi56 | 882) {
                    case 530:
                        var v103 bool = v1
                        _ = v103
                    case -828:
                        if v1 {
                            v15 = "\n\t"
                            v104 := GenS4{F0: (!(-49 <= 42))}
                            _ = v104
                        } else {
                            fmt.Println(v2, v15, v1)
                            if v2 {
                                v105 := []int{fi55, (fi55 - fi56)}
                                _ = v105
                            }
                            v106 := []int{}
                            _ = v106
                        }
                    default:
                        var v107 int32 = -1
                        var v108 float32 = v3
                        fmt.Println(v4, v3)
                        var v109 int64 = 42
                        var v110 string = v15
                        _ = v107
                        _ = v108
                        _ = v109
                        _ = v110
                    }
                default:
                    var v111 uint = 2
                    _ = v111
                }
                v112 := []int{(fi55 | 4294967295), fi56, 128}
                _ = v112
            }
            switch v15 {
            case "\n\t":
                switch v15 {
                case "a":
                    v113 := GenS3{F0: 42}
                    _ = v113
                case "\n\t":
                    v1 = false
                case "foo bar":
                    v114 := []int{(-(fi55)), 128, (^(fi56))}
                    _ = v114
                default:
                    v115 := []int{fi55, 9223372036854775807, fi55}
                    _ = v115
                }
            default:
                fmt.Println(v1, v4)
            }
            _ = fi56
        }
        _ = fi55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
