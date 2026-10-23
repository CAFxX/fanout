package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint8; F1 int8; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint8; F1 bool; F2 int }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int; F1 uint32; F2 uint64 }

func main() {
    var v1 int8 = 127
    var v2 string = (("a" + ("a" + "foo bar")) + "\n\t")
    v3 := []string{v2, v2}
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    _ = len(v3)
    var v7 float32 = (float32((-2.25 + -2.25)) / -2.25)
    switch ((v2 + "hello") + v2) {
    case "":
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 748}
                fmt.Println(fi8, v7, v2)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "foo bar":
        var v11 int = 42
        _ = v11
    case "\n\t":
        switch 21 {
        case -2147483648:
            v2 = v2
        case 2147483647:
            fmt.Println(v1, v7)
        case 4294967295:
            _ = len(v5)
        }
    default:
        var v12 uint32 = 446
        _ = v12
    }
    v2 = v2
    var v13 int = int(v1)
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            if (161 >= 724) {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                var v16 uint16 = 506
                if true {
                    v3 = append(v3, v2)
                    v17 := v5[(("" + v2) + "")]
                    v2 = v2
                    _ = v17
                } else {
                    var v18 uint = 788
                    switch v2 {
                    case "\n\t":
                        v16 -= v16
                    case "":
                        var v19 int64 = 115
                        _ = v19
                    }
                    _ = v18
                }
                _ = v16
            } else {
                fmt.Println(v1)
                var v20 float32 = float32(3.14159)
                if true {
                    fmt.Println(v2, v7)
                    v7 += float32((3.14159 * 1.5))
                    var v21 float32 = (((((v20 + ((v7 * 0.0) * (v20 / v7))) + -2.25) * v7) / v7) + ((0.1 - -2.25) - v7))
                    _ = v21
                }
                _ = v20
            }
            if true {
                if false {
                    var v22 uint16 = 814
                    v7 -= ((float32(1.5) / (v7 + (((v7 + (((v7 / v7) / v7) / float32(0.1))) - v7) + (v7 / v7)))) * float32((((3.14159 / 0.1) + (0.0 / (1.5 / 1e10))) / 1.0)))
                    v23 := GenS1{F0: 9223372036854775807}
                    _ = v22
                    _ = v23
                }
                if false {
                    v13 -= fi14
                } else {
                    var v24 string = v2
                    _ = v24
                }
            }
            if false {
                var v25 string = ""
                v25 = v2
                _ = v25
            }
            _ = fi15
        }
        v26 := GenS2{F0: false, F1: 42}
        v13 += 1
        _ = fi14
        _ = v26
    }
    fmt.Println(v2, v2)
    v3 = append(v3, "\n\t")
    var v27 bool = (150 > 127)
    if false {
        switch v2 {
        case "hello":
            v28 := v5[(v2 + (v2 + "a"))]
            _ = v28
        }
    } else {
        v2 = v2
        switch int(v1) {
        case 2147483647:
            v3 = append(v3, ((v2 + "hello") + v2))
        case 158:
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(v1, v1)
                v30 := GenS2{F0: (!(285 != 452)), F1: 9223372036854775807}
                _ = fi29
                _ = v30
            }
        case 9223372036854775807:
            for fi31 := 0; fi31 < 10; fi31++ {
                if true {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    if (4294967295 <= 2) {
                        if v27 {
                            for fi32 := 0; fi32 < 19; fi32++ {
                                v2 = ("a" + (("a" + "a") + v2))
                                _ = fi32
                            }
                            var v33 float32 = 0.1
                            _ = v33
                        } else {
                            var v34 int8 = 127
                            if (-684 <= 379) {
                                v1 = v34
                                v27 = false
                            } else {
                                var v35 bool = (v1 == (v1 * v34))
                                if (3.14159 < 3.14159) {
                                    v36 := GenS1{F0: 1}
                                    var v37 int64 = -564
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        switch v2 {
                                        case "hello":
                                            if v27 {
                                                if (!true) {
                                                    v3 = append(v3, v2)
                                                    if v35 {
                                                        if ((255 == 756) || (!v27)) {
                                                            if false {
                                                                v35 = (32767 == 0)
                                                                if v27 {
                                                                    v13 -= fi38
                                                                    v39 := GenS2{F0: (42 == 1), F1: -936}
                                                                    _ = v39
                                                                }
                                                                switch fi38 {
                                                                case 0:
                                                                    if v27 {
                                                                        delete(v5, "a")
                                                                    }
                                                                default:
                                                                    var v40 int = fi38
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                v41 := GenS1{F0: (v37 >> ((2) & 63))}
                                                                v42 := GenS3{F0: 87, F1: 0, F2: 9223372036854775807}
                                                                var v43 int = 42
                                                                _ = v41
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        } else {
                                                            var v44 float32 = -2.25
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v44
                                                        }
                                                        var v45 uint8 = 3
                                                        _ = v45
                                                    } else {
                                                        v34 = v1
                                                        v46 := v5[v2]
                                                        _ = v46
                                                    }
                                                }
                                                v47 := GenS2{F0: (2 != 42), F1: v37}
                                                switch (fi31 & 127) {
                                                case 32768:
                                                    fmt.Println(v7)
                                                    v1 += 100
                                                default:
                                                    if v27 {
                                                        var v48 uint32 = 42
                                                        v48 += v48
                                                        v2 = (v2 + "\n\t")
                                                        var v49 uint16 = 0
                                                        if v27 {
                                                            var v50 uint32 = v48
                                                            delete(v5, v2)
                                                            v3 = append(v3, v2)
                                                            _ = v50
                                                        } else {
                                                            v34 += v1
                                                            switch v2 {
                                                            case "\n\t":
                                                                var v51 uint16 = (^(v49))
                                                                _ = v51
                                                            }
                                                        }
                                                        v52 := GenS2{F0: (!(-32768 >= -563)), F1: (v37 * 1)}
                                                        _ = v48
                                                        _ = v49
                                                        _ = v52
                                                    }
                                                }
                                                _ = v47
                                            }
                                        case "\n\t":
                                            var v53 float64 = ((-2.25 - float64(0.0)) * (1.5 + 0.1))
                                            _ = v53
                                        case "foo bar":
                                            fmt.Println(v2, v34, fi31)
                                        default:
                                            if ((!(!v27)) && (false && false)) {
                                                v27 = (-1 < 311)
                                            } else {
                                                v54 := GenS4{F0: 17, F1: 65535, F2: 32767}
                                                _ = v54
                                            }
                                        }
                                        _ = fi38
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                fmt.Println(v13, v13)
                                _ = v35
                            }
                            fmt.Println(v1, v1)
                            _ = v34
                        }
                    }
                    var v55 uint8 = 127
                    _ = v55
                }
                _ = fi31
            }
        }
    }
    v56 := GenS2{F0: v27, F1: 1}
    var v57 string = v2
    _ = len(v5)
    var v58 uint32 = 2147483647
    var v59 float32 = v7
    switch v57 {
    case "":
        for fi60 := 0; fi60 < 20; fi60++ {
            if (42 != 0) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v61 string = v2
                v57 = v61
                _ = v61
            }
            switch -363 {
            case 42:
                v3 = append(v3, "")
            case 32768:
                var v62 uint16 = 0
                _ = v62
            case -535:
                v2 = v2
            }
            v7 = v59
            _ = fi60
        }
    }
    if v27 {
        var v63 float32 = (float32(-2.25) / v59)
        v6[(^(v13))] = v27
        _ = v63
    } else {
        if true {
            v1 += (v1 + v1)
        } else {
            v64 := GenS5{F0: -1, F1: 43}
            switch v2 {
            case "\n\t":
                if (1 >= 873) {
                    if v27 {
                        v27 = (true && false)
                        v58 = (^(v58))
                        fmt.Println(v58)
                    }
                    var v65 int64 = -1
                    for fi66 := 0; fi66 < 2; fi66++ {
                        if true {
                            var v67 uint16 = 0
                            _ = v67
                        } else {
                            for fi68 := 0; fi68 < 16; fi68++ {
                                v69 := GenS6{F0: 255, F1: v27, F2: fi68}
                                var v70 int8 = (v1 | v1)
                                v71 := GenS7{F0: 0, F1: -1}
                                _ = fi68
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                            v72 := GenS8{F0: -81, F1: (v58 - v58), F2: 440}
                            var v73 uint64 = 508
                            _ = v72
                            _ = v73
                        }
                        if (!false) {
                            var v74 string = v57
                            _ = v74
                        }
                        v13 += fi66
                        _ = fi66
                    }
                    _ = v65
                }
            case "foo bar":
                switch ((v57 + "\n\t") + v57) {
                case "\n\t":
                    switch v2 {
                    case "":
                        fmt.Println(v1)
                    case "hello":
                        v57 = (("\n\t" + "foo bar") + v57)
                    case "\n\t":
                        var v75 int = (v13 & v13)
                        _ = v75
                    }
                }
            case "a":
                fmt.Println(v59, v2)
            }
            _ = v64
        }
        var v76 uint64 = 1
        _ = v76
    }
    switch "\n\t" {
    case "hello":
        v77 := GenS7{F0: 916, F1: 194}
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v7
}
