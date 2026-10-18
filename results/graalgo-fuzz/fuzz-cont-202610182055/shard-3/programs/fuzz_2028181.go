package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint8; F1 uint64; F2 string }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint64; F2 uint }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 int64 }
func genFn0(p0 int16, p1 int) int16 {
    v7 := GenS1{F0: int64(p1)}
    v8 := []int{}
    p0 -= 0
    v9 := GenS1{F0: 42}
    if (((p1 / ((p1) | 1)) <= (p1 | -1)) && false) {
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint8 = 127
    v5 := []int{384}
    var v6 bool = true
    _ = len(v5)
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := v2[fi10]
        v12 := GenS1{F0: -9223372036854775808}
        switch (fi10 << ((1) & 63)) {
        case 4294967295:
            v13 := GenS1{F0: -9223372036854775808}
            _ = v13
        case -2147483648:
            fmt.Println(v4)
        case -1:
            v6 = true
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v6)
            _ = fi15
        }
        v4 = 42
        var v16 float64 = 1.5
        _ = fi14
        _ = v16
    }
    if v6 {
        var v17 uint64 = 18446744073709551615
        v6 = v6
        var v18 bool = false
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v6 = v6
        _ = fi19
    }
    v6 = v6
    v20 := GenS1{F0: 1}
    v21 := GenS2{F0: -32768, F1: 1, F2: -2147483648}
    v4 += (v4 << ((732) & 63))
    v1[("hello" + ("\n\t" + ("foo bar" + "\n\t")))] = ("\n\t" + ("hello" + "foo bar"))
    v22 := v2[32767]
    v4 -= 231
    for fi23 := 0; fi23 < 8; fi23++ {
        fmt.Println(v4, v6, v4)
        _ = fi23
    }
    if ((!v6) || v6) {
        var v24 float32 = ((3.14159 + 3.14159) + 0.0)
        _ = v24
    } else {
        for fi25 := 0; fi25 < 13; fi25++ {
            if false {
                switch "" {
                case "foo bar":
                    v6 = ((2147483647 >= 1) || (2147483647 < 508))
                case "hello":
                    v22 = 0
                case "":
                    if (636 != 0) {
                        for fi26 := 0; fi26 < 10; fi26++ {
                            var v27 int8 = -107
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(fi25, v4, v6)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi28
                        }
                    } else {
                        v22 += (v22 * -774)
                    }
                default:
                    if false {
                        v29 := GenS2{F0: 32767, F1: 100, F2: 0}
                        _ = v29
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v30 int16 = -713
                        var v31 uint64 = 894
                        _ = v30
                        _ = v31
                    }
                }
                switch "" {
                case "foo bar":
                    var v32 int8 = -54
                    _ = v32
                default:
                    v6 = v6
                }
                for fi33 := 0; fi33 < 4; fi33++ {
                    if (!(((1e10 - 1.5) / 1.0) >= 3.14159)) {
                        _ = len(v5)
                        if (2147483647 < int32(v22)) {
                            if v6 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v6 = (1 >= 0)
                            }
                        } else {
                            switch "hello" {
                            case "hello":
                                v5 = append(v5, v22)
                            case "":
                                if (100 > 1) {
                                    fmt.Println(v22, fi25, v22)
                                } else {
                                    v6 = (!((v4 % ((v4) | 1)) != 128))
                                    fmt.Println(v22, fi33, v4)
                                    v22 = int(v4)
                                    var v34 bool = (0 >= 42)
                                    var v35 bool = (!v34)
                                    _ = v34
                                    _ = v35
                                }
                            }
                            if false {
                                v6 = (v6 || true)
                            } else {
                                v36 := GenS1{F0: -9223372036854775808}
                                switch (("foo bar" + "\n\t") + (("\n\t" + "hello") + "\n\t")) {
                                case "":
                                    v6 = (647 > 63)
                                case "hello":
                                    if ((v4 <= v4) && false) {
                                        if true {
                                            switch fi33 {
                                            case 32767:
                                                fmt.Println(v22, fi25)
                                            default:
                                                var v37 int8 = 0
                                                _ = v37
                                            }
                                            v38 := GenS2{F0: 32767, F1: 42, F2: 720}
                                            fmt.Println(v4)
                                            fmt.Println(fi33, v4)
                                            if true {
                                                var v39 int32 = 0
                                                v5 = append(v5, (fi25 - -2147483648))
                                                _ = v39
                                            }
                                            _ = v38
                                        }
                                        v40 := GenS2{F0: int16(fi33), F1: 61, F2: -2147483648}
                                        var v41 uint32 = 2
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if (0 > 42) {
                                            switch v22 {
                                            case 127:
                                                var v42 int16 = 1000
                                                _ = v42
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v43 uint16 = 42
                                                    _ = v43
                                                case "hello":
                                                    v44 := GenS3{F0: 100, F1: (fi33 + v22)}
                                                    _ = v44
                                                case "foo bar":
                                                    v4 = v4
                                                }
                                            }
                                        } else {
                                            var v45 uint = 931
                                            v22 -= (fi33 | -2147483648)
                                            _ = v45
                                        }
                                        v22 -= (fi33 << ((0) & 63))
                                    }
                                default:
                                    v46 := GenS4{F0: v4, F1: 956, F2: ((("hello" + ((("a" + ("foo bar" + "")) + (("a" + (("hello" + "foo bar") + ("hello" + "a"))) + "a")) + "\n\t")) + ((("a" + "hello") + "a") + ("hello" + ""))) + "")}
                                    var v47 int32 = -2147483648
                                    if v6 {
                                        v22 = fi25
                                        var v48 int = v22
                                        fmt.Println(v4, v6, v48)
                                        _ = v48
                                    } else {
                                        v4 = v4
                                        v49 := GenS4{F0: v4, F1: 0, F2: "foo bar"}
                                        if true {
                                            var v50 int8 = 100
                                            v5 = append(v5, (-(fi25)))
                                            _ = v50
                                        }
                                        _ = v49
                                    }
                                    _ = v46
                                    _ = v47
                                }
                                var v51 float32 = -2.25
                                _ = v36
                                _ = v51
                            }
                            delete(v3, (fi25 - fi25))
                            v52 := GenS4{F0: v4, F1: 825, F2: "a"}
                            fmt.Println(fi25, v4)
                            _ = v52
                        }
                        if (-2147483648 != -54) {
                            if (!v6) {
                                v4 -= v4
                                var v53 uint = 0
                                _ = v53
                            } else {
                                var v54 int32 = 2147483647
                                switch ("hello" + "hello") {
                                case "a":
                                    v55 := GenS5{F0: -756, F1: 9223372036854775807}
                                    _ = v55
                                case "\n\t":
                                    fmt.Println(v4)
                                default:
                                    switch ((("hello" + "hello") + ("foo bar" + "foo bar")) + "hello") {
                                    case "":
                                        v5 = append(v5, 423)
                                    case "hello":
                                        v6 = v6
                                    }
                                    var v56 int8 = -127
                                    if ((457 >= (fi25 ^ fi25)) && v6) {
                                        var v57 int = fi25
                                        var v58 int8 = v56
                                        _ = v57
                                        _ = v58
                                    } else {
                                        var v59 float32 = 3.14159
                                        if true {
                                            v2[(fi33 % (((fi25 + fi25)) | 1))] = fi25
                                            v54 -= 174
                                        } else {
                                            var v60 float32 = (1.5 / (-2.25 - v59))
                                            _ = v60
                                        }
                                        _ = v59
                                    }
                                    _ = v56
                                }
                                switch 127 {
                                case 9223372036854775807:
                                    var v61 bool = (v6 || ((v6 && (824 <= 575)) && (2147483648 > 320)))
                                    _ = v61
                                case 128:
                                    var v62 float64 = (1e10 - (-2.25 * 0.0))
                                    _ = v62
                                case 779:
                                    delete(v2, fi25)
                                default:
                                    var v63 int16 = -1
                                    _ = v63
                                }
                                v2[(-(fi33))] = 2147483647
                                _ = v54
                            }
                            fmt.Println(v22, fi25, fi33)
                            v64 := GenS6{F0: 1}
                            _ = v64
                        }
                        v65 := GenS3{F0: 127, F1: (v22 & v22)}
                        switch 2147483647 {
                        case 466:
                            v66 := GenS7{F0: 255, F1: 9223372036854775807, F2: v4}
                            if (!false) {
                                var v67 int16 = 42
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                                _ = v67
                            } else {
                                fmt.Println(v6)
                                switch ("foo bar" + "a") {
                                case "a":
                                    v6 = v6
                                }
                            }
                            _ = v66
                        case 32768:
                            v22 = (fi33 - 32767)
                        }
                        _ = v65
                    } else {
                        v22 = v22
                        if (((0.0 - 1.5) / 3.14159) == 3.14159) {
                            if true {
                                v6 = true
                                v22 = v22
                            } else {
                                v22 -= v22
                            }
                            if (v6 || true) {
                                v68 := GenS8{F0: 244, F1: 916, F2: 42}
                                fmt.Println(v6, v6, fi33)
                                _ = v68
                            } else {
                                v69 := GenS4{F0: (v4 % (((v4 | (v4 >> ((786) & 63)))) | 1)), F1: 4294967295, F2: "hello"}
                                if true {
                                    v70 := GenS7{F0: 32767, F1: 772, F2: v4}
                                    _ = v70
                                }
                                _ = v69
                            }
                        }
                    }
                    _ = fi33
                }
            } else {
                v6 = v6
                for fi71 := 0; fi71 < 1; fi71++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if v6 {
                        if v6 {
                            switch (("a" + "") + "") {
                            case "\n\t":
                                var v72 uint = 4294967295
                                _ = v72
                            default:
                                v73 := GenS9{F0: -1}
                                _ = v73
                            }
                            if true {
                                var v74 uint = 18446744073709551615
                                v2[-297] = fi71
                                _ = v74
                            } else {
                                if v6 {
                                    if (231 >= 0) {
                                        var v75 uint64 = 1
                                        var v76 float64 = 1.5
                                        var v77 int32 = 1
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    } else {
                                        var v78 uint32 = 4294967295
                                        v79 := GenS8{F0: 4294967295, F1: 26, F2: 1}
                                        _ = v78
                                        _ = v79
                                    }
                                    var v80 bool = (1000 <= 0)
                                    if (v80 || v80) {
                                        switch fi25 {
                                        case 39:
                                            v81 := GenS1{F0: 811}
                                            _ = v81
                                        }
                                        v4 += v4
                                    }
                                    var v82 uint16 = 0
                                    _ = v80
                                    _ = v82
                                } else {
                                    v6 = (v6 || (-843 == (fi71 | fi71)))
                                    v83 := GenS10{F0: v4, F1: int64(v4)}
                                    var v84 bool = false
                                    v85 := GenS2{F0: -1, F1: -128, F2: -993}
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                }
                                fmt.Println(v4)
                                _ = len(v1)
                            }
                            var v86 int64 = -1
                            _ = v86
                        }
                    }
                    _ = fi71
                }
            }
            v87 := v1["hello"]
            _ = fi25
            _ = v87
        }
        for fi88 := 0; fi88 < 7; fi88++ {
            v2[fi88] = -2147483648
            fmt.Println(v22, v6, v4)
            _ = fi88
        }
    }
    v89 := GenS6{F0: 171}
    var v90 uint32 = 95
    var v91 int8 = 42
    for fi92 := 0; fi92 < 1; fi92++ {
        var v93 uint64 = 357
        fmt.Println(v4, fi92)
        for fi94 := 0; fi94 < 16; fi94++ {
            var v95 float64 = ((1e10 + 1e10) + (1.5 + float64(-2.25)))
            _ = fi94
            _ = v95
        }
        _ = fi92
        _ = v93
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
