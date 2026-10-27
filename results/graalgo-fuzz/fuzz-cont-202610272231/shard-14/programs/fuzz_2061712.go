package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 bool; F1 string; F2 uint64 }
type GenS7 struct { F0 uint; F1 int; F2 int8 }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint64) int8 {
    p1 = (p1 << ((4294967295) & 63))
    p0 = p0
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int8) int64 {
    v7 := GenS1{F0: 42, F1: 127, F2: 42}
    v8 := []int{-9223372036854775808, -9223372036854775808, 128, 390}
    switch "" {
    case "hello":
        var v9 float32 = 0.1
        _ = v9
    default:
        if len(v8) > 0 {
            v8[0] = 127
        }
    }
    var v10 int64 = 42
    if true {
        var v11 uint32 = 110
        var v12 uint8 = 255
        _ = v11
        _ = v12
    } else {
        _ = len(v8)
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    return (v10 + (v10 + v10))
}

func main() {
    var v1 string = ""
    var v2 float32 = ((3.14159 + float32(1.5)) - 3.14159)
    v3 := make(map[string]bool)
    var v4 int8 = 127
    v5 := []int{127, 9223372036854775807}
    var v6 uint64 = 4294967295
    var v13 bool = ((1 != -1) || (-568 < -1))
    v2 += -2.25
    if v13 {
        _ = len(v3)
        switch v1 {
        case "":
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    _ = len(v3)
                    delete(v3, v1)
                    v2 += v2
                    _ = fi15
                }
                _ = fi14
            }
        case "a":
            var v16 int = -9223372036854775808
            _ = v16
        default:
            var v17 float32 = v2
            _ = v17
        }
    } else {
        var v18 bool = v13
        _ = v18
    }
    _ = len(v3)
    var v19 float64 = 3.14159
    if (255 > 1) {
        v4 += 100
        if v13 {
            if (!false) {
                v4 = 32
                v13 = v13
            } else {
                if ((-9223372036854775808 != 32767) || (820 < 141)) {
                    v13 = false
                    switch -1 {
                    case -1:
                        if (18446744073709551615 != 18446744073709551615) {
                            v1 = (v1 + v1)
                        } else {
                            var v20 float32 = float32(float64((0.1 / 1.0)))
                            var v21 int32 = 1
                            _ = v20
                            _ = v21
                        }
                    case -9223372036854775808:
                        if v13 {
                            if (!true) {
                                fmt.Println(v1, v13, v4)
                            }
                            v22 := GenS1{F0: -1, F1: 127, F2: 1}
                            v23 := GenS1{F0: 1, F1: 210, F2: -9223372036854775808}
                            _ = v22
                            _ = v23
                        }
                    case 2147483647:
                        for fi24 := 0; fi24 < 18; fi24++ {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                v26 := GenS2{F0: (-(v4)), F1: -313, F2: 42}
                                _ = fi25
                                _ = v26
                            }
                            v19 = v19
                            for fi27 := 0; fi27 < 14; fi27++ {
                                switch (v1 + ("" + ((v1 + "\n\t") + ("hello" + ((("a" + "\n\t") + "\n\t") + (((("" + (v1 + (v1 + "hello"))) + "hello") + v1) + (v1 + ("\n\t" + "\n\t")))))))) {
                                case "foo bar":
                                    v3[(((v1 + "") + (("a" + ("\n\t" + v1)) + v1)) + (v1 + v1))] = (v6 > (v6 << ((18446744073709551615) & 63)))
                                case "\n\t":
                                    switch 549 {
                                    case 32768:
                                        if false {
                                            if (fi24 >= fi24) {
                                                delete(v3, "")
                                            }
                                            v19 = v19
                                            fmt.Println(v13)
                                        }
                                    case 42:
                                        if v13 {
                                            var v28 int64 = -914
                                            switch 32768 {
                                            case -1:
                                                var v29 int = (fi24 * int(v28))
                                                _ = v29
                                            default:
                                                v30 := GenS3{F0: 0, F1: (v28 & (v28 | v28)), F2: genFn1(v4)}
                                                _ = v30
                                            }
                                            _ = v28
                                        }
                                    case 9223372036854775807:
                                        v31 := GenS4{F0: 50, F1: 0, F2: fi24}
                                        _ = v31
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                case "a":
                                    if (!v13) {
                                        fmt.Println(v6, fi24)
                                        var v32 int16 = int16(v4)
                                        _ = v32
                                    } else {
                                        v33 := GenS5{F0: -1, F1: v4, F2: (^(v6))}
                                        v34 := GenS6{F0: v13, F1: v1, F2: v6}
                                        v5 = append(v5, (fi27 ^ 4294967295))
                                        _ = v33
                                        _ = v34
                                    }
                                default:
                                    if true {
                                        v35 := GenS2{F0: genFn0(39, (v6 ^ v6)), F1: 23, F2: 0}
                                        _ = v35
                                    } else {
                                        v2 += 0.0
                                        var v36 int8 = 1
                                        _ = v36
                                    }
                                }
                                if false {
                                    var v37 uint64 = (v6 ^ v6)
                                    _ = v37
                                } else {
                                    v4 += v4
                                    v38 := GenS1{F0: 42, F1: 128, F2: 1}
                                    switch "foo bar" {
                                    case "\n\t":
                                        v1 = ("a" + "hello")
                                    case "foo bar":
                                        var v39 uint8 = 109
                                        _ = v39
                                    default:
                                        var v40 string = v1
                                        _ = v40
                                    }
                                    _ = v38
                                }
                                _ = fi27
                            }
                            _ = fi24
                        }
                    }
                } else {
                    v13 = v13
                    v2 += (v2 - 0.1)
                }
            }
        }
    } else {
        var v41 float32 = (v2 / v2)
        fmt.Println(v19)
        v42 := GenS5{F0: int(v4), F1: genFn0(0, 4294967295), F2: 4294967295}
        _ = v41
        _ = v42
    }
    delete(v3, "\n\t")
    v43 := GenS7{F0: 653, F1: 32767, F2: genFn0(2147483648, (v6 | (v6 % ((18446744073709551615) | 1))))}
    v3[("\n\t" + ((((((v1 + (v1 + v1)) + v1) + v1) + (v1 + v1)) + (v1 + (v1 + v1))) + "\n\t"))] = (!(false || true))
    var v44 uint32 = 2147483648
    var v45 int = -964
    if (v44 == 898) {
        var v46 float32 = v2
        if v13 {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v47 string = "foo bar"
            var v48 int8 = (v4 % ((38) | 1))
            _ = v47
            _ = v48
        } else {
            switch v1 {
            case "":
                switch (^(v45)) {
                case -1:
                    if v13 {
                        var v49 uint16 = 568
                        if (319 < -488) {
                            v3[v1] = (32767 != int16(v45))
                            for fi50 := 0; fi50 < 4; fi50++ {
                                var v51 int = -2147483648
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    var v53 int32 = 697
                                    if v13 {
                                        var v54 uint64 = v6
                                        _ = v54
                                    }
                                    var v55 int = (-(fi52))
                                    _ = fi52
                                    _ = v53
                                    _ = v55
                                }
                                _ = fi50
                                _ = v51
                            }
                        }
                        _ = v49
                    }
                default:
                    v56 := v3[v1]
                    _ = v56
                }
            case "\n\t":
                v2 = v46
            case "foo bar":
                if false {
                    v19 = v19
                } else {
                    if ((float32(v19) * (((v46 / (1.5 / 0.1)) + v2) - float32(1e10))) < 0.0) {
                        v6 -= 517
                        v57 := GenS8{F0: -9223372036854775808}
                        _ = v57
                    }
                    if v13 {
                        var v58 int64 = 0
                        _ = v58
                    } else {
                        var v59 int8 = 127
                        _ = v59
                    }
                    var v60 int = v45
                    _ = v60
                }
            }
            var v61 float32 = (-2.25 + (v2 * v2))
            fmt.Println(v45, v13)
            _ = v61
        }
        _ = v46
    }
    v2 += (v2 / 1.5)
    for fi62 := 0; fi62 < 7; fi62++ {
        v1 = v1
        _ = fi62
    }
    fmt.Println(v6, v19, v44)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
