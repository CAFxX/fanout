package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 uint32; F1 int32; F2 int16 }
type GenS10 struct { F0 bool; F1 int16; F2 int64 }
type GenS11 struct { F0 int16; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 uint32) string {
    p1 -= p1
    if (0 > 9223372036854775807) {
        switch 698 {
        case -81:
            v6 := GenS1{F0: 42, F1: (1000 == 1), F2: true}
            _ = v6
        case -2147483648:
            var v7 float64 = 1e10
            _ = v7
        case 2147483647:
            v8 := []int{32768, 32767, 4294967295, 127}
            _ = v8
        }
    }
    var v9 uint = p0
    var v10 float32 = ((1e10 / 0.1) + -2.25)
    var v11 float64 = 0.1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    var v4 float32 = 1e10
    var v5 uint8 = 175
    v12 := GenS1{F0: 127, F1: (!true), F2: true}
    v4 += 1.5
    fmt.Println(v5, v5)
    _ = len(v3)
    fmt.Println(v2)
    if ((42 <= 32767) && (int16(v5) > 42)) {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{}
            v15 := GenS1{F0: 0, F1: false, F2: true}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 string = ("hello" + "")
    if ((100 != -27) && (((v4 * (((v4 + (v4 * ((((1.5 + (((((v4 * (v4 - v4)) * v4) + (v4 / -2.25)) * (float32(1e10) * 3.14159)) / (v4 - v4))) + (1e10 + v4)) / (v4 - v4)) + v4))) * float32(v1)) + (-2.25 - -2.25))) + v4) >= v4)) {
        v17 := GenS1{F0: -1, F1: false, F2: (!(868 <= 32767))}
        var v18 uint8 = 2
        _ = v17
        _ = v18
    } else {
        fmt.Println(v5, v5)
        var v19 uint16 = 42
        _ = v19
    }
    v16 = v16
    if true {
        if (100 > 100) {
            v1 += (v1 - v1)
            v1 = (0.0 * 1e10)
        } else {
            if false {
                switch 1 {
                case 2147483647:
                    v5 = 0
                case 373:
                    v20 := GenS2{F0: 1, F1: -189, F2: -370}
                    _ = v20
                case 351:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if (2147483647 < 226) {
                            var v22 int64 = (v2 << ((0) & 63))
                            _ = len(v3)
                            switch v16 {
                            case "hello":
                                v23 := GenS1{F0: -128, F1: true, F2: (!false)}
                                _ = v23
                            case "\n\t":
                                v16 = "a"
                            }
                            _ = v22
                        } else {
                            v24 := []int{127, -9223372036854775808}
                            _ = v24
                        }
                        _ = fi21
                    }
                }
                if false {
                    v4 = float32(0.0)
                    v1 += v1
                } else {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        fmt.Println(v1)
                        var v26 uint8 = v5
                        if true {
                            var v27 uint8 = v5
                            v28 := GenS3{F0: (^(fi25)), F1: 636}
                            v29 := GenS2{F0: 1, F1: 42, F2: 1000}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = fi25
                        _ = v26
                    }
                    var v30 float32 = (1e10 / ((v4 + (float32(((float64(v4) * 1.5) - v1)) * float32(v1))) + (-2.25 + v4)))
                    v2 += v2
                    _ = v30
                }
                v31 := GenS1{F0: 127, F1: true, F2: (v2 <= v2)}
                _ = v31
            } else {
                if (!(!false)) {
                    if false {
                        var v32 uint32 = uint32(v5)
                        v33 := GenS4{F0: -467}
                        _ = v32
                        _ = v33
                    }
                    if false {
                        v34 := v3["a"]
                        _ = v34
                    } else {
                        var v35 float64 = 3.14159
                        fmt.Println(v2)
                        v4 = ((v4 / ((1.5 / (v4 + ((((v4 - ((v4 + (((0.0 + v4) - 0.0) / (v4 / (v4 * v4)))) + v4)) + ((v4 - float32((3.14159 / 1e10))) + ((0.1 * 3.14159) / v4))) * 3.14159) - v4))) / v4)) - 1e10)
                        _ = v35
                    }
                }
                var v36 int8 = -128
                _ = v36
            }
            fmt.Println(v2, v4)
            switch v16 {
            case "\n\t":
                v1 = (v1 + v1)
            case "hello":
                v1 = v1
            case "foo bar":
                var v37 int = 1
                _ = v37
            default:
                _ = len(v3)
            }
        }
        fmt.Println(v4, v16, v5)
    } else {
        fmt.Println(v16)
        v2 -= (v2 - 56)
    }
    var v38 int16 = -1
    v39 := GenS5{F0: 1, F1: 0}
    switch -104 {
    case 128:
        for fi40 := 0; fi40 < 1; fi40++ {
            v2 = -1
            if false {
                var v41 int8 = -128
                _ = v41
            }
            v1 -= v1
            _ = fi40
        }
    case 0:
        v42 := GenS6{F0: "\n\t"}
        _ = v42
    default:
        v2 = v2
    }
    var v43 uint64 = 205
    v44 := GenS7{F0: -1, F1: 0}
    if false {
        v45 := GenS8{F0: int64(v38), F1: -503}
        v46 := GenS9{F0: 2147483647, F1: 58, F2: v38}
        switch 214 {
        case 2147483647:
            var v47 bool = (!((255 <= 2) || false))
            _ = v47
        case 32767:
            for fi48 := 0; fi48 < 11; fi48++ {
                switch "" {
                case "hello":
                    v49 := GenS2{F0: 165, F1: -2147483648, F2: (v38 / ((v38) | 1))}
                    _ = v49
                case "\n\t":
                    var v50 int32 = -757
                    _ = v50
                default:
                    var v51 int = fi48
                    _ = v51
                }
                switch "\n\t" {
                case "":
                    v52 := GenS8{F0: 0, F1: -2147483648}
                    _ = v52
                case "\n\t":
                    for fi53 := 0; fi53 < 9; fi53++ {
                        var v54 bool = (0.1 <= v1)
                        _ = fi53
                        _ = v54
                    }
                case "hello":
                    if false {
                        fmt.Println(v38, v38)
                    } else {
                        fmt.Println(v5, fi48)
                    }
                }
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 string = (((("foo bar" + v16) + v16) + v16) + v16)
                    if false {
                        v57 := GenS10{F0: ((fi55 >> ((255) & 63)) <= (fi48 >> ((172) & 63))), F1: (v38 % ((v38) | 1)), F2: (v2 / (((v2 >> ((2) & 63))) | 1))}
                        _ = v57
                    }
                    _ = fi55
                    _ = v56
                }
                _ = fi48
            }
        }
        _ = v45
        _ = v46
    } else {
        var v58 bool = ((^(v2)) > (v2 >> ((2) & 63)))
        v59 := GenS11{F0: 670, F1: v43, F2: (127 >= -76)}
        _ = v58
        _ = v59
    }
    v60 := []int{32767, 4294967295, 32767}
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v60
}
