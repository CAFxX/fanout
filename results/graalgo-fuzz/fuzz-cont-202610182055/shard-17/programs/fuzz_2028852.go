package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint64; F1 int32; F2 int8 }
func genFn0(p0 uint, p1 int16, p2 int16) int32 {
    var v6 float64 = 0.0
    p2 = int16(p0)
    var v7 uint = 4294967295
    fmt.Println(v7, p2, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -471
}
func genFn1(p0 int, p1 uint, p2 bool) uint64 {
    v8 := GenS1{F0: 32767, F1: 127}
    var v9 int32 = 59
    v10 := GenS1{F0: 390, F1: 127}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    if p2 {
        var v13 string = "hello"
        p1 += (^(p1))
        var v14 bool = (94 != uint64(p1))
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    var v3 int64 = (-(v2))
    v4 := []string{"hello", "a"}
    v5 := make(map[int]int)
    if false {
        var v15 uint64 = 262
        if (!(false || (genFn0(639, 0, 1000) > -2147483648))) {
            if (!false) {
                v16 := GenS2{F0: 1000, F1: 127}
                v2 = (^(v3))
                _ = v16
            }
            v17 := GenS1{F0: 0, F1: 128}
            switch 1 {
            case 32768:
                var v18 uint64 = 4294967295
                _ = v18
            default:
                _ = len(v4)
            }
            _ = v17
        } else {
            v19 := GenS3{F0: (v2 | -229)}
            if true {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v2, v15)
                    var v21 uint8 = 1
                    _ = fi20
                    _ = v21
                }
            }
            _ = v19
        }
        _ = v15
    } else {
        if false {
            switch "hello" {
            case "foo bar":
                v2 = (v2 / ((v3) | 1))
            case "a":
                fmt.Println(v3)
            }
            if false {
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch "" {
                    case "a":
                        var v23 int8 = -1
                        _ = v23
                    default:
                        v24 := GenS4{F0: 2, F1: 0}
                        _ = v24
                    }
                    _ = fi22
                }
                var v25 uint32 = 2147483647
                v26 := GenS1{F0: 458, F1: 176}
                _ = v25
                _ = v26
            }
            switch 918 {
            case 0:
                fmt.Println(v2, v2)
            case 851:
                var v27 uint32 = 4294967295
                _ = v27
            default:
                if true {
                    v2 = (v3 ^ v3)
                    fmt.Println(v2)
                    for fi28 := 0; fi28 < 19; fi28++ {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            var v30 int64 = 42
                            v31 := GenS5{F0: 2147483648}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        fmt.Println(v3, v3, v3)
                        v32 := GenS1{F0: -64, F1: 1}
                        _ = fi28
                        _ = v32
                    }
                }
            }
        } else {
            v2 = 42
            v2 = (-(v3))
            v33 := GenS2{F0: 104, F1: 42}
            _ = v33
        }
        if (0 >= 42) {
            switch "\n\t" {
            case "":
                v34 := v1["hello"]
                _ = v34
            case "hello":
                if (32767 <= 629) {
                    fmt.Println(v3)
                    for fi35 := 0; fi35 < 2; fi35++ {
                        switch (fi35 / (((fi35 << ((1) & 63))) | 1)) {
                        case -739:
                            for fi36 := 0; fi36 < 3; fi36++ {
                                switch (^(fi35)) {
                                case -233:
                                    var v37 uint8 = 74
                                    _ = v37
                                case 847:
                                    v2 = (-(v2))
                                case 9223372036854775807:
                                    v38 := GenS6{F0: 98}
                                    _ = v38
                                default:
                                    if (65535 > 2) {
                                        var v39 uint = 578
                                        v40 := GenS3{F0: v3}
                                        _ = v39
                                        _ = v40
                                    } else {
                                        if (0 < 0) {
                                            v3 -= (v2 % ((v2) | 1))
                                            v41 := GenS6{F0: 0}
                                            v2 += (v3 | (v3 | (v3 >> ((0) & 63))))
                                            _ = v41
                                        } else {
                                            if false {
                                                v42 := GenS1{F0: -92, F1: 2}
                                                if (!(v2 != v3)) {
                                                    v43 := v1[((((("\n\t" + "\n\t") + "foo bar") + "") + (("foo bar" + ("foo bar" + (("hello" + "a") + ""))) + "")) + "hello")]
                                                    v43 = v43
                                                    _ = v43
                                                }
                                                _ = v42
                                            }
                                        }
                                        v3 = 9223372036854775807
                                    }
                                }
                                if (0 >= 127) {
                                    if (false && (42 != 555)) {
                                        fmt.Println(fi36, v2)
                                    }
                                    v44 := GenS4{F0: 42, F1: 941}
                                    v3 += (v3 % ((int64(fi35)) | 1))
                                    v45 := GenS7{F0: 367, F1: (!(v2 != v3))}
                                    _ = v44
                                    _ = v45
                                }
                                var v46 float32 = (1.5 - (1e10 - 0.1))
                                _ = fi36
                                _ = v46
                            }
                        case 584:
                            fmt.Println(v2, v2, v3)
                        default:
                            var v47 bool = (123 <= 255)
                            _ = v47
                        }
                        _ = fi35
                    }
                }
            }
        } else {
            var v48 string = "foo bar"
            v49 := v1["\n\t"]
            _ = v48
            _ = v49
        }
    }
    if false {
        fmt.Println(v3)
        var v50 float64 = 3.14159
        if (true || false) {
            v51 := GenS3{F0: -9223372036854775808}
            v4 = append(v4, "a")
            v52 := GenS7{F0: 42, F1: false}
            _ = v51
            _ = v52
        }
        _ = v50
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ((((("" + ("hello" + "hello")) + "") + "hello") + ("foo bar" + "\n\t")) + "") {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = ("foo bar" + "")
        }
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        for fi53 := 0; fi53 < 9; fi53++ {
            var v54 int32 = 2147483647
            if true {
                v2 -= 1000
                if false {
                    fmt.Println(v54)
                    v54 -= (v54 >> ((1) & 63))
                    delete(v1, ("foo bar" + "a"))
                }
                v4 = append(v4, (("hello" + ("hello" + "")) + "a"))
            }
            v3 = (^(v2))
            _ = fi53
            _ = v54
        }
    }
    var v55 float32 = 3.14159
    v56 := GenS1{F0: 1000, F1: 0}
    if true {
        v57 := GenS8{F0: 2, F1: 2, F2: 65535}
        v55 += (v55 / v55)
        if (127 != 255) {
            for fi58 := 0; fi58 < 3; fi58++ {
                v3 = v3
                switch (fi58 >> ((894) & 63)) {
                case 641:
                    var v59 bool = (!true)
                    _ = v59
                case 4294967295:
                    _ = len(v1)
                case -9223372036854775808:
                    var v60 int8 = -128
                    _ = v60
                default:
                    if (!false) {
                        if false {
                            var v61 uint8 = 128
                            var v62 string = ("foo bar" + (("a" + "foo bar") + (("foo bar" + (("" + "hello") + "\n\t")) + "foo bar")))
                            v3 = 0
                            _ = v61
                            _ = v62
                        }
                        if true {
                            if (0 == -1) {
                                var v63 int8 = 100
                                _ = v63
                            }
                        }
                        if (2 <= 201) {
                            v55 = ((v55 + ((0.0 / 0.1) / v55)) + ((v55 + v55) / v55))
                        } else {
                            v3 = (^(v3))
                            _ = len(v4)
                            v64 := GenS9{F0: 4294967295, F1: 42, F2: -1}
                            _ = v64
                        }
                    }
                }
                fmt.Println(v2, v3, fi58)
                _ = fi58
            }
            var v65 uint = 2
            var v66 uint8 = 128
            _ = v65
            _ = v66
        }
        _ = v57
    } else {
        var v67 int8 = -72
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = v67
    }
    for fi68 := 0; fi68 < 2; fi68++ {
        if (!true) {
            v69 := GenS7{F0: 42, F1: true}
            for fi70 := 0; fi70 < 13; fi70++ {
                v3 += (^(v3))
                _ = fi70
            }
            var v71 uint64 = 0
            _ = v69
            _ = v71
        } else {
            var v72 uint32 = 2
            _ = v72
        }
        var v73 int8 = 127
        _ = fi68
        _ = v73
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        if (3.14159 > v55) {
            v55 = float32((0.1 * (0.0 + (0.0 + 1e10))))
            v75 := GenS4{F0: 87, F1: 1}
            for fi76 := 0; fi76 < 15; fi76++ {
                var v77 uint16 = 255
                _ = fi76
                _ = v77
            }
            _ = v75
        }
        if true {
            v2 += -534
            for fi78 := 0; fi78 < 3; fi78++ {
                v3 = v2
                fmt.Println(v55, v2, fi78)
                v2 = v3
                _ = fi78
            }
        }
        v79 := GenS5{F0: 0}
        _ = fi74
        _ = v79
    }
    _ = len(v5)
    v2 -= (-(v2))
    var v80 uint32 = 1
    v80 = 357
    var v81 float32 = float32(-2.25)
    var v82 int64 = v2
    v82 -= v3
    if (100 < -128) {
        v83 := GenS2{F0: -1, F1: 241}
        _ = v83
    } else {
        v84 := GenS5{F0: v80}
        _ = v84
    }
    v2 = 207
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v80
    _ = v81
    _ = v82
}
