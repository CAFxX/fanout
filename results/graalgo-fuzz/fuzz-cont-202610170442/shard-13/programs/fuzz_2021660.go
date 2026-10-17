package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 uint8, p1 int64, p2 int16) uint64 {
    v7 := []int{}
    v8 := GenS1{F0: 127}
    v9 := GenS1{F0: int8(p2)}
    var v10 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 738
}
func genFn1(p0 int32, p1 int) string {
    switch ("a" + ("foo bar" + ("foo bar" + ("\n\t" + "\n\t")))) {
    case "foo bar":
        var v11 uint = 2
        _ = v11
    case "hello":
        p1 = (p1 + (p1 % (((p1 % (((^(p1))) | 1))) | 1)))
    default:
        switch "" {
        case "\n\t":
            v12 := []int{(p1 * -9223372036854775808), p1}
            _ = v12
        case "foo bar":
            if false {
                var v13 uint16 = 2
                v14 := []int{(^(p1))}
                p0 += p0
                _ = v13
                _ = v14
            } else {
                fmt.Println(p1)
            }
        case "":
            var v15 string = "a"
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
    }
    if (42 < 314) {
        switch (-(p1)) {
        case 32767:
            if ((false && true) || false) {
                if (-9223372036854775808 > 0) {
                    v17 := GenS2{F0: -621}
                    _ = v17
                } else {
                    p1 += 32768
                    var v18 int64 = -971
                    v18 += v18
                    _ = v18
                }
                if false {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS3{F0: 42, F1: 18446744073709551615, F2: ((3.14159 + -2.25) != 1.5)}
                            p1 -= 4294967295
                            switch (fi19 | (p1 * fi19)) {
                            case 32768:
                                if ((3.14159 * 1e10) < 1.5) {
                                    p1 += (-(fi19))
                                    var v22 uint64 = 4294967295
                                    if false {
                                        v22 += v22
                                        fmt.Println(fi20, p0, v22)
                                    } else {
                                        switch 32767 {
                                        case -216:
                                            if ((false && (((!false) && true) && true)) || false) {
                                                if (42 >= 1) {
                                                    var v23 int32 = p0
                                                    var v24 uint32 = 2147483648
                                                    v25 := GenS4{F0: -9223372036854775808, F1: 193, F2: v22}
                                                    _ = v23
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v22 = (^(v22))
                                                p1 = 4294967295
                                            } else {
                                                v26 := GenS2{F0: 32767}
                                                v27 := GenS1{F0: 100}
                                                _ = v26
                                                _ = v27
                                            }
                                        case 4294967295:
                                            switch (p1 + (^(fi19))) {
                                            case 4294967295:
                                                if true {
                                                    fmt.Println(p1)
                                                } else {
                                                    var v28 float32 = 3.14159
                                                    v29 := GenS5{F0: (false || false)}
                                                    _ = v28
                                                    _ = v29
                                                }
                                            default:
                                                switch (fi20 | -201) {
                                                case 1:
                                                    v30 := []int{32767, fi19}
                                                    _ = v30
                                                case 286:
                                                    v31 := []int{p1, p1, (p1 % ((fi19) | 1)), (fi20 ^ fi20)}
                                                    if true {
                                                        if true {
                                                            if true {
                                                                var v32 uint64 = v22
                                                                if (!(0.0 > -2.25)) {
                                                                    v33 := GenS4{F0: -702, F1: 1, F2: (v22 | (v22 % ((v32) | 1)))}
                                                                    if true {
                                                                        v32 -= v32
                                                                    } else {
                                                                        var v34 uint32 = 295
                                                                        var v35 uint16 = 2
                                                                        v32 = (^(v32))
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                    v36 := GenS4{F0: 9223372036854775807, F1: 0, F2: (v22 >> ((0) & 63))}
                                                                    _ = v33
                                                                    _ = v36
                                                                }
                                                                v22 = v32
                                                                _ = v32
                                                            }
                                                        }
                                                        fmt.Println(v22)
                                                    }
                                                    v37 := GenS4{F0: 229, F1: 127, F2: (v22 & (v22 / (((v22 / ((v22) | 1))) | 1)))}
                                                    var v38 int64 = 42
                                                    fmt.Println(v38, fi20)
                                                    if true {
                                                        v39 := GenS4{F0: (v38 % ((v38) | 1)), F1: 1, F2: genFn0(226, (v38 / ((int64(fi19)) | 1)), 32767)}
                                                        var v40 uint = 0
                                                        v41 := GenS4{F0: 42, F1: uint8(v22), F2: v22}
                                                        var v42 int64 = (v38 >> ((v40) & 63))
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    _ = v31
                                                    _ = v37
                                                    _ = v38
                                                default:
                                                    var v43 string = (("a" + "hello") + "\n\t")
                                                    _ = v43
                                                }
                                            }
                                        default:
                                            var v44 int64 = 799
                                            _ = v44
                                        }
                                    }
                                    _ = v22
                                } else {
                                    var v45 uint = 275
                                    _ = v45
                                }
                            case 42:
                                switch "\n\t" {
                                case "\n\t":
                                    p0 = (^(p0))
                                    v46 := GenS6{F0: 46, F1: 389}
                                    v47 := []int{p1, (fi19 & p1), 401}
                                    _ = v46
                                    _ = v47
                                }
                            case 128:
                                var v48 bool = true
                                _ = v48
                            default:
                                v49 := GenS3{F0: 2, F1: 362, F2: true}
                                _ = v49
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = fi19
                    }
                } else {
                    fmt.Println(p1, p0)
                    var v50 uint64 = 1
                    fmt.Println(v50, p1)
                    _ = v50
                }
            }
        case -563:
            switch "foo bar" {
            case "\n\t":
                if (p1 >= 987) {
                    for fi51 := 0; fi51 < 8; fi51++ {
                        v52 := []int{32768}
                        p1 = p1
                        var v53 float64 = (3.14159 + 0.1)
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 uint64 = 903
                    var v55 string = ""
                    _ = v54
                    _ = v55
                } else {
                    p0 += (-(p0))
                }
            }
        case 662:
            for fi56 := 0; fi56 < 9; fi56++ {
                if (0 >= 18446744073709551615) {
                    var v57 int16 = 42
                    for fi58 := 0; fi58 < 9; fi58++ {
                        v59 := []int{fi58, -341, 128, (^(p1))}
                        _ = fi58
                        _ = v59
                    }
                    _ = v57
                }
                p0 = int32(fi56)
                switch "\n\t" {
                case "foo bar":
                    var v60 uint8 = 1
                    _ = v60
                default:
                    var v61 string = "\n\t"
                    _ = v61
                }
                _ = fi56
            }
        default:
            p1 += p1
        }
        if true {
            if false {
                switch p1 {
                case -989:
                    v62 := []int{}
                    _ = v62
                case 4294967295:
                    if ((p0 << ((18446744073709551615) & 63)) > -2147483648) {
                        var v63 uint8 = 42
                        _ = v63
                    }
                default:
                    var v64 bool = false
                    _ = v64
                }
            }
        }
        fmt.Println(p0, p1)
    } else {
        p1 -= p1
        switch (p1 + (p1 / (((p1 << ((780) & 63))) | 1))) {
        case 1:
            p1 += 797
        case -850:
            var v65 int32 = (p0 - -1)
            _ = v65
        case -1:
            p1 += (p1 % (((p1 + (p1 - (p1 / (((p1 & p1)) | 1))))) | 1))
        default:
            for fi66 := 0; fi66 < 11; fi66++ {
                if ((p0 * p0) < (p0 & 42)) {
                    p1 = (p1 << ((4294967295) & 63))
                } else {
                    v67 := []int{(p1 % (((p1 >> ((255) & 63))) | 1)), -1, int(p0), p1}
                    if len(v67) > 0 {
                        v67[0] = -449
                    }
                    _ = v67
                }
                if (0 >= 2147483648) {
                    v68 := []int{(fi66 - 576), (p1 + p1)}
                    if len(v68) > 0 {
                        v68[0] = 4294967295
                    }
                    p0 += p0
                    _ = v68
                } else {
                    p0 -= (-(p0))
                    v69 := GenS7{F0: fi66}
                    for fi70 := 0; fi70 < 20; fi70++ {
                        p1 += (p1 & fi70)
                        var v71 float64 = 3.14159
                        _ = fi70
                        _ = v71
                    }
                    _ = v69
                }
                _ = fi66
            }
        }
    }
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{0, 0, 166, 2147483648}
    v3 := []uint32{1}
    var v4 int16 = -1
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    v1 = v1
    var v72 int = 0
    if ((!false) && true) {
        v1 = (v1 | -972)
        switch (-(v72)) {
        case 128:
            v4 = (-(v4))
        }
    } else {
        var v73 int8 = 6
        _ = v73
    }
    var v74 bool = (0.0 < float32((3.14159 / (0.0 - -2.25))))
    _ = len(v3)
    if true {
        if true {
            v4 = v4
            fmt.Println(v72, v4)
            if true {
                for fi75 := 0; fi75 < 15; fi75++ {
                    _ = len(v3)
                    v1 -= v1
                    _ = len(v3)
                    _ = fi75
                }
                v74 = ((("foo bar" + ("\n\t" + "")) != "foo bar") || ((false && (3.14159 != (1.5 + 1e10))) && ("hello" != ("" + "foo bar"))))
                if false {
                    v76 := GenS3{F0: 0, F1: 2, F2: ((!(v74 || (-9223372036854775808 < -116))) && true)}
                    if (true && v74) {
                        _ = len(v6)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if true {
                        v72 = 1
                        var v77 uint64 = 4294967295
                        if v74 {
                            fmt.Println(v72)
                            v78 := GenS8{F0: 0, F1: v74}
                            for fi79 := 0; fi79 < 3; fi79++ {
                                fmt.Println(v74, v4)
                                _ = fi79
                            }
                            _ = v78
                        }
                        _ = v77
                    } else {
                        for fi80 := 0; fi80 < 4; fi80++ {
                            v81 := GenS9{F0: v1, F1: -836, F2: 0}
                            _ = fi80
                            _ = v81
                        }
                        var v82 uint16 = 32767
                        _ = v82
                    }
                    _ = v76
                }
            }
        }
    }
    v1 = -289
    v83 := GenS3{F0: genFn0(0, 9223372036854775807, (v4 >> ((0) & 63))), F1: 0, F2: v74}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v74
    _ = v83
}
