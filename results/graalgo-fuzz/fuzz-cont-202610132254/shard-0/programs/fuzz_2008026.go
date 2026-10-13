package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint64 }
func genFn0(p0 bool) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 % ((fi6) | 1)), (fi6 & fi6), fi6, fi6}
        var v8 int16 = 0
        v9 := GenS1{F0: -1, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 584
    _ = p0
    _ = v10
    return 1
}
func genFn1() uint16 {
    v11 := GenS1{F0: -2147483648, F1: -654}
    var v12 uint = 370
    v12 = (v12 | v12)
    if false {
        v12 -= 297
    } else {
        v13 := GenS2{F0: 352, F1: (32767 > 1000)}
        _ = v13
    }
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    var v4 string = "foo bar"
    var v5 uint8 = 2
    v1 = (v1 * v1)
    for fi14 := 0; fi14 < 3; fi14++ {
        if (737 <= genFn1()) {
            fmt.Println(v4, v1, v4)
            var v15 int16 = 32767
            var v16 bool = ((((v1 / -2.25) >= 1e10) || false) && (1 <= 2))
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS3{F0: ((-94 != genFn0((708 <= 65535))) || false), F1: -121}
    v18 := GenS2{F0: 163, F1: false}
    v4 = v3
    var v19 float32 = (v1 * v1)
    switch 99 {
    case 32768:
        var v20 uint16 = 2
        _ = v20
    case 9223372036854775807:
        _ = len(v2)
    default:
        v21 := GenS3{F0: false, F1: 32768}
        _ = v21
    }
    switch 4294967295 {
    case 1:
        v1 -= v19
    case 42:
        if (false || (-2147483648 != 42)) {
            if false {
                v22 := GenS4{F0: (!(-61 != 0)), F1: 0}
                _ = v22
            } else {
                fmt.Println(v4)
                v1 = -2.25
            }
            var v23 bool = (306 != 453)
            if (2147483648 > 2) {
                v19 = ((v19 / v1) / (((((((((1.5 + float32(1.5)) + v1) / -2.25) - (((((3.14159 / v1) + -2.25) / v19) / v19) + v1)) / 3.14159) / 1.5) * ((v19 * (-2.25 - v1)) / (v1 * v19))) / 1e10) / (0.0 + (v19 * 0.0))))
            } else {
                if (0.1 == 0.1) {
                    if (4294967295 >= 1) {
                        for fi24 := 0; fi24 < 11; fi24++ {
                            for fi25 := 0; fi25 < 13; fi25++ {
                                if ((!(0 != 364)) && (0.0 <= -2.25)) {
                                    var v26 uint8 = v5
                                    switch v3 {
                                    case "a":
                                        v27 := GenS2{F0: 65535, F1: false}
                                        _ = v27
                                    }
                                    switch fi25 {
                                    case 336:
                                        fmt.Println(v4, v26, v4)
                                    case -385:
                                        v28 := []int{(fi24 - (fi24 ^ 92)), (fi24 - -914), fi25, (fi24 - (fi24 | (fi25 | fi24)))}
                                        _ = v28
                                    case 1:
                                        v4 = v3
                                    default:
                                        if ((!(4294967295 >= 65535)) || v23) {
                                            if true {
                                                var v29 uint = 42
                                                var v30 uint16 = 32767
                                                var v31 int16 = 42
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                    }
                                    _ = v26
                                }
                                var v32 int32 = -698
                                _ = fi25
                                _ = v32
                            }
                            if v23 {
                                v4 = v4
                            }
                            var v33 int = fi24
                            _ = fi24
                            _ = v33
                        }
                        v19 = v19
                    } else {
                        v4 = v3
                    }
                    if true {
                        var v34 string = ((v3 + "foo bar") + v3)
                        v35 := GenS5{F0: 42, F1: -1}
                        _ = v34
                        _ = v35
                    }
                } else {
                    v5 -= (^(v5))
                    _ = len(v2)
                    if v23 {
                        var v36 string = "\n\t"
                        if (v36 <= v4) {
                            var v37 uint64 = 2
                            _ = v37
                        }
                        v23 = v23
                        _ = v36
                    } else {
                        v38 := GenS3{F0: true, F1: 127}
                        fmt.Println(v3, v19)
                        v39 := GenS2{F0: 255, F1: v23}
                        _ = v38
                        _ = v39
                    }
                }
                fmt.Println(v23, v5)
            }
            _ = v23
        }
    case 127:
        v40 := GenS2{F0: 0, F1: ((2147483647 != 1) || false)}
        _ = v40
    }
    v19 = (-2.25 + v19)
    v41 := v2[(v4 + "foo bar")]
    for fi42 := 0; fi42 < 4; fi42++ {
        _ = len(v2)
        _ = fi42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v44 := v2[""]
        var v45 int16 = 0
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 uint64 = 0
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint = 0
            fmt.Println(v5, v47)
            var v50 uint = 0
            var v51 bool = true
            switch fi48 {
            case 128:
                var v52 int64 = -168
                _ = v52
            case 2147483647:
                v53 := GenS6{F0: 9223372036854775807, F1: v51}
                _ = v53
            default:
                var v54 uint32 = 42
                _ = v54
            }
            v5 = v5
            v55 := GenS1{F0: 1, F1: -2147483648}
            if true {
                v3 = (v41 + v3)
                v56 := GenS7{F0: (v47 * v47)}
                switch (("hello" + (v41 + ("" + ("foo bar" + (v4 + (v41 + (v41 + v41))))))) + v3) {
                case "":
                    var v57 float32 = v1
                    _ = v57
                default:
                    v49 -= (v49 / ((65535) | 1))
                    switch v4 {
                    case "hello":
                        var v58 uint32 = 689
                        _ = v58
                    case "\n\t":
                        v59 := []int{(fi46 / ((int(v5)) | 1)), (fi48 * (-(fi48)))}
                        _ = v59
                    default:
                        fmt.Println(v51, v4, v5)
                    }
                }
                _ = v56
            }
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
            _ = v55
        }
        for fi60 := 0; fi60 < 12; fi60++ {
            v5 -= (v5 * v5)
            _ = fi60
        }
        _ = fi46
        _ = v47
    }
    v5 = v5
    var v61 float64 = float64((v1 - (float32(0.0) / float32(1e10))))
    var v62 float64 = (1.5 / (1.5 + v61))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v61
    _ = v62
}
