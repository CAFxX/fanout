package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int8; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 uint16, p2 uint8) uint32 {
    switch -87 {
    case 2147483647:
        for fi4 := 0; fi4 < 15; fi4++ {
            v5 := GenS1{F0: 2}
            switch "" {
            case "a":
                var v6 uint8 = p2
                _ = v6
            case "foo bar":
                p2 = (p0 << ((0) & 63))
            case "hello":
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := GenS1{F0: 2147483648}
                    _ = fi7
                    _ = v8
                }
            default:
                for fi9 := 0; fi9 < 17; fi9++ {
                    if true {
                        p0 -= p0
                        var v10 uint8 = 42
                        _ = v10
                    } else {
                        p2 = p0
                        v11 := []int{fi9, 132, 128}
                        _ = v11
                    }
                    _ = fi9
                }
            }
            fmt.Println(p2, p2)
            _ = fi4
            _ = v5
        }
    case -737:
        if (-751 <= -167) {
            switch 476 {
            case -847:
                p2 = p0
            case 918:
                var v12 int8 = 42
                _ = v12
            }
            var v13 int = 759
            fmt.Println(v13)
            _ = v13
        } else {
            fmt.Println(p1, p1, p0)
            v14 := GenS1{F0: 2}
            p0 -= p0
            _ = v14
        }
    case -2147483648:
        p1 = p1
    default:
        v15 := []int{1}
        _ = v15
    }
    v16 := GenS2{F0: 18446744073709551615, F1: 1000, F2: 764}
    v17 := []int{0, 128, -9223372036854775808, -9223372036854775808}
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := GenS3{F0: 785, F1: "\n\t"}
        _ = fi18
        _ = v19
    }
    var v20 float32 = float32((((3.14159 - (1.5 / float64(0.1))) - -2.25) + float64(1.5)))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v20
    return 958
}
func genFn1(p0 int64, p1 int32, p2 int) uint32 {
    for fi21 := 0; fi21 < 7; fi21++ {
        if (true || (-1 > -1)) {
            for fi22 := 0; fi22 < 11; fi22++ {
                p0 = (p0 - p0)
                _ = fi22
            }
            p2 = (fi21 ^ (p2 % (((-(fi21))) | 1)))
        } else {
            v23 := GenS2{F0: 488, F1: 893, F2: 369}
            switch "" {
            case "hello":
                for fi24 := 0; fi24 < 19; fi24++ {
                    p2 = -256
                    var v25 int32 = -538
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                v26 := GenS1{F0: genFn0(127, 1, 1)}
                _ = v26
            default:
                v27 := []int{1, p2, fi21, (fi21 % ((p2) | 1))}
                _ = v27
            }
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(p0, fi28)
                fmt.Println(p1)
                p2 -= 2147483647
                _ = fi28
            }
            _ = v23
        }
        v29 := []int{}
        _ = fi21
        _ = v29
    }
    v30 := GenS1{F0: 841}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 2147483648
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]bool)
    v3 := []uint32{407, 42, 2147483647}
    var v31 uint8 = v1
    fmt.Println(v31, v1)
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 string = "foo bar"
        var v34 string = (v33 + v33)
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v1 = (^(v31))
        v31 = v1
        var v36 int8 = -69
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        v2["a"] = (-963 == 0)
        _ = fi37
    }
    if ("\n\t" > ("a" + "hello")) {
        fmt.Println(v1, v31)
        v38 := GenS2{F0: 42, F1: 1, F2: int32(v31)}
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                v31 = v1
                _ = len(v2)
                _ = fi40
            }
            _ = fi39
        }
        if false {
            fmt.Println(v31)
            v31 += v31
            if false {
                v31 += 217
            } else {
                var v41 uint32 = 937
                _ = v41
            }
        }
    }
    fmt.Println(v31, v1)
    var v42 int32 = 495
    var v43 float64 = (float64(-2.25) - float64((1.5 / (3.14159 * (1.5 * ((1e10 / 1.0) - 1.5))))))
    v44 := v2[("hello" + "hello")]
    v45 := GenS4{F0: 127, F1: 32767, F2: 0}
    if v44 {
        if true {
            for fi46 := 0; fi46 < 11; fi46++ {
                var v47 int32 = v42
                fmt.Println(fi46, v31)
                if (v44 || v44) {
                    var v48 float64 = (((v43 + 0.1) / (0.0 - (0.1 + (v43 * v43)))) * (1.5 * (v43 - v43)))
                    fmt.Println(v47, v31)
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v44 = ((((v44 && false) && (4294967295 > 723)) || (!((1 == 926) && (875 != 0)))) || ((v44 || (1000 >= -629)) && ((^(v31)) != 0)))
                        v50 := GenS1{F0: 0}
                        _ = fi49
                        _ = v50
                    }
                    fmt.Println(v43, v47)
                    v1 = v1
                }
                _ = fi46
                _ = v47
            }
            var v51 float32 = 1e10
            v52 := GenS2{F0: 255, F1: 1000, F2: v42}
            _ = v51
            _ = v52
        } else {
            var v53 int = 32768
            var v54 float32 = -2.25
            _ = v53
            _ = v54
        }
    }
    v55 := GenS3{F0: 974, F1: (("foo bar" + "hello") + "hello")}
    var v56 float32 = float32(v43)
    if (!v44) {
        for fi57 := 0; fi57 < 16; fi57++ {
            v58 := GenS4{F0: -128, F1: (-(fi57)), F2: v42}
            _ = fi57
            _ = v58
        }
        fmt.Println(v31, v1, v44)
        switch "\n\t" {
        case "":
            if (-9223372036854775808 >= 804) {
                v43 = 0.0
                if (496 < genFn1(42, -257, 32767)) {
                    switch 42 {
                    case 4294967295:
                        if (false || false) {
                            var v59 int8 = 100
                            _ = v59
                        } else {
                            if false {
                                for fi60 := 0; fi60 < 17; fi60++ {
                                    switch 4294967295 {
                                    case 0:
                                        if (!v44) {
                                            v42 = 2147483647
                                            fmt.Println(v31, v56, v43)
                                            v61 := v2[("foo bar" + ("" + "hello"))]
                                            _ = v61
                                        } else {
                                            v43 += ((float64(v56) * 0.0) + 1.5)
                                        }
                                    case 32768:
                                        v56 -= 1.5
                                    case 127:
                                        v62 := GenS4{F0: int8(fi60), F1: fi60, F2: 42}
                                        _ = v62
                                    }
                                    var v63 uint64 = 42
                                    _ = fi60
                                    _ = v63
                                }
                                _ = len(v2)
                                fmt.Println(v1)
                            } else {
                                var v64 int64 = 816
                                _ = v64
                            }
                            v2["\n\t"] = true
                        }
                    }
                    var v65 uint8 = 255
                    _ = v65
                } else {
                    if v44 {
                        fmt.Println(v42)
                    }
                    for fi66 := 0; fi66 < 20; fi66++ {
                        switch "hello" {
                        case "hello":
                            if (-243 == -9223372036854775808) {
                                fmt.Println(v31)
                            }
                        case "\n\t":
                            _ = len(v3)
                        case "":
                            v67 := GenS1{F0: 0}
                            _ = v67
                        default:
                            v44 = ((-2.25 / 3.14159) <= float64(((float32(0.0) + v56) / 1e10)))
                        }
                        _ = fi66
                    }
                }
                for fi68 := 0; fi68 < 17; fi68++ {
                    v56 = v56
                    var v69 int16 = 0
                    fmt.Println(v44, v69, v56)
                    _ = fi68
                    _ = v69
                }
            }
        case "a":
            delete(v2, (("foo bar" + "\n\t") + ("" + "foo bar")))
        }
    }
    v43 += (v43 * v43)
    for fi70 := 0; fi70 < 10; fi70++ {
        v71 := v2[""]
        _ = fi70
        _ = v71
    }
    if v44 {
        for fi72 := 0; fi72 < 18; fi72++ {
            if ((!true) && true) {
                v56 += v56
                var v73 int16 = -1
                _ = v73
            } else {
                v74 := GenS1{F0: 168}
                v3 = append(v3, 148)
                var v75 int32 = v42
                _ = v74
                _ = v75
            }
            v42 = v42
            _ = fi72
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v55
    _ = v56
}
