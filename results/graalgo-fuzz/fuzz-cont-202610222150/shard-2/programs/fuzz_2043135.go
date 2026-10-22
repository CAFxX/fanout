package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 int16; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS6 struct { F0 string; F1 int16; F2 int8 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int16, p1 int64, p2 uint32) uint16 {
    switch "a" {
    case "hello":
        v4 := []int{32768, -346}
        _ = v4
    case "a":
        var v5 uint = 2
        _ = v5
    default:
        v6 := GenS1{F0: 2}
        _ = v6
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 bool, p1 uint64, p2 int16) int8 {
    var v7 float64 = float64((0.1 * -2.25))
    var v8 string = "foo bar"
    var v9 uint = 738
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 108
}

func main() {
    var v1 int8 = -111
    var v2 int16 = 904
    v3 := make(map[int]bool)
    v10 := GenS2{F0: 666, F1: 42}
    if true {
        if ((!false) || (678 < 199)) {
            var v11 uint8 = 128
            v12 := GenS2{F0: 2, F1: 4294967295}
            _ = v11
            _ = v12
        }
        v13 := []int{-1, 32767, -649, 32768}
        _ = v13
    } else {
        fmt.Println(v1, v2)
    }
    _ = len(v3)
    if false {
        v14 := GenS1{F0: 255}
        _ = v14
    } else {
        v15 := GenS2{F0: 2, F1: 937}
        var v16 int32 = -2147483648
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := []int{fi18, 2147483647, fi17, fi17}
            v2 += v2
            switch fi18 {
            case 127:
                switch fi18 {
                case 1:
                    var v20 int8 = v1
                    _ = v20
                case 0:
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case -386:
                    delete(v3, fi18)
                }
            default:
                if ((true || true) || (false && (false || (false && (true && (!(!(58 != 111)))))))) {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    var v21 uint = 4294967295
                    var v22 float32 = -2.25
                    var v23 uint64 = 124
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v2)
                    fmt.Println(v1, v2)
                }
            }
            v2 += v2
            if (!(!true)) {
                fmt.Println(fi17, v2, fi18)
                v24 := GenS1{F0: 83}
                fmt.Println(fi17, v1, v2)
                _ = v24
            }
            if true {
                v25 := GenS1{F0: 127}
                _ = len(v3)
                _ = v25
            }
            var v26 uint = 4294967295
            var v27 bool = (127 > 137)
            if len(v19) > 0 {
                v19[0] = fi18
            }
            _ = fi18
            _ = v19
            _ = v26
            _ = v27
        }
        _ = fi17
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS1{F0: 2}
        v2 = 32767
        v30 := GenS3{F0: fi28, F1: (v2 + v2), F2: 65535}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 bool = false
    v32 := GenS1{F0: 168}
    v33 := GenS3{F0: 42, F1: v2, F2: 42}
    fmt.Println(v31, v1, v2)
    var v34 int16 = 32767
    for fi35 := 0; fi35 < 16; fi35++ {
        switch (^(fi35)) {
        case 128:
            delete(v3, fi35)
        }
        v36 := GenS2{F0: 18446744073709551615, F1: 2}
        var v37 int = 2147483647
        _ = fi35
        _ = v36
        _ = v37
    }
    if v31 {
        for fi38 := 0; fi38 < 18; fi38++ {
            v2 = v2
            v39 := []int{243, fi38, (fi38 / ((fi38) | 1))}
            _ = fi38
            _ = v39
        }
        if (169 < 0) {
            for fi40 := 0; fi40 < 6; fi40++ {
                if true {
                    var v41 uint8 = 127
                    if v31 {
                        switch fi40 {
                        case 215:
                            v42 := GenS1{F0: (v41 + (v41 & (v41 * 42)))}
                            _ = v42
                        case -2147483648:
                            fmt.Println(v1, v2)
                        case 431:
                            v43 := GenS4{F0: 109}
                            _ = v43
                        default:
                            switch fi40 {
                            case 127:
                                v44 := GenS2{F0: 1, F1: 0}
                                _ = v44
                            default:
                                var v45 string = ("foo bar" + "")
                                _ = v45
                            }
                        }
                    }
                    _ = v41
                } else {
                    var v46 int64 = 9223372036854775807
                    v47 := GenS2{F0: 1, F1: 0}
                    _ = v46
                    _ = v47
                }
                if v31 {
                    fmt.Println(v1, v34)
                    switch (fi40 / (((fi40 - (fi40 * 666))) | 1)) {
                    case -438:
                        var v48 uint32 = 42
                        _ = v48
                    case 0:
                        v49 := v3[4294967295]
                        _ = v49
                    case 9223372036854775807:
                        if (-613 > 0) {
                            var v50 float64 = ((0.0 / (float64(0.0) + (1e10 + (1e10 + 3.14159)))) + ((1e10 + ((0.0 / 1e10) - 3.14159)) - 0.1))
                            v51 := GenS5{F0: 255, F1: 0, F2: 0}
                            _ = v50
                            _ = v51
                        } else {
                            var v52 float32 = 3.14159
                            v53 := v3[fi40]
                            var v54 uint = 255
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    default:
                        v1 = genFn1((133 != 54), 140, 1000)
                    }
                    if (127 < v1) {
                        fmt.Println(v31, fi40)
                    } else {
                        if true {
                            switch "a" {
                            case "":
                                v1 += (v1 & v1)
                            default:
                                delete(v3, 378)
                            }
                            for fi55 := 0; fi55 < 16; fi55++ {
                                v31 = true
                                if false {
                                    if (!(127 != v1)) {
                                        switch "a" {
                                        case "hello":
                                            v56 := []int{(fi55 >> ((911) & 63)), (fi55 & 32767), 127}
                                            _ = v56
                                        case "":
                                            delete(v3, (-(fi55)))
                                        default:
                                            v1 = genFn1((((3.14159 - 3.14159) * (((3.14159 + float64(1.5)) + float64(0.0)) / ((0.0 * 1e10) + 0.1))) < 0.1), 2, (-(v34)))
                                        }
                                        v1 = genFn1((0 == 1), 920, 0)
                                        if (42 != (v34 / ((0) | 1))) {
                                            v3[(fi40 ^ -1)] = (42 != 2)
                                        } else {
                                            v57 := GenS4{F0: -514}
                                            v34 += v34
                                            if (42 != 4294967295) {
                                                switch ("a" + ("\n\t" + ("foo bar" + ((("\n\t" + (((("foo bar" + "hello") + "\n\t") + "a") + "foo bar")) + (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))) + (("foo bar" + "a") + ((("hello" + "\n\t") + "hello") + "a")))))) {
                                                case "hello":
                                                    fmt.Println(v34)
                                                case "\n\t":
                                                    fmt.Println(v31, v31, v31)
                                                default:
                                                    switch "hello" {
                                                    case "hello":
                                                        v58 := []int{}
                                                        _ = v58
                                                    case "":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            fmt.Println(v1, fi40, v2)
                                                        case "":
                                                            v59 := GenS3{F0: fi55, F1: v34, F2: 415}
                                                            _ = v59
                                                        case "a":
                                                            if ((!v31) && (!(820 <= 298))) {
                                                                var v60 int8 = 100
                                                                _ = v60
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                v2 = 1
                                                if v31 {
                                                    v3[(fi55 << ((255) & 63))] = ((v31 && v31) && (-2147483648 >= 2147483647))
                                                    var v61 uint64 = 34
                                                    _ = v61
                                                } else {
                                                    if v31 {
                                                        v34 -= 42
                                                        v1 = int8(v2)
                                                    } else {
                                                        var v62 int16 = (^(v34))
                                                        var v63 float64 = 0.0
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                v1 += v1
                                            }
                                            v31 = v31
                                            fmt.Println(fi55, v1)
                                            _ = v57
                                        }
                                        var v64 uint8 = 255
                                        var v65 int8 = genFn1(v31, 1, v2)
                                        v66 := GenS6{F0: "hello", F1: 1, F2: int8(v2)}
                                        v65 = (v1 >> ((18446744073709551615) & 63))
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    } else {
                                        v67 := GenS5{F0: 1, F1: 0, F2: -2147483648}
                                        v68 := []int{}
                                        v69 := GenS6{F0: ("\n\t" + "a"), F1: (v2 % (((v2 | 437)) | 1)), F2: v1}
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    v70 := GenS7{F0: 18446744073709551615}
                                    _ = v70
                                }
                                var v71 int16 = v34
                                _ = fi55
                                _ = v71
                            }
                            v31 = true
                        }
                        v72 := GenS6{F0: ((((("hello" + ("\n\t" + ((("a" + "\n\t") + "hello") + ""))) + "a") + "foo bar") + "hello") + "a"), F1: -32768, F2: (v1 % (((v1 + (v1 | 100))) | 1))}
                        _ = v72
                    }
                }
                _ = fi40
            }
        } else {
            v73 := GenS3{F0: -925, F1: (v2 & 609), F2: 0}
            _ = v73
        }
        if ("\n\t" < "\n\t") {
            v31 = (1.5 != 0.1)
        } else {
            v74 := GenS7{F0: 942}
            _ = v74
        }
    }
    v3[32767] = true
    if (0 >= 253) {
        delete(v3, 0)
    } else {
        var v75 uint32 = 2
        v76 := []int{9223372036854775807}
        var v77 int32 = 171
        _ = v75
        _ = v76
        _ = v77
    }
    if v31 {
        fmt.Println(v34)
        v78 := GenS3{F0: 396, F1: v34, F2: 42}
        var v79 string = "hello"
        _ = v78
        _ = v79
    } else {
        v80 := GenS1{F0: 0}
        switch 299 {
        case 4294967295:
            v81 := GenS6{F0: "", F1: v34, F2: (-(v1))}
            _ = v81
        default:
            v82 := v3[-649]
            _ = v82
        }
        _ = v80
    }
    v83 := GenS7{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v83
}
