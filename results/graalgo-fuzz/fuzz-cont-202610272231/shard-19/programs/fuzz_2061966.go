package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 int; F2 int8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0() int {
    var v6 int64 = 42
    v7 := []int{-979, -1, 9223372036854775807, 127}
    _ = v6
    _ = v7
    return 32767
}
func genFn1() string {
    for fi8 := 0; fi8 < 8; fi8++ {
        switch 811 {
        case -9223372036854775808:
            v9 := GenS1{F0: 4294967295}
            _ = v9
        }
        var v10 float32 = 0.0
        _ = fi8
        _ = v10
    }
    switch ("\n\t" + "") {
    case "hello":
        if false {
            v11 := []int{127, -1, -152, -598}
            var v12 uint8 = 127
            _ = v11
            _ = v12
        } else {
            if (2 >= 42) {
                switch 488 {
                case 128:
                    for fi13 := 0; fi13 < 4; fi13++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v14 int8 = 57
                            _ = v14
                        case "hello":
                            if (1 <= 42) {
                                var v15 int64 = -438
                                _ = v15
                            }
                        case "":
                            v16 := []int{(fi13 * 690)}
                            _ = v16
                        default:
                            v17 := GenS2{F0: 4294967295, F1: (fi13 >> ((1) & 63))}
                            _ = v17
                        }
                        _ = fi13
                    }
                case 0:
                    var v18 int16 = 1000
                    _ = v18
                }
                var v19 int = 42
                switch "hello" {
                case "":
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v21 := GenS2{F0: 535, F1: 2147483647}
                        _ = fi20
                        _ = v21
                    }
                case "hello":
                    var v22 uint16 = 65535
                    _ = v22
                case "foo bar":
                    v23 := GenS3{F0: 9223372036854775807, F1: v19, F2: 1}
                    _ = v23
                default:
                    var v24 bool = false
                    _ = v24
                }
                _ = v19
            } else {
                if true {
                    var v25 bool = false
                    _ = v25
                } else {
                    var v26 uint16 = 96
                    v26 = 2
                    v26 = 42
                    _ = v26
                }
            }
            v27 := []int{genFn0(), 4294967295, -101}
            _ = v27
        }
    case "":
        var v28 int = -670
        _ = v28
    case "foo bar":
        if true {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 float32 = (1e10 / float32(3.14159))
                _ = fi29
                _ = v30
            }
            var v31 uint = 948
            v31 -= v31
            _ = v31
        }
    }
    v32 := GenS2{F0: 18446744073709551615, F1: 0}
    v33 := []int{300, 128, 128, genFn0()}
    _ = v32
    _ = v33
    return ""
}

