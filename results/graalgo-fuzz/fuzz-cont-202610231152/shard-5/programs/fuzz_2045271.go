package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint64; F2 int8 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 bool }
func genFn0(p0 int32, p1 int32) uint64 {
    v6 := GenS1{F0: false}
    fmt.Println(p0)
    var v7 bool = true
    v7 = (127 >= 115)
    if true {
        if v7 {
            var v8 uint8 = 90
            var v9 int = 9223372036854775807
            switch "\n\t" {
            case "foo bar":
                var v10 uint64 = 1
                _ = v10
            case "hello":
                var v11 int64 = 9223372036854775807
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            var v12 int8 = -108
            p0 = (p1 % (((-(p0))) | 1))
            _ = v12
        }
    } else {
        var v13 int16 = 32767
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 449
}

func main() {
    v1 := []string{"foo bar", "foo bar", ("a" + ""), "foo bar"}
    var v2 bool = true
    var v3 float64 = (float64(0.1) / (1e10 * (float64(-2.25) - ((0.1 - 3.14159) * float64((0.1 * float32(1e10)))))))
    v4 := make(map[string]uint32)
    var v5 uint32 = 2147483647
    var v14 float32 = -2.25
    var v15 uint16 = 2
    for fi16 := 0; fi16 < 3; fi16++ {
        if (69 >= 0) {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
        v17 := GenS1{F0: (337 == 462)}
        for fi18 := 0; fi18 < 13; fi18++ {
            if v2 {
                switch (-(fi16)) {
                case 0:
                    var v19 int = 1
                    _ = v19
                case -2147483648:
                    v5 += v5
                case -1:
                    fmt.Println(v5)
                }
                var v20 int64 = -9223372036854775808
                v20 = (v20 & v20)
                switch "a" {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    if v2 {
                        if v2 {
                            _ = len(v4)
                            fmt.Println(v5, v2)
                            switch "\n\t" {
                            case "":
                                if v2 {
                                    if (42 > 83) {
                                        v3 += ((v3 - float64(((0.1 + (v14 + (1.5 * ((float32(v3) / 1e10) / ((v14 - -2.25) * (1e10 + 1.5)))))) * v14))) * (v3 * v3))
                                        if ((("\n\t" + "") + "foo bar") < "") {
                                            if v2 {
                                                var v21 float32 = (0.0 * 1e10)
                                                fmt.Println(v20)
                                                if (128 > 0) {
                                                    var v22 uint8 = 40
                                                    switch "" {
                                                    case "hello":
                                                        var v23 uint32 = v5
                                                        _ = v23
                                                    case "":
                                                        _ = len(v4)
                                                    case "\n\t":
                                                        delete(v4, "\n\t")
                                                    }
                                                    fmt.Println(v21)
                                                    _ = v22
                                                }
                                                _ = v21
                                            } else {
                                                var v24 int16 = -402
                                                _ = v24
                                            }
                                        } else {
                                            var v25 float64 = ((((v3 / float64((v14 + v14))) + 0.0) * v3) + 0.0)
                                            fmt.Println(fi16, v3, v2)
                                            if len(v1) > 0 {
                                                v1[0] = ("\n\t" + ("" + "foo bar"))
                                            }
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = "foo bar"
                                                }
                                            }
                                            _ = v25
                                        }
                                        v26 := GenS1{F0: (-37 != 121)}
                                        _ = v26
                                    } else {
                                        var v27 int16 = 32767
                                        _ = v27
                                    }
                                    fmt.Println(v20, v2, v3)
                                    var v28 uint64 = 0
                                    _ = v28
                                }
                            case "hello":
                                v5 += (v5 * v5)
                            case "\n\t":
                                v29 := GenS2{F0: 255}
                                fmt.Println(v14, v20, v2)
                                var v30 uint64 = genFn0(820, -1)
                                _ = v29
                                _ = v30
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                        v31 := GenS2{F0: 255}
                        if (!(v2 || (!v2))) {
                            if v2 {
                                var v32 string = "a"
                                var v33 float64 = v3
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = len(v4)
                        _ = v31
                    }
                default:
                    v34 := GenS3{F0: "a"}
                    _ = v34
                }
                _ = len(v4)
                switch "a" {
                case "foo bar":
                    if ((v5 + v5) >= 42) {
                        if true {
                            v15 += 32767
                            var v35 float64 = v3
                            v1 = append(v1, ((("" + "") + "foo bar") + ("foo bar" + (("" + ("a" + "foo bar")) + "hello"))))
                            _ = v35
                        } else {
                            var v36 uint8 = 232
                            _ = v36
                        }
                    }
                default:
                    var v37 float32 = v14
                    _ = v37
                }
                fmt.Println(fi18)
                delete(v4, ("a" + ("\n\t" + (("hello" + "foo bar") + "foo bar"))))
                _ = v20
            } else {
                v2 = ((4294967295 > 18446744073709551615) && true)
            }
            var v38 uint = 18446744073709551615
            _ = fi18
            _ = v38
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v14)
    var v39 float64 = (v3 * v3)
    if (v2 && (v15 >= v15)) {
        for fi40 := 0; fi40 < 10; fi40++ {
            for fi41 := 0; fi41 < 12; fi41++ {
                if (!(0 <= 100)) {
                    var v42 uint16 = v15
                    fmt.Println(v5, v39)
                    v43 := GenS3{F0: "a"}
                    _ = v42
                    _ = v43
                } else {
                    v44 := GenS4{F0: 18446744073709551615, F1: 0, F2: uint32(v15)}
                    v45 := GenS3{F0: "foo bar"}
                    fmt.Println(v39, v14)
                    _ = v44
                    _ = v45
                }
                _ = fi41
            }
            var v46 uint64 = 853
            _ = fi40
            _ = v46
        }
    }
    _ = len(v1)
    if false {
        switch ("a" + ("" + ("foo bar" + "\n\t"))) {
        case "foo bar":
            v47 := GenS5{F0: "hello", F1: 1}
            _ = v47
        case "\n\t":
            v39 -= (-2.25 * v3)
        default:
            for fi48 := 0; fi48 < 15; fi48++ {
                v5 = (v5 * 2147483648)
                if v2 {
                    if true {
                        v3 -= v3
                        if v2 {
                            v5 += v5
                            for fi49 := 0; fi49 < 15; fi49++ {
                                v2 = v2
                                fmt.Println(fi49)
                                fmt.Println(v3, v3, v39)
                                var v50 int8 = 1
                                if (32767 != -32768) {
                                    var v51 float32 = v14
                                    if (!v2) {
                                        if (!(-577 <= 9223372036854775807)) {
                                            fmt.Println(v15)
                                            v52 := v4[("\n\t" + "\n\t")]
                                            v3 = 0.1
                                            _ = v52
                                        } else {
                                            fmt.Println(v39, v51)
                                            v15 += v15
                                            var v53 int = fi49
                                            _ = v53
                                        }
                                    } else {
                                        switch ("" + "a") {
                                        case "\n\t":
                                            _ = len(v4)
                                        case "":
                                            var v54 int = (fi48 * fi48)
                                            _ = v54
                                        default:
                                            fmt.Println(v15, v2, v50)
                                        }
                                        _ = len(v1)
                                        v2 = (true || v2)
                                        if (4294967295 < 42) {
                                            switch -912 {
                                            case 1:
                                                var v55 uint8 = 7
                                                _ = v55
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                            var v56 int32 = 2147483647
                                            _ = v56
                                        } else {
                                            fmt.Println(fi48)
                                        }
                                        fmt.Println(v14, v50)
                                    }
                                    fmt.Println(v50)
                                    switch ("a" + (("hello" + ("\n\t" + "\n\t")) + (("foo bar" + ("foo bar" + "foo bar")) + ("hello" + "foo bar")))) {
                                    case "":
                                        var v57 uint64 = 0
                                        _ = v57
                                    }
                                    _ = v51
                                } else {
                                    var v58 int32 = -1
                                    if ((true && false) || (false && (v2 && (v2 && (-759 >= -9223372036854775808))))) {
                                        fmt.Println(v5)
                                    } else {
                                        v39 -= ((3.14159 - 0.0) - ((v39 - v3) - v39))
                                        delete(v4, "a")
                                        v3 = v3
                                    }
                                    _ = v58
                                }
                                _ = fi49
                                _ = v50
                            }
                            v59 := GenS6{F0: 0, F1: genFn0(851, -869), F2: 53}
                            _ = v59
                        }
                    }
                    switch "" {
                    case "hello":
                        fmt.Println(v2, v14, v2)
                    case "":
                        fmt.Println(v39)
                    default:
                        for fi60 := 0; fi60 < 18; fi60++ {
                            v14 -= v14
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v61 float32 = float32(float64(v14))
                            _ = fi60
                            _ = v61
                        }
                    }
                } else {
                    var v62 int64 = 9223372036854775807
                    if true {
                        v39 = 1e10
                        v3 = 0.1
                    }
                    switch 2147483647 {
                    case 128:
                        var v63 bool = (true || v2)
                        _ = v63
                    case 454:
                        v3 += 1.5
                    default:
                        v1 = append(v1, "foo bar")
                    }
                    _ = v62
                }
                _ = fi48
            }
        }
    } else {
        v64 := GenS7{F0: 0, F1: 2}
        var v65 uint16 = v15
        _ = v64
        _ = v65
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        fmt.Println(v15)
        v67 := GenS2{F0: 1}
        _ = fi66
        _ = v67
    }
    _ = len(v1)
    if (v2 || v2) {
        for fi68 := 0; fi68 < 2; fi68++ {
            switch ("hello" + "foo bar") {
            case "":
                if (!(v2 || v2)) {
                    switch (("a" + "hello") + "hello") {
                    case "\n\t":
                        fmt.Println(v2, v3)
                    case "a":
                        fmt.Println(v14)
                    case "foo bar":
                        for fi69 := 0; fi69 < 3; fi69++ {
                            v3 -= (float64(v14) * float64(float32(v3)))
                            v14 -= v14
                            fmt.Println(v3, fi68)
                            _ = fi69
                        }
                    default:
                        var v70 float32 = -2.25
                        _ = v70
                    }
                    _ = len(v4)
                    v14 += v14
                } else {
                    if v2 {
                        var v71 uint32 = (v5 % ((v5) | 1))
                        fmt.Println(v14)
                        for fi72 := 0; fi72 < 9; fi72++ {
                            if true {
                                var v73 string = "\n\t"
                                _ = v73
                            } else {
                                v74 := GenS5{F0: ("" + "\n\t"), F1: 2}
                                _ = v74
                            }
                            _ = fi72
                        }
                        _ = v71
                    } else {
                        _ = len(v1)
                    }
                    v39 += (v39 + (float64(v14) / v3))
                }
            case "\n\t":
                v1 = append(v1, "")
            }
            v4["\n\t"] = 2
            v75 := v4[(("" + "") + "foo bar")]
            _ = fi68
            _ = v75
        }
    }
    var v76 int32 = -442
    v5 -= v5
    for fi77 := 0; fi77 < 9; fi77++ {
        v4["\n\t"] = (v5 >> ((0) & 63))
        switch "a" {
        case "\n\t":
            fmt.Println(v14)
        default:
            v76 = -200
        }
        _ = fi77
    }
    var v78 uint = 4294967295
    switch "a" {
    case "hello":
        for fi79 := 0; fi79 < 10; fi79++ {
            for fi80 := 0; fi80 < 1; fi80++ {
                v3 -= (((v39 / v39) * v3) + 1e10)
                fmt.Println(v39, fi80)
                v4["a"] = (^(v5))
                _ = fi80
            }
            fmt.Println(v76)
            _ = fi79
        }
    case "\n\t":
        var v81 string = "hello"
        _ = v81
    case "":
        var v82 int = 0
        _ = v82
    default:
        v83 := GenS8{F0: v2}
        _ = v83
    }
    v76 = (v76 | v76)
    v14 -= (v14 * v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v76
    _ = v78
}
