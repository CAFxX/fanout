package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }
type GenS6 struct { F0 uint16 }
func genFn0() int16 {
    if ("foo bar" == "a") {
        var v4 uint32 = 340
        _ = v4
    } else {
        if false {
            var v5 int16 = 422
            _ = v5
        } else {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 int8 = 84
                v8 := []int{(fi6 - fi6), (-(fi6))}
                v9 := GenS1{F0: 2}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{32767, 9223372036854775807}
            _ = v10
        }
        v11 := []int{}
        _ = v11
    }
    var v12 bool = (65535 > 255)
    v12 = false
    _ = v12
    return 0
}

func main() {
    var v1 uint = 864
    v2 := []string{"\n\t", "\n\t", ""}
    v3 := make(map[int]uint32)
    v1 = (v1 ^ v1)
    var v13 int32 = 791
    v13 = 2147483647
    var v14 int64 = 0
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 float64 = (float64(0.0) * 3.14159)
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: -32768, F1: v1}
    v3[9223372036854775807] = 42
    var v18 uint32 = 42
    if true {
        switch 1 {
        case 128:
            fmt.Println(v1, v18, v13)
        case -641:
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int64 = (v14 & int64(v18))
                v21 := GenS3{F0: (v13 & (^(v13))), F1: v13}
                var v22 int = fi19
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        case 0:
            v23 := GenS3{F0: -2147483648, F1: v13}
            _ = v23
        }
        fmt.Println(v13)
    } else {
        var v24 int8 = -59
        var v25 int16 = 978
        if true {
            if ((v25 ^ -73) > v25) {
                fmt.Println(v13, v18, v13)
                switch (("a" + ("\n\t" + "foo bar")) + ("a" + ("a" + "a"))) {
                case "foo bar":
                    if false {
                        switch "a" {
                        case "hello":
                            v13 = 396
                        default:
                            if true {
                                switch (((("\n\t" + "a") + "hello") + ("hello" + "\n\t")) + "foo bar") {
                                case "hello":
                                    var v26 uint16 = 32767
                                    _ = v26
                                }
                                var v27 string = "foo bar"
                                _ = v27
                            }
                        }
                    }
                case "a":
                    for fi28 := 0; fi28 < 9; fi28++ {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            var v30 int64 = (v14 & (v14 | (v14 >> ((v1) & 63))))
                            v25 -= v25
                            fmt.Println(v18)
                            var v31 int = 42
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        if true {
                            v32 := v3[fi28]
                            switch "a" {
                            case "a":
                                var v33 int32 = (v13 ^ (v13 / (((v13 ^ (v13 | v13))) | 1)))
                                _ = v33
                            case "hello":
                                v34 := GenS3{F0: (-(v13)), F1: (v13 * (^(v13)))}
                                _ = v34
                            case "":
                                var v35 uint32 = v18
                                _ = v35
                            }
                            _ = v32
                        } else {
                            v36 := GenS4{F0: 2}
                            if true {
                                v37 := GenS2{F0: genFn0(), F1: v1}
                                var v38 bool = false
                                _ = v37
                                _ = v38
                            } else {
                                v1 -= v1
                            }
                            v39 := GenS5{F0: 543, F1: 255, F2: ("" == "foo bar")}
                            _ = v36
                            _ = v39
                        }
                        _ = fi28
                    }
                default:
                    fmt.Println(v25)
                }
                v40 := GenS2{F0: (v25 % (((v25 + v25)) | 1)), F1: 18446744073709551615}
                _ = v40
            } else {
                v25 -= v25
                var v41 uint64 = 0
                switch "\n\t" {
                case "\n\t":
                    switch 127 {
                    case 2147483647:
                        var v42 uint32 = 0
                        _ = v42
                    }
                case "foo bar":
                    if true {
                        var v43 float64 = (float64(float32(1e10)) * 1.5)
                        _ = v43
                    } else {
                        var v44 string = ("\n\t" + "foo bar")
                        switch v44 {
                        case "hello":
                            for fi45 := 0; fi45 < 15; fi45++ {
                                v14 = v14
                                fmt.Println(fi45, v41, v24)
                                _ = fi45
                            }
                        case "":
                            v14 += v14
                        case "\n\t":
                            v1 = v1
                        default:
                            v41 += v41
                        }
                        _ = v44
                    }
                case "hello":
                    if (v14 == v14) {
                        v14 -= v14
                        for fi46 := 0; fi46 < 1; fi46++ {
                            if false {
                                v3[32767] = 2
                                if true {
                                    v47 := GenS6{F0: 9}
                                    if (v14 < -272) {
                                        v48 := GenS4{F0: 282}
                                        var v49 uint16 = 0
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 uint8 = 42
                                    _ = v47
                                    _ = v50
                                } else {
                                    v1 -= v1
                                    if (!true) {
                                        var v51 int16 = v25
                                        v1 = (v1 << ((v1) & 63))
                                        for fi52 := 0; fi52 < 7; fi52++ {
                                            v13 += 1
                                            _ = fi52
                                        }
                                        _ = v51
                                    } else {
                                        for fi53 := 0; fi53 < 10; fi53++ {
                                            if false {
                                                v41 -= v41
                                            }
                                            v54 := GenS5{F0: v41, F1: 42, F2: ((255 > 32767) || false)}
                                            if (613 >= 261) {
                                                v41 += (v41 + (v41 >> ((255) & 63)))
                                                v14 = (-(v14))
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                        var v55 int32 = 0
                                        if false {
                                            delete(v3, (fi46 << (((v1 + 18446744073709551615)) & 63)))
                                            v14 = v14
                                            v14 = (v14 + (v14 ^ (v14 | v14)))
                                        }
                                        _ = v55
                                    }
                                }
                            }
                            if false {
                                v56 := GenS1{F0: 128}
                                fmt.Println(v18)
                                _ = v56
                            }
                            v57 := GenS1{F0: 127}
                            _ = fi46
                            _ = v57
                        }
                        var v58 int64 = 0
                        _ = v58
                    } else {
                        _ = len(v3)
                        for fi59 := 0; fi59 < 12; fi59++ {
                            switch ("a" + (("" + "a") + "\n\t")) {
                            case "\n\t":
                                v60 := v3[-447]
                                _ = v60
                            case "hello":
                                v61 := GenS5{F0: 18446744073709551615, F1: 151, F2: (255 > 165)}
                                _ = v61
                            default:
                                v25 += (v25 % ((genFn0()) | 1))
                            }
                            v62 := GenS6{F0: 341}
                            if false {
                                var v63 int32 = (v13 / (((-(v13))) | 1))
                                var v64 int64 = 0
                                _ = v63
                                _ = v64
                            } else {
                                fmt.Println(v41)
                            }
                            _ = fi59
                            _ = v62
                        }
                    }
                }
                _ = v41
            }
            delete(v3, 32767)
            var v65 int8 = v24
            _ = v65
        }
        _ = v24
        _ = v25
    }
    if false {
        v66 := GenS2{F0: genFn0(), F1: (v1 - (v1 << ((v1) & 63)))}
        var v67 int32 = (v13 % ((v13) | 1))
        _ = v66
        _ = v67
    }
    v14 = (-(v14))
    for fi68 := 0; fi68 < 20; fi68++ {
        v69 := GenS4{F0: 521}
        _ = fi68
        _ = v69
    }
    _ = len(v3)
    v70 := GenS1{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v70
}
