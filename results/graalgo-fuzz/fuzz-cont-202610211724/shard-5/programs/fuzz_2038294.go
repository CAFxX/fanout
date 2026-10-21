package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int; F2 bool }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 int; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int8; F1 bool }
func genFn0(p0 int8) string {
    var v6 int32 = 42
    v6 -= v6
    _ = p0
    _ = v6
    return "hello"
}
func genFn1() uint {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float64 = float64(float32(1e10))
        v9 := GenS1{F0: 65535, F1: fi7}
        v8 = (float64((0.1 * (0.0 - ((-2.25 + float32((1e10 - -2.25))) / 3.14159)))) / 1.5)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-567, -2147483648, 0, 2147483647}
    var v11 float64 = 1e10
    v10 = append(v10, 2147483647)
    v11 += float64(float32((v11 - 0.0)))
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint = 437
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint32 = 65535
    var v5 int64 = 874
    switch "\n\t" {
    case "a":
        var v12 uint64 = 42
        _ = v12
    case "":
        v5 = v5
    default:
        v5 -= 0
    }
    switch -698 {
    case -9223372036854775808:
        v13 := GenS1{F0: 255, F1: -603}
        _ = v13
    case 0:
        v1 = (^(v1))
    case 32767:
        v14 := GenS1{F0: 255, F1: 128}
        _ = v14
    }
    v15 := GenS1{F0: 235, F1: -378}
    if true {
        var v16 uint64 = 4294967295
        _ = v16
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            v4 = (v4 ^ v4)
            _ = fi17
        }
    }
    fmt.Println(v1, v1, v1)
    v18 := GenS1{F0: uint16(v1), F1: 824}
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(v1)
        v20 := GenS1{F0: 931, F1: (^(fi19))}
        _ = fi19
        _ = v20
    }
    if ((!false) || true) {
        for fi21 := 0; fi21 < 11; fi21++ {
            v1 -= (v1 * (v1 ^ genFn1()))
            fmt.Println(v5, v4)
            if true {
                var v22 int8 = -128
                fmt.Println(v1, v5)
                v23 := GenS1{F0: 65535, F1: (fi21 << ((v1) & 63))}
                _ = v22
                _ = v23
            } else {
                v1 = 65535
                for fi24 := 0; fi24 < 13; fi24++ {
                    switch "a" {
                    case "a":
                        var v25 uint8 = 42
                        _ = v25
                    case "hello":
                        v26 := []int{(fi24 >> (((^(v1))) & 63)), fi24, fi24, 32767}
                        _ = v26
                    case "":
                        fmt.Println(v4)
                    }
                    _ = fi24
                }
                fmt.Println(v4, v5)
            }
            _ = fi21
        }
        if true {
            var v27 int = -310
            fmt.Println(v1, v5)
            if false {
                v28 := GenS1{F0: 2, F1: v27}
                v5 = v5
                _ = v28
            }
            _ = v27
        }
    } else {
        var v29 bool = false
        v5 += (v5 | (v5 >> ((489) & 63)))
        v4 += 666
        _ = v29
    }
    if true {
        delete(v2, ("foo bar" + "hello"))
        if true {
            v30 := GenS1{F0: 2, F1: 914}
            _ = v30
        } else {
            if (42 <= 4294967295) {
                if true {
                    v31 := GenS2{F0: -706}
                    var v32 uint16 = 0
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        var v34 int64 = v5
                        _ = fi33
                        _ = v34
                    }
                    var v35 int16 = -32768
                    v35 = v35
                    _ = v35
                }
                v36 := []int{32767, 0, 4294967295, -9223372036854775808}
                _ = v36
            } else {
                var v37 int = -798
                var v38 uint32 = (v4 * 65535)
                if (false && (int16(v1) <= 0)) {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        var v40 int16 = 196
                        switch "a" {
                        case "foo bar":
                            for fi41 := 0; fi41 < 5; fi41++ {
                                switch (fi41 % ((-24) | 1)) {
                                case 693:
                                    v2["hello"] = "foo bar"
                                case 32768:
                                    v42 := GenS3{F0: v38, F1: 42, F2: -1}
                                    var v43 int32 = -2147483648
                                    if true {
                                        v44 := GenS4{F0: 42, F1: 42, F2: (578 > 42)}
                                        _ = v44
                                    }
                                    _ = v42
                                    _ = v43
                                case 127:
                                    v45 := []int{fi41, (^(fi41)), int(v1), (fi39 % (((fi41 / ((fi41) | 1))) | 1))}
                                    _ = v45
                                default:
                                    var v46 bool = (((!true) && false) || ((v1 % (((v1 ^ 794)) | 1)) != (v1 - 42)))
                                    _ = v46
                                }
                                _ = fi41
                            }
                        case "":
                            fmt.Println(v1, v4, v4)
                        case "hello":
                            v47 := GenS5{F0: (v37 ^ v37), F1: -1}
                            _ = v47
                        default:
                            var v48 uint = v1
                            _ = v48
                        }
                        var v49 bool = ((!false) || true)
                        _ = fi39
                        _ = v40
                        _ = v49
                    }
                    v50 := []int{(v37 * v37), (v37 + -1)}
                    if (1e10 > -2.25) {
                        v51 := GenS6{F0: 2147483647}
                        var v52 int32 = 1
                        _ = v51
                        _ = v52
                    }
                    _ = v50
                }
                _ = v37
                _ = v38
            }
            v4 = (v4 + (^(v4)))
            if true {
                v5 = 9223372036854775807
            } else {
                v53 := []int{-1, 418, -572, -565}
                _ = v53
            }
        }
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 int32 = -2147483648
            _ = fi54
            _ = v55
        }
    } else {
        v56 := GenS7{F0: "hello", F1: 4294967295, F2: (2147483647 < -160)}
        var v57 int64 = (v5 >> (((v1 & 18446744073709551615)) & 63))
        var v58 int64 = (v5 % (((v5 | 0)) | 1))
        _ = v56
        _ = v57
        _ = v58
    }
    v59 := []int{1, 679, 133}
    fmt.Println(v5, v1, v5)
    switch 32768 {
    case 42:
        fmt.Println(v1, v5, v4)
    case 652:
        if (2147483648 <= (v4 + v4)) {
            if len(v59) > 0 {
                v59[0] = -751
            }
            v60 := GenS6{F0: 32767}
            _ = v60
        } else {
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
        }
    default:
        for fi61 := 0; fi61 < 6; fi61++ {
            delete(v2, ("foo bar" + ("hello" + "hello")))
            for fi62 := 0; fi62 < 1; fi62++ {
                switch "a" {
                case "foo bar":
                    var v63 float32 = (1e10 * (1.5 + 1.5))
                    fmt.Println(v63, v4)
                    if (3.14159 >= ((1e10 / -2.25) - (((1.5 + (float64(float32((-2.25 * ((0.0 / float64(1e10)) * 0.0)))) * 3.14159)) / 3.14159) * 0.0))) {
                        fmt.Println(v4, fi62)
                        v64 := v2["a"]
                        _ = v64
                    } else {
                        v65 := GenS1{F0: 0, F1: fi61}
                        v63 += 0.1
                        v5 += (v5 / (((v5 / (((v5 << ((v1) & 63))) | 1))) | 1))
                        _ = v65
                    }
                    _ = v63
                case "hello":
                    if (65535 >= 2) {
                        switch "hello" {
                        case "hello":
                            if ((3.14159 > -2.25) || (!true)) {
                                var v66 uint64 = 42
                                v67 := GenS1{F0: 42, F1: -9223372036854775808}
                                v4 = v4
                                _ = v66
                                _ = v67
                            }
                        case "\n\t":
                            switch 4294967295 {
                            case -2147483648:
                                if true {
                                    var v68 bool = (int8(fi61) < -113)
                                    v1 = (v1 / ((v1) | 1))
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                    _ = v68
                                } else {
                                    var v69 int16 = 32767
                                    v4 += v4
                                    var v70 string = ("a" + "foo bar")
                                    _ = v69
                                    _ = v70
                                }
                            default:
                                var v71 int = -418
                                _ = v71
                            }
                            switch "" {
                            case "a":
                                v72 := GenS8{F0: v5, F1: false}
                                _ = v72
                            case "foo bar":
                                var v73 string = "hello"
                                _ = v73
                            }
                        default:
                            v74 := GenS1{F0: 1, F1: fi61}
                            _ = v74
                        }
                        var v75 int8 = 42
                        var v76 bool = true
                        _ = v75
                        _ = v76
                    } else {
                        if len(v59) > 0 {
                            v59[0] = fi62
                        }
                    }
                case "a":
                    v3[(fi61 * (^(fi62)))] = fi61
                default:
                    v5 = (v5 / (((v5 >> (((v1 % ((255) | 1))) & 63))) | 1))
                }
                v1 -= 18446744073709551615
                v77 := v2[("a" + ("\n\t" + "a"))]
                if len(v59) > 0 {
                    v59[0] = fi62
                }
                var v78 int = fi62
                _ = fi62
                _ = v77
                _ = v78
            }
            var v79 uint64 = 42
            _ = fi61
            _ = v79
        }
    }
    fmt.Println(v4, v1, v5)
    fmt.Println(v1, v5, v4)
    fmt.Println(v4, v5)
    if (0 <= (^(v4))) {
        v4 += (v4 << ((v1) & 63))
    } else {
        if false {
            v4 = uint32(v1)
            switch ("" + ((("hello" + ("hello" + "foo bar")) + "foo bar") + "")) {
            case "hello":
                var v80 int = -2147483648
                _ = v80
            case "":
                switch (("foo bar" + "foo bar") + ("" + (("" + ("\n\t" + "hello")) + "a"))) {
                case "\n\t":
                    if false {
                        if true {
                            for fi81 := 0; fi81 < 17; fi81++ {
                                var v82 uint64 = 4294967295
                                v83 := v3[fi81]
                                v84 := GenS9{F0: -375, F1: 1, F2: v4}
                                _ = fi81
                                _ = v82
                                _ = v83
                                _ = v84
                            }
                        }
                    }
                case "a":
                    fmt.Println(v1)
                case "hello":
                    switch -1 {
                    case -2147483648:
                        if true {
                            var v85 uint8 = 2
                            _ = v85
                        }
                    case 127:
                        var v86 uint = genFn1()
                        _ = v86
                    case -1:
                        var v87 uint32 = (v4 ^ v4)
                        _ = v87
                    }
                }
            case "a":
                for fi88 := 0; fi88 < 12; fi88++ {
                    if false {
                        v89 := GenS10{F0: -1, F1: true}
                        var v90 int32 = -974
                        for fi91 := 0; fi91 < 16; fi91++ {
                            v92 := GenS2{F0: v90}
                            v90 += -2147483648
                            _ = fi91
                            _ = v92
                        }
                        _ = v89
                        _ = v90
                    } else {
                        v4 = 0
                        v4 -= 1
                        if len(v59) >= 2 {
                            v59 = v59[1:len(v59)]
                        }
                    }
                    v1 += genFn1()
                    v93 := GenS4{F0: 2, F1: 20, F2: (!true)}
                    _ = fi88
                    _ = v93
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
}
