package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int32; F2 string }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 bool }

func main() {
    var v1 int8 = 48
    var v2 float32 = (-2.25 / 0.1)
    var v3 uint32 = 1
    var v4 uint = 65535
    var v5 uint16 = 1
    switch "a" {
    case "hello":
        v4 = 2
    case "foo bar":
        switch (("" + ("a" + (("\n\t" + ("a" + "")) + ("foo bar" + "a")))) + "a") {
        case "":
            var v6 int8 = (-(v1))
            _ = v6
        case "foo bar":
            fmt.Println(v4, v1, v5)
        }
    default:
        var v7 uint32 = v3
        _ = v7
    }
    v2 += (v2 / 1.0)
    v8 := GenS1{F0: ""}
    var v9 int32 = -221
    if true {
        v2 = v2
        var v10 uint = 255
        _ = v10
    } else {
        v11 := GenS2{F0: int64(v5)}
        _ = v11
    }
    fmt.Println(v2, v4)
    if false {
        v12 := []int{-2147483648, 9223372036854775807}
        v13 := GenS3{F0: 667}
        _ = v12
        _ = v13
    } else {
        v14 := []int{127, 128}
        switch "foo bar" {
        case "foo bar":
            if (!false) {
                var v15 uint64 = 123
                _ = v15
            } else {
                _ = len(v14)
                var v16 uint = v4
                if false {
                    var v17 int8 = v1
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v19 := GenS2{F0: 1}
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v16 = v16
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            if (false || (true && (32767 < 42))) {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                            _ = fi20
                        }
                        v5 += 1
                        _ = fi18
                        _ = v19
                    }
                    if len(v14) > 0 {
                        v14[0] = -1
                    }
                    _ = v17
                } else {
                    if ((v2 / (v2 * v2)) == (((3.14159 + v2) + (v2 + (v2 - v2))) + 1.5)) {
                        if (v1 != 121) {
                            if (!false) {
                                switch 9223372036854775807 {
                                case -1:
                                    var v21 int64 = 9223372036854775807
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    v14 = append(v14, (fi22 * 1))
                                    if false {
                                        if (v2 != -2.25) {
                                            fmt.Println(v9, v2, v2)
                                        }
                                    } else {
                                        for fi23 := 0; fi23 < 9; fi23++ {
                                            if len(v14) > 0 {
                                                v14[0] = 32768
                                            }
                                            if ((true || true) || true) {
                                                if (273 < 167) {
                                                    fmt.Println(v5)
                                                    var v24 int32 = v9
                                                    _ = v24
                                                } else {
                                                    if (((v1 ^ (^(v1))) < 57) && (-301 > 32768)) {
                                                        v25 := GenS4{F0: -756, F1: (v9 * (-(v9))), F2: "a"}
                                                        v26 := GenS5{F0: 327, F1: 277}
                                                        v14 = append(v14, fi22)
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        fmt.Println(v4)
                                                        v4 += 2
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        var v27 int64 = 1
                                                        _ = v27
                                                    case "foo bar":
                                                        var v28 float64 = 3.14159
                                                        if (-69 >= -2147483648) {
                                                            switch int(v9) {
                                                            case -2147483648:
                                                                if len(v14) > 0 {
                                                                    v14[0] = fi22
                                                                }
                                                            case 4294967295:
                                                                var v29 bool = true
                                                                _ = v29
                                                            case 2147483647:
                                                                v30 := GenS2{F0: -157}
                                                                var v31 int8 = -71
                                                                v14 = append(v14, fi22)
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        }
                                                        fmt.Println(v4)
                                                        _ = v28
                                                    default:
                                                        v5 -= v5
                                                    }
                                                    var v32 uint32 = (v3 >> ((0) & 63))
                                                    var v33 string = ("\n\t" + (("foo bar" + ((("hello" + "hello") + "hello") + "\n\t")) + ("hello" + "")))
                                                    var v34 int8 = v1
                                                    v1 -= -1
                                                    var v35 int32 = -983
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                                if true {
                                                    v36 := GenS6{F0: v5}
                                                    switch "a" {
                                                    case "a":
                                                        v14 = append(v14, (fi23 * fi22))
                                                    default:
                                                        if len(v14) > 0 {
                                                            v14[0] = (fi22 * 4294967295)
                                                        }
                                                    }
                                                    _ = v36
                                                } else {
                                                    var v37 float64 = 0.1
                                                    v38 := GenS4{F0: 0, F1: v9, F2: "hello"}
                                                    fmt.Println(fi22, v37)
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v39 := GenS4{F0: -449, F1: v9, F2: "hello"}
                                                _ = v39
                                            }
                                            if false {
                                                var v40 uint64 = 18446744073709551615
                                                fmt.Println(v2)
                                                if true {
                                                    _ = len(v14)
                                                }
                                                _ = v40
                                            } else {
                                                fmt.Println(v2)
                                            }
                                            _ = fi23
                                        }
                                        v4 -= (v16 << (((v16 ^ v4)) & 63))
                                        fmt.Println(v16)
                                    }
                                    if true {
                                        v41 := GenS2{F0: 42}
                                        fmt.Println(fi22)
                                        var v42 uint64 = 954
                                        _ = v41
                                        _ = v42
                                    } else {
                                        v43 := GenS7{F0: v16}
                                        _ = v43
                                    }
                                    _ = fi22
                                }
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    v4 = (^(v16))
                                    var v45 string = (("a" + (("" + "\n\t") + "hello")) + ("" + "foo bar"))
                                    _ = fi44
                                    _ = v45
                                }
                            } else {
                                v46 := GenS2{F0: int64(v1)}
                                for fi47 := 0; fi47 < 7; fi47++ {
                                    var v48 uint64 = 1
                                    v14 = append(v14, fi47)
                                    if false {
                                        if len(v14) > 0 {
                                            v14[0] = fi47
                                        }
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 9; fi49++ {
                                    v14 = append(v14, (-(fi49)))
                                    var v50 uint8 = 2
                                    _ = fi49
                                    _ = v50
                                }
                                _ = v46
                            }
                            switch "a" {
                            case "":
                                switch ("" + "foo bar") {
                                case "hello":
                                    v9 = v9
                                case "a":
                                    var v51 int16 = -1
                                    _ = v51
                                }
                            case "hello":
                                v4 += 255
                            case "a":
                                fmt.Println(v16, v2)
                            default:
                                var v52 float32 = v2
                                _ = v52
                            }
                            var v53 uint64 = 248
                            _ = v53
                        } else {
                            v54 := GenS8{F0: -1, F1: 0}
                            var v55 uint32 = (^(v3))
                            if false {
                                fmt.Println(v2, v4)
                            } else {
                                if len(v14) > 0 {
                                    v14[0] = 54
                                }
                            }
                            _ = v54
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 1; fi56++ {
                            if (v1 > v1) {
                                var v57 int32 = (v9 - v9)
                                var v58 int16 = -1
                                switch fi56 {
                                case 2147483647:
                                    var v59 int64 = 0
                                    _ = v59
                                case 32767:
                                    v60 := GenS5{F0: v16, F1: -9223372036854775808}
                                    _ = v60
                                default:
                                    v61 := GenS9{F0: 42}
                                    _ = v61
                                }
                                _ = v57
                                _ = v58
                            }
                            var v62 int8 = -128
                            fmt.Println(v2)
                            _ = fi56
                            _ = v62
                        }
                    } else {
                        for fi63 := 0; fi63 < 3; fi63++ {
                            for fi64 := 0; fi64 < 7; fi64++ {
                                fmt.Println(v16, v5)
                                _ = fi64
                            }
                            _ = fi63
                        }
                    }
                }
                _ = v16
            }
        case "\n\t":
            v65 := GenS4{F0: 32767, F1: v9, F2: "a"}
            _ = v65
        case "":
            if true {
                if true {
                    var v66 string = ("a" + "")
                    _ = v66
                }
            } else {
                v4 = 18446744073709551615
                var v67 uint64 = 278
                _ = v67
            }
        }
        _ = v14
    }
    v1 = v1
    if true {
        var v68 int8 = (v1 + 0)
        v2 = ((((((0.1 + -2.25) + v2) / ((float32(0.0) + v2) + v2)) * float32(float64(0.1))) * (v2 / 1.5)) / v2)
        v3 += v3
        _ = v68
    }
    v69 := GenS3{F0: 0}
    v70 := GenS5{F0: (v4 + 4294967295), F1: -9223372036854775808}
    fmt.Println(v2)
    if false {
        switch -407 {
        case 128:
            var v71 uint64 = 63
            _ = v71
        case -414:
            if ((v9 % (((v9 + (v9 + v9))) | 1)) >= (v9 * (v9 % ((595) | 1)))) {
                fmt.Println(v2, v4)
                var v72 uint16 = 0
                var v73 bool = false
                _ = v72
                _ = v73
            }
        case 42:
            var v74 float64 = (float64((float32(0.1) + (v2 + float32(-2.25)))) / 1.0)
            _ = v74
        default:
            for fi75 := 0; fi75 < 11; fi75++ {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v76 := GenS10{F0: 103, F1: (uint16(v4) == (v5 >> ((v4) & 63)))}
                    _ = v76
                case "\n\t":
                    var v77 uint8 = 127
                    _ = v77
                }
                for fi78 := 0; fi78 < 11; fi78++ {
                    v9 += -758
                    v79 := GenS1{F0: (("\n\t" + ("\n\t" + "")) + "\n\t")}
                    _ = fi78
                    _ = v79
                }
                _ = fi75
            }
        }
        v4 = v4
    }
    v80 := []int{32767, 76, 42, -390}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v8
    _ = v80
    _ = v9
}
