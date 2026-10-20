package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint8; F1 int16; F2 uint64 }

func main() {
    var v1 bool = true
    var v2 uint32 = 65535
    v3 := []string{("foo bar" + ("\n\t" + ("a" + "a"))), "a", ("" + ""), ""}
    v4 := make(map[uint32]int)
    v2 = v2
    if (32767 > 2147483647) {
        switch ("hello" + ("a" + "hello")) {
        case "\n\t":
            var v5 int32 = 989
            _ = v5
        default:
            if true {
                v2 += 4294967295
            } else {
                _ = len(v3)
                var v6 bool = v1
                var v7 uint8 = 111
                _ = v6
                _ = v7
            }
        }
    }
    if false {
        if (1 <= 1) {
            switch -2147483648 {
            case 797:
                v2 = (v2 | (v2 & v2))
            case 667:
                _ = len(v3)
            case 92:
                switch 2147483647 {
                case -802:
                    switch (("foo bar" + "") + "hello") {
                    case "hello":
                        for fi8 := 0; fi8 < 14; fi8++ {
                            var v9 int64 = 1
                            v9 = -799
                            _ = fi8
                            _ = v9
                        }
                    default:
                        fmt.Println(v1, v1)
                    }
                case 128:
                    v2 += v2
                case 890:
                    for fi10 := 0; fi10 < 20; fi10++ {
                        var v11 int32 = -1
                        _ = fi10
                        _ = v11
                    }
                }
            default:
                if v1 {
                    switch 127 {
                    case 9223372036854775807:
                        if (!(!(false && (v1 && v1)))) {
                            v2 -= 2147483647
                            for fi12 := 0; fi12 < 18; fi12++ {
                                var v13 string = "foo bar"
                                fmt.Println(v1, fi12, fi12)
                                _ = fi12
                                _ = v13
                            }
                            v2 = v2
                        }
                    default:
                        delete(v4, v2)
                    }
                }
            }
        }
        switch "\n\t" {
        case "foo bar":
            v14 := v4[42]
            _ = v14
        case "":
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint16 = 65535
                fmt.Println(v2)
                _ = fi15
                _ = v16
            }
        }
    }
    if (125 == 47) {
        switch 4294967295 {
        case 2147483647:
            if v1 {
                switch "foo bar" {
                case "\n\t":
                    v1 = true
                case "a":
                    v3 = append(v3, (("a" + "a") + ("\n\t" + "hello")))
                case "foo bar":
                    v2 -= (v2 % ((2147483647) | 1))
                }
                v2 = 203
            } else {
                if (-1 < 42) {
                    var v17 uint64 = 4294967295
                    _ = v17
                }
            }
        case 32768:
            v1 = (4294967295 >= 42)
        case 127:
            delete(v4, v2)
        default:
            var v18 uint64 = 18446744073709551615
            _ = v18
        }
        switch 32767 {
        case 981:
            v1 = true
        default:
            v2 = v2
        }
        if (0 <= 42) {
            var v19 uint64 = 0
            v4[(v2 % ((136) | 1))] = -925
            _ = v19
        } else {
            v4[1] = 1
        }
    }
    var v20 int8 = 0
    if v1 {
        if (-610 == 937) {
            if false {
                var v21 int64 = int64(v2)
                var v22 int32 = 0
                _ = v21
                _ = v22
            } else {
                fmt.Println(v20, v2)
                v23 := GenS1{F0: 589, F1: (1 >= 2147483647)}
                _ = v23
            }
            fmt.Println(v1, v2)
        } else {
            if (-539 < 1) {
                v20 -= (v20 >> ((4294967295) & 63))
                for fi24 := 0; fi24 < 19; fi24++ {
                    v1 = v1
                    _ = fi24
                }
                var v25 int = 4294967295
                _ = v25
            }
        }
    } else {
        var v26 uint = 1
        v27 := GenS1{F0: -818, F1: (2 != 1)}
        _ = v26
        _ = v27
    }
    if (494 < 544) {
        fmt.Println(v20, v20, v2)
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int32 = 679
            var v30 float32 = 0.0
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    var v31 int32 = 586
    switch "a" {
    case "a":
        fmt.Println(v20)
    case "":
        if false {
            fmt.Println(v31, v1)
        }
    case "\n\t":
        switch "" {
        case "foo bar":
            v32 := GenS2{F0: 9223372036854775807, F1: 32767, F2: 2}
            _ = v32
        case "a":
            fmt.Println(v31)
        case "hello":
            v2 += v2
        default:
            var v33 float64 = (0.1 / -2.25)
            _ = v33
        }
    default:
        v34 := GenS3{F0: v31, F1: 42}
        _ = v34
    }
    fmt.Println(v2)
    switch 2147483647 {
    case -1:
        v1 = (!(-32768 == 17))
    default:
        if v1 {
            for fi35 := 0; fi35 < 6; fi35++ {
                switch ("hello" + "\n\t") {
                case "":
                    v36 := GenS3{F0: -330, F1: -9223372036854775808}
                    _ = v36
                default:
                    switch (fi35 ^ (^(fi35))) {
                    case 42:
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v31 -= 222
                            if v1 {
                                _ = len(v3)
                                var v38 int = (fi37 << ((1) & 63))
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                _ = v38
                            } else {
                                fmt.Println(v1)
                            }
                            v1 = true
                            v1 = (!true)
                            v1 = ((^(v2)) > v2)
                            _ = fi37
                        }
                    case -9223372036854775808:
                        var v39 float64 = float64((0.1 + 0.1))
                        _ = v39
                    case 1:
                        var v40 float64 = 0.0
                        _ = v40
                    }
                }
                if v1 {
                    if v1 {
                        v41 := GenS4{F0: v20}
                        fmt.Println(v20)
                        _ = v41
                    } else {
                        v20 += (-(v20))
                        var v42 int16 = 1000
                        v43 := GenS4{F0: -1}
                        _ = v42
                        _ = v43
                    }
                    if v1 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    if v1 {
                        v20 = (^(v20))
                        var v44 bool = v1
                        for fi45 := 0; fi45 < 2; fi45++ {
                            if v1 {
                                if (609 == 255) {
                                    v31 = (^(v31))
                                    v44 = true
                                    v31 -= 2147483647
                                }
                                _ = len(v3)
                                if (-2.25 == 3.14159) {
                                    v46 := GenS4{F0: 1}
                                    _ = v46
                                }
                            }
                            var v47 bool = ((((0.0 + 0.0) / (float64(((float32((3.14159 - 1e10)) - 1e10) * (1.5 - (float32((1.5 - float64(-2.25))) / -2.25)))) / 3.14159)) / 0.1) <= 3.14159)
                            if (float64(0.0) >= (3.14159 * (((1e10 / 1.0) - -2.25) / 3.14159))) {
                                switch -9223372036854775808 {
                                case 771:
                                    v44 = (v47 && v44)
                                default:
                                    if false {
                                        var v48 int32 = (v31 ^ (v31 & v31))
                                        _ = v48
                                    }
                                }
                                var v49 uint8 = 0
                                if true {
                                    fmt.Println(v47, v47, v49)
                                    v50 := GenS4{F0: (v20 >> ((65535) & 63))}
                                    v1 = v44
                                    _ = v50
                                } else {
                                    fmt.Println(v1, v44)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v51 int8 = v20
                                    _ = v51
                                }
                                _ = v49
                            }
                            fmt.Println(v2, v31, v20)
                            var v52 int8 = v20
                            fmt.Println(v1)
                            _ = fi45
                            _ = v47
                            _ = v52
                        }
                        _ = v44
                    } else {
                        var v53 int16 = -953
                        v31 += (v31 * v31)
                        for fi54 := 0; fi54 < 10; fi54++ {
                            fmt.Println(v31, fi54, fi54)
                            if ((42 >= -9223372036854775808) || (!(!false))) {
                                var v55 int8 = (v20 * -128)
                                _ = v55
                            } else {
                                switch fi35 {
                                case 9223372036854775807:
                                    if v1 {
                                        var v56 int8 = 91
                                        var v57 bool = v1
                                        fmt.Println(v53, v2)
                                        switch "" {
                                        case "a":
                                            switch 127 {
                                            case 2147483647:
                                                var v58 int64 = 140
                                                _ = v58
                                            case 4294967295:
                                                var v59 uint32 = 2147483648
                                                _ = v59
                                            case -9223372036854775808:
                                                var v60 string = "hello"
                                                _ = v60
                                            default:
                                                if (!v57) {
                                                    fmt.Println(v53)
                                                    if ((((0.1 - -2.25) <= float32((0.0 - 0.1))) && (-916 != 256)) && v1) {
                                                        var v61 int16 = (v53 >> ((255) & 63))
                                                        switch (fi54 + fi54) {
                                                        case -828:
                                                            v62 := GenS1{F0: -12, F1: v57}
                                                            _ = v62
                                                        default:
                                                            var v63 int = (fi54 + 0)
                                                            _ = v63
                                                        }
                                                        fmt.Println(v31, fi35, v56)
                                                        _ = v61
                                                    }
                                                    fmt.Println(v57, v20, fi35)
                                                } else {
                                                    v1 = true
                                                    v31 = (-(v31))
                                                    v64 := GenS4{F0: -128}
                                                    _ = v64
                                                }
                                            }
                                        }
                                        _ = v56
                                        _ = v57
                                    } else {
                                        fmt.Println(v20, v53)
                                        if ((-9223372036854775808 < fi35) || (v1 || (592 >= 32767))) {
                                            var v65 float32 = 0.1
                                            fmt.Println(v1, v31)
                                            _ = v65
                                        } else {
                                            v2 = v2
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v66 uint = 1
                                    v2 -= (v2 % ((v2) | 1))
                                    _ = v66
                                case 4294967295:
                                    v67 := GenS5{F0: 18446744073709551615, F1: 42}
                                    _ = v67
                                case 77:
                                    _ = len(v3)
                                }
                                fmt.Println(v1, v53, v20)
                                v68 := GenS4{F0: 100}
                                _ = v68
                            }
                            if false {
                                v69 := GenS4{F0: (v20 + v20)}
                                _ = v69
                            }
                            fmt.Println(v53, fi35)
                            _ = fi54
                        }
                        _ = v53
                    }
                    v2 = v2
                    v20 += 1
                }
                var v70 float32 = float32((3.14159 + 0.0))
                _ = fi35
                _ = v70
            }
            v20 = v20
            if v1 {
                for fi71 := 0; fi71 < 20; fi71++ {
                    if (!v1) {
                        v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                    }
                    _ = fi71
                }
            } else {
                fmt.Println(v1, v20)
            }
        }
    }
    v72 := v4[v2]
    v73 := GenS6{F0: 0, F1: ("hello" + ("" + ("foo bar" + "")))}
    switch 32767 {
    case -32:
        var v74 float32 = (1e10 + (0.0 * 0.0))
        _ = v74
    case -1:
        v4[(v2 & 2147483647)] = v72
    default:
        v20 = (^(v20))
    }
    fmt.Println(v20, v31, v31)
    if (18446744073709551615 == 42) {
        v31 = v31
        var v75 uint32 = v2
        _ = v75
    } else {
        v20 = v20
        var v76 int64 = 9223372036854775807
        switch "\n\t" {
        case "foo bar":
            for fi77 := 0; fi77 < 6; fi77++ {
                for fi78 := 0; fi78 < 19; fi78++ {
                    v79 := GenS4{F0: v20}
                    v80 := GenS7{F0: 32767, F1: (v20 * (v20 + -128))}
                    var v81 int8 = (v20 >> ((2) & 63))
                    v82 := GenS8{F0: -911}
                    _ = fi78
                    _ = v79
                    _ = v80
                    _ = v81
                    _ = v82
                }
                v83 := GenS9{F0: v31, F1: 255, F2: 42}
                if ((1e10 + (1e10 + (3.14159 + float64(1e10)))) >= (3.14159 / (float64(0.1) + (1e10 / 1e10)))) {
                    var v84 uint8 = 113
                    switch 0 {
                    case -2147483648:
                        var v85 uint16 = 255
                        _ = v85
                    case -872:
                        for fi86 := 0; fi86 < 4; fi86++ {
                            fmt.Println(fi86, v72, v76)
                            _ = fi86
                        }
                    default:
                        for fi87 := 0; fi87 < 13; fi87++ {
                            fmt.Println(fi77, v31)
                            var v88 float64 = -2.25
                            _ = fi87
                            _ = v88
                        }
                    }
                    _ = v84
                } else {
                    switch "" {
                    case "":
                        if v1 {
                            v89 := GenS5{F0: 1, F1: 42}
                            var v90 int32 = v31
                            switch (fi77 | 4294967295) {
                            case -610:
                                for fi91 := 0; fi91 < 9; fi91++ {
                                    if (0.1 >= 1.5) {
                                        v20 += (-(v20))
                                        v20 += 100
                                        v76 = v76
                                    }
                                    _ = fi91
                                }
                            case 32768:
                                var v92 int8 = (v20 * v20)
                                _ = v92
                            case -1:
                                var v93 int16 = 1
                                _ = v93
                            default:
                                for fi94 := 0; fi94 < 12; fi94++ {
                                    if (true && (0 < 42)) {
                                        var v95 int = v72
                                        _ = v95
                                    }
                                    v76 = 0
                                    _ = fi94
                                }
                            }
                            _ = v89
                            _ = v90
                        } else {
                            v96 := GenS10{F0: 255, F1: -1, F2: 2}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch int(v76) {
                            case -362:
                                v97 := v4[v2]
                                _ = v97
                            default:
                                switch ("\n\t" + ("\n\t" + "a")) {
                                case "a":
                                    fmt.Println(v2)
                                default:
                                    v1 = v1
                                }
                            }
                            _ = v96
                        }
                    default:
                        for fi98 := 0; fi98 < 5; fi98++ {
                            var v99 uint = 0
                            _ = fi98
                            _ = v99
                        }
                    }
                }
                _ = fi77
                _ = v83
            }
        }
        _ = v76
    }
    for fi100 := 0; fi100 < 17; fi100++ {
        v4[981] = (v72 | (fi100 & 889))
        v101 := v4[(v2 ^ (v2 / ((v2) | 1)))]
        _ = fi100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v72
    _ = v73
}
