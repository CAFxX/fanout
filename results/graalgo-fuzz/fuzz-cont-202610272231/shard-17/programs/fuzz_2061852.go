package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int; F1 int64; F2 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 int64; F2 uint }
func genFn0() int16 {
    switch 32768 {
    case -2147483648:
        v7 := []int{468, 9223372036854775807}
        _ = v7
    case 2147483647:
        var v8 float64 = (3.14159 + (-2.25 / ((1e10 - -2.25) * (0.1 / -2.25))))
        _ = v8
    default:
        var v9 int64 = 42
        _ = v9
    }
    var v10 int64 = 766
    v11 := []int{-2147483648, -9223372036854775808}
    v12 := GenS1{F0: 65535, F1: 127, F2: false}
    _ = v10
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 int32) uint8 {
    v13 := GenS1{F0: 0, F1: 127, F2: (!(18446744073709551615 >= 394))}
    v14 := GenS1{F0: 393, F1: 100, F2: ((-1 <= 32767) || ((-1 >= 9223372036854775807) || ("" <= "")))}
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := []int{fi15, (fi15 + fi15), fi15}
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS1{F0: 980, F1: 127, F2: (1 < 100)}
            var v19 uint16 = 397
            var v20 float32 = 0.0
            switch (fi15 & (^(fi15))) {
            case 32768:
                var v21 int32 = (p0 & (^(p0)))
                if (-1 == genFn0()) {
                    var v22 uint16 = (^(v19))
                    v23 := GenS1{F0: 65535, F1: -61, F2: (3.14159 == 3.14159)}
                    var v24 uint16 = 65535
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    if false {
                        if ((!false) || ((false || false) && (float32(float64(1.5)) > float32((1.5 * -2.25))))) {
                            switch "foo bar" {
                            case "foo bar":
                                v21 -= 247
                            case "":
                                v25 := GenS1{F0: (v19 + (v19 | 611)), F1: 100, F2: (!((v19 + 114) < v19))}
                                _ = v25
                            default:
                                var v26 float64 = 0.1
                                _ = v26
                            }
                            v27 := GenS1{F0: (v19 + v19), F1: 127, F2: (184 != 76)}
                            p0 = 103
                            _ = v27
                        } else {
                            v20 = v20
                        }
                    } else {
                        v28 := GenS2{F0: "", F1: false, F2: -72}
                        _ = v28
                    }
                    if true {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    } else {
                        fmt.Println(p0)
                        v19 = (^(v19))
                    }
                    var v29 uint16 = (^(v19))
                    _ = v29
                }
                v30 := GenS2{F0: ("\n\t" + ""), F1: (true || false), F2: 97}
                _ = v21
                _ = v30
            case 1:
                if true {
                    v20 = ((float32(0.0) * (v20 - (1.5 + v20))) * ((((v20 - 3.14159) / 1e10) * (v20 / (0.1 + v20))) + (v20 + v20)))
                    p0 -= -9
                } else {
                    v20 -= float32((((1.5 - 0.1) + -2.25) / 1.5))
                    var v31 bool = (18446744073709551615 == 2)
                    v19 = (v19 * (v19 * v19))
                    v20 += v20
                    _ = v31
                }
            }
            var v32 int64 = 315
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v32
        }
        switch ("a" + "") {
        case "\n\t":
            switch (fi15 ^ (fi15 % (((fi15 % (((fi15 & fi15)) | 1))) | 1))) {
            case -992:
                v33 := GenS2{F0: "foo bar", F1: false, F2: 1}
                _ = v33
            case -1:
                v16 = append(v16, (fi15 - 42))
            case 394:
                v34 := GenS3{F0: 360, F1: (p0 + p0), F2: -3}
                _ = v34
            }
        default:
            var v35 string = "\n\t"
            _ = v35
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int = 32767
    var v4 uint8 = 0
    v5 := []int{}
    v6 := make(map[string]uint32)
    var v36 string = "foo bar"
    for fi37 := 0; fi37 < 15; fi37++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi37
    }
    switch 217 {
    case 32768:
        var v38 bool = false
        _ = v38
    }
    var v39 int16 = -32768
    var v40 string = "foo bar"
    v36 = ""
    for fi41 := 0; fi41 < 16; fi41++ {
        if true {
            v6["hello"] = 65535
            for fi42 := 0; fi42 < 11; fi42++ {
                if false {
                    fmt.Println(v36, fi41)
                    v40 = v40
                    v39 = v39
                } else {
                    v3 += (-(v3))
                    v43 := GenS3{F0: 9223372036854775807, F1: -2147483648, F2: 1}
                    _ = v43
                }
                v36 = v36
                _ = fi42
            }
            var v44 uint16 = 258
            _ = v44
        }
        for fi45 := 0; fi45 < 7; fi45++ {
            v39 = (v39 * 611)
            _ = fi45
        }
        _ = fi41
    }
    switch (v3 ^ v3) {
    case 0:
        fmt.Println(v3)
    case 32768:
        v46 := GenS1{F0: 32767, F1: 20, F2: (v3 <= v3)}
        _ = v46
    }
    v47 := GenS1{F0: 511, F1: 79, F2: ((v39 ^ genFn0()) >= 1000)}
    if false {
        if (v39 == (v39 * v39)) {
            v1[(v3 ^ v3)] = (v40 + v36)
            v48 := GenS2{F0: (v36 + v40), F1: (!(v3 <= (^(v3)))), F2: 114}
            _ = v48
        }
        if (true || false) {
            v49 := GenS4{F0: 654, F1: 2}
            _ = v49
        }
        if true {
            var v50 string = v36
            v4 -= v4
            _ = v50
        } else {
            v51 := GenS3{F0: 677, F1: 660, F2: -128}
            for fi52 := 0; fi52 < 15; fi52++ {
                v53 := GenS5{F0: v4, F1: 4294967295}
                fmt.Println(v40, v4, v4)
                var v54 uint64 = 4294967295
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = v51
        }
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        v56 := GenS6{F0: 32767, F1: 277, F2: 2}
        var v57 int16 = 0
        switch v40 {
        case "\n\t":
            fmt.Println(v57)
        default:
            v36 = "foo bar"
        }
        _ = fi55
        _ = v56
        _ = v57
    }
    if (true || (!true)) {
        v58 := GenS3{F0: -546, F1: 994, F2: 0}
        _ = v58
    } else {
        if false {
            v59 := GenS2{F0: "", F1: (0 >= 42), F2: 42}
            _ = v59
        } else {
            switch (v3 | v3) {
            case 1:
                var v60 float32 = (0.1 / float32(((3.14159 * 1.5) * (3.14159 / -2.25))))
                _ = v60
            case 32768:
                v39 += (v39 * v39)
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        switch v36 {
        case "\n\t":
            fmt.Println(v39, v4, v3)
        case "a":
            v61 := GenS7{F0: (v3 - (v3 % (((v3 - v3)) | 1)))}
            _ = v61
        case "hello":
            if (!(false && false)) {
                if true {
                    v4 += (^(v4))
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v40 = v36
                        v4 += 0
                        _ = fi62
                    }
                } else {
                    if (0.1 != (((0.1 / (1.5 - 0.1)) - 1.5) / 1.0)) {
                        v3 = (v3 + (v3 + v3))
                        if (4294967295 != 42) {
                            v3 += -857
                            for fi63 := 0; fi63 < 8; fi63++ {
                                var v64 float32 = 0.0
                                _ = len(v5)
                                var v65 float32 = v64
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            v36 = v36
                        }
                        v3 = v3
                    }
                }
                for fi66 := 0; fi66 < 5; fi66++ {
                    if ((224 <= 214) || (true || (false || (1 != -621)))) {
                        var v67 bool = false
                        for fi68 := 0; fi68 < 19; fi68++ {
                            v5 = append(v5, -2147483648)
                            switch v36 {
                            case "\n\t":
                                var v69 int32 = 1
                                _ = v69
                            case "":
                                v4 += (v4 - v4)
                            }
                            switch (fi68 ^ v3) {
                            case 32768:
                                if v67 {
                                    var v70 int = -47
                                    _ = v70
                                } else {
                                    if ((v39 * (v39 * (-(v39)))) >= -541) {
                                        v71 := v1[278]
                                        var v72 uint16 = 380
                                        _ = v71
                                        _ = v72
                                    }
                                    if (-473 != 772) {
                                        var v73 int32 = 369
                                        _ = v73
                                    }
                                    var v74 float32 = ((1e10 / (float32(0.1) + 1.5)) - 1.5)
                                    _ = v74
                                }
                            case -9223372036854775808:
                                delete(v1, (fi66 | fi66))
                            case 128:
                                var v75 int32 = -2147483648
                                _ = v75
                            }
                            _ = fi68
                        }
                        switch ((v36 + v40) + v36) {
                        case "a":
                            fmt.Println(v40, v36)
                        case "hello":
                            var v76 int = (fi66 - fi66)
                            _ = v76
                        }
                        _ = v67
                    }
                    for fi77 := 0; fi77 < 11; fi77++ {
                        fmt.Println(v36)
                        _ = fi77
                    }
                    _ = fi66
                }
                for fi78 := 0; fi78 < 10; fi78++ {
                    v79 := GenS5{F0: (v4 << ((65535) & 63)), F1: 835}
                    var v80 int64 = 42
                    _ = fi78
                    _ = v79
                    _ = v80
                }
            }
        }
        v4 += 1
    }
    if (true || true) {
        if false {
            if true {
                v39 -= (v39 & (v39 % ((genFn0()) | 1)))
            }
            var v81 float64 = 3.14159
            _ = v81
        } else {
            v5 = append(v5, (-(v3)))
        }
        v36 = v40
        delete(v2, v40)
    }
    v82 := GenS8{F0: 42, F1: 9223372036854775807, F2: 18446744073709551615}
    var v83 int = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
    _ = v82
    _ = v83
}
