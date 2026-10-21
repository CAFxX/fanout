package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int8; F1 bool }
func genFn0() int32 {
    var v4 bool = true
    var v5 uint8 = 46
    v6 := GenS1{F0: 1}
    if v4 {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS2{F0: v4}
            v9 := []int{fi7, fi7, 32768, fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint64 = 1
        _ = v10
    }
    if (514 != 65535) {
        switch "a" {
        case "a":
            v5 += 42
        case "foo bar":
            if (!true) {
                var v11 uint64 = 0
                var v12 int = 0
                if (!(("a" + "\n\t") <= "hello")) {
                    fmt.Println(v11, v4)
                    v12 = v12
                }
                _ = v11
                _ = v12
            } else {
                var v13 int = -959
                v4 = true
                fmt.Println(v5, v13, v13)
                _ = v13
            }
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return -973
}

func main() {
    var v1 uint64 = 121
    var v2 string = "a"
    v3 := make(map[string]string)
    if true {
        fmt.Println(v2)
        var v14 float32 = 1e10
        _ = v14
    } else {
        v1 = v1
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS3{F0: 42, F1: 1, F2: -9223372036854775808}
        _ = fi15
        _ = v16
    }
    v17 := []int{762, 651, -1}
    if false {
        if true {
            switch ((v2 + "hello") + "a") {
            case "":
                var v18 uint32 = 2147483647
                _ = v18
            case "hello":
                var v19 int64 = -9223372036854775808
                _ = v19
            case "foo bar":
                if true {
                    v17 = append(v17, 0)
                    if (true || (-1 != -128)) {
                        fmt.Println(v1, v2)
                        v2 = v2
                        v20 := GenS4{F0: 1, F1: 0, F2: -586}
                        _ = v20
                    }
                    var v21 int32 = -2147483648
                    _ = v21
                }
            default:
                v22 := GenS3{F0: 58, F1: 65535, F2: 0}
                _ = v22
            }
            fmt.Println(v1)
            v23 := GenS5{F0: 127}
            _ = v23
        }
        v1 += 4294967295
        for fi24 := 0; fi24 < 15; fi24++ {
            delete(v3, v2)
            v25 := GenS3{F0: 0, F1: 127, F2: -9223372036854775808}
            _ = fi24
            _ = v25
        }
    } else {
        fmt.Println(v1, v2)
        v1 = v1
    }
    fmt.Println(v2)
    var v26 int32 = 726
    if false {
        if (-9223372036854775808 < 128) {
            var v27 uint32 = 0
            if false {
                _ = len(v17)
                v28 := GenS2{F0: true}
                _ = v28
            }
            v27 = (v27 - (v27 * (v27 / ((42) | 1))))
            _ = v27
        } else {
            v29 := GenS6{F0: -246, F1: 127}
            _ = v29
        }
        var v30 uint = 263
        _ = v30
    } else {
        var v31 uint64 = v1
        _ = v31
    }
    v2 = v2
    switch v2 {
    case "a":
        if false {
            for fi32 := 0; fi32 < 9; fi32++ {
                for fi33 := 0; fi33 < 1; fi33++ {
                    v34 := GenS6{F0: (v26 & (v26 ^ genFn0())), F1: 0}
                    v26 = (-(v26))
                    _ = fi33
                    _ = v34
                }
                var v35 uint = 1
                _ = fi32
                _ = v35
            }
        }
    case "\n\t":
        v36 := v3[(v2 + "foo bar")]
        _ = v36
    case "":
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 bool = true
            v26 = (v26 * 42)
            _ = fi37
            _ = v38
        }
    default:
        v39 := GenS6{F0: (v26 / ((v26) | 1)), F1: 119}
        _ = v39
    }
    v1 = v1
    for fi40 := 0; fi40 < 3; fi40++ {
        var v41 float64 = 1.5
        fmt.Println(v26)
        v1 += v1
        _ = fi40
        _ = v41
    }
    v26 = (v26 | (v26 % ((v26) | 1)))
    for fi42 := 0; fi42 < 6; fi42++ {
        for fi43 := 0; fi43 < 3; fi43++ {
            v1 -= v1
            _ = fi43
        }
        switch (fi42 >> ((680) & 63)) {
        case -9223372036854775808:
            for fi44 := 0; fi44 < 2; fi44++ {
                switch (("foo bar" + "foo bar") + v2) {
                case "a":
                    var v45 float32 = (0.1 / 1e10)
                    _ = v45
                case "hello":
                    v1 = 1
                    if true {
                        v46 := v3[((v2 + "") + (("\n\t" + (("" + ("foo bar" + v2)) + v2)) + "\n\t"))]
                        switch 0 {
                        case 1:
                            fmt.Println(v2)
                        default:
                            if (v46 <= ((v2 + "hello") + v46)) {
                                var v47 uint = 4294967295
                                v17 = append(v17, fi42)
                                _ = v47
                            } else {
                                if false {
                                    fmt.Println(fi44)
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        if len(v17) > 0 {
                                            v17[0] = (fi42 % (((fi44 - (-(fi42)))) | 1))
                                        }
                                    }
                                }
                                v48 := GenS5{F0: (fi42 * fi44)}
                                _ = v48
                            }
                        }
                        v49 := GenS7{F0: 42, F1: true}
                        var v50 bool = false
                        fmt.Println(v26)
                        _ = v46
                        _ = v49
                        _ = v50
                    } else {
                        fmt.Println(fi44, v1)
                        _ = len(v17)
                        v51 := GenS2{F0: (771 < 0)}
                        _ = v51
                    }
                case "foo bar":
                    var v52 uint32 = 2147483648
                    _ = v52
                default:
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
                var v53 float64 = 1e10
                fmt.Println(v2, v53, v2)
                _ = fi44
                _ = v53
            }
        case 2147483647:
            if false {
                if false {
                    var v54 int16 = 42
                    for fi55 := 0; fi55 < 15; fi55++ {
                        v54 = v54
                        switch (fi42 / (((fi42 + fi55)) | 1)) {
                        case -832:
                            switch (-(fi42)) {
                            case 32768:
                                v56 := GenS7{F0: -1, F1: (!(!((1e10 - (1e10 - float32(3.14159))) != (0.0 - (float32(0.1) + float32(-2.25))))))}
                                _ = v56
                            default:
                                if (1e10 <= float32(1.5)) {
                                    v2 = ("\n\t" + v2)
                                    var v57 int8 = int8(v54)
                                    _ = v57
                                } else {
                                    v26 += (v26 * (v26 | (v26 << ((1) & 63))))
                                    v2 = (v2 + (v2 + v2))
                                }
                            }
                        case 4294967295:
                            var v58 float64 = (1.5 - float64(3.14159))
                            _ = v58
                        case 899:
                            var v59 bool = true
                            _ = v59
                        }
                        _ = fi55
                    }
                    _ = v54
                } else {
                    fmt.Println(v2, fi42)
                    v26 += (v26 << ((65535) & 63))
                    v2 = "foo bar"
                }
                var v60 uint32 = 975
                _ = v60
            }
        }
        _ = fi42
    }
    var v61 uint64 = v1
    if false {
        v26 += (v26 * 0)
        fmt.Println(v2, v1)
        if (float32(-2.25) <= 1e10) {
            switch 2147483647 {
            case 42:
                v62 := GenS1{F0: v26}
                _ = v62
            case 32768:
                var v63 int16 = 1
                _ = v63
            default:
                var v64 int16 = 0
                _ = v64
            }
        }
    } else {
        v61 -= (^(v61))
        if false {
            switch -9223372036854775808 {
            case 128:
                var v65 int32 = v26
                _ = v65
            case 42:
                v61 -= 42
            default:
                var v66 uint16 = 2
                _ = v66
            }
            v26 = genFn0()
        }
        var v67 float64 = ((0.1 - 1e10) - 1.5)
        _ = v67
    }
    v26 = -189
    if true {
        if len(v17) > 0 {
            v17[0] = -605
        }
        var v68 int16 = 1
        _ = v68
    } else {
        var v69 float64 = -2.25
        var v70 int32 = -2147483648
        v26 = 0
        _ = v69
        _ = v70
    }
    if false {
        v2 = v2
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = len(v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v61
}
