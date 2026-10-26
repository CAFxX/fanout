package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint = 18446744073709551615
    var v3 string = ""
    v4 := []string{"foo bar", v3}
    v5 := make(map[uint32]bool)
    v6 := []string{"", v3}
    v1 += (v1 - v1)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 186, F1: false, F2: -9223372036854775808}
        if false {
            var v9 int32 = 1
            v3 = v3
            if (false && true) {
                if true {
                    switch ((((v3 + "foo bar") + (((("foo bar" + (v3 + "hello")) + "") + (v3 + (((v3 + v3) + v3) + v3))) + "a")) + ("" + (v3 + v3))) + v3) {
                    case "hello":
                        if true {
                            for fi10 := 0; fi10 < 10; fi10++ {
                                if false {
                                    var v11 bool = false
                                    if v11 {
                                        if v11 {
                                            var v12 int16 = 0
                                            _ = v12
                                        }
                                        v13 := GenS1{F0: v2, F1: v11, F2: int64(fi10)}
                                        _ = v13
                                    }
                                    fmt.Println(v11, fi10, v11)
                                    _ = v11
                                }
                                v2 = v2
                                var v14 float32 = 1.5
                                _ = fi10
                                _ = v14
                            }
                            var v15 uint32 = 4294967295
                            v9 += v9
                            _ = v15
                        }
                    default:
                        v16 := GenS2{F0: v9, F1: -528}
                        _ = v16
                    }
                }
                switch 0 {
                case 4294967295:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if (1 >= -9223372036854775808) {
                            v18 := GenS1{F0: v2, F1: (int(v9) != (v1 % (((fi17 << (((v2 | v2)) & 63))) | 1))), F2: -1}
                            if (145 >= 25) {
                                var v19 uint16 = 0
                                _ = v19
                            } else {
                                v20 := GenS3{F0: 503, F1: 65535}
                                v3 = "\n\t"
                                v2 -= (v2 + 42)
                                _ = v20
                            }
                            var v21 uint64 = 761
                            _ = v18
                            _ = v21
                        } else {
                            v22 := GenS3{F0: uint16(v2), F1: 2}
                            fmt.Println(fi7)
                            if len(v6) > 0 {
                                v6[0] = v3
                            }
                            _ = v22
                        }
                        v3 = ""
                        if (!(((float64((1e10 * 3.14159)) * 1.5) / 0.1) != (-2.25 / (float64(-2.25) * 0.1)))) {
                            _ = len(v6)
                            v9 += v9
                        }
                        v6 = append(v6, "a")
                        _ = fi17
                    }
                case -204:
                    switch (v3 + "") {
                    case "":
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    case "hello":
                        var v24 float32 = (float32(float64(0.1)) * float32((float64(((0.0 * (0.0 + float32(0.1))) - (((((((-2.25 + 3.14159) + 0.1) + 0.1) - float32((((1.5 / 1e10) * (((((3.14159 + 0.1) * (float64(3.14159) - (3.14159 + 0.0))) / ((0.1 + 3.14159) * ((float64((0.1 + 0.1)) - float64(3.14159)) + float64((1e10 / 1e10))))) * (float64(0.0) / 1.0)) + 0.0)) - (-2.25 + ((((-2.25 / 0.1) - 0.1) - 1.5) * 0.1))))) * (1.5 - 0.1)) + 1.5) + 1e10))) * 3.14159)))
                        _ = v24
                    }
                }
                v25 := GenS3{F0: 339, F1: 255}
                _ = v25
            }
            _ = v9
        } else {
            v1 = fi7
            var v26 int8 = 42
            v27 := GenS4{F0: 128, F1: (v2 | 564)}
            _ = v26
            _ = v27
        }
        _ = fi7
        _ = v8
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        if false {
            for fi29 := 0; fi29 < 6; fi29++ {
                switch 128 {
                case -9223372036854775808:
                    var v30 int32 = -1
                    _ = v30
                }
                var v31 uint64 = 4294967295
                _ = fi29
                _ = v31
            }
            var v32 string = "hello"
            _ = v32
        } else {
            v33 := GenS4{F0: 42, F1: v2}
            var v34 int64 = -672
            _ = v33
            _ = v34
        }
        v2 = v2
        v5[2147483648] = true
        _ = fi28
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 int8 = 42
        switch v3 {
        case "\n\t":
            v37 := GenS5{F0: v36}
            _ = v37
        case "hello":
            for fi38 := 0; fi38 < 13; fi38++ {
                if true {
                    if true {
                        v36 = 100
                    } else {
                        if false {
                            var v39 bool = true
                            var v40 uint64 = 18446744073709551615
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(v2, v36, v3)
                        fmt.Println(v2, v1, v36)
                        v3 = v3
                        v6 = append(v6, "\n\t")
                        switch "\n\t" {
                        case "a":
                            v2 -= v2
                        case "\n\t":
                            if false {
                                v2 += (v2 + 0)
                                switch ("hello" + v3) {
                                case "a":
                                    var v41 int64 = 0
                                    _ = v41
                                case "hello":
                                    var v42 bool = ((v36 >> (((^(v2))) & 63)) < (v36 % (((-(v36))) | 1)))
                                    _ = v42
                                case "":
                                    fmt.Println(v1)
                                }
                                var v43 uint8 = 1
                                _ = v43
                            }
                        case "":
                            var v44 int8 = (v36 << (((v2 % ((255) | 1))) & 63))
                            _ = v44
                        default:
                            if (false || true) {
                                v2 = v2
                            } else {
                                _ = len(v5)
                                var v45 uint = (v2 + 392)
                                var v46 int = (fi35 / ((-534) | 1))
                                _ = v45
                                _ = v46
                            }
                        }
                    }
                    if (!((v2 % ((v2) | 1)) == v2)) {
                        _ = len(v6)
                        _ = len(v4)
                        var v47 uint8 = 42
                        var v48 uint64 = 772
                        switch (v3 + "\n\t") {
                        case "\n\t":
                            if (v3 == v3) {
                                var v49 uint16 = 1
                                v2 = v2
                                v3 = v3
                                _ = v49
                            } else {
                                var v50 float64 = (((1.5 / (3.14159 - 0.0)) + -2.25) * (0.0 - float64(1.5)))
                                _ = v50
                            }
                        case "":
                            v51 := GenS6{F0: false, F1: 386}
                            _ = v51
                        default:
                            if (!(v3 > v3)) {
                                v3 = v3
                            }
                        }
                        _ = v47
                        _ = v48
                    }
                } else {
                    v5[2] = true
                    if ((42 <= 755) && (265 > 2)) {
                        var v52 float64 = float64(1e10)
                        var v53 uint64 = 888
                        switch -129 {
                        case 2147483647:
                            v1 += fi38
                        }
                        fmt.Println(v52)
                        fmt.Println(v1, v52)
                        _ = v52
                        _ = v53
                    } else {
                        var v54 bool = false
                        var v55 bool = (42 > 0)
                        v1 += -67
                        var v56 uint = 2
                        v2 = (v2 * (v2 / ((v2) | 1)))
                        var v57 uint32 = 697
                        _ = v54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                }
                var v58 int64 = 42
                delete(v5, 461)
                v59 := GenS1{F0: (v2 * v2), F1: true, F2: (v58 >> (((v2 & v2)) & 63))}
                var v60 uint8 = 0
                v1 = (-(fi38))
                if len(v4) > 0 {
                    v4[0] = (v3 + (("foo bar" + v3) + ""))
                }
                _ = fi38
                _ = v58
                _ = v59
                _ = v60
            }
        default:
            var v61 uint8 = 219
            _ = v61
        }
        switch (v3 + v3) {
        case "a":
            for fi62 := 0; fi62 < 14; fi62++ {
                if true {
                    if (2147483648 > 725) {
                        v3 = (v3 + v3)
                        v2 = (^(v2))
                    }
                }
                _ = fi62
            }
        case "\n\t":
            for fi63 := 0; fi63 < 16; fi63++ {
                var v64 int = 501
                _ = fi63
                _ = v64
            }
        case "foo bar":
            for fi65 := 0; fi65 < 10; fi65++ {
                switch fi65 {
                case -9223372036854775808:
                    v66 := GenS5{F0: (^(v36))}
                    _ = v66
                case -534:
                    v6 = append(v6, (v3 + (("\n\t" + v3) + "foo bar")))
                default:
                    v1 = 551
                }
                v67 := GenS5{F0: (v36 / ((1) | 1))}
                v1 += fi35
                v1 -= (fi35 - -353)
                var v68 int32 = -683
                var v69 uint64 = 70
                var v70 uint32 = 42
                var v71 uint64 = (v69 - (v69 + v69))
                v72 := GenS7{F0: v2, F1: 4}
                fmt.Println(v1, fi65, v68)
                switch "foo bar" {
                case "foo bar":
                    v73 := GenS7{F0: (v2 - (^(v2))), F1: 127}
                    _ = v73
                case "\n\t":
                    var v74 uint8 = 127
                    _ = v74
                }
                var v75 uint8 = 42
                _ = fi65
                _ = v67
                _ = v68
                _ = v69
                _ = v70
                _ = v71
                _ = v72
                _ = v75
            }
        }
        _ = fi35
        _ = v36
    }
    v3 = (v3 + (((v3 + "hello") + v3) + v3))
    var v76 bool = ((v1 % ((v1) | 1)) > 2147483647)
    v76 = (4294967295 >= 660)
    var v77 uint64 = 326
    delete(v5, 4294967295)
    for fi78 := 0; fi78 < 18; fi78++ {
        if false {
            var v79 uint32 = 42
            v76 = (v77 < 0)
            for fi80 := 0; fi80 < 12; fi80++ {
                if true {
                    var v81 string = v3
                    if v76 {
                        _ = len(v4)
                        if v76 {
                            v3 = (v81 + "\n\t")
                            var v82 int8 = 127
                            var v83 int16 = int16(v77)
                            _ = v82
                            _ = v83
                        } else {
                            fmt.Println(fi80, v3)
                        }
                        var v84 uint = 255
                        _ = v84
                    }
                    _ = v81
                } else {
                    _ = len(v5)
                    var v85 int8 = -1
                    _ = v85
                }
                delete(v5, (v79 >> (((^(v2))) & 63)))
                fmt.Println(fi78, fi78, v2)
                var v86 uint = (v2 << (((v2 | (v2 & 65535))) & 63))
                _ = fi80
                _ = v86
            }
            _ = v79
        }
        v87 := GenS7{F0: v2, F1: 134}
        _ = fi78
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
