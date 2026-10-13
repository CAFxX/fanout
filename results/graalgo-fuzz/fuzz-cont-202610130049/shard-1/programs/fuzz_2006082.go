package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 int8; F2 string }
type GenS9 struct { F0 int32; F1 uint32 }
func genFn0() int16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 808, F1: "\n\t"}
        if len(v7) > 0 {
            v7[0] = fi8
        }
        if (2 != 1) {
            if (255 == 42) {
                switch "hello" {
                case "\n\t":
                    var v10 uint16 = 255
                    _ = v10
                case "a":
                    var v11 int = -231
                    _ = v11
                default:
                    switch (fi8 / ((127) | 1)) {
                    case -10:
                        for fi12 := 0; fi12 < 11; fi12++ {
                            var v13 uint8 = 45
                            v14 := GenS1{F0: -373, F1: ("foo bar" + ("" + ""))}
                            v15 := GenS1{F0: -996, F1: "\n\t"}
                            if false {
                                v13 = v13
                                var v16 float64 = 1.5
                                var v17 int = -9223372036854775808
                                _ = v16
                                _ = v17
                            }
                            if true {
                                if false {
                                    v13 -= (^(v13))
                                    _ = len(v7)
                                } else {
                                    v13 += v13
                                }
                                var v18 uint = 111
                                _ = v18
                            } else {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    var v19 bool = (false || true)
                                    _ = v19
                                case "\n\t":
                                    var v20 uint16 = 255
                                    _ = v20
                                default:
                                    var v21 uint32 = 1
                                    _ = v21
                                }
                                if (2147483647 == 913) {
                                    v13 = v13
                                }
                                v22 := GenS1{F0: 597, F1: "a"}
                                _ = v22
                            }
                            v23 := GenS1{F0: -602, F1: ("hello" + "hello")}
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v23
                        }
                    case 702:
                        switch "hello" {
                        case "hello":
                            switch "hello" {
                            case "foo bar":
                                v24 := GenS2{F0: 65535}
                                _ = v24
                            case "hello":
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v26 uint64 = 42
                                    var v27 float32 = float32(-2.25)
                                    if false {
                                        v27 = v27
                                    } else {
                                        if false {
                                            fmt.Println(fi25, v26)
                                            if (3.14159 <= float64(v27)) {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                fmt.Println(v26, fi25, v27)
                                                if (343 >= -32768) {
                                                    var v28 int32 = 965
                                                    _ = v28
                                                }
                                            } else {
                                                if len(v7) > 0 {
                                                    v7[0] = fi25
                                                }
                                                if false {
                                                    v7 = append(v7, fi8)
                                                    switch (fi25 * fi8) {
                                                    case 1:
                                                        v29 := GenS3{F0: -714, F1: (3.14159 < 0.0)}
                                                        _ = v29
                                                    }
                                                }
                                            }
                                            var v30 string = ("\n\t" + ((((("" + "") + "hello") + "\n\t") + "a") + (("hello" + "foo bar") + ("a" + "a"))))
                                            _ = v30
                                        } else {
                                            if true {
                                                var v31 int32 = 50
                                                fmt.Println(v31, fi25)
                                                _ = v31
                                            } else {
                                                v26 += v26
                                                var v32 float32 = 1e10
                                                v27 = (3.14159 - 1.5)
                                                _ = v32
                                            }
                                            var v33 float64 = (1e10 * 1.5)
                                            var v34 int = -226
                                            v7 = append(v7, 42)
                                            var v35 int32 = 457
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        v27 += v27
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case "a":
                            if (2 == 440) {
                                var v36 string = "\n\t"
                                if len(v7) > 0 {
                                    v7[0] = (^(fi8))
                                }
                                _ = v36
                            } else {
                                if (false || false) {
                                    if true {
                                        var v37 int8 = -18
                                        _ = v37
                                    } else {
                                        var v38 string = "foo bar"
                                        _ = v38
                                    }
                                    var v39 int64 = 1
                                    switch "\n\t" {
                                    case "a":
                                        if (1.5 >= -2.25) {
                                            v39 -= v39
                                        }
                                    case "":
                                        v40 := GenS4{F0: -1, F1: 100, F2: -295}
                                        _ = v40
                                    }
                                    _ = v39
                                }
                            }
                        }
                    case -448:
                        var v41 int64 = 9223372036854775807
                        _ = v41
                    }
                }
            }
            var v42 uint64 = 369
            _ = v42
        } else {
            v43 := GenS1{F0: int64(fi8), F1: "a"}
            _ = v43
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return -32768
}

func main() {
    v1 := []int{324, -367, 2147483647, 200}
    v2 := make(map[uint32]string)
    var v3 float32 = (-2.25 * 1.5)
    var v4 int = 626
    v5 := make(map[uint32]bool)
    var v6 float32 = (3.14159 / 1.0)
    if true {
        v1 = append(v1, v4)
        v6 += (1e10 / (((((float32(1e10) + (0.1 * v3)) + (1.5 * ((float32(float64((1.5 * 0.1))) + 1e10) / 3.14159))) * v6) + 1.5) - float32(0.1)))
    } else {
        switch "foo bar" {
        case "hello":
            v44 := GenS5{F0: ""}
            _ = v44
        case "foo bar":
            var v45 int16 = 1000
            _ = v45
        case "":
            if true {
                v3 -= v3
            } else {
                var v46 float32 = (-2.25 * v6)
                if false {
                    var v47 float64 = 3.14159
                    _ = v47
                }
                _ = v46
            }
        }
        _ = len(v2)
    }
    var v48 uint16 = 2
    v4 += (v4 >> ((736) & 63))
    switch 606 {
    case 1:
        switch 128 {
        case -2147483648:
            delete(v5, uint32(v48))
        }
    case 2147483647:
        var v49 int32 = -2147483648
        _ = v49
    }
    fmt.Println(v6)
    var v50 float64 = (3.14159 / 1.5)
    var v51 uint8 = 1
    if (0 != 4294967295) {
        for fi52 := 0; fi52 < 12; fi52++ {
            switch 128 {
            case -466:
                delete(v5, 406)
            case 2147483647:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -9223372036854775808:
                if false {
                    var v53 uint16 = 255
                    _ = v53
                } else {
                    v51 = v51
                    v4 = (-(v4))
                    v51 += 128
                }
            default:
                if false {
                    v51 = v51
                    v54 := GenS4{F0: -128, F1: -1, F2: 0}
                    if (v48 < v48) {
                        for fi55 := 0; fi55 < 5; fi55++ {
                            v51 = (v51 - (^(v51)))
                            v56 := v5[4294967295]
                            switch fi55 {
                            case 32768:
                                v50 = 0.1
                            }
                            _ = fi55
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 14; fi57++ {
                            if ((fi52 ^ 32768) > fi57) {
                                fmt.Println(v3, v6)
                                v58 := GenS1{F0: -932, F1: "hello"}
                                _ = v58
                            }
                            v59 := GenS6{F0: v48}
                            switch "a" {
                            case "":
                                if true {
                                    var v60 string = "a"
                                    var v61 uint32 = 524
                                    _ = v60
                                    _ = v61
                                }
                            case "\n\t":
                                if (0 == genFn0()) {
                                    if (!true) {
                                        var v62 int64 = 9223372036854775807
                                        var v63 uint32 = 279
                                        if (v51 >= (v51 * (v51 + v51))) {
                                            if false {
                                                delete(v5, 190)
                                                var v64 int8 = 127
                                                _ = v64
                                            } else {
                                                var v65 float64 = ((-2.25 + v50) * 0.0)
                                                _ = v65
                                            }
                                            if true {
                                                var v66 uint16 = (v48 >> ((255) & 63))
                                                v67 := GenS4{F0: -94, F1: 0, F2: 490}
                                                _ = v66
                                                _ = v67
                                            } else {
                                                var v68 bool = true
                                                _ = v68
                                            }
                                        }
                                        switch "hello" {
                                        case "a":
                                            v48 = v48
                                        }
                                        v4 = (fi57 | v4)
                                        switch -2147483648 {
                                        case 32768:
                                            _ = len(v1)
                                        default:
                                            if ((!false) && true) {
                                                v3 = (1.5 / (((v3 - (v3 - 1e10)) + (v6 * v3)) - 1e10))
                                            } else {
                                                var v69 uint16 = 255
                                                _ = v69
                                            }
                                        }
                                        if (false && false) {
                                            if (false && true) {
                                                fmt.Println(fi57, v4)
                                            }
                                        } else {
                                            v51 = (v51 * (v51 % ((v51) | 1)))
                                        }
                                        _ = v62
                                        _ = v63
                                    } else {
                                        v6 += v6
                                    }
                                    fmt.Println(v4, v48, v51)
                                }
                            default:
                                switch "a" {
                                case "a":
                                    var v70 int64 = -311
                                    _ = v70
                                case "foo bar":
                                    v71 := GenS3{F0: -422, F1: (-1 < 1)}
                                    _ = v71
                                case "":
                                    v48 += 710
                                default:
                                    v72 := GenS3{F0: 0, F1: true}
                                    _ = v72
                                }
                            }
                            _ = fi57
                            _ = v59
                        }
                    } else {
                        var v73 uint64 = 4294967295
                        _ = v73
                    }
                    _ = v54
                } else {
                    for fi74 := 0; fi74 < 17; fi74++ {
                        v6 = v3
                        v4 = -9223372036854775808
                        _ = fi74
                    }
                    for fi75 := 0; fi75 < 7; fi75++ {
                        v76 := v5[2147483647]
                        if v76 {
                            fmt.Println(v48)
                        } else {
                            v5[2] = (1.5 == v50)
                        }
                        v50 = ((-2.25 + (float64(v3) * v50)) / v50)
                        _ = fi75
                        _ = v76
                    }
                }
            }
            _ = fi52
        }
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                v77 := v2[2147483647]
                _ = v77
            }
        case "hello":
            for fi78 := 0; fi78 < 13; fi78++ {
                if (v4 != fi78) {
                    v79 := GenS2{F0: v48}
                    if false {
                        var v80 uint32 = 122
                        switch "a" {
                        case "a":
                            v81 := GenS1{F0: 9223372036854775807, F1: "a"}
                            _ = v81
                        case "foo bar":
                            v4 = fi78
                        case "":
                            fmt.Println(v3, v48, v51)
                        }
                        _ = v80
                    }
                    var v82 uint64 = 2
                    _ = v79
                    _ = v82
                } else {
                    v83 := GenS7{F0: 42}
                    _ = v83
                }
                if true {
                    v84 := GenS8{F0: int8(v51), F1: 0, F2: "\n\t"}
                    _ = v84
                }
                _ = fi78
            }
        }
    }
    v85 := GenS1{F0: -544, F1: "\n\t"}
    var v86 int8 = -28
    var v87 int32 = -780
    var v88 uint8 = (v51 * (v51 / ((v51) | 1)))
    if ((v87 / ((v87) | 1)) == v87) {
        v89 := GenS9{F0: (^(v87)), F1: 2147483647}
        _ = len(v5)
        _ = v89
    } else {
        v4 = (v4 ^ v4)
    }
    _ = len(v1)
    _ = len(v2)
    var v90 string = ("foo bar" + "\n\t")
    v91 := v2[2147483648]
    v51 = v51
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v90
    _ = v91
}
