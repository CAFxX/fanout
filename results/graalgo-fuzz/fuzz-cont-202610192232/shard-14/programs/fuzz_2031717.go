package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 string }
type GenS5 struct { F0 uint64; F1 string; F2 int32 }
type GenS6 struct { F0 uint16; F1 int; F2 uint64 }
type GenS7 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 int, p1 uint32, p2 uint) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                var v9 uint32 = p1
                var v10 int8 = -109
                v11 := []int{fi7}
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := GenS1{F0: 1, F1: 188}
                v13 := []int{(fi7 + (fi7 & -703)), 32767, fi7, fi8}
                v14 := GenS2{F0: false}
                fmt.Println(p2)
                p0 += p0
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
        switch "hello" {
        case "a":
            fmt.Println(p0, p1)
        default:
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS2{F0: true}
                var v17 int32 = 42
                if (!true) {
                    p1 = (p1 - (p1 >> ((p2) & 63)))
                    if (32767 > 32767) {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            var v18 uint = p2
                            _ = v18
                        case "hello":
                            v17 -= (v17 - (v17 - v17))
                        case "a":
                            switch "\n\t" {
                            case "hello":
                                v19 := GenS1{F0: 109, F1: 1}
                                _ = v19
                            }
                        default:
                            if ((false && true) || false) {
                                var v20 uint16 = 836
                                v21 := []int{(-(fi15)), (fi7 - 981)}
                                if (true && ((32767 > -722) && true)) {
                                    fmt.Println(p1, p1)
                                    p1 = p1
                                } else {
                                    fmt.Println(fi15, v20, fi7)
                                }
                                _ = v20
                                _ = v21
                            }
                        }
                        v22 := []int{(p0 % ((p0) | 1))}
                        if len(v22) > 0 {
                            v22[0] = fi7
                        }
                        _ = v22
                    }
                    var v23 uint = p2
                    _ = v23
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi7
    }
    v24 := GenS3{F0: 42, F1: -32768}
    var v25 uint32 = p1
    v26 := GenS4{F0: 749, F1: 127, F2: ""}
    for fi27 := 0; fi27 < 16; fi27++ {
        v25 -= (v25 << (((p2 / (((p2 << ((p2) & 63))) | 1))) & 63))
        var v28 float32 = (((((3.14159 / 1.0) / 1.5) * float32((float64(0.0) * 0.0))) / (1.5 / 1e10)) + 0.0)
        _ = fi27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v26
    return 1
}
func genFn1(p0 int) uint16 {
    switch "hello" {
    case "\n\t":
        if false {
            var v29 int32 = 42
            _ = v29
        } else {
            var v30 int8 = -1
            var v31 uint8 = 2
            p0 = (p0 / ((127) | 1))
            _ = v30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 3; fi32++ {
            switch (p0 + (fi32 - fi32)) {
            case -32:
                switch "foo bar" {
                case "\n\t":
                    v33 := GenS2{F0: true}
                    _ = v33
                case "a":
                    var v34 int8 = -1
                    _ = v34
                case "foo bar":
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v36 := []int{fi32, -9223372036854775808}
                        _ = fi35
                        _ = v36
                    }
                default:
                    var v37 uint64 = 0
                    _ = v37
                }
            case -9223372036854775808:
                fmt.Println(p0, p0)
            }
            fmt.Println(p0)
            v38 := GenS5{F0: 2, F1: "hello", F2: -1}
            _ = fi32
            _ = v38
        }
    }
    var v39 bool = (0.1 > -2.25)
    _ = p0
    _ = v39
    return 486
}

