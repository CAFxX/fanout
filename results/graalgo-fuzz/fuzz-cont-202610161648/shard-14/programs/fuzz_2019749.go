package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int; F2 bool }
type GenS6 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 int64) int16 {
    v4 := []int{1, 127, 2147483647, 32768}
    if true {
        switch ("" + "foo bar") {
        case "hello":
            v5 := GenS1{F0: 255, F1: 0, F2: 42}
            _ = v5
        case "":
            switch "" {
            case "a":
                for fi6 := 0; fi6 < 4; fi6++ {
                    v7 := GenS1{F0: 42, F1: 42, F2: -128}
                    _ = fi6
                    _ = v7
                }
            case "foo bar":
                var v8 float64 = 1e10
                _ = v8
            case "hello":
                fmt.Println(p0)
            default:
                var v9 bool = true
                _ = v9
            }
        case "\n\t":
            var v10 int64 = (p0 & p0)
            _ = v10
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                switch fi11 {
                case -135:
                    p0 = -515
                case 2147483647:
                    v12 := GenS2{F0: 921}
                    _ = v12
                case -9223372036854775808:
                    fmt.Println(fi11, fi11)
                }
                var v13 float32 = 1e10
                var v14 int64 = p0
                _ = fi11
                _ = v13
                _ = v14
            }
        }
        v15 := GenS2{F0: 2}
        _ = v15
    } else {
        if true {
            p0 += 42
            v4 = append(v4, -1)
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int32 = 2147483647
                v19 := GenS1{F0: 2, F1: (fi16 - (fi17 / ((fi17) | 1))), F2: -128}
                _ = len(v4)
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS3{F0: 65535}
            _ = fi16
            _ = v20
        }
        p0 = (p0 * (p0 | 1))
    }
    var v21 uint32 = 224
    if (!(-2.25 >= (((0.1 * 1e10) - 0.1) - 1.5))) {
        var v22 int32 = -1000
        var v23 int = -9223372036854775808
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v21
    _ = v4
    return 1
}
func genFn1(p0 int64, p1 uint) int16 {
    p0 -= (p0 + -1)
    p1 = 0
    var v24 uint64 = 101
    _ = p0
    _ = p1
    _ = v24
    return 1000
}

