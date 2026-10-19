package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int8; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 string }
func genFn0(p0 int64) int16 {
    if (32767 <= 32767) {
        fmt.Println(p0)
    } else {
        v7 := []int{32768, 32768, -234}
        var v8 float64 = (0.0 / 0.1)
        var v9 float64 = 3.14159
        _ = v7
        _ = v8
        _ = v9
    }
    if false {
        var v10 uint32 = 540
        v11 := GenS1{F0: 42, F1: v10, F2: 128}
        var v12 float32 = (-2.25 + 0.1)
        _ = v10
        _ = v11
        _ = v12
    } else {
        if true {
            v13 := GenS1{F0: -128, F1: 65535, F2: 224}
            v14 := []int{100}
            p0 = (p0 ^ (p0 >> ((948) & 63)))
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint16 = 255
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := []int{(fi15 * -641), (-(fi17)), fi15}
                fmt.Println(fi15)
                var v19 uint = 112
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                if true {
                    v16 = 255
                    if false {
                        if (!false) {
                            v21 := GenS2{F0: "hello"}
                            fmt.Println(fi15, v16)
                            _ = v21
                        }
                        v22 := []int{fi15, (fi20 & (fi15 % ((fi20) | 1))), -9223372036854775808}
                        _ = v22
                    }
                } else {
                    switch "\n\t" {
                    case "":
                        var v23 uint64 = 2
                        _ = v23
                    case "foo bar":
                        if false {
                            var v24 uint32 = 42
                            v16 -= (v16 | (v16 / (((^(v16))) | 1)))
                            _ = v24
                        } else {
                            var v25 float64 = 1.5
                            _ = v25
                        }
                        var v26 uint8 = 255
                        v27 := GenS1{F0: 42, F1: 2147483647, F2: v26}
                        switch (fi20 % ((fi15) | 1)) {
                        case 2147483647:
                            switch "a" {
                            case "":
                                var v28 uint8 = 0
                                _ = v28
                            case "\n\t":
                                var v29 int16 = -779
                                _ = v29
                            case "foo bar":
                                var v30 int = (fi20 - 2147483647)
                                _ = v30
                            }
                        }
                        if false {
                            var v31 int = fi20
                            var v32 float32 = 1e10
                            v33 := []int{}
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            if (false && true) {
                                var v34 int8 = 100
                                v35 := GenS3{F0: 255}
                                v34 += v34
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = v26
                        _ = v27
                    }
                    if true {
                        var v36 bool = (0 <= 255)
                        _ = v36
                    }
                    fmt.Println(fi20, fi20, p0)
                }
                _ = fi20
            }
            _ = fi15
            _ = v16
        }
    }
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{405, 2, 2147483647, 2}
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 float64 = (0.0 * float64(((1e10 * (float32((((0.0 + 1e10) * 1.5) / (-2.25 + ((0.0 / 1.0) / -2.25)))) - 0.1)) * (1e10 / 1.5))))
    v5 := []int{-2147483648}
    var v6 uint16 = 2
    v6 += (v6 >> ((957) & 63))
    for fi37 := 0; fi37 < 20; fi37++ {
        if (false && (!(!(409 < 0)))) {
            var v38 float64 = 3.14159
            if (true || false) {
                v4 = (v4 / (0.1 + 1e10))
                var v39 float64 = -2.25
                switch fi37 {
                case 2147483647:
                    _ = len(v5)
                }
                _ = v39
            }
            _ = v38
        }
        _ = fi37
    }
    fmt.Println(v2, v4)
    _ = len(v3)
    v2 = ((v2 + v2) + v2)
    for fi40 := 0; fi40 < 1; fi40++ {
        for fi41 := 0; fi41 < 3; fi41++ {
            v42 := GenS4{F0: 0, F1: -2147483648}
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    switch "foo bar" {
    case "a":
        v6 -= 59
    case "hello":
        fmt.Println(v2, v2, v6)
    case "foo bar":
        v4 = (0.1 * (3.14159 * v4))
    }
    v2 = ""
    for fi43 := 0; fi43 < 6; fi43++ {
        if false {
            for fi44 := 0; fi44 < 6; fi44++ {
                if (true || true) {
                    var v45 bool = true
                    var v46 int8 = -1
                    _ = v45
                    _ = v46
                } else {
                    fmt.Println(v4)
                    v2 = "foo bar"
                    var v47 bool = (0 <= 956)
                    _ = v47
                }
                _ = fi44
            }
            switch fi43 {
            case 2147483647:
                if true {
                    _ = len(v3)
                    v5 = append(v5, 9223372036854775807)
                    v48 := GenS5{F0: 1, F1: -274, F2: 4294967295}
                    _ = v48
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch (v2 + v2) {
                    case "foo bar":
                        if false {
                            v2 = v2
                            v4 += v4
                        } else {
                            for fi49 := 0; fi49 < 17; fi49++ {
                                var v50 uint = 446
                                v2 = "foo bar"
                                _ = fi49
                                _ = v50
                            }
                            v6 = (^(v6))
                            var v51 int8 = -1
                            _ = v51
                        }
                    case "hello":
                        v52 := GenS5{F0: 0, F1: 985, F2: 2147483648}
                        _ = v52
                    case "\n\t":
                        if (1e10 == v4) {
                            switch fi43 {
                            case 328:
                                var v53 uint64 = 4294967295
                                _ = v53
                            default:
                                if true {
                                    v54 := GenS6{F0: 0, F1: 505, F2: -815}
                                    v4 = v4
                                    _ = v54
                                }
                            }
                            fmt.Println(v2)
                        }
                    default:
                        var v55 uint8 = 128
                        _ = v55
                    }
                    if (true && false) {
                        var v56 bool = true
                        if v56 {
                            v2 = v2
                        } else {
                            for fi57 := 0; fi57 < 16; fi57++ {
                                switch v2 {
                                case "a":
                                    v2 = ((("foo bar" + "") + "hello") + ((v2 + v2) + (v2 + "foo bar")))
                                case "":
                                    v58 := GenS7{F0: v2}
                                    _ = v58
                                case "foo bar":
                                    if (v56 || v56) {
                                        v2 = (("\n\t" + (("hello" + "\n\t") + v2)) + "")
                                        v3[497] = 4294967295
                                    }
                                default:
                                    var v59 int32 = 2147483647
                                    _ = v59
                                }
                                v56 = false
                                _ = fi57
                            }
                            v4 += v4
                        }
                        _ = v56
                    } else {
                        var v60 int = (-(fi43))
                        v4 += float64((0.0 / 3.14159))
                        _ = v60
                    }
                }
            case -57:
                for fi61 := 0; fi61 < 5; fi61++ {
                    fmt.Println(v6, v4)
                    fmt.Println(v2, v2, fi43)
                    _ = fi61
                }
            case 127:
                var v62 float32 = 3.14159
                _ = v62
            }
        }
        _ = fi43
    }
    v6 = 983
    for fi63 := 0; fi63 < 4; fi63++ {
        v6 = (v6 & 65535)
        v2 = ("foo bar" + v2)
        v64 := GenS8{F0: -69, F1: 2147483647}
        _ = fi63
        _ = v64
    }
    fmt.Println(v2, v2)
    if (v2 != (v2 + (v2 + "\n\t"))) {
        var v65 float64 = v4
        for fi66 := 0; fi66 < 6; fi66++ {
            v6 -= 532
            var v67 uint16 = v6
            _ = fi66
            _ = v67
        }
        _ = v65
    }
    fmt.Println(v2, v4, v4)
    for fi68 := 0; fi68 < 7; fi68++ {
        switch (fi68 % ((int(v6)) | 1)) {
        case 4294967295:
            switch "\n\t" {
            case "\n\t":
                switch fi68 {
                case 923:
                    fmt.Println(fi68)
                case 2147483647:
                    switch -759 {
                    case 290:
                        v69 := GenS9{F0: 4294967295}
                        _ = v69
                    case 32768:
                        v70 := GenS10{F0: (^(v6)), F1: "\n\t"}
                        _ = v70
                    case 0:
                        var v71 int64 = 0
                        _ = v71
                    }
                }
            }
        case 32768:
            v72 := GenS5{F0: 0, F1: 0, F2: 42}
            _ = v72
        case -38:
            fmt.Println(v2, v6, v2)
        }
        _ = fi68
    }
    var v73 int32 = -1
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
