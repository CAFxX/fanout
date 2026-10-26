package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8; F2 string }
type GenS5 struct { F0 int32; F1 int64; F2 int64 }
func genFn0(p0 int64, p1 int8, p2 string) uint8 {
    var v6 int64 = p0
    v7 := []int{1, -432, 558}
    v7 = append(v7, 32767)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 56
}

func main() {
    var v1 uint16 = 379
    var v2 int32 = 2147483647
    var v3 float32 = float32((3.14159 + -2.25))
    v4 := make(map[int]uint32)
    var v5 float32 = v3
    v8 := v4[-969]
    fmt.Println(v2, v8)
    var v9 uint16 = 2
    switch -9223372036854775808 {
    case -928:
        v9 += (^(v9))
    case -2147483648:
        v8 -= v8
    case -961:
        var v10 uint = 1
        _ = v10
    default:
        if true {
            var v11 string = "foo bar"
            _ = v11
        } else {
            v1 += v9
            v12 := v4[977]
            if false {
                v13 := []int{}
                _ = v13
            }
            _ = v12
        }
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS1{F0: 42}
        _ = fi14
        _ = v15
    }
    var v16 uint = 1
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int64 = -908
        fmt.Println(fi17, fi17)
        _ = fi17
        _ = v18
    }
    if (!true) {
        for fi19 := 0; fi19 < 11; fi19++ {
            v9 = v1
            var v20 int8 = 0
            _ = fi19
            _ = v20
        }
    }
    switch "" {
    case "foo bar":
        var v21 float32 = 3.14159
        _ = v21
    case "a":
        v22 := GenS2{F0: 742, F1: true}
        _ = v22
    case "\n\t":
        v1 -= 68
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int = fi23
            v25 := GenS2{F0: 126, F1: false}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 int8 = 0
    var v27 string = ((("\n\t" + ("\n\t" + "foo bar")) + ("hello" + (("a" + "a") + ("hello" + ("" + "a"))))) + "")
    var v28 uint16 = (v9 / ((v1) | 1))
    var v29 int8 = (v26 >> ((1) & 63))
    if (false && false) {
        v27 = v27
    }
    if (4294967295 < 18446744073709551615) {
        if true {
            v3 = (v3 * v3)
            if (42 < 2) {
                v30 := []int{365}
                _ = v30
            }
            var v31 int8 = v29
            _ = v31
        } else {
            switch 127 {
            case -447:
                fmt.Println(v9, v2, v1)
            case -1:
                if ((0.1 / 1e10) < float64((v5 / v3))) {
                    if true {
                        switch 9223372036854775807 {
                        case 2147483647:
                            switch -9223372036854775808 {
                            case 42:
                                v32 := GenS3{F0: (^(v8))}
                                _ = v32
                            case 4294967295:
                                v2 = v2
                            case 127:
                                v5 += ((1.5 * (v3 * ((v3 / v5) - (v3 - (1e10 * 1e10))))) * (((((v3 + (v5 * v5)) * (v3 / v5)) - v3) * float32(0.1)) + 1e10))
                            }
                        default:
                            var v33 uint32 = v8
                            _ = v33
                        }
                        v1 = 32767
                    } else {
                        delete(v4, 949)
                        v5 -= v3
                        v29 = 127
                    }
                    var v34 float32 = float32(1.5)
                    switch -9223372036854775808 {
                    case 137:
                        var v35 uint64 = 789
                        _ = v35
                    case 147:
                        var v36 bool = false
                        _ = v36
                    case 0:
                        switch "\n\t" {
                        case "":
                            if (18446744073709551615 >= (^(v16))) {
                                v37 := GenS2{F0: 673, F1: (!(v27 > ("hello" + ((v27 + v27) + v27))))}
                                _ = v37
                            }
                        case "a":
                            switch 42 {
                            case 1:
                                delete(v4, 32768)
                            case 900:
                                var v38 int64 = 1
                                _ = v38
                            default:
                                v39 := GenS4{F0: (v16 >> (((v16 + v16)) & 63)), F1: 127, F2: "foo bar"}
                                _ = v39
                            }
                        default:
                            if ((!(false && false)) && (false || false)) {
                                fmt.Println(v28, v5, v16)
                                for fi40 := 0; fi40 < 6; fi40++ {
                                    v41 := []int{32768}
                                    for fi42 := 0; fi42 < 19; fi42++ {
                                        var v43 int16 = 32767
                                        _ = fi42
                                        _ = v43
                                    }
                                    _ = fi40
                                    _ = v41
                                }
                            } else {
                                v1 -= uint16(v8)
                            }
                        }
                    default:
                        for fi44 := 0; fi44 < 14; fi44++ {
                            var v45 uint16 = v28
                            if false {
                                v1 = 32767
                                var v46 uint16 = v9
                                _ = v46
                            } else {
                                v47 := GenS2{F0: 32, F1: (float32(1e10) < 3.14159)}
                                v3 = -2.25
                                _ = v47
                            }
                            switch v27 {
                            case "\n\t":
                                v8 = uint32(v16)
                            case "hello":
                                v48 := GenS3{F0: 4294967295}
                                _ = v48
                            default:
                                v29 += (v29 << ((v16) & 63))
                            }
                            _ = fi44
                            _ = v45
                        }
                    }
                    _ = v34
                } else {
                    for fi49 := 0; fi49 < 13; fi49++ {
                        for fi50 := 0; fi50 < 7; fi50++ {
                            v5 = v3
                            var v51 float32 = ((float32(float64(((v3 / float32(3.14159)) + v3))) / v3) - 0.0)
                            fmt.Println(fi50)
                            var v52 float64 = 0.1
                            if false {
                                var v53 uint8 = genFn0(0, v26, v27)
                                _ = v53
                            }
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v54 := []int{}
                        for fi55 := 0; fi55 < 13; fi55++ {
                            var v56 float64 = 0.1
                            v57 := v4[fi55]
                            switch fi55 {
                            case -9223372036854775808:
                                v9 = (^(v9))
                            default:
                                fmt.Println(v8, v8)
                            }
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                        _ = fi49
                        _ = v54
                    }
                    for fi58 := 0; fi58 < 13; fi58++ {
                        var v59 uint16 = (^(v9))
                        _ = fi58
                        _ = v59
                    }
                    var v60 int32 = v2
                    _ = v60
                }
            }
            if (!(!(v26 > (-(v29))))) {
                for fi61 := 0; fi61 < 12; fi61++ {
                    var v62 int64 = -9223372036854775808
                    v8 += 4294967295
                    v63 := GenS5{F0: int32(v62), F1: v62, F2: (v62 - (v62 + -241))}
                    _ = fi61
                    _ = v62
                    _ = v63
                }
            }
            _ = len(v4)
        }
    }
    var v64 int8 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v8
    _ = v9
}
