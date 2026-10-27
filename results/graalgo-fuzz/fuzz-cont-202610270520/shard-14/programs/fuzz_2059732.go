package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0() uint {
    if (279 < 551) {
        var v4 int = 32768
        if (false || true) {
            if false {
                v5 := GenS1{F0: 175, F1: 196}
                v4 = v4
                _ = v5
            } else {
                v4 += v4
                if (3.14159 > 1e10) {
                    if (3.14159 == (1.5 + 0.0)) {
                        v4 = v4
                        if false {
                            fmt.Println(v4)
                            var v6 int32 = 104
                            if (!true) {
                                var v7 uint = 685
                                v8 := []int{}
                                v9 := GenS1{F0: 95, F1: 255}
                                _ = v7
                                _ = v8
                                _ = v9
                            } else {
                                v10 := []int{v4}
                                for fi11 := 0; fi11 < 17; fi11++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v12 int64 = -304
                                        _ = v12
                                    case "a":
                                        fmt.Println(v6, v4)
                                    case "\n\t":
                                        if false {
                                            var v13 string = "\n\t"
                                            if ((!false) || true) {
                                                v4 += 32768
                                            } else {
                                                if len(v10) > 0 {
                                                    v10[0] = -202
                                                }
                                                v13 = v13
                                            }
                                            switch fi11 {
                                            case 127:
                                                v14 := GenS2{F0: 1}
                                                _ = v14
                                            case -121:
                                                v13 = v13
                                            case 32767:
                                                v15 := GenS2{F0: 32767}
                                                _ = v15
                                            }
                                            _ = v13
                                        } else {
                                            fmt.Println(v6)
                                        }
                                    default:
                                        var v16 int16 = 0
                                        _ = v16
                                    }
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        fmt.Println(fi11)
                                    }
                                    v17 := GenS2{F0: 1}
                                    _ = fi11
                                    _ = v17
                                }
                                switch (("" + (("foo bar" + ("" + "a")) + "a")) + "\n\t") {
                                case "":
                                    if len(v10) > 0 {
                                        v10[0] = v4
                                    }
                                default:
                                    v18 := GenS3{F0: 2147483647}
                                    _ = v18
                                }
                                _ = v10
                            }
                            _ = v6
                        }
                        var v19 int16 = 32767
                        _ = v19
                    } else {
                        var v20 string = (("" + ("a" + "\n\t")) + "")
                        if true {
                            v21 := []int{}
                            var v22 int16 = 1
                            v23 := GenS2{F0: 2}
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            v4 += v4
                            v24 := GenS4{F0: true, F1: 1, F2: (v20 + v20)}
                            _ = v24
                        }
                        v20 = ((v20 + ("\n\t" + (v20 + "a"))) + v20)
                        _ = v20
                    }
                    if (0 <= 371) {
                        v25 := []int{v4, v4, (v4 / (((v4 | -411)) | 1)), v4}
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v4 += 127
                            v27 := GenS1{F0: 128, F1: 2}
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    }
                }
                v4 += 32767
            }
        } else {
            var v28 int16 = 1
            switch "hello" {
            case "":
                switch 32768 {
                case -826:
                    var v29 int32 = 137
                    _ = v29
                case 9223372036854775807:
                    var v30 int8 = -24
                    _ = v30
                default:
                    if (696 != 1) {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            var v32 uint64 = 4294967295
                            switch ("" + ("\n\t" + "a")) {
                            case "a":
                                var v33 int = fi31
                                _ = v33
                            }
                            _ = fi31
                            _ = v32
                        }
                    }
                }
            case "foo bar":
                fmt.Println(v4, v28)
            }
            _ = v28
        }
        _ = v4
    } else {
        var v34 int32 = 42
        v35 := GenS4{F0: false, F1: 174, F2: "hello"}
        _ = v34
        _ = v35
    }
    v36 := []int{-314, 42, -2147483648}
    v37 := GenS2{F0: 1}
    _ = v36
    _ = v37
    return 42
}

