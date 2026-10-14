package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 int }
func genFn0(p0 string, p1 uint64) uint32 {
    var v5 float64 = 1e10
    var v6 bool = ((^(p1)) == (p1 + p1))
    switch -9223372036854775808 {
    case 451:
        v7 := GenS1{F0: 498, F1: 42}
        _ = v7
    case 293:
        switch "hello" {
        case "foo bar":
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint32 = 65535
                if v6 {
                    if v6 {
                        v10 := []int{(fi8 >> ((4294967295) & 63)), fi8, int(p1)}
                        _ = v10
                    }
                } else {
                    v11 := GenS2{F0: (p0 + p0)}
                    v12 := GenS1{F0: 65535, F1: 0}
                    v13 := GenS1{F0: 4294967295, F1: 42}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := []int{829, fi8, int(v9), (fi8 + (fi8 % ((fi8) | 1)))}
                _ = fi8
                _ = v14
                _ = v9
            }
        case "a":
            p0 = p0
        case "":
            fmt.Println(v5, p1)
        }
    }
    p0 = "a"
    var v15 uint64 = p1
    _ = p0
    _ = p1
    _ = v15
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 int64, p2 uint32) int16 {
    switch "" {
    case "a":
        fmt.Println(p0, p2)
    default:
        var v16 int16 = 32767
        _ = v16
    }
    var v17 int8 = 103
    if (1e10 >= float64(0.0)) {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS1{F0: 65535, F1: 2}
            v20 := GenS2{F0: ("" + "a")}
            _ = fi18
            _ = v19
            _ = v20
        }
        switch 4294967295 {
        case -106:
            v21 := []int{}
            _ = v21
        case -498:
            switch -983 {
            case -1:
                var v22 float64 = 0.1
                _ = v22
            default:
                p0 += p0
            }
        default:
            v23 := GenS3{F0: uint(p0), F1: 1}
            _ = v23
        }
        fmt.Println(p2, p2)
    } else {
        var v24 uint16 = 506
        _ = v24
    }
    p1 += p1
    v25 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v25
    return 0
}

