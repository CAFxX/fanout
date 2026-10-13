package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint; F1 uint16; F2 int16 }
type GenS7 struct { F0 int32; F1 int64; F2 int8 }
type GenS8 struct { F0 int16; F1 uint; F2 string }
type GenS9 struct { F0 uint }

func main() {
    var v1 int32 = 1
    var v2 bool = (762 <= 995)
    v3 := []string{("\n\t" + "a"), "", "", ""}
    switch "a" {
    case "a":
        if (126 != 203) {
            if (!false) {
                for fi4 := 0; fi4 < 19; fi4++ {
                    if (((int8(fi4) < -128) || (true && (32767 != -1))) && (false || false)) {
                        for fi5 := 0; fi5 < 3; fi5++ {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            } else {
                                fmt.Println(v1, fi5)
                                v6 := GenS1{F0: 255, F1: (!false)}
                                _ = v6
                            }
                            fmt.Println(v2, fi5, fi5)
                            _ = fi5
                        }
                        fmt.Println(fi4)
                    }
                    if false {
                        if (!(786 == fi4)) {
                            for fi7 := 0; fi7 < 12; fi7++ {
                                var v8 uint = 0
                                if true {
                                    v9 := GenS1{F0: 42, F1: (v8 <= v8)}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v9
                                }
                                var v10 int8 = 81
                                _ = fi7
                                _ = v10
                                _ = v8
                            }
                            var v11 int = (fi4 * fi4)
                            _ = v11
                        }
                        fmt.Println(v2)
                    } else {
                        var v12 float64 = (1e10 + 3.14159)
                        _ = v12
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi4
                }
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v2 = v2
                        v1 += v1
                        _ = fi14
                    }
                    _ = fi13
                }
            }
            v2 = v2
        }
    }
    if (-1 != 719) {
        if (v2 && false) {
            if true {
                if (!(-278 != -53)) {
                    v1 += v1
                    v15 := GenS1{F0: 1, F1: ((2 != 42) && v2)}
                    var v16 float64 = 1e10
                    _ = v15
                    _ = v16
                } else {
                    v2 = false
                    var v17 uint = 1
                    switch "\n\t" {
                    case "\n\t":
                        if v2 {
                            switch "hello" {
                            case "a":
                                v18 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: true}
                                _ = v18
                            case "hello":
                                if false {
                                    v19 := GenS3{F0: 42, F1: v1}
                                    v20 := GenS2{F0: 0, F1: -9223372036854775808, F2: (1.5 < (-2.25 + ((0.1 + 0.0) + 0.1)))}
                                    _ = v19
                                    _ = v20
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "hello":
                        var v21 string = ""
                        _ = v21
                    case "a":
                        v22 := GenS3{F0: 631, F1: v1}
                        _ = v22
                    }
                    _ = v17
                }
            } else {
                switch 128 {
                case -611:
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v3 = append(v3, "hello")
                        _ = fi23
                    }
                case -9223372036854775808:
                    fmt.Println(v2)
                case 42:
                    v1 -= (^(v1))
                default:
                    if (!false) {
                        fmt.Println(v2)
                        if (941 < 775) {
                            var v24 uint64 = 4294967295
                            _ = v24
                        }
                    }
                }
                var v25 uint8 = 2
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v25
            }
        } else {
            v2 = (false || v2)
            fmt.Println(v1, v1)
            switch "\n\t" {
            case "hello":
                v26 := GenS2{F0: -689, F1: -9223372036854775808, F2: false}
                _ = v26
            default:
                for fi27 := 0; fi27 < 1; fi27++ {
                    if (1 >= 572) {
                        v28 := GenS3{F0: 1, F1: v1}
                        if (v2 && v2) {
                            var v29 uint64 = 4294967295
                            _ = v29
                        } else {
                            v1 = (-(v1))
                            var v30 float64 = 0.1
                            _ = v30
                        }
                        v31 := GenS2{F0: 9223372036854775807, F1: 42, F2: v2}
                        _ = v28
                        _ = v31
                    }
                    if v2 {
                        if (3.14159 < 3.14159) {
                            var v32 float64 = (1e10 / 1.0)
                            v32 += v32
                            for fi33 := 0; fi33 < 16; fi33++ {
                                v34 := GenS3{F0: 2, F1: 0}
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        } else {
                            v2 = (v2 && (v1 == (v1 >> ((546) & 63))))
                        }
                        var v35 uint8 = 0
                        v35 = (v35 + (v35 | (v35 * 33)))
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 8; fi36++ {
                            var v37 int32 = (v1 | -1)
                            _ = fi36
                            _ = v37
                        }
                    }
                    for fi38 := 0; fi38 < 6; fi38++ {
                        switch "" {
                        case "a":
                            fmt.Println(v2)
                        default:
                            var v39 uint32 = 2
                            v1 += (-(v1))
                            v2 = (2 != 255)
                            fmt.Println(fi27, fi27)
                            _ = v39
                        }
                        _ = len(v3)
                        fmt.Println(fi27, v1, v1)
                        _ = fi38
                    }
                    _ = fi27
                }
            }
        }
    } else {
        v40 := GenS4{F0: 255}
        fmt.Println(v2, v2)
        _ = v40
    }
    v2 = (("" + "foo bar") <= "")
    v1 = v1
    var v41 bool = v2
    var v42 int16 = -32768
    switch "hello" {
    case "foo bar":
        if true {
            if v2 {
                var v43 int16 = (v42 * v42)
                v44 := GenS5{F0: -750, F1: 65535}
                _ = v43
                _ = v44
            }
            fmt.Println(v41, v41, v42)
        } else {
            v2 = (!(!true))
            var v45 int64 = -359
            if v41 {
                v46 := GenS6{F0: 255, F1: 32767, F2: v42}
                fmt.Println(v45)
                _ = v46
            } else {
                for fi47 := 0; fi47 < 18; fi47++ {
                    v1 = (v1 - v1)
                    var v48 float64 = -2.25
                    _ = fi47
                    _ = v48
                }
                var v49 bool = v2
                switch 310 {
                case 9223372036854775807:
                    var v50 uint64 = 2
                    _ = v50
                default:
                    var v51 uint32 = 42
                    _ = v51
                }
                _ = v49
            }
            _ = v45
        }
    case "a":
        var v52 uint = 1
        _ = v52
    case "\n\t":
        v53 := GenS1{F0: 2, F1: (0.0 <= 0.1)}
        _ = v53
    default:
        var v54 uint32 = 488
        _ = v54
    }
    if (3.14159 != 0.1) {
        var v55 uint64 = 819
        v56 := GenS7{F0: v1, F1: 122, F2: -1}
        v57 := GenS8{F0: v42, F1: 18446744073709551615, F2: "a"}
        _ = v55
        _ = v56
        _ = v57
    } else {
        if (0 < 1) {
            fmt.Println(v42, v1, v42)
        } else {
            var v58 uint64 = 4294967295
            switch 32767 {
            case -253:
                var v59 uint8 = 245
                _ = v59
            default:
                if (!v41) {
                    var v60 int64 = -1
                    var v61 int8 = 37
                    if (v2 || v2) {
                        var v62 uint8 = 10
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v62
                    }
                    _ = v60
                    _ = v61
                }
            }
            _ = v58
        }
        v2 = v2
        v63 := GenS9{F0: 242}
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
}
