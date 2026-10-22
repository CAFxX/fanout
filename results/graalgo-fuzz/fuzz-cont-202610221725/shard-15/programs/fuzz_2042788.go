package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int64; F1 uint16; F2 string }
type GenS8 struct { F0 int32; F1 uint64; F2 bool }
type GenS9 struct { F0 string; F1 int8; F2 uint32 }
type GenS10 struct { F0 int16; F1 uint8; F2 string }
type GenS11 struct { F0 uint; F1 int64; F2 uint8 }
func genFn0(p0 uint16) bool {
    v6 := []int{-2147483648, -9223372036854775808, 0}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v6
    return true
}
func genFn1(p0 bool) int32 {
    fmt.Println(p0)
    p0 = true
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 0.0
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]string)
    v5 := []uint32{277, 2147483648}
    switch 56 {
    case 4294967295:
        if (238 == 164) {
            switch "\n\t" {
            case "\n\t":
                v1 += v1
            case "a":
                v7 := GenS1{F0: (42 <= 423), F1: 128}
                _ = v7
            case "hello":
                v3 = 42
            }
            fmt.Println(v1, v3)
            v5 = append(v5, 2147483647)
        } else {
            v3 += (v3 & (v3 | v3))
            v8 := GenS1{F0: ((!(32767 > 42)) || true), F1: -1}
            for fi9 := 0; fi9 < 5; fi9++ {
                if true {
                    var v10 int8 = 42
                    _ = v10
                } else {
                    var v11 int8 = 0
                    _ = v11
                }
                v3 = (v3 * (v3 ^ v3))
                if len(v5) > 0 {
                    v5[0] = uint32(v1)
                }
                _ = fi9
            }
            _ = v8
        }
    case 9223372036854775807:
        var v12 uint32 = 0
        _ = v12
    }
    if (3.14159 > 0.0) {
        for fi13 := 0; fi13 < 15; fi13++ {
            v2 -= float32(1e10)
            switch fi13 {
            case 32768:
                if (!((v1 ^ v1) <= v1)) {
                    var v14 string = "foo bar"
                    _ = v14
                }
            case 128:
                _ = len(v4)
            default:
                var v15 string = "\n\t"
                _ = v15
            }
            _ = fi13
        }
    } else {
        if ("foo bar" >= ("" + (("" + ("hello" + ("\n\t" + ("a" + (("a" + "a") + (("\n\t" + "hello") + ("foo bar" + ("hello" + "\n\t")))))))) + ("a" + ("foo bar" + "a"))))) {
            for fi16 := 0; fi16 < 20; fi16++ {
                if false {
                    if (!false) {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            if true {
                                _ = len(v4)
                                if (!(!(2 <= 2))) {
                                    v18 := GenS1{F0: true, F1: fi16}
                                    v5 = append(v5, 1)
                                    v19 := v4["foo bar"]
                                    _ = v18
                                    _ = v19
                                } else {
                                    v2 = v2
                                    v20 := GenS2{F0: -2147483648, F1: fi16}
                                    var v21 uint32 = 0
                                    if (732 == 1) {
                                        delete(v4, "")
                                        switch (("foo bar" + ("a" + "foo bar")) + "") {
                                        case "hello":
                                            if false {
                                                var v22 float64 = 0.0
                                                switch fi17 {
                                                case -931:
                                                    v23 := GenS3{F0: v21, F1: 18446744073709551615, F2: (v21 | (^(v21)))}
                                                    _ = v23
                                                case 0:
                                                    v24 := GenS2{F0: 0, F1: (fi16 & (fi17 ^ fi17))}
                                                    _ = v24
                                                default:
                                                    var v25 int32 = genFn1(true)
                                                    _ = v25
                                                }
                                                _ = v22
                                            } else {
                                                if true {
                                                    if true {
                                                        v26 := GenS1{F0: false, F1: -523}
                                                        v27 := GenS4{F0: "a"}
                                                        v28 := GenS2{F0: 42, F1: fi17}
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                } else {
                                                    switch int(v1) {
                                                    case 2147483647:
                                                        _ = len(v5)
                                                    default:
                                                        v3 += (v3 >> ((0) & 63))
                                                    }
                                                    v5 = append(v5, v21)
                                                }
                                            }
                                        }
                                        v3 -= v3
                                    } else {
                                        var v29 uint = 4294967295
                                        v5 = append(v5, (^(v21)))
                                        _ = v29
                                    }
                                    _ = v20
                                    _ = v21
                                }
                            }
                            _ = fi17
                        }
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v31 := GenS3{F0: 2, F1: v1, F2: 1}
                            v32 := GenS3{F0: 727, F1: 18446744073709551615, F2: 2147483647}
                            v3 -= -1
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 13; fi33++ {
                            fmt.Println(fi33)
                            v34 := GenS5{F0: 42}
                            v35 := GenS1{F0: true, F1: 32768}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    } else {
                        fmt.Println(v1, v2, v1)
                    }
                    v36 := GenS2{F0: 49, F1: 32768}
                    switch (fi16 >> ((4294967295) & 63)) {
                    case -9223372036854775808:
                        v1 = (v1 & (v1 << ((4294967295) & 63)))
                    case -2147483648:
                        var v37 int8 = -1
                        _ = v37
                    }
                    _ = v36
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi16
            }
            var v38 int = 9223372036854775807
            v1 = 0
            _ = v38
        } else {
            v5 = append(v5, 269)
            for fi39 := 0; fi39 < 10; fi39++ {
                for fi40 := 0; fi40 < 5; fi40++ {
                    v2 = 3.14159
                    _ = fi40
                }
                switch "" {
                case "a":
                    fmt.Println(v1)
                case "":
                    var v41 bool = (((!true) || false) && false)
                    _ = v41
                case "foo bar":
                    var v42 int32 = 572
                    _ = v42
                }
                v43 := GenS1{F0: (-763 > -2147483648), F1: int(v1)}
                _ = fi39
                _ = v43
            }
        }
        switch "" {
        case "\n\t":
            switch 9223372036854775807 {
            case 4294967295:
                v3 += -537
            default:
                for fi44 := 0; fi44 < 16; fi44++ {
                    if true {
                        v2 -= (0.0 + (v2 * (3.14159 / float32(1e10))))
                    }
                    switch "foo bar" {
                    case "a":
                        for fi45 := 0; fi45 < 11; fi45++ {
                            if true {
                                if (false || true) {
                                    fmt.Println(fi44)
                                    v3 -= 258
                                } else {
                                    v46 := GenS6{F0: -1, F1: -495}
                                    var v47 uint16 = 113
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                var v48 int32 = 2147483647
                                _ = len(v4)
                                _ = v48
                            }
                            _ = fi45
                        }
                    default:
                        if true {
                            v3 = v3
                            if (-1 != genFn1((uint16(v1) < 254))) {
                                _ = len(v4)
                                for fi49 := 0; fi49 < 16; fi49++ {
                                    fmt.Println(v2)
                                    if (false && true) {
                                        v50 := GenS6{F0: int8(fi49), F1: 42}
                                        if (!false) {
                                            var v51 float64 = 0.0
                                            _ = v51
                                        }
                                        _ = v50
                                    } else {
                                        if (!(27 > 42)) {
                                            switch "" {
                                            case "foo bar":
                                                switch fi49 {
                                                case -9223372036854775808:
                                                    v52 := GenS7{F0: v3, F1: 967, F2: "foo bar"}
                                                    v3 += (v3 & (v3 / ((789) | 1)))
                                                    _ = v52
                                                default:
                                                    fmt.Println(v1, fi44)
                                                }
                                            default:
                                                fmt.Println(v3)
                                            }
                                            v1 += (v1 % (((^(v1))) | 1))
                                        }
                                        fmt.Println(v1, v3)
                                        if true {
                                            var v53 float64 = (1.5 + (float64(v2) - 1.5))
                                            var v54 uint32 = 0
                                            _ = v53
                                            _ = v54
                                        } else {
                                            v3 += v3
                                        }
                                    }
                                    if (!true) {
                                        fmt.Println(v3, v1)
                                        v3 = v3
                                        var v55 string = ("a" + "a")
                                        fmt.Println(fi49, v55)
                                        _ = v55
                                    } else {
                                        var v56 int = (fi49 & (fi49 / ((1) | 1)))
                                        var v57 uint8 = 255
                                        if true {
                                            switch "hello" {
                                            case "hello":
                                                var v58 uint16 = 42
                                                _ = len(v5)
                                                _ = v58
                                            case "":
                                                switch (^(fi44)) {
                                                case -1:
                                                    var v59 float64 = 0.1
                                                    _ = v59
                                                case 2147483647:
                                                    if true {
                                                        var v60 int = (fi49 | (v56 + (fi44 / (((v56 / (((v56 ^ (fi49 * v56))) | 1))) | 1))))
                                                        _ = v60
                                                    } else {
                                                        v2 = float32(((-2.25 - (1e10 + (-2.25 - 0.0))) / 1.5))
                                                        v56 = fi49
                                                        fmt.Println(v1)
                                                    }
                                                case 42:
                                                    v61 := GenS8{F0: -1, F1: 1, F2: (935 <= 0)}
                                                    _ = v61
                                                }
                                            case "\n\t":
                                                v3 -= (v3 + v3)
                                            }
                                            var v62 int64 = (v3 >> ((4294967295) & 63))
                                            _ = v62
                                        }
                                        _ = v56
                                        _ = v57
                                    }
                                    v63 := GenS9{F0: (("hello" + "a") + "a"), F1: -6, F2: 65535}
                                    v64 := GenS10{F0: -1, F1: 241, F2: ("" + "hello")}
                                    _ = fi49
                                    _ = v63
                                    _ = v64
                                }
                            }
                        }
                    }
                    _ = fi44
                }
            }
        }
    }
    v65 := GenS7{F0: (v3 ^ -781), F1: 42, F2: ((("a" + ("a" + "a")) + "foo bar") + ("hello" + (("a" + ("a" + "hello")) + "")))}
    v5 = append(v5, 42)
    v2 += -2.25
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    v4[""] = "a"
    var v66 float64 = float64((float32((3.14159 - 1e10)) + v2))
    switch "hello" {
    case "foo bar":
        v2 += v2
    case "\n\t":
        v67 := GenS2{F0: -2147483648, F1: -2147483648}
        _ = v67
    case "hello":
        var v68 uint64 = (v1 & 2)
        _ = v68
    default:
        for fi69 := 0; fi69 < 10; fi69++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (!false) {
                v70 := GenS6{F0: 42, F1: 42}
                v2 = (v2 - v2)
                var v71 string = ("hello" + "foo bar")
                _ = v70
                _ = v71
            }
            _ = fi69
        }
    }
    for fi72 := 0; fi72 < 11; fi72++ {
        for fi73 := 0; fi73 < 2; fi73++ {
            if (!true) {
                v2 -= 3.14159
                var v74 float64 = (float64(v2) - -2.25)
                v75 := GenS7{F0: v3, F1: 0, F2: "hello"}
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                if ((!false) && (!false)) {
                    fmt.Println(v74)
                    fmt.Println(v74)
                    if (!(true && (fi73 == fi72))) {
                        v3 = -868
                        if (true || true) {
                            var v76 uint64 = (^(v1))
                            fmt.Println(fi72, v76, fi73)
                            _ = v76
                        }
                        var v77 float32 = v2
                        _ = v77
                    }
                }
                _ = v74
                _ = v75
            } else {
                var v78 uint8 = 191
                delete(v4, "hello")
                var v79 float64 = v66
                v80 := GenS8{F0: 1, F1: 261, F2: true}
                switch (^(fi73)) {
                case -1:
                    v81 := GenS11{F0: 324, F1: (v3 << ((0) & 63)), F2: v78}
                    _ = v81
                case -660:
                    if true {
                        fmt.Println(v3, v1)
                        v2 += float32((v79 / (float64(v2) / float64((0.0 - float32((float64((v2 + (1e10 + v2))) - float64(v2))))))))
                        var v82 float32 = (float32(v79) + v2)
                        _ = v82
                    }
                case 0:
                    v3 -= v3
                }
                _ = v78
                _ = v79
                _ = v80
            }
            v66 = 1.5
            if (!(!true)) {
                v3 += (-(v3))
            } else {
                v3 = v3
            }
            _ = fi73
        }
        for fi83 := 0; fi83 < 6; fi83++ {
            fmt.Println(v2, fi72, fi72)
            if true {
                if (127 != -1) {
                    var v84 uint64 = (v1 % ((v1) | 1))
                    switch fi72 {
                    case 128:
                        var v85 int32 = 0
                        _ = v85
                    case 842:
                        switch 1 {
                        case 1:
                            var v86 uint = 65535
                            _ = v86
                        case 554:
                            v5 = append(v5, 697)
                        case -306:
                            v87 := GenS2{F0: int32(fi72), F1: fi72}
                            _ = v87
                        default:
                            var v88 string = ""
                            _ = v88
                        }
                    }
                    _ = v84
                }
                v3 = (v3 << ((42) & 63))
                v1 = (^(v1))
                v2 = v2
                v2 -= float32(((0.0 + 1.5) + (0.0 - (float64((1e10 / (v2 + float32(v66)))) + (v66 / v66)))))
            }
            var v89 uint16 = 34
            _ = fi83
            _ = v89
        }
        for fi90 := 0; fi90 < 19; fi90++ {
            var v91 int = fi72
            _ = fi90
            _ = v91
        }
        _ = fi72
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v92 uint = 4294967295
    var v93 int = 982
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v92
    _ = v93
}
