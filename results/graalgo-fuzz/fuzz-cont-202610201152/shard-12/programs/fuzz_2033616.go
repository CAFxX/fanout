package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 int16; F1 uint64; F2 string }
type GenS8 struct { F0 uint8; F1 uint16; F2 uint }
type GenS9 struct { F0 int32; F1 uint16 }
func genFn0(p0 int32) uint32 {
    p0 = (p0 % ((p0) | 1))
    if false {
        v6 := []int{9223372036854775807, 2147483647, 0, 42}
        p0 += (p0 % ((-2147483648) | 1))
        v7 := GenS1{F0: false, F1: 65535, F2: 226}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = ("hello" + "foo bar")
        switch ((v9 + "foo bar") + (v9 + v9)) {
        case "foo bar":
            v9 = "a"
        case "\n\t":
            v10 := GenS1{F0: ((-2.25 - (0.0 * 1.5)) > 3.14159), F1: 3, F2: 2}
            _ = v10
        }
        if true {
            v11 := GenS2{F0: 1, F1: 4294967295}
            v12 := GenS3{F0: 7, F1: 1}
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v13 := GenS3{F0: 79, F1: -46}
    _ = p0
    _ = v13
    return 0
}
func genFn1(p0 uint32, p1 bool, p2 int32) int {
    var v14 float32 = 1e10
    v15 := GenS3{F0: 1, F1: 32767}
    v16 := []int{-118, 42, 591, 127}
    var v17 bool = (float64(v14) >= float64((v14 / (1e10 + v14))))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 560
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 127
    var v3 uint16 = 42
    var v4 int32 = -2147483648
    var v5 float32 = 3.14159
    var v18 uint16 = v3
    v3 -= v18
    v1 = v1
    var v19 bool = (1 < -8)
    v18 = v18
    var v20 string = v1
    if true {
        if true {
            if ((0 > -1) || v19) {
                switch 42 {
                case 4294967295:
                    switch ("foo bar" + (v20 + "\n\t")) {
                    case "":
                        v21 := GenS4{F0: (v2 / (((v2 * 0)) | 1))}
                        _ = v21
                    case "a":
                        var v22 int16 = 1
                        _ = v22
                    }
                case -2147483648:
                    var v23 uint16 = (v18 + (v3 + v18))
                    _ = v23
                case 128:
                    var v24 float64 = (3.14159 * 0.1)
                    _ = v24
                default:
                    switch v1 {
                    case "foo bar":
                        var v25 float64 = float64((((((v5 / (-2.25 + ((v5 / ((((v5 / v5) * v5) / v5) - v5)) - 1e10))) / (-2.25 * v5)) * v5) / 1.0) * v5))
                        _ = v25
                    case "":
                        v26 := GenS4{F0: (v2 & v2)}
                        _ = v26
                    default:
                        var v27 float64 = 0.0
                        _ = v27
                    }
                }
                v18 = 65535
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 uint64 = 0
                    if ("foo bar" >= (v1 + v1)) {
                        fmt.Println(v19, v3)
                    } else {
                        if ((-1 > (^(v4))) && v19) {
                            v30 := []int{fi28, 9223372036854775807, fi28}
                            _ = v30
                        }
                        v31 := []int{fi28, genFn1(2147483648, (true || v19), (v4 / ((v4) | 1))), genFn1(2147483647, v19, -608), (fi28 << ((2) & 63))}
                        var v32 int8 = -65
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                }
            } else {
                v19 = (!(v2 <= (v2 & (v2 & (v2 >> ((65535) & 63))))))
                if (128 == (v2 % ((v2) | 1))) {
                    var v33 uint16 = v3
                    var v34 int16 = -84
                    var v35 int32 = (v4 * v4)
                    _ = v33
                    _ = v34
                    _ = v35
                }
            }
            v36 := GenS2{F0: 1, F1: 591}
            switch v20 {
            case "hello":
                switch v1 {
                case "\n\t":
                    if (v19 && v19) {
                        fmt.Println(v19, v3)
                        switch 32768 {
                        case 127:
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v19 = v19
                                if v19 {
                                    v38 := GenS5{F0: (v3 << ((18446744073709551615) & 63)), F1: 1}
                                    var v39 int32 = 42
                                    _ = v38
                                    _ = v39
                                }
                                var v40 int16 = -77
                                _ = fi37
                                _ = v40
                            }
                        }
                    } else {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            if v19 {
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    v43 := GenS5{F0: (v18 - 2), F1: 1000}
                                    var v44 int32 = v4
                                    v45 := GenS1{F0: (v19 && v19), F1: 821, F2: (v2 | (^(v2)))}
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                if (((float64(v5) == 0.0) && v19) && (!(v1 != v1))) {
                                    var v46 float32 = v5
                                    var v47 int = -365
                                    var v48 uint = 2
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                } else {
                                    v18 = (v3 | v3)
                                    v18 = 42
                                    var v49 bool = (-1 >= -1)
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 5; fi50++ {
                                    if v19 {
                                        var v51 float64 = 0.0
                                        fmt.Println(fi50, v19, v4)
                                        fmt.Println(v19, v20, v19)
                                        _ = v51
                                    } else {
                                        var v52 bool = ((49 == 255) && (!true))
                                        if v52 {
                                            v53 := GenS2{F0: 18446744073709551615, F1: 382}
                                            switch v1 {
                                            case "hello":
                                                v54 := GenS4{F0: (v2 + 1)}
                                                _ = v54
                                            case "\n\t":
                                                v55 := GenS6{F0: 65535, F1: genFn1(525, false, v4)}
                                                _ = v55
                                            case "foo bar":
                                                v4 = v4
                                            }
                                            fmt.Println(v18, v19)
                                            _ = v53
                                        } else {
                                            var v56 bool = false
                                            _ = v56
                                        }
                                        var v57 float64 = (-2.25 + (0.0 * 1e10))
                                        _ = v52
                                        _ = v57
                                    }
                                    v18 -= (v18 / (((v18 / ((v18) | 1))) | 1))
                                    v58 := GenS6{F0: 399, F1: fi41}
                                    var v59 uint16 = 143
                                    _ = fi50
                                    _ = v58
                                    _ = v59
                                }
                            } else {
                                v60 := []int{genFn1(4294967295, (127 != 0), v4), fi41}
                                v61 := GenS3{F0: 1, F1: 32767}
                                fmt.Println(v3)
                                _ = v60
                                _ = v61
                            }
                            _ = fi41
                        }
                        v62 := []int{32768, -973, 127, -844}
                        v63 := GenS7{F0: 42, F1: 0, F2: v20}
                        _ = v62
                        _ = v63
                    }
                }
            }
            _ = v36
        }
        if (!(1 < (v2 * 106))) {
            v5 = v5
        } else {
            switch 128 {
            case 580:
                v64 := GenS8{F0: (v2 % ((v2) | 1)), F1: 1, F2: 128}
                _ = v64
            }
        }
        v65 := []int{32768}
        _ = v65
    } else {
        if true {
            if v19 {
                v19 = v19
                v19 = (-900 >= 952)
            } else {
                if ((v2 ^ v2) > v2) {
                    fmt.Println(v5)
                } else {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        if (!true) {
                            v67 := GenS9{F0: (v4 + (v4 << ((255) & 63))), F1: (v18 >> ((99) & 63))}
                            switch v1 {
                            case "foo bar":
                                v68 := []int{(fi66 - fi66), int(v2)}
                                _ = v68
                            case "":
                                for fi69 := 0; fi69 < 17; fi69++ {
                                    v70 := []int{fi66, (fi69 | (fi66 >> ((2) & 63))), fi69}
                                    _ = fi69
                                    _ = v70
                                }
                            default:
                                fmt.Println(v18, fi66)
                            }
                            _ = v67
                        } else {
                            var v71 uint32 = 65535
                            for fi72 := 0; fi72 < 18; fi72++ {
                                var v73 int64 = -9223372036854775808
                                fmt.Println(v3, v20)
                                fmt.Println(v3, v5, v5)
                                _ = fi72
                                _ = v73
                            }
                            _ = v71
                        }
                        var v74 int64 = -8
                        _ = fi66
                        _ = v74
                    }
                }
                fmt.Println(v3, v18)
                switch 42 {
                case 42:
                    var v75 float32 = ((float32((3.14159 / 1.0)) - v5) * -2.25)
                    _ = v75
                case 127:
                    fmt.Println(v2)
                case 128:
                    var v76 int8 = -1
                    _ = v76
                default:
                    for fi77 := 0; fi77 < 6; fi77++ {
                        if (genFn1(2147483647, v19, v4) >= fi77) {
                            v5 += v5
                            v4 = v4
                        } else {
                            v78 := []int{genFn1(407, v19, v4), fi77}
                            _ = v78
                        }
                        v79 := []int{genFn1(716, true, (v4 / ((v4) | 1))), (-(fi77))}
                        v19 = true
                        _ = fi77
                        _ = v79
                    }
                }
            }
        }
        for fi80 := 0; fi80 < 12; fi80++ {
            v19 = true
            _ = fi80
        }
    }
    v81 := GenS2{F0: 505, F1: 18446744073709551615}
    for fi82 := 0; fi82 < 2; fi82++ {
        switch fi82 {
        case 394:
            v83 := []int{4294967295, (fi82 & fi82)}
            _ = v83
        case -425:
            if v19 {
                if true {
                    switch fi82 {
                    case 0:
                        fmt.Println(v5, v18)
                    case 9223372036854775807:
                        v1 = v20
                    case -1:
                        v2 = v2
                    default:
                        for fi84 := 0; fi84 < 16; fi84++ {
                            v4 += (^(v4))
                            v20 = v20
                            _ = fi84
                        }
                    }
                }
            } else {
                v19 = v19
                if true {
                    var v85 int = -433
                    _ = v85
                }
                if v19 {
                    for fi86 := 0; fi86 < 16; fi86++ {
                        switch fi86 {
                        case 127:
                            v19 = v19
                            fmt.Println(v4, v5, v5)
                            v3 -= 65535
                        default:
                            var v87 uint16 = v18
                            _ = v87
                        }
                        v88 := []int{(fi82 + (fi86 + fi86))}
                        var v89 bool = v19
                        _ = fi86
                        _ = v88
                        _ = v89
                    }
                    v2 += 103
                }
            }
        default:
            var v90 int32 = v4
            _ = v90
        }
        var v91 int16 = 425
        v20 = ""
        _ = fi82
        _ = v91
    }
    v92 := []int{2147483647, 380}
    if v19 {
        v4 = (v4 * -2147483648)
        for fi93 := 0; fi93 < 16; fi93++ {
            for fi94 := 0; fi94 < 8; fi94++ {
                v20 = ("a" + v20)
                _ = fi94
            }
            for fi95 := 0; fi95 < 3; fi95++ {
                if len(v92) >= 2 {
                    v92 = v92[1:len(v92)]
                }
                var v96 uint32 = 1
                _ = fi95
                _ = v96
            }
            v97 := GenS3{F0: 18446744073709551615, F1: 63}
            _ = fi93
            _ = v97
        }
        v2 = v2
    }
    var v98 bool = (((0.0 + float64((((v5 + v5) - v5) / v5))) - -2.25) < 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v92
    _ = v98
}
