package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 int) int32 {
    v5 := []int{p0, 127, 42, -1}
    switch (p0 & 1) {
    case 32767:
        p0 += p0
    case -797:
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 20; fi6++ {
                p0 = (fi6 + 32768)
                p0 = 186
                if true {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi7
                    }
                } else {
                    if false {
                        if false {
                            if true {
                                var v8 int32 = -267
                                _ = v8
                            }
                            var v9 uint = 0
                            _ = v9
                        } else {
                            if false {
                                for fi10 := 0; fi10 < 9; fi10++ {
                                    p0 -= fi6
                                    var v11 uint = 18446744073709551615
                                    _ = fi10
                                    _ = v11
                                }
                                v12 := GenS1{F0: 2147483647, F1: true}
                                _ = v12
                            } else {
                                var v13 float64 = 0.1
                                v13 = float64(0.1)
                                var v14 string = (("foo bar" + "") + "")
                                _ = v13
                                _ = v14
                            }
                            var v15 bool = (!(255 == 255))
                            for fi16 := 0; fi16 < 17; fi16++ {
                                p0 += fi16
                                if false {
                                    v17 := GenS2{F0: 125, F1: 70, F2: 140}
                                    v18 := GenS1{F0: 834, F1: (((0.1 <= 1.5) && (664 != 9223372036854775807)) && v15)}
                                    switch "foo bar" {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v15, fi16, v15)
                                        }
                                    default:
                                        var v19 int32 = 2147483647
                                        _ = v19
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi16
                            }
                            _ = v15
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        p0 = 32768
                    }
                    var v20 int16 = 32767
                    _ = v20
                }
                _ = fi6
            }
        case "foo bar":
            for fi21 := 0; fi21 < 12; fi21++ {
                p0 -= p0
                _ = fi21
            }
        case "\n\t":
            v22 := GenS2{F0: 0, F1: 4, F2: 1}
            _ = v22
        default:
            if false {
                var v23 uint64 = 18446744073709551615
                var v24 int16 = 0
                var v25 int64 = 0
                _ = v23
                _ = v24
                _ = v25
            }
        }
    default:
        if (0.1 == float64(float32(1.5))) {
            var v26 uint32 = 2147483648
            var v27 uint64 = 2
            var v28 float32 = 1e10
            _ = v26
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int = -2147483648
                v31 := GenS3{F0: 439, F1: 746, F2: 2}
                if ((!(true || true)) && false) {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        fmt.Println(p0, v30, v30)
                        v33 := GenS2{F0: 42, F1: 42, F2: 99}
                        _ = fi32
                        _ = v33
                    }
                } else {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        var v35 uint64 = 0
                        var v36 uint16 = 302
                        var v37 float64 = (3.14159 + 1.5)
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    var v38 uint8 = 1
                    _ = v38
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1() bool {
    var v39 float64 = 3.14159
    if ((!(!(!(42 == 223)))) && (1.5 < (((1e10 - (1.5 - 0.0)) - 1e10) * (0.1 * float32(((0.0 / (((float64(0.0) + v39) * (1e10 * v39)) * -2.25)) * v39)))))) {
        v40 := GenS3{F0: 42, F1: 65535, F2: 2}
        _ = v40
    }
    v41 := GenS4{F0: 4294967295, F1: 580, F2: 2147483648}
    _ = v39
    _ = v41
    return false
}

func main() {
    var v1 bool = (true && (!false))
    v2 := []string{"foo bar", "hello", ""}
    var v3 uint64 = 0
    v4 := make(map[string]uint32)
    v1 = v1
    v1 = v1
    delete(v4, "foo bar")
    if (-2147483648 >= 128) {
        if ((!v1) && v1) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
    }
    v3 += 42
    fmt.Println(v3, v1)
    switch "\n\t" {
    case "a":
        var v42 uint16 = 109
        _ = v42
    case "hello":
        v43 := v4[("hello" + "foo bar")]
        _ = v43
    case "foo bar":
        switch ("foo bar" + ("\n\t" + "foo bar")) {
        case "\n\t":
            v44 := GenS1{F0: 127, F1: v1}
            _ = v44
        case "":
            if (2 > 4294967295) {
                var v45 uint32 = 601
                fmt.Println(v1, v45, v45)
                v45 = v45
                _ = v45
            } else {
                if v1 {
                    fmt.Println(v3, v1)
                    var v46 int8 = -115
                    _ = v46
                } else {
                    fmt.Println(v3, v1)
                }
                _ = len(v2)
            }
        case "a":
            fmt.Println(v3, v3)
        }
    default:
        for fi47 := 0; fi47 < 20; fi47++ {
            v1 = (v1 || v1)
            if v1 {
                if (!false) {
                    v3 -= (v3 ^ v3)
                    v3 = (v3 ^ v3)
                }
                for fi48 := 0; fi48 < 2; fi48++ {
                    var v49 float32 = 0.0
                    v3 += (v3 >> ((42) & 63))
                    _ = fi48
                    _ = v49
                }
                for fi50 := 0; fi50 < 5; fi50++ {
                    fmt.Println(fi50)
                    var v51 float32 = float32(float64((0.1 / float32(-2.25))))
                    _ = fi50
                    _ = v51
                }
            } else {
                var v52 uint8 = 213
                _ = v52
            }
            fmt.Println(v1, fi47, v1)
            _ = fi47
        }
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v3, v1)
    for fi53 := 0; fi53 < 13; fi53++ {
        var v54 uint32 = 4294967295
        if false {
            fmt.Println(v54, v3, fi53)
            if v1 {
                if v1 {
                    v54 += v54
                } else {
                    fmt.Println(v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        } else {
            v54 = v54
            v55 := GenS4{F0: v54, F1: 32767, F2: (v54 << ((65535) & 63))}
            _ = v55
        }
        _ = fi53
        _ = v54
    }
    v56 := GenS5{F0: (255 != 255), F1: 0}
    fmt.Println(v3, v3)
    for fi57 := 0; fi57 < 9; fi57++ {
        if (42 < 1) {
            v58 := GenS5{F0: (!((3.14159 > 1e10) && (!v1))), F1: 2}
            switch "\n\t" {
            case "foo bar":
                var v59 uint8 = 42
                _ = v59
            default:
                v60 := GenS5{F0: (153 > 42), F1: uint(v3)}
                _ = v60
            }
            if (0 > 0) {
                if len(v2) > 0 {
                    v2[0] = ""
                }
            } else {
                for fi61 := 0; fi61 < 11; fi61++ {
                    v1 = (true && (v1 || (!(!false))))
                    var v62 float64 = -2.25
                    switch "a" {
                    case "hello":
                        var v63 uint64 = v3
                        switch (fi57 * -807) {
                        case -1:
                            if (v1 || ((v3 >> ((665) & 63)) > (v3 * (v3 ^ (v63 % ((442) | 1)))))) {
                                var v64 int16 = -160
                                var v65 uint64 = v3
                                v66 := GenS5{F0: v1, F1: 18446744073709551615}
                                _ = v64
                                _ = v65
                                _ = v66
                            } else {
                                switch "" {
                                case "":
                                    if (v1 || (((("hello" + ("" + ((("" + "hello") + "") + ("" + (("a" + ("a" + (("" + (("" + "\n\t") + "a")) + ("a" + "")))) + "a"))))) > "a") || (-910 > -815)) || v1)) {
                                        v67 := GenS6{F0: 626}
                                        if false {
                                            if v1 {
                                                _ = len(v4)
                                                _ = len(v2)
                                            }
                                        } else {
                                            v3 = v63
                                            if (v1 || (!v1)) {
                                                fmt.Println(v62, fi57)
                                                var v68 float64 = (v62 - v62)
                                                v69 := GenS1{F0: fi57, F1: v1}
                                                _ = v68
                                                _ = v69
                                            }
                                            v70 := GenS7{F0: genFn0(fi61)}
                                            _ = v70
                                        }
                                        _ = v67
                                    }
                                    var v71 int16 = 1
                                    _ = v71
                                }
                            }
                        case 127:
                            if (2147483648 <= 1) {
                                fmt.Println(fi61)
                            }
                        case -192:
                            var v72 uint = 4294967295
                            _ = v72
                        }
                        _ = v63
                    case "":
                        var v73 uint = 42
                        _ = v73
                    case "\n\t":
                        var v74 float32 = 1.5
                        switch "foo bar" {
                        case "\n\t":
                            v3 = (v3 ^ v3)
                        default:
                            var v75 uint64 = 2
                            _ = v75
                        }
                        if true {
                            fmt.Println(fi61, v1, v3)
                            fmt.Println(v62, fi57)
                        } else {
                            if (-63 == 0) {
                                v1 = (2147483647 >= -1)
                                var v76 uint16 = 32767
                                fmt.Println(v62, v76)
                                _ = v76
                            }
                        }
                        _ = v74
                    }
                    _ = fi61
                    _ = v62
                }
            }
            _ = v58
        }
        v77 := GenS2{F0: 255, F1: 18446744073709551615, F2: 1}
        fmt.Println(fi57)
        _ = fi57
        _ = v77
    }
    v78 := GenS8{F0: 2, F1: "foo bar", F2: 127}
    v3 -= v3
    var v79 string = ("hello" + ((("" + "") + "hello") + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v78
    _ = v79
}
