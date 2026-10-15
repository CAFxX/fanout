package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int64; F1 int64; F2 int32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int32; F1 bool; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 uint64; F2 int8 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint8) int32 {
    p0 = 1
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 12; fi5++ {
            p0 = (p0 ^ p0)
            v6 := GenS1{F0: 2147483647, F1: -68}
            var v7 int16 = -32768
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 2147483647, F1: 134}
        if false {
            var v9 uint16 = 0
            v10 := []int{(-(fi4)), (-(fi4)), (fi4 / ((fi4) | 1))}
            _ = len(v10)
            _ = v10
            _ = v9
        } else {
            p0 += p0
            var v11 uint8 = p0
            v12 := GenS2{F0: 0}
            _ = v11
            _ = v12
        }
        _ = fi4
        _ = v8
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 bool = false
    var v2 bool = (!v1)
    var v3 uint8 = 0
    var v13 int32 = 445
    var v14 uint = 322
    for fi15 := 0; fi15 < 5; fi15++ {
        switch "foo bar" {
        case "hello":
            if v1 {
                v16 := GenS1{F0: 1, F1: 32767}
                v17 := GenS3{F0: 915, F1: -818}
                _ = v16
                _ = v17
            }
        case "":
            var v18 uint8 = (v3 >> (((v14 >> ((18446744073709551615) & 63))) & 63))
            _ = v18
        case "\n\t":
            if (((v13 * (v13 * (v13 + 866))) > v13) && v1) {
                if ((((v3 > v3) || v1) && v1) || false) {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if (2 <= 2147483647) {
                            v1 = ((^(v14)) == (v14 & 2))
                            v13 += 838
                            if (186 > 1) {
                                v20 := GenS3{F0: 42, F1: 14}
                                fmt.Println(v2)
                                v21 := []int{162, fi15, 9223372036854775807, fi19}
                                _ = v20
                                _ = v21
                            }
                        }
                        _ = fi19
                    }
                }
                if false {
                    fmt.Println(fi15, v2)
                    fmt.Println(v1, v14)
                    v22 := []int{(fi15 + (fi15 + (^(fi15))))}
                    _ = v22
                } else {
                    fmt.Println(v2)
                    v23 := GenS4{F0: v14, F1: true}
                    var v24 uint16 = uint16(v3)
                    _ = v23
                    _ = v24
                }
                if v2 {
                    if false {
                        var v25 int16 = 351
                        switch "a" {
                        case "":
                            v1 = (1 >= 65535)
                        case "foo bar":
                            var v26 int16 = v25
                            _ = v26
                        default:
                            var v27 float64 = 1.5
                            _ = v27
                        }
                        _ = v25
                    }
                }
            } else {
                var v28 string = ""
                var v29 int8 = 42
                v30 := GenS5{F0: (v28 + ""), F1: -1}
                _ = v28
                _ = v29
                _ = v30
            }
        }
        var v31 bool = (-2147483648 >= (fi15 + (fi15 & 411)))
        _ = fi15
        _ = v31
    }
    v32 := []int{567, 9223372036854775807, -411}
    v13 -= v13
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        if (true && ((v2 && true) || (1 > 255))) {
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            v2 = ((v13 >> ((379) & 63)) != (-(v13)))
            v34 := GenS6{F0: 1, F1: 42}
            _ = v34
        } else {
            v32 = append(v32, (fi33 + 128))
            v13 += (v13 * v13)
            if v2 {
                for fi35 := 0; fi35 < 4; fi35++ {
                    fmt.Println(fi33)
                    v36 := GenS7{F0: -1, F1: -9223372036854775808, F2: v13}
                    v1 = true
                    if v2 {
                        fmt.Println(fi35, v1, v14)
                        v32 = append(v32, int(v3))
                    } else {
                        if (!v2) {
                            var v37 string = "foo bar"
                            switch v37 {
                            case "a":
                                if v1 {
                                    fmt.Println(v3, fi33, v13)
                                    v13 = (^(v13))
                                    if ((v37 + (v37 + v37)) >= v37) {
                                        if len(v32) > 0 {
                                            v32[0] = (fi35 >> ((255) & 63))
                                        }
                                        switch fi35 {
                                        case 32767:
                                            if v2 {
                                                if len(v32) > 0 {
                                                    v32[0] = (fi33 % ((fi33) | 1))
                                                }
                                                v14 = v14
                                                var v38 int = (fi33 | (^(fi35)))
                                                _ = v38
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    v1 = v2
                                                default:
                                                    fmt.Println(v37, v37)
                                                }
                                                fmt.Println(fi35, v37)
                                                if (!true) {
                                                    v39 := GenS4{F0: (^(v14)), F1: (((false && (("foo bar" + v37) <= ("" + (v37 + "hello")))) && (1 < 2147483648)) && (((-2.25 - 1.5) != 1e10) || v1))}
                                                    if (-842 >= (v13 * v13)) {
                                                        v32 = append(v32, 32767)
                                                        v1 = (!v1)
                                                        v1 = ((1.5 != (3.14159 - 3.14159)) && v1)
                                                    } else {
                                                        switch (v37 + v37) {
                                                        case "":
                                                            v1 = (42 <= -1)
                                                        case "\n\t":
                                                            v1 = (((false && false) && true) || (!(!v1)))
                                                            if len(v32) > 0 {
                                                                v32[0] = fi33
                                                            }
                                                            v1 = ((("\n\t" + v37) + (((v37 + v37) + "") + "hello")) <= "hello")
                                                        }
                                                        v40 := GenS8{F0: (fi35 ^ (-(fi33))), F1: ((1e10 - ((1e10 / 1.0) * 0.0)) != 1e10)}
                                                        var v41 uint32 = 693
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                } else {
                                                    fmt.Println(fi33, v14, v1)
                                                    var v42 string = v37
                                                    _ = v42
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v43 := GenS7{F0: -1, F1: -670, F2: (v13 << (((v14 & 0)) & 63))}
                                    _ = v43
                                }
                            case "\n\t":
                                v44 := GenS5{F0: "a", F1: -773}
                                _ = v44
                            case "":
                                var v45 int16 = 42
                                _ = v45
                            }
                            _ = v37
                        }
                    }
                    _ = len(v32)
                    _ = fi35
                    _ = v36
                }
            } else {
                if len(v32) > 0 {
                    v32[0] = fi33
                }
                switch ((("foo bar" + ("\n\t" + "")) + "hello") + "a") {
                case "":
                    var v46 int32 = (v13 / ((161) | 1))
                    _ = v46
                }
            }
        }
        var v47 bool = v2
        v47 = (0 == 1000)
        _ = fi33
        _ = v47
    }
    v48 := GenS9{F0: 42, F1: v1, F2: 769}
    var v49 int16 = -1
    v50 := GenS10{F0: 42}
    fmt.Println(v1)
    v51 := GenS11{F0: false, F1: 18446744073709551615, F2: -1}
    if ("foo bar" >= "") {
        if (42 <= v13) {
            if (!(!true)) {
                var v52 int16 = (v49 & v49)
                var v53 int32 = -902
                _ = v52
                _ = v53
            }
            if v2 {
                v32 = append(v32, 890)
                v3 += 42
                if (519 >= 1) {
                    if len(v32) > 0 {
                        v32[0] = 127
                    }
                    _ = len(v32)
                    v14 -= (v14 ^ (v14 ^ (v14 | v14)))
                } else {
                    if (!true) {
                        var v54 string = ("\n\t" + "hello")
                        switch v54 {
                        case "":
                            for fi55 := 0; fi55 < 17; fi55++ {
                                v2 = ((!(288 > 0)) || v1)
                                _ = fi55
                            }
                        }
                        _ = v54
                    } else {
                        var v56 uint64 = 2
                        for fi57 := 0; fi57 < 8; fi57++ {
                            if (42 < 1) {
                                v14 += v14
                                if (!v2) {
                                    for fi58 := 0; fi58 < 19; fi58++ {
                                        var v59 float64 = 0.1
                                        v1 = ((1.5 >= -2.25) && (32767 > 32767))
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                        v32 = append(v32, 1)
                                        v60 := GenS10{F0: 2147483648}
                                        _ = fi58
                                        _ = v59
                                        _ = v60
                                    }
                                    v32 = append(v32, 32768)
                                } else {
                                    fmt.Println(v49, v13)
                                    var v61 uint8 = v3
                                    v14 -= v14
                                    _ = v61
                                }
                                if len(v32) > 0 {
                                    v32[0] = fi57
                                }
                            }
                            _ = fi57
                        }
                        _ = v56
                    }
                    var v62 int8 = 76
                    _ = v62
                }
            } else {
                fmt.Println(v13, v3)
                var v63 int8 = 1
                _ = v63
            }
        } else {
            v13 -= genFn0(v3)
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            v13 += (v13 * (v13 ^ v13))
        }
    } else {
        v64 := GenS3{F0: 83, F1: v49}
        _ = v64
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        var v66 int16 = int16(v3)
        switch 49 {
        case 4294967295:
            switch 128 {
            case 128:
                if ((fi65 % ((fi65) | 1)) == (fi65 >> ((v14) & 63))) {
                    switch "foo bar" {
                    case "foo bar":
                        var v67 uint = 255
                        _ = v67
                    case "\n\t":
                        var v68 uint32 = 1
                        _ = v68
                    }
                }
            case -46:
                var v69 uint32 = 455
                _ = v69
            case 32767:
                var v70 uint16 = 918
                _ = v70
            default:
                v71 := GenS12{F0: 156}
                _ = v71
            }
        case 57:
            var v72 uint16 = 1
            _ = v72
        }
        fmt.Println(v1, v66, v1)
        _ = fi65
        _ = v66
    }
    v32 = append(v32, 32768)
    v14 = v14
    v32 = append(v32, 335)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v48
    _ = v49
    _ = v50
    _ = v51
}
