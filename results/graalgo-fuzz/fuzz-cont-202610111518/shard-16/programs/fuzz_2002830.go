package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS9 struct { F0 bool; F1 uint8; F2 bool }
type GenS10 struct { F0 uint8; F1 uint8; F2 int32 }
func genFn0(p0 int64, p1 string, p2 int) uint32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        if (255 > 0) {
            var v6 int8 = -1
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := []int{fi5, fi7, fi7, p2}
                var v9 float64 = (float64(float32((1.5 - (3.14159 * -2.25)))) / 1.5)
                var v10 uint8 = 128
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            v11 := []int{p2}
            _ = v11
            _ = v6
        }
        _ = fi5
    }
    v12 := GenS1{F0: 672, F1: 819, F2: 4294967295}
    v13 := GenS1{F0: 42, F1: 0, F2: 565}
    p0 = (p0 ^ (p0 - (p0 ^ p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}
func genFn1(p0 uint) int32 {
    v14 := []int{32768, -9223372036854775808}
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint8 = 42
        for fi17 := 0; fi17 < 5; fi17++ {
            switch (("hello" + ("" + "hello")) + (("a" + "\n\t") + ("foo bar" + ("\n\t" + ("a" + "a"))))) {
            case "":
                v18 := GenS1{F0: p0, F1: 0, F2: 639}
                switch ("a" + "a") {
                case "hello":
                    if len(v14) > 0 {
                        v14[0] = fi15
                    }
                case "\n\t":
                    _ = len(v14)
                case "a":
                    v19 := GenS2{F0: ""}
                    _ = v19
                default:
                    if true {
                        v20 := GenS3{F0: 42, F1: 1}
                        v16 -= v16
                        fmt.Println(fi17, v16, fi15)
                        _ = v20
                    }
                }
                _ = v18
            default:
                switch ("" + "") {
                case "a":
                    v16 += (v16 >> (((p0 & (p0 % ((p0) | 1)))) & 63))
                    v21 := GenS4{F0: 699}
                    _ = v21
                case "":
                    if (!true) {
                        var v22 float64 = -2.25
                        _ = v22
                    } else {
                        _ = len(v14)
                    }
                default:
                    var v23 uint = p0
                    _ = v23
                }
            }
            if true {
                var v24 uint64 = 18446744073709551615
                var v25 uint16 = 671
                if true {
                    switch "foo bar" {
                    case "a":
                        if false {
                            v26 := GenS5{F0: "hello", F1: 42}
                            if true {
                                fmt.Println(v25, v24, fi17)
                            } else {
                                _ = len(v14)
                                v27 := GenS3{F0: 469, F1: (v25 - v25)}
                                var v28 bool = (-1 < -68)
                                _ = v27
                                _ = v28
                            }
                            _ = v26
                        } else {
                            var v29 uint32 = 302
                            switch (fi15 << (((p0 << ((1) & 63))) & 63)) {
                            case 127:
                                var v30 float64 = 3.14159
                                _ = v30
                            case 2147483647:
                                if false {
                                    var v31 bool = false
                                    v31 = v31
                                    _ = v31
                                }
                            }
                            _ = v29
                        }
                    }
                    v25 = v25
                } else {
                    var v32 int8 = 121
                    _ = v32
                }
                _ = v24
                _ = v25
            } else {
                p0 = p0
            }
            _ = fi17
        }
        var v33 float32 = (-2.25 - ((((float32(3.14159) - 0.1) * -2.25) - (1.5 + -2.25)) / 0.1))
        _ = fi15
        _ = v16
        _ = v33
    }
    _ = p0
    _ = v14
    return -262
}

func main() {
    var v1 int16 = 18
    var v2 int8 = -1
    var v3 uint16 = 255
    var v4 int8 = (v2 ^ (v2 - v2))
    v34 := GenS4{F0: 0}
    v35 := GenS3{F0: 18446744073709551615, F1: (v3 >> ((584) & 63))}
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := []int{391, (-(fi36)), 2147483647, fi36}
        switch fi36 {
        case 671:
            v38 := GenS6{F0: (-(v1)), F1: 65535}
            _ = v38
        case 4294967295:
            _ = len(v37)
        case -469:
            for fi39 := 0; fi39 < 12; fi39++ {
                var v40 bool = true
                var v41 bool = v40
                var v42 int16 = v1
                v41 = v41
                switch -2147483648 {
                case -947:
                    var v43 int = (fi36 ^ -519)
                    _ = v43
                default:
                    var v44 int = (fi36 % (((fi36 >> ((156) & 63))) | 1))
                    _ = v44
                }
                if (42 != 42) {
                    v41 = false
                    v4 = v4
                }
                var v45 uint16 = (^(v3))
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
                _ = v45
            }
        }
        _ = fi36
        _ = v37
    }
    if (true && (!(9223372036854775807 > 4))) {
        for fi46 := 0; fi46 < 1; fi46++ {
            v4 = v2
            _ = fi46
        }
        v4 = (v4 ^ 100)
    } else {
        v47 := GenS1{F0: 42, F1: 2, F2: 4294967295}
        var v48 int16 = v1
        _ = v47
        _ = v48
    }
    v2 -= (v4 % ((v2) | 1))
    v49 := []int{-1, 9223372036854775807}
    for fi50 := 0; fi50 < 18; fi50++ {
        v1 += v1
        switch "hello" {
        case "a":
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
        default:
            var v51 uint32 = genFn0(1, "foo bar", (fi50 - fi50))
            _ = v51
        }
        _ = fi50
    }
    if len(v49) > 0 {
        v49[0] = 128
    }
    v3 = (v3 / ((531) | 1))
    var v52 float64 = (((3.14159 - float64(3.14159)) * 3.14159) / 1.0)
    v4 += (v4 >> ((1) & 63))
    if (!(2 >= 255)) {
        if false {
            for fi53 := 0; fi53 < 17; fi53++ {
                for fi54 := 0; fi54 < 2; fi54++ {
                    fmt.Println(fi53, v4, v3)
                    if len(v49) > 0 {
                        v49[0] = (fi53 << ((0) & 63))
                    }
                    _ = fi54
                }
                v55 := GenS1{F0: 240, F1: 32767, F2: 18446744073709551615}
                v56 := GenS1{F0: 4294967295, F1: (v3 & (^(v3))), F2: 299}
                _ = fi53
                _ = v55
                _ = v56
            }
            for fi57 := 0; fi57 < 5; fi57++ {
                var v58 float64 = v52
                _ = fi57
                _ = v58
            }
        } else {
            v49 = append(v49, int(v2))
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            fmt.Println(v3)
        }
        switch "hello" {
        case "hello":
            v3 -= v3
        case "a":
            switch 1 {
            case 127:
                fmt.Println(v3, v3, v52)
            default:
                v3 = v3
            }
        default:
            for fi59 := 0; fi59 < 7; fi59++ {
                v60 := GenS7{F0: true}
                _ = fi59
                _ = v60
            }
        }
        if (true && ((-961 != 32767) && (1 < -63))) {
            v61 := GenS8{F0: 0, F1: 989, F2: 921}
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            _ = v61
        }
    }
    var v62 uint64 = 4294967295
    if true {
        var v63 int32 = -155
        switch "" {
        case "\n\t":
            if true {
                switch 32768 {
                case -932:
                    v49 = append(v49, -9223372036854775808)
                default:
                    var v64 bool = (3.14159 >= ((((-2.25 - -2.25) / 1e10) - ((float32((1.5 + v52)) - (0.0 + -2.25)) + (((3.14159 / 0.1) + (-2.25 / -2.25)) + 1e10))) + (-2.25 * (0.1 + 0.0))))
                    _ = v64
                }
            }
        }
        v49 = append(v49, 1)
        _ = v63
    }
    if (0 <= 2) {
        for fi65 := 0; fi65 < 10; fi65++ {
            for fi66 := 0; fi66 < 11; fi66++ {
                var v67 string = ""
                var v68 uint8 = 0
                switch "hello" {
                case "hello":
                    switch v67 {
                    case "":
                        v68 = v68
                    default:
                        v69 := GenS1{F0: 65535, F1: (v3 >> ((2) & 63)), F2: v62}
                        _ = v69
                    }
                case "a":
                    v70 := GenS5{F0: ((v67 + "a") + ((v67 + (v67 + "a")) + "")), F1: (v1 % ((v1) | 1))}
                    _ = v70
                default:
                    if (!false) {
                        switch int(v68) {
                        case 4294967295:
                            if false {
                                var v71 bool = (true && false)
                                var v72 int64 = 9223372036854775807
                                _ = v71
                                _ = v72
                            }
                        case 1:
                            switch (fi66 % ((fi66) | 1)) {
                            case -119:
                                var v73 string = v67
                                _ = v73
                            case 128:
                                v68 -= 114
                            case -46:
                                var v74 float32 = ((0.0 * 0.1) - 0.1)
                                _ = v74
                            default:
                                var v75 bool = (true && ((!true) || (128 < (v68 << ((65535) & 63)))))
                                _ = v75
                            }
                        default:
                            v1 += (v1 * v1)
                        }
                    } else {
                        _ = len(v49)
                        if (1 != v3) {
                            var v76 int64 = -23
                            switch -380 {
                            case 127:
                                v77 := GenS7{F0: false}
                                _ = v77
                            case 32768:
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                            }
                            v62 -= 0
                            _ = v76
                        }
                        switch v67 {
                        case "":
                            v78 := GenS9{F0: ((fi66 ^ fi65) < -975), F1: (v68 ^ (v68 * v68)), F2: (!(((v52 + v52) - v52) > v52))}
                            _ = v78
                        case "hello":
                            v79 := GenS10{F0: v68, F1: (v68 >> ((18446744073709551615) & 63)), F2: genFn1(42)}
                            _ = v79
                        case "foo bar":
                            v52 += 1.5
                        default:
                            fmt.Println(v52, fi66, v1)
                        }
                        if ((v68 * v68) != 1) {
                            v68 = v68
                            var v80 int32 = int32(v1)
                            switch v67 {
                            case "":
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                            case "a":
                                _ = len(v49)
                            default:
                                fmt.Println(v1, v52)
                            }
                            _ = v80
                        } else {
                            v4 += v4
                            if len(v49) > 0 {
                                v49[0] = fi66
                            }
                        }
                    }
                }
                _ = fi66
                _ = v67
                _ = v68
            }
            _ = fi65
        }
        switch 32767 {
        case 9223372036854775807:
            if false {
                v62 -= 18446744073709551615
            } else {
                fmt.Println(v52, v52)
                for fi81 := 0; fi81 < 7; fi81++ {
                    var v82 float64 = v52
                    _ = fi81
                    _ = v82
                }
                var v83 uint = 440
                _ = v83
            }
        case -330:
            _ = len(v49)
        case 127:
            fmt.Println(v3, v4)
        }
        var v84 int = 1
        _ = v84
    }
    v85 := GenS8{F0: 372, F1: v62, F2: 0}
    _ = len(v49)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v52:", v52)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v49
    _ = v52
    _ = v62
    _ = v85
}
