package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
func genFn0(p0 int32, p1 int8) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 255, F1: 127}
        p1 = p1
        v7 := []int{-1, (fi5 + 128), fi5, (fi5 ^ (fi5 >> ((65535) & 63)))}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = 32768
    _ = p0
    _ = p1
    _ = v8
    return p1
}
func genFn1() string {
    v9 := GenS1{F0: 18446744073709551615, F1: 4294967295}
    var v10 uint32 = 577
    fmt.Println(v10)
    fmt.Println(v10)
    _ = v10
    _ = v9
    return ("a" + "hello")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -1
    v3 := []int{}
    v4 := []uint32{0, 20}
    if true {
        v1[638] = 0
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    var v12 uint8 = 21
    if false {
        if len(v4) > 0 {
            v4[0] = 80
        }
    } else {
        var v13 int64 = 0
        switch -729 {
        case -130:
            v3 = append(v3, 419)
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: 255, F1: (fi14 + (fi14 & 2147483647))}
                switch "foo bar" {
                case "hello":
                    var v16 uint8 = v12
                    _ = v16
                case "":
                    switch fi14 {
                    case 117:
                        v2 = v2
                    case -9223372036854775808:
                        var v17 string = "hello"
                        _ = v17
                    }
                case "\n\t":
                    v2 += (v2 * v2)
                }
                for fi18 := 0; fi18 < 12; fi18++ {
                    delete(v1, 905)
                    switch -931 {
                    case 682:
                        v13 -= 42
                    case 2147483647:
                        var v19 int64 = v13
                        _ = v19
                    case 821:
                        fmt.Println(v2, v12, v2)
                    }
                    _ = fi18
                }
                _ = fi14
                _ = v15
            }
        }
        _ = v13
    }
    var v20 int = 0
    var v21 uint8 = (v12 + 1)
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint16 = 2
        _ = fi22
        _ = v23
    }
    if ((0.0 == (1.5 * 0.0)) || (140 == 1)) {
        for fi24 := 0; fi24 < 19; fi24++ {
            fmt.Println(v21)
            _ = fi24
        }
        switch 32768 {
        case 32768:
            var v25 uint8 = 42
            _ = v25
        case 599:
            var v26 uint8 = v12
            _ = v26
        default:
            if false {
                switch (^(v20)) {
                case 4294967295:
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if (0.0 <= (0.0 + -2.25)) {
                            var v28 uint = 255
                            switch "" {
                            case "foo bar":
                                var v29 bool = ("hello" <= "\n\t")
                                _ = v29
                            case "hello":
                                if (147 >= 4294967295) {
                                    fmt.Println(v28, v28, fi27)
                                    fmt.Println(v28, fi27)
                                    v30 := GenS2{F0: (v28 % (((v28 << (((v28 % ((v28) | 1))) & 63))) | 1))}
                                    _ = v30
                                }
                            }
                            v28 = 18446744073709551615
                            _ = v28
                        }
                        _ = fi27
                    }
                case -863:
                    fmt.Println(v2, v2)
                }
                var v31 bool = false
                _ = v31
            } else {
                switch "" {
                case "\n\t":
                    var v32 uint = 279
                    _ = v32
                }
                v20 = (v20 >> ((2) & 63))
                if (!false) {
                    switch (^(v20)) {
                    case 32768:
                        var v33 uint = 890
                        _ = v33
                    case 0:
                        switch "" {
                        case "foo bar":
                            v4 = append(v4, 575)
                        case "\n\t":
                            fmt.Println(v12, v12, v20)
                        }
                    case 960:
                        v21 = v21
                    }
                    for fi34 := 0; fi34 < 16; fi34++ {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            if true {
                                v36 := GenS2{F0: uint(v21)}
                                if false {
                                    v4 = append(v4, 4294967295)
                                    v37 := v1[2147483647]
                                    fmt.Println(fi35, v20)
                                    switch v37 {
                                    case 135:
                                        var v38 int32 = (v2 << ((0) & 63))
                                        _ = v38
                                    default:
                                        var v39 int64 = -1
                                        _ = v39
                                    }
                                    if (1.5 == -2.25) {
                                        v40 := GenS2{F0: 660}
                                        fmt.Println(fi34, v2, v37)
                                        v20 += (fi35 << ((4294967295) & 63))
                                        _ = v40
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            var v41 float64 = 3.14159
                                            _ = v41
                                        }
                                        fmt.Println(v2, v12, v12)
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v2)
                                    case "\n\t":
                                        var v42 int = 32768
                                        _ = v42
                                    case "":
                                        v20 = v37
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            var v43 bool = (!true)
                                            _ = v43
                                        case "foo bar":
                                            var v44 float32 = (1e10 - 0.0)
                                            _ = v44
                                        default:
                                            var v45 bool = (90 == 42)
                                            _ = v45
                                        }
                                    }
                                    v46 := GenS2{F0: 4294967295}
                                    var v47 int16 = int16(v12)
                                    delete(v1, 2147483647)
                                    switch "hello" {
                                    case "":
                                        v37 -= fi35
                                        v12 += (v12 >> ((65535) & 63))
                                    case "\n\t":
                                        var v48 int = (fi34 >> ((2) & 63))
                                        _ = v48
                                    case "a":
                                        v1[0] = v37
                                    }
                                    if false {
                                        v49 := GenS3{F0: (32767 <= 2), F1: v2}
                                        var v50 float32 = 1e10
                                        if false {
                                            switch 32768 {
                                            case -696:
                                                v2 = (v2 & v2)
                                                if false {
                                                    var v51 bool = (!true)
                                                    switch ("foo bar" + ("foo bar" + "hello")) {
                                                    case "foo bar":
                                                        var v52 int8 = -85
                                                        _ = v52
                                                    }
                                                    _ = v51
                                                }
                                                var v53 float32 = (1e10 * v50)
                                                switch fi35 {
                                                case 4294967295:
                                                    if (fi34 > 32767) {
                                                        v50 -= (v53 - (((v53 + (float32(1.5) + v53)) / 3.14159) * (v53 / float32(-2.25))))
                                                        var v54 float64 = 1e10
                                                        _ = v54
                                                    }
                                                default:
                                                    v47 = v47
                                                }
                                                _ = v53
                                            }
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                switch "a" {
                                                case "hello":
                                                    fmt.Println(v21)
                                                }
                                            case "\n\t":
                                                v12 = v21
                                            case "":
                                                v50 = ((1e10 / (v50 - ((v50 - -2.25) / v50))) + v50)
                                            }
                                            var v55 float64 = 0.0
                                            var v56 uint = 4294967295
                                            _ = v55
                                            _ = v56
                                        }
                                        fmt.Println(v20, v2, v12)
                                        _ = v49
                                        _ = v50
                                    }
                                    var v57 int = (fi34 % ((int(v2)) | 1))
                                    if false {
                                        fmt.Println(fi34, fi35)
                                        v58 := GenS1{F0: 42, F1: (^(fi35))}
                                        v59 := GenS1{F0: 65535, F1: fi35}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v37
                                    _ = v46
                                    _ = v47
                                    _ = v57
                                }
                                var v60 int = fi35
                                _ = v36
                                _ = v60
                            } else {
                                v61 := GenS3{F0: true, F1: (v2 | 42)}
                                var v62 float32 = 0.0
                                if (0 <= 595) {
                                    v62 = 0.0
                                    v62 = 3.14159
                                    var v63 int8 = -1
                                    _ = v63
                                }
                                _ = v61
                                _ = v62
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi35
                        }
                        _ = fi34
                    }
                    v12 += 42
                }
            }
        }
        if (!true) {
            v12 = (v12 | 2)
            for fi64 := 0; fi64 < 20; fi64++ {
                if false {
                    v65 := GenS4{F0: false, F1: -1, F2: 9223372036854775807}
                    _ = v65
                }
                _ = fi64
            }
            v66 := GenS2{F0: 65535}
            _ = v66
        }
    } else {
        var v67 float64 = 1e10
        _ = v67
    }
    switch v20 {
    case 2147483647:
        var v68 uint64 = 97
        _ = v68
    default:
        switch 32768 {
        case 42:
            if false {
                if (!true) {
                    if len(v3) > 0 {
                        v3[0] = (v20 + (^(v20)))
                    }
                    _ = len(v3)
                    if ((94 == 42) && ((v20 % (((v20 % ((v20) | 1))) | 1)) < v20)) {
                        if false {
                            v69 := GenS4{F0: (4294967295 > 773), F1: 127, F2: (v20 * v20)}
                            _ = v69
                        } else {
                            v21 = v12
                        }
                    }
                } else {
                    var v70 string = ("a" + "\n\t")
                    for fi71 := 0; fi71 < 6; fi71++ {
                        for fi72 := 0; fi72 < 6; fi72++ {
                            v73 := GenS2{F0: 255}
                            switch v20 {
                            case 32768:
                                switch "" {
                                case "a":
                                    var v74 uint64 = 884
                                    _ = v74
                                case "\n\t":
                                    var v75 uint8 = 2
                                    _ = v75
                                }
                            }
                            _ = fi72
                            _ = v73
                        }
                        v4 = append(v4, 721)
                        _ = fi71
                    }
                    _ = v70
                }
                for fi76 := 0; fi76 < 7; fi76++ {
                    var v77 uint64 = 4294967295
                    _ = fi76
                    _ = v77
                }
            }
        case 824:
            fmt.Println(v21, v20, v12)
        case 0:
            var v78 int64 = -9223372036854775808
            _ = v78
        }
    }
    var v79 string = ("foo bar" + "")
    v80 := GenS4{F0: false, F1: 1, F2: int(v2)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v79
    _ = v80
}