func main() {
    var v1 bool = (((42 == 42) || false) && (false || false))
    var v2 float32 = 3.14159
    v3 := make(map[uint32]int)
    v2 += 0.1
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 9; fi39++ {
            fmt.Println(v2)
            if ((42 >= 119) && v1) {
                var v40 uint64 = 4294967295
                switch "" {
                case "foo bar":
                    v1 = v1
                case "\n\t":
                    var v41 float32 = ((v2 + v2) / v2)
                    _ = v41
                }
                var v42 int16 = -32768
                _ = v40
                _ = v42
            } else {
                var v43 float32 = v2
                v1 = v1
                _ = v43
            }
            _ = fi39
        }
        v44 := v3[737]
        _ = fi38
        _ = v44
    }
    v45 := GenS1{F0: 128, F1: 127}
    if v1 {
        var v46 int = -2147483648
        if (-1 >= 1) {
            v46 = -9223372036854775808
            var v47 uint8 = 127
            var v48 uint32 = 99
            _ = v47
            _ = v48
        } else {
            v49 := GenS2{F0: 685}
            if false {
                var v50 uint = 18446744073709551615
                _ = v50
            } else {
                if (v1 && (v46 < v46)) {
                    if (2 > 127) {
                        var v51 int = 4294967295
                        v52 := GenS2{F0: 65535}
                        _ = v51
                        _ = v52
                    } else {
                        fmt.Println(v2)
                        v53 := GenS5{F0: genFn0()}
                        _ = v53
                    }
                }
            }
            _ = v49
        }
        var v54 int16 = 42
        _ = v46
        _ = v54
    }
    v2 = v2
    v1 = true
    switch 727 {
    case 0:
        for fi55 := 0; fi55 < 15; fi55++ {
            var v56 uint32 = 1
            _ = fi55
            _ = v56
        }
    }
    v57 := []int{-69}
    var v58 float32 = 0.0
    v3[454] = -755
    var v59 float64 = (((0.1 * 1e10) + -2.25) / 3.14159)
    fmt.Println(v2, v2)
    if (!v1) {
        if (!(((genFn0() < 42) && false) && (false && (v1 && v1)))) {
            v1 = v1
        } else {
            for fi60 := 0; fi60 < 3; fi60++ {
                var v61 int32 = -1
                switch "foo bar" {
                case "hello":
                    if (977 != 42) {
                        var v62 float64 = v59
                        _ = v62
                    }
                }
                _ = fi60
                _ = v61
            }
            if v1 {
                _ = len(v57)
                if (((!v1) || v1) || false) {
                    v2 = float32(float64((v2 - -2.25)))
                } else {
                    switch "foo bar" {
                    case "a":
                        switch 32767 {
                        case 8:
                            if len(v57) > 0 {
                                v57[0] = 2147483647
                            }
                        case 1:
                            v59 += (0.1 - (v59 + (float64((0.0 / (float32(float64((v2 / 1.5))) / 1e10))) * v59)))
                        case 128:
                            if v1 {
                                v58 = v58
                            }
                        }
                    case "\n\t":
                        fmt.Println(v58, v58, v58)
                    default:
                        fmt.Println(v58, v59)
                    }
                }
                if v1 {
                    v63 := GenS1{F0: 1, F1: 255}
                    if false {
                        var v64 uint16 = 0
                        _ = len(v3)
                        _ = v64
                    } else {
                        if v1 {
                            for fi65 := 0; fi65 < 20; fi65++ {
                                var v66 uint = genFn0()
                                _ = fi65
                                _ = v66
                            }
                        } else {
                            switch -253 {
                            case -9223372036854775808:
                                fmt.Println(v2, v58, v58)
                            case -33:
                                switch "\n\t" {
                                case "a":
                                    v67 := GenS5{F0: 4294967295}
                                    _ = v67
                                case "hello":
                                    if (v2 > float32(-2.25)) {
                                        v59 += (v59 - v59)
                                    } else {
                                        fmt.Println(v1, v58)
                                        var v68 uint8 = 1
                                        for fi69 := 0; fi69 < 16; fi69++ {
                                            for fi70 := 0; fi70 < 11; fi70++ {
                                                if ((v68 == v68) || (!v1)) {
                                                    v58 = ((v2 / v58) + (v2 * (v2 * 1.5)))
                                                    if len(v57) > 0 {
                                                        v57[0] = (-(fi70))
                                                    }
                                                    v1 = v1
                                                }
                                                var v71 float64 = v59
                                                var v72 int16 = -1
                                                _ = fi70
                                                _ = v71
                                                _ = v72
                                            }
                                            v59 -= 0.1
                                            _ = fi69
                                        }
                                        _ = v68
                                    }
                                case "":
                                    var v73 float32 = float32(v59)
                                    _ = v73
                                }
                            case 2147483647:
                                var v74 uint64 = 755
                                _ = v74
                            default:
                                var v75 uint8 = 127
                                _ = v75
                            }
                            fmt.Println(v59)
                            for fi76 := 0; fi76 < 10; fi76++ {
                                if (((!false) && (v1 || (!true))) && true) {
                                    var v77 uint16 = 220
                                    var v78 uint8 = 0
                                    _ = v77
                                    _ = v78
                                }
                                v57 = append(v57, (fi76 - (fi76 % ((fi76) | 1))))
                                _ = fi76
                            }
                        }
                    }
                    v79 := GenS1{F0: 221, F1: 162}
                    _ = v63
                    _ = v79
                } else {
                    switch "" {
                    case "":
                        for fi80 := 0; fi80 < 12; fi80++ {
                            fmt.Println(fi80, v59)
                            for fi81 := 0; fi81 < 14; fi81++ {
                                if (v59 > 0.0) {
                                    if (true || (42 == 1)) {
                                        v59 = -2.25
                                        if len(v57) > 0 {
                                            v57[0] = fi80
                                        }
                                    }
                                    var v82 uint8 = 2
                                    _ = v82
                                } else {
                                    var v83 string = ""
                                    v58 = ((v2 + float32(v59)) * v58)
                                    _ = v83
                                }
                                _ = fi81
                            }
                            switch ("hello" + "\n\t") {
                            case "":
                                for fi84 := 0; fi84 < 1; fi84++ {
                                    v2 -= (v58 / 3.14159)
                                    _ = fi84
                                }
                            case "a":
                                v85 := GenS4{F0: false, F1: 4294967295, F2: ("a" + (("hello" + "\n\t") + "\n\t"))}
                                _ = v85
                            }
                            _ = fi80
                        }
                    case "hello":
                        var v86 float64 = v59
                        _ = v86
                    }
                    v59 = float64(1.5)
                }
            }
        }
    } else {
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v87 int32 = 42
            _ = v87
        case -9223372036854775808:
            for fi88 := 0; fi88 < 11; fi88++ {
                v89 := GenS6{F0: 14, F1: 2147483647}
                var v90 int16 = int16(fi88)
                _ = fi88
                _ = v89
                _ = v90
            }
        }
        var v91 int32 = -1
        _ = v91
    }
    v92 := GenS7{F0: 0, F1: -96, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v57
    _ = v58
    _ = v59
    _ = v92
}
