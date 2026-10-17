package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint; F2 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0() int8 {
    var v7 int = -1
    for fi8 := 0; fi8 < 9; fi8++ {
        if (4294967295 > 2) {
            v7 = fi8
            switch (("" + "\n\t") + "\n\t") {
            case "a":
                for fi9 := 0; fi9 < 8; fi9++ {
                    v7 += (v7 % ((fi9) | 1))
                    _ = fi9
                }
            default:
                v7 = fi8
            }
            v10 := GenS1{F0: -759}
            _ = v10
        } else {
            if false {
                var v11 bool = (748 > 65535)
                fmt.Println(v11)
                _ = v11
            } else {
                v12 := GenS2{F0: 814, F1: ((-2.25 - 0.1) < 1.5)}
                _ = v12
            }
            var v13 uint8 = 143
            _ = v13
        }
        v7 -= 0
        _ = fi8
    }
    v14 := GenS3{F0: -1, F1: 42}
    _ = v14
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 32767
    v3 := []uint32{}
    var v4 uint32 = 2
    v5 := []string{"a", ("\n\t" + "\n\t"), "hello"}
    var v6 int8 = -128
    v15 := GenS4{F0: 2, F1: 1}
    v1 += (v1 & v1)
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS3{F0: 1, F1: 42}
        var v18 int32 = 42
        v19 := GenS5{F0: (-9223372036854775808 == -9223372036854775808), F1: 500}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS6{F0: 0}
    v4 = uint32(v2)
    v21 := GenS4{F0: 524, F1: -9223372036854775808}
    for fi22 := 0; fi22 < 14; fi22++ {
        v5 = append(v5, "foo bar")
        if true {
            for fi23 := 0; fi23 < 19; fi23++ {
                if (0.1 < ((-2.25 + 0.1) - float32(-2.25))) {
                    if true {
                        v1 = v1
                        fmt.Println(fi23, v6, v1)
                    }
                    v4 += (v4 * v4)
                    switch "\n\t" {
                    case "foo bar":
                        var v24 uint = 263
                        _ = v24
                    case "a":
                        v1 = v1
                    default:
                        if true {
                            if false {
                                switch 9223372036854775807 {
                                case 978:
                                    var v25 float32 = (1e10 + (0.1 - 3.14159))
                                    _ = v25
                                case 764:
                                    fmt.Println(fi22, v1, fi22)
                                default:
                                    v1 -= (v1 & v1)
                                }
                                v2 -= (^(v2))
                                fmt.Println(fi23, v1, fi23)
                            }
                            var v26 float64 = -2.25
                            _ = v26
                        } else {
                            fmt.Println(fi23, v6)
                            switch (fi23 % ((fi22) | 1)) {
                            case 1:
                                v6 += v6
                            case -2147483648:
                                v27 := GenS7{F0: "a", F1: 255, F2: 924}
                                _ = v27
                            case 2147483647:
                                switch "" {
                                case "hello":
                                    if (!true) {
                                        v2 = 829
                                        v4 += (v4 << ((uint(v4)) & 63))
                                    } else {
                                        var v28 int = (fi23 ^ (fi23 | 4294967295))
                                        v4 = (^(v4))
                                        _ = v28
                                    }
                                case "\n\t":
                                    v5 = append(v5, "a")
                                }
                            default:
                                if (false && ("" > "foo bar")) {
                                    switch (fi23 % (((fi23 | fi23)) | 1)) {
                                    case -2147483648:
                                        v2 = 32767
                                    }
                                    v29 := GenS8{F0: 1}
                                    if len(v5) > 0 {
                                        v5[0] = "foo bar"
                                    }
                                    _ = v29
                                } else {
                                    fmt.Println(v4, fi23, v4)
                                }
                            }
                            v3 = append(v3, (v4 * (^(v4))))
                        }
                    }
                } else {
                    v6 += (v6 >> ((429) & 63))
                    if ("" <= "") {
                        if (1e10 != 0.0) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v3 = append(v3, v4)
                        }
                        v4 = (v4 << ((42) & 63))
                        if false {
                            v4 = (v4 | 42)
                            var v30 int16 = -32768
                            v31 := GenS1{F0: int(v4)}
                            _ = v30
                            _ = v31
                        }
                    } else {
                        fmt.Println(fi22)
                        var v32 bool = (!false)
                        switch 193 {
                        case -138:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 9223372036854775807:
                            v4 = v4
                        case -1:
                            fmt.Println(v1, v2, v32)
                        }
                        _ = v32
                    }
                    if false {
                        switch "" {
                        case "":
                            if (false && true) {
                                v3 = append(v3, v4)
                            }
                        default:
                            if true {
                                fmt.Println(fi23)
                                if false {
                                    v33 := GenS5{F0: false, F1: 0}
                                    if true {
                                        var v34 uint8 = uint8(v4)
                                        if false {
                                            v35 := GenS8{F0: -488}
                                            _ = v35
                                        } else {
                                            var v36 float32 = 0.1
                                            var v37 int16 = -32768
                                            switch (-(fi23)) {
                                            case 128:
                                                if len(v3) > 0 {
                                                    v3[0] = (^(v4))
                                                }
                                            case 127:
                                                if true {
                                                    if true {
                                                        v38 := GenS9{F0: uint8(v4), F1: ("foo bar" + "")}
                                                        if false {
                                                            var v39 int32 = -1
                                                            v34 += v34
                                                            _ = v39
                                                        } else {
                                                            var v40 float32 = v36
                                                            v40 = (0.1 / (1.5 * 0.1))
                                                            var v41 float32 = (v36 / float32(((0.0 * (3.14159 * 1.5)) - (1.5 / 1.0))))
                                                            v5 = append(v5, "a")
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        var v42 int16 = v37
                                                        _ = v38
                                                        _ = v42
                                                    }
                                                    var v43 uint32 = (v4 % ((v4) | 1))
                                                    _ = v43
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        if true {
                                                            v36 += v36
                                                            v44 := GenS10{F0: 42, F1: 42}
                                                            var v45 uint32 = v4
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        var v46 int16 = (v37 / ((v37) | 1))
                                                        v36 += v36
                                                        _ = v46
                                                    case "":
                                                        var v47 uint64 = v1
                                                        _ = v47
                                                    case "\n\t":
                                                        var v48 float32 = -2.25
                                                        _ = v48
                                                    }
                                                    var v49 uint32 = 42
                                                    v50 := GenS11{F0: uint64(v2)}
                                                    _ = v49
                                                    _ = v50
                                                }
                                            case -2147483648:
                                                var v51 uint8 = (v34 ^ v34)
                                                _ = v51
                                            default:
                                                if (v36 <= ((((float32(float64(3.14159)) - ((float32(1e10) * float32(((1e10 - 0.1) / 1e10))) + (((v36 + (1.5 * 3.14159)) + v36) * v36))) - v36) * v36) * v36)) {
                                                    v52 := GenS8{F0: -32768}
                                                    _ = v52
                                                } else {
                                                    v5 = append(v5, ("hello" + (("" + "a") + "foo bar")))
                                                }
                                            }
                                            _ = v36
                                            _ = v37
                                        }
                                        fmt.Println(v34, v4, v6)
                                        v34 = v34
                                        v53 := GenS12{F0: -308, F1: 2}
                                        _ = v34
                                        _ = v53
                                    }
                                    _ = v33
                                } else {
                                    if (!false) {
                                        v54 := GenS13{F0: 42, F1: "hello"}
                                        _ = v54
                                    }
                                    var v55 string = "hello"
                                    _ = v55
                                }
                            } else {
                                fmt.Println(v2, v2, fi23)
                                var v56 int64 = 798
                                var v57 int64 = -509
                                _ = v56
                                _ = v57
                            }
                        }
                    } else {
                        if (!(!(-1 == 42))) {
                            var v58 uint = 36
                            _ = v58
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    fmt.Println(v6, fi23)
                }
                v1 = v1
                _ = fi23
            }
        } else {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
        }
        v59 := GenS14{F0: 1, F1: 755, F2: 1}
        _ = fi22
        _ = v59
    }
    v6 -= genFn0()
    fmt.Println(v2, v1)
    if false {
        switch 128 {
        case 42:
            _ = len(v5)
        default:
            v2 = 36
        }
    }
    fmt.Println(v6)
    if (1 < -9223372036854775808) {
        v60 := GenS12{F0: -2147483648, F1: 18446744073709551615}
        _ = v60
    } else {
        for fi61 := 0; fi61 < 14; fi61++ {
            v4 = v4
            _ = fi61
        }
        if (((false || false) && (850 < 2)) || false) {
            var v62 int16 = -503
            var v63 uint64 = v1
            _ = v62
            _ = v63
        }
    }
    v2 = (v2 - (v2 & uint16(v1)))
    v2 = (^(v2))
    v1 = (^(v1))
    v6 += -1
    var v64 uint8 = 235
    switch -623 {
    case 1:
        var v65 int64 = 1
        _ = v65
    case 32768:
        v4 -= v4
    default:
        for fi66 := 0; fi66 < 10; fi66++ {
            v6 = v6
            _ = fi66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
