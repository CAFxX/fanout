package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 uint8 }
func genFn0(p0 uint, p1 int64, p2 int8) int32 {
    for fi4 := 0; fi4 < 15; fi4++ {
        switch "hello" {
        case "hello":
            if true {
                if false {
                    var v5 uint32 = 841
                    for fi6 := 0; fi6 < 8; fi6++ {
                        v5 += (^(v5))
                        _ = fi6
                    }
                    _ = v5
                }
            } else {
                if ((0 != p0) && true) {
                    var v7 bool = true
                    v8 := []int{(fi4 / (((fi4 * (fi4 >> ((p0) & 63)))) | 1))}
                    _ = v7
                    _ = v8
                } else {
                    p1 += p1
                    var v9 float64 = 0.0
                    v10 := []int{}
                    _ = v10
                    _ = v9
                }
                if false {
                    fmt.Println(p2)
                } else {
                    v11 := GenS1{F0: 1}
                    v12 := GenS2{F0: p0}
                    var v13 int16 = -32768
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 int16 = 1000
                _ = v14
            }
        case "a":
            if (true || false) {
                p1 -= p1
                p2 += p2
                if ((!true) || ("a" < "")) {
                    v15 := []int{fi4}
                    p0 = p0
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS3{F0: ""}
                    fmt.Println(p2, fi16, p0)
                    fmt.Println(p1, p2, p2)
                    _ = fi16
                    _ = v17
                }
                var v18 int8 = (p2 + p2)
                if false {
                    v19 := GenS3{F0: ("a" + "")}
                    _ = v19
                } else {
                    if (16 != 32767) {
                        v20 := GenS4{F0: (p0 << ((p0) & 63))}
                        if (!((true && false) && (true && (!((65535 >= 552) && false))))) {
                            v21 := GenS5{F0: 42, F1: int8(p1)}
                            _ = v21
                        } else {
                            if (!false) {
                                v22 := GenS3{F0: (("" + "\n\t") + "\n\t")}
                                switch "" {
                                case "hello":
                                    switch 1 {
                                    case -1:
                                        for fi23 := 0; fi23 < 2; fi23++ {
                                            v24 := []int{fi4, fi23, -1, fi23}
                                            _ = fi23
                                            _ = v24
                                        }
                                    default:
                                        for fi25 := 0; fi25 < 16; fi25++ {
                                            v18 = (^(p2))
                                            _ = fi25
                                        }
                                    }
                                case "":
                                    p0 += p0
                                case "\n\t":
                                    p1 = p1
                                default:
                                    fmt.Println(p1, v18)
                                }
                                v18 += -128
                                _ = v22
                            }
                        }
                        _ = v20
                    } else {
                        v18 += p2
                        p2 = (p2 / ((p2) | 1))
                        var v26 uint64 = 816
                        _ = v26
                    }
                }
                _ = v18
            }
        case "\n\t":
            for fi27 := 0; fi27 < 8; fi27++ {
                switch (-(fi4)) {
                case -9223372036854775808:
                    p0 -= p0
                }
                if (1 >= fi4) {
                    p1 += 9223372036854775807
                    var v28 uint64 = 1
                    var v29 uint32 = 2
                    _ = v28
                    _ = v29
                } else {
                    v30 := []int{fi4}
                    switch 4294967295 {
                    case -2147483648:
                        var v31 uint8 = 127
                        _ = v31
                    default:
                        var v32 bool = (!(2 != 18446744073709551615))
                        _ = v32
                    }
                    var v33 uint16 = 65535
                    var v34 int64 = (p1 << ((p0) & 63))
                    fmt.Println(fi27)
                    _ = v30
                    _ = v33
                    _ = v34
                }
                _ = fi27
            }
        default:
            v35 := GenS3{F0: ("\n\t" + "foo bar")}
            _ = v35
        }
        v36 := GenS3{F0: "a"}
        _ = fi4
        _ = v36
    }
    switch 93 {
    case 429:
        var v37 uint = p0
        _ = v37
    }
    var v38 uint64 = 528
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    return 495
}
func genFn1(p0 uint, p1 uint) int64 {
    v39 := GenS2{F0: p0}
    if true {
        p0 = 18446744073709551615
        switch 742 {
        case 32767:
            v40 := GenS6{F0: 32767, F1: genFn0((^(p1)), 0, 28), F2: 599}
            _ = v40
        case 9223372036854775807:
            var v41 string = ("hello" + ((("a" + ("foo bar" + "hello")) + (("a" + "hello") + ("" + ((("" + (("" + (("" + (("hello" + ("foo bar" + "\n\t")) + "foo bar")) + "a")) + ("a" + "\n\t"))) + ("foo bar" + "hello")) + ("foo bar" + ""))))) + "a"))
            _ = v41
        }
    } else {
        p0 -= p0
        if ((0.1 * (((1e10 - float64(-2.25)) - 3.14159) + (float64(1.5) / (3.14159 * ((1e10 + 1e10) * 0.1))))) <= 1.5) {
            switch "\n\t" {
            case "foo bar":
                for fi42 := 0; fi42 < 3; fi42++ {
                    switch (fi42 & fi42) {
                    case -163:
                        p1 = 4294967295
                    case 32767:
                        if (false || true) {
                            for fi43 := 0; fi43 < 6; fi43++ {
                                p0 = p0
                                v44 := []int{fi42, (^(fi42)), fi43, (^(fi43))}
                                _ = fi43
                                _ = v44
                            }
                            v45 := GenS1{F0: 42}
                            _ = v45
                        } else {
                            fmt.Println(p1, p0, p0)
                            p1 += 255
                        }
                    case 662:
                        v46 := GenS5{F0: 467, F1: 0}
                        _ = v46
                    }
                    var v47 int16 = -1
                    var v48 int32 = 0
                    _ = fi42
                    _ = v47
                    _ = v48
                }
            case "\n\t":
                p0 += p0
            default:
                switch "\n\t" {
                case "hello":
                    if false {
                        v49 := []int{42, -275}
                        var v50 bool = (false && (true && false))
                        if v50 {
                            if false {
                                v51 := GenS6{F0: 1000, F1: -1, F2: 158}
                                if ((1e10 + -2.25) == -2.25) {
                                    v52 := GenS5{F0: 1, F1: 100}
                                    fmt.Println(p1)
                                    for fi53 := 0; fi53 < 1; fi53++ {
                                        var v54 uint64 = 57
                                        fmt.Println(fi53, v54)
                                        fmt.Println(v54, v54, p0)
                                        _ = fi53
                                        _ = v54
                                    }
                                    _ = v52
                                } else {
                                    switch 42 {
                                    case -342:
                                        if len(v49) > 0 {
                                            v49[0] = -9223372036854775808
                                        }
                                    case 2147483647:
                                        v55 := GenS1{F0: 1}
                                        _ = v55
                                    case 1:
                                        var v56 uint32 = 2
                                        _ = v56
                                    }
                                }
                                v57 := GenS5{F0: 32767, F1: 127}
                                _ = v51
                                _ = v57
                            }
                            v58 := GenS7{F0: true, F1: 286}
                            p0 = 0
                            _ = v58
                        } else {
                            var v59 uint16 = 572
                            for fi60 := 0; fi60 < 6; fi60++ {
                                v49 = append(v49, fi60)
                                _ = fi60
                            }
                            _ = v59
                        }
                        _ = v49
                        _ = v50
                    }
                case "foo bar":
                    p1 = 1
                }
            }
            var v61 int64 = -9223372036854775808
            p1 = (^(p0))
            _ = v61
        }
    }
    if true {
        v62 := []int{1}
        v63 := GenS7{F0: (0 < 9223372036854775807), F1: 32767}
        _ = v62
        _ = v63
    } else {
        v64 := GenS8{F0: 360, F1: 1}
        _ = v64
    }
    p1 += p0
    p1 += (p1 - p1)
    _ = p0
    _ = p1
    _ = v39
    return -1
}

func main() {
    var v1 uint8 = 40
    v2 := make(map[string]uint32)
    var v3 int = 301
    var v65 float64 = 1.5
    v66 := GenS7{F0: (2 != 1), F1: 794}
    var v67 uint64 = 800
    var v68 string = "\n\t"
    delete(v2, v68)
    var v69 uint32 = 2147483648
    fmt.Println(v65, v68, v68)
    var v70 uint = 887
    if false {
        for fi71 := 0; fi71 < 18; fi71++ {
            switch v68 {
            case "\n\t":
                for fi72 := 0; fi72 < 19; fi72++ {
                    v73 := GenS8{F0: -582, F1: 42}
                    switch "hello" {
                    case "\n\t":
                        if true {
                            fmt.Println(v67, v68, v70)
                        }
                        v74 := GenS8{F0: 0, F1: 255}
                        v75 := GenS9{F0: (v67 >= v67), F1: 1}
                        _ = v74
                        _ = v75
                    case "":
                        v76 := GenS10{F0: 32767, F1: (fi71 - fi72)}
                        _ = v76
                    default:
                        switch (v3 * (fi71 & 9223372036854775807)) {
                        case -155:
                            var v77 uint64 = 2
                            _ = v77
                        case 32767:
                            switch -666 {
                            case 32768:
                                var v78 int32 = -1
                                _ = v78
                            }
                        case 127:
                            v65 -= v65
                        default:
                            if (fi71 != (fi71 >> ((v70) & 63))) {
                                fmt.Println(v1, v70, fi71)
                            }
                        }
                    }
                    _ = fi72
                    _ = v73
                }
            case "a":
                var v79 float32 = 1.5
                _ = v79
            default:
                var v80 int64 = 968
                _ = v80
            }
            if false {
                if (false && false) {
                    var v81 int16 = 32767
                    var v82 int64 = 0
                    v83 := []int{fi71, (fi71 / ((int(v70)) | 1)), fi71, v3}
                    _ = v81
                    _ = v82
                    _ = v83
                } else {
                    if true {
                        var v84 float64 = (-2.25 * v65)
                        v68 = "\n\t"
                        _ = v84
                    } else {
                        var v85 bool = false
                        _ = v85
                    }
                }
            }
            _ = fi71
        }
    } else {
        if ((v3 * (v3 & -699)) != (v3 % ((v3) | 1))) {
            switch v68 {
            case "a":
                v70 = (^(v70))
            case "hello":
                switch v68 {
                case "hello":
                    _ = len(v2)
                case "\n\t":
                    v86 := []int{}
                    _ = v86
                }
            case "\n\t":
                var v87 float64 = v65
                _ = v87
            }
            fmt.Println(v70)
            v1 = (v1 << ((v70) & 63))
        } else {
            var v88 float32 = float32(v65)
            _ = v88
        }
        v1 = (v1 >> ((199) & 63))
    }
    v89 := GenS11{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v89
}
