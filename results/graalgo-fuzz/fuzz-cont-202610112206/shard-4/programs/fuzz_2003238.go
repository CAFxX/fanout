package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int8; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 bool; F2 int8 }
type GenS7 struct { F0 bool; F1 int; F2 int }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint8; F1 int; F2 int32 }
func genFn0(p0 int, p1 uint16, p2 uint16) int32 {
    var v5 uint16 = (^(p1))
    v5 = (v5 % (((^(p2))) | 1))
    var v6 int64 = int64(p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(p1, p1)
        p0 = (-(p0))
        if false {
            p2 -= 255
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -597
}
func genFn1() int16 {
    if false {
        v8 := GenS1{F0: 0, F1: 85}
        v9 := []int{}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint = 984
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        if true {
            var v13 int16 = 42
            _ = v13
        } else {
            switch "foo bar" {
            case "foo bar":
                switch 32768 {
                case 2147483647:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            switch ("\n\t" + "a") {
                            case "foo bar":
                                v16 := GenS1{F0: 1, F1: 925}
                                _ = v16
                            case "\n\t":
                                var v17 int8 = -1
                                _ = v17
                            default:
                                v18 := []int{(fi14 << ((65535) & 63))}
                                _ = v18
                            }
                            var v19 string = ""
                            v20 := GenS1{F0: 42, F1: 1}
                            _ = fi15
                            _ = v19
                            _ = v20
                        }
                        var v21 int64 = 9223372036854775807
                        _ = fi14
                        _ = v21
                    }
                case 127:
                    var v22 int = 32767
                    _ = v22
                case 753:
                    v23 := GenS2{F0: 519, F1: false, F2: 1}
                    _ = v23
                default:
                    v24 := GenS3{F0: 589}
                    _ = v24
                }
            }
        }
    }
    var v25 uint64 = 107
    v26 := GenS2{F0: 4294967295, F1: ((1.5 > (1e10 * -2.25)) || (0.0 > float64(-2.25))), F2: 0}
    _ = v25
    _ = v26
    return -32768
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 819
    var v3 uint = 4294967295
    var v4 float32 = float32((1.5 - 3.14159))
    v27 := GenS2{F0: (v2 % ((42) | 1)), F1: true, F2: 341}
    fmt.Println(v3, v3, v3)
    v28 := []int{2147483647}
    v29 := GenS2{F0: 2147483648, F1: false, F2: v2}
    var v30 uint = (v3 % (((v3 / (((v3 | 1)) | 1))) | 1))
    var v31 uint64 = 456
    fmt.Println(v30, v2, v4)
    switch 127 {
    case 619:
        var v32 int32 = 2147483647
        _ = v32
    case 262:
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS4{F0: 45, F1: -121, F2: false}
            _ = fi33
            _ = v34
        }
    case 4294967295:
        fmt.Println(v31, v30, v3)
    default:
        var v35 uint = 18446744073709551615
        _ = v35
    }
    if (!(!(false && false))) {
        if (2147483648 <= (v2 * 2147483647)) {
            switch 192 {
            case 128:
                if (float32(0.0) < float32(float64(0.0))) {
                    var v36 int32 = 42
                    fmt.Println(v1, v1, v30)
                    _ = v36
                } else {
                    fmt.Println(v31)
                    switch "" {
                    case "foo bar":
                        for fi37 := 0; fi37 < 19; fi37++ {
                            v1 = v1
                            if false {
                                v30 -= (^(v30))
                                if false {
                                    if (42 == genFn1()) {
                                        for fi38 := 0; fi38 < 12; fi38++ {
                                            v31 = v31
                                            var v39 int = (^(fi37))
                                            v40 := GenS5{F0: (v3 >> ((255) & 63))}
                                            _ = fi38
                                            _ = v39
                                            _ = v40
                                        }
                                        var v41 uint = uint(v2)
                                        fmt.Println(v1)
                                        _ = v41
                                    } else {
                                        _ = len(v28)
                                    }
                                    if len(v28) > 0 {
                                        v28[0] = (^(fi37))
                                    }
                                    v28 = append(v28, (fi37 << ((18446744073709551615) & 63)))
                                }
                                if ((fi37 < (fi37 + fi37)) || false) {
                                    if (true || (4294967295 < 138)) {
                                        v3 = (v3 | uint(v2))
                                    } else {
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            v3 = (^(v30))
                                        case "":
                                            v1 += genFn1()
                                        case "hello":
                                            fmt.Println(fi37)
                                        default:
                                            if len(v28) > 0 {
                                                v28[0] = (fi37 + (fi37 | (fi37 - fi37)))
                                            }
                                        }
                                        var v42 float32 = (v4 * (v4 / v4))
                                        _ = v42
                                    }
                                } else {
                                    fmt.Println(v2)
                                    if (!(1 <= 255)) {
                                        v28 = append(v28, fi37)
                                        var v43 uint32 = uint32(v31)
                                        v1 -= v1
                                        _ = v43
                                    }
                                }
                            } else {
                                v1 = int16(v2)
                                if false {
                                    _ = len(v28)
                                    for fi44 := 0; fi44 < 19; fi44++ {
                                        v4 = v4
                                        v4 -= (v4 - ((float32((0.0 / 1e10)) + (v4 / 1e10)) - v4))
                                        switch fi44 {
                                        case 0:
                                            v2 = v2
                                        }
                                        _ = fi44
                                    }
                                } else {
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        var v46 int16 = -1
                                        _ = fi45
                                        _ = v46
                                    }
                                    v31 -= 18446744073709551615
                                }
                                if (13 >= -1) {
                                    v3 += (v3 / ((v30) | 1))
                                    v47 := GenS6{F0: -122, F1: (!true), F2: -128}
                                    v4 = 0.0
                                    _ = v47
                                } else {
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        v49 := GenS5{F0: v30}
                                        var v50 int = (fi37 * (fi37 << ((v3) & 63)))
                                        if true {
                                            v2 -= 0
                                            var v51 bool = false
                                            _ = v51
                                        } else {
                                            if (!((((float32(float64(v4)) * ((v4 * float32((((float64(float32(0.0)) / (-2.25 - 3.14159)) - 3.14159) * (0.1 / ((3.14159 * (1e10 - (-2.25 * 1.5))) / (((1e10 - 1.5) + ((0.0 - 1.5) + 1.5)) * (0.0 - 3.14159))))))) / float32(3.14159))) * (v4 + v4)) >= 0.1) || (!false))) {
                                                if false {
                                                    if len(v28) > 0 {
                                                        v28[0] = 127
                                                    }
                                                    v2 = v2
                                                } else {
                                                    v52 := GenS3{F0: 9223372036854775807}
                                                    var v53 int32 = int32(v3)
                                                    var v54 uint16 = 573
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                var v55 int16 = v1
                                                _ = v55
                                            }
                                            v30 += uint(v31)
                                            v56 := GenS1{F0: 127, F1: (v2 >> ((v3) & 63))}
                                            _ = v56
                                        }
                                        v57 := GenS7{F0: ((!false) || true), F1: 613, F2: fi37}
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                        _ = v57
                                    }
                                }
                            }
                            if true {
                                v31 += (v31 & uint64(v2))
                                v28 = append(v28, 4294967295)
                                v58 := GenS7{F0: ((255 != 1) || (v2 == v2)), F1: -9223372036854775808, F2: fi37}
                                _ = v58
                            } else {
                                if false {
                                    var v59 bool = false
                                    _ = v59
                                } else {
                                    v28 = append(v28, 1)
                                }
                                v4 = (float32(((-2.25 + (-2.25 * float64(-2.25))) - -2.25)) * float32((0.0 / ((-2.25 * ((float64(v4) + (0.0 * 1e10)) * (-2.25 + float64(v4)))) * ((((0.0 * float64(((3.14159 / v4) * (float32(3.14159) * 1.5)))) / -2.25) * (((0.1 * 0.1) + 1e10) * (float64((((((v4 + 0.0) / v4) * (v4 + 3.14159)) + (v4 / ((v4 + (((v4 + v4) * (1.5 + float32(0.0))) * 0.1)) + 0.1))) + v4)) * 3.14159))) - (1e10 - ((3.14159 - 1e10) * -2.25)))))))
                                fmt.Println(v31, v3)
                            }
                            _ = fi37
                        }
                    case "a":
                        v60 := GenS4{F0: genFn1(), F1: 127, F2: (0.1 != 3.14159)}
                        _ = v60
                    }
                    fmt.Println(v2, v3, v3)
                }
            default:
                fmt.Println(v30)
            }
            if (51 == 232) {
                v61 := GenS8{F0: v3, F1: ("a" + "\n\t")}
                var v62 int = 32768
                _ = v61
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 19; fi63++ {
            var v64 int64 = -1
            _ = fi63
            _ = v64
        }
        var v65 uint64 = (v31 - v31)
        _ = v65
    }
    v1 += -1
    v66 := GenS9{F0: 2, F1: 0, F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v66
}
