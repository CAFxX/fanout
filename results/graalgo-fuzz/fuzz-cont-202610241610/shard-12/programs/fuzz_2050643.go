package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int16; F2 int64 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 string }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 int8; F1 int8 }
func genFn0(p0 int, p1 int) int32 {
    fmt.Println(p0, p1)
    if (1.5 < (0.0 + 3.14159)) {
        v5 := GenS1{F0: ""}
        p0 = 42
        fmt.Println(p0, p1)
        _ = v5
    } else {
        var v6 uint8 = 42
        fmt.Println(p0, p1)
        _ = v6
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1() uint {
    v7 := GenS1{F0: "a"}
    switch "a" {
    case "":
        v8 := []int{42, 2147483647, 261, -950}
        _ = v8
    }
    var v9 int16 = -147
    var v10 int = -9223372036854775808
    _ = v10
    _ = v7
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 string = ""
    var v2 uint32 = 1
    v3 := make(map[uint32]string)
    v4 := []uint32{(v2 & v2)}
    if (65535 <= 2) {
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS1{F0: ((v1 + (v1 + "hello")) + ((((v1 + v1) + "\n\t") + (v1 + v1)) + "\n\t"))}
            _ = fi11
            _ = v12
        }
        switch 0 {
        case 1:
            var v13 uint16 = 32767
            _ = v13
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        v2 = 1
        v14 := GenS1{F0: "\n\t"}
        v15 := GenS1{F0: v1}
        _ = v14
        _ = v15
    }
    v1 = (v1 + (v1 + v1))
    v16 := GenS1{F0: (v1 + (v1 + "a"))}
    var v17 uint32 = (v2 | (^(v2)))
    if len(v4) > 0 {
        v4[0] = v2
    }
    switch (v1 + ("" + ("hello" + (v1 + "a")))) {
    case "hello":
        var v18 int = 602
        _ = v18
    case "foo bar":
        v1 = v1
    case "\n\t":
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v1)
            var v20 bool = true
            _ = fi19
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 19; fi22++ {
                if (v2 <= 42) {
                    if false {
                        var v23 int8 = 80
                        _ = v23
                    }
                    v24 := GenS1{F0: ((v1 + v1) + (("\n\t" + v1) + v1))}
                    v17 -= (v2 * (v17 + (v17 - (v2 / (((^(v2))) | 1)))))
                    if (0 >= -1) {
                        v17 = (^(v2))
                        var v25 uint16 = 65535
                        var v26 uint32 = v2
                        _ = v25
                        _ = v26
                    }
                    delete(v3, v2)
                    _ = v24
                } else {
                    fmt.Println(fi22, v2)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v27 uint32 = (v2 % (((v17 / ((v17) | 1))) | 1))
                fmt.Println(v2)
                var v28 bool = ((^(fi22)) > int(v27))
                v29 := GenS2{F0: v1, F1: 621, F2: 1}
                fmt.Println(v17, fi21, v27)
                if v28 {
                    _ = len(v4)
                }
                var v30 uint64 = 1
                var v31 string = v1
                v1 = ("" + "foo bar")
                switch -801 {
                case -2147483648:
                    delete(v3, 169)
                case -9223372036854775808:
                    var v32 int32 = 434
                    _ = v32
                case 0:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    fmt.Println(fi22)
                }
                fmt.Println(v27, v17, v27)
                _ = fi22
                _ = v27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            v3[(v17 * (v2 % (((^(v2))) | 1)))] = v1
            _ = fi21
        }
    }
    v33 := GenS1{F0: v1}
    fmt.Println(v2)
    v34 := GenS1{F0: "a"}
    v2 = v2
    if false {
        for fi35 := 0; fi35 < 10; fi35++ {
            v1 = (v1 + (v1 + "hello"))
            var v36 bool = true
            _ = fi35
            _ = v36
        }
        v1 = v1
        switch -250 {
        case 9223372036854775807:
            for fi37 := 0; fi37 < 3; fi37++ {
                switch (v1 + (v1 + (v1 + v1))) {
                case "a":
                    var v38 int16 = -362
                    _ = v38
                case "foo bar":
                    var v39 uint8 = 128
                    _ = v39
                }
                _ = fi37
            }
        case -1:
            v2 = v17
        }
    } else {
        v40 := GenS3{F0: -1}
        var v41 bool = true
        for fi42 := 0; fi42 < 20; fi42++ {
            for fi43 := 0; fi43 < 19; fi43++ {
                v1 = v1
                if v41 {
                    if v41 {
                        if (!(1e10 != 0.1)) {
                            var v44 int16 = -295
                            v45 := GenS4{F0: (!((((^(v2)) <= v2) && true) || v41)), F1: v44, F2: 42}
                            _ = v44
                            _ = v45
                        }
                        switch fi42 {
                        case 32768:
                            v46 := GenS1{F0: ("foo bar" + v1)}
                            _ = v46
                        case 1:
                            switch ((v1 + v1) + "hello") {
                            case "hello":
                                v17 = 42
                            case "foo bar":
                                v17 -= v17
                            case "a":
                                if (!(0.1 < -2.25)) {
                                    if (!(!v41)) {
                                        v3[v2] = "\n\t"
                                        v47 := GenS5{F0: -906, F1: 4294967295}
                                        v3[(^(v17))] = ("a" + ("" + v1))
                                        if false {
                                            v41 = (v41 || (true || (-9223372036854775808 <= 0)))
                                        }
                                        if false {
                                            v41 = v41
                                            if (v17 > v2) {
                                                switch v1 {
                                                case "a":
                                                    var v48 string = v1
                                                    _ = v48
                                                case "":
                                                    fmt.Println(v2, v2)
                                                    if (!(!v41)) {
                                                        v41 = v41
                                                    } else {
                                                        var v49 uint16 = 55
                                                        _ = v49
                                                    }
                                                }
                                            }
                                        } else {
                                            var v50 float32 = 1e10
                                            v51 := v3[(v2 | v17)]
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v47
                                    }
                                    v52 := GenS3{F0: 98}
                                    var v53 int64 = -266
                                    fmt.Println(v41, fi42, fi42)
                                    _ = v52
                                    _ = v53
                                }
                            default:
                                var v54 uint = 0
                                _ = v54
                            }
                        case 127:
                            v3[v17] = v1
                        }
                        v41 = (58 != 0)
                    } else {
                        var v55 float32 = float32(((float64(0.1) * 3.14159) - 0.0))
                        v4 = append(v4, 2147483648)
                        _ = v55
                    }
                    if true {
                        v41 = v41
                    } else {
                        if v41 {
                            var v56 uint16 = 2
                            _ = len(v4)
                            _ = v56
                        }
                    }
                }
                if ((float64(1.5) / (0.1 / 0.1)) > (0.1 / 1e10)) {
                    var v57 uint32 = v17
                    var v58 float32 = 1e10
                    _ = v57
                    _ = v58
                } else {
                    if (1 < 42) {
                        fmt.Println(v1, v2, v2)
                        switch (^(fi42)) {
                        case 1:
                            v17 = (v17 % (((v17 - v17)) | 1))
                        case -224:
                            var v59 uint8 = 127
                            _ = v59
                        default:
                            switch v1 {
                            case "hello":
                                v3[(v2 & (v17 + (v2 % ((2147483648) | 1))))] = ((v1 + v1) + v1)
                            default:
                                v17 = 2147483648
                            }
                        }
                    } else {
                        if v41 {
                            if ((fi42 / ((37) | 1)) <= (fi43 % ((fi42) | 1))) {
                                v1 = "\n\t"
                                var v60 int16 = 32767
                                _ = v60
                            } else {
                                if false {
                                    v61 := GenS6{F0: fi42, F1: false, F2: (v1 + v1)}
                                    _ = len(v4)
                                    _ = v61
                                }
                            }
                        }
                    }
                }
                _ = fi43
            }
            v4 = append(v4, 2)
            _ = fi42
        }
        _ = v40
        _ = v41
    }
    if (true && (1e10 > (1e10 / (1e10 * 3.14159)))) {
        v1 = v1
        v62 := GenS3{F0: 127}
        _ = v62
    } else {
        v2 += 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            fmt.Println(v1, v17)
            switch v1 {
            case "foo bar":
                for fi63 := 0; fi63 < 13; fi63++ {
                    v2 -= 65535
                    v64 := v3[(v17 + (v17 ^ v2))]
                    var v65 int32 = 2147483647
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            case "hello":
                v2 = (v17 >> ((1) & 63))
            case "\n\t":
                v66 := GenS7{F0: 128, F1: 127}
                _ = v66
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    if true {
        if (-108 <= 375) {
            switch v1 {
            case "a":
                v67 := GenS1{F0: ""}
                _ = v67
            case "foo bar":
                var v68 bool = true
                _ = v68
            case "\n\t":
                var v69 uint16 = 2
                _ = v69
            }
        } else {
            v2 -= 65535
            for fi70 := 0; fi70 < 9; fi70++ {
                v71 := GenS3{F0: 127}
                _ = fi70
                _ = v71
            }
        }
        var v72 uint = 412
        v17 = v17
        _ = v72
    } else {
        var v73 float64 = (1e10 / -2.25)
        _ = v73
    }
    var v74 uint64 = 18446744073709551615
    if (!(!(false && true))) {
        v74 = v74
        for fi75 := 0; fi75 < 20; fi75++ {
            v2 -= 297
            v17 -= 2147483647
            _ = fi75
        }
    } else {
        if (((455 > 165) && false) || (false && true)) {
            for fi76 := 0; fi76 < 13; fi76++ {
                fmt.Println(v74, v1, v1)
                switch fi76 {
                case -9223372036854775808:
                    _ = len(v4)
                default:
                    var v77 int8 = 100
                    _ = v77
                }
                _ = fi76
            }
            var v78 int16 = 42
            _ = v78
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (v1 + (v1 + v1)) {
            case "\n\t":
                v79 := v3[(v2 | v17)]
                _ = v79
            default:
                switch v1 {
                case "\n\t":
                    switch 0 {
                    case -926:
                        var v80 float32 = (3.14159 - (1e10 - ((1.5 - 0.0) / 0.1)))
                        _ = v80
                    case 235:
                        for fi81 := 0; fi81 < 20; fi81++ {
                            _ = len(v3)
                            fmt.Println(v17, fi81)
                            _ = fi81
                        }
                    default:
                        switch v1 {
                        case "\n\t":
                            switch 698 {
                            case 128:
                                switch 32767 {
                                case 0:
                                    v82 := GenS8{F0: -1, F1: -87}
                                    _ = v82
                                case 42:
                                    var v83 bool = (!(v2 != (v17 & v17)))
                                    _ = v83
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                        default:
                            switch 127 {
                            case -253:
                                if ((float32((0.1 - float64(3.14159))) / float32(3.14159)) == 0.1) {
                                    v84 := GenS2{F0: (("foo bar" + ("\n\t" + v1)) + v1), F1: 1, F2: 1}
                                    v85 := GenS9{F0: -128, F1: int8(v2)}
                                    _ = v84
                                    _ = v85
                                } else {
                                    var v86 float64 = float64(1.5)
                                    var v87 bool = true
                                    _ = v86
                                    _ = v87
                                }
                            default:
                                fmt.Println(v2, v2)
                            }
                        }
                    }
                }
            }
            _ = len(v4)
        }
        if (((v74 * v74) == v74) || (100 > 42)) {
            v2 -= (^(v17))
        }
    }
    var v88 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v74:", v74)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v74
    _ = v88
}
