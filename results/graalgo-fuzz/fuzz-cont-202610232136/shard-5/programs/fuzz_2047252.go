package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint; F2 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint; F1 int32; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 string; F1 int }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 string; F1 int8; F2 uint8 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8; F1 int64; F2 int16 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64; F1 uint64 }
type GenS17 struct { F0 int32; F1 int64; F2 int32 }
type GenS18 struct { F0 uint32; F1 int8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int8 }
func genFn0() uint16 {
    v5 := []int{}
    v6 := GenS1{F0: ("foo bar" + "hello")}
    _ = v5
    _ = v6
    return 638
}
func genFn1(p0 bool, p1 uint) int8 {
    v7 := GenS1{F0: "a"}
    if (42 >= 663) {
        for fi8 := 0; fi8 < 8; fi8++ {
            if ((fi8 - fi8) != fi8) {
                p1 -= 4294967295
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: ""}
                    p0 = (1e10 != float32((1.5 + ((1e10 * -2.25) + 1e10))))
                    if (3.14159 == (0.0 + 0.1)) {
                        v11 := GenS2{F0: 203, F1: (("foo bar" + "hello") + ("" + ""))}
                        var v12 int16 = 32767
                        var v13 uint8 = 255
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch (-(fi8)) {
                    case -166:
                        if ("a" > "foo bar") {
                            var v14 uint16 = genFn0()
                            if p0 {
                                v15 := []int{}
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                switch ("\n\t" + "") {
                                case "":
                                    v16 := GenS3{F0: 0}
                                    _ = v16
                                }
                                _ = v15
                            }
                            var v17 float32 = 3.14159
                            _ = v14
                            _ = v17
                        }
                    case -796:
                        var v18 int8 = -82
                        _ = v18
                    }
                    var v19 int = -9223372036854775808
                    _ = fi9
                    _ = v10
                    _ = v19
                }
                v20 := GenS3{F0: 1}
                _ = v20
            } else {
                v21 := GenS4{F0: 42, F1: (^(p1)), F2: -299}
                var v22 bool = (!p0)
                _ = v21
                _ = v22
            }
            var v23 uint32 = 458
            _ = fi8
            _ = v23
        }
        v24 := []int{716, 649}
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(p1, p0, fi25)
        switch -9223372036854775808 {
        case 127:
            var v26 uint8 = 128
            _ = v26
        default:
            if (("hello" + "a") > "\n\t") {
                switch -9223372036854775808 {
                case 158:
                    fmt.Println(fi25, p0, fi25)
                case 128:
                    v27 := GenS5{F0: 578, F1: -128, F2: 18446744073709551615}
                    _ = v27
                case -2147483648:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        switch (fi28 - fi25) {
                        case 42:
                            if (p0 || true) {
                                if (!p0) {
                                    p0 = false
                                    v29 := GenS6{F0: (p1 + 1), F1: -2147483648, F2: false}
                                    _ = v29
                                } else {
                                    var v30 uint32 = 595
                                    _ = v30
                                }
                                var v31 uint8 = 2
                                _ = v31
                            }
                        case 2147483647:
                            fmt.Println(fi25, p0, p1)
                        case -739:
                            v32 := []int{fi25, fi28}
                            _ = v32
                        }
                        fmt.Println(fi28, p0)
                        _ = fi28
                    }
                default:
                    var v33 float32 = (float32(3.14159) - 3.14159)
                    _ = v33
                }
            } else {
                p1 = (p1 - (^(p1)))
            }
        }
        v34 := GenS7{F0: 1}
        _ = fi25
        _ = v34
    }
    p0 = (297 <= uint32(p1))
    _ = p0
    _ = p1
    _ = v7
    return 107
}

