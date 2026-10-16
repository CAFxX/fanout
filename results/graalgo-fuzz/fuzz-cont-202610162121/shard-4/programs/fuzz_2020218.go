package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
func genFn0(p0 int8, p1 uint16, p2 uint16) int16 {
    fmt.Println(p0)
    if false {
        switch 42 {
        case 32768:
            var v7 uint8 = 0
            _ = v7
        case 4294967295:
            var v8 uint8 = 42
            _ = v8
        }
        switch -9223372036854775808 {
        case 2147483647:
            var v9 int64 = -9223372036854775808
            _ = v9
        }
    }
    if (1 > 42) {
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: 42}
                var v12 uint32 = 2147483647
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 12; fi14++ {
                    p2 = (^(p1))
                    fmt.Println(fi14, fi13)
                    if (57 == 142) {
                        var v15 int32 = 1
                        var v16 uint = 255
                        _ = v15
                        _ = v16
                    } else {
                        p1 = p1
                        p1 += 42
                        if false {
                            v17 := []int{fi14, (-(fi14))}
                            _ = v17
                        }
                    }
                    var v18 uint16 = 2
                    var v19 float32 = 1e10
                    p2 += v18
                    var v20 float32 = v19
                    fmt.Println(p0, p2)
                    _ = fi14
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 int32 = 0
                _ = fi13
                _ = v21
            }
        } else {
            p2 += (p1 * (p2 << ((uint(p1)) & 63)))
        }
    } else {
        p2 = p1
        v22 := GenS1{F0: 2}
        var v23 bool = (false || (-439 != 0))
        _ = v22
        _ = v23
    }
    var v24 uint32 = 42
    v25 := []int{-1, 32768, 1, -987}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 153
}
func genFn1() uint16 {
    var v26 uint64 = 1
    if (false || false) {
        var v27 float64 = 0.0
        _ = v27
    } else {
        var v28 int = 952
        _ = v28
    }
    _ = v26
    return 255
}

