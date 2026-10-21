package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int32 }

func main() {
    var v1 uint = 42
    var v2 int = 127
    var v3 int = (v2 & v2)
    var v4 int = (v2 % ((-9223372036854775808) | 1))
    v5 := make(map[string]bool)
    var v6 int = (^(v3))
    switch "a" {
    case "\n\t":
        if false {
            switch -9223372036854775808 {
            case 32768:
                var v7 int8 = 28
                _ = v7
            case 2147483647:
                var v8 int32 = 42
                _ = v8
            default:
                v1 -= 174
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                if (1000 <= -643) {
                    var v10 string = ("foo bar" + "a")
                    v6 = 128
                    v11 := v5[v10]
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            var v12 int64 = 151
            _ = v12
        } else {
            if false {
                v13 := GenS1{F0: 745, F1: 2147483647, F2: 9223372036854775807}
                for fi14 := 0; fi14 < 17; fi14++ {
                    v5[(("a" + "\n\t") + "")] = true
                    var v15 int64 = -9223372036854775808
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    if (!((-40 != -889) || ((2 == v1) || false))) {
                        v17 := GenS1{F0: 678, F1: -57, F2: -9223372036854775808}
                        v18 := []int{v4, v3}
                        switch 1 {
                        case 4294967295:
                            switch (v6 + (v4 >> ((v1) & 63))) {
                            case -170:
                                var v19 string = ("" + ("\n\t" + ("foo bar" + "\n\t")))
                                _ = v19
                            default:
                                if (42 < 109) {
                                    switch 42 {
                                    case -1:
                                        for fi20 := 0; fi20 < 6; fi20++ {
                                            v21 := GenS2{F0: ("\n\t" + "\n\t")}
                                            var v22 uint64 = 1
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                    case -528:
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch 128 {
                                            case -2147483648:
                                                var v23 int = (v3 * fi16)
                                                _ = v23
                                            case -8:
                                                var v24 int32 = 1
                                                _ = v24
                                            }
                                        case "hello":
                                            var v25 int64 = 1
                                            _ = v25
                                        case "\n\t":
                                            var v26 int = 128
                                            _ = v26
                                        }
                                    case 127:
                                        v4 -= fi16
                                    default:
                                        var v27 int32 = 0
                                        _ = v27
                                    }
                                    var v28 int32 = -2147483648
                                    _ = v28
                                }
                            }
                        case 32767:
                            if true {
                                v2 = v4
                                var v29 int32 = 0
                                _ = v29
                            } else {
                                fmt.Println(v2)
                                if (!(false || ((!(!(true && ((3.14159 >= (-2.25 / (float32(0.1) / ((1.5 / float32(1.5)) + 1.5)))) && false)))) && (1 > 0)))) {
                                    v30 := GenS2{F0: ""}
                                    _ = v30
                                } else {
                                    fmt.Println(fi16, v4)
                                    v4 -= (v6 >> ((v1) & 63))
                                    var v31 uint16 = 681
                                    _ = v31
                                }
                                _ = len(v5)
                            }
                        case -101:
                            for fi32 := 0; fi32 < 4; fi32++ {
                                if ((!((1.5 + -2.25) >= float64((1.5 + -2.25)))) || (!((0.0 != 1e10) || false))) {
                                    var v33 int32 = -2147483648
                                    delete(v5, "")
                                    _ = v33
                                } else {
                                    v34 := GenS3{F0: -1, F1: 65535, F2: 70}
                                    v4 = 127
                                    _ = v34
                                }
                                var v35 uint64 = 1
                                v4 = v4
                                _ = fi32
                                _ = v35
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v2, v6)
                        v2 -= 32768
                        for fi36 := 0; fi36 < 3; fi36++ {
                            v37 := GenS4{F0: -1}
                            switch "" {
                            case "\n\t":
                                if (!false) {
                                    fmt.Println(fi36, v2, v2)
                                } else {
                                    var v38 int32 = -662
                                    var v39 uint = (v1 ^ v1)
                                    _ = v38
                                    _ = v39
                                }
                                var v40 float64 = ((0.1 * 1.5) * 1.5)
                                _ = v40
                            case "foo bar":
                                _ = len(v5)
                            }
                            if (!true) {
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v2, v2, v2)
                                }
                                var v41 uint64 = 0
                                v1 = v1
                                _ = v41
                            }
                            v42 := v5["a"]
                            _ = fi36
                            _ = v37
                            _ = v42
                        }
                    }
                    if (!true) {
                        v1 -= 255
                    }
                    var v43 int32 = 42
                    _ = fi16
                    _ = v43
                }
                _ = v13
            } else {
                if (!true) {
                    if true {
                        v4 = v3
                        if true {
                            v44 := GenS4{F0: 93}
                            _ = v44
                        } else {
                            v3 = v3
                            v45 := []int{}
                            _ = v45
                        }
                    }
                    var v46 int = v3
                    var v47 int = (v3 ^ v4)
                    _ = v46
                    _ = v47
                }
            }
        }
    case "a":
        v48 := GenS5{F0: 100}
        _ = v48
    default:
        if false {
            v49 := []int{v2, (v4 % ((v6) | 1))}
            _ = v49
        } else {
            var v50 uint = 65535
            fmt.Println(v2, v1, v6)
            v51 := GenS6{F0: -32768, F1: v50, F2: 32767}
            _ = v50
            _ = v51
        }
    }
    fmt.Println(v2, v2)
    v4 = v2
    var v52 int8 = 0
    fmt.Println(v1, v3, v1)
    for fi53 := 0; fi53 < 3; fi53++ {
        fmt.Println(v3, fi53, fi53)
        var v54 int8 = (v52 / ((v52) | 1))
        for fi55 := 0; fi55 < 19; fi55++ {
            v5["\n\t"] = true
            _ = fi55
        }
        _ = fi53
        _ = v54
    }
    if (548 > 894) {
        var v56 int8 = v52
        _ = v56
    } else {
        fmt.Println(v6, v4, v4)
    }
    for fi57 := 0; fi57 < 16; fi57++ {
        var v58 float32 = 1.5
        v59 := GenS7{F0: 255}
        _ = fi57
        _ = v58
        _ = v59
    }
    if (false && true) {
        fmt.Println(v6, v3)
        delete(v5, "hello")
        for fi60 := 0; fi60 < 4; fi60++ {
            for fi61 := 0; fi61 < 10; fi61++ {
                v6 = fi61
                if true {
                    switch "foo bar" {
                    case "":
                        if (441 == 32767) {
                            v6 = fi61
                            v62 := GenS4{F0: -2147483648}
                            _ = v62
                        }
                    case "foo bar":
                        var v63 uint32 = 0
                        _ = v63
                    }
                    v64 := []int{(^(v2)), v3, 42, v2}
                    _ = v64
                } else {
                    var v65 int32 = -2147483648
                    v66 := GenS1{F0: 2, F1: -2147483648, F2: -1}
                    var v67 int = (v6 << (((v1 | 570)) & 63))
                    _ = v65
                    _ = v66
                    _ = v67
                }
                v1 += v1
                var v68 uint16 = 65535
                _ = fi61
                _ = v68
            }
            for fi69 := 0; fi69 < 8; fi69++ {
                v70 := []int{(v2 >> ((558) & 63)), (v3 / ((-286) | 1)), fi60}
                _ = fi69
                _ = v70
            }
            _ = fi60
        }
    }
    for fi71 := 0; fi71 < 8; fi71++ {
        var v72 uint8 = 39
        if (0 < (v52 ^ int8(v6))) {
            var v73 int16 = -32768
            switch (^(v2)) {
            case 294:
                switch "" {
                case "":
                    var v74 float64 = -2.25
                    _ = v74
                case "a":
                    v3 += (v3 - (v4 & (v4 + (v4 * (v4 & -247)))))
                case "hello":
                    for fi75 := 0; fi75 < 15; fi75++ {
                        var v76 uint = 65535
                        v77 := GenS1{F0: 2, F1: -2147483648, F2: -9223372036854775808}
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                }
            case 2147483647:
                if (!((v1 & 660) != (v1 / ((155) | 1)))) {
                    v78 := GenS8{F0: 458, F1: 42}
                    for fi79 := 0; fi79 < 14; fi79++ {
                        v5[("foo bar" + ("\n\t" + "hello"))] = false
                        v2 -= v2
                        var v80 uint16 = uint16(v72)
                        var v81 uint = (v1 / ((255) | 1))
                        var v82 uint = (v1 * v1)
                        v72 += (v72 * v72)
                        switch "" {
                        case "a":
                            if true {
                                switch (v3 ^ (v4 & (v2 ^ 696))) {
                                case 9223372036854775807:
                                    fmt.Println(v6, v4)
                                }
                            } else {
                                v80 = v80
                            }
                        }
                        _ = fi79
                        _ = v80
                        _ = v81
                        _ = v82
                    }
                    v83 := GenS2{F0: ("" + "")}
                    _ = v78
                    _ = v83
                } else {
                    if (v1 == (v1 << ((4294967295) & 63))) {
                        var v84 uint16 = 255
                        _ = v84
                    }
                    v4 = v4
                    v1 -= (v1 & v1)
                }
            }
            delete(v5, "hello")
            _ = v73
        } else {
            v85 := GenS4{F0: -2147483648}
            v86 := GenS8{F0: 2, F1: 345}
            _ = v85
            _ = v86
        }
        _ = fi71
        _ = v72
    }
    var v87 int32 = 0
    v88 := []int{(-(v2)), (v2 >> ((v1) & 63)), 4294967295}
    v89 := GenS7{F0: 42}
    for fi90 := 0; fi90 < 11; fi90++ {
        if false {
            switch fi90 {
            case 2147483647:
                v2 -= v2
            case 127:
                v91 := GenS9{F0: (v87 | 0)}
                _ = v91
            case 616:
                v4 = int(v87)
            default:
                if ((v1 >> (((v1 + v1)) & 63)) < v1) {
                    for fi92 := 0; fi92 < 6; fi92++ {
                        _ = len(v88)
                        _ = fi92
                    }
                }
            }
            switch "hello" {
            case "\n\t":
                var v93 uint8 = 127
                _ = v93
            case "":
                var v94 float64 = 3.14159
                _ = v94
            case "hello":
                v2 = fi90
            default:
                _ = len(v88)
            }
        } else {
            if false {
                fmt.Println(v2, v4, v52)
                var v95 uint = v1
                v4 = (fi90 | 26)
                _ = v95
            }
            var v96 uint64 = 42
            var v97 int8 = (v52 + v52)
            _ = v96
            _ = v97
        }
        _ = fi90
    }
    v5["a"] = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("slicelen v88:", len(v88), v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v87
    _ = v88
    _ = v89
}
