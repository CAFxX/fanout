package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 string; F2 int16 }
type GenS7 struct { F0 string; F1 bool; F2 int64 }
func genFn0(p0 int8, p1 uint8) uint {
    fmt.Println(p0, p1)
    if (!false) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(fi6 % ((42) | 1))}
            v7 = append(v7, fi6)
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            if true {
                if true {
                    var v9 float64 = (0.0 / 1.0)
                    _ = v9
                } else {
                    v10 := []int{4294967295, 685}
                    p1 = (p1 / ((p1) | 1))
                    _ = v10
                }
                v11 := GenS1{F0: 2147483647, F1: 2147483647}
                _ = v11
            } else {
                var v12 uint64 = 4294967295
                _ = v12
            }
            _ = fi8
        }
    }
    if false {
        var v13 float64 = ((((((0.1 / ((1.5 - (1.5 * (1e10 * 1.5))) - 0.0)) + 1.5) / -2.25) + 1.5) - 0.0) / 1.0)
        p0 = 127
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint8 = 223
            var v17 uint32 = 4294967295
            if (((!true) && false) && ("a" > (("hello" + ("foo bar" + ("a" + "\n\t"))) + ("hello" + "\n\t")))) {
                p0 = (p0 / (((-(p0))) | 1))
                v18 := []int{}
                var v19 bool = (-9223372036854775808 <= 1)
                _ = v18
                _ = v19
            } else {
                fmt.Println(fi14, v17, p1)
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        var v20 uint64 = 1
        v21 := []int{fi14, 32768}
        _ = fi14
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    return 526
}
func genFn1(p0 uint8, p1 bool, p2 int16) int64 {
    for fi22 := 0; fi22 < 10; fi22++ {
        if (!("a" < "\n\t")) {
            v23 := []int{fi22}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            var v24 bool = p1
            _ = v23
            _ = v24
        } else {
            p2 = 1
        }
        _ = fi22
    }
    if (-194 > -2147483648) {
        var v25 uint32 = 4294967295
        if ((v25 << ((65535) & 63)) >= (v25 ^ 2)) {
            var v26 string = ((("a" + ("\n\t" + "a")) + ("hello" + ("" + ("a" + "a")))) + "")
            _ = v26
        }
        _ = v25
    }
    if (true && (1e10 >= 1.5)) {
        if p1 {
            v27 := GenS2{F0: 974}
            fmt.Println(p1, p2, p2)
            v28 := GenS3{F0: 1, F1: 42}
            _ = v27
            _ = v28
        }
        var v29 uint32 = 725
        var v30 float32 = 0.0
        _ = v29
        _ = v30
    } else {
        fmt.Println(p0, p0)
        v31 := GenS4{F0: 4294967295, F1: 0}
        var v32 uint16 = 255
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    var v5 string = "foo bar"
    v33 := GenS2{F0: 1}
    var v34 int32 = 42
    for fi35 := 0; fi35 < 14; fi35++ {
        switch fi35 {
        case 9223372036854775807:
            var v36 int16 = int16(v34)
            _ = v36
        case 1:
            switch "" {
            case "a":
                v34 = (v34 + 2147483647)
            case "\n\t":
                v37 := GenS2{F0: 0}
                _ = v37
            case "":
                if (genFn1(2, false, 42) < 9223372036854775807) {
                    v38 := []int{}
                    v39 := GenS5{F0: (v34 & (v34 | (v34 >> ((64) & 63))))}
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 uint64 = (v2 % (((^(v2))) | 1))
                        v42 := GenS1{F0: 4294967295, F1: 2147483647}
                        if false {
                            if true {
                                v41 = 0
                                v43 := GenS4{F0: 0, F1: 651}
                                switch -9223372036854775808 {
                                case 318:
                                    v2 = 0
                                }
                                _ = v43
                            } else {
                                v41 += (v2 | v41)
                            }
                            var v44 int64 = 9223372036854775807
                            _ = v44
                        }
                        v3[v5] = 571
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v38
                    _ = v39
                }
            default:
                var v45 uint16 = 32767
                _ = v45
            }
        default:
            var v46 string = ("foo bar" + v5)
            _ = v46
        }
        _ = fi35
    }
    switch v5 {
    case "hello":
        fmt.Println(v5)
    case "a":
        var v47 uint64 = 42
        _ = v47
    default:
        v5 = "\n\t"
    }
    switch v5 {
    case "\n\t":
        if (!(0 < 0)) {
            v4 = (1.5 / (-2.25 + v4))
            for fi48 := 0; fi48 < 20; fi48++ {
                var v49 uint32 = 1
                _ = fi48
                _ = v49
            }
        } else {
            fmt.Println(v5, v34, v5)
            delete(v3, v5)
        }
    case "hello":
        for fi50 := 0; fi50 < 13; fi50++ {
            v34 = v34
            _ = fi50
        }
    }
    fmt.Println(v2, v34, v2)
    var v51 int8 = 42
    fmt.Println(v34, v34, v51)
    v52 := GenS6{F0: v2, F1: v5, F2: 32767}
    v4 = float32(float64((((0.1 + v4) + ((((v4 + 1e10) + float32(3.14159)) + 0.1) / -2.25)) + 0.1)))
    v53 := v3[v5]
    if ((-2.25 + ((float64(0.1) / ((((float64(v4) + 0.1) * 3.14159) + -2.25) / (0.0 - (3.14159 - 0.1)))) - 3.14159)) == (1.5 / float64(v4))) {
        var v54 uint8 = 0
        var v55 uint16 = 42
        _ = v54
        _ = v55
    }
    fmt.Println(v34)
    switch 42 {
    case 612:
        v5 = ((v5 + ((v5 + (v5 + (("" + v5) + v5))) + v5)) + v5)
    case 4294967295:
        if ((1.5 == v4) || true) {
            v3[(v5 + v5)] = (v53 ^ 2147483647)
            fmt.Println(v5)
            if (v5 <= "a") {
                switch 521 {
                case 877:
                    var v56 int64 = genFn1(127, true, -32768)
                    _ = v56
                case -65:
                    v57 := []int{-1, -196, 9223372036854775807}
                    _ = v57
                case -2147483648:
                    v34 = (v34 / (((v34 + (v34 * (v34 & (-(v34)))))) | 1))
                default:
                    for fi58 := 0; fi58 < 17; fi58++ {
                        var v59 uint16 = 2
                        v1[fi58] = (v53 & v53)
                        var v60 int = 1
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
            }
        } else {
            var v61 int16 = -301
            v62 := GenS1{F0: 42, F1: 1}
            v4 += v4
            _ = v61
            _ = v62
        }
    case 308:
        switch v5 {
        case "hello":
            fmt.Println(v34)
        case "\n\t":
            for fi63 := 0; fi63 < 18; fi63++ {
                v5 = v5
                _ = fi63
            }
        }
    default:
        switch (((v5 + v5) + "hello") + (((v5 + v5) + v5) + v5)) {
        case "foo bar":
            v53 = v53
        case "":
            switch 657 {
            case 32767:
                for fi64 := 0; fi64 < 14; fi64++ {
                    v34 -= v34
                    _ = fi64
                }
            case 1:
                var v65 uint8 = 180
                _ = v65
            case 804:
                if false {
                    for fi66 := 0; fi66 < 7; fi66++ {
                        switch fi66 {
                        case 1:
                            v51 = 0
                        default:
                            var v67 int16 = int16(fi66)
                            _ = v67
                        }
                        _ = fi66
                    }
                } else {
                    v68 := GenS5{F0: (v34 % (((v34 * 42)) | 1))}
                    v69 := []int{1, 32768, 128, 42}
                    _ = v68
                    _ = v69
                }
            default:
                if (int64(v2) == 880) {
                    var v70 uint32 = v53
                    switch 32767 {
                    case 128:
                        fmt.Println(v70, v51)
                    }
                    v70 += v70
                    _ = v70
                } else {
                    if true {
                        v71 := GenS5{F0: -45}
                        switch 42 {
                        case -893:
                            v72 := []int{-770}
                            _ = v72
                        case -2147483648:
                            delete(v3, "")
                        case 256:
                            if (true && false) {
                                if false {
                                    var v73 string = v5
                                    _ = v73
                                } else {
                                    var v74 uint8 = 248
                                    v53 = v53
                                    _ = v74
                                }
                                v34 -= 463
                                for fi75 := 0; fi75 < 12; fi75++ {
                                    v76 := []int{(fi75 ^ (fi75 & 0))}
                                    delete(v3, ("hello" + v5))
                                    _ = fi75
                                    _ = v76
                                }
                            } else {
                                fmt.Println(v4, v2)
                            }
                        default:
                            var v77 int64 = -525
                            _ = v77
                        }
                        _ = v71
                    } else {
                        if false {
                            v5 = (("hello" + v5) + ("a" + (("\n\t" + v5) + (v5 + "\n\t"))))
                            if (!((575 < 9223372036854775807) && false)) {
                                _ = len(v3)
                            }
                        } else {
                            var v78 uint = 544
                            _ = v78
                        }
                    }
                }
            }
        }
    }
    v79 := GenS7{F0: (v5 + ("hello" + v5)), F1: (v2 >= (v2 ^ 141)), F2: 42}
    v2 = v2
    var v80 int8 = 34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v79
    _ = v80
}
