package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint32; F2 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int32; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint64; F1 uint }
func genFn0(p0 bool, p1 uint64, p2 uint32) int32 {
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 = (true || ((-273 <= 1000) || p0))
        var v6 string = "hello"
        v7 := GenS1{F0: (p1 * (p1 + (p1 % ((uint64(p2)) | 1)))), F1: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = 41
    v9 := GenS1{F0: (^(p1)), F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 233
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    v3 -= 0.1
    v3 = v3
    if true {
        var v10 int = -2147483648
        switch (v10 / (((v10 >> ((2) & 63))) | 1)) {
        case -228:
            v11 := GenS1{F0: 4294967295, F1: 838}
            _ = v11
        case 32767:
            if ("" >= ("hello" + "a")) {
                if (false || (0 >= 819)) {
                    v12 := []int{int(v2), (v10 << ((906) & 63)), int(v2)}
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 bool = false
                        var v15 uint16 = 2
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                }
                v16 := []int{(v10 % (((v10 << ((571) & 63))) | 1)), 127, (v10 - (v10 % (((v10 * v10)) | 1)))}
                _ = v16
            }
        }
        v17 := GenS2{F0: 4294967295, F1: -1}
        _ = v10
        _ = v17
    }
    var v18 uint8 = 140
    v2 = -9223372036854775808
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            v4[("hello" + ("" + "\n\t"))] = (("" + "foo bar") + "\n\t")
            for fi20 := 0; fi20 < 7; fi20++ {
                if true {
                    switch ("" + "a") {
                    case "a":
                        var v21 uint32 = 689
                        _ = v21
                    case "foo bar":
                        v22 := GenS1{F0: 4294967295, F1: 1}
                        _ = v22
                    default:
                        v23 := GenS1{F0: 862, F1: 65535}
                        _ = v23
                    }
                } else {
                    _ = len(v1)
                    v24 := GenS1{F0: 2, F1: 65535}
                    v1[fi20] = 395
                    var v25 float32 = ((3.14159 * ((v3 - (1e10 * (v3 - (v3 / 1.0)))) - v3)) / (0.1 / 1e10))
                    _ = v24
                    _ = v25
                }
                v3 = float32(0.0)
                _ = fi20
            }
        } else {
            var v26 int = (-(fi19))
            _ = v26
        }
        _ = fi19
    }
    v18 -= v18
    if (false && true) {
        var v27 bool = true
        v28 := GenS3{F0: 4294967295, F1: -1}
        v29 := GenS2{F0: 65535, F1: 975}
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: -648, F1: -383}
    switch ("foo bar" + "\n\t") {
    case "":
        var v31 string = "a"
        _ = v31
    case "\n\t":
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS2{F0: uint(v18), F1: -884}
            _ = fi32
            _ = v33
        }
    case "a":
        v34 := GenS1{F0: 2, F1: 98}
        _ = v34
    }
    v35 := GenS2{F0: 65535, F1: -1}
    fmt.Println(v3, v2)
    var v36 float32 = (v3 - 0.1)
    if true {
        v37 := []int{}
        v38 := GenS4{F0: int16(v2)}
        switch 32767 {
        case -1:
            if false {
                if true {
                    fmt.Println(v2)
                    v3 = v36
                } else {
                    switch 4294967295 {
                    case -9223372036854775808:
                        for fi39 := 0; fi39 < 17; fi39++ {
                            if len(v37) > 0 {
                                v37[0] = fi39
                            }
                            _ = fi39
                        }
                    case 127:
                        var v40 int64 = (v2 >> ((4294967295) & 63))
                        _ = v40
                    case 4294967295:
                        v18 -= v18
                    default:
                        var v41 float64 = 0.0
                        _ = v41
                    }
                    var v42 uint8 = 0
                    v3 += v36
                    _ = v42
                }
                var v43 string = ((("hello" + ("\n\t" + (("\n\t" + "\n\t") + (((("hello" + "\n\t") + "foo bar") + "") + "hello")))) + ("\n\t" + "")) + "")
                _ = v43
            }
        case 4294967295:
            fmt.Println(v36, v2, v3)
        }
        _ = v37
        _ = v38
    }
    switch ((("" + "a") + "") + "a") {
    case "":
        if false {
            if true {
                if true {
                    v44 := GenS1{F0: 939, F1: 0}
                    _ = v44
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    v36 += (1e10 + -2.25)
                    _ = fi45
                }
            } else {
                var v46 uint = 1
                var v47 string = "\n\t"
                v36 = 0.1
                _ = v46
                _ = v47
            }
            if (!true) {
                fmt.Println(v3, v2)
                for fi48 := 0; fi48 < 13; fi48++ {
                    v49 := GenS3{F0: -2147483648, F1: 32768}
                    if (!true) {
                        for fi50 := 0; fi50 < 16; fi50++ {
                            if false {
                                switch (fi48 + fi48) {
                                case 9223372036854775807:
                                    var v51 float32 = v3
                                    _ = v51
                                default:
                                    fmt.Println(v36)
                                }
                                var v52 uint = 18446744073709551615
                                fmt.Println(fi48, v52)
                                _ = v52
                            }
                            _ = fi50
                        }
                        v2 = (v2 + v2)
                    } else {
                        fmt.Println(v18, v3, v2)
                    }
                    _ = fi48
                    _ = v49
                }
                v53 := GenS4{F0: 1}
                _ = v53
            } else {
                for fi54 := 0; fi54 < 13; fi54++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v2, v36, v3)
                    case "hello":
                        v3 = (float32(0.0) / v3)
                    case "\n\t":
                        fmt.Println(v3, fi54, fi54)
                    }
                    _ = fi54
                }
                if false {
                    if false {
                        for fi55 := 0; fi55 < 6; fi55++ {
                            var v56 uint = 2
                            _ = fi55
                            _ = v56
                        }
                        _ = len(v1)
                    } else {
                        delete(v4, (("hello" + "hello") + "foo bar"))
                    }
                    var v57 float32 = (v36 - v3)
                    v57 -= v57
                    _ = v57
                } else {
                    v58 := v4["foo bar"]
                    _ = v58
                }
            }
            switch 128 {
            case 869:
                for fi59 := 0; fi59 < 20; fi59++ {
                    fmt.Println(v36, v3)
                    v2 -= v2
                    v60 := []int{(fi59 + fi59)}
                    _ = fi59
                    _ = v60
                }
            case 32768:
                switch 4294967295 {
                case 128:
                    v36 += v36
                default:
                    if true {
                        v2 += (^(v2))
                        for fi61 := 0; fi61 < 13; fi61++ {
                            v3 += (v3 - ((3.14159 + -2.25) * (((v3 * v36) / 0.1) + (v3 * (1.5 * v36)))))
                            _ = fi61
                        }
                        v36 += 0.1
                    }
                }
            case 9223372036854775807:
                var v62 uint = 4294967295
                _ = v62
            default:
                for fi63 := 0; fi63 < 16; fi63++ {
                    v64 := GenS3{F0: fi63, F1: -934}
                    v4[("\n\t" + "hello")] = "a"
                    delete(v1, (fi63 - fi63))
                    _ = fi63
                    _ = v64
                }
            }
        }
    case "hello":
        var v65 int64 = 0
        _ = v65
    case "foo bar":
        var v66 uint32 = 2147483648
        _ = v66
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            v2 -= (v2 >> ((0) & 63))
        case "hello":
            fmt.Println(v2)
        default:
            fmt.Println(v3)
        }
        if (-2147483648 < 982) {
            fmt.Println(v2, v18)
            v67 := GenS2{F0: 2, F1: 235}
            fmt.Println(v2, v36)
            _ = v67
        } else {
            v18 = v18
            switch "" {
            case "foo bar":
                v18 -= 58
            default:
                v68 := GenS5{F0: false, F1: 2, F2: false}
                _ = v68
            }
            v2 -= (-(v2))
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            if true {
                var v69 int8 = 113
                _ = v69
            } else {
                switch -113 {
                case 2147483647:
                    v36 -= (v36 / -2.25)
                case 874:
                    if (!true) {
                        if (false && true) {
                            var v70 int = -1
                            var v71 uint64 = 509
                            v72 := GenS6{F0: 127}
                            _ = v70
                            _ = v71
                            _ = v72
                        }
                    } else {
                        for fi73 := 0; fi73 < 11; fi73++ {
                            var v74 int32 = genFn0(false, 18446744073709551615, 2147483648)
                            _ = fi73
                            _ = v74
                        }
                        for fi75 := 0; fi75 < 10; fi75++ {
                            fmt.Println(v36)
                            if false {
                                for fi76 := 0; fi76 < 7; fi76++ {
                                    switch (fi76 >> ((4294967295) & 63)) {
                                    case 325:
                                        v3 += v3
                                    case 9223372036854775807:
                                        v77 := GenS7{F0: 255}
                                        _ = v77
                                    default:
                                        if (true && (-2147483648 == -2147483648)) {
                                            v78 := GenS1{F0: 4294967295, F1: 0}
                                            v36 += (float32(((float64(v3) / 3.14159) + 1e10)) / 1e10)
                                            _ = v78
                                        } else {
                                            var v79 uint16 = 2
                                            v80 := []int{fi75, fi75}
                                            v81 := GenS8{F0: false}
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                    _ = fi76
                                }
                                v3 -= 3.14159
                            }
                            v82 := []int{}
                            _ = fi75
                            _ = v82
                        }
                    }
                case -147:
                    if true {
                        if ((1e10 >= float64((((v3 * 1.5) - 0.1) + 0.1))) || (!(3.14159 != 1e10))) {
                            v83 := GenS8{F0: true}
                            delete(v4, ("\n\t" + ("" + "\n\t")))
                            fmt.Println(v3, v3)
                            _ = v83
                        }
                        v18 = 42
                        v84 := GenS9{F0: 627, F1: ((v18 < v18) && (false && false))}
                        _ = v84
                    }
                default:
                    v36 += float32(1.5)
                }
                v1[-439] = 2147483647
            }
        case "":
            v85 := []int{4294967295, 127, -9223372036854775808}
            _ = v85
        }
    }
    switch 1 {
    case 42:
        for fi86 := 0; fi86 < 20; fi86++ {
            for fi87 := 0; fi87 < 3; fi87++ {
                var v88 int8 = 0
                v89 := GenS10{F0: 523, F1: (^(v88)), F2: (v18 % (((v18 % ((254) | 1))) | 1))}
                v90 := GenS11{F0: 42, F1: 255}
                _ = fi87
                _ = v88
                _ = v89
                _ = v90
            }
            v91 := []int{(fi86 ^ (fi86 >> ((42) & 63))), fi86, fi86, fi86}
            _ = fi86
            _ = v91
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
}
