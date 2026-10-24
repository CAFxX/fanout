package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 uint; F2 string }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint64; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int32; F1 int8 }
func genFn0(p0 int8, p1 uint) int16 {
    v6 := []int{}
    switch 1 {
    case 0:
        for fi7 := 0; fi7 < 14; fi7++ {
            if (42 != -197) {
                switch fi7 {
                case 4294967295:
                    if true {
                        if true {
                            v8 := GenS1{F0: "foo bar", F1: p0}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (1.5 < -2.25) {
                                fmt.Println(fi7, p1)
                                var v9 int16 = 1000
                                switch (fi7 >> (((p1 - p1)) & 63)) {
                                case -9223372036854775808:
                                    var v10 float64 = 0.0
                                    _ = v10
                                case 198:
                                    var v11 uint8 = 12
                                    _ = v11
                                case -1:
                                    p1 -= 1
                                default:
                                    v12 := GenS2{F0: (^(p1)), F1: p1, F2: "foo bar"}
                                    _ = v12
                                }
                                _ = v9
                            }
                            _ = v8
                        }
                    }
                case -1:
                    p1 -= p1
                case 9223372036854775807:
                    p1 = (^(p1))
                default:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        p1 -= p1
                        var v14 int16 = 0
                        _ = fi13
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 4; fi15++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v16 := GenS1{F0: ((("hello" + ("foo bar" + "\n\t")) + "a") + ("\n\t" + ((("\n\t" + ("" + "hello")) + "a") + ("hello" + (("foo bar" + ((((("\n\t" + "hello") + "\n\t") + ("hello" + "\n\t")) + "") + "foo bar")) + ("\n\t" + ("a" + ("\n\t" + "foo bar")))))))), F1: 1}
                    if false {
                        v17 := GenS1{F0: "foo bar", F1: p0}
                        v6 = append(v6, int(p1))
                        var v18 string = ("\n\t" + ("hello" + "foo bar"))
                        _ = v17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
            }
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int16 = 1000
                p0 = 42
                p0 = 79
                _ = fi19
                _ = v20
            }
            if len(v6) > 0 {
                v6[0] = -1
            }
            _ = fi7
        }
    default:
        v21 := GenS2{F0: (p1 | (p1 << ((p1) & 63))), F1: (p1 >> (((p1 - p1)) & 63)), F2: "\n\t"}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1() int {
    if false {
        var v22 uint = 682
        fmt.Println(v22)
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := []int{fi23, fi23, fi23}
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 uint16 = 32767
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    switch 25 {
    case -818:
        v27 := GenS3{F0: 127, F1: 9, F2: 124}
        _ = v27
    }
    if false {
        var v28 string = (("a" + ("" + "")) + "\n\t")
        _ = v28
    } else {
        var v29 uint = 65535
        switch -2147483648 {
        case 42:
            v29 += v29
        case 32767:
            if false {
                var v30 uint64 = 2
                v31 := GenS4{F0: 100, F1: v30, F2: -206}
                _ = v30
                _ = v31
            }
        case -580:
            v32 := []int{}
            _ = v32
        }
        v33 := []int{-263, 37, 127, 9223372036854775807}
        _ = v29
        _ = v33
    }
    if (42 != 32767) {
        var v34 string = "foo bar"
        var v35 int64 = 947
        _ = v34
        _ = v35
    }
    return 42
}

func main() {
    var v1 string = (("foo bar" + "hello") + "hello")
    var v2 int16 = 42
    var v3 string = ("hello" + (("\n\t" + v1) + v1))
    v4 := []int{}
    var v5 uint = 205
    switch -1 {
    case 1:
        if len(v4) > 0 {
            v4[0] = genFn1()
        }
    case -2147483648:
        var v36 uint16 = 104
        _ = v36
    case 32768:
        v2 = (v2 + v2)
    }
    if true {
        v37 := GenS5{F0: 1}
        _ = len(v4)
        _ = v37
    } else {
        if (!true) {
            var v38 int64 = -1
            _ = v38
        }
        switch (((v1 + ((v1 + v3) + v1)) + v3) + v1) {
        case "":
            var v39 bool = false
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            if (1 > -9223372036854775808) {
                v41 := GenS6{F0: 126, F1: 4294967295, F2: fi40}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v42 int16 = v2
                _ = v41
                _ = v42
            } else {
                v43 := GenS2{F0: (v5 << (((v5 | (v5 % ((42) | 1)))) & 63)), F1: v5, F2: "a"}
                v44 := GenS3{F0: -110, F1: 42, F2: 255}
                _ = v43
                _ = v44
            }
            _ = fi40
        }
    }
    _ = len(v4)
    if ((!true) && true) {
        var v45 int8 = 0
        switch -1 {
        case -597:
            v1 = "hello"
        }
        _ = v45
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        switch fi46 {
        case -1:
            for fi47 := 0; fi47 < 13; fi47++ {
                var v48 bool = true
                var v49 int8 = 26
                _ = fi47
                _ = v48
                _ = v49
            }
        case 4294967295:
            v50 := GenS5{F0: 0}
            _ = v50
        case 633:
            v51 := GenS2{F0: 263, F1: 18446744073709551615, F2: v1}
            _ = v51
        }
        v52 := GenS6{F0: 128, F1: 255, F2: (fi46 << (((^(v5))) & 63))}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi46
        _ = v52
    }
    if (-49 > 1) {
        _ = len(v4)
        v53 := GenS4{F0: 100, F1: 1, F2: 0}
        _ = v53
    } else {
        v1 = v1
        if ((1e10 < 1.5) && true) {
            v4 = append(v4, 32768)
            var v54 float64 = -2.25
            v1 = (v1 + (v1 + "a"))
            _ = v54
        }
        switch v1 {
        case "\n\t":
            if ((0 == -610) || (169 == 4294967295)) {
                fmt.Println(v3)
            }
        case "foo bar":
            for fi55 := 0; fi55 < 19; fi55++ {
                v4 = append(v4, fi55)
                if (!true) {
                    for fi56 := 0; fi56 < 14; fi56++ {
                        var v57 uint32 = 0
                        v3 = v3
                        _ = fi56
                        _ = v57
                    }
                } else {
                    v2 += -32768
                    v3 = v1
                    v5 += 157
                }
                _ = fi55
            }
        default:
            switch v1 {
            case "hello":
                _ = len(v4)
            }
        }
    }
    var v58 uint16 = 1
    _ = len(v4)
    fmt.Println(v3, v1, v3)
    var v59 uint = (^(v5))
    v59 += (v5 - (v59 * v5))
    switch 127 {
    case -558:
        fmt.Println(v5, v2, v59)
    default:
        var v60 int64 = 9223372036854775807
        _ = v60
    }
    switch 127 {
    case 4294967295:
        v61 := GenS3{F0: -1, F1: 1, F2: v58}
        _ = v61
    case -238:
        var v62 uint = v5
        _ = v62
    default:
        if (18446744073709551615 < 42) {
            if ((!false) || (-1 == 42)) {
                if true {
                    v63 := GenS7{F0: (v5 % ((v5) | 1))}
                    _ = v63
                } else {
                    v58 = 32767
                    var v64 uint8 = 2
                    for fi65 := 0; fi65 < 10; fi65++ {
                        var v66 uint64 = 42
                        for fi67 := 0; fi67 < 12; fi67++ {
                            switch v3 {
                            case "a":
                                if true {
                                    v58 += 0
                                    v68 := GenS7{F0: 255}
                                    v4 = append(v4, fi65)
                                    _ = v68
                                } else {
                                    var v69 int64 = -1
                                    v70 := GenS8{F0: 0, F1: v64, F2: genFn0(-38, (v59 + 319))}
                                    v69 += 307
                                    _ = v69
                                    _ = v70
                                }
                            default:
                                if false {
                                    if false {
                                        var v71 uint8 = (^(v64))
                                        _ = v71
                                    } else {
                                        var v72 uint = 824
                                        v64 -= uint8(v72)
                                        switch genFn1() {
                                        case -305:
                                            fmt.Println(v64, v59)
                                        case 1:
                                            if true {
                                                v73 := GenS9{F0: (v1 + ("foo bar" + v3))}
                                                fmt.Println(v59, fi67)
                                                v74 := GenS10{F0: v1}
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                        v5 = 255
                                        _ = v72
                                    }
                                }
                            }
                            var v75 int8 = 59
                            v64 -= v64
                            var v76 int = fi67
                            if len(v4) > 0 {
                                v4[0] = v76
                            }
                            _ = fi67
                            _ = v75
                            _ = v76
                        }
                        _ = fi65
                        _ = v66
                    }
                    _ = v64
                }
            }
            if (true && (!false)) {
                switch v3 {
                case "\n\t":
                    for fi77 := 0; fi77 < 11; fi77++ {
                        if true {
                            for fi78 := 0; fi78 < 12; fi78++ {
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v59 = (v5 >> ((v59) & 63))
                                }
                                v79 := GenS10{F0: v3}
                                v80 := GenS6{F0: uint8(v58), F1: v5, F2: fi77}
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                            var v81 uint32 = 0
                            for fi82 := 0; fi82 < 1; fi82++ {
                                var v83 uint = 4294967295
                                _ = fi82
                                _ = v83
                            }
                            _ = v81
                        } else {
                            v4 = append(v4, genFn1())
                        }
                        v58 -= v58
                        for fi84 := 0; fi84 < 6; fi84++ {
                            switch (fi77 + fi77) {
                            case 4294967295:
                                v85 := GenS3{F0: 127, F1: -128, F2: v58}
                                _ = v85
                            case -729:
                                var v86 float64 = 1e10
                                _ = v86
                            default:
                                var v87 bool = false
                                _ = v87
                            }
                            v4 = append(v4, 0)
                            v3 = ""
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi84
                        }
                        _ = fi77
                    }
                case "":
                    var v88 int = -1
                    _ = v88
                }
            } else {
                switch v1 {
                case "hello":
                    v4 = append(v4, 753)
                case "a":
                    fmt.Println(v5, v3)
                case "\n\t":
                    v89 := GenS6{F0: 83, F1: (v5 / (((^(v5))) | 1)), F2: genFn1()}
                    _ = v89
                default:
                    v59 -= (v59 / ((v5) | 1))
                }
                if true {
                    var v90 uint = (v5 << ((v5) & 63))
                    var v91 int16 = v2
                    var v92 int16 = genFn0(1, (v5 ^ (^(v5))))
                    _ = v90
                    _ = v91
                    _ = v92
                } else {
                    v93 := GenS11{F0: -418, F1: 112}
                    fmt.Println(v1, v1)
                    _ = v93
                }
                v58 += v58
            }
        }
    }
    var v94 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v94
}
