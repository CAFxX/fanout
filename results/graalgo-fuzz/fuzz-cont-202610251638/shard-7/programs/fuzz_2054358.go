package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int; F2 string }
func genFn0(p0 int32) bool {
    var v5 float32 = 0.0
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch (fi6 + (-(fi7))) {
            case 127:
                var v8 float64 = 1e10
                _ = v8
            case -2147483648:
                p0 = p0
            case 181:
                v5 -= v5
            }
            var v9 bool = ((1.5 * v5) < 1e10)
            _ = fi7
            _ = v9
        }
        var v10 string = "hello"
        _ = fi6
        _ = v10
    }
    if (248 != 173) {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: -1, F1: 42}
            v13 := []int{fi11, fi11, fi11, fi11}
            if false {
                var v14 uint64 = 4294967295
                for fi15 := 0; fi15 < 9; fi15++ {
                    v5 = v5
                    _ = fi15
                }
                v14 = (^(v14))
                _ = v14
            } else {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int8 = -16
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS2{F0: (("hello" + "") + ("foo bar" + "a")), F1: -1}
            v5 = -2.25
            var v20 bool = (42 < 65535)
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint64 = 2
        _ = v21
    }
    _ = p0
    _ = v5
    return (32767 > -1)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-320}
    var v3 bool = true
    var v4 int16 = -1
    v4 += (v4 | (v4 << ((1) & 63)))
    fmt.Println(v3, v4)
    v22 := GenS3{F0: 0}
    if (v3 && true) {
        v3 = v3
        var v23 string = ("hello" + "\n\t")
        if true {
            for fi24 := 0; fi24 < 8; fi24++ {
                if false {
                    if (42 == 4294967295) {
                        v4 += -32768
                        if true {
                            fmt.Println(v3, v3)
                            v25 := GenS3{F0: int8(fi24)}
                            if (!(!((true && (false || v3)) && true))) {
                                switch (^(fi24)) {
                                case 127:
                                    _ = len(v2)
                                case -9223372036854775808:
                                    for fi26 := 0; fi26 < 3; fi26++ {
                                        _ = len(v1)
                                        _ = fi26
                                    }
                                case 989:
                                    fmt.Println(v23, v4, v4)
                                default:
                                    fmt.Println(v4)
                                }
                                v4 = 32767
                                v27 := v1[fi24]
                                _ = v27
                            }
                            _ = v25
                        }
                    } else {
                        delete(v1, (fi24 % ((fi24) | 1)))
                        var v28 float64 = 0.0
                        _ = v28
                    }
                    fmt.Println(v3, v23)
                }
                v29 := GenS1{F0: 42, F1: 1}
                v23 = v23
                _ = fi24
                _ = v29
            }
            var v30 uint = 255
            if v3 {
                fmt.Println(v3)
                if (false || (float64(-2.25) > (1e10 + (((float64(-2.25) / 1e10) + (1e10 + 0.1)) - float64((((((float32(1e10) + (1.5 / -2.25)) - (-2.25 - (1.5 / (1e10 + 3.14159)))) / 1.5) + (0.1 + 0.0)) / 3.14159)))))) {
                    v4 += v4
                }
            } else {
                switch v23 {
                case "":
                    var v31 float32 = (0.0 - (-2.25 - 1.5))
                    _ = v31
                case "\n\t":
                    if false {
                        if (!true) {
                            v30 -= 18446744073709551615
                        } else {
                            v32 := GenS2{F0: v23, F1: 3}
                            for fi33 := 0; fi33 < 7; fi33++ {
                                var v34 bool = ((fi33 % ((fi33) | 1)) < fi33)
                                var v35 float64 = (0.1 * 0.1)
                                fmt.Println(v23)
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            v23 = v23
                            _ = v32
                        }
                        v4 += int16(v30)
                        if (v30 > v30) {
                            var v36 int16 = (v4 ^ (v4 & -549))
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 15; fi37++ {
                                v38 := GenS4{F0: -72}
                                var v39 int = 32768
                                _ = len(v1)
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            var v40 float32 = 0.1
                            v41 := v1[128]
                            _ = v40
                            _ = v41
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                    }
                case "foo bar":
                    if v3 {
                        _ = len(v2)
                        var v42 int64 = 715
                        fmt.Println(v3)
                        _ = v42
                    } else {
                        v30 = 65535
                        v43 := GenS5{F0: 128, F1: 22}
                        _ = v43
                    }
                }
                v4 = v4
                var v44 float64 = (3.14159 + 1e10)
                _ = v44
            }
            _ = v30
        } else {
            for fi45 := 0; fi45 < 13; fi45++ {
                switch (fi45 + fi45) {
                case 765:
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v3)
                    case "hello":
                        switch v23 {
                        case "":
                            v4 = -431
                        }
                    case "":
                        for fi46 := 0; fi46 < 8; fi46++ {
                            v1[fi45] = 0
                            _ = fi46
                        }
                    }
                case 9223372036854775807:
                    fmt.Println(fi45)
                case -2147483648:
                    var v47 string = v23
                    _ = v47
                default:
                    v4 += 32767
                }
                _ = fi45
            }
            fmt.Println(v23)
            if ((-1 == -9223372036854775808) || (!v3)) {
                fmt.Println(v23, v23, v23)
                switch v23 {
                case "\n\t":
                    var v48 uint32 = 1
                    _ = v48
                default:
                    if false {
                        var v49 uint8 = 0
                        v49 += 127
                        fmt.Println(v3)
                        _ = v49
                    }
                }
            } else {
                for fi50 := 0; fi50 < 14; fi50++ {
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := GenS6{F0: 42}
                        v4 -= (v4 - (v4 & v4))
                        _ = len(v2)
                        _ = fi51
                        _ = v52
                    }
                    _ = fi50
                }
            }
        }
        _ = v23
    } else {
        v4 -= (v4 - (v4 - v4))
        for fi53 := 0; fi53 < 8; fi53++ {
            var v54 float32 = (float32(3.14159) * (0.0 / (0.1 + (3.14159 / -2.25))))
            _ = fi53
            _ = v54
        }
    }
    var v55 uint32 = 2147483647
    v56 := GenS7{F0: v55, F1: 127, F2: (!(0 >= 0))}
    fmt.Println(v4)
    if (v55 != v55) {
        var v57 int16 = -722
        if false {
            switch (("hello" + ("a" + ("\n\t" + "a"))) + "a") {
            case "a":
                v58 := GenS4{F0: 42}
                _ = v58
            }
            v55 = 65535
            fmt.Println(v55, v57, v3)
        } else {
            fmt.Println(v57, v3, v4)
        }
        _ = v57
    }
    if ((v4 - (v4 % (((v4 & v4)) | 1))) >= v4) {
        for fi59 := 0; fi59 < 12; fi59++ {
            for fi60 := 0; fi60 < 6; fi60++ {
                if v3 {
                    switch "a" {
                    case "":
                        var v61 int64 = 42
                        _ = v61
                    case "foo bar":
                        switch fi60 {
                        case 285:
                            fmt.Println(fi60, v3)
                        case 630:
                            if (-89 == 2147483647) {
                                var v62 float32 = 0.1
                                _ = v62
                            }
                        case 2147483647:
                            switch fi59 {
                            case 536:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -9223372036854775808:
                                if true {
                                    v3 = v3
                                    var v63 bool = v3
                                    _ = v63
                                } else {
                                    v64 := GenS2{F0: ("foo bar" + "a"), F1: -35}
                                    v65 := GenS1{F0: 0, F1: 65535}
                                    _ = v64
                                    _ = v65
                                }
                            }
                            fmt.Println(fi59)
                        }
                    }
                    if (!((v4 - v4) == v4)) {
                        var v66 string = ""
                        v55 += (^(v55))
                        _ = v66
                    }
                } else {
                    var v67 int8 = -128
                    var v68 uint32 = 465
                    var v69 int = fi60
                    v70 := GenS8{F0: 0}
                    v71 := GenS1{F0: -855, F1: 42}
                    _ = v67
                    _ = v68
                    _ = v69
                    _ = v70
                    _ = v71
                }
                _ = fi60
            }
            if (v55 < v55) {
                if true {
                    for fi72 := 0; fi72 < 16; fi72++ {
                        delete(v1, (fi72 * (^(fi72))))
                        v3 = (v55 > v55)
                        v4 = v4
                        _ = fi72
                    }
                    if (("hello" < "hello") && v3) {
                        if len(v2) > 0 {
                            v2[0] = fi59
                        }
                        if false {
                            v2 = append(v2, -641)
                            fmt.Println(v55, v55, v55)
                        }
                        fmt.Println(v3, v4, v55)
                    } else {
                        var v73 int = fi59
                        _ = v73
                    }
                    v55 -= (v55 - v55)
                } else {
                    var v74 int = fi59
                    _ = v74
                }
                switch (("\n\t" + "a") + "hello") {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = (fi59 & (fi59 % (((fi59 / (((fi59 >> ((18446744073709551615) & 63))) | 1))) | 1)))
                    }
                case "foo bar":
                    v55 -= v55
                default:
                    for fi75 := 0; fi75 < 9; fi75++ {
                        v3 = true
                        var v76 bool = (((!true) || (("" + "\n\t") <= "hello")) && (!(232 == 9223372036854775807)))
                        v77 := GenS9{F0: ""}
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                }
            } else {
                if (0 >= 94) {
                    var v78 uint64 = 759
                    var v79 int32 = 0
                    _ = v78
                    _ = v79
                } else {
                    v80 := GenS4{F0: -1}
                    var v81 uint = 252
                    _ = v80
                    _ = v81
                }
                switch "hello" {
                case "a":
                    v4 -= (v4 + (v4 / ((v4) | 1)))
                case "hello":
                    var v82 int16 = v4
                    _ = v82
                case "":
                    if (v3 && v3) {
                        v55 = v55
                    }
                }
            }
            _ = fi59
        }
    }
    v83 := GenS10{F0: v3, F1: -347, F2: "a"}
    v55 = 2147483648
    v3 = (("" <= "hello") || (!((42 == 187) && (1 < 255))))
    v84 := GenS10{F0: true, F1: 127, F2: "hello"}
    var v85 float64 = (float64(0.1) * (0.0 + 1e10))
    var v86 int64 = 9223372036854775807
    var v87 int16 = (v4 | v4)
    var v88 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
