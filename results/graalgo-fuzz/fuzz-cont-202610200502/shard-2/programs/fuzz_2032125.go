package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS3 struct { F0 string; F1 string; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 string; F1 uint8; F2 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int; F1 int32; F2 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 uint32; F1 int64; F2 string }
type GenS13 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 106
    v4 := make(map[uint32]int)
    var v5 uint = 10
    var v6 uint = (v5 & 2)
    if (true && true) {
        fmt.Println(v2, v2, v5)
    }
    v7 := GenS1{F0: -128, F1: 42, F2: 2}
    for fi8 := 0; fi8 < 10; fi8++ {
        switch "" {
        case "hello":
            if (1 <= -9223372036854775808) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if (65535 > 0) {
                        var v10 uint64 = 1
                        _ = v10
                    } else {
                        v5 = 42
                        if true {
                            v5 += v5
                            v5 = 1
                        }
                        v6 = (^(v5))
                    }
                    _ = fi9
                }
                v5 = 255
            } else {
                if true {
                    switch "a" {
                    case "\n\t":
                        switch fi8 {
                        case -379:
                            var v11 int16 = 1000
                            _ = v11
                        case 42:
                            switch 239 {
                            case 127:
                                _ = len(v1)
                            case -744:
                                if true {
                                    for fi12 := 0; fi12 < 8; fi12++ {
                                        var v13 float32 = (-2.25 - 1e10)
                                        fmt.Println(fi12)
                                        v14 := GenS1{F0: 127, F1: -1, F2: 2}
                                        _ = fi12
                                        _ = v13
                                        _ = v14
                                    }
                                } else {
                                    for fi15 := 0; fi15 < 12; fi15++ {
                                        v3 = 688
                                        v3 = (v3 & (v3 + v3))
                                        _ = fi15
                                    }
                                    for fi16 := 0; fi16 < 6; fi16++ {
                                        v5 += v5
                                        v17 := v1[("a" + "")]
                                        var v18 uint64 = 2
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            }
                        case -319:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v5 += v6
                                if (!true) {
                                    if false {
                                        v5 += 255
                                        v20 := GenS1{F0: 0, F1: -409, F2: 4294967295}
                                        v21 := GenS2{F0: ((18446744073709551615 != 18446744073709551615) || ((0 < 656) || (-64 >= 2147483647))), F1: 710, F2: 1}
                                        _ = v20
                                        _ = v21
                                    }
                                    if true {
                                        var v22 int64 = 1
                                        if false {
                                            var v23 uint8 = 0
                                            var v24 uint = 4294967295
                                            v5 += 897
                                            _ = v23
                                            _ = v24
                                        } else {
                                            var v25 float32 = ((1.5 + (float32(float64(3.14159)) + 0.1)) + -2.25)
                                            fmt.Println(fi8, v2)
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v6, v6)
                                                var v26 string = ("a" + "hello")
                                                _ = v26
                                            default:
                                                v27 := GenS1{F0: -26, F1: -1, F2: 18446744073709551615}
                                                _ = v27
                                            }
                                            if false {
                                                v2 -= (-(v2))
                                                v3 += v3
                                                v3 = (v3 + (v3 % ((536) | 1)))
                                            }
                                            _ = v25
                                        }
                                        _ = v22
                                    } else {
                                        var v28 int16 = 79
                                        var v29 int32 = -974
                                        _ = v28
                                        _ = v29
                                    }
                                    v30 := GenS3{F0: ("foo bar" + "hello"), F1: "\n\t", F2: (1 != -664)}
                                    v31 := v1[(("a" + "hello") + "hello")]
                                    _ = v30
                                    _ = v31
                                } else {
                                    if false {
                                        fmt.Println(fi8, v6, v3)
                                        v32 := GenS3{F0: "\n\t", F1: ("hello" + "\n\t"), F2: true}
                                        var v33 uint = v6
                                        _ = v32
                                        _ = v33
                                    } else {
                                        if (!true) {
                                            switch (fi8 % ((fi8) | 1)) {
                                            case 9223372036854775807:
                                                v34 := GenS1{F0: 119, F1: -2147483648, F2: 0}
                                                _ = v34
                                            case -9223372036854775808:
                                                fmt.Println(v6)
                                            case 0:
                                                if ((1e10 * ((1.5 * -2.25) - 0.0)) < -2.25) {
                                                    v5 = 2
                                                }
                                                switch (fi8 + (fi8 & (^(fi19)))) {
                                                case 0:
                                                    if (!true) {
                                                        v35 := GenS4{F0: (-9223372036854775808 >= (v2 - (v2 * (v2 % ((v2) | 1)))))}
                                                        var v36 int64 = v2
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                case 4294967295:
                                                    fmt.Println(v6, v6, v5)
                                                }
                                                if ((128 != 127) || true) {
                                                    v37 := v4[604]
                                                    _ = v37
                                                } else {
                                                    if false {
                                                        v5 -= 2
                                                    } else {
                                                        fmt.Println(v5)
                                                        var v38 float64 = float64((0.0 - (3.14159 - 0.1)))
                                                        _ = len(v4)
                                                        if (((^(v2)) != (v2 * v2)) && (1000 > -123)) {
                                                            v39 := GenS2{F0: ((!(v2 <= v2)) && true), F1: 706, F2: 151}
                                                            _ = v39
                                                        } else {
                                                            var v40 string = "hello"
                                                            v41 := GenS3{F0: v40, F1: ((v40 + ("\n\t" + "\n\t")) + v40), F2: false}
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        v5 = 18446744073709551615
                                                        v42 := GenS5{F0: "a", F1: 0, F2: uint64(v6)}
                                                        v43 := GenS6{F0: ("\n\t" + ("a" + "")), F1: 176, F2: (fi19 % ((fi19) | 1))}
                                                        _ = v38
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    var v44 uint = (v5 - (^(v6)))
                                                    _ = v44
                                                }
                                            }
                                        } else {
                                            if false {
                                                v45 := []int{}
                                                if false {
                                                    var v46 bool = true
                                                    _ = v46
                                                } else {
                                                    if true {
                                                        var v47 float64 = 0.1
                                                        _ = v47
                                                    }
                                                    fmt.Println(v2)
                                                    var v48 int = (fi8 << (((v5 << ((255) & 63))) & 63))
                                                    switch "a" {
                                                    case "":
                                                        var v49 uint32 = 105
                                                        _ = v49
                                                    case "foo bar":
                                                        var v50 uint16 = 298
                                                        _ = v50
                                                    case "a":
                                                        if true {
                                                            v5 -= (v6 ^ 65535)
                                                            v51 := v1[("\n\t" + "a")]
                                                            if ("foo bar" <= (("hello" + "") + "hello")) {
                                                                fmt.Println(fi8)
                                                                var v52 float32 = 1.5
                                                                _ = v52
                                                            }
                                                            _ = v51
                                                        } else {
                                                            v53 := GenS7{F0: 772}
                                                            _ = v53
                                                        }
                                                    default:
                                                        if true {
                                                            _ = len(v45)
                                                            fmt.Println(v48, v6, v2)
                                                            v2 = -1
                                                        } else {
                                                            if false {
                                                                if (255 < 0) {
                                                                    v54 := GenS8{F0: (v2 - v2), F1: 0}
                                                                    fmt.Println(v48)
                                                                    v48 = v48
                                                                    _ = v54
                                                                } else {
                                                                    v5 = (v6 << (((v6 % ((v6) | 1))) & 63))
                                                                }
                                                                v48 -= 127
                                                                v2 = v2
                                                                delete(v4, (^(v3)))
                                                                if false {
                                                                    var v55 uint = 65535
                                                                    _ = v55
                                                                }
                                                            } else {
                                                                v56 := GenS3{F0: "hello", F1: "hello", F2: true}
                                                                v57 := GenS7{F0: (v3 % ((v3) | 1))}
                                                                fmt.Println(fi8, fi8)
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            v58 := GenS1{F0: 127, F1: -557, F2: 2}
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if true {
                                                                    if true {
                                                                        v59 := v4[(v3 << (((^(v6))) & 63))]
                                                                        var v60 float32 = 1e10
                                                                        _ = v59
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 int = fi19
                                                                        switch ((("\n\t" + ("a" + "hello")) + "a") + "\n\t") {
                                                                        case "\n\t":
                                                                            if false {
                                                                                v62 := GenS1{F0: 42, F1: 1, F2: 42}
                                                                                fmt.Println(fi19, v2, v6)
                                                                                v63 := GenS9{F0: (fi19 ^ 0), F1: 42, F2: 42}
                                                                                _ = v62
                                                                                _ = v63
                                                                            } else {
                                                                                v64 := GenS4{F0: false}
                                                                                if true {
                                                                                    v5 -= 65535
                                                                                    if false {
                                                                                        v65 := GenS10{F0: v3}
                                                                                        _ = v65
                                                                                    }
                                                                                    v2 += int64(v61)
                                                                                }
                                                                                var v66 float64 = 1e10
                                                                                _ = v64
                                                                                _ = v66
                                                                            }
                                                                        case "a":
                                                                            v3 += v3
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    var v67 uint32 = (v3 << ((v5) & 63))
                                                                    _ = v67
                                                                } else {
                                                                    var v68 uint16 = 844
                                                                    _ = v68
                                                                }
                                                                v2 = (v2 % (((v2 - v2)) | 1))
                                                            }
                                                            _ = v58
                                                        }
                                                    }
                                                    fmt.Println(v6)
                                                    var v69 float32 = (1e10 - (3.14159 / 1e10))
                                                    _ = v48
                                                    _ = v69
                                                }
                                                _ = v45
                                            }
                                            if (!true) {
                                                v70 := v4[v3]
                                                _ = v70
                                            }
                                        }
                                        fmt.Println(v2)
                                        fmt.Println(v5)
                                    }
                                    var v71 int8 = -43
                                    _ = v71
                                }
                                switch ("foo bar" + "") {
                                case "hello":
                                    v72 := []int{(^(fi19)), fi8, 4294967295, (fi19 / (((fi8 * 1)) | 1))}
                                    _ = v72
                                default:
                                    fmt.Println(v5)
                                }
                                v73 := []int{(fi8 % ((fi19) | 1)), fi19, (fi8 + (fi19 << (((v6 - 18446744073709551615)) & 63)))}
                                v74 := GenS11{F0: true, F1: 995}
                                _ = fi19
                                _ = v73
                                _ = v74
                            }
                        }
                    case "":
                        switch ("" + "foo bar") {
                        case "a":
                            fmt.Println(v6, v3, fi8)
                        case "hello":
                            var v75 string = "foo bar"
                            _ = v75
                        case "":
                            fmt.Println(fi8, fi8)
                        default:
                            var v76 string = ("\n\t" + "\n\t")
                            _ = v76
                        }
                    default:
                        for fi77 := 0; fi77 < 15; fi77++ {
                            var v78 string = ("\n\t" + "foo bar")
                            switch (v78 + (v78 + v78)) {
                            case "\n\t":
                                var v79 uint64 = 2
                                _ = v79
                            case "hello":
                                var v80 int16 = 307
                                _ = v80
                            case "a":
                                var v81 int8 = 1
                                _ = v81
                            default:
                                v82 := []int{}
                                if ((-2.25 > -2.25) || true) {
                                    switch (v78 + v78) {
                                    case "a":
                                        if true {
                                            var v83 int = fi77
                                            var v84 uint = 1
                                            if true {
                                                switch "a" {
                                                case "":
                                                    if false {
                                                        v1[(("\n\t" + (v78 + v78)) + "\n\t")] = false
                                                        var v85 int64 = v2
                                                        _ = v85
                                                    } else {
                                                        var v86 int32 = 0
                                                        v87 := GenS4{F0: false}
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                default:
                                                    if (v2 <= (^(v2))) {
                                                        v84 = v5
                                                    }
                                                }
                                                switch fi8 {
                                                case -9223372036854775808:
                                                    v3 += (v3 / ((v3) | 1))
                                                }
                                            }
                                            _ = v83
                                            _ = v84
                                        }
                                    case "\n\t":
                                        fmt.Println(v5, v6, v6)
                                    case "hello":
                                        switch (v78 + (v78 + "foo bar")) {
                                        case "hello":
                                            var v88 int32 = -2147483648
                                            _ = v88
                                        case "foo bar":
                                            if false {
                                                _ = len(v82)
                                            } else {
                                                var v89 string = v78
                                                fmt.Println(fi8)
                                                var v90 uint64 = 18446744073709551615
                                                _ = v89
                                                _ = v90
                                            }
                                        default:
                                            v3 -= (^(v3))
                                        }
                                    }
                                    v5 -= v6
                                    switch fi77 {
                                    case 9223372036854775807:
                                        v2 -= 1
                                    }
                                }
                                fmt.Println(v5, v78, v6)
                                _ = v82
                            }
                            switch (fi8 & 32767) {
                            case 226:
                                fmt.Println(v3, v3)
                            case 140:
                                fmt.Println(v2)
                            default:
                                switch "hello" {
                                case "\n\t":
                                    v78 = ""
                                case "":
                                    v6 += (v5 | uint(v3))
                                case "hello":
                                    if (42 <= 42) {
                                        v78 = v78
                                        var v91 int16 = -1
                                        _ = v91
                                    }
                                }
                                if (false && true) {
                                    v5 = 111
                                } else {
                                    v3 -= (v3 % ((0) | 1))
                                    fmt.Println(v6, fi8, fi77)
                                }
                                var v92 uint = v5
                                _ = v92
                            }
                            _ = fi77
                            _ = v78
                        }
                    }
                }
            }
        case "foo bar":
            v93 := GenS7{F0: 1}
            _ = v93
        case "":
            v6 = v6
        default:
            if true {
                v3 = v3
            } else {
                var v94 string = "foo bar"
                v2 = v2
                var v95 int16 = 1
                _ = v94
                _ = v95
            }
        }
        _ = fi8
    }
    v96 := GenS12{F0: 2, F1: 0, F2: ""}
    v97 := GenS13{F0: 4294967295}
    switch 0 {
    case 32767:
        v2 += v2
    case 127:
        var v98 int32 = 0
        _ = v98
    }
    for fi99 := 0; fi99 < 14; fi99++ {
        var v100 uint64 = 905
        var v101 bool = false
        v102 := v1[("\n\t" + "\n\t")]
        _ = fi99
        _ = v100
        _ = v101
        _ = v102
    }
    for fi103 := 0; fi103 < 5; fi103++ {
        v2 += v2
        fmt.Println(v6, v6, v3)
        _ = len(v1)
        _ = fi103
    }
    v1["hello"] = (626 <= 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v96
    _ = v97
}
