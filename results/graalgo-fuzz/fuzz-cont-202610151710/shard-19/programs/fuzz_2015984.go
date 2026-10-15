package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int64; F1 int8; F2 int }

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 uint8 = 1
    var v5 string = "a"
    v6 := make(map[int]int)
    v4 += v4
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    if true {
        if (!(42 == 0)) {
            delete(v3, (v5 + ((("hello" + v5) + v5) + (v5 + (v5 + v5)))))
        }
        var v9 float64 = 0.1
        _ = v9
    }
    v10 := GenS1{F0: false}
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: (((v4 ^ 42) <= 128) || (!(fi11 == fi11)))}
        _ = fi11
        _ = v12
    }
    var v13 bool = true
    v5 = ((v5 + (v5 + (v5 + "hello"))) + (("" + v5) + "hello"))
    delete(v3, v5)
    for fi14 := 0; fi14 < 20; fi14++ {
        v3[(v5 + v5)] = v5
        _ = fi14
    }
    switch "a" {
    case "hello":
        var v15 uint32 = uint32(v4)
        _ = v15
    case "foo bar":
        var v16 float64 = 0.1
        _ = v16
    case "\n\t":
        v17 := GenS1{F0: v13}
        _ = v17
    default:
        var v18 int32 = 2147483647
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if v13 {
            switch ((v5 + v5) + v5) {
            case "a":
                if (!v13) {
                    switch (fi19 >> ((710) & 63)) {
                    case 2147483647:
                        if (!v13) {
                            v20 := GenS1{F0: (v13 || v13)}
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v4 += v4
                                if v13 {
                                    switch v5 {
                                    case "foo bar":
                                        _ = len(v6)
                                        var v22 uint = 4294967295
                                        v23 := GenS1{F0: ((^(v22)) < 65535)}
                                        _ = v22
                                        _ = v23
                                    case "hello":
                                        v13 = true
                                        delete(v2, (fi21 * 4294967295))
                                    case "\n\t":
                                        v24 := GenS1{F0: v13}
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: (!(!(true && ((255 > 986) || (-1 > 1)))))}
                                    _ = v25
                                }
                                v26 := GenS2{F0: -1}
                                _ = fi21
                                _ = v26
                            }
                            var v27 float32 = (-2.25 / (((1e10 - 0.1) * (0.1 * ((0.0 * 0.1) - (((3.14159 - -2.25) + (-2.25 / (-2.25 + 3.14159))) + -2.25)))) * 1.5))
                            _ = v20
                            _ = v27
                        }
                    case 816:
                        switch fi19 {
                        case 277:
                            v13 = v13
                        case -285:
                            for fi28 := 0; fi28 < 19; fi28++ {
                                if v13 {
                                    v29 := GenS1{F0: (!(2 >= 578))}
                                    switch fi19 {
                                    case 42:
                                        if true {
                                            var v30 int8 = 0
                                            fmt.Println(v30, v5, fi28)
                                            var v31 string = "a"
                                            v32 := GenS1{F0: (42 != 1000)}
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        } else {
                                            var v33 uint64 = 18446744073709551615
                                            v34 := GenS2{F0: 573}
                                            if true {
                                                var v35 int32 = 722
                                                switch 946 {
                                                case 287:
                                                    fmt.Println(v33)
                                                case 2147483647:
                                                    v36 := GenS1{F0: (v13 || (2147483647 < (fi19 * fi28)))}
                                                    _ = v36
                                                case 42:
                                                    fmt.Println(v4, v13, fi28)
                                                }
                                                v33 = (v33 * (v33 + v33))
                                                _ = v35
                                            }
                                            var v37 int16 = -1
                                            _ = v33
                                            _ = v34
                                            _ = v37
                                        }
                                    default:
                                        var v38 bool = v13
                                        _ = v38
                                    }
                                    if v13 {
                                        v5 = "foo bar"
                                        v13 = (42 == 572)
                                    }
                                    _ = v29
                                }
                                var v39 uint8 = (v4 + (v4 - 215))
                                v40 := GenS3{F0: 592, F1: -9223372036854775808}
                                v41 := GenS3{F0: 977, F1: 0}
                                _ = fi28
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        default:
                            if (v4 >= (^(v4))) {
                                var v42 float64 = (((1e10 - ((0.1 - float64(1e10)) / -2.25)) + float64(-2.25)) + ((3.14159 * float64(1.5)) * float64(float32(3.14159))))
                                v5 = (v5 + "\n\t")
                                if ("\n\t" != v5) {
                                    if v13 {
                                        var v43 bool = ((fi19 * (^(fi19))) < fi19)
                                        switch "" {
                                        case "\n\t":
                                            for fi44 := 0; fi44 < 16; fi44++ {
                                                v13 = (0 > 0)
                                                v4 -= v4
                                                v4 -= 127
                                                _ = fi44
                                            }
                                        default:
                                            delete(v2, -890)
                                        }
                                        fmt.Println(fi19, v42)
                                        _ = v43
                                    }
                                    v45 := v3["a"]
                                    _ = v45
                                } else {
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        v42 += float64(0.0)
                                        var v47 bool = ((v5 + ("hello" + "hello")) <= v5)
                                        var v48 uint32 = 4294967295
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v42
                            } else {
                                fmt.Println(v5)
                                if len(v1) > 0 {
                                    v1[0] = 720
                                }
                            }
                        }
                    default:
                        var v49 float32 = 1e10
                        _ = v49
                    }
                    v50 := GenS3{F0: 2, F1: -555}
                    fmt.Println(v5, v5)
                    _ = v50
                } else {
                    switch v5 {
                    case "foo bar":
                        for fi51 := 0; fi51 < 10; fi51++ {
                            _ = len(v6)
                            _ = fi51
                        }
                    default:
                        var v52 uint64 = 1
                        _ = v52
                    }
                }
            case "foo bar":
                if ((v5 + "\n\t") > v5) {
                    var v53 string = v5
                    v13 = v13
                    _ = len(v1)
                    _ = v53
                }
            case "\n\t":
                var v54 int16 = 32767
                _ = v54
            default:
                var v55 uint32 = 4294967295
                _ = v55
            }
            var v56 int8 = -1
            if v13 {
                if false {
                    if (!(122 == 0)) {
                        if (0 >= 567) {
                            if (v13 || (!(42 <= 0))) {
                                switch fi19 {
                                case 128:
                                    v13 = (v13 && v13)
                                case -1:
                                    v57 := GenS1{F0: (-591 > 124)}
                                    _ = v57
                                case -9223372036854775808:
                                    var v58 uint8 = 42
                                    _ = v58
                                default:
                                    if false {
                                        delete(v3, ("a" + "\n\t"))
                                        if false {
                                            v59 := GenS4{F0: (fi19 * (fi19 + fi19)), F1: 275, F2: -1}
                                            var v60 float32 = -2.25
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                }
                                if (v56 == 127) {
                                    _ = len(v6)
                                    v56 += 5
                                } else {
                                    fmt.Println(v13)
                                }
                                v61 := GenS5{F0: 1, F1: 18446744073709551615, F2: 2}
                                _ = v61
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                v1 = append(v1, 270)
                            default:
                                var v62 bool = ((308 < 255) && true)
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 8; fi63++ {
                                switch fi63 {
                                case 4294967295:
                                    fmt.Println(fi63, fi19)
                                default:
                                    var v64 float32 = -2.25
                                    _ = v64
                                }
                                var v65 string = v5
                                _ = fi63
                                _ = v65
                            }
                        }
                        v13 = v13
                    } else {
                        v56 = (v56 >> ((2) & 63))
                        if false {
                            if v13 {
                                v66 := GenS5{F0: (^(v56)), F1: 2, F2: 65535}
                                v5 = ("" + (v5 + "hello"))
                                _ = v66
                            }
                            if (4294967295 <= 2) {
                                v5 = "\n\t"
                                if (!(true || v13)) {
                                    for fi67 := 0; fi67 < 14; fi67++ {
                                        v68 := GenS6{F0: 77, F1: 0}
                                        var v69 int8 = -1
                                        var v70 uint16 = 336
                                        v71 := GenS4{F0: fi19, F1: 0, F2: v69}
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            }
                            if false {
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                    var v73 uint8 = v4
                                    _ = fi72
                                    _ = v73
                                }
                                v13 = false
                            } else {
                                v13 = true
                            }
                        } else {
                            var v74 float32 = -2.25
                            _ = v74
                        }
                    }
                }
                v5 = (("\n\t" + (("\n\t" + ((v5 + "foo bar") + v5)) + ("\n\t" + "\n\t"))) + (v5 + ((v5 + v5) + (v5 + v5))))
                for fi75 := 0; fi75 < 10; fi75++ {
                    if v13 {
                        var v76 bool = true
                        _ = v76
                    }
                    v13 = (109 >= 642)
                    var v77 float32 = (1.5 * 0.0)
                    _ = fi75
                    _ = v77
                }
            } else {
                var v78 float64 = (-2.25 * (((3.14159 / ((0.0 / (((((0.0 / 1e10) - ((float64(1e10) / -2.25) + -2.25)) + (1e10 * ((1e10 - 0.1) + 0.0))) + 0.1) * 3.14159)) + (1e10 * (0.1 / 1.0)))) * ((3.14159 - 1e10) - 1.5)) * float64(0.1)))
                for fi79 := 0; fi79 < 12; fi79++ {
                    v56 = v56
                    if v13 {
                        var v80 int = fi19
                        _ = v80
                    }
                    switch 2147483647 {
                    case 128:
                        if true {
                            var v81 uint16 = 32767
                            _ = v81
                        }
                    case -1:
                        v5 = (v5 + v5)
                    }
                    v82 := GenS3{F0: 2, F1: -1}
                    _ = fi79
                    _ = v82
                }
                _ = v78
            }
            _ = v56
        }
        var v83 float64 = 1.5
        for fi84 := 0; fi84 < 8; fi84++ {
            v85 := GenS7{F0: -9223372036854775808, F1: 1, F2: fi19}
            var v86 int64 = -1
            _ = fi84
            _ = v85
            _ = v86
        }
        _ = fi19
        _ = v83
    }
    v4 = v4
    _ = len(v1)
    v5 = v5
    if v13 {
        for fi87 := 0; fi87 < 5; fi87++ {
            fmt.Println(v13)
            _ = fi87
        }
    } else {
        var v88 uint8 = v4
        for fi89 := 0; fi89 < 17; fi89++ {
            fmt.Println(v88, v5, v13)
            var v90 uint = 65535
            _ = fi89
            _ = v90
        }
        _ = v88
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
