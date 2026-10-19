package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint; F2 bool }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 int; F1 uint16; F2 uint8 }
type GenS10 struct { F0 int; F1 string; F2 uint32 }
func genFn0(p0 int64) uint8 {
    v5 := GenS1{F0: 1, F1: 843}
    var v6 int8 = 0
    if false {
        v7 := GenS1{F0: 32767, F1: 2}
        v8 := []int{489, int(p0)}
        _ = v7
        _ = v8
    }
    var v9 int32 = 1
    var v10 bool = true
    _ = p0
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return 0
}
func genFn1(p0 int32, p1 int16, p2 int64) int8 {
    var v11 uint64 = 1
    fmt.Println(p2, p1)
    v12 := []int{1, 42, 1, -9223372036854775808}
    v13 := GenS1{F0: (p1 | (p1 | p1)), F1: 4294967295}
    var v14 int64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int16 = -215
    var v2 float64 = (1e10 + 1e10)
    var v3 uint16 = 32767
    var v4 int64 = 252
    for fi15 := 0; fi15 < 7; fi15++ {
        switch (fi15 + fi15) {
        case 2147483647:
            v16 := []int{fi15, (fi15 | (^(fi15))), fi15, 1}
            _ = v16
        case -223:
            switch -9223372036854775808 {
            case -303:
                v17 := GenS2{F0: -457, F1: 0}
                _ = v17
            default:
                if (!(176 <= v3)) {
                    var v18 int8 = 0
                    if (v18 == 1) {
                        if ("hello" > ("a" + ("" + "\n\t"))) {
                            switch (fi15 + fi15) {
                            case 4294967295:
                                fmt.Println(v18)
                            case -1:
                                var v19 int8 = genFn1(-2147483648, v1, v4)
                                _ = v19
                            case -2147483648:
                                if true {
                                    v20 := GenS3{F0: -2147483648, F1: 877, F2: (!false)}
                                    var v21 int32 = 625
                                    v22 := GenS1{F0: (v1 - (v1 | v1)), F1: 844}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        var v24 int32 = 2147483647
                                        _ = fi23
                                        _ = v24
                                    }
                                    v25 := []int{int(v3), (fi15 ^ fi15)}
                                    _ = v25
                                }
                            }
                        }
                        fmt.Println(v4)
                        fmt.Println(fi15, fi15, v1)
                    }
                    _ = v18
                } else {
                    switch "foo bar" {
                    case "a":
                        var v26 int32 = -1
                        _ = v26
                    }
                    v27 := GenS4{F0: (!(!(4294967295 <= 0))), F1: -802}
                    if ((!(false && (!true))) && false) {
                        if true {
                            if ("foo bar" <= (("foo bar" + "") + "a")) {
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    v29 := GenS2{F0: 42, F1: 128}
                                    if false {
                                        switch "\n\t" {
                                        case "a":
                                            var v30 uint = 602
                                            _ = v30
                                        default:
                                            if ((1e10 / (3.14159 * (((0.1 / 0.1) + (float32(v2) - 1.5)) * -2.25))) > (1e10 / ((3.14159 + 1e10) + 1e10))) {
                                                v3 += v3
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v31 int16 = -702
                                                        _ = v31
                                                    case "hello":
                                                        if true {
                                                            v32 := GenS5{F0: (v3 << ((65535) & 63)), F1: 255, F2: v1}
                                                            _ = v32
                                                        } else {
                                                            var v33 string = ("hello" + "a")
                                                            v34 := GenS5{F0: v3, F1: 228, F2: v1}
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    }
                                                case "":
                                                    var v35 uint32 = 330
                                                    _ = v35
                                                case "hello":
                                                    switch ("foo bar" + "a") {
                                                    case "foo bar":
                                                        var v36 int8 = -16
                                                        fmt.Println(v2, v3)
                                                        _ = v36
                                                    case "hello":
                                                        v2 -= float64((0.1 / (float32(v2) / (1e10 - float32(0.0)))))
                                                    case "":
                                                        v37 := GenS6{F0: 2, F1: v4}
                                                        _ = v37
                                                    }
                                                default:
                                                    v4 = -403
                                                }
                                            } else {
                                                v4 = (v4 | (^(v4)))
                                                var v38 int16 = v1
                                                _ = v38
                                            }
                                        }
                                        v2 += 1.5
                                    } else {
                                        if false {
                                            var v39 int64 = (v4 | v4)
                                            if true {
                                                v4 = v39
                                                var v40 float32 = ((3.14159 + -2.25) * (3.14159 - -2.25))
                                                var v41 int8 = -128
                                                _ = v40
                                                _ = v41
                                            } else {
                                                v42 := []int{(^(fi15)), (fi15 & fi28)}
                                                var v43 int64 = (v39 ^ v39)
                                                _ = v42
                                                _ = v43
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                v44 := GenS3{F0: 42, F1: 1, F2: false}
                                                _ = v44
                                            default:
                                                fmt.Println(fi28)
                                            }
                                            var v45 string = ("hello" + ("\n\t" + "a"))
                                            var v46 int16 = 42
                                            _ = v39
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v47 := GenS7{F0: 411, F1: "a"}
                                            _ = v47
                                        }
                                        var v48 uint = 255
                                        switch fi15 {
                                        case 9223372036854775807:
                                            var v49 int32 = 2147483647
                                            _ = v49
                                        default:
                                            var v50 int16 = (v1 ^ int16(fi28))
                                            _ = v50
                                        }
                                        _ = v48
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                fmt.Println(v1)
                            } else {
                                if false {
                                    v4 = v4
                                    var v51 int = 42
                                    var v52 int32 = 2147483647
                                    _ = v51
                                    _ = v52
                                } else {
                                    fmt.Println(v1)
                                }
                                switch 639 {
                                case 1:
                                    var v53 uint = 794
                                    _ = v53
                                default:
                                    if (0 > -1) {
                                        v1 -= v1
                                    } else {
                                        v54 := GenS7{F0: 42, F1: "a"}
                                        v55 := GenS2{F0: -310, F1: genFn0((^(v4)))}
                                        switch fi15 {
                                        case 4294967295:
                                            var v56 uint32 = 591
                                            _ = v56
                                        case -551:
                                            fmt.Println(v4)
                                        case 128:
                                            for fi57 := 0; fi57 < 13; fi57++ {
                                                v58 := GenS8{F0: -1, F1: 862}
                                                var v59 string = ""
                                                fmt.Println(v1, v1, fi15)
                                                var v60 int = int(v3)
                                                v61 := []int{(v60 >> ((18446744073709551615) & 63)), 42, -9223372036854775808}
                                                _ = fi57
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            }
                                        default:
                                            v4 += v4
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                for fi62 := 0; fi62 < 2; fi62++ {
                                    var v63 uint16 = v3
                                    v64 := GenS6{F0: 65535, F1: v4}
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                            }
                            var v65 int32 = 0
                            _ = v65
                        } else {
                            v66 := []int{fi15, (fi15 / ((fi15) | 1)), (fi15 + 346), (fi15 | fi15)}
                            if len(v66) >= 2 {
                                v66 = v66[1:len(v66)]
                            }
                            _ = v66
                        }
                        v2 = float64(0.1)
                        v1 -= 0
                    } else {
                        fmt.Println(v2)
                        v2 = v2
                    }
                    _ = v27
                }
            }
        default:
            v2 -= (v2 * 1.5)
        }
        v67 := GenS9{F0: fi15, F1: (^(v3)), F2: 96}
        v68 := GenS1{F0: (v1 ^ (v1 - v1)), F1: 79}
        _ = fi15
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 11; fi69++ {
        var v70 int32 = -1
        if (v1 >= (^(v1))) {
            fmt.Println(v70, v4)
            v71 := []int{fi69}
            _ = v71
        }
        v72 := []int{-9223372036854775808}
        _ = fi69
        _ = v70
        _ = v72
    }
    var v73 float32 = float32((0.1 - float64((1e10 * (-2.25 / float32(float64(0.1)))))))
    var v74 uint64 = 34
    v74 += (^(v74))
    var v75 float64 = v2
    var v76 uint = 821
    v77 := []int{}
    if (!(v3 != v3)) {
        v1 -= int16(v3)
    } else {
        var v78 uint32 = 2147483647
        var v79 float32 = (v73 - float32((0.0 / ((v75 / float64(float32(v2))) + (1e10 + (0.0 / (v2 / float64(float32((3.14159 - ((v2 / v75) - (float64(((((v73 + v73) / (v73 + -2.25)) / v73) + 1e10)) * v75))))))))))))
        _ = v78
        _ = v79
    }
    fmt.Println(v76)
    v76 += uint(v74)
    fmt.Println(v1, v75)
    for fi80 := 0; fi80 < 12; fi80++ {
        for fi81 := 0; fi81 < 6; fi81++ {
            if true {
                fmt.Println(fi80)
            } else {
                if (1 > v3) {
                    v73 = ((((v73 / float32((((float64(1.5) + (v2 + v2)) + v2) * 3.14159))) * (float32(((-2.25 + (((v2 / v2) / v75) - v75)) + 3.14159)) + -2.25)) + ((((-2.25 + v73) / -2.25) - -2.25) - 1.5)) * -2.25)
                } else {
                    var v82 int64 = int64(v74)
                    _ = len(v77)
                    var v83 string = ("a" + "\n\t")
                    _ = v82
                    _ = v83
                }
                var v84 int32 = 0
                switch (-(fi80)) {
                case 649:
                    v74 = 42
                case 0:
                    v1 += (^(v1))
                case 32768:
                    v85 := GenS3{F0: (v84 & -1), F1: 4294967295, F2: true}
                    var v86 bool = true
                    var v87 string = ""
                    var v88 float32 = (v73 - ((v73 / (((3.14159 - v73) + ((float32(v2) + v73) - (v73 - v73))) + 0.1)) / v73))
                    v89 := GenS6{F0: 42, F1: (v4 >> ((v76) & 63))}
                    v90 := GenS1{F0: v1, F1: 2}
                    _ = v85
                    _ = v86
                    _ = v87
                    _ = v88
                    _ = v89
                    _ = v90
                }
                _ = v84
            }
            _ = fi81
        }
        fmt.Println(fi80)
        v74 = v74
        _ = fi80
    }
    _ = len(v77)
    if ((^(v74)) <= (v74 - v74)) {
        if len(v77) > 0 {
            v77[0] = 32767
        }
        fmt.Println(v73, v4)
    } else {
        if (-80 > -2147483648) {
            if (224 < (v1 % ((32767) | 1))) {
                v76 = v76
                for fi91 := 0; fi91 < 14; fi91++ {
                    fmt.Println(v76)
                    _ = fi91
                }
                var v92 uint32 = 65535
                _ = v92
            }
        } else {
            for fi93 := 0; fi93 < 1; fi93++ {
                var v94 bool = ((("\n\t" + "\n\t") + "\n\t") != "a")
                for fi95 := 0; fi95 < 18; fi95++ {
                    v74 = uint64(v3)
                    fmt.Println(v1, v94, v75)
                    _ = fi95
                }
                switch (fi93 & (fi93 % (((fi93 / (((fi93 >> (((v76 / (((v76 | 42)) | 1))) & 63))) | 1))) | 1))) {
                case 42:
                    var v96 int8 = int8(v74)
                    _ = v96
                default:
                    v97 := GenS10{F0: (fi93 % ((fi93) | 1)), F1: "a", F2: 367}
                    _ = v97
                }
                _ = fi93
                _ = v94
            }
            v73 = (3.14159 / v73)
            v98 := GenS8{F0: 1, F1: (v3 - (v3 + v3))}
            _ = v98
        }
        var v99 int = -582
        v2 = v2
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("slicelen v77:", len(v77), v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
