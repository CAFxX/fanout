package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 string; F1 int; F2 int64 }
type GenS9 struct { F0 uint32; F1 string; F2 int64 }
type GenS10 struct { F0 int64; F1 int64; F2 bool }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 1
    var v3 uint = 65535
    var v4 uint16 = 255
    var v5 int16 = 1000
    var v6 uint32 = 208
    fmt.Println(v2, v5, v6)
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 = uint(v2)
        fmt.Println(fi7, v6, v6)
        var v8 uint8 = (v2 >> (((^(v3))) & 63))
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v2)
    if false {
        if true {
            if (-1 >= int8(v6)) {
                switch "" {
                case "\n\t":
                    if (!false) {
                        var v9 float64 = 1.5
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            if true {
                                v10 := GenS1{F0: (v6 % ((v6) | 1))}
                                switch 128 {
                                case -2147483648:
                                    var v11 int16 = (v5 ^ v5)
                                    _ = v11
                                }
                                var v12 uint64 = uint64(v2)
                                _ = v10
                                _ = v12
                            }
                        default:
                            v2 += v2
                        }
                        _ = v9
                    } else {
                        v13 := []int{-9223372036854775808, 275, -1, -19}
                        fmt.Println(v2, v2)
                        if true {
                            v14 := GenS1{F0: 2147483647}
                            var v15 uint = (^(v3))
                            switch -9223372036854775808 {
                            case 81:
                                switch "" {
                                case "foo bar":
                                    if true {
                                        fmt.Println(v15, v6)
                                        var v16 int8 = -128
                                        _ = v16
                                    }
                                default:
                                    v17 := GenS1{F0: (v6 & (v6 >> ((v3) & 63)))}
                                    _ = v17
                                }
                            case 32767:
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    if (-671 != (v1 * v1)) {
                                        for fi19 := 0; fi19 < 5; fi19++ {
                                            v2 = (v2 % ((v2) | 1))
                                            var v20 uint64 = 18446744073709551615
                                            fmt.Println(v1, v1)
                                            v21 := GenS1{F0: uint32(v20)}
                                            var v22 int8 = -72
                                            fmt.Println(v22)
                                            _ = fi19
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    } else {
                                        if (984 <= 4294967295) {
                                            fmt.Println(fi18, v4, v1)
                                        }
                                        if false {
                                            for fi23 := 0; fi23 < 16; fi23++ {
                                                if true {
                                                    if (255 <= 65535) {
                                                        v4 = v4
                                                        if (!(v2 > 0)) {
                                                            v5 = 32767
                                                            if len(v13) > 0 {
                                                                v13[0] = (fi23 | (fi18 >> (((v15 * 65535)) & 63)))
                                                            }
                                                            switch fi18 {
                                                            case -9223372036854775808:
                                                                var v24 uint64 = 47
                                                                _ = v24
                                                            case 649:
                                                                var v25 int16 = v5
                                                                _ = v25
                                                            case 147:
                                                                fmt.Println(v15, v3, v15)
                                                            }
                                                        } else {
                                                            var v26 uint8 = (v2 / (((v2 % ((2) | 1))) | 1))
                                                            var v27 int = 32767
                                                            v1 -= (v1 / ((-1) | 1))
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        switch (fi18 * (fi18 << ((842) & 63))) {
                                                        case -2147483648:
                                                            var v28 uint16 = 255
                                                            _ = v28
                                                        case 32767:
                                                            v29 := GenS2{F0: fi18, F1: uint32(v2), F2: 2}
                                                            var v30 int8 = 100
                                                            fmt.Println(v2)
                                                            _ = v29
                                                            _ = v30
                                                        case 0:
                                                            var v31 uint64 = 0
                                                            _ = v31
                                                        default:
                                                            if false {
                                                                _ = len(v13)
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v3)
                                                    }
                                                } else {
                                                    var v32 bool = (!((v4 >> ((v15) & 63)) > (v4 - 2)))
                                                    _ = v32
                                                }
                                                fmt.Println(fi23)
                                                switch "a" {
                                                case "\n\t":
                                                    v33 := GenS2{F0: int(v3), F1: v6, F2: 1}
                                                    _ = v33
                                                case "a":
                                                    fmt.Println(v6, fi23, v5)
                                                }
                                                _ = fi23
                                            }
                                        }
                                        if (v3 <= v3) {
                                            var v34 int = 235
                                            var v35 int8 = 0
                                            var v36 int64 = int64(v6)
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        } else {
                                            if len(v13) > 0 {
                                                v13[0] = (fi18 / ((fi18) | 1))
                                            }
                                            _ = len(v13)
                                            var v37 uint64 = uint64(v2)
                                            _ = v37
                                        }
                                    }
                                    _ = fi18
                                }
                            case 471:
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    if (1 >= v2) {
                                        var v39 float32 = (0.0 + (((1.5 * (0.1 - -2.25)) - -2.25) / -2.25))
                                        v39 -= 0.0
                                        _ = v39
                                    }
                                    _ = fi38
                                }
                            default:
                                v4 = v4
                            }
                            _ = v14
                            _ = v15
                        } else {
                            switch 4294967295 {
                            case -9223372036854775808:
                                v40 := GenS3{F0: (v2 + (v2 - v2)), F1: -1}
                                _ = v40
                            case -2147483648:
                                v41 := GenS4{F0: false, F1: ("\n\t" + ("a" + "hello"))}
                                _ = v41
                            default:
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    for fi43 := 0; fi43 < 14; fi43++ {
                                        v44 := GenS2{F0: (fi42 - fi43), F1: 0, F2: 1}
                                        _ = fi43
                                        _ = v44
                                    }
                                    var v45 bool = true
                                    v46 := GenS5{F0: (v6 * (v6 & (v6 / ((2) | 1))))}
                                    _ = fi42
                                    _ = v45
                                    _ = v46
                                }
                            }
                            for fi47 := 0; fi47 < 5; fi47++ {
                                for fi48 := 0; fi48 < 15; fi48++ {
                                    fmt.Println(v4)
                                    var v49 float64 = -2.25
                                    if true {
                                        switch (fi47 ^ fi48) {
                                        case -565:
                                            switch "" {
                                            case "a":
                                                switch -604 {
                                                case 424:
                                                    if ("" == "") {
                                                        v13 = append(v13, 842)
                                                        fmt.Println(v4)
                                                        var v50 bool = (!((^(v2)) > v2))
                                                        if (false || (((!true) || (!true)) && ((3.14159 - 1e10) == 0.1))) {
                                                            if len(v13) > 0 {
                                                                v13[0] = 571
                                                            }
                                                            v51 := GenS6{F0: 4294967295, F1: -485, F2: 376}
                                                            _ = v51
                                                        } else {
                                                            var v52 uint = (v3 * 255)
                                                            _ = v52
                                                        }
                                                        _ = v50
                                                    } else {
                                                        fmt.Println(fi47, v49, fi48)
                                                        v2 = v2
                                                    }
                                                default:
                                                    switch "foo bar" {
                                                    case "a":
                                                        if (("a" >= ("a" + ("foo bar" + ("foo bar" + "foo bar")))) || false) {
                                                            v13 = append(v13, fi48)
                                                        }
                                                    case "":
                                                        v5 += v5
                                                    default:
                                                        v6 = 647
                                                    }
                                                }
                                            case "foo bar":
                                                var v53 bool = false
                                                _ = v53
                                            }
                                            v6 += v6
                                            var v54 int8 = -122
                                            v55 := GenS2{F0: fi47, F1: (^(v6)), F2: 18446744073709551615}
                                            _ = v54
                                            _ = v55
                                        case 42:
                                            var v56 int16 = (^(v5))
                                            _ = v56
                                        }
                                        v57 := GenS1{F0: v6}
                                        var v58 int = fi48
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = fi48
                                    _ = v49
                                }
                                _ = fi47
                            }
                        }
                        _ = v13
                    }
                case "foo bar":
                    fmt.Println(v6)
                default:
                    v59 := GenS5{F0: (v6 | (v6 ^ (v6 << (((^(v3))) & 63))))}
                    _ = v59
                }
                for fi60 := 0; fi60 < 5; fi60++ {
                    var v61 uint16 = v4
                    v61 -= v61
                    _ = fi60
                    _ = v61
                }
            }
        }
        switch -1 {
        case 9223372036854775807:
            if true {
                v62 := []int{9223372036854775807, 127}
                _ = v62
            } else {
                if (32767 < -9223372036854775808) {
                    if (false || true) {
                        v63 := GenS2{F0: 9223372036854775807, F1: (v6 & uint32(v4)), F2: 16}
                        _ = v63
                    } else {
                        v2 = v2
                        for fi64 := 0; fi64 < 18; fi64++ {
                            v65 := GenS7{F0: 1, F1: 1}
                            switch ("hello" + "") {
                            case "hello":
                                for fi66 := 0; fi66 < 4; fi66++ {
                                    var v67 uint = v3
                                    _ = fi66
                                    _ = v67
                                }
                            case "a":
                                v4 = v4
                            case "foo bar":
                                v68 := GenS1{F0: v6}
                                _ = v68
                            }
                            var v69 int64 = v1
                            _ = fi64
                            _ = v65
                            _ = v69
                        }
                        v70 := GenS2{F0: 0, F1: uint32(v2), F2: 4294967295}
                        _ = v70
                    }
                }
            }
        case -9223372036854775808:
            for fi71 := 0; fi71 < 3; fi71++ {
                if false {
                    var v72 int16 = -32768
                    fmt.Println(v3, v1)
                    switch "hello" {
                    case "hello":
                        v6 -= v6
                    case "":
                        v73 := GenS2{F0: fi71, F1: 563, F2: 0}
                        _ = v73
                    default:
                        var v74 string = "foo bar"
                        _ = v74
                    }
                    _ = v72
                }
                v5 -= v5
                _ = fi71
            }
        case 42:
            if (!(true && (108 < 46))) {
                var v75 uint16 = uint16(v6)
                _ = v75
            }
        default:
            var v76 uint16 = v4
            _ = v76
        }
        var v77 uint64 = 0
        _ = v77
    } else {
        if false {
            if false {
                for fi78 := 0; fi78 < 15; fi78++ {
                    v1 += v1
                    v79 := GenS2{F0: (-(fi78)), F1: v6, F2: 18446744073709551615}
                    _ = fi78
                    _ = v79
                }
                v2 += v2
                switch -9223372036854775808 {
                case -9223372036854775808:
                    var v80 int32 = 1
                    _ = v80
                case -2147483648:
                    fmt.Println(v5, v2)
                default:
                    v81 := GenS4{F0: true, F1: "\n\t"}
                    _ = v81
                }
            }
            fmt.Println(v6, v3, v5)
        } else {
            v3 = uint(v6)
            var v82 uint64 = 2
            fmt.Println(v3, v2, v1)
            _ = v82
        }
    }
    fmt.Println(v6)
    var v83 bool = true
    v84 := GenS8{F0: ("a" + ("a" + "")), F1: 4294967295, F2: (^(v1))}
    fmt.Println(v5, v5)
    v85 := GenS5{F0: 4294967295}
    if v83 {
        var v86 uint8 = v2
        v5 = v5
        var v87 int32 = 42
        _ = v86
        _ = v87
    }
    var v88 uint64 = 18446744073709551615
    v2 += (v2 + uint8(v4))
    var v89 uint8 = (v2 - 255)
    if true {
        switch "a" {
        case "":
            switch "\n\t" {
            case "":
                switch 42 {
                case 118:
                    for fi90 := 0; fi90 < 8; fi90++ {
                        var v91 int32 = -2147483648
                        v92 := []int{(fi90 / (((fi90 | int(v4))) | 1))}
                        if (("\n\t" + (("a" + "foo bar") + "")) == "") {
                            var v93 int16 = v5
                            v94 := GenS7{F0: 97, F1: 244}
                            _ = v93
                            _ = v94
                        } else {
                            var v95 int64 = (v1 >> ((v3) & 63))
                            _ = v95
                        }
                        _ = fi90
                        _ = v91
                        _ = v92
                    }
                case 0:
                    if (-2147483648 >= -2147483648) {
                        var v96 uint = (v3 / (((v3 * v3)) | 1))
                        _ = v96
                    }
                default:
                    var v97 string = "foo bar"
                    _ = v97
                }
            case "foo bar":
                for fi98 := 0; fi98 < 18; fi98++ {
                    v99 := GenS9{F0: v6, F1: "\n\t", F2: v1}
                    if true {
                        v1 -= 732
                    }
                    if false {
                        v100 := GenS10{F0: int64(v4), F1: v1, F2: (v1 >= (v1 << ((v3) & 63)))}
                        _ = v100
                    }
                    _ = fi98
                    _ = v99
                }
            case "\n\t":
                var v101 int64 = (v1 % ((v1) | 1))
                _ = v101
            }
        }
        v102 := GenS7{F0: -58, F1: (^(v88))}
        v103 := []int{-250, 32768, 9223372036854775807, 698}
        _ = v102
        _ = v103
    } else {
        v6 = uint32(v89)
        switch "foo bar" {
        case "":
            v104 := []int{9223372036854775807}
            _ = v104
        case "\n\t":
            var v105 uint64 = v88
            _ = v105
        }
        v4 = uint16(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v85
    _ = v88
    _ = v89
}
