package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint8; F2 uint }
type GenS8 struct { F0 string; F1 int64; F2 bool }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint }
func genFn0(p0 uint32, p1 bool, p2 int) int32 {
    v7 := []int{p2, (-(p2)), 9223372036854775807, p2}
    if len(v7) > 0 {
        v7[0] = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 48
    var v3 int8 = 0
    v4 := make(map[string]string)
    var v5 float32 = -2.25
    v6 := []uint32{4294967295}
    var v8 uint16 = 0
    if (-9223372036854775808 > 1) {
        var v9 string = "\n\t"
        v10 := GenS1{F0: 518, F1: v2, F2: v3}
        if false {
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v3, v8, v8)
                v3 += v3
                _ = fi11
            }
            var v12 string = v9
            var v13 float64 = 0.0
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(v3, v9, fi14)
                v15 := GenS1{F0: 2147483648, F1: v2, F2: 118}
                _ = fi14
                _ = v15
            }
            fmt.Println(v2, v8, v9)
            var v16 string = v9
            _ = v16
        }
        _ = v10
        _ = v9
    }
    v17 := GenS2{F0: (4294967295 > 826), F1: 2, F2: 2}
    v8 = v8
    switch (("a" + ("a" + "foo bar")) + ("foo bar" + ("hello" + ""))) {
    case "\n\t":
        if false {
            fmt.Println(v5, v5, v3)
            if (!("hello" < ("foo bar" + ("\n\t" + "\n\t")))) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v3 += v3
                    fmt.Println(v8, v3)
                    _ = fi18
                }
                v19 := GenS3{F0: ("\n\t" + "hello"), F1: -2147483648}
                _ = v19
            }
        }
    }
    v8 -= v8
    v8 -= v8
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v3, fi20, v2)
        _ = fi20
    }
    v5 += v5
    v21 := GenS2{F0: false, F1: 65535, F2: 4294967295}
    for fi22 := 0; fi22 < 6; fi22++ {
        if ((fi22 >> ((4294967295) & 63)) < (fi22 * (fi22 >> ((44) & 63)))) {
            v3 += v3
        }
        _ = fi22
    }
    fmt.Println(v2)
    var v23 int8 = 0
    var v24 float32 = ((v5 + 1.5) / v5)
    switch int(v8) {
    case 2147483647:
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 string = "foo bar"
            var v27 uint = 1
            _ = fi25
            _ = v26
            _ = v27
        }
    case 42:
        if (746 <= 2) {
            v2 += v2
            v3 -= v23
        }
    case -2147483648:
        var v28 uint64 = 18446744073709551615
        _ = v28
    default:
        var v29 uint = 42
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v24 = -2.25
        _ = fi30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        if false {
            v24 = 0.0
            switch fi31 {
            case 127:
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 bool = (460 != 42)
                    fmt.Println(v8)
                    _ = fi32
                    _ = v33
                }
            case 32768:
                if (42 < genFn0(1, (1 <= -178), (-(fi31)))) {
                    v34 := GenS4{F0: 66, F1: v3, F2: v3}
                    v35 := GenS5{F0: v8}
                    fmt.Println(v23)
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint32 = 2147483647
                    if ("foo bar" >= ("a" + "")) {
                        if false {
                            v1[(fi31 | fi31)] = fi31
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v38 := v1[(fi31 << ((255) & 63))]
                                if (v24 < v5) {
                                    if (v5 == 1e10) {
                                        var v39 uint = 42
                                        _ = v39
                                    } else {
                                        if false {
                                            v40 := GenS3{F0: "", F1: 676}
                                            v36 += (v36 << ((0) & 63))
                                            fmt.Println(v23, fi31, v38)
                                            var v41 int = fi31
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                    var v42 int16 = v2
                                    v38 = v38
                                    _ = v42
                                }
                                _ = fi37
                                _ = v38
                            }
                            if ((0.0 - 0.1) <= -2.25) {
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    v44 := GenS4{F0: 42, F1: 0, F2: v3}
                                    delete(v1, (fi31 >> ((0) & 63)))
                                    _ = fi43
                                    _ = v44
                                }
                                if (v8 > v8) {
                                    var v45 int16 = 59
                                    _ = v45
                                } else {
                                    v2 = (^(v2))
                                    var v46 uint8 = 209
                                    _ = v46
                                }
                                fmt.Println(v2)
                            } else {
                                var v47 uint16 = v8
                                v48 := GenS6{F0: 4294967295}
                                _ = v47
                                _ = v48
                            }
                        } else {
                            v8 = (v8 - v8)
                        }
                        var v49 int8 = v23
                        v8 = v8
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 9; fi50++ {
                            var v51 int16 = 42
                            v52 := GenS5{F0: (^(v8))}
                            switch fi31 {
                            case 86:
                                if (42 != (v36 % ((2) | 1))) {
                                    var v53 float64 = (float64(((float32((((0.0 * 1e10) / 3.14159) - 1.5)) + ((0.0 - float32(float64(1.5))) - float32(0.0))) + 0.1)) + 1e10)
                                    if ((!true) && (1 != 757)) {
                                        v36 = v36
                                        v54 := GenS7{F0: -9223372036854775808, F1: 255, F2: 2}
                                        if (!false) {
                                            v8 = 2
                                            v55 := GenS8{F0: "hello", F1: 42, F2: true}
                                            var v56 int = int(v51)
                                            switch -915 {
                                            case -239:
                                                _ = len(v6)
                                            default:
                                                v57 := GenS5{F0: v8}
                                                _ = v57
                                            }
                                            switch (v56 >> ((4294967295) & 63)) {
                                            case -289:
                                                v3 -= (v3 << ((979) & 63))
                                            case 1:
                                                v58 := GenS9{F0: 647, F1: 386}
                                                _ = v58
                                            default:
                                                fmt.Println(v8)
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                        var v59 int = (^(fi31))
                                        v60 := GenS3{F0: "foo bar", F1: 42}
                                        _ = v54
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = v53
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        var v61 uint64 = 2
                                        _ = v61
                                    case "a":
                                        switch "" {
                                        case "\n\t":
                                            v62 := GenS10{F0: 128, F1: 32767}
                                            _ = v62
                                        case "a":
                                            v63 := GenS2{F0: (!false), F1: (v36 ^ v36), F2: 18446744073709551615}
                                            _ = v63
                                        case "foo bar":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        default:
                                            v64 := GenS11{F0: 42}
                                            _ = v64
                                        }
                                    }
                                    v24 = float32(0.0)
                                }
                            case 128:
                                fmt.Println(fi31, fi50, fi31)
                            default:
                                var v65 uint64 = 4294967295
                                _ = v65
                            }
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v66 := v1[fi31]
                        var v67 uint32 = (v36 << ((1) & 63))
                        _ = v66
                        _ = v67
                    }
                    fmt.Println(v24)
                    _ = v36
                }
            }
        } else {
            if true {
                for fi68 := 0; fi68 < 2; fi68++ {
                    v8 = v8
                    _ = fi68
                }
            }
            v24 = (((0.0 * v24) - v24) / ((v5 - 0.1) / ((-2.25 * (-2.25 - float32(0.0))) + v24)))
        }
        var v69 uint32 = 524
        _ = fi31
        _ = v69
    }
    v4["a"] = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
