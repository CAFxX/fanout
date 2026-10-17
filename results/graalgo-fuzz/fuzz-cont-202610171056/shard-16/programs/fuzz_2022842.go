package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int, p1 bool) uint16 {
    v6 := []int{p0, p0}
    var v7 int32 = -935
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int8, p1 string) int32 {
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p1)
        fmt.Println(fi8)
        _ = fi8
    }
    if (true || (!(!(-1 < -1)))) {
        v9 := []int{9223372036854775807, -1, -528}
        _ = v9
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 int32 = 42
    var v2 int8 = 127
    v3 := []int{127, 2147483647}
    v4 := make(map[uint32]int)
    var v5 string = ("a" + "a")
    delete(v4, 291)
    v10 := GenS1{F0: (true || false)}
    v1 = 0
    v1 -= (-(v1))
    var v11 float64 = (0.0 * float64(0.0))
    v11 -= float64(-2.25)
    if (!true) {
        var v12 int = 4294967295
        switch ((("a" + (v5 + "")) + "hello") + "") {
        case "hello":
            switch (v12 / ((v12) | 1)) {
            case 9223372036854775807:
                v11 += v11
            default:
                if (v5 <= v5) {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v15 := GenS2{F0: (v1 + v1), F1: 42}
                            _ = len(v3)
                            v5 = ""
                            var v16 uint32 = 2147483647
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 float64 = (float64((0.1 * 0.0)) - float64(float32(v11)))
                        switch (fi13 ^ 0) {
                        case 127:
                            v18 := GenS1{F0: true}
                            _ = v18
                        case 42:
                            v12 += (v12 << ((970) & 63))
                        case -902:
                            var v19 int = (-(fi13))
                            _ = v19
                        default:
                            if ((false && (-782 != 9223372036854775807)) && (!true)) {
                                switch fi13 {
                                case -2147483648:
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        v21 := v4[2]
                                        _ = fi20
                                        _ = v21
                                    }
                                default:
                                    v4[739] = v12
                                }
                                v22 := GenS3{F0: 32767, F1: 4294967295, F2: -77}
                                _ = v22
                            }
                        }
                        _ = fi13
                        _ = v17
                    }
                    var v23 string = v5
                    _ = v23
                }
            }
        case "foo bar":
            var v24 uint = 764
            _ = v24
        default:
            v25 := GenS3{F0: 42, F1: 185, F2: (v2 + v2)}
            _ = v25
        }
        switch (v12 + v12) {
        case 128:
            var v26 int8 = (v2 - 127)
            _ = v26
        case 42:
            switch 32767 {
            case 1:
                if (65535 < 42) {
                    fmt.Println(v1)
                    var v27 float64 = 0.0
                    _ = v27
                }
            }
        }
        _ = v12
    } else {
        v2 -= v2
        var v28 int = 9223372036854775807
        fmt.Println(v2, v1)
        _ = v28
    }
    var v29 int = 9223372036854775807
    v5 = v5
    switch (v29 + (v29 + v29)) {
    case 471:
        v29 = -473
    case -9223372036854775808:
        if (!(genFn1((v2 * (v2 | (v2 / (((v2 & v2)) | 1)))), v5) == -2147483648)) {
            v30 := v4[2147483647]
            v31 := GenS4{F0: 0, F1: 2147483647, F2: 44}
            if (-10 > -1) {
                var v32 int8 = (v2 % (((v2 - v2)) | 1))
                _ = v32
            } else {
                v30 += v29
            }
            _ = v30
            _ = v31
        } else {
            v11 += v11
            for fi33 := 0; fi33 < 20; fi33++ {
                if (true || (1e10 != 3.14159)) {
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v2 = (v2 | (v2 & v2))
                        v35 := GenS2{F0: (-(v1)), F1: 65535}
                        if (true || (1 != 701)) {
                            switch (fi33 << ((18446744073709551615) & 63)) {
                            case -412:
                                v1 = 1
                            case 4294967295:
                                fmt.Println(v1, fi34)
                            }
                            v5 = v5
                            v2 = 127
                        } else {
                            v36 := GenS2{F0: 0, F1: 667}
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                } else {
                    v11 += float64(1e10)
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 int64 = 9223372036854775807
                        fmt.Println(v11, v1, v11)
                        _ = fi37
                        _ = v38
                    }
                    switch v5 {
                    case "foo bar":
                        v2 += (v2 & v2)
                    case "\n\t":
                        if false {
                            v11 = v11
                            fmt.Println(v11, v5)
                        }
                    case "a":
                        var v39 float64 = v11
                        _ = v39
                    default:
                        fmt.Println(v2)
                    }
                }
                var v40 bool = (false || true)
                v29 = (v29 - 477)
                _ = fi33
                _ = v40
            }
        }
    case 113:
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS2{F0: int32(v29), F1: 896}
            if (true || (1 > 2)) {
                if (338 >= 229) {
                    if ((true && (false || true)) || (1 > 2147483647)) {
                        if true {
                            var v43 int16 = -1
                            if true {
                                var v44 int64 = 1
                                var v45 bool = (false && true)
                                _ = v44
                                _ = v45
                            } else {
                                var v46 string = "hello"
                                v47 := GenS3{F0: 1, F1: 108, F2: v2}
                                _ = v46
                                _ = v47
                            }
                            _ = v43
                        } else {
                            for fi48 := 0; fi48 < 12; fi48++ {
                                var v49 uint8 = 128
                                v50 := GenS5{F0: -1, F1: 512}
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v3 = append(v3, int(v1))
                            if true {
                                for fi51 := 0; fi51 < 11; fi51++ {
                                    if true {
                                        var v52 uint64 = 42
                                        _ = v52
                                    }
                                    _ = fi51
                                }
                                for fi53 := 0; fi53 < 14; fi53++ {
                                    var v54 uint64 = 338
                                    v5 = ("a" + "")
                                    _ = fi53
                                    _ = v54
                                }
                                fmt.Println(v2, fi41)
                            }
                        }
                        _ = len(v4)
                    }
                }
                v29 += (-(v29))
                v55 := GenS6{F0: (^(v1)), F1: 42, F2: -9223372036854775808}
                _ = v55
            } else {
                var v56 uint16 = 42
                switch (v5 + ("a" + "")) {
                case "":
                    _ = len(v3)
                default:
                    var v57 uint8 = 1
                    _ = v57
                }
                _ = v56
            }
            v58 := GenS3{F0: 109, F1: 52, F2: (v2 * v2)}
            _ = fi41
            _ = v42
            _ = v58
        }
    }
    for fi59 := 0; fi59 < 3; fi59++ {
        switch "foo bar" {
        case "":
            v60 := GenS7{F0: 224, F1: 255}
            _ = v60
        case "a":
            if (42 >= 450) {
                v61 := GenS8{F0: 1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v61
            } else {
                for fi62 := 0; fi62 < 20; fi62++ {
                    var v63 bool = true
                    if (444 <= 1) {
                        fmt.Println(v2, v11)
                        var v64 bool = v63
                        if (!v63) {
                            var v65 uint32 = 65535
                            if len(v3) > 0 {
                                v3[0] = (v29 * (fi62 % (((^(v29))) | 1)))
                            }
                            if len(v3) > 0 {
                                v3[0] = fi62
                            }
                            _ = v65
                        }
                        _ = v64
                    }
                    var v66 float32 = 1e10
                    _ = fi62
                    _ = v63
                    _ = v66
                }
            }
        case "foo bar":
            v11 = (((0.1 / float64(3.14159)) * 1.5) + 1e10)
        }
        if true {
            if ((v2 << ((859) & 63)) > (-(v2))) {
                if false {
                    var v67 float64 = v11
                    for fi68 := 0; fi68 < 5; fi68++ {
                        if (v2 <= (v2 % ((v2) | 1))) {
                            v11 -= (0.0 - v11)
                            var v69 uint16 = 32767
                            _ = v69
                        } else {
                            v29 += -447
                            _ = len(v4)
                            v67 -= 0.1
                        }
                        v70 := GenS8{F0: genFn0(v29, true)}
                        _ = fi68
                        _ = v70
                    }
                    _ = len(v3)
                    _ = v67
                } else {
                    v1 -= (v1 << ((4294967295) & 63))
                }
            }
            if true {
                var v71 int32 = genFn1(-28, v5)
                var v72 uint8 = 42
                _ = v71
                _ = v72
            }
        }
        _ = fi59
    }
    var v73 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v73
}
