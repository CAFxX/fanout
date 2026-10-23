package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint32) int64 {
    v4 := GenS1{F0: 0, F1: 2147483647, F2: 928}
    p0 = p0
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 += (p0 >> ((716) & 63))
        fmt.Println(fi5, p0)
        _ = fi5
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        v6 := GenS1{F0: 0, F1: 42, F2: 605}
        _ = v6
    case "":
        for fi7 := 0; fi7 < 2; fi7++ {
            p0 = 42
            p0 -= p0
            v8 := []int{fi7, fi7, fi7}
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = 961
            switch "\n\t" {
            case "a":
                p0 += 21
            default:
                switch fi9 {
                case 4294967295:
                    v11 := []int{(fi9 * fi9)}
                    _ = v11
                }
            }
            _ = fi9
            _ = v10
        }
    }
    if false {
        p0 = (^(p0))
        fmt.Println(p0)
        p0 = p0
    }
    _ = p0
    _ = v4
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v12 uint16 = 685
    if true {
        v12 = v12
    } else {
        fmt.Println(v12)
    }
    if false {
        var v13 uint = 4294967295
        _ = v13
    }
    delete(v2, "\n\t")
    var v14 float32 = 3.14159
    for fi15 := 0; fi15 < 7; fi15++ {
        switch fi15 {
        case 2147483647:
            v16 := []int{(fi15 + -1), (fi15 + fi15), fi15, (fi15 * (fi15 >> ((65535) & 63)))}
            _ = v16
        case -9223372036854775808:
            v12 = (v12 - (v12 % ((65535) | 1)))
        case 0:
            var v17 uint16 = (^(v12))
            _ = v17
        }
        _ = fi15
    }
    if true {
        if false {
            switch -2147483648 {
            case 834:
                v14 = 1.5
            }
            if false {
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v14, fi18, v14)
                    v2[("" + "\n\t")] = "\n\t"
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if false {
                            var v20 int16 = 32767
                            _ = v20
                        } else {
                            v14 -= v14
                            var v21 uint64 = 2
                            var v22 uint16 = v12
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{}
                        _ = fi19
                        _ = v23
                    }
                    _ = fi18
                }
                switch 9223372036854775807 {
                case -842:
                    v24 := GenS1{F0: 42, F1: 42, F2: 1}
                    _ = v24
                case -1:
                    switch ("" + ("foo bar" + ("\n\t" + (("\n\t" + "foo bar") + "a")))) {
                    case "foo bar":
                        fmt.Println(v14, v14)
                    }
                case 9223372036854775807:
                    v25 := GenS1{F0: 0, F1: 42, F2: 0}
                    _ = v25
                default:
                    switch 2147483647 {
                    case 4294967295:
                        if (!true) {
                            var v26 int8 = 1
                            switch -439 {
                            case 2147483647:
                                if false {
                                    var v27 uint32 = 42
                                    _ = v27
                                } else {
                                    var v28 uint64 = 2
                                    _ = v28
                                }
                            case -2147483648:
                                for fi29 := 0; fi29 < 19; fi29++ {
                                    var v30 bool = true
                                    v14 = v14
                                    v31 := GenS2{F0: -529, F1: -1}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            var v32 uint = 42
                            _ = v26
                            _ = v32
                        }
                    case 127:
                        if (!(!false)) {
                            switch "a" {
                            case "":
                                switch "" {
                                case "foo bar":
                                    v33 := GenS3{F0: (v12 * v12), F1: 205}
                                    _ = v33
                                case "hello":
                                    v34 := []int{41, -866}
                                    _ = v34
                                default:
                                    v35 := GenS4{F0: 255, F1: (1 == 373)}
                                    _ = v35
                                }
                            case "foo bar":
                                v12 = (v12 >> ((1) & 63))
                            default:
                                v14 = -2.25
                            }
                            if false {
                                fmt.Println(v14, v14)
                                for fi36 := 0; fi36 < 2; fi36++ {
                                    v14 = ((v14 / 0.1) - 1.5)
                                    v14 += 3.14159
                                    var v37 int = -453
                                    _ = fi36
                                    _ = v37
                                }
                            } else {
                                v38 := []int{2147483647, 9223372036854775807}
                                if (("a" + "") <= "") {
                                    if len(v38) > 0 {
                                        v38[0] = -9223372036854775808
                                    }
                                }
                                _ = v38
                            }
                            v39 := GenS5{F0: 724, F1: -94}
                            _ = v39
                        }
                    case 128:
                        v40 := GenS5{F0: -1, F1: -9223372036854775808}
                        _ = v40
                    default:
                        if (!true) {
                            v41 := GenS2{F0: 32767, F1: 0}
                            var v42 int8 = 127
                            fmt.Println(v14, v12)
                            _ = v41
                            _ = v42
                        } else {
                            v43 := GenS5{F0: -578, F1: -1}
                            v44 := []int{}
                            switch -449 {
                            case -161:
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    var v46 bool = true
                                    if false {
                                        var v47 uint = 42
                                        _ = v47
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                            case 9223372036854775807:
                                var v48 uint16 = v12
                                _ = v48
                            case 127:
                                for fi49 := 0; fi49 < 20; fi49++ {
                                    if ((-78 > -127) && (v12 <= 42)) {
                                        if (393 <= 65535) {
                                            fmt.Println(v14, v14)
                                        }
                                    }
                                    _ = fi49
                                }
                            default:
                                switch -9223372036854775808 {
                                case 2147483647:
                                    v14 = (0.1 / v14)
                                case 345:
                                    var v50 int32 = -1
                                    _ = v50
                                default:
                                    if (0.1 <= 0.0) {
                                        var v51 float32 = ((0.1 - v14) + 3.14159)
                                        _ = v51
                                    } else {
                                        var v52 uint = 4294967295
                                        if (false || false) {
                                            v44 = append(v44, 386)
                                            var v53 uint16 = 531
                                            _ = v53
                                        } else {
                                            v54 := GenS3{F0: v12, F1: (v12 << ((v52) & 63))}
                                            v55 := GenS5{F0: 0, F1: -714}
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v52
                                    }
                                }
                            }
                            _ = v43
                            _ = v44
                        }
                    }
                }
            } else {
                var v56 uint = 0
                for fi57 := 0; fi57 < 2; fi57++ {
                    v58 := []int{}
                    v14 = -2.25
                    _ = fi57
                    _ = v58
                }
                _ = v56
            }
            fmt.Println(v14, v12)
        } else {
            v14 = 1e10
        }
    } else {
        if true {
            v12 = (v12 / ((1) | 1))
            v59 := GenS6{F0: 1}
            v12 -= (^(v12))
            _ = v59
        } else {
            fmt.Println(v14)
        }
        delete(v1, 2147483647)
        for fi60 := 0; fi60 < 1; fi60++ {
            var v61 uint64 = 509
            _ = fi60
            _ = v61
        }
    }
    v12 += 32767
    v62 := []int{127, 32767, -9223372036854775808, 9223372036854775807}
    _ = len(v62)
    v63 := GenS7{F0: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v62
    _ = v63
}
