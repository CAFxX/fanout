package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int }
func genFn0(p0 int, p1 int16, p2 uint32) uint64 {
    fmt.Println(p0)
    p0 -= 1
    if false {
        var v6 uint8 = 48
        v7 := []int{-296}
        v6 = v6
        _ = v6
        _ = v7
    }
    var v8 uint8 = 2
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    var v3 float32 = -2.25
    var v4 uint64 = 1
    v5 := make(map[uint32]uint32)
    v11 := []int{}
    if false {
        if (!true) {
            _ = len(v11)
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            v3 = 0.0
            v1[""] = true
            v3 = -2.25
            _ = fi12
        }
    } else {
        if (!(true && (42 <= genFn0(v2, -1, 4294967295)))) {
            v13 := GenS2{F0: 127, F1: 255, F2: (-(v2))}
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 string = "foo bar"
            switch (fi14 / ((-196) | 1)) {
            case 955:
                if false {
                    if len(v11) > 0 {
                        v11[0] = 32768
                    }
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 bool = true
                        if v17 {
                            var v18 uint8 = 1
                            var v19 uint8 = 128
                            fmt.Println(v19, v2)
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int32 = -1
                            delete(v5, 42)
                            if v17 {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                v21 := GenS1{F0: 0, F1: 9223372036854775807, F2: 361}
                                _ = v21
                            }
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                    }
                    if (720 <= 1) {
                        v4 = genFn0(fi14, 32767, 42)
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := GenS3{F0: 42, F1: 42, F2: 0}
                            fmt.Println(v4)
                            _ = fi22
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v5[740] = 533
                            v5[65535] = 571
                            _ = fi24
                        }
                    }
                }
            case -127:
                if (2147483647 < 655) {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v11 = append(v11, (^(fi14)))
                        _ = fi25
                    }
                    var v26 float64 = float64(v3)
                    var v27 int32 = 871
                    _ = v26
                    _ = v27
                }
            case -383:
                _ = len(v11)
            default:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            switch v15 {
            case "a":
                for fi28 := 0; fi28 < 5; fi28++ {
                    if len(v11) > 0 {
                        v11[0] = (fi14 + (fi28 / ((v2) | 1)))
                    }
                    v2 -= (v2 * -624)
                    var v29 string = "\n\t"
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi14
            _ = v15
        }
        if (((522 == 42) || true) || true) {
            var v30 uint32 = 2147483647
            if true {
                if len(v11) > 0 {
                    v11[0] = (v2 & (v2 - v2))
                }
                for fi31 := 0; fi31 < 5; fi31++ {
                    if (uint8(v30) > 107) {
                        for fi32 := 0; fi32 < 8; fi32++ {
                            var v33 int8 = 0
                            v3 = 1.5
                            _ = fi32
                            _ = v33
                        }
                        var v34 int8 = -128
                        v35 := GenS1{F0: 1, F1: 42, F2: fi31}
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS2{F0: -94, F1: 553, F2: (fi31 + fi31)}
                    fmt.Println(v4, v30)
                    _ = fi31
                    _ = v36
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    if true {
                        v38 := v5[v30]
                        fmt.Println(v30)
                        _ = v38
                    } else {
                        v39 := GenS4{F0: 42, F1: 1, F2: (v30 | (^(v30)))}
                        _ = v39
                    }
                    switch 42 {
                    case 475:
                        v30 += v30
                    default:
                        v30 = 2
                    }
                    _ = fi37
                }
            } else {
                fmt.Println(v3)
            }
            var v40 float32 = v3
            _ = v30
            _ = v40
        } else {
            var v41 float32 = ((((((v3 - ((1e10 * v3) / v3)) * v3) - v3) / 3.14159) * v3) - 3.14159)
            var v42 int8 = 0
            v43 := GenS2{F0: v42, F1: 1, F2: (v2 + v2)}
            _ = v41
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        var v45 int32 = 0
        var v46 float32 = v3
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        v3 = 3.14159
        _ = fi47
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        if false {
            fmt.Println(v4, v2, fi48)
        } else {
            for fi49 := 0; fi49 < 6; fi49++ {
                v50 := GenS1{F0: 0, F1: 9223372036854775807, F2: -67}
                fmt.Println(fi49, fi49)
                v3 = v3
                _ = fi49
                _ = v50
            }
            v4 = v4
            v11 = append(v11, (fi48 + (-(fi48))))
        }
        if false {
            v51 := GenS5{F0: 22, F1: 65535, F2: (true || true)}
            var v52 string = "\n\t"
            _ = v51
            _ = v52
        } else {
            var v53 int = int(v4)
            v54 := GenS2{F0: 100, F1: 255, F2: (v2 & (v53 / (((v2 & v53)) | 1)))}
            if false {
                v53 -= fi48
                switch (v53 + (fi48 + v2)) {
                case 520:
                    var v55 uint = 255
                    _ = v55
                case 9223372036854775807:
                    v56 := GenS6{F0: 4294967295}
                    _ = v56
                default:
                    v4 -= 42
                }
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v2, v53, v3)
                }
            }
            _ = v53
            _ = v54
        }
        _ = fi48
    }
    v4 += genFn0(v2, 1, 42)
    var v57 uint = 18446744073709551615
    v57 = (v57 % ((v57) | 1))
    for fi58 := 0; fi58 < 2; fi58++ {
        switch "" {
        case "":
            v5[2147483648] = 42
        case "\n\t":
            switch ("a" + "foo bar") {
            case "":
                var v59 uint64 = genFn0((v2 - v2), -32768, 693)
                _ = v59
            }
        }
        _ = fi58
    }
    var v60 float64 = float64((v3 + (float32(3.14159) + (v3 + float32(0.0)))))
    var v61 int16 = 32767
    var v62 uint32 = uint32(v4)
    fmt.Println(v57)
    for fi63 := 0; fi63 < 13; fi63++ {
        v62 += (v62 ^ 797)
        v57 = v57
        _ = fi63
    }
    delete(v5, v62)
    v64 := GenS7{F0: 34, F1: (v2 | (v2 & v2))}
    if false {
        if true {
            if true {
                if (v60 > ((float64(v3) / v60) - float64(v3))) {
                    for fi65 := 0; fi65 < 5; fi65++ {
                        var v66 int = v2
                        _ = fi65
                        _ = v66
                    }
                    var v67 int8 = 1
                    _ = v67
                } else {
                    v62 += v62
                }
            }
        } else {
            if false {
                var v68 int32 = -889
                if (0 < 82) {
                    switch "" {
                    case "\n\t":
                        var v69 bool = (v68 < v68)
                        _ = v69
                    case "foo bar":
                        if (46 < 2) {
                            fmt.Println(v2)
                            v3 = ((v3 + (0.0 * v3)) + v3)
                        } else {
                            fmt.Println(v4, v4)
                            var v70 int32 = 2147483647
                            _ = v70
                        }
                    default:
                        fmt.Println(v3)
                    }
                    var v71 int8 = 42
                    for fi72 := 0; fi72 < 20; fi72++ {
                        _ = len(v11)
                        var v73 uint = 44
                        switch fi72 {
                        case 42:
                            v5[(v62 >> (((v73 ^ uint(v4))) & 63))] = (v62 << ((uint(v4)) & 63))
                        }
                        _ = fi72
                        _ = v73
                    }
                    _ = v71
                }
                _ = v68
            } else {
                switch (v2 >> ((328) & 63)) {
                case 128:
                    v2 = (v2 << ((525) & 63))
                case -157:
                    var v74 int64 = 9223372036854775807
                    _ = v74
                case 317:
                    var v75 string = "\n\t"
                    _ = v75
                }
            }
            v76 := GenS4{F0: (v4 & 18446744073709551615), F1: (v57 ^ v57), F2: (v62 & (v62 * 4294967295))}
            var v77 int64 = -9223372036854775808
            _ = v76
            _ = v77
        }
    } else {
        switch "hello" {
        case "a":
            for fi78 := 0; fi78 < 9; fi78++ {
                var v79 uint8 = 136
                _ = fi78
                _ = v79
            }
        default:
            v3 -= (v3 + v3)
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v60
    _ = v61
    _ = v62
    _ = v64
}
