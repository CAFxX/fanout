package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int8 }
func genFn0(p0 int8, p1 int32, p2 int32) int32 {
    v7 := GenS1{F0: 0}
    v8 := []int{4294967295, 108, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 >> ((328) & 63))
}
func genFn1(p0 int32, p1 int32) uint32 {
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            if false {
                p0 -= (^(p0))
                p0 = (^(p0))
                var v11 uint16 = 42
                switch fi9 {
                case -2147483648:
                    fmt.Println(p0, fi10, v11)
                default:
                    if (745 > 32767) {
                        v12 := GenS2{F0: 0}
                        fmt.Println(fi9)
                        p0 = p0
                        _ = v12
                    }
                }
                _ = v11
            }
            var v13 float32 = 1e10
            _ = fi10
            _ = v13
        }
        fmt.Println(p1, fi9)
        for fi14 := 0; fi14 < 1; fi14++ {
            p1 = p0
            if (-4 != 124) {
                var v15 int8 = 42
                if false {
                    var v16 int64 = -9223372036854775808
                    _ = v16
                }
                _ = v15
            }
            var v17 uint64 = 2
            _ = fi14
            _ = v17
        }
        _ = fi9
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 bool = false
        var v20 float64 = ((1.5 * float64(0.1)) * 0.1)
        p0 = (^(p1))
        _ = fi18
        _ = v19
        _ = v20
    }
    p1 -= genFn0(92, p0, 2147483647)
    var v21 int8 = 34
    _ = p0
    _ = p1
    _ = v21
    return 65535
}

