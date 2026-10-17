package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 bool; F1 bool; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint32, p1 int8) int {
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{int(p0), (fi6 + (^(fi6))), (^(fi6))}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi6
        _ = v7
    }
    if (-3 >= 0) {
        p1 = p1
    } else {
        p1 = p1
        var v8 int = 128
        _ = v8
    }
    _ = p0
    _ = p1
    return 836
}
func genFn1() uint8 {
    v9 := GenS1{F0: -32768, F1: 4294967295}
    switch ((("hello" + "a") + "") + "a") {
    case "\n\t":
        switch 127 {
        case 42:
            var v10 uint32 = 1
            _ = v10
        case 0:
            v11 := GenS2{F0: 0}
            _ = v11
        case 128:
            var v12 uint32 = 0
            _ = v12
        }
    case "":
        var v13 uint64 = 4294967295
        _ = v13
    case "hello":
        v14 := []int{128, 9223372036854775807, 940}
        _ = v14
    default:
        if (!(-1 != 32767)) {
            var v15 float64 = ((((0.0 + float64(3.14159)) - 1e10) - (1.5 + 1.5)) * 1.5)
            var v16 float64 = (((v15 / v15) / 3.14159) * v15)
            if ("hello" != ((("hello" + "hello") + ("hello" + "\n\t")) + ("" + "foo bar"))) {
                v17 := []int{32768, 32767, -1, 1}
                v18 := GenS3{F0: 0, F1: 727, F2: 32767}
                if ((!false) && false) {
                    var v19 int8 = 1
                    v19 = (v19 / (((v19 - (v19 * 59))) | 1))
                    _ = v19
                } else {
                    v15 += v16
                    v20 := GenS1{F0: -1, F1: 18446744073709551615}
                    _ = v20
                }
                _ = v17
                _ = v18
            }
            _ = v15
            _ = v16
        } else {
            var v21 string = "hello"
            v21 = ((("hello" + ((v21 + (v21 + v21)) + v21)) + v21) + v21)
            _ = v21
        }
    }
    v22 := GenS4{F0: (true || (0 >= 0)), F1: (1.5 <= float32(0.0)), F2: 595}
    if true {
        if false {
            var v23 float32 = -2.25
            if false {
                var v24 string = "foo bar"
                v25 := []int{42, 9223372036854775807, 2147483647}
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: 1}
                v23 -= -2.25
                _ = v26
            }
            v27 := GenS5{F0: true}
            _ = v23
            _ = v27
        }
    }
    if true {
        var v28 int64 = 1
        for fi29 := 0; fi29 < 6; fi29++ {
            v28 = 9223372036854775807
            var v30 uint64 = 18446744073709551615
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 uint8 = 255
                var v33 int8 = 127
                var v34 int16 = -32768
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi29
            _ = v30
        }
        v28 = 42
        _ = v28
    }
    _ = v22
    _ = v9
    return 255
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint = 525
    v3 := make(map[int]int)
    var v4 uint16 = 42
    var v35 int16 = 42
    v36 := v3[-2147483648]
    var v37 bool = false
    fmt.Println(v36)
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 uint64 = 2
        v35 -= (v35 * (v35 | v35))
        _ = fi38
        _ = v39
    }
    if v37 {
        v2 -= 821
        v40 := v3[genFn0(249, 71)]
        _ = v40
    } else {
        if (!((v4 / ((v4) | 1)) != 65535)) {
            for fi41 := 0; fi41 < 14; fi41++ {
                for fi42 := 0; fi42 < 20; fi42++ {
                    var v43 int8 = 100
                    var v44 int64 = -1
                    if (genFn0(42, int8(fi42)) != (-(v36))) {
                        v43 += 0
                        if (!(!(3.14159 >= ((3.14159 * 3.14159) - -2.25)))) {
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    if v37 {
                                        fmt.Println(v2, v44, v37)
                                        if (!v37) {
                                            if (-1 > 164) {
                                                v36 = v36
                                                fmt.Println(v35, v43, fi41)
                                                fmt.Println(v36, fi42)
                                                var v45 uint8 = 42
                                                _ = v45
                                            }
                                            if v37 {
                                                v46 := GenS4{F0: ((2 != 440) && v37), F1: ((0 <= uint32(v4)) && (42 > -2147483648)), F2: v35}
                                                if ((v4 + (v4 % (((v4 % ((v4) | 1))) | 1))) <= uint16(v2)) {
                                                    v44 -= v44
                                                    if true {
                                                        var v47 string = "\n\t"
                                                        v48 := GenS3{F0: 88, F1: 855, F2: (v36 << ((v2) & 63))}
                                                        v43 -= (v43 >> ((v2) & 63))
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        v35 = v35
                                                    }
                                                    var v49 int64 = int64(v2)
                                                    _ = v49
                                                } else {
                                                    var v50 uint = v2
                                                    _ = v50
                                                }
                                                _ = v46
                                            }
                                        }
                                    } else {
                                        v43 += -47
                                        switch "hello" {
                                        case "":
                                            var v51 int16 = (v35 / ((393) | 1))
                                            _ = v51
                                        case "a":
                                            var v52 uint16 = (v4 | 432)
                                            _ = v52
                                        case "hello":
                                            v53 := GenS5{F0: ((!v37) && false)}
                                            _ = v53
                                        }
                                        if false {
                                            fmt.Println(v4, v43)
                                            v54 := GenS5{F0: (!(!(0.0 > (1e10 + ((-2.25 + 0.1) + (-2.25 / 1.0))))))}
                                            _ = len(v3)
                                            v4 -= (v4 & v4)
                                            var v55 int = (^(fi41))
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 uint16 = (v4 & uint16(v2))
                                            v2 = 4294967295
                                            _ = v56
                                        }
                                    }
                                    if (v43 >= (-(v43))) {
                                        var v57 uint32 = 84
                                        var v58 int32 = 42
                                        if (!(v37 || (!(!true)))) {
                                            switch (fi42 >> ((v2) & 63)) {
                                            case -2147483648:
                                                var v59 uint8 = genFn1()
                                                v2 = 255
                                                var v60 float64 = 1e10
                                                v57 -= (v57 ^ v57)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v37 = false
                                                _ = v59
                                                _ = v60
                                            case 1:
                                                var v61 uint32 = (v57 / (((v57 + (v57 % (((v57 * (^(v57)))) | 1)))) | 1))
                                                _ = v61
                                            case 214:
                                                v43 = 45
                                            default:
                                                switch fi42 {
                                                case -617:
                                                    switch "" {
                                                    case "foo bar":
                                                        v62 := GenS4{F0: ((v57 / ((v57) | 1)) >= (v57 | v57)), F1: (1 == v43), F2: (^(v35))}
                                                        _ = v62
                                                    case "a":
                                                        var v63 int8 = (v43 << ((887) & 63))
                                                        v57 = 2147483648
                                                        _ = v63
                                                    case "hello":
                                                        var v64 int64 = v44
                                                        _ = v64
                                                    default:
                                                        v65 := GenS6{F0: "a"}
                                                        _ = v65
                                                    }
                                                case 9223372036854775807:
                                                    v66 := GenS3{F0: 255, F1: 42, F2: (fi42 >> ((4294967295) & 63))}
                                                    _ = v66
                                                default:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v44 = v44
                                                    case "a":
                                                        var v67 int16 = (v35 - (-(v35)))
                                                        _ = v67
                                                    case "":
                                                        if v37 {
                                                            var v68 uint16 = (v4 >> (((v2 >> ((18446744073709551615) & 63))) & 63))
                                                            _ = v68
                                                        }
                                                    }
                                                    v69 := GenS1{F0: 42, F1: uint(v57)}
                                                    v2 = v2
                                                    _ = v69
                                                }
                                            }
                                        } else {
                                            var v70 uint32 = v57
                                            var v71 float32 = (float32((float64(0.0) / ((float64((0.1 + 0.1)) * 0.1) * 1e10))) - (0.1 * (3.14159 * ((0.0 * (float32(((float64(float32(1e10)) / 3.14159) + float64((1.5 * (0.1 / (1.5 + 0.1)))))) / 1e10)) / -2.25))))
                                            switch fi42 {
                                            case 9223372036854775807:
                                                v1 = append(v1, (v57 << ((v2) & 63)))
                                            case 25:
                                                v71 += v71
                                            default:
                                                v72 := v3[v36]
                                                _ = v72
                                            }
                                            v44 -= v44
                                            fmt.Println(v43, v35, v71)
                                            v73 := GenS7{F0: (v43 ^ (v43 | v43))}
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v74 uint = v2
                                                _ = v74
                                            case "":
                                                var v75 int64 = 646
                                                _ = v75
                                            case "a":
                                                v76 := GenS2{F0: 456}
                                                _ = v76
                                            default:
                                                v1 = append(v1, (v57 | (^(v70))))
                                            }
                                            _ = v70
                                            _ = v71
                                            _ = v73
                                        }
                                        _ = v57
                                        _ = v58
                                    } else {
                                        v2 = 0
                                        var v77 uint64 = 2
                                        v1 = append(v1, 2147483648)
                                        _ = v77
                                    }
                                    var v78 int32 = 0
                                    _ = v78
                                } else {
                                    var v79 int32 = 42
                                    v4 = 1
                                    v79 = (v79 >> ((v2) & 63))
                                    _ = v79
                                }
                            case "":
                                if v37 {
                                    v36 = 127
                                } else {
                                    v80 := GenS5{F0: false}
                                    _ = v80
                                }
                            default:
                                fmt.Println(v37, v43)
                            }
                            var v81 string = ""
                            _ = v81
                        }
                        v43 += (v43 / ((int8(v4)) | 1))
                    } else {
                        var v82 uint32 = 4294967295
                        if (v37 && true) {
                            var v83 uint16 = (v4 << (((v2 ^ (v2 << ((v2) & 63)))) & 63))
                            var v84 int16 = (v35 << (((v2 % (((v2 / (((^(v2))) | 1))) | 1))) & 63))
                            fmt.Println(v43, fi41, v4)
                            _ = v83
                            _ = v84
                        }
                        _ = v82
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                if (!true) {
                    if v37 {
                        v4 = v4
                        if (32767 >= 32767) {
                            var v85 int = v36
                            _ = v85
                        } else {
                            fmt.Println(v36, fi41)
                            for fi86 := 0; fi86 < 13; fi86++ {
                                v87 := GenS7{F0: 1}
                                v36 = 9223372036854775807
                                v88 := GenS2{F0: 1}
                                var v89 uint32 = 2147483648
                                v37 = (v37 && v37)
                                _ = fi86
                                _ = v87
                                _ = v88
                                _ = v89
                            }
                        }
                        _ = len(v1)
                    } else {
                        var v90 string = "a"
                        var v91 uint = v2
                        v1 = append(v1, 216)
                        _ = v90
                        _ = v91
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v92 int64 = -655
                _ = fi41
                _ = v92
            }
            var v93 int = v36
            var v94 uint32 = 1
            _ = v93
            _ = v94
        }
        v4 -= v4
    }
    if false {
        if v37 {
            if true {
                var v95 float32 = -2.25
                var v96 float32 = 0.1
                _ = v95
                _ = v96
            }
            fmt.Println(v37, v37, v35)
            switch "a" {
            case "\n\t":
                var v97 string = "\n\t"
                _ = v97
            case "":
                v98 := GenS7{F0: 0}
                _ = v98
            }
        }
        for fi99 := 0; fi99 < 4; fi99++ {
            delete(v3, (v36 ^ fi99))
            delete(v3, v36)
            for fi100 := 0; fi100 < 14; fi100++ {
                v37 = v37
                _ = fi100
            }
            _ = fi99
        }
    }
    fmt.Println(v37, v37, v37)
    v2 = v2
    var v101 int = v36
    v102 := GenS5{F0: ((0.1 >= (0.0 / -2.25)) && false)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
