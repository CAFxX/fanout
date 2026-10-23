package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int8; F2 int64 }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]bool)
    v3 := []string{(("a" + "foo bar") + "\n\t")}
    var v4 uint8 = 13
    var v5 int64 = -9223372036854775808
    var v6 float32 = (-2.25 + float32((3.14159 * 1.5)))
    if (((-(v5)) == (v5 % (((-(v5))) | 1))) && (455 > 1)) {
        v7 := v2["a"]
        var v8 int8 = 100
        _ = v7
        _ = v8
    }
    if (!true) {
        v9 := GenS1{F0: 32767, F1: false}
        switch "" {
        case "a":
            for fi10 := 0; fi10 < 1; fi10++ {
                for fi11 := 0; fi11 < 5; fi11++ {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v5)
                    default:
                        switch "" {
                        case "hello":
                            if (0.1 == 0.0) {
                                v4 = (v4 & 127)
                                switch -1 {
                                case -9223372036854775808:
                                    v1 -= 971
                                case 128:
                                    if ((v1 + uint32(v4)) < (v1 - (^(v1)))) {
                                        v12 := GenS1{F0: 0, F1: (!false)}
                                        switch fi11 {
                                        case 103:
                                            v13 := GenS1{F0: 0, F1: ((0.1 > 0.1) || true)}
                                            _ = v13
                                        case 260:
                                            v14 := GenS2{F0: 1}
                                            _ = v14
                                        }
                                        var v15 int32 = -2147483648
                                        _ = v12
                                        _ = v15
                                    } else {
                                        var v16 uint64 = 0
                                        _ = v16
                                    }
                                case 32767:
                                    v17 := GenS2{F0: 286}
                                    _ = v17
                                default:
                                    v6 += (v6 / v6)
                                }
                            }
                        case "foo bar":
                            v4 = 42
                        }
                    }
                    v5 += v5
                    var v18 int = (fi11 % ((fi10) | 1))
                    _ = fi11
                    _ = v18
                }
                var v19 uint16 = 32767
                _ = fi10
                _ = v19
            }
        case "\n\t":
            var v20 uint64 = 479
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 11; fi21++ {
                switch ("foo bar" + "foo bar") {
                case "":
                    if true {
                        switch 32768 {
                        case -564:
                            v4 += v4
                        }
                        var v22 int64 = 42
                        v5 += v5
                        _ = v22
                    }
                case "\n\t":
                    var v23 uint32 = (v1 << ((2) & 63))
                    _ = v23
                case "hello":
                    v24 := GenS1{F0: 42, F1: true}
                    _ = v24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    v6 = v6
                    if false {
                        if false {
                            v6 = 1e10
                            if false {
                                if (((v1 << ((93) & 63)) < (^(v1))) && true) {
                                    fmt.Println(v6, v5, fi21)
                                    var v26 int32 = -287
                                    _ = v26
                                } else {
                                    var v27 uint64 = 1
                                    _ = v27
                                }
                            }
                            v6 = ((((v6 + v6) / 3.14159) + v6) * v6)
                        }
                        v28 := GenS2{F0: 179}
                        _ = v28
                    }
                    fmt.Println(v1, fi21)
                    _ = fi25
                }
                var v29 uint = 255
                _ = fi21
                _ = v29
            }
        default:
            var v30 int32 = 375
            _ = v30
        }
        fmt.Println(v5, v1)
        _ = v9
    } else {
        var v31 uint16 = 806
        v32 := GenS3{F0: (v4 * 1), F1: 1, F2: 815}
        v5 = (v5 * (v5 / ((v5) | 1)))
        _ = v31
        _ = v32
    }
    v33 := GenS3{F0: (v4 + v4), F1: 1, F2: 18446744073709551615}
    if true {
        var v34 uint = 18446744073709551615
        v35 := GenS1{F0: 102, F1: false}
        switch -1 {
        case 0:
            var v36 uint8 = v4
            _ = v36
        case 1:
            if true {
                if (42 <= 766) {
                    if (42 <= 0) {
                        var v37 int32 = 0
                        _ = v37
                    }
                    v34 = v34
                    v6 += v6
                } else {
                    var v38 uint = (v34 >> ((v34) & 63))
                    var v39 uint64 = 18446744073709551615
                    _ = len(v2)
                    _ = v38
                    _ = v39
                }
                v40 := GenS3{F0: v4, F1: 540, F2: 840}
                var v41 int = 9223372036854775807
                _ = v40
                _ = v41
            } else {
                switch 127 {
                case 32768:
                    switch "hello" {
                    case "a":
                        delete(v2, "hello")
                    case "\n\t":
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v3 = append(v3, "")
                            delete(v2, ("foo bar" + "a"))
                            _ = fi42
                        }
                    }
                default:
                    v43 := GenS3{F0: (v4 | (v4 / ((v4) | 1))), F1: (^(v34)), F2: 844}
                    _ = v43
                }
                fmt.Println(v1, v6)
            }
        case 703:
            v5 -= v5
        default:
            for fi44 := 0; fi44 < 3; fi44++ {
                if true {
                    v45 := GenS2{F0: 0}
                    for fi46 := 0; fi46 < 4; fi46++ {
                        var v47 uint = 511
                        v1 = (^(v1))
                        if true {
                            v4 -= 218
                            v6 = v6
                            if true {
                                v48 := GenS4{F0: 469}
                                var v49 uint64 = 4294967295
                                var v50 bool = (-773 != int16(v1))
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                            _ = len(v3)
                        } else {
                            _ = len(v2)
                            v3 = append(v3, ("" + "hello"))
                            var v51 int32 = -230
                            _ = v51
                        }
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                } else {
                    v1 -= v1
                }
                var v52 float64 = 1e10
                for fi53 := 0; fi53 < 18; fi53++ {
                    v2["a"] = (false || false)
                    _ = fi53
                }
                _ = fi44
                _ = v52
            }
        }
        _ = v34
        _ = v35
    }
    v1 = (^(v1))
    v1 = (v1 & (v1 / ((v1) | 1)))
    var v54 int16 = -551
    if (v1 < (v1 * v1)) {
        fmt.Println(v4, v54, v1)
        fmt.Println(v5, v1)
    } else {
        var v55 uint32 = v1
        var v56 int16 = v54
        _ = v55
        _ = v56
    }
    v57 := GenS5{F0: -32768, F1: 1, F2: (^(v5))}
    var v58 int16 = v54
    var v59 int16 = v58
    v1 += v1
    v5 += v5
    if (true || (2 <= 210)) {
        v6 = (0.1 / ((v6 - ((float32(3.14159) * v6) + 0.0)) + ((1e10 * (v6 * 0.0)) + v6)))
        var v60 uint = 42
        _ = v60
    } else {
        v4 += v4
        for fi61 := 0; fi61 < 9; fi61++ {
            var v62 bool = false
            v4 = (v4 >> ((4294967295) & 63))
            fmt.Println(v4)
            _ = fi61
            _ = v62
        }
    }
    var v63 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v63
}