func main() {
    var v1 int16 = 0
    v2 := []int{}
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    v5 := make(map[int]bool)
    if (!false) {
        var v34 uint32 = 42
        _ = v34
    } else {
        _ = len(v2)
        _ = len(v5)
    }
    if (-1 >= v4) {
        switch 42 {
        case 1:
            fmt.Println(v4)
        default:
            var v35 int64 = 0
            _ = v35
        }
    }
    if (-128 < v4) {
        v36 := GenS1{F0: 1}
        var v37 int64 = 1
        v4 = (v4 ^ v4)
        _ = v36
        _ = v37
    }
    switch -404 {
    case 32767:
        fmt.Println(v4)
    case 884:
        if false {
            v1 = (^(v1))
            fmt.Println(v1)
            var v38 string = (("a" + "foo bar") + "hello")
            _ = v38
        } else {
            v4 = (v4 << ((42) & 63))
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := GenS3{F0: -313, F1: (^(fi39)), F2: 1}
        var v41 int8 = (^(v4))
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint32 = 0
    v43 := GenS1{F0: 42}
    if len(v2) > 0 {
        v2[0] = 127
    }
    v4 = v4
    if false {
        fmt.Println(v1, v4)
    } else {
        v44 := GenS1{F0: 4294967295}
        if (1e10 >= 3.14159) {
            v1 -= (v1 & (v1 ^ -1))
            v45 := GenS4{F0: 4294967295, F1: (-(v4))}
            _ = v45
        } else {
            for fi46 := 0; fi46 < 2; fi46++ {
                if true {
                    var v47 uint64 = 634
                    switch "foo bar" {
                    case "\n\t":
                        switch -1 {
                        case 42:
                            for fi48 := 0; fi48 < 12; fi48++ {
                                if (false || ((2 >= 2) || (2 <= 2))) {
                                    var v49 uint = 42
                                    if false {
                                        v50 := GenS4{F0: fi46, F1: v4}
                                        var v51 bool = ("\n\t" < "a")
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v49
                                }
                                if len(v2) > 0 {
                                    v2[0] = (fi46 + genFn0())
                                }
                                _ = fi48
                            }
                        case 4294967295:
                            var v52 float64 = 1.5
                            _ = v52
                        case 0:
                            fmt.Println(fi46, fi46, v4)
                        default:
                            for fi53 := 0; fi53 < 14; fi53++ {
                                var v54 uint32 = v42
                                _ = fi53
                                _ = v54
                            }
                        }
                    case "":
                        var v55 uint = 2
                        _ = v55
                    case "hello":
                        for fi56 := 0; fi56 < 7; fi56++ {
                            fmt.Println(v47, v47, fi46)
                            delete(v5, (fi46 - 4294967295))
                            var v57 uint8 = 130
                            var v58 uint64 = (v47 - v47)
                            var v59 string = (("" + "foo bar") + (((("a" + "") + (("hello" + "a") + (("\n\t" + "\n\t") + ""))) + "") + "foo bar"))
                            fmt.Println(v1, v47, v59)
                            _ = fi56
                            _ = v57
                            _ = v58
                            _ = v59
                        }
                    default:
                        if (1 == 942) {
                            var v60 float64 = 1.5
                            if (((("hello" + "foo bar") + "foo bar") + (("" + "\n\t") + "a")) > "a") {
                                var v61 uint32 = v42
                                var v62 uint32 = (v61 | (v61 - v61))
                                for fi63 := 0; fi63 < 1; fi63++ {
                                    if (3.14159 == 1e10) {
                                        v47 = v47
                                    }
                                    _ = fi63
                                }
                                _ = v61
                                _ = v62
                            }
                            fmt.Println(v60, fi46, v4)
                            _ = v60
                        }
                    }
                    _ = v47
                } else {
                    var v64 bool = true
                    var v65 uint64 = 4294967295
                    _ = v64
                    _ = v65
                }
                v66 := v3[(v42 / (((v42 - v42)) | 1))]
                v1 += v1
                _ = fi46
                _ = v66
            }
            v42 = 367
        }
        if true {
            v42 += v42
            if (!true) {
                var v67 string = "a"
                for fi68 := 0; fi68 < 7; fi68++ {
                    delete(v3, v42)
                    var v69 uint64 = 2
                    _ = fi68
                    _ = v69
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v67
            }
            if false {
                v4 -= -128
                v5[-2147483648] = (true || true)
            }
        } else {
            if (true || (false && (((2147483648 == (v42 / (((v42 + (v42 | v42))) | 1))) || (false && (42 <= 4294967295))) && (!false)))) {
                v1 += (v1 % ((550) | 1))
            }
            var v70 int64 = 59
            v71 := GenS5{F0: (v42 - (v42 % ((4294967295) | 1)))}
            _ = v70
            _ = v71
        }
        _ = v44
    }
    v1 += (v1 - v1)
    for fi72 := 0; fi72 < 16; fi72++ {
        v4 -= -127
        _ = fi72
    }
    for fi73 := 0; fi73 < 16; fi73++ {
        delete(v3, v42)
        var v74 int16 = 193
        for fi75 := 0; fi75 < 8; fi75++ {
            var v76 int8 = 15
            var v77 uint8 = 255
            _ = fi75
            _ = v76
            _ = v77
        }
        _ = fi73
        _ = v74
    }
    fmt.Println(v4)
    var v78 int8 = 0
    for fi79 := 0; fi79 < 1; fi79++ {
        v80 := GenS4{F0: (fi79 >> ((239) & 63)), F1: 97}
        fmt.Println(v78, v42, fi79)
        _ = fi79
        _ = v80
    }
    for fi81 := 0; fi81 < 15; fi81++ {
        v42 = (v42 << ((2) & 63))
        v82 := v3[0]
        _ = fi81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v78
}
