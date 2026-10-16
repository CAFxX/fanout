package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint; F1 int8 }
func genFn0() uint8 {
    if true {
        var v4 float32 = 0.1
        v4 += ((-2.25 - ((float32((0.1 - float64(v4))) * 0.1) + v4)) + float32(1.5))
        _ = v4
    }
    v5 := []int{-994, -9223372036854775808, 127, 32768}
    _ = v5
    return 46
}
func genFn1() uint8 {
    var v6 int8 = -95
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 -= 74
        v8 := GenS1{F0: ("" + "foo bar"), F1: (fi7 >> ((255) & 63))}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v6 = (^(v6))
        v10 := GenS1{F0: "", F1: fi9}
        _ = fi9
        _ = v10
    }
    var v11 int8 = (v6 / ((v6) | 1))
    v12 := []int{32768, 128, 127, 2147483647}
    _ = v11
    _ = v12
    _ = v6
    return 148
}

func main() {
    v1 := []string{"a", ("foo bar" + "\n\t"), ("foo bar" + ("\n\t" + "hello")), "\n\t"}
    v2 := make(map[uint32]int)
    var v3 uint64 = 269
    var v13 uint8 = genFn1()
    var v14 int8 = 19
    v3 -= v3
    var v15 uint32 = 42
    var v16 uint64 = v3
    var v17 uint64 = (^(v16))
    if true {
        switch "" {
        case "\n\t":
            for fi18 := 0; fi18 < 5; fi18++ {
                for fi19 := 0; fi19 < 12; fi19++ {
                    if (32767 == 255) {
                        fmt.Println(v13)
                        var v20 int8 = (v14 << ((42) & 63))
                        v17 = v17
                        v21 := GenS1{F0: "foo bar", F1: fi19}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint64 = (v17 % (((v17 - (v3 << ((0) & 63)))) | 1))
                        _ = len(v1)
                        _ = v22
                    }
                    var v23 uint8 = 2
                    fmt.Println(v16)
                    _ = fi19
                    _ = v23
                }
                fmt.Println(fi18, v14, v17)
                switch 2147483647 {
                case 693:
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if true {
                            v25 := GenS2{F0: 155, F1: 359, F2: 42}
                            _ = v25
                        }
                        if (float64(0.1) != 0.1) {
                            if (0 != 270) {
                                if true {
                                    var v26 float64 = (-2.25 * float64(0.0))
                                    var v27 bool = true
                                    v28 := GenS2{F0: 0, F1: uint(v15), F2: 410}
                                    var v29 string = ("a" + ("hello" + "\n\t"))
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                var v30 string = ""
                                _ = v30
                            }
                        }
                        _ = fi24
                    }
                case 0:
                    v31 := GenS3{F0: 255, F1: 158, F2: (-(v14))}
                    _ = v31
                }
                _ = fi18
            }
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            fmt.Println(v14, v3)
        default:
            v32 := GenS1{F0: "\n\t", F1: 9223372036854775807}
            _ = v32
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            v15 += v15
            var v34 uint8 = (^(v13))
            _ = fi33
            _ = v34
        }
        var v35 int8 = (v14 + v14)
        _ = v35
    }
    var v36 uint = 315
    switch (("" + "hello") + "\n\t") {
    case "":
        if false {
            if false {
                _ = len(v2)
            }
            if (uint64(v15) < uint64(v15)) {
                fmt.Println(v14, v3, v3)
                for fi37 := 0; fi37 < 4; fi37++ {
                    var v38 int64 = -1
                    var v39 uint = v36
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                v17 += v16
            } else {
                for fi40 := 0; fi40 < 3; fi40++ {
                    v1 = append(v1, "foo bar")
                    for fi41 := 0; fi41 < 20; fi41++ {
                        var v42 int16 = 357
                        var v43 uint64 = v16
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    _ = fi40
                }
            }
            if (!true) {
                if (true || ((v15 - v15) <= v15)) {
                    var v44 float32 = 0.0
                    _ = v44
                }
            }
        }
    case "hello":
        if ((v36 >> ((4294967295) & 63)) <= v36) {
            for fi45 := 0; fi45 < 1; fi45++ {
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 int16 = -1
                    _ = fi46
                    _ = v47
                }
                v1 = append(v1, "a")
                _ = fi45
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                for fi49 := 0; fi49 < 4; fi49++ {
                    v14 += -128
                    fmt.Println(fi48)
                    v50 := GenS2{F0: -180, F1: v36, F2: 0}
                    _ = fi49
                    _ = v50
                }
                _ = fi48
            }
        }
    case "foo bar":
        var v51 uint64 = (v17 ^ (v17 % ((v3) | 1)))
        _ = v51
    default:
        if ((float64(1e10) / (float64(-2.25) + (1e10 + (1e10 / 1e10)))) != float64((0.0 / 1.0))) {
            for fi52 := 0; fi52 < 6; fi52++ {
                v53 := GenS2{F0: 0, F1: v36, F2: -494}
                var v54 int64 = -1
                for fi55 := 0; fi55 < 8; fi55++ {
                    v36 -= v36
                    v54 = (v54 << ((v36) & 63))
                    var v56 string = "foo bar"
                    _ = fi55
                    _ = v56
                }
                _ = fi52
                _ = v53
                _ = v54
            }
            v16 += v17
        } else {
            switch 128 {
            case -2147483648:
                switch "" {
                case "a":
                    v15 = 1
                default:
                    v15 += (^(v15))
                }
            default:
                var v57 uint16 = 642
                _ = v57
            }
            var v58 int8 = v14
            for fi59 := 0; fi59 < 18; fi59++ {
                v15 -= 65535
                _ = fi59
            }
            _ = v58
        }
    }
    v1 = append(v1, "a")
    for fi60 := 0; fi60 < 3; fi60++ {
        if true {
            v61 := GenS3{F0: (^(v36)), F1: -1, F2: v14}
            _ = v61
        } else {
            _ = len(v1)
            fmt.Println(v14, v17, fi60)
            v62 := GenS1{F0: ("a" + "foo bar"), F1: fi60}
            _ = v62
        }
        v15 = (v15 ^ v15)
        _ = fi60
    }
    if (!true) {
        var v63 float64 = 0.1
        for fi64 := 0; fi64 < 7; fi64++ {
            if true {
                v65 := GenS4{F0: int(v36), F1: v15}
                for fi66 := 0; fi66 < 8; fi66++ {
                    if (v15 > (v15 ^ (v15 * v15))) {
                        v67 := v2[(v15 % ((2147483648) | 1))]
                        v14 -= int8(fi64)
                        if (false && true) {
                            var v68 uint = 255
                            _ = v68
                        }
                        _ = v67
                    } else {
                        v14 -= (v14 * v14)
                        switch fi66 {
                        case 128:
                            v69 := GenS5{F0: (v3 % ((v3) | 1)), F1: 249}
                            _ = v69
                        case 221:
                            var v70 float32 = 0.1
                            _ = v70
                        }
                    }
                    _ = fi66
                }
                v71 := v2[v15]
                _ = v65
                _ = v71
            }
            var v72 int64 = 42
            _ = fi64
            _ = v72
        }
        if true {
            v16 = v16
        } else {
            if true {
                if (true && ("\n\t" == (("\n\t" + ("hello" + ((((((("\n\t" + "a") + "a") + "a") + ("a" + "hello")) + ((("hello" + ("a" + ("a" + ("foo bar" + "")))) + "hello") + "")) + "hello") + ("foo bar" + ("hello" + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar")))))) + "hello"))) {
                    _ = len(v2)
                } else {
                    switch 127 {
                    case 9223372036854775807:
                        for fi73 := 0; fi73 < 3; fi73++ {
                            if true {
                                v15 = 194
                            }
                            fmt.Println(v15, v36)
                            _ = fi73
                        }
                    }
                }
            } else {
                _ = len(v1)
                var v74 float64 = (v63 - (0.1 * v63))
                v13 -= 176
                _ = v74
            }
        }
        _ = v63
    }
    switch -2147483648 {
    case 42:
        var v75 string = "\n\t"
        _ = v75
    default:
        if true {
            v36 = v36
        }
    }
    var v76 int = 1
    if ("hello" < "a") {
        v76 += 32768
        var v77 int8 = v14
        _ = v77
    } else {
        if true {
            v78 := GenS6{F0: (v36 | v36), F1: v14}
            _ = v78
        }
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v3 = (v3 ^ v17)
    v36 -= v36
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v76
}