func main() {
    var v1 int32 = -1
    var v2 bool = (-128 < -1)
    var v3 uint16 = 818
    var v25 uint16 = (^(v3))
    for fi26 := 0; fi26 < 14; fi26++ {
        if (99 <= -1) {
            v27 := GenS3{F0: v25}
            v28 := []int{}
            var v29 uint8 = 1
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 bool = v2
            _ = v30
        }
        v31 := []int{}
        v3 += (v25 * v3)
        _ = fi26
        _ = v31
    }
    if (false && ((-2.25 - -2.25) >= 1.5)) {
        var v32 uint8 = 1
        if true {
            fmt.Println(v1, v1)
            v1 -= v1
            v25 += v3
        } else {
            v33 := []int{0, 2147483647, -1}
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 uint32 = 42
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 20; fi36++ {
                v33 = append(v33, (^(fi36)))
                _ = fi36
            }
            _ = v33
        }
        v25 -= v25
        _ = v32
    } else {
        if (!(!true)) {
            var v37 int8 = 0
            _ = v37
        } else {
            v3 = 42
        }
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 int = (fi38 ^ -849)
            for fi40 := 0; fi40 < 12; fi40++ {
                if (100 != 100) {
                    v39 = fi38
                    if v2 {
                        if false {
                            var v41 int = v39
                            var v42 int = (^(fi40))
                            var v43 int32 = (v1 << ((65535) & 63))
                            v44 := []int{v42, fi38, (v39 + (v42 + fi40))}
                            v45 := GenS3{F0: v3}
                            if ((v41 | (-(v39))) <= v42) {
                                var v46 uint64 = 42
                                _ = v46
                            } else {
                                var v47 uint16 = 42
                                _ = v47
                            }
                            _ = len(v44)
                            if (true && v2) {
                                if v2 {
                                    v41 += (^(v41))
                                    v39 += (-(v41))
                                }
                                var v48 float32 = (1.5 - 1.5)
                                _ = v48
                            } else {
                                if v2 {
                                    var v49 int16 = int16(v42)
                                    v50 := GenS3{F0: v25}
                                    var v51 string = ("" + ("hello" + "hello"))
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                            }
                            var v52 int16 = -32768
                            _ = v41
                            _ = v42
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v52
                        }
                        if (42 < genFn1(42, 42)) {
                            var v53 float64 = 0.0
                            _ = v53
                        }
                        if v2 {
                            var v54 float32 = 0.0
                            var v55 bool = (!(!(!(!((v2 && (((v25 >> ((65535) & 63)) == (v3 * (v25 | v3))) && (65535 != 42))) || (v2 || (1e10 < 1e10)))))))
                            v1 = v1
                            _ = v54
                            _ = v55
                        } else {
                            var v56 uint16 = v3
                            v2 = v2
                            _ = v56
                        }
                        if (!true) {
                            v2 = ((v3 % (((v3 & (v25 << ((18446744073709551615) & 63)))) | 1)) == v25)
                            if (!false) {
                                var v57 int32 = v1
                                if ((3.14159 / 1.0) == (0.1 * -2.25)) {
                                    v1 -= v57
                                } else {
                                    var v58 int8 = 33
                                    v59 := GenS2{F0: 255}
                                    v60 := []int{-485, fi40, 32768}
                                    var v61 int = v39
                                    if len(v60) > 0 {
                                        v60[0] = fi38
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if (18446744073709551615 > 618) {
                                            switch fi40 {
                                            case -2147483648:
                                                var v62 string = "hello"
                                                _ = v62
                                            case 42:
                                                if v2 {
                                                    var v63 float64 = 1e10
                                                    _ = v63
                                                } else {
                                                    var v64 int16 = 42
                                                    fmt.Println(v25)
                                                    var v65 float64 = 0.1
                                                    fmt.Println(v2)
                                                    if v2 {
                                                        if v2 {
                                                            v2 = (!(fi40 < -9223372036854775808))
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                var v66 uint64 = 2
                                                                var v67 uint32 = 637
                                                                _ = v66
                                                                _ = v67
                                                            case "hello":
                                                                _ = len(v60)
                                                            default:
                                                                v68 := GenS4{F0: 65535}
                                                                _ = v68
                                                            }
                                                            var v69 int8 = (^(v58))
                                                            if v2 {
                                                                v65 = v65
                                                                fmt.Println(v65)
                                                            } else {
                                                                _ = len(v60)
                                                                fmt.Println(fi38, v61, v39)
                                                                if len(v60) >= 2 {
                                                                    v60 = v60[1:len(v60)]
                                                                }
                                                                var v70 uint16 = 255
                                                                v64 = v64
                                                                _ = len(v60)
                                                                if ((v69 % ((83) | 1)) < (-(v69))) {
                                                                    v3 -= v70
                                                                    var v71 float32 = 0.1
                                                                    if (!true) {
                                                                        var v72 uint32 = 901
                                                                        v1 = (v1 / (((^(v57))) | 1))
                                                                        if (!((v64 & (v64 >> ((65535) & 63))) < genFn0(743))) {
                                                                            var v73 int8 = v58
                                                                            v64 += genFn1(0, 662)
                                                                            var v74 int32 = v1
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                        _ = v72
                                                                    } else {
                                                                        var v75 string = "hello"
                                                                        if (!v2) {
                                                                            if len(v60) >= 2 {
                                                                                v60 = v60[1:len(v60)]
                                                                            }
                                                                        } else {
                                                                            if len(v60) > 0 {
                                                                                v60[0] = (fi38 / ((-353) | 1))
                                                                            }
                                                                        }
                                                                        v60 = append(v60, -2147483648)
                                                                        _ = v75
                                                                    }
                                                                    _ = v71
                                                                }
                                                                _ = v70
                                                            }
                                                            v65 = (v65 - v65)
                                                            _ = v69
                                                        }
                                                        v76 := GenS1{F0: 565, F1: (v61 | v39), F2: (^(v58))}
                                                        fmt.Println(v58, v65)
                                                        _ = v76
                                                    }
                                                    var v77 string = "hello"
                                                    _ = v64
                                                    _ = v65
                                                    _ = v77
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        switch 258 {
                                        case 343:
                                            if (false || v2) {
                                                v39 += 331
                                                if (v1 > (-(v57))) {
                                                    var v78 uint64 = 1
                                                    v78 = v78
                                                    var v79 float32 = float32((-2.25 * 3.14159))
                                                    _ = v78
                                                    _ = v79
                                                }
                                            } else {
                                                fmt.Println(v39, v58)
                                                v3 = 423
                                                v80 := GenS1{F0: v3, F1: fi38, F2: (v58 ^ 33)}
                                                v60 = append(v60, v39)
                                                if len(v60) > 0 {
                                                    v60[0] = 690
                                                }
                                                var v81 int8 = v58
                                                var v82 float64 = 1e10
                                                var v83 uint32 = 450
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            }
                                        case -2147483648:
                                            var v84 int64 = 431
                                            _ = v84
                                        }
                                    case "a":
                                        if len(v60) >= 2 {
                                            v60 = v60[1:len(v60)]
                                        }
                                    default:
                                        v85 := GenS5{F0: 258, F1: fi38, F2: (v3 == 32767)}
                                        _ = v85
                                    }
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = v57
                            }
                        } else {
                            v86 := []int{(fi38 * (v39 + (fi40 & v39))), 127, (-(fi38))}
                            fmt.Println(v2, v2)
                            _ = v86
                        }
                        var v87 uint64 = 18446744073709551615
                        v88 := []int{}
                        _ = v87
                        _ = v88
                    }
                    v89 := []int{-699, 2147483647}
                    _ = v89
                } else {
                    var v90 int64 = -608
                    v90 = 1
                    v91 := []int{(fi38 % ((v39) | 1)), fi40}
                    _ = v90
                    _ = v91
                }
                _ = fi40
            }
            fmt.Println(v25, v2, v2)
            _ = fi38
            _ = v39
        }
    }
    v3 -= (v3 + 388)
    if false {
        v3 = 32767
        var v92 int8 = 1
        v1 -= v1
        _ = v92
    } else {
        v3 = v25
    }
    if (3.14159 != (0.1 + -2.25)) {
        var v93 int = -560
        switch (v93 / (((v93 % ((128) | 1))) | 1)) {
        case 1:
            v94 := GenS5{F0: 1, F1: (v93 / ((-9223372036854775808) | 1)), F2: true}
            _ = v94
        case 2147483647:
            v25 = v3
        default:
            switch v93 {
            case -9223372036854775808:
                fmt.Println(v1, v3, v25)
            case -412:
                v95 := []int{(v93 + v93), (v93 ^ (v93 - v93))}
                _ = v95
            default:
                v96 := []int{}
                _ = v96
            }
        }
        _ = v93
    } else {
        for fi97 := 0; fi97 < 16; fi97++ {
            for fi98 := 0; fi98 < 15; fi98++ {
                fmt.Println(v25)
                switch ("hello" + "hello") {
                case "hello":
                    fmt.Println(v3, fi98)
                case "foo bar":
                    if (18446744073709551615 <= 18446744073709551615) {
                        v99 := []int{(fi98 >> ((46) & 63)), fi97, fi98}
                        v1 = v1
                        var v100 int64 = 42
                        var v101 uint16 = v25
                        _ = v100
                        _ = v101
                        _ = v99
                    }
                case "a":
                    v102 := GenS6{F0: ("\n\t" + "foo bar"), F1: 2, F2: 0}
                    _ = v102
                default:
                    switch fi97 {
                    case -669:
                        fmt.Println(v1, v1)
                    default:
                        var v103 float32 = (0.0 / ((0.1 * -2.25) * float32((3.14159 / -2.25))))
                        _ = v103
                    }
                }
                _ = fi98
            }
            v1 = v1
            for fi104 := 0; fi104 < 3; fi104++ {
                fmt.Println(v3, fi104, v3)
                var v105 uint8 = 128
                _ = fi104
                _ = v105
            }
            _ = fi97
        }
    }
    for fi106 := 0; fi106 < 13; fi106++ {
        fmt.Println(fi106)
        _ = fi106
    }
    var v107 int8 = -1
    fmt.Println(v2, v1, v25)
    fmt.Println(v107)
    for fi108 := 0; fi108 < 8; fi108++ {
        var v109 uint16 = 32767
        _ = fi108
        _ = v109
    }
    var v110 uint32 = 1
    var v111 uint = 557
    if (false || v2) {
        v1 = (v1 + v1)
    }
    var v112 uint8 = 145
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v107
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v25
    _ = v3
}