func main() {
    v1 := []int{2147483647, -781, -905, -1}
    v2 := []uint32{}
    var v3 uint32 = 42
    v4 := make(map[int]string)
    v26 := GenS4{F0: -1, F1: 4294967295, F2: -2147483648}
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 int = (fi27 << ((18446744073709551615) & 63))
        for fi29 := 0; fi29 < 2; fi29++ {
            v3 = 2147483648
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    var v30 uint16 = 32767
    switch "a" {
    case "foo bar":
        var v31 uint64 = 1
        _ = v31
    case "a":
        switch "" {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                v30 = v30
            }
        case "":
            v3 -= v3
        case "a":
            if (823 == 1) {
                var v32 float64 = 1e10
                var v33 uint64 = 0
                if len(v2) > 0 {
                    v2[0] = v3
                }
                _ = v32
                _ = v33
            }
        default:
            delete(v4, 1)
        }
    case "\n\t":
        _ = len(v4)
    }
    v34 := v4[9223372036854775807]
    var v35 uint32 = v3
    v34 = "foo bar"
    v3 += genFn0(v34, 597)
    var v36 uint32 = v35
    v37 := GenS5{F0: true, F1: v35, F2: 0}
    switch ("a" + v34) {
    case "hello":
        if (!(!(583 < -266))) {
            if ((^(v35)) < (^(v36))) {
                var v38 uint8 = 255
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v34 = v34
                _ = v38
            } else {
                v39 := GenS6{F0: 138, F1: -561, F2: genFn1(0, -25, 188)}
                _ = v39
            }
            v34 = ""
            if true {
                fmt.Println(v35)
                if ((float32(3.14159) + 3.14159) >= 0.1) {
                    v30 = v30
                    for fi40 := 0; fi40 < 17; fi40++ {
                        fmt.Println(v35)
                        _ = fi40
                    }
                    if (2 != 128) {
                        switch 1 {
                        case -2147483648:
                            if true {
                                var v41 bool = ((!(!false)) || (!true))
                                v35 = v35
                                switch ("foo bar" + v34) {
                                case "foo bar":
                                    var v42 uint = 85
                                    _ = v42
                                default:
                                    v43 := GenS1{F0: 195, F1: (v30 + uint16(v35))}
                                    _ = v43
                                }
                                _ = v41
                            } else {
                                v34 = v34
                            }
                        }
                        v2 = append(v2, genFn0((v34 + (v34 + (v34 + v34))), 18446744073709551615))
                    }
                } else {
                    v30 += (v30 & v30)
                    for fi44 := 0; fi44 < 4; fi44++ {
                        var v45 int64 = 0
                        for fi46 := 0; fi46 < 13; fi46++ {
                            var v47 int32 = 1
                            _ = fi46
                            _ = v47
                        }
                        v45 = v45
                        _ = fi44
                        _ = v45
                    }
                    v48 := GenS7{F0: 65535, F1: 1000, F2: 42}
                    _ = v48
                }
                var v49 string = "foo bar"
                _ = v49
            }
        } else {
            v34 = v34
        }
    case "foo bar":
        fmt.Println(v3)
    }
    if true {
        v50 := v4[0]
        switch 2147483647 {
        case -554:
            if (1.5 > 0.0) {
                switch (v34 + v34) {
                case "a":
                    var v51 uint8 = 0
                    _ = v51
                default:
                    for fi52 := 0; fi52 < 3; fi52++ {
                        delete(v4, fi52)
                        _ = fi52
                    }
                }
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := GenS2{F0: ("\n\t" + "a")}
                    _ = fi53
                    _ = v54
                }
                fmt.Println(v36)
            } else {
                _ = len(v4)
                if false {
                    v35 -= 4294967295
                    for fi55 := 0; fi55 < 17; fi55++ {
                        if true {
                            fmt.Println(fi55, v3)
                            if (false || ((-2.25 + (float32((((3.14159 * 1e10) / 3.14159) * 0.1)) + 0.1)) <= 1.5)) {
                                var v56 uint8 = 2
                                fmt.Println(v35)
                                _ = v56
                            } else {
                                var v57 float32 = 3.14159
                                v58 := GenS8{F0: 1, F1: v3}
                                var v59 string = v50
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            switch (v34 + (v50 + (("\n\t" + v34) + v50))) {
                            case "":
                                v50 = "\n\t"
                            case "\n\t":
                                if false {
                                    v35 += v36
                                } else {
                                    var v60 string = v50
                                    v34 = "\n\t"
                                    for fi61 := 0; fi61 < 19; fi61++ {
                                        v34 = ((v50 + "foo bar") + "a")
                                        _ = fi61
                                    }
                                    _ = v60
                                }
                            default:
                                for fi62 := 0; fi62 < 14; fi62++ {
                                    var v63 uint = 255
                                    v64 := GenS9{F0: (fi55 | 337)}
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi55
                    }
                } else {
                    for fi65 := 0; fi65 < 12; fi65++ {
                        var v66 float64 = 0.0
                        v3 = 0
                        v3 -= 65535
                        _ = fi65
                        _ = v66
                    }
                }
                var v67 uint64 = 747
                _ = v67
            }
        case -235:
            v68 := GenS9{F0: 4294967295}
            _ = v68
        default:
            var v69 int8 = 0
            _ = v69
        }
        _ = v50
    } else {
        v3 -= genFn0(v34, 0)
    }
    for fi70 := 0; fi70 < 1; fi70++ {
        v36 += (v35 % ((v3) | 1))
        v3 += (v36 * (v3 >> ((1) & 63)))
        v30 = (v30 << ((255) & 63))
        _ = fi70
    }
    v36 = 0
    for fi71 := 0; fi71 < 20; fi71++ {
        delete(v4, fi71)
        _ = fi71
    }
    if (v34 <= v34) {
        if len(v2) > 0 {
            v2[0] = 65535
        }
    } else {
        fmt.Println(v3, v34)
    }
    var v72 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v72
}
