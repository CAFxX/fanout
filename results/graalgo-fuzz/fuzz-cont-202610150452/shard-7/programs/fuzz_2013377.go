package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS6 struct { F0 int32; F1 int16; F2 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint8; F2 int8 }
type GenS9 struct { F0 int64; F1 int16; F2 int8 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint32 }
func genFn0() string {
    if false {
        switch ("" + "a") {
        case "a":
            var v7 uint = 42
            _ = v7
        case "":
            var v8 int32 = 66
            _ = v8
        case "foo bar":
            var v9 int16 = 1000
            _ = v9
        }
        v10 := []int{127, 32768, 2147483647}
        v11 := GenS1{F0: 32767, F1: "", F2: -2147483648}
        _ = v10
        _ = v11
    } else {
        switch "foo bar" {
        case "":
            var v12 int64 = 22
            _ = v12
        case "\n\t":
            var v13 float32 = -2.25
            _ = v13
        default:
            var v14 uint8 = 255
            _ = v14
        }
    }
    var v15 string = "foo bar"
    if ((false && true) || false) {
        var v16 string = "hello"
        for fi17 := 0; fi17 < 12; fi17++ {
            switch 42 {
            case 136:
                var v18 uint32 = 2147483648
                _ = v18
            case 9223372036854775807:
                if (v16 <= "") {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v15 = (v15 + "a")
                        _ = fi19
                    }
                    v20 := GenS1{F0: 32767, F1: ("hello" + "\n\t"), F2: 42}
                    _ = v20
                }
            default:
                v21 := []int{}
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                switch (fi17 << ((65535) & 63)) {
                case -507:
                    v15 = (v16 + ("" + "foo bar"))
                default:
                    switch fi17 {
                    case 42:
                        var v23 int = fi17
                        _ = v23
                    case 127:
                        v15 = v15
                    case 948:
                        v15 = "hello"
                    }
                }
                var v24 uint64 = 0
                v24 = v24
                _ = fi22
                _ = v24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 uint = 2
                fmt.Println(fi17, v15, v15)
                if true {
                    fmt.Println(v16, v16)
                    v27 := []int{127, fi25, (^(fi25)), (fi17 << (((v26 ^ (v26 | (^(v26))))) & 63))}
                    _ = len(v27)
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            _ = fi17
        }
        _ = v16
    }
    var v28 float32 = 0.0
    _ = v15
    _ = v28
    return v15
}
func genFn1() int32 {
    var v29 int8 = -10
    for fi30 := 0; fi30 < 18; fi30++ {
        if true {
            var v31 float64 = 1e10
            if true {
                if true {
                    switch "hello" {
                    case "hello":
                        v32 := GenS2{F0: 1, F1: 241}
                        _ = v32
                    case "a":
                        fmt.Println(fi30)
                    }
                    v33 := []int{-9223372036854775808, fi30}
                    v31 += v31
                    _ = v33
                } else {
                    var v34 string = "hello"
                    if (false || ((float64((1.5 / 1.0)) / (v31 * ((float64(0.0) / v31) - v31))) >= ((v31 - (((v31 + v31) / (1.5 / 1.0)) / v31)) - v31))) {
                        if true {
                            if (false && (v29 >= v29)) {
                                if (!false) {
                                    fmt.Println(fi30, v34, v29)
                                    var v35 int8 = -58
                                    _ = v35
                                } else {
                                    v36 := []int{(fi30 & fi30), fi30}
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 18; fi37++ {
                                    var v38 int = fi37
                                    _ = fi37
                                    _ = v38
                                }
                            }
                        }
                        v39 := []int{fi30}
                        _ = len(v39)
                        _ = v39
                    } else {
                        v31 += float64(1.5)
                        var v40 uint32 = 65535
                        _ = v40
                    }
                    switch 0 {
                    case 871:
                        switch "" {
                        case "a":
                            switch 1 {
                            case 9223372036854775807:
                                var v41 uint32 = 42
                                _ = v41
                            case -401:
                                v42 := []int{}
                                _ = v42
                            default:
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    v31 -= 1e10
                                    v44 := []int{}
                                    _ = fi43
                                    _ = v44
                                }
                            }
                        }
                    case 9223372036854775807:
                        v31 = 3.14159
                    }
                    _ = v34
                }
                v45 := GenS3{F0: 18446744073709551615, F1: 1, F2: v29}
                _ = v45
            }
            switch "a" {
            case "hello":
                fmt.Println(fi30, fi30, v31)
            case "":
                fmt.Println(v31, v29, v31)
            }
            _ = v31
        } else {
            for fi46 := 0; fi46 < 11; fi46++ {
                if (1 == 1) {
                    v47 := GenS1{F0: 42, F1: "hello", F2: 577}
                    var v48 string = "hello"
                    _ = v47
                    _ = v48
                } else {
                    v29 += 107
                    v49 := GenS2{F0: 227, F1: 128}
                    v50 := GenS2{F0: 134, F1: 128}
                    _ = v49
                    _ = v50
                }
                var v51 int32 = 42
                _ = fi46
                _ = v51
            }
            v52 := GenS4{F0: (!false)}
            _ = v52
        }
        v53 := []int{}
        v54 := GenS5{F0: 1, F1: int32(v29), F2: 8}
        _ = fi30
        _ = v53
        _ = v54
    }
    v55 := GenS6{F0: 42, F1: -881, F2: -973}
    switch ((("foo bar" + "foo bar") + "foo bar") + ("\n\t" + "a")) {
    case "\n\t":
        v56 := GenS5{F0: 4294967295, F1: 2147483647, F2: 128}
        _ = v56
    case "hello":
        if false {
            switch 0 {
            case 32767:
                v57 := GenS7{F0: 42}
                _ = v57
            default:
                v29 -= v29
            }
        }
    }
    _ = v29
    _ = v55
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 611
    var v3 uint16 = 32767
    v4 := []uint32{1, 2147483648, 306}
    var v5 uint8 = 118
    var v6 float64 = (1e10 + -2.25)
    v3 = 2
    var v58 float64 = v6
    var v59 int64 = 9223372036854775807
    v1[65535] = "a"
    switch 128 {
    case -522:
        var v60 uint8 = (v5 | v5)
        _ = v60
    case 0:
        switch "hello" {
        case "foo bar":
            var v61 uint = 41
            _ = v61
        default:
            v58 = 1e10
        }
    case 42:
        v59 += v59
    }
    switch -1 {
    case 32767:
        if (v2 != v2) {
            switch "" {
            case "hello":
                if (794 >= 1) {
                    var v62 float32 = 0.1
                    _ = v62
                } else {
                    switch 9223372036854775807 {
                    case 2147483647:
                        v1[65535] = ("hello" + "hello")
                    case 771:
                        for fi63 := 0; fi63 < 19; fi63++ {
                            v5 = (v5 & 2)
                            var v64 float32 = 1e10
                            for fi65 := 0; fi65 < 9; fi65++ {
                                if ((!(false && false)) && ((0.1 - 0.0) == 3.14159)) {
                                    v6 -= 1e10
                                    v6 = v58
                                    var v66 int64 = 819
                                    _ = v66
                                } else {
                                    v67 := GenS1{F0: v3, F1: "foo bar", F2: (v2 & (v2 - (v2 / (((v2 + v2)) | 1))))}
                                    _ = len(v4)
                                    _ = v67
                                }
                                if false {
                                    var v68 float32 = v64
                                    _ = v68
                                }
                                var v69 uint8 = (v5 % (((v5 % ((v5) | 1))) | 1))
                                _ = fi65
                                _ = v69
                            }
                            _ = fi63
                            _ = v64
                        }
                    }
                }
            case "\n\t":
                fmt.Println(v59, v59, v6)
            default:
                v59 -= v59
            }
        }
    default:
        if true {
            switch "hello" {
            case "a":
                fmt.Println(v3)
            case "foo bar":
                v70 := GenS8{F0: 9223372036854775807, F1: v5, F2: -83}
                _ = v70
            default:
                var v71 int32 = genFn1()
                _ = v71
            }
            fmt.Println(v3, v3, v3)
            var v72 float32 = float32((v58 * (3.14159 - (0.1 / (((v6 / v58) + v6) / 3.14159)))))
            _ = v72
        } else {
            v5 = 168
        }
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        var v74 bool = true
        _ = fi73
        _ = v74
    }
    for fi75 := 0; fi75 < 11; fi75++ {
        v3 = v3
        if (false || ((false && (!((!false) || (((v5 <= v5) && (false && true)) || ((1e10 - v6) <= 0.1))))) || (57 != -128))) {
            var v76 uint8 = 221
            _ = len(v4)
            v77 := v1[2]
            _ = v76
            _ = v77
        } else {
            var v78 int8 = -103
            var v79 float32 = 3.14159
            _ = v78
            _ = v79
        }
        _ = fi75
    }
    switch 1 {
    case -2147483648:
        v80 := GenS2{F0: v5, F1: (v5 >> ((2) & 63))}
        _ = v80
    case 1:
        if ((v5 & 127) != v5) {
            if false {
                var v81 uint32 = 209
                v82 := GenS9{F0: -38, F1: 1, F2: 23}
                var v83 int8 = int8(v59)
                _ = v81
                _ = v82
                _ = v83
            } else {
                if false {
                    var v84 int8 = 0
                    v5 = 193
                    v85 := v1[uint32(v3)]
                    _ = v84
                    _ = v85
                } else {
                    if true {
                        if true {
                            v3 = (v3 >> ((255) & 63))
                            fmt.Println(v2)
                        } else {
                            _ = len(v1)
                        }
                        _ = len(v4)
                    } else {
                        v5 = 128
                        v86 := GenS10{F0: v2, F1: 2}
                        _ = v86
                    }
                    v87 := GenS11{F0: 0}
                    v88 := GenS9{F0: (v59 + 0), F1: 0, F2: 1}
                    _ = v87
                    _ = v88
                }
                for fi89 := 0; fi89 < 12; fi89++ {
                    var v90 uint64 = 2
                    _ = fi89
                    _ = v90
                }
                v91 := GenS1{F0: (v3 % ((v3) | 1)), F1: "", F2: genFn1()}
                _ = v91
            }
        } else {
            var v92 uint32 = 457
            v93 := GenS7{F0: 2147483647}
            var v94 int64 = int64(v3)
            _ = v92
            _ = v93
            _ = v94
        }
    default:
        v6 = -2.25
    }
    if (1000 < 32767) {
        v59 = 0
    }
    v2 = (^(v2))
    v3 = (v3 % ((65535) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
