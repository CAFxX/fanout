package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 string; F1 uint16; F2 int16 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int64 = -228
    v4 := make(map[uint32]uint32)
    v5 := []int{0, -2147483648, 4294967295}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float32 = 0.0
        var v8 int32 = 319
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        if false {
            v3 -= (v3 >> ((897) & 63))
            var v9 uint = 18446744073709551615
            _ = v9
        }
    } else {
        v3 += v3
        if (932 != 438) {
            v10 := GenS1{F0: -664, F1: 296, F2: -22}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS2{F0: (!(!(v3 > 0))), F1: -280}
                    if len(v5) > 0 {
                        v5[0] = (fi11 ^ fi12)
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v3 += (v3 | v3)
            _ = fi14
        }
    }
    var v15 string = (("\n\t" + ((("a" + ("" + ("" + ("hello" + "foo bar")))) + (("\n\t" + "foo bar") + "a")) + "\n\t")) + "hello")
    v15 = "hello"
    var v16 float32 = (float32(((0.0 * -2.25) / 1e10)) / ((1e10 / ((1e10 + 1e10) / (1.5 - 1e10))) * 1.5))
    var v17 uint8 = 127
    switch v15 {
    case "":
        v15 = v15
    default:
        if false {
            var v18 uint = 1
            switch v15 {
            case "":
                if false {
                    v15 = (v15 + (v15 + "hello"))
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int8 = 42
                        var v21 float32 = v16
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v18 = v18
                }
            default:
                v22 := GenS2{F0: false, F1: 42}
                _ = v22
            }
            _ = v18
        } else {
            if false {
                delete(v1, -9223372036854775808)
                v23 := GenS1{F0: -32768, F1: 65535, F2: -1}
                var v24 uint = 1
                _ = v23
                _ = v24
            } else {
                if (0 > 37) {
                    v25 := GenS3{F0: 9, F1: (-901 < v3)}
                    v26 := GenS4{F0: (v15 + v15), F1: -128}
                    v27 := GenS2{F0: true, F1: 0}
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v17 -= (v17 & v17)
                var v28 uint8 = v17
                _ = v28
            }
            v29 := GenS2{F0: (false || (false && false)), F1: 273}
            if false {
                if len(v5) > 0 {
                    v5[0] = -458
                }
                var v30 uint = 2
                if ((-2.25 > 1e10) && (true || (!(!true)))) {
                    switch ("foo bar" + v15) {
                    case "foo bar":
                        if (uint64(v17) < 4294967295) {
                            v17 -= v17
                            if (((-2147483648 != 127) || (float64(v16) < 0.0)) || true) {
                                v31 := GenS5{F0: v15, F1: 255, F2: 1000}
                                _ = v31
                            } else {
                                v32 := v4[137]
                                v33 := v1[199]
                                v32 += v32
                                _ = v32
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 7; fi34++ {
                                switch fi34 {
                                case -265:
                                    var v35 float32 = v16
                                    _ = v35
                                case 622:
                                    v1[fi34] = "\n\t"
                                case -9223372036854775808:
                                    var v36 int = (fi34 & 2147483647)
                                    _ = v36
                                }
                                _ = len(v2)
                                _ = fi34
                            }
                        } else {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                fmt.Println(v15)
                                var v38 uint64 = 4294967295
                                _ = fi37
                                _ = v38
                            }
                        }
                    case "a":
                        switch -571 {
                        case -744:
                            for fi39 := 0; fi39 < 3; fi39++ {
                                v15 = ((v15 + ("" + (v15 + ("a" + ("foo bar" + v15))))) + (v15 + v15))
                                _ = fi39
                            }
                        case -1:
                            var v40 int16 = -32768
                            _ = v40
                        case 9223372036854775807:
                            switch 32768 {
                            case -621:
                                if len(v5) > 0 {
                                    v5[0] = -40
                                }
                            case -205:
                                var v41 int = 32768
                                _ = v41
                            default:
                                fmt.Println(v30, v3)
                            }
                        }
                    }
                    for fi42 := 0; fi42 < 6; fi42++ {
                        if (-526 > 32767) {
                            v17 -= 2
                            var v43 int64 = (v3 * v3)
                            _ = v43
                        } else {
                            if true {
                                var v44 float64 = (0.1 - 1e10)
                                if (false || (0 <= 1)) {
                                    var v45 uint64 = 42
                                    if false {
                                        v46 := GenS6{F0: 42, F1: 65535}
                                        if (!true) {
                                            v5 = append(v5, 9223372036854775807)
                                            v45 = 18446744073709551615
                                            var v47 bool = (-2.25 < (-2.25 * v44))
                                            _ = v47
                                        } else {
                                            var v48 uint8 = v17
                                            v15 = "\n\t"
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v48
                                        }
                                        _ = v46
                                    }
                                    _ = v45
                                }
                                _ = v44
                            } else {
                                v30 = (^(v30))
                            }
                            for fi49 := 0; fi49 < 5; fi49++ {
                                v5 = append(v5, -161)
                                _ = fi49
                            }
                            v50 := GenS3{F0: -15, F1: (false && (((v16 / v16) / 1e10) == v16))}
                            _ = v50
                        }
                        v17 -= 88
                        if (!(!true)) {
                            switch fi42 {
                            case 128:
                                for fi51 := 0; fi51 < 15; fi51++ {
                                    _ = len(v5)
                                    _ = fi51
                                }
                            case -351:
                                var v52 uint32 = 446
                                _ = v52
                            default:
                                var v53 uint = (v30 / ((v30) | 1))
                                _ = v53
                            }
                            var v54 int16 = 1000
                            _ = v54
                        } else {
                            var v55 bool = true
                            for fi56 := 0; fi56 < 19; fi56++ {
                                var v57 int8 = 1
                                var v58 uint = (v30 % (((v30 ^ (^(v30)))) | 1))
                                if (v55 || ((v15 <= (("\n\t" + v15) + "foo bar")) || (!v55))) {
                                    var v59 int16 = 1000
                                    _ = len(v5)
                                    if true {
                                        v16 = v16
                                        var v60 string = v15
                                        v61 := GenS1{F0: v59, F1: 793, F2: (v57 ^ (v57 & 118))}
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v59
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = (fi56 * 42)
                                    }
                                }
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            _ = v55
                        }
                        _ = fi42
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    var v62 int32 = 0
                    fmt.Println(v15, v17, v17)
                    _ = v62
                }
                _ = v30
            }
            _ = v29
        }
    }
    v63 := GenS7{F0: (true || true), F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