func main() {
    var v1 uint16 = 255
    var v2 int = 682
    var v3 float64 = -2.25
    var v4 uint32 = 2
    var v5 uint64 = 610
    var v6 int = v2
    v40 := GenS6{F0: 2, F1: v2, F2: (v5 ^ v5)}
    v4 -= v4
    fmt.Println(v2)
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := []int{(^(v6)), v6}
        var v43 float64 = v3
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 uint8 = 127
    v2 -= (v2 ^ (v6 + (v2 - (v2 - (v6 - (v6 << ((65535) & 63)))))))
    switch ("foo bar" + "foo bar") {
    case "hello":
        v45 := GenS4{F0: 1, F1: 127, F2: ("" + "hello")}
        _ = v45
    case "\n\t":
        if false {
            for fi46 := 0; fi46 < 5; fi46++ {
                v47 := GenS5{F0: 766, F1: ("hello" + "hello"), F2: 1}
                if ((fi46 + 128) < (v6 << ((4294967295) & 63))) {
                    v48 := GenS3{F0: -128, F1: 290}
                    _ = v48
                }
                _ = fi46
                _ = v47
            }
            for fi49 := 0; fi49 < 20; fi49++ {
                var v50 uint = 2
                var v51 uint32 = 513
                _ = fi49
                _ = v50
                _ = v51
            }
            fmt.Println(v6, v5, v6)
        } else {
            v52 := []int{9223372036854775807, 4294967295, (-(v2))}
            _ = v52
        }
    case "foo bar":
        for fi53 := 0; fi53 < 13; fi53++ {
            var v54 uint32 = (v4 ^ 65535)
            var v55 int64 = 0
            _ = fi53
            _ = v54
            _ = v55
        }
    default:
        v56 := []int{281}
        _ = v56
    }
    v57 := []int{}
    v44 = (v44 - (^(v44)))
    if ((v1 & uint16(v5)) > (v1 >> ((1) & 63))) {
        switch "\n\t" {
        case "a":
            if (!false) {
                var v58 uint32 = uint32(v1)
                _ = v58
            } else {
                v4 = 949
                v59 := GenS6{F0: (v1 << ((765) & 63)), F1: (v6 % ((0) | 1)), F2: v5}
                for fi60 := 0; fi60 < 14; fi60++ {
                    var v61 float64 = (v3 - 0.1)
                    v1 = v1
                    _ = fi60
                    _ = v61
                }
                _ = v59
            }
        case "":
            if (((!(true || (false && (-778 == 42)))) || (!false)) && true) {
                fmt.Println(v4, v2, v44)
            }
        default:
            if len(v57) > 0 {
                v57[0] = (v6 ^ (-(v2)))
            }
        }
        var v62 float64 = v3
        _ = v62
    }
    _ = len(v57)
    v5 += (^(v5))
    for fi63 := 0; fi63 < 14; fi63++ {
        var v64 uint = 255
        if false {
            if ((0 <= 1) && (!(!true))) {
                var v65 int16 = 32767
                _ = v65
            } else {
                var v66 float64 = (v3 + v3)
                for fi67 := 0; fi67 < 8; fi67++ {
                    if len(v57) >= 2 {
                        v57 = v57[1:len(v57)]
                    }
                    _ = len(v57)
                    v2 = v6
                    _ = fi67
                }
                fmt.Println(v44, v2, v1)
                _ = v66
            }
            switch fi63 {
            case 9223372036854775807:
                fmt.Println(v2)
            case 849:
                if (42 < (v5 << ((uint(v4)) & 63))) {
                    v57 = append(v57, (fi63 % ((v2) | 1)))
                } else {
                    var v68 float32 = 1e10
                    var v69 int8 = 0
                    _ = v68
                    _ = v69
                }
            case -962:
                v44 += uint8(v64)
            }
            v70 := GenS4{F0: v64, F1: 126, F2: "a"}
            _ = v70
        } else {
            _ = len(v57)
            v64 = v64
        }
        v44 = v44
        _ = fi63
        _ = v64
    }
    for fi71 := 0; fi71 < 20; fi71++ {
        if (399 < 1) {
            for fi72 := 0; fi72 < 19; fi72++ {
                v1 = v1
                v73 := GenS3{F0: -128, F1: 1000}
                var v74 bool = (("a" + "") >= "hello")
                v57 = append(v57, -754)
                var v75 float64 = v3
                _ = fi72
                _ = v73
                _ = v74
                _ = v75
            }
        }
        v76 := GenS7{F0: 0, F1: 0, F2: 1}
        v77 := GenS4{F0: 18446744073709551615, F1: 128, F2: (("\n\t" + "hello") + "hello")}
        _ = fi71
        _ = v76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v5
    _ = v57
    _ = v6
}
