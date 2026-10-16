package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint32; F2 uint8 }
func genFn0() uint32 {
    var v4 int64 = 680
    if true {
        var v5 bool = false
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int8 = -128
            var v8 int8 = v7
            var v9 string = ("\n\t" + ("foo bar" + ("\n\t" + "foo bar")))
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = v5
    } else {
        if (!true) {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 uint8 = 168
                v11 = v11
                v11 += 128
                _ = fi10
                _ = v11
            }
        } else {
            if true {
                fmt.Println(v4)
                var v12 int16 = -99
                _ = v12
            }
        }
        fmt.Println(v4)
        var v13 bool = true
        _ = v13
    }
    _ = v4
    return 4294967295
}
func genFn1() uint16 {
    v14 := GenS1{F0: -709}
    var v15 string = "\n\t"
    v15 = v15
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS2{F0: 18446744073709551615}
        var v18 int8 = 127
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = v14
    _ = v15
    return 65535
}

func main() {
    var v1 int32 = -708
    var v2 float32 = 1e10
    var v3 uint16 = 42
    v1 = v1
    if (90 >= -24) {
        v2 = 3.14159
        v3 -= (v3 ^ 65535)
        var v19 float64 = (-2.25 * -2.25)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = 294
            _ = fi20
            _ = v21
        }
        fmt.Println(v2)
    }
    if false {
        if (-128 < 100) {
            v22 := []int{313, 9223372036854775807, 32768}
            if ((v3 + 556) == v3) {
                v23 := GenS2{F0: 595}
                switch 532 {
                case 32767:
                    var v24 uint64 = 4294967295
                    _ = v24
                case 127:
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                case 0:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 uint64 = 18446744073709551615
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 int16 = 1000
                            if (-128 != 1) {
                                if false {
                                    if (127 >= 24) {
                                        fmt.Println(fi25)
                                    } else {
                                        v26 -= (v26 & (v26 + v26))
                                    }
                                    v29 := GenS2{F0: 312}
                                    v22 = append(v22, 32768)
                                    _ = v29
                                } else {
                                    switch fi27 {
                                    case 943:
                                        v30 := GenS3{F0: fi27}
                                        _ = v30
                                    case 0:
                                        if (0 >= v28) {
                                            var v31 int = 4294967295
                                            var v32 uint16 = (v3 * v3)
                                            v33 := GenS1{F0: (fi25 % (((v31 + (fi27 & fi25))) | 1))}
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := GenS2{F0: 944}
                                        var v35 bool = false
                                        _ = v34
                                        _ = v35
                                    default:
                                        var v36 string = ""
                                        _ = v36
                                    }
                                    if ((("foo bar" + "\n\t") < ("foo bar" + ("\n\t" + ("hello" + ("" + "\n\t"))))) && true) {
                                        fmt.Println(v28)
                                    } else {
                                        v37 := GenS1{F0: fi25}
                                        v38 := GenS2{F0: (v26 % ((12) | 1))}
                                        _ = v37
                                        _ = v38
                                    }
                                    v39 := GenS4{F0: -265, F1: 2147483648}
                                    _ = v39
                                }
                                if true {
                                    var v40 uint8 = 18
                                    _ = len(v22)
                                    _ = v40
                                } else {
                                    if true {
                                        v41 := GenS1{F0: fi27}
                                        v2 = 0.1
                                        v26 = (^(v26))
                                        _ = v41
                                    }
                                    v42 := GenS3{F0: fi25}
                                    var v43 uint8 = uint8(v26)
                                    _ = v42
                                    _ = v43
                                }
                                if (2147483648 >= 0) {
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                } else {
                                    var v44 bool = (true || (1.5 <= ((((-2.25 / 3.14159) - (0.0 + (float64(v2) / 0.1))) / 3.14159) * (1e10 * 0.1))))
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = v44
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        if len(v22) > 0 {
                            v22[0] = 0
                        }
                        _ = fi25
                        _ = v26
                    }
                default:
                    var v45 float64 = 0.0
                    _ = v45
                }
                _ = v23
            } else {
                var v46 float64 = 0.1
                _ = v46
            }
            _ = v22
        }
        var v47 uint = 1
        for fi48 := 0; fi48 < 4; fi48++ {
            v49 := []int{fi48, fi48, (fi48 | fi48), fi48}
            v49 = append(v49, fi48)
            v2 += 1.5
            _ = fi48
            _ = v49
        }
        _ = v47
    } else {
        switch (("\n\t" + "a") + "\n\t") {
        case "a":
            switch 4294967295 {
            case -2147483648:
                var v50 float64 = 1.5
                _ = v50
            case 128:
                for fi51 := 0; fi51 < 14; fi51++ {
                    v52 := GenS4{F0: -789, F1: 2}
                    var v53 uint32 = 4294967295
                    fmt.Println(v3)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            default:
                var v54 int64 = 42
                _ = v54
            }
        default:
            var v55 int64 = 0
            _ = v55
        }
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 string = ("hello" + "hello")
            if ((!false) || true) {
                fmt.Println(v57)
                switch -243 {
                case -927:
                    var v58 string = ""
                    _ = v58
                case 0:
                    fmt.Println(v2, fi56)
                default:
                    fmt.Println(v3)
                }
                var v59 uint = uint(v3)
                _ = v59
            }
            var v60 uint32 = genFn0()
            _ = fi56
            _ = v57
            _ = v60
        }
    }
    switch ("hello" + "") {
    case "a":
        switch -341 {
        case 2147483647:
            if ((-1 < 42) || true) {
                for fi61 := 0; fi61 < 20; fi61++ {
                    v62 := GenS2{F0: 139}
                    v63 := GenS5{F0: 0}
                    var v64 int32 = -2147483648
                    _ = fi61
                    _ = v62
                    _ = v63
                    _ = v64
                }
                var v65 uint64 = 629
                _ = v65
            } else {
                v66 := GenS5{F0: -1}
                for fi67 := 0; fi67 < 17; fi67++ {
                    v68 := []int{}
                    var v69 uint8 = 1
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = v66
            }
        case 32767:
            var v70 int = -9223372036854775808
            _ = v70
        case 702:
            switch "\n\t" {
            case "hello":
                var v71 float32 = v2
                _ = v71
            case "a":
                if false {
                    for fi72 := 0; fi72 < 12; fi72++ {
                        v73 := []int{fi72, 1, 864, fi72}
                        _ = fi72
                        _ = v73
                    }
                }
            default:
                if false {
                    v3 -= v3
                } else {
                    v74 := GenS3{F0: 32767}
                    fmt.Println(v2, v1)
                    v75 := GenS6{F0: 127}
                    _ = v74
                    _ = v75
                }
            }
        }
    }
    for fi76 := 0; fi76 < 18; fi76++ {
        for fi77 := 0; fi77 < 1; fi77++ {
            v78 := GenS1{F0: 171}
            v79 := []int{890, -1, 9223372036854775807}
            v80 := GenS3{F0: fi76}
            _ = fi77
            _ = v78
            _ = v79
            _ = v80
        }
        _ = fi76
    }
    v1 += v1
    for fi81 := 0; fi81 < 7; fi81++ {
        if ((fi81 << ((65535) & 63)) <= fi81) {
            v3 = v3
        } else {
            v2 -= v2
            fmt.Println(v3, v1, v1)
        }
        v2 = ((1.5 / (((v2 + (v2 - 1e10)) * v2) + v2)) * 0.1)
        _ = fi81
    }
    fmt.Println(v3)
    v1 = v1
    v82 := GenS7{F0: 314, F1: 289, F2: 42}
    fmt.Println(v3, v2, v3)
    v1 = v1
    v83 := GenS8{F0: 65535}
    if true {
        var v84 string = (("hello" + "a") + ("foo bar" + "\n\t"))
        switch (("a" + "hello") + "hello") {
        case "\n\t":
            var v85 int = 754
            _ = v85
        case "hello":
            var v86 int64 = 9223372036854775807
            _ = v86
        case "a":
            if (4294967295 != 47) {
                if (v1 >= (^(v1))) {
                    v87 := GenS9{F0: 319, F1: 569}
                    var v88 uint = 18446744073709551615
                    for fi89 := 0; fi89 < 11; fi89++ {
                        switch v84 {
                        case "a":
                            var v90 int = (fi89 - (-(fi89)))
                            _ = v90
                        case "\n\t":
                            v91 := GenS8{F0: v88}
                            _ = v91
                        case "foo bar":
                            var v92 int8 = 100
                            _ = v92
                        default:
                            v93 := []int{(fi89 << ((726) & 63)), 42}
                            _ = v93
                        }
                        if false {
                            switch 127 {
                            case 9223372036854775807:
                                v88 += (^(v88))
                            }
                        } else {
                            fmt.Println(v88, v2)
                            switch v84 {
                            case "a":
                                v94 := []int{}
                                _ = v94
                            case "":
                                switch fi89 {
                                case 128:
                                    var v95 uint64 = 154
                                    _ = v95
                                case 42:
                                    var v96 int32 = (-(v1))
                                    _ = v96
                                case 32767:
                                    switch "hello" {
                                    case "":
                                        v88 = (v88 * v88)
                                    }
                                }
                            default:
                                fmt.Println(v88, v2)
                            }
                            v3 = (v3 + (v3 ^ (v3 + (v3 * v3))))
                        }
                        var v97 uint64 = 706
                        _ = fi89
                        _ = v97
                    }
                    _ = v87
                    _ = v88
                }
                if ((((v84 + v84) + v84) + "") > (("foo bar" + v84) + "\n\t")) {
                    v2 -= v2
                    v98 := GenS1{F0: 739}
                    v99 := GenS10{F0: -128, F1: 4294967295, F2: 83}
                    _ = v98
                    _ = v99
                }
            }
        }
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v82
    _ = v83
}
