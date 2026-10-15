package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int8; F2 string }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 int }
func genFn0(p0 string, p1 string, p2 int16) int64 {
    v5 := GenS1{F0: 127, F1: 255}
    switch 32767 {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int = fi6
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 bool = (643 < 0)
                p2 -= (p2 | p2)
                var v10 uint8 = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case 127:
        v11 := []int{-9223372036854775808, -334}
        _ = v11
    case 32767:
        var v12 int = 127
        _ = v12
    }
    p1 = ("hello" + "hello")
    v13 := []int{42, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return -469
}
func genFn1() int16 {
    var v14 int = -448
    fmt.Println(v14)
    _ = v14
    return 0
}

func main() {
    var v1 int64 = 42
    var v2 int16 = -32768
    v3 := make(map[int]int)
    var v4 int16 = (v2 & 42)
    v4 += v4
    delete(v3, 719)
    v15 := GenS2{F0: 2, F1: 0}
    fmt.Println(v4, v1, v4)
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            if true {
                if ((!true) || true) {
                    switch fi16 {
                    case -25:
                        v1 -= (^(v1))
                    }
                } else {
                    if true {
                        v18 := []int{fi17, (fi16 - 4294967295)}
                        var v19 uint8 = 128
                        _ = v18
                        _ = v19
                    }
                    var v20 int8 = 1
                    _ = v20
                }
            } else {
                v21 := GenS3{F0: 0}
                v2 += (v2 * (v2 & v4))
                fmt.Println(v4, v2, fi17)
                _ = v21
            }
            var v22 bool = ((true || ((false && true) && (false && false))) && false)
            v2 = int16(v1)
            _ = fi17
            _ = v22
        }
        _ = fi16
    }
    v23 := GenS1{F0: 32768, F1: 0}
    if false {
        v24 := GenS3{F0: 559}
        fmt.Println(v2, v4, v4)
        v2 += v2
        _ = v24
    } else {
        v25 := GenS3{F0: 4294967295}
        if (false || true) {
            if (40 == 775) {
                var v26 bool = true
                fmt.Println(v26)
                v1 += -1
                _ = v26
            } else {
                v27 := GenS4{F0: v4, F1: 1, F2: ((("hello" + "a") + ((("foo bar" + "foo bar") + "hello") + ("\n\t" + (("hello" + "\n\t") + (((("foo bar" + "\n\t") + (((("hello" + "") + "a") + "") + ("hello" + ""))) + "a") + "\n\t"))))) + "")}
                v1 += (v1 >> ((18446744073709551615) & 63))
                fmt.Println(v1, v4, v1)
                _ = v27
            }
        } else {
            fmt.Println(v1, v1, v4)
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 uint32 = 596
                if false {
                    if false {
                        v1 -= 837
                    }
                    fmt.Println(v1, v1)
                    v29 = v29
                }
                for fi30 := 0; fi30 < 17; fi30++ {
                    fmt.Println(fi30, fi28, v4)
                    var v31 float32 = 0.1
                    _ = fi30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
        }
        _ = v25
    }
    v32 := GenS4{F0: (v2 & 0), F1: 125, F2: ""}
    var v33 uint16 = 32767
    for fi34 := 0; fi34 < 9; fi34++ {
        v35 := GenS2{F0: 19, F1: 2}
        v2 = v4
        v36 := []int{fi34, (fi34 % (((^(fi34))) | 1))}
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS5{F0: true, F1: 0}
        if true {
            v3[fi37] = (fi37 | -1)
        }
        _ = fi37
        _ = v38
    }
    v4 -= v4
    if false {
        v39 := GenS6{F0: v4}
        v40 := []int{-471}
        _ = v39
        _ = v40
    } else {
        var v41 string = "hello"
        fmt.Println(v33, v4, v4)
        _ = v41
    }
    if false {
        v42 := GenS5{F0: true, F1: 2}
        v43 := GenS7{F0: -235, F1: 945}
        v44 := GenS1{F0: 42, F1: v33}
        _ = v42
        _ = v43
        _ = v44
    } else {
        var v45 int8 = -128
        _ = v45
    }
    fmt.Println(v1, v2, v33)
    v1 = genFn0("a", "a", (v4 & (v4 | genFn1())))
    for fi46 := 0; fi46 < 11; fi46++ {
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint16 = v33
            var v49 uint16 = v33
            if true {
                var v50 int32 = -277
                var v51 uint16 = (v48 % ((v33) | 1))
                var v52 uint64 = 749
                _ = v50
                _ = v51
                _ = v52
            }
            var v53 int = (fi46 * -2147483648)
            _ = fi47
            _ = v48
            _ = v49
            _ = v53
        }
        var v54 string = "hello"
        for fi55 := 0; fi55 < 18; fi55++ {
            if (!true) {
                var v56 uint16 = 255
                if true {
                    var v57 float32 = 1e10
                    _ = len(v3)
                    fmt.Println(v4, fi46)
                    _ = v57
                } else {
                    var v58 int8 = 34
                    switch fi46 {
                    case 42:
                        v59 := GenS5{F0: (1e10 != 0.1), F1: 65535}
                        _ = v59
                    }
                    if true {
                        _ = len(v3)
                        v54 = v54
                        v60 := GenS4{F0: genFn1(), F1: (^(v58)), F2: v54}
                        _ = v60
                    } else {
                        if false {
                            v61 := GenS7{F0: int32(v1), F1: 927}
                            switch v54 {
                            case "\n\t":
                                if false {
                                    var v62 string = v54
                                    v63 := GenS3{F0: 2147483647}
                                    _ = v62
                                    _ = v63
                                } else {
                                    var v64 uint16 = 369
                                    switch ("hello" + ((v54 + "foo bar") + v54)) {
                                    case "\n\t":
                                        v2 += (v4 ^ v4)
                                    case "foo bar":
                                        v65 := GenS5{F0: (!(!true)), F1: 1}
                                        _ = v65
                                    default:
                                        var v66 float64 = (3.14159 * 0.0)
                                        _ = v66
                                    }
                                    fmt.Println(v33)
                                    _ = v64
                                }
                                fmt.Println(v33)
                                switch v54 {
                                case "\n\t":
                                    var v67 int32 = 42
                                    _ = v67
                                case "":
                                    if (!((1e10 - 0.0) == 0.1)) {
                                        if true {
                                            v4 -= (^(v4))
                                            var v68 int16 = (-(v2))
                                            var v69 bool = false
                                            v68 = 1000
                                            v70 := GenS5{F0: v69, F1: 722}
                                            v33 -= v33
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                        } else {
                                            fmt.Println(v56, v54)
                                        }
                                    }
                                case "a":
                                    fmt.Println(v4, v33)
                                    if ((v58 ^ v58) <= (v58 ^ v58)) {
                                        fmt.Println(v4, v54)
                                    } else {
                                        v1 = genFn0("hello", (("foo bar" + "") + ("foo bar" + (((v54 + v54) + v54) + v54))), (v4 ^ -782))
                                        v71 := GenS8{F0: 65535, F1: 18446744073709551615, F2: 65535}
                                        v1 = v1
                                        switch 385 {
                                        case 9223372036854775807:
                                            v56 = (v33 ^ 32767)
                                        case -433:
                                            var v72 bool = false
                                            if false {
                                                v73 := GenS9{F0: 171, F1: 42}
                                                _ = v73
                                            }
                                            _ = v72
                                        case 1:
                                            v74 := GenS10{F0: 32767}
                                            _ = v74
                                        }
                                        v54 = ("" + ((v54 + (v54 + v54)) + ("hello" + ((v54 + v54) + (v54 + v54)))))
                                        _ = v71
                                    }
                                    if false {
                                        if (633 == 1) {
                                            v75 := GenS6{F0: v4}
                                            if (!true) {
                                                var v76 int32 = -1
                                                var v77 uint = 909
                                                var v78 uint32 = 4294967295
                                                fmt.Println(fi55)
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            } else {
                                                switch fi55 {
                                                case 127:
                                                    v79 := GenS1{F0: 575, F1: 589}
                                                    _ = v79
                                                case 32768:
                                                    v4 -= (v2 ^ (v4 & -1))
                                                default:
                                                    v54 = ""
                                                }
                                                if ((255 > 313) || ((!true) && true)) {
                                                    v80 := []int{fi55, fi46, fi55}
                                                    _ = v80
                                                }
                                                v56 = (^(v33))
                                                var v81 uint8 = 0
                                                _ = v81
                                            }
                                            _ = v75
                                        } else {
                                            fmt.Println(v33)
                                            var v82 uint16 = v56
                                            _ = v82
                                        }
                                    } else {
                                        v83 := []int{fi46, fi55}
                                        v84 := GenS9{F0: 1, F1: 451}
                                        fmt.Println(v4)
                                        _ = v83
                                        _ = v84
                                    }
                                default:
                                    v85 := GenS6{F0: (v4 | (v2 % (((v4 / (((v2 & v2)) | 1))) | 1)))}
                                    _ = v85
                                }
                            case "a":
                                v86 := GenS8{F0: 1, F1: 0, F2: 223}
                                _ = v86
                            default:
                                v4 += (v2 & v4)
                            }
                            var v87 string = ("" + "a")
                            _ = v61
                            _ = v87
                        } else {
                            var v88 int = 731
                            _ = v88
                        }
                        var v89 uint16 = (v56 + v56)
                        _ = v89
                    }
                    _ = v58
                }
                fmt.Println(v2, fi46, v54)
                _ = v56
            }
            fmt.Println(v1, fi55)
            _ = fi55
        }
        _ = fi46
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
