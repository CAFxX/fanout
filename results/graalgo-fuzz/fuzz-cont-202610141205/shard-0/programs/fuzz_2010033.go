package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint; F1 uint64; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool; F1 bool }
type GenS14 struct { F0 string; F1 int64; F2 bool }
func genFn0(p0 uint16, p1 int32) uint16 {
    v5 := []int{1, 32767, 128, -9223372036854775808}
    if (106 != 1) {
        switch 42 {
        case 42:
            v6 := GenS1{F0: 270}
            _ = v6
        case 32767:
            var v7 int64 = 9223372036854775807
            _ = v7
        case 241:
            var v8 uint8 = 140
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
        p0 = 0
    }
    if (!false) {
        if (false && (-915 != -67)) {
            var v9 uint = 4294967295
            _ = v9
        } else {
            v5 = append(v5, -872)
        }
        switch "a" {
        case "a":
            if true {
                _ = len(v5)
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(p1, p1, p1)
                    p1 = int32(p0)
                    _ = fi10
                }
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 float64 = -2.25
                    p0 = 32767
                    _ = len(v5)
                    _ = fi11
                    _ = v12
                }
            }
        case "":
            v13 := GenS2{F0: true, F1: 2}
            _ = v13
        case "foo bar":
            if (-3 != -1) {
                var v14 uint16 = (p0 << ((649) & 63))
                _ = v14
            } else {
                p1 -= -1
                if len(v5) > 0 {
                    v5[0] = -608
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int16 = -1
                    if (p1 > p1) {
                        v5 = append(v5, (fi15 - (-(fi15))))
                        v16 = (v16 % ((32767) | 1))
                        fmt.Println(v16)
                    }
                    for fi17 := 0; fi17 < 12; fi17++ {
                        p0 -= 1
                        _ = fi17
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    var v18 int = 0
    _ = p0
    _ = p1
    _ = v18
    _ = v5
    return (p0 | (p0 - 255))
}
func genFn1(p0 bool, p1 uint8, p2 uint32) uint8 {
    if false {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 float64 = float64((-2.25 - 3.14159))
            p1 += (^(p1))
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := GenS3{F0: 18446744073709551615, F1: 127}
                v23 := []int{797, (fi19 * (fi21 * fi21)), (^(fi21)), (fi19 * -9223372036854775808)}
                v24 := GenS2{F0: p0, F1: 73}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v20
        }
    } else {
        var v25 uint16 = genFn0(uint16(p1), 1)
        for fi26 := 0; fi26 < 7; fi26++ {
            if p0 {
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 uint16 = genFn0((v25 & v25), -1)
                    if p0 {
                        var v29 uint64 = 333
                        if (!(true && false)) {
                            p0 = false
                            if true {
                                var v30 int8 = -85
                                _ = v30
                            } else {
                                if ("foo bar" != ("a" + "hello")) {
                                    if (!(-32768 > 867)) {
                                        v31 := GenS4{F0: 1000}
                                        var v32 int = (fi26 / ((fi26) | 1))
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 int = 42
                                        fmt.Println(fi26, v29)
                                        p2 += (p2 - (p2 / ((p2) | 1)))
                                        _ = v33
                                    }
                                }
                                v34 := GenS1{F0: 824}
                                v35 := []int{(fi26 & fi26), fi26, (fi27 % ((fi26) | 1)), fi26}
                                v36 := GenS5{F0: 42, F1: "a"}
                                if (!((3.14159 == 0.1) && true)) {
                                    _ = len(v35)
                                    v28 += (v25 / (((v25 + (v28 & (v25 << ((1) & 63))))) | 1))
                                } else {
                                    v28 = 50
                                    switch 32768 {
                                    case 55:
                                        fmt.Println(p2, v29, v28)
                                    case -1:
                                        if (0 == 867) {
                                            if true {
                                                v25 = v28
                                                var v37 int32 = 2147483647
                                                _ = v37
                                            } else {
                                                p0 = ((fi26 << ((1) & 63)) <= (-(fi27)))
                                                v38 := GenS6{F0: 65535, F1: v29, F2: ("\n\t" + "hello")}
                                                v39 := GenS7{F0: 32767}
                                                _ = v38
                                                _ = v39
                                            }
                                            fmt.Println(fi27, v29, p2)
                                            if ((0.0 * -2.25) < 0.0) {
                                                if (false && (1.5 == (1e10 * 1.5))) {
                                                    if (p1 == 128) {
                                                        if len(v35) > 0 {
                                                            v35[0] = (fi27 / (((fi26 % ((fi27) | 1))) | 1))
                                                        }
                                                        var v40 string = ((("a" + (((("\n\t" + "foo bar") + "") + ((("" + ("hello" + "hello")) + "a") + "foo bar")) + "")) + "a") + ("\n\t" + "\n\t"))
                                                        _ = v40
                                                    } else {
                                                        var v41 uint64 = v29
                                                        v28 -= (^(v28))
                                                        v42 := GenS2{F0: (false || true), F1: p1}
                                                        v43 := GenS8{F0: -867, F1: 42}
                                                        v25 -= 0
                                                        fmt.Println(v41)
                                                        if ((p2 == 1) && p0) {
                                                            v44 := GenS7{F0: 0}
                                                            _ = len(v35)
                                                            var v45 float64 = -2.25
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    if len(v35) > 0 {
                                                        v35[0] = (fi27 >> ((756) & 63))
                                                    }
                                                } else {
                                                    v46 := GenS4{F0: 184}
                                                    _ = v46
                                                }
                                            }
                                        }
                                        v28 += uint16(p2)
                                        if p0 {
                                            switch fi26 {
                                            case 9223372036854775807:
                                                v47 := GenS6{F0: 65535, F1: (v29 / ((uint64(v28)) | 1)), F2: ((("\n\t" + "a") + "\n\t") + (("" + "\n\t") + "foo bar"))}
                                                _ = v47
                                            default:
                                                fmt.Println(v25)
                                            }
                                            v48 := GenS3{F0: (v29 >> ((360) & 63)), F1: p1}
                                            v35 = append(v35, fi27)
                                            _ = v48
                                        } else {
                                            switch 343 {
                                            case 32767:
                                                v35 = append(v35, fi26)
                                            case 1:
                                                if true {
                                                    if p0 {
                                                        v49 := GenS8{F0: -1, F1: -32768}
                                                        _ = len(v35)
                                                        _ = v49
                                                    } else {
                                                        _ = len(v35)
                                                        if len(v35) > 0 {
                                                            v35[0] = fi27
                                                        }
                                                        var v50 uint32 = (p2 % ((p2) | 1))
                                                        _ = v50
                                                    }
                                                } else {
                                                    switch (^(fi26)) {
                                                    case 4294967295:
                                                        if len(v35) >= 2 {
                                                            v35 = v35[1:len(v35)]
                                                        }
                                                    case 842:
                                                        v35 = append(v35, (fi27 % ((fi26) | 1)))
                                                    default:
                                                        var v51 float64 = ((0.1 - (1e10 + (-2.25 / -2.25))) * 0.0)
                                                        _ = v51
                                                    }
                                                }
                                            default:
                                                var v52 uint8 = 42
                                                _ = v52
                                            }
                                            if len(v35) > 0 {
                                                v35[0] = (fi27 ^ (fi26 & 1))
                                            }
                                            var v53 string = "foo bar"
                                            _ = v53
                                        }
                                        p2 += 4294967295
                                        var v54 int64 = 9223372036854775807
                                        _ = v54
                                    default:
                                        var v55 int32 = 1
                                        _ = v55
                                    }
                                }
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        }
                        v56 := []int{(fi26 | (fi26 & 959)), (fi26 - (fi27 / ((fi26) | 1))), fi27, fi26}
                        switch (fi27 & (-(fi26))) {
                        case -861:
                            _ = len(v56)
                        case 32767:
                            switch (("foo bar" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t") {
                            case "":
                                fmt.Println(v28, v29, fi26)
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    switch "a" {
                                    case "":
                                        var v57 uint16 = v28
                                        var v58 uint64 = (v29 >> ((1) & 63))
                                        fmt.Println(v29, fi27, p2)
                                        _ = v57
                                        _ = v58
                                    case "a":
                                        var v59 uint64 = v29
                                        _ = v59
                                    default:
                                        v56 = append(v56, fi26)
                                    }
                                default:
                                    fmt.Println(p1)
                                }
                            case "\n\t":
                                v28 += (v28 + v25)
                            }
                            fmt.Println(fi27)
                        }
                        v56 = append(v56, (-(fi26)))
                        _ = v29
                        _ = v56
                    } else {
                        switch 0 {
                        case -396:
                            var v60 bool = (p0 || p0)
                            _ = v60
                        case -1:
                            var v61 int16 = 1
                            _ = v61
                        }
                        var v62 uint64 = 4294967295
                        _ = v62
                    }
                    _ = fi27
                    _ = v28
                }
                v63 := []int{(fi26 % ((128) | 1)), fi26, (fi26 >> ((255) & 63)), fi26}
                if false {
                    v64 := GenS4{F0: 763}
                    _ = v64
                }
                _ = v63
            } else {
                v65 := GenS9{F0: 18446744073709551615}
                fmt.Println(p1, p2)
                p0 = (!true)
                _ = v65
            }
            _ = fi26
        }
        p1 = p1
        _ = v25
    }
    p2 -= uint32(p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int32 = 279
    v2 := make(map[uint32]uint32)
    var v3 int32 = (v1 << ((864) & 63))
    var v4 int16 = 42
    v4 = 1000
    var v66 int32 = (v1 + 42)
    v67 := []int{911}
    if true {
        v68 := GenS10{F0: 1}
        v3 += (v1 & v66)
        _ = v68
    }
    v1 += (v66 * (v3 >> ((18446744073709551615) & 63)))
    if (18 < 1) {
        for fi69 := 0; fi69 < 17; fi69++ {
            switch fi69 {
            case -330:
                delete(v2, 2147483647)
            default:
                for fi70 := 0; fi70 < 8; fi70++ {
                    var v71 int64 = 0
                    var v72 int16 = v4
                    v4 = 32767
                    _ = fi70
                    _ = v71
                    _ = v72
                }
            }
            switch (("\n\t" + "hello") + "hello") {
            case "\n\t":
                _ = len(v2)
            case "a":
                switch (fi69 << ((2) & 63)) {
                case -2147483648:
                    var v73 int16 = v4
                    _ = v73
                case 127:
                    if (((!(!true)) || false) || (65535 != 0)) {
                        if true {
                            v4 += (v4 >> ((2) & 63))
                            v74 := GenS11{F0: 0}
                            _ = v74
                        } else {
                            fmt.Println(v66, v4)
                        }
                    } else {
                        fmt.Println(fi69, v1, v4)
                        var v75 string = ("\n\t" + "a")
                        _ = v75
                    }
                case 42:
                    for fi76 := 0; fi76 < 3; fi76++ {
                        var v77 uint64 = 2
                        _ = fi76
                        _ = v77
                    }
                default:
                    if (true && false) {
                        if len(v67) >= 2 {
                            v67 = v67[1:len(v67)]
                        }
                        for fi78 := 0; fi78 < 9; fi78++ {
                            v79 := GenS12{F0: 65535}
                            _ = fi78
                            _ = v79
                        }
                        if false {
                            var v80 float64 = -2.25
                            _ = v80
                        }
                    } else {
                        v81 := GenS12{F0: 1}
                        _ = v81
                    }
                }
            }
            _ = fi69
        }
        if false {
            fmt.Println(v4)
            for fi82 := 0; fi82 < 8; fi82++ {
                var v83 uint16 = 1
                v84 := GenS13{F0: false, F1: false}
                _ = fi82
                _ = v83
                _ = v84
            }
        }
    }
    switch "foo bar" {
    case "":
        if (("" + "hello") >= ("" + (("\n\t" + ("" + "a")) + "foo bar"))) {
            fmt.Println(v1)
        } else {
            v67 = append(v67, 128)
            _ = len(v2)
        }
    case "\n\t":
        if (54 < 255) {
            fmt.Println(v4)
        }
    default:
        var v85 int32 = (v3 - 855)
        _ = v85
    }
    v86 := GenS2{F0: true, F1: 1}
    var v87 int32 = (v66 >> ((0) & 63))
    v88 := GenS14{F0: ("hello" + ("foo bar" + "foo bar")), F1: -237, F2: (!(255 <= 65535))}
    if false {
        for fi89 := 0; fi89 < 10; fi89++ {
            fmt.Println(v1, v66)
            var v90 bool = true
            _ = fi89
            _ = v90
        }
    }
    v91 := GenS9{F0: 42}
    var v92 bool = (2 == 254)
    if len(v67) > 0 {
        v67[0] = -1
    }
    var v93 uint32 = 110
    v94 := GenS14{F0: "\n\t", F1: 1, F2: (1.5 > 0.0)}
    v95 := GenS8{F0: v4, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v87:", v87)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v86
    _ = v87
    _ = v88
    _ = v91
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
