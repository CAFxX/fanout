package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0() uint8 {
    var v7 int64 = 1
    v7 -= 1
    _ = v7
    return 1
}
func genFn1(p0 int) int32 {
    v8 := []int{(p0 << ((2) & 63)), (p0 ^ (p0 & (p0 + (p0 - p0)))), (-(p0))}
    switch (p0 / ((p0) | 1)) {
    case 127:
        p0 += (p0 + p0)
    case 9223372036854775807:
        v9 := GenS1{F0: 1, F1: 876}
        _ = v9
    default:
        p0 = (p0 + (p0 | 533))
    }
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 int = -1
    v2 := make(map[int]int)
    var v3 uint64 = 981
    var v4 float32 = -2.25
    var v5 int = (v1 | (v1 * v1))
    var v6 int8 = 1
    switch (("" + "a") + "") {
    case "a":
        v6 += (v6 * v6)
    case "":
        fmt.Println(v1, v3, v6)
    case "hello":
        if false {
            if (-182 < v1) {
                v6 = (v6 | v6)
                v10 := []int{-831, v5}
                if false {
                    v1 = 1
                } else {
                    _ = len(v10)
                }
                _ = v10
            }
            var v11 float32 = (float32(3.14159) / (float32(((1.5 + float64(float32(float64((1e10 - float32(1e10)))))) + 3.14159)) * v4))
            _ = v11
        } else {
            var v12 uint16 = 65535
            _ = v12
        }
    default:
        v3 = v3
    }
    var v13 float32 = v4
    delete(v2, (v1 + v1))
    var v14 uint8 = 42
    v4 = float32(0.1)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int64 = 42
        v5 += (-(v1))
        _ = fi15
        _ = v16
    }
    v17 := []int{32768}
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v18 int = v5
        _ = v18
    case "a":
        v19 := GenS2{F0: 93}
        _ = v19
    case "":
        if (!false) {
            v20 := GenS2{F0: 65535}
            var v21 uint16 = 0
            if false {
                fmt.Println(v4)
            } else {
                switch 128 {
                case 128:
                    if (((-2.25 / (-2.25 - float64(v4))) == (((((-2.25 * (float64(v4) / float64(-2.25))) * (float64(v13) + float64(0.1))) * 1.5) + (3.14159 + (0.1 / 1.0))) / 1.5)) || (v3 >= v3)) {
                        v13 = 3.14159
                        fmt.Println(v4, v4, v1)
                    }
                case -1:
                    if false {
                        v1 -= (^(v5))
                        var v22 bool = (-2.25 > 3.14159)
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = v22
                    } else {
                        v23 := GenS1{F0: 42, F1: 2147483647}
                        v13 = 1.5
                        if (1.5 < 0.0) {
                            if true {
                                if false {
                                    fmt.Println(v1)
                                    v6 += v6
                                    var v24 bool = (!(!(false || true)))
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 18; fi25++ {
                                        for fi26 := 0; fi26 < 13; fi26++ {
                                            var v27 int8 = int8(v21)
                                            _ = fi26
                                            _ = v27
                                        }
                                        _ = fi25
                                    }
                                    var v28 float32 = (v4 - (float32(1.5) + (v13 + v13)))
                                    _ = v28
                                }
                                v29 := GenS2{F0: 1}
                                _ = v29
                            } else {
                                v4 += ((1.5 - float32(float64((0.0 - v4)))) / 1.5)
                                var v30 uint16 = v21
                                _ = v30
                            }
                            v1 += 9223372036854775807
                        } else {
                            var v31 uint64 = (v3 - v3)
                            fmt.Println(v3)
                            v17 = append(v17, (v5 + (v5 % (((v1 - (v1 >> ((42) & 63)))) | 1))))
                            _ = v31
                        }
                        _ = v23
                    }
                default:
                    fmt.Println(v1, v13)
                }
                switch -879 {
                case -2147483648:
                    v32 := GenS2{F0: 1}
                    _ = v32
                case -1:
                    var v33 int16 = -1
                    _ = v33
                case 42:
                    if false {
                        v34 := GenS2{F0: 433}
                        v35 := v2[(v1 | -1)]
                        _ = v34
                        _ = v35
                    } else {
                        delete(v2, (v5 - v1))
                        v36 := GenS1{F0: 42, F1: v1}
                        _ = v36
                    }
                default:
                    for fi37 := 0; fi37 < 4; fi37++ {
                        var v38 int32 = -2147483648
                        fmt.Println(v4, v1)
                        _ = fi37
                        _ = v38
                    }
                }
                if true {
                    for fi39 := 0; fi39 < 4; fi39++ {
                        v40 := GenS2{F0: 415}
                        fmt.Println(v3, fi39)
                        _ = fi39
                        _ = v40
                    }
                } else {
                    var v41 string = ""
                    for fi42 := 0; fi42 < 15; fi42++ {
                        for fi43 := 0; fi43 < 19; fi43++ {
                            switch v41 {
                            case "":
                                fmt.Println(v6)
                            case "a":
                                v44 := GenS3{F0: 42, F1: (false || (true && false))}
                                v14 += 2
                                _ = v44
                            case "foo bar":
                                var v45 int32 = 2147483647
                                _ = v45
                            default:
                                var v46 int = (fi43 & v1)
                                _ = v46
                            }
                            _ = fi43
                        }
                        var v47 int64 = 883
                        _ = fi42
                        _ = v47
                    }
                    v5 = 4294967295
                    _ = v41
                }
            }
            _ = v20
            _ = v21
        } else {
            v48 := GenS3{F0: -9223372036854775808, F1: false}
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 bool = false
            v2[127] = fi49
            _ = fi50
            _ = v51
        }
        v14 = (v14 / ((v14) | 1))
        _ = fi49
    }
    v52 := GenS1{F0: -32768, F1: (v5 / ((1) | 1))}
    v6 = v6
    switch 1 {
    case 310:
        fmt.Println(v13, v1)
    case 1:
        if true {
            v2[v1] = v5
            v53 := GenS4{F0: v6}
            _ = v53
        } else {
            for fi54 := 0; fi54 < 12; fi54++ {
                if (4294967295 == 65535) {
                    for fi55 := 0; fi55 < 20; fi55++ {
                        v56 := GenS1{F0: -32768, F1: -2147483648}
                        v57 := GenS5{F0: false}
                        var v58 int16 = 863
                        v4 = v4
                        v6 = (v6 / (((-(v6))) | 1))
                        _ = fi55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    if false {
                        switch "foo bar" {
                        case "a":
                            var v59 int16 = 42
                            _ = v59
                        case "\n\t":
                            var v60 uint = 2
                            _ = v60
                        }
                    }
                    var v61 uint8 = (v14 % ((genFn0()) | 1))
                    _ = v61
                } else {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v13)
                    case "a":
                        v62 := GenS4{F0: (v6 ^ (v6 * (^(v6))))}
                        _ = v62
                    case "hello":
                        if false {
                            var v63 uint64 = v3
                            _ = v63
                        }
                    }
                    if (!false) {
                        v64 := GenS6{F0: true, F1: false, F2: uint32(v14)}
                        _ = v64
                    }
                }
                _ = fi54
            }
        }
    case 793:
        v13 = float32(0.1)
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        var v66 uint8 = (v14 & 234)
        v14 = 255
        _ = fi65
        _ = v66
    }
    for fi67 := 0; fi67 < 6; fi67++ {
        switch int(v6) {
        case 1:
            var v68 int8 = v6
            _ = v68
        case -9223372036854775808:
            fmt.Println(fi67)
        }
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
