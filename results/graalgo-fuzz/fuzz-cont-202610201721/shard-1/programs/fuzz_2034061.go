package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 int8) string {
    if true {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 bool = false
            v7 = true
            _ = fi6
            _ = v7
        }
        var v8 int8 = (p0 | p0)
        v9 := []int{-502}
        _ = v8
        _ = v9
    }
    v10 := []int{42}
    p0 += (p0 | p0)
    v11 := GenS1{F0: 12, F1: 92}
    _ = p0
    _ = v10
    _ = v11
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2, 2147483647, 2147483647}
    var v3 float32 = ((float32(((0.0 / 3.14159) - 0.0)) / ((0.1 * 0.1) - ((1.5 * 0.1) + (1e10 * 0.1)))) - 3.14159)
    var v4 int8 = 89
    v5 := make(map[uint32]uint32)
    var v12 int = 127
    if true {
        var v13 float32 = ((float32((0.1 - 1e10)) / v3) - (v3 / v3))
        v3 -= (-2.25 * ((v13 + (float32((float64((v3 * (((0.1 * v3) - v13) / 1e10))) + 3.14159)) + (0.1 + v3))) + (v3 * 0.1)))
        v13 = 1.5
        _ = v13
    } else {
        var v14 float32 = (v3 * (-2.25 + ((v3 * v3) + v3)))
        _ = len(v5)
        _ = v14
    }
    v4 = v4
    var v15 bool = false
    var v16 int = v12
    switch (v12 * (^(v16))) {
    case 539:
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                v15 = false
                var v19 uint = 2
                _ = fi18
                _ = v19
            }
            var v20 int8 = v4
            _ = fi17
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint64 = 42
        v23 := GenS2{F0: 32767, F1: 2}
        _ = fi21
        _ = v22
        _ = v23
    }
    v15 = (v15 && false)
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint64 = 767
        if (false || (!(!v15))) {
            if (464 != -881) {
                if ((1 <= 42) || ((v12 ^ (v12 + (v16 ^ 32767))) == (v12 << ((725) & 63)))) {
                    fmt.Println(v16, v15)
                    fmt.Println(v25)
                } else {
                    v4 -= (v4 << ((65535) & 63))
                    v16 -= v12
                }
                var v26 int16 = -1
                fmt.Println(v4)
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                switch "foo bar" {
                case "foo bar":
                    var v28 int32 = -1
                    _ = v28
                case "\n\t":
                    var v29 string = ""
                    _ = v29
                default:
                    switch 9223372036854775807 {
                    case 127:
                        if v15 {
                            fmt.Println(v15, fi24)
                        } else {
                            fmt.Println(v3, v15, v4)
                            if v15 {
                                if v15 {
                                    var v30 int64 = -478
                                    _ = len(v5)
                                    _ = v30
                                } else {
                                    v31 := GenS2{F0: -622, F1: 1}
                                    if v15 {
                                        v25 -= (v25 ^ (v25 / (((v25 % (((v25 >> ((1) & 63))) | 1))) | 1)))
                                        if (!(!v15)) {
                                            v32 := GenS3{F0: 846, F1: 2}
                                            _ = v32
                                        }
                                        switch 127 {
                                        case 487:
                                            fmt.Println(fi27, fi24, v16)
                                        default:
                                            var v33 int16 = 32767
                                            _ = v33
                                        }
                                    }
                                    var v34 int32 = -139
                                    v1[618] = 2147483647
                                    v35 := GenS2{F0: 0, F1: 1}
                                    var v36 bool = v15
                                    var v37 int32 = -132
                                    _ = v31
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                if ((((0 <= 850) || (v12 != (v12 + (fi24 << ((18446744073709551615) & 63))))) || v15) && (v15 || (v15 && v15))) {
                                    v38 := GenS4{F0: 1, F1: 1}
                                    var v39 string = "foo bar"
                                    v4 -= v4
                                    _ = v38
                                    _ = v39
                                } else {
                                    v40 := v5[298]
                                    _ = v40
                                }
                                v25 = v25
                                if ((v25 * v25) < v25) {
                                    var v41 uint8 = 255
                                    _ = v41
                                }
                            }
                            fmt.Println(v16, fi24, v15)
                        }
                    case -208:
                        fmt.Println(v12, fi27, v25)
                    case -123:
                        v42 := GenS5{F0: false, F1: 127, F2: 30}
                        _ = v42
                    }
                }
                _ = fi27
            }
            fmt.Println(v16, v15)
        }
        v25 = v25
        _ = fi24
        _ = v25
    }
    v12 -= 1
    v3 -= v3
    if (!((v15 && (!false)) && true)) {
        var v43 string = ""
        v44 := GenS4{F0: 0, F1: 684}
        _ = v43
        _ = v44
    }
    v3 = v3
    switch "hello" {
    case "\n\t":
        for fi45 := 0; fi45 < 19; fi45++ {
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 int = v12
                if v15 {
                    v3 = v3
                } else {
                    delete(v1, 42)
                }
                switch 128 {
                case 1:
                    var v48 uint64 = 134
                    _ = v48
                }
                _ = fi46
                _ = v47
            }
            v2 = append(v2, 2147483647)
            fmt.Println(v16)
            _ = fi45
        }
    case "hello":
        fmt.Println(v4)
    case "a":
        var v49 int16 = -560
        _ = v49
    default:
        v16 -= v12
    }
    var v50 int = (v12 | (-(v16)))
    var v51 int = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
