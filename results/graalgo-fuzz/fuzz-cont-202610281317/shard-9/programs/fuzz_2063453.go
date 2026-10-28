package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 string; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64; F2 string }
func genFn0(p0 uint8) int16 {
    var v4 int16 = 1
    var v5 uint16 = 0
    p0 = p0
    var v6 int64 = 208
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1(p0 string, p1 bool, p2 uint8) int16 {
    if (((((p0 + p0) + (p0 + "\n\t")) + "a") > (((p0 + "\n\t") + ((p0 + (p0 + "foo bar")) + (p0 + ("" + "a")))) + "a")) && true) {
        p2 = 147
        p0 = (p0 + "a")
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                if (p1 && false) {
                    var v9 uint16 = 65535
                    _ = v9
                } else {
                    var v10 bool = ((!true) || (false && (((3.14159 / 0.1) / ((float64((float32(1e10) + 0.1)) * ((0.1 / 1.0) * (0.1 - 0.0))) + float64((0.1 - (3.14159 + 3.14159))))) > 1e10)))
                    _ = v10
                }
                _ = fi8
            }
            v11 := GenS1{F0: 0}
            _ = fi7
            _ = v11
        }
    }
    p2 = 1
    if (((p0 <= ((p0 + (p0 + p0)) + "a")) && (0 < 2)) && false) {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int = (fi12 ^ (fi12 & -1))
            v14 := []int{-2147483648, v13, (v13 % ((-1) | 1))}
            if true {
                if len(v14) > 0 {
                    v14[0] = (^(fi12))
                }
                var v15 uint32 = 180
                if len(v14) > 0 {
                    v14[0] = (fi12 | v13)
                }
                _ = v15
            } else {
                var v16 int8 = 43
                fmt.Println(p0, p1, v13)
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        p2 += p2
        v17 := GenS1{F0: 65535}
        _ = v17
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 1
    v2 := []string{"\n\t"}
    var v3 bool = true
    if (1 >= -986) {
        v1 -= 1000
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "hello":
        v1 = v1
    case "":
        v2 = append(v2, "")
    case "\n\t":
        var v18 int8 = 0
        _ = v18
    default:
        switch 9223372036854775807 {
        case 32767:
            var v19 int64 = 151
            _ = v19
        case -9223372036854775808:
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS2{F0: 491, F1: 4294967295, F2: 42}
                var v22 uint32 = 231
                _ = fi20
                _ = v21
                _ = v22
            }
        default:
            if true {
                switch 414 {
                case 2147483647:
                    v3 = true
                case 42:
                    var v23 uint8 = 249
                    _ = v23
                }
                if (!v3) {
                    v3 = v3
                    v1 = (^(v1))
                } else {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if (!(-128 < 53)) {
                            var v25 uint32 = 42
                            var v26 int16 = (-(v1))
                            _ = v25
                            _ = v26
                        }
                        _ = fi24
                    }
                    if (-9223372036854775808 != 575) {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            v28 := GenS2{F0: 386, F1: 65535, F2: 0}
                            v1 -= (v1 | v1)
                            fmt.Println(fi27, v3)
                            _ = fi27
                            _ = v28
                        }
                    }
                    var v29 uint16 = 32767
                    _ = v29
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    if (v3 || ((false && (false || (-281 != 0))) || (v1 >= 1000))) {
                        if (255 <= 204) {
                            v1 = v1
                            v31 := GenS3{F0: v3, F1: ("\n\t" + "\n\t"), F2: 127}
                            v32 := GenS1{F0: 918}
                            _ = v31
                            _ = v32
                        }
                        var v33 uint = 255
                        _ = v33
                    }
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    var v34 uint32 = 2147483647
                    _ = fi30
                    _ = v34
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    if (!(!v3)) {
                        var v35 float32 = 1e10
                        _ = v35
                    }
                default:
                    v2 = append(v2, "\n\t")
                }
            }
        }
    }
    v36 := GenS4{F0: 1000, F1: 42}
    v3 = v3
    v2 = append(v2, ("\n\t" + ("foo bar" + "foo bar")))
    v1 += (v1 / ((-823) | 1))
    switch 1 {
    case 32767:
        if (2147483647 < 287) {
            v37 := GenS3{F0: (-1 == -819), F1: "a", F2: 127}
            var v38 uint8 = 42
            for fi39 := 0; fi39 < 13; fi39++ {
                if v3 {
                    if ("\n\t" != "foo bar") {
                        var v40 uint32 = 1
                        switch 4294967295 {
                        case 534:
                            for fi41 := 0; fi41 < 5; fi41++ {
                                var v42 int8 = 0
                                v1 = v1
                                switch (fi41 ^ (-(fi39))) {
                                case 4294967295:
                                    var v43 int32 = 84
                                    _ = v43
                                default:
                                    v40 += (v40 + v40)
                                }
                                switch 128 {
                                case 32768:
                                    if false {
                                        var v44 int = (fi39 << ((255) & 63))
                                        var v45 float64 = (1e10 * ((float64(((((3.14159 * 0.0) - 1e10) * 0.0) * 3.14159)) * ((float64((float32(float64(3.14159)) / 1.5)) / 1e10) * (float64(-2.25) - -2.25))) - 0.0))
                                        v44 = (fi41 * 128)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        fmt.Println(v42, v1, v42)
                                        v46 := GenS5{F0: 42, F1: 879, F2: "hello"}
                                        if len(v2) > 0 {
                                            v2[0] = (("a" + ("a" + "hello")) + "a")
                                        }
                                        v1 = genFn1("", (true || v3), (v38 - (v38 | v38)))
                                        _ = v46
                                    }
                                case 286:
                                    if v3 {
                                        v2 = append(v2, ((("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + ("\n\t" + ((("" + "hello") + "a") + "hello"))))) + "a") + "a"))
                                        if (3.14159 <= (1e10 * 3.14159)) {
                                            var v47 uint = 42
                                            var v48 int8 = (-(v42))
                                            _ = v47
                                            _ = v48
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                default:
                                    switch "" {
                                    case "\n\t":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                }
                                var v49 int32 = 0
                                v50 := GenS1{F0: 0}
                                var v51 float64 = 0.0
                                var v52 uint32 = 65535
                                var v53 uint16 = 42
                                _ = fi41
                                _ = v42
                                _ = v49
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        default:
                            var v54 string = ("\n\t" + ("" + ""))
                            _ = v54
                        }
                        var v55 int16 = -1
                        _ = v40
                        _ = v55
                    }
                }
                var v56 int8 = -1
                v1 -= (v1 | (v1 >> ((2) & 63)))
                _ = fi39
                _ = v56
            }
            _ = v37
            _ = v38
        } else {
            v1 -= (^(v1))
        }
    case 0:
        v3 = (32768 <= 0)
    }
    switch "" {
    case "":
        switch -82 {
        case 9223372036854775807:
            var v57 int16 = v1
            _ = v57
        }
    case "foo bar":
        v3 = (v3 || false)
    default:
        fmt.Println(v3)
    }
    for fi58 := 0; fi58 < 1; fi58++ {
        for fi59 := 0; fi59 < 15; fi59++ {
            v2 = append(v2, "foo bar")
            if (float32(1.5) <= 0.1) {
                v3 = false
                if (((!v3) || (float32(1e10) > 0.1)) || v3) {
                    v60 := GenS3{F0: (!(!true)), F1: ("a" + ("hello" + "")), F2: 42}
                    var v61 int16 = 0
                    _ = v60
                    _ = v61
                } else {
                    var v62 int = 128
                    _ = v62
                }
            }
            if v3 {
                v3 = ((fi59 & -9223372036854775808) == 32767)
                v3 = false
            }
            _ = fi59
        }
        v3 = v3
        var v63 uint8 = 42
        _ = fi58
        _ = v63
    }
    v1 += 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
