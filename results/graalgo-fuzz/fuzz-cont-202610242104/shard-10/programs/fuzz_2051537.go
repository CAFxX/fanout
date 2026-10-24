package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int; F2 uint8 }
type GenS9 struct { F0 int32; F1 int16 }
func genFn0(p0 uint16, p1 uint8, p2 int16) string {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int16 = p2
        v9 := []int{(fi7 % ((fi7) | 1)), fi7, fi7}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "foo bar"
}
func genFn1(p0 uint) int16 {
    p0 = (p0 ^ 18446744073709551615)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 1, F1: 42}
        _ = fi11
        _ = v12
    }
    v13 := []int{int(p0), 1}
    v13 = append(v13, 1)
    _ = p0
    _ = v13
    return 32767
}

func main() {
    var v1 string = "hello"
    v2 := []string{"hello", v1, v1, (v1 + "")}
    v3 := make(map[int]int)
    v4 := []string{}
    var v5 int32 = 2147483647
    v6 := []uint32{}
    delete(v3, 32767)
    switch v1 {
    case "foo bar":
        v1 = "hello"
    case "hello":
        v5 += (v5 + (^(v5)))
    default:
        if true {
            v5 = (v5 / (((v5 ^ 0)) | 1))
        }
    }
    switch v1 {
    case "\n\t":
        var v14 int = -660
        _ = v14
    default:
        fmt.Println(v5, v5)
    }
    var v15 uint16 = 0
    v1 = v1
    var v16 int = 128
    fmt.Println(v5)
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int32 = (v5 | 42)
        if ((3.14159 * 3.14159) >= (1e10 / 1.5)) {
            switch (v16 - v16) {
            case 32768:
                v18 += (^(v5))
            case 748:
                if true {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 int32 = v5
                        _ = fi19
                        _ = v20
                    }
                    if (true && (356 < (v18 - (v5 | v18)))) {
                        var v21 uint64 = 2
                        _ = v21
                    }
                    v2 = append(v2, (v1 + ""))
                } else {
                    _ = len(v3)
                    fmt.Println(fi17, v5, v18)
                    delete(v3, 42)
                }
            case 128:
                var v22 int = fi17
                _ = v22
            }
        } else {
            _ = len(v4)
        }
        v18 = (v18 / ((v5) | 1))
        _ = fi17
        _ = v18
    }
    fmt.Println(v15)
    var v23 float64 = -2.25
    if (!(1e10 >= (float64(-2.25) - (float64(-2.25) - (((v23 - (float64(float32(v23)) * v23)) * (v23 * -2.25)) / v23))))) {
        v24 := GenS2{F0: 4294967295}
        if ((0 >= 0) || false) {
            var v25 int8 = 71
            _ = v25
        }
        _ = v24
    } else {
        var v26 string = "\n\t"
        _ = v26
    }
    if (!(true && (!(1 >= 1)))) {
        if false {
            v3[v16] = v16
            var v27 uint8 = 96
            v28 := GenS3{F0: (^(v16)), F1: (v15 * (v15 / (((v15 / ((v15) | 1))) | 1)))}
            _ = v27
            _ = v28
        } else {
            if (v15 != v15) {
                switch (((v1 + v1) + v1) + v1) {
                case "":
                    if (false && true) {
                        switch v1 {
                        case "\n\t":
                            var v29 int16 = genFn1(1)
                            _ = v29
                        }
                        var v30 int8 = -57
                        if (!true) {
                            if false {
                                fmt.Println(v15)
                            } else {
                                _ = len(v3)
                                v31 := GenS4{F0: 255}
                                fmt.Println(v30, v15)
                                _ = v31
                            }
                            v32 := GenS5{F0: (v15 << ((18446744073709551615) & 63)), F1: 1, F2: (v15 * (v15 & v15))}
                            _ = v32
                        }
                        _ = v30
                    } else {
                        if true {
                            v15 += v15
                            v5 -= (^(v5))
                            var v33 float64 = v23
                            _ = v33
                        }
                        var v34 uint = 151
                        v35 := GenS6{F0: v16}
                        _ = v34
                        _ = v35
                    }
                case "a":
                    switch v1 {
                    case "":
                        v15 -= v15
                    case "hello":
                        if false {
                            var v36 string = v1
                            if true {
                                for fi37 := 0; fi37 < 4; fi37++ {
                                    v38 := GenS7{F0: 968}
                                    _ = fi37
                                    _ = v38
                                }
                                delete(v3, (v16 | (v16 / ((-1) | 1))))
                                fmt.Println(v1, v15, v15)
                            }
                            v5 += v5
                            _ = v36
                        }
                    default:
                        var v39 string = ""
                        _ = v39
                    }
                case "\n\t":
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v23 -= ((v23 * v23) + (float64((float32(v23) - 3.14159)) / 3.14159))
                        var v41 uint8 = 186
                        var v42 int64 = -1
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                default:
                    v15 += v15
                }
                if true {
                    v23 = v23
                    _ = len(v6)
                    for fi43 := 0; fi43 < 10; fi43++ {
                        if false {
                            v44 := GenS5{F0: v15, F1: 2, F2: 32767}
                            v6 = append(v6, 119)
                            _ = v44
                        }
                        _ = fi43
                    }
                } else {
                    v23 -= v23
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v45 := GenS5{F0: (^(v15)), F1: 1, F2: v15}
                    _ = v45
                }
                v46 := GenS3{F0: int(v5), F1: v15}
                _ = v46
            } else {
                v47 := GenS1{F0: (v15 / ((v15) | 1)), F1: (v15 + v15)}
                if false {
                    v48 := GenS8{F0: -872, F1: (v16 >> ((255) & 63)), F2: 0}
                    v16 += v16
                    _ = v48
                } else {
                    var v49 uint8 = 128
                    _ = v49
                }
                var v50 int32 = (v5 + (v5 << ((0) & 63)))
                _ = v47
                _ = v50
            }
            switch v16 {
            case 363:
                switch v1 {
                case "a":
                    fmt.Println(v23, v5, v1)
                case "foo bar":
                    _ = len(v2)
                default:
                    v51 := GenS4{F0: v15}
                    _ = v51
                }
            case -1:
                var v52 bool = (!(true || false))
                _ = v52
            default:
                fmt.Println(v23, v15, v15)
            }
        }
    }
    switch v1 {
    case "\n\t":
        var v53 int64 = -403
        _ = v53
    case "foo bar":
        delete(v3, v16)
    }
    switch (v16 / ((v16) | 1)) {
    case 75:
        v1 = "a"
    case -9223372036854775808:
        switch 9223372036854775807 {
        case -2147483648:
            if false {
                v15 -= (v15 - 65535)
                for fi54 := 0; fi54 < 2; fi54++ {
                    var v55 bool = false
                    if v55 {
                        fmt.Println(v16)
                        switch fi54 {
                        case 128:
                            switch -1 {
                            case 32768:
                                v55 = (genFn1(1) >= -32768)
                            }
                        case 4294967295:
                            v2 = append(v2, (v1 + "\n\t"))
                        case 0:
                            for fi56 := 0; fi56 < 6; fi56++ {
                                if (v23 != 3.14159) {
                                    fmt.Println(v55, v23, v23)
                                    var v57 bool = v55
                                    v58 := GenS2{F0: 95}
                                    _ = len(v3)
                                    _ = v57
                                    _ = v58
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = v1
                                    }
                                    var v59 uint8 = 208
                                    v1 = ((((v1 + "a") + ("" + v1)) + (v1 + v1)) + "a")
                                    _ = v59
                                }
                                fmt.Println(v23, v16, v16)
                                v15 -= 32767
                                var v60 uint = 42
                                _ = fi56
                                _ = v60
                            }
                        }
                        v15 -= 32767
                    }
                    if (-1 == 32767) {
                        v5 = v5
                    }
                    _ = fi54
                    _ = v55
                }
                var v61 uint16 = 2
                _ = v61
            } else {
                switch "\n\t" {
                case "\n\t":
                    if (504 <= 2) {
                        v5 = v5
                    } else {
                        if true {
                            v23 += 1e10
                        }
                        for fi62 := 0; fi62 < 11; fi62++ {
                            switch fi62 {
                            case 42:
                                v5 -= v5
                            default:
                                var v63 int64 = -426
                                _ = v63
                            }
                            _ = fi62
                        }
                    }
                case "a":
                    var v64 uint32 = 2147483647
                    _ = v64
                case "foo bar":
                    var v65 uint64 = 476
                    _ = v65
                default:
                    _ = len(v3)
                }
                switch v16 {
                case 603:
                    delete(v3, 32767)
                }
                v16 = (v16 % ((-337) | 1))
            }
        case 128:
            var v66 uint16 = (v15 - v15)
            _ = v66
        case 32768:
            switch 32768 {
            case 81:
                v67 := GenS3{F0: -9223372036854775808, F1: v15}
                _ = v67
            default:
                v15 += (v15 & v15)
            }
        }
    case 9223372036854775807:
        delete(v3, v16)
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        for fi69 := 0; fi69 < 7; fi69++ {
            var v70 uint16 = 2
            v71 := GenS9{F0: (v5 * 42), F1: 559}
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
