package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int16 }
type GenS6 struct { F0 int16; F1 int16; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0() int8 {
    var v5 float64 = 3.14159
    v6 := []int{-692, 42, 42}
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = int8(v1)
    v3 := make(map[int]bool)
    var v4 uint16 = v1
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v4)
    case "a":
        var v7 int64 = -1
        _ = v7
    case "":
        v2 = genFn0()
    default:
        var v8 float64 = (0.1 - float64(-2.25))
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v2 = v2
            if false {
                if (uint8(v4) != 130) {
                    v11 := GenS1{F0: v2, F1: -849}
                    fmt.Println(v2, fi10, v4)
                    _ = v11
                } else {
                    fmt.Println(fi9)
                    v12 := []int{fi9, 2147483647}
                    v2 += genFn0()
                    _ = v12
                }
                v13 := GenS1{F0: v2, F1: -32768}
                v14 := v3[fi10]
                if v14 {
                    v2 -= v2
                    v3[fi10] = false
                } else {
                    v2 = genFn0()
                    switch (^(fi10)) {
                    case 999:
                        switch "\n\t" {
                        case "a":
                            v14 = v14
                        default:
                            if true {
                                var v15 int32 = -978
                                v16 := GenS2{F0: 32767, F1: 128}
                                v17 := []int{2147483647}
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                switch 556 {
                                case 1:
                                    v1 = (^(v1))
                                default:
                                    var v18 int64 = 302
                                    _ = v18
                                }
                                v19 := []int{fi9, (fi9 | (fi10 | (^(fi9)))), (fi10 % (((fi9 + 127)) | 1))}
                                v19 = append(v19, fi9)
                                var v20 int = (fi9 ^ fi9)
                                if ((v20 & (^(v20))) <= fi9) {
                                    v2 = -65
                                } else {
                                    v2 -= v2
                                    var v21 uint = 2
                                    _ = v21
                                }
                                v1 -= (v4 & (v1 << ((0) & 63)))
                                fmt.Println(fi9, v14, v1)
                                _ = v19
                                _ = v20
                            }
                        }
                    case 227:
                        delete(v3, int(v2))
                    default:
                        var v22 int32 = -533
                        _ = v22
                    }
                }
                v2 = (v2 & -128)
                fmt.Println(v4, v4)
                v23 := GenS3{F0: (false && (v14 && (255 < 243))), F1: (v2 & 0), F2: -695}
                switch "foo bar" {
                case "\n\t":
                    switch (fi9 >> ((1) & 63)) {
                    case 42:
                        v24 := GenS4{F0: v1, F1: (v1 >> ((18446744073709551615) & 63)), F2: 32767}
                        _ = v24
                    case -1:
                        switch (("foo bar" + ("hello" + "hello")) + "") {
                        case "\n\t":
                            fmt.Println(v2, fi10, fi10)
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    if (("hello" < "a") || v14) {
                                        v4 = 255
                                    }
                                } else {
                                    var v25 int16 = 42
                                    _ = v25
                                }
                            case "":
                                var v26 int8 = 0
                                _ = v26
                            }
                        }
                        if (true || ((v4 >= (v4 >> ((0) & 63))) || false)) {
                            var v27 uint64 = 18446744073709551615
                            var v28 bool = (2147483647 == 65535)
                            _ = v27
                            _ = v28
                        } else {
                            if v14 {
                                if true {
                                    var v29 int16 = 1
                                    v3[42] = v14
                                    _ = v29
                                } else {
                                    var v30 uint8 = 2
                                    if (!(0.0 == (3.14159 - float32(((3.14159 - -2.25) + float64(float32((float64(1e10) * 3.14159)))))))) {
                                        var v31 int32 = -915
                                        _ = v31
                                    }
                                    _ = v30
                                }
                                v32 := []int{fi10, fi9, (fi10 << ((2) & 63))}
                                _ = v32
                            } else {
                                var v33 uint8 = 80
                                v4 = 661
                                if v14 {
                                    v33 -= 163
                                    var v34 int64 = 731
                                    if v14 {
                                        v35 := []int{(^(fi10)), (fi9 ^ (^(fi9)))}
                                        v36 := GenS5{F0: 0, F1: v33, F2: -1}
                                        if false {
                                            var v37 string = ""
                                            var v38 uint = 42
                                            if (fi9 <= fi10) {
                                                switch fi9 {
                                                case 42:
                                                    var v39 uint8 = (v33 | (^(v33)))
                                                    _ = v39
                                                case 32768:
                                                    switch (fi10 - (fi9 | fi9)) {
                                                    case 32767:
                                                        v14 = v14
                                                    case 42:
                                                        var v40 uint = 79
                                                        _ = v40
                                                    }
                                                case -167:
                                                    var v41 uint16 = (v1 >> (((v38 / ((v38) | 1))) & 63))
                                                    _ = v41
                                                }
                                                v3[int(v34)] = false
                                                var v42 string = v37
                                                _ = v42
                                            }
                                            _ = v37
                                            _ = v38
                                        } else {
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                var v43 uint16 = v1
                                                _ = v43
                                            case "hello":
                                                v44 := GenS6{F0: 967, F1: -1, F2: 0}
                                                _ = v44
                                            default:
                                                if v14 {
                                                    _ = len(v35)
                                                    var v45 uint64 = 125
                                                    _ = v45
                                                }
                                            }
                                            v46 := GenS4{F0: 846, F1: v4, F2: 32767}
                                            _ = v46
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                _ = v33
                            }
                        }
                    case 128:
                        v47 := GenS4{F0: v1, F1: (v1 | (v4 + v1)), F2: -32768}
                        _ = v47
                    default:
                        fmt.Println(v2)
                    }
                case "a":
                    v14 = (!(1 > 0))
                }
                _ = v13
                _ = v14
                _ = v23
            } else {
                if (true && ((float64(((1.5 - ((float32(0.0) + 1e10) + 1e10)) - 1e10)) == ((0.0 - 0.0) - 0.1)) || true)) {
                    var v48 string = "hello"
                    var v49 float32 = (1e10 - ((-2.25 * -2.25) / 3.14159))
                    fmt.Println(v48)
                    _ = v48
                    _ = v49
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            if (837 >= 1) {
                                if (-436 != -2147483648) {
                                    var v50 uint8 = 241
                                    _ = v50
                                } else {
                                    if false {
                                        var v51 bool = (473 == -628)
                                        v52 := GenS7{F0: (^(v1))}
                                        _ = v51
                                        _ = v52
                                    } else {
                                        var v53 int64 = int64(v4)
                                        if (!true) {
                                            v53 = (v53 - v53)
                                            var v54 int64 = v53
                                            v55 := []int{(fi9 & fi9), (fi10 * fi9), fi10}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v1 = (v1 << ((255) & 63))
                                        }
                                        _ = v53
                                    }
                                }
                                var v56 int32 = 472
                                v57 := []int{fi9, fi10, fi9, fi9}
                                _ = v56
                                _ = v57
                            } else {
                                var v58 uint = 42
                                var v59 int32 = -32
                                v60 := v3[fi10]
                                fmt.Println(v60, fi10, v2)
                                switch "\n\t" {
                                case "":
                                    if v60 {
                                        switch "a" {
                                        case "":
                                            var v61 uint16 = 217
                                            _ = v61
                                        case "foo bar":
                                            var v62 int8 = (v2 & v2)
                                            _ = v62
                                        case "a":
                                            v2 = genFn0()
                                        }
                                        fmt.Println(v59)
                                    }
                                case "\n\t":
                                    delete(v3, fi9)
                                case "foo bar":
                                    v63 := []int{}
                                    _ = v63
                                default:
                                    var v64 uint64 = 2
                                    _ = v64
                                }
                                fmt.Println(fi9, v60)
                                var v65 int8 = v2
                                _ = v58
                                _ = v59
                                _ = v60
                                _ = v65
                            }
                        default:
                            v1 -= (^(v1))
                        }
                    default:
                        v66 := []int{fi9, (^(fi10)), (fi9 % ((fi9) | 1)), fi10}
                        _ = v66
                    }
                    var v67 uint16 = (^(v1))
                    v68 := v3[fi9]
                    _ = v67
                    _ = v68
                }
            }
            _ = fi10
        }
        v69 := GenS7{F0: v1}
        _ = fi9
        _ = v69
    }
    if true {
        fmt.Println(v2, v4, v1)
        v70 := GenS2{F0: 1000, F1: 137}
        _ = v70
    } else {
        for fi71 := 0; fi71 < 1; fi71++ {
            v1 = v4
            var v72 uint64 = 839
            _ = fi71
            _ = v72
        }
        v4 = 65535
        v73 := v3[0]
        _ = v73
    }
    switch "a" {
    case "":
        v1 = 71
    }
    switch ("" + "a") {
    case "a":
        if true {
            v74 := GenS8{F0: 2147483647, F1: (v2 * (v2 - 42)), F2: -685}
            if (-795 > -2147483648) {
                v75 := []int{32767, 4294967295}
                var v76 float64 = 0.0
                var v77 int16 = -1
                _ = v75
                _ = v76
                _ = v77
            } else {
                v78 := v3[0]
                fmt.Println(v2, v2)
                fmt.Println(v2)
                _ = v78
            }
            _ = v74
        } else {
            fmt.Println(v1)
            switch 32768 {
            case 1:
                v79 := GenS3{F0: false, F1: (-(v2)), F2: 9223372036854775807}
                _ = v79
            default:
                var v80 int = -2147483648
                _ = v80
            }
        }
    case "foo bar":
        if false {
            v2 -= (v2 | 1)
            v81 := GenS3{F0: (42 != -2147483648), F1: int8(v4), F2: -9223372036854775808}
            _ = v81
        }
    }
    fmt.Println(v4)
    var v82 int64 = -1
    fmt.Println(v4, v1)
    var v83 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
}
