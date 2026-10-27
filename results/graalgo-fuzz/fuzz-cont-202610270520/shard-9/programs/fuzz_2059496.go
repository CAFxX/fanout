package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int8; F2 uint }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int, p1 bool, p2 bool) int32 {
    p2 = p2
    v7 := []int{(p0 / (((p0 * 4294967295)) | 1))}
    v8 := GenS1{F0: 2147483647}
    if ((!false) && (p2 && (!p1))) {
        p0 += p0
    }
    p1 = ((("hello" + "foo bar") + "hello") > ("" + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 uint8) uint16 {
    p0 = (p0 / ((p0) | 1))
    v9 := []int{127, -1}
    v10 := GenS2{F0: (("\n\t" + "a") + "hello"), F1: 2147483647}
    fmt.Println(p0)
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS3{F0: false, F1: 103, F2: ((3.14159 / float32(3.14159)) != 0.1)}
        switch 32768 {
        case 1:
            var v13 uint = 2
            _ = v13
        default:
            if len(v9) > 0 {
                v9[0] = fi11
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -223
    var v3 uint32 = 4294967295
    v4 := []string{}
    var v5 int8 = 42
    v6 := make(map[int]string)
    v3 += v3
    v1 = v1
    if ((v1 / ((938) | 1)) != v1) {
        v5 = -84
    } else {
        var v14 int = -2147483648
        for fi15 := 0; fi15 < 13; fi15++ {
            v4 = append(v4, ("hello" + "foo bar"))
            fmt.Println(fi15, v2)
            _ = fi15
        }
        v2 = (v2 << ((65535) & 63))
        _ = v14
    }
    v3 -= v3
    if (!(!(!(!true)))) {
        if false {
            fmt.Println(v1)
        } else {
            v5 = v5
            switch "\n\t" {
            case "hello":
                v16 := GenS2{F0: "hello", F1: 788}
                _ = v16
            case "":
                switch -45 {
                case 127:
                    switch "hello" {
                    case "a":
                        v1 -= (v1 | (v1 << ((4294967295) & 63)))
                    case "":
                        var v17 float64 = float64(3.14159)
                        _ = v17
                    }
                case -353:
                    v18 := GenS3{F0: false, F1: 0, F2: (uint64(v3) <= 18446744073709551615)}
                    _ = v18
                }
            }
        }
    }
    if true {
        v2 = v2
        var v19 uint = 1
        switch int(v3) {
        case 32767:
            v3 = (v3 * 2)
        case 32768:
            _ = len(v6)
        case 42:
            switch 984 {
            case 2147483647:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 457:
                v19 = uint(v3)
            case -2147483648:
                v20 := GenS3{F0: true, F1: 128, F2: true}
                _ = v20
            }
        }
        _ = v19
    }
    v1 = v1
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 float32 = 0.0
        _ = fi21
        _ = v22
    }
    var v23 string = "\n\t"
    fmt.Println(v3, v2)
    for fi24 := 0; fi24 < 7; fi24++ {
        v3 -= uint32(v1)
        _ = fi24
    }
    v25 := GenS2{F0: v23, F1: 42}
    if (1e10 <= 3.14159) {
        v3 = (v3 << ((169) & 63))
        switch -1 {
        case 1:
            for fi26 := 0; fi26 < 2; fi26++ {
                switch fi26 {
                case 32767:
                    if true {
                        for fi27 := 0; fi27 < 19; fi27++ {
                            v1 += 592
                            _ = fi27
                        }
                        v28 := GenS4{F0: v1, F1: 850}
                        v29 := GenS2{F0: v23, F1: v3}
                        _ = v28
                        _ = v29
                    } else {
                        if (-899 >= -896) {
                            for fi30 := 0; fi30 < 5; fi30++ {
                                fmt.Println(v5, fi26, fi26)
                                _ = fi30
                            }
                            v31 := v6[(fi26 >> ((4294967295) & 63))]
                            if (0 == 1) {
                                _ = len(v6)
                            }
                            _ = v31
                        } else {
                            v1 += v1
                            var v32 int32 = 42
                            _ = v32
                        }
                        switch v23 {
                        case "hello":
                            var v33 uint8 = 1
                            _ = v33
                        case "foo bar":
                            v3 = 869
                        case "":
                            v23 = v23
                        default:
                            for fi34 := 0; fi34 < 16; fi34++ {
                                var v35 uint16 = (v1 + v1)
                                v23 = "foo bar"
                                v3 += (v3 / (((v3 - v3)) | 1))
                                v36 := GenS3{F0: false, F1: 127, F2: (1e10 >= 0.0)}
                                v1 += v35
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                        if false {
                            var v37 bool = true
                            v38 := GenS1{F0: genFn0(fi26, (417 != uint(v3)), ((1e10 / 1.0) != 3.14159))}
                            v23 = "foo bar"
                            _ = v37
                            _ = v38
                        }
                    }
                case 42:
                    if ((939 > 42) || false) {
                        v1 += (v1 - v1)
                    } else {
                        v2 = (v2 << ((2) & 63))
                        _ = len(v6)
                    }
                case 4294967295:
                    if true {
                        var v39 float32 = (1e10 - 1e10)
                        var v40 float64 = float64(0.1)
                        v23 = (v23 + v23)
                        _ = v39
                        _ = v40
                    }
                default:
                    _ = len(v4)
                }
                for fi41 := 0; fi41 < 2; fi41++ {
                    var v42 int8 = 102
                    v1 -= v1
                    v43 := GenS5{F0: 255, F1: 0, F2: 42}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                for fi44 := 0; fi44 < 20; fi44++ {
                    v45 := GenS6{F0: 0}
                    if len(v4) > 0 {
                        v4[0] = v23
                    }
                    if (!("hello" == v23)) {
                        _ = len(v4)
                    } else {
                        v46 := v6[435]
                        switch fi26 {
                        case 517:
                            var v47 bool = ((v1 | v1) < v1)
                            _ = v47
                        case -9223372036854775808:
                            _ = len(v4)
                        default:
                            var v48 int8 = 127
                            _ = v48
                        }
                        var v49 string = (v46 + v46)
                        _ = v46
                        _ = v49
                    }
                    _ = fi44
                    _ = v45
                }
                _ = fi26
            }
        case 128:
            var v50 string = "a"
            _ = v50
        }
        var v51 bool = true
        _ = v51
    }
    v2 += (v2 - v2)
    var v52 int16 = 757
    v1 = (v1 | uint16(v3))
    switch 9223372036854775807 {
    case -2147483648:
        if (0.1 != 1.5) {
            v53 := GenS5{F0: 89, F1: 65535, F2: 0}
            v54 := v6[42]
            v55 := GenS7{F0: v5, F1: 123, F2: 940}
            _ = v53
            _ = v54
            _ = v55
        } else {
            v56 := GenS3{F0: true, F1: 162, F2: (!true)}
            var v57 float32 = 1.5
            if (1e10 >= 0.0) {
                for fi58 := 0; fi58 < 6; fi58++ {
                    for fi59 := 0; fi59 < 15; fi59++ {
                        if true {
                            v6[(fi59 ^ fi58)] = "a"
                            v57 = 0.0
                            var v60 float32 = v57
                            fmt.Println(fi59, v23, v2)
                            var v61 int = 798
                            _ = v60
                            _ = v61
                        } else {
                            v1 = v1
                            var v62 int16 = v52
                            _ = len(v4)
                            _ = v62
                        }
                        _ = fi59
                    }
                    switch -1 {
                    case 672:
                        var v63 int8 = (v5 >> ((953) & 63))
                        _ = v63
                    case 702:
                        v52 = (v52 ^ v52)
                    case 128:
                        var v64 float64 = ((3.14159 / float64(v57)) * (1e10 / 0.1))
                        _ = v64
                    default:
                        v5 = 100
                    }
                    for fi65 := 0; fi65 < 11; fi65++ {
                        switch v23 {
                        case "foo bar":
                            switch v23 {
                            case "\n\t":
                                v1 = v1
                            case "hello":
                                v52 = v52
                            default:
                                if (255 >= 39) {
                                    var v66 uint16 = (v1 - 32767)
                                    v23 = (v23 + v23)
                                    _ = v66
                                } else {
                                    var v67 float64 = 0.0
                                    _ = v67
                                }
                            }
                            fmt.Println(v52)
                            if ((v5 >> ((uint(v1)) & 63)) <= 127) {
                                var v68 uint32 = 2
                                _ = v68
                            } else {
                                v6[(^(fi58))] = "hello"
                                switch "" {
                                case "\n\t":
                                    if (((v23 + v23) + v23) > ((v23 + v23) + v23)) {
                                        v23 = "foo bar"
                                        v69 := GenS8{F0: genFn1(127)}
                                        v3 -= v3
                                        _ = v69
                                    }
                                    fmt.Println(fi58)
                                    v5 += (v5 ^ 1)
                                case "foo bar":
                                    var v70 uint16 = genFn1(128)
                                    _ = v70
                                }
                                switch (fi65 & 4294967295) {
                                case -179:
                                    var v71 uint = 42
                                    _ = v71
                                case -483:
                                    v52 = -175
                                default:
                                    fmt.Println(v52, v2, v3)
                                }
                            }
                            var v72 int64 = v2
                            _ = v72
                        case "a":
                            var v73 uint8 = 42
                            _ = v73
                        default:
                            if (!false) {
                                v52 = v52
                                v3 += v3
                                v74 := v6[fi65]
                                v3 += (v3 % ((v3) | 1))
                                _ = v74
                            }
                        }
                        v23 = v23
                        _ = fi65
                    }
                    _ = fi58
                }
            }
            _ = v56
            _ = v57
        }
    default:
        var v75 int = 32768
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
