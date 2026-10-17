package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 bool; F1 int }
func genFn0() int {
    var v6 int16 = 1000
    switch "a" {
    case "a":
        var v7 bool = false
        _ = v7
    case "foo bar":
        v8 := []int{}
        _ = v8
    case "hello":
        var v9 uint16 = 255
        _ = v9
    }
    if (0.0 > float64(0.0)) {
        for fi10 := 0; fi10 < 11; fi10++ {
            if true {
                switch ((("" + "") + (("foo bar" + (("" + "") + "hello")) + "a")) + "") {
                case "hello":
                    v11 := GenS1{F0: 65535}
                    _ = v11
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        if ((3.14159 * 0.1) == 1.5) {
                            var v12 uint16 = 255
                            fmt.Println(v6, v12, fi10)
                            _ = v12
                        } else {
                            v6 = v6
                            v13 := GenS1{F0: 2}
                            v14 := GenS1{F0: 494}
                            _ = v13
                            _ = v14
                        }
                    case "a":
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 int32 = 516
                            switch "" {
                            case "foo bar":
                                v16 -= (v16 & v16)
                            default:
                                fmt.Println(v16)
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                }
            } else {
                if false {
                    fmt.Println(fi10)
                    v17 := GenS1{F0: 129}
                    _ = v17
                }
                v18 := []int{}
                var v19 bool = (0 == 2147483648)
                _ = v18
                _ = v19
            }
            var v20 uint32 = 999
            v21 := GenS1{F0: 255}
            _ = fi10
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := []int{128, fi22, fi22, fi22}
            _ = fi22
            _ = v23
        }
        var v24 float64 = 1.5
        _ = v24
    }
    _ = v6
    return 42
}
func genFn1(p0 int, p1 uint32) uint32 {
    fmt.Println(p1)
    var v25 int32 = 704
    var v26 uint16 = 42
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 0
}

func main() {
    v1 := []int{215, 32767}
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 bool = (false && false)
    v5 := make(map[string]int)
    for fi27 := 0; fi27 < 15; fi27++ {
        if v4 {
            v5["hello"] = 2147483647
            for fi28 := 0; fi28 < 5; fi28++ {
                v4 = false
                _ = fi28
            }
            v29 := GenS1{F0: 1}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 8; fi30++ {
                v4 = v4
                v31 := GenS1{F0: 2}
                var v32 uint64 = 42
                fmt.Println(fi27)
                _ = fi30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 9; fi33++ {
                v34 := GenS2{F0: 65535}
                v4 = (0.0 == 0.0)
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 14; fi35++ {
                v1 = append(v1, fi35)
                if (true && true) {
                    v4 = v4
                    v36 := GenS3{F0: -1}
                    if false {
                        fmt.Println(fi27)
                        v3[fi35] = 0
                    }
                    var v37 float64 = 3.14159
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v38 := GenS4{F0: 1, F1: -96, F2: 0}
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi35
            }
        }
        _ = fi27
    }
    switch 32767 {
    case 698:
        v39 := v2["\n\t"]
        _ = v39
    case 32768:
        var v40 int8 = 42
        _ = v40
    }
    _ = len(v1)
    if v4 {
        var v41 int = -874
        _ = v41
    }
    fmt.Println(v4)
    if len(v1) > 0 {
        v1[0] = 913
    }
    v4 = (((0.1 + float64((-2.25 / -2.25))) * (-2.25 / 3.14159)) >= 1.5)
    v42 := GenS3{F0: 119}
    if (47 > 0) {
        if (-1 <= -45) {
            v43 := GenS1{F0: 255}
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 11; fi44++ {
            for fi45 := 0; fi45 < 2; fi45++ {
                var v46 float32 = 0.0
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
        var v47 uint64 = 4294967295
        if len(v1) > 0 {
            v1[0] = 32767
        }
        _ = v47
    }
    switch ("a" + ("\n\t" + "\n\t")) {
    case "\n\t":
        var v48 uint32 = 434
        _ = v48
    case "hello":
        v4 = ((127 > 128) || (v4 || ((0 > -629) && true)))
    default:
        v4 = (-1 == 0)
    }
    v4 = (65535 < 282)
    for fi49 := 0; fi49 < 20; fi49++ {
        v4 = (false || (0.0 <= (1.5 + (float64((0.1 - (float32(1e10) / float32(1.5)))) * 3.14159))))
        var v50 int64 = 42
        _ = fi49
        _ = v50
    }
    if v4 {
        switch ("foo bar" + ("\n\t" + "\n\t")) {
        case "a":
            var v51 uint = 361
            _ = v51
        case "foo bar":
            if ("" > (("a" + "hello") + "hello")) {
                switch "hello" {
                case "":
                    if len(v1) > 0 {
                        v1[0] = 364
                    }
                default:
                    v5[((("foo bar" + "a") + ("hello" + "\n\t")) + (("" + "\n\t") + "hello"))] = 2147483647
                }
                for fi52 := 0; fi52 < 9; fi52++ {
                    fmt.Println(fi52, v4)
                    v4 = (!false)
                    if ((("" + "foo bar") + "") < ((("\n\t" + ("\n\t" + "foo bar")) + "") + "")) {
                        if (!(1000 >= 32767)) {
                            v53 := GenS5{F0: fi52}
                            switch "a" {
                            case "foo bar":
                                for fi54 := 0; fi54 < 4; fi54++ {
                                    switch -48 {
                                    case 885:
                                        var v55 bool = true
                                        _ = v55
                                    case 32767:
                                        v56 := GenS2{F0: 42}
                                        _ = v56
                                    default:
                                        v57 := GenS6{F0: 0, F1: 0}
                                        _ = v57
                                    }
                                    v4 = v4
                                    _ = fi54
                                }
                            case "hello":
                                v4 = (v4 || true)
                            case "\n\t":
                                if (88 >= 32) {
                                    if (genFn0() > (fi52 - (fi52 & (fi52 + 32767)))) {
                                        v4 = (!(("\n\t" > "a") && (42 < 0)))
                                        v1 = append(v1, fi52)
                                    } else {
                                        v4 = v4
                                    }
                                    switch 32767 {
                                    case 737:
                                        if (!v4) {
                                            if (((fi52 - 4294967295) >= genFn0()) || (true && false)) {
                                                var v58 float64 = 1e10
                                                var v59 float32 = -2.25
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v4 = (!(!(846 == 610)))
                                                var v60 int32 = 698
                                                _ = v60
                                            }
                                            v61 := v3[(fi52 >> ((2) & 63))]
                                            _ = v61
                                        }
                                    }
                                    switch (fi52 * -9223372036854775808) {
                                    case 986:
                                        v4 = (true && v4)
                                    case -207:
                                        v4 = (!true)
                                    default:
                                        v4 = v4
                                    }
                                } else {
                                    for fi62 := 0; fi62 < 4; fi62++ {
                                        if (2147483648 < 2) {
                                            switch (fi52 * fi52) {
                                            case 32768:
                                                v4 = v4
                                            case 127:
                                                switch fi62 {
                                                case 153:
                                                    switch "hello" {
                                                    case "a":
                                                        v4 = (!((0 >= 878) && v4))
                                                    }
                                                case 0:
                                                    if ((false || v4) && (1e10 != -2.25)) {
                                                        fmt.Println(v4)
                                                    }
                                                default:
                                                    v63 := v5["\n\t"]
                                                    _ = v63
                                                }
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    var v64 uint16 = 1
                                                    _ = v64
                                                default:
                                                    var v65 float32 = 1e10
                                                    _ = v65
                                                }
                                            }
                                            v66 := GenS7{F0: (1 != 4294967295), F1: fi62}
                                            v67 := GenS2{F0: 42}
                                            var v68 int8 = 127
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                        var v69 uint8 = 176
                                        _ = fi62
                                        _ = v69
                                    }
                                }
                            }
                            var v70 int64 = -406
                            _ = v53
                            _ = v70
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v4 = ((1 == 2) && v4)
                            var v71 uint = 897
                            _ = v71
                        }
                        v4 = (!(fi52 > fi52))
                        var v72 float32 = ((3.14159 + 1.5) / 0.1)
                        _ = v72
                    }
                    _ = fi52
                }
            }
        }
        var v73 float32 = ((1e10 / 1.0) + 3.14159)
        _ = v73
    } else {
        v74 := GenS7{F0: (32767 < 255), F1: -133}
        _ = v74
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
