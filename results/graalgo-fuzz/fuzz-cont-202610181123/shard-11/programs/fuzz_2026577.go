package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int) uint {
    v5 := []int{(^(p0)), p0, p0}
    var v6 uint = 18446744073709551615
    if true {
        if true {
            v6 = (v6 << ((2) & 63))
            var v7 uint16 = 1
            _ = v7
        } else {
            if ((1.5 - 3.14159) > (float32(float64(0.0)) / 1e10)) {
                var v8 uint16 = 0
                v6 -= (^(v6))
                p0 -= 83
                _ = v8
            }
        }
    }
    v9 := GenS1{F0: 156}
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return (v6 / ((42) | 1))
}
func genFn1(p0 uint, p1 uint64, p2 uint) int {
    var v10 float32 = 0.1
    fmt.Println(v10, p2)
    fmt.Println(p2)
    p1 += 87
    v11 := []int{-2147483648, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 bool = (((0 != 189) || true) && false)
    var v2 float32 = 1e10
    var v3 int = -168
    v4 := []string{"", "foo bar"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (18446744073709551615 > 224) {
        if (!true) {
            v1 = (("a" + (((("foo bar" + "foo bar") + "hello") + "a") + ("a" + "a"))) == "foo bar")
            v4 = append(v4, "foo bar")
            var v12 uint = 2
            _ = v12
        }
        fmt.Println(v1, v2)
        _ = len(v4)
    }
    if false {
        if true {
            var v13 int = 2147483647
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: 4294967295}
                    var v17 string = "a"
                    v17 = v17
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if v1 {
                    if (((v2 + v2) - 1.5) != v2) {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            var v19 float32 = (1.5 - v2)
                            _ = fi18
                            _ = v19
                        }
                        v1 = v1
                        if true {
                            v20 := GenS2{F0: (!v1)}
                            v3 += v3
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 11; fi21++ {
                                switch "foo bar" {
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "a":
                                    switch (v3 % ((fi21) | 1)) {
                                    case 79:
                                        var v22 uint8 = 186
                                        _ = v22
                                    }
                                case "hello":
                                    v23 := GenS2{F0: false}
                                    _ = v23
                                default:
                                    v13 -= fi21
                                }
                                v24 := GenS1{F0: 1}
                                var v25 uint64 = 18446744073709551615
                                if v1 {
                                    v3 -= (-(fi21))
                                    v2 = (float32(0.1) / v2)
                                    if (-32768 < -32768) {
                                        var v26 string = "a"
                                        fmt.Println(v2)
                                        switch "a" {
                                        case "":
                                            var v27 uint = 255
                                            _ = v27
                                        default:
                                            switch v3 {
                                            case 128:
                                                v28 := GenS3{F0: 1}
                                                _ = v28
                                            case 297:
                                                fmt.Println(v1, v25)
                                            case 1:
                                                v13 -= genFn1(311, (v25 + (v25 ^ v25)), 2)
                                            default:
                                                v29 := GenS4{F0: genFn1(1, v25, 42)}
                                                _ = v29
                                            }
                                        }
                                        _ = v26
                                    } else {
                                        v30 := GenS4{F0: (v3 + (fi21 >> ((18446744073709551615) & 63)))}
                                        var v31 uint32 = 0
                                        v4 = append(v4, "hello")
                                        _ = v30
                                        _ = v31
                                    }
                                } else {
                                    v3 = genFn1(genFn0((fi14 % ((fi21) | 1))), 1, genFn0(-452))
                                    v32 := GenS5{F0: (("foo bar" < "") && false)}
                                    v33 := GenS3{F0: 0}
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi21
                                _ = v24
                                _ = v25
                            }
                            var v34 int16 = int16(v13)
                            _ = v34
                        }
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v36 := GenS6{F0: "hello", F1: 914}
                        var v37 uint = 4294967295
                        var v38 bool = (820 <= -32768)
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    fmt.Println(v13, v1, fi14)
                }
                _ = fi14
            }
            if ((!v1) && (-677 < 42)) {
                v39 := GenS1{F0: 42}
                if v1 {
                    v4 = append(v4, "foo bar")
                    var v40 int64 = 0
                    v4 = append(v4, "a")
                    _ = v40
                }
                _ = v39
            }
            _ = v13
        }
        v2 = 0.0
    } else {
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 float64 = 0.1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi41
            _ = v42
        }
    }
    var v43 uint64 = 4294967295
    for fi44 := 0; fi44 < 6; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            v3 += (v3 << ((42) & 63))
            v46 := GenS1{F0: uint32(v43)}
            switch fi45 {
            case 4294967295:
                var v47 string = ""
                _ = v47
            case 128:
                switch ("" + (("foo bar" + "hello") + "foo bar")) {
                case "a":
                    switch fi44 {
                    case -2147483648:
                        v48 := GenS3{F0: 969}
                        _ = v48
                    case 127:
                        switch genFn1(18446744073709551615, (v43 % (((v43 * v43)) | 1)), 175) {
                        case 703:
                            if v1 {
                                fmt.Println(v2, v43)
                                v43 -= (v43 * (^(v43)))
                                v49 := GenS7{F0: 1, F1: -128}
                                _ = v49
                            }
                        case -955:
                            var v50 string = ((("hello" + "a") + "foo bar") + ("foo bar" + ("a" + "hello")))
                            _ = v50
                        case 2147483647:
                            if ((4294967295 != 42) && (1.5 != v2)) {
                                v43 = 4294967295
                                switch "foo bar" {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = ("hello" + "a")
                                    }
                                }
                            }
                        default:
                            fmt.Println(fi45, v2)
                        }
                        var v51 bool = (191 != 241)
                        _ = v51
                    default:
                        v43 -= v43
                    }
                case "foo bar":
                    v3 += 2147483647
                case "\n\t":
                    var v52 uint32 = 4294967295
                    _ = v52
                }
            case 0:
                if ((("" + ("foo bar" + "\n\t")) + "a") != "hello") {
                    var v53 int8 = 127
                    switch ("" + ("" + "a")) {
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "foo bar":
                        fmt.Println(fi45)
                    case "":
                        v1 = (!v1)
                    default:
                        var v54 uint = 2
                        _ = v54
                    }
                    v55 := GenS8{F0: -446}
                    _ = v53
                    _ = v55
                } else {
                    if (!false) {
                        v56 := GenS6{F0: "\n\t", F1: 1}
                        v3 = fi45
                        _ = v56
                    }
                }
            default:
                if (42 <= 652) {
                    v2 = v2
                    var v57 int32 = 1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v57
                }
            }
            _ = fi45
            _ = v46
        }
        var v58 float64 = (float64(v2) + 3.14159)
        _ = fi44
        _ = v58
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        v1 = (1 >= -32768)
        var v60 float64 = 3.14159
        _ = fi59
        _ = v60
    }
    switch (("foo bar" + "") + ("\n\t" + "")) {
    case "a":
        var v61 float32 = v2
        _ = v61
    case "":
        var v62 uint32 = 0
        _ = v62
    }
    var v63 uint16 = 65535
    v4 = append(v4, "hello")
    fmt.Println(v43, v3)
    v63 += 1
    v64 := GenS7{F0: 699, F1: int8(v43)}
    var v65 float32 = 1.5
    v63 -= 255
    var v66 int64 = 741
    v4 = append(v4, "")
    var v67 float64 = (float64(1.5) - -2.25)
    fmt.Println(v67, v1, v66)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
