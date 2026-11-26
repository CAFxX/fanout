package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int8; F1 int; F2 uint }
type GenS9 struct { F0 uint }

func main() {
    var v1 int32 = 42
    var v2 int = 2147483647
    v3 := make(map[uint32]int)
    if (!(42 < 42)) {
        fmt.Println(v1, v2)
    } else {
        for fi4 := 0; fi4 < 13; fi4++ {
            for fi5 := 0; fi5 < 3; fi5++ {
                var v6 int8 = 127
                _ = fi5
                _ = v6
            }
            if true {
                var v7 uint = 594
                _ = v7
            }
            _ = fi4
        }
        v1 += 1
        v8 := GenS1{F0: 244}
        _ = v8
    }
    switch "a" {
    case "\n\t":
        v9 := GenS1{F0: 440}
        _ = v9
    case "":
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS1{F0: 1}
        _ = v12
    }
    v13 := GenS2{F0: 162, F1: 1}
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            v2 = -2147483648
            if true {
                fmt.Println(v2, fi14, fi14)
                v1 = -2147483648
                v16 := GenS3{F0: 913}
                _ = v16
            } else {
                switch ("hello" + ("a" + "a")) {
                case "hello":
                    if (true || ((^(v1)) > v1)) {
                        var v17 int16 = 0
                        if false {
                            v3[4294967295] = 516
                            if false {
                                var v18 uint8 = 128
                                v17 = (-(v17))
                                var v19 float32 = 0.1
                                _ = v18
                                _ = v19
                            } else {
                                v20 := []int{fi15}
                                if (true || ((!false) && true)) {
                                    switch (v2 - (fi14 << ((18446744073709551615) & 63))) {
                                    case 4294967295:
                                        v1 = -2147483648
                                    case -9223372036854775808:
                                        delete(v3, 270)
                                    default:
                                        switch "" {
                                        case "hello":
                                            v21 := GenS4{F0: 4294967295, F1: 65535, F2: 65535}
                                            _ = v21
                                        case "a":
                                            if len(v20) > 0 {
                                                v20[0] = (-(fi15))
                                            }
                                        case "\n\t":
                                            v22 := v3[65535]
                                            _ = v22
                                        default:
                                            var v23 float64 = float64(0.1)
                                            _ = v23
                                        }
                                    }
                                    if false {
                                        fmt.Println(v17, fi15)
                                        v1 = 0
                                        var v24 int64 = 1
                                        _ = v24
                                    }
                                    fmt.Println(fi15, v2)
                                    switch ("" + ("hello" + ("foo bar" + ("a" + "\n\t")))) {
                                    case "a":
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    case "":
                                        v1 += v1
                                    }
                                }
                                _ = v20
                            }
                            switch v2 {
                            case -9223372036854775808:
                                v25 := GenS5{F0: 65535, F1: 994, F2: 18446744073709551615}
                                _ = v25
                            }
                        } else {
                            switch 32767 {
                            case -2147483648:
                                var v26 bool = true
                                _ = v26
                            case 42:
                                if true {
                                    v2 = (-(fi14))
                                    v17 = v17
                                } else {
                                    var v27 bool = (1.5 != (float32((0.0 * 0.1)) / 3.14159))
                                    var v28 uint32 = 65535
                                    _ = v27
                                    _ = v28
                                }
                                switch "hello" {
                                case "a":
                                    if (true || (!((v17 <= -1) || false))) {
                                        fmt.Println(fi15, v2, v2)
                                        v29 := GenS2{F0: 2, F1: 42}
                                        if false {
                                            v2 = (-(fi15))
                                            v30 := GenS6{F0: -306, F1: 100}
                                            _ = v30
                                        }
                                        v31 := GenS5{F0: 603, F1: 1, F2: 919}
                                        var v32 float64 = float64((0.0 * 1.5))
                                        _ = v29
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS3{F0: (^(v17))}
                                        _ = v33
                                    }
                                case "hello":
                                    var v34 float64 = ((-2.25 * 1.5) - 3.14159)
                                    _ = v34
                                default:
                                    var v35 int64 = -765
                                    _ = v35
                                }
                            default:
                                v1 = (v1 & v1)
                            }
                        }
                        v1 = (v1 / ((v1) | 1))
                        _ = v17
                    }
                    var v36 int8 = 100
                    _ = v36
                case "a":
                    var v37 bool = true
                    _ = v37
                case "\n\t":
                    v3[2] = (fi15 >> ((255) & 63))
                default:
                    if (42 < 1000) {
                        fmt.Println(v1, fi14)
                        if false {
                            fmt.Println(fi14, v1)
                        }
                    }
                }
                switch ("a" + (("" + "foo bar") + ("\n\t" + "\n\t"))) {
                case "hello":
                    v38 := GenS5{F0: 2147483647, F1: 656, F2: 4294967295}
                    _ = v38
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi15)
                    default:
                        if true {
                            switch (v2 & 127) {
                            case -460:
                                var v39 int = 1
                                _ = v39
                            case 98:
                                v40 := v3[42]
                                _ = v40
                            case 4294967295:
                                v41 := GenS2{F0: 255, F1: 359}
                                var v42 float32 = (-2.25 + 0.1)
                                _ = v41
                                _ = v42
                            }
                            v2 = (fi15 & fi14)
                        } else {
                            if (v1 == (v1 * int32(v2))) {
                                var v43 uint32 = 65535
                                v2 = int(v1)
                                var v44 int = (fi14 - -9223372036854775808)
                                _ = v43
                                _ = v44
                            } else {
                                var v45 uint = 255
                                v46 := GenS7{F0: 9223372036854775807, F1: ("a" + "foo bar")}
                                if (100 <= -128) {
                                    if (true || (3.14159 == (1.5 * -2.25))) {
                                        var v47 bool = ((1000 < 1) && (!(0 >= 0)))
                                        _ = v47
                                    }
                                    fmt.Println(fi14, fi14)
                                    v2 += v2
                                }
                                _ = v45
                                _ = v46
                            }
                            fmt.Println(v1, fi14, v1)
                            v48 := []int{(fi15 >> ((18446744073709551615) & 63))}
                            _ = v48
                        }
                    }
                case "foo bar":
                    var v49 int32 = 42
                    _ = v49
                default:
                    var v50 int64 = int64(v1)
                    _ = v50
                }
            }
            if false {
                if false {
                    if ("foo bar" < "hello") {
                        if false {
                            v1 -= v1
                            v51 := GenS8{F0: -1, F1: fi15, F2: 42}
                            var v52 uint8 = 127
                            _ = v51
                            _ = v52
                        }
                        var v53 float32 = float32(3.14159)
                        v54 := GenS1{F0: 0}
                        _ = v53
                        _ = v54
                    } else {
                        v1 = 42
                    }
                }
            } else {
                var v55 string = "a"
                if false {
                    if true {
                        var v56 uint8 = 2
                        var v57 uint32 = 715
                        _ = v56
                        _ = v57
                    }
                    var v58 uint32 = 2
                    _ = v58
                } else {
                    if true {
                        v59 := GenS8{F0: -10, F1: fi15, F2: 203}
                        var v60 float32 = (-2.25 + 1e10)
                        _ = v59
                        _ = v60
                    } else {
                        v1 += -1
                        switch "a" {
                        case "a":
                            v1 = v1
                            v61 := v3[2147483647]
                            if true {
                                v62 := v3[0]
                                var v63 float32 = (1e10 * float32(3.14159))
                                _ = v62
                                _ = v63
                            } else {
                                switch v55 {
                                case "a":
                                    v1 = v1
                                case "\n\t":
                                    v64 := []int{v2, (^(fi15)), fi14, (fi14 ^ (v2 + (v61 & v61)))}
                                    _ = v64
                                }
                                if (-860 > (v1 - 207)) {
                                    v65 := v3[321]
                                    fmt.Println(v61, fi14)
                                    _ = v65
                                } else {
                                    var v66 uint8 = 46
                                    _ = v66
                                }
                                if true {
                                    v3[2147483647] = (v2 + (^(fi14)))
                                    var v67 int64 = 1
                                    v68 := []int{v2, 259}
                                    _ = v67
                                    _ = v68
                                }
                            }
                            _ = v61
                        case "\n\t":
                            v3[4294967295] = (-(fi14))
                        }
                        v2 = (fi14 - (-(v2)))
                        var v69 int64 = -9223372036854775808
                        var v70 int8 = 78
                        _ = v69
                        _ = v70
                    }
                    var v71 uint16 = 32767
                    _ = v71
                }
                switch v55 {
                case "\n\t":
                    v2 = (fi14 >> ((164) & 63))
                }
                var v72 uint32 = 979
                _ = v55
                _ = v72
            }
            fmt.Println(fi15, fi15)
            v2 = (v2 * 127)
            _ = fi15
        }
        var v73 uint32 = 2147483647
        _ = fi14
        _ = v73
    }
    switch "hello" {
    case "\n\t":
        for fi74 := 0; fi74 < 7; fi74++ {
            switch 32768 {
            case -712:
                v75 := []int{729}
                _ = v75
            case -762:
                v76 := GenS8{F0: -128, F1: v2, F2: 953}
                _ = v76
            case 263:
                var v77 bool = true
                _ = v77
            }
            _ = fi74
        }
    }
    var v78 string = ""
    for fi79 := 0; fi79 < 6; fi79++ {
        v1 = v1
        switch ("foo bar" + v78) {
        case "":
            v80 := v3[1]
            _ = v80
        case "\n\t":
            var v81 int64 = -1
            _ = v81
        default:
            for fi82 := 0; fi82 < 8; fi82++ {
                v83 := []int{450, 9223372036854775807, 490, fi82}
                if (true || false) {
                    if len(v83) >= 2 {
                        v83 = v83[1:len(v83)]
                    }
                    v78 = "hello"
                    _ = len(v83)
                } else {
                    if true {
                        delete(v3, 73)
                        v2 += (v2 * (v2 ^ fi82))
                        fmt.Println(v1, fi79, fi82)
                    }
                }
                if (304 < 2) {
                    v84 := GenS9{F0: 65535}
                    switch (v2 * (fi79 | (fi82 - 908))) {
                    case -9223372036854775808:
                        var v85 int32 = v1
                        _ = v85
                    default:
                        v2 -= v2
                    }
                    _ = v84
                }
                _ = fi82
                _ = v83
            }
        }
        delete(v3, 480)
        _ = fi79
    }
    var v86 int = 0
    if (((v2 - v86) < (-(v2))) || true) {
        v87 := GenS9{F0: 64}
        var v88 bool = (v1 != v1)
        _ = v87
        _ = v88
    }
    v1 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v78:", v78)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v78
    _ = v86
}