func main() {
    var v1 uint = 743
    var v2 uint8 = 1
    var v3 int32 = -1
    var v4 uint8 = v2
    v5 := []string{"\n\t", (("\n\t" + (("\n\t" + "hello") + (("\n\t" + "a") + "\n\t"))) + "foo bar"), "a", ""}
    v6 := make(map[int]string)
    delete(v6, -918)
    if false {
        v6[2147483647] = ((((("" + (("hello" + "\n\t") + "")) + ("foo bar" + "a")) + (("a" + "\n\t") + "a")) + "a") + "hello")
        v22 := GenS1{F0: 0}
        _ = v22
    }
    v3 -= (v3 % (((v3 | 0)) | 1))
    fmt.Println(v3, v2, v2)
    for fi23 := 0; fi23 < 18; fi23++ {
        if true {
            fmt.Println(v3)
            for fi24 := 0; fi24 < 2; fi24++ {
                fmt.Println(fi24, v3)
                _ = fi24
            }
        }
        _ = len(v6)
        _ = fi23
    }
    var v25 bool = true
    if (42 <= 0) {
        if (0.1 > ((1e10 + 0.1) * 3.14159)) {
            if (0 > 223) {
                fmt.Println(v25)
                v1 -= (v1 + (v1 * v1))
            }
            var v26 bool = ((-649 != 42) || v25)
            v27 := GenS2{F0: (v4 | 0)}
            _ = v26
            _ = v27
        } else {
            if (!(!v25)) {
                v5 = append(v5, (("a" + ("hello" + (("foo bar" + ("" + (("" + "hello") + "foo bar"))) + "a"))) + ""))
                if ((v4 - 255) >= v4) {
                    if v25 {
                        switch "\n\t" {
                        case "":
                            v28 := v6[42]
                            _ = v28
                        }
                    }
                    if false {
                        v3 = 1
                        var v29 int64 = 0
                        _ = v29
                    }
                }
                var v30 uint = 1
                _ = v30
            }
            for fi31 := 0; fi31 < 8; fi31++ {
                for fi32 := 0; fi32 < 12; fi32++ {
                    if v25 {
                        v33 := GenS3{F0: genFn0(-74, -640, (v3 * (v3 / ((v3) | 1)))), F1: genFn0(1, (^(v3)), v3)}
                        fmt.Println(v25, v25)
                        var v34 uint8 = (^(v4))
                        _ = v33
                        _ = v34
                    }
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    var v35 uint = 1
                    _ = fi32
                    _ = v35
                }
                v36 := GenS2{F0: (v2 / ((v2) | 1))}
                _ = len(v6)
                _ = fi31
                _ = v36
            }
            for fi37 := 0; fi37 < 18; fi37++ {
                var v38 uint64 = 42
                _ = fi37
                _ = v38
            }
        }
    }
    v39 := GenS4{F0: "a"}
    v40 := v6[-466]
    switch 9223372036854775807 {
    case 127:
        var v41 uint16 = 42
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        v2 -= (v2 | v2)
        if ((false && ((v40 != v40) && (42 == 65535))) || (1 > -14)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            switch 477 {
            case -9223372036854775808:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 9223372036854775807:
                var v43 int8 = int8(v1)
                _ = v43
            default:
                v44 := GenS3{F0: v3, F1: (v3 % (((v3 << ((v1) & 63))) | 1))}
                _ = v44
            }
            v5 = append(v5, v40)
        }
        _ = fi42
    }
    var v45 float64 = 0.0
    if false {
        delete(v6, -482)
        var v46 int = -138
        _ = v46
    } else {
        if false {
            fmt.Println(v40, v2)
            for fi47 := 0; fi47 < 16; fi47++ {
                if v25 {
                    v48 := GenS5{F0: (v3 & v3), F1: -22}
                    if (!(false || true)) {
                        _ = len(v5)
                    }
                    for fi49 := 0; fi49 < 16; fi49++ {
                        if (false || ((genFn1((v3 / (((v3 * 1)) | 1)), (v3 + (v3 >> (((v1 + (v1 - v1))) & 63)))) >= 2147483647) || true)) {
                            if (1e10 >= 0.0) {
                                v1 -= v1
                                var v50 float64 = (v45 / v45)
                                switch 42 {
                                case 128:
                                    var v51 uint = (^(v1))
                                    _ = v51
                                case 4294967295:
                                    if (!(42 != 0)) {
                                        v50 -= ((((v45 - 0.0) / ((v50 - -2.25) / ((v50 * v50) - float64(1.5)))) * v45) / v45)
                                        switch v40 {
                                        case "":
                                            v52 := v6[(fi47 / (((fi49 ^ (fi49 / ((fi49) | 1)))) | 1))]
                                            _ = v52
                                        case "\n\t":
                                            v3 -= v3
                                        default:
                                            var v53 uint = v1
                                            _ = v53
                                        }
                                    } else {
                                        fmt.Println(v1, v25, v50)
                                        v54 := GenS2{F0: (v4 & 43)}
                                        var v55 float64 = (((1e10 * 3.14159) / (v45 - v50)) / (v50 - (float64(-2.25) / v50)))
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                if v25 {
                                    v56 := v6[fi47]
                                    _ = v56
                                }
                                _ = v50
                            } else {
                                v25 = v25
                                if ((v25 && v25) || (!(!v25))) {
                                    if false {
                                        _ = len(v5)
                                        var v57 uint64 = 2
                                        v25 = (!(42 <= 65535))
                                        _ = v57
                                    }
                                    v3 = 0
                                    var v58 int64 = 883
                                    if false {
                                        v40 = ((((v40 + "foo bar") + v40) + (v40 + v40)) + (((v40 + (v40 + ((v40 + v40) + v40))) + "") + ((v40 + "") + (v40 + (((v40 + v40) + "\n\t") + (("\n\t" + "hello") + ((v40 + ((("" + v40) + (v40 + v40)) + ("\n\t" + "foo bar"))) + v40)))))))
                                        v5 = append(v5, (v40 + v40))
                                    } else {
                                        v59 := GenS2{F0: v2}
                                        _ = v59
                                    }
                                    _ = v58
                                } else {
                                    var v60 int = -949
                                    fmt.Println(v40, v60, v1)
                                    _ = v60
                                }
                            }
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                            v45 -= v45
                            v1 -= (v1 / ((271) | 1))
                            v25 = true
                        }
                        v5 = append(v5, "a")
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        _ = fi49
                    }
                    _ = v48
                } else {
                    v61 := GenS3{F0: genFn0(-29, v3, genFn0(108, (v3 & v3), -21)), F1: v3}
                    for fi62 := 0; fi62 < 5; fi62++ {
                        fmt.Println(v1)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v25 = ((v25 || (1e10 < 3.14159)) && (1 == 2))
                        v6[fi62] = v40
                        _ = fi62
                    }
                    switch fi47 {
                    case 32767:
                        v63 := v6[fi47]
                        _ = v63
                    }
                    _ = v61
                }
                for fi64 := 0; fi64 < 16; fi64++ {
                    fmt.Println(v3, fi47)
                    _ = fi64
                }
                var v65 bool = (!v25)
                _ = fi47
                _ = v65
            }
        } else {
            v45 = 0.1
            fmt.Println(v4, v45)
            var v66 int8 = 0
            _ = v66
        }
        v67 := GenS3{F0: (v3 - -2147483648), F1: v3}
        _ = len(v6)
        _ = v67
    }
    var v68 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
    _ = v68
}
