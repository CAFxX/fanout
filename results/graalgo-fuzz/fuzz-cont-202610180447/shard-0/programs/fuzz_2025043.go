package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int8; F2 int16 }
type GenS7 struct { F0 int64; F1 int64; F2 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 string; F2 bool }
func genFn0(p0 string) string {
    fmt.Println(p0)
    var v5 uint32 = 65535
    if ((!false) && true) {
        v6 := GenS1{F0: 0, F1: 42, F2: uint16(v5)}
        if false {
            p0 = (p0 + "foo bar")
        }
        _ = v6
    }
    _ = p0
    _ = v5
    return (p0 + p0)
}
func genFn1(p0 uint64) int64 {
    v7 := []int{-9223372036854775808, 9223372036854775807, 1}
    if (42 < -1) {
        v8 := GenS2{F0: false, F1: -806, F2: 65535}
        _ = v8
    } else {
        p0 -= (p0 & p0)
    }
    if (!false) {
        v9 := GenS1{F0: 1, F1: 350, F2: 255}
        _ = v9
    } else {
        v10 := GenS1{F0: -1, F1: -32768, F2: 2}
        _ = v10
    }
    var v11 int64 = -1
    _ = p0
    _ = v11
    _ = v7
    return v11
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -305
    v3 := []string{((("\n\t" + "a") + "") + ("hello" + ("\n\t" + ("" + "")))), ""}
    var v4 bool = (3.14159 < 1e10)
    var v12 int32 = (^(v2))
    var v13 int32 = (v12 / ((-1) | 1))
    v14 := GenS3{F0: 1, F1: 42, F2: 653}
    _ = len(v3)
    if v4 {
        fmt.Println(v4, v2, v2)
    } else {
        switch ("foo bar" + ("hello" + "")) {
        case "a":
            v13 = v13
        case "foo bar":
            fmt.Println(v4)
        }
    }
    fmt.Println(v4, v13)
    v15 := GenS4{F0: 131}
    v12 = (v13 & v13)
    switch ("\n\t" + "foo bar") {
    case "":
        if (v4 && true) {
            v16 := GenS1{F0: 0, F1: 32767, F2: 65535}
            var v17 int64 = -1
            var v18 uint = 18
            _ = v16
            _ = v17
            _ = v18
        } else {
            v19 := GenS5{F0: "foo bar"}
            var v20 uint = 0
            v21 := GenS6{F0: v20, F1: 42, F2: 0}
            _ = v19
            _ = v20
            _ = v21
        }
    case "foo bar":
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS2{F0: true, F1: 1, F2: 60}
            for fi24 := 0; fi24 < 8; fi24++ {
                v2 = v12
                if (9223372036854775807 != -699) {
                    fmt.Println(v12)
                    fmt.Println(v2, v12, v2)
                    if (v4 || v4) {
                        if false {
                            v25 := GenS3{F0: 4294967295, F1: 203, F2: 638}
                            fmt.Println(v12)
                            var v26 int64 = 902
                            v27 := GenS7{F0: v26, F1: 1, F2: v4}
                            fmt.Println(v12, fi24, v26)
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        v2 += v2
                    }
                }
                _ = fi24
            }
            switch "" {
            case "hello":
                switch ("a" + "") {
                case "hello":
                    if (true && v4) {
                        v28 := v1[fi22]
                        fmt.Println(v28)
                        delete(v1, v28)
                        _ = v28
                    }
                case "a":
                    var v29 uint = 42
                    _ = v29
                case "":
                    var v30 int16 = 1
                    _ = v30
                }
            case "\n\t":
                v1[4294967295] = 128
            case "":
                switch fi22 {
                case 733:
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v2 = v2
                        v32 := GenS3{F0: 1, F1: 42, F2: 65535}
                        if (!true) {
                            v13 = v13
                            v33 := GenS5{F0: "foo bar"}
                            v3 = append(v3, "foo bar")
                            _ = v33
                        } else {
                            switch (fi22 >> ((4294967295) & 63)) {
                            case 128:
                                v12 -= (v13 % (((v12 / ((v2) | 1))) | 1))
                            case 524:
                                if true {
                                    switch (fi31 | fi22) {
                                    case 32768:
                                        v34 := v1[(fi22 | (fi22 ^ 2147483647))]
                                        var v35 int16 = 1000
                                        if ((((v35 >> ((42) & 63)) >= (v35 >> ((583) & 63))) || true) && v4) {
                                            var v36 int = (fi31 + v34)
                                            v2 += v2
                                            _ = v36
                                        }
                                        var v37 int8 = -1
                                        switch (fi31 << ((0) & 63)) {
                                        case 11:
                                            var v38 float32 = 1e10
                                            _ = v38
                                        }
                                        _ = v34
                                        _ = v35
                                        _ = v37
                                    case 4294967295:
                                        var v39 int8 = -128
                                        var v40 int = 9223372036854775807
                                        _ = v39
                                        _ = v40
                                    case 128:
                                        var v41 uint32 = 477
                                        switch fi31 {
                                        case -497:
                                            if v4 {
                                                v42 := GenS8{F0: 0}
                                                _ = v42
                                            } else {
                                                var v43 float32 = 1e10
                                                _ = v43
                                            }
                                        default:
                                            delete(v1, fi31)
                                        }
                                        _ = v41
                                    default:
                                        fmt.Println(v13, v4, fi22)
                                    }
                                } else {
                                    v12 = v2
                                }
                                fmt.Println(v2, v13, v12)
                                delete(v1, fi31)
                            case 32768:
                                v44 := GenS8{F0: 2}
                                _ = v44
                            default:
                                v45 := GenS6{F0: 65535, F1: 0, F2: 967}
                                _ = v45
                            }
                            v46 := GenS7{F0: 9223372036854775807, F1: 9223372036854775807, F2: v4}
                            if (v4 && false) {
                                var v47 int8 = 19
                                v4 = (0 <= -376)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v48 uint64 = 4294967295
                                switch "" {
                                case "hello":
                                    var v49 bool = (1 <= 1)
                                    _ = v49
                                default:
                                    v13 += v12
                                }
                                _ = v47
                                _ = v48
                            }
                            _ = v46
                        }
                        _ = fi31
                        _ = v32
                    }
                case 128:
                    fmt.Println(fi22, v12, fi22)
                default:
                    v50 := GenS8{F0: 4294967295}
                    _ = v50
                }
            default:
                for fi51 := 0; fi51 < 18; fi51++ {
                    _ = len(v1)
                    var v52 int32 = (v13 ^ v13)
                    delete(v1, 128)
                    _ = fi51
                    _ = v52
                }
            }
            _ = fi22
            _ = v23
        }
    case "a":
        _ = len(v1)
    default:
        var v53 uint16 = 0
        _ = v53
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        if v4 {
            v12 = v2
        }
        _ = fi54
    }
    if (!v4) {
        if ("foo bar" >= "hello") {
            var v55 uint32 = 1
            _ = v55
        }
        v4 = v4
    }
    switch 4294967295 {
    case 1:
        v56 := GenS2{F0: (false && (-2147483648 != 745)), F1: 2147483647, F2: 0}
        _ = v56
    default:
        if (int(v2) > -249) {
            var v57 bool = (42 != 0)
            switch "hello" {
            case "":
                switch "" {
                case "":
                    v2 += (v2 | (v12 << ((701) & 63)))
                case "hello":
                    v58 := v1[32767]
                    _ = v58
                default:
                    _ = len(v1)
                }
            case "\n\t":
                v59 := GenS7{F0: -1, F1: 413, F2: (!false)}
                _ = v59
            case "hello":
                if v4 {
                    v13 -= (v2 - (v2 - 1))
                    for fi60 := 0; fi60 < 3; fi60++ {
                        if len(v3) > 0 {
                            v3[0] = ("a" + "hello")
                        }
                        if v4 {
                            if (0 < 255) {
                                _ = len(v3)
                            }
                            v4 = v4
                            var v61 uint16 = 854
                            _ = v61
                        }
                        _ = fi60
                    }
                    if v57 {
                        var v62 int = -971
                        var v63 uint8 = 90
                        _ = v62
                        _ = v63
                    } else {
                        var v64 uint16 = 735
                        v57 = true
                        _ = v64
                    }
                } else {
                    if v57 {
                        var v65 uint16 = 408
                        if v4 {
                            if (("foo bar" + "foo bar") >= "\n\t") {
                                v66 := GenS9{F0: 4294967295, F1: "\n\t", F2: ((!((!(!(42 == 575))) || (v13 == v2))) && (v2 <= v13))}
                                v65 -= (^(v65))
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    switch "\n\t" {
                                    case "a":
                                        var v68 int64 = -9223372036854775808
                                        _ = v68
                                    default:
                                        fmt.Println(v2, v12)
                                    }
                                    _ = len(v3)
                                    var v69 string = "foo bar"
                                    _ = fi67
                                    _ = v69
                                }
                                _ = v66
                            }
                            for fi70 := 0; fi70 < 10; fi70++ {
                                v1[fi70] = (fi70 + (-(fi70)))
                                _ = fi70
                            }
                            fmt.Println(v12, v13, v12)
                        }
                        var v71 bool = v4
                        _ = v65
                        _ = v71
                    } else {
                        v2 += -101
                        for fi72 := 0; fi72 < 6; fi72++ {
                            v57 = v57
                            _ = len(v1)
                            _ = fi72
                        }
                        if (1.5 >= 3.14159) {
                            if v57 {
                                _ = len(v3)
                            } else {
                                switch "\n\t" {
                                case "":
                                    var v73 int16 = -121
                                    _ = v73
                                case "foo bar":
                                    if (!false) {
                                        var v74 uint32 = 65535
                                        _ = v74
                                    } else {
                                        fmt.Println(v13)
                                    }
                                }
                                if (v57 && v57) {
                                    if ("a" == "") {
                                        if (!true) {
                                            _ = len(v3)
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                            v57 = (!(!((1.5 < 0.0) && (2 != 2))))
                                        }
                                    } else {
                                        switch -1 {
                                        case -9223372036854775808:
                                            var v75 int16 = 32767
                                            _ = v75
                                        case 0:
                                            v76 := v1[1]
                                            _ = v76
                                        case -508:
                                            fmt.Println(v2, v4, v12)
                                        }
                                        v1[-48] = 4294967295
                                    }
                                } else {
                                    switch "" {
                                    case "hello":
                                        v12 = v13
                                    case "foo bar":
                                        switch "a" {
                                        case "foo bar":
                                            var v77 uint8 = 2
                                            _ = v77
                                        }
                                    case "":
                                        v4 = ((2 < 1) && ((436 <= 273) || (255 != 1)))
                                    }
                                    for fi78 := 0; fi78 < 12; fi78++ {
                                        var v79 uint = 18446744073709551615
                                        _ = fi78
                                        _ = v79
                                    }
                                    var v80 int = 405
                                    _ = v80
                                }
                            }
                        } else {
                            var v81 int32 = (v12 % ((v2) | 1))
                            _ = v81
                        }
                    }
                }
            default:
                v4 = v4
            }
            fmt.Println(v57)
            _ = v57
        }
    }
    fmt.Println(v2)
    v2 = (v12 >> ((2) & 63))
    v2 += v12
    v2 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
