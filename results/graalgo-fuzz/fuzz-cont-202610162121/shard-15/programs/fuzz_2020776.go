package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 4294967295
    var v3 uint32 = uint32(v2)
    v4 := []uint32{(v3 >> ((65535) & 63)), (v3 & v3)}
    var v5 int16 = 0
    var v6 uint64 = 1
    v6 = v6
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 204
            v10 := GenS1{F0: (v2 + 799)}
            if (3.14159 <= 1.5) {
                if false {
                    v9 -= v9
                    if ((!true) || (!(!true))) {
                        if (int32(v3) >= 0) {
                            if false {
                                if ((-2.25 - (1.5 - 1.5)) < 3.14159) {
                                    fmt.Println(v9)
                                    var v11 float32 = (1e10 - ((float32(0.0) / 1e10) / 3.14159))
                                    var v12 uint64 = v2
                                    _ = v11
                                    _ = v12
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v4 = append(v4, (v3 >> ((42) & 63)))
                                }
                                if false {
                                    var v13 string = ""
                                    if ((v5 % ((-540) | 1)) > -1) {
                                        v14 := GenS2{F0: v13}
                                        var v15 int8 = 121
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v14
                                        _ = v15
                                    }
                                    var v16 float32 = float32(-2.25)
                                    _ = v13
                                    _ = v16
                                } else {
                                    v6 -= v6
                                    if (-1 != fi8) {
                                        v6 += (v2 + 4294967295)
                                        v1 -= -1
                                        _ = len(v4)
                                    } else {
                                        if ("\n\t" > ("hello" + "foo bar")) {
                                            var v17 int16 = v5
                                            v3 += (v3 ^ v3)
                                            fmt.Println(v17, v9, v3)
                                            _ = v17
                                        } else {
                                            v9 = 103
                                            v18 := GenS2{F0: (("hello" + "a") + "\n\t")}
                                            v3 += 2147483648
                                            _ = v18
                                        }
                                        var v19 int32 = -2147483648
                                        var v20 int16 = (v5 << ((6) & 63))
                                        _ = v19
                                        _ = v20
                                    }
                                    var v21 int32 = 1
                                    v1 = 1
                                    var v22 uint8 = 195
                                    _ = v21
                                    _ = v22
                                }
                                var v23 uint64 = (v2 % ((v2) | 1))
                                var v24 uint16 = 65535
                                _ = v23
                                _ = v24
                            } else {
                                v25 := GenS2{F0: "foo bar"}
                                _ = v25
                            }
                            if true {
                                if false {
                                    var v26 string = "hello"
                                    if false {
                                        var v27 uint32 = 2
                                        if len(v4) > 0 {
                                            v4[0] = v27
                                        }
                                        _ = v27
                                    }
                                    _ = v26
                                } else {
                                    if true {
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = (^(v3))
                                            }
                                        } else {
                                            v2 -= v2
                                        }
                                    } else {
                                        var v28 int64 = (v1 + v1)
                                        var v29 float64 = 1e10
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 float64 = 1.5
                                    _ = v30
                                }
                            }
                            v9 = v9
                        }
                        v31 := GenS3{F0: v1}
                        var v32 string = "hello"
                        _ = v31
                        _ = v32
                    }
                    if (v1 > v1) {
                        switch "foo bar" {
                        case "foo bar":
                            var v33 int = (fi8 >> ((2) & 63))
                            _ = v33
                        case "\n\t":
                            var v34 uint = 926
                            _ = v34
                        }
                    } else {
                        var v35 int = fi7
                        var v36 int32 = 42
                        if false {
                            fmt.Println(fi8)
                            fmt.Println(v5, fi8)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            if false {
                                fmt.Println(v9, v5)
                                var v37 bool = false
                                _ = v37
                            } else {
                                var v38 uint = 607
                                _ = v38
                            }
                            if true {
                                v4 = append(v4, (v3 | 0))
                                v4 = append(v4, uint32(v6))
                                var v39 bool = false
                                _ = v39
                            }
                            var v40 uint64 = v2
                            _ = v40
                        }
                        _ = v35
                        _ = v36
                    }
                } else {
                    v3 += (v3 + (v3 >> ((4294967295) & 63)))
                    switch fi7 {
                    case 128:
                        var v41 int = fi8
                        if false {
                            if ((1 == v5) || (-2147483648 <= 2147483647)) {
                                v41 -= 0
                                v4 = append(v4, 0)
                            }
                        } else {
                            if (true || false) {
                                v9 += v9
                            } else {
                                var v42 int8 = -128
                                v43 := GenS3{F0: 629}
                                var v44 string = ("hello" + "foo bar")
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            v9 -= 255
                        }
                        v5 += v5
                        _ = v41
                    case -785:
                        if false {
                            v4 = append(v4, v3)
                            v45 := GenS2{F0: ""}
                            v9 = (v9 * 211)
                            _ = v45
                        }
                    default:
                        var v46 uint = 65535
                        _ = v46
                    }
                    var v47 int = 128
                    v47 -= 42
                    _ = v47
                }
                v3 = 2147483648
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v48 := GenS4{F0: (fi7 - 128)}
        if len(v4) > 0 {
            v4[0] = (^(v3))
        }
        _ = fi7
        _ = v48
    }
    v4 = append(v4, (v3 ^ (v3 & v3)))
    v5 += v5
    var v49 uint = 18446744073709551615
    fmt.Println(v1)
    v50 := GenS5{F0: 255, F1: (^(v5)), F2: (^(v2))}
    fmt.Println(v2, v5)
    if true {
        var v51 int16 = (v5 / ((-991) | 1))
        _ = v51
    }
    if false {
        var v52 bool = false
        if (!((^(v5)) > (v5 / (((v5 % (((v5 ^ (v5 << ((925) & 63)))) | 1))) | 1)))) {
            if ((v1 & (v1 % ((v1) | 1))) != v1) {
                var v53 int8 = 127
                v54 := GenS5{F0: 127, F1: v5, F2: (v2 | (^(v2)))}
                switch -7 {
                case 771:
                    switch "\n\t" {
                    case "\n\t":
                        v5 = (v5 / ((-1) | 1))
                    case "hello":
                        fmt.Println(v1, v53)
                    default:
                        v55 := GenS5{F0: 12, F1: v5, F2: v6}
                        _ = v55
                    }
                case 917:
                    for fi56 := 0; fi56 < 2; fi56++ {
                        fmt.Println(v2, v6)
                        _ = fi56
                    }
                }
                _ = v53
                _ = v54
            } else {
                v57 := GenS6{F0: v5, F1: -455}
                if len(v4) > 0 {
                    v4[0] = (v3 << ((uint(v6)) & 63))
                }
                v6 = (v2 / (((v6 ^ (v6 >> (((v49 | v49)) & 63)))) | 1))
                _ = v57
            }
        }
        for fi58 := 0; fi58 < 12; fi58++ {
            for fi59 := 0; fi59 < 10; fi59++ {
                v1 -= (v1 ^ (v1 % ((v1) | 1)))
                v6 = v2
                _ = fi59
            }
            var v60 uint8 = 2
            fmt.Println(v3, v60, v49)
            _ = fi58
            _ = v60
        }
        _ = v52
    } else {
        v61 := GenS7{F0: (-(v1)), F1: true}
        switch "foo bar" {
        case "hello":
            if (uint16(v49) == 1) {
                if false {
                    for fi62 := 0; fi62 < 16; fi62++ {
                        fmt.Println(v3, v3, v2)
                        _ = fi62
                    }
                    v49 += 18446744073709551615
                } else {
                    v1 = -9223372036854775808
                    v63 := GenS4{F0: 0}
                    _ = v63
                }
            } else {
                var v64 float64 = 1.5
                if len(v4) > 0 {
                    v4[0] = (v3 % (((v3 ^ (v3 / ((2147483647) | 1)))) | 1))
                }
                v4 = append(v4, (v3 >> (((^(v49))) & 63)))
                _ = v64
            }
        case "":
            var v65 int = 128
            _ = v65
        default:
            v6 = v6
        }
        _ = v61
    }
    v66 := GenS8{F0: 42}
    for fi67 := 0; fi67 < 1; fi67++ {
        fmt.Println(v5, v49)
        _ = fi67
    }
    v5 -= -672
    var v68 bool = (1 > 1)
    if (v1 < v1) {
        switch 127 {
        case 128:
            v69 := GenS5{F0: 127, F1: (-(v5)), F2: (v6 * (v6 + (v2 % (((v2 << ((v49) & 63))) | 1))))}
            _ = v69
        case 0:
            if len(v4) > 0 {
                v4[0] = (v3 * uint32(v49))
            }
        }
    }
    if v68 {
        v70 := GenS9{F0: (-(v1)), F1: (v6 - v2)}
        if v68 {
            var v71 uint8 = 126
            _ = v71
        }
        if (!(v68 || true)) {
            for fi72 := 0; fi72 < 16; fi72++ {
                v73 := GenS9{F0: v1, F1: 4294967295}
                var v74 int64 = v1
                v49 = 65535
                _ = fi72
                _ = v73
                _ = v74
            }
        }
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v66
    _ = v68
}
