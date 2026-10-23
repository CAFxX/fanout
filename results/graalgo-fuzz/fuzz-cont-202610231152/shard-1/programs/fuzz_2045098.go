package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int; F1 int; F2 int64 }
type GenS9 struct { F0 string; F1 string; F2 string }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint }
func genFn0(p0 uint16) int32 {
    v7 := GenS1{F0: "hello"}
    var v8 float32 = -2.25
    v8 = (float32(1e10) + float32(((float64(3.14159) + (-2.25 - 0.0)) / float64((0.1 * float32(1e10))))))
    fmt.Println(v8, p0)
    var v9 uint8 = 128
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint32) uint64 {
    v10 := GenS2{F0: 1, F1: 1, F2: 2}
    var v11 uint64 = 4294967295
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(p0, p0, v11)
        p0 += 1
        _ = fi12
    }
    fmt.Println(v11, v11)
    if (2 < 0) {
        var v13 string = "\n\t"
        if (-36 < genFn0(2)) {
            v14 := GenS1{F0: "\n\t"}
            _ = v14
        } else {
            switch v13 {
            case "":
                switch 128 {
                case 2147483647:
                    if (-1 > -945) {
                        fmt.Println(v11, v11)
                    }
                case -643:
                    v11 -= v11
                case 128:
                    v11 -= 42
                default:
                    if ((322 <= v11) || true) {
                        var v15 bool = (!true)
                        if ((^(p0)) <= 879) {
                            for fi16 := 0; fi16 < 1; fi16++ {
                                var v17 float32 = (-2.25 / 3.14159)
                                var v18 uint32 = p0
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = v15
                    }
                }
            case "hello":
                v19 := []int{157, -959, 128, -2147483648}
                _ = v19
            default:
                if true {
                    var v20 int32 = 1
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v11 -= v11
                        p0 = (p0 * p0)
                        v11 += v11
                        _ = fi21
                    }
                    _ = v20
                }
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 int32 = 2147483647
                v23 -= v23
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                for fi25 := 0; fi25 < 14; fi25++ {
                    switch 848 {
                    case -198:
                        if false {
                            v11 = (v11 << ((520) & 63))
                            var v26 int64 = -291
                            _ = v26
                        }
                    }
                    v27 := GenS2{F0: -1, F1: 1, F2: 50}
                    _ = fi25
                    _ = v27
                }
                fmt.Println(fi24, p0, fi24)
                _ = fi24
            }
        }
        if false {
            var v28 uint = 255
            _ = v28
        }
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 string = ("\n\t" + "")
    v4 := []string{v3, "hello"}
    v5 := []string{"a", v3, v3, (v3 + v3)}
    var v6 uint16 = 114
    fmt.Println(v3, v6)
    var v29 uint32 = 42
    if false {
        var v30 bool = false
        v2[(v29 + 143)] = 2147483647
        _ = v30
    } else {
        if true {
            v31 := GenS2{F0: 42, F1: -1, F2: (v6 + v6)}
            _ = v31
        }
        v6 += 65535
        if (!(0 <= 1)) {
            switch 946 {
            case -921:
                if (v3 <= "a") {
                    v32 := GenS3{F0: 128, F1: 2}
                    _ = v32
                }
            }
            var v33 uint8 = 10
            if len(v1) > 0 {
                v1[0] = 1
            }
            _ = v33
        }
    }
    fmt.Println(v3)
    v5 = append(v5, "")
    v3 = v3
    v34 := GenS2{F0: 2147483647, F1: 0, F2: 32767}
    fmt.Println(v3, v29, v29)
    switch ("" + v3) {
    case "foo bar":
        var v35 float64 = (-2.25 - (1.5 / (1e10 - (0.1 / (-2.25 + ((1e10 - (1e10 * 0.0)) + float64(float32((0.1 - 3.14159)))))))))
        _ = v35
    case "":
        v36 := GenS4{F0: genFn1(v29), F1: 2, F2: (1e10 < (0.1 / (1e10 + (1.5 * 1e10))))}
        _ = v36
    case "a":
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 int32 = 662
            v6 = 65535
            _ = fi37
            _ = v38
        }
    default:
        switch 128 {
        case 156:
            v39 := GenS5{F0: 1, F1: 42}
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 int = fi40
        switch (v41 << ((2) & 63)) {
        case 2147483647:
            fmt.Println(v6)
        case 32767:
            if true {
                if true {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        if (9223372036854775807 >= int64(v29)) {
                            if (42 <= 18446744073709551615) {
                                if ((v29 != v29) && (genFn1((^(v29))) > 4294967295)) {
                                    var v43 int32 = -2147483648
                                    _ = v43
                                }
                                v44 := GenS6{F0: (^(v6))}
                                _ = v44
                            } else {
                                fmt.Println(v29)
                            }
                            switch (((v3 + v3) + "\n\t") + v3) {
                            case "\n\t":
                                v29 = (v29 % ((uint32(v6)) | 1))
                            case "foo bar":
                                v45 := GenS7{F0: (v41 & fi42), F1: 0}
                                _ = v45
                            default:
                                v3 = ((v3 + v3) + "a")
                            }
                        } else {
                            _ = len(v4)
                        }
                        _ = fi42
                    }
                    v6 = v6
                    var v46 int16 = 1
                    _ = v46
                }
                v47 := GenS5{F0: 42, F1: 9223372036854775807}
                _ = v47
            } else {
                if (true || true) {
                    switch fi40 {
                    case -9223372036854775808:
                        if (false || ((1.5 - 3.14159) >= (float64((-2.25 + ((float32(-2.25) + 1.5) - 1.5))) - 0.0))) {
                            _ = len(v5)
                        }
                    case -2147483648:
                        var v48 uint32 = (v29 - (v29 | (^(v29))))
                        _ = v48
                    case 9223372036854775807:
                        var v49 int = -9223372036854775808
                        _ = v49
                    default:
                        v50 := v2[v29]
                        _ = v50
                    }
                }
                if (int8(fi40) > 1) {
                    var v51 int16 = 477
                    switch v41 {
                    case -272:
                        v52 := v2[v29]
                        _ = v52
                    case 42:
                        switch v3 {
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                for fi53 := 0; fi53 < 8; fi53++ {
                                    v41 += int(v51)
                                    var v54 float64 = (0.1 * -2.25)
                                    _ = fi53
                                    _ = v54
                                }
                            default:
                                v2[(^(v29))] = (v41 & fi40)
                            }
                        default:
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v51 = -32768
                                if true {
                                    v51 = (v51 << ((42) & 63))
                                } else {
                                    _ = len(v4)
                                    fmt.Println(v41)
                                    v51 -= -478
                                }
                                if ((-(v41)) != (-(v41))) {
                                    v6 -= v6
                                } else {
                                    v56 := GenS8{F0: (v41 + v41), F1: v41, F2: 399}
                                    v57 := GenS9{F0: "foo bar", F1: v3, F2: v3}
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi55
                            }
                        }
                    case 9223372036854775807:
                        var v58 int64 = 625
                        _ = v58
                    default:
                        var v59 uint8 = 231
                        _ = v59
                    }
                    switch v3 {
                    case "a":
                        for fi60 := 0; fi60 < 20; fi60++ {
                            if (true || true) {
                                fmt.Println(v6, v29, v51)
                                if (!false) {
                                    var v61 float32 = (-2.25 / 3.14159)
                                    switch v3 {
                                    case "":
                                        v62 := v2[969]
                                        _ = v62
                                    case "hello":
                                        v29 = uint32(v6)
                                    }
                                    v61 -= 0.0
                                    _ = v61
                                }
                                fmt.Println(v51)
                            } else {
                                var v63 uint16 = v6
                                _ = v63
                            }
                            _ = fi60
                        }
                    case "foo bar":
                        if (!false) {
                            if true {
                                var v64 float64 = ((0.1 / (0.1 * float64(3.14159))) / float64((((0.0 / 1.0) + 1.5) + (-2.25 - (((1.5 * (-2.25 + 0.1)) - (1e10 - (1.5 - 1e10))) / 1.5)))))
                                v41 = (fi40 + fi40)
                                _ = v64
                            } else {
                                v65 := GenS2{F0: -476, F1: 2147483647, F2: 65535}
                                switch (fi40 ^ v41) {
                                case 42:
                                    var v66 uint = 42
                                    _ = v66
                                }
                                if (!(float64((1e10 - (3.14159 * (1.5 / 3.14159)))) == 0.0)) {
                                    var v67 uint8 = 127
                                    _ = v67
                                }
                                _ = v65
                            }
                            v51 = v51
                        } else {
                            v6 = (v6 - (v6 | 0))
                            v68 := GenS9{F0: v3, F1: "", F2: "a"}
                            _ = v68
                        }
                    default:
                        if (!(false || (true || false))) {
                            if false {
                                fmt.Println(v51, v41)
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    v3 = v3
                                    _ = fi69
                                }
                            } else {
                                if (42 > 1) {
                                    var v70 int16 = (v51 * 32767)
                                    switch ((v3 + v3) + v3) {
                                    case "\n\t":
                                        if (-1 != -128) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v51, v70)
                                    case "hello":
                                        switch fi40 {
                                        case -9223372036854775808:
                                            for fi71 := 0; fi71 < 9; fi71++ {
                                                v72 := GenS6{F0: (v6 % ((386) | 1))}
                                                v73 := GenS10{F0: 597}
                                                _ = fi71
                                                _ = v72
                                                _ = v73
                                            }
                                        case 2147483647:
                                            v74 := GenS11{F0: 1, F1: (v6 + 42), F2: genFn0((v6 >> ((42) & 63)))}
                                            _ = v74
                                        case 566:
                                            fmt.Println(v51, v3, v6)
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = (fi40 / ((754) | 1))
                                            }
                                        }
                                    }
                                    var v75 int = (fi40 % (((fi40 | (v41 >> ((uint(v6)) & 63)))) | 1))
                                    _ = v70
                                    _ = v75
                                } else {
                                    var v76 float64 = (-2.25 - 1e10)
                                    _ = v76
                                }
                                var v77 uint = 255
                                _ = v77
                            }
                        }
                    }
                    _ = v51
                }
                v78 := GenS12{F0: (v6 >> ((2) & 63)), F1: 569, F2: 2}
                _ = v78
            }
        }
        fmt.Println(fi40, v6, v41)
        _ = fi40
        _ = v41
    }
    for fi79 := 0; fi79 < 20; fi79++ {
        fmt.Println(v3)
        fmt.Println(fi79, v6)
        _ = fi79
    }
    if false {
        switch 704 {
        case -2147483648:
            v5 = append(v5, (v3 + "a"))
        default:
            v3 = (v3 + v3)
        }
        _ = len(v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
