package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 uint64) uint16 {
    v6 := []int{0}
    fmt.Println(p0)
    var v7 float32 = 3.14159
    v7 = (float32(float64(v7)) * -2.25)
    _ = p0
    _ = v6
    _ = v7
    return 255
}
func genFn1() string {
    v8 := GenS1{F0: false, F1: ""}
    var v9 int16 = 1000
    v10 := []int{-1, -719}
    if false {
        var v11 int32 = -87
        var v12 float64 = 3.14159
        var v13 uint32 = 4294967295
        _ = v11
        _ = v12
        _ = v13
    } else {
        fmt.Println(v9)
    }
    _ = v10
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 79
    var v3 uint16 = 32767
    var v4 int64 = -1
    var v5 uint32 = 4294967295
    if false {
        v14 := GenS1{F0: false, F1: "a"}
        v15 := GenS2{F0: genFn0(994), F1: 42}
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: true, F1: "hello"}
        v17 := []int{32767}
        if (1e10 < 0.1) {
            v3 = (v3 % ((genFn0(42)) | 1))
        } else {
            var v18 uint32 = (v5 | (v5 * v5))
            if len(v17) > 0 {
                v17[0] = -383
            }
            _ = v18
        }
        _ = v16
        _ = v17
    }
    v1 = 499
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int8 = 24
        if true {
            v21 := GenS2{F0: v3, F1: v20}
            if (0.0 != 0.1) {
                v22 := []int{9223372036854775807, fi19, int(v2)}
                for fi23 := 0; fi23 < 1; fi23++ {
                    if ((4294967295 >= 875) && true) {
                        switch "foo bar" {
                        case "hello":
                            if len(v22) > 0 {
                                v22[0] = (fi23 * fi19)
                            }
                        case "\n\t":
                            v5 -= 0
                        case "foo bar":
                            switch -9223372036854775808 {
                            case 1:
                                if false {
                                    v22 = append(v22, 32768)
                                    v1 = (v1 >> ((255) & 63))
                                    fmt.Println(v3, v2)
                                } else {
                                    fmt.Println(fi23, v4)
                                    v24 := GenS3{F0: false, F1: false, F2: v5}
                                    if true {
                                        v25 := GenS3{F0: true, F1: false, F2: (v5 << ((4294967295) & 63))}
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            case 42:
                                v22 = append(v22, (fi23 | fi23))
                            default:
                                v5 -= 42
                            }
                        }
                        fmt.Println(v4, v4)
                        var v26 int16 = -264
                        v27 := GenS4{F0: 42, F1: v3, F2: (^(v26))}
                        _ = v26
                        _ = v27
                    }
                    v1 = (v1 | v1)
                    _ = fi23
                }
                _ = v22
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v4 = v4
                    var v29 uint8 = (v2 << ((42) & 63))
                    _ = fi28
                    _ = v29
                }
            }
            var v30 int64 = v4
            _ = v21
            _ = v30
        } else {
            for fi31 := 0; fi31 < 12; fi31++ {
                v32 := GenS4{F0: 988, F1: genFn0(42), F2: 1}
                if (0 > 2) {
                    var v33 float64 = 3.14159
                    if false {
                        v34 := GenS5{F0: 556, F1: (fi31 - (^(fi19)))}
                        var v35 int16 = 42
                        v1 += int32(v20)
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                } else {
                    fmt.Println(v1, v2, fi31)
                    if true {
                        if (int16(v5) != 42) {
                            v20 -= 42
                        } else {
                            v3 += v3
                            var v36 int = (fi19 & fi31)
                            v37 := []int{int(v4), fi19}
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS3{F0: true, F1: false, F2: v5}
                        if (true && false) {
                            switch (fi19 * -2147483648) {
                            case 42:
                                var v39 uint64 = 1
                                _ = v39
                            case 155:
                                v1 = (v1 - v1)
                            case 49:
                                v40 := GenS3{F0: ((v2 >> ((0) & 63)) >= v2), F1: false, F2: 42}
                                _ = v40
                            default:
                                v2 += v2
                            }
                            fmt.Println(v1, fi31)
                        }
                        _ = v38
                    } else {
                        var v41 int16 = 42
                        var v42 uint64 = 42
                        fmt.Println(v20, v20)
                        _ = v41
                        _ = v42
                    }
                    v20 = (v20 << ((551) & 63))
                }
                v3 -= (v3 % ((65535) | 1))
                _ = fi31
                _ = v32
            }
        }
        if (0.1 < 0.1) {
            if true {
                switch "\n\t" {
                case "":
                    var v43 uint8 = (v2 - v2)
                    _ = v43
                case "hello":
                    v2 -= (v2 & v2)
                case "\n\t":
                    v44 := []int{fi19, 127, (fi19 / ((fi19) | 1)), 0}
                    _ = v44
                default:
                    for fi45 := 0; fi45 < 13; fi45++ {
                        var v46 float32 = -2.25
                        var v47 int16 = 32767
                        var v48 int8 = 11
                        _ = fi45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                }
                for fi49 := 0; fi49 < 10; fi49++ {
                    v5 = uint32(v2)
                    _ = fi49
                }
                for fi50 := 0; fi50 < 4; fi50++ {
                    v51 := GenS3{F0: (uint32(v2) == 317), F1: (!(int64(v5) >= (v4 ^ (v4 >> ((255) & 63))))), F2: (v5 - 42)}
                    v2 -= v2
                    v52 := []int{fi50, (-(fi19)), fi50, (fi50 ^ 2147483647)}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            }
        }
        _ = fi19
        _ = v20
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        for fi54 := 0; fi54 < 9; fi54++ {
            if true {
                var v55 uint8 = (v2 * 0)
                v3 += (v3 * uint16(v5))
                v56 := []int{fi54, 45, (fi54 ^ fi53)}
                v57 := GenS6{F0: (("hello" + "\n\t") + ("" + "a"))}
                _ = v55
                _ = v56
                _ = v57
            }
            _ = fi54
        }
        _ = fi53
    }
    if false {
        switch 56 {
        case 128:
            var v58 int8 = -128
            _ = v58
        case 0:
            switch "\n\t" {
            case "a":
                var v59 uint32 = v5
                _ = v59
            case "foo bar":
                if true {
                    v60 := GenS7{F0: ((!(4294967295 < 18446744073709551615)) || (v1 > (-(v1)))), F1: ((("foo bar" + "") + "hello") + "a"), F2: 18446744073709551615}
                    _ = v60
                }
            }
        case 32767:
            v61 := []int{}
            _ = v61
        }
        var v62 int = -862
        v63 := []int{(v62 | 0), v62, (v62 << ((820) & 63))}
        _ = v62
        _ = v63
    }
    var v64 uint32 = 1
    if (false || (-9223372036854775808 > 9223372036854775807)) {
        v64 += v5
        if true {
            v4 -= v4
            var v65 uint8 = (v2 & 81)
            _ = v65
        }
    }
    fmt.Println(v64)
    var v66 int8 = 1
    var v67 uint16 = (v3 / ((v3) | 1))
    for fi68 := 0; fi68 < 13; fi68++ {
        if (-32768 > 42) {
            switch (fi68 + fi68) {
            case 0:
                var v69 bool = false
                _ = v69
            case 969:
                v4 += (v4 - 42)
            case -572:
                if false {
                    v70 := []int{32767, (fi68 / (((fi68 | 2147483647)) | 1)), -423, fi68}
                    v71 := GenS4{F0: 0, F1: (v3 / ((genFn0(18446744073709551615)) | 1)), F2: 32767}
                    _ = v70
                    _ = v71
                }
            }
            for fi72 := 0; fi72 < 14; fi72++ {
                v73 := GenS7{F0: true, F1: ("" + ""), F2: 4294967295}
                var v74 float64 = (1e10 / (float64(1.5) + 0.1))
                fmt.Println(fi72, fi72, v67)
                v4 -= (v4 | (v4 & v4))
                _ = fi72
                _ = v73
                _ = v74
            }
            v4 -= 872
        }
        _ = fi68
    }
    var v75 int16 = -32768
    var v76 int32 = v1
    v5 = v64
    var v77 int16 = (v75 / ((v75) | 1))
    v67 += v3
    v78 := []int{47, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v66
    _ = v67
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
