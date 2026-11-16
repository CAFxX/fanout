package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS13 struct { F0 uint32; F1 string }
func genFn0(p0 bool, p1 uint32, p2 string) bool {
    var v7 uint16 = 255
    v7 -= 42
    var v8 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int32) uint64 {
    v9 := GenS1{F0: 142, F1: 4294967295}
    var v10 string = (("foo bar" + "foo bar") + "foo bar")
    switch 42 {
    case 4294967295:
        v11 := []int{710, 697, 236, 32767}
        _ = v11
    case -544:
        switch 42 {
        case 128:
            v12 := GenS2{F0: 42}
            _ = v12
        case -2147483648:
            if false {
                var v13 int32 = -2147483648
                if ((p0 != p0) && false) {
                    v14 := GenS2{F0: 2}
                    switch ("hello" + "") {
                    case "":
                        fmt.Println(p0)
                    case "foo bar":
                        var v15 int32 = 0
                        _ = v15
                    case "hello":
                        p0 = v13
                    }
                    var v16 float64 = 0.0
                    _ = v14
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if true {
                            for fi18 := 0; fi18 < 14; fi18++ {
                                p0 = -522
                                v19 := []int{fi18, fi17, (fi18 >> ((65535) & 63))}
                                switch "" {
                                case "hello":
                                    switch fi17 {
                                    case -253:
                                        if false {
                                            fmt.Println(fi18, v10)
                                            if true {
                                                fmt.Println(p0, v10, v10)
                                            }
                                            p0 = v13
                                        }
                                    case 127:
                                        if false {
                                            var v20 uint32 = 4294967295
                                            switch (fi17 - (fi18 & fi17)) {
                                            case 0:
                                                fmt.Println(v13)
                                            case -340:
                                                fmt.Println(p0, p0)
                                            case 32767:
                                                var v21 uint8 = 128
                                                _ = v21
                                            }
                                            v13 = (v13 ^ v13)
                                            fmt.Println(fi18)
                                            _ = v20
                                        } else {
                                            if len(v19) > 0 {
                                                v19[0] = (fi18 / ((-974) | 1))
                                            }
                                            if len(v19) > 0 {
                                                v19[0] = (fi17 / ((fi18) | 1))
                                            }
                                            v22 := GenS2{F0: 42}
                                            _ = v22
                                        }
                                    case 0:
                                        var v23 float32 = (((0.0 * (((-2.25 * 1e10) + ((0.0 - 1.5) / 1.0)) - float32(-2.25))) * 0.0) + -2.25)
                                        _ = v23
                                    default:
                                        var v24 int = fi17
                                        _ = v24
                                    }
                                case "a":
                                    if (!true) {
                                        if true {
                                            fmt.Println(fi17)
                                        }
                                        switch v10 {
                                        case "a":
                                            if true {
                                                var v25 int = 0
                                                _ = v25
                                            }
                                        }
                                    } else {
                                        fmt.Println(v13, fi17)
                                    }
                                case "foo bar":
                                    if (42 == 127) {
                                        fmt.Println(fi17, fi17)
                                        var v26 uint16 = 1
                                        v27 := GenS3{F0: 4294967295, F1: v10}
                                        _ = v26
                                        _ = v27
                                    }
                                default:
                                    v10 = ""
                                }
                                _ = fi18
                                _ = v19
                            }
                            var v28 int64 = 7
                            _ = v28
                        } else {
                            v29 := GenS4{F0: 56}
                            var v30 int32 = 2147483647
                            _ = v29
                            _ = v30
                        }
                        _ = fi17
                    }
                    v13 = v13
                }
                p0 = p0
                _ = v13
            }
        case 4294967295:
            v31 := []int{836, 0, 0, -188}
            _ = v31
        default:
            for fi32 := 0; fi32 < 4; fi32++ {
                for fi33 := 0; fi33 < 9; fi33++ {
                    fmt.Println(p0)
                    _ = fi33
                }
                _ = fi32
            }
        }
    case 14:
        v10 = v10
    default:
        p0 -= (p0 % ((p0) | 1))
    }
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 bool = (!(180 > 95))
    v2 := []string{}
    var v3 uint = 878
    v4 := []int{9223372036854775807, 0, 4294967295, 196}
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    v3 = v3
    if ((1e10 * float64(float32(((0.1 * 0.1) / 1.0)))) < float64((1.5 + (1.5 * (1e10 - (0.0 * ((3.14159 + (-2.25 + ((float32((0.1 - 1e10)) * (0.0 - (0.0 / 1.5))) - 0.0))) + (0.1 / 1.0)))))))) {
        v3 = v3
        fmt.Println(v3)
        var v34 int32 = 2147483647
        _ = v34
    } else {
        if v1 {
            fmt.Println(v1)
            v3 += 4294967295
            v1 = (255 < 2)
        } else {
            v1 = (("\n\t" + ("\n\t" + "foo bar")) <= "\n\t")
            v35 := GenS3{F0: 540, F1: "hello"}
            _ = v35
        }
    }
    v36 := GenS3{F0: 654, F1: "foo bar"}
    for fi37 := 0; fi37 < 6; fi37++ {
        fmt.Println(v1, v3)
        switch 32768 {
        case -2147483648:
            for fi38 := 0; fi38 < 1; fi38++ {
                fmt.Println(fi37, v1)
                switch fi37 {
                case 4294967295:
                    v3 -= 42
                }
                _ = fi38
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi37
    }
    if (((1e10 + -2.25) - 0.1) <= (0.1 + 1e10)) {
        var v39 uint8 = 1
        _ = v39
    } else {
        for fi40 := 0; fi40 < 14; fi40++ {
            for fi41 := 0; fi41 < 19; fi41++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                if (2147483647 > 118) {
                    if v1 {
                        v42 := v6[1]
                        v43 := GenS4{F0: 42}
                        var v44 uint64 = uint64(v3)
                        var v45 float64 = (0.0 / 3.14159)
                        _ = v42
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    v46 := GenS5{F0: 1, F1: 719}
                    _ = v46
                }
                var v47 uint = (v3 & (v3 | v3))
                _ = fi41
                _ = v47
            }
            _ = fi40
        }
        v48 := GenS6{F0: -1, F1: 0, F2: -233}
        _ = v48
    }
    switch "foo bar" {
    case "a":
        if (274 > v3) {
            _ = len(v4)
            if (42 == 382) {
                for fi49 := 0; fi49 < 3; fi49++ {
                    fmt.Println(fi49)
                    _ = fi49
                }
                if (!v1) {
                    var v50 string = ("foo bar" + "foo bar")
                    _ = v50
                } else {
                    fmt.Println(v3, v1)
                    if true {
                        fmt.Println(v3, v3)
                    } else {
                        if (-128 > 12) {
                            var v51 uint16 = 42
                            var v52 bool = (1 <= 381)
                            _ = v51
                            _ = v52
                        } else {
                            if (!false) {
                                v1 = v1
                            }
                        }
                        var v53 uint16 = 1
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v1, v53, v53)
                        case "\n\t":
                            v54 := GenS7{F0: (v53 >> (((v3 % (((v3 * 501)) | 1))) & 63))}
                            _ = v54
                        case "a":
                            v55 := GenS8{F0: 2147483648, F1: 274}
                            _ = v55
                        default:
                            for fi56 := 0; fi56 < 6; fi56++ {
                                for fi57 := 0; fi57 < 20; fi57++ {
                                    delete(v6, 65535)
                                    v3 -= v3
                                    _ = fi57
                                }
                                v1 = v1
                                v58 := GenS2{F0: 41}
                                _ = fi56
                                _ = v58
                            }
                        }
                        _ = v53
                    }
                    v1 = (1000 > 0)
                }
                v3 = (v3 / ((v3) | 1))
            }
        } else {
            fmt.Println(v3)
            var v59 uint = (^(v3))
            _ = v59
        }
    case "\n\t":
        v60 := GenS1{F0: 128, F1: uint64(v3)}
        _ = v60
    default:
        v61 := GenS9{F0: 32767, F1: ("" < "")}
        _ = v61
    }
    if true {
        var v62 int16 = 0
        v63 := GenS10{F0: 42, F1: 1, F2: 2147483648}
        _ = v62
        _ = v63
    }
    v3 += 173
    var v64 uint = v3
    for fi65 := 0; fi65 < 20; fi65++ {
        _ = len(v2)
        for fi66 := 0; fi66 < 18; fi66++ {
            if (!v1) {
                fmt.Println(fi65)
                switch -9223372036854775808 {
                case 157:
                    switch fi66 {
                    case 0:
                        var v67 float64 = 0.1
                        _ = v67
                    case -349:
                        switch fi66 {
                        case -1:
                            v68 := v5[fi66]
                            _ = v68
                        }
                    case -967:
                        var v69 float64 = 3.14159
                        if v1 {
                            if v1 {
                                fmt.Println(v64)
                                fmt.Println(fi65)
                            } else {
                                v69 += v69
                                if (240 >= 0) {
                                    fmt.Println(fi65, v1)
                                } else {
                                    v3 = (v3 - v3)
                                }
                            }
                            var v70 uint16 = 42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v69 -= 0.1
                            if (0.1 == 1.5) {
                                fmt.Println(v69, fi65)
                                v2 = append(v2, (((("" + ("foo bar" + "")) + "hello") + ((("foo bar" + "hello") + "foo bar") + ("\n\t" + (("" + "foo bar") + "")))) + "a"))
                                var v71 int = -2147483648
                                _ = v71
                            }
                            _ = v70
                        }
                        v64 = (^(v64))
                        _ = v69
                    }
                case -2147483648:
                    fmt.Println(v64)
                case 32768:
                    v4 = append(v4, (fi66 ^ (fi65 ^ (fi66 ^ fi65))))
                }
            } else {
                v72 := GenS11{F0: "foo bar"}
                v73 := GenS12{F0: 42, F1: ((!v1) && v1), F2: 0}
                fmt.Println(fi65, v3, v64)
                _ = v72
                _ = v73
            }
            switch fi66 {
            case -4:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -988:
                var v74 int = (fi66 | 0)
                _ = v74
            case -2147483648:
                if true {
                    v3 = 0
                    if v1 {
                        if (false && ((fi66 - fi65) < fi65)) {
                            delete(v5, fi65)
                            v75 := GenS2{F0: 2}
                            v76 := GenS10{F0: -1, F1: 128, F2: 0}
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v64, fi65)
                            var v77 uint = 255
                            _ = v77
                        }
                        fmt.Println(fi66, fi65)
                    } else {
                        v64 = 0
                        var v78 int16 = 0
                        if true {
                            _ = len(v6)
                        } else {
                            v64 -= (v3 & v64)
                            var v79 int8 = 0
                            if v1 {
                                var v80 uint16 = 65535
                                _ = v80
                            }
                            var v81 string = "foo bar"
                            switch (fi66 - fi65) {
                            case 32768:
                                v5[(fi65 >> (((v3 ^ (^(v3)))) & 63))] = 42
                            }
                            _ = v79
                            _ = v81
                        }
                        _ = v78
                    }
                    switch (fi66 + (fi66 ^ 4294967295)) {
                    case -774:
                        v82 := GenS13{F0: 4294967295, F1: "hello"}
                        _ = v82
                    case -399:
                        var v83 float32 = (0.0 * float32(float64(1.5)))
                        _ = v83
                    }
                }
            default:
                if true {
                    if v1 {
                        var v84 float32 = 1.5
                        if v1 {
                            v2 = append(v2, ((("\n\t" + "foo bar") + "a") + (("hello" + ("a" + "a")) + "\n\t")))
                            var v85 bool = v1
                            _ = v85
                        } else {
                            v1 = v1
                        }
                        var v86 float64 = 1e10
                        _ = v84
                        _ = v86
                    } else {
                        switch "a" {
                        case "foo bar":
                            v1 = true
                        default:
                            v1 = false
                        }
                    }
                }
            }
            _ = fi66
        }
        _ = fi65
    }
    v87 := v5[-1]
    v87 = v87
    delete(v6, 1)
    for fi88 := 0; fi88 < 4; fi88++ {
        if false {
            v4 = append(v4, fi88)
            var v89 uint8 = 42
            var v90 uint16 = 300
            _ = v89
            _ = v90
        } else {
            for fi91 := 0; fi91 < 11; fi91++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi91
            }
        }
        _ = fi88
    }
    if (false || false) {
        fmt.Println(v1, v3)
    }
    var v92 int32 = -865
    var v93 uint8 = 86
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v87:", v87)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v87
    _ = v92
    _ = v93
}