func main() {
    var v1 bool = (65535 >= 255)
    v2 := []int{168, 0, 32768, 2147483647}
    var v3 uint64 = 592
    v4 := make(map[uint32]bool)
    v1 = ((0 <= 32) && (v3 != v3))
    var v35 string = "\n\t"
    var v36 int16 = 1
    fmt.Println(v36)
    for fi37 := 0; fi37 < 19; fi37++ {
        for fi38 := 0; fi38 < 19; fi38++ {
            if false {
                var v39 int8 = 127
                v40 := v4[2]
                if v1 {
                    v39 = v39
                    v41 := GenS2{F0: 255, F1: v35}
                    var v42 float64 = float64(1e10)
                    v43 := GenS8{F0: (v1 && true), F1: 1}
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    switch fi38 {
                    case 127:
                        if v1 {
                            var v44 float32 = 0.1
                            v35 = ""
                            _ = v44
                        } else {
                            v45 := GenS9{F0: v35, F1: int(v3)}
                            var v46 int16 = v36
                            v1 = v1
                            _ = v45
                            _ = v46
                        }
                    case 9223372036854775807:
                        if v1 {
                            fmt.Println(fi37, v1)
                            var v47 int = (^(fi37))
                            _ = v47
                        }
                    case 0:
                        if (255 >= 42) {
                            v36 += (v36 / (((v36 / ((32767) | 1))) | 1))
                            if (float64((1.5 * 1.5)) == -2.25) {
                                v40 = v1
                                v48 := GenS10{F0: (v36 % (((v36 / ((-32768) | 1))) | 1)), F1: 42}
                                var v49 int = 0
                                _ = v48
                                _ = v49
                            }
                        }
                    default:
                        if (!(32767 > 42)) {
                            var v50 uint32 = 2147483648
                            _ = v50
                        } else {
                            var v51 int = fi37
                            _ = v51
                        }
                    }
                    if (0 != 371) {
                        v40 = v1
                    } else {
                        v52 := GenS11{F0: v35, F1: (v39 >> ((105) & 63)), F2: 42}
                        var v53 int = (fi38 / (((fi37 ^ 9223372036854775807)) | 1))
                        fmt.Println(v35, v53)
                        _ = v52
                        _ = v53
                    }
                    v54 := GenS12{F0: 822, F1: genFn1(false, 0)}
                    if (((-(v39)) > genFn1(v40, 65535)) || ((v36 <= 0) || true)) {
                        switch fi38 {
                        case 32767:
                            v36 += (v36 * int16(v3))
                        case 2147483647:
                            v35 = ((v35 + ("a" + v35)) + "\n\t")
                        case -1:
                            v35 = "\n\t"
                        default:
                            var v55 uint = 853
                            _ = v55
                        }
                        var v56 uint64 = v3
                        _ = v56
                    } else {
                        if (742 <= 18446744073709551615) {
                            var v57 uint32 = 2
                            _ = v57
                        }
                    }
                    _ = v54
                }
                _ = v39
                _ = v40
            } else {
                var v58 uint8 = 98
                switch fi37 {
                case -288:
                    if (!v1) {
                        v58 += 128
                        v35 = v35
                    } else {
                        var v59 float32 = 3.14159
                        var v60 uint16 = 42
                        _ = v59
                        _ = v60
                    }
                case 32767:
                    fmt.Println(fi37, v36, v58)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                fmt.Println(v58, v35)
                _ = v58
            }
            fmt.Println(v35, v35, v1)
            v3 -= (v3 / (((v3 | (v3 + 468))) | 1))
            _ = fi38
        }
        var v61 uint32 = 42
        _ = fi37
        _ = v61
    }
    v36 = (v36 / (((-(v36))) | 1))
    if v1 {
        v62 := GenS1{F0: v35}
        fmt.Println(v1)
        _ = v62
    } else {
        var v63 float32 = float32(3.14159)
        _ = v63
    }
    if v1 {
        var v64 int64 = -950
        if ((12 > 990) || (true || v1)) {
            switch 32768 {
            case -1:
                v65 := GenS13{F0: v3}
                _ = v65
            }
        }
        var v66 float32 = 1.5
        _ = v64
        _ = v66
    } else {
        var v67 uint32 = 2147483648
        if (183 == 128) {
            v67 += 2147483647
            v68 := GenS14{F0: 70, F1: 377, F2: (v36 % (((v36 & 0)) | 1))}
            for fi69 := 0; fi69 < 5; fi69++ {
                v70 := GenS15{F0: int(v67)}
                _ = fi69
                _ = v70
            }
            _ = v68
        }
        if ((-2.25 * 3.14159) > 0.1) {
            var v71 int = 9223372036854775807
            var v72 uint32 = v67
            fmt.Println(v72, v67, v71)
            _ = v71
            _ = v72
        } else {
            for fi73 := 0; fi73 < 3; fi73++ {
                var v74 uint16 = 32767
                for fi75 := 0; fi75 < 18; fi75++ {
                    var v76 uint32 = (v67 - (v67 & (v67 % ((v67) | 1))))
                    v77 := GenS9{F0: v35, F1: (^(fi75))}
                    v78 := GenS5{F0: -182, F1: -41, F2: v3}
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                }
                _ = fi73
                _ = v74
            }
            v36 -= v36
            v79 := GenS16{F0: -9223372036854775808, F1: v3}
            _ = v79
        }
        _ = v67
    }
    fmt.Println(v1, v1)
    v80 := GenS1{F0: (v35 + "a")}
    for fi81 := 0; fi81 < 11; fi81++ {
        for fi82 := 0; fi82 < 6; fi82++ {
            v83 := GenS17{F0: -1, F1: -1, F2: 42}
            _ = fi82
            _ = v83
        }
        fmt.Println(v3, v3)
        _ = fi81
    }
    v35 = ((v35 + ("\n\t" + v35)) + "hello")
    v84 := GenS18{F0: 2, F1: -102}
    var v85 float32 = (1e10 + 3.14159)
    switch v35 {
    case "":
        var v86 uint = uint(v3)
        _ = v86
    case "foo bar":
        for fi87 := 0; fi87 < 8; fi87++ {
            v88 := GenS2{F0: 128, F1: v35}
            for fi89 := 0; fi89 < 13; fi89++ {
                v90 := GenS19{F0: 100}
                _ = fi89
                _ = v90
            }
            _ = fi87
            _ = v88
        }
    default:
        var v91 float64 = float64((0.0 * v85))
        _ = v91
    }
    if (2 != 0) {
        fmt.Println(v85, v1)
        v92 := GenS20{F0: 42}
        _ = v92
    }
    for fi93 := 0; fi93 < 17; fi93++ {
        v94 := GenS21{F0: 1}
        _ = fi93
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v80
    _ = v84
    _ = v85
}
