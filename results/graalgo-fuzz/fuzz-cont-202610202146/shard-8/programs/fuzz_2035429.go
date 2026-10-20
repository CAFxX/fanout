package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint; F2 int8 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0(p0 string) uint8 {
    fmt.Println(p0)
    v7 := []int{32767, -9223372036854775808, 2147483647, 127}
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    v8 := GenS1{F0: 502, F1: 255, F2: -447}
    switch 19 {
    case -3:
        var v9 int = 9223372036854775807
        _ = v9
    default:
        var v10 uint = 42
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 189
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295, 717}
    var v3 int64 = 42
    var v4 uint8 = 0
    v5 := make(map[uint32]uint32)
    var v6 int8 = 72
    fmt.Println(v6, v4, v6)
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            if ("\n\t" == "foo bar") {
                fmt.Println(fi12, fi12)
            }
            switch "a" {
            case "foo bar":
                v13 := GenS1{F0: v3, F1: genFn0((("a" + ("\n\t" + "foo bar")) + "a")), F2: v3}
                _ = v13
            case "hello":
                v6 = v6
            case "a":
                var v14 int32 = -2147483648
                _ = v14
            default:
                if (!(952 > 2)) {
                    switch fi11 {
                    case 0:
                        switch fi11 {
                        case -9223372036854775808:
                            var v15 uint8 = (v4 % ((176) | 1))
                            _ = v15
                        }
                    case 32768:
                        if (!true) {
                            if (!(((!(32767 > 418)) && true) || false)) {
                                fmt.Println(v3, fi11)
                            }
                            v4 += (v4 * v4)
                            if true {
                                if false {
                                    if (fi12 < -9223372036854775808) {
                                        var v16 uint64 = 2
                                        v3 -= -9223372036854775808
                                        fmt.Println(v4, v16)
                                        _ = v16
                                    }
                                } else {
                                    var v17 int64 = v3
                                    switch fi11 {
                                    case 4294967295:
                                        switch -862 {
                                        case 2147483647:
                                            fmt.Println(fi12, v6)
                                        }
                                    case -76:
                                        v3 += 42
                                    }
                                    var v18 bool = false
                                    _ = v17
                                    _ = v18
                                }
                            }
                        } else {
                            delete(v5, 4294967295)
                            if false {
                                var v19 int = fi11
                                fmt.Println(fi11, fi11)
                                v4 = 0
                                _ = v19
                            }
                            v6 += v6
                            v20 := v5[12]
                            v21 := GenS2{F0: false}
                            _ = v20
                            _ = v21
                        }
                    case 4294967295:
                        switch "hello" {
                        case "":
                            v3 += v3
                        case "a":
                            v6 = -128
                        default:
                            v22 := GenS1{F0: 1, F1: v4, F2: (v3 << ((2) & 63))}
                            _ = v22
                        }
                    }
                    v3 += v3
                    var v23 float32 = float32((float64(0.1) * (3.14159 * (0.1 * float64(-2.25)))))
                    _ = v23
                } else {
                    var v24 float32 = 0.0
                    v24 -= (((1.5 + float32(1.5)) - v24) * v24)
                    var v25 uint32 = 2147483648
                    v26 := GenS3{F0: "", F1: v4}
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi12
        }
        v27 := GenS4{F0: 985}
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi11
        _ = v27
    }
    var v28 uint64 = 1
    var v29 float32 = 3.14159
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 float64 = -2.25
            if false {
                if true {
                    fmt.Println(v6, v6)
                    if false {
                        v29 = 0.1
                    }
                    switch ("hello" + ("\n\t" + "a")) {
                    case "\n\t":
                        v3 -= v3
                    case "a":
                        v29 = v29
                    case "":
                        v33 := GenS2{F0: ((167 < -662) && (104 == v6))}
                        _ = v33
                    default:
                        if (!(4294967295 >= 4294967295)) {
                            v6 += (^(v6))
                            v32 = ((float64(3.14159) / v32) - 3.14159)
                            switch ("hello" + ("\n\t" + (("\n\t" + "") + "a"))) {
                            case "a":
                                v34 := GenS5{F0: v6}
                                _ = v34
                            case "\n\t":
                                fmt.Println(fi30, v29)
                            default:
                                v4 -= v4
                            }
                            v29 = (v29 / float32(v32))
                            if false {
                                var v35 uint16 = 255
                                switch "hello" {
                                case "a":
                                    v36 := GenS4{F0: 508}
                                    _ = v36
                                case "foo bar":
                                    v4 = (v4 << ((0) & 63))
                                default:
                                    v37 := GenS6{F0: true, F1: 263, F2: -113}
                                    _ = v37
                                }
                                _ = v35
                            }
                        } else {
                            v38 := GenS3{F0: "", F1: (v4 << ((1) & 63))}
                            var v39 float32 = float32((1e10 / v32))
                            _ = v38
                            _ = v39
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v40 := GenS2{F0: false}
                    if false {
                        v32 = 0.1
                        v4 = (v4 % (((^(v4))) | 1))
                    }
                    _ = v40
                }
                var v41 uint32 = 1
                if (255 >= 327) {
                    if (false || true) {
                        var v42 int32 = 1
                        v43 := GenS3{F0: "\n\t", F1: (v4 - (^(v4)))}
                        _ = v42
                        _ = v43
                    } else {
                        v28 -= (v28 << ((255) & 63))
                        var v44 int32 = -830
                        _ = v44
                    }
                }
                _ = v41
            } else {
                v45 := v1[32767]
                _ = v45
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    switch 0 {
    case -2147483648:
        var v46 uint8 = (v4 | (v4 | (v4 + 255)))
        _ = v46
    }
    var v47 int = 32768
    v5[uint32(v28)] = 2147483647
    if true {
        v48 := GenS4{F0: 489}
        _ = v48
    } else {
        for fi49 := 0; fi49 < 4; fi49++ {
            v50 := GenS1{F0: 42, F1: genFn0("\n\t"), F2: (^(v3))}
            if (-2147483648 < 512) {
                for fi51 := 0; fi51 < 15; fi51++ {
                    v52 := GenS5{F0: 3}
                    if true {
                        if true {
                            switch ("" + "") {
                            case "":
                                v2 = append(v2, 162)
                            }
                            var v53 int = fi49
                            _ = v53
                        } else {
                            if false {
                                var v54 uint8 = (v4 - genFn0(("a" + "a")))
                                var v55 int16 = 42
                                switch fi51 {
                                case 128:
                                    var v56 int64 = v3
                                    _ = v56
                                case 133:
                                    v47 -= 1
                                case 379:
                                    var v57 float32 = -2.25
                                    _ = v57
                                default:
                                    var v58 bool = (!(uint32(v28) < 587))
                                    _ = v58
                                }
                                _ = v54
                                _ = v55
                            } else {
                                v28 += (v28 * (v28 % ((4294967295) | 1)))
                                v6 = (v6 << ((370) & 63))
                            }
                            switch fi51 {
                            case -9223372036854775808:
                                v28 = (v28 / (((v28 | 4294967295)) | 1))
                            case 4294967295:
                                v59 := GenS7{F0: 32767, F1: v6}
                                _ = v59
                            }
                        }
                    } else {
                        var v60 int = 0
                        v61 := GenS8{F0: (v3 << ((406) & 63)), F1: 4294967295, F2: -2147483648}
                        _ = v60
                        _ = v61
                    }
                    _ = fi51
                    _ = v52
                }
            }
            _ = fi49
            _ = v50
        }
    }
    v28 -= v28
    switch (v47 + v47) {
    case -1:
        if (-2147483648 > 404) {
            v62 := GenS5{F0: -1}
            _ = v62
        } else {
            v47 += (v47 - (v47 / (((v47 / ((v47) | 1))) | 1)))
            v63 := GenS7{F0: 1, F1: (^(v6))}
            var v64 float32 = v29
            _ = v63
            _ = v64
        }
    }
    var v65 int32 = 1
    v29 = v29
    var v66 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
