package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 int8, p1 int16) bool {
    p1 += p1
    v4 := []int{-335}
    v5 := GenS1{F0: 1, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (!true)
}
func genFn1() string {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    switch ("a" + "hello") {
    case "foo bar":
        var v8 float32 = 3.14159
        _ = v8
    case "":
        var v9 int32 = 42
        _ = v9
    case "hello":
        if (((0 <= -1) || true) && false) {
            if (((!(!(9223372036854775807 == 2147483647))) || false) && false) {
                var v10 int32 = 1
                if (false || false) {
                    if false {
                        switch 848 {
                        case 494:
                            var v11 uint = 2
                            _ = v11
                        case -241:
                            var v12 int8 = -1
                            _ = v12
                        default:
                            v10 = 42
                        }
                        v13 := GenS2{F0: 18446744073709551615, F1: -1}
                        var v14 uint8 = 224
                        _ = v13
                        _ = v14
                    }
                    if (42 < 4294967295) {
                        v15 := GenS1{F0: 0, F1: 1}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 15; fi16++ {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            if (!(1 >= 389)) {
                                v10 = (v10 - (^(v10)))
                                var v18 uint16 = 0
                                _ = v18
                            }
                            var v19 uint16 = 32767
                            var v20 int8 = 0
                            _ = fi17
                            _ = v19
                            _ = v20
                        }
                        var v21 float32 = 3.14159
                        _ = fi16
                        _ = v21
                    }
                } else {
                    if ((!false) && (119 == -13)) {
                        switch "a" {
                        case "":
                            v22 := GenS3{F0: 2, F1: 65535}
                            _ = v22
                        case "foo bar":
                            var v23 uint32 = 4294967295
                            _ = v23
                        default:
                            for fi24 := 0; fi24 < 18; fi24++ {
                                var v25 uint32 = 65535
                                var v26 int32 = (v10 & (^(v10)))
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        }
                        switch 695 {
                        case 9223372036854775807:
                            v27 := []int{9223372036854775807}
                            _ = v27
                        }
                        v28 := []int{}
                        _ = v28
                    } else {
                        v29 := []int{4294967295, 173}
                        var v30 bool = (true && (!true))
                        _ = v29
                        _ = v30
                    }
                    v31 := []int{-53, 2147483647}
                    _ = v31
                }
                v32 := GenS4{F0: -1}
                _ = v10
                _ = v32
            } else {
                switch "hello" {
                case "a":
                    v33 := GenS2{F0: 1, F1: 0}
                    _ = v33
                case "hello":
                    switch -1 {
                    case 821:
                        v34 := GenS5{F0: 883, F1: 18446744073709551615}
                        _ = v34
                    }
                }
                v35 := GenS1{F0: 152, F1: 501}
                _ = v35
            }
            var v36 int64 = 42
            _ = v36
        }
    }
    return ""
}

func main() {
    var v1 uint16 = 1
    var v2 int8 = 42
    var v3 uint8 = 127
    var v37 string = "hello"
    for fi38 := 0; fi38 < 14; fi38++ {
        v2 = v2
        _ = fi38
    }
    fmt.Println(v2, v37, v3)
    if false {
        v39 := GenS6{F0: 42}
        _ = v39
    }
    if false {
        var v40 int16 = 294
        v41 := GenS7{F0: (v37 + v37)}
        _ = v40
        _ = v41
    } else {
        if ((18446744073709551615 > 42) && ((v1 | 255) <= v1)) {
            if (v37 == v37) {
                v42 := GenS2{F0: 2, F1: 0}
                _ = v42
            } else {
                var v43 uint8 = (v3 & v3)
                v2 = (v2 >> ((255) & 63))
                _ = v43
            }
            var v44 uint8 = (^(v3))
            _ = v44
        } else {
            v45 := []int{4294967295, -2147483648}
            if (false || (((0.0 - (3.14159 / 1.0)) - 3.14159) < 1.5)) {
                var v46 uint = 42
                v47 := GenS1{F0: 1, F1: 1}
                if len(v45) > 0 {
                    v45[0] = -1
                }
                _ = v46
                _ = v47
            }
            _ = v45
        }
    }
    v3 -= v3
    var v48 float64 = 0.1
    if (2 < 227) {
        var v49 uint8 = (^(v3))
        fmt.Println(v49, v3, v49)
        var v50 bool = false
        _ = v49
        _ = v50
    } else {
        v51 := GenS7{F0: "hello"}
        var v52 float32 = float32(0.1)
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 19; fi53++ {
        for fi54 := 0; fi54 < 20; fi54++ {
            if false {
                v55 := GenS6{F0: 0}
                _ = v55
            } else {
                fmt.Println(v37, v2)
                v56 := []int{fi54, fi53}
                if false {
                    v2 += (v2 - v2)
                    if true {
                        v48 = (-2.25 * v48)
                        fmt.Println(v1, v1)
                        switch v37 {
                        case "":
                            if len(v56) >= 2 {
                                v56 = v56[1:len(v56)]
                            }
                        }
                    }
                    switch "" {
                    case "":
                        v57 := GenS3{F0: 20, F1: 18446744073709551615}
                        _ = v57
                    case "hello":
                        var v58 int64 = 0
                        _ = v58
                    }
                } else {
                    var v59 string = (v37 + v37)
                    v60 := GenS2{F0: 1, F1: 42}
                    v2 += (v2 | (v2 - (v2 / ((v2) | 1))))
                    switch v37 {
                    case "foo bar":
                        v56 = append(v56, -17)
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v2, fi53, v1)
                        case "hello":
                            if false {
                                v48 += v48
                                fmt.Println(v1, v59)
                            } else {
                                var v61 bool = (1 != -140)
                                _ = v61
                            }
                        }
                        var v62 uint8 = v3
                        _ = v62
                    case "\n\t":
                        var v63 int8 = (-(v2))
                        _ = v63
                    }
                    _ = v59
                    _ = v60
                }
                _ = v56
            }
            v64 := GenS8{F0: v2, F1: 1000}
            if (-9223372036854775808 != 42) {
                if (0.1 <= 0.0) {
                    v65 := GenS7{F0: (v37 + v37)}
                    v37 = (v37 + (v37 + (v37 + (v37 + v37))))
                    switch fi53 {
                    case -9223372036854775808:
                        var v66 string = v37
                        _ = v66
                    case 127:
                        v48 += (v48 + 3.14159)
                    case 433:
                        switch 42 {
                        case 1:
                            if (!false) {
                                v67 := []int{fi54, (fi54 << ((166) & 63)), fi54, fi53}
                                var v68 uint32 = 4294967295
                                var v69 uint64 = uint64(v68)
                                _ = v67
                                _ = v68
                                _ = v69
                            } else {
                                v48 = v48
                                v70 := []int{(fi54 + (fi53 * 4294967295)), 9223372036854775807, (fi54 * -970), 599}
                                var v71 float64 = v48
                                _ = v70
                                _ = v71
                            }
                        case -692:
                            var v72 uint16 = (v1 % ((2) | 1))
                            _ = v72
                        }
                    default:
                        v3 += (v3 / ((v3) | 1))
                    }
                    _ = v65
                }
            } else {
                v73 := []int{}
                var v74 bool = (796 > 0)
                _ = v73
                _ = v74
            }
            _ = fi54
            _ = v64
        }
        v75 := GenS8{F0: v2, F1: 602}
        _ = fi53
        _ = v75
    }
    v76 := GenS4{F0: 1000}
    v77 := []int{32768, 792, 42, 2147483647}
    var v78 uint32 = 612
    v79 := GenS2{F0: 377, F1: 42}
    var v80 float32 = ((3.14159 / 1e10) * 0.1)
    var v81 uint = 65535
    if (true && true) {
        if (((v3 & (v3 % (((v3 | 0)) | 1))) < (v3 | (v3 ^ 255))) || (353 > int(v1))) {
            v1 = (^(v1))
            v82 := GenS4{F0: 32767}
            _ = v82
        } else {
            var v83 uint = (v81 * v81)
            switch ("hello" + v37) {
            case "a":
                var v84 uint64 = uint64(v81)
                _ = v84
            }
            _ = v83
        }
        v85 := GenS9{F0: uint32(v3), F1: 1, F2: 2}
        _ = v85
    } else {
        v48 = 0.1
        v78 = 855
        if false {
            v37 = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v48
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
}