func main() {
    v1 := []uint32{934}
    var v2 int8 = 42
    v3 := []uint32{2147483647, 42}
    v4 := []uint32{42, 4294967295}
    v5 := []string{}
    var v6 int8 = (^(v2))
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            switch (fi29 | fi29) {
            case -274:
                v2 = (v6 ^ (v2 << ((0) & 63)))
            default:
                switch fi30 {
                case -2147483648:
                    if (-942 < 1) {
                        if ((42 >= 801) || (65535 != 2147483647)) {
                            v2 += -98
                            fmt.Println(v6)
                            v6 = (^(v2))
                            fmt.Println(fi30, fi29, v2)
                            v2 += v2
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                var v31 int32 = -2147483648
                                _ = v31
                            default:
                                v32 := GenS2{F0: 0}
                                _ = v32
                            }
                        }
                        switch (fi29 ^ -2147483648) {
                        case 32767:
                            if true {
                                v2 -= v6
                                v33 := GenS2{F0: -619}
                                _ = v33
                            } else {
                                if true {
                                    if ((v2 / ((-128) | 1)) > v6) {
                                        v34 := GenS3{F0: 0, F1: 0}
                                        switch (("" + "a") + "a") {
                                        case "":
                                            v6 = (v2 >> ((720) & 63))
                                        case "foo bar":
                                            v6 -= 1
                                        case "a":
                                            if true {
                                                switch fi30 {
                                                case 128:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                var v35 int64 = 0
                                                if true {
                                                    var v36 int = fi30
                                                    v37 := GenS2{F0: (v35 % (((v35 + 9223372036854775807)) | 1))}
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    fmt.Println(fi29)
                                                    switch ("hello" + ("hello" + "")) {
                                                    case "foo bar":
                                                        var v38 uint32 = 42
                                                        _ = v38
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        v35 += v35
                                                        fmt.Println(v6)
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v35
                                            }
                                        }
                                        if true {
                                            fmt.Println(v2)
                                            v2 = -4
                                            switch "foo bar" {
                                            case "":
                                                if false {
                                                    var v39 int = fi29
                                                    _ = v39
                                                }
                                            case "a":
                                                var v40 string = "hello"
                                                _ = v40
                                            case "\n\t":
                                                switch "" {
                                                case "hello":
                                                    v41 := GenS1{F0: 42}
                                                    var v42 bool = false
                                                    _ = v41
                                                    _ = v42
                                                case "a":
                                                    v2 -= (v2 % (((v2 ^ (-(v6)))) | 1))
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = 0
                                                    }
                                                }
                                            }
                                        } else {
                                            v43 := GenS4{F0: 108}
                                            v44 := GenS5{F0: -1, F1: 42, F2: 2147483647}
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 int64 = -9223372036854775808
                                        v46 := GenS1{F0: 307}
                                        _ = v34
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                var v47 float32 = 1.5
                                v6 += v2
                                fmt.Println(fi29, v6, v2)
                                _ = v47
                            }
                        case 9223372036854775807:
                            v6 -= (-(v6))
                        default:
                            if (-2147483648 < 1) {
                                var v48 int16 = genFn0(42, 2, 280)
                                var v49 uint64 = 618
                                _ = v48
                                _ = v49
                            } else {
                                if (0 < genFn1()) {
                                    var v50 float32 = 3.14159
                                    switch "" {
                                    case "\n\t":
                                        var v51 int64 = 9223372036854775807
                                        _ = v51
                                    }
                                    _ = v50
                                } else {
                                    var v52 float64 = 0.1
                                    v53 := GenS5{F0: -32768, F1: 0, F2: 65535}
                                    _ = v52
                                    _ = v53
                                }
                                var v54 float64 = 1e10
                                var v55 uint = 542
                                if (-2147483648 != 0) {
                                    v6 -= -123
                                    if true {
                                        v56 := GenS5{F0: 931, F1: 0, F2: 865}
                                        _ = v56
                                    }
                                    v55 = v55
                                }
                                _ = v54
                                _ = v55
                            }
                        }
                    } else {
                        fmt.Println(v6)
                        _ = len(v5)
                        if (65535 == 42) {
                            v3 = append(v3, 1)
                        }
                    }
                }
                var v57 uint8 = 133
                _ = v57
            }
            _ = fi30
        }
        v6 = 0
        _ = fi29
    }
    if false {
        v3 = append(v3, 42)
        _ = len(v3)
        v6 -= (v2 >> ((0) & 63))
    } else {
        v58 := GenS6{F0: -42, F1: 42}
        v59 := GenS4{F0: 42}
        _ = v58
        _ = v59
    }
    v2 -= (v2 / ((v6) | 1))
    v2 = v2
    v60 := GenS3{F0: 1, F1: 42}
    v2 += v6
    for fi61 := 0; fi61 < 13; fi61++ {
        for fi62 := 0; fi62 < 3; fi62++ {
            if (1 == 42) {
                _ = len(v3)
                switch "" {
                case "a":
                    fmt.Println(v2, fi61)
                default:
                    if false {
                        var v63 uint32 = 42
                        v64 := GenS3{F0: 42, F1: 198}
                        var v65 int16 = 1000
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                }
            }
            if (!true) {
                v6 += v6
                if ((0 != v2) || true) {
                    v6 -= int8(fi61)
                    v3 = append(v3, 2)
                    switch (fi61 ^ fi61) {
                    case 9223372036854775807:
                        var v66 uint = 0
                        var v67 string = (("hello" + ("" + "")) + "\n\t")
                        switch fi62 {
                        case 9223372036854775807:
                            fmt.Println(v67, fi61)
                        default:
                            if true {
                                if false {
                                    var v68 uint32 = 0
                                    _ = v68
                                } else {
                                    v67 = "a"
                                    if len(v5) > 0 {
                                        v5[0] = v67
                                    }
                                }
                                if true {
                                    var v69 string = "foo bar"
                                    v2 += (v2 * (v6 / (((^(v6))) | 1)))
                                    _ = v69
                                }
                            } else {
                                v70 := GenS5{F0: -30, F1: 0, F2: uint32(v66)}
                                if (!false) {
                                    v71 := GenS7{F0: 4294967295}
                                    _ = v71
                                }
                                if (!(false && true)) {
                                    v72 := GenS1{F0: 42}
                                    _ = v72
                                } else {
                                    if (-9223372036854775808 < 1) {
                                        v67 = v67
                                        switch v67 {
                                        case "a":
                                            fmt.Println(fi61)
                                        default:
                                            v6 = v2
                                        }
                                        fmt.Println(v6, v6)
                                    }
                                }
                                _ = v70
                            }
                        }
                        v66 = v66
                        if len(v4) > 0 {
                            v4[0] = 620
                        }
                        _ = v66
                        _ = v67
                    case 2147483647:
                        v5 = append(v5, "foo bar")
                    case 32767:
                        var v73 uint16 = 32767
                        _ = v73
                    }
                    v74 := GenS8{F0: 119, F1: 815, F2: 0}
                    _ = v74
                }
                fmt.Println(fi61, fi61)
            } else {
                switch (fi61 | (fi61 ^ -685)) {
                case 0:
                    var v75 int8 = (v6 << ((0) & 63))
                    _ = v75
                case 9223372036854775807:
                    fmt.Println(fi61, fi62)
                case -9223372036854775808:
                    v6 = 100
                }
                v6 += 0
                switch "\n\t" {
                case "\n\t":
                    switch (((("\n\t" + "a") + "hello") + "foo bar") + "a") {
                    case "hello":
                        v6 = -8
                    case "a":
                        v76 := GenS3{F0: genFn0(-49, 955, 1), F1: -246}
                        _ = v76
                    default:
                        if false {
                            fmt.Println(fi61, fi62)
                            var v77 uint64 = 254
                            _ = v77
                        }
                    }
                }
            }
            _ = fi62
        }
        var v78 uint32 = 1
        _ = fi61
        _ = v78
    }
    v6 = (v2 << ((452) & 63))
    for fi79 := 0; fi79 < 6; fi79++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = append(v1, 2)
        _ = fi79
    }
    var v80 int64 = 549
    var v81 bool = ((42 <= 1) || false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v80
    _ = v81
}
