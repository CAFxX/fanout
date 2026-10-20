package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 14; fi4++ {
        if (true || false) {
            if true {
                var v5 uint8 = 127
                if (!(1000 > 42)) {
                    var v6 float64 = (1e10 + (-2.25 + 0.1))
                    var v7 uint64 = 18446744073709551615
                    var v8 float32 = (float32((v6 - ((-2.25 / v6) * v6))) * -2.25)
                    _ = v6
                    _ = v7
                    _ = v8
                } else {
                    if false {
                        var v9 int = fi4
                        var v10 float64 = ((-2.25 / -2.25) / 0.1)
                        _ = v10
                        _ = v9
                    } else {
                        if false {
                            v5 = 42
                            v5 = (v5 << ((uint(v5)) & 63))
                        } else {
                            v5 = 255
                        }
                        var v11 int16 = 1
                        _ = v11
                    }
                    fmt.Println(v5)
                }
                v2["\n\t"] = (v5 >= v5)
                _ = v5
            }
            var v12 string = ("hello" + "\n\t")
            _ = v12
        }
        if (!true) {
            if false {
                var v13 uint32 = 65535
                switch fi4 {
                case -2147483648:
                    v14 := GenS1{F0: -39}
                    _ = v14
                case 32767:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        if true {
                            if (18446744073709551615 < 98) {
                                if false {
                                    v13 = (v13 | (^(v13)))
                                    v1[521] = false
                                    fmt.Println(fi15, fi4)
                                }
                                v16 := GenS2{F0: "\n\t", F1: (fi4 / ((fi15) | 1)), F2: false}
                                var v17 uint64 = 4294967295
                                _ = v16
                                _ = v17
                            } else {
                                var v18 uint16 = 0
                                _ = v18
                            }
                        } else {
                            v13 = 2147483647
                            fmt.Println(v13, fi4)
                            var v19 uint32 = v13
                            _ = v19
                        }
                        v20 := GenS2{F0: ("\n\t" + "\n\t"), F1: (fi4 ^ 62), F2: (!true)}
                        _ = fi15
                        _ = v20
                    }
                default:
                    v13 -= v13
                }
                var v21 uint = 273
                _ = v13
                _ = v21
            } else {
                v22 := GenS1{F0: 100}
                _ = v22
            }
            var v23 uint8 = 2
            _ = v23
        } else {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v24 int16 = 1000
                _ = v24
            default:
                if (2 <= 42) {
                    if (false || (!true)) {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            fmt.Println(fi25, fi25)
                            var v26 uint = 18446744073709551615
                            _ = fi25
                            _ = v26
                        }
                    }
                    var v27 int64 = 9223372036854775807
                    if false {
                        switch ("a" + ("foo bar" + "foo bar")) {
                        case "hello":
                            if (float64(0.1) != 1e10) {
                                if (false && (508 != 814)) {
                                    v27 = (^(v27))
                                }
                                if (!true) {
                                    var v28 int64 = (-(v27))
                                    v29 := GenS1{F0: 100}
                                    _ = v28
                                    _ = v29
                                }
                            }
                        case "":
                            for fi30 := 0; fi30 < 1; fi30++ {
                                fmt.Println(fi4)
                                var v31 string = "hello"
                                v32 := v2["foo bar"]
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = len(v2)
                    }
                    _ = v27
                } else {
                    var v33 int64 = 1
                    v33 = 1
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 bool = false
                v35 = true
                v35 = false
                var v36 int64 = 169
                switch (fi34 >> ((275) & 63)) {
                case 504:
                    switch 27 {
                    case 811:
                        v37 := GenS1{F0: -1}
                        _ = v37
                    case -788:
                        v36 = v36
                    case 1:
                        switch 0 {
                        case 32768:
                            v35 = (0 < 554)
                        }
                    }
                case 42:
                    var v38 bool = v35
                    _ = v38
                default:
                    switch -9223372036854775808 {
                    case 446:
                        fmt.Println(v36)
                    case 1:
                        v39 := []int{-9223372036854775808, (fi34 | 959), (fi34 >> ((255) & 63))}
                        _ = v39
                    }
                }
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi4
    }
    v40 := GenS3{F0: 32767, F1: 1}
    var v41 int8 = -34
    v41 = (v41 & v41)
    if true {
        for fi42 := 0; fi42 < 18; fi42++ {
            for fi43 := 0; fi43 < 8; fi43++ {
                if true {
                    var v44 float64 = -2.25
                    _ = v44
                }
                var v45 int32 = -1
                fmt.Println(v45, v45, fi42)
                _ = fi43
                _ = v45
            }
            v46 := v3[204]
            _ = fi42
            _ = v46
        }
    }
    fmt.Println(v41)
    var v47 uint = 255
    delete(v1, 869)
    v47 = v47
    v47 -= (v47 << (((^(v47))) & 63))
    fmt.Println(v41)
    var v48 uint32 = 0
    for fi49 := 0; fi49 < 9; fi49++ {
        v50 := GenS1{F0: (v41 / (((v41 % ((v41) | 1))) | 1))}
        var v51 int32 = 2147483647
        _ = fi49
        _ = v50
        _ = v51
    }
    var v52 uint16 = uint16(v47)
    for fi53 := 0; fi53 < 8; fi53++ {
        for fi54 := 0; fi54 < 13; fi54++ {
            if ((395 != 1) || (true || ((!(74 < 1)) && false))) {
                switch ("\n\t" + ("hello" + ("" + ("hello" + ("hello" + ("hello" + "a")))))) {
                case "foo bar":
                    switch (fi54 & 127) {
                    case 32768:
                        var v55 uint64 = 18446744073709551615
                        _ = v55
                    case 42:
                        if false {
                            var v56 int16 = -816
                            if false {
                                var v57 float64 = (1.5 - float64(1e10))
                                _ = v57
                            } else {
                                var v58 int8 = v41
                                _ = v58
                            }
                            _ = v56
                        }
                    case 128:
                        switch fi54 {
                        case 2147483647:
                            var v59 int = fi54
                            _ = v59
                        case -1:
                            v47 = (^(v47))
                        }
                        v52 += (v52 / ((v52) | 1))
                        v60 := []int{fi53, (fi53 / (((^(fi54))) | 1))}
                        _ = v60
                    }
                case "\n\t":
                    if (!false) {
                        var v61 float64 = float64(((3.14159 * -2.25) / (1.5 + (1.5 * float32(((0.1 * 1.5) - (1e10 - 1e10)))))))
                        _ = v61
                    }
                case "a":
                    var v62 uint32 = v48
                    _ = v62
                }
                v63 := []int{fi53, 2147483647}
                var v64 int64 = -1
                _ = v63
                _ = v64
            } else {
                if true {
                    v47 -= v47
                    switch "a" {
                    case "":
                        v65 := GenS2{F0: "", F1: fi53, F2: false}
                        _ = v65
                    case "a":
                        v2["a"] = (127 <= 42)
                    case "foo bar":
                        if false {
                            v48 -= v48
                            var v66 uint32 = v48
                            delete(v2, "hello")
                            _ = v66
                        } else {
                            var v67 float64 = (((3.14159 + ((1e10 * (0.1 - 0.1)) - 0.0)) / 0.1) + 3.14159)
                            v47 = (v47 >> ((v47) & 63))
                            var v68 int32 = -1
                            _ = v67
                            _ = v68
                        }
                    default:
                        var v69 int16 = -751
                        _ = v69
                    }
                } else {
                    if (-1 <= 1) {
                        v52 -= (v52 + (v52 / ((v52) | 1)))
                    }
                    v70 := v3[100]
                    switch ("foo bar" + "hello") {
                    case "hello":
                        if (!true) {
                            v71 := GenS4{F0: 2, F1: 42, F2: (^(v47))}
                            if (-384 > v70) {
                                v52 += (v52 - 243)
                                var v72 int8 = 1
                                _ = v72
                            }
                            _ = v71
                        } else {
                            if true {
                                if (uint16(v48) > 1) {
                                    if false {
                                        var v73 int16 = 0
                                        var v74 int16 = v73
                                        _ = v73
                                        _ = v74
                                    }
                                } else {
                                    var v75 uint64 = 858
                                    v41 = int8(v48)
                                    var v76 int32 = 33
                                    _ = v75
                                    _ = v76
                                }
                            }
                            if (!((v47 | (v47 | v47)) == v47)) {
                                var v77 int32 = 106
                                fmt.Println(fi53, v48)
                                v78 := []int{}
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v78) > 0 {
                                        v78[0] = fi53
                                    }
                                }
                                _ = v77
                                _ = v78
                            }
                            v79 := v3[86]
                            _ = v79
                        }
                    default:
                        v47 += v47
                    }
                    _ = v70
                }
                if (false && true) {
                    switch "foo bar" {
                    case "hello":
                        switch (("" + "hello") + "foo bar") {
                        case "":
                            fmt.Println(v52, v48, v52)
                        case "hello":
                            v80 := GenS5{F0: false}
                            _ = v80
                        }
                    case "a":
                        var v81 uint = (^(v47))
                        _ = v81
                    }
                    var v82 float64 = ((0.0 - 1.5) / float64(1e10))
                    var v83 bool = false
                    _ = v82
                    _ = v83
                }
                var v84 uint = 867
                _ = v84
            }
            v85 := v1[65535]
            _ = fi54
            _ = v85
        }
        _ = fi53
    }
    v41 += v41
    if false {
        if (((^(v52)) == v52) && false) {
            switch (("\n\t" + "a") + "foo bar") {
            case "a":
                fmt.Println(v41)
            case "":
                v86 := GenS4{F0: 128, F1: 42, F2: (v47 - (v47 / ((v47) | 1)))}
                _ = v86
            default:
                v41 -= -55
            }
            v52 = 255
        } else {
            if (true && (664 == 1)) {
                var v87 uint = v47
                switch "\n\t" {
                case "hello":
                    var v88 int16 = 32767
                    _ = v88
                default:
                    var v89 float64 = 1.5
                    _ = v89
                }
                _ = v87
            }
        }
        v52 = (v52 ^ uint16(v48))
        v90 := GenS6{F0: 2}
        _ = v90
    }
    switch 1 {
    case 127:
        var v91 uint64 = 1
        _ = v91
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v47
    _ = v48
    _ = v52
}
