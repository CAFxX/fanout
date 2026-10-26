package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 int; F2 uint32 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 uint; F1 int16 }
type GenS13 struct { F0 uint16; F1 uint64 }
type GenS14 struct { F0 uint32; F1 int; F2 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 int; F2 int16 }
func genFn0(p0 uint16, p1 int16) uint {
    var v7 uint16 = p0
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 float32 = (1e10 / 1.0)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 uint = 255
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int = (fi12 ^ fi12)
            var v15 float64 = (float64((3.14159 / -2.25)) / (1e10 + 3.14159))
            v16 := []int{}
            if len(v16) > 0 {
                v16[0] = fi12
            }
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS1{F0: 4294967295, F1: p1, F2: 128}
            p1 = (p1 | (p1 * int16(p0)))
            v19 := GenS2{F0: -748}
            v20 := GenS3{F0: 100}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        if (0.1 < -2.25) {
            v21 := GenS4{F0: fi12}
            _ = v21
        } else {
            fmt.Println(p0, v11)
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = -32768
    var v3 uint16 = 2
    var v4 int = 514
    var v5 uint64 = 2
    v6 := []uint32{283, 1}
    fmt.Println(v3)
    var v22 uint = genFn0((v3 + 255), 338)
    switch 128 {
    case 2147483647:
        var v23 float64 = 3.14159
        _ = v23
    case 127:
        fmt.Println(v4)
    case 1:
        v24 := GenS5{F0: -1, F1: (("hello" + ((("\n\t" + ("\n\t" + "hello")) + "\n\t") + ("\n\t" + "hello"))) + "foo bar"), F2: 398}
        _ = v24
    }
    v2 = -1
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS6{F0: v2}
        switch fi25 {
        case -949:
            v27 := GenS5{F0: -674, F1: ((("" + ("hello" + "\n\t")) + "") + "\n\t"), F2: -700}
            _ = v27
        case 647:
            if true {
                v28 := GenS1{F0: 743, F1: int16(v4), F2: 229}
                v1 = (v1 + 1.5)
                v29 := GenS5{F0: 495, F1: "", F2: 9223372036854775807}
                _ = v28
                _ = v29
            }
        }
        _ = fi25
        _ = v26
    }
    v22 = v22
    var v30 string = ("" + "a")
    if (1e10 >= 1e10) {
        v31 := GenS7{F0: 42, F1: v4, F2: uint32(v5)}
        _ = v31
    } else {
        if (191 == 133) {
            if len(v6) > 0 {
                v6[0] = 42
            }
        } else {
            v30 = ("a" + (v30 + v30))
            fmt.Println(v3, v1)
        }
        v22 += 18446744073709551615
        for fi32 := 0; fi32 < 17; fi32++ {
            v33 := GenS8{F0: 937, F1: (v3 + v3)}
            v4 = 2147483647
            if ((!true) || (99 != 255)) {
                _ = len(v6)
            }
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v22 = genFn0((^(v3)), (v2 ^ (v2 - v2)))
        for fi35 := 0; fi35 < 3; fi35++ {
            if (true || (false || false)) {
                v36 := GenS3{F0: -4}
                _ = v36
            } else {
                switch fi34 {
                case 32768:
                    if (0 == 42) {
                        v22 = (v22 << (((^(v22))) & 63))
                        fmt.Println(v22, v3, v1)
                    } else {
                        v22 += genFn0((v3 / (((v3 + 42)) | 1)), (v2 ^ 32767))
                        if (-1 <= 455) {
                            switch int(v5) {
                            case 4294967295:
                                if (1 < 255) {
                                    switch "" {
                                    case "hello":
                                        var v37 string = v30
                                        _ = v37
                                    case "a":
                                        switch v4 {
                                        case 539:
                                            fmt.Println(v5, v22, v1)
                                        }
                                    case "foo bar":
                                        fmt.Println(v1, v3)
                                    default:
                                        v38 := GenS9{F0: 127, F1: 1000}
                                        _ = v38
                                    }
                                    v4 = (fi35 ^ (fi34 >> (((v22 * (^(v22)))) & 63)))
                                } else {
                                    v39 := GenS6{F0: (v2 >> ((v22) & 63))}
                                    var v40 int64 = 1
                                    v41 := GenS7{F0: v3, F1: (^(v4)), F2: 327}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                            case 134:
                                var v42 uint16 = v3
                                _ = v42
                            }
                        } else {
                            v22 += v22
                            v30 = v30
                            fmt.Println(v3)
                        }
                    }
                case 9223372036854775807:
                    var v43 uint = v22
                    _ = v43
                default:
                    v30 = (v30 + "hello")
                }
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
                if ("a" != "hello") {
                    switch fi34 {
                    case 127:
                        switch v30 {
                        case "foo bar":
                            v30 = ""
                            if false {
                                if false {
                                    _ = len(v6)
                                } else {
                                    if (1 != 0) {
                                        v3 = v3
                                        if false {
                                            if ((v2 * (v2 / ((v2) | 1))) < (v2 / ((v2) | 1))) {
                                                v30 = v30
                                                if len(v6) > 0 {
                                                    v6[0] = 2
                                                }
                                                var v44 uint64 = v5
                                                if len(v6) > 0 {
                                                    v6[0] = 2147483647
                                                }
                                                var v45 uint64 = 42
                                                switch v30 {
                                                case "":
                                                    v30 = v30
                                                case "foo bar":
                                                    v44 -= v5
                                                }
                                                _ = v44
                                                _ = v45
                                            }
                                            v4 -= (fi35 - fi35)
                                        }
                                        v46 := GenS10{F0: 1}
                                        _ = v46
                                    }
                                    fmt.Println(v22)
                                }
                            }
                        case "a":
                            v47 := GenS6{F0: v2}
                            _ = v47
                        case "\n\t":
                            switch 0 {
                            case 127:
                                if false {
                                    var v48 int8 = -122
                                    _ = v48
                                } else {
                                    var v49 float32 = (1.5 - (-2.25 * v1))
                                    _ = v49
                                }
                            case 586:
                                if (float64(v1) > 3.14159) {
                                    switch (v30 + v30) {
                                    case "a":
                                        v1 = ((3.14159 + 1.5) - 3.14159)
                                    case "":
                                        var v50 int16 = -510
                                        _ = v50
                                    case "foo bar":
                                        if true {
                                            v3 -= v3
                                        } else {
                                            if ("hello" > v30) {
                                                switch (v4 % (((^(fi34))) | 1)) {
                                                case 959:
                                                    v4 = (v4 + 4294967295)
                                                default:
                                                    _ = len(v6)
                                                }
                                            } else {
                                                var v51 int = (fi35 & 256)
                                                v52 := GenS11{F0: v30, F1: (^(v3))}
                                                if true {
                                                    v5 -= 629
                                                    _ = len(v6)
                                                    v3 -= v3
                                                } else {
                                                    if (1 >= 2147483647) {
                                                        v53 := GenS12{F0: 255, F1: (^(v2))}
                                                        v54 := GenS11{F0: v30, F1: (v3 | (v3 % (((v3 >> ((42) & 63))) | 1)))}
                                                        if (1.5 != v1) {
                                                            v1 = 3.14159
                                                            v51 = fi34
                                                            if (!false) {
                                                                var v55 int = fi34
                                                                _ = v55
                                                            } else {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            }
                                                        } else {
                                                            var v56 uint16 = v3
                                                            v6 = append(v6, 1)
                                                            _ = v56
                                                        }
                                                        fmt.Println(v5, v2)
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v4 += (^(fi34))
                                                        var v57 int16 = (v2 & (^(v2)))
                                                        _ = v57
                                                    }
                                                }
                                                var v58 uint64 = uint64(v3)
                                                v1 -= v1
                                                _ = v51
                                                _ = v52
                                                _ = v58
                                            }
                                            var v59 string = v30
                                            var v60 int32 = -1
                                            _ = v59
                                            _ = v60
                                        }
                                    default:
                                        v61 := GenS13{F0: (v3 - v3), F1: v5}
                                        _ = v61
                                    }
                                    var v62 uint64 = 0
                                    v3 = v3
                                    _ = v62
                                } else {
                                    var v63 uint = uint(v3)
                                    _ = v63
                                }
                            default:
                                var v64 float64 = 0.0
                                fmt.Println(v30, v64, v30)
                                _ = v64
                            }
                        default:
                            var v65 string = (("hello" + "") + ("hello" + (v30 + v30)))
                            _ = v65
                        }
                    }
                    switch "a" {
                    case "":
                        var v66 float64 = float64(v1)
                        _ = v66
                    case "\n\t":
                        if false {
                            var v67 bool = false
                            v68 := GenS9{F0: 127, F1: 42}
                            _ = v67
                            _ = v68
                        }
                    case "hello":
                        fmt.Println(v5, v22)
                    default:
                        var v69 uint8 = 1
                        _ = v69
                    }
                    v5 = v5
                }
                v1 -= 0.1
            }
            _ = fi35
        }
        _ = fi34
    }
    v3 = 255
    switch v4 {
    case 4294967295:
        if false {
            var v70 int32 = -1
            _ = v70
        } else {
            v22 = (^(v22))
            if true {
                if false {
                    if true {
                        v71 := GenS14{F0: 664, F1: 42, F2: 255}
                        _ = len(v6)
                        _ = v71
                    } else {
                        fmt.Println(v3)
                        if (!false) {
                            if (((v5 + v5) != 683) || false) {
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi72
                                }
                                v73 := GenS10{F0: 184}
                                if (!((0.0 + 0.0) <= 3.14159)) {
                                    var v74 int32 = -1
                                    _ = v74
                                }
                                _ = v73
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = 65535
                                }
                                v4 = (v4 | (v4 >> (((v22 >> (((v22 ^ genFn0((v3 / ((uint16(v22)) | 1)), 42))) & 63))) & 63)))
                            }
                            switch v30 {
                            case "hello":
                                var v75 uint16 = (v3 + (v3 >> (((v22 | 4294967295)) & 63)))
                                _ = v75
                            default:
                                fmt.Println(v1, v4, v5)
                            }
                        }
                        switch v4 {
                        case 32767:
                            _ = len(v6)
                        }
                    }
                }
            } else {
                v22 = 4294967295
                var v76 int16 = v2
                var v77 int = (v4 % (((-(v4))) | 1))
                _ = v76
                _ = v77
            }
            var v78 float64 = (0.0 * ((1.5 + 3.14159) + (3.14159 * ((((0.0 - 0.1) - 0.0) + ((3.14159 + (3.14159 / 1e10)) / float64((float32(float64(float32(1e10))) + 0.1)))) / (((float64(1e10) + 1.5) - 1.5) / float64(3.14159))))))
            _ = v78
        }
    case 32767:
        v2 += -1
    default:
        v79 := GenS4{F0: v4}
        _ = v79
    }
    var v80 int16 = v2
    var v81 uint32 = 65535
    v82 := GenS6{F0: v2}
    switch (v30 + v30) {
    case "\n\t":
        if (v5 != (v5 ^ 42)) {
            var v83 int8 = -19
            v84 := GenS15{F0: (-32768 >= (v2 / ((v80) | 1)))}
            _ = v83
            _ = v84
        }
    case "foo bar":
        for fi85 := 0; fi85 < 15; fi85++ {
            v22 -= v22
            _ = fi85
        }
    case "":
        switch v30 {
        case "":
            var v86 int = v4
            _ = v86
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = v81
            }
        default:
            var v87 uint32 = (v81 % (((v81 + v81)) | 1))
            _ = v87
        }
    }
    var v88 string = "hello"
    if false {
        if len(v6) > 0 {
            v6[0] = 62
        }
        for fi89 := 0; fi89 < 16; fi89++ {
            v30 = "a"
            var v90 int32 = 470
            _ = fi89
            _ = v90
        }
    }
    if true {
        var v91 bool = (("hello" + v30) == "foo bar")
        v92 := GenS16{F0: int8(v22), F1: (v4 << ((v22) & 63)), F2: (v2 + int16(v22))}
        _ = v91
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v81
    _ = v82
    _ = v88
}
