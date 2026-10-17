package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint8) uint8 {
    var v5 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    var v8 uint8 = (p0 | (p0 | p0))
    v9 := GenS1{F0: int16(v8), F1: -1}
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 uint64, p1 int) uint32 {
    v10 := GenS2{F0: 187, F1: 220, F2: -669}
    v11 := []int{127, (p1 ^ p1)}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 bool = false
    var v2 uint = 246
    var v3 int16 = 1000
    var v4 int64 = 0
    v12 := GenS3{F0: 701}
    if v1 {
        v13 := []int{498, 584, 346}
        var v14 int8 = -1
        _ = len(v13)
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 2147483648
        for fi16 := 0; fi16 < 15; fi16++ {
            if ((255 == 255) && false) {
                var v17 float64 = (1.5 - (1e10 - 1e10))
                if ((^(v15)) >= (v15 | 0)) {
                    v18 := []int{}
                    _ = v18
                }
                _ = v17
            } else {
                fmt.Println(v15, v2)
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := GenS3{F0: v2}
                switch "foo bar" {
                case "foo bar":
                    switch "hello" {
                    case "\n\t":
                        var v21 float32 = (-2.25 - 0.1)
                        _ = v21
                    case "hello":
                        switch (fi16 ^ (-(fi16))) {
                        case 0:
                            if ((v15 / (((v15 * v15)) | 1)) != genFn1(2, fi16)) {
                                v2 -= (v2 * 1)
                                v22 := GenS4{F0: 73}
                                switch -9223372036854775808 {
                                case 32767:
                                    if ("foo bar" >= "\n\t") {
                                        v3 = v3
                                        v23 := GenS1{F0: 399, F1: (v4 / ((v4) | 1))}
                                        _ = v23
                                    }
                                default:
                                    v24 := GenS2{F0: v3, F1: 380, F2: v3}
                                    _ = v24
                                }
                                var v25 uint64 = 18446744073709551615
                                _ = v22
                                _ = v25
                            }
                        case -1:
                            fmt.Println(v2)
                        case 32768:
                            var v26 int32 = int32(fi16)
                            _ = v26
                        }
                    case "a":
                        switch (("hello" + "a") + ("foo bar" + ("a" + "a"))) {
                        case "foo bar":
                            v3 -= -1
                        case "a":
                            fmt.Println(v4, fi19, fi19)
                        default:
                            var v27 uint64 = 340
                            _ = v27
                        }
                    default:
                        if (!v1) {
                            v28 := GenS1{F0: (v3 - (v3 ^ v3)), F1: -612}
                            var v29 string = ("a" + "a")
                            var v30 bool = false
                            if (((((((((0.0 - 0.0) * (1e10 / -2.25)) - (0.0 + 0.1)) * 0.0) / ((-2.25 + 1.5) / (3.14159 + 1e10))) / 1.5) - (1.5 + (float64(-2.25) - 1.5))) / 1.0) <= 1e10) {
                                v3 -= v3
                                v3 += (v3 & v3)
                            } else {
                                var v31 uint = v2
                                _ = v31
                            }
                            if (127 == 32) {
                                var v32 int64 = (v4 * (-(v4)))
                                if (65 > 255) {
                                    v15 += v15
                                } else {
                                    if (v30 && true) {
                                        if ((42 <= (^(v2))) && (!(v1 || true))) {
                                            v29 = ((("" + (v29 + "")) + v29) + "")
                                            v15 += 254
                                            v32 = -116
                                        }
                                        var v33 int64 = (^(v32))
                                        _ = v33
                                    } else {
                                        if v1 {
                                            v34 := GenS4{F0: genFn0(128)}
                                            var v35 uint64 = 4294967295
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if (v30 || v1) {
                                                v15 -= (^(v15))
                                            }
                                            var v36 bool = v1
                                            _ = v36
                                        }
                                        v37 := GenS5{F0: -128, F1: 42}
                                        v3 = v3
                                        _ = v37
                                    }
                                    v38 := []int{fi16, fi16, 516}
                                    _ = v38
                                }
                                var v39 int8 = 31
                                _ = v32
                                _ = v39
                            } else {
                                v40 := GenS6{F0: v4, F1: 42, F2: 68}
                                v41 := []int{(fi16 * (fi19 % (((fi16 & (fi19 & fi16))) | 1))), fi16, fi16}
                                _ = v40
                                _ = v41
                            }
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            v42 := []int{(fi16 ^ 32768), (fi16 + (fi16 >> ((v2) & 63))), (fi16 >> ((829) & 63))}
                            if v1 {
                                v2 -= (v2 + (v2 + (v2 << ((v2) & 63))))
                                if len(v42) >= 2 {
                                    v42 = v42[1:len(v42)]
                                }
                                v43 := GenS1{F0: 1, F1: (v4 << (((v2 * 927)) & 63))}
                                _ = v43
                            }
                            _ = v42
                        }
                    }
                default:
                    var v44 int16 = (v3 << (((v2 - v2)) & 63))
                    _ = v44
                }
                v45 := []int{}
                var v46 uint64 = 851
                _ = fi19
                _ = v20
                _ = v45
                _ = v46
            }
            _ = fi16
        }
        _ = v15
    }
    var v47 int64 = 660
    v48 := []int{782, 292}
    for fi49 := 0; fi49 < 7; fi49++ {
        var v50 bool = (!false)
        if v1 {
            if v50 {
                switch ("foo bar" + "hello") {
                case "a":
                    v51 := GenS7{F0: "hello", F1: 14, F2: 42}
                    _ = v51
                case "hello":
                    v52 := GenS8{F0: (v4 & (v47 & int64(fi49)))}
                    _ = v52
                }
                v53 := GenS9{F0: v2, F1: 190}
                _ = v53
            }
            var v54 int16 = 42
            _ = v54
        } else {
            v47 += (-(v4))
            fmt.Println(fi49, v47, v2)
            for fi55 := 0; fi55 < 13; fi55++ {
                if len(v48) > 0 {
                    v48[0] = fi49
                }
                v56 := GenS4{F0: 2}
                _ = fi55
                _ = v56
            }
        }
        _ = fi49
        _ = v50
    }
    if true {
        var v57 string = "a"
        var v58 float64 = ((((1e10 * 3.14159) + ((1.5 * 1e10) / (1.5 + float64(-2.25)))) * (-2.25 + (1e10 - 0.1))) + (0.0 / (3.14159 / (-2.25 / 3.14159))))
        _ = v57
        _ = v58
    }
    if len(v48) > 0 {
        v48[0] = 128
    }
    v59 := GenS7{F0: ("\n\t" + "a"), F1: 0, F2: genFn0(42)}
    switch -2147483648 {
    case 127:
        var v60 uint64 = 42
        _ = v60
    case -697:
        var v61 float32 = 3.14159
        _ = v61
    default:
        switch 32768 {
        case -186:
            v62 := GenS7{F0: "hello", F1: 1, F2: 2}
            _ = v62
        }
    }
    fmt.Println(v47)
    for fi63 := 0; fi63 < 18; fi63++ {
        for fi64 := 0; fi64 < 5; fi64++ {
            v47 += 1
            var v65 uint32 = 756
            _ = fi64
            _ = v65
        }
        var v66 uint64 = 811
        _ = fi63
        _ = v66
    }
    if false {
        var v67 uint16 = 577
        v2 = v2
        _ = v67
    } else {
        v47 = (v47 * 42)
        var v68 uint32 = 2147483647
        _ = v68
    }
    var v69 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v59
    _ = v69
}
