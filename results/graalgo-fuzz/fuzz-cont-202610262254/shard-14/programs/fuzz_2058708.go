package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
func genFn0() uint16 {
    var v6 uint16 = 2
    var v7 int16 = 0
    fmt.Println(v6)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v6 = v6
            _ = fi9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            v6 -= (v6 & v6)
            v11 := GenS1{F0: "hello", F1: 2147483647}
            _ = fi10
            _ = v11
        }
        fmt.Println(v7, fi8)
        _ = fi8
    }
    v12 := GenS2{F0: 0, F1: "a", F2: (v6 - 396)}
    _ = v12
    _ = v6
    _ = v7
    return (v6 / (((v6 << ((4294967295) & 63))) | 1))
}
func genFn1(p0 uint32, p1 int8, p2 uint32) string {
    var v13 float32 = (1.5 * (1e10 * 0.1))
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int16 = 967
        switch fi14 {
        case 509:
            v16 := GenS2{F0: -1, F1: "foo bar", F2: genFn0()}
            _ = v16
        case 42:
            var v17 uint32 = 4294967295
            _ = v17
        case 989:
            fmt.Println(v15, v13, v15)
        default:
            if (p1 != (p1 % (((p1 ^ p1)) | 1))) {
                v15 = (v15 >> ((255) & 63))
                switch (fi14 % (((fi14 >> ((4294967295) & 63))) | 1)) {
                case 1:
                    var v18 float32 = (float32(0.1) - (3.14159 + (((v13 - v13) * v13) + v13)))
                    _ = v18
                case -115:
                    fmt.Println(v13, p2)
                default:
                    if (!((!true) && (((!true) || ((v15 << ((0) & 63)) < v15)) || true))) {
                        var v19 uint16 = 65535
                        _ = v19
                    }
                }
                if (42 < 951) {
                    switch ("a" + "") {
                    case "":
                        fmt.Println(fi14, v13, fi14)
                    default:
                        for fi20 := 0; fi20 < 11; fi20++ {
                            if true {
                                p0 = 0
                                v21 := GenS1{F0: "a", F1: 0}
                                switch fi20 {
                                case -1:
                                    if false {
                                        p0 -= p0
                                    }
                                case 0:
                                    var v22 float64 = 0.1
                                    _ = v22
                                case 4294967295:
                                    v15 += (v15 << ((2) & 63))
                                default:
                                    var v23 float32 = v13
                                    _ = v23
                                }
                                var v24 int16 = (v15 + (v15 << ((uint(p0)) & 63)))
                                v25 := []int{(fi20 >> ((0) & 63)), fi20, fi14, (fi20 % ((fi20) | 1))}
                                _ = v21
                                _ = v24
                                _ = v25
                            } else {
                                var v26 bool = true
                                v27 := []int{fi20, (fi20 ^ (-(fi20))), (fi14 | (fi14 * (fi20 << ((2) & 63)))), (fi14 - 4294967295)}
                                _ = v26
                                _ = v27
                            }
                            p1 = 81
                            _ = fi20
                        }
                    }
                } else {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 int32 = 42
                        var v30 int16 = 42
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
            } else {
                p2 -= p2
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 int = (fi31 << ((1) & 63))
                    v32 -= (v32 ^ fi31)
                    if (false && (1e10 > v13)) {
                        var v33 string = "\n\t"
                        v34 := GenS2{F0: 329, F1: "foo bar", F2: 0}
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                v35 := []int{fi14, (fi14 >> ((552) & 63))}
                _ = v35
            }
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ("" + "\n\t")
}

func main() {
    var v1 float32 = -2.25
    var v2 string = ""
    var v3 float64 = float64(float32(float64(v1)))
    v4 := []string{((((v2 + ((v2 + "a") + (v2 + (v2 + v2)))) + v2) + (v2 + (v2 + v2))) + ""), v2}
    var v5 uint = 42
    v36 := GenS1{F0: "\n\t", F1: 2147483647}
    var v37 float32 = (3.14159 + float32(v3))
    v3 += ((((float64(float32((1e10 + (((v3 * ((v3 * float64(3.14159)) - v3)) + (((-2.25 - (v3 / 1.0)) - (3.14159 + float64(v37))) / (float64(v1) * (-2.25 - (float64(v37) - v3))))) + float64(v37))))) / 1e10) * v3) - v3) * -2.25)
    v38 := GenS1{F0: (v2 + "a"), F1: 42}
    v3 += v3
    v5 -= 65535
    for fi39 := 0; fi39 < 4; fi39++ {
        if (((v37 - 3.14159) / 1.0) != (v37 + ((v37 / (0.0 - ((v1 - float32(v3)) / ((v1 + 3.14159) * (float32(v3) + -2.25))))) * v37))) {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            if ((-461 != 965) || false) {
                if (32767 >= -574) {
                    switch v2 {
                    case "":
                        v5 = 0
                    case "hello":
                        var v40 float64 = 3.14159
                        _ = v40
                    default:
                        switch ("hello" + v2) {
                        case "\n\t":
                            v5 = v5
                        case "":
                            if len(v4) > 0 {
                                v4[0] = v2
                            }
                        }
                    }
                    var v41 int8 = 80
                    _ = v41
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = v2
                }
                fmt.Println(v37, v3, fi39)
                var v42 int16 = -775
                _ = v42
            }
            v43 := GenS3{F0: 1, F1: 0}
            _ = v43
        } else {
            if true {
                fmt.Println(fi39)
                _ = len(v4)
                var v44 uint32 = 42
                _ = v44
            }
            switch "foo bar" {
            case "foo bar":
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 int = -1
                    var v47 uint = 42
                    var v48 int16 = -1
                    fmt.Println(v46, v46, v2)
                    v2 = "foo bar"
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
            case "\n\t":
                v5 = v5
            case "":
                _ = len(v4)
            default:
                var v49 int16 = -1
                _ = v49
            }
            if false {
                var v50 int64 = 124
                v50 -= 42
                switch "hello" {
                case "foo bar":
                    v50 = 860
                case "hello":
                    v1 += -2.25
                default:
                    if true {
                        var v51 float64 = (-2.25 + v3)
                        switch "" {
                        case "foo bar":
                            fmt.Println(fi39)
                        }
                        _ = v51
                    } else {
                        v3 -= ((3.14159 / (v3 + (v3 + v3))) * (1.5 / 0.1))
                        var v52 uint16 = 0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v52
                    }
                }
                _ = v50
            } else {
                v53 := GenS1{F0: "\n\t", F1: 42}
                v2 = v2
                _ = v53
            }
        }
        v3 -= 0.1
        switch (fi39 % (((fi39 / ((fi39) | 1))) | 1)) {
        case 0:
            var v54 string = v2
            _ = v54
        case -9223372036854775808:
            fmt.Println(fi39, v1)
        case 2147483647:
            var v55 uint32 = 2
            _ = v55
        }
        _ = fi39
    }
    if (42 <= 255) {
        var v56 uint64 = 1
        v57 := GenS4{F0: 0, F1: v56, F2: 42}
        _ = v56
        _ = v57
    } else {
        v37 += float32(v3)
    }
    v3 = v3
    for fi58 := 0; fi58 < 8; fi58++ {
        var v59 uint16 = 356
        v60 := GenS5{F0: 987, F1: (fi58 % (((fi58 ^ 9223372036854775807)) | 1)), F2: -587}
        _ = fi58
        _ = v59
        _ = v60
    }
    switch (v2 + v2) {
    case "hello":
        if ((true || ((1.5 + (v3 - v3)) == v3)) || (242 > genFn0())) {
            v37 += v37
        }
    case "a":
        for fi61 := 0; fi61 < 9; fi61++ {
            if ((!(-375 <= 1)) || true) {
                v1 += float32(-2.25)
            } else {
                _ = len(v4)
                fmt.Println(fi61, v3, v2)
                var v62 int64 = 0
                _ = v62
            }
            _ = fi61
        }
    }
    var v63 float64 = v3
    v2 = "\n\t"
    var v64 uint64 = 577
    var v65 uint16 = genFn0()
    v66 := GenS5{F0: 1, F1: int(v5), F2: 42}
    var v67 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
